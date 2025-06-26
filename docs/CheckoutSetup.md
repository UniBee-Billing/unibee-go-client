# \CheckoutSetup

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckoutArchivePost**](CheckoutSetup.md#CheckoutArchivePost) | **Post** /merchant/checkout/archive | Archive Merchant Checkout
[**CheckoutDetailGet**](CheckoutSetup.md#CheckoutDetailGet) | **Get** /merchant/checkout/detail | Get Merchant Checkout Detail
[**CheckoutDetailPost**](CheckoutSetup.md#CheckoutDetailPost) | **Post** /merchant/checkout/detail | Get Merchant Checkout Detail
[**CheckoutEditCheckoutPost**](CheckoutSetup.md#CheckoutEditCheckoutPost) | **Post** /merchant/checkout/edit_checkout | Edit Merchant Checkout
[**CheckoutGetLinkGet**](CheckoutSetup.md#CheckoutGetLinkGet) | **Get** /merchant/checkout/get_link | Get Merchant Checkout Link
[**CheckoutGetLinkPost**](CheckoutSetup.md#CheckoutGetLinkPost) | **Post** /merchant/checkout/get_link | Get Merchant Checkout Link
[**CheckoutListGet**](CheckoutSetup.md#CheckoutListGet) | **Get** /merchant/checkout/list | Get Merchant Checkout list
[**CheckoutListPost**](CheckoutSetup.md#CheckoutListPost) | **Post** /merchant/checkout/list | Get Merchant Checkout list
[**CheckoutNewCheckoutPost**](CheckoutSetup.md#CheckoutNewCheckoutPost) | **Post** /merchant/checkout/new_checkout | New Merchant Checkout



## CheckoutArchivePost

> MerchantAuthSsoClearTotpPost200Response CheckoutArchivePost(ctx).UnibeeApiMerchantCheckoutArchiveReq(unibeeApiMerchantCheckoutArchiveReq).Execute()

Archive Merchant Checkout

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
	unibeeApiMerchantCheckoutArchiveReq := *openapiclient.NewUnibeeApiMerchantCheckoutArchiveReq(int64(123)) // UnibeeApiMerchantCheckoutArchiveReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutArchivePost(context.Background()).UnibeeApiMerchantCheckoutArchiveReq(unibeeApiMerchantCheckoutArchiveReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutArchivePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutArchivePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutArchivePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutArchivePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutArchiveReq** | [**UnibeeApiMerchantCheckoutArchiveReq**](UnibeeApiMerchantCheckoutArchiveReq.md) |  | 

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


## CheckoutDetailGet

> MerchantCheckoutDetailGet200Response CheckoutDetailGet(ctx).CheckoutId(checkoutId).Execute()

Get Merchant Checkout Detail

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
	checkoutId := int64(789) // int64 | checkout id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutDetailGet(context.Background()).CheckoutId(checkoutId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutDetailGet`: MerchantCheckoutDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutId** | **int64** | checkout id | 

### Return type

[**MerchantCheckoutDetailGet200Response**](MerchantCheckoutDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutDetailPost

> MerchantCheckoutDetailGet200Response CheckoutDetailPost(ctx).UnibeeApiMerchantCheckoutDetailReq(unibeeApiMerchantCheckoutDetailReq).Execute()

Get Merchant Checkout Detail

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
	unibeeApiMerchantCheckoutDetailReq := *openapiclient.NewUnibeeApiMerchantCheckoutDetailReq(int64(123)) // UnibeeApiMerchantCheckoutDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutDetailPost(context.Background()).UnibeeApiMerchantCheckoutDetailReq(unibeeApiMerchantCheckoutDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutDetailPost`: MerchantCheckoutDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutDetailReq** | [**UnibeeApiMerchantCheckoutDetailReq**](UnibeeApiMerchantCheckoutDetailReq.md) |  | 

### Return type

[**MerchantCheckoutDetailGet200Response**](MerchantCheckoutDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutEditCheckoutPost

> MerchantCheckoutDetailGet200Response CheckoutEditCheckoutPost(ctx).UnibeeApiMerchantCheckoutEditReq(unibeeApiMerchantCheckoutEditReq).Execute()

Edit Merchant Checkout

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
	unibeeApiMerchantCheckoutEditReq := *openapiclient.NewUnibeeApiMerchantCheckoutEditReq(int64(123)) // UnibeeApiMerchantCheckoutEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutEditCheckoutPost(context.Background()).UnibeeApiMerchantCheckoutEditReq(unibeeApiMerchantCheckoutEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutEditCheckoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutEditCheckoutPost`: MerchantCheckoutDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutEditCheckoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutEditCheckoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutEditReq** | [**UnibeeApiMerchantCheckoutEditReq**](UnibeeApiMerchantCheckoutEditReq.md) |  | 

### Return type

[**MerchantCheckoutDetailGet200Response**](MerchantCheckoutDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutGetLinkGet

> MerchantCheckoutGetLinkGet200Response CheckoutGetLinkGet(ctx).CheckoutId(checkoutId).PlanId(planId).Execute()

Get Merchant Checkout Link

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
	checkoutId := int64(789) // int64 | checkout id
	planId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutGetLinkGet(context.Background()).CheckoutId(checkoutId).PlanId(planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutGetLinkGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutGetLinkGet`: MerchantCheckoutGetLinkGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutGetLinkGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutGetLinkGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkoutId** | **int64** | checkout id | 
 **planId** | **int64** |  | 

### Return type

[**MerchantCheckoutGetLinkGet200Response**](MerchantCheckoutGetLinkGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutGetLinkPost

> MerchantCheckoutGetLinkGet200Response CheckoutGetLinkPost(ctx).UnibeeApiMerchantCheckoutGetLinkReq(unibeeApiMerchantCheckoutGetLinkReq).Execute()

Get Merchant Checkout Link

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
	unibeeApiMerchantCheckoutGetLinkReq := *openapiclient.NewUnibeeApiMerchantCheckoutGetLinkReq(int64(123), int64(123)) // UnibeeApiMerchantCheckoutGetLinkReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutGetLinkPost(context.Background()).UnibeeApiMerchantCheckoutGetLinkReq(unibeeApiMerchantCheckoutGetLinkReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutGetLinkPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutGetLinkPost`: MerchantCheckoutGetLinkGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutGetLinkPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutGetLinkPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutGetLinkReq** | [**UnibeeApiMerchantCheckoutGetLinkReq**](UnibeeApiMerchantCheckoutGetLinkReq.md) |  | 

### Return type

[**MerchantCheckoutGetLinkGet200Response**](MerchantCheckoutGetLinkGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutListGet

> MerchantCheckoutListGet200Response CheckoutListGet(ctx).SearchKey(searchKey).Execute()

Get Merchant Checkout list

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
	searchKey := "searchKey_example" // string | Search checkout id|name|description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutListGet(context.Background()).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutListGet`: MerchantCheckoutListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | Search checkout id|name|description | 

### Return type

[**MerchantCheckoutListGet200Response**](MerchantCheckoutListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutListPost

> MerchantCheckoutListGet200Response CheckoutListPost(ctx).UnibeeApiMerchantCheckoutListReq(unibeeApiMerchantCheckoutListReq).Execute()

Get Merchant Checkout list

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
	unibeeApiMerchantCheckoutListReq := *openapiclient.NewUnibeeApiMerchantCheckoutListReq() // UnibeeApiMerchantCheckoutListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutListPost(context.Background()).UnibeeApiMerchantCheckoutListReq(unibeeApiMerchantCheckoutListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutListPost`: MerchantCheckoutListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutListReq** | [**UnibeeApiMerchantCheckoutListReq**](UnibeeApiMerchantCheckoutListReq.md) |  | 

### Return type

[**MerchantCheckoutListGet200Response**](MerchantCheckoutListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckoutNewCheckoutPost

> MerchantCheckoutDetailGet200Response CheckoutNewCheckoutPost(ctx).UnibeeApiMerchantCheckoutNewReq(unibeeApiMerchantCheckoutNewReq).Execute()

New Merchant Checkout

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
	unibeeApiMerchantCheckoutNewReq := *openapiclient.NewUnibeeApiMerchantCheckoutNewReq() // UnibeeApiMerchantCheckoutNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CheckoutSetup.CheckoutNewCheckoutPost(context.Background()).UnibeeApiMerchantCheckoutNewReq(unibeeApiMerchantCheckoutNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CheckoutSetup.CheckoutNewCheckoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckoutNewCheckoutPost`: MerchantCheckoutDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CheckoutSetup.CheckoutNewCheckoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckoutNewCheckoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCheckoutNewReq** | [**UnibeeApiMerchantCheckoutNewReq**](UnibeeApiMerchantCheckoutNewReq.md) |  | 

### Return type

[**MerchantCheckoutDetailGet200Response**](MerchantCheckoutDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

