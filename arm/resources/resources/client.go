// Package resources implements the Azure ARM Resources service API version
// 2015-11-01.
//
package resources

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

const (
	// APIVersion is the version of the Resources
	APIVersion = "2015-11-01"

	// DefaultBaseURI is the default URI used for the service Resources
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Resources.
type ManagementClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the ManagementClient client.
func New(subscriptionID string) ManagementClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the ManagementClient client.
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return ManagementClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// CheckExistence checks whether resource exists.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ManagementClient) CheckExistence(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result autorest.Response, err error) {
	req, err := client.CheckExistencePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "CheckExistence", nil, "Failure preparing request")
	}

	resp, err := client.CheckExistenceSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "CheckExistence", resp, "Failure sending request")
	}

	result, err = client.CheckExistenceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "CheckExistence", resp, "Failure responding to request")
	}

	return
}

// CheckExistencePreparer prepares the CheckExistence request.
func (client ManagementClient) CheckExistencePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CheckExistenceSender sends the CheckExistence request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CheckExistenceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckExistenceResponder handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (client ManagementClient) CheckExistenceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CreateOrUpdate create a resource.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity. parameters is create or
// update resource parameters.
func (client ManagementClient) CreateOrUpdate(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string, parameters GenericResource) (result GenericResource, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion, parameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagementClient) CreateOrUpdatePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string, parameters GenericResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagementClient) CreateOrUpdateResponder(resp *http.Response) (result GenericResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete resource and all of its resources.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ManagementClient) Delete(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagementClient) DeletePreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a resource belonging to a resource group.
//
// resourceGroupName is the name of the resource group. The name is case
// insensitive. resourceProviderNamespace is resource identity.
// parentResourcePath is resource identity. resourceType is resource
// identity. resourceName is resource identity.
func (client ManagementClient) Get(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (result GenericResource, err error) {
	req, err := client.GetPreparer(resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, apiVersion)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementClient) GetPreparer(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, apiVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         url.QueryEscape(resourceGroupName),
		"resourceName":              url.QueryEscape(resourceName),
		"resourceProviderNamespace": url.QueryEscape(resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetResponder(resp *http.Response) (result GenericResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all of the resources under a subscription.
//
// filter is the filter to apply on the operation. top is query parameters. If
// null is passed returns all resource groups.
func (client ManagementClient) List(filter string, top *int32) (result ResourceListResult, err error) {
	req, err := client.ListPreparer(filter, top)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagementClient) ListPreparer(filter string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = filter
	}
	if top != nil {
		queryParameters["$top"] = top
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resources"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagementClient) ListResponder(resp *http.Response) (result ResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ManagementClient) ListNextResults(lastResults ResourceListResult) (result ResourceListResult, err error) {
	req, err := lastResults.ResourceListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// MoveResources begin moving resources.To determine whether the operation has
// finished processing the request, call GetLongRunningOperationStatus. This
// method handles polling itself for long-running operations. User can cancel
// polling for long-running calls by passing cancel channel as an argument to
// this method (cancel <-chan struct{}). This channel won't cancel the
// operation, it will only stop the polling.
//
// sourceResourceGroupName is source resource group name. parameters is move
// resources' parameters.
func (client ManagementClient) MoveResources(sourceResourceGroupName string, parameters MoveInfo, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.MoveResourcesPreparer(sourceResourceGroupName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "MoveResources", nil, "Failure preparing request")
	}

	resp, err := client.MoveResourcesSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "resources/ManagementClient", "MoveResources", resp, "Failure sending request")
	}

	result, err = client.MoveResourcesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources/ManagementClient", "MoveResources", resp, "Failure responding to request")
	}

	return
}

// MoveResourcesPreparer prepares the MoveResources request.
func (client ManagementClient) MoveResourcesPreparer(sourceResourceGroupName string, parameters MoveInfo, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceResourceGroupName": url.QueryEscape(sourceResourceGroupName),
		"subscriptionId":          url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{Cancel: cancel}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{sourceResourceGroupName}/moveResources"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// MoveResourcesSender sends the MoveResources request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) MoveResourcesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(autorest.DefaultPollingDelay))
}

// MoveResourcesResponder handles the response to the MoveResources request. The method always
// closes the http.Response Body.
func (client ManagementClient) MoveResourcesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
