# \User

All URIs are relative to *http://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserFrozenUserPost**](User.md#UserFrozenUserPost) | **Post** /merchant/user/frozen_user | FrozenUser
[**UserGetGet**](User.md#UserGetGet) | **Get** /merchant/user/get | GetUserProfile
[**UserListGet**](User.md#UserListGet) | **Get** /merchant/user/list | UserList
[**UserListPost**](User.md#UserListPost) | **Post** /merchant/user/list | UserList
[**UserNewPost**](User.md#UserNewPost) | **Post** /merchant/user/new | NewUser
[**UserReleaseUserPost**](User.md#UserReleaseUserPost) | **Post** /merchant/user/release_user | ReleaseUser
[**UserSearchGet**](User.md#UserSearchGet) | **Get** /merchant/user/search | UserSearch
[**UserSearchPost**](User.md#UserSearchPost) | **Post** /merchant/user/search | UserSearch
[**UserUpdatePost**](User.md#UserUpdatePost) | **Post** /merchant/user/update | UpdateUserProfile



## UserFrozenUserPost

> MerchantAuthSsoLoginOTPPost200Response UserFrozenUserPost(ctx).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()

FrozenUser

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
	unibeeApiMerchantUserFrozenReq := *openapiclient.NewUnibeeApiMerchantUserFrozenReq() // UnibeeApiMerchantUserFrozenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserFrozenUserPost(context.Background()).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserFrozenUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserFrozenUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserFrozenUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserFrozenUserPostRequest struct via the builder pattern


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


## UserGetGet

> MerchantUserGetGet200Response UserGetGet(ctx).UserId(userId).Execute()

GetUserProfile

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

> MerchantUserListGet200Response UserListGet(ctx).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

UserList

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
	userId := int32(56) // int32 | Filter UserId (optional)
	firstName := "firstName_example" // string | Search FirstName (optional)
	lastName := "lastName_example" // string | Search LastName (optional)
	email := "email_example" // string | Search Filter Email (optional)
	status := []int32{int32(123)} // []int32 | Status, 0-Active｜2-Frozen (optional)
	deleteInclude := true // bool | Deleted Involved，Need Admin (optional)
	sortField := "sortField_example" // string | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page,Start 0 (optional)
	count := int32(56) // int32 | Count OF Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserListGet(context.Background()).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
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
 **userId** | **int32** | Filter UserId | 
 **firstName** | **string** | Search FirstName | 
 **lastName** | **string** | Search LastName | 
 **email** | **string** | Search Filter Email | 
 **status** | **[]int32** | Status, 0-Active｜2-Frozen | 
 **deleteInclude** | **bool** | Deleted Involved，Need Admin | 
 **sortField** | **string** | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page,Start 0 | 
 **count** | **int32** | Count OF Page | 

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

UserList

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


## UserNewPost

> MerchantUserGetGet200Response UserNewPost(ctx).UnibeeApiMerchantUserNewReq(unibeeApiMerchantUserNewReq).Execute()

NewUser



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
	unibeeApiMerchantUserNewReq := *openapiclient.NewUnibeeApiMerchantUserNewReq("Email_example", "ExternalUserId_example") // UnibeeApiMerchantUserNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserNewPost(context.Background()).UnibeeApiMerchantUserNewReq(unibeeApiMerchantUserNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserNewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserNewPost`: MerchantUserGetGet200Response
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

[**MerchantUserGetGet200Response**](MerchantUserGetGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserReleaseUserPost

> MerchantAuthSsoLoginOTPPost200Response UserReleaseUserPost(ctx).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()

ReleaseUser

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
	unibeeApiMerchantUserReleaseReq := *openapiclient.NewUnibeeApiMerchantUserReleaseReq() // UnibeeApiMerchantUserReleaseReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserReleaseUserPost(context.Background()).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserReleaseUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserReleaseUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `User.UserReleaseUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserReleaseUserPostRequest struct via the builder pattern


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

UserSearch

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

UserSearch

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


## UserUpdatePost

> MerchantAuthSsoLoginOTPPost200Response UserUpdatePost(ctx).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()

UpdateUserProfile

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
	unibeeApiMerchantUserUpdateReq := *openapiclient.NewUnibeeApiMerchantUserUpdateReq("Address_example", "CountryCode_example", "CountryName_example", "Email_example", "FirstName_example", "LastName_example", int64(123)) // UnibeeApiMerchantUserUpdateReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.User.UserUpdatePost(context.Background()).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `User.UserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserUpdatePost`: MerchantAuthSsoLoginOTPPost200Response
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

[**MerchantAuthSsoLoginOTPPost200Response**](MerchantAuthSsoLoginOTPPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

