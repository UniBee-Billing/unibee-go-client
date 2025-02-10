# UnibeeApiMerchantDiscountQuantityDecrementReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | The discount quantity Amount to decrease, should greater than 0 | [optional] 
**Id** | **int64** | The discount&#39;s Id | 

## Methods

### NewUnibeeApiMerchantDiscountQuantityDecrementReq

`func NewUnibeeApiMerchantDiscountQuantityDecrementReq(id int64, ) *UnibeeApiMerchantDiscountQuantityDecrementReq`

NewUnibeeApiMerchantDiscountQuantityDecrementReq instantiates a new UnibeeApiMerchantDiscountQuantityDecrementReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountQuantityDecrementReqWithDefaults

`func NewUnibeeApiMerchantDiscountQuantityDecrementReqWithDefaults() *UnibeeApiMerchantDiscountQuantityDecrementReq`

NewUnibeeApiMerchantDiscountQuantityDecrementReqWithDefaults instantiates a new UnibeeApiMerchantDiscountQuantityDecrementReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantDiscountQuantityDecrementReq) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


