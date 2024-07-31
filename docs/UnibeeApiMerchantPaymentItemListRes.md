# UnibeeApiMerchantPaymentItemListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentItems** | Pointer to [**[]UnibeeApiBeanPaymentItem**](UnibeeApiBeanPaymentItem.md) | Payment Item Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentItemListRes

`func NewUnibeeApiMerchantPaymentItemListRes() *UnibeeApiMerchantPaymentItemListRes`

NewUnibeeApiMerchantPaymentItemListRes instantiates a new UnibeeApiMerchantPaymentItemListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentItemListResWithDefaults

`func NewUnibeeApiMerchantPaymentItemListResWithDefaults() *UnibeeApiMerchantPaymentItemListRes`

NewUnibeeApiMerchantPaymentItemListResWithDefaults instantiates a new UnibeeApiMerchantPaymentItemListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentItems

`func (o *UnibeeApiMerchantPaymentItemListRes) GetPaymentItems() []UnibeeApiBeanPaymentItem`

GetPaymentItems returns the PaymentItems field if non-nil, zero value otherwise.

### GetPaymentItemsOk

`func (o *UnibeeApiMerchantPaymentItemListRes) GetPaymentItemsOk() (*[]UnibeeApiBeanPaymentItem, bool)`

GetPaymentItemsOk returns a tuple with the PaymentItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentItems

`func (o *UnibeeApiMerchantPaymentItemListRes) SetPaymentItems(v []UnibeeApiBeanPaymentItem)`

SetPaymentItems sets PaymentItems field to given value.

### HasPaymentItems

`func (o *UnibeeApiMerchantPaymentItemListRes) HasPaymentItems() bool`

HasPaymentItems returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantPaymentItemListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantPaymentItemListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantPaymentItemListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantPaymentItemListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


