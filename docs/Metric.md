# \Metric

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetricDeletePost**](Metric.md#MetricDeletePost) | **Post** /merchant/metric/delete | Delete Merchant Metric
[**MetricDetailPost**](Metric.md#MetricDetailPost) | **Post** /merchant/metric/detail | Merchant Metric Detail
[**MetricEditPost**](Metric.md#MetricEditPost) | **Post** /merchant/metric/edit | Edit Merchant Metric
[**MetricListGet**](Metric.md#MetricListGet) | **Get** /merchant/metric/list | Get Merchant Metric list
[**MetricListPost**](Metric.md#MetricListPost) | **Post** /merchant/metric/list | Get Merchant Metric list
[**MetricNewPost**](Metric.md#MetricNewPost) | **Post** /merchant/metric/new | New Merchant Metric
[**MetricPlanLimitDeletePost**](Metric.md#MetricPlanLimitDeletePost) | **Post** /merchant/metric/plan/limit/delete | Delete Merchant Metric Plan TotalLimit
[**MetricPlanLimitEditPost**](Metric.md#MetricPlanLimitEditPost) | **Post** /merchant/metric/plan/limit/edit | Edit Merchant Metric Plan TotalLimit
[**MetricPlanLimitNewPost**](Metric.md#MetricPlanLimitNewPost) | **Post** /merchant/metric/plan/limit/new | New Merchant Metric Plan TotalLimit



## MetricDeletePost

> MerchantMetricDeletePost200Response MetricDeletePost(ctx).UnibeeApiMerchantMetricDeleteReq(unibeeApiMerchantMetricDeleteReq).Execute()

Delete Merchant Metric

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
	unibeeApiMerchantMetricDeleteReq := *openapiclient.NewUnibeeApiMerchantMetricDeleteReq(int64(123)) // UnibeeApiMerchantMetricDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricDeletePost(context.Background()).UnibeeApiMerchantMetricDeleteReq(unibeeApiMerchantMetricDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricDeletePost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDeleteReq** | [**UnibeeApiMerchantMetricDeleteReq**](UnibeeApiMerchantMetricDeleteReq.md) |  | 

### Return type

[**MerchantMetricDeletePost200Response**](MerchantMetricDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricDetailPost

> MerchantMetricDeletePost200Response MetricDetailPost(ctx).UnibeeApiMerchantMetricDetailReq(unibeeApiMerchantMetricDetailReq).Execute()

Merchant Metric Detail

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
	unibeeApiMerchantMetricDetailReq := *openapiclient.NewUnibeeApiMerchantMetricDetailReq(int64(123)) // UnibeeApiMerchantMetricDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricDetailPost(context.Background()).UnibeeApiMerchantMetricDetailReq(unibeeApiMerchantMetricDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricDetailPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDetailReq** | [**UnibeeApiMerchantMetricDetailReq**](UnibeeApiMerchantMetricDetailReq.md) |  | 

### Return type

[**MerchantMetricDeletePost200Response**](MerchantMetricDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricEditPost

> MerchantMetricDeletePost200Response MetricEditPost(ctx).UnibeeApiMerchantMetricEditReq(unibeeApiMerchantMetricEditReq).Execute()

Edit Merchant Metric

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
	unibeeApiMerchantMetricEditReq := *openapiclient.NewUnibeeApiMerchantMetricEditReq(int64(123), "MetricName_example") // UnibeeApiMerchantMetricEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricEditPost(context.Background()).UnibeeApiMerchantMetricEditReq(unibeeApiMerchantMetricEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricEditPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricEditReq** | [**UnibeeApiMerchantMetricEditReq**](UnibeeApiMerchantMetricEditReq.md) |  | 

### Return type

[**MerchantMetricDeletePost200Response**](MerchantMetricDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricListGet

> MerchantMetricListGet200Response MetricListGet(ctx).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Merchant Metric list

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
	sortField := "sortField_example" // string | Sort，user_id|gmt_create，Default gmt_create (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count OF Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricListGet(context.Background()).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricListGet`: MerchantMetricListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortField** | **string** | Sort，user_id|gmt_create，Default gmt_create | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count OF Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantMetricListGet200Response**](MerchantMetricListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricListPost

> MerchantMetricListGet200Response MetricListPost(ctx).UnibeeApiMerchantMetricListReq(unibeeApiMerchantMetricListReq).Execute()

Get Merchant Metric list

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
	unibeeApiMerchantMetricListReq := *openapiclient.NewUnibeeApiMerchantMetricListReq() // UnibeeApiMerchantMetricListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricListPost(context.Background()).UnibeeApiMerchantMetricListReq(unibeeApiMerchantMetricListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricListPost`: MerchantMetricListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricListReq** | [**UnibeeApiMerchantMetricListReq**](UnibeeApiMerchantMetricListReq.md) |  | 

### Return type

[**MerchantMetricListGet200Response**](MerchantMetricListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricNewPost

> MerchantMetricDeletePost200Response MetricNewPost(ctx).UnibeeApiMerchantMetricNewReq(unibeeApiMerchantMetricNewReq).Execute()

New Merchant Metric

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
	unibeeApiMerchantMetricNewReq := *openapiclient.NewUnibeeApiMerchantMetricNewReq("Code_example", "MetricName_example") // UnibeeApiMerchantMetricNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricNewPost(context.Background()).UnibeeApiMerchantMetricNewReq(unibeeApiMerchantMetricNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricNewPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricNewReq** | [**UnibeeApiMerchantMetricNewReq**](UnibeeApiMerchantMetricNewReq.md) |  | 

### Return type

[**MerchantMetricDeletePost200Response**](MerchantMetricDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricPlanLimitDeletePost

> MerchantMetricPlanLimitDeletePost200Response MetricPlanLimitDeletePost(ctx).UnibeeApiMerchantMetricDeletePlanLimitReq(unibeeApiMerchantMetricDeletePlanLimitReq).Execute()

Delete Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricDeletePlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricDeletePlanLimitReq(int64(123)) // UnibeeApiMerchantMetricDeletePlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricPlanLimitDeletePost(context.Background()).UnibeeApiMerchantMetricDeletePlanLimitReq(unibeeApiMerchantMetricDeletePlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricPlanLimitDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricPlanLimitDeletePost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricPlanLimitDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricPlanLimitDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDeletePlanLimitReq** | [**UnibeeApiMerchantMetricDeletePlanLimitReq**](UnibeeApiMerchantMetricDeletePlanLimitReq.md) |  | 

### Return type

[**MerchantMetricPlanLimitDeletePost200Response**](MerchantMetricPlanLimitDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricPlanLimitEditPost

> MerchantMetricPlanLimitDeletePost200Response MetricPlanLimitEditPost(ctx).UnibeeApiMerchantMetricEditPlanLimitReq(unibeeApiMerchantMetricEditPlanLimitReq).Execute()

Edit Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricEditPlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricEditPlanLimitReq(int64(123), int64(123)) // UnibeeApiMerchantMetricEditPlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricPlanLimitEditPost(context.Background()).UnibeeApiMerchantMetricEditPlanLimitReq(unibeeApiMerchantMetricEditPlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricPlanLimitEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricPlanLimitEditPost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricPlanLimitEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricPlanLimitEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricEditPlanLimitReq** | [**UnibeeApiMerchantMetricEditPlanLimitReq**](UnibeeApiMerchantMetricEditPlanLimitReq.md) |  | 

### Return type

[**MerchantMetricPlanLimitDeletePost200Response**](MerchantMetricPlanLimitDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetricPlanLimitNewPost

> MerchantMetricPlanLimitDeletePost200Response MetricPlanLimitNewPost(ctx).UnibeeApiMerchantMetricNewPlanLimitReq(unibeeApiMerchantMetricNewPlanLimitReq).Execute()

New Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricNewPlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricNewPlanLimitReq(int64(123), int64(123), int64(123)) // UnibeeApiMerchantMetricNewPlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Metric.MetricPlanLimitNewPost(context.Background()).UnibeeApiMerchantMetricNewPlanLimitReq(unibeeApiMerchantMetricNewPlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Metric.MetricPlanLimitNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MetricPlanLimitNewPost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Metric.MetricPlanLimitNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetricPlanLimitNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricNewPlanLimitReq** | [**UnibeeApiMerchantMetricNewPlanLimitReq**](UnibeeApiMerchantMetricNewPlanLimitReq.md) |  | 

### Return type

[**MerchantMetricPlanLimitDeletePost200Response**](MerchantMetricPlanLimitDeletePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

