# \MerchantSubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSubscriptionAddNewTrialStartPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionAddNewTrialStartPost) | **Post** /merchant/subscription/add_new_trial_start | Merchant Edit Subscription-add appendTrialEndHour For Free
[**MerchantSubscriptionCancelAtPeriodEndPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_at_period_end | Merchant Edit Subscription-Set Cancel Ad Period End
[**MerchantSubscriptionCancelLastCancelAtPeriodEndPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionCancelLastCancelAtPeriodEndPost) | **Post** /merchant/subscription/cancel_last_cancel_at_period_end | Merchant Edit Subscription-Cancel Last CancelAtPeriod
[**MerchantSubscriptionCancelPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionCancelPost) | **Post** /merchant/subscription/cancel | Merchant Cancel Subscription Immediately (Will Not Generate Proration Invoice)
[**MerchantSubscriptionDetailGet**](MerchantSubscriptionAPI.md#MerchantSubscriptionDetailGet) | **Get** /merchant/subscription/detail | Subscription Detail
[**MerchantSubscriptionDetailPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionDetailPost) | **Post** /merchant/subscription/detail | Subscription Detail
[**MerchantSubscriptionListGet**](MerchantSubscriptionAPI.md#MerchantSubscriptionListGet) | **Get** /merchant/subscription/list | Subscription List
[**MerchantSubscriptionListPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionListPost) | **Post** /merchant/subscription/list | Subscription List
[**MerchantSubscriptionResumePost**](MerchantSubscriptionAPI.md#MerchantSubscriptionResumePost) | **Post** /merchant/subscription/resume | Merchant Edit Subscription-Resume
[**MerchantSubscriptionSuspendPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionSuspendPost) | **Post** /merchant/subscription/suspend | Merchant Edit Subscription-Stop
[**MerchantSubscriptionUpdatePreviewPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionUpdatePreviewPost) | **Post** /merchant/subscription/update_preview | Merchant Update Subscription Preview
[**MerchantSubscriptionUpdateSubmitPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionUpdateSubmitPost) | **Post** /merchant/subscription/update_submit | Merchant Update Subscription Submit
[**MerchantSubscriptionUserSubscriptionDetailGet**](MerchantSubscriptionAPI.md#MerchantSubscriptionUserSubscriptionDetailGet) | **Get** /merchant/subscription/user_subscription_detail | Subscription Detail
[**MerchantSubscriptionUserSubscriptionDetailPost**](MerchantSubscriptionAPI.md#MerchantSubscriptionUserSubscriptionDetailPost) | **Post** /merchant/subscription/user_subscription_detail | Subscription Detail



## MerchantSubscriptionAddNewTrialStartPost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionAddNewTrialStartPost(ctx).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()

Merchant Edit Subscription-add appendTrialEndHour For Free

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
	unibeeApiMerchantSubscriptionAddNewTrialStartReq := *openapiclient.NewUnibeeApiMerchantSubscriptionAddNewTrialStartReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionAddNewTrialStartReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionAddNewTrialStartPost(context.Background()).UnibeeApiMerchantSubscriptionAddNewTrialStartReq(unibeeApiMerchantSubscriptionAddNewTrialStartReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionAddNewTrialStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionAddNewTrialStartPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionAddNewTrialStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionAddNewTrialStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionAddNewTrialStartReq** | [**UnibeeApiMerchantSubscriptionAddNewTrialStartReq**](UnibeeApiMerchantSubscriptionAddNewTrialStartReq.md) |  | 

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


## MerchantSubscriptionCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()

Merchant Edit Subscription-Set Cancel Ad Period End

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
	unibeeApiMerchantSubscriptionCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq.md) |  | 

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


## MerchantSubscriptionCancelLastCancelAtPeriodEndPost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionCancelLastCancelAtPeriodEndPost(ctx).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()

Merchant Edit Subscription-Cancel Last CancelAtPeriod

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
	unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionCancelLastCancelAtPeriodEndPost(context.Background()).UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionCancelLastCancelAtPeriodEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionCancelLastCancelAtPeriodEndPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionCancelLastCancelAtPeriodEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionCancelLastCancelAtPeriodEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq** | [**UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq**](UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq.md) |  | 

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


## MerchantSubscriptionCancelPost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionCancelPost(ctx).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()

Merchant Cancel Subscription Immediately (Will Not Generate Proration Invoice)

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
	unibeeApiMerchantSubscriptionCancelReq := *openapiclient.NewUnibeeApiMerchantSubscriptionCancelReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionCancelReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionCancelPost(context.Background()).UnibeeApiMerchantSubscriptionCancelReq(unibeeApiMerchantSubscriptionCancelReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionCancelPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionCancelPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionCancelReq** | [**UnibeeApiMerchantSubscriptionCancelReq**](UnibeeApiMerchantSubscriptionCancelReq.md) |  | 

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


## MerchantSubscriptionDetailGet

> MerchantSubscriptionDetailGet200Response MerchantSubscriptionDetailGet(ctx).SubscriptionId(subscriptionId).Execute()

Subscription Detail

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
	subscriptionId := "subscriptionId_example" // string | SubscriptionId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionDetailGet(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionDetailGet`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionDetailPost

> MerchantSubscriptionDetailGet200Response MerchantSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()

Subscription Detail

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
	unibeeApiMerchantSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionDetailReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionDetailReq(unibeeApiMerchantSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionDetailPost`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionListGet

> MerchantSubscriptionListGet200Response MerchantSubscriptionListGet(ctx).UserId(userId).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Subscription List

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
	userId := int64(789) // int64 | UserId (optional)
	status := []int32{int32(123)} // []int32 | Filter, Default All，Status，0-Init | 1-Create｜2-Active｜3-Suspend | 4-Cancel | 5-Expire (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionListGet(context.Background()).UserId(userId).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionListGet`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 
 **status** | **[]int32** | Filter, Default All，Status，0-Init | 1-Create｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start WIth 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionListPost

> MerchantSubscriptionListGet200Response MerchantSubscriptionListPost(ctx).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()

Subscription List

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
	unibeeApiMerchantSubscriptionListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionListReq() // UnibeeApiMerchantSubscriptionListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionListPost(context.Background()).UnibeeApiMerchantSubscriptionListReq(unibeeApiMerchantSubscriptionListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionListPost`: MerchantSubscriptionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionListReq** | [**UnibeeApiMerchantSubscriptionListReq**](UnibeeApiMerchantSubscriptionListReq.md) |  | 

### Return type

[**MerchantSubscriptionListGet200Response**](MerchantSubscriptionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionResumePost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionResumePost(ctx).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()

Merchant Edit Subscription-Resume

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
	unibeeApiMerchantSubscriptionResumeReq := *openapiclient.NewUnibeeApiMerchantSubscriptionResumeReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionResumeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionResumePost(context.Background()).UnibeeApiMerchantSubscriptionResumeReq(unibeeApiMerchantSubscriptionResumeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionResumePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionResumePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionResumePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionResumePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionResumeReq** | [**UnibeeApiMerchantSubscriptionResumeReq**](UnibeeApiMerchantSubscriptionResumeReq.md) |  | 

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


## MerchantSubscriptionSuspendPost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionSuspendPost(ctx).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()

Merchant Edit Subscription-Stop

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
	unibeeApiMerchantSubscriptionSuspendReq := *openapiclient.NewUnibeeApiMerchantSubscriptionSuspendReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionSuspendReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionSuspendPost(context.Background()).UnibeeApiMerchantSubscriptionSuspendReq(unibeeApiMerchantSubscriptionSuspendReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionSuspendPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionSuspendPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionSuspendPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionSuspendPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionSuspendReq** | [**UnibeeApiMerchantSubscriptionSuspendReq**](UnibeeApiMerchantSubscriptionSuspendReq.md) |  | 

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


## MerchantSubscriptionUpdatePreviewPost

> MerchantSubscriptionUpdatePreviewPost200Response MerchantSubscriptionUpdatePreviewPost(ctx).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()

Merchant Update Subscription Preview

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
	unibeeApiMerchantSubscriptionUpdatePreviewReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdatePreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionUpdatePreviewPost(context.Background()).UnibeeApiMerchantSubscriptionUpdatePreviewReq(unibeeApiMerchantSubscriptionUpdatePreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionUpdatePreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionUpdatePreviewPost`: MerchantSubscriptionUpdatePreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionUpdatePreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionUpdatePreviewPostRequest struct via the builder pattern


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


## MerchantSubscriptionUpdateSubmitPost

> MerchantSubscriptionUpdateSubmitPost200Response MerchantSubscriptionUpdateSubmitPost(ctx).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()

Merchant Update Subscription Submit

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
	unibeeApiMerchantSubscriptionUpdateReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUpdateReq("ConfirmCurrency_example", int64(123), int64(123), int64(123), "SubscriptionId_example") // UnibeeApiMerchantSubscriptionUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionUpdateSubmitPost(context.Background()).UnibeeApiMerchantSubscriptionUpdateReq(unibeeApiMerchantSubscriptionUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionUpdateSubmitPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionUpdateSubmitPost`: MerchantSubscriptionUpdateSubmitPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionUpdateSubmitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionUpdateSubmitPostRequest struct via the builder pattern


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


## MerchantSubscriptionUserSubscriptionDetailGet

> MerchantSubscriptionDetailGet200Response MerchantSubscriptionUserSubscriptionDetailGet(ctx).UserId(userId).Execute()

Subscription Detail

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
	userId := int64(789) // int64 | UserId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailGet(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionUserSubscriptionDetailGet`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionUserSubscriptionDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionUserSubscriptionDetailPost

> MerchantSubscriptionDetailGet200Response MerchantSubscriptionUserSubscriptionDetailPost(ctx).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()

Subscription Detail

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
	unibeeApiMerchantSubscriptionUserSubscriptionDetailReq := *openapiclient.NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(int64(123)) // UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailPost(context.Background()).UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq(unibeeApiMerchantSubscriptionUserSubscriptionDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionUserSubscriptionDetailPost`: MerchantSubscriptionDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionAPI.MerchantSubscriptionUserSubscriptionDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionUserSubscriptionDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionUserSubscriptionDetailReq** | [**UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq**](UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq.md) |  | 

### Return type

[**MerchantSubscriptionDetailGet200Response**](MerchantSubscriptionDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

