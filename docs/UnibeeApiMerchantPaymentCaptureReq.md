# UnibeeApiMerchantPaymentCaptureReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaptureAmount** | **int64** | CaptureAmount, Cent | 
**Currency** | **string** | Currency | 
**ExternalCaptureId** | **string** | ExternalCaptureId | 
**PaymentId** | **string** | PaymentId | 

## Methods

### NewUnibeeApiMerchantPaymentCaptureReq

`func NewUnibeeApiMerchantPaymentCaptureReq(captureAmount int64, currency string, externalCaptureId string, paymentId string, ) *UnibeeApiMerchantPaymentCaptureReq`

NewUnibeeApiMerchantPaymentCaptureReq instantiates a new UnibeeApiMerchantPaymentCaptureReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentCaptureReqWithDefaults

`func NewUnibeeApiMerchantPaymentCaptureReqWithDefaults() *UnibeeApiMerchantPaymentCaptureReq`

NewUnibeeApiMerchantPaymentCaptureReqWithDefaults instantiates a new UnibeeApiMerchantPaymentCaptureReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptureAmount

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetCaptureAmount() int64`

GetCaptureAmount returns the CaptureAmount field if non-nil, zero value otherwise.

### GetCaptureAmountOk

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetCaptureAmountOk() (*int64, bool)`

GetCaptureAmountOk returns a tuple with the CaptureAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureAmount

`func (o *UnibeeApiMerchantPaymentCaptureReq) SetCaptureAmount(v int64)`

SetCaptureAmount sets CaptureAmount field to given value.


### GetCurrency

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentCaptureReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExternalCaptureId

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetExternalCaptureId() string`

GetExternalCaptureId returns the ExternalCaptureId field if non-nil, zero value otherwise.

### GetExternalCaptureIdOk

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetExternalCaptureIdOk() (*string, bool)`

GetExternalCaptureIdOk returns a tuple with the ExternalCaptureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalCaptureId

`func (o *UnibeeApiMerchantPaymentCaptureReq) SetExternalCaptureId(v string)`

SetExternalCaptureId sets ExternalCaptureId field to given value.


### GetPaymentId

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentCaptureReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentCaptureReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


