# UnibeeApiCheckoutPaymentDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**PaymentStatus** | Pointer to **int32** | Payment Status，10-pending，20-success，30-failure, 40-cancel | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiCheckoutPaymentDetailRes

`func NewUnibeeApiCheckoutPaymentDetailRes() *UnibeeApiCheckoutPaymentDetailRes`

NewUnibeeApiCheckoutPaymentDetailRes instantiates a new UnibeeApiCheckoutPaymentDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutPaymentDetailResWithDefaults

`func NewUnibeeApiCheckoutPaymentDetailResWithDefaults() *UnibeeApiCheckoutPaymentDetailRes`

NewUnibeeApiCheckoutPaymentDetailResWithDefaults instantiates a new UnibeeApiCheckoutPaymentDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiCheckoutPaymentDetailRes) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiCheckoutPaymentDetailRes) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetPaymentStatus() int32`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetPaymentStatusOk() (*int32, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *UnibeeApiCheckoutPaymentDetailRes) SetPaymentStatus(v int32)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *UnibeeApiCheckoutPaymentDetailRes) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiCheckoutPaymentDetailRes) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiCheckoutPaymentDetailRes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


