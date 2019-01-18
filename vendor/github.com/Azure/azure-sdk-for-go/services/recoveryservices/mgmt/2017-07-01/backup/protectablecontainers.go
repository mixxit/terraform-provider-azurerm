package backup

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProtectableContainersClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ProtectableContainersClient struct {
	BaseClient
}

// NewProtectableContainersClient creates an instance of the ProtectableContainersClient client.
func NewProtectableContainersClient(subscriptionID string) ProtectableContainersClient {
	return NewProtectableContainersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectableContainersClientWithBaseURI creates an instance of the ProtectableContainersClient client.
func NewProtectableContainersClientWithBaseURI(baseURI string, subscriptionID string) ProtectableContainersClient {
	return ProtectableContainersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the containers that can be registered to Recovery Services Vault.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the container.
// filter - oData filter options.
func (client ProtectableContainersClient) List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result ProtectableContainerResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectableContainersClient.List")
		defer func() {
			sc := -1
			if result.pcrl.Response.Response != nil {
				sc = result.pcrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, fabricName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pcrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "List", resp, "Failure sending request")
		return
	}

	result.pcrl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProtectableContainersClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":        autorest.Encode("path", fabricName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectableContainers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectableContainersClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectableContainersClient) ListResponder(resp *http.Response) (result ProtectableContainerResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ProtectableContainersClient) listNextResults(ctx context.Context, lastResults ProtectableContainerResourceList) (result ProtectableContainerResourceList, err error) {
	req, err := lastResults.protectableContainerResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ProtectableContainersClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProtectableContainersClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, filter string) (result ProtectableContainerResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProtectableContainersClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, fabricName, filter)
	return
}
