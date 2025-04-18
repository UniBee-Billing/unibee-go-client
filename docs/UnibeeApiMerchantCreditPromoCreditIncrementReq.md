# UnibeeApiMerchantCreditPromoCreditIncrementReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The amount to increase, should greater than 0 | 
**Currency** | **string** | currency of recharge | 
**Description** | Pointer to **string** | description of increase action | [optional] 
**Name** | Pointer to **string** | name of increase action | [optional] 
**UserId** | **int64** | filter id of user | 

## Methods

### NewUnibeeApiMerchantCreditPromoCreditIncrementReq

`func NewUnibeeApiMerchantCreditPromoCreditIncrementReq(amount int64, currency string, userId int64, ) *UnibeeApiMerchantCreditPromoCreditIncrementReq`

NewUnibeeApiMerchantCreditPromoCreditIncrementReq instantiates a new UnibeeApiMerchantCreditPromoCreditIncrementReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditPromoCreditIncrementReqWithDefaults

`func NewUnibeeApiMerchantCreditPromoCreditIncrementReqWithDefaults() *UnibeeApiMerchantCreditPromoCreditIncrementReq`

NewUnibeeApiMerchantCreditPromoCreditIncrementReqWithDefaults instantiates a new UnibeeApiMerchantCreditPromoCreditIncrementReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantCreditPromoCreditIncrementReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


