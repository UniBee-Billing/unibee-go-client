# \User

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAdminNoteListGet**](User.md#UserAdminNoteListGet) | **Get** /merchant/user/admin_note_list | Get User Admin Note List
[**UserAdminNoteListPost**](User.md#UserAdminNoteListPost) | **Post** /merchant/user/admin_note_list | Get User Admin Note List
[**UserChangeEmailPost**](User.md#UserChangeEmailPost) | **Post** /merchant/user/change_email | Change User Email
[**UserChangeGatewayPost**](User.md#UserChangeGatewayPost) | **Post** /merchant/user/change_gateway | Change User Default Gateway
[**UserClearAutoChargeMethodPost**](User.md#UserClearAutoChargeMethodPost) | **Post** /merchant/user/clear_auto_charge_method | Clear AutoCharge Method
[**UserGetGet**](User.md#UserGetGet) | **Get** /merchant/user/get | Get User Profile
[**UserListGet**](User.md#UserListGet) | **Get** /merchant/user/list | User List
[**UserListPost**](User.md#UserListPost) | **Post** /merchant/user/list | User List
[**UserNewAdminNotePost**](User.md#UserNewAdminNotePost) | **Post** /merchant/user/new_admin_note | New Admin Note
[**UserNewPost**](User.md#UserNewPost) | **Post** /merchant/user/new | New User
[**UserResumeUserPost**](User.md#UserResumeUserPost) | **Post** /merchant/user/resume_user | Resume User
[**UserSearchGet**](User.md#UserSearchGet) | **Get** /merchant/user/search | User Search
[**UserSearchPost**](User.md#UserSearchPost) | **Post** /merchant/user/search | User Search
[**UserSuspendUserPost**](User.md#UserSuspendUserPost) | **Post** /merchant/user/suspend_user | Suspend User
[**UserUpdatePost**](User.md#UserUpdatePost) | **Post** /merchant/user/update | Update User Profile



## UserAdminNoteListGet

> MerchantUserAdminNoteListGet200Response UserAdminNoteListGet(ctx).UserId(userId).Page(page).Count(count).Execute()

Get User Admin Note List

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
	userId := int64(789) // int64 | The id of user, either ExternalUserId or UserId needed
	page := int32(56) // int32 | Page, Start With 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserAdminNoteListGet(context.Background()).UserId(userId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserAdminNoteListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAdminNoteListGet`: MerchantUserAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserAdminNoteListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAdminNoteListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | The id of user, either ExternalUserId or UserId needed | 
 **page** | **int32** | Page, Start With 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantUserAdminNoteListGet200Response**](MerchantUserAdminNoteListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAdminNoteListPost

> MerchantUserAdminNoteListGet200Response UserAdminNoteListPost(ctx).UnibeeApiMerchantUserAdminNoteListReq(unibeeApiMerchantUserAdminNoteListReq).Execute()

Get User Admin Note List

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
	unibeeApiMerchantUserAdminNoteListReq := *openapiclient.NewUnibeeApiMerchantUserAdminNoteListReq(int64(123)) // UnibeeApiMerchantUserAdminNoteListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserAdminNoteListPost(context.Background()).UnibeeApiMerchantUserAdminNoteListReq(unibeeApiMerchantUserAdminNoteListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserAdminNoteListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAdminNoteListPost`: MerchantUserAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserAdminNoteListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAdminNoteListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserAdminNoteListReq** | [**UnibeeApiMerchantUserAdminNoteListReq**](UnibeeApiMerchantUserAdminNoteListReq.md) |  | 

### Return type

[**MerchantUserAdminNoteListGet200Response**](MerchantUserAdminNoteListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChangeEmailPost

> MerchantAuthSsoLoginOTPPost200Response UserChangeEmailPost(ctx).UnibeeApiMerchantUserChangeEmailReq(unibeeApiMerchantUserChangeEmailReq).Execute()

Change User Email

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
	unibeeApiMerchantUserChangeEmailReq := *openapiclient.NewUnibeeApiMerchantUserChangeEmailReq("NewEmail_example") // UnibeeApiMerchantUserChangeEmailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserChangeEmailPost(context.Background()).UnibeeApiMerchantUserChangeEmailReq(unibeeApiMerchantUserChangeEmailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserChangeEmailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChangeEmailPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserChangeEmailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChangeEmailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserChangeEmailReq** | [**UnibeeApiMerchantUserChangeEmailReq**](UnibeeApiMerchantUserChangeEmailReq.md) |  | 

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


## UserChangeGatewayPost

> MerchantAuthSsoLoginOTPPost200Response UserChangeGatewayPost(ctx).UnibeeApiMerchantUserChangeGatewayReq(unibeeApiMerchantUserChangeGatewayReq).Execute()

Change User Default Gateway

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
	unibeeApiMerchantUserChangeGatewayReq := *openapiclient.NewUnibeeApiMerchantUserChangeGatewayReq(int64(123), int64(123)) // UnibeeApiMerchantUserChangeGatewayReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserChangeGatewayPost(context.Background()).UnibeeApiMerchantUserChangeGatewayReq(unibeeApiMerchantUserChangeGatewayReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserChangeGatewayPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChangeGatewayPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserChangeGatewayPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChangeGatewayPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserChangeGatewayReq** | [**UnibeeApiMerchantUserChangeGatewayReq**](UnibeeApiMerchantUserChangeGatewayReq.md) |  | 

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


## UserClearAutoChargeMethodPost

> MerchantAuthSsoLoginOTPPost200Response UserClearAutoChargeMethodPost(ctx).UnibeeApiMerchantUserClearAutoChargeMethodReq(unibeeApiMerchantUserClearAutoChargeMethodReq).Execute()

Clear AutoCharge Method

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
	unibeeApiMerchantUserClearAutoChargeMethodReq := *openapiclient.NewUnibeeApiMerchantUserClearAutoChargeMethodReq(int64(123)) // UnibeeApiMerchantUserClearAutoChargeMethodReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserClearAutoChargeMethodPost(context.Background()).UnibeeApiMerchantUserClearAutoChargeMethodReq(unibeeApiMerchantUserClearAutoChargeMethodReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserClearAutoChargeMethodPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserClearAutoChargeMethodPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserClearAutoChargeMethodPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserClearAutoChargeMethodPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserClearAutoChargeMethodReq** | [**UnibeeApiMerchantUserClearAutoChargeMethodReq**](UnibeeApiMerchantUserClearAutoChargeMethodReq.md) |  | 

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


## UserGetGet

> MerchantUserGetGet200Response UserGetGet(ctx).UserId(userId).Execute()

Get User Profile

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
	userId := int64(789) // int64 | UserId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserGetGet(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserGetGet`: MerchantUserGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | UserId | 

### Return type

[**MerchantUserGetGet200Response**](MerchantUserGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserListGet

> MerchantUserListGet200Response UserListGet(ctx).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).PlanIds(planIds).SubscriptionId(subscriptionId).SubStatus(subStatus).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

User List

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
	userId := int64(789) // int64 | Filter UserId (optional)
	firstName := "firstName_example" // string | Search FirstName (optional)
	lastName := "lastName_example" // string | Search LastName (optional)
	email := "email_example" // string | Search Filter Email (optional)
	planIds := []int32{int32(123)} // []int32 | PlanIds, Search Filter PlanIds (optional)
	subscriptionId := "subscriptionId_example" // string | Search Filter SubscriptionId (optional)
	subStatus := []int32{int32(123)} // []int32 | Filter, Default All，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed (optional)
	status := []int32{int32(123)} // []int32 | Status, 0-Active｜2-Frozen (optional)
	deleteInclude := true // bool | Deleted Involved，Need Admin (optional)
	sortField := "sortField_example" // string | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count OF Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserListGet(context.Background()).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).PlanIds(planIds).SubscriptionId(subscriptionId).SubStatus(subStatus).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserListGet`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | Filter UserId | 
 **firstName** | **string** | Search FirstName | 
 **lastName** | **string** | Search LastName | 
 **email** | **string** | Search Filter Email | 
 **planIds** | **[]int32** | PlanIds, Search Filter PlanIds | 
 **subscriptionId** | **string** | Search Filter SubscriptionId | 
 **subStatus** | **[]int32** | Filter, Default All，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | 
 **status** | **[]int32** | Status, 0-Active｜2-Frozen | 
 **deleteInclude** | **bool** | Deleted Involved，Need Admin | 
 **sortField** | **string** | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count OF Page | 
 **createTimeStart** | **int64** | CreateTimeStart | 
 **createTimeEnd** | **int64** | CreateTimeEnd | 

### Return type

[**MerchantUserListGet200Response**](MerchantUserListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserListPost

> MerchantUserListGet200Response UserListPost(ctx).UnibeeApiMerchantUserListReq(unibeeApiMerchantUserListReq).Execute()

User List

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
	unibeeApiMerchantUserListReq := *openapiclient.NewUnibeeApiMerchantUserListReq() // UnibeeApiMerchantUserListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserListPost(context.Background()).UnibeeApiMerchantUserListReq(unibeeApiMerchantUserListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserListPost`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserListReq** | [**UnibeeApiMerchantUserListReq**](UnibeeApiMerchantUserListReq.md) |  | 

### Return type

[**MerchantUserListGet200Response**](MerchantUserListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserNewAdminNotePost

> MerchantAuthSsoLoginOTPPost200Response UserNewAdminNotePost(ctx).UnibeeApiMerchantUserNewAdminNoteReq(unibeeApiMerchantUserNewAdminNoteReq).Execute()

New Admin Note

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
	unibeeApiMerchantUserNewAdminNoteReq := *openapiclient.NewUnibeeApiMerchantUserNewAdminNoteReq("Note_example", int64(123)) // UnibeeApiMerchantUserNewAdminNoteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserNewAdminNotePost(context.Background()).UnibeeApiMerchantUserNewAdminNoteReq(unibeeApiMerchantUserNewAdminNoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserNewAdminNotePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserNewAdminNotePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserNewAdminNotePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserNewAdminNotePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserNewAdminNoteReq** | [**UnibeeApiMerchantUserNewAdminNoteReq**](UnibeeApiMerchantUserNewAdminNoteReq.md) |  | 

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


## UserNewPost

> MerchantUserNewPost200Response UserNewPost(ctx).UnibeeApiMerchantUserNewReq(unibeeApiMerchantUserNewReq).Execute()

New User



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
	unibeeApiMerchantUserNewReq := *openapiclient.NewUnibeeApiMerchantUserNewReq("Email_example") // UnibeeApiMerchantUserNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserNewPost(context.Background()).UnibeeApiMerchantUserNewReq(unibeeApiMerchantUserNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserNewPost`: MerchantUserNewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserNewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserNewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserNewReq** | [**UnibeeApiMerchantUserNewReq**](UnibeeApiMerchantUserNewReq.md) |  | 

### Return type

[**MerchantUserNewPost200Response**](MerchantUserNewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserResumeUserPost

> MerchantAuthSsoLoginOTPPost200Response UserResumeUserPost(ctx).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()

Resume User

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
	unibeeApiMerchantUserReleaseReq := *openapiclient.NewUnibeeApiMerchantUserReleaseReq() // UnibeeApiMerchantUserReleaseReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserResumeUserPost(context.Background()).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserResumeUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserResumeUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserResumeUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserResumeUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserReleaseReq** | [**UnibeeApiMerchantUserReleaseReq**](UnibeeApiMerchantUserReleaseReq.md) |  | 

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


## UserSearchGet

> MerchantUserSearchGet200Response UserSearchGet(ctx).SearchKey(searchKey).Execute()

User Search

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
	searchKey := "searchKey_example" // string | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserSearchGet(context.Background()).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSearchGet`: MerchantUserSearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId | 

### Return type

[**MerchantUserSearchGet200Response**](MerchantUserSearchGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSearchPost

> MerchantUserSearchGet200Response UserSearchPost(ctx).UnibeeApiMerchantUserSearchReq(unibeeApiMerchantUserSearchReq).Execute()

User Search

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
	unibeeApiMerchantUserSearchReq := *openapiclient.NewUnibeeApiMerchantUserSearchReq() // UnibeeApiMerchantUserSearchReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserSearchPost(context.Background()).UnibeeApiMerchantUserSearchReq(unibeeApiMerchantUserSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSearchPost`: MerchantUserSearchGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserSearchReq** | [**UnibeeApiMerchantUserSearchReq**](UnibeeApiMerchantUserSearchReq.md) |  | 

### Return type

[**MerchantUserSearchGet200Response**](MerchantUserSearchGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSuspendUserPost

> MerchantAuthSsoLoginOTPPost200Response UserSuspendUserPost(ctx).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()

Suspend User

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
	unibeeApiMerchantUserFrozenReq := *openapiclient.NewUnibeeApiMerchantUserFrozenReq() // UnibeeApiMerchantUserFrozenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserSuspendUserPost(context.Background()).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserSuspendUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSuspendUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserSuspendUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSuspendUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserFrozenReq** | [**UnibeeApiMerchantUserFrozenReq**](UnibeeApiMerchantUserFrozenReq.md) |  | 

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


## UserUpdatePost

> MerchantUserGetGet200Response UserUpdatePost(ctx).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()

Update User Profile

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
	unibeeApiMerchantUserUpdateReq := *openapiclient.NewUnibeeApiMerchantUserUpdateReq() // UnibeeApiMerchantUserUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserUpdatePost(context.Background()).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUpdatePost`: MerchantUserGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserUpdateReq** | [**UnibeeApiMerchantUserUpdateReq**](UnibeeApiMerchantUserUpdateReq.md) |  | 

### Return type

[**MerchantUserGetGet200Response**](MerchantUserGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

