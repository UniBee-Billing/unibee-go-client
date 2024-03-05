# \MerchantPlanAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantPlanActivatePost**](MerchantPlanAPI.md#MerchantPlanActivatePost) | **Post** /merchant/plan/activate | Plan Sync To Gateway And Activate
[**MerchantPlanAddonsBindingPost**](MerchantPlanAPI.md#MerchantPlanAddonsBindingPost) | **Post** /merchant/plan/addons_binding | Plan Binding Addons
[**MerchantPlanDeletePost**](MerchantPlanAPI.md#MerchantPlanDeletePost) | **Post** /merchant/plan/delete | Delete A Plan Before Activate
[**MerchantPlanDetailGet**](MerchantPlanAPI.md#MerchantPlanDetailGet) | **Get** /merchant/plan/detail | Plan Detail
[**MerchantPlanDetailPost**](MerchantPlanAPI.md#MerchantPlanDetailPost) | **Post** /merchant/plan/detail | Plan Detail
[**MerchantPlanEditPost**](MerchantPlanAPI.md#MerchantPlanEditPost) | **Post** /merchant/plan/edit | Edit Plan
[**MerchantPlanExpirePost**](MerchantPlanAPI.md#MerchantPlanExpirePost) | **Post** /merchant/plan/expire | Expire A Plan
[**MerchantPlanListGet**](MerchantPlanAPI.md#MerchantPlanListGet) | **Get** /merchant/plan/list | Plan List
[**MerchantPlanListPost**](MerchantPlanAPI.md#MerchantPlanListPost) | **Post** /merchant/plan/list | Plan List
[**MerchantPlanNewPost**](MerchantPlanAPI.md#MerchantPlanNewPost) | **Post** /merchant/plan/new | Create Plan
[**MerchantPlanPublishPost**](MerchantPlanAPI.md#MerchantPlanPublishPost) | **Post** /merchant/plan/publish | Publish Plan，Will Be Visible To UserPortal
[**MerchantPlanUnpublishedPost**](MerchantPlanAPI.md#MerchantPlanUnpublishedPost) | **Post** /merchant/plan/unpublished | UnPublish Plan



## MerchantPlanActivatePost

> MerchantAuthSsoLoginOTPPost200Response MerchantPlanActivatePost(ctx).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()

Plan Sync To Gateway And Activate

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
	unibeeApiMerchantPlanActivateReq := *openapiclient.NewUnibeeApiMerchantPlanActivateReq(int32(123)) // UnibeeApiMerchantPlanActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanActivatePost(context.Background()).UnibeeApiMerchantPlanActivateReq(unibeeApiMerchantPlanActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanActivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanActivatePostRequest struct via the builder pattern


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


## MerchantPlanAddonsBindingPost

> MerchantPlanAddonsBindingPost200Response MerchantPlanAddonsBindingPost(ctx).UnibeeApiMerchantPlanAddonsBindingReq(unibeeApiMerchantPlanAddonsBindingReq).Execute()

Plan Binding Addons

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
	unibeeApiMerchantPlanAddonsBindingReq := *openapiclient.NewUnibeeApiMerchantPlanAddonsBindingReq(int64(123), []int64{int64(123)}, int32(123)) // UnibeeApiMerchantPlanAddonsBindingReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanAddonsBindingPost(context.Background()).UnibeeApiMerchantPlanAddonsBindingReq(unibeeApiMerchantPlanAddonsBindingReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanAddonsBindingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanAddonsBindingPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanAddonsBindingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanAddonsBindingPostRequest struct via the builder pattern


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


## MerchantPlanDeletePost

> MerchantAuthSsoLoginOTPPost200Response MerchantPlanDeletePost(ctx).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()

Delete A Plan Before Activate

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
	unibeeApiMerchantPlanDeleteReq := *openapiclient.NewUnibeeApiMerchantPlanDeleteReq(int32(123)) // UnibeeApiMerchantPlanDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanDeletePost(context.Background()).UnibeeApiMerchantPlanDeleteReq(unibeeApiMerchantPlanDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanDeletePostRequest struct via the builder pattern


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


## MerchantPlanDetailGet

> MerchantPlanDetailGet200Response MerchantPlanDetailGet(ctx).PlanId(planId).Execute()

Plan Detail

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
	planId := int32(56) // int32 | PlanId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanDetailGet(context.Background()).PlanId(planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanDetailGet`: MerchantPlanDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planId** | **int32** | PlanId | 

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


## MerchantPlanDetailPost

> MerchantPlanDetailGet200Response MerchantPlanDetailPost(ctx).UnibeeApiMerchantPlanDetailReq(unibeeApiMerchantPlanDetailReq).Execute()

Plan Detail

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
	unibeeApiMerchantPlanDetailReq := *openapiclient.NewUnibeeApiMerchantPlanDetailReq(int32(123)) // UnibeeApiMerchantPlanDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanDetailPost(context.Background()).UnibeeApiMerchantPlanDetailReq(unibeeApiMerchantPlanDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanDetailPost`: MerchantPlanDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanDetailPostRequest struct via the builder pattern


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


## MerchantPlanEditPost

> MerchantPlanAddonsBindingPost200Response MerchantPlanEditPost(ctx).UnibeeApiMerchantPlanEditReq(unibeeApiMerchantPlanEditReq).Execute()

Edit Plan

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
	unibeeApiMerchantPlanEditReq := *openapiclient.NewUnibeeApiMerchantPlanEditReq(int64(123), "Currency_example", "IntervalUnit_example", int32(123), "PlanName_example") // UnibeeApiMerchantPlanEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanEditPost(context.Background()).UnibeeApiMerchantPlanEditReq(unibeeApiMerchantPlanEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanEditPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanEditPostRequest struct via the builder pattern


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


## MerchantPlanExpirePost

> MerchantAuthSsoLoginOTPPost200Response MerchantPlanExpirePost(ctx).UnibeeApiMerchantPlanExpireReq(unibeeApiMerchantPlanExpireReq).Execute()

Expire A Plan

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
	unibeeApiMerchantPlanExpireReq := *openapiclient.NewUnibeeApiMerchantPlanExpireReq(int64(123), int32(123)) // UnibeeApiMerchantPlanExpireReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanExpirePost(context.Background()).UnibeeApiMerchantPlanExpireReq(unibeeApiMerchantPlanExpireReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanExpirePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanExpirePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanExpirePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanExpirePostRequest struct via the builder pattern


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


## MerchantPlanListGet

> MerchantPlanListGet200Response MerchantPlanListGet(ctx).Type_(type_).Status(status).PublishStatus(publishStatus).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Plan List

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
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanListGet(context.Background()).Type_(type_).Status(status).PublishStatus(publishStatus).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanListGet`: MerchantPlanListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## MerchantPlanListPost

> MerchantPlanListGet200Response MerchantPlanListPost(ctx).UnibeeApiMerchantPlanListReq(unibeeApiMerchantPlanListReq).Execute()

Plan List

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
	unibeeApiMerchantPlanListReq := *openapiclient.NewUnibeeApiMerchantPlanListReq() // UnibeeApiMerchantPlanListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanListPost(context.Background()).UnibeeApiMerchantPlanListReq(unibeeApiMerchantPlanListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanListPost`: MerchantPlanListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanListPostRequest struct via the builder pattern


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


## MerchantPlanNewPost

> MerchantPlanAddonsBindingPost200Response MerchantPlanNewPost(ctx).UnibeeApiMerchantPlanNewReq(unibeeApiMerchantPlanNewReq).Execute()

Create Plan

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
	unibeeApiMerchantPlanNewReq := *openapiclient.NewUnibeeApiMerchantPlanNewReq(int64(123), "Currency_example", "IntervalUnit_example", "PlanName_example") // UnibeeApiMerchantPlanNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanNewPost(context.Background()).UnibeeApiMerchantPlanNewReq(unibeeApiMerchantPlanNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanNewPost`: MerchantPlanAddonsBindingPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanNewPostRequest struct via the builder pattern


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


## MerchantPlanPublishPost

> MerchantAuthSsoLoginOTPPost200Response MerchantPlanPublishPost(ctx).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()

Publish Plan，Will Be Visible To UserPortal

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
	unibeeApiMerchantPlanPublishReq := *openapiclient.NewUnibeeApiMerchantPlanPublishReq(int32(123)) // UnibeeApiMerchantPlanPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanPublishPost(context.Background()).UnibeeApiMerchantPlanPublishReq(unibeeApiMerchantPlanPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanPublishPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanPublishPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanPublishPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanPublishPostRequest struct via the builder pattern


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


## MerchantPlanUnpublishedPost

> MerchantAuthSsoLoginOTPPost200Response MerchantPlanUnpublishedPost(ctx).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()

UnPublish Plan

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
	unibeeApiMerchantPlanUnPublishReq := *openapiclient.NewUnibeeApiMerchantPlanUnPublishReq(int32(123)) // UnibeeApiMerchantPlanUnPublishReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantPlanAPI.MerchantPlanUnpublishedPost(context.Background()).UnibeeApiMerchantPlanUnPublishReq(unibeeApiMerchantPlanUnPublishReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantPlanAPI.MerchantPlanUnpublishedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantPlanUnpublishedPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantPlanAPI.MerchantPlanUnpublishedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantPlanUnpublishedPostRequest struct via the builder pattern


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

