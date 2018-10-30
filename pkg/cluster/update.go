package cluster

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/log"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
	"github.com/openshift/openshift-azure/pkg/util/managedcluster"
)

func (u *simpleUpgrader) Update(ctx context.Context, cs *api.OpenShiftManagedCluster, azuretemplate map[string]interface{}, deployFn api.DeployFn) error {
	err := u.createClients(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepClientCreation}
	}
	// deployFn() may change the number of VMs.  If we can see that any VMs are
	// about to be deleted, drain them first.  Record which VMs are visible now
	// so that we can detect newly created VMs and wait for them to become ready.
	vmsBefore, err := u.getNodesAndDrain(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepDrain}
	}
	err = deployFn(ctx, azuretemplate)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepDeploy}
	}
	err = u.Initialize(ctx, cs)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepInitialize}
	}
	err = managedcluster.WaitForHealthz(ctx, cs.Config.AdminKubeconfig)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepWaitForWaitForOpenShiftAPI}
	}
	err = u.waitForNewNodes(ctx, cs, vmsBefore)
	if err != nil {
		return &api.PluginError{Err: err, Step: api.PluginStepWaitForNodes}
	}

	// For PP day 1, scale is permitted but not any other sort of update.  When
	// we enable configuration changes and/or upgrades, uncomment this code.  At
	// the same time, current thinking is that we will add a hash-based
	// mechanism to avoid unnecessary VM rotations as well.
	if os.Getenv("RUNNING_UNDER_TEST") != "" {
		err = u.updateInPlace(ctx, cs, api.AgentPoolProfileRoleMaster)
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepUpdateMasterVMRotation}
		}
		err = u.updatePlusOne(ctx, cs, api.AgentPoolProfileRoleInfra)
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepUpdateInfraVMRotation}
		}
		err = u.updatePlusOne(ctx, cs, api.AgentPoolProfileRoleCompute)
		if err != nil {
			return &api.PluginError{Err: err, Step: api.PluginStepUpdateComputeVMRotation}
		}
	}
	return nil
}

func (u *simpleUpgrader) getNodesAndDrain(ctx context.Context, cs *api.OpenShiftManagedCluster) (map[string]struct{}, error) {
	vmsBefore := map[string]struct{}{}

	for _, agent := range cs.Properties.AgentPoolProfiles {
		vms, err := listVMs(ctx, cs, u.vmc, agent.Role)
		if err != nil {
			return nil, err
		}

		for i, vm := range vms {
			if i < agent.Count {
				vmsBefore[*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName] = struct{}{}
			} else {
				err = u.delete(ctx, cs, agent.Role, *vm.InstanceID, *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return vmsBefore, nil
}

func (u *simpleUpgrader) waitForNewNodes(ctx context.Context, cs *api.OpenShiftManagedCluster, nodes map[string]struct{}) error {
	for _, agent := range cs.Properties.AgentPoolProfiles {
		vms, err := listVMs(ctx, cs, u.vmc, agent.Role)
		if err != nil {
			return err
		}

		// wait for newly created VMs to reach readiness
		for _, vm := range vms {
			hostname := *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName
			if _, found := nodes[hostname]; !found {
				log.Infof("waiting for %s to be ready", hostname)
				err = waitForReady(ctx, cs, agent.Role, hostname, u.kubeclient)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func getCount(cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole) int {
	for _, app := range cs.Properties.AgentPoolProfiles {
		if app.Role == role {
			return app.Count
		}
	}

	panic("invalid role")
}

func listVMs(ctx context.Context, cs *api.OpenShiftManagedCluster, vmc azureclient.VirtualMachineScaleSetVMsClient, role api.AgentPoolProfileRole) ([]compute.VirtualMachineScaleSetVM, error) {
	vmPages, err := vmc.List(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), "", "", "")
	if err != nil {
		return nil, err
	}

	var vms []compute.VirtualMachineScaleSetVM
	for vmPages.NotDone() {
		vms = append(vms, vmPages.Values()...)

		err = vmPages.Next()
		if err != nil {
			return nil, err
		}
	}

	return vms, nil
}

// updatePlusOne creates new VMs and removes old VMs one by one.
func (u *simpleUpgrader) updatePlusOne(ctx context.Context, cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole) error {
	count := getCount(cs, role)

	// store a list of all the VM instances now, so that if we end up creating
	// new ones (in the crash recovery case, we might not), we can detect which
	// they are
	oldVMs, err := listVMs(ctx, cs, u.vmc, role)
	if err != nil {
		return err
	}

	ss, err := u.ssc.Get(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role))
	if err != nil {
		return err
	}

	cm, err := u.readConfigMap()
	if err != nil {
		return err
	}

	// Filter out VMs that do not need to get upgraded. Should speed
	// up retrying failed upgrades.
	oldVMs, err = u.filterOldVMs(ctx, cs.Properties.AzProfile.ResourceGroup, ss, oldVMs, role, cm)
	if err != nil {
		return err
	}
	vmsBefore := map[string]struct{}{}
	for _, vm := range oldVMs {
		vmsBefore[*vm.InstanceID] = struct{}{}
	}

	for _, vm := range oldVMs {
		log.Infof("setting ss-%s capacity to %d", role, count+1)
		future, err := u.ssc.Update(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), compute.VirtualMachineScaleSetUpdate{
			Sku: &compute.Sku{
				Capacity: to.Int64Ptr(int64(count) + 1),
			},
		})
		if err != nil {
			return err
		}

		if err := future.WaitForCompletionRef(ctx, u.ssc.Client()); err != nil {
			return err
		}

		updatedList, err := listVMs(ctx, cs, u.vmc, role)
		if err != nil {
			return err
		}

		// wait for newly created VMs to reach readiness (n.b. one alternative to
		// this approach would be for the CSE to not return until the node is
		// ready, but that is also problematic)
		for _, updated := range updatedList {
			if _, found := vmsBefore[*updated.InstanceID]; !found {
				log.Infof("waiting for %s to be ready", *updated.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
				err = waitForReady(ctx, cs, role, *updated.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName, u.kubeclient)
				if err != nil {
					return err
				}
				vmsBefore[*updated.InstanceID] = struct{}{}
				if hash := ss.Tags[arm.HashKey]; hash != nil && *hash != "" {
					cm.Data[*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName] = *hash
					if err := u.updateConfigMap(cm.Data); err != nil {
						log.Warn(err)
					}
				}
			}
		}

		if err := u.delete(ctx, cs, role, *vm.InstanceID, *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName); err != nil {
			return err
		}
		delete(cm.Data, *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
		if err := u.updateConfigMap(cm.Data); err != nil {
			log.Warn(err)
		}
	}

	return nil
}

func (u *simpleUpgrader) filterOldVMs(ctx context.Context, resourceGroup string, ss compute.VirtualMachineScaleSet, vms []compute.VirtualMachineScaleSetVM, role api.AgentPoolProfileRole, cm *v1.ConfigMap) ([]compute.VirtualMachineScaleSetVM, error) {
	var oldVMs []compute.VirtualMachineScaleSetVM
	if ssTag := ss.Tags["scaleset-checksum"]; ssTag != nil && *ssTag != "" {
		for _, vm := range vms {
			if !vmMatchesScaleSet(vm, ss, cm) {
				oldVMs = append(oldVMs, vm)
			} else {
				log.Debugf("skipping vm %q since it's already updated", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			}
		}
	} else {
		log.Warnf("scale set %q is not tagged with its hash - unnecessary VM rotations may occur", "ss-"+string(role))
		oldVMs = vms
	}
	return oldVMs, nil
}

func vmMatchesScaleSet(vm compute.VirtualMachineScaleSetVM, ss compute.VirtualMachineScaleSet, cm *v1.ConfigMap) bool {
	vmHash := cm.Data[*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName]
	ssHash := ss.Tags["scaleset-checksum"]
	return vmHash != "" && ssHash != nil && vmHash == *ssHash
}

// updateInPlace updates one by one all the VMs of a scale set, in place.
func (u *simpleUpgrader) updateInPlace(ctx context.Context, cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole) error {
	vms, err := listVMs(ctx, cs, u.vmc, role)
	if err != nil {
		return err
	}

	ss, err := u.ssc.Get(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role))
	if err != nil {
		return err
	}

	cm, err := u.readConfigMap()
	if err != nil {
		return err
	}

	sorted, err := sortMasterVMsByHealth(vms, cs, u.kubeclient)
	if err != nil {
		return err
	}

	for _, vm := range sorted {
		if isUpdated := vmMatchesScaleSet(vm, ss, cm); isUpdated {
			log.Infof("skipping already updated %q", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			continue
		}

		log.Infof("draining %s", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
		err = u.drain(ctx, cs, role, *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
		if err != nil {
			return err
		}

		{
			log.Infof("deallocating %s (%s)", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName, *vm.InstanceID)
			future, err := u.vmc.Deallocate(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), *vm.InstanceID)
			if err != nil {
				return err
			}

			err = future.WaitForCompletionRef(ctx, u.vmc.Client())
			if err != nil {
				return err
			}
		}

		{
			log.Infof("updating %s", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			future, err := u.ssc.UpdateInstances(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), compute.VirtualMachineScaleSetVMInstanceRequiredIDs{
				InstanceIds: &[]string{*vm.InstanceID},
			})
			if err != nil {
				return err
			}

			err = future.WaitForCompletionRef(ctx, u.ssc.Client())
			if err != nil {
				return err
			}
		}

		{
			log.Infof("reimaging %s", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			future, err := u.vmc.Reimage(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), *vm.InstanceID)
			if err != nil {
				return err
			}

			err = future.WaitForCompletionRef(ctx, u.vmc.Client())
			if err != nil {
				return err
			}
		}

		{
			log.Infof("starting %s", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
			future, err := u.vmc.Start(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), *vm.InstanceID)
			if err != nil {
				return err
			}

			err = future.WaitForCompletionRef(ctx, u.vmc.Client())
			if err != nil {
				return err
			}
		}

		log.Infof("waiting for %s to be ready", *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName)
		err = waitForReady(ctx, cs, role, *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName, u.kubeclient)
		if err != nil {
			return err
		}

		if hash := ss.Tags[arm.HashKey]; hash != nil && *hash != "" {
			cm.Data[*vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName] = *hash
			if err := u.updateConfigMap(cm.Data); err != nil {
				log.Warn(err)
			}
		}
	}

	return nil
}

func sortMasterVMsByHealth(vms []compute.VirtualMachineScaleSetVM, cs *api.OpenShiftManagedCluster, kc kubernetes.Interface) ([]compute.VirtualMachineScaleSetVM, error) {
	var ready, unready []compute.VirtualMachineScaleSetVM
	for _, vm := range vms {
		nodeName := *vm.VirtualMachineScaleSetVMProperties.OsProfile.ComputerName
		isReady, err := masterIsReady(kc, nodeName)
		if err != nil {
			return nil, fmt.Errorf("cannot get health for %q: %v", nodeName, err)
		}
		if isReady {
			ready = append(ready, vm)
		} else {
			unready = append(unready, vm)
		}
	}

	return append(unready, ready...), nil
}

func (u *simpleUpgrader) delete(ctx context.Context, cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole, instanceID, nodeName string) error {
	log.Infof("draining %s", nodeName)
	if err := u.drain(ctx, cs, role, nodeName); err != nil {
		return err
	}

	log.Infof("deleting %s", nodeName)
	future, err := u.vmc.Delete(ctx, cs.Properties.AzProfile.ResourceGroup, "ss-"+string(role), instanceID)
	if err != nil {
		return err
	}

	return future.WaitForCompletionRef(ctx, u.vmc.Client())
}
