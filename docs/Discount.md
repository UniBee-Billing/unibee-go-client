# \Discount

All URIs are relative to *http://api.unibee.top*

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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
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

> MerchantDiscountDetailGet200Response DiscountEditPost(ctx).UnibeeApiMerchantDiscountEditReq(unibeeApiMerchantDiscountEditReq).Execute()

EditDiscountCode



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
	unibeeApiMerchantDiscountEditReq := *openapiclient.NewUnibeeApiMerchantDiscountEditReq(int64(123)) // UnibeeApiMerchantDiscountEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountEditPost(context.Background()).UnibeeApiMerchantDiscountEditReq(unibeeApiMerchantDiscountEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountEditPost`: MerchantDiscountDetailGet200Response
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

[**MerchantDiscountDetailGet200Response**](MerchantDiscountDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountListGet

> MerchantDiscountListGet200Response DiscountListGet(ctx).Execute()

DiscountCodeList



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
	resp, r, err := apiClient.Discount.DiscountListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountListGet`: MerchantDiscountListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Discount.DiscountListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDiscountListGetRequest struct via the builder pattern


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

> MerchantDiscountDetailGet200Response DiscountNewPost(ctx).UnibeeApiMerchantDiscountNewReq(unibeeApiMerchantDiscountNewReq).Execute()

NewDiscountCode



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
	unibeeApiMerchantDiscountNewReq := *openapiclient.NewUnibeeApiMerchantDiscountNewReq(int32(123), "Code_example", int32(123), int64(123), int64(123)) // UnibeeApiMerchantDiscountNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Discount.DiscountNewPost(context.Background()).UnibeeApiMerchantDiscountNewReq(unibeeApiMerchantDiscountNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Discount.DiscountNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountNewPost`: MerchantDiscountDetailGet200Response
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

[**MerchantDiscountDetailGet200Response**](MerchantDiscountDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

