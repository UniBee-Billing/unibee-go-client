# MerchantPaymentRefundNewPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRefundId** | **string** | ExternalRefundId | 
**PaymentId** | **string** | PaymentId | 
**Reason** | Pointer to **string** | Reason | [optional] 

## Methods

### NewMerchantPaymentRefundNewPostRequest

`func NewMerchantPaymentRefundNewPostRequest(externalRefundId string, paymentId string, ) *MerchantPaymentRefundNewPostRequest`

NewMerchantPaymentRefundNewPostRequest instantiates a new MerchantPaymentRefundNewPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaymentRefundNewPostRequestWithDefaults

`func NewMerchantPaymentRefundNewPostRequestWithDefaults() *MerchantPaymentRefundNewPostRequest`

NewMerchantPaymentRefundNewPostRequestWithDefaults instantiates a new MerchantPaymentRefundNewPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRefundId

`func (o *MerchantPaymentRefundNewPostRequest) GetExternalRefundId() string`

GetExternalRefundId returns the ExternalRefundId field if non-nil, zero value otherwise.

### GetExternalRefundIdOk

`func (o *MerchantPaymentRefundNewPostRequest) GetExternalRefundIdOk() (*string, bool)`

GetExternalRefundIdOk returns a tuple with the ExternalRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefundId

`func (o *MerchantPaymentRefundNewPostRequest) SetExternalRefundId(v string)`

SetExternalRefundId sets ExternalRefundId field to given value.


### GetPaymentId

`func (o *MerchantPaymentRefundNewPostRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *MerchantPaymentRefundNewPostRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *MerchantPaymentRefundNewPostRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetReason

`func (o *MerchantPaymentRefundNewPostRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MerchantPaymentRefundNewPostRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MerchantPaymentRefundNewPostRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MerchantPaymentRefundNewPostRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


