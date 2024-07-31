# \Discount

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscountActivatePost**](Discount.md#DiscountActivatePost) | **Post** /merchant/discount/activate | ActivateDiscountCode
[**DiscountDeactivatePost**](Discount.md#DiscountDeactivatePost) | **Post** /merchant/discount/deactivate | DeactivateDiscountCode
[**DiscountDeletePost**](Discount.md#DiscountDeletePost) | **Post** /merchant/discount/delete | DeleteDiscountCode
[**DiscountDetailGet**](Discount.md#DiscountDetailGet) | **Get** /merchant/discount/detail | Merchant Discount Detail
[**DiscountDetailPost**](Discount.md#DiscountDetailPost) | **Post** /merchant/discount/detail | Merchant Discount Detail
[**DiscountEditPost**](Discount.md#DiscountEditPost) | **Post** /merchant/discount/edit | EditDiscountCode
[**DiscountListGet**](Discount.md#DiscountListGet) | **Get** /merchant/discount/list | DiscountCodeList
[**DiscountNewPost**](Discount.md#DiscountNewPost) | **Post** /merchant/discount/new | NewDiscountCode
[**DiscountPlanApplyPreviewPost**](Discount.md#DiscountPlanApplyPreviewPost) | **Post** /merchant/discount/plan_apply_preview | PlanApplyPreview
[**DiscountUserDiscountListGet**](Discount.md#DiscountUserDiscountListGet) | **Get** /merchant/discount/user_discount_list | UserDiscountCodeList



## DiscountActivatePost

> MerchantAuthSsoLoginOTPPost200Response DiscountActivatePost(ctx).UnibeeApiMerchantDiscountActivateReq(unibeeApiMerchantDiscountActivateReq).Execute()

ActivateDiscountCode



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
	unibeeApiMerchantDiscountActivateReq := *openapiclient.NewUnibeeApiMerchantDiscountActivateReq(int64(123)) // UnibeeApiMerchantDiscountActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountActivatePost(context.Background()).UnibeeApiMerchantDiscountActivateReq(unibeeApiMerchantDiscountActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountActivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountActivateReq** | [**UnibeeApiMerchantDiscountActivateReq**](UnibeeApiMerchantDiscountActivateReq.md) |  | 

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


## DiscountDeactivatePost

> MerchantAuthSsoLoginOTPPost200Response DiscountDeactivatePost(ctx).UnibeeApiMerchantDiscountDeactivateReq(unibeeApiMerchantDiscountDeactivateReq).Execute()

DeactivateDiscountCode



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
	unibeeApiMerchantDiscountDeactivateReq := *openapiclient.NewUnibeeApiMerchantDiscountDeactivateReq(int64(123)) // UnibeeApiMerchantDiscountDeactivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountDeactivatePost(context.Background()).UnibeeApiMerchantDiscountDeactivateReq(unibeeApiMerchantDiscountDeactivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountDeactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountDeactivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountDeactivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountDeactivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountDeactivateReq** | [**UnibeeApiMerchantDiscountDeactivateReq**](UnibeeApiMerchantDiscountDeactivateReq.md) |  | 

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


## DiscountDeletePost

> MerchantAuthSsoLoginOTPPost200Response DiscountDeletePost(ctx).UnibeeApiMerchantDiscountDeleteReq(unibeeApiMerchantDiscountDeleteReq).Execute()

DeleteDiscountCode



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
	unibeeApiMerchantDiscountDeleteReq := *openapiclient.NewUnibeeApiMerchantDiscountDeleteReq(int64(123)) // UnibeeApiMerchantDiscountDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountDeletePost(context.Background()).UnibeeApiMerchantDiscountDeleteReq(unibeeApiMerchantDiscountDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountDeleteReq** | [**UnibeeApiMerchantDiscountDeleteReq**](UnibeeApiMerchantDiscountDeleteReq.md) |  | 

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


## DiscountDetailGet

> MerchantDiscountDetailGet200Response DiscountDetailGet(ctx).Id(id).Execute()

Merchant Discount Detail

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
	id := int64(789) // int64 | The discount's Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountDetailGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountDetailGet`: MerchantDiscountDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The discount&#39;s Id | 

### Return type

[**MerchantDiscountDetailGet200Response**](MerchantDiscountDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountDetailPost

> MerchantDiscountDetailGet200Response DiscountDetailPost(ctx).UnibeeApiMerchantDiscountDetailReq(unibeeApiMerchantDiscountDetailReq).Execute()

Merchant Discount Detail

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
	unibeeApiMerchantDiscountDetailReq := *openapiclient.NewUnibeeApiMerchantDiscountDetailReq(int64(123)) // UnibeeApiMerchantDiscountDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountDetailPost(context.Background()).UnibeeApiMerchantDiscountDetailReq(unibeeApiMerchantDiscountDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountDetailPost`: MerchantDiscountDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountDetailReq** | [**UnibeeApiMerchantDiscountDetailReq**](UnibeeApiMerchantDiscountDetailReq.md) |  | 

### Return type

[**MerchantDiscountDetailGet200Response**](MerchantDiscountDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountEditPost

> MerchantDiscountEditPost200Response DiscountEditPost(ctx).UnibeeApiMerchantDiscountEditReq(unibeeApiMerchantDiscountEditReq).Execute()

EditDiscountCode



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
	unibeeApiMerchantDiscountEditReq := *openapiclient.NewUnibeeApiMerchantDiscountEditReq(int64(123)) // UnibeeApiMerchantDiscountEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountEditPost(context.Background()).UnibeeApiMerchantDiscountEditReq(unibeeApiMerchantDiscountEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountEditPost`: MerchantDiscountEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountEditReq** | [**UnibeeApiMerchantDiscountEditReq**](UnibeeApiMerchantDiscountEditReq.md) |  | 

### Return type

[**MerchantDiscountEditPost200Response**](MerchantDiscountEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountListGet

> MerchantDiscountListGet200Response DiscountListGet(ctx).DiscountType(discountType).BillingType(billingType).Status(status).Code(code).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

DiscountCodeList



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
	discountType := []int32{int32(123)} // []int32 | discount_type, 1-percentage, 2-fixed_amount (optional)
	billingType := []int32{int32(123)} // []int32 | billing_type, 1-one-time, 2-recurring (optional)
	status := []int32{int32(123)} // []int32 | status, 1-editable, 2-active, 3-deactive, 4-expire (optional)
	code := "code_example" // string | Filter Code (optional)
	currency := "currency_example" // string | Filter Currency (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountListGet(context.Background()).DiscountType(discountType).BillingType(billingType).Status(status).Code(code).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountListGet`: MerchantDiscountListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discountType** | **[]int32** | discount_type, 1-percentage, 2-fixed_amount | 
 **billingType** | **[]int32** | billing_type, 1-one-time, 2-recurring | 
 **status** | **[]int32** | status, 1-editable, 2-active, 3-deactive, 4-expire | 
 **code** | **string** | Filter Code | 
 **currency** | **string** | Filter Currency | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart | 
 **createTimeEnd** | **int64** | CreateTimeEnd | 

### Return type

[**MerchantDiscountListGet200Response**](MerchantDiscountListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountNewPost

> MerchantDiscountEditPost200Response DiscountNewPost(ctx).UnibeeApiMerchantDiscountNewReq(unibeeApiMerchantDiscountNewReq).Execute()

NewDiscountCode



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
	unibeeApiMerchantDiscountNewReq := *openapiclient.NewUnibeeApiMerchantDiscountNewReq(int32(123), "Code_example", int32(123), int64(123), int64(123)) // UnibeeApiMerchantDiscountNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountNewPost(context.Background()).UnibeeApiMerchantDiscountNewReq(unibeeApiMerchantDiscountNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountNewPost`: MerchantDiscountEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountNewReq** | [**UnibeeApiMerchantDiscountNewReq**](UnibeeApiMerchantDiscountNewReq.md) |  | 

### Return type

[**MerchantDiscountEditPost200Response**](MerchantDiscountEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountPlanApplyPreviewPost

> MerchantDiscountPlanApplyPreviewPost200Response DiscountPlanApplyPreviewPost(ctx).UnibeeApiMerchantDiscountPlanApplyPreviewReq(unibeeApiMerchantDiscountPlanApplyPreviewReq).Execute()

PlanApplyPreview



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
	unibeeApiMerchantDiscountPlanApplyPreviewReq := *openapiclient.NewUnibeeApiMerchantDiscountPlanApplyPreviewReq("Code_example") // UnibeeApiMerchantDiscountPlanApplyPreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountPlanApplyPreviewPost(context.Background()).UnibeeApiMerchantDiscountPlanApplyPreviewReq(unibeeApiMerchantDiscountPlanApplyPreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountPlanApplyPreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountPlanApplyPreviewPost`: MerchantDiscountPlanApplyPreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountPlanApplyPreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountPlanApplyPreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountPlanApplyPreviewReq** | [**UnibeeApiMerchantDiscountPlanApplyPreviewReq**](UnibeeApiMerchantDiscountPlanApplyPreviewReq.md) |  | 

### Return type

[**MerchantDiscountPlanApplyPreviewPost200Response**](MerchantDiscountPlanApplyPreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountUserDiscountListGet

> MerchantDiscountUserDiscountListGet200Response DiscountUserDiscountListGet(ctx).Id(id).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

UserDiscountCodeList



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
	id := int64(789) // int64 | The discount's Id
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountUserDiscountListGet(context.Background()).Id(id).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountUserDiscountListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountUserDiscountListGet`: MerchantDiscountUserDiscountListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountUserDiscountListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountUserDiscountListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The discount&#39;s Id | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart | 
 **createTimeEnd** | **int64** | CreateTimeEnd | 

### Return type

[**MerchantDiscountUserDiscountListGet200Response**](MerchantDiscountUserDiscountListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

