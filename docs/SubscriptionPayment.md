# \SubscriptionPayment

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionNewOnetimeAddonPaymentPost**](SubscriptionPayment.md#SubscriptionNewOnetimeAddonPaymentPost) | **Post** /merchant/subscription/new_onetime_addon_payment | New Subscription Onetime Addon Payment
[**SubscriptionOnetimeAddonListGet**](SubscriptionPayment.md#SubscriptionOnetimeAddonListGet) | **Get** /merchant/subscription/onetime_addon_list | Get Subscription Onetime Addon List
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


## SubscriptionOnetimeAddonListGet

> MerchantSubscriptionOnetimeAddonListGet200Response SubscriptionOnetimeAddonListGet(ctx).UserId(userId).Page(page).Count(count).Execute()

Get Subscription Onetime Addon List

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
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionPayment.SubscriptionOnetimeAddonListGet(context.Background()).UserId(userId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPayment.SubscriptionOnetimeAddonListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionOnetimeAddonListGet`: MerchantSubscriptionOnetimeAddonListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPayment.SubscriptionOnetimeAddonListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionOnetimeAddonListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionOnetimeAddonListGet200Response**](MerchantSubscriptionOnetimeAddonListGet200Response.md)

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

