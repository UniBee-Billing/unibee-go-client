# \MerchantSessionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MerchantSessionNewSessionPost**](MerchantSessionAPI.md#MerchantSessionNewSessionPost) | **Post** /merchant/session/new_session | New User Portal Session



## MerchantSessionNewSessionPost

> MerchantSessionNewSessionPost200Response MerchantSessionNewSessionPost(ctx).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()

New User Portal Session

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
	unibeeApiMerchantSessionNewReq := *openapiclient.NewUnibeeApiMerchantSessionNewReq("Email_example", "ExternalUserId_example") // UnibeeApiMerchantSessionNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantSessionAPI.MerchantSessionNewSessionPost(context.Background()).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantSessionAPI.MerchantSessionNewSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MerchantSessionNewSessionPost`: MerchantSessionNewSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `MerchantSessionAPI.MerchantSessionNewSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMerchantSessionNewSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantSessionNewReq** | [**UnibeeApiMerchantSessionNewReq**](UnibeeApiMerchantSessionNewReq.md) |  | 

### Return type

[**MerchantSessionNewSessionPost200Response**](MerchantSessionNewSessionPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

