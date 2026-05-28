# UnibeeApiMerchantCreditPromoCreditDecrementReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | The Amount to decrease, should greater than 0 | 
**Currency** | **string** | currency of recharge | 
**Description** | Pointer to **string** | description of increase action | [optional] 
**IdempotencyKey** | Pointer to **string** | Optional. Client-supplied idempotency key, max 256 chars (UUID recommended). Retrying the same key returns the original credit transaction and does NOT decrease the balance a second time. Reusing the same key with a different amount or currency is rejected. Increment and decrement use separate namespaces and are isolated per merchant and user, so reusing the same key on the increment endpoint, under another merchant, or for another user does not collide. If two requests with the same key are still in-flight, the later one returns an error and should be retried by the caller. | [optional] 
**Name** | Pointer to **string** | name of increase action | [optional] 
**UserId** | **int64** | filter id of user | 

## Methods

### NewUnibeeApiMerchantCreditPromoCreditDecrementReq

`func NewUnibeeApiMerchantCreditPromoCreditDecrementReq(amount int64, currency string, userId int64, ) *UnibeeApiMerchantCreditPromoCreditDecrementReq`

NewUnibeeApiMerchantCreditPromoCreditDecrementReq instantiates a new UnibeeApiMerchantCreditPromoCreditDecrementReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditPromoCreditDecrementReqWithDefaults

`func NewUnibeeApiMerchantCreditPromoCreditDecrementReqWithDefaults() *UnibeeApiMerchantCreditPromoCreditDecrementReq`

NewUnibeeApiMerchantCreditPromoCreditDecrementReqWithDefaults instantiates a new UnibeeApiMerchantCreditPromoCreditDecrementReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantCreditPromoCreditDecrementReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


