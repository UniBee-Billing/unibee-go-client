# UnibeeApiMerchantSubscriptionCancelReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNow** | Pointer to **bool** | Default false | [optional] 
**Prorate** | Pointer to **bool** | Prorate Generate Invoice，Default false | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 

## Methods

### NewUnibeeApiMerchantSubscriptionCancelReq

`func NewUnibeeApiMerchantSubscriptionCancelReq(subscriptionId string, ) *UnibeeApiMerchantSubscriptionCancelReq`

NewUnibeeApiMerchantSubscriptionCancelReq instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelReq`

NewUnibeeApiMerchantSubscriptionCancelReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNow

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNow() bool`

GetInvoiceNow returns the InvoiceNow field if non-nil, zero value otherwise.

### GetInvoiceNowOk

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetInvoiceNowOk() (*bool, bool)`

GetInvoiceNowOk returns a tuple with the InvoiceNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNow

`func (o *UnibeeApiMerchantSubscriptionCancelReq) SetInvoiceNow(v bool)`

SetInvoiceNow sets InvoiceNow field to given value.

### HasInvoiceNow

`func (o *UnibeeApiMerchantSubscriptionCancelReq) HasInvoiceNow() bool`

HasInvoiceNow returns a boolean if a field has been set.

### GetProrate

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrate() bool`

GetProrate returns the Prorate field if non-nil, zero value otherwise.

### GetProrateOk

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetProrateOk() (*bool, bool)`

GetProrateOk returns a tuple with the Prorate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrate

`func (o *UnibeeApiMerchantSubscriptionCancelReq) SetProrate(v bool)`

SetProrate sets Prorate field to given value.

### HasProrate

`func (o *UnibeeApiMerchantSubscriptionCancelReq) HasProrate() bool`

HasProrate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


