# \MemberAuthentication

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthSsoLoginOTPPost**](MemberAuthentication.md#AuthSsoLoginOTPPost) | **Post** /merchant/auth/sso/loginOTP | OTP Login
[**AuthSsoLoginOTPVerifyPost**](MemberAuthentication.md#AuthSsoLoginOTPVerifyPost) | **Post** /merchant/auth/sso/loginOTPVerify | OTP Login Code Verification
[**AuthSsoLoginPost**](MemberAuthentication.md#AuthSsoLoginPost) | **Post** /merchant/auth/sso/login | Password Login
[**AuthSsoPasswordForgetOTPPost**](MemberAuthentication.md#AuthSsoPasswordForgetOTPPost) | **Post** /merchant/auth/sso/passwordForgetOTP | OTP Password Forget
[**AuthSsoPasswordForgetOTPVerifyPost**](MemberAuthentication.md#AuthSsoPasswordForgetOTPVerifyPost) | **Post** /merchant/auth/sso/passwordForgetOTPVerify | OTP Password Forget Code Verification
[**AuthSsoPasswordSetupPost**](MemberAuthentication.md#AuthSsoPasswordSetupPost) | **Post** /merchant/auth/sso/passwordSetup | Password Setup
[**AuthSsoRegisterPost**](MemberAuthentication.md#AuthSsoRegisterPost) | **Post** /merchant/auth/sso/register | Register
[**AuthSsoRegisterVerifyPost**](MemberAuthentication.md#AuthSsoRegisterVerifyPost) | **Post** /merchant/auth/sso/registerVerify | Register Verify



## AuthSsoLoginOTPPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoLoginOTPPost(ctx).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()

OTP Login



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
	unibeeApiMerchantAuthLoginOtpReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpReq("Email_example") // UnibeeApiMerchantAuthLoginOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoLoginOTPPost(context.Background()).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoLoginOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoLoginOTPPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoLoginOTPPostRequest struct via the builder pattern


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


## AuthSsoLoginOTPVerifyPost

> MerchantAuthSsoLoginPost200Response AuthSsoLoginOTPVerifyPost(ctx).UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq).Execute()

OTP Login Code Verification



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
	unibeeApiMerchantAuthLoginOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthLoginOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoLoginOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoLoginOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginOTPVerifyPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoLoginOTPVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoLoginOTPVerifyPostRequest struct via the builder pattern


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


## AuthSsoLoginPost

> MerchantAuthSsoLoginPost200Response AuthSsoLoginPost(ctx).UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq).Execute()

Password Login



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
	unibeeApiMerchantAuthLoginReq := *openapiclient.NewUnibeeApiMerchantAuthLoginReq("Email_example", "Password_example") // UnibeeApiMerchantAuthLoginReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoLoginPost(context.Background()).UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoLoginPostRequest struct via the builder pattern


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


## AuthSsoPasswordForgetOTPPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoPasswordForgetOTPPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()

OTP Password Forget



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
	unibeeApiMerchantAuthPasswordForgetOtpReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpReq("Email_example") // UnibeeApiMerchantAuthPasswordForgetOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoPasswordForgetOTPPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoPasswordForgetOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordForgetOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoPasswordForgetOTPPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoPasswordForgetOTPPostRequest struct via the builder pattern


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


## AuthSsoPasswordForgetOTPVerifyPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoPasswordForgetOTPVerifyPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()

OTP Password Forget Code Verification



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
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq("Email_example", "NewPassword_example", "VerificationCode_example") // UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoPasswordForgetOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoPasswordForgetOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordForgetOTPVerifyPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoPasswordForgetOTPVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoPasswordForgetOTPVerifyPostRequest struct via the builder pattern


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


## AuthSsoPasswordSetupPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoPasswordSetupPost(ctx).UnibeeApiMerchantAuthPasswordSetupOtpReq(unibeeApiMerchantAuthPasswordSetupOtpReq).Execute()

Password Setup



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
	unibeeApiMerchantAuthPasswordSetupOtpReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordSetupOtpReq("Email_example", "NewPassword_example", "SetupToken_example") // UnibeeApiMerchantAuthPasswordSetupOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoPasswordSetupPost(context.Background()).UnibeeApiMerchantAuthPasswordSetupOtpReq(unibeeApiMerchantAuthPasswordSetupOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoPasswordSetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordSetupPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoPasswordSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoPasswordSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthPasswordSetupOtpReq** | [**UnibeeApiMerchantAuthPasswordSetupOtpReq**](UnibeeApiMerchantAuthPasswordSetupOtpReq.md) |  | 

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


## AuthSsoRegisterPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoRegisterPost(ctx).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()

Register



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
	unibeeApiMerchantAuthRegisterReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterReq("Email_example", "FirstName_example", "LastName_example", "Password_example") // UnibeeApiMerchantAuthRegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoRegisterPost(context.Background()).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoRegisterPostRequest struct via the builder pattern


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


## AuthSsoRegisterVerifyPost

> MerchantAuthSsoRegisterVerifyPost200Response AuthSsoRegisterVerifyPost(ctx).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()

Register Verify



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
	unibeeApiMerchantAuthRegisterVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthRegisterVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoRegisterVerifyPost(context.Background()).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoRegisterVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterVerifyPost`: MerchantAuthSsoRegisterVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoRegisterVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoRegisterVerifyPostRequest struct via the builder pattern


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

