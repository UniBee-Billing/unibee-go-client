# \SubscriptionPendingUpdate

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionPendingUpdateListGet**](SubscriptionPendingUpdate.md#SubscriptionPendingUpdateListGet) | **Get** /merchant/subscription/pending_update_list | SubscriptionPendingUpdateList
[**SubscriptionPendingUpdateListPost**](SubscriptionPendingUpdate.md#SubscriptionPendingUpdateListPost) | **Post** /merchant/subscription/pending_update_list | SubscriptionPendingUpdateList



## SubscriptionPendingUpdateListGet

> MerchantSubscriptionPendingUpdateListGet200Response SubscriptionPendingUpdateListGet(ctx).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

SubscriptionPendingUpdateList

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
	resp, r, err := apiClient.SubscriptionPendingUpdate.SubscriptionPendingUpdateListGet(context.Background()).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPendingUpdate.SubscriptionPendingUpdateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPendingUpdateListGet`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPendingUpdate.SubscriptionPendingUpdateListGet`: %v\n", resp)
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

SubscriptionPendingUpdateList

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
	resp, r, err := apiClient.SubscriptionPendingUpdate.SubscriptionPendingUpdateListPost(context.Background()).UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPendingUpdate.SubscriptionPendingUpdateListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionPendingUpdateListPost`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionPendingUpdate.SubscriptionPendingUpdateListPost`: %v\n", resp)
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

