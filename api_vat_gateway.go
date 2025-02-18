/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
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


// VatGatewayService VatGateway service
type VatGatewayService service

type VatGatewayVatCountryListGetRequest struct {
	ctx context.Context
	ApiService *VatGatewayService
}

func (r VatGatewayVatCountryListGetRequest) Execute() (*MerchantVatCountryListGet200Response, *http.Response, error) {
	return r.ApiService.VatCountryListGetExecute(r)
}

/*
VatCountryListGet Get Vat Country List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return VatGatewayVatCountryListGetRequest
*/
func (a *VatGatewayService) VatCountryListGet(ctx context.Context) VatGatewayVatCountryListGetRequest {
	return VatGatewayVatCountryListGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantVatCountryListGet200Response
func (a *VatGatewayService) VatCountryListGetExecute(r VatGatewayVatCountryListGetRequest) (*MerchantVatCountryListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantVatCountryListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VatGatewayService.VatCountryListGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/vat/country_list"

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

type VatGatewayVatCountryListPostRequest struct {
	ctx context.Context
	ApiService *VatGatewayService
	body *map[string]interface{}
}

func (r VatGatewayVatCountryListPostRequest) Body(body map[string]interface{}) VatGatewayVatCountryListPostRequest {
	r.body = &body
	return r
}

func (r VatGatewayVatCountryListPostRequest) Execute() (*MerchantVatCountryListGet200Response, *http.Response, error) {
	return r.ApiService.VatCountryListPostExecute(r)
}

/*
VatCountryListPost Get Vat Country List

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return VatGatewayVatCountryListPostRequest
*/
func (a *VatGatewayService) VatCountryListPost(ctx context.Context) VatGatewayVatCountryListPostRequest {
	return VatGatewayVatCountryListPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantVatCountryListGet200Response
func (a *VatGatewayService) VatCountryListPostExecute(r VatGatewayVatCountryListPostRequest) (*MerchantVatCountryListGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantVatCountryListGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VatGatewayService.VatCountryListPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/vat/country_list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type VatGatewayVatInitDefaultGatewayPostRequest struct {
	ctx context.Context
	ApiService *VatGatewayService
	body *map[string]interface{}
}

func (r VatGatewayVatInitDefaultGatewayPostRequest) Body(body map[string]interface{}) VatGatewayVatInitDefaultGatewayPostRequest {
	r.body = &body
	return r
}

func (r VatGatewayVatInitDefaultGatewayPostRequest) Execute() (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	return r.ApiService.VatInitDefaultGatewayPostExecute(r)
}

/*
VatInitDefaultGatewayPost Init Default Vat Gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return VatGatewayVatInitDefaultGatewayPostRequest
*/
func (a *VatGatewayService) VatInitDefaultGatewayPost(ctx context.Context) VatGatewayVatInitDefaultGatewayPostRequest {
	return VatGatewayVatInitDefaultGatewayPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantAuthSsoLoginOTPPost200Response
func (a *VatGatewayService) VatInitDefaultGatewayPostExecute(r VatGatewayVatInitDefaultGatewayPostRequest) (*MerchantAuthSsoLoginOTPPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantAuthSsoLoginOTPPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VatGatewayService.VatInitDefaultGatewayPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/vat/init_default_gateway"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type VatGatewayVatSetupGatewayPostRequest struct {
	ctx context.Context
	ApiService *VatGatewayService
	unibeeApiMerchantVatSetupGatewayReq *UnibeeApiMerchantVatSetupGatewayReq
}

func (r VatGatewayVatSetupGatewayPostRequest) UnibeeApiMerchantVatSetupGatewayReq(unibeeApiMerchantVatSetupGatewayReq UnibeeApiMerchantVatSetupGatewayReq) VatGatewayVatSetupGatewayPostRequest {
	r.unibeeApiMerchantVatSetupGatewayReq = &unibeeApiMerchantVatSetupGatewayReq
	return r
}

func (r VatGatewayVatSetupGatewayPostRequest) Execute() (*MerchantEmailGatewaySetupPost200Response, *http.Response, error) {
	return r.ApiService.VatSetupGatewayPostExecute(r)
}

/*
VatSetupGatewayPost Vat Gateway Setup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return VatGatewayVatSetupGatewayPostRequest
*/
func (a *VatGatewayService) VatSetupGatewayPost(ctx context.Context) VatGatewayVatSetupGatewayPostRequest {
	return VatGatewayVatSetupGatewayPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantEmailGatewaySetupPost200Response
func (a *VatGatewayService) VatSetupGatewayPostExecute(r VatGatewayVatSetupGatewayPostRequest) (*MerchantEmailGatewaySetupPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantEmailGatewaySetupPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VatGatewayService.VatSetupGatewayPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/vat/setup_gateway"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantVatSetupGatewayReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantVatSetupGatewayReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantVatSetupGatewayReq
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

type VatGatewayVatVatNumberValidatePostRequest struct {
	ctx context.Context
	ApiService *VatGatewayService
	unibeeApiMerchantVatNumberValidateReq *UnibeeApiMerchantVatNumberValidateReq
}

func (r VatGatewayVatVatNumberValidatePostRequest) UnibeeApiMerchantVatNumberValidateReq(unibeeApiMerchantVatNumberValidateReq UnibeeApiMerchantVatNumberValidateReq) VatGatewayVatVatNumberValidatePostRequest {
	r.unibeeApiMerchantVatNumberValidateReq = &unibeeApiMerchantVatNumberValidateReq
	return r
}

func (r VatGatewayVatVatNumberValidatePostRequest) Execute() (*MerchantVatVatNumberValidatePost200Response, *http.Response, error) {
	return r.ApiService.VatVatNumberValidatePostExecute(r)
}

/*
VatVatNumberValidatePost Vat Number Validation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return VatGatewayVatVatNumberValidatePostRequest
*/
func (a *VatGatewayService) VatVatNumberValidatePost(ctx context.Context) VatGatewayVatVatNumberValidatePostRequest {
	return VatGatewayVatVatNumberValidatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantVatVatNumberValidatePost200Response
func (a *VatGatewayService) VatVatNumberValidatePostExecute(r VatGatewayVatVatNumberValidatePostRequest) (*MerchantVatVatNumberValidatePost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantVatVatNumberValidatePost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VatGatewayService.VatVatNumberValidatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/vat/vat_number_validate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantVatNumberValidateReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantVatNumberValidateReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantVatNumberValidateReq
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
