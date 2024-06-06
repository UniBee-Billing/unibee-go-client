# UnibeeApiMerchantPaymentTimeLineListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentTimeLines** | Pointer to [**[]UnibeeApiBeanDetailPaymentTimelineDetail**](UnibeeApiBeanDetailPaymentTimelineDetail.md) | Payment TimeLine Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentTimeLineListRes

`func NewUnibeeApiMerchantPaymentTimeLineListRes() *UnibeeApiMerchantPaymentTimeLineListRes`

NewUnibeeApiMerchantPaymentTimeLineListRes instantiates a new UnibeeApiMerchantPaymentTimeLineListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentTimeLineListResWithDefaults

`func NewUnibeeApiMerchantPaymentTimeLineListResWithDefaults() *UnibeeApiMerchantPaymentTimeLineListRes`

NewUnibeeApiMerchantPaymentTimeLineListResWithDefaults instantiates a new UnibeeApiMerchantPaymentTimeLineListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentTimeLines

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetPaymentTimeLines() []UnibeeApiBeanDetailPaymentTimelineDetail`

GetPaymentTimeLines returns the PaymentTimeLines field if non-nil, zero value otherwise.

### GetPaymentTimeLinesOk

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetPaymentTimeLinesOk() (*[]UnibeeApiBeanDetailPaymentTimelineDetail, bool)`

GetPaymentTimeLinesOk returns a tuple with the PaymentTimeLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTimeLines

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) SetPaymentTimeLines(v []UnibeeApiBeanDetailPaymentTimelineDetail)`

SetPaymentTimeLines sets PaymentTimeLines field to given value.

### HasPaymentTimeLines

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) HasPaymentTimeLines() bool`

HasPaymentTimeLines returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantPaymentTimeLineListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


