# \Auth

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthSsoLoginOTPPost**](Auth.md#AuthSsoLoginOTPPost) | **Post** /merchant/auth/sso/loginOTP | LoginOTP
[**AuthSsoLoginOTPVerifyPost**](Auth.md#AuthSsoLoginOTPVerifyPost) | **Post** /merchant/auth/sso/loginOTPVerify | LoginOTPVerify
[**AuthSsoLoginPost**](Auth.md#AuthSsoLoginPost) | **Post** /merchant/auth/sso/login | Login
[**AuthSsoPasswordForgetOTPPost**](Auth.md#AuthSsoPasswordForgetOTPPost) | **Post** /merchant/auth/sso/passwordForgetOTP | PasswordForgetOTP
[**AuthSsoPasswordForgetOTPVerifyPost**](Auth.md#AuthSsoPasswordForgetOTPVerifyPost) | **Post** /merchant/auth/sso/passwordForgetOTPVerify | PasswordForgetOTPVerify
[**AuthSsoRegisterPost**](Auth.md#AuthSsoRegisterPost) | **Post** /merchant/auth/sso/register | Register
[**AuthSsoRegisterVerifyPost**](Auth.md#AuthSsoRegisterVerifyPost) | **Post** /merchant/auth/sso/registerVerify | RegisterVerify



## AuthSsoLoginOTPPost

> MerchantAuthSsoLoginOTPPost200Response AuthSsoLoginOTPPost(ctx).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()

LoginOTP



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
	unibeeApiMerchantAuthLoginOtpReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpReq("Email_example") // UnibeeApiMerchantAuthLoginOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoLoginOTPPost(context.Background()).UnibeeApiMerchantAuthLoginOtpReq(unibeeApiMerchantAuthLoginOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoLoginOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoLoginOTPPost`: %v\n", resp)
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

LoginOTPVerify



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
	unibeeApiMerchantAuthLoginOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthLoginOtpVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthLoginOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoLoginOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthLoginOtpVerifyReq(unibeeApiMerchantAuthLoginOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoLoginOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginOTPVerifyPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoLoginOTPVerifyPost`: %v\n", resp)
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

Login



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
	unibeeApiMerchantAuthLoginReq := *openapiclient.NewUnibeeApiMerchantAuthLoginReq("Email_example", "Password_example") // UnibeeApiMerchantAuthLoginReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoLoginPost(context.Background()).UnibeeApiMerchantAuthLoginReq(unibeeApiMerchantAuthLoginReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoLoginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoLoginPost`: MerchantAuthSsoLoginPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoLoginPost`: %v\n", resp)
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

PasswordForgetOTP



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
	unibeeApiMerchantAuthPasswordForgetOtpReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpReq("Email_example") // UnibeeApiMerchantAuthPasswordForgetOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoPasswordForgetOTPPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpReq(unibeeApiMerchantAuthPasswordForgetOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoPasswordForgetOTPPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordForgetOTPPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoPasswordForgetOTPPost`: %v\n", resp)
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

PasswordForgetOTPVerify



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
	unibeeApiMerchantAuthPasswordForgetOtpVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthPasswordForgetOtpVerifyReq("Email_example", "NewPassword_example", "VerificationCode_example") // UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoPasswordForgetOTPVerifyPost(context.Background()).UnibeeApiMerchantAuthPasswordForgetOtpVerifyReq(unibeeApiMerchantAuthPasswordForgetOtpVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoPasswordForgetOTPVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoPasswordForgetOTPVerifyPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoPasswordForgetOTPVerifyPost`: %v\n", resp)
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
	openapiclient "github.com/UniB-e-e/unibee-go-client"
)

func main() {
	unibeeApiMerchantAuthRegisterReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterReq("Email_example", "FirstName_example", "LastName_example", "Password_example") // UnibeeApiMerchantAuthRegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoRegisterPost(context.Background()).UnibeeApiMerchantAuthRegisterReq(unibeeApiMerchantAuthRegisterReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoRegisterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoRegisterPost`: %v\n", resp)
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

RegisterVerify



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
	unibeeApiMerchantAuthRegisterVerifyReq := *openapiclient.NewUnibeeApiMerchantAuthRegisterVerifyReq("Email_example", "VerificationCode_example") // UnibeeApiMerchantAuthRegisterVerifyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Auth.AuthSsoRegisterVerifyPost(context.Background()).UnibeeApiMerchantAuthRegisterVerifyReq(unibeeApiMerchantAuthRegisterVerifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Auth.AuthSsoRegisterVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthSsoRegisterVerifyPost`: MerchantAuthSsoRegisterVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Auth.AuthSsoRegisterVerifyPost`: %v\n", resp)
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

