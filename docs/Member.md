# \Member

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberListGet**](Member.md#MemberListGet) | **Get** /merchant/member/list | GetMemberList
[**MemberListPost**](Member.md#MemberListPost) | **Post** /merchant/member/list | GetMemberList
[**MemberLogoutPost**](Member.md#MemberLogoutPost) | **Post** /merchant/member/logout | MemberLogout
[**MemberNewMemberPost**](Member.md#MemberNewMemberPost) | **Post** /merchant/member/new_member | Invite member
[**MemberOperationLogListGet**](Member.md#MemberOperationLogListGet) | **Get** /merchant/member/operation_log_list | GetMemberOperationLogList
[**MemberPasswordResetPost**](Member.md#MemberPasswordResetPost) | **Post** /merchant/member/passwordReset | MemberResetPassword
[**MemberProfileGet**](Member.md#MemberProfileGet) | **Get** /merchant/member/profile | GetMemberProfile
[**MemberResumeMemberPost**](Member.md#MemberResumeMemberPost) | **Post** /merchant/member/resume_member | ResumeMember
[**MemberSuspendMemberPost**](Member.md#MemberSuspendMemberPost) | **Post** /merchant/member/suspend_member | SuspendMember
[**MemberUpdateMemberRolePost**](Member.md#MemberUpdateMemberRolePost) | **Post** /merchant/member/update_member_role | UpdateMemberRole



## MemberListGet

> MerchantMemberListGet200Response MemberListGet(ctx).RoleIds(roleIds).Page(page).Count(count).Execute()

GetMemberList

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
	roleIds := []int64{int64(123)} // []int64 | The member's roleId if specified' (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberListGet(context.Background()).RoleIds(roleIds).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberListGet`: MerchantMemberListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleIds** | **[]int64** | The member&#39;s roleId if specified&#39; | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

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


## MemberListPost

> MerchantMemberListGet200Response MemberListPost(ctx).UnibeeApiMerchantMemberListReq(unibeeApiMerchantMemberListReq).Execute()

GetMemberList

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
	unibeeApiMerchantMemberListReq := *openapiclient.NewUnibeeApiMerchantMemberListReq() // UnibeeApiMerchantMemberListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberListPost(context.Background()).UnibeeApiMerchantMemberListReq(unibeeApiMerchantMemberListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberListPost`: MerchantMemberListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberListReq** | [**UnibeeApiMerchantMemberListReq**](UnibeeApiMerchantMemberListReq.md) |  | 

### Return type

[**MerchantMemberListGet200Response**](MerchantMemberListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberLogoutPost

> MerchantAuthSsoLoginOTPPost200Response MemberLogoutPost(ctx).Body(body).Execute()

MemberLogout

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
	openapiclient "github.com/UniBee-Billing/unibee-go-client"
)

func main() {
	unibeeApiMerchantMemberNewMemberReq := *openapiclient.NewUnibeeApiMerchantMemberNewMemberReq("Email_example", []int64{int64(123)}) // UnibeeApiMerchantMemberNewMemberReq | 

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


## MemberOperationLogListGet

> MerchantMemberOperationLogListGet200Response MemberOperationLogListGet(ctx).MemberFirstName(memberFirstName).MemberLastName(memberLastName).MemberEmail(memberEmail).FirstName(firstName).LastName(lastName).Email(email).SubscriptionId(subscriptionId).InvoiceId(invoiceId).PlanId(planId).DiscountCode(discountCode).Page(page).Count(count).Execute()

GetMemberOperationLogList

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
	memberFirstName := "memberFirstName_example" // string | Filter Member's FirstName Default All (optional)
	memberLastName := "memberLastName_example" // string | Filter Member's LastName, Default All (optional)
	memberEmail := "memberEmail_example" // string | Filter Member's Email, Default All (optional)
	firstName := "firstName_example" // string | FirstName (optional)
	lastName := "lastName_example" // string | LastName (optional)
	email := "email_example" // string | Email (optional)
	subscriptionId := "subscriptionId_example" // string | subscription_id (optional)
	invoiceId := "invoiceId_example" // string | invoice id (optional)
	planId := int64(789) // int64 | plan id (optional)
	discountCode := "discountCode_example" // string | discount_code (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberOperationLogListGet(context.Background()).MemberFirstName(memberFirstName).MemberLastName(memberLastName).MemberEmail(memberEmail).FirstName(firstName).LastName(lastName).Email(email).SubscriptionId(subscriptionId).InvoiceId(invoiceId).PlanId(planId).DiscountCode(discountCode).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberOperationLogListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberOperationLogListGet`: MerchantMemberOperationLogListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberOperationLogListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberOperationLogListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **memberFirstName** | **string** | Filter Member&#39;s FirstName Default All | 
 **memberLastName** | **string** | Filter Member&#39;s LastName, Default All | 
 **memberEmail** | **string** | Filter Member&#39;s Email, Default All | 
 **firstName** | **string** | FirstName | 
 **lastName** | **string** | LastName | 
 **email** | **string** | Email | 
 **subscriptionId** | **string** | subscription_id | 
 **invoiceId** | **string** | invoice id | 
 **planId** | **int64** | plan id | 
 **discountCode** | **string** | discount_code | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantMemberOperationLogListGet200Response**](MerchantMemberOperationLogListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberPasswordResetPost

> MerchantAuthSsoLoginOTPPost200Response MemberPasswordResetPost(ctx).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()

MemberResetPassword

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

GetMemberProfile

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


## MemberResumeMemberPost

> MerchantAuthSsoLoginOTPPost200Response MemberResumeMemberPost(ctx).UnibeeApiMerchantMemberReleaseReq(unibeeApiMerchantMemberReleaseReq).Execute()

ResumeMember

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
	unibeeApiMerchantMemberReleaseReq := *openapiclient.NewUnibeeApiMerchantMemberReleaseReq() // UnibeeApiMerchantMemberReleaseReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberResumeMemberPost(context.Background()).UnibeeApiMerchantMemberReleaseReq(unibeeApiMerchantMemberReleaseReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberResumeMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberResumeMemberPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberResumeMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberResumeMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberReleaseReq** | [**UnibeeApiMerchantMemberReleaseReq**](UnibeeApiMerchantMemberReleaseReq.md) |  | 

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


## MemberSuspendMemberPost

> MerchantAuthSsoLoginOTPPost200Response MemberSuspendMemberPost(ctx).UnibeeApiMerchantMemberFrozenReq(unibeeApiMerchantMemberFrozenReq).Execute()

SuspendMember

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
	unibeeApiMerchantMemberFrozenReq := *openapiclient.NewUnibeeApiMerchantMemberFrozenReq() // UnibeeApiMerchantMemberFrozenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Member.MemberSuspendMemberPost(context.Background()).UnibeeApiMerchantMemberFrozenReq(unibeeApiMerchantMemberFrozenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Member.MemberSuspendMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberSuspendMemberPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Member.MemberSuspendMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberSuspendMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberFrozenReq** | [**UnibeeApiMerchantMemberFrozenReq**](UnibeeApiMerchantMemberFrozenReq.md) |  | 

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


## MemberUpdateMemberRolePost

> MerchantAuthSsoLoginOTPPost200Response MemberUpdateMemberRolePost(ctx).UnibeeApiMerchantMemberUpdateMemberRoleReq(unibeeApiMerchantMemberUpdateMemberRoleReq).Execute()

UpdateMemberRole

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

