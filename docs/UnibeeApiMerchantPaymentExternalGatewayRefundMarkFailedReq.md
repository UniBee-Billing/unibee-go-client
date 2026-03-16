# UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRefundId** | **string** | External refund transaction id from gateway, used as idempotent key | 
**FailureReason** | **string** | Human-readable failure reason from external gateway | 
**GatewayId** | Pointer to **int64** | External gateway id (optional), if omitted will be resolved from refundId | [optional] 
**RefundId** | **string** | The UniBee refundId | 
**Signature** | **string** | HMAC-SHA256 hex signature using External Gateway API Key over &#39;refundId|externalRefundId|timestamp&#39; | 
**Timestamp** | **int64** | Unix timestamp in seconds, used for signature and anti-replay | 

## Methods

### NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq

`func NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq(externalRefundId string, failureReason string, refundId string, signature string, timestamp int64, ) *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq`

NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq instantiates a new UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReqWithDefaults

`func NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReqWithDefaults() *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq`

NewUnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReqWithDefaults instantiates a new UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRefundId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetExternalRefundId() string`

GetExternalRefundId returns the ExternalRefundId field if non-nil, zero value otherwise.

### GetExternalRefundIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetExternalRefundIdOk() (*string, bool)`

GetExternalRefundIdOk returns a tuple with the ExternalRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefundId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetExternalRefundId(v string)`

SetExternalRefundId sets ExternalRefundId field to given value.


### GetFailureReason

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.


### GetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UnibeeApiMerchantPaymentExternalGatewayRefundMarkFailedReq) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


