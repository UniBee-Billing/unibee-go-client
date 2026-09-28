# \UserChurn

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserChurnBulkCheckPost**](UserChurn.md#UserChurnBulkCheckPost) | **Post** /merchant/user_churn/bulk_check | Bulk Check Currently Churned Users
[**UserChurnListGet**](UserChurn.md#UserChurnListGet) | **Get** /merchant/user_churn/list | Churned User List
[**UserChurnListPost**](UserChurn.md#UserChurnListPost) | **Post** /merchant/user_churn/list | Churned User List



## UserChurnBulkCheckPost

> MerchantUserChurnBulkCheckPost200Response UserChurnBulkCheckPost(ctx).UnibeeApiMerchantUserChurnBulkCheckReq(unibeeApiMerchantUserChurnBulkCheckReq).Execute()

Bulk Check Currently Churned Users



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
	unibeeApiMerchantUserChurnBulkCheckReq := *openapiclient.NewUnibeeApiMerchantUserChurnBulkCheckReq("Emails_example") // UnibeeApiMerchantUserChurnBulkCheckReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserChurn.UserChurnBulkCheckPost(context.Background()).UnibeeApiMerchantUserChurnBulkCheckReq(unibeeApiMerchantUserChurnBulkCheckReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserChurn.UserChurnBulkCheckPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChurnBulkCheckPost`: MerchantUserChurnBulkCheckPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserChurn.UserChurnBulkCheckPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChurnBulkCheckPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserChurnBulkCheckReq** | [**UnibeeApiMerchantUserChurnBulkCheckReq**](UnibeeApiMerchantUserChurnBulkCheckReq.md) |  | 

### Return type

[**MerchantUserChurnBulkCheckPost200Response**](MerchantUserChurnBulkCheckPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChurnListGet

> MerchantUserChurnListGet200Response UserChurnListGet(ctx).View(view).ChurnTimeStart(churnTimeStart).ChurnTimeEnd(churnTimeEnd).ChurnType(churnType).ReasonCategory(reasonCategory).ProductIds(productIds).PlanIds(planIds).UserId(userId).Email(email).ExternalUserId(externalUserId).SearchKey(searchKey).CountryCode(countryCode).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()

Churned User List

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
	view := "view_example" // string | pending|current|reactivated|all, default all (optional)
	churnTimeStart := int64(789) // int64 | ChurnTimeStart, UTC seconds, inclusive (optional)
	churnTimeEnd := int64(789) // int64 | ChurnTimeEnd, UTC seconds, exclusive [start,end) (optional)
	churnType := "churnType_example" // string | Cancelled|Expired (optional)
	reasonCategory := "reasonCategory_example" // string | Normalized reason category (optional)
	productIds := []int64{int64(123)} // []int64 | Filter ProductIds (optional)
	planIds := []int64{int64(123)} // []int64 | Filter PlanIds (optional)
	userId := int64(789) // int64 | Filter UserId (optional)
	email := "email_example" // string | Filter Email (optional)
	externalUserId := "externalUserId_example" // string | Filter ExternalUserId (optional)
	searchKey := "searchKey_example" // string | Search Email|ExternalUserId|SubscriptionId|UserId (optional)
	countryCode := "countryCode_example" // string | Filter CountryCode (optional)
	currency := "currency_example" // string | Filter Currency (optional)
	sortField := "sortField_example" // string | churn_time|confirm_time|reactivated_time|id, default churn_time (optional)
	sortType := "sortType_example" // string | asc|desc, default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserChurn.UserChurnListGet(context.Background()).View(view).ChurnTimeStart(churnTimeStart).ChurnTimeEnd(churnTimeEnd).ChurnType(churnType).ReasonCategory(reasonCategory).ProductIds(productIds).PlanIds(planIds).UserId(userId).Email(email).ExternalUserId(externalUserId).SearchKey(searchKey).CountryCode(countryCode).Currency(currency).SortField(sortField).SortType(sortType).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserChurn.UserChurnListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChurnListGet`: MerchantUserChurnListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserChurn.UserChurnListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChurnListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | pending|current|reactivated|all, default all | 
 **churnTimeStart** | **int64** | ChurnTimeStart, UTC seconds, inclusive | 
 **churnTimeEnd** | **int64** | ChurnTimeEnd, UTC seconds, exclusive [start,end) | 
 **churnType** | **string** | Cancelled|Expired | 
 **reasonCategory** | **string** | Normalized reason category | 
 **productIds** | **[]int64** | Filter ProductIds | 
 **planIds** | **[]int64** | Filter PlanIds | 
 **userId** | **int64** | Filter UserId | 
 **email** | **string** | Filter Email | 
 **externalUserId** | **string** | Filter ExternalUserId | 
 **searchKey** | **string** | Search Email|ExternalUserId|SubscriptionId|UserId | 
 **countryCode** | **string** | Filter CountryCode | 
 **currency** | **string** | Filter Currency | 
 **sortField** | **string** | churn_time|confirm_time|reactivated_time|id, default churn_time | 
 **sortType** | **string** | asc|desc, default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 

### Return type

[**MerchantUserChurnListGet200Response**](MerchantUserChurnListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserChurnListPost

> MerchantUserChurnListGet200Response UserChurnListPost(ctx).UnibeeApiMerchantUserChurnListReq(unibeeApiMerchantUserChurnListReq).Execute()

Churned User List

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
	unibeeApiMerchantUserChurnListReq := *openapiclient.NewUnibeeApiMerchantUserChurnListReq() // UnibeeApiMerchantUserChurnListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserChurn.UserChurnListPost(context.Background()).UnibeeApiMerchantUserChurnListReq(unibeeApiMerchantUserChurnListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserChurn.UserChurnListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChurnListPost`: MerchantUserChurnListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserChurn.UserChurnListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserChurnListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantUserChurnListReq** | [**UnibeeApiMerchantUserChurnListReq**](UnibeeApiMerchantUserChurnListReq.md) |  | 

### Return type

[**MerchantUserChurnListGet200Response**](MerchantUserChurnListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

