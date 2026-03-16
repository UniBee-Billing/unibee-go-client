# UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External payment transaction id from gateway, used as idempotent key | 
**GatewayId** | Pointer to **int64** | External gateway id (optional), if omitted will be resolved from paymentId | [optional] 
**PaymentId** | **string** | The UniBee paymentId | 
**PaymentLink** | **string** | External checkout URL for this payment | 
**SendInvoice** | Pointer to **bool** | Whether to send invoice email to customer after updating link | [optional] 
**Signature** | **string** | HMAC-SHA256 hex signature using External Gateway API Key over &#39;paymentId|externalTransactionId|timestamp&#39; | 
**Timestamp** | **int64** | Unix timestamp in seconds, used for signature and anti-replay | 

## Methods

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq(externalTransactionId string, paymentId string, paymentLink string, signature string, timestamp int64, ) *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReqWithDefaults

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReqWithDefaults() *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReqWithDefaults instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetPaymentLink

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.


### GetSendInvoice

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetSendInvoice() bool`

GetSendInvoice returns the SendInvoice field if non-nil, zero value otherwise.

### GetSendInvoiceOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetSendInvoiceOk() (*bool, bool)`

GetSendInvoiceOk returns a tuple with the SendInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvoice

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetSendInvoice(v bool)`

SetSendInvoice sets SendInvoice field to given value.

### HasSendInvoice

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) HasSendInvoice() bool`

HasSendInvoice returns a boolean if a field has been set.

### GetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentUpdateLinkReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


