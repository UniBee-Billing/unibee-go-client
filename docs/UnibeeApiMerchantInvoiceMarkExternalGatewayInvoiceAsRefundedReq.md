# UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External refund transaction id from gateway | 
**GatewayId** | Pointer to **int64** | Optional. External gateway id; omit to use invoice&#39;s current gateway for auth | [optional] 
**InvoiceId** | **string** | Invoice id | 
**Signature** | **string** | HMAC-SHA256(gatewayKey, invoiceId|externalTransactionId|timestamp), hex | 
**Timestamp** | **int64** | Request timestamp, UTC seconds, used for expiry check | 

## Methods

### NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq

`func NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq(externalTransactionId string, invoiceId string, signature string, timestamp int64, ) *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq`

NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq instantiates a new UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReqWithDefaults

`func NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReqWithDefaults() *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq`

NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetSignature

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsRefundedReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


