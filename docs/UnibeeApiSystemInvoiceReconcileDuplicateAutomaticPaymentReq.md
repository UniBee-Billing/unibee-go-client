# UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DuplicatePaymentId** | **string** | The successful automatic payment to move onto a new reconciliation invoice | 
**OriginalInvoiceId** | **string** | The paid invoice that currently contains both the normal payment and duplicate automatic payment | 
**SendWebhook** | Pointer to **bool** | Whether to enqueue invoice.created and invoice.paid webhooks, default true. Set false for one-time addon reconciliation to avoid granting entitlements again | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq

`func NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq(duplicatePaymentId string, originalInvoiceId string, ) *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq`

NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq instantiates a new UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReqWithDefaults

`func NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReqWithDefaults() *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq`

NewUnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReqWithDefaults instantiates a new UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicatePaymentId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetDuplicatePaymentId() string`

GetDuplicatePaymentId returns the DuplicatePaymentId field if non-nil, zero value otherwise.

### GetDuplicatePaymentIdOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetDuplicatePaymentIdOk() (*string, bool)`

GetDuplicatePaymentIdOk returns a tuple with the DuplicatePaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatePaymentId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) SetDuplicatePaymentId(v string)`

SetDuplicatePaymentId sets DuplicatePaymentId field to given value.


### GetOriginalInvoiceId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetOriginalInvoiceId() string`

GetOriginalInvoiceId returns the OriginalInvoiceId field if non-nil, zero value otherwise.

### GetOriginalInvoiceIdOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetOriginalInvoiceIdOk() (*string, bool)`

GetOriginalInvoiceIdOk returns a tuple with the OriginalInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInvoiceId

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) SetOriginalInvoiceId(v string)`

SetOriginalInvoiceId sets OriginalInvoiceId field to given value.


### GetSendWebhook

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetSendWebhook() bool`

GetSendWebhook returns the SendWebhook field if non-nil, zero value otherwise.

### GetSendWebhookOk

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) GetSendWebhookOk() (*bool, bool)`

GetSendWebhookOk returns a tuple with the SendWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendWebhook

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) SetSendWebhook(v bool)`

SetSendWebhook sets SendWebhook field to given value.

### HasSendWebhook

`func (o *UnibeeApiSystemInvoiceReconcileDuplicateAutomaticPaymentReq) HasSendWebhook() bool`

HasSendWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


