# MerchantPaymentRefundListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundDetails** | Pointer to [**[]UnibeeApiBeanDetailRefundDetail**](UnibeeApiBeanDetailRefundDetail.md) | Refund Detail Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantPaymentRefundListGet200ResponseData

`func NewMerchantPaymentRefundListGet200ResponseData() *MerchantPaymentRefundListGet200ResponseData`

NewMerchantPaymentRefundListGet200ResponseData instantiates a new MerchantPaymentRefundListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaymentRefundListGet200ResponseDataWithDefaults

`func NewMerchantPaymentRefundListGet200ResponseDataWithDefaults() *MerchantPaymentRefundListGet200ResponseData`

NewMerchantPaymentRefundListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentRefundListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundDetails

`func (o *MerchantPaymentRefundListGet200ResponseData) GetRefundDetails() []UnibeeApiBeanDetailRefundDetail`

GetRefundDetails returns the RefundDetails field if non-nil, zero value otherwise.

### GetRefundDetailsOk

`func (o *MerchantPaymentRefundListGet200ResponseData) GetRefundDetailsOk() (*[]UnibeeApiBeanDetailRefundDetail, bool)`

GetRefundDetailsOk returns a tuple with the RefundDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundDetails

`func (o *MerchantPaymentRefundListGet200ResponseData) SetRefundDetails(v []UnibeeApiBeanDetailRefundDetail)`

SetRefundDetails sets RefundDetails field to given value.

### HasRefundDetails

`func (o *MerchantPaymentRefundListGet200ResponseData) HasRefundDetails() bool`

HasRefundDetails returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantPaymentRefundListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantPaymentRefundListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantPaymentRefundListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantPaymentRefundListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


