# \AdminMember

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberClearMemberTotpPost**](AdminMember.md#MemberClearMemberTotpPost) | **Post** /merchant/member/clear_member_totp | Admin Owner Clear Member 2FA Key
[**MemberClearOauthPost**](AdminMember.md#MemberClearOauthPost) | **Post** /merchant/member/clear_oauth | Clear Member OAuth Account
[**MemberConfirmTotpKeyPost**](AdminMember.md#MemberConfirmTotpKeyPost) | **Post** /merchant/member/confirm_totp_key | Admin Member Confirm 2FA Key
[**MemberDeleteTotpDevicePost**](AdminMember.md#MemberDeleteTotpDevicePost) | **Post** /merchant/member/delete_totp_device | Admin Owner Or Admin Delete Own&#39;s&#39; 2FA Device
[**MemberGetTotpKeyPost**](AdminMember.md#MemberGetTotpKeyPost) | **Post** /merchant/member/get_totp_key | Admin Member Get 2FA Key
[**MemberListGet**](AdminMember.md#MemberListGet) | **Get** /merchant/member/list | Get Member List
[**MemberListPost**](AdminMember.md#MemberListPost) | **Post** /merchant/member/list | Get Member List
[**MemberLogoutPost**](AdminMember.md#MemberLogoutPost) | **Post** /merchant/member/logout | Logout
[**MemberNewMemberPost**](AdminMember.md#MemberNewMemberPost) | **Post** /merchant/member/new_member | Invite member
[**MemberOperationLogListGet**](AdminMember.md#MemberOperationLogListGet) | **Get** /merchant/member/operation_log_list | Get Member Operation Log List
[**MemberPasswordResetPost**](AdminMember.md#MemberPasswordResetPost) | **Post** /merchant/member/passwordReset | Member Reset Password
[**MemberProfileGet**](AdminMember.md#MemberProfileGet) | **Get** /merchant/member/profile | Get Member Profile
[**MemberResetTotpPost**](AdminMember.md#MemberResetTotpPost) | **Post** /merchant/member/reset_totp | Admin Member Reset Member 2FA Key
[**MemberResumeMemberPost**](AdminMember.md#MemberResumeMemberPost) | **Post** /merchant/member/resume_member | Resume Member
[**MemberSuspendMemberPost**](AdminMember.md#MemberSuspendMemberPost) | **Post** /merchant/member/suspend_member | Suspend Member
[**MemberUpdateMemberRolePost**](AdminMember.md#MemberUpdateMemberRolePost) | **Post** /merchant/member/update_member_role | Update Member Role
[**MemberUpdateOauthPost**](AdminMember.md#MemberUpdateOauthPost) | **Post** /merchant/member/update_oauth | Update Member OAuth Account
[**MemberUpdatePost**](AdminMember.md#MemberUpdatePost) | **Post** /merchant/member/update | Update Member Profile



## MemberClearMemberTotpPost

> MerchantAuthSsoClearTotpPost200Response MemberClearMemberTotpPost(ctx).UnibeeApiMerchantMemberClearTotpReq(unibeeApiMerchantMemberClearTotpReq).Execute()

Admin Owner Clear Member 2FA Key

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
	unibeeApiMerchantMemberClearTotpReq := *openapiclient.NewUnibeeApiMerchantMemberClearTotpReq() // UnibeeApiMerchantMemberClearTotpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberClearMemberTotpPost(context.Background()).UnibeeApiMerchantMemberClearTotpReq(unibeeApiMerchantMemberClearTotpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberClearMemberTotpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberClearMemberTotpPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberClearMemberTotpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberClearMemberTotpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberClearTotpReq** | [**UnibeeApiMerchantMemberClearTotpReq**](UnibeeApiMerchantMemberClearTotpReq.md) |  | 

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


## MemberClearOauthPost

> MerchantAuthSsoClearTotpPost200Response MemberClearOauthPost(ctx).UnibeeApiMerchantMemberClearOAuthReq(unibeeApiMerchantMemberClearOAuthReq).Execute()

Clear Member OAuth Account

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
	unibeeApiMerchantMemberClearOAuthReq := *openapiclient.NewUnibeeApiMerchantMemberClearOAuthReq("Provider_example") // UnibeeApiMerchantMemberClearOAuthReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberClearOauthPost(context.Background()).UnibeeApiMerchantMemberClearOAuthReq(unibeeApiMerchantMemberClearOAuthReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberClearOauthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberClearOauthPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberClearOauthPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberClearOauthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberClearOAuthReq** | [**UnibeeApiMerchantMemberClearOAuthReq**](UnibeeApiMerchantMemberClearOAuthReq.md) |  | 

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


## MemberConfirmTotpKeyPost

> MerchantAuthSsoClearTotpPost200Response MemberConfirmTotpKeyPost(ctx).UnibeeApiMerchantMemberConfirmTotpKeyReq(unibeeApiMerchantMemberConfirmTotpKeyReq).Execute()

Admin Member Confirm 2FA Key

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
	unibeeApiMerchantMemberConfirmTotpKeyReq := *openapiclient.NewUnibeeApiMerchantMemberConfirmTotpKeyReq() // UnibeeApiMerchantMemberConfirmTotpKeyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberConfirmTotpKeyPost(context.Background()).UnibeeApiMerchantMemberConfirmTotpKeyReq(unibeeApiMerchantMemberConfirmTotpKeyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberConfirmTotpKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberConfirmTotpKeyPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberConfirmTotpKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberConfirmTotpKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberConfirmTotpKeyReq** | [**UnibeeApiMerchantMemberConfirmTotpKeyReq**](UnibeeApiMerchantMemberConfirmTotpKeyReq.md) |  | 

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


## MemberDeleteTotpDevicePost

> MerchantMemberDeleteTotpDevicePost200Response MemberDeleteTotpDevicePost(ctx).UnibeeApiMerchantMemberDeleteDeviceReq(unibeeApiMerchantMemberDeleteDeviceReq).Execute()

Admin Owner Or Admin Delete Own's' 2FA Device

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
	unibeeApiMerchantMemberDeleteDeviceReq := *openapiclient.NewUnibeeApiMerchantMemberDeleteDeviceReq() // UnibeeApiMerchantMemberDeleteDeviceReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberDeleteTotpDevicePost(context.Background()).UnibeeApiMerchantMemberDeleteDeviceReq(unibeeApiMerchantMemberDeleteDeviceReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberDeleteTotpDevicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberDeleteTotpDevicePost`: MerchantMemberDeleteTotpDevicePost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberDeleteTotpDevicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberDeleteTotpDevicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberDeleteDeviceReq** | [**UnibeeApiMerchantMemberDeleteDeviceReq**](UnibeeApiMerchantMemberDeleteDeviceReq.md) |  | 

### Return type

[**MerchantMemberDeleteTotpDevicePost200Response**](MerchantMemberDeleteTotpDevicePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberGetTotpKeyPost

> MerchantMemberGetTotpKeyPost200Response MemberGetTotpKeyPost(ctx).UnibeeApiMerchantMemberGetTotpKeyReq(unibeeApiMerchantMemberGetTotpKeyReq).Execute()

Admin Member Get 2FA Key

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
	unibeeApiMerchantMemberGetTotpKeyReq := *openapiclient.NewUnibeeApiMerchantMemberGetTotpKeyReq() // UnibeeApiMerchantMemberGetTotpKeyReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberGetTotpKeyPost(context.Background()).UnibeeApiMerchantMemberGetTotpKeyReq(unibeeApiMerchantMemberGetTotpKeyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberGetTotpKeyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberGetTotpKeyPost`: MerchantMemberGetTotpKeyPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberGetTotpKeyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberGetTotpKeyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberGetTotpKeyReq** | [**UnibeeApiMerchantMemberGetTotpKeyReq**](UnibeeApiMerchantMemberGetTotpKeyReq.md) |  | 

### Return type

[**MerchantMemberGetTotpKeyPost200Response**](MerchantMemberGetTotpKeyPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberListGet

> MerchantMemberListGet200Response MemberListGet(ctx).Email(email).RoleIds(roleIds).Page(page).Count(count).Execute()

Get Member List

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
	email := "email_example" // string | Search Filter Email (optional)
	roleIds := []int64{int64(123)} // []int64 | The member roleId if specified' (optional)
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberListGet(context.Background()).Email(email).RoleIds(roleIds).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberListGet`: MerchantMemberListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | Search Filter Email | 
 **roleIds** | **[]int64** | The member roleId if specified&#39; | 
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

Get Member List

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
	resp, r, err := apiClient.AdminMember.MemberListPost(context.Background()).UnibeeApiMerchantMemberListReq(unibeeApiMerchantMemberListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberListPost`: MerchantMemberListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberListPost`: %v\n", resp)
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

> MerchantAuthSsoClearTotpPost200Response MemberLogoutPost(ctx).Body(body).Execute()

Logout

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
	resp, r, err := apiClient.AdminMember.MemberLogoutPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberLogoutPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberLogoutPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberLogoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## MemberNewMemberPost

> MerchantAuthSsoClearTotpPost200Response MemberNewMemberPost(ctx).UnibeeApiMerchantMemberNewMemberReq(unibeeApiMerchantMemberNewMemberReq).Execute()

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
	resp, r, err := apiClient.AdminMember.MemberNewMemberPost(context.Background()).UnibeeApiMerchantMemberNewMemberReq(unibeeApiMerchantMemberNewMemberReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberNewMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberNewMemberPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberNewMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberNewMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberNewMemberReq** | [**UnibeeApiMerchantMemberNewMemberReq**](UnibeeApiMerchantMemberNewMemberReq.md) |  | 

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


## MemberOperationLogListGet

> MerchantMemberOperationLogListGet200Response MemberOperationLogListGet(ctx).MemberFirstName(memberFirstName).MemberLastName(memberLastName).MemberEmail(memberEmail).FirstName(firstName).LastName(lastName).Email(email).SubscriptionId(subscriptionId).InvoiceId(invoiceId).PlanId(planId).DiscountCode(discountCode).Page(page).Count(count).Execute()

Get Member Operation Log List

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
	resp, r, err := apiClient.AdminMember.MemberOperationLogListGet(context.Background()).MemberFirstName(memberFirstName).MemberLastName(memberLastName).MemberEmail(memberEmail).FirstName(firstName).LastName(lastName).Email(email).SubscriptionId(subscriptionId).InvoiceId(invoiceId).PlanId(planId).DiscountCode(discountCode).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberOperationLogListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberOperationLogListGet`: MerchantMemberOperationLogListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberOperationLogListGet`: %v\n", resp)
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

> MerchantAuthSsoClearTotpPost200Response MemberPasswordResetPost(ctx).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()

Member Reset Password

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
	resp, r, err := apiClient.AdminMember.MemberPasswordResetPost(context.Background()).UnibeeApiMerchantMemberPasswordResetReq(unibeeApiMerchantMemberPasswordResetReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberPasswordResetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberPasswordResetPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberPasswordResetPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberPasswordResetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberPasswordResetReq** | [**UnibeeApiMerchantMemberPasswordResetReq**](UnibeeApiMerchantMemberPasswordResetReq.md) |  | 

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


## MemberProfileGet

> MerchantMemberProfileGet200Response MemberProfileGet(ctx).Execute()

Get Member Profile

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
	resp, r, err := apiClient.AdminMember.MemberProfileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberProfileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberProfileGet`: MerchantMemberProfileGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberProfileGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMemberProfileGetRequest struct via the builder pattern


### Return type

[**MerchantMemberProfileGet200Response**](MerchantMemberProfileGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberResetTotpPost

> MerchantMemberProfileGet200Response MemberResetTotpPost(ctx).UnibeeApiMerchantMemberResetTotpReq(unibeeApiMerchantMemberResetTotpReq).Execute()

Admin Member Reset Member 2FA Key

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
	unibeeApiMerchantMemberResetTotpReq := *openapiclient.NewUnibeeApiMerchantMemberResetTotpReq() // UnibeeApiMerchantMemberResetTotpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberResetTotpPost(context.Background()).UnibeeApiMerchantMemberResetTotpReq(unibeeApiMerchantMemberResetTotpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberResetTotpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberResetTotpPost`: MerchantMemberProfileGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberResetTotpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberResetTotpPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberResetTotpReq** | [**UnibeeApiMerchantMemberResetTotpReq**](UnibeeApiMerchantMemberResetTotpReq.md) |  | 

### Return type

[**MerchantMemberProfileGet200Response**](MerchantMemberProfileGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberResumeMemberPost

> MerchantAuthSsoClearTotpPost200Response MemberResumeMemberPost(ctx).UnibeeApiMerchantMemberReleaseReq(unibeeApiMerchantMemberReleaseReq).Execute()

Resume Member

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
	resp, r, err := apiClient.AdminMember.MemberResumeMemberPost(context.Background()).UnibeeApiMerchantMemberReleaseReq(unibeeApiMerchantMemberReleaseReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberResumeMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberResumeMemberPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberResumeMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberResumeMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberReleaseReq** | [**UnibeeApiMerchantMemberReleaseReq**](UnibeeApiMerchantMemberReleaseReq.md) |  | 

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


## MemberSuspendMemberPost

> MerchantAuthSsoClearTotpPost200Response MemberSuspendMemberPost(ctx).UnibeeApiMerchantMemberFrozenReq(unibeeApiMerchantMemberFrozenReq).Execute()

Suspend Member

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
	resp, r, err := apiClient.AdminMember.MemberSuspendMemberPost(context.Background()).UnibeeApiMerchantMemberFrozenReq(unibeeApiMerchantMemberFrozenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberSuspendMemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberSuspendMemberPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberSuspendMemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberSuspendMemberPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberFrozenReq** | [**UnibeeApiMerchantMemberFrozenReq**](UnibeeApiMerchantMemberFrozenReq.md) |  | 

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


## MemberUpdateMemberRolePost

> MerchantAuthSsoClearTotpPost200Response MemberUpdateMemberRolePost(ctx).UnibeeApiMerchantMemberUpdateMemberRoleReq(unibeeApiMerchantMemberUpdateMemberRoleReq).Execute()

Update Member Role

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
	resp, r, err := apiClient.AdminMember.MemberUpdateMemberRolePost(context.Background()).UnibeeApiMerchantMemberUpdateMemberRoleReq(unibeeApiMerchantMemberUpdateMemberRoleReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberUpdateMemberRolePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberUpdateMemberRolePost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberUpdateMemberRolePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberUpdateMemberRolePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberUpdateMemberRoleReq** | [**UnibeeApiMerchantMemberUpdateMemberRoleReq**](UnibeeApiMerchantMemberUpdateMemberRoleReq.md) |  | 

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


## MemberUpdateOauthPost

> MerchantAuthSsoClearTotpPost200Response MemberUpdateOauthPost(ctx).Body(body).Execute()

Update Member OAuth Account



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
	resp, r, err := apiClient.AdminMember.MemberUpdateOauthPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberUpdateOauthPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberUpdateOauthPost`: MerchantAuthSsoClearTotpPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberUpdateOauthPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberUpdateOauthPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

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


## MemberUpdatePost

> MerchantMemberProfileGet200Response MemberUpdatePost(ctx).UnibeeApiMerchantMemberUpdateReq(unibeeApiMerchantMemberUpdateReq).Execute()

Update Member Profile

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
	unibeeApiMerchantMemberUpdateReq := *openapiclient.NewUnibeeApiMerchantMemberUpdateReq() // UnibeeApiMerchantMemberUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminMember.MemberUpdatePost(context.Background()).UnibeeApiMerchantMemberUpdateReq(unibeeApiMerchantMemberUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminMember.MemberUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberUpdatePost`: MerchantMemberProfileGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AdminMember.MemberUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMemberUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantMemberUpdateReq** | [**UnibeeApiMerchantMemberUpdateReq**](UnibeeApiMerchantMemberUpdateReq.md) |  | 

### Return type

[**MerchantMemberProfileGet200Response**](MerchantMemberProfileGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

