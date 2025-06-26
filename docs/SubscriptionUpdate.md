# \SubscriptionUpdate

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionActiveTemporarilyPost**](SubscriptionUpdate.md#SubscriptionActiveTemporarilyPost) | **Post** /merchant/subscription/active_temporarily | Subscription Active Temporarily
[**SubscriptionPendingUpdateDetailGet**](SubscriptionUpdate.md#SubscriptionPendingUpdateDetailGet) | **Get** /merchant/subscription/pending_update_detail | Subscription Pending Update Detail
[**SubscriptionPendingUpdateListGet**](SubscriptionUpdate.md#SubscriptionPendingUpdateListGet) | **Get** /merchant/subscription/pending_update_list | Get Subscription Pending Update List
[**SubscriptionPendingUpdateListPost**](SubscriptionUpdate.md#SubscriptionPendingUpdateListPost) | **Post** /merchant/subscription/pending_update_list | Get Subscription Pending Update List
[**SubscriptionRenewPost**](SubscriptionUpdate.md#SubscriptionRenewPost) | **Post** /merchant/subscription/renew | Renew Subscription
[**SubscriptionUpdatePreviewPost**](SubscriptionUpdate.md#SubscriptionUpdatePreviewPost) | **Post** /merchant/subscription/update_preview | Update Subscription Preview
[**SubscriptionUpdateSubmitPost**](SubscriptionUpdate.md#SubscriptionUpdateSubmitPost) | **Post** /merchant/subscription/update_submit | Update Subscription



## SubscriptionActiveTemporarilyPost

> MerchantAuthSsoClearTotpPost200Response SubscriptionActiveTemporarilyPost(ctx).UnibeeApiMerchantSubscriptionActiveTemporarilyReq(unibeeApiMerchantSubscriptionActiveTemporarilyReq).Execute()

Subscription Active Temporarily



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
	unibeeApiMerchantSubscriptionActiveTemporarilyReq := *openapiclient.NewUnibeeApiMerchantSubscriptionActiveTemporarilyReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionActiveTemporarilyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionActiveTemporarilyPost(context.Background()).UnibeeApiMerchantSubscriptionActiveTemporarilyReq(unibeeApiMerchantSubscriptionActiveTemporarilyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionActiveTemporarilyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionActiveTemporarilyPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionActiveTemporarilyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionActiveTemporarilyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionActiveTemporarilyReq** | [**UnibeeApiMerchantSubscriptionActiveTemporarilyReq**](UnibeeApiMerchantSubscriptionActiveTemporarilyReq.md) |  | 

### Return type

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPendingUpdateDetailGet

> MerchantSubscriptionPendingUpdateDetailGet200Response SubscriptionPendingUpdateDetailGet(ctx).SubscriptionPendingUpdateId(subscriptionPendingUpdateId).Execute()

Subscription Pending Update Detail

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
	subscriptionPendingUpdateId := "subscriptionPendingUpdateId_example" // string | SubscriptionPendingUpdateId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionPendingUpdateDetailGet(context.Background()).SubscriptionPendingUpdateId(subscriptionPendingUpdateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionPendingUpdateDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPendingUpdateDetailGet`: MerchantSubscriptionPendingUpdateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionPendingUpdateDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPendingUpdateDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionPendingUpdateId** | **string** | SubscriptionPendingUpdateId | 

### Return type

[**MerchantSubscriptionPendingUpdateDetailGet200Response**](MerchantSubscriptionPendingUpdateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPendingUpdateListGet

> MerchantSubscriptionPendingUpdateListGet200Response SubscriptionPendingUpdateListGet(ctx).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Get Subscription Pending Update List

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
	subscriptionId := "subscriptionId_example" // string | SubscriptionId
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionPendingUpdateListGet(context.Background()).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionPendingUpdateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPendingUpdateListGet`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionPendingUpdateListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPendingUpdateListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionPendingUpdateListGet200Response**](MerchantSubscriptionPendingUpdateListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPendingUpdateListPost

> MerchantSubscriptionPendingUpdateListGet200Response SubscriptionPendingUpdateListPost(ctx).UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq).Execute()

Get Subscription Pending Update List

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
	unibeeApiMerchantSubscriptionPendingUpdateListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionPendingUpdateListReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionPendingUpdateListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionPendingUpdateListPost(context.Background()).UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionPendingUpdateListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPendingUpdateListPost`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionPendingUpdateListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPendingUpdateListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionPendingUpdateListReq** | [**UnibeeApiMerchantSubscriptionPendingUpdateListReq**](UnibeeApiMerchantSubscriptionPendingUpdateListReq.md) |  | 

### Return type

[**MerchantSubscriptionPendingUpdateListGet200Response**](MerchantSubscriptionPendingUpdateListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionRenewPost

> MerchantSubscriptionRenewPost200Response SubscriptionRenewPost(ctx).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()

Renew Subscription



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
	unibeeApiMerchantSubscriptionRenewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionRenewReq() // UnibeeApiMerchantSubscriptionRenewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionRenewPost(context.Background()).UnibeeApiMerchantSubscriptionRenewReq(unibeeApiMerchantSubscriptionRenewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionRenewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRenewPost`: MerchantSubscriptionRenewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionRenewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRenewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionRenewReq** | [**UnibeeApiMerchantSubscriptionRenewReq**](UnibeeApiMerchantSubscriptionRenewReq.md) |  | 

### Return type

[**MerchantSubscriptionRenewPost200Response**](MerchantSubscriptionRenewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdatePreviewPost

> MerchantSubscriptionUpdatePreviewPost200Response SubscriptionUpdatePreviewPost(ctx).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()

Update Subscription Preview

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
	unibeeApiMerchantSubscriptionUpdatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionUpdatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionUpdatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdatePreviewPost`: MerchantSubscriptionUpdatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionUpdatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdatePreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdatePreviewReq** | [**UnibeeApiMerchantSubscriptionUpdatePreviewReq**](UnibeeApiMerchantSubscriptionUpdatePreviewReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdatePreviewPost200Response**](MerchantSubscriptionUpdatePreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdateSubmitPost

> MerchantSubscriptionUpdateSubmitPost200Response SubscriptionUpdateSubmitPost(ctx).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()

Update Subscription

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
	unibeeApiMerchantSubscriptionUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdateReq(int64(123), int64(123)) // UnibeeApiMerchantSubscriptionUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUpdate.SubscriptionUpdateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUpdate.SubscriptionUpdateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdateSubmitPost`: MerchantSubscriptionUpdateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUpdate.SubscriptionUpdateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateSubmitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUpdateReq** | [**UnibeeApiMerchantSubscriptionUpdateReq**](UnibeeApiMerchantSubscriptionUpdateReq.md) |  | 

### Return type

[**MerchantSubscriptionUpdateSubmitPost200Response**](MerchantSubscriptionUpdateSubmitPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

