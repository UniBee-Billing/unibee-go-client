# UnibeeApiMerchantDiscountQuantityIncrementReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The discount quantity amount to increase, should greater than 0 | [optional] 
**Id** | **int64** | The discount&#39;s Id | 

## Methods

### NewUnibeeApiMerchantDiscountQuantityIncrementReq

`func NewUnibeeApiMerchantDiscountQuantityIncrementReq(id int64, ) *UnibeeApiMerchantDiscountQuantityIncrementReq`

NewUnibeeApiMerchantDiscountQuantityIncrementReq instantiates a new UnibeeApiMerchantDiscountQuantityIncrementReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountQuantityIncrementReqWithDefaults

`func NewUnibeeApiMerchantDiscountQuantityIncrementReqWithDefaults() *UnibeeApiMerchantDiscountQuantityIncrementReq`

NewUnibeeApiMerchantDiscountQuantityIncrementReqWithDefaults instantiates a new UnibeeApiMerchantDiscountQuantityIncrementReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantDiscountQuantityIncrementReq) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


