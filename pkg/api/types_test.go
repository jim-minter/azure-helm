package api

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

// testContainerService is defined in converterfromv20180930preview_test.go.

var testContainerServiceJSON = []byte(`{
	"id": "id",
	"location": "location",
	"name": "name",
	"plan": {
		"name": "plan.name",
		"product": "plan.product",
		"promotionCode": "plan.promotionCode",
		"publisher": "plan.publisher"
	},
	"tags": {
		"tags.k1": "v1",
		"tags.k2": "v2"
	},
	"type": "type",
	"properties": {
		"provisioningState": "properties.provisioningState",
		"openShiftVersion": "properties.openShiftVersion",
		"publicHostname": "properties.publicHostname",
		"fqdn": "properties.fqdn",
		"routerProfiles": [
			{
				"name": "properties.routerProfiles.0.name",
				"publicSubdomain": "properties.routerProfiles.0.publicSubdomain",
				"fqdn": "properties.routerProfiles.0.fqdn"
			},
			{
				"name": "properties.routerProfiles.1.name",
				"publicSubdomain": "properties.routerProfiles.1.publicSubdomain",
				"fqdn": "properties.routerProfiles.1.fqdn"
			}
		],
		"agentPoolProfiles": {
			"compute": {
				"name": "properties.agentPoolProfiles.compute.name",
				"count": 3,
				"vmSize": "properties.agentPoolProfiles.compute.vmSize",
				"vnetSubnetID": "properties.agentPoolProfiles.compute.vnetSubnetID",
				"osType": "properties.agentPoolProfiles.compute.osType"
			},
			"infra": {
				"name": "properties.agentPoolProfiles.infra.name",
				"count": 2,
				"vmSize": "properties.agentPoolProfiles.infra.vmSize",
				"vnetSubnetID": "properties.agentPoolProfiles.infra.vnetSubnetID",
				"osType": "properties.agentPoolProfiles.infra.osType"
			},
			"master": {
				"name": "properties.agentPoolProfiles.master.name",
				"count": 1,
				"vmSize": "properties.agentPoolProfiles.master.vmSize",
				"vnetSubnetID": "properties.agentPoolProfiles.master.vnetSubnetID",
				"osType": "properties.agentPoolProfiles.master.osType"
			}
		},
		"authProfile": {
			"identityProviders": [
				{
					"name": "properties.authProfile.identityProviders.0.name",
					"provider": {
						"kind": "AADIdentityProvider",
						"clientId": "properties.authProfile.identityProviders.0.provider.clientId",
						"secret": "properties.authProfile.identityProviders.0.provider.secret"
					}
				}
			]
		},
		"servicePrincipalProfile": {
			"clientId": "properties.servicePrincipalProfile.clientId",
			"secret": "properties.servicePrincipalProfile.secret"
		}
	}
}`)

func TestMarshal(t *testing.T) {
	b, err := json.MarshalIndent(testContainerService, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, testContainerServiceJSON) {
		t.Errorf("json.MarshalIndent returned unexpected result\n%s\n", string(b))
	}
}

func TestUnmarshal(t *testing.T) {
	var oc *OpenShiftManagedCluster
	err := json.Unmarshal(testContainerServiceJSON, &oc)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(oc, testContainerService) {
		t.Errorf("json.Unmarshal returned unexpected result\n%#v\n", oc)
	}
}
