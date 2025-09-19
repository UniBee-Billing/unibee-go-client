# UnibeeApiSystemPaymentDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**PaymentStatus** | Pointer to **int32** | Payment Status，10-pending，20-success，30-failure, 40-cancel | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiSystemPaymentDetailRes

`func NewUnibeeApiSystemPaymentDetailRes() *UnibeeApiSystemPaymentDetailRes`

NewUnibeeApiSystemPaymentDetailRes instantiates a new UnibeeApiSystemPaymentDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemPaymentDetailResWithDefaults

`func NewUnibeeApiSystemPaymentDetailResWithDefaults() *UnibeeApiSystemPaymentDetailRes`

NewUnibeeApiSystemPaymentDetailResWithDefaults instantiates a new UnibeeApiSystemPaymentDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiSystemPaymentDetailRes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiSystemPaymentDetailRes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiSystemPaymentDetailRes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiSystemPaymentDetailRes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiSystemPaymentDetailRes) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiSystemPaymentDetailRes) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiSystemPaymentDetailRes) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiSystemPaymentDetailRes) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *UnibeeApiSystemPaymentDetailRes) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *UnibeeApiSystemPaymentDetailRes) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *UnibeeApiSystemPaymentDetailRes) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *UnibeeApiSystemPaymentDetailRes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiSystemPaymentDetailRes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiSystemPaymentDetailRes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiSystemPaymentDetailRes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiSystemPaymentDetailRes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


