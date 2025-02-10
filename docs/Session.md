# \Session

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionNewSessionPost**](Session.md#SessionNewSessionPost) | **Post** /merchant/session/new_session | New Session



## SessionNewSessionPost

> MerchantSessionNewSessionPost200Response SessionNewSessionPost(ctx).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()

New Session



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
	unibeeApiMerchantSessionNewReq := *openapiclient.NewUnibeeApiMerchantSessionNewReq("Email_example", "ExternalUserId_example") // UnibeeApiMerchantSessionNewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Session.SessionNewSessionPost(context.Background()).UnibeeApiMerchantSessionNewReq(unibeeApiMerchantSessionNewReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Session.SessionNewSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionNewSessionPost`: MerchantSessionNewSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Session.SessionNewSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionNewSessionPostRequest struct via the builder pattern


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

