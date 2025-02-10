# \PromoCredit

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditEditPromoConfigPost**](PromoCredit.md#CreditEditPromoConfigPost) | **Post** /merchant/credit/edit_promo_config | Edit Promo Credit Config
[**CreditGetPromoConfigGet**](PromoCredit.md#CreditGetPromoConfigGet) | **Get** /merchant/credit/get_promo_config | Get Promo Credit Config
[**CreditGetPromoConfigPost**](PromoCredit.md#CreditGetPromoConfigPost) | **Post** /merchant/credit/get_promo_config | Get Promo Credit Config
[**CreditGetPromoConfigStatisticsGet**](PromoCredit.md#CreditGetPromoConfigStatisticsGet) | **Get** /merchant/credit/get_promo_config_statistics | Get Promo Credit Config Statistics
[**CreditGetPromoConfigStatisticsPost**](PromoCredit.md#CreditGetPromoConfigStatisticsPost) | **Post** /merchant/credit/get_promo_config_statistics | Get Promo Credit Config Statistics
[**CreditPromoCreditDecrementPost**](PromoCredit.md#CreditPromoCreditDecrementPost) | **Post** /merchant/credit/promo_credit_decrement | Promo Credit Decrement
[**CreditPromoCreditIncrementPost**](PromoCredit.md#CreditPromoCreditIncrementPost) | **Post** /merchant/credit/promo_credit_increment | Promo Credit Increment



## CreditEditPromoConfigPost

> MerchantCreditEditConfigPost200Response CreditEditPromoConfigPost(ctx).UnibeeApiMerchantCreditEditPromoConfigReq(unibeeApiMerchantCreditEditPromoConfigReq).Execute()

Edit Promo Credit Config

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
	unibeeApiMerchantCreditEditPromoConfigReq := *openapiclient.NewUnibeeApiMerchantCreditEditPromoConfigReq("Currency_example") // UnibeeApiMerchantCreditEditPromoConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditEditPromoConfigPost(context.Background()).UnibeeApiMerchantCreditEditPromoConfigReq(unibeeApiMerchantCreditEditPromoConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditEditPromoConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditEditPromoConfigPost`: MerchantCreditEditConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditEditPromoConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditEditPromoConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditEditPromoConfigReq** | [**UnibeeApiMerchantCreditEditPromoConfigReq**](UnibeeApiMerchantCreditEditPromoConfigReq.md) |  | 

### Return type

[**MerchantCreditEditConfigPost200Response**](MerchantCreditEditConfigPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditGetPromoConfigGet

> MerchantCreditEditConfigPost200Response CreditGetPromoConfigGet(ctx).Currency(currency).Execute()

Get Promo Credit Config

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
	currency := "currency_example" // string | currency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditGetPromoConfigGet(context.Background()).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditGetPromoConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditGetPromoConfigGet`: MerchantCreditEditConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditGetPromoConfigGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditGetPromoConfigGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | currency | 

### Return type

[**MerchantCreditEditConfigPost200Response**](MerchantCreditEditConfigPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditGetPromoConfigPost

> MerchantCreditEditConfigPost200Response CreditGetPromoConfigPost(ctx).UnibeeApiMerchantCreditPromoConfigReq(unibeeApiMerchantCreditPromoConfigReq).Execute()

Get Promo Credit Config

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
	unibeeApiMerchantCreditPromoConfigReq := *openapiclient.NewUnibeeApiMerchantCreditPromoConfigReq() // UnibeeApiMerchantCreditPromoConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditGetPromoConfigPost(context.Background()).UnibeeApiMerchantCreditPromoConfigReq(unibeeApiMerchantCreditPromoConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditGetPromoConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditGetPromoConfigPost`: MerchantCreditEditConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditGetPromoConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditGetPromoConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditPromoConfigReq** | [**UnibeeApiMerchantCreditPromoConfigReq**](UnibeeApiMerchantCreditPromoConfigReq.md) |  | 

### Return type

[**MerchantCreditEditConfigPost200Response**](MerchantCreditEditConfigPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditGetPromoConfigStatisticsGet

> MerchantCreditGetPromoConfigStatisticsGet200Response CreditGetPromoConfigStatisticsGet(ctx).Currency(currency).Execute()

Get Promo Credit Config Statistics

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
	currency := "currency_example" // string | currency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditGetPromoConfigStatisticsGet(context.Background()).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditGetPromoConfigStatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditGetPromoConfigStatisticsGet`: MerchantCreditGetPromoConfigStatisticsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditGetPromoConfigStatisticsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditGetPromoConfigStatisticsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | currency | 

### Return type

[**MerchantCreditGetPromoConfigStatisticsGet200Response**](MerchantCreditGetPromoConfigStatisticsGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditGetPromoConfigStatisticsPost

> MerchantCreditGetPromoConfigStatisticsGet200Response CreditGetPromoConfigStatisticsPost(ctx).UnibeeApiMerchantCreditPromoConfigStatisticsReq(unibeeApiMerchantCreditPromoConfigStatisticsReq).Execute()

Get Promo Credit Config Statistics

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
	unibeeApiMerchantCreditPromoConfigStatisticsReq := *openapiclient.NewUnibeeApiMerchantCreditPromoConfigStatisticsReq() // UnibeeApiMerchantCreditPromoConfigStatisticsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditGetPromoConfigStatisticsPost(context.Background()).UnibeeApiMerchantCreditPromoConfigStatisticsReq(unibeeApiMerchantCreditPromoConfigStatisticsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditGetPromoConfigStatisticsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditGetPromoConfigStatisticsPost`: MerchantCreditGetPromoConfigStatisticsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditGetPromoConfigStatisticsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditGetPromoConfigStatisticsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditPromoConfigStatisticsReq** | [**UnibeeApiMerchantCreditPromoConfigStatisticsReq**](UnibeeApiMerchantCreditPromoConfigStatisticsReq.md) |  | 

### Return type

[**MerchantCreditGetPromoConfigStatisticsGet200Response**](MerchantCreditGetPromoConfigStatisticsGet200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditPromoCreditDecrementPost

> MerchantCreditPromoCreditDecrementPost200Response CreditPromoCreditDecrementPost(ctx).UnibeeApiMerchantCreditPromoCreditDecrementReq(unibeeApiMerchantCreditPromoCreditDecrementReq).Execute()

Promo Credit Decrement



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
	unibeeApiMerchantCreditPromoCreditDecrementReq := *openapiclient.NewUnibeeApiMerchantCreditPromoCreditDecrementReq(int64(123), "Currency_example", int64(123)) // UnibeeApiMerchantCreditPromoCreditDecrementReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditPromoCreditDecrementPost(context.Background()).UnibeeApiMerchantCreditPromoCreditDecrementReq(unibeeApiMerchantCreditPromoCreditDecrementReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditPromoCreditDecrementPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditPromoCreditDecrementPost`: MerchantCreditPromoCreditDecrementPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditPromoCreditDecrementPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditPromoCreditDecrementPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditPromoCreditDecrementReq** | [**UnibeeApiMerchantCreditPromoCreditDecrementReq**](UnibeeApiMerchantCreditPromoCreditDecrementReq.md) |  | 

### Return type

[**MerchantCreditPromoCreditDecrementPost200Response**](MerchantCreditPromoCreditDecrementPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditPromoCreditIncrementPost

> MerchantCreditPromoCreditDecrementPost200Response CreditPromoCreditIncrementPost(ctx).UnibeeApiMerchantCreditPromoCreditIncrementReq(unibeeApiMerchantCreditPromoCreditIncrementReq).Execute()

Promo Credit Increment



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
	unibeeApiMerchantCreditPromoCreditIncrementReq := *openapiclient.NewUnibeeApiMerchantCreditPromoCreditIncrementReq(int64(123), "Currency_example", int64(123)) // UnibeeApiMerchantCreditPromoCreditIncrementReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PromoCredit.CreditPromoCreditIncrementPost(context.Background()).UnibeeApiMerchantCreditPromoCreditIncrementReq(unibeeApiMerchantCreditPromoCreditIncrementReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PromoCredit.CreditPromoCreditIncrementPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditPromoCreditIncrementPost`: MerchantCreditPromoCreditDecrementPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PromoCredit.CreditPromoCreditIncrementPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditPromoCreditIncrementPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditPromoCreditIncrementReq** | [**UnibeeApiMerchantCreditPromoCreditIncrementReq**](UnibeeApiMerchantCreditPromoCreditIncrementReq.md) |  | 

### Return type

[**MerchantCreditPromoCreditDecrementPost200Response**](MerchantCreditPromoCreditDecrementPost200Response.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

