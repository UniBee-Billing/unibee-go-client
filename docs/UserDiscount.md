# \UserDiscount

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscountPlanApplyPreviewPost**](UserDiscount.md#DiscountPlanApplyPreviewPost) | **Post** /merchant/discount/plan_apply_preview | Plan Apply Preview
[**DiscountUserDiscountListGet**](UserDiscount.md#DiscountUserDiscountListGet) | **Get** /merchant/discount/user_discount_list | Get User Discount Code List



## DiscountPlanApplyPreviewPost

> MerchantDiscountPlanApplyPreviewPost200Response DiscountPlanApplyPreviewPost(ctx).UnibeeApiMerchantDiscountPlanApplyPreviewReq(unibeeApiMerchantDiscountPlanApplyPreviewReq).Execute()

Plan Apply Preview



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
	unibeeApiMerchantDiscountPlanApplyPreviewReq := *openapiclient.NewUnibeeApiMerchantDiscountPlanApplyPreviewReq("Code_example") // UnibeeApiMerchantDiscountPlanApplyPreviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDiscount.DiscountPlanApplyPreviewPost(context.Background()).UnibeeApiMerchantDiscountPlanApplyPreviewReq(unibeeApiMerchantDiscountPlanApplyPreviewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDiscount.DiscountPlanApplyPreviewPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountPlanApplyPreviewPost`: MerchantDiscountPlanApplyPreviewPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserDiscount.DiscountPlanApplyPreviewPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountPlanApplyPreviewPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantDiscountPlanApplyPreviewReq** | [**UnibeeApiMerchantDiscountPlanApplyPreviewReq**](UnibeeApiMerchantDiscountPlanApplyPreviewReq.md) |  | 

### Return type

[**MerchantDiscountPlanApplyPreviewPost200Response**](MerchantDiscountPlanApplyPreviewPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscountUserDiscountListGet

> MerchantDiscountUserDiscountListGet200Response DiscountUserDiscountListGet(ctx).Id(id).UserIds(userIds).Email(email).PlanIds(planIds).SubscriptionIds(subscriptionIds).Status(status).ChildCode(childCode).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get User Discount Code List



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
	id := int64(789) // int64 | The discount's Id
	userIds := []int64{int64(123)} // []int64 | Filter UserIds Default All (optional)
	email := "email_example" // string | Filter Email Default All (optional)
	planIds := []int64{int64(123)} // []int64 | Filter PlanIds Default All (optional)
	subscriptionIds := []string{"Inner_example"} // []string | Filter SubscriptionIds Default All (optional)
	status := []int32{int32(123)} // []int32 | Filter Status Default All, 1-normal, 2-rollback (optional)
	childCode := "childCode_example" // string | Filter Child Code (fuzzy search), only available when Id is batch template ID (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDiscount.DiscountUserDiscountListGet(context.Background()).Id(id).UserIds(userIds).Email(email).PlanIds(planIds).SubscriptionIds(subscriptionIds).Status(status).ChildCode(childCode).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDiscount.DiscountUserDiscountListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscountUserDiscountListGet`: MerchantDiscountUserDiscountListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserDiscount.DiscountUserDiscountListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscountUserDiscountListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The discount&#39;s Id | 
 **userIds** | **[]int64** | Filter UserIds Default All | 
 **email** | **string** | Filter Email Default All | 
 **planIds** | **[]int64** | Filter PlanIds Default All | 
 **subscriptionIds** | **[]string** | Filter SubscriptionIds Default All | 
 **status** | **[]int32** | Filter Status Default All, 1-normal, 2-rollback | 
 **childCode** | **string** | Filter Child Code (fuzzy search), only available when Id is batch template ID | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantDiscountUserDiscountListGet200Response**](MerchantDiscountUserDiscountListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

