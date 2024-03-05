# \MerchantUserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantUserFrozenUserPost**](MerchantUserAPI.md#MerchantUserFrozenUserPost) | **Post** /merchant/user/frozen_user | Merchant Frozen User
[**MerchantUserGetGet**](MerchantUserAPI.md#MerchantUserGetGet) | **Get** /merchant/user/get | Get User Profile
[**MerchantUserListGet**](MerchantUserAPI.md#MerchantUserListGet) | **Get** /merchant/user/list | User List
[**MerchantUserListPost**](MerchantUserAPI.md#MerchantUserListPost) | **Post** /merchant/user/list | User List
[**MerchantUserReleaseUserPost**](MerchantUserAPI.md#MerchantUserReleaseUserPost) | **Post** /merchant/user/release_user | Merchant Release User
[**MerchantUserSearchGet**](MerchantUserAPI.md#MerchantUserSearchGet) | **Get** /merchant/user/search | User Search
[**MerchantUserSearchPost**](MerchantUserAPI.md#MerchantUserSearchPost) | **Post** /merchant/user/search | User Search
[**MerchantUserUpdatePost**](MerchantUserAPI.md#MerchantUserUpdatePost) | **Post** /merchant/user/update | Update User Profile



## MerchantUserFrozenUserPost

> MerchantAuthSsoLoginOTPPost200Response MerchantUserFrozenUserPost(ctx).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()

Merchant Frozen User

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserFrozenUserPost(context.Background()).UnibeeApiMerchantUserFrozenReq(unibeeApiMerchantUserFrozenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserFrozenUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserFrozenUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserFrozenUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserFrozenUserPostRequest struct via the builder pattern


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


## MerchantUserGetGet

> MerchantUserGetGet200Response MerchantUserGetGet(ctx).UserId(userId).Execute()

Get User Profile

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserGetGet(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserGetGet`: MerchantUserGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserGetGetRequest struct via the builder pattern


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


## MerchantUserListGet

> MerchantUserListGet200Response MerchantUserListGet(ctx).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

User List

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserListGet(context.Background()).UserId(userId).FirstName(firstName).LastName(lastName).Email(email).Status(status).DeleteInclude(deleteInclude).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserListGet`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserListGetRequest struct via the builder pattern


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


## MerchantUserListPost

> MerchantUserListGet200Response MerchantUserListPost(ctx).UnibeeApiMerchantUserListReq(unibeeApiMerchantUserListReq).Execute()

User List

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserListPost(context.Background()).UnibeeApiMerchantUserListReq(unibeeApiMerchantUserListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserListPost`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserListPostRequest struct via the builder pattern


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


## MerchantUserReleaseUserPost

> MerchantAuthSsoLoginOTPPost200Response MerchantUserReleaseUserPost(ctx).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()

Merchant Release User

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserReleaseUserPost(context.Background()).UnibeeApiMerchantUserReleaseReq(unibeeApiMerchantUserReleaseReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserReleaseUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserReleaseUserPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserReleaseUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserReleaseUserPostRequest struct via the builder pattern


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


## MerchantUserSearchGet

> MerchantUserListGet200Response MerchantUserSearchGet(ctx).SearchKey(searchKey).Execute()

User Search

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserSearchGet(context.Background()).SearchKey(searchKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserSearchGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserSearchGet`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserSearchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserSearchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchKey** | **string** | SearchKey, Will Search UserId|Email|UserName|CompanyName|SubscriptionId|VatNumber|InvoiceId||PaymentId | 

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


## MerchantUserSearchPost

> MerchantUserListGet200Response MerchantUserSearchPost(ctx).UnibeeApiMerchantUserSearchReq(unibeeApiMerchantUserSearchReq).Execute()

User Search

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserSearchPost(context.Background()).UnibeeApiMerchantUserSearchReq(unibeeApiMerchantUserSearchReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserSearchPost`: MerchantUserListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserSearchReq** | [**UnibeeApiMerchantUserSearchReq**](UnibeeApiMerchantUserSearchReq.md) |  | 

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


## MerchantUserUpdatePost

> MerchantAuthSsoLoginOTPPost200Response MerchantUserUpdatePost(ctx).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()

Update User Profile

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
	resp, r, err := apiClient.MerchantUserAPI.MerchantUserUpdatePost(context.Background()).UnibeeApiMerchantUserUpdateReq(unibeeApiMerchantUserUpdateReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantUserAPI.MerchantUserUpdatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantUserUpdatePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantUserAPI.MerchantUserUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantUserUpdatePostRequest struct via the builder pattern


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

