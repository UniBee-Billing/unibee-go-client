# \SubscriptionNote

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionAdminNoteListGet**](SubscriptionNote.md#SubscriptionAdminNoteListGet) | **Get** /merchant/subscription/admin_note_list | Merchant Subscription Note List
[**SubscriptionAdminNoteListPost**](SubscriptionNote.md#SubscriptionAdminNoteListPost) | **Post** /merchant/subscription/admin_note_list | Merchant Subscription Note List
[**SubscriptionNewAdminNotePost**](SubscriptionNote.md#SubscriptionNewAdminNotePost) | **Post** /merchant/subscription/new_admin_note | Merchant New Subscription Note



## SubscriptionAdminNoteListGet

> MerchantSubscriptionAdminNoteListGet200Response SubscriptionAdminNoteListGet(ctx).SubscriptionId(subscriptionId).Page(page).Count(count).Execute()

Merchant Subscription Note List

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
	subscriptionId := "subscriptionId_example" // string | SubscriptionId
	page := int32(56) // int32 | Page, Start WIth 0 (optional)
	count := int32(56) // int32 | Count Of Page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionNote.SubscriptionAdminNoteListGet(context.Background()).SubscriptionId(subscriptionId).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionNote.SubscriptionAdminNoteListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAdminNoteListGet`: MerchantSubscriptionAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionNote.SubscriptionAdminNoteListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAdminNoteListGetRequest struct via the builder pattern


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


## SubscriptionAdminNoteListPost

> MerchantSubscriptionAdminNoteListGet200Response SubscriptionAdminNoteListPost(ctx).UnibeeApiMerchantSubscriptionAdminNoteListReq(unibeeApiMerchantSubscriptionAdminNoteListReq).Execute()

Merchant Subscription Note List

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
	unibeeApiMerchantSubscriptionAdminNoteListReq := *openapiclient.NewUnibeeApiMerchantSubscriptionAdminNoteListReq("SubscriptionId_example") // UnibeeApiMerchantSubscriptionAdminNoteListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionNote.SubscriptionAdminNoteListPost(context.Background()).UnibeeApiMerchantSubscriptionAdminNoteListReq(unibeeApiMerchantSubscriptionAdminNoteListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionNote.SubscriptionAdminNoteListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionAdminNoteListPost`: MerchantSubscriptionAdminNoteListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionNote.SubscriptionAdminNoteListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionAdminNoteListPostRequest struct via the builder pattern


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


## SubscriptionNewAdminNotePost

> MerchantAuthSsoLoginOTPPost200Response SubscriptionNewAdminNotePost(ctx).UnibeeApiMerchantSubscriptionNewAdminNoteReq(unibeeApiMerchantSubscriptionNewAdminNoteReq).Execute()

Merchant New Subscription Note

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
	unibeeApiMerchantSubscriptionNewAdminNoteReq := *openapiclient.NewUnibeeApiMerchantSubscriptionNewAdminNoteReq("Note_example", "SubscriptionId_example") // UnibeeApiMerchantSubscriptionNewAdminNoteReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionNote.SubscriptionNewAdminNotePost(context.Background()).UnibeeApiMerchantSubscriptionNewAdminNoteReq(unibeeApiMerchantSubscriptionNewAdminNoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionNote.SubscriptionNewAdminNotePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionNewAdminNotePost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionNote.SubscriptionNewAdminNotePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionNewAdminNotePostRequest struct via the builder pattern


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

