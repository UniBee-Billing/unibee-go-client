# \MerchantAuthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantAuthSsoLoginOTPPost**](MerchantAuthAPI.md#MerchantAuthSsoLoginOTPPost) | **Post** /merchant/auth/sso/loginOTP | Login OTP
[**MerchantAuthSsoLoginOTPVerifyPost**](MerchantAuthAPI.md#MerchantAuthSsoLoginOTPVerifyPost) | **Post** /merchant/auth/sso/loginOTPVerify | Merchant User OTP Login Verify
[**MerchantAuthSsoLoginPost**](MerchantAuthAPI.md#MerchantAuthSsoLoginPost) | **Post** /merchant/auth/sso/login | Login
[**MerchantAuthSsoPasswordForgetOTPPost**](MerchantAuthAPI.md#MerchantAuthSsoPasswordForgetOTPPost) | **Post** /merchant/auth/sso/passwordForgetOTP | Merchant Password Forget OTP
[**MerchantAuthSsoPasswordForgetOTPVerifyPost**](MerchantAuthAPI.md#MerchantAuthSsoPasswordForgetOTPVerifyPost) | **Post** /merchant/auth/sso/passwordForgetOTPVerify | Merchant Password Forget OTP Verify
[**MerchantAuthSsoRegisterPost**](MerchantAuthAPI.md#MerchantAuthSsoRegisterPost) | **Post** /merchant/auth/sso/register | Merchant Register
[**MerchantAuthSsoRegisterVerifyPost**](MerchantAuthAPI.md#MerchantAuthSsoRegisterVerifyPost) | **Post** /merchant/auth/sso/registerVerify | Merchant Register Verify



## MerchantAuthSsoLoginOTPPost

> MerchantAuthSsoLoginOTPPost200Response MerchantAuthSsoLoginOTPPost(ctx).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()

Login OTP

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
	unibeeApiMerchantAuthLoginOtpReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpReq("Email_example") // UnibeeApiMerchantAuthLoginOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoLoginOTPPost(context.Background()).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoLoginOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoLoginOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoLoginOTPPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoLoginOTPPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthLoginOtpReq** | [**UnibeeApiMerchantAuthLoginOtpReq**](UnibeeApiMerchantAuthLoginOtpReq.md) |  | 

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


## MerchantAuthSsoLoginOTPVerifyPost

> MerchantAuthSsoLoginPost200Response MerchantAuthSsoLoginOTPVerifyPost(ctx).UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq).Execute()

Merchant User OTP Login Verify

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
	unibeeApiMerchantAuthLoginOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthLoginOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoLoginOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoLoginOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoLoginOTPVerifyPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoLoginOTPVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoLoginOTPVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthLoginOtpVerifyReq** | [**UnibeeApiMerchantAuthLoginOtpVerifyReq**](UnibeeApiMerchantAuthLoginOtpVerifyReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginPost200Response**](MerchantAuthSsoLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantAuthSsoLoginPost

> MerchantAuthSsoLoginPost200Response MerchantAuthSsoLoginPost(ctx).UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq).Execute()

Login

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
	unibeeApiMerchantAuthLoginReq := *openapiclient.NewUnibeeApiMerchantAuthLoginReq("Email_example", "Password_example") // UnibeeApiMerchantAuthLoginReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoLoginPost(context.Background()).UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoLoginPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthLoginReq** | [**UnibeeApiMerchantAuthLoginReq**](UnibeeApiMerchantAuthLoginReq.md) |  | 

### Return type

[**MerchantAuthSsoLoginPost200Response**](MerchantAuthSsoLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantAuthSsoPasswordForgetOTPPost

> MerchantAuthSsoLoginOTPPost200Response MerchantAuthSsoPasswordForgetOTPPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()

Merchant Password Forget OTP

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
	unibeeApiMerchantAuthPasswordForgetOtpReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpReq("Email_example") // UnibeeApiMerchantAuthPasswordForgetOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoPasswordForgetOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoPasswordForgetOTPPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthPasswordForgetOtpReq** | [**UnibeeApiMerchantAuthPasswordForgetOtpReq**](UnibeeApiMerchantAuthPasswordForgetOtpReq.md) |  | 

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


## MerchantAuthSsoPasswordForgetOTPVerifyPost

> MerchantAuthSsoLoginOTPPost200Response MerchantAuthSsoPasswordForgetOTPVerifyPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()

Merchant Password Forget OTP Verify

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
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq("Email_example", "NewPassword_example", "VerificationCode_example") // UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoPasswordForgetOTPVerifyPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoPasswordForgetOTPVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoPasswordForgetOTPVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthPasswordForgetOtpVerifyReq** | [**UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq**](UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq.md) |  | 

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


## MerchantAuthSsoRegisterPost

> MerchantAuthSsoLoginOTPPost200Response MerchantAuthSsoRegisterPost(ctx).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()

Merchant Register

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
	unibeeApiMerchantAuthRegisterReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterReq("Email_example", "FirstName_example", "LastName_example", "Password_example") // UnibeeApiMerchantAuthRegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoRegisterPost(context.Background()).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoRegisterPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthRegisterReq** | [**UnibeeApiMerchantAuthRegisterReq**](UnibeeApiMerchantAuthRegisterReq.md) |  | 

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


## MerchantAuthSsoRegisterVerifyPost

> MerchantAuthSsoRegisterVerifyPost200Response MerchantAuthSsoRegisterVerifyPost(ctx).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()

Merchant Register Verify

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
	unibeeApiMerchantAuthRegisterVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthRegisterVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAuthAPI.MerchantAuthSsoRegisterVerifyPost(context.Background()).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAuthAPI.MerchantAuthSsoRegisterVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantAuthSsoRegisterVerifyPost`: MerchantAuthSsoRegisterVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantAuthAPI.MerchantAuthSsoRegisterVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantAuthSsoRegisterVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthRegisterVerifyReq** | [**UnibeeApiMerchantAuthRegisterVerifyReq**](UnibeeApiMerchantAuthRegisterVerifyReq.md) |  | 

### Return type

[**MerchantAuthSsoRegisterVerifyPost200Response**](MerchantAuthSsoRegisterVerifyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

