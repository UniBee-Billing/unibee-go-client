# UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External payment transaction id from gateway, used as idempotent key | 
**GatewayId** | Pointer to **int64** | External gateway id (optional), if omitted will be resolved from paymentId | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Additional metadata from external gateway | [optional] 
**PaidTime** | Pointer to **int64** | Payment success time (unix timestamp in seconds, optional) | [optional] 
**PaymentId** | **string** | The UniBee paymentId | 
**Signature** | **string** | HMAC-SHA256 hex signature using External Gateway API Key over &#39;paymentId|externalTransactionId|timestamp&#39; | 
**Timestamp** | **int64** | Unix timestamp in seconds, used for signature and anti-replay | 

## Methods

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq(externalTransactionId string, paymentId string, signature string, timestamp int64, ) *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReqWithDefaults

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReqWithDefaults() *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReqWithDefaults instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkPaidReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


