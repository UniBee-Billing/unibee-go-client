# UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** | InvoiceId to regenerate, defaults to subscription latest invoice | [optional] 
**Reason** | **string** | Operation reason for audit log | 
**SubscriptionId** | **string** | SubscriptionId | 

## Methods

### NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq

`func NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq(reason string, subscriptionId string, ) *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq`

NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq instantiates a new UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReqWithDefaults() *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq`

NewUnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetReason

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRegenerateBillingCycleInvoiceReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


