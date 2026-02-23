# \MetricEvent

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricEventCurrentValuePost**](MetricEvent.md#MetricEventCurrentValuePost) | **Post** /merchant/metric/event/current_value | Get Merchant Metric Event Current Value
[**MetricEventDeletePost**](MetricEvent.md#MetricEventDeletePost) | **Post** /merchant/metric/event/delete | Del Merchant Metric Event
[**MetricEventListGet**](MetricEvent.md#MetricEventListGet) | **Get** /merchant/metric/event_list | Metric Event List
[**MetricEventListPost**](MetricEvent.md#MetricEventListPost) | **Post** /merchant/metric/event_list | Metric Event List
[**MetricEventNewPost**](MetricEvent.md#MetricEventNewPost) | **Post** /merchant/metric/event/new | New Merchant Metric Event



## MetricEventCurrentValuePost

> MerchantMetricEventCurrentValuePost200Response MetricEventCurrentValuePost(ctx).UnibeeApiMerchantMetricEventCurrentValueReq(unibeeApiMerchantMetricEventCurrentValueReq).Execute()

Get Merchant Metric Event Current Value

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
	unibeeApiMerchantMetricEventCurrentValueReq := *openapiclient.NewUnibeeApiMerchantMetricEventCurrentValueReq("MetricCode_example") // UnibeeApiMerchantMetricEventCurrentValueReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricEvent.MetricEventCurrentValuePost(context.Background()).UnibeeApiMerchantMetricEventCurrentValueReq(unibeeApiMerchantMetricEventCurrentValueReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricEvent.MetricEventCurrentValuePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventCurrentValuePost`: MerchantMetricEventCurrentValuePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricEvent.MetricEventCurrentValuePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventCurrentValuePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricEventCurrentValueReq** | [**UnibeeApiMerchantMetricEventCurrentValueReq**](UnibeeApiMerchantMetricEventCurrentValueReq.md) |  | 

### Return type

[**MerchantMetricEventCurrentValuePost200Response**](MerchantMetricEventCurrentValuePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricEventDeletePost

> MerchantAuthSsoClearTotpPost200Response MetricEventDeletePost(ctx).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()

Del Merchant Metric Event

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
	unibeeApiMerchantMetricDeleteEventReq := *openapiclient.NewUnibeeApiMerchantMetricDeleteEventReq("ExternalEventId_example", "MetricCode_example") // UnibeeApiMerchantMetricDeleteEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricEvent.MetricEventDeletePost(context.Background()).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricEvent.MetricEventDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventDeletePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricEvent.MetricEventDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDeleteEventReq** | [**UnibeeApiMerchantMetricDeleteEventReq**](UnibeeApiMerchantMetricDeleteEventReq.md) |  | 

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


## MetricEventListGet

> MerchantMetricEventListGet200Response MetricEventListGet(ctx).UserIds(userIds).MetricIds(metricIds).SubscriptionIds(subscriptionIds).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).InvoiceId(invoiceId).Execute()

Metric Event List

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
	userIds := []int64{int64(123)} // []int64 | Filter UserIds, Default All (optional)
	metricIds := []int64{int64(123)} // []int64 | Filter MetricIds, Default All (optional)
	subscriptionIds := []string{"Inner_example"} // []string | Filter SubscriptionIds, Default All (optional)
	sortField := "sortField_example" // string | Sort，user_id|gmt_create，Default gmt_create (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count OF Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)
	invoiceId := "invoiceId_example" // string | When set, list events for this invoice period (match subscription_period_start/end with invoice PeriodStart/PeriodEnd); overrides CreateTimeStart/CreateTimeEnd for period scope (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricEvent.MetricEventListGet(context.Background()).UserIds(userIds).MetricIds(metricIds).SubscriptionIds(subscriptionIds).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).InvoiceId(invoiceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricEvent.MetricEventListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventListGet`: MerchantMetricEventListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricEvent.MetricEventListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | **[]int64** | Filter UserIds, Default All | 
 **metricIds** | **[]int64** | Filter MetricIds, Default All | 
 **subscriptionIds** | **[]string** | Filter SubscriptionIds, Default All | 
 **sortField** | **string** | Sort，user_id|gmt_create，Default gmt_create | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count OF Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 
 **invoiceId** | **string** | When set, list events for this invoice period (match subscription_period_start/end with invoice PeriodStart/PeriodEnd); overrides CreateTimeStart/CreateTimeEnd for period scope | 

### Return type

[**MerchantMetricEventListGet200Response**](MerchantMetricEventListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricEventListPost

> MerchantMetricEventListGet200Response MetricEventListPost(ctx).UnibeeApiMerchantMetricEventListReq(unibeeApiMerchantMetricEventListReq).Execute()

Metric Event List

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
	unibeeApiMerchantMetricEventListReq := *openapiclient.NewUnibeeApiMerchantMetricEventListReq() // UnibeeApiMerchantMetricEventListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricEvent.MetricEventListPost(context.Background()).UnibeeApiMerchantMetricEventListReq(unibeeApiMerchantMetricEventListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricEvent.MetricEventListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventListPost`: MerchantMetricEventListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricEvent.MetricEventListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricEventListReq** | [**UnibeeApiMerchantMetricEventListReq**](UnibeeApiMerchantMetricEventListReq.md) |  | 

### Return type

[**MerchantMetricEventListGet200Response**](MerchantMetricEventListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricEventNewPost

> MerchantMetricEventNewPost200Response MetricEventNewPost(ctx).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()

New Merchant Metric Event

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
	unibeeApiMerchantMetricNewEventReq := *openapiclient.NewUnibeeApiMerchantMetricNewEventReq("ExternalEventId_example", "MetricCode_example") // UnibeeApiMerchantMetricNewEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricEvent.MetricEventNewPost(context.Background()).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricEvent.MetricEventNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEventNewPost`: MerchantMetricEventNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MetricEvent.MetricEventNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEventNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricNewEventReq** | [**UnibeeApiMerchantMetricNewEventReq**](UnibeeApiMerchantMetricNewEventReq.md) |  | 

### Return type

[**MerchantMetricEventNewPost200Response**](MerchantMetricEventNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

