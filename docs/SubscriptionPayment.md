# \SubscriptionPayment

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionNewOnetimeAddonPaymentPost**](SubscriptionPayment.md#SubscriptionNewOnetimeAddonPaymentPost) | **Post** /merchant/subscription/new_onetime_addon_payment | New Subscription Onetime Addon Payment
[**SubscriptionNewOnetimeAddonPreviewPost**](SubscriptionPayment.md#SubscriptionNewOnetimeAddonPreviewPost) | **Post** /merchant/subscription/new_onetime_addon_preview | New Subscription Onetime Addon Preview
[**SubscriptionOnetimeAddonPurchaseListGet**](SubscriptionPayment.md#SubscriptionOnetimeAddonPurchaseListGet) | **Get** /merchant/subscription/onetime_addon_purchase_list | Get Subscription Onetime Addon Purchase History List
[**SubscriptionPaymentNewPost**](SubscriptionPayment.md#SubscriptionPaymentNewPost) | **Post** /merchant/subscription/payment/new | New Subscription Payment



## SubscriptionNewOnetimeAddonPaymentPost

> MerchantSubscriptionNewOnetimeAddonPaymentPost200Response SubscriptionNewOnetimeAddonPaymentPost(ctx).UnibeeApiMerchantSubscriptionOnetimeAddonNewReq(unibeeApiMerchantSubscriptionOnetimeAddonNewReq).Execute()

New Subscription Onetime Addon Payment



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
	unibeeApiMerchantSubscriptionOnetimeAddonNewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq(int64(123), int64(123)) // UnibeeApiMerchantSubscriptionOnetimeAddonNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPayment.SubscriptionNewOnetimeAddonPaymentPost(context.Background()).UnibeeApiMerchantSubscriptionOnetimeAddonNewReq(unibeeApiMerchantSubscriptionOnetimeAddonNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPayment.SubscriptionNewOnetimeAddonPaymentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionNewOnetimeAddonPaymentPost`: MerchantSubscriptionNewOnetimeAddonPaymentPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPayment.SubscriptionNewOnetimeAddonPaymentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionNewOnetimeAddonPaymentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionOnetimeAddonNewReq** | [**UnibeeApiMerchantSubscriptionOnetimeAddonNewReq**](UnibeeApiMerchantSubscriptionOnetimeAddonNewReq.md) |  | 

### Return type

[**MerchantSubscriptionNewOnetimeAddonPaymentPost200Response**](MerchantSubscriptionNewOnetimeAddonPaymentPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionNewOnetimeAddonPreviewPost

> MerchantSubscriptionNewOnetimeAddonPreviewPost200Response SubscriptionNewOnetimeAddonPreviewPost(ctx).UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq(unibeeApiMerchantSubscriptionOnetimeAddonPreviewReq).Execute()

New Subscription Onetime Addon Preview



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
	unibeeApiMerchantSubscriptionOnetimeAddonPreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq(int64(123), int64(123)) // UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPayment.SubscriptionNewOnetimeAddonPreviewPost(context.Background()).UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq(unibeeApiMerchantSubscriptionOnetimeAddonPreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPayment.SubscriptionNewOnetimeAddonPreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionNewOnetimeAddonPreviewPost`: MerchantSubscriptionNewOnetimeAddonPreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPayment.SubscriptionNewOnetimeAddonPreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionNewOnetimeAddonPreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionOnetimeAddonPreviewReq** | [**UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq**](UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionNewOnetimeAddonPreviewPost200Response**](MerchantSubscriptionNewOnetimeAddonPreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionOnetimeAddonPurchaseListGet

> MerchantSubscriptionOnetimeAddonPurchaseListGet200Response SubscriptionOnetimeAddonPurchaseListGet(ctx).UserId(userId).Page(page).Count(count).Execute()

Get Subscription Onetime Addon Purchase History List

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
	userId := int64(789) // int64 | UserId
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page，Default 100  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPayment.SubscriptionOnetimeAddonPurchaseListGet(context.Background()).UserId(userId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPayment.SubscriptionOnetimeAddonPurchaseListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOnetimeAddonPurchaseListGet`: MerchantSubscriptionOnetimeAddonPurchaseListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPayment.SubscriptionOnetimeAddonPurchaseListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOnetimeAddonPurchaseListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page，Default 100  | 

### Return type

[**MerchantSubscriptionOnetimeAddonPurchaseListGet200Response**](MerchantSubscriptionOnetimeAddonPurchaseListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPaymentNewPost

> MerchantPaymentNewPost200Response SubscriptionPaymentNewPost(ctx).UnibeeApiMerchantSubscriptionNewPaymentReq(unibeeApiMerchantSubscriptionNewPaymentReq).Execute()

New Subscription Payment

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
	unibeeApiMerchantSubscriptionNewPaymentReq := *openapiclient.NewUnibeeApiMerchantSubscriptionNewPaymentReq(int64(123)) // UnibeeApiMerchantSubscriptionNewPaymentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPayment.SubscriptionPaymentNewPost(context.Background()).UnibeeApiMerchantSubscriptionNewPaymentReq(unibeeApiMerchantSubscriptionNewPaymentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPayment.SubscriptionPaymentNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPaymentNewPost`: MerchantPaymentNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPayment.SubscriptionPaymentNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPaymentNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionNewPaymentReq** | [**UnibeeApiMerchantSubscriptionNewPaymentReq**](UnibeeApiMerchantSubscriptionNewPaymentReq.md) |  | 

### Return type

[**MerchantPaymentNewPost200Response**](MerchantPaymentNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

