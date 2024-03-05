# \MerchantSubscriptionNoteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSubscriptionAdminNoteListGet**](MerchantSubscriptionNoteAPI.md#MerchantSubscriptionAdminNoteListGet) | **Get** /merchant/subscription/admin_note_list | Merchant Subscription Note List
[**MerchantSubscriptionAdminNoteListPost**](MerchantSubscriptionNoteAPI.md#MerchantSubscriptionAdminNoteListPost) | **Post** /merchant/subscription/admin_note_list | Merchant Subscription Note List
[**MerchantSubscriptionNewAdminNotePost**](MerchantSubscriptionNoteAPI.md#MerchantSubscriptionNewAdminNotePost) | **Post** /merchant/subscription/new_admin_note | Merchant New Subscription Note



## MerchantSubscriptionAdminNoteListGet

> MerchantSubscriptionAdminNoteListGet200Response MerchantSubscriptionAdminNoteListGet(ctx).SubscriptionId(subscriptionId).Page(page).Count(count).Execute()

Merchant Subscription Note List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | SubscriptionId
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListGet(context.Background()).SubscriptionId(subscriptionId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionAdminNoteListGet`: MerchantSubscriptionAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionAdminNoteListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | SubscriptionId | 
 **page** | **int32** | Page, Start WIth 0 | 
 **count** | **int32** | Count Of Page | 

### Return type

[**MerchantSubscriptionAdminNoteListGet200Response**](MerchantSubscriptionAdminNoteListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionAdminNoteListPost

> MerchantSubscriptionAdminNoteListGet200Response MerchantSubscriptionAdminNoteListPost(ctx).UnibeeApiMerchantSubscriptionAdminNoteListReq(unibeeApiMerchantSubscriptionAdminNoteListReq).Execute()

Merchant Subscription Note List

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantSubscriptionAdminNoteListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionAdminNoteListReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionAdminNoteListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListPost(context.Background()).UnibeeApiMerchantSubscriptionAdminNoteListReq(unibeeApiMerchantSubscriptionAdminNoteListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionAdminNoteListPost`: MerchantSubscriptionAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionNoteAPI.MerchantSubscriptionAdminNoteListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionAdminNoteListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionAdminNoteListReq** | [**UnibeeApiMerchantSubscriptionAdminNoteListReq**](UnibeeApiMerchantSubscriptionAdminNoteListReq.md) |  | 

### Return type

[**MerchantSubscriptionAdminNoteListGet200Response**](MerchantSubscriptionAdminNoteListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MerchantSubscriptionNewAdminNotePost

> MerchantAuthSsoLoginOTPPost200Response MerchantSubscriptionNewAdminNotePost(ctx).UnibeeApiMerchantSubscriptionNewAdminNoteReq(unibeeApiMerchantSubscriptionNewAdminNoteReq).Execute()

Merchant New Subscription Note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	unibeeApiMerchantSubscriptionNewAdminNoteReq := *openapiclient.NewUnibeeApiMerchantSubscriptionNewAdminNoteReq(int64(123), "Note_example", "SubscriptionId_example") // UnibeeApiMerchantSubscriptionNewAdminNoteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSubscriptionNoteAPI.MerchantSubscriptionNewAdminNotePost(context.Background()).UnibeeApiMerchantSubscriptionNewAdminNoteReq(unibeeApiMerchantSubscriptionNewAdminNoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSubscriptionNoteAPI.MerchantSubscriptionNewAdminNotePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSubscriptionNewAdminNotePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSubscriptionNoteAPI.MerchantSubscriptionNewAdminNotePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSubscriptionNewAdminNotePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSubscriptionNewAdminNoteReq** | [**UnibeeApiMerchantSubscriptionNewAdminNoteReq**](UnibeeApiMerchantSubscriptionNewAdminNoteReq.md) |  | 

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

