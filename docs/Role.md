# \Role

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoleDeletePost**](Role.md#RoleDeletePost) | **Post** /merchant/role/delete | DeleteRole
[**RoleEditPost**](Role.md#RoleEditPost) | **Post** /merchant/role/edit | EditRole
[**RoleListGet**](Role.md#RoleListGet) | **Get** /merchant/role/list | RoleList
[**RoleNewPost**](Role.md#RoleNewPost) | **Post** /merchant/role/new | NewRole



## RoleDeletePost

> MerchantAuthSsoLoginOTPPost200Response RoleDeletePost(ctx).UnibeeApiMerchantRoleDeleteReq(unibeeApiMerchantRoleDeleteReq).Execute()

DeleteRole

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
	unibeeApiMerchantRoleDeleteReq := *openapiclient.NewUnibeeApiMerchantRoleDeleteReq("Role_example") // UnibeeApiMerchantRoleDeleteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Role.RoleDeletePost(context.Background()).UnibeeApiMerchantRoleDeleteReq(unibeeApiMerchantRoleDeleteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Role.RoleDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleDeletePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Role.RoleDeletePost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleEditPost

> MerchantAuthSsoLoginOTPPost200Response RoleEditPost(ctx).UnibeeApiMerchantRoleEditReq(unibeeApiMerchantRoleEditReq).Execute()

EditRole

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
	unibeeApiMerchantRoleEditReq := *openapiclient.NewUnibeeApiMerchantRoleEditReq([]openapiclient.UnibeeApiBeanMerchantRolePermission{*openapiclient.NewUnibeeApiBeanMerchantRolePermission()}, "Role_example") // UnibeeApiMerchantRoleEditReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Role.RoleEditPost(context.Background()).UnibeeApiMerchantRoleEditReq(unibeeApiMerchantRoleEditReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Role.RoleEditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleEditPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Role.RoleEditPost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleListGet

> MerchantRoleListGet200Response RoleListGet(ctx).Execute()

RoleList

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
	resp, r, err := apiClient.Role.RoleListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Role.RoleListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleListGet`: MerchantRoleListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Role.RoleListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRoleListGetRequest struct via the builder pattern


### Return type

[**MerchantRoleListGet200Response**](MerchantRoleListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleNewPost

> MerchantAuthSsoLoginOTPPost200Response RoleNewPost(ctx).UnibeeApiMerchantRoleNewReq(unibeeApiMerchantRoleNewReq).Execute()

NewRole

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
	unibeeApiMerchantRoleNewReq := *openapiclient.NewUnibeeApiMerchantRoleNewReq([]openapiclient.UnibeeApiBeanMerchantRolePermission{*openapiclient.NewUnibeeApiBeanMerchantRolePermission()}, "Role_example") // UnibeeApiMerchantRoleNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Role.RoleNewPost(context.Background()).UnibeeApiMerchantRoleNewReq(unibeeApiMerchantRoleNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Role.RoleNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RoleNewPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Role.RoleNewPost`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

