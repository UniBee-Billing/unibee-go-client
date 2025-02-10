# \AdminRole

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDeletePost**](AdminRole.md#RoleDeletePost) | **Post** /merchant/role/delete | Delete Role
[**RoleEditPost**](AdminRole.md#RoleEditPost) | **Post** /merchant/role/edit | Edit Role
[**RoleListGet**](AdminRole.md#RoleListGet) | **Get** /merchant/role/list | Get Role List
[**RoleNewPost**](AdminRole.md#RoleNewPost) | **Post** /merchant/role/new | New Role



## RoleDeletePost

> MerchantAuthSsoLoginOTPPost200Response RoleDeletePost(ctx).UnibeeApiMerchantRoleDeleteReq(unibeeApiMerchantRoleDeleteReq).Execute()

Delete Role

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
	unibeeApiMerchantRoleDeleteReq := *openapiclient.NewUnibeeApiMerchantRoleDeleteReq(int64(123)) // UnibeeApiMerchantRoleDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminRole.RoleDeletePost(context.Background()).UnibeeApiMerchantRoleDeleteReq(unibeeApiMerchantRoleDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminRole.RoleDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminRole.RoleDeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantRoleDeleteReq** | [**UnibeeApiMerchantRoleDeleteReq**](UnibeeApiMerchantRoleDeleteReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleEditPost

> MerchantAuthSsoLoginOTPPost200Response RoleEditPost(ctx).UnibeeApiMerchantRoleEditReq(unibeeApiMerchantRoleEditReq).Execute()

Edit Role

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
	unibeeApiMerchantRoleEditReq := *openapiclient.NewUnibeeApiMerchantRoleEditReq(int64(123), []openapiclient.UnibeeApiBeanMerchantRolePermission{*openapiclient.NewUnibeeApiBeanMerchantRolePermission()}, "Role_example") // UnibeeApiMerchantRoleEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminRole.RoleEditPost(context.Background()).UnibeeApiMerchantRoleEditReq(unibeeApiMerchantRoleEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminRole.RoleEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleEditPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminRole.RoleEditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantRoleEditReq** | [**UnibeeApiMerchantRoleEditReq**](UnibeeApiMerchantRoleEditReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleListGet

> MerchantRoleListGet200Response RoleListGet(ctx).Execute()

Get Role List

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminRole.RoleListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminRole.RoleListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleListGet`: MerchantRoleListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminRole.RoleListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRoleListGetRequest struct via the builder pattern


### Return type

[**MerchantRoleListGet200Response**](MerchantRoleListGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleNewPost

> MerchantAuthSsoLoginOTPPost200Response RoleNewPost(ctx).UnibeeApiMerchantRoleNewReq(unibeeApiMerchantRoleNewReq).Execute()

New Role

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
	unibeeApiMerchantRoleNewReq := *openapiclient.NewUnibeeApiMerchantRoleNewReq([]openapiclient.UnibeeApiBeanMerchantRolePermission{*openapiclient.NewUnibeeApiBeanMerchantRolePermission()}, "Role_example") // UnibeeApiMerchantRoleNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminRole.RoleNewPost(context.Background()).UnibeeApiMerchantRoleNewReq(unibeeApiMerchantRoleNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminRole.RoleNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleNewPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminRole.RoleNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantRoleNewReq** | [**UnibeeApiMerchantRoleNewReq**](UnibeeApiMerchantRoleNewReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

