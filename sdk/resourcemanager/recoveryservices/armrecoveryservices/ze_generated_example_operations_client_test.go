//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/ListOperations.json
func ExampleOperationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewOperationsClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/GetOperationStatus.json
func ExampleOperationsClient_OperationStatusGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewOperationsClient("<subscription-id>", cred, nil)
	res, err := client.OperationStatusGet(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationsClientOperationStatusGetResult)
}

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/GetOperationResult.json
func ExampleOperationsClient_GetOperationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewOperationsClient("<subscription-id>", cred, nil)
	res, err := client.GetOperationResult(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.OperationsClientGetOperationResultResult)
}