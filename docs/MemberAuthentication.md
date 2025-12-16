# \MemberAuthentication

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthSessionLoginPost**](MemberAuthentication.md#AuthSessionLoginPost) | **Post** /merchant/auth/session_login | Session Login
[**AuthSsoClearTotpPost**](MemberAuthentication.md#AuthSsoClearTotpPost) | **Post** /merchant/auth/sso/clear_totp | Admin Member Clear Member 2FA Key With ResumeCode
[**AuthSsoLoginOTPPost**](MemberAuthentication.md#AuthSsoLoginOTPPost) | **Post** /merchant/auth/sso/loginOTP | OTP Login
[**AuthSsoLoginOTPVerifyPost**](MemberAuthentication.md#AuthSsoLoginOTPVerifyPost) | **Post** /merchant/auth/sso/loginOTPVerify | OTP Login Code Verification
[**AuthSsoLoginPost**](MemberAuthentication.md#AuthSsoLoginPost) | **Post** /merchant/auth/sso/login | Merchant Password Login
[**AuthSsoOauthGithubGet**](MemberAuthentication.md#AuthSsoOauthGithubGet) | **Get** /merchant/auth/sso/oauth/github | Get Oauth Github
[**AuthSsoOauthGoogleGet**](MemberAuthentication.md#AuthSsoOauthGoogleGet) | **Get** /merchant/auth/sso/oauth/google | Get Oauth Google
[**AuthSsoOauthLoginPost**](MemberAuthentication.md#AuthSsoOauthLoginPost) | **Post** /merchant/auth/sso/oauth/login | OAuth Login
[**AuthSsoOauthMembersGet**](MemberAuthentication.md#AuthSsoOauthMembersGet) | **Get** /merchant/auth/sso/oauth/members | Get Oauth Members
[**AuthSsoOauthRegisterPost**](MemberAuthentication.md#AuthSsoOauthRegisterPost) | **Post** /merchant/auth/sso/oauth/register | Register OAuth
[**AuthSsoOauthSetupPost**](MemberAuthentication.md#AuthSsoOauthSetupPost) | **Post** /merchant/auth/sso/oauth/setup | OAuth Setup
[**AuthSsoPasswordForgetOTPPost**](MemberAuthentication.md#AuthSsoPasswordForgetOTPPost) | **Post** /merchant/auth/sso/passwordForgetOTP | OTP Password Forget
[**AuthSsoPasswordForgetOTPVerifyPost**](MemberAuthentication.md#AuthSsoPasswordForgetOTPVerifyPost) | **Post** /merchant/auth/sso/passwordForgetOTPVerify | OTP Password Forget Code Verification
[**AuthSsoPasswordForgetTotpVerifyPost**](MemberAuthentication.md#AuthSsoPasswordForgetTotpVerifyPost) | **Post** /merchant/auth/sso/passwordForgetTotpVerify | 2FA Password Forget Code Verification
[**AuthSsoPasswordSetupPost**](MemberAuthentication.md#AuthSsoPasswordSetupPost) | **Post** /merchant/auth/sso/passwordSetup | Password Setup
[**AuthSsoRegisterEmailCheckPost**](MemberAuthentication.md#AuthSsoRegisterEmailCheckPost) | **Post** /merchant/auth/sso/register_email_check | Register Email Check
[**AuthSsoRegisterPost**](MemberAuthentication.md#AuthSsoRegisterPost) | **Post** /merchant/auth/sso/register | Register
[**AuthSsoRegisterVerifyPost**](MemberAuthentication.md#AuthSsoRegisterVerifyPost) | **Post** /merchant/auth/sso/registerVerify | Register Verify
[**AuthSsoRegisterWithoutVerificationPost**](MemberAuthentication.md#AuthSsoRegisterWithoutVerificationPost) | **Post** /merchant/auth/sso/register_without_verification | RegisterWithoutCode



## AuthSessionLoginPost

> MerchantAuthSessionLoginPost200Response AuthSessionLoginPost(ctx).UnibeeApiMerchantAuthSessionReq(unibeeApiMerchantAuthSessionReq).Execute()

Session Login



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
	unibeeApiMerchantAuthSessionReq := *openapiclient.NewUnibeeApiMerchantAuthSessionReq("Session_example") // UnibeeApiMerchantAuthSessionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSessionLoginPost(context.Background()).UnibeeApiMerchantAuthSessionReq(unibeeApiMerchantAuthSessionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSessionLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSessionLoginPost`: MerchantAuthSessionLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSessionLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSessionLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthSessionReq** | [**UnibeeApiMerchantAuthSessionReq**](UnibeeApiMerchantAuthSessionReq.md) |  | 

### Return type

[**MerchantAuthSessionLoginPost200Response**](MerchantAuthSessionLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoClearTotpPost

> MerchantAuthSsoClearTotpPost200Response AuthSsoClearTotpPost(ctx).UnibeeApiMerchantAuthClearTotpReq(unibeeApiMerchantAuthClearTotpReq).Execute()

Admin Member Clear Member 2FA Key With ResumeCode

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
	unibeeApiMerchantAuthClearTotpReq := *openapiclient.NewUnibeeApiMerchantAuthClearTotpReq("Email_example", "TotpResumeCode_example") // UnibeeApiMerchantAuthClearTotpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoClearTotpPost(context.Background()).UnibeeApiMerchantAuthClearTotpReq(unibeeApiMerchantAuthClearTotpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoClearTotpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoClearTotpPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoClearTotpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoClearTotpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthClearTotpReq** | [**UnibeeApiMerchantAuthClearTotpReq**](UnibeeApiMerchantAuthClearTotpReq.md) |  | 

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


## AuthSsoLoginOTPPost

> MerchantAuthSsoClearTotpPost200Response AuthSsoLoginOTPPost(ctx).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()

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
	// response from `AuthSsoLoginOTPPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

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

Merchant Password Login



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


## AuthSsoOauthGithubGet

> MerchantAuthSsoOauthGithubGet200Response AuthSsoOauthGithubGet(ctx).GithubCode(githubCode).Execute()

Get Oauth Github



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
	githubCode := "githubCode_example" // string | Github Code (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthGithubGet(context.Background()).GithubCode(githubCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthGithubGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthGithubGet`: MerchantAuthSsoOauthGithubGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthGithubGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthGithubGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **githubCode** | **string** | Github Code | 

### Return type

[**MerchantAuthSsoOauthGithubGet200Response**](MerchantAuthSsoOauthGithubGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoOauthGoogleGet

> MerchantAuthSsoOauthGoogleGet200Response AuthSsoOauthGoogleGet(ctx).GoogleCode(googleCode).RedirectUri(redirectUri).Execute()

Get Oauth Google



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
	googleCode := "googleCode_example" // string | Google Code (optional)
	redirectUri := "redirectUri_example" // string | The Google Redirect Uri (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthGoogleGet(context.Background()).GoogleCode(googleCode).RedirectUri(redirectUri).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthGoogleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthGoogleGet`: MerchantAuthSsoOauthGoogleGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthGoogleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthGoogleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleCode** | **string** | Google Code | 
 **redirectUri** | **string** | The Google Redirect Uri | 

### Return type

[**MerchantAuthSsoOauthGoogleGet200Response**](MerchantAuthSsoOauthGoogleGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoOauthLoginPost

> MerchantAuthSsoLoginPost200Response AuthSsoOauthLoginPost(ctx).UnibeeApiMerchantAuthLoginOAuthReq(unibeeApiMerchantAuthLoginOAuthReq).Execute()

OAuth Login



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
	unibeeApiMerchantAuthLoginOAuthReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOAuthReq("Email_example") // UnibeeApiMerchantAuthLoginOAuthReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthLoginPost(context.Background()).UnibeeApiMerchantAuthLoginOAuthReq(unibeeApiMerchantAuthLoginOAuthReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthLoginPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthLoginOAuthReq** | [**UnibeeApiMerchantAuthLoginOAuthReq**](UnibeeApiMerchantAuthLoginOAuthReq.md) |  | 

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


## AuthSsoOauthMembersGet

> MerchantAuthSsoOauthMembersGet200Response AuthSsoOauthMembersGet(ctx).Execute()

Get Oauth Members



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
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthMembersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthMembersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthMembersGet`: MerchantAuthSsoOauthMembersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthMembersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthMembersGetRequest struct via the builder pattern


### Return type

[**MerchantAuthSsoOauthMembersGet200Response**](MerchantAuthSsoOauthMembersGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoOauthRegisterPost

> MerchantAuthSsoLoginPost200Response AuthSsoOauthRegisterPost(ctx).UnibeeApiMerchantAuthRegisterOAuthReq(unibeeApiMerchantAuthRegisterOAuthReq).Execute()

Register OAuth



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
	unibeeApiMerchantAuthRegisterOAuthReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterOAuthReq("Email_example") // UnibeeApiMerchantAuthRegisterOAuthReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthRegisterPost(context.Background()).UnibeeApiMerchantAuthRegisterOAuthReq(unibeeApiMerchantAuthRegisterOAuthReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthRegisterPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthRegisterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthRegisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthRegisterOAuthReq** | [**UnibeeApiMerchantAuthRegisterOAuthReq**](UnibeeApiMerchantAuthRegisterOAuthReq.md) |  | 

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


## AuthSsoOauthSetupPost

> MerchantAuthSsoLoginPost200Response AuthSsoOauthSetupPost(ctx).UnibeeApiMerchantAuthSetupOAuthReq(unibeeApiMerchantAuthSetupOAuthReq).Execute()

OAuth Setup



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
	unibeeApiMerchantAuthSetupOAuthReq := *openapiclient.NewUnibeeApiMerchantAuthSetupOAuthReq("Email_example", "SetupToken_example") // UnibeeApiMerchantAuthSetupOAuthReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoOauthSetupPost(context.Background()).UnibeeApiMerchantAuthSetupOAuthReq(unibeeApiMerchantAuthSetupOAuthReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoOauthSetupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoOauthSetupPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoOauthSetupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoOauthSetupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthSetupOAuthReq** | [**UnibeeApiMerchantAuthSetupOAuthReq**](UnibeeApiMerchantAuthSetupOAuthReq.md) |  | 

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

> MerchantAuthSsoClearTotpPost200Response AuthSsoPasswordForgetOTPPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()

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
	// response from `AuthSsoPasswordForgetOTPPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoPasswordForgetOTPVerifyPost

> MerchantAuthSsoClearTotpPost200Response AuthSsoPasswordForgetOTPVerifyPost(ctx).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()

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
	// response from `AuthSsoPasswordForgetOTPVerifyPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoPasswordForgetTotpVerifyPost

> MerchantAuthSsoClearTotpPost200Response AuthSsoPasswordForgetTotpVerifyPost(ctx).UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq(unibeeApiMerchantAuthPasswordForgetTotpVerifyReq).Execute()

2FA Password Forget Code Verification



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
	unibeeApiMerchantAuthPasswordForgetTotpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetTotpVerifyReq("Email_example", "NewPassword_example", "TotpCode_example") // UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoPasswordForgetTotpVerifyPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq(unibeeApiMerchantAuthPasswordForgetTotpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoPasswordForgetTotpVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordForgetTotpVerifyPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoPasswordForgetTotpVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoPasswordForgetTotpVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthPasswordForgetTotpVerifyReq** | [**UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq**](UnibeeApiMerchantAuthPasswordForgetTotpVerifyReq.md) |  | 

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


## AuthSsoPasswordSetupPost

> MerchantAuthSsoLoginPost200Response AuthSsoPasswordSetupPost(ctx).UnibeeApiMerchantAuthPasswordSetupOtpReq(unibeeApiMerchantAuthPasswordSetupOtpReq).Execute()

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
	// response from `AuthSsoPasswordSetupPost`: MerchantAuthSsoLoginPost200Response
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

[**MerchantAuthSsoLoginPost200Response**](MerchantAuthSsoLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoRegisterEmailCheckPost

> MerchantAuthSsoRegisterEmailCheckPost200Response AuthSsoRegisterEmailCheckPost(ctx).UnibeeApiMerchantAuthRegisterEmailCheckReq(unibeeApiMerchantAuthRegisterEmailCheckReq).Execute()

Register Email Check



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
	unibeeApiMerchantAuthRegisterEmailCheckReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterEmailCheckReq("Email_example") // UnibeeApiMerchantAuthRegisterEmailCheckReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoRegisterEmailCheckPost(context.Background()).UnibeeApiMerchantAuthRegisterEmailCheckReq(unibeeApiMerchantAuthRegisterEmailCheckReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoRegisterEmailCheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterEmailCheckPost`: MerchantAuthSsoRegisterEmailCheckPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoRegisterEmailCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoRegisterEmailCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthRegisterEmailCheckReq** | [**UnibeeApiMerchantAuthRegisterEmailCheckReq**](UnibeeApiMerchantAuthRegisterEmailCheckReq.md) |  | 

### Return type

[**MerchantAuthSsoRegisterEmailCheckPost200Response**](MerchantAuthSsoRegisterEmailCheckPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoRegisterPost

> MerchantAuthSsoClearTotpPost200Response AuthSsoRegisterPost(ctx).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()

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
	// response from `AuthSsoRegisterPost`: MerchantAuthSsoClearTotpPost200Response
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

[**MerchantAuthSsoClearTotpPost200Response**](MerchantAuthSsoClearTotpPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoRegisterVerifyPost

> MerchantAuthSsoLoginPost200Response AuthSsoRegisterVerifyPost(ctx).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()

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
	// response from `AuthSsoRegisterVerifyPost`: MerchantAuthSsoLoginPost200Response
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

[**MerchantAuthSsoLoginPost200Response**](MerchantAuthSsoLoginPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthSsoRegisterWithoutVerificationPost

> MerchantAuthSsoLoginPost200Response AuthSsoRegisterWithoutVerificationPost(ctx).UnibeeApiMerchantAuthRegisterWithoutVerificationReq(unibeeApiMerchantAuthRegisterWithoutVerificationReq).Execute()

RegisterWithoutCode



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
	unibeeApiMerchantAuthRegisterWithoutVerificationReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterWithoutVerificationReq("Email_example", "FirstName_example", "LastName_example", "Password_example") // UnibeeApiMerchantAuthRegisterWithoutVerificationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MemberAuthentication.AuthSsoRegisterWithoutVerificationPost(context.Background()).UnibeeApiMerchantAuthRegisterWithoutVerificationReq(unibeeApiMerchantAuthRegisterWithoutVerificationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberAuthentication.AuthSsoRegisterWithoutVerificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterWithoutVerificationPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MemberAuthentication.AuthSsoRegisterWithoutVerificationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthSsoRegisterWithoutVerificationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantAuthRegisterWithoutVerificationReq** | [**UnibeeApiMerchantAuthRegisterWithoutVerificationReq**](UnibeeApiMerchantAuthRegisterWithoutVerificationReq.md) |  | 

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

