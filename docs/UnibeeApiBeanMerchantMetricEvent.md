# UnibeeApiBeanMerchantMetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationPropertyData** | Pointer to **string** | aggregation property data (Json) | [optional] 
**AggregationPropertyInt** | Pointer to **int64** | aggregation property int, use for metric of max|sum type | [optional] 
**AggregationPropertyString** | Pointer to **string** | aggregation property string, use for metric of count|count_unique type | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**ExternalEventId** | Pointer to **string** | external_event_id, should be unique | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MetricId** | Pointer to **int64** | metric_id | [optional] 
**MetricLimit** | Pointer to **int64** |  | [optional] 
**SubscriptionIds** | Pointer to **string** |  | [optional] 
**SubscriptionPeriodEnd** | Pointer to **int64** | matched subscription&#39;s current_period_end | [optional] 
**SubscriptionPeriodStart** | Pointer to **int64** | matched subscription&#39;s current_period_start | [optional] 
**Used** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMetricEvent

`func NewUnibeeApiBeanMerchantMetricEvent() *UnibeeApiBeanMerchantMetricEvent`

NewUnibeeApiBeanMerchantMetricEvent instantiates a new UnibeeApiBeanMerchantMetricEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMetricEventWithDefaults

`func NewUnibeeApiBeanMerchantMetricEventWithDefaults() *UnibeeApiBeanMerchantMetricEvent`

NewUnibeeApiBeanMerchantMetricEventWithDefaults instantiates a new UnibeeApiBeanMerchantMetricEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyData() string`

GetAggregationPropertyData returns the AggregationPropertyData field if non-nil, zero value otherwise.

### GetAggregationPropertyDataOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyDataOk() (*string, bool)`

GetAggregationPropertyDataOk returns a tuple with the AggregationPropertyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEvent) SetAggregationPropertyData(v string)`

SetAggregationPropertyData sets AggregationPropertyData field to given value.

### HasAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEvent) HasAggregationPropertyData() bool`

HasAggregationPropertyData returns a boolean if a field has been set.

### GetAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyInt() int64`

GetAggregationPropertyInt returns the AggregationPropertyInt field if non-nil, zero value otherwise.

### GetAggregationPropertyIntOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyIntOk() (*int64, bool)`

GetAggregationPropertyIntOk returns a tuple with the AggregationPropertyInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEvent) SetAggregationPropertyInt(v int64)`

SetAggregationPropertyInt sets AggregationPropertyInt field to given value.

### HasAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEvent) HasAggregationPropertyInt() bool`

HasAggregationPropertyInt returns a boolean if a field has been set.

### GetAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyString() string`

GetAggregationPropertyString returns the AggregationPropertyString field if non-nil, zero value otherwise.

### GetAggregationPropertyStringOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetAggregationPropertyStringOk() (*string, bool)`

GetAggregationPropertyStringOk returns a tuple with the AggregationPropertyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEvent) SetAggregationPropertyString(v string)`

SetAggregationPropertyString sets AggregationPropertyString field to given value.

### HasAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEvent) HasAggregationPropertyString() bool`

HasAggregationPropertyString returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMetricEvent) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMetricEvent) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMetricEvent) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEvent) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEvent) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.

### HasExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEvent) HasExternalEventId() bool`

HasExternalEventId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMetricEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMetricEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMetricEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMetricEvent) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMetricEvent) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanMerchantMetricEvent) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanMerchantMetricEvent) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEvent) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEvent) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionIds() string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionIdsOk() (*string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEvent) SetSubscriptionIds(v string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEvent) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionPeriodEnd() int64`

GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field if non-nil, zero value otherwise.

### GetSubscriptionPeriodEndOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionPeriodEndOk() (*int64, bool)`

GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEvent) SetSubscriptionPeriodEnd(v int64)`

SetSubscriptionPeriodEnd sets SubscriptionPeriodEnd field to given value.

### HasSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEvent) HasSubscriptionPeriodEnd() bool`

HasSubscriptionPeriodEnd returns a boolean if a field has been set.

### GetSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionPeriodStart() int64`

GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field if non-nil, zero value otherwise.

### GetSubscriptionPeriodStartOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetSubscriptionPeriodStartOk() (*int64, bool)`

GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEvent) SetSubscriptionPeriodStart(v int64)`

SetSubscriptionPeriodStart sets SubscriptionPeriodStart field to given value.

### HasSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEvent) HasSubscriptionPeriodStart() bool`

HasSubscriptionPeriodStart returns a boolean if a field has been set.

### GetUsed

`func (o *UnibeeApiBeanMerchantMetricEvent) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UnibeeApiBeanMerchantMetricEvent) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UnibeeApiBeanMerchantMetricEvent) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanMerchantMetricEvent) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanMerchantMetricEvent) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanMerchantMetricEvent) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanMerchantMetricEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


