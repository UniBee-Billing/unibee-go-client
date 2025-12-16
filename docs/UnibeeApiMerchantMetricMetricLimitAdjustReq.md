# UnibeeApiMerchantMetricMetricLimitAdjustReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** | Adjustment amount (positive to increase, negative to decrease, cannot be 0) | 
**MetricCode** | **string** | Metric Code | 
**ProductId** | Pointer to **int64** | Product ID, default 0 for default product | [optional] 
**Reason** | **string** | Reason for adjustment | 
**SubscriptionId** | Pointer to **string** | Subscription ID (priority, use this if you know the subscription) | [optional] 
**UserId** | Pointer to **int32** | User ID (alternative to subscriptionId) | [optional] 

## Methods

### NewUnibeeApiMerchantMetricMetricLimitAdjustReq

`func NewUnibeeApiMerchantMetricMetricLimitAdjustReq(amount int64, metricCode string, reason string, ) *UnibeeApiMerchantMetricMetricLimitAdjustReq`

NewUnibeeApiMerchantMetricMetricLimitAdjustReq instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricMetricLimitAdjustReqWithDefaults

`func NewUnibeeApiMerchantMetricMetricLimitAdjustReqWithDefaults() *UnibeeApiMerchantMetricMetricLimitAdjustReq`

NewUnibeeApiMerchantMetricMetricLimitAdjustReqWithDefaults instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetMetricCode

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.


### GetProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReason

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


