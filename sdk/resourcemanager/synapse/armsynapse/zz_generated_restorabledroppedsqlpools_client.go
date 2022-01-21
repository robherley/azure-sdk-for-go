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

// RestorableDroppedSQLPoolsClient contains the methods for the RestorableDroppedSQLPools group.
// Don't use this type directly, use NewRestorableDroppedSQLPoolsClient() instead.
type RestorableDroppedSQLPoolsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableDroppedSQLPoolsClient creates a new instance of RestorableDroppedSQLPoolsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableDroppedSQLPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableDroppedSQLPoolsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RestorableDroppedSQLPoolsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// Get - Gets a deleted sql pool that can be restored
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// restorableDroppedSQLPoolID - The id of the deleted Sql Pool in the form of sqlPoolName,deletionTimeInFileTimeFormat
// options - RestorableDroppedSQLPoolsClientGetOptions contains the optional parameters for the RestorableDroppedSQLPoolsClient.Get
// method.
func (client *RestorableDroppedSQLPoolsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, restorableDroppedSQLPoolID string, options *RestorableDroppedSQLPoolsClientGetOptions) (RestorableDroppedSQLPoolsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, restorableDroppedSQLPoolID, options)
	if err != nil {
		return RestorableDroppedSQLPoolsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableDroppedSQLPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDroppedSQLPoolsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorableDroppedSQLPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, restorableDroppedSQLPoolID string, options *RestorableDroppedSQLPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/restorableDroppedSqlPools/{restorableDroppedSqlPoolId}"
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
	if restorableDroppedSQLPoolID == "" {
		return nil, errors.New("parameter restorableDroppedSQLPoolID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedSqlPoolId}", url.PathEscape(restorableDroppedSQLPoolID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorableDroppedSQLPoolsClient) getHandleResponse(resp *http.Response) (RestorableDroppedSQLPoolsClientGetResponse, error) {
	result := RestorableDroppedSQLPoolsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedSQLPool); err != nil {
		return RestorableDroppedSQLPoolsClientGetResponse{}, err
	}
	return result, nil
}

// ListByWorkspace - Gets a list of deleted Sql pools that can be restored
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - RestorableDroppedSQLPoolsClientListByWorkspaceOptions contains the optional parameters for the RestorableDroppedSQLPoolsClient.ListByWorkspace
// method.
func (client *RestorableDroppedSQLPoolsClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *RestorableDroppedSQLPoolsClientListByWorkspaceOptions) (RestorableDroppedSQLPoolsClientListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return RestorableDroppedSQLPoolsClientListByWorkspaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableDroppedSQLPoolsClientListByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDroppedSQLPoolsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *RestorableDroppedSQLPoolsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *RestorableDroppedSQLPoolsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/restorableDroppedSqlPools"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *RestorableDroppedSQLPoolsClient) listByWorkspaceHandleResponse(resp *http.Response) (RestorableDroppedSQLPoolsClientListByWorkspaceResponse, error) {
	result := RestorableDroppedSQLPoolsClientListByWorkspaceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedSQLPoolListResult); err != nil {
		return RestorableDroppedSQLPoolsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}