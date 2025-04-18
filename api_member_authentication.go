/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// MemberAuthenticationService MemberAuthentication service
type MemberAuthenticationService service

type MemberAuthenticationAuthSessionLoginPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthSessionReq *UnibeeApiMerchantAuthSessionReq
}

func (r MemberAuthenticationAuthSessionLoginPostRequest) UnibeeApiMerchantAuthSessionReq(unibeeApiMerchantAuthSessionReq UnibeeApiMerchantAuthSessionReq) MemberAuthenticationAuthSessionLoginPostRequest {
	r.unibeeApiMerchantAuthSessionReq = &unibeeApiMerchantAuthSessionReq
	return r
}

func (r MemberAuthenticationAuthSessionLoginPostRequest) Execute() (*MerchantAuthSessionLoginPost200Response, *http.Response, error) {
	return r.ApiService.AuthSessionLoginPostExecute(r)
}

/*
AuthSessionLoginPost Session Login

Session login

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSessionLoginPostRequest
*/
func (a *MemberAuthenticationService) AuthSessionLoginPost(ctx context.Context) MemberAuthenticationAuthSessionLoginPostRequest {
	return MemberAuthenticationAuthSessionLoginPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSessionLoginPost200Response
func (a *MemberAuthenticationService) AuthSessionLoginPostExecute(r MemberAuthenticationAuthSessionLoginPostRequest) (*MerchantAuthSessionLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSessionLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSessionLoginPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/session_login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthSessionReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthSessionReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthSessionReq
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

type MemberAuthenticationAuthSsoLoginOTPPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthLoginOtpReq *UnibeeApiMerchantAuthLoginOtpReq
}

func (r MemberAuthenticationAuthSsoLoginOTPPostRequest) UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq UnibeeApiMerchantAuthLoginOtpReq) MemberAuthenticationAuthSsoLoginOTPPostRequest {
	r.unibeeApiMerchantAuthLoginOtpReq = &unibeeApiMerchantAuthLoginOtpReq
	return r
}

func (r MemberAuthenticationAuthSsoLoginOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginOTPPostExecute(r)
}

/*
AuthSsoLoginOTPPost OTP Login

Send email to member with OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoLoginOTPPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoLoginOTPPost(ctx context.Context) MemberAuthenticationAuthSsoLoginOTPPostRequest {
	return MemberAuthenticationAuthSsoLoginOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MemberAuthenticationService) AuthSsoLoginOTPPostExecute(r MemberAuthenticationAuthSsoLoginOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoLoginOTPPost")
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

type MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthLoginOtpVerifyReq *UnibeeApiMerchantAuthLoginOtpVerifyReq
}

func (r MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest) UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq UnibeeApiMerchantAuthLoginOtpVerifyReq) MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthLoginOtpVerifyReq = &unibeeApiMerchantAuthLoginOtpVerifyReq
	return r
}

func (r MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginOTPVerifyPostExecute(r)
}

/*
AuthSsoLoginOTPVerifyPost OTP Login Code Verification

OTP login for member, verify OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoLoginOTPVerifyPost(ctx context.Context) MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest {
	return MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *MemberAuthenticationService) AuthSsoLoginOTPVerifyPostExecute(r MemberAuthenticationAuthSsoLoginOTPVerifyPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoLoginOTPVerifyPost")
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

type MemberAuthenticationAuthSsoLoginPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthLoginReq *UnibeeApiMerchantAuthLoginReq
}

func (r MemberAuthenticationAuthSsoLoginPostRequest) UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq UnibeeApiMerchantAuthLoginReq) MemberAuthenticationAuthSsoLoginPostRequest {
	r.unibeeApiMerchantAuthLoginReq = &unibeeApiMerchantAuthLoginReq
	return r
}

func (r MemberAuthenticationAuthSsoLoginPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginPostExecute(r)
}

/*
AuthSsoLoginPost Password Login

Password login

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoLoginPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoLoginPost(ctx context.Context) MemberAuthenticationAuthSsoLoginPostRequest {
	return MemberAuthenticationAuthSsoLoginPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *MemberAuthenticationService) AuthSsoLoginPostExecute(r MemberAuthenticationAuthSsoLoginPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoLoginPost")
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

type MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthPasswordForgetOtpReq *UnibeeApiMerchantAuthPasswordForgetOtpReq
}

func (r MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq UnibeeApiMerchantAuthPasswordForgetOtpReq) MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpReq = &unibeeApiMerchantAuthPasswordForgetOtpReq
	return r
}

func (r MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoPasswordForgetOTPPostExecute(r)
}

/*
AuthSsoPasswordForgetOTPPost OTP Password Forget

Send email to member with OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoPasswordForgetOTPPost(ctx context.Context) MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest {
	return MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MemberAuthenticationService) AuthSsoPasswordForgetOTPPostExecute(r MemberAuthenticationAuthSsoPasswordForgetOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoPasswordForgetOTPPost")
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

type MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq
}

func (r MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpVerifyReq = &unibeeApiMerchantAuthPasswordForgetOtpVerifyReq
	return r
}

func (r MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoPasswordForgetOTPVerifyPostExecute(r)
}

/*
AuthSsoPasswordForgetOTPVerifyPost OTP Password Forget Code Verification

Password forget OTP process, verify OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoPasswordForgetOTPVerifyPost(ctx context.Context) MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest {
	return MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MemberAuthenticationService) AuthSsoPasswordForgetOTPVerifyPostExecute(r MemberAuthenticationAuthSsoPasswordForgetOTPVerifyPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoPasswordForgetOTPVerifyPost")
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

type MemberAuthenticationAuthSsoPasswordSetupPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthPasswordSetupOtpReq *UnibeeApiMerchantAuthPasswordSetupOtpReq
}

func (r MemberAuthenticationAuthSsoPasswordSetupPostRequest) UnibeeApiMerchantAuthPasswordSetupOtpReq(unibeeApiMerchantAuthPasswordSetupOtpReq UnibeeApiMerchantAuthPasswordSetupOtpReq) MemberAuthenticationAuthSsoPasswordSetupPostRequest {
	r.unibeeApiMerchantAuthPasswordSetupOtpReq = &unibeeApiMerchantAuthPasswordSetupOtpReq
	return r
}

func (r MemberAuthenticationAuthSsoPasswordSetupPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoPasswordSetupPostExecute(r)
}

/*
AuthSsoPasswordSetupPost Password Setup

Member Password Setup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoPasswordSetupPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoPasswordSetupPost(ctx context.Context) MemberAuthenticationAuthSsoPasswordSetupPostRequest {
	return MemberAuthenticationAuthSsoPasswordSetupPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MemberAuthenticationService) AuthSsoPasswordSetupPostExecute(r MemberAuthenticationAuthSsoPasswordSetupPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoPasswordSetupPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/auth/sso/passwordSetup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantAuthPasswordSetupOtpReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantAuthPasswordSetupOtpReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantAuthPasswordSetupOtpReq
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

type MemberAuthenticationAuthSsoRegisterPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthRegisterReq *UnibeeApiMerchantAuthRegisterReq
}

func (r MemberAuthenticationAuthSsoRegisterPostRequest) UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq UnibeeApiMerchantAuthRegisterReq) MemberAuthenticationAuthSsoRegisterPostRequest {
	r.unibeeApiMerchantAuthRegisterReq = &unibeeApiMerchantAuthRegisterReq
	return r
}

func (r MemberAuthenticationAuthSsoRegisterPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoRegisterPostExecute(r)
}

/*
AuthSsoRegisterPost Register

Register with owner permission, send email with OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoRegisterPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoRegisterPost(ctx context.Context) MemberAuthenticationAuthSsoRegisterPostRequest {
	return MemberAuthenticationAuthSsoRegisterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *MemberAuthenticationService) AuthSsoRegisterPostExecute(r MemberAuthenticationAuthSsoRegisterPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoRegisterPost")
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

type MemberAuthenticationAuthSsoRegisterVerifyPostRequest struct {
	ctx context.Context
	ApiService *MemberAuthenticationService
	unibeeApiMerchantAuthRegisterVerifyReq *UnibeeApiMerchantAuthRegisterVerifyReq
}

func (r MemberAuthenticationAuthSsoRegisterVerifyPostRequest) UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq UnibeeApiMerchantAuthRegisterVerifyReq) MemberAuthenticationAuthSsoRegisterVerifyPostRequest {
	r.unibeeApiMerchantAuthRegisterVerifyReq = &unibeeApiMerchantAuthRegisterVerifyReq
	return r
}

func (r MemberAuthenticationAuthSsoRegisterVerifyPostRequest) Execute() (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoRegisterVerifyPostExecute(r)
}

/*
AuthSsoRegisterVerifyPost Register Verify

Merchant Register, verify OTP code 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MemberAuthenticationAuthSsoRegisterVerifyPostRequest
*/
func (a *MemberAuthenticationService) AuthSsoRegisterVerifyPost(ctx context.Context) MemberAuthenticationAuthSsoRegisterVerifyPostRequest {
	return MemberAuthenticationAuthSsoRegisterVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoRegisterVerifyPost200Response
func (a *MemberAuthenticationService) AuthSsoRegisterVerifyPostExecute(r MemberAuthenticationAuthSsoRegisterVerifyPostRequest) (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoRegisterVerifyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MemberAuthenticationService.AuthSsoRegisterVerifyPost")
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
