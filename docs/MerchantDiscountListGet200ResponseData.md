# MerchantDiscountListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discounts** | Pointer to [**[]UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) | Discount Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantDiscountListGet200ResponseData

`func NewMerchantDiscountListGet200ResponseData() *MerchantDiscountListGet200ResponseData`

NewMerchantDiscountListGet200ResponseData instantiates a new MerchantDiscountListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDiscountListGet200ResponseDataWithDefaults

`func NewMerchantDiscountListGet200ResponseDataWithDefaults() *MerchantDiscountListGet200ResponseData`

NewMerchantDiscountListGet200ResponseDataWithDefaults instantiates a new MerchantDiscountListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscounts

`func (o *MerchantDiscountListGet200ResponseData) GetDiscounts() []UnibeeApiBeanMerchantDiscountCode`

GetDiscounts returns the Discounts field if non-nil, zero value otherwise.

### GetDiscountsOk

`func (o *MerchantDiscountListGet200ResponseData) GetDiscountsOk() (*[]UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountsOk returns a tuple with the Discounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscounts

`func (o *MerchantDiscountListGet200ResponseData) SetDiscounts(v []UnibeeApiBeanMerchantDiscountCode)`

SetDiscounts sets Discounts field to given value.

### HasDiscounts

`func (o *MerchantDiscountListGet200ResponseData) HasDiscounts() bool`

HasDiscounts returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantDiscountListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantDiscountListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantDiscountListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantDiscountListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


