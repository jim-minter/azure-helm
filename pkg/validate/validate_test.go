package validate

import (
	"errors"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ghodss/yaml"

	"github.com/openshift/openshift-azure/pkg/api"
	v20180930preview "github.com/openshift/openshift-azure/pkg/api/2018-09-30-preview/api"
)

var testOpenShiftClusterYAML = []byte(`---
location: eastus
name: openshift
properties:
  openShiftVersion: v3.10
  publicHostname: openshift.test.example.com
  fqdn: "www.example.com"
  authProfile:
    identityProviders:
    - name: Azure AAD
      provider:
        kind: AADIdentityProvider
        clientId: aadClientId
        secret: aadClientSecret
  routerProfiles:
  - name: default
    publicSubdomain: test.example.com
  masterPoolProfile:
    name: master
    count: 3
    vmSize: Standard_D2s_v3
    osType: Linux
  agentPoolProfiles:
  - name: infra
    role: infra
    count: 2
    vmSize: Standard_D2s_v3
    osType: Linux
  - name: compute
    role: compute
    count: 1
    vmSize: Standard_D2s_v3
    osType: Linux
  servicePrincipalProfile:
    clientID: client_id
    secret: client_secret
`)

func TestValidate(t *testing.T) {
	tests := map[string]struct {
		f            func(*api.OpenShiftManagedCluster)
		expectedErrs []error
		externalOnly bool
	}{
		"test yaml parsing": { // test yaml parsing

		},
		"location": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Location = "" },
			expectedErrs: []error{errors.New(`invalid location ""`)},
		},
		"name": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Name = "" },
			expectedErrs: []error{errors.New(`invalid name ""`)},
		},
		"nil properties": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties = nil },
			expectedErrs: []error{errors.New(`properties cannot be nil`)},
		},
		"openshift config invalid api fqdn": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.FQDN = ""
			},
			expectedErrs: []error{errors.New(`invalid properties.fqdn ""`)},
		},
		"test external only false - invalid fqdn fails": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.FQDN = "()" },
			expectedErrs: []error{errors.New(`invalid properties.fqdn "()"`)},
			externalOnly: false,
		},
		"provisioning state bad": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "bad" },
			expectedErrs: []error{errors.New(`invalid properties.provisioningState "bad"`)},
		},
		"provisioning state Creating": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Creating" },
		},
		"provisioning state Failed": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Failed" },
		},
		"provisioning state Updating": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Updating" },
		},
		"provisioning state Succeeded": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Succeeded" },
		},
		"provisioning state Deleting": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Deleting" },
		},
		"provisioning state Migrating": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Migrating" },
		},
		"provisioning state Upgrading": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "Upgrading" },
		},
		"provisioning state empty": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.ProvisioningState = "" },
		},
		"openshift version good": {
			f: func(oc *api.OpenShiftManagedCluster) { oc.Properties.OpenShiftVersion = "v3.10" },
		},
		"openshift version bad": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.OpenShiftVersion = "" },
			expectedErrs: []error{errors.New(`invalid properties.openShiftVersion ""`)},
		},
		"openshift config empty public hostname": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.PublicHostname = ""
			},
		},
		"openshift config valid public hostname": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.PublicHostname = "www.example.com"
			},
		},
		"openshift config invalid public hostname": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.PublicHostname = "()"
			},
			expectedErrs: []error{errors.New(`invalid properties.publicHostname "()"`)},
		},
		"router profile duplicate names": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles =
					append(oc.Properties.RouterProfiles,
						oc.Properties.RouterProfiles[0])
			},
			expectedErrs: []error{errors.New(`duplicate properties.routerProfiles "default"`)},
		},
		"router profile invalid name": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].Name = "foo"
			},
			// two errors expected here because we require the default profile
			expectedErrs: []error{errors.New(`invalid properties.routerProfiles["foo"]`),
				errors.New(`invalid properties.routerProfiles["default"]`)},
		},
		"router profile empty name": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].Name = ""
			},
			// same as above with 2 errors but additional validate on the individual profile yeilds a third
			// this is not very user friendly but testing as is for now
			// TODO fix
			expectedErrs: []error{errors.New(`invalid properties.routerProfiles[""]`),
				errors.New(`invalid properties.routerProfiles[""].name ""`),
				errors.New(`invalid properties.routerProfiles["default"]`)},
		},
		"router empty public subdomain": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].PublicSubdomain = ""
			},
		},
		"router invalid public subdomain": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].PublicSubdomain = "()"
			},
			expectedErrs: []error{errors.New(`invalid properties.routerProfiles["default"].publicSubdomain "()"`)},
		},
		"router valid public subdomain": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].PublicSubdomain = "example.com"
			},
		},
		"test external only true - unset router profile does not fail": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles = nil
			},
			externalOnly: true,
		},
		"test external only false - unset router profile does fail": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles = nil
			},
			expectedErrs: []error{errors.New(`invalid properties.routerProfiles["default"]`)},
			externalOnly: false,
		},
		"test external only false - invalid router profile does fail": {
			f: func(oc *api.OpenShiftManagedCluster) {
				oc.Properties.RouterProfiles[0].FQDN = "()"
			},
			expectedErrs: []error{errors.New(`invalid properties.routerProfiles["default"].fqdn "()"`)},
			externalOnly: false,
		},
		"agent pool profile invalid name": {
			f: func(oc *api.OpenShiftManagedCluster) {
				app := oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster]
				app.Name = "foo"
				oc.Properties.AgentPoolProfiles["foo"] = app
			},
			expectedErrs: []error{
				errors.New(`invalid properties.agentPoolProfiles["foo"]`),
				errors.New(`invalid properties.agentPoolProfiles["foo"].role "foo"`),
			},
		},
		"agent pool unmatched vnet subnet id": {
			f: func(oc *api.OpenShiftManagedCluster) {
				app := oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster]
				app.VnetSubnetID = "/subscriptions/a/resourceGroups/a/providers/Microsoft.Network/virtualNetworks/a/subnets/a"
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster] = app

				app = oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleInfra]
				app.VnetSubnetID = "/subscriptions/a/resourceGroups/a/providers/Microsoft.Network/virtualNetworks/a/subnets/a"
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleInfra] = app
			},
			expectedErrs: []error{errors.New(`invalid properties.agentPoolProfiles.vnetSubnetID "": all subnets must match when using vnetSubnetID`)},
		},
		"agent pool bad vnet subnet id": {
			f: func(oc *api.OpenShiftManagedCluster) {
				app := oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster]
				app.VnetSubnetID = "foo"
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster] = app

				app = oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleInfra]
				app.VnetSubnetID = "/subscriptions/a/resourceGroups/a/providers/Microsoft.Network/virtualNetworks/a/subnets/a"
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleInfra] = app

				app = oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleCompute]
				app.VnetSubnetID = "/subscriptions/a/resourceGroups/a/providers/Microsoft.Network/virtualNetworks/a/subnets/a"
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleCompute] = app
			},
			expectedErrs: []error{
				errors.New(`invalid properties.agentPoolProfiles["infra"].vnetSubnetID "foo"`),
				errors.New(`invalid properties.agentPoolProfiles.vnetSubnetID "/subscriptions/a/resourceGroups/a/providers/Microsoft.Network/virtualNetworks/a/subnets/a": all subnets must match when using vnetSubnetID`),
			},
		},
		"agent pool bad master count": {
			f: func(oc *api.OpenShiftManagedCluster) {
				app := oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster]
				app.Count = 1
				oc.Properties.AgentPoolProfiles[api.AgentPoolProfileRoleMaster] = app
			},
			expectedErrs: []error{errors.New(`invalid masterPoolProfile.count 1`)},
		},
		"sp nil": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.ServicePrincipalProfile = nil },
			expectedErrs: []error{errors.New(`servicePrincipalProfile cannot be nil`)},
		},
		"sp empty client id": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.ServicePrincipalProfile.ClientID = "" },
			expectedErrs: []error{errors.New(`invalid properties.servicePrincipalProfile.clientId ""`)},
		},
		"sp empty secret": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.ServicePrincipalProfile.Secret = "" },
			expectedErrs: []error{errors.New(`invalid properties.servicePrincipalProfile.secret ""`)},
		},
		//we dont check authProfile because it is non pointer struct. Which is all zero values.
		"authProfile.identityProviders empty": {
			f:            func(oc *api.OpenShiftManagedCluster) { oc.Properties.AuthProfile = &api.AuthProfile{} },
			expectedErrs: []error{errors.New(`invalid properties.authProfile.identityProviders length`)},
		},
		"AADIdentityProvider secret empty": {
			f: func(oc *api.OpenShiftManagedCluster) {
				aadIdentityProvider := &api.AADIdentityProvider{
					ClientID: "test",
					Kind:     "AADIdentityProvider",
					Secret:   "",
				}
				oc.Properties.AuthProfile.IdentityProviders[0].Provider = aadIdentityProvider
				oc.Properties.AuthProfile.IdentityProviders[0].Name = "Azure AD"
			},
			expectedErrs: []error{errors.New(`invalid properties.authProfile.AADIdentityProvider clientId ""`)},
		},
		"AADIdentityProvider clientId empty": {
			f: func(oc *api.OpenShiftManagedCluster) {
				aadIdentityProvider := &api.AADIdentityProvider{
					ClientID: "",
					Kind:     "AADIdentityProvider",
					Secret:   "aadClientSecret",
				}
				oc.Properties.AuthProfile.IdentityProviders[0].Provider = aadIdentityProvider
				oc.Properties.AuthProfile.IdentityProviders[0].Name = "Azure AD"
			},
			expectedErrs: []error{errors.New(`invalid properties.authProfile.AADIdentityProvider clientId ""`)},
		},
	}

	for name, test := range tests {
		var oc *v20180930preview.OpenShiftManagedCluster
		err := yaml.Unmarshal(testOpenShiftClusterYAML, &oc)
		if err != nil {
			t.Fatal(err)
		}

		// TODO we're hoping conversion is correct. Change this to a known valid config
		cs := api.ConvertFromV20180930preview(oc)
		if test.f != nil {
			test.f(cs)
		}
		errs := Validate(cs, nil, test.externalOnly)
		if !reflect.DeepEqual(errs, test.expectedErrs) {
			t.Errorf("%q: expected errors\n%s\nbut received\n%s", name, spew.Sprint(test.expectedErrs), spew.Sprint(errs))
		}
	}
}
