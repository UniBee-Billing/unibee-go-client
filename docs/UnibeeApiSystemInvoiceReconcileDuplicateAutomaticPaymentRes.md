# UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyReconciled** | Pointer to **bool** | Whether the payment had already been reconciled by an earlier call | [optional] 
**InvoiceId** | Pointer to **string** | The reconciliation invoice id | [optional] 
**PaymentId** | Pointer to **string** | The duplicate payment id | [optional] 
**WebhooksSent** | Pointer to **bool** | Whether invoice.created and invoice.paid were enqueued by this call | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes

`func NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes() *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes`

NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes instantiates a new UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentResWithDefaults

`func NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentResWithDefaults() *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes`

NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentResWithDefaults instantiates a new UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyReconciled

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetAlreadyReconciled() bool`

GetAlreadyReconciled returns the AlreadyReconciled field if non-nil, zero value otherwise.

### GetAlreadyReconciledOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetAlreadyReconciledOk() (*bool, bool)`

GetAlreadyReconciledOk returns a tuple with the AlreadyReconciled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyReconciled

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) SetAlreadyReconciled(v bool)`

SetAlreadyReconciled sets AlreadyReconciled field to given value.

### HasAlreadyReconciled

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) HasAlreadyReconciled() bool`

HasAlreadyReconciled returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetWebhooksSent

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetWebhooksSent() bool`

GetWebhooksSent returns the WebhooksSent field if non-nil, zero value otherwise.

### GetWebhooksSentOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) GetWebhooksSentOk() (*bool, bool)`

GetWebhooksSentOk returns a tuple with the WebhooksSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooksSent

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) SetWebhooksSent(v bool)`

SetWebhooksSent sets WebhooksSent field to given value.

### HasWebhooksSent

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentRes) HasWebhooksSent() bool`

HasWebhooksSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


