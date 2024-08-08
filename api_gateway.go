/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
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


// GatewayService Gateway service
type GatewayService service

type GatewayGatewayEditCountryConfigPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewayEditCountryConfigReq *UnibeeApiMerchantGatewayEditCountryConfigReq
}

func (r GatewayGatewayEditCountryConfigPostRequest) UnibeeApiMerchantGatewayEditCountryConfigReq(unibeeApiMerchantGatewayEditCountryConfigReq UnibeeApiMerchantGatewayEditCountryConfigReq) GatewayGatewayEditCountryConfigPostRequest {
	r.unibeeApiMerchantGatewayEditCountryConfigReq = &unibeeApiMerchantGatewayEditCountryConfigReq
	return r
}

func (r GatewayGatewayEditCountryConfigPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.GatewayEditCountryConfigPostExecute(r)
}

/*
GatewayEditCountryConfigPost PaymentGatewayCountryConfigEdit

Edit country config for payment gateway, to enable or disable the payment for countryCode, default is enable

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewayEditCountryConfigPostRequest
*/
func (a *GatewayService) GatewayEditCountryConfigPost(ctx context.Context) GatewayGatewayEditCountryConfigPostRequest {
	return GatewayGatewayEditCountryConfigPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *GatewayService) GatewayEditCountryConfigPostExecute(r GatewayGatewayEditCountryConfigPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewayEditCountryConfigPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/edit_country_config"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewayEditCountryConfigReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewayEditCountryConfigReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewayEditCountryConfigReq
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

type GatewayGatewayEditPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewayEditReq *UnibeeApiMerchantGatewayEditReq
}

func (r GatewayGatewayEditPostRequest) UnibeeApiMerchantGatewayEditReq(unibeeApiMerchantGatewayEditReq UnibeeApiMerchantGatewayEditReq) GatewayGatewayEditPostRequest {
	r.unibeeApiMerchantGatewayEditReq = &unibeeApiMerchantGatewayEditReq
	return r
}

func (r GatewayGatewayEditPostRequest) Execute() (*MerchantGatewayEditPost200Response, *http.Response, error) {
	return r.ApiService.GatewayEditPostExecute(r)
}

/*
GatewayEditPost PaymentGatewayEdit

edit the exist payment gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewayEditPostRequest
*/
func (a *GatewayService) GatewayEditPost(ctx context.Context) GatewayGatewayEditPostRequest {
	return GatewayGatewayEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantGatewayEditPost200Response
func (a *GatewayService) GatewayEditPostExecute(r GatewayGatewayEditPostRequest) (*MerchantGatewayEditPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantGatewayEditPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewayEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewayEditReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewayEditReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewayEditReq
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

type GatewayGatewayListGetRequest struct {
	ctx context.Context
	ApiService *GatewayService
}

func (r GatewayGatewayListGetRequest) Execute() (*MerchantGatewayListGet200Response, *http.Response, error) {
	return r.ApiService.GatewayListGetExecute(r)
}

/*
GatewayListGet PaymentGatewayList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewayListGetRequest
*/
func (a *GatewayService) GatewayListGet(ctx context.Context) GatewayGatewayListGetRequest {
	return GatewayGatewayListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantGatewayListGet200Response
func (a *GatewayService) GatewayListGetExecute(r GatewayGatewayListGetRequest) (*MerchantGatewayListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantGatewayListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewayListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/list"

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

type GatewayGatewaySetupPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewaySetupReq *UnibeeApiMerchantGatewaySetupReq
}

func (r GatewayGatewaySetupPostRequest) UnibeeApiMerchantGatewaySetupReq(unibeeApiMerchantGatewaySetupReq UnibeeApiMerchantGatewaySetupReq) GatewayGatewaySetupPostRequest {
	r.unibeeApiMerchantGatewaySetupReq = &unibeeApiMerchantGatewaySetupReq
	return r
}

func (r GatewayGatewaySetupPostRequest) Execute() (*MerchantGatewayEditPost200Response, *http.Response, error) {
	return r.ApiService.GatewaySetupPostExecute(r)
}

/*
GatewaySetupPost PaymentGatewaySetup

Setup the payment gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewaySetupPostRequest
*/
func (a *GatewayService) GatewaySetupPost(ctx context.Context) GatewayGatewaySetupPostRequest {
	return GatewayGatewaySetupPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantGatewayEditPost200Response
func (a *GatewayService) GatewaySetupPostExecute(r GatewayGatewaySetupPostRequest) (*MerchantGatewayEditPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantGatewayEditPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewaySetupPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/setup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewaySetupReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewaySetupReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewaySetupReq
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

type GatewayGatewaySetupWebhookPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewaySetupWebhookReq *UnibeeApiMerchantGatewaySetupWebhookReq
}

func (r GatewayGatewaySetupWebhookPostRequest) UnibeeApiMerchantGatewaySetupWebhookReq(unibeeApiMerchantGatewaySetupWebhookReq UnibeeApiMerchantGatewaySetupWebhookReq) GatewayGatewaySetupWebhookPostRequest {
	r.unibeeApiMerchantGatewaySetupWebhookReq = &unibeeApiMerchantGatewaySetupWebhookReq
	return r
}

func (r GatewayGatewaySetupWebhookPostRequest) Execute() (*MerchantGatewaySetupWebhookPost200Response, *http.Response, error) {
	return r.ApiService.GatewaySetupWebhookPostExecute(r)
}

/*
GatewaySetupWebhookPost PaymentGatewayWebhookSetup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewaySetupWebhookPostRequest
*/
func (a *GatewayService) GatewaySetupWebhookPost(ctx context.Context) GatewayGatewaySetupWebhookPostRequest {
	return GatewayGatewaySetupWebhookPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantGatewaySetupWebhookPost200Response
func (a *GatewayService) GatewaySetupWebhookPostExecute(r GatewayGatewaySetupWebhookPostRequest) (*MerchantGatewaySetupWebhookPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantGatewaySetupWebhookPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewaySetupWebhookPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/setup_webhook"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewaySetupWebhookReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewaySetupWebhookReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewaySetupWebhookReq
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

type GatewayGatewayWireTransferEditPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewayWireTransferEditReq *UnibeeApiMerchantGatewayWireTransferEditReq
}

func (r GatewayGatewayWireTransferEditPostRequest) UnibeeApiMerchantGatewayWireTransferEditReq(unibeeApiMerchantGatewayWireTransferEditReq UnibeeApiMerchantGatewayWireTransferEditReq) GatewayGatewayWireTransferEditPostRequest {
	r.unibeeApiMerchantGatewayWireTransferEditReq = &unibeeApiMerchantGatewayWireTransferEditReq
	return r
}

func (r GatewayGatewayWireTransferEditPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.GatewayWireTransferEditPostExecute(r)
}

/*
GatewayWireTransferEditPost WireTransferEdit

Edit the wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewayWireTransferEditPostRequest
*/
func (a *GatewayService) GatewayWireTransferEditPost(ctx context.Context) GatewayGatewayWireTransferEditPostRequest {
	return GatewayGatewayWireTransferEditPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *GatewayService) GatewayWireTransferEditPostExecute(r GatewayGatewayWireTransferEditPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewayWireTransferEditPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/wire_transfer_edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewayWireTransferEditReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewayWireTransferEditReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewayWireTransferEditReq
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

type GatewayGatewayWireTransferSetupPostRequest struct {
	ctx context.Context
	ApiService *GatewayService
	unibeeApiMerchantGatewayWireTransferSetupReq *UnibeeApiMerchantGatewayWireTransferSetupReq
}

func (r GatewayGatewayWireTransferSetupPostRequest) UnibeeApiMerchantGatewayWireTransferSetupReq(unibeeApiMerchantGatewayWireTransferSetupReq UnibeeApiMerchantGatewayWireTransferSetupReq) GatewayGatewayWireTransferSetupPostRequest {
	r.unibeeApiMerchantGatewayWireTransferSetupReq = &unibeeApiMerchantGatewayWireTransferSetupReq
	return r
}

func (r GatewayGatewayWireTransferSetupPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.GatewayWireTransferSetupPostExecute(r)
}

/*
GatewayWireTransferSetupPost WireTransferSetup

Setup the wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GatewayGatewayWireTransferSetupPostRequest
*/
func (a *GatewayService) GatewayWireTransferSetupPost(ctx context.Context) GatewayGatewayWireTransferSetupPostRequest {
	return GatewayGatewayWireTransferSetupPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *GatewayService) GatewayWireTransferSetupPostExecute(r GatewayGatewayWireTransferSetupPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GatewayService.GatewayWireTransferSetupPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/gateway/wire_transfer_setup"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantGatewayWireTransferSetupReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantGatewayWireTransferSetupReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantGatewayWireTransferSetupReq
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
