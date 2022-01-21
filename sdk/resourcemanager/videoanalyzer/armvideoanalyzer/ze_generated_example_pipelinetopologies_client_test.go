//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer"
)

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-list.json
func ExamplePipelineTopologiesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<account-name>",
		&armvideoanalyzer.PipelineTopologiesClientListOptions{Filter: nil,
			Top: to.Int32Ptr(2),
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-get-by-name.json
func ExamplePipelineTopologiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineTopologiesClientGetResult)
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-create.json
func ExamplePipelineTopologiesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		armvideoanalyzer.PipelineTopology{
			Kind: armvideoanalyzer.Kind("Live").ToPtr(),
			Properties: &armvideoanalyzer.PipelineTopologyProperties{
				Description: to.StringPtr("<description>"),
				Parameters: []*armvideoanalyzer.ParameterDeclaration{
					{
						Name:        to.StringPtr("<name>"),
						Type:        armvideoanalyzer.ParameterType("String").ToPtr(),
						Description: to.StringPtr("<description>"),
						Default:     to.StringPtr("<default>"),
					},
					{
						Name:        to.StringPtr("<name>"),
						Type:        armvideoanalyzer.ParameterType("SecretString").ToPtr(),
						Description: to.StringPtr("<description>"),
						Default:     to.StringPtr("<default>"),
					}},
				Sinks: []armvideoanalyzer.SinkNodeBaseClassification{
					&armvideoanalyzer.VideoSink{
						Name: to.StringPtr("<name>"),
						Type: to.StringPtr("<type>"),
						Inputs: []*armvideoanalyzer.NodeInput{
							{
								NodeName: to.StringPtr("<node-name>"),
							}},
						VideoCreationProperties: &armvideoanalyzer.VideoCreationProperties{
							Description:   to.StringPtr("<description>"),
							SegmentLength: to.StringPtr("<segment-length>"),
							Title:         to.StringPtr("<title>"),
						},
						VideoName: to.StringPtr("<video-name>"),
						VideoPublishingOptions: &armvideoanalyzer.VideoPublishingOptions{
							DisableArchive:        to.StringPtr("<disable-archive>"),
							DisableRtspPublishing: to.StringPtr("<disable-rtsp-publishing>"),
						},
					}},
				Sources: []armvideoanalyzer.SourceNodeBaseClassification{
					&armvideoanalyzer.RtspSource{
						Name: to.StringPtr("<name>"),
						Type: to.StringPtr("<type>"),
						Endpoint: &armvideoanalyzer.UnsecuredEndpoint{
							Type: to.StringPtr("<type>"),
							Credentials: &armvideoanalyzer.UsernamePasswordCredentials{
								Type:     to.StringPtr("<type>"),
								Password: to.StringPtr("<password>"),
								Username: to.StringPtr("<username>"),
							},
							URL: to.StringPtr("<url>"),
						},
						Transport: armvideoanalyzer.RtspTransport("Http").ToPtr(),
					}},
			},
			SKU: &armvideoanalyzer.SKU{
				Name: armvideoanalyzer.SKUName("Live_S1").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineTopologiesClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-delete.json
func ExamplePipelineTopologiesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-patch.json
func ExamplePipelineTopologiesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvideoanalyzer.NewPipelineTopologiesClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<pipeline-topology-name>",
		armvideoanalyzer.PipelineTopologyUpdate{
			Properties: &armvideoanalyzer.PipelineTopologyPropertiesUpdate{
				Description: to.StringPtr("<description>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineTopologiesClientUpdateResult)
}