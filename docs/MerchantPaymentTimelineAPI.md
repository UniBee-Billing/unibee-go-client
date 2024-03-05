# \MerchantPaymentTimelineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantPaymentPaymentTimelineListGet**](MerchantPaymentTimelineAPI.md#MerchantPaymentPaymentTimelineListGet) | **Get** /merchant/payment/payment_timeline_list | Payment TimeLine List
[**MerchantPaymentPaymentTimelineListPost**](MerchantPaymentTimelineAPI.md#MerchantPaymentPaymentTimelineListPost) | **Post** /merchant/payment/payment_timeline_list | Payment TimeLine List



## MerchantPaymentPaymentTimelineListGet

> MerchantPaymentPaymentTimelineListGet200Response MerchantPaymentPaymentTimelineListGet(ctx).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Payment TimeLine List

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
	userId := int32(56) // int32 | Filter UserId, Default All (optional)
	sortField := "sortField_example" // string | Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListGet(context.Background()).UserId(userId).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPaymentPaymentTimelineListGet`: MerchantPaymentPaymentTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPaymentPaymentTimelineListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** | Filter UserId, Default All | 
 **sortField** | **string** | Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantPaymentPaymentTimelineListGet200Response**](MerchantPaymentPaymentTimelineListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantPaymentPaymentTimelineListPost

> MerchantPaymentPaymentTimelineListGet200Response MerchantPaymentPaymentTimelineListPost(ctx).UnibeeApiMerchantPaymentTimeLineListReq(unibeeApiMerchantPaymentTimeLineListReq).Execute()

Payment TimeLine List

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
	unibeeApiMerchantPaymentTimeLineListReq := *openapiclient.NewUnibeeApiMerchantPaymentTimeLineListReq() // UnibeeApiMerchantPaymentTimeLineListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListPost(context.Background()).UnibeeApiMerchantPaymentTimeLineListReq(unibeeApiMerchantPaymentTimeLineListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPaymentPaymentTimelineListPost`: MerchantPaymentPaymentTimelineListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPaymentTimelineAPI.MerchantPaymentPaymentTimelineListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPaymentPaymentTimelineListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPaymentTimeLineListReq** | [**UnibeeApiMerchantPaymentTimeLineListReq**](UnibeeApiMerchantPaymentTimeLineListReq.md) |  | 

### Return type

[**MerchantPaymentPaymentTimelineListGet200Response**](MerchantPaymentPaymentTimelineListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

