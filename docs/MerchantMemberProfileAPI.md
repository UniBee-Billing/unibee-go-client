# \MerchantMemberProfileAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantMemberLogoutPost**](MerchantMemberProfileAPI.md#MerchantMemberLogoutPost) | **Post** /merchant/member/logout | Merchant Member Logout
[**MerchantMemberPasswordResetPost**](MerchantMemberProfileAPI.md#MerchantMemberPasswordResetPost) | **Post** /merchant/member/passwordReset | Merchant Member Reset Password
[**MerchantMemberProfileGet**](MerchantMemberProfileAPI.md#MerchantMemberProfileGet) | **Get** /merchant/member/profile | Get Merchant Member Profile



## MerchantMemberLogoutPost

> MerchantAuthSsoLoginOTPPost200Response MerchantMemberLogoutPost(ctx).Body(body).Execute()

Merchant Member Logout

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMemberProfileAPI.MerchantMemberLogoutPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMemberProfileAPI.MerchantMemberLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMemberLogoutPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMemberProfileAPI.MerchantMemberLogoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMemberLogoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## MerchantMemberPasswordResetPost

> MerchantAuthSsoLoginOTPPost200Response MerchantMemberPasswordResetPost(ctx).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()

Merchant Member Reset Password

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
	unibeeApiMerchantMemberPasswordResetReq := *openapiclient.NewUnibeeApiMerchantMemberPasswordResetReq("NewPassword_example", "OldPassword_example") // UnibeeApiMerchantMemberPasswordResetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMemberProfileAPI.MerchantMemberPasswordResetPost(context.Background()).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMemberProfileAPI.MerchantMemberPasswordResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMemberPasswordResetPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMemberProfileAPI.MerchantMemberPasswordResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMemberPasswordResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberPasswordResetReq** | [**UnibeeApiMerchantMemberPasswordResetReq**](UnibeeApiMerchantMemberPasswordResetReq.md) |  | 

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


## MerchantMemberProfileGet

> MerchantAuthSsoRegisterVerifyPost200Response MerchantMemberProfileGet(ctx).Execute()

Get Merchant Member Profile

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantMemberProfileAPI.MerchantMemberProfileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantMemberProfileAPI.MerchantMemberProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantMemberProfileGet`: MerchantAuthSsoRegisterVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantMemberProfileAPI.MerchantMemberProfileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMerchantMemberProfileGetRequest struct via the builder pattern


### Return type

[**MerchantAuthSsoRegisterVerifyPost200Response**](MerchantAuthSsoRegisterVerifyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

