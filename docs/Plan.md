# \Plan

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlanActivatePost**](Plan.md#PlanActivatePost) | **Post** /merchant/plan/activate | ActivatePlan
[**PlanAddonsBindingPost**](Plan.md#PlanAddonsBindingPost) | **Post** /merchant/plan/addons_binding | AddonBinding
[**PlanCopyPost**](Plan.md#PlanCopyPost) | **Post** /merchant/plan/copy | CopyPlan
[**PlanDeletePost**](Plan.md#PlanDeletePost) | **Post** /merchant/plan/delete | DeletePlan
[**PlanDetailGet**](Plan.md#PlanDetailGet) | **Get** /merchant/plan/detail | PlanDetail
[**PlanDetailPost**](Plan.md#PlanDetailPost) | **Post** /merchant/plan/detail | PlanDetail
[**PlanEditPost**](Plan.md#PlanEditPost) | **Post** /merchant/plan/edit | EditPlan
[**PlanExpirePost**](Plan.md#PlanExpirePost) | **Post** /merchant/plan/expire | ExpirePlan
[**PlanListGet**](Plan.md#PlanListGet) | **Get** /merchant/plan/list | PlanList
[**PlanListPost**](Plan.md#PlanListPost) | **Post** /merchant/plan/list | PlanList
[**PlanNewPost**](Plan.md#PlanNewPost) | **Post** /merchant/plan/new | CreatePlan
[**PlanPublishPost**](Plan.md#PlanPublishPost) | **Post** /merchant/plan/publish | PublishPlan
[**PlanUnpublishedPost**](Plan.md#PlanUnpublishedPost) | **Post** /merchant/plan/unpublished | UnPublishPlan



## PlanActivatePost

> MerchantAuthSsoLoginOTPPost200Response PlanActivatePost(ctx).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()

ActivatePlan

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
	unibeeApiMerchantPlanActivateReq := *openapiclient.NewUnibeeApiMerchantPlanActivateReq(int64(123)) // UnibeeApiMerchantPlanActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanActivatePost(context.Background()).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanActivatePost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

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

AddonBinding

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


## PlanCopyPost

> MerchantPlanAddonsBindingPost200Response PlanCopyPost(ctx).UnibeeApiMerchantPlanCopyReq(unibeeApiMerchantPlanCopyReq).Execute()

CopyPlan

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
	unibeeApiMerchantPlanCopyReq := *openapiclient.NewUnibeeApiMerchantPlanCopyReq(int64(123)) // UnibeeApiMerchantPlanCopyReq | 

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

> MerchantAuthSsoLoginOTPPost200Response PlanDeletePost(ctx).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()

DeletePlan

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
	unibeeApiMerchantPlanDeleteReq := *openapiclient.NewUnibeeApiMerchantPlanDeleteReq(int64(123)) // UnibeeApiMerchantPlanDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanDeletePost(context.Background()).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanDeletePost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanDetailGet

> MerchantPlanDetailGet200Response PlanDetailGet(ctx).PlanId(planId).Execute()

PlanDetail

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
	planId := int64(789) // int64 | PlanId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanDetailGet(context.Background()).PlanId(planId).Execute()
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
 **planId** | **int64** | PlanId | 

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

PlanDetail

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
	unibeeApiMerchantPlanDetailReq := *openapiclient.NewUnibeeApiMerchantPlanDetailReq(int64(123)) // UnibeeApiMerchantPlanDetailReq | 

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

EditPlan



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


## PlanExpirePost

> MerchantAuthSsoLoginOTPPost200Response PlanExpirePost(ctx).UnibeeApiMerchantPlanExpireReq(unibeeApiMerchantPlanExpireReq).Execute()

ExpirePlan

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
	unibeeApiMerchantPlanExpireReq := *openapiclient.NewUnibeeApiMerchantPlanExpireReq(int64(123), int64(123)) // UnibeeApiMerchantPlanExpireReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanExpirePost(context.Background()).UnibeeApiMerchantPlanExpireReq(unibeeApiMerchantPlanExpireReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanExpirePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanExpirePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Plan.PlanExpirePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanExpirePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantPlanExpireReq** | [**UnibeeApiMerchantPlanExpireReq**](UnibeeApiMerchantPlanExpireReq.md) |  | 

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


## PlanListGet

> MerchantPlanListGet200Response PlanListGet(ctx).ProductIds(productIds).Type_(type_).Status(status).PublishStatus(publishStatus).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

PlanList

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
	productIds := []int64{int64(123)} // []int64 | filter id list of product, default all (optional)
	type_ := []int32{int32(123)} // []int32 | 1-main plan，2-addon plan (optional)
	status := []int32{int32(123)} // []int32 | Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-Expired (optional)
	publishStatus := int32(56) // int32 | Filter, Default All，PublishStatus，1-UnPublished，2-Published (optional)
	currency := "currency_example" // string | Filter Currency (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanListGet(context.Background()).ProductIds(productIds).Type_(type_).Status(status).PublishStatus(publishStatus).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
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
 **productIds** | **[]int64** | filter id list of product, default all | 
 **type_** | **[]int32** | 1-main plan，2-addon plan | 
 **status** | **[]int32** | Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-Expired | 
 **publishStatus** | **int32** | Filter, Default All，PublishStatus，1-UnPublished，2-Published | 
 **currency** | **string** | Filter Currency | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
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

PlanList

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

CreatePlan

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

> MerchantAuthSsoLoginOTPPost200Response PlanPublishPost(ctx).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()

PublishPlan



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
	unibeeApiMerchantPlanPublishReq := *openapiclient.NewUnibeeApiMerchantPlanPublishReq(int64(123)) // UnibeeApiMerchantPlanPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanPublishPost(context.Background()).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanPublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanPublishPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlanUnpublishedPost

> MerchantAuthSsoLoginOTPPost200Response PlanUnpublishedPost(ctx).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()

UnPublishPlan

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
	unibeeApiMerchantPlanUnPublishReq := *openapiclient.NewUnibeeApiMerchantPlanUnPublishReq(int64(123)) // UnibeeApiMerchantPlanUnPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Plan.PlanUnpublishedPost(context.Background()).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Plan.PlanUnpublishedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlanUnpublishedPost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

