# \MerchantMetricAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantMetricDeletePost**](MerchantMetricAPI.md#MerchantMetricDeletePost) | **Post** /merchant/metric/delete | Delete Merchant Metric
[**MerchantMetricDetailPost**](MerchantMetricAPI.md#MerchantMetricDetailPost) | **Post** /merchant/metric/detail | Merchant Metric Detail
[**MerchantMetricEditPost**](MerchantMetricAPI.md#MerchantMetricEditPost) | **Post** /merchant/metric/edit | Edit Merchant Metric
[**MerchantMetricEventDeletePost**](MerchantMetricAPI.md#MerchantMetricEventDeletePost) | **Post** /merchant/metric/event/delete | Del Merchant Metric Event
[**MerchantMetricEventNewPost**](MerchantMetricAPI.md#MerchantMetricEventNewPost) | **Post** /merchant/metric/event/new | Merchant Metric Event
[**MerchantMetricListGet**](MerchantMetricAPI.md#MerchantMetricListGet) | **Get** /merchant/metric/list | Merchant Metric list
[**MerchantMetricNewPost**](MerchantMetricAPI.md#MerchantMetricNewPost) | **Post** /merchant/metric/new | New Merchant Metric
[**MerchantMetricPlanLimitDeletePost**](MerchantMetricAPI.md#MerchantMetricPlanLimitDeletePost) | **Post** /merchant/metric/plan/limit/delete | Delete Merchant Metric Plan TotalLimit
[**MerchantMetricPlanLimitEditPost**](MerchantMetricAPI.md#MerchantMetricPlanLimitEditPost) | **Post** /merchant/metric/plan/limit/edit | Edit Merchant Metric Plan TotalLimit
[**MerchantMetricPlanLimitNewPost**](MerchantMetricAPI.md#MerchantMetricPlanLimitNewPost) | **Post** /merchant/metric/plan/limit/new | New Merchant Metric Plan TotalLimit



## MerchantMetricDeletePost

> MerchantMetricDeletePost200Response MerchantMetricDeletePost(ctx).UnibeeApiMerchantMetricDeleteReq(unibeeApiMerchantMetricDeleteReq).Execute()

Delete Merchant Metric

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
	unibeeApiMerchantMetricDeleteReq := *openapiclient.NewUnibeeApiMerchantMetricDeleteReq(int64(123)) // UnibeeApiMerchantMetricDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricDeletePost(context.Background()).UnibeeApiMerchantMetricDeleteReq(unibeeApiMerchantMetricDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricDeletePost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricDeletePostRequest struct via the builder pattern


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


## MerchantMetricDetailPost

> MerchantMetricDeletePost200Response MerchantMetricDetailPost(ctx).UnibeeApiMerchantMetricDetailReq(unibeeApiMerchantMetricDetailReq).Execute()

Merchant Metric Detail

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
	unibeeApiMerchantMetricDetailReq := *openapiclient.NewUnibeeApiMerchantMetricDetailReq(int64(123)) // UnibeeApiMerchantMetricDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricDetailPost(context.Background()).UnibeeApiMerchantMetricDetailReq(unibeeApiMerchantMetricDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricDetailPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricDetailPostRequest struct via the builder pattern


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


## MerchantMetricEditPost

> MerchantMetricDeletePost200Response MerchantMetricEditPost(ctx).UnibeeApiMerchantMetricEditReq(unibeeApiMerchantMetricEditReq).Execute()

Edit Merchant Metric

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
	unibeeApiMerchantMetricEditReq := *openapiclient.NewUnibeeApiMerchantMetricEditReq(int64(123), "MetricName_example") // UnibeeApiMerchantMetricEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricEditPost(context.Background()).UnibeeApiMerchantMetricEditReq(unibeeApiMerchantMetricEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricEditPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricEditPostRequest struct via the builder pattern


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


## MerchantMetricEventDeletePost

> MerchantAuthSsoLoginOTPPost200Response MerchantMetricEventDeletePost(ctx).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()

Del Merchant Metric Event

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
	unibeeApiMerchantMetricDeleteEventReq := *openapiclient.NewUnibeeApiMerchantMetricDeleteEventReq("ExternalEventId_example", "ExternalUserId_example", "MetricCode_example") // UnibeeApiMerchantMetricDeleteEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricEventDeletePost(context.Background()).UnibeeApiMerchantMetricDeleteEventReq(unibeeApiMerchantMetricDeleteEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricEventDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricEventDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricEventDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricEventDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMetricDeleteEventReq** | [**UnibeeApiMerchantMetricDeleteEventReq**](UnibeeApiMerchantMetricDeleteEventReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantMetricEventNewPost

> MerchantMetricEventNewPost200Response MerchantMetricEventNewPost(ctx).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()

Merchant Metric Event

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
	unibeeApiMerchantMetricNewEventReq := *openapiclient.NewUnibeeApiMerchantMetricNewEventReq("ExternalEventId_example", "ExternalUserId_example", "MetricCode_example") // UnibeeApiMerchantMetricNewEventReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricEventNewPost(context.Background()).UnibeeApiMerchantMetricNewEventReq(unibeeApiMerchantMetricNewEventReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricEventNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricEventNewPost`: MerchantMetricEventNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricEventNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricEventNewPostRequest struct via the builder pattern


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


## MerchantMetricListGet

> MerchantMetricListGet200Response MerchantMetricListGet(ctx).Execute()

Merchant Metric list

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricListGet`: MerchantMetricListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricListGetRequest struct via the builder pattern


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


## MerchantMetricNewPost

> MerchantMetricDeletePost200Response MerchantMetricNewPost(ctx).UnibeeApiMerchantMetricNewReq(unibeeApiMerchantMetricNewReq).Execute()

New Merchant Metric

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
	unibeeApiMerchantMetricNewReq := *openapiclient.NewUnibeeApiMerchantMetricNewReq("Code_example", "MetricName_example") // UnibeeApiMerchantMetricNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricNewPost(context.Background()).UnibeeApiMerchantMetricNewReq(unibeeApiMerchantMetricNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricNewPost`: MerchantMetricDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricNewPostRequest struct via the builder pattern


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


## MerchantMetricPlanLimitDeletePost

> MerchantMetricPlanLimitDeletePost200Response MerchantMetricPlanLimitDeletePost(ctx).UnibeeApiMerchantMetricDeletePlanLimitReq(unibeeApiMerchantMetricDeletePlanLimitReq).Execute()

Delete Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricDeletePlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricDeletePlanLimitReq(int64(123)) // UnibeeApiMerchantMetricDeletePlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricPlanLimitDeletePost(context.Background()).UnibeeApiMerchantMetricDeletePlanLimitReq(unibeeApiMerchantMetricDeletePlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricPlanLimitDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricPlanLimitDeletePost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricPlanLimitDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricPlanLimitDeletePostRequest struct via the builder pattern


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


## MerchantMetricPlanLimitEditPost

> MerchantMetricPlanLimitDeletePost200Response MerchantMetricPlanLimitEditPost(ctx).UnibeeApiMerchantMetricEditPlanLimitReq(unibeeApiMerchantMetricEditPlanLimitReq).Execute()

Edit Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricEditPlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricEditPlanLimitReq(int64(123), int64(123)) // UnibeeApiMerchantMetricEditPlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricPlanLimitEditPost(context.Background()).UnibeeApiMerchantMetricEditPlanLimitReq(unibeeApiMerchantMetricEditPlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricPlanLimitEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricPlanLimitEditPost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricPlanLimitEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricPlanLimitEditPostRequest struct via the builder pattern


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


## MerchantMetricPlanLimitNewPost

> MerchantMetricPlanLimitDeletePost200Response MerchantMetricPlanLimitNewPost(ctx).UnibeeApiMerchantMetricNewPlanLimitReq(unibeeApiMerchantMetricNewPlanLimitReq).Execute()

New Merchant Metric Plan TotalLimit

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
	unibeeApiMerchantMetricNewPlanLimitReq := *openapiclient.NewUnibeeApiMerchantMetricNewPlanLimitReq(int64(123), int64(123), int64(123)) // UnibeeApiMerchantMetricNewPlanLimitReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMetricAPI.MerchantMetricPlanLimitNewPost(context.Background()).UnibeeApiMerchantMetricNewPlanLimitReq(unibeeApiMerchantMetricNewPlanLimitReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMetricAPI.MerchantMetricPlanLimitNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMetricPlanLimitNewPost`: MerchantMetricPlanLimitDeletePost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMetricAPI.MerchantMetricPlanLimitNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMetricPlanLimitNewPostRequest struct via the builder pattern


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

