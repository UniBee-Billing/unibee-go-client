# UnibeeApiMerchantPaymentListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentDetails** | Pointer to [**[]UnibeeApiBeanDetailPaymentDetail**](UnibeeApiBeanDetailPaymentDetail.md) | Payment Detail Object List | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentListRes

`func NewUnibeeApiMerchantPaymentListRes() *UnibeeApiMerchantPaymentListRes`

NewUnibeeApiMerchantPaymentListRes instantiates a new UnibeeApiMerchantPaymentListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentListResWithDefaults

`func NewUnibeeApiMerchantPaymentListResWithDefaults() *UnibeeApiMerchantPaymentListRes`

NewUnibeeApiMerchantPaymentListResWithDefaults instantiates a new UnibeeApiMerchantPaymentListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentDetails

`func (o *UnibeeApiMerchantPaymentListRes) GetPaymentDetails() []UnibeeApiBeanDetailPaymentDetail`

GetPaymentDetails returns the PaymentDetails field if non-nil, zero value otherwise.

### GetPaymentDetailsOk

`func (o *UnibeeApiMerchantPaymentListRes) GetPaymentDetailsOk() (*[]UnibeeApiBeanDetailPaymentDetail, bool)`

GetPaymentDetailsOk returns a tuple with the PaymentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDetails

`func (o *UnibeeApiMerchantPaymentListRes) SetPaymentDetails(v []UnibeeApiBeanDetailPaymentDetail)`

SetPaymentDetails sets PaymentDetails field to given value.

### HasPaymentDetails

`func (o *UnibeeApiMerchantPaymentListRes) HasPaymentDetails() bool`

HasPaymentDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


