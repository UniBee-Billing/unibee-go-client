# UnibeeApiSystemInvoiceCreatePaymentRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The payment link or hosted UI link (when amount limit, redirect to split payment) | [optional] 
**PaymentId** | Pointer to **string** | The payment id created (when success) | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceCreatePaymentRes

`func NewUnibeeApiSystemInvoiceCreatePaymentRes() *UnibeeApiSystemInvoiceCreatePaymentRes`

NewUnibeeApiSystemInvoiceCreatePaymentRes instantiates a new UnibeeApiSystemInvoiceCreatePaymentRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceCreatePaymentResWithDefaults

`func NewUnibeeApiSystemInvoiceCreatePaymentResWithDefaults() *UnibeeApiSystemInvoiceCreatePaymentRes`

NewUnibeeApiSystemInvoiceCreatePaymentResWithDefaults instantiates a new UnibeeApiSystemInvoiceCreatePaymentRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiSystemInvoiceCreatePaymentRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


