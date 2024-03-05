# \MerchantSubscriptionTimelineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSubscriptionTimelineListGet**](MerchantSubscriptionTimelineAPI.md#MerchantSubscriptionTimelineListGet) | **Get** /merchant/subscription/timeline_list | Merchant Subscription TimeLine List
[**MerchantSubscriptionTimelineListPost**](MerchantSubscriptionTimelineAPI.md#MerchantSubscriptionTimelineListPost) | **Post** /merchant/subscription/timeline_list | Merchant Subscription TimeLine List



## MerchantSubscriptionTimelineListGet

> MerchantSubscriptionTimelineListGet200Response MerchantSubscriptionTimelineListGet(ctx).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Merchant Subscription TimeLine List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	userId := int32(56) // int32 | Filter UserId, Default All  (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListGet(context.Background()).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionTimelineListGet`: MerchantSubscriptionTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionTimelineListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** | Filter UserId, Default All  | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start WIth 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionTimelineListGet200Response**](MerchantSubscriptionTimelineListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionTimelineListPost

> MerchantSubscriptionTimelineListGet200Response MerchantSubscriptionTimelineListPost(ctx).UnibeeApiMerchantSubscriptionTimeLineListReq(unibeeApiMerchantSubscriptionTimeLineListReq).Execute()

Merchant Subscription TimeLine List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantSubscriptionTimeLineListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionTimeLineListReq() // UnibeeApiMerchantSubscriptionTimeLineListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListPost(context.Background()).UnibeeApiMerchantSubscriptionTimeLineListReq(unibeeApiMerchantSubscriptionTimeLineListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionTimelineListPost`: MerchantSubscriptionTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionTimelineAPI.MerchantSubscriptionTimelineListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionTimelineListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionTimeLineListReq** | [**UnibeeApiMerchantSubscriptionTimeLineListReq**](UnibeeApiMerchantSubscriptionTimeLineListReq.md) |  | 

### Return type

[**MerchantSubscriptionTimelineListGet200Response**](MerchantSubscriptionTimelineListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

