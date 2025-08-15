# \Plan

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlanActivatePost**](Plan.md#PlanActivatePost) | **Post** /merchant/plan/activate | Activate Plan
[**PlanAddonsBindingPost**](Plan.md#PlanAddonsBindingPost) | **Post** /merchant/plan/addons_binding | Addon Binding
[**PlanArchivePost**](Plan.md#PlanArchivePost) | **Post** /merchant/plan/archive | Archive Plan
[**PlanCopyPost**](Plan.md#PlanCopyPost) | **Post** /merchant/plan/copy | Copy Plan
[**PlanDeletePost**](Plan.md#PlanDeletePost) | **Post** /merchant/plan/delete | Delete Plan
[**PlanDetailGet**](Plan.md#PlanDetailGet) | **Get** /merchant/plan/detail | Plan Detail
[**PlanDetailPost**](Plan.md#PlanDetailPost) | **Post** /merchant/plan/detail | Plan Detail
[**PlanEditPost**](Plan.md#PlanEditPost) | **Post** /merchant/plan/edit | Edit Plan
[**PlanListGet**](Plan.md#PlanListGet) | **Get** /merchant/plan/list | Get Plan List
[**PlanListPost**](Plan.md#PlanListPost) | **Post** /merchant/plan/list | Get Plan List
[**PlanNewPost**](Plan.md#PlanNewPost) | **Post** /merchant/plan/new | Create Plan
[**PlanPublishPost**](Plan.md#PlanPublishPost) | **Post** /merchant/plan/publish | Publish Plan
[**PlanUnpublishedPost**](Plan.md#PlanUnpublishedPost) | **Post** /merchant/plan/unpublished | UnPublish Plan



## PlanActivatePost

> MerchantAuthSsoClearTotpPost200Response PlanActivatePost(ctx).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()

Activate Plan

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
	unibeeApiMerchantPlanActivateReq := *openapiclient.NewUnibeeApiMerchantPlanActivateReq() // UnibeeApiMerchantPlanActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanActivatePost(context.Background()).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanActivatePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanActivateReq** | [**UnibeeApiMerchantPlanActivateReq**](UnibeeApiMerchantPlanActivateReq.md) |  | 

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


## PlanAddonsBindingPost

> MerchantPlanAddonsBindingPost200Response PlanAddonsBindingPost(ctx).UnibeeApiMerchantPlanAddonsBindingReq(unibeeApiMerchantPlanAddonsBindingReq).Execute()

Addon Binding

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
	unibeeApiMerchantPlanAddonsBindingReq := *openapiclient.NewUnibeeApiMerchantPlanAddonsBindingReq(int64(123), []int64{int64(123)}, []int64{int64(123)}, int64(123)) // UnibeeApiMerchantPlanAddonsBindingReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanAddonsBindingPost(context.Background()).UnibeeApiMerchantPlanAddonsBindingReq(unibeeApiMerchantPlanAddonsBindingReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanAddonsBindingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanAddonsBindingPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanAddonsBindingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanAddonsBindingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanAddonsBindingReq** | [**UnibeeApiMerchantPlanAddonsBindingReq**](UnibeeApiMerchantPlanAddonsBindingReq.md) |  | 

### Return type

[**MerchantPlanAddonsBindingPost200Response**](MerchantPlanAddonsBindingPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanArchivePost

> MerchantAuthSsoClearTotpPost200Response PlanArchivePost(ctx).UnibeeApiMerchantPlanArchiveReq(unibeeApiMerchantPlanArchiveReq).Execute()

Archive Plan

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
	unibeeApiMerchantPlanArchiveReq := *openapiclient.NewUnibeeApiMerchantPlanArchiveReq() // UnibeeApiMerchantPlanArchiveReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanArchivePost(context.Background()).UnibeeApiMerchantPlanArchiveReq(unibeeApiMerchantPlanArchiveReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanArchivePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanArchiveReq** | [**UnibeeApiMerchantPlanArchiveReq**](UnibeeApiMerchantPlanArchiveReq.md) |  | 

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


## PlanCopyPost

> MerchantPlanAddonsBindingPost200Response PlanCopyPost(ctx).UnibeeApiMerchantPlanCopyReq(unibeeApiMerchantPlanCopyReq).Execute()

Copy Plan

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
	unibeeApiMerchantPlanCopyReq := *openapiclient.NewUnibeeApiMerchantPlanCopyReq() // UnibeeApiMerchantPlanCopyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanCopyPost(context.Background()).UnibeeApiMerchantPlanCopyReq(unibeeApiMerchantPlanCopyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanCopyPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanCopyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanCopyReq** | [**UnibeeApiMerchantPlanCopyReq**](UnibeeApiMerchantPlanCopyReq.md) |  | 

### Return type

[**MerchantPlanAddonsBindingPost200Response**](MerchantPlanAddonsBindingPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanDeletePost

> MerchantAuthSsoClearTotpPost200Response PlanDeletePost(ctx).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()

Delete Plan

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
	unibeeApiMerchantPlanDeleteReq := *openapiclient.NewUnibeeApiMerchantPlanDeleteReq() // UnibeeApiMerchantPlanDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanDeletePost(context.Background()).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanDeletePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanDeleteReq** | [**UnibeeApiMerchantPlanDeleteReq**](UnibeeApiMerchantPlanDeleteReq.md) |  | 

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


## PlanDetailGet

> MerchantPlanDetailGet200Response PlanDetailGet(ctx).PlanId(planId).ExternalPlanId(externalPlanId).Execute()

Plan Detail

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
	planId := int64(789) // int64 | The Id of plan, either planId or externalPlanId should be set (optional)
	externalPlanId := "externalPlanId_example" // string | The ExternalId of plan, either externalPlanId or planId should be set (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanDetailGet(context.Background()).PlanId(planId).ExternalPlanId(externalPlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanDetailGet`: MerchantPlanDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planId** | **int64** | The Id of plan, either planId or externalPlanId should be set | 
 **externalPlanId** | **string** | The ExternalId of plan, either externalPlanId or planId should be set | 

### Return type

[**MerchantPlanDetailGet200Response**](MerchantPlanDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanDetailPost

> MerchantPlanDetailGet200Response PlanDetailPost(ctx).UnibeeApiMerchantPlanDetailReq(unibeeApiMerchantPlanDetailReq).Execute()

Plan Detail

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
	unibeeApiMerchantPlanDetailReq := *openapiclient.NewUnibeeApiMerchantPlanDetailReq() // UnibeeApiMerchantPlanDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanDetailPost(context.Background()).UnibeeApiMerchantPlanDetailReq(unibeeApiMerchantPlanDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanDetailPost`: MerchantPlanDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanDetailReq** | [**UnibeeApiMerchantPlanDetailReq**](UnibeeApiMerchantPlanDetailReq.md) |  | 

### Return type

[**MerchantPlanDetailGet200Response**](MerchantPlanDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanEditPost

> MerchantPlanAddonsBindingPost200Response PlanEditPost(ctx).UnibeeApiMerchantPlanEditReq(unibeeApiMerchantPlanEditReq).Execute()

Edit Plan



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
	unibeeApiMerchantPlanEditReq := *openapiclient.NewUnibeeApiMerchantPlanEditReq(int64(123)) // UnibeeApiMerchantPlanEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanEditPost(context.Background()).UnibeeApiMerchantPlanEditReq(unibeeApiMerchantPlanEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanEditPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanEditReq** | [**UnibeeApiMerchantPlanEditReq**](UnibeeApiMerchantPlanEditReq.md) |  | 

### Return type

[**MerchantPlanAddonsBindingPost200Response**](MerchantPlanAddonsBindingPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanListGet

> MerchantPlanListGet200Response PlanListGet(ctx).PlanIds(planIds).ProductIds(productIds).Tqype(tqype).Status(status).PublishStatus(publishStatus).Currency(currency).SearchKey(searchKey).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Get Plan List

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
	planIds := []int64{int64(123)} // []int64 | filter id list of plan, default all (optional)
	productIds := []int64{int64(123)} // []int64 | filter id list of product, default product(0) used if not specified (optional)
	tqype := []int32{int32(123)} // []int32 | 1-main plan，2-addon plan,3-onetime (optional)
	status := []int32{int32(123)} // []int32 | Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-SoftArchive, 5-HardArchive (optional)
	publishStatus := int32(56) // int32 | Filter, Default All，PublishStatus，1-UnPublished，2-Published (optional)
	currency := "currency_example" // string | Filter Currency (optional)
	searchKey := "searchKey_example" // string | Search Key, plan name or description (optional)
	sortField := "sortField_example" // string | Sort Field，plan_name|gmt_create|gmt_modify，Default gmt_create (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanListGet(context.Background()).PlanIds(planIds).ProductIds(productIds).Tqype(tqype).Status(status).PublishStatus(publishStatus).Currency(currency).SearchKey(searchKey).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanListGet`: MerchantPlanListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planIds** | **[]int64** | filter id list of plan, default all | 
 **productIds** | **[]int64** | filter id list of product, default product(0) used if not specified | 
 **tqype** | **[]int32** | 1-main plan，2-addon plan,3-onetime | 
 **status** | **[]int32** | Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-SoftArchive, 5-HardArchive | 
 **publishStatus** | **int32** | Filter, Default All，PublishStatus，1-UnPublished，2-Published | 
 **currency** | **string** | Filter Currency | 
 **searchKey** | **string** | Search Key, plan name or description | 
 **sortField** | **string** | Sort Field，plan_name|gmt_create|gmt_modify，Default gmt_create | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 

### Return type

[**MerchantPlanListGet200Response**](MerchantPlanListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanListPost

> MerchantPlanListGet200Response PlanListPost(ctx).UnibeeApiMerchantPlanListReq(unibeeApiMerchantPlanListReq).Execute()

Get Plan List

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
	unibeeApiMerchantPlanListReq := *openapiclient.NewUnibeeApiMerchantPlanListReq() // UnibeeApiMerchantPlanListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanListPost(context.Background()).UnibeeApiMerchantPlanListReq(unibeeApiMerchantPlanListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanListPost`: MerchantPlanListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanListReq** | [**UnibeeApiMerchantPlanListReq**](UnibeeApiMerchantPlanListReq.md) |  | 

### Return type

[**MerchantPlanListGet200Response**](MerchantPlanListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanNewPost

> MerchantPlanAddonsBindingPost200Response PlanNewPost(ctx).UnibeeApiMerchantPlanNewReq(unibeeApiMerchantPlanNewReq).Execute()

Create Plan

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
	unibeeApiMerchantPlanNewReq := *openapiclient.NewUnibeeApiMerchantPlanNewReq(int64(123), "Currency_example", "PlanName_example") // UnibeeApiMerchantPlanNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanNewPost(context.Background()).UnibeeApiMerchantPlanNewReq(unibeeApiMerchantPlanNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanNewPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanNewReq** | [**UnibeeApiMerchantPlanNewReq**](UnibeeApiMerchantPlanNewReq.md) |  | 

### Return type

[**MerchantPlanAddonsBindingPost200Response**](MerchantPlanAddonsBindingPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanPublishPost

> MerchantAuthSsoClearTotpPost200Response PlanPublishPost(ctx).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()

Publish Plan



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
	unibeeApiMerchantPlanPublishReq := *openapiclient.NewUnibeeApiMerchantPlanPublishReq() // UnibeeApiMerchantPlanPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanPublishPost(context.Background()).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanPublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanPublishPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanPublishPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanPublishPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanPublishReq** | [**UnibeeApiMerchantPlanPublishReq**](UnibeeApiMerchantPlanPublishReq.md) |  | 

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


## PlanUnpublishedPost

> MerchantAuthSsoClearTotpPost200Response PlanUnpublishedPost(ctx).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()

UnPublish Plan

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
	unibeeApiMerchantPlanUnPublishReq := *openapiclient.NewUnibeeApiMerchantPlanUnPublishReq() // UnibeeApiMerchantPlanUnPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanUnpublishedPost(context.Background()).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanUnpublishedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanUnpublishedPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanUnpublishedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanUnpublishedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanUnPublishReq** | [**UnibeeApiMerchantPlanUnPublishReq**](UnibeeApiMerchantPlanUnPublishReq.md) |  | 

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

