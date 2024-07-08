# \Track

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackSetupSegmentPost**](Track.md#TrackSetupSegmentPost) | **Post** /merchant/track/setup_segment | SegmentSetup



## TrackSetupSegmentPost

> MerchantAuthSsoLoginOTPPost200Response TrackSetupSegmentPost(ctx).UnibeeApiMerchantTrackSetupSegmentReq(unibeeApiMerchantTrackSetupSegmentReq).Execute()

SegmentSetup

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
	unibeeApiMerchantTrackSetupSegmentReq := *openapiclient.NewUnibeeApiMerchantTrackSetupSegmentReq("ServerSideSecret_example", "UserPortalSecret_example") // UnibeeApiMerchantTrackSetupSegmentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Track.TrackSetupSegmentPost(context.Background()).UnibeeApiMerchantTrackSetupSegmentReq(unibeeApiMerchantTrackSetupSegmentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Track.TrackSetupSegmentPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackSetupSegmentPost`: MerchantAuthSsoLoginOTPPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Track.TrackSetupSegmentPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackSetupSegmentPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantTrackSetupSegmentReq** | [**UnibeeApiMerchantTrackSetupSegmentReq**](UnibeeApiMerchantTrackSetupSegmentReq.md) |  | 

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

