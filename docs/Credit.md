# \Credit

All URIs are relative to *https://api.unibee.top*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreditConfigListGet**](Credit.md#CreditConfigListGet) | **Get** /merchant/credit/config_list | Get Credit Config list
[**CreditConfigListPost**](Credit.md#CreditConfigListPost) | **Post** /merchant/credit/config_list | Get Credit Config list
[**CreditCreditAccountListGet**](Credit.md#CreditCreditAccountListGet) | **Get** /merchant/credit/credit_account_list | Get Credit Account List
[**CreditCreditTransactionListGet**](Credit.md#CreditCreditTransactionListGet) | **Get** /merchant/credit/credit_transaction_list | Get Credit Transaction List
[**CreditCreditTransactionListPost**](Credit.md#CreditCreditTransactionListPost) | **Post** /merchant/credit/credit_transaction_list | Get Credit Transaction List
[**CreditDetailGet**](Credit.md#CreditDetailGet) | **Get** /merchant/credit/detail | Credit Account Detail
[**CreditDetailPost**](Credit.md#CreditDetailPost) | **Post** /merchant/credit/detail | Credit Account Detail
[**CreditEditConfigPost**](Credit.md#CreditEditConfigPost) | **Post** /merchant/credit/edit_config | Edit Credit Config
[**CreditEditCreditAccountPost**](Credit.md#CreditEditCreditAccountPost) | **Post** /merchant/credit/edit_credit_account | Edit User Credit Account Config
[**CreditNewConfigPost**](Credit.md#CreditNewConfigPost) | **Post** /merchant/credit/new_config | Setup Credit Config
[**CreditNewCreditRechargePost**](Credit.md#CreditNewCreditRechargePost) | **Post** /merchant/credit/new_credit_recharge | New Credit Recharge



## CreditConfigListGet

> MerchantCreditConfigListGet200Response CreditConfigListGet(ctx).Types(types).Currency(currency).Execute()

Get Credit Config list

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
	types := []int32{int32(123)} // []int32 | type list of credit account, 1-main account, 2-promo credit account (optional)
	currency := "currency_example" // string | currency (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditConfigListGet(context.Background()).Types(types).Currency(currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditConfigListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditConfigListGet`: MerchantCreditConfigListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditConfigListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditConfigListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]int32** | type list of credit account, 1-main account, 2-promo credit account | 
 **currency** | **string** | currency | 

### Return type

[**MerchantCreditConfigListGet200Response**](MerchantCreditConfigListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditConfigListPost

> MerchantCreditConfigListGet200Response CreditConfigListPost(ctx).UnibeeApiMerchantCreditConfigListReq(unibeeApiMerchantCreditConfigListReq).Execute()

Get Credit Config list

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
	unibeeApiMerchantCreditConfigListReq := *openapiclient.NewUnibeeApiMerchantCreditConfigListReq() // UnibeeApiMerchantCreditConfigListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditConfigListPost(context.Background()).UnibeeApiMerchantCreditConfigListReq(unibeeApiMerchantCreditConfigListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditConfigListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditConfigListPost`: MerchantCreditConfigListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditConfigListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditConfigListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditConfigListReq** | [**UnibeeApiMerchantCreditConfigListReq**](UnibeeApiMerchantCreditConfigListReq.md) |  | 

### Return type

[**MerchantCreditConfigListGet200Response**](MerchantCreditConfigListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditCreditAccountListGet

> MerchantCreditCreditAccountListGet200Response CreditCreditAccountListGet(ctx).UserId(userId).Email(email).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Credit Account List



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
	userId := int64(789) // int64 | filter id of user (optional)
	email := "email_example" // string | filter email of user (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditCreditAccountListGet(context.Background()).UserId(userId).Email(email).SortField(sortField).SortType(sortType).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditCreditAccountListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditCreditAccountListGet`: MerchantCreditCreditAccountListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditCreditAccountListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditCreditAccountListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int64** | filter id of user | 
 **email** | **string** | filter email of user | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantCreditCreditAccountListGet200Response**](MerchantCreditCreditAccountListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditCreditTransactionListGet

> MerchantCreditCreditTransactionListGet200Response CreditCreditTransactionListGet(ctx).AccountType(accountType).UserId(userId).Email(email).Currency(currency).SortField(sortField).SortType(sortType).TransactionTypes(transactionTypes).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()

Get Credit Transaction List



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
	accountType := int32(56) // int32 | filter type of account, 1-main account, 2-promo credit account
	userId := int64(789) // int64 | filter id of user (optional)
	email := "email_example" // string | filter email of user (optional)
	currency := "currency_example" // string | filter currency of account (optional)
	sortField := "sortField_example" // string | Sort Field，gmt_create|gmt_modify，Default gmt_modify (optional)
	sortType := "sortType_example" // string | Sort Type，asc|desc，Default desc (optional)
	transactionTypes := []int32{int32(123)} // []int32 | transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out (optional)
	page := int32(56) // int32 | Page, Start 0 (optional)
	count := int32(56) // int32 | Count Of Per Page (optional)
	createTimeStart := int64(789) // int64 | CreateTimeStart，UTC timestamp，seconds (optional)
	createTimeEnd := int64(789) // int64 | CreateTimeEnd，UTC timestamp，seconds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditCreditTransactionListGet(context.Background()).AccountType(accountType).UserId(userId).Email(email).Currency(currency).SortField(sortField).SortType(sortType).TransactionTypes(transactionTypes).Page(page).Count(count).CreateTimeStart(createTimeStart).CreateTimeEnd(createTimeEnd).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditCreditTransactionListGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditCreditTransactionListGet`: MerchantCreditCreditTransactionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditCreditTransactionListGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditCreditTransactionListGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountType** | **int32** | filter type of account, 1-main account, 2-promo credit account | 
 **userId** | **int64** | filter id of user | 
 **email** | **string** | filter email of user | 
 **currency** | **string** | filter currency of account | 
 **sortField** | **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | 
 **sortType** | **string** | Sort Type，asc|desc，Default desc | 
 **transactionTypes** | **[]int32** | transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out | 
 **page** | **int32** | Page, Start 0 | 
 **count** | **int32** | Count Of Per Page | 
 **createTimeStart** | **int64** | CreateTimeStart，UTC timestamp，seconds | 
 **createTimeEnd** | **int64** | CreateTimeEnd，UTC timestamp，seconds | 

### Return type

[**MerchantCreditCreditTransactionListGet200Response**](MerchantCreditCreditTransactionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditCreditTransactionListPost

> MerchantCreditCreditTransactionListGet200Response CreditCreditTransactionListPost(ctx).UnibeeApiMerchantCreditCreditTransactionListReq(unibeeApiMerchantCreditCreditTransactionListReq).Execute()

Get Credit Transaction List



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
	unibeeApiMerchantCreditCreditTransactionListReq := *openapiclient.NewUnibeeApiMerchantCreditCreditTransactionListReq(int32(123)) // UnibeeApiMerchantCreditCreditTransactionListReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditCreditTransactionListPost(context.Background()).UnibeeApiMerchantCreditCreditTransactionListReq(unibeeApiMerchantCreditCreditTransactionListReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditCreditTransactionListPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditCreditTransactionListPost`: MerchantCreditCreditTransactionListGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditCreditTransactionListPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditCreditTransactionListPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditCreditTransactionListReq** | [**UnibeeApiMerchantCreditCreditTransactionListReq**](UnibeeApiMerchantCreditCreditTransactionListReq.md) |  | 

### Return type

[**MerchantCreditCreditTransactionListGet200Response**](MerchantCreditCreditTransactionListGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditDetailGet

> MerchantCreditDetailGet200Response CreditDetailGet(ctx).Id(id).Execute()

Credit Account Detail

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
	id := int64(789) // int64 | The credit account Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditDetailGet(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditDetailGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditDetailGet`: MerchantCreditDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditDetailGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditDetailGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64** | The credit account Id | 

### Return type

[**MerchantCreditDetailGet200Response**](MerchantCreditDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditDetailPost

> MerchantCreditDetailGet200Response CreditDetailPost(ctx).UnibeeApiMerchantCreditDetailReq(unibeeApiMerchantCreditDetailReq).Execute()

Credit Account Detail

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
	unibeeApiMerchantCreditDetailReq := *openapiclient.NewUnibeeApiMerchantCreditDetailReq(int64(123)) // UnibeeApiMerchantCreditDetailReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditDetailPost(context.Background()).UnibeeApiMerchantCreditDetailReq(unibeeApiMerchantCreditDetailReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditDetailPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditDetailPost`: MerchantCreditDetailGet200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditDetailPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditDetailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditDetailReq** | [**UnibeeApiMerchantCreditDetailReq**](UnibeeApiMerchantCreditDetailReq.md) |  | 

### Return type

[**MerchantCreditDetailGet200Response**](MerchantCreditDetailGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditEditConfigPost

> MerchantCreditEditConfigPost200Response CreditEditConfigPost(ctx).UnibeeApiMerchantCreditEditConfigReq(unibeeApiMerchantCreditEditConfigReq).Execute()

Edit Credit Config

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
	unibeeApiMerchantCreditEditConfigReq := *openapiclient.NewUnibeeApiMerchantCreditEditConfigReq("Currency_example", int32(123)) // UnibeeApiMerchantCreditEditConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditEditConfigPost(context.Background()).UnibeeApiMerchantCreditEditConfigReq(unibeeApiMerchantCreditEditConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditEditConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditEditConfigPost`: MerchantCreditEditConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditEditConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditEditConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditEditConfigReq** | [**UnibeeApiMerchantCreditEditConfigReq**](UnibeeApiMerchantCreditEditConfigReq.md) |  | 

### Return type

[**MerchantCreditEditConfigPost200Response**](MerchantCreditEditConfigPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditEditCreditAccountPost

> MerchantCreditEditCreditAccountPost200Response CreditEditCreditAccountPost(ctx).UnibeeApiMerchantCreditEditCreditAccountReq(unibeeApiMerchantCreditEditCreditAccountReq).Execute()

Edit User Credit Account Config



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
	unibeeApiMerchantCreditEditCreditAccountReq := *openapiclient.NewUnibeeApiMerchantCreditEditCreditAccountReq(int64(123)) // UnibeeApiMerchantCreditEditCreditAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditEditCreditAccountPost(context.Background()).UnibeeApiMerchantCreditEditCreditAccountReq(unibeeApiMerchantCreditEditCreditAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditEditCreditAccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditEditCreditAccountPost`: MerchantCreditEditCreditAccountPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditEditCreditAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditEditCreditAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditEditCreditAccountReq** | [**UnibeeApiMerchantCreditEditCreditAccountReq**](UnibeeApiMerchantCreditEditCreditAccountReq.md) |  | 

### Return type

[**MerchantCreditEditCreditAccountPost200Response**](MerchantCreditEditCreditAccountPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNewConfigPost

> MerchantCreditEditConfigPost200Response CreditNewConfigPost(ctx).UnibeeApiMerchantCreditNewConfigReq(unibeeApiMerchantCreditNewConfigReq).Execute()

Setup Credit Config

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
	unibeeApiMerchantCreditNewConfigReq := *openapiclient.NewUnibeeApiMerchantCreditNewConfigReq("Currency_example", int32(123)) // UnibeeApiMerchantCreditNewConfigReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditNewConfigPost(context.Background()).UnibeeApiMerchantCreditNewConfigReq(unibeeApiMerchantCreditNewConfigReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditNewConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNewConfigPost`: MerchantCreditEditConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditNewConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditNewConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditNewConfigReq** | [**UnibeeApiMerchantCreditNewConfigReq**](UnibeeApiMerchantCreditNewConfigReq.md) |  | 

### Return type

[**MerchantCreditEditConfigPost200Response**](MerchantCreditEditConfigPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreditNewCreditRechargePost

> MerchantCreditNewCreditRechargePost200Response CreditNewCreditRechargePost(ctx).UnibeeApiMerchantCreditNewCreditRechargeReq(unibeeApiMerchantCreditNewCreditRechargeReq).Execute()

New Credit Recharge



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
	unibeeApiMerchantCreditNewCreditRechargeReq := *openapiclient.NewUnibeeApiMerchantCreditNewCreditRechargeReq(int64(123), int64(123)) // UnibeeApiMerchantCreditNewCreditRechargeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.Credit.CreditNewCreditRechargePost(context.Background()).UnibeeApiMerchantCreditNewCreditRechargeReq(unibeeApiMerchantCreditNewCreditRechargeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Credit.CreditNewCreditRechargePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreditNewCreditRechargePost`: MerchantCreditNewCreditRechargePost200Response
	fmt.Fprintf(os.Stdout, "Response from `Credit.CreditNewCreditRechargePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreditNewCreditRechargePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unibeeApiMerchantCreditNewCreditRechargeReq** | [**UnibeeApiMerchantCreditNewCreditRechargeReq**](UnibeeApiMerchantCreditNewCreditRechargeReq.md) |  | 

### Return type

[**MerchantCreditNewCreditRechargePost200Response**](MerchantCreditNewCreditRechargePost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

