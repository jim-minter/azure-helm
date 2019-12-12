package apimanagement

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
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UserIdentitiesClient is the apiManagement Client
type UserIdentitiesClient struct {
	BaseClient
}

// NewUserIdentitiesClient creates an instance of the UserIdentitiesClient client.
func NewUserIdentitiesClient(subscriptionID string) UserIdentitiesClient {
	return NewUserIdentitiesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUserIdentitiesClientWithBaseURI creates an instance of the UserIdentitiesClient client.
func NewUserIdentitiesClientWithBaseURI(baseURI string, subscriptionID string) UserIdentitiesClient {
	return UserIdentitiesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByUsers lists all user identities.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// UID - user identifier. Must be unique in the current API Management service instance.
func (client UserIdentitiesClient) ListByUsers(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result UserIdentityCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UserIdentitiesClient.ListByUsers")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: UID,
			Constraints: []validation.Constraint{{Target: "UID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "UID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "UID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.UserIdentitiesClient", "ListByUsers", err.Error())
	}

	req, err := client.ListByUsersPreparer(ctx, resourceGroupName, serviceName, UID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "ListByUsers", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByUsersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "ListByUsers", resp, "Failure sending request")
		return
	}

	result, err = client.ListByUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.UserIdentitiesClient", "ListByUsers", resp, "Failure responding to request")
	}

	return
}

// ListByUsersPreparer prepares the ListByUsers request.
func (client UserIdentitiesClient) ListByUsersPreparer(ctx context.Context, resourceGroupName string, serviceName string, UID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"uid":               autorest.Encode("path", UID),
	}

	const APIVersion = "2016-10-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{uid}/identities", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByUsersSender sends the ListByUsers request. The method will close the
// http.Response Body if it receives an error.
func (client UserIdentitiesClient) ListByUsersSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListByUsersResponder handles the response to the ListByUsers request. The method always
// closes the http.Response Body.
func (client UserIdentitiesClient) ListByUsersResponder(resp *http.Response) (result UserIdentityCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
