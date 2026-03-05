# UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External payment transaction id from gateway (for linking) | 
**GatewayId** | Pointer to **int64** | Optional. External gateway id; omit to keep current gateway and use it for auth | [optional] 
**InvoiceId** | **string** | Invoice id | 
**PaymentLink** | **string** | External payment/checkout url shown to user in invoice | 
**SendInvoice** | Pointer to **bool** | Whether to (re)send invoice email after payment link is updated | [optional] 
**Signature** | **string** | HMAC-SHA256(gatewayKey, invoiceId|externalTransactionId|timestamp), hex | 
**Timestamp** | **int64** | Request timestamp, UTC seconds, used for expiry check | 

## Methods

### NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq

`func NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq(externalTransactionId string, invoiceId string, paymentLink string, signature string, timestamp int64, ) *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq`

NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq instantiates a new UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReqWithDefaults

`func NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReqWithDefaults() *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq`

NewUnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetPaymentLink

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.


### GetSendInvoice

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetSendInvoice() bool`

GetSendInvoice returns the SendInvoice field if non-nil, zero value otherwise.

### GetSendInvoiceOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetSendInvoiceOk() (*bool, bool)`

GetSendInvoiceOk returns a tuple with the SendInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvoice

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetSendInvoice(v bool)`

SetSendInvoice sets SendInvoice field to given value.

### HasSendInvoice

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) HasSendInvoice() bool`

HasSendInvoice returns a boolean if a field has been set.

### GetSignature

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantInvoiceUpdateExternalGatewayInvoicePaymentLinkReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


