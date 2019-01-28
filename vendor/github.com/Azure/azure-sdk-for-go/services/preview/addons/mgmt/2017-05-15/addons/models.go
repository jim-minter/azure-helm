package addons

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/addons/mgmt/2017-05-15/addons"

// PlanTypeName enumerates the values for plan type name.
type PlanTypeName string

const (
	// Advanced ...
	Advanced PlanTypeName = "Advanced"
	// Essential ...
	Essential PlanTypeName = "Essential"
	// Standard ...
	Standard PlanTypeName = "Standard"
)

// PossiblePlanTypeNameValues returns an array of possible values for the PlanTypeName const type.
func PossiblePlanTypeNameValues() []PlanTypeName {
	return []PlanTypeName{Advanced, Essential, Standard}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// Cancelling ...
	Cancelling ProvisioningState = "Cancelling"
	// Downgrading ...
	Downgrading ProvisioningState = "Downgrading"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Purchasing ...
	Purchasing ProvisioningState = "Purchasing"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Upgrading ...
	Upgrading ProvisioningState = "Upgrading"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Cancelled, Cancelling, Downgrading, Failed, Purchasing, Succeeded, Upgrading}
}

// SupportPlanType enumerates the values for support plan type.
type SupportPlanType string

const (
	// SupportPlanTypeAdvanced ...
	SupportPlanTypeAdvanced SupportPlanType = "advanced"
	// SupportPlanTypeEssential ...
	SupportPlanTypeEssential SupportPlanType = "essential"
	// SupportPlanTypeStandard ...
	SupportPlanTypeStandard SupportPlanType = "standard"
)

// PossibleSupportPlanTypeValues returns an array of possible values for the SupportPlanType const type.
func PossibleSupportPlanTypeValues() []SupportPlanType {
	return []SupportPlanType{SupportPlanTypeAdvanced, SupportPlanTypeEssential, SupportPlanTypeStandard}
}

// CanonicalSupportPlanProperties the properties of the Canonical support plan.
type CanonicalSupportPlanProperties struct {
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Failed', 'Cancelled', 'Purchasing', 'Downgrading', 'Cancelling', 'Upgrading'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
}

// CanonicalSupportPlanResponseEnvelope the status of the Canonical support plan.
type CanonicalSupportPlanResponseEnvelope struct {
	autorest.Response `json:"-"`
	// ID - The id of the ARM resource, e.g. "/subscriptions/{id}/providers/Microsoft.Addons/supportProvider/{supportProviderName}/supportPlanTypes/{planTypeName}".
	ID *string `json:"id,omitempty"`
	// Name - The name of the Canonical support plan, i.e. "essential", "standard" or "advanced".
	Name *string `json:"name,omitempty"`
	// Type - Microsoft.Addons/supportProvider
	Type *string `json:"type,omitempty"`
	// CanonicalSupportPlanProperties - Describes Canonical support plan type and status.
	*CanonicalSupportPlanProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for CanonicalSupportPlanResponseEnvelope.
func (cspre CanonicalSupportPlanResponseEnvelope) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cspre.ID != nil {
		objectMap["id"] = cspre.ID
	}
	if cspre.Name != nil {
		objectMap["name"] = cspre.Name
	}
	if cspre.Type != nil {
		objectMap["type"] = cspre.Type
	}
	if cspre.CanonicalSupportPlanProperties != nil {
		objectMap["properties"] = cspre.CanonicalSupportPlanProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CanonicalSupportPlanResponseEnvelope struct.
func (cspre *CanonicalSupportPlanResponseEnvelope) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cspre.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cspre.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				cspre.Type = &typeVar
			}
		case "properties":
			if v != nil {
				var canonicalSupportPlanProperties CanonicalSupportPlanProperties
				err = json.Unmarshal(*v, &canonicalSupportPlanProperties)
				if err != nil {
					return err
				}
				cspre.CanonicalSupportPlanProperties = &canonicalSupportPlanProperties
			}
		}
	}

	return nil
}

// CanonicalSupportPlanStatusItem ...
type CanonicalSupportPlanStatusItem struct {
	// SupportPlanType - Support plan type. Possible values include: 'SupportPlanTypeEssential', 'SupportPlanTypeStandard', 'SupportPlanTypeAdvanced'
	SupportPlanType SupportPlanType `json:"supportPlanType,omitempty"`
	// Enabled - Whether the support plan is enabled for this subscription.
	Enabled *bool `json:"enabled,omitempty"`
	// WillEmitOneTimeChargeWhenEnabled - This indicates that when this support plan is enabled if AddonsRP will emit a one-time charge.
	WillEmitOneTimeChargeWhenEnabled *bool `json:"willEmitOneTimeChargeWhenEnabled,omitempty"`
	// WillEmitOneTimeChargeIfReEnabled - This indicates that when this support plan is cancelled and then enabled that AddonsRP will emit a one-time charge.
	WillEmitOneTimeChargeIfReEnabled *bool `json:"willEmitOneTimeChargeIfReEnabled,omitempty"`
}

// ErrorDefinition error description and code explaining why an operation failed.
type ErrorDefinition struct {
	// Message - Description of the error.
	Message *string `json:"message,omitempty"`
	// Code - Service specific error code which serves as the substatus for the HTTP error code.
	Code *int32 `json:"code,omitempty"`
}

// ListCanonicalSupportPlanStatusItem ...
type ListCanonicalSupportPlanStatusItem struct {
	autorest.Response `json:"-"`
	Value             *[]CanonicalSupportPlanStatusItem `json:"value,omitempty"`
}

// ListOperationsDefinition ...
type ListOperationsDefinition struct {
	autorest.Response `json:"-"`
	Value             *[]OperationsDefinition `json:"value,omitempty"`
}

// OperationsDefinition definition object with the name and properties of an operation.
type OperationsDefinition struct {
	// Name - Name of the operation.
	Name *string `json:"name,omitempty"`
	// Display - Display object with properties of the operation.
	Display *OperationsDisplayDefinition `json:"display,omitempty"`
}

// OperationsDisplayDefinition display object with properties of the operation.
type OperationsDisplayDefinition struct {
	// Provider - Resource provider of the operation.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource for the operation.
	Resource *string `json:"resource,omitempty"`
	// Operation - Short description of the operation.
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation.
	Description *string `json:"description,omitempty"`
}

// SupportPlanTypesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type SupportPlanTypesCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SupportPlanTypesCreateOrUpdateFuture) Result(client SupportPlanTypesClient) (cspre CanonicalSupportPlanResponseEnvelope, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("addons.SupportPlanTypesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cspre.Response.Response, err = future.GetResult(sender); err == nil && cspre.Response.Response.StatusCode != http.StatusNoContent {
		cspre, err = client.CreateOrUpdateResponder(cspre.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesCreateOrUpdateFuture", "Result", cspre.Response.Response, "Failure responding to request")
		}
	}
	return
}

// SupportPlanTypesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type SupportPlanTypesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *SupportPlanTypesDeleteFuture) Result(client SupportPlanTypesClient) (cspre CanonicalSupportPlanResponseEnvelope, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("addons.SupportPlanTypesDeleteFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if cspre.Response.Response, err = future.GetResult(sender); err == nil && cspre.Response.Response.StatusCode != http.StatusNoContent {
		cspre, err = client.DeleteResponder(cspre.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "addons.SupportPlanTypesDeleteFuture", "Result", cspre.Response.Response, "Failure responding to request")
		}
	}
	return
}
