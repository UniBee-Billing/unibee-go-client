# UnibeeApiSystemPaymentStripeSmartRetryDetailReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** | UniBee invoice ID; queries all related payments when paymentId is empty | [optional] 
**PaymentId** | Pointer to **string** | Optional UniBee payment ID; limits the query to one payment | [optional] 

## Methods

### NewUnibeeApiSystemPaymentStripeSmartRetryDetailReq

`func NewUnibeeApiSystemPaymentStripeSmartRetryDetailReq() *UnibeeApiSystemPaymentStripeSmartRetryDetailReq`

NewUnibeeApiSystemPaymentStripeSmartRetryDetailReq instantiates a new UnibeeApiSystemPaymentStripeSmartRetryDetailReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemPaymentStripeSmartRetryDetailReqWithDefaults

`func NewUnibeeApiSystemPaymentStripeSmartRetryDetailReqWithDefaults() *UnibeeApiSystemPaymentStripeSmartRetryDetailReq`

NewUnibeeApiSystemPaymentStripeSmartRetryDetailReqWithDefaults instantiates a new UnibeeApiSystemPaymentStripeSmartRetryDetailReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiSystemPaymentStripeSmartRetryDetailReq) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


