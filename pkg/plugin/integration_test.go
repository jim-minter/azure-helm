package plugin

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/ghodss/yaml"
	securityapi "github.com/openshift/api/security/v1"
	fakesec "github.com/openshift/client-go/security/clientset/versioned/fake"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes/fake"
	restclient "k8s.io/client-go/rest"
	k8stesting "k8s.io/client-go/testing"

	"github.com/openshift/openshift-azure/pkg/api"
	pluginapi "github.com/openshift/openshift-azure/pkg/api/plugin"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/cluster"
	"github.com/openshift/openshift-azure/pkg/cluster/kubeclient"
	"github.com/openshift/openshift-azure/pkg/cluster/scaler"
	"github.com/openshift/openshift-azure/pkg/cluster/updateblob"
	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/startup"
	"github.com/openshift/openshift-azure/pkg/sync"
	fakecloud "github.com/openshift/openshift-azure/pkg/util/azureclient/fake"
	"github.com/openshift/openshift-azure/pkg/util/random"
	"github.com/openshift/openshift-azure/pkg/util/resourceid"
	"github.com/openshift/openshift-azure/pkg/util/tls"
	"github.com/openshift/openshift-azure/pkg/util/wait"
	testtls "github.com/openshift/openshift-azure/test/util/tls"
)

const (
	vaultKeyNamePublicHostname = "PublicHostname"
	vaultKeyNameRouter         = "Router"
)

func getFakeDeployer(log *logrus.Entry, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) api.DeployFn {
	return func(ctx context.Context, azuretemplate map[string]interface{}) error {
		log.Info("applying arm template deployment")

		err := az.DeploymentsClient.CreateOrUpdateAndWait(ctx, cs.Properties.AzProfile.ResourceGroup, "azuredeploy", resources.Deployment{
			Properties: &resources.DeploymentProperties{
				Template: azuretemplate,
				Mode:     resources.Incremental,
			},
		})
		if err != nil {
			log.Warnf("deployment failed: %#v", err)
			return err
		}

		return nil
	}
}

func enrich(cs *api.OpenShiftManagedCluster) error {
	rg := "testRG"
	dnsDomain := "cloudapp.azure.com"
	tenantID := uuid.NewV4().String()
	clientID := uuid.NewV4().String()
	secret := "suspicious"
	cs.Properties.AzProfile = api.AzProfile{
		TenantID:       tenantID,
		SubscriptionID: uuid.NewV4().String(),
		ResourceGroup:  rg,
	}

	cs.Properties.AuthProfile.IdentityProviders = make([]api.IdentityProvider, 1)
	cs.Properties.AuthProfile.IdentityProviders[0].Name = "Azure AD"
	cs.Properties.AuthProfile.IdentityProviders[0].Provider = &api.AADIdentityProvider{
		Kind:     "AADIdentityProvider",
		ClientID: clientID,
		Secret:   secret,
		TenantID: tenantID,
	}

	cs.Properties.MasterServicePrincipalProfile = api.ServicePrincipalProfile{
		ClientID: uuid.NewV4().String(),
		Secret:   uuid.NewV4().String(),
	}
	cs.Properties.WorkerServicePrincipalProfile = api.ServicePrincipalProfile{
		ClientID: uuid.NewV4().String(),
		Secret:   uuid.NewV4().String(),
	}

	// /subscriptions/{subscription}/resourcegroups/{resource_group}/providers/Microsoft.ContainerService/openshiftmanagedClusters/{cluster_name}
	cs.ID = resourceid.ResourceID(cs.Properties.AzProfile.SubscriptionID, rg, "Microsoft.ContainerService/openshiftmanagedClusters", cs.Name)

	if len(cs.Properties.RouterProfiles) == 0 {
		cs.Properties.RouterProfiles = []api.RouterProfile{
			{
				Name: "default",
			},
		}
	}

	var vaultURL string
	var err error
	vaultURL, err = random.VaultURL("kv-")
	if err != nil {
		return err
	}

	cs.Properties.APICertProfile.KeyVaultSecretURL = vaultURL + "/secrets/" + vaultKeyNamePublicHostname
	cs.Properties.RouterProfiles[0].RouterCertProfile.KeyVaultSecretURL = vaultURL + "/secrets/" + vaultKeyNameRouter

	cs.Properties.PublicHostname = "openshift." + rg + "." + dnsDomain
	cs.Properties.RouterProfiles[0].PublicSubdomain = "apps." + rg + "." + dnsDomain

	if cs.Properties.FQDN == "" {
		cs.Properties.FQDN, err = random.FQDN(cs.Location+"."+dnsDomain, 20)
		if err != nil {
			return err
		}
	}

	if cs.Properties.RouterProfiles[0].FQDN == "" {
		cs.Properties.RouterProfiles[0].FQDN, err = random.FQDN(cs.Location+"."+dnsDomain, 20)
		if err != nil {
			return err
		}
	}

	return nil
}

func (w fakeResponseWrapper) DoRaw() ([]byte, error) {
	return w, nil
}

func (w fakeResponseWrapper) Stream() (io.ReadCloser, error) {
	return nil, nil
}

func newFakeResponseWrapper(raw []byte) restclient.ResponseWrapper {
	var fr fakeResponseWrapper = raw
	return fr
}

type fakeResponseWrapper []byte

func getFakeHTTPClient(cs *api.OpenShiftManagedCluster) wait.SimpleHTTPClient {
	return wait.NewFakeHTTPClient()
}

func newFakeUpgrader(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, testConfig api.TestConfig, kubeclient kubeclient.Interface, azs *fakecloud.AzureCloud) (cluster.Upgrader, error) {
	arm, err := arm.New(ctx, log, cs, testConfig)
	if err != nil {
		return nil, err
	}

	u := &cluster.Upgrade{
		Interface: kubeclient,

		TestConfig:     testConfig,
		AccountsClient: azs.AccountsClient,
		StorageClient:  azs.StorageClient,
		Vmc:            azs.VirtualMachineScaleSetVMsClient,
		Ssc:            azs.VirtualMachineScaleSetsClient,
		Kvc:            azs.KeyVaultClient,
		Log:            log,
		ScalerFactory:  scaler.NewFactory(),
		Hasher: &cluster.Hash{
			Log:            log,
			TestConfig:     testConfig,
			StartupFactory: startup.New,
			Arm:            arm,
		},
		Arm:                arm,
		GetConsoleClient:   getFakeHTTPClient,
		GetAPIServerClient: getFakeHTTPClient,
		Cs:                 cs,
	}

	u.Cs.Config.ConfigStorageAccountKey = "config"
	u.Cs.Config.ConfigStorageAccountKey = uuid.NewV4().String()
	bsc := u.StorageClient.GetBlobService()
	u.UpdateBlobService = updateblob.NewBlobService(bsc)

	return u, nil
}

func newFakeKubeclient(log *logrus.Entry, cli *fake.Clientset, seccli *fakesec.Clientset) kubeclient.Interface {
	return &kubeclient.Kubeclientset{
		Log:    log,
		Client: cli,
		Seccli: seccli,
	}
}

func setupNewCluster(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) (*plugin, *fake.Clientset, error) {
	data, err := ioutil.ReadFile("../../pluginconfig/pluginconfig-311.yaml")
	if err != nil {
		return nil, nil, err
	}
	var template *pluginapi.Config
	if err := yaml.Unmarshal(data, &template); err != nil {
		return nil, nil, err
	}

	template.Certificates.GenevaLogging.Cert = testtls.DummyCertificate
	template.Certificates.GenevaLogging.Key = testtls.DummyPrivateKey
	template.Certificates.GenevaMetrics.Cert = testtls.DummyCertificate
	template.Certificates.GenevaMetrics.Key = testtls.DummyPrivateKey
	template.GenevaImagePullSecret = []byte("pullSecret")
	template.ImagePullSecret = []byte("imagePullSecret")

	cli := fake.NewSimpleClientset()
	cli.PrependReactor("get", "deployments", func(action k8stesting.Action) (bool, runtime.Object, error) {
		get := action.(k8stesting.GetAction)
		d := &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      get.GetName(),
				Namespace: get.GetNamespace(),
			},
			Status: appsv1.DeploymentStatus{
				AvailableReplicas: 1,
				UpdatedReplicas:   1,
			},
		}
		return true, d, nil
	})
	cli.PrependReactor("get", "pods", func(action k8stesting.Action) (bool, runtime.Object, error) {
		get := action.(k8stesting.GetAction)
		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      get.GetName(),
				Namespace: get.GetNamespace(),
			},
			Status: corev1.PodStatus{
				Conditions: []corev1.PodCondition{
					{
						Type:   corev1.PodReady,
						Status: corev1.ConditionTrue,
					},
				},
			},
		}
		return true, pod, nil
	})
	cli.PrependReactor("get", "nodes", func(action k8stesting.Action) (bool, runtime.Object, error) {
		get := action.(k8stesting.GetAction)
		node := &corev1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: get.GetName(),
			},
			Status: corev1.NodeStatus{
				Conditions: []corev1.NodeCondition{
					{
						Type:   corev1.NodeReady,
						Status: corev1.ConditionTrue,
					},
				},
			},
		}
		return true, node, nil
	})

	cli.AddProxyReactor("services", func(action k8stesting.Action) (handled bool, ret restclient.ResponseWrapper, err error) {
		return true, newFakeResponseWrapper(nil), nil
	})

	priv := securityapi.SecurityContextConstraints{
		TypeMeta: metav1.TypeMeta{
			Kind: "securitycontextconstraints",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "privileged",
		},
		AllowPrivilegedContainer: true,
		Users:                    []string{"system:admin"},
	}
	seccli := fakesec.NewSimpleClientset()
	// securitycontextconstraints gets mistakenly converted to securitycontextconstraints"es" by the generic
	// plural converter and it's not in the list of exceptions as it's an openshift type. So
	// we are using a reactor to return value.
	seccli.PrependReactor("get", "*", func(action k8stesting.Action) (bool, runtime.Object, error) {
		get := action.(k8stesting.GetAction)
		if get.GetResource().Resource == "securitycontextconstraints" {
			return true, &priv, nil
		}
		return false, nil, fmt.Errorf("does not exist")
	})
	seccli.PrependReactor("update", "*", func(action k8stesting.Action) (bool, runtime.Object, error) {
		get := action.(k8stesting.UpdateAction)
		if get.GetResource().Resource == "securitycontextconstraints" {
			return true, &priv, nil
		}
		return false, nil, fmt.Errorf("does not exist")
	})
	kc := newFakeKubeclient(log, cli, seccli)
	p := &plugin{
		pluginConfig: template,
		testConfig:   api.TestConfig{RunningUnderTest: true},
		upgraderFactory: func(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, initializeStorageClients, disableKeepAlives bool, testConfig api.TestConfig) (cluster.Upgrader, error) {
			return newFakeUpgrader(ctx, log, cs, testConfig, kc, az)
		},
		configInterfaceFactory: config.New,
		log:                    log,
		now:                    time.Now,
	}
	err = enrich(cs)
	if err != nil {
		return nil, nil, err
	}

	az.ComputeRP.Cs = cs
	err = p.GenerateConfig(ctx, cs, false)
	if err != nil {
		return nil, nil, err
	}

	if err := p.CreateOrUpdate(ctx, cs, false, getFakeDeployer(log, cs, az)); err != nil {
		return nil, nil, err
	}
	return p, cli, nil
}

func newTestCs() *api.OpenShiftManagedCluster {
	return &api.OpenShiftManagedCluster{
		Name:     "integrationTest",
		Location: "eastus",
		Config: api.Config{
			ConfigStorageAccount:      "config",
			RegistryStorageAccount:    "registry",
			RegistryStorageAccountKey: "foo",

			SSHKey: testtls.DummyPrivateKey,
			Certificates: api.CertificateConfig{
				Ca:            api.CertKeyPair{Cert: testtls.DummyCertificate, Key: testtls.DummyPrivateKey},
				NodeBootstrap: api.CertKeyPair{Cert: testtls.DummyCertificate, Key: testtls.DummyPrivateKey},
			},
		},
		Properties: api.Properties{
			OpenShiftVersion: "v3.11",
			AgentPoolProfiles: []api.AgentPoolProfile{
				{Role: api.AgentPoolProfileRoleMaster, Count: 3, Name: "master", VMSize: "Standard_D2s_v3", SubnetCIDR: "10.0.0.0/24", OSType: "Linux"},
				{Role: api.AgentPoolProfileRoleCompute, Count: 1, Name: "compute", VMSize: "Standard_D2s_v3", SubnetCIDR: "10.0.0.0/24", OSType: "Linux"},
				{Role: api.AgentPoolProfileRoleInfra, Count: 2, Name: "infra", VMSize: "Standard_D2s_v3", SubnetCIDR: "10.0.0.0/24", OSType: "Linux"},
			},
			NetworkProfile: api.NetworkProfile{VnetCIDR: "10.0.0.0/8"},
		},
	}
}

func newFakeAzureCloud(log *logrus.Entry) *fakecloud.AzureCloud {
	bKey, _ := tls.PrivateKeyAsBytes(testtls.DummyPrivateKey)
	bCert, _ := tls.CertAsBytes(testtls.DummyCertificate)
	secret := "PRIVATE KEY\n" + string(bKey) + "CERTIFICATE\n" + string(bCert)
	secrets := []keyvault.SecretBundle{
		{
			ID:    to.StringPtr("PublicHostname"),
			Value: to.StringPtr(secret),
		},
		{
			ID:    to.StringPtr("Router"),
			Value: to.StringPtr(secret),
		},
	}

	return fakecloud.NewFakeAzureCloud(log, secrets)
}

func getHashes(az *fakecloud.AzureCloud, cs *api.OpenShiftManagedCluster) (*updateblob.UpdateBlob, string, error) {
	bsc := az.StorageClient.GetBlobService()
	updateBlobService := updateblob.NewBlobService(bsc)
	blob, err := updateBlobService.Read()
	if err != nil {
		return nil, "", err
	}
	// get sync deployment checksum annotation
	// FIXME: only doing it this way as fake kube Get on the deployment
	// always returns an empty string on the annotation
	syncer, err := sync.New(az.ComputeRP.Log, cs, false)
	if err != nil {
		return nil, "", err
	}
	syncChecksum, err := syncer.Hash()
	if err != nil {
		return nil, "", err
	}
	return blob, hex.EncodeToString(syncChecksum), nil
}

type rotationType string

const (
	rotationCompute rotationType = "compute"
	rotationInfra   rotationType = "infra"
	rotationMaster  rotationType = "master"
	rotationSync    rotationType = "sync"
)

func getRotations(beforeBlob, afterBlob *updateblob.UpdateBlob, beforeSyncChecksum, afterSyncChecksum string) map[rotationType]bool {
	rotations := map[rotationType]bool{rotationCompute: false, rotationInfra: false, rotationMaster: false, rotationSync: false}
	for host := range beforeBlob.HostnameHashes {
		rotated := !reflect.DeepEqual(beforeBlob.HostnameHashes[host], afterBlob.HostnameHashes[host])
		if rotated {
			rotations[rotationMaster] = true
		}
	}
	for scaleset := range beforeBlob.ScalesetHashes {
		rotated := !reflect.DeepEqual(beforeBlob.ScalesetHashes[scaleset], afterBlob.ScalesetHashes[scaleset])
		if strings.Contains(scaleset, "compute") && rotated {
			rotations[rotationCompute] = true
		} else if strings.Contains(scaleset, "infra") && rotated {
			rotations[rotationInfra] = true
		}
	}
	if beforeSyncChecksum != afterSyncChecksum {
		rotations[rotationSync] = true
	}
	return rotations
}

func getNodeCountFromAz(az *fakecloud.AzureCloud) map[rotationType]int {
	nodeCount := map[rotationType]int{rotationCompute: 0, rotationInfra: 0, rotationMaster: 0}
	for _, scaleset := range az.ComputeRP.State {
		for _, role := range []rotationType{rotationCompute, rotationInfra, rotationMaster} {
			if strings.Contains(scaleset.Name, string(role)) {
				nodeCount[role] += len(scaleset.Vms)
				break
			}
		}
	}

	return nodeCount
}

func TestHowAdminConfigChangesCausesRotations(t *testing.T) {
	tests := []struct {
		name           string
		change         func(cs *api.OpenShiftManagedCluster)
		expectRotation map[rotationType]bool
	}{
		{
			name:           "no changes",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			change:         func(cs *api.OpenShiftManagedCluster) {},
		},
		{
			name:           "change vm image",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: true, rotationSync: true, rotationCompute: true},
			change:         func(cs *api.OpenShiftManagedCluster) { cs.Config.ImageVersion = "311.12.12345678" },
		},
		{
			name:           "change controller loglevel",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: false, rotationSync: false, rotationCompute: false},
			change:         func(cs *api.OpenShiftManagedCluster) { cs.Config.ComponentLogLevel.ControllerManager = to.IntPtr(5) },
		},
		{
			name:           "change container image",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: true, rotationCompute: false},
			change:         func(cs *api.OpenShiftManagedCluster) { cs.Config.Images.WebConsole = "newImage" },
		},
	}

	log := logrus.NewEntry(logrus.StandardLogger())
	ctx := context.Background()
	cs := newTestCs()
	az := newFakeAzureCloud(log)
	defer az.Cleanup()
	p, _, err := setupNewCluster(ctx, log, cs, az)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Infof("--- Test: %s", tt.name)
			beforeBlob, beforeSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}
			oldCs := cs.DeepCopy()
			tt.change(cs)

			errs := p.ValidateAdmin(ctx, cs, oldCs)
			if errs != nil {
				t.Fatal(errs)
			}
			perr := p.CreateOrUpdate(ctx, cs, true, getFakeDeployer(log, cs, az))
			if perr != nil {
				t.Fatal(perr)
			}

			afterBlob, afterSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}
			rotations := getRotations(beforeBlob, afterBlob, beforeSyncChecksum, afterSyncChecksum)
			if !reflect.DeepEqual(tt.expectRotation, rotations) {
				t.Fatalf("rotation mismatch: expected %v, got %v", tt.expectRotation, rotations)
			}
		})
	}
}

func TestHowUserConfigChangesCausesRotations(t *testing.T) {
	tests := []struct {
		name           string
		change         func(cs *api.OpenShiftManagedCluster)
		expectRotation map[rotationType]bool
	}{
		{
			name:           "no changes",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			change:         func(cs *api.OpenShiftManagedCluster) {},
		},
		{
			// Note: master and infra must be changed together.
			name:           "change master and infra vm size",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: true, rotationSync: false, rotationCompute: false},
			change: func(cs *api.OpenShiftManagedCluster) {
				for i := range cs.Properties.AgentPoolProfiles {
					if cs.Properties.AgentPoolProfiles[i].Role != api.AgentPoolProfileRoleCompute {
						cs.Properties.AgentPoolProfiles[i].VMSize = "Standard_D16s_v3"
					}
				}
			},
		},
		{
			// Note: master is rotating here, is this expected?
			name:           "change compute vm size",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: false, rotationSync: false, rotationCompute: true},
			change: func(cs *api.OpenShiftManagedCluster) {
				for i := range cs.Properties.AgentPoolProfiles {
					if cs.Properties.AgentPoolProfiles[i].Role == api.AgentPoolProfileRoleCompute {
						cs.Properties.AgentPoolProfiles[i].VMSize = "Standard_F16s_v2"
					}
				}
			},
		},
		{
			name:           "change AADIdentityProvider",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: false, rotationSync: true, rotationCompute: false},
			change: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.AuthProfile.IdentityProviders[0].Provider.(*api.AADIdentityProvider).Secret = "new"
			},
		},
	}

	log := logrus.NewEntry(logrus.StandardLogger())
	ctx := context.Background()
	cs := newTestCs()
	az := newFakeAzureCloud(log)
	p, _, err := setupNewCluster(ctx, log, cs, az)
	if err != nil {
		t.Fatal(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeBlob, beforeSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}
			oldCs := cs.DeepCopy()
			tt.change(cs)

			errs := p.Validate(ctx, cs, oldCs, false)
			if errs != nil {
				t.Fatal(errs)
			}
			perr := p.CreateOrUpdate(ctx, cs, true, getFakeDeployer(log, cs, az))
			if perr != nil {
				t.Fatal(perr)
			}

			afterBlob, afterSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}
			rotations := getRotations(beforeBlob, afterBlob, beforeSyncChecksum, afterSyncChecksum)
			if !reflect.DeepEqual(tt.expectRotation, rotations) {
				t.Fatalf("rotation mismatch: expected %v, got %v", tt.expectRotation, rotations)
			}
		})
	}
}

func TestHowActionsCauseRotations(t *testing.T) {
	log := logrus.NewEntry(logrus.StandardLogger())
	ctx := context.Background()
	tests := []struct {
		name           string
		change         func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error
		expectRotation map[rotationType]bool
		expectNodes    map[rotationType]int
		expectCalls    []string
	}{
		{
			name:           "scale up",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 6, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				oldCs := cs.DeepCopy()
				for i, p := range cs.Properties.AgentPoolProfiles {
					if p.Role == api.AgentPoolProfileRoleCompute {
						cs.Properties.AgentPoolProfiles[i].Count = 6
					}
				}
				errs := p.Validate(ctx, cs, oldCs, false)
				if errs != nil {
					return kerrors.NewAggregate(errs)
				}
				perr := p.CreateOrUpdate(ctx, cs, true, getFakeDeployer(log, cs, az))
				if perr != nil {
					return perr
				}
				return nil
			},
		},
		{
			name:           "rotate cluster secrets",
			expectRotation: map[rotationType]bool{rotationMaster: true, rotationInfra: true, rotationSync: true, rotationCompute: true},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				perr := p.RotateClusterSecrets(ctx, cs, getFakeDeployer(log, cs, az))
				if perr != nil {
					return perr
				}
				return nil
			},
		},
		{
			name:           "GetPluginVersion() cause no rotations",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				p.GetPluginVersion(ctx)
				return nil
			},
		},
		{
			name:           "ListClusterVMs() cause no rotations",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				_, perr := p.ListClusterVMs(ctx, cs)
				if perr != nil {
					return perr
				}
				return nil
			},
		},
		{
			name:           "ListEtcdBackups() cause no rotations",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				_, perr := p.ListEtcdBackups(ctx, cs)
				if perr != nil {
					return perr
				}
				return nil
			},
		},
		{
			name:           "GetControlPlanePods() cause no rotations",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				_, perr := p.GetControlPlanePods(ctx, cs)
				if perr != nil {
					return perr
				}
				return nil
			},
		},
		{
			name:           "runcommand - no rotations and correct call",
			expectRotation: map[rotationType]bool{rotationMaster: false, rotationInfra: false, rotationSync: false, rotationCompute: false},
			expectNodes:    map[rotationType]int{rotationCompute: 1, rotationInfra: 2, rotationMaster: 3},
			expectCalls:    []string{"VirtualMachineScaleSetVMsClient:RunCommand:ss-master:1"},
			change: func(p *plugin, cs *api.OpenShiftManagedCluster, az *fakecloud.AzureCloud) error {
				perr := p.RunCommand(ctx, cs, "master-000001", "RestartDocker")
				if perr != nil {
					return perr
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// for this test we always start with a new cluster (unlike the config change test above)
			cs := newTestCs()
			az := newFakeAzureCloud(log)
			defer az.Cleanup()
			p, _, err := setupNewCluster(ctx, log, cs, az)

			beforeBlob, beforeSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}

			// clear the calls
			az.ComputeRP.Calls = []string{}

			err = tt.change(p, cs, az)
			if err != nil {
				t.Fatal(err)
			}

			for _, ec := range tt.expectCalls {
				found := false
				for _, call := range az.ComputeRP.Calls {
					if call == ec {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("call %s not found in %v", ec, az.ComputeRP.Calls)
				}
			}
			nodeCount := getNodeCountFromAz(az)
			if !reflect.DeepEqual(tt.expectNodes, nodeCount) {
				t.Fatalf("node mismatch: expected %v, got %v", tt.expectNodes, nodeCount)
			}

			afterBlob, afterSyncChecksum, err := getHashes(az, cs)
			if err != nil {
				t.Fatal(err)
			}
			rotations := getRotations(beforeBlob, afterBlob, beforeSyncChecksum, afterSyncChecksum)
			if !reflect.DeepEqual(tt.expectRotation, rotations) {
				t.Fatalf("rotation mismatch: expected %v, got %v", tt.expectRotation, rotations)
			}
		})
	}
}

func fileHash(path string) (string, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(path)
	hasher.Write(s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func hashCurrentFiles(az *fakecloud.AzureCloud, log *logrus.Entry) func() map[string]map[string][]string {
	type bothHashes struct {
		before string
		after  string
	}
	hashes := map[string]bothHashes{}
	prefix := ""
	parentDir := ""
	beforeCallback := func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		hash, err := fileHash(path)
		if err != nil {
			return err
		}
		hashes[prefix+strings.TrimPrefix(path, parentDir)] = bothHashes{before: hash}
		return nil
	}
	afterCallback := func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		hash, err := fileHash(path)
		if err != nil {
			return err
		}
		var bh bothHashes
		bh, ok := hashes[prefix+strings.TrimPrefix(path, parentDir)]
		if !ok {
			bh = bothHashes{}
		}
		bh.after = hash
		hashes[prefix+strings.TrimPrefix(path, parentDir)] = bh
		return nil
	}
	for _, scaleset := range az.ComputeRP.State {
		for _, vm := range scaleset.Vms {
			prefix = fmt.Sprintf("%s:%s:", scaleset.Name, *vm.InstanceID)
			parentDir = scaleset.VmsDir[*vm.InstanceID]
			filepath.Walk(scaleset.VmsDir[*vm.InstanceID], beforeCallback)
		}
	}

	return func() map[string]map[string][]string {
		for _, scaleset := range az.ComputeRP.State {
			for _, vm := range scaleset.Vms {
				prefix = fmt.Sprintf("%s:%s:", scaleset.Name, *vm.InstanceID)
				parentDir = scaleset.VmsDir[*vm.InstanceID]
				filepath.Walk(scaleset.VmsDir[*vm.InstanceID], afterCallback)
			}
		}
		diff := map[string]map[string][]string{}
		for key, val := range hashes {
			if val.before != val.after {
				keySplit := strings.Split(key, ":")
				_, ok := diff[keySplit[0]]
				if !ok {
					diff[keySplit[0]] = map[string][]string{}
				}
				_, ok = diff[keySplit[0]][keySplit[1]]
				if !ok {
					diff[keySplit[0]][keySplit[1]] = []string{}
				}
				diff[keySplit[0]][keySplit[1]] = append(diff[keySplit[0]][keySplit[1]], keySplit[2])
			}
		}
		return diff
	}
}

func TestHowRotateClusterSecretsChangesFiles(t *testing.T) {
	log := logrus.NewEntry(logrus.StandardLogger())
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})
	ctx := context.Background()

	cs := newTestCs()
	az := newFakeAzureCloud(log)
	defer az.Cleanup()
	p, _, err := setupNewCluster(ctx, log, cs, az)
	if err != nil {
		t.Fatal(err)
	}

	whatFilesChanged := hashCurrentFiles(az, log)

	perr := p.RotateClusterSecrets(ctx, cs, getFakeDeployer(log, cs, az))
	if perr != nil {
		t.Fatal(perr)
	}

	// For now we are ignoring the instance
	expected := map[string][]string{
		"ss-master": {
			"/etc/origin/master/master.proxy-client.key",
			"/etc/origin/master/session-secrets.yaml",
			"/etc/origin/node/node.kubeconfig",
			"/etc/origin/master/master.server.key",
			"/etc/origin/master/openshift-master.kubeconfig",
			"/etc/origin/master/master.proxy-client.crt",
			"/root/.kube/config",
			"/etc/origin/master/admin.key",
			"/etc/origin/master/admin.kubeconfig",
			"/etc/origin/master/master.kubelet-client.crt",
			"/etc/origin/master/admin.crt",
			"/etc/origin/node/sdn.kubeconfig",
			"/etc/origin/master/master.kubelet-client.key",
			"/etc/origin/master/master.server.crt",
		},
		"ss-infra": {
			"/etc/origin/node/pods/sdn.yaml",
			"/etc/origin/node/ca.crt",
			"/etc/origin/node/node-bootstrapper.key",
			"/etc/sysconfig/atomic-openshift-node",
			"/var/lib/origin/.docker/config.json",
			"/etc/origin/node/bootstrap.kubeconfig",
			"/etc/origin/node/node-bootstrapper.crt",
			"/etc/origin/node/resolv.conf",
			"/etc/pki/ca-trust/source/anchors/openshift-ca.crt",
			"/etc/origin/node/pods/ovs.yaml",
			"/etc/origin/node/node-config.yaml",
			"/etc/origin/cloudprovider/azure.conf",
			"/etc/origin/node/sdn.kubeconfig",
		},
		"ss-compute": {
			"/etc/origin/node/sdn.kubeconfig",
			"/etc/origin/cloudprovider/azure.conf",
			"/etc/origin/node/pods/ovs.yaml",
			"/etc/origin/node/node-bootstrapper.crt",
			"/etc/origin/node/node-config.yaml",
			"/etc/origin/node/resolv.conf",
			"/etc/origin/node/bootstrap.kubeconfig",
			"/etc/sysconfig/atomic-openshift-node",
			"/etc/origin/node/node-bootstrapper.key",
			"/var/lib/origin/.docker/config.json",
			"/etc/origin/node/ca.crt",
			"/etc/origin/node/pods/sdn.yaml",
			"/etc/pki/ca-trust/source/anchors/openshift-ca.crt",
		},
	}
	actual := whatFilesChanged()
	for ssFullName := range actual {
		found := false
		for ss := range expected {
			if strings.HasPrefix(ssFullName, ss) {
				found = true
				for instance := range actual[ssFullName] {
					sort.Slice(actual[ssFullName][instance], func(i, j int) bool { return actual[ssFullName][instance][i] < actual[ssFullName][instance][j] })
					sort.Slice(expected[ss], func(i, j int) bool { return expected[ss][i] < expected[ss][j] })
					if !reflect.DeepEqual(actual[ssFullName][instance], expected[ss]) {
						t.Errorf("unexpected %v != %v", actual[ssFullName][instance], expected[ss])
					}
				}
				break
			}
		}
		if !found {
			t.Errorf("unexpected scaleset %s", ssFullName)
		}
	}
}
