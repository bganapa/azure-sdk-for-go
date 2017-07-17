package policy

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

// DefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type DefinitionsClient struct {
    ManagementClient
}
// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
        return NewDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
        }

// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client.
    func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
        return DefinitionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a policy definition.
//
// policyDefinitionName is the name of the policy definition to create. parameters is the policy definition properties.
func (client DefinitionsClient) CreateOrUpdate(policyDefinitionName string, parameters Definition) (result Definition, err error) {
    req, err := client.CreateOrUpdatePreparer(policyDefinitionName, parameters)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", nil , "Failure preparing request")
        return
    }

    resp, err := client.CreateOrUpdateSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
    }

    result, err = client.CreateOrUpdateResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
    }

    return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DefinitionsClient) CreateOrUpdatePreparer(policyDefinitionName string, parameters Definition) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "policyDefinitionName": autorest.Encode("path",policyDefinitionName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2015-10-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPut(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}",pathParameters),
                        autorest.WithJSON(parameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result Definition, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Delete deletes a policy definition.
//
// policyDefinitionName is the name of the policy definition to delete.
func (client DefinitionsClient) Delete(policyDefinitionName string) (result autorest.Response, err error) {
    req, err := client.DeletePreparer(policyDefinitionName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", nil , "Failure preparing request")
        return
    }

    resp, err := client.DeleteSender(req)
    if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure sending request")
        return
    }

    result, err = client.DeleteResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure responding to request")
    }

    return
}

// DeletePreparer prepares the Delete request.
func (client DefinitionsClient) DeletePreparer(policyDefinitionName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "policyDefinitionName": autorest.Encode("path",policyDefinitionName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2015-10-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsDelete(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusNoContent,http.StatusOK),
            autorest.ByClosing())
    result.Response = resp
    return
}

// Get gets the policy definition.
//
// policyDefinitionName is the name of the policy definition to get.
func (client DefinitionsClient) Get(policyDefinitionName string) (result Definition, err error) {
    req, err := client.GetPreparer(policyDefinitionName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(policyDefinitionName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "policyDefinitionName": autorest.Encode("path",policyDefinitionName),
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2015-10-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// List gets all the policy definitions for a subscription.
//
// filter is the filter to apply on the operation.
func (client DefinitionsClient) List(filter string) (result DefinitionListResult, err error) {
    req, err := client.ListPreparer(filter)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil , "Failure preparing request")
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending request")
        return
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to request")
    }

    return
}

// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer(filter string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "subscriptionId": autorest.Encode("path",client.SubscriptionID),
    }

        const APIVersion = "2015-10-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if len(filter) > 0 {
        queryParameters["$filter"] = autorest.Encode("query",filter)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionListResult, err error) {
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
func (client DefinitionsClient) ListNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
    req, err := lastResults.DefinitionListResultPreparer()
    if err != nil {
        return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil , "Failure preparing next results request")
    }
    if req == nil {
        return
    }

    resp, err := client.ListSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending next results request")
    }

    result, err = client.ListResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to next results request")
    }

    return
}

// ListComplete gets all elements from the list without paging.
func (client DefinitionsClient) ListComplete (filter string, cancel <-chan struct{}) (<-chan Definition, <-chan error) {
    resultChan := make(chan Definition)
    errChan := make(chan error, 1)
    go func() {
        defer func() {
			close(resultChan)
			close(errChan)
		}()
        list, err := client.List(filter)
        if err != nil {
            errChan <- err
            return
        }
        if list.Value != nil {
            for _, item := range *list.Value {
                select {
                case <- cancel:
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
                            case <- cancel:
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

