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
// Code generated by Microsoft (R) AutoRest Code Generator 1.2.0.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualNetworkGatewayConnectionsClient is the network Client
type VirtualNetworkGatewayConnectionsClient struct {
	ManagementClient
}

// NewVirtualNetworkGatewayConnectionsClient creates an instance of the VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClient(subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return NewVirtualNetworkGatewayConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkGatewayConnectionsClientWithBaseURI creates an instance of the
// VirtualNetworkGatewayConnectionsClient client.
func NewVirtualNetworkGatewayConnectionsClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkGatewayConnectionsClient {
	return VirtualNetworkGatewayConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a virtual network gateway connection in the specified resource group. This method
// may poll for completion. Polling can be canceled by passing the cancel channel argument. The channel will be used to
// cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection. parameters is parameters supplied to the create or update virtual network gateway
// connection operation.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdate(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, cancel <-chan struct{}) (<-chan VirtualNetworkGatewayConnection, <-chan error) {
	resultChan := make(chan VirtualNetworkGatewayConnection, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualNetworkGatewayConnection
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters VirtualNetworkGatewayConnection, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": autorest.Encode("path", virtualNetworkGatewayConnectionName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetworkGatewayConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified virtual network Gateway connection. This method may poll for completion. Polling can be
// canceled by passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP
// requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Delete(resourceGroupName string, virtualNetworkGatewayConnectionName string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, virtualNetworkGatewayConnectionName, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworkGatewayConnectionsClient) DeletePreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": autorest.Encode("path", virtualNetworkGatewayConnectionName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified virtual network gateway connection by resource group.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the name of the virtual
// network gateway connection.
func (client VirtualNetworkGatewayConnectionsClient) Get(resourceGroupName string, virtualNetworkGatewayConnectionName string) (result VirtualNetworkGatewayConnection, err error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkGatewayConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworkGatewayConnectionsClient) GetPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": autorest.Encode("path", virtualNetworkGatewayConnectionName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetResponder(resp *http.Response) (result VirtualNetworkGatewayConnection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSharedKey the Get VirtualNetworkGatewayConnectionSharedKey operation retrieves information about the specified
// virtual network gateway connection shared key through Network resource provider.
//
// resourceGroupName is the name of the resource group. connectionSharedKeyName is the virtual network gateway
// connection shared key name.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKey(resourceGroupName string, connectionSharedKeyName string) (result ConnectionSharedKeyResult, err error) {
	req, err := client.GetSharedKeyPreparer(resourceGroupName, connectionSharedKeyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSharedKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure sending request")
		return
	}

	result, err = client.GetSharedKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "GetSharedKey", resp, "Failure responding to request")
	}

	return
}

// GetSharedKeyPreparer prepares the GetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyPreparer(resourceGroupName string, connectionSharedKeyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"connectionSharedKeyName": autorest.Encode("path", connectionSharedKeyName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{connectionSharedKeyName}/sharedkey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSharedKeySender sends the GetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetSharedKeyResponder handles the response to the GetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) GetSharedKeyResponder(resp *http.Response) (result ConnectionSharedKeyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the List VirtualNetworkGatewayConnections operation retrieves all the virtual network gateways connections
// created.
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworkGatewayConnectionsClient) List(resourceGroupName string) (result VirtualNetworkGatewayConnectionListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkGatewayConnectionsClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ListResponder(resp *http.Response) (result VirtualNetworkGatewayConnectionListResult, err error) {
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
func (client VirtualNetworkGatewayConnectionsClient) ListNextResults(lastResults VirtualNetworkGatewayConnectionListResult) (result VirtualNetworkGatewayConnectionListResult, err error) {
	req, err := lastResults.VirtualNetworkGatewayConnectionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client VirtualNetworkGatewayConnectionsClient) ListComplete(resourceGroupName string, cancel <-chan struct{}) (<-chan VirtualNetworkGatewayConnection, <-chan error) {
	resultChan := make(chan VirtualNetworkGatewayConnection)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceGroupName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ResetSharedKey the VirtualNetworkGatewayConnectionResetSharedKey operation resets the virtual network gateway
// connection shared key for passed virtual network gateway connection in the specified resource group through Network
// resource provider. This method may poll for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the virtual network
// gateway connection reset shared key Name. parameters is parameters supplied to the begin reset virtual network
// gateway connection shared key operation through network resource provider.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKey(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, cancel <-chan struct{}) (<-chan ConnectionResetSharedKey, <-chan error) {
	resultChan := make(chan ConnectionResetSharedKey, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ConnectionResetSharedKey
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.ResetSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", nil, "Failure preparing request")
			return
		}

		resp, err := client.ResetSharedKeySender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", resp, "Failure sending request")
			return
		}

		result, err = client.ResetSharedKeyResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "ResetSharedKey", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// ResetSharedKeyPreparer prepares the ResetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionResetSharedKey, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": autorest.Encode("path", virtualNetworkGatewayConnectionName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey/reset", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// ResetSharedKeySender sends the ResetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// ResetSharedKeyResponder handles the response to the ResetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) ResetSharedKeyResponder(resp *http.Response) (result ConnectionResetSharedKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// SetSharedKey the Put VirtualNetworkGatewayConnectionSharedKey operation sets the virtual network gateway connection
// shared key for passed virtual network gateway connection in the specified resource group through Network resource
// provider. This method may poll for completion. Polling can be canceled by passing the cancel channel argument. The
// channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkGatewayConnectionName is the virtual network
// gateway connection name. parameters is parameters supplied to the Begin Set Virtual Network Gateway conection Shared
// key operation throughNetwork resource provider.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKey(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, cancel <-chan struct{}) (<-chan ConnectionSharedKey, <-chan error) {
	resultChan := make(chan ConnectionSharedKey, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result ConnectionSharedKey
		defer func() {
			if err != nil {
				errChan <- err
			}
			resultChan <- result
			close(resultChan)
			close(errChan)
		}()
		req, err := client.SetSharedKeyPreparer(resourceGroupName, virtualNetworkGatewayConnectionName, parameters, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "SetSharedKey", nil, "Failure preparing request")
			return
		}

		resp, err := client.SetSharedKeySender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "SetSharedKey", resp, "Failure sending request")
			return
		}

		result, err = client.SetSharedKeyResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "network.VirtualNetworkGatewayConnectionsClient", "SetSharedKey", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// SetSharedKeyPreparer prepares the SetSharedKey request.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyPreparer(resourceGroupName string, virtualNetworkGatewayConnectionName string, parameters ConnectionSharedKey, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"virtualNetworkGatewayConnectionName": autorest.Encode("path", virtualNetworkGatewayConnectionName),
	}

	const APIVersion = "2015-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections/{virtualNetworkGatewayConnectionName}/sharedkey", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// SetSharedKeySender sends the SetSharedKey request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// SetSharedKeyResponder handles the response to the SetSharedKey request. The method always
// closes the http.Response Body.
func (client VirtualNetworkGatewayConnectionsClient) SetSharedKeyResponder(resp *http.Response) (result ConnectionSharedKey, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
