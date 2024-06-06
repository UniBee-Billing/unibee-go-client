# UnibeeApiMerchantDiscountListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discounts** | Pointer to [**[]UnibeeApiBeanMerchantDiscountCodeSimplify**](UnibeeApiBeanMerchantDiscountCodeSimplify.md) | Discount Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountListRes

`func NewUnibeeApiMerchantDiscountListRes() *UnibeeApiMerchantDiscountListRes`

NewUnibeeApiMerchantDiscountListRes instantiates a new UnibeeApiMerchantDiscountListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountListResWithDefaults

`func NewUnibeeApiMerchantDiscountListResWithDefaults() *UnibeeApiMerchantDiscountListRes`

NewUnibeeApiMerchantDiscountListResWithDefaults instantiates a new UnibeeApiMerchantDiscountListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscounts

`func (o *UnibeeApiMerchantDiscountListRes) GetDiscounts() []UnibeeApiBeanMerchantDiscountCodeSimplify`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *UnibeeApiMerchantDiscountListRes) GetDiscountsOk() (*[]UnibeeApiBeanMerchantDiscountCodeSimplify, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *UnibeeApiMerchantDiscountListRes) SetDiscounts(v []UnibeeApiBeanMerchantDiscountCodeSimplify)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *UnibeeApiMerchantDiscountListRes) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantDiscountListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantDiscountListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantDiscountListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantDiscountListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


