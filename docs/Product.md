# \Product

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductActivatePost**](Product.md#ProductActivatePost) | **Post** /merchant/product/activate | ActivateProduct
[**ProductCopyPost**](Product.md#ProductCopyPost) | **Post** /merchant/product/copy | CopyProduct
[**ProductDeletePost**](Product.md#ProductDeletePost) | **Post** /merchant/product/delete | DeleteProduct
[**ProductDetailGet**](Product.md#ProductDetailGet) | **Get** /merchant/product/detail | ProductDetail
[**ProductDetailPost**](Product.md#ProductDetailPost) | **Post** /merchant/product/detail | ProductDetail
[**ProductEditPost**](Product.md#ProductEditPost) | **Post** /merchant/product/edit | EditProduct
[**ProductInactivatePost**](Product.md#ProductInactivatePost) | **Post** /merchant/product/inactivate | InactivateProduct
[**ProductListGet**](Product.md#ProductListGet) | **Get** /merchant/product/list | ProductList
[**ProductListPost**](Product.md#ProductListPost) | **Post** /merchant/product/list | ProductList
[**ProductNewPost**](Product.md#ProductNewPost) | **Post** /merchant/product/new | CreateProduct



## ProductActivatePost

> MerchantAuthSsoLoginOTPPost200Response ProductActivatePost(ctx).UnibeeApiMerchantProductActivateReq(unibeeApiMerchantProductActivateReq).Execute()

ActivateProduct

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
	unibeeApiMerchantProductActivateReq := *openapiclient.NewUnibeeApiMerchantProductActivateReq(int64(123)) // UnibeeApiMerchantProductActivateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductActivatePost(context.Background()).UnibeeApiMerchantProductActivateReq(unibeeApiMerchantProductActivateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductActivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductActivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductActivateReq** | [**UnibeeApiMerchantProductActivateReq**](UnibeeApiMerchantProductActivateReq.md) |  | 

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


## ProductCopyPost

> MerchantProductCopyPost200Response ProductCopyPost(ctx).UnibeeApiMerchantProductCopyReq(unibeeApiMerchantProductCopyReq).Execute()

CopyProduct

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
	unibeeApiMerchantProductCopyReq := *openapiclient.NewUnibeeApiMerchantProductCopyReq(int64(123)) // UnibeeApiMerchantProductCopyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductCopyPost(context.Background()).UnibeeApiMerchantProductCopyReq(unibeeApiMerchantProductCopyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductCopyPost`: MerchantProductCopyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductCopyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductCopyReq** | [**UnibeeApiMerchantProductCopyReq**](UnibeeApiMerchantProductCopyReq.md) |  | 

### Return type

[**MerchantProductCopyPost200Response**](MerchantProductCopyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDeletePost

> MerchantAuthSsoLoginOTPPost200Response ProductDeletePost(ctx).UnibeeApiMerchantProductDeleteReq(unibeeApiMerchantProductDeleteReq).Execute()

DeleteProduct



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
	unibeeApiMerchantProductDeleteReq := *openapiclient.NewUnibeeApiMerchantProductDeleteReq(int64(123)) // UnibeeApiMerchantProductDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductDeletePost(context.Background()).UnibeeApiMerchantProductDeleteReq(unibeeApiMerchantProductDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductDeleteReq** | [**UnibeeApiMerchantProductDeleteReq**](UnibeeApiMerchantProductDeleteReq.md) |  | 

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


## ProductDetailGet

> MerchantProductCopyPost200Response ProductDetailGet(ctx).ProductId(productId).Execute()

ProductDetail

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
	productId := int64(789) // int64 | ProductId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductDetailGet(context.Background()).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDetailGet`: MerchantProductCopyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **int64** | ProductId | 

### Return type

[**MerchantProductCopyPost200Response**](MerchantProductCopyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductDetailPost

> MerchantProductCopyPost200Response ProductDetailPost(ctx).UnibeeApiMerchantProductDetailReq(unibeeApiMerchantProductDetailReq).Execute()

ProductDetail

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
	unibeeApiMerchantProductDetailReq := *openapiclient.NewUnibeeApiMerchantProductDetailReq(int64(123)) // UnibeeApiMerchantProductDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductDetailPost(context.Background()).UnibeeApiMerchantProductDetailReq(unibeeApiMerchantProductDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductDetailPost`: MerchantProductCopyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductDetailReq** | [**UnibeeApiMerchantProductDetailReq**](UnibeeApiMerchantProductDetailReq.md) |  | 

### Return type

[**MerchantProductCopyPost200Response**](MerchantProductCopyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductEditPost

> MerchantProductCopyPost200Response ProductEditPost(ctx).UnibeeApiMerchantProductEditReq(unibeeApiMerchantProductEditReq).Execute()

EditProduct



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
	unibeeApiMerchantProductEditReq := *openapiclient.NewUnibeeApiMerchantProductEditReq(int64(123)) // UnibeeApiMerchantProductEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductEditPost(context.Background()).UnibeeApiMerchantProductEditReq(unibeeApiMerchantProductEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductEditPost`: MerchantProductCopyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductEditReq** | [**UnibeeApiMerchantProductEditReq**](UnibeeApiMerchantProductEditReq.md) |  | 

### Return type

[**MerchantProductCopyPost200Response**](MerchantProductCopyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductInactivatePost

> MerchantAuthSsoLoginOTPPost200Response ProductInactivatePost(ctx).UnibeeApiMerchantProductInactiveReq(unibeeApiMerchantProductInactiveReq).Execute()

InactivateProduct

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
	unibeeApiMerchantProductInactiveReq := *openapiclient.NewUnibeeApiMerchantProductInactiveReq(int64(123)) // UnibeeApiMerchantProductInactiveReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductInactivatePost(context.Background()).UnibeeApiMerchantProductInactiveReq(unibeeApiMerchantProductInactiveReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductInactivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductInactivatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductInactivatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductInactivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductInactiveReq** | [**UnibeeApiMerchantProductInactiveReq**](UnibeeApiMerchantProductInactiveReq.md) |  | 

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


## ProductListGet

> MerchantProductListGet200Response ProductListGet(ctx).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

ProductList

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
	status := []int32{int32(123)} // []int32 | Filter, Default All，,Status，1-active，2-inactive (optional)
	sortField := "sortField_example" // string | Sort Field，id|create_time|gmt_modify，Default id (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductListGet(context.Background()).Status(status).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductListGet`: MerchantProductListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]int32** | Filter, Default All，,Status，1-active，2-inactive | 
 **sortField** | **string** | Sort Field，id|create_time|gmt_modify，Default id | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 

### Return type

[**MerchantProductListGet200Response**](MerchantProductListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductListPost

> MerchantProductListGet200Response ProductListPost(ctx).UnibeeApiMerchantProductListReq(unibeeApiMerchantProductListReq).Execute()

ProductList

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
	unibeeApiMerchantProductListReq := *openapiclient.NewUnibeeApiMerchantProductListReq() // UnibeeApiMerchantProductListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductListPost(context.Background()).UnibeeApiMerchantProductListReq(unibeeApiMerchantProductListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductListPost`: MerchantProductListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductListReq** | [**UnibeeApiMerchantProductListReq**](UnibeeApiMerchantProductListReq.md) |  | 

### Return type

[**MerchantProductListGet200Response**](MerchantProductListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductNewPost

> MerchantProductCopyPost200Response ProductNewPost(ctx).UnibeeApiMerchantProductNewReq(unibeeApiMerchantProductNewReq).Execute()

CreateProduct

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
	unibeeApiMerchantProductNewReq := *openapiclient.NewUnibeeApiMerchantProductNewReq() // UnibeeApiMerchantProductNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Product.ProductNewPost(context.Background()).UnibeeApiMerchantProductNewReq(unibeeApiMerchantProductNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Product.ProductNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductNewPost`: MerchantProductCopyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Product.ProductNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantProductNewReq** | [**UnibeeApiMerchantProductNewReq**](UnibeeApiMerchantProductNewReq.md) |  | 

### Return type

[**MerchantProductCopyPost200Response**](MerchantProductCopyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

