# UnibeeApiMerchantCreditDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAccount** | Pointer to [**UnibeeApiBeanDetailCreditAccountDetail**](UnibeeApiBeanDetailCreditAccountDetail.md) |  | [optional] 
**CreditTransactions** | Pointer to [**[]UnibeeApiBeanCreditTransaction**](UnibeeApiBeanCreditTransaction.md) | Credit Transaction List | [optional] 

## Methods

### NewUnibeeApiMerchantCreditDetailRes

`func NewUnibeeApiMerchantCreditDetailRes() *UnibeeApiMerchantCreditDetailRes`

NewUnibeeApiMerchantCreditDetailRes instantiates a new UnibeeApiMerchantCreditDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditDetailResWithDefaults

`func NewUnibeeApiMerchantCreditDetailResWithDefaults() *UnibeeApiMerchantCreditDetailRes`

NewUnibeeApiMerchantCreditDetailResWithDefaults instantiates a new UnibeeApiMerchantCreditDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAccount

`func (o *UnibeeApiMerchantCreditDetailRes) GetCreditAccount() UnibeeApiBeanDetailCreditAccountDetail`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *UnibeeApiMerchantCreditDetailRes) GetCreditAccountOk() (*UnibeeApiBeanDetailCreditAccountDetail, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *UnibeeApiMerchantCreditDetailRes) SetCreditAccount(v UnibeeApiBeanDetailCreditAccountDetail)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *UnibeeApiMerchantCreditDetailRes) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetCreditTransactions

`func (o *UnibeeApiMerchantCreditDetailRes) GetCreditTransactions() []UnibeeApiBeanCreditTransaction`

GetCreditTransactions returns the CreditTransactions field if non-nil, zero value otherwise.

### GetCreditTransactionsOk

`func (o *UnibeeApiMerchantCreditDetailRes) GetCreditTransactionsOk() (*[]UnibeeApiBeanCreditTransaction, bool)`

GetCreditTransactionsOk returns a tuple with the CreditTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTransactions

`func (o *UnibeeApiMerchantCreditDetailRes) SetCreditTransactions(v []UnibeeApiBeanCreditTransaction)`

SetCreditTransactions sets CreditTransactions field to given value.

### HasCreditTransactions

`func (o *UnibeeApiMerchantCreditDetailRes) HasCreditTransactions() bool`

HasCreditTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


