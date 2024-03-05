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


// MerchantBalanceAPIService MerchantBalanceAPI service
type MerchantBalanceAPIService service

type ApiMerchantBalanceMerchantBalanceQueryGetRequest struct {
	ctx context.Context
	ApiService *MerchantBalanceAPIService
	gatewayId *int64
}

// gatewayId
func (r ApiMerchantBalanceMerchantBalanceQueryGetRequest) GatewayId(gatewayId int64) ApiMerchantBalanceMerchantBalanceQueryGetRequest {
	r.gatewayId = &gatewayId
	return r
}

func (r ApiMerchantBalanceMerchantBalanceQueryGetRequest) Execute() (*MerchantBalanceMerchantBalanceQueryGet200Response, *http.Response, error) {
	return r.ApiService.MerchantBalanceMerchantBalanceQueryGetExecute(r)
}

/*
MerchantBalanceMerchantBalanceQueryGet Query Merchant Gateway Balance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantBalanceMerchantBalanceQueryGetRequest
*/
func (a *MerchantBalanceAPIService) MerchantBalanceMerchantBalanceQueryGet(ctx context.Context) ApiMerchantBalanceMerchantBalanceQueryGetRequest {
	return ApiMerchantBalanceMerchantBalanceQueryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantBalanceMerchantBalanceQueryGet200Response
func (a *MerchantBalanceAPIService) MerchantBalanceMerchantBalanceQueryGetExecute(r ApiMerchantBalanceMerchantBalanceQueryGetRequest) (*MerchantBalanceMerchantBalanceQueryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantBalanceMerchantBalanceQueryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantBalanceAPIService.MerchantBalanceMerchantBalanceQueryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/balance/merchant_balance_query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gatewayId == nil {
		return localVarReturnValue, nil, reportError("gatewayId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayId", r.gatewayId, "")
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

type ApiMerchantBalanceMerchantBalanceQueryPostRequest struct {
	ctx context.Context
	ApiService *MerchantBalanceAPIService
	unibeeApiMerchantBalanceDetailQueryReq *UnibeeApiMerchantBalanceDetailQueryReq
}

func (r ApiMerchantBalanceMerchantBalanceQueryPostRequest) UnibeeApiMerchantBalanceDetailQueryReq(unibeeApiMerchantBalanceDetailQueryReq UnibeeApiMerchantBalanceDetailQueryReq) ApiMerchantBalanceMerchantBalanceQueryPostRequest {
	r.unibeeApiMerchantBalanceDetailQueryReq = &unibeeApiMerchantBalanceDetailQueryReq
	return r
}

func (r ApiMerchantBalanceMerchantBalanceQueryPostRequest) Execute() (*MerchantBalanceMerchantBalanceQueryGet200Response, *http.Response, error) {
	return r.ApiService.MerchantBalanceMerchantBalanceQueryPostExecute(r)
}

/*
MerchantBalanceMerchantBalanceQueryPost Query Merchant Gateway Balance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantBalanceMerchantBalanceQueryPostRequest
*/
func (a *MerchantBalanceAPIService) MerchantBalanceMerchantBalanceQueryPost(ctx context.Context) ApiMerchantBalanceMerchantBalanceQueryPostRequest {
	return ApiMerchantBalanceMerchantBalanceQueryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantBalanceMerchantBalanceQueryGet200Response
func (a *MerchantBalanceAPIService) MerchantBalanceMerchantBalanceQueryPostExecute(r ApiMerchantBalanceMerchantBalanceQueryPostRequest) (*MerchantBalanceMerchantBalanceQueryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantBalanceMerchantBalanceQueryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantBalanceAPIService.MerchantBalanceMerchantBalanceQueryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/balance/merchant_balance_query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantBalanceDetailQueryReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantBalanceDetailQueryReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantBalanceDetailQueryReq
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

type ApiMerchantBalanceUserBalanceQueryGetRequest struct {
	ctx context.Context
	ApiService *MerchantBalanceAPIService
	userId *int64
	gatewayId *int64
}

// userId
func (r ApiMerchantBalanceUserBalanceQueryGetRequest) UserId(userId int64) ApiMerchantBalanceUserBalanceQueryGetRequest {
	r.userId = &userId
	return r
}

// gatewayId
func (r ApiMerchantBalanceUserBalanceQueryGetRequest) GatewayId(gatewayId int64) ApiMerchantBalanceUserBalanceQueryGetRequest {
	r.gatewayId = &gatewayId
	return r
}

func (r ApiMerchantBalanceUserBalanceQueryGetRequest) Execute() (*MerchantBalanceUserBalanceQueryGet200Response, *http.Response, error) {
	return r.ApiService.MerchantBalanceUserBalanceQueryGetExecute(r)
}

/*
MerchantBalanceUserBalanceQueryGet Query User Balance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantBalanceUserBalanceQueryGetRequest
*/
func (a *MerchantBalanceAPIService) MerchantBalanceUserBalanceQueryGet(ctx context.Context) ApiMerchantBalanceUserBalanceQueryGetRequest {
	return ApiMerchantBalanceUserBalanceQueryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantBalanceUserBalanceQueryGet200Response
func (a *MerchantBalanceAPIService) MerchantBalanceUserBalanceQueryGetExecute(r ApiMerchantBalanceUserBalanceQueryGetRequest) (*MerchantBalanceUserBalanceQueryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantBalanceUserBalanceQueryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantBalanceAPIService.MerchantBalanceUserBalanceQueryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/balance/user_balance_query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userId == nil {
		return localVarReturnValue, nil, reportError("userId is required and must be specified")
	}
	if r.gatewayId == nil {
		return localVarReturnValue, nil, reportError("gatewayId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "gatewayId", r.gatewayId, "")
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

type ApiMerchantBalanceUserBalanceQueryPostRequest struct {
	ctx context.Context
	ApiService *MerchantBalanceAPIService
	unibeeApiMerchantBalanceUserDetailQueryReq *UnibeeApiMerchantBalanceUserDetailQueryReq
}

func (r ApiMerchantBalanceUserBalanceQueryPostRequest) UnibeeApiMerchantBalanceUserDetailQueryReq(unibeeApiMerchantBalanceUserDetailQueryReq UnibeeApiMerchantBalanceUserDetailQueryReq) ApiMerchantBalanceUserBalanceQueryPostRequest {
	r.unibeeApiMerchantBalanceUserDetailQueryReq = &unibeeApiMerchantBalanceUserDetailQueryReq
	return r
}

func (r ApiMerchantBalanceUserBalanceQueryPostRequest) Execute() (*MerchantBalanceUserBalanceQueryGet200Response, *http.Response, error) {
	return r.ApiService.MerchantBalanceUserBalanceQueryPostExecute(r)
}

/*
MerchantBalanceUserBalanceQueryPost Query User Balance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMerchantBalanceUserBalanceQueryPostRequest
*/
func (a *MerchantBalanceAPIService) MerchantBalanceUserBalanceQueryPost(ctx context.Context) ApiMerchantBalanceUserBalanceQueryPostRequest {
	return ApiMerchantBalanceUserBalanceQueryPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantBalanceUserBalanceQueryGet200Response
func (a *MerchantBalanceAPIService) MerchantBalanceUserBalanceQueryPostExecute(r ApiMerchantBalanceUserBalanceQueryPostRequest) (*MerchantBalanceUserBalanceQueryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantBalanceUserBalanceQueryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantBalanceAPIService.MerchantBalanceUserBalanceQueryPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/balance/user_balance_query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantBalanceUserDetailQueryReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantBalanceUserDetailQueryReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantBalanceUserDetailQueryReq
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
