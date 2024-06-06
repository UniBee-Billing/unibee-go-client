# MerchantPaymentListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentDetails** | Pointer to [**[]UnibeeApiBeanDetailPaymentDetail**](UnibeeApiBeanDetailPaymentDetail.md) | Payment Detail Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantPaymentListGet200ResponseData

`func NewMerchantPaymentListGet200ResponseData() *MerchantPaymentListGet200ResponseData`

NewMerchantPaymentListGet200ResponseData instantiates a new MerchantPaymentListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaymentListGet200ResponseDataWithDefaults

`func NewMerchantPaymentListGet200ResponseDataWithDefaults() *MerchantPaymentListGet200ResponseData`

NewMerchantPaymentListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentDetails

`func (o *MerchantPaymentListGet200ResponseData) GetPaymentDetails() []UnibeeApiBeanDetailPaymentDetail`

GetPaymentDetails returns the PaymentDetails field if non-nil, zero value otherwise.

### GetPaymentDetailsOk

`func (o *MerchantPaymentListGet200ResponseData) GetPaymentDetailsOk() (*[]UnibeeApiBeanDetailPaymentDetail, bool)`

GetPaymentDetailsOk returns a tuple with the PaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetails

`func (o *MerchantPaymentListGet200ResponseData) SetPaymentDetails(v []UnibeeApiBeanDetailPaymentDetail)`

SetPaymentDetails sets PaymentDetails field to given value.

### HasPaymentDetails

`func (o *MerchantPaymentListGet200ResponseData) HasPaymentDetails() bool`

HasPaymentDetails returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantPaymentListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantPaymentListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantPaymentListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantPaymentListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


