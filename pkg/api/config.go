package api

import (
	"github.com/satori/go.uuid"
	"k8s.io/client-go/tools/clientcmd/api/v1"

	"github.com/openshift/openshift-azure/pkg/tls"
)

//go:generate go get k8s.io/gengo/...
//go:generate deepcopy-gen --input-dirs . -O zz_generated.deepcopy

// +k8s:deepcopy-gen=true
type Config struct {
	Version int

	ImageOffer     string
	ImagePublisher string
	ImageSKU       string
	ImageVersion   string

	// for development
	ImageResourceGroup string
	ImageResourceName  string

	Certificates CertificateConfig

	// container images for pods
	MasterEtcdImage             string
	ControlPlaneImage           string
	NodeImage                   string
	ServiceCatalogImage         string
	SyncImage                   string
	TemplateServiceBrokerImage  string
	PrometheusNodeExporterImage string
	RegistryImage               string
	RouterImage                 string
	AzureCLIImage               string
	RegistryConsoleImage        string
	AnsibleServiceBrokerImage   string
	WebConsoleImage             string
	OAuthProxyImage             string
	PrometheusImage             string
	PrometheusAlertBufferImage  string
	PrometheusAlertManagerImage string
	LogBridgeImage              string

	// kubeconfigs
	AdminKubeconfig              *v1.Config
	MasterKubeconfig             *v1.Config
	NodeBootstrapKubeconfig      *v1.Config
	AzureClusterReaderKubeconfig *v1.Config

	// misc control plane configurables
	ServiceAccountKey *tls.PrivateKey
	SessionSecretAuth []byte
	SessionSecretEnc  []byte
	HtPasswd          []byte
	ImageConfigFormat string

	// misc node configurables
	SSHKey *tls.PrivateKey

	// misc infra configurables
	RegistryHTTPSecret             []byte
	AlertManagerProxySessionSecret []byte
	AlertsProxySessionSecret       []byte
	PrometheusProxySessionSecret   []byte
	ServiceCatalogClusterID        UUID
	// random string based configurables
	RegistryStorageAccount     string
	RegistryConsoleOAuthSecret string
	RouterStatsPassword        string
	LoggingWorkspace           string // workspace for Azure Log Analytics resource

	// DNS configurables
	RouterLBCNamePrefix string
	MasterLBCNamePrefix string

	// enriched values which are not present in the external API representation
	TenantID       string
	SubscriptionID string
	ResourceGroup  string

	CloudProviderConf []byte

	ConfigStorageAccount string
}

// CertificateConfig contains all certificate configuration for the cluster.
// +k8s:deepcopy-gen=true
type CertificateConfig struct {
	// CAs
	EtcdCa           CertKeyPair
	Ca               CertKeyPair
	FrontProxyCa     CertKeyPair
	ServiceSigningCa CertKeyPair
	ServiceCatalogCa CertKeyPair

	// etcd certificates
	EtcdServer CertKeyPair
	EtcdPeer   CertKeyPair
	EtcdClient CertKeyPair

	// control plane certificates
	MasterServer         CertKeyPair
	OpenshiftConsole     CertKeyPair
	Admin                CertKeyPair
	AggregatorFrontProxy CertKeyPair
	MasterKubeletClient  CertKeyPair
	MasterProxyClient    CertKeyPair
	OpenShiftMaster      CertKeyPair
	NodeBootstrap        CertKeyPair

	// infra certificates
	Registry                CertKeyPair
	Router                  CertKeyPair
	ServiceCatalogServer    CertKeyPair
	ServiceCatalogAPIClient CertKeyPair

	// misc certificates
	AzureClusterReader CertKeyPair
}

// CertKeyPair is an rsa private key and x509 certificate pair.
type CertKeyPair struct {
	Key  *tls.PrivateKey
	Cert *tls.Certificate
}

func (in *CertKeyPair) DeepCopyInto(out *CertKeyPair) {
	if out == nil {
		out = new(CertKeyPair)
	}
	if in.Key != nil {
		out.Key = in.Key.DeepCopy()
	}
	if in.Cert != nil {
		out.Cert = in.Cert.DeepCopy()
	}
}

type UUID struct {
	uuid.UUID
}

func (in *UUID) DeepCopyInto(out *UUID) {
	*out = *in
}
