//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// IntegrationRuntimeObjectMetadataClient contains the methods for the IntegrationRuntimeObjectMetadata group.
// Don't use this type directly, use NewIntegrationRuntimeObjectMetadataClient() instead.
type IntegrationRuntimeObjectMetadataClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationRuntimeObjectMetadataClient creates a new instance of IntegrationRuntimeObjectMetadataClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationRuntimeObjectMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationRuntimeObjectMetadataClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &IntegrationRuntimeObjectMetadataClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Get object metadata from an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// options - IntegrationRuntimeObjectMetadataClientListOptions contains the optional parameters for the IntegrationRuntimeObjectMetadataClient.List
// method.
func (client *IntegrationRuntimeObjectMetadataClient) List(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientListOptions) (IntegrationRuntimeObjectMetadataClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeObjectMetadataClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *IntegrationRuntimeObjectMetadataClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/getObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.GetMetadataRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.GetMetadataRequest)
	}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationRuntimeObjectMetadataClient) listHandleResponse(resp *http.Response) (IntegrationRuntimeObjectMetadataClientListResponse, error) {
	result := IntegrationRuntimeObjectMetadataClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SsisObjectMetadataListResponse); err != nil {
		return IntegrationRuntimeObjectMetadataClientListResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refresh the object metadata in an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// integrationRuntimeName - Integration runtime name
// options - IntegrationRuntimeObjectMetadataClientBeginRefreshOptions contains the optional parameters for the IntegrationRuntimeObjectMetadataClient.BeginRefresh
// method.
func (client *IntegrationRuntimeObjectMetadataClient) BeginRefresh(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (IntegrationRuntimeObjectMetadataClientRefreshPollerResponse, error) {
	resp, err := client.refresh(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientRefreshPollerResponse{}, err
	}
	result := IntegrationRuntimeObjectMetadataClientRefreshPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationRuntimeObjectMetadataClient.Refresh", "", resp, client.pl)
	if err != nil {
		return IntegrationRuntimeObjectMetadataClientRefreshPollerResponse{}, err
	}
	result.Poller = &IntegrationRuntimeObjectMetadataClientRefreshPoller{
		pt: pt,
	}
	return result, nil
}

// Refresh - Refresh the object metadata in an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
func (client *IntegrationRuntimeObjectMetadataClient) refresh(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *IntegrationRuntimeObjectMetadataClient) refreshCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/refreshObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}