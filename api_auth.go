/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
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


// AuthService Auth service
type AuthService service

type AuthAuthSsoLoginOTPPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthLoginOtpReq *UnibeeApiMerchantAuthLoginOtpReq
}

func (r AuthAuthSsoLoginOTPPostRequest) UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq UnibeeApiMerchantAuthLoginOtpReq) AuthAuthSsoLoginOTPPostRequest {
	r.unibeeApiMerchantAuthLoginOtpReq = &unibeeApiMerchantAuthLoginOtpReq
	return r
}

func (r AuthAuthSsoLoginOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginOTPPostExecute(r)
}

/*
AuthSsoLoginOTPPost LoginOTP

OTP login for merchant member, send email to member's email address with OTP code'

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoLoginOTPPostRequest
*/
func (a *AuthService) AuthSsoLoginOTPPost(ctx context.Context) AuthAuthSsoLoginOTPPostRequest {
	return AuthAuthSsoLoginOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *AuthService) AuthSsoLoginOTPPostExecute(r AuthAuthSsoLoginOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoLoginOTPPost")
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

type AuthAuthSsoLoginOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthLoginOtpVerifyReq *UnibeeApiMerchantAuthLoginOtpVerifyReq
}

func (r AuthAuthSsoLoginOTPVerifyPostRequest) UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq UnibeeApiMerchantAuthLoginOtpVerifyReq) AuthAuthSsoLoginOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthLoginOtpVerifyReq = &unibeeApiMerchantAuthLoginOtpVerifyReq
	return r
}

func (r AuthAuthSsoLoginOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginOTPVerifyPostExecute(r)
}

/*
AuthSsoLoginOTPVerifyPost LoginOTPVerify

OTP login for merchant member, verify OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoLoginOTPVerifyPostRequest
*/
func (a *AuthService) AuthSsoLoginOTPVerifyPost(ctx context.Context) AuthAuthSsoLoginOTPVerifyPostRequest {
	return AuthAuthSsoLoginOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *AuthService) AuthSsoLoginOTPVerifyPostExecute(r AuthAuthSsoLoginOTPVerifyPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoLoginOTPVerifyPost")
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

type AuthAuthSsoLoginPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthLoginReq *UnibeeApiMerchantAuthLoginReq
}

func (r AuthAuthSsoLoginPostRequest) UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq UnibeeApiMerchantAuthLoginReq) AuthAuthSsoLoginPostRequest {
	r.unibeeApiMerchantAuthLoginReq = &unibeeApiMerchantAuthLoginReq
	return r
}

func (r AuthAuthSsoLoginPostRequest) Execute() (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoLoginPostExecute(r)
}

/*
AuthSsoLoginPost Login

Password login for merchant member'

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoLoginPostRequest
*/
func (a *AuthService) AuthSsoLoginPost(ctx context.Context) AuthAuthSsoLoginPostRequest {
	return AuthAuthSsoLoginPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginPost200Response
func (a *AuthService) AuthSsoLoginPostExecute(r AuthAuthSsoLoginPostRequest) (*MerchantAuthSsoLoginPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoLoginPost")
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

type AuthAuthSsoPasswordForgetOTPPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthPasswordForgetOtpReq *UnibeeApiMerchantAuthPasswordForgetOtpReq
}

func (r AuthAuthSsoPasswordForgetOTPPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq UnibeeApiMerchantAuthPasswordForgetOtpReq) AuthAuthSsoPasswordForgetOTPPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpReq = &unibeeApiMerchantAuthPasswordForgetOtpReq
	return r
}

func (r AuthAuthSsoPasswordForgetOTPPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoPasswordForgetOTPPostExecute(r)
}

/*
AuthSsoPasswordForgetOTPPost PasswordForgetOTP

Merchant member's password forget OTP process,, send email to member's email address with OTP code'

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoPasswordForgetOTPPostRequest
*/
func (a *AuthService) AuthSsoPasswordForgetOTPPost(ctx context.Context) AuthAuthSsoPasswordForgetOTPPostRequest {
	return AuthAuthSsoPasswordForgetOTPPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *AuthService) AuthSsoPasswordForgetOTPPostExecute(r AuthAuthSsoPasswordForgetOTPPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoPasswordForgetOTPPost")
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

type AuthAuthSsoPasswordForgetOTPVerifyPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq *UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq
}

func (r AuthAuthSsoPasswordForgetOTPVerifyPostRequest) UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq) AuthAuthSsoPasswordForgetOTPVerifyPostRequest {
	r.unibeeApiMerchantAuthPasswordForgetOtpVerifyReq = &unibeeApiMerchantAuthPasswordForgetOtpVerifyReq
	return r
}

func (r AuthAuthSsoPasswordForgetOTPVerifyPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoPasswordForgetOTPVerifyPostExecute(r)
}

/*
AuthSsoPasswordForgetOTPVerifyPost PasswordForgetOTPVerify

Merchant member's password forget OTP process, verify OTP code

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoPasswordForgetOTPVerifyPostRequest
*/
func (a *AuthService) AuthSsoPasswordForgetOTPVerifyPost(ctx context.Context) AuthAuthSsoPasswordForgetOTPVerifyPostRequest {
	return AuthAuthSsoPasswordForgetOTPVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *AuthService) AuthSsoPasswordForgetOTPVerifyPostExecute(r AuthAuthSsoPasswordForgetOTPVerifyPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoPasswordForgetOTPVerifyPost")
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

type AuthAuthSsoRegisterPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthRegisterReq *UnibeeApiMerchantAuthRegisterReq
}

func (r AuthAuthSsoRegisterPostRequest) UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq UnibeeApiMerchantAuthRegisterReq) AuthAuthSsoRegisterPostRequest {
	r.unibeeApiMerchantAuthRegisterReq = &unibeeApiMerchantAuthRegisterReq
	return r
}

func (r AuthAuthSsoRegisterPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoRegisterPostExecute(r)
}

/*
AuthSsoRegisterPost Register

Register merchant with owner, send email to owner's email address with OTP code, only open for cloud version; the owner account will create automatic for standalone version

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoRegisterPostRequest
*/
func (a *AuthService) AuthSsoRegisterPost(ctx context.Context) AuthAuthSsoRegisterPostRequest {
	return AuthAuthSsoRegisterPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *AuthService) AuthSsoRegisterPostExecute(r AuthAuthSsoRegisterPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoRegisterPost")
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

type AuthAuthSsoRegisterVerifyPostRequest struct {
	ctx context.Context
	ApiService *AuthService
	unibeeApiMerchantAuthRegisterVerifyReq *UnibeeApiMerchantAuthRegisterVerifyReq
}

func (r AuthAuthSsoRegisterVerifyPostRequest) UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq UnibeeApiMerchantAuthRegisterVerifyReq) AuthAuthSsoRegisterVerifyPostRequest {
	r.unibeeApiMerchantAuthRegisterVerifyReq = &unibeeApiMerchantAuthRegisterVerifyReq
	return r
}

func (r AuthAuthSsoRegisterVerifyPostRequest) Execute() (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	return r.ApiService.AuthSsoRegisterVerifyPostExecute(r)
}

/*
AuthSsoRegisterVerifyPost RegisterVerify

Merchant Register, verify OTP code 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AuthAuthSsoRegisterVerifyPostRequest
*/
func (a *AuthService) AuthSsoRegisterVerifyPost(ctx context.Context) AuthAuthSsoRegisterVerifyPostRequest {
	return AuthAuthSsoRegisterVerifyPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoRegisterVerifyPost200Response
func (a *AuthService) AuthSsoRegisterVerifyPostExecute(r AuthAuthSsoRegisterVerifyPostRequest) (*MerchantAuthSsoRegisterVerifyPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoRegisterVerifyPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthService.AuthSsoRegisterVerifyPost")
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
