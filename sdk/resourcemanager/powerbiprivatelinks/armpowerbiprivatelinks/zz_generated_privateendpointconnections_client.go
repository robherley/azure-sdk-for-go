//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	host                string
	subscriptionID      string
	resourceGroupName   string
	azureResourceName   string
	privateEndpointName string
	pl                  runtime.Pipeline
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// resourceGroupName - The name of the resource group.
// azureResourceName - The name of the Azure resource.
// privateEndpointName - The name of the private endpoint.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, resourceGroupName string, azureResourceName string, privateEndpointName string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateEndpointConnectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID:      subscriptionID,
		resourceGroupName:   resourceGroupName,
		azureResourceName:   azureResourceName,
		privateEndpointName: privateEndpointName,
		host:                string(cp.Endpoint),
		pl:                  armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Create - Updates the status of Private Endpoint Connection object. Used to approve or reject a connection.
// If the operation fails it returns an *azcore.ResponseError type.
// privateEndpointConnection - Private endpoint connection object to update.
// options - PrivateEndpointConnectionsClientCreateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Create
// method.
func (client *PrivateEndpointConnectionsClient) Create(ctx context.Context, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientCreateOptions) (PrivateEndpointConnectionsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, privateEndpointConnection, options)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return PrivateEndpointConnectionsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *PrivateEndpointConnectionsClient) createCreateRequest(ctx context.Context, privateEndpointConnection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}/privateEndpointConnections/{privateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	if client.privateEndpointName == "" {
		return nil, errors.New("parameter client.privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(client.privateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, privateEndpointConnection)
}

// createHandleResponse handles the Create response.
func (client *PrivateEndpointConnectionsClient) createHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientCreateResponse, error) {
	result := PrivateEndpointConnectionsClientCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a private endpoint connection for Power BI by private endpoint name.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateEndpointConnectionsClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.BeginDelete
// method.
func (client *PrivateEndpointConnectionsClient) BeginDelete(ctx context.Context, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (PrivateEndpointConnectionsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeletePollerResponse{}, err
	}
	result := PrivateEndpointConnectionsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateEndpointConnectionsClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PrivateEndpointConnectionsClientDeletePollerResponse{}, err
	}
	result.Poller = &PrivateEndpointConnectionsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a private endpoint connection for Power BI by private endpoint name.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateEndpointConnectionsClient) deleteOperation(ctx context.Context, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionsClient) deleteCreateRequest(ctx context.Context, options *PrivateEndpointConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}/privateEndpointConnections/{privateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	if client.privateEndpointName == "" {
		return nil, errors.New("parameter client.privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(client.privateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get a specific private endpoint connection for Power BI by private endpoint name.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
func (client *PrivateEndpointConnectionsClient) Get(ctx context.Context, options *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsClient) getCreateRequest(ctx context.Context, options *PrivateEndpointConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}/privateEndpointConnections/{privateEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.azureResourceName == "" {
		return nil, errors.New("parameter client.azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(client.azureResourceName))
	if client.privateEndpointName == "" {
		return nil, errors.New("parameter client.privateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointName}", url.PathEscape(client.privateEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetResponse, error) {
	result := PrivateEndpointConnectionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResource - Gets private endpoint connection for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription.
// azureResourceName - The name of the powerbi resource.
// options - PrivateEndpointConnectionsClientListByResourceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.ListByResource
// method.
func (client *PrivateEndpointConnectionsClient) ListByResource(resourceGroupName string, azureResourceName string, options *PrivateEndpointConnectionsClientListByResourceOptions) *PrivateEndpointConnectionsClientListByResourcePager {
	return &PrivateEndpointConnectionsClientListByResourcePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceCreateRequest(ctx, resourceGroupName, azureResourceName, options)
		},
		advancer: func(ctx context.Context, resp PrivateEndpointConnectionsClientListByResourceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PrivateEndpointConnectionListResult.NextLink)
		},
	}
}

// listByResourceCreateRequest creates the ListByResource request.
func (client *PrivateEndpointConnectionsClient) listByResourceCreateRequest(ctx context.Context, resourceGroupName string, azureResourceName string, options *PrivateEndpointConnectionsClientListByResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBI/privateLinkServicesForPowerBI/{azureResourceName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureResourceName == "" {
		return nil, errors.New("parameter azureResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureResourceName}", url.PathEscape(azureResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceHandleResponse handles the ListByResource response.
func (client *PrivateEndpointConnectionsClient) listByResourceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByResourceResponse, error) {
	result := PrivateEndpointConnectionsClientListByResourceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionsClientListByResourceResponse{}, err
	}
	return result, nil
}