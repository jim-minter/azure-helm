package fake

import (
	"os"

	azcompute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	azkeyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	azstorage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2018-02-01/storage"
	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/util/azureclient/compute"
	fakecompute "github.com/openshift/openshift-azure/pkg/util/azureclient/fake/compute"
	fakekeyvault "github.com/openshift/openshift-azure/pkg/util/azureclient/fake/keyvault"
	fakeresources "github.com/openshift/openshift-azure/pkg/util/azureclient/fake/resources"
	fakestorage "github.com/openshift/openshift-azure/pkg/util/azureclient/fake/storage"
	"github.com/openshift/openshift-azure/pkg/util/azureclient/keyvault"
	"github.com/openshift/openshift-azure/pkg/util/azureclient/resources"
	"github.com/openshift/openshift-azure/pkg/util/azureclient/storage"
)

type AzureCloud struct {
	fakecompute.ComputeRP
	fakestorage.StorageRP
	fakekeyvault.VaultRP

	AccountsClient                  storage.AccountsClient
	StorageClient                   storage.Client
	DeploymentsClient               resources.DeploymentsClient
	KeyVaultClient                  keyvault.KeyVaultClient
	VirtualMachineScaleSetVMsClient compute.VirtualMachineScaleSetVMsClient
	VirtualMachineScaleSetsClient   compute.VirtualMachineScaleSetsClient
}

// Cleanup all resources
func (a *AzureCloud) Cleanup() {
	for _, scaleset := range a.ComputeRP.State {
		for instanceID := range scaleset.VmsDir {
			os.RemoveAll(scaleset.VmsDir[instanceID])
		}
	}
}

func NewFakeAzureCloud(log *logrus.Entry, secrets []azkeyvault.SecretBundle) *AzureCloud {
	az := &AzureCloud{
		ComputeRP: fakecompute.ComputeRP{
			State: []fakecompute.ScaleSetState{},
			Ssc:   []azcompute.VirtualMachineScaleSet{},
			Calls: []string{},
			Log:   log,
		},
		VaultRP: fakekeyvault.VaultRP{
			Log:     log,
			Calls:   []string{},
			Secrets: secrets,
		},
		StorageRP: fakestorage.StorageRP{
			Log:   log,
			Calls: []string{},
			Accts: []azstorage.Account{},
			Blobs: map[string]map[string][]byte{},
		},
	}
	az.AccountsClient = fakestorage.NewFakeAccountsClient(&az.StorageRP)
	az.StorageClient = fakestorage.NewFakeStorageClient(&az.StorageRP)
	az.KeyVaultClient = fakekeyvault.NewFakeKeyVaultClient(&az.VaultRP)
	az.VirtualMachineScaleSetVMsClient = fakecompute.NewFakeVirtualMachineScaleSetVMsClient(az.KeyVaultClient, &az.ComputeRP)
	az.VirtualMachineScaleSetsClient = fakecompute.NewFakeVirtualMachineScaleSetsClient(az.VirtualMachineScaleSetVMsClient, &az.ComputeRP)
	az.DeploymentsClient = fakeresources.NewFakeDeploymentsClient(az.VirtualMachineScaleSetsClient, &az.StorageRP)
	return az
}
