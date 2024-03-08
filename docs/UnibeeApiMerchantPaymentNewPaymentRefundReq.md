# UnibeeApiMerchantPaymentNewPaymentRefundReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency | 
**ExternalRefundId** | **string** | ExternalRefundId | 
**PaymentId** | **string** | PaymentId | 
**Reason** | Pointer to **string** | Reason | [optional] 
**RefundAmount** | **int64** | RefundAmount, Cent | 

## Methods

### NewUnibeeApiMerchantPaymentNewPaymentRefundReq

`func NewUnibeeApiMerchantPaymentNewPaymentRefundReq(currency string, externalRefundId string, paymentId string, refundAmount int64, ) *UnibeeApiMerchantPaymentNewPaymentRefundReq`

NewUnibeeApiMerchantPaymentNewPaymentRefundReq instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentNewPaymentRefundReqWithDefaults

`func NewUnibeeApiMerchantPaymentNewPaymentRefundReqWithDefaults() *UnibeeApiMerchantPaymentNewPaymentRefundReq`

NewUnibeeApiMerchantPaymentNewPaymentRefundReqWithDefaults instantiates a new UnibeeApiMerchantPaymentNewPaymentRefundReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExternalRefundId

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetExternalRefundId() string`

GetExternalRefundId returns the ExternalRefundId field if non-nil, zero value otherwise.

### GetExternalRefundIdOk

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetExternalRefundIdOk() (*string, bool)`

GetExternalRefundIdOk returns a tuple with the ExternalRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefundId

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetExternalRefundId(v string)`

SetExternalRefundId sets ExternalRefundId field to given value.


### GetPaymentId

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetReason

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeApiMerchantPaymentNewPaymentRefundReq) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


