# MerchantPaymentTimelineListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentTimeLines** | Pointer to [**[]UnibeeApiBeanDetailPaymentTimelineDetail**](UnibeeApiBeanDetailPaymentTimelineDetail.md) | Payment TimeLine Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantPaymentTimelineListGet200ResponseData

`func NewMerchantPaymentTimelineListGet200ResponseData() *MerchantPaymentTimelineListGet200ResponseData`

NewMerchantPaymentTimelineListGet200ResponseData instantiates a new MerchantPaymentTimelineListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaymentTimelineListGet200ResponseDataWithDefaults

`func NewMerchantPaymentTimelineListGet200ResponseDataWithDefaults() *MerchantPaymentTimelineListGet200ResponseData`

NewMerchantPaymentTimelineListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentTimelineListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentTimeLines

`func (o *MerchantPaymentTimelineListGet200ResponseData) GetPaymentTimeLines() []UnibeeApiBeanDetailPaymentTimelineDetail`

GetPaymentTimeLines returns the PaymentTimeLines field if non-nil, zero value otherwise.

### GetPaymentTimeLinesOk

`func (o *MerchantPaymentTimelineListGet200ResponseData) GetPaymentTimeLinesOk() (*[]UnibeeApiBeanDetailPaymentTimelineDetail, bool)`

GetPaymentTimeLinesOk returns a tuple with the PaymentTimeLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTimeLines

`func (o *MerchantPaymentTimelineListGet200ResponseData) SetPaymentTimeLines(v []UnibeeApiBeanDetailPaymentTimelineDetail)`

SetPaymentTimeLines sets PaymentTimeLines field to given value.

### HasPaymentTimeLines

`func (o *MerchantPaymentTimelineListGet200ResponseData) HasPaymentTimeLines() bool`

HasPaymentTimeLines returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantPaymentTimelineListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantPaymentTimelineListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantPaymentTimelineListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantPaymentTimelineListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


