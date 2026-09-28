# \UserTag

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserTagEnsurePost**](UserTag.md#UserTagEnsurePost) | **Post** /merchant/user/tag/ensure | Ensure computed user tags are calculated
[**UserTagListGet**](UserTag.md#UserTagListGet) | **Get** /merchant/user/tag/list | List computed user tags for one user
[**UserTagListPost**](UserTag.md#UserTagListPost) | **Post** /merchant/user/tag/list | List computed user tags for one user
[**UserTagListUsersPost**](UserTag.md#UserTagListUsersPost) | **Post** /merchant/user/tag/list_users | List users by computed tag codes



## UserTagEnsurePost

> MerchantUserTagEnsurePost200Response UserTagEnsurePost(ctx).UnibeeApiMerchantUserTagEnsureReq(unibeeApiMerchantUserTagEnsureReq).Execute()

Ensure computed user tags are calculated

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
	unibeeApiMerchantUserTagEnsureReq := *openapiclient.NewUnibeeApiMerchantUserTagEnsureReq(int64(123)) // UnibeeApiMerchantUserTagEnsureReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTag.UserTagEnsurePost(context.Background()).UnibeeApiMerchantUserTagEnsureReq(unibeeApiMerchantUserTagEnsureReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTag.UserTagEnsurePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTagEnsurePost`: MerchantUserTagEnsurePost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTag.UserTagEnsurePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTagEnsurePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserTagEnsureReq** | [**UnibeeApiMerchantUserTagEnsureReq**](UnibeeApiMerchantUserTagEnsureReq.md) |  | 

### Return type

[**MerchantUserTagEnsurePost200Response**](MerchantUserTagEnsurePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTagListGet

> MerchantUserTagListGet200Response UserTagListGet(ctx).UserId(userId).Execute()

List computed user tags for one user

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
	userId := int64(789) // int64 | User id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTag.UserTagListGet(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTag.UserTagListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTagListGet`: MerchantUserTagListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTag.UserTagListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTagListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | User id | 

### Return type

[**MerchantUserTagListGet200Response**](MerchantUserTagListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTagListPost

> MerchantUserTagListGet200Response UserTagListPost(ctx).UnibeeApiMerchantUserTagListReq(unibeeApiMerchantUserTagListReq).Execute()

List computed user tags for one user

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
	unibeeApiMerchantUserTagListReq := *openapiclient.NewUnibeeApiMerchantUserTagListReq(int64(123)) // UnibeeApiMerchantUserTagListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTag.UserTagListPost(context.Background()).UnibeeApiMerchantUserTagListReq(unibeeApiMerchantUserTagListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTag.UserTagListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTagListPost`: MerchantUserTagListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTag.UserTagListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTagListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserTagListReq** | [**UnibeeApiMerchantUserTagListReq**](UnibeeApiMerchantUserTagListReq.md) |  | 

### Return type

[**MerchantUserTagListGet200Response**](MerchantUserTagListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTagListUsersPost

> MerchantUserTagListUsersPost200Response UserTagListUsersPost(ctx).UnibeeApiMerchantUserTagListUsersReq(unibeeApiMerchantUserTagListUsersReq).Execute()

List users by computed tag codes

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
	unibeeApiMerchantUserTagListUsersReq := *openapiclient.NewUnibeeApiMerchantUserTagListUsersReq([]string{"TagCodes_example"}) // UnibeeApiMerchantUserTagListUsersReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTag.UserTagListUsersPost(context.Background()).UnibeeApiMerchantUserTagListUsersReq(unibeeApiMerchantUserTagListUsersReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTag.UserTagListUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTagListUsersPost`: MerchantUserTagListUsersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTag.UserTagListUsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTagListUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserTagListUsersReq** | [**UnibeeApiMerchantUserTagListUsersReq**](UnibeeApiMerchantUserTagListUsersReq.md) |  | 

### Return type

[**MerchantUserTagListUsersPost200Response**](MerchantUserTagListUsersPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

