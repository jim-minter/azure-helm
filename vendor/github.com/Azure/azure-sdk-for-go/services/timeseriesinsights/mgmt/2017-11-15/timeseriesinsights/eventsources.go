package timeseriesinsights

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

// EventSourcesClient is the time Series Insights client
type EventSourcesClient struct {
	BaseClient
}

// NewEventSourcesClient creates an instance of the EventSourcesClient client.
func NewEventSourcesClient(subscriptionID string) EventSourcesClient {
	return NewEventSourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEventSourcesClientWithBaseURI creates an instance of the EventSourcesClient client.
func NewEventSourcesClientWithBaseURI(baseURI string, subscriptionID string) EventSourcesClient {
	return EventSourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update an event source under the specified environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// eventSourceName - name of the event source.
// parameters - parameters for creating an event source resource.
func (client EventSourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, parameters BasicEventSourceCreateOrUpdateParameters) (result EventSourceResourceModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventSourcesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: eventSourceName,
			Constraints: []validation.Constraint{{Target: "eventSourceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "eventSourceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "eventSourceName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("timeseriesinsights.EventSourcesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, environmentName, eventSourceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EventSourcesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, parameters BasicEventSourceCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"eventSourceName":   autorest.Encode("path", eventSourceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EventSourcesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EventSourcesClient) CreateOrUpdateResponder(resp *http.Response) (result EventSourceResourceModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the event source with the specified name in the specified subscription, resource group, and
// environment
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// eventSourceName - the name of the Time Series Insights event source associated with the specified
// environment.
func (client EventSourcesClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventSourcesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, environmentName, eventSourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EventSourcesClient) DeletePreparer(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"eventSourceName":   autorest.Encode("path", eventSourceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EventSourcesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EventSourcesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the event source with the specified name in the specified environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// eventSourceName - the name of the Time Series Insights event source associated with the specified
// environment.
func (client EventSourcesClient) Get(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (result EventSourceResourceModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventSourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, environmentName, eventSourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EventSourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"eventSourceName":   autorest.Encode("path", eventSourceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EventSourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EventSourcesClient) GetResponder(resp *http.Response) (result EventSourceResourceModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEnvironment lists all the available event sources associated with the subscription and within the specified
// resource group and environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
func (client EventSourcesClient) ListByEnvironment(ctx context.Context, resourceGroupName string, environmentName string) (result EventSourceListResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventSourcesClient.ListByEnvironment")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByEnvironmentPreparer(ctx, resourceGroupName, environmentName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "ListByEnvironment", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEnvironmentSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "ListByEnvironment", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEnvironmentResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "ListByEnvironment", resp, "Failure responding to request")
	}

	return
}

// ListByEnvironmentPreparer prepares the ListByEnvironment request.
func (client EventSourcesClient) ListByEnvironmentPreparer(ctx context.Context, resourceGroupName string, environmentName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEnvironmentSender sends the ListByEnvironment request. The method will close the
// http.Response Body if it receives an error.
func (client EventSourcesClient) ListByEnvironmentSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByEnvironmentResponder handles the response to the ListByEnvironment request. The method always
// closes the http.Response Body.
func (client EventSourcesClient) ListByEnvironmentResponder(resp *http.Response) (result EventSourceListResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates the event source with the specified name in the specified subscription, resource group, and
// environment.
// Parameters:
// resourceGroupName - name of an Azure Resource group.
// environmentName - the name of the Time Series Insights environment associated with the specified resource
// group.
// eventSourceName - the name of the Time Series Insights event source associated with the specified
// environment.
// eventSourceUpdateParameters - request object that contains the updated information for the event source.
func (client EventSourcesClient) Update(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, eventSourceUpdateParameters EventSourceUpdateParameters) (result EventSourceResourceModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventSourcesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, environmentName, eventSourceName, eventSourceUpdateParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "timeseriesinsights.EventSourcesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EventSourcesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, environmentName string, eventSourceName string, eventSourceUpdateParameters EventSourceUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"environmentName":   autorest.Encode("path", environmentName),
		"eventSourceName":   autorest.Encode("path", eventSourceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-11-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TimeSeriesInsights/environments/{environmentName}/eventSources/{eventSourceName}", pathParameters),
		autorest.WithJSON(eventSourceUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EventSourcesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EventSourcesClient) UpdateResponder(resp *http.Response) (result EventSourceResourceModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
