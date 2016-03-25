package network

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

// ExpressRouteCircuitsClient is the the Microsoft Azure Network management
// API provides a RESTful set of web services that interact with Microsoft
// Azure Networks service to manage your network resrources. The API has
// entities that capture the relationship between an end user and the
// Microsoft Azure Networks service.
type ExpressRouteCircuitsClient struct {
	ManagementClient
}

// NewExpressRouteCircuitsClient creates an instance of the
// ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClient(subscriptionID string) ExpressRouteCircuitsClient {
	return NewExpressRouteCircuitsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExpressRouteCircuitsClientWithBaseURI creates an instance of the
// ExpressRouteCircuitsClient client.
func NewExpressRouteCircuitsClientWithBaseURI(baseURI string, subscriptionID string) ExpressRouteCircuitsClient {
	return ExpressRouteCircuitsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put ExpressRouteCircuit operation creates/updates a
// ExpressRouteCircuit This method handles polling itself for long-running
// operations. User can cancel polling for long-running calls by passing
// cancel channel as an argument to this method (cancel <-chan struct{}).
// This channel won't cancel the operation, it will only stop the polling.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit. parameters is parameters supplied to the
// create/delete ExpressRouteCircuit operation
func (client ExpressRouteCircuitsClient) CreateOrUpdate(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, circuitName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ExpressRouteCircuitsClient) CreateOrUpdatePreparer(resourceGroupName string, circuitName string, parameters ExpressRouteCircuit, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{Cancel: cancel}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(autorest.DefaultPollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the delete ExpressRouteCircuit operation deletes the specified
// ExpressRouteCircuit. This method handles polling itself for long-running
// operations. User can cancel polling for long-running calls by passing
// cancel channel as an argument to this method (cancel <-chan struct{}).
// This channel won't cancel the operation, it will only stop the polling.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the express route Circuit.
func (client ExpressRouteCircuitsClient) Delete(resourceGroupName string, circuitName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, circuitName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ExpressRouteCircuitsClient) DeletePreparer(resourceGroupName string, circuitName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{Cancel: cancel}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(autorest.DefaultPollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusAccepted, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get ExpressRouteCircuit operation retreives information about the
// specified ExpressRouteCircuit.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsClient) Get(resourceGroupName string, circuitName string) (result ExpressRouteCircuit, err error) {
	req, err := client.GetPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExpressRouteCircuitsClient) GetPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) GetResponder(resp *http.Response) (result ExpressRouteCircuit, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List ExpressRouteCircuit opertion retrieves all the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client ExpressRouteCircuitsClient) List(resourceGroupName string) (result ExpressRouteCircuitListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ExpressRouteCircuitsClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
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
func (client ExpressRouteCircuitsClient) ListNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.ExpressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListAll the List ExpressRouteCircuit opertion retrieves all the
// ExpressRouteCircuits in a subscription.
func (client ExpressRouteCircuitsClient) ListAll() (result ExpressRouteCircuitListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", nil, "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", resp, "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client ExpressRouteCircuitsClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteCircuits"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListAllResponder(resp *http.Response) (result ExpressRouteCircuitListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) ListAllNextResults(lastResults ExpressRouteCircuitListResult) (result ExpressRouteCircuitListResult, err error) {
	req, err := lastResults.ExpressRouteCircuitListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", resp, "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListAll", resp, "Failure responding to next results request request")
	}

	return
}

// ListArpTable the ListArpTable from ExpressRouteCircuit opertion retrieves
// the currently advertised arp table associated with the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsClient) ListArpTable(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsArpTableListResult, err error) {
	req, err := client.ListArpTablePreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", nil, "Failure preparing request")
	}

	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure sending request")
	}

	result, err = client.ListArpTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure responding to request")
	}

	return
}

// ListArpTablePreparer prepares the ListArpTable request.
func (client ExpressRouteCircuitsClient) ListArpTablePreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/arpTable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListArpTableSender sends the ListArpTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListArpTableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListArpTableResponder handles the response to the ListArpTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListArpTableResponder(resp *http.Response) (result ExpressRouteCircuitsArpTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListArpTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) ListArpTableNextResults(lastResults ExpressRouteCircuitsArpTableListResult) (result ExpressRouteCircuitsArpTableListResult, err error) {
	req, err := lastResults.ExpressRouteCircuitsArpTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListArpTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure sending next results request request")
	}

	result, err = client.ListArpTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListArpTable", resp, "Failure responding to next results request request")
	}

	return
}

// ListRoutesTable the ListRoutesTable from ExpressRouteCircuit opertion
// retrieves the currently advertised routes table associated with the
// ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the circuit.
func (client ExpressRouteCircuitsClient) ListRoutesTable(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	req, err := client.ListRoutesTablePreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", nil, "Failure preparing request")
	}

	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure sending request")
	}

	result, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure responding to request")
	}

	return
}

// ListRoutesTablePreparer prepares the ListRoutesTable request.
func (client ExpressRouteCircuitsClient) ListRoutesTablePreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/routesTable"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListRoutesTableSender sends the ListRoutesTable request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListRoutesTableSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListRoutesTableResponder handles the response to the ListRoutesTable request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListRoutesTableResponder(resp *http.Response) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListRoutesTableNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) ListRoutesTableNextResults(lastResults ExpressRouteCircuitsRoutesTableListResult) (result ExpressRouteCircuitsRoutesTableListResult, err error) {
	req, err := lastResults.ExpressRouteCircuitsRoutesTableListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListRoutesTableSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure sending next results request request")
	}

	result, err = client.ListRoutesTableResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListRoutesTable", resp, "Failure responding to next results request request")
	}

	return
}

// ListStats the Liststats ExpressRouteCircuit opertion retrieves all the
// stats from a ExpressRouteCircuits in a resource group.
//
// resourceGroupName is the name of the resource group. circuitName is the
// name of the loadBalancer.
func (client ExpressRouteCircuitsClient) ListStats(resourceGroupName string, circuitName string) (result ExpressRouteCircuitsStatsListResult, err error) {
	req, err := client.ListStatsPreparer(resourceGroupName, circuitName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", nil, "Failure preparing request")
	}

	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", resp, "Failure sending request")
	}

	result, err = client.ListStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", resp, "Failure responding to request")
	}

	return
}

// ListStatsPreparer prepares the ListStats request.
func (client ExpressRouteCircuitsClient) ListStatsPreparer(resourceGroupName string, circuitName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"circuitName":       url.QueryEscape(circuitName),
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteCircuits/{circuitName}/stats"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListStatsSender sends the ListStats request. The method will close the
// http.Response Body if it receives an error.
func (client ExpressRouteCircuitsClient) ListStatsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListStatsResponder handles the response to the ListStats request. The method always
// closes the http.Response Body.
func (client ExpressRouteCircuitsClient) ListStatsResponder(resp *http.Response) (result ExpressRouteCircuitsStatsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListStatsNextResults retrieves the next set of results, if any.
func (client ExpressRouteCircuitsClient) ListStatsNextResults(lastResults ExpressRouteCircuitsStatsListResult) (result ExpressRouteCircuitsStatsListResult, err error) {
	req, err := lastResults.ExpressRouteCircuitsStatsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListStatsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", resp, "Failure sending next results request request")
	}

	result, err = client.ListStatsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network/ExpressRouteCircuitsClient", "ListStats", resp, "Failure responding to next results request request")
	}

	return
}
