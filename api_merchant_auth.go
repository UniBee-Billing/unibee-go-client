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


// MerchantAuthAPIService MerchantAuthAPI service
type MerchantAuthAPIService service

type MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthLoginOtpReq *UnibeeApiMerchantAuthLoginOtpReq
}

func (r MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest) UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq UnibeeApiMerchantAuthLoginOtpReq) MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest {
	r.unibeeApiMerchantAuthLoginOtpReq = &unibeeApiMerchantAuthLoginOtpReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoLoginOTPPostExecute(r)
}

/*
MerchantAuthSsoLoginOTPPost Login OTP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginOTPPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest {
	return MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginOTPPostExecute(r MerchantAuthAPIMerchantAuthSsoLoginOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoLoginOTPPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/loginOTP"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthLoginOtpReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthLoginOtpReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthLoginOtpReq
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

type MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthLoginOtpVerifyReq *UnibeeApiMerchantAuthLoginOtpVerifyReq
}

func (r MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest) UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq UnibeeApiMerchantAuthLoginOtpVerifyReq) MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthLoginOtpVerifyReq = &unibeeApiMerchantAuthLoginOtpVerifyReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoLoginOTPVerifyPostExecute(r)
}

/*
MerchantAuthSsoLoginOTPVerifyPost Merchant User OTP Login Verify

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginOTPVerifyPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest {
	return MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginOTPVerifyPostExecute(r MerchantAuthAPIMerchantAuthSsoLoginOTPVerifyPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoLoginOTPVerifyPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/loginOTPVerify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthLoginOtpVerifyReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthLoginOtpVerifyReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthLoginOtpVerifyReq
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

type MerchantAuthAPIMerchantAuthSsoLoginPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthLoginReq *UnibeeApiMerchantAuthLoginReq
}

func (r MerchantAuthAPIMerchantAuthSsoLoginPostRequest) UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq UnibeeApiMerchantAuthLoginReq) MerchantAuthAPIMerchantAuthSsoLoginPostRequest {
	r.unibeeApiMerchantAuthLoginReq = &unibeeApiMerchantAuthLoginReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoLoginPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoLoginPostExecute(r)
}

/*
MerchantAuthSsoLoginPost Login

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoLoginPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoLoginPostRequest {
	return MerchantAuthAPIMerchantAuthSsoLoginPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoLoginPostExecute(r MerchantAuthAPIMerchantAuthSsoLoginPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoLoginPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthLoginReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthLoginReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthLoginReq
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

type MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthPasswordForgetOtpReq *UnibeeApiMerchantAuthPasswordForgetOtpReq
}

func (r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq UnibeeApiMerchantAuthPasswordForgetOtpReq) MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpReq = &unibeeApiMerchantAuthPasswordForgetOtpReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoPasswordForgetOTPPostExecute(r)
}

/*
MerchantAuthSsoPasswordForgetOTPPost Merchant Password Forget OTP

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoPasswordForgetOTPPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest {
	return MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoPasswordForgetOTPPostExecute(r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoPasswordForgetOTPPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/passwordForgetOTP"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthPasswordForgetOtpReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthPasswordForgetOtpReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthPasswordForgetOtpReq
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

type MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq
}

func (r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpVerifyReq = &unibeeApiMerchantAuthPasswordForgetOtpVerifyReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoPasswordForgetOTPVerifyPostExecute(r)
}

/*
MerchantAuthSsoPasswordForgetOTPVerifyPost Merchant Password Forget OTP Verify

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoPasswordForgetOTPVerifyPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest {
	return MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoPasswordForgetOTPVerifyPostExecute(r MerchantAuthAPIMerchantAuthSsoPasswordForgetOTPVerifyPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoPasswordForgetOTPVerifyPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/passwordForgetOTPVerify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthPasswordForgetOtpVerifyReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthPasswordForgetOtpVerifyReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthPasswordForgetOtpVerifyReq
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

type MerchantAuthAPIMerchantAuthSsoRegisterPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthRegisterReq *UnibeeApiMerchantAuthRegisterReq
}

func (r MerchantAuthAPIMerchantAuthSsoRegisterPostRequest) UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq UnibeeApiMerchantAuthRegisterReq) MerchantAuthAPIMerchantAuthSsoRegisterPostRequest {
	r.unibeeApiMerchantAuthRegisterReq = &unibeeApiMerchantAuthRegisterReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoRegisterPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoRegisterPostExecute(r)
}

/*
MerchantAuthSsoRegisterPost Merchant Register

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoRegisterPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoRegisterPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoRegisterPostRequest {
	return MerchantAuthAPIMerchantAuthSsoRegisterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoRegisterPostExecute(r MerchantAuthAPIMerchantAuthSsoRegisterPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoRegisterPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthRegisterReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthRegisterReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthRegisterReq
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

type MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest struct {
	ctx context.Context
	ApiService *MerchantAuthAPIService
	unibeeApiMerchantAuthRegisterVerifyReq *UnibeeApiMerchantAuthRegisterVerifyReq
}

func (r MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest) UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq UnibeeApiMerchantAuthRegisterVerifyReq) MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest {
	r.unibeeApiMerchantAuthRegisterVerifyReq = &unibeeApiMerchantAuthRegisterVerifyReq
	return r
}

func (r MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest) Execute() (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	return r.ApiService.MerchantAuthSsoRegisterVerifyPostExecute(r)
}

/*
MerchantAuthSsoRegisterVerifyPost Merchant Register Verify

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest
*/
func (a *MerchantAuthAPIService) MerchantAuthSsoRegisterVerifyPost(ctx context.Context) MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest {
	return MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoRegisterVerifyPost200Response
func (a *MerchantAuthAPIService) MerchantAuthSsoRegisterVerifyPostExecute(r MerchantAuthAPIMerchantAuthSsoRegisterVerifyPostRequest) (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoRegisterVerifyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerchantAuthAPIService.MerchantAuthSsoRegisterVerifyPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/registerVerify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthRegisterVerifyReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthRegisterVerifyReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthRegisterVerifyReq
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
