# \Tag

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagDeletePost**](Tag.md#TagDeletePost) | **Post** /merchant/tag/delete | Delete Tag
[**TagDetailGet**](Tag.md#TagDetailGet) | **Get** /merchant/tag/detail | Tag Detail
[**TagDetailPost**](Tag.md#TagDetailPost) | **Post** /merchant/tag/detail | Tag Detail
[**TagEditPost**](Tag.md#TagEditPost) | **Post** /merchant/tag/edit | Edit Tag
[**TagListGet**](Tag.md#TagListGet) | **Get** /merchant/tag/list | Tag List
[**TagListPost**](Tag.md#TagListPost) | **Post** /merchant/tag/list | Tag List
[**TagNewPost**](Tag.md#TagNewPost) | **Post** /merchant/tag/new | Create Tag



## TagDeletePost

> MerchantAuthSsoClearTotpPost200Response TagDeletePost(ctx).UnibeeApiMerchantTagDeleteReq(unibeeApiMerchantTagDeleteReq).Execute()

Delete Tag

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
	unibeeApiMerchantTagDeleteReq := *openapiclient.NewUnibeeApiMerchantTagDeleteReq(int64(123)) // UnibeeApiMerchantTagDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagDeletePost(context.Background()).UnibeeApiMerchantTagDeleteReq(unibeeApiMerchantTagDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagDeletePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTagDeleteReq** | [**UnibeeApiMerchantTagDeleteReq**](UnibeeApiMerchantTagDeleteReq.md) |  | 

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


## TagDetailGet

> MerchantTagDetailGet200Response TagDetailGet(ctx).TagId(tagId).Execute()

Tag Detail

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
	tagId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagDetailGet(context.Background()).TagId(tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagDetailGet`: MerchantTagDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagId** | **int64** |  | 

### Return type

[**MerchantTagDetailGet200Response**](MerchantTagDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagDetailPost

> MerchantTagDetailGet200Response TagDetailPost(ctx).UnibeeApiMerchantTagDetailReq(unibeeApiMerchantTagDetailReq).Execute()

Tag Detail

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
	unibeeApiMerchantTagDetailReq := *openapiclient.NewUnibeeApiMerchantTagDetailReq(int64(123)) // UnibeeApiMerchantTagDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagDetailPost(context.Background()).UnibeeApiMerchantTagDetailReq(unibeeApiMerchantTagDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagDetailPost`: MerchantTagDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTagDetailReq** | [**UnibeeApiMerchantTagDetailReq**](UnibeeApiMerchantTagDetailReq.md) |  | 

### Return type

[**MerchantTagDetailGet200Response**](MerchantTagDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagEditPost

> MerchantTagEditPost200Response TagEditPost(ctx).UnibeeApiMerchantTagEditReq(unibeeApiMerchantTagEditReq).Execute()

Edit Tag

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
	unibeeApiMerchantTagEditReq := *openapiclient.NewUnibeeApiMerchantTagEditReq(int64(123)) // UnibeeApiMerchantTagEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagEditPost(context.Background()).UnibeeApiMerchantTagEditReq(unibeeApiMerchantTagEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagEditPost`: MerchantTagEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTagEditReq** | [**UnibeeApiMerchantTagEditReq**](UnibeeApiMerchantTagEditReq.md) |  | 

### Return type

[**MerchantTagEditPost200Response**](MerchantTagEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagListGet

> MerchantTagListGet200Response TagListGet(ctx).Scope(scope).SearchKey(searchKey).Page(page).Count(count).Execute()

Tag List

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
	scope := int32(56) // int32 | 0-plan
	searchKey := "searchKey_example" // string |  (optional)
	page := int32(56) // int32 | Page, start 0 (optional)
	count := int32(56) // int32 | Count per page, default 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagListGet(context.Background()).Scope(scope).SearchKey(searchKey).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagListGet`: MerchantTagListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **int32** | 0-plan | 
 **searchKey** | **string** |  | 
 **page** | **int32** | Page, start 0 | 
 **count** | **int32** | Count per page, default 100 | 

### Return type

[**MerchantTagListGet200Response**](MerchantTagListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagListPost

> MerchantTagListGet200Response TagListPost(ctx).UnibeeApiMerchantTagListReq(unibeeApiMerchantTagListReq).Execute()

Tag List

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
	unibeeApiMerchantTagListReq := *openapiclient.NewUnibeeApiMerchantTagListReq(int32(123)) // UnibeeApiMerchantTagListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagListPost(context.Background()).UnibeeApiMerchantTagListReq(unibeeApiMerchantTagListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagListPost`: MerchantTagListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTagListReq** | [**UnibeeApiMerchantTagListReq**](UnibeeApiMerchantTagListReq.md) |  | 

### Return type

[**MerchantTagListGet200Response**](MerchantTagListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TagNewPost

> MerchantTagEditPost200Response TagNewPost(ctx).UnibeeApiMerchantTagNewReq(unibeeApiMerchantTagNewReq).Execute()

Create Tag

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
	unibeeApiMerchantTagNewReq := *openapiclient.NewUnibeeApiMerchantTagNewReq("Name_example", int32(123)) // UnibeeApiMerchantTagNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Tag.TagNewPost(context.Background()).UnibeeApiMerchantTagNewReq(unibeeApiMerchantTagNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Tag.TagNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TagNewPost`: MerchantTagEditPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Tag.TagNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTagNewReq** | [**UnibeeApiMerchantTagNewReq**](UnibeeApiMerchantTagNewReq.md) |  | 

### Return type

[**MerchantTagEditPost200Response**](MerchantTagEditPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

