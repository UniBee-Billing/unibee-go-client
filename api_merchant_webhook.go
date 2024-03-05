/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MerchantWebhookAPIService MerchantWebhookAPI service
type MerchantWebhookAPIService service

type ApiMerchantWebhookDeleteEndpointPostRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
	unibeeApiMerchantWebhookDeleteEndpointReq *UnibeeApiMerchantWebhookDeleteEndpointReq
}

func (r ApiMerchantWebhookDeleteEndpointPostRequest) UnibeeApiMerchantWebhookDeleteEndpointReq(unibeeApiMerchantWebhookDeleteEndpointReq UnibeeApiMerchantWebhookDeleteEndpointReq) ApiMerchantWebhookDeleteEndpointPostRequest {
	r.unibeeApiMerchantWebhookDeleteEndpointReq = &unibeeApiMerchantWebhookDeleteEndpointReq
	return r
}

func (r ApiMerchantWebhookDeleteEndpointPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookDeleteEndpointPostExecute(r)
}

/*
MerchantWebhookDeleteEndpointPost Merchant Delete Webhook Endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookDeleteEndpointPostRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookDeleteEndpointPost(ctx context.Context) ApiMerchantWebhookDeleteEndpointPostRequest {
	return ApiMerchantWebhookDeleteEndpointPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantWebhookAPIService) MerchantWebhookDeleteEndpointPostExecute(r ApiMerchantWebhookDeleteEndpointPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookDeleteEndpointPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/delete_endpoint"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantWebhookDeleteEndpointReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantWebhookDeleteEndpointReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantWebhookDeleteEndpointReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMerchantWebhookEndpointListGetRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
}

func (r ApiMerchantWebhookEndpointListGetRequest) Execute() (*MerchantWebhookEndpointListGet200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookEndpointListGetExecute(r)
}

/*
MerchantWebhookEndpointListGet Merchant Webhook Endpoint list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookEndpointListGetRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookEndpointListGet(ctx context.Context) ApiMerchantWebhookEndpointListGetRequest {
	return ApiMerchantWebhookEndpointListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantWebhookEndpointListGet200Response
func (a *MerchantWebhookAPIService) MerchantWebhookEndpointListGetExecute(r ApiMerchantWebhookEndpointListGetRequest) (*MerchantWebhookEndpointListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantWebhookEndpointListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookEndpointListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/endpoint_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMerchantWebhookEndpointLogListGetRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
	endpointId *int64
	page *int32
	count *int32
}

// EndpointId
func (r ApiMerchantWebhookEndpointLogListGetRequest) EndpointId(endpointId int64) ApiMerchantWebhookEndpointLogListGetRequest {
	r.endpointId = &endpointId
	return r
}

// Page, Start WIth 0
func (r ApiMerchantWebhookEndpointLogListGetRequest) Page(page int32) ApiMerchantWebhookEndpointLogListGetRequest {
	r.page = &page
	return r
}

// Count Of Page
func (r ApiMerchantWebhookEndpointLogListGetRequest) Count(count int32) ApiMerchantWebhookEndpointLogListGetRequest {
	r.count = &count
	return r
}

func (r ApiMerchantWebhookEndpointLogListGetRequest) Execute() (*MerchantWebhookEndpointLogListGet200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookEndpointLogListGetExecute(r)
}

/*
MerchantWebhookEndpointLogListGet Merchant Webhook Endpoint Log list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookEndpointLogListGetRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookEndpointLogListGet(ctx context.Context) ApiMerchantWebhookEndpointLogListGetRequest {
	return ApiMerchantWebhookEndpointLogListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantWebhookEndpointLogListGet200Response
func (a *MerchantWebhookAPIService) MerchantWebhookEndpointLogListGetExecute(r ApiMerchantWebhookEndpointLogListGetRequest) (*MerchantWebhookEndpointLogListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantWebhookEndpointLogListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookEndpointLogListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/endpoint_log_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.endpointId == nil {
		return localVarReturnValue, nil, reportError("endpointId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "endpointId", r.endpointId, "")
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMerchantWebhookEventListGetRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
}

func (r ApiMerchantWebhookEventListGetRequest) Execute() (*MerchantWebhookEventListGet200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookEventListGetExecute(r)
}

/*
MerchantWebhookEventListGet Webhook Event list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookEventListGetRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookEventListGet(ctx context.Context) ApiMerchantWebhookEventListGetRequest {
	return ApiMerchantWebhookEventListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantWebhookEventListGet200Response
func (a *MerchantWebhookAPIService) MerchantWebhookEventListGetExecute(r ApiMerchantWebhookEventListGetRequest) (*MerchantWebhookEventListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantWebhookEventListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookEventListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/event_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMerchantWebhookNewEndpointPostRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
	unibeeApiMerchantWebhookNewEndpointReq *UnibeeApiMerchantWebhookNewEndpointReq
}

func (r ApiMerchantWebhookNewEndpointPostRequest) UnibeeApiMerchantWebhookNewEndpointReq(unibeeApiMerchantWebhookNewEndpointReq UnibeeApiMerchantWebhookNewEndpointReq) ApiMerchantWebhookNewEndpointPostRequest {
	r.unibeeApiMerchantWebhookNewEndpointReq = &unibeeApiMerchantWebhookNewEndpointReq
	return r
}

func (r ApiMerchantWebhookNewEndpointPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookNewEndpointPostExecute(r)
}

/*
MerchantWebhookNewEndpointPost Merchant New Webhook Endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookNewEndpointPostRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookNewEndpointPost(ctx context.Context) ApiMerchantWebhookNewEndpointPostRequest {
	return ApiMerchantWebhookNewEndpointPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantWebhookAPIService) MerchantWebhookNewEndpointPostExecute(r ApiMerchantWebhookNewEndpointPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookNewEndpointPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/new_endpoint"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantWebhookNewEndpointReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantWebhookNewEndpointReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantWebhookNewEndpointReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMerchantWebhookUpdateEndpointPostRequest struct {
	ctx context.Context
	ApiService *MerchantWebhookAPIService
	unibeeApiMerchantWebhookUpdateEndpointReq *UnibeeApiMerchantWebhookUpdateEndpointReq
}

func (r ApiMerchantWebhookUpdateEndpointPostRequest) UnibeeApiMerchantWebhookUpdateEndpointReq(unibeeApiMerchantWebhookUpdateEndpointReq UnibeeApiMerchantWebhookUpdateEndpointReq) ApiMerchantWebhookUpdateEndpointPostRequest {
	r.unibeeApiMerchantWebhookUpdateEndpointReq = &unibeeApiMerchantWebhookUpdateEndpointReq
	return r
}

func (r ApiMerchantWebhookUpdateEndpointPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantWebhookUpdateEndpointPostExecute(r)
}

/*
MerchantWebhookUpdateEndpointPost Merchant Update Webhook Endpoint

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantWebhookUpdateEndpointPostRequest
*/
func (a *MerchantWebhookAPIService) MerchantWebhookUpdateEndpointPost(ctx context.Context) ApiMerchantWebhookUpdateEndpointPostRequest {
	return ApiMerchantWebhookUpdateEndpointPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantWebhookAPIService) MerchantWebhookUpdateEndpointPostExecute(r ApiMerchantWebhookUpdateEndpointPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantWebhookAPIService.MerchantWebhookUpdateEndpointPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/webhook/update_endpoint"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantWebhookUpdateEndpointReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantWebhookUpdateEndpointReq is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.unibeeApiMerchantWebhookUpdateEndpointReq
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
