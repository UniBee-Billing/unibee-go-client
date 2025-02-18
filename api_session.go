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


// SessionService Session service
type SessionService service

type SessionSessionNewSessionPostRequest struct {
	ctx context.Context
	ApiService *SessionService
	unibeeApiMerchantSessionNewReq *UnibeeApiMerchantSessionNewReq
}

func (r SessionSessionNewSessionPostRequest) UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq UnibeeApiMerchantSessionNewReq) SessionSessionNewSessionPostRequest {
	r.unibeeApiMerchantSessionNewReq = &unibeeApiMerchantSessionNewReq
	return r
}

func (r SessionSessionNewSessionPostRequest) Execute() (*MerchantSessionNewSessionPost200Response, *http.Response, error) {
	return r.ApiService.SessionNewSessionPostExecute(r)
}

/*
SessionNewSessionPost New Session

New session for user portal or web component

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return SessionSessionNewSessionPostRequest
*/
func (a *SessionService) SessionNewSessionPost(ctx context.Context) SessionSessionNewSessionPostRequest {
	return SessionSessionNewSessionPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MerchantSessionNewSessionPost200Response
func (a *SessionService) SessionNewSessionPostExecute(r SessionSessionNewSessionPostRequest) (*MerchantSessionNewSessionPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MerchantSessionNewSessionPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionService.SessionNewSessionPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/merchant/session/new_session"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.unibeeApiMerchantSessionNewReq == nil {
		return localVarReturnValue, nil, reportError("unibeeApiMerchantSessionNewReq is required and must be specified")
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
	localVarPostBody = r.unibeeApiMerchantSessionNewReq
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
