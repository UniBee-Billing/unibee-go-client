# \BatchDiscount

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscountBatchChildrenDetailGet**](BatchDiscount.md#DiscountBatchChildrenDetailGet) | **Get** /merchant/discount/batch/children/detail | Get Batch Child Code Detail
[**DiscountBatchChildrenDetailPost**](BatchDiscount.md#DiscountBatchChildrenDetailPost) | **Post** /merchant/discount/batch/children/detail | Get Batch Child Code Detail
[**DiscountBatchChildrenGeneratePost**](BatchDiscount.md#DiscountBatchChildrenGeneratePost) | **Post** /merchant/discount/batch/children/generate | Generate Batch Child Codes
[**DiscountBatchChildrenListGet**](BatchDiscount.md#DiscountBatchChildrenListGet) | **Get** /merchant/discount/batch/children/list | Get Batch Child Codes List
[**DiscountBatchTemplateDetailGet**](BatchDiscount.md#DiscountBatchTemplateDetailGet) | **Get** /merchant/discount/batch/template/detail | Get Batch Template Detail
[**DiscountBatchTemplateDetailPost**](BatchDiscount.md#DiscountBatchTemplateDetailPost) | **Post** /merchant/discount/batch/template/detail | Get Batch Template Detail
[**DiscountBatchTemplateEditPost**](BatchDiscount.md#DiscountBatchTemplateEditPost) | **Post** /merchant/discount/batch/template/edit | Edit Batch Discount Template
[**DiscountBatchTemplateListGet**](BatchDiscount.md#DiscountBatchTemplateListGet) | **Get** /merchant/discount/batch/template/list | Get Batch Template List
[**DiscountBatchTemplateNewPost**](BatchDiscount.md#DiscountBatchTemplateNewPost) | **Post** /merchant/discount/batch/template/new | Create Batch Discount Template
[**DiscountBatchTemplateQuantityIncrementPost**](BatchDiscount.md#DiscountBatchTemplateQuantityIncrementPost) | **Post** /merchant/discount/batch/template/quantity_increment | Increment Template Quantity



## DiscountBatchChildrenDetailGet

> MerchantDiscountBatchChildrenDetailGet200Response DiscountBatchChildrenDetailGet(ctx).Id(id).Execute()

Get Batch Child Code Detail



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
	id := int64(789) // int64 | The child code's Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchChildrenDetailGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchChildrenDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchChildrenDetailGet`: MerchantDiscountBatchChildrenDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchChildrenDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchChildrenDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The child code&#39;s Id | 

### Return type

[**MerchantDiscountBatchChildrenDetailGet200Response**](MerchantDiscountBatchChildrenDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchChildrenDetailPost

> MerchantDiscountBatchChildrenDetailGet200Response DiscountBatchChildrenDetailPost(ctx).UnibeeApiMerchantDiscountBatchChildrenDetailReq(unibeeApiMerchantDiscountBatchChildrenDetailReq).Execute()

Get Batch Child Code Detail



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
	unibeeApiMerchantDiscountBatchChildrenDetailReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchChildrenDetailReq(int64(123)) // UnibeeApiMerchantDiscountBatchChildrenDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchChildrenDetailPost(context.Background()).UnibeeApiMerchantDiscountBatchChildrenDetailReq(unibeeApiMerchantDiscountBatchChildrenDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchChildrenDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchChildrenDetailPost`: MerchantDiscountBatchChildrenDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchChildrenDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchChildrenDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchChildrenDetailReq** | [**UnibeeApiMerchantDiscountBatchChildrenDetailReq**](UnibeeApiMerchantDiscountBatchChildrenDetailReq.md) |  | 

### Return type

[**MerchantDiscountBatchChildrenDetailGet200Response**](MerchantDiscountBatchChildrenDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchChildrenGeneratePost

> MerchantDiscountBatchChildrenGeneratePost200Response DiscountBatchChildrenGeneratePost(ctx).UnibeeApiMerchantDiscountBatchChildrenGenerateReq(unibeeApiMerchantDiscountBatchChildrenGenerateReq).Execute()

Generate Batch Child Codes



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
	unibeeApiMerchantDiscountBatchChildrenGenerateReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchChildrenGenerateReq() // UnibeeApiMerchantDiscountBatchChildrenGenerateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchChildrenGeneratePost(context.Background()).UnibeeApiMerchantDiscountBatchChildrenGenerateReq(unibeeApiMerchantDiscountBatchChildrenGenerateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchChildrenGeneratePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchChildrenGeneratePost`: MerchantDiscountBatchChildrenGeneratePost200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchChildrenGeneratePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchChildrenGeneratePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchChildrenGenerateReq** | [**UnibeeApiMerchantDiscountBatchChildrenGenerateReq**](UnibeeApiMerchantDiscountBatchChildrenGenerateReq.md) |  | 

### Return type

[**MerchantDiscountBatchChildrenGeneratePost200Response**](MerchantDiscountBatchChildrenGeneratePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchChildrenListGet

> MerchantDiscountBatchChildrenListGet200Response DiscountBatchChildrenListGet(ctx).TemplateId(templateId).Code(code).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Batch Child Codes List



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
	templateId := int64(789) // int64 | The template's Id
	code := "code_example" // string | Filter by child code, fuzzy search (optional)
	sortField := "sortField_example" // string | Sort field, gmt_create|gmt_modify, default gmt_create (optional)
	sortType := "sortType_example" // string | Sort type, asc|desc, default desc (optional)
	page := int32(56) // int32 | Page number, start from 0 (optional)
	count := int32(56) // int32 | Count per page (optional)
	createTimeStart := int64(789) // int64 | Filter by create time start, UTC timestamp in seconds (optional)
	createTimeEnd := int64(789) // int64 | Filter by create time end, UTC timestamp in seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchChildrenListGet(context.Background()).TemplateId(templateId).Code(code).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchChildrenListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchChildrenListGet`: MerchantDiscountBatchChildrenListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchChildrenListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchChildrenListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateId** | **int64** | The template&#39;s Id | 
 **code** | **string** | Filter by child code, fuzzy search | 
 **sortField** | **string** | Sort field, gmt_create|gmt_modify, default gmt_create | 
 **sortType** | **string** | Sort type, asc|desc, default desc | 
 **page** | **int32** | Page number, start from 0 | 
 **count** | **int32** | Count per page | 
 **createTimeStart** | **int64** | Filter by create time start, UTC timestamp in seconds | 
 **createTimeEnd** | **int64** | Filter by create time end, UTC timestamp in seconds | 

### Return type

[**MerchantDiscountBatchChildrenListGet200Response**](MerchantDiscountBatchChildrenListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateDetailGet

> MerchantDiscountBatchTemplateDetailGet200Response DiscountBatchTemplateDetailGet(ctx).Id(id).Execute()

Get Batch Template Detail



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
	id := int64(789) // int64 | The template's Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateDetailGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateDetailGet`: MerchantDiscountBatchTemplateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The template&#39;s Id | 

### Return type

[**MerchantDiscountBatchTemplateDetailGet200Response**](MerchantDiscountBatchTemplateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateDetailPost

> MerchantDiscountBatchTemplateDetailGet200Response DiscountBatchTemplateDetailPost(ctx).UnibeeApiMerchantDiscountBatchTemplateDetailReq(unibeeApiMerchantDiscountBatchTemplateDetailReq).Execute()

Get Batch Template Detail



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
	unibeeApiMerchantDiscountBatchTemplateDetailReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchTemplateDetailReq(int64(123)) // UnibeeApiMerchantDiscountBatchTemplateDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateDetailPost(context.Background()).UnibeeApiMerchantDiscountBatchTemplateDetailReq(unibeeApiMerchantDiscountBatchTemplateDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateDetailPost`: MerchantDiscountBatchTemplateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchTemplateDetailReq** | [**UnibeeApiMerchantDiscountBatchTemplateDetailReq**](UnibeeApiMerchantDiscountBatchTemplateDetailReq.md) |  | 

### Return type

[**MerchantDiscountBatchTemplateDetailGet200Response**](MerchantDiscountBatchTemplateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateEditPost

> MerchantDiscountBatchTemplateDetailGet200Response DiscountBatchTemplateEditPost(ctx).UnibeeApiMerchantDiscountBatchTemplateEditReq(unibeeApiMerchantDiscountBatchTemplateEditReq).Execute()

Edit Batch Discount Template



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
	unibeeApiMerchantDiscountBatchTemplateEditReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchTemplateEditReq("CodePrefix_example", int64(123)) // UnibeeApiMerchantDiscountBatchTemplateEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateEditPost(context.Background()).UnibeeApiMerchantDiscountBatchTemplateEditReq(unibeeApiMerchantDiscountBatchTemplateEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateEditPost`: MerchantDiscountBatchTemplateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchTemplateEditReq** | [**UnibeeApiMerchantDiscountBatchTemplateEditReq**](UnibeeApiMerchantDiscountBatchTemplateEditReq.md) |  | 

### Return type

[**MerchantDiscountBatchTemplateDetailGet200Response**](MerchantDiscountBatchTemplateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateListGet

> MerchantDiscountBatchTemplateListGet200Response DiscountBatchTemplateListGet(ctx).DiscountType(discountType).BillingType(billingType).Status(status).CodePrefix(codePrefix).SearchKey(searchKey).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Batch Template List



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
	discountType := []int32{int32(123)} // []int32 | Filter by discount_type, 1-percentage, 2-fixed_amount (optional)
	billingType := []int32{int32(123)} // []int32 | Filter by billing_type, 1-one-time, 2-recurring (optional)
	status := []int32{int32(123)} // []int32 | Filter by status, 1-editable, 2-active, 3-deactive, 4-expire, 10-archive (optional)
	codePrefix := "codePrefix_example" // string | Filter by code prefix (optional)
	searchKey := "searchKey_example" // string | Search by code prefix or name (optional)
	currency := "currency_example" // string | Filter by currency (optional)
	sortField := "sortField_example" // string | Sort field, gmt_create|gmt_modify, default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort type, asc|desc, default desc (optional)
	page := int32(56) // int32 | Page number, start from 0 (optional)
	count := int32(56) // int32 | Count per page (optional)
	createTimeStart := int64(789) // int64 | Filter by create time start, UTC timestamp in seconds (optional)
	createTimeEnd := int64(789) // int64 | Filter by create time end, UTC timestamp in seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateListGet(context.Background()).DiscountType(discountType).BillingType(billingType).Status(status).CodePrefix(codePrefix).SearchKey(searchKey).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateListGet`: MerchantDiscountBatchTemplateListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **discountType** | **[]int32** | Filter by discount_type, 1-percentage, 2-fixed_amount | 
 **billingType** | **[]int32** | Filter by billing_type, 1-one-time, 2-recurring | 
 **status** | **[]int32** | Filter by status, 1-editable, 2-active, 3-deactive, 4-expire, 10-archive | 
 **codePrefix** | **string** | Filter by code prefix | 
 **searchKey** | **string** | Search by code prefix or name | 
 **currency** | **string** | Filter by currency | 
 **sortField** | **string** | Sort field, gmt_create|gmt_modify, default gmt_modify | 
 **sortType** | **string** | Sort type, asc|desc, default desc | 
 **page** | **int32** | Page number, start from 0 | 
 **count** | **int32** | Count per page | 
 **createTimeStart** | **int64** | Filter by create time start, UTC timestamp in seconds | 
 **createTimeEnd** | **int64** | Filter by create time end, UTC timestamp in seconds | 

### Return type

[**MerchantDiscountBatchTemplateListGet200Response**](MerchantDiscountBatchTemplateListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateNewPost

> MerchantDiscountBatchTemplateDetailGet200Response DiscountBatchTemplateNewPost(ctx).UnibeeApiMerchantDiscountBatchTemplateNewReq(unibeeApiMerchantDiscountBatchTemplateNewReq).Execute()

Create Batch Discount Template



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
	unibeeApiMerchantDiscountBatchTemplateNewReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchTemplateNewReq(int32(123), "CodePrefix_example", int32(123), int32(123), int32(123), int32(123)) // UnibeeApiMerchantDiscountBatchTemplateNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateNewPost(context.Background()).UnibeeApiMerchantDiscountBatchTemplateNewReq(unibeeApiMerchantDiscountBatchTemplateNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateNewPost`: MerchantDiscountBatchTemplateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchTemplateNewReq** | [**UnibeeApiMerchantDiscountBatchTemplateNewReq**](UnibeeApiMerchantDiscountBatchTemplateNewReq.md) |  | 

### Return type

[**MerchantDiscountBatchTemplateDetailGet200Response**](MerchantDiscountBatchTemplateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountBatchTemplateQuantityIncrementPost

> MerchantDiscountBatchTemplateDetailGet200Response DiscountBatchTemplateQuantityIncrementPost(ctx).UnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq(unibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq).Execute()

Increment Template Quantity



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
	unibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq := *openapiclient.NewUnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq(int64(123), int64(123)) // UnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchDiscount.DiscountBatchTemplateQuantityIncrementPost(context.Background()).UnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq(unibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchDiscount.DiscountBatchTemplateQuantityIncrementPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountBatchTemplateQuantityIncrementPost`: MerchantDiscountBatchTemplateDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `BatchDiscount.DiscountBatchTemplateQuantityIncrementPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountBatchTemplateQuantityIncrementPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq** | [**UnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq**](UnibeeApiMerchantDiscountBatchTemplateQuantityIncrementReq.md) |  | 

### Return type

[**MerchantDiscountBatchTemplateDetailGet200Response**](MerchantDiscountBatchTemplateDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

