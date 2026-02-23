# \Session

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionCustomerPortalUrlGet**](Session.md#SessionCustomerPortalUrlGet) | **Get** /merchant/session/customer_portal_url | Get Customer Portal Url
[**SessionCustomerPortalUrlPost**](Session.md#SessionCustomerPortalUrlPost) | **Post** /merchant/session/customer_portal_url | Get Customer Portal Url
[**SessionNewSessionPost**](Session.md#SessionNewSessionPost) | **Post** /merchant/session/new_session | New Checkout Session
[**SessionUserSubUpdateUrlGet**](Session.md#SessionUserSubUpdateUrlGet) | **Get** /merchant/session/user_sub_update_url | Get User Subscription Update Page Url
[**SessionUserSubUpdateUrlPost**](Session.md#SessionUserSubUpdateUrlPost) | **Post** /merchant/session/user_sub_update_url | Get User Subscription Update Page Url



## SessionCustomerPortalUrlGet

> MerchantGetLicenseUpdateUrlGet200Response SessionCustomerPortalUrlGet(ctx).Email(email).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).PlanId(planId).VatCountryCode(vatCountryCode).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()

Get Customer Portal Url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	email := "email_example" // string | Email, unique, either ExternalUserId&Email or UserId needed (optional)
	userId := int64(789) // int64 | UserId, unique, either ExternalUserId&Email or UserId needed (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, unique, either ExternalUserId&Email or UserId needed (optional)
	productId := int64(789) // int64 | Id of product, default product will use if productId not specified (optional)
	planId := int64(789) // int64 | Id of plan to update (optional)
	vatCountryCode := "vatCountryCode_example" // string | Vat Country Code (optional)
	returnUrl := "returnUrl_example" // string | ReturnUrl (optional)
	cancelUrl := "cancelUrl_example" // string | CancelUrl (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionCustomerPortalUrlGet(context.Background()).Email(email).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).PlanId(planId).VatCountryCode(vatCountryCode).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionCustomerPortalUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionCustomerPortalUrlGet`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionCustomerPortalUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionCustomerPortalUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email, unique, either ExternalUserId&amp;Email or UserId needed | 
 **userId** | **int64** | UserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **externalUserId** | **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **productId** | **int64** | Id of product, default product will use if productId not specified | 
 **planId** | **int64** | Id of plan to update | 
 **vatCountryCode** | **string** | Vat Country Code | 
 **returnUrl** | **string** | ReturnUrl | 
 **cancelUrl** | **string** | CancelUrl | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionCustomerPortalUrlPost

> MerchantGetLicenseUpdateUrlGet200Response SessionCustomerPortalUrlPost(ctx).UnibeeApiMerchantSessionCustomerPortalUrlReq(unibeeApiMerchantSessionCustomerPortalUrlReq).Execute()

Get Customer Portal Url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	unibeeApiMerchantSessionCustomerPortalUrlReq := *openapiclient.NewUnibeeApiMerchantSessionCustomerPortalUrlReq() // UnibeeApiMerchantSessionCustomerPortalUrlReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionCustomerPortalUrlPost(context.Background()).UnibeeApiMerchantSessionCustomerPortalUrlReq(unibeeApiMerchantSessionCustomerPortalUrlReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionCustomerPortalUrlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionCustomerPortalUrlPost`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionCustomerPortalUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionCustomerPortalUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSessionCustomerPortalUrlReq** | [**UnibeeApiMerchantSessionCustomerPortalUrlReq**](UnibeeApiMerchantSessionCustomerPortalUrlReq.md) |  | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionNewSessionPost

> MerchantSessionNewSessionPost200Response SessionNewSessionPost(ctx).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()

New Checkout Session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	unibeeApiMerchantSessionNewReq := *openapiclient.NewUnibeeApiMerchantSessionNewReq("Email_example") // UnibeeApiMerchantSessionNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionNewSessionPost(context.Background()).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionNewSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionNewSessionPost`: MerchantSessionNewSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionNewSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionNewSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSessionNewReq** | [**UnibeeApiMerchantSessionNewReq**](UnibeeApiMerchantSessionNewReq.md) |  | 

### Return type

[**MerchantSessionNewSessionPost200Response**](MerchantSessionNewSessionPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionUserSubUpdateUrlGet

> MerchantGetLicenseUpdateUrlGet200Response SessionUserSubUpdateUrlGet(ctx).Email(email).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).PlanId(planId).VatCountryCode(vatCountryCode).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()

Get User Subscription Update Page Url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	email := "email_example" // string | Email, unique, either ExternalUserId&Email or UserId needed (optional)
	userId := int64(789) // int64 | UserId, unique, either ExternalUserId&Email or UserId needed (optional)
	externalUserId := "externalUserId_example" // string | ExternalUserId, unique, either ExternalUserId&Email or UserId needed (optional)
	productId := int64(789) // int64 | Id of product,default product will use if productId not specified (optional)
	planId := int64(789) // int64 | Id of plan to update (optional)
	vatCountryCode := "vatCountryCode_example" // string | Vat Country Code (optional)
	returnUrl := "returnUrl_example" // string | ReturnUrl (optional)
	cancelUrl := "cancelUrl_example" // string | CancelUrl (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionUserSubUpdateUrlGet(context.Background()).Email(email).UserId(userId).ExternalUserId(externalUserId).ProductId(productId).PlanId(planId).VatCountryCode(vatCountryCode).ReturnUrl(returnUrl).CancelUrl(cancelUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionUserSubUpdateUrlGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionUserSubUpdateUrlGet`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionUserSubUpdateUrlGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionUserSubUpdateUrlGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Email, unique, either ExternalUserId&amp;Email or UserId needed | 
 **userId** | **int64** | UserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **externalUserId** | **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | 
 **productId** | **int64** | Id of product,default product will use if productId not specified | 
 **planId** | **int64** | Id of plan to update | 
 **vatCountryCode** | **string** | Vat Country Code | 
 **returnUrl** | **string** | ReturnUrl | 
 **cancelUrl** | **string** | CancelUrl | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionUserSubUpdateUrlPost

> MerchantGetLicenseUpdateUrlGet200Response SessionUserSubUpdateUrlPost(ctx).UnibeeApiMerchantSessionNewSubUpdatePageReq(unibeeApiMerchantSessionNewSubUpdatePageReq).Execute()

Get User Subscription Update Page Url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	unibeeApiMerchantSessionNewSubUpdatePageReq := *openapiclient.NewUnibeeApiMerchantSessionNewSubUpdatePageReq() // UnibeeApiMerchantSessionNewSubUpdatePageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionUserSubUpdateUrlPost(context.Background()).UnibeeApiMerchantSessionNewSubUpdatePageReq(unibeeApiMerchantSessionNewSubUpdatePageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionUserSubUpdateUrlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionUserSubUpdateUrlPost`: MerchantGetLicenseUpdateUrlGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionUserSubUpdateUrlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionUserSubUpdateUrlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSessionNewSubUpdatePageReq** | [**UnibeeApiMerchantSessionNewSubUpdatePageReq**](UnibeeApiMerchantSessionNewSubUpdatePageReq.md) |  | 

### Return type

[**MerchantGetLicenseUpdateUrlGet200Response**](MerchantGetLicenseUpdateUrlGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

