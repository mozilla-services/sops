package web

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

// ProviderClient is the webSite Management Client
type ProviderClient struct {
	BaseClient
}

// NewProviderClient creates an instance of the ProviderClient client.
func NewProviderClient(subscriptionID string) ProviderClient {
	return NewProviderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProviderClientWithBaseURI creates an instance of the ProviderClient client.
func NewProviderClientWithBaseURI(baseURI string, subscriptionID string) ProviderClient {
	return ProviderClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetPublishingUser sends the get publishing user request.
func (client ProviderClient) GetPublishingUser(ctx context.Context) (result User, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetPublishingUser")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPublishingUserPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", resp, "Failure sending request")
		return
	}

	result, err = client.GetPublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetPublishingUser", resp, "Failure responding to request")
	}

	return
}

// GetPublishingUserPreparer prepares the GetPublishingUser request.
func (client ProviderClient) GetPublishingUserPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPublishingUserSender sends the GetPublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetPublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetPublishingUserResponder handles the response to the GetPublishingUser request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetPublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSourceControl sends the get source control request.
// Parameters:
// sourceControlType - type of source control
func (client ProviderClient) GetSourceControl(ctx context.Context, sourceControlType string) (result SourceControl, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetSourceControl")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSourceControlPreparer(ctx, sourceControlType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSourceControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", resp, "Failure sending request")
		return
	}

	result, err = client.GetSourceControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControl", resp, "Failure responding to request")
	}

	return
}

// GetSourceControlPreparer prepares the GetSourceControl request.
func (client ProviderClient) GetSourceControlPreparer(ctx context.Context, sourceControlType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceControlType": autorest.Encode("path", sourceControlType),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSourceControlSender sends the GetSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetSourceControlSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSourceControlResponder handles the response to the GetSourceControl request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetSourceControlResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSourceControls sends the get source controls request.
func (client ProviderClient) GetSourceControls(ctx context.Context) (result SourceControlCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetSourceControls")
		defer func() {
			sc := -1
			if result.scc.Response.Response != nil {
				sc = result.scc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getSourceControlsNextResults
	req, err := client.GetSourceControlsPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSourceControlsSender(req)
	if err != nil {
		result.scc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", resp, "Failure sending request")
		return
	}

	result.scc, err = client.GetSourceControlsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "GetSourceControls", resp, "Failure responding to request")
	}

	return
}

// GetSourceControlsPreparer prepares the GetSourceControls request.
func (client ProviderClient) GetSourceControlsPreparer(ctx context.Context) (*http.Request, error) {
	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/sourcecontrols"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSourceControlsSender sends the GetSourceControls request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) GetSourceControlsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSourceControlsResponder handles the response to the GetSourceControls request. The method always
// closes the http.Response Body.
func (client ProviderClient) GetSourceControlsResponder(resp *http.Response) (result SourceControlCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getSourceControlsNextResults retrieves the next set of results, if any.
func (client ProviderClient) getSourceControlsNextResults(ctx context.Context, lastResults SourceControlCollection) (result SourceControlCollection, err error) {
	req, err := lastResults.sourceControlCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getSourceControlsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSourceControlsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.ProviderClient", "getSourceControlsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetSourceControlsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "getSourceControlsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetSourceControlsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProviderClient) GetSourceControlsComplete(ctx context.Context) (result SourceControlCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.GetSourceControls")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetSourceControls(ctx)
	return
}

// UpdatePublishingUser sends the update publishing user request.
// Parameters:
// requestMessage - details of publishing user
func (client ProviderClient) UpdatePublishingUser(ctx context.Context, requestMessage User) (result User, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.UpdatePublishingUser")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePublishingUserPreparer(ctx, requestMessage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePublishingUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePublishingUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdatePublishingUser", resp, "Failure responding to request")
	}

	return
}

// UpdatePublishingUserPreparer prepares the UpdatePublishingUser request.
func (client ProviderClient) UpdatePublishingUserPreparer(ctx context.Context, requestMessage User) (*http.Request, error) {
	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Web/publishingUsers/web"),
		autorest.WithJSON(requestMessage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePublishingUserSender sends the UpdatePublishingUser request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) UpdatePublishingUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdatePublishingUserResponder handles the response to the UpdatePublishingUser request. The method always
// closes the http.Response Body.
func (client ProviderClient) UpdatePublishingUserResponder(resp *http.Response) (result User, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateSourceControl sends the update source control request.
// Parameters:
// sourceControlType - type of source control
// requestMessage - source control token information
func (client ProviderClient) UpdateSourceControl(ctx context.Context, sourceControlType string, requestMessage SourceControl) (result SourceControl, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProviderClient.UpdateSourceControl")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateSourceControlPreparer(ctx, sourceControlType, requestMessage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSourceControlSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateSourceControlResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.ProviderClient", "UpdateSourceControl", resp, "Failure responding to request")
	}

	return
}

// UpdateSourceControlPreparer prepares the UpdateSourceControl request.
func (client ProviderClient) UpdateSourceControlPreparer(ctx context.Context, sourceControlType string, requestMessage SourceControl) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"sourceControlType": autorest.Encode("path", sourceControlType),
	}

	const APIVersion = "2015-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Web/sourcecontrols/{sourceControlType}", pathParameters),
		autorest.WithJSON(requestMessage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSourceControlSender sends the UpdateSourceControl request. The method will close the
// http.Response Body if it receives an error.
func (client ProviderClient) UpdateSourceControlSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateSourceControlResponder handles the response to the UpdateSourceControl request. The method always
// closes the http.Response Body.
func (client ProviderClient) UpdateSourceControlResponder(resp *http.Response) (result SourceControl, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
