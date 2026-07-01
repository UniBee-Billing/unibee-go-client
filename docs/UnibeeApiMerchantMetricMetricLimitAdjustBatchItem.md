# UnibeeApiMerchantMetricMetricLimitAdjustBatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | Adjustment amount (positive to increase, negative to decrease, cannot be 0) | [optional] 
**IdempotencyKey** | Pointer to **string** | Optional item-level idempotency key, max 256 chars. Has higher priority than batch idempotencyKey. | [optional] 
**MetricCode** | Pointer to **string** | Metric Code | [optional] 
**ProductId** | Pointer to **int64** | Product ID, default 0 for default product | [optional] 
**Reason** | Pointer to **string** | Reason for adjustment | [optional] 
**SubscriptionId** | Pointer to **string** | Subscription ID (priority, use this if you know the subscription) | [optional] 
**UserId** | Pointer to **int32** | User ID (alternative to subscriptionId) | [optional] 

## Methods

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItem

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItem() *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItem instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItemWithDefaults

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItemWithDefaults() *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchItemWithDefaults instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.

### HasMetricCode

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasMetricCode() bool`

HasMetricCode returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReason

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


