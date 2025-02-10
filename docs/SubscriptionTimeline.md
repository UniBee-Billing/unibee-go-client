# \SubscriptionTimeline

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionTimelineListGet**](SubscriptionTimeline.md#SubscriptionTimelineListGet) | **Get** /merchant/subscription/timeline_list | Get Subscription TimeLine List
[**SubscriptionTimelineListPost**](SubscriptionTimeline.md#SubscriptionTimelineListPost) | **Post** /merchant/subscription/timeline_list | Get Subscription TimeLine List



## SubscriptionTimelineListGet

> MerchantSubscriptionTimelineListGet200Response SubscriptionTimelineListGet(ctx).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Get Subscription TimeLine List

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
	userId := int64(789) // int64 | Filter UserId, Default All  (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionTimeline.SubscriptionTimelineListGet(context.Background()).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionTimeline.SubscriptionTimelineListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionTimelineListGet`: MerchantSubscriptionTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionTimeline.SubscriptionTimelineListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionTimelineListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | Filter UserId, Default All  | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionTimelineListGet200Response**](MerchantSubscriptionTimelineListGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionTimelineListPost

> MerchantSubscriptionTimelineListGet200Response SubscriptionTimelineListPost(ctx).UnibeeApiMerchantSubscriptionTimeLineListReq(unibeeApiMerchantSubscriptionTimeLineListReq).Execute()

Get Subscription TimeLine List

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
	unibeeApiMerchantSubscriptionTimeLineListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionTimeLineListReq() // UnibeeApiMerchantSubscriptionTimeLineListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionTimeline.SubscriptionTimelineListPost(context.Background()).UnibeeApiMerchantSubscriptionTimeLineListReq(unibeeApiMerchantSubscriptionTimeLineListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionTimeline.SubscriptionTimelineListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionTimelineListPost`: MerchantSubscriptionTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionTimeline.SubscriptionTimelineListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionTimelineListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionTimeLineListReq** | [**UnibeeApiMerchantSubscriptionTimeLineListReq**](UnibeeApiMerchantSubscriptionTimeLineListReq.md) |  | 

### Return type

[**MerchantSubscriptionTimelineListGet200Response**](MerchantSubscriptionTimelineListGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

