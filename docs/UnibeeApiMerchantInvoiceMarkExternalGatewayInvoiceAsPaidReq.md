# UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External payment transaction id from gateway, used as idempotent key | 
**GatewayId** | Pointer to **int64** | Optional. External gateway id; omit to use invoice&#39;s current gateway for auth | [optional] 
**InvoiceId** | **string** | Invoice id | 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Additional metadata from external gateway | [optional] 
**PaidTime** | Pointer to **int64** | Paid time, UTC timestamp in seconds (optional, default now) | [optional] 
**Signature** | **string** | HMAC-SHA256(gatewayKey, invoiceId|externalTransactionId|timestamp), hex-encoded | 
**Timestamp** | **int64** | Request timestamp, UTC seconds, used for expiry check | 

## Methods

### NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq

`func NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq(externalTransactionId string, invoiceId string, signature string, timestamp int64, ) *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq`

NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq instantiates a new UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReqWithDefaults

`func NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReqWithDefaults() *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq`

NewUnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetSignature

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantInvoiceMarkExternalGatewayInvoiceAsPaidReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


