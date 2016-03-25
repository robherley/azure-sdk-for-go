// Package intune implements the Azure ARM Intune service API version
// 2015-01-14-preview.
//
// Microsoft.Intune Resource provider Api features in the swagger-2.0
// specification
package intune

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
	// APIVersion is the version of the Intune
	APIVersion = "2015-01-14-preview"

	// DefaultBaseURI is the default URI used for the service Intune
	DefaultBaseURI = "https://management.azure.com"
)

// ManagementClient is the base client for Intune.
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

// GetApps returns Intune Manageable apps.
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client ManagementClient) GetApps(hostName string, filter string, top *int32, selectParameter string) (result ApplicationCollection, err error) {
	req, err := client.GetAppsPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", nil, "Failure preparing request")
	}

	resp, err := client.GetAppsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", resp, "Failure sending request")
	}

	result, err = client.GetAppsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", resp, "Failure responding to request")
	}

	return
}

// GetAppsPreparer prepares the GetApps request.
func (client ManagementClient) GetAppsPreparer(hostName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/apps"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetAppsSender sends the GetApps request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetAppsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAppsResponder handles the response to the GetApps request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetAppsResponder(resp *http.Response) (result ApplicationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAppsNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetAppsNextResults(lastResults ApplicationCollection) (result ApplicationCollection, err error) {
	req, err := lastResults.ApplicationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetAppsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", resp, "Failure sending next results request request")
	}

	result, err = client.GetAppsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetApps", resp, "Failure responding to next results request request")
	}

	return
}

// GetLocationByHostName returns location for given tenant.
func (client ManagementClient) GetLocationByHostName() (result Location, err error) {
	req, err := client.GetLocationByHostNamePreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocationByHostName", nil, "Failure preparing request")
	}

	resp, err := client.GetLocationByHostNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocationByHostName", resp, "Failure sending request")
	}

	result, err = client.GetLocationByHostNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocationByHostName", resp, "Failure responding to request")
	}

	return
}

// GetLocationByHostNamePreparer prepares the GetLocationByHostName request.
func (client ManagementClient) GetLocationByHostNamePreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/hostName"),
		autorest.WithQueryParameters(queryParameters))
}

// GetLocationByHostNameSender sends the GetLocationByHostName request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetLocationByHostNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetLocationByHostNameResponder handles the response to the GetLocationByHostName request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetLocationByHostNameResponder(resp *http.Response) (result Location, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLocations returns location for user tenant.
func (client ManagementClient) GetLocations() (result LocationCollection, err error) {
	req, err := client.GetLocationsPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", nil, "Failure preparing request")
	}

	resp, err := client.GetLocationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", resp, "Failure sending request")
	}

	result, err = client.GetLocationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", resp, "Failure responding to request")
	}

	return
}

// GetLocationsPreparer prepares the GetLocations request.
func (client ManagementClient) GetLocationsPreparer() (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations"),
		autorest.WithQueryParameters(queryParameters))
}

// GetLocationsSender sends the GetLocations request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetLocationsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetLocationsResponder handles the response to the GetLocations request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetLocationsResponder(resp *http.Response) (result LocationCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLocationsNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetLocationsNextResults(lastResults LocationCollection) (result LocationCollection, err error) {
	req, err := lastResults.LocationCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetLocationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", resp, "Failure sending next results request request")
	}

	result, err = client.GetLocationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetLocations", resp, "Failure responding to next results request request")
	}

	return
}

// GetMAMFlaggedUserByName returns Intune flagged user details
//
// hostName is location hostName for the tenant userName is flagged userName
// selectParameter is select specific fields in entity.
func (client ManagementClient) GetMAMFlaggedUserByName(hostName string, userName string, selectParameter string) (result FlaggedUser, err error) {
	req, err := client.GetMAMFlaggedUserByNamePreparer(hostName, userName, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUserByName", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMFlaggedUserByNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUserByName", resp, "Failure sending request")
	}

	result, err = client.GetMAMFlaggedUserByNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUserByName", resp, "Failure responding to request")
	}

	return
}

// GetMAMFlaggedUserByNamePreparer prepares the GetMAMFlaggedUserByName request.
func (client ManagementClient) GetMAMFlaggedUserByNamePreparer(hostName string, userName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"userName": url.QueryEscape(userName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers/{userName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMFlaggedUserByNameSender sends the GetMAMFlaggedUserByName request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMFlaggedUserByNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMFlaggedUserByNameResponder handles the response to the GetMAMFlaggedUserByName request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMFlaggedUserByNameResponder(resp *http.Response) (result FlaggedUser, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMFlaggedUsers returns Intune flagged user collection
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client ManagementClient) GetMAMFlaggedUsers(hostName string, filter string, top *int32, selectParameter string) (result FlaggedUserCollection, err error) {
	req, err := client.GetMAMFlaggedUsersPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMFlaggedUsersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", resp, "Failure sending request")
	}

	result, err = client.GetMAMFlaggedUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", resp, "Failure responding to request")
	}

	return
}

// GetMAMFlaggedUsersPreparer prepares the GetMAMFlaggedUsers request.
func (client ManagementClient) GetMAMFlaggedUsersPreparer(hostName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMFlaggedUsersSender sends the GetMAMFlaggedUsers request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMFlaggedUsersSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMFlaggedUsersResponder handles the response to the GetMAMFlaggedUsers request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMFlaggedUsersResponder(resp *http.Response) (result FlaggedUserCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMFlaggedUsersNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetMAMFlaggedUsersNextResults(lastResults FlaggedUserCollection) (result FlaggedUserCollection, err error) {
	req, err := lastResults.FlaggedUserCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMFlaggedUsersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", resp, "Failure sending next results request request")
	}

	result, err = client.GetMAMFlaggedUsersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMFlaggedUsers", resp, "Failure responding to next results request request")
	}

	return
}

// GetMAMStatuses returns Intune Tenant level statuses.
//
// hostName is location hostName for the tenant
func (client ManagementClient) GetMAMStatuses(hostName string) (result StatusesDefault, err error) {
	req, err := client.GetMAMStatusesPreparer(hostName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMStatusesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", resp, "Failure sending request")
	}

	result, err = client.GetMAMStatusesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", resp, "Failure responding to request")
	}

	return
}

// GetMAMStatusesPreparer prepares the GetMAMStatuses request.
func (client ManagementClient) GetMAMStatusesPreparer(hostName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/statuses/default"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMStatusesSender sends the GetMAMStatuses request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMStatusesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMStatusesResponder handles the response to the GetMAMStatuses request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMStatusesResponder(resp *http.Response) (result StatusesDefault, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMStatusesNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetMAMStatusesNextResults(lastResults StatusesDefault) (result StatusesDefault, err error) {
	req, err := lastResults.StatusesDefaultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMStatusesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", resp, "Failure sending next results request request")
	}

	result, err = client.GetMAMStatusesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMStatuses", resp, "Failure responding to next results request request")
	}

	return
}

// GetMAMUserDeviceByDeviceName get a unique device for a user.
//
// hostName is location hostName for the tenant userName is unique user name
// deviceName is device name selectParameter is select specific fields in
// entity.
func (client ManagementClient) GetMAMUserDeviceByDeviceName(hostName string, userName string, deviceName string, selectParameter string) (result Device, err error) {
	req, err := client.GetMAMUserDeviceByDeviceNamePreparer(hostName, userName, deviceName, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDeviceByDeviceName", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMUserDeviceByDeviceNameSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDeviceByDeviceName", resp, "Failure sending request")
	}

	result, err = client.GetMAMUserDeviceByDeviceNameResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDeviceByDeviceName", resp, "Failure responding to request")
	}

	return
}

// GetMAMUserDeviceByDeviceNamePreparer prepares the GetMAMUserDeviceByDeviceName request.
func (client ManagementClient) GetMAMUserDeviceByDeviceNamePreparer(hostName string, userName string, deviceName string, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName": url.QueryEscape(deviceName),
		"hostName":   url.QueryEscape(hostName),
		"userName":   url.QueryEscape(userName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/users/{userName}/devices/{deviceName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMUserDeviceByDeviceNameSender sends the GetMAMUserDeviceByDeviceName request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMUserDeviceByDeviceNameSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMUserDeviceByDeviceNameResponder handles the response to the GetMAMUserDeviceByDeviceName request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMUserDeviceByDeviceNameResponder(resp *http.Response) (result Device, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMUserDevices get devices for a user.
//
// hostName is location hostName for the tenant userName is user unique Name
// filter is the filter to apply on the operation. selectParameter is select
// specific fields in entity.
func (client ManagementClient) GetMAMUserDevices(hostName string, userName string, filter string, top *int32, selectParameter string) (result DeviceCollection, err error) {
	req, err := client.GetMAMUserDevicesPreparer(hostName, userName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMUserDevicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", resp, "Failure sending request")
	}

	result, err = client.GetMAMUserDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", resp, "Failure responding to request")
	}

	return
}

// GetMAMUserDevicesPreparer prepares the GetMAMUserDevices request.
func (client ManagementClient) GetMAMUserDevicesPreparer(hostName string, userName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"userName": url.QueryEscape(userName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/users/{userName}/devices"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMUserDevicesSender sends the GetMAMUserDevices request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMUserDevicesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMUserDevicesResponder handles the response to the GetMAMUserDevices request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMUserDevicesResponder(resp *http.Response) (result DeviceCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMUserDevicesNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetMAMUserDevicesNextResults(lastResults DeviceCollection) (result DeviceCollection, err error) {
	req, err := lastResults.DeviceCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMUserDevicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", resp, "Failure sending next results request request")
	}

	result, err = client.GetMAMUserDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserDevices", resp, "Failure responding to next results request request")
	}

	return
}

// GetMAMUserFlaggedEnrolledApps returns Intune flagged enrolled app
// collection for the User
//
// hostName is location hostName for the tenant userName is user name for the
// tenant filter is the filter to apply on the operation. selectParameter is
// select specific fields in entity.
func (client ManagementClient) GetMAMUserFlaggedEnrolledApps(hostName string, userName string, filter string, top *int32, selectParameter string) (result FlaggedEnrolledAppCollection, err error) {
	req, err := client.GetMAMUserFlaggedEnrolledAppsPreparer(hostName, userName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", nil, "Failure preparing request")
	}

	resp, err := client.GetMAMUserFlaggedEnrolledAppsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", resp, "Failure sending request")
	}

	result, err = client.GetMAMUserFlaggedEnrolledAppsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", resp, "Failure responding to request")
	}

	return
}

// GetMAMUserFlaggedEnrolledAppsPreparer prepares the GetMAMUserFlaggedEnrolledApps request.
func (client ManagementClient) GetMAMUserFlaggedEnrolledAppsPreparer(hostName string, userName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
		"userName": url.QueryEscape(userName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/flaggedUsers/{userName}/flaggedEnrolledApps"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetMAMUserFlaggedEnrolledAppsSender sends the GetMAMUserFlaggedEnrolledApps request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetMAMUserFlaggedEnrolledAppsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetMAMUserFlaggedEnrolledAppsResponder handles the response to the GetMAMUserFlaggedEnrolledApps request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetMAMUserFlaggedEnrolledAppsResponder(resp *http.Response) (result FlaggedEnrolledAppCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetMAMUserFlaggedEnrolledAppsNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetMAMUserFlaggedEnrolledAppsNextResults(lastResults FlaggedEnrolledAppCollection) (result FlaggedEnrolledAppCollection, err error) {
	req, err := lastResults.FlaggedEnrolledAppCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetMAMUserFlaggedEnrolledAppsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", resp, "Failure sending next results request request")
	}

	result, err = client.GetMAMUserFlaggedEnrolledAppsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetMAMUserFlaggedEnrolledApps", resp, "Failure responding to next results request request")
	}

	return
}

// GetOperationResults returns operationResults.
//
// hostName is location hostName for the tenant filter is the filter to apply
// on the operation. selectParameter is select specific fields in entity.
func (client ManagementClient) GetOperationResults(hostName string, filter string, top *int32, selectParameter string) (result OperationResultCollection, err error) {
	req, err := client.GetOperationResultsPreparer(hostName, filter, top, selectParameter)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", nil, "Failure preparing request")
	}

	resp, err := client.GetOperationResultsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", resp, "Failure sending request")
	}

	result, err = client.GetOperationResultsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", resp, "Failure responding to request")
	}

	return
}

// GetOperationResultsPreparer prepares the GetOperationResults request.
func (client ManagementClient) GetOperationResultsPreparer(hostName string, filter string, top *int32, selectParameter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"hostName": url.QueryEscape(hostName),
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
	if len(selectParameter) > 0 {
		queryParameters["$select"] = selectParameter
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/operationResults"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetOperationResultsSender sends the GetOperationResults request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) GetOperationResultsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetOperationResultsResponder handles the response to the GetOperationResults request. The method always
// closes the http.Response Body.
func (client ManagementClient) GetOperationResultsResponder(resp *http.Response) (result OperationResultCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetOperationResultsNextResults retrieves the next set of results, if any.
func (client ManagementClient) GetOperationResultsNextResults(lastResults OperationResultCollection) (result OperationResultCollection, err error) {
	req, err := lastResults.OperationResultCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.GetOperationResultsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", resp, "Failure sending next results request request")
	}

	result, err = client.GetOperationResultsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "GetOperationResults", resp, "Failure responding to next results request request")
	}

	return
}

// WipeMAMUserDevice wipe a device for a user.
//
// hostName is location hostName for the tenant userName is unique user name
// deviceName is device name
func (client ManagementClient) WipeMAMUserDevice(hostName string, userName string, deviceName string) (result WipeDeviceOperationResult, err error) {
	req, err := client.WipeMAMUserDevicePreparer(hostName, userName, deviceName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "WipeMAMUserDevice", nil, "Failure preparing request")
	}

	resp, err := client.WipeMAMUserDeviceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "intune/ManagementClient", "WipeMAMUserDevice", resp, "Failure sending request")
	}

	result, err = client.WipeMAMUserDeviceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "intune/ManagementClient", "WipeMAMUserDevice", resp, "Failure responding to request")
	}

	return
}

// WipeMAMUserDevicePreparer prepares the WipeMAMUserDevice request.
func (client ManagementClient) WipeMAMUserDevicePreparer(hostName string, userName string, deviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deviceName": url.QueryEscape(deviceName),
		"hostName":   url.QueryEscape(hostName),
		"userName":   url.QueryEscape(userName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	req := http.Request{}
	return autorest.Prepare(&req,
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Intune/locations/{hostName}/users/{userName}/devices/{deviceName}/wipe"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// WipeMAMUserDeviceSender sends the WipeMAMUserDevice request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementClient) WipeMAMUserDeviceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// WipeMAMUserDeviceResponder handles the response to the WipeMAMUserDevice request. The method always
// closes the http.Response Body.
func (client ManagementClient) WipeMAMUserDeviceResponder(resp *http.Response) (result WipeDeviceOperationResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
