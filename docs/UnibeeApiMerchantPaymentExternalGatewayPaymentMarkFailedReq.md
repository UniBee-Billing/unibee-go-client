# UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalTransactionId** | **string** | External payment transaction id from gateway, used for correlation | 
**FailureReason** | Pointer to **string** | Human-readable failure reason from external gateway | [optional] 
**GatewayId** | Pointer to **int64** | External gateway id (optional), if omitted will be resolved from paymentId | [optional] 
**PaymentId** | **string** | The UniBee paymentId | 
**Signature** | **string** | HMAC-SHA256 hex signature using External Gateway API Key over &#39;paymentId|externalTransactionId|timestamp&#39; | 
**Timestamp** | **int64** | Unix timestamp in seconds, used for signature and anti-replay | 

## Methods

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq(externalTransactionId string, paymentId string, signature string, timestamp int64, ) *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReqWithDefaults

`func NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReqWithDefaults() *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq`

NewUnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReqWithDefaults instantiates a new UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.


### GetFailureReason

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayPaymentMarkFailedReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


