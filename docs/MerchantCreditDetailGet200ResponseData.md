# MerchantCreditDetailGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAccount** | Pointer to [**UnibeeApiBeanDetailCreditAccountDetail**](UnibeeApiBeanDetailCreditAccountDetail.md) |  | [optional] 
**CreditTransactions** | Pointer to [**[]UnibeeApiBeanCreditTransaction**](UnibeeApiBeanCreditTransaction.md) | Credit Transaction List | [optional] 

## Methods

### NewMerchantCreditDetailGet200ResponseData

`func NewMerchantCreditDetailGet200ResponseData() *MerchantCreditDetailGet200ResponseData`

NewMerchantCreditDetailGet200ResponseData instantiates a new MerchantCreditDetailGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantCreditDetailGet200ResponseDataWithDefaults

`func NewMerchantCreditDetailGet200ResponseDataWithDefaults() *MerchantCreditDetailGet200ResponseData`

NewMerchantCreditDetailGet200ResponseDataWithDefaults instantiates a new MerchantCreditDetailGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAccount

`func (o *MerchantCreditDetailGet200ResponseData) GetCreditAccount() UnibeeApiBeanDetailCreditAccountDetail`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *MerchantCreditDetailGet200ResponseData) GetCreditAccountOk() (*UnibeeApiBeanDetailCreditAccountDetail, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *MerchantCreditDetailGet200ResponseData) SetCreditAccount(v UnibeeApiBeanDetailCreditAccountDetail)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *MerchantCreditDetailGet200ResponseData) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetCreditTransactions

`func (o *MerchantCreditDetailGet200ResponseData) GetCreditTransactions() []UnibeeApiBeanCreditTransaction`

GetCreditTransactions returns the CreditTransactions field if non-nil, zero value otherwise.

### GetCreditTransactionsOk

`func (o *MerchantCreditDetailGet200ResponseData) GetCreditTransactionsOk() (*[]UnibeeApiBeanCreditTransaction, bool)`

GetCreditTransactionsOk returns a tuple with the CreditTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditTransactions

`func (o *MerchantCreditDetailGet200ResponseData) SetCreditTransactions(v []UnibeeApiBeanCreditTransaction)`

SetCreditTransactions sets CreditTransactions field to given value.

### HasCreditTransactions

`func (o *MerchantCreditDetailGet200ResponseData) HasCreditTransactions() bool`

HasCreditTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


