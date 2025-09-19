# UnibeeApiMerchantCreditCreditTransactionListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **int32** | filter type of account, 1-main account, 2-promo credit account | 
**Count** | Pointer to **int32** | Count Of Per Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**Currency** | Pointer to **string** | filter currency of account | [optional] 
**Email** | Pointer to **string** | filter email of user | [optional] 
**Page** | Pointer to **int32** | Page, Start 0 | [optional] 
**SortField** | Pointer to **string** | Sort Field，gmt_create|gmt_modify，Default gmt_modify | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**TransactionTypes** | Pointer to **[]int32** | transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out | [optional] 
**UserId** | Pointer to **int64** | filter id of user | [optional] 

## Methods

### NewUnibeeApiMerchantCreditCreditTransactionListReq

`func NewUnibeeApiMerchantCreditCreditTransactionListReq(accountType int32, ) *UnibeeApiMerchantCreditCreditTransactionListReq`

NewUnibeeApiMerchantCreditCreditTransactionListReq instantiates a new UnibeeApiMerchantCreditCreditTransactionListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditCreditTransactionListReqWithDefaults

`func NewUnibeeApiMerchantCreditCreditTransactionListReqWithDefaults() *UnibeeApiMerchantCreditCreditTransactionListReq`

NewUnibeeApiMerchantCreditCreditTransactionListReqWithDefaults instantiates a new UnibeeApiMerchantCreditCreditTransactionListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.


### GetCount

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetTransactionTypes

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetTransactionTypes() []int32`

GetTransactionTypes returns the TransactionTypes field if non-nil, zero value otherwise.

### GetTransactionTypesOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetTransactionTypesOk() (*[]int32, bool)`

GetTransactionTypesOk returns a tuple with the TransactionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypes

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetTransactionTypes(v []int32)`

SetTransactionTypes sets TransactionTypes field to given value.

### HasTransactionTypes

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasTransactionTypes() bool`

HasTransactionTypes returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantCreditCreditTransactionListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


