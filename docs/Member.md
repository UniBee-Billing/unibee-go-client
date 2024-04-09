# \Member

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberListGet**](Member.md#MemberListGet) | **Get** /merchant/member/list | Get Merchant Member List
[**MemberLogoutPost**](Member.md#MemberLogoutPost) | **Post** /merchant/member/logout | Merchant Member Logout
[**MemberNewMemberPost**](Member.md#MemberNewMemberPost) | **Post** /merchant/member/new_member | Invite member
[**MemberPasswordResetPost**](Member.md#MemberPasswordResetPost) | **Post** /merchant/member/passwordReset | Merchant Member Reset Password
[**MemberProfileGet**](Member.md#MemberProfileGet) | **Get** /merchant/member/profile | Get Merchant Member Profile
[**MemberUpdateMemberRolePost**](Member.md#MemberUpdateMemberRolePost) | **Post** /merchant/member/update_member_role | Update Member Role



## MemberListGet

> MerchantMemberListGet200Response MemberListGet(ctx).Execute()

Get Merchant Member List

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
	resp, r, err := apiClient.Member.MemberListGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberListGet`: MerchantMemberListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberListGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMemberListGetRequest struct via the builder pattern


### Return type

[**MerchantMemberListGet200Response**](MerchantMemberListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberLogoutPost

> MerchantAuthSsoLoginOTPPost200Response MemberLogoutPost(ctx).Body(body).Execute()

Merchant Member Logout

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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberLogoutPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberLogoutPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberLogoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberLogoutPostRequest struct via the builder pattern


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


## MemberNewMemberPost

> MerchantAuthSsoLoginOTPPost200Response MemberNewMemberPost(ctx).UnibeeApiMerchantMemberNewMemberReq(unibeeApiMerchantMemberNewMemberReq).Execute()

Invite member



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
	unibeeApiMerchantMemberNewMemberReq := *openapiclient.NewUnibeeApiMerchantMemberNewMemberReq("Email_example", "Role_example") // UnibeeApiMerchantMemberNewMemberReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberNewMemberPost(context.Background()).UnibeeApiMerchantMemberNewMemberReq(unibeeApiMerchantMemberNewMemberReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberNewMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberNewMemberPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberNewMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberNewMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberNewMemberReq** | [**UnibeeApiMerchantMemberNewMemberReq**](UnibeeApiMerchantMemberNewMemberReq.md) |  | 

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


## MemberPasswordResetPost

> MerchantAuthSsoLoginOTPPost200Response MemberPasswordResetPost(ctx).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()

Merchant Member Reset Password

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
	unibeeApiMerchantMemberPasswordResetReq := *openapiclient.NewUnibeeApiMerchantMemberPasswordResetReq("NewPassword_example", "OldPassword_example") // UnibeeApiMerchantMemberPasswordResetReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberPasswordResetPost(context.Background()).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberPasswordResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberPasswordResetPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberPasswordResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberPasswordResetPostRequest struct via the builder pattern


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


## MemberProfileGet

> MerchantAuthSsoRegisterVerifyPost200Response MemberProfileGet(ctx).Execute()

Get Merchant Member Profile

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
	resp, r, err := apiClient.Member.MemberProfileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberProfileGet`: MerchantAuthSsoRegisterVerifyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberProfileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMemberProfileGetRequest struct via the builder pattern


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


## MemberUpdateMemberRolePost

> MerchantAuthSsoLoginOTPPost200Response MemberUpdateMemberRolePost(ctx).UnibeeApiMerchantMemberUpdateMemberRoleReq(unibeeApiMerchantMemberUpdateMemberRoleReq).Execute()

Update Member Role

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
	unibeeApiMerchantMemberUpdateMemberRoleReq := *openapiclient.NewUnibeeApiMerchantMemberUpdateMemberRoleReq() // UnibeeApiMerchantMemberUpdateMemberRoleReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberUpdateMemberRolePost(context.Background()).UnibeeApiMerchantMemberUpdateMemberRoleReq(unibeeApiMerchantMemberUpdateMemberRoleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberUpdateMemberRolePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberUpdateMemberRolePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberUpdateMemberRolePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberUpdateMemberRolePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberUpdateMemberRoleReq** | [**UnibeeApiMerchantMemberUpdateMemberRoleReq**](UnibeeApiMerchantMemberUpdateMemberRoleReq.md) |  | 

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

