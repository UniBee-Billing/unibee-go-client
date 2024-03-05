# \MerchantSubscriptionPendingUpdateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSubscriptionPendingUpdateListGet**](MerchantSubscriptionPendingUpdateAPI.md#MerchantSubscriptionPendingUpdateListGet) | **Get** /merchant/subscription/pending_update_list | Merchant SubscriptionPendingUpdate List
[**MerchantSubscriptionPendingUpdateListPost**](MerchantSubscriptionPendingUpdateAPI.md#MerchantSubscriptionPendingUpdateListPost) | **Post** /merchant/subscription/pending_update_list | Merchant SubscriptionPendingUpdate List



## MerchantSubscriptionPendingUpdateListGet

> MerchantSubscriptionPendingUpdateListGet200Response MerchantSubscriptionPendingUpdateListGet(ctx).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Merchant SubscriptionPendingUpdate List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | SubscriptionId
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListGet(context.Background()).SubscriptionId(subscriptionId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionPendingUpdateListGet`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionPendingUpdateListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start WIth 0 | 
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


## MerchantSubscriptionPendingUpdateListPost

> MerchantSubscriptionPendingUpdateListGet200Response MerchantSubscriptionPendingUpdateListPost(ctx).UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq).Execute()

Merchant SubscriptionPendingUpdate List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantSubscriptionPendingUpdateListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionPendingUpdateListReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionPendingUpdateListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListPost(context.Background()).UnibeeApiMerchantSubscriptionPendingUpdateListReq(unibeeApiMerchantSubscriptionPendingUpdateListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionPendingUpdateListPost`: MerchantSubscriptionPendingUpdateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionPendingUpdateAPI.MerchantSubscriptionPendingUpdateListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionPendingUpdateListPostRequest struct via the builder pattern


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

