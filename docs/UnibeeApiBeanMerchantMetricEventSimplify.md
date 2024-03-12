# UnibeeApiBeanMerchantMetricEventSimplify

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

### NewUnibeeApiBeanMerchantMetricEventSimplify

`func NewUnibeeApiBeanMerchantMetricEventSimplify() *UnibeeApiBeanMerchantMetricEventSimplify`

NewUnibeeApiBeanMerchantMetricEventSimplify instantiates a new UnibeeApiBeanMerchantMetricEventSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMetricEventSimplifyWithDefaults

`func NewUnibeeApiBeanMerchantMetricEventSimplifyWithDefaults() *UnibeeApiBeanMerchantMetricEventSimplify`

NewUnibeeApiBeanMerchantMetricEventSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantMetricEventSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyData() string`

GetAggregationPropertyData returns the AggregationPropertyData field if non-nil, zero value otherwise.

### GetAggregationPropertyDataOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyDataOk() (*string, bool)`

GetAggregationPropertyDataOk returns a tuple with the AggregationPropertyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyData(v string)`

SetAggregationPropertyData sets AggregationPropertyData field to given value.

### HasAggregationPropertyData

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyData() bool`

HasAggregationPropertyData returns a boolean if a field has been set.

### GetAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyInt() int64`

GetAggregationPropertyInt returns the AggregationPropertyInt field if non-nil, zero value otherwise.

### GetAggregationPropertyIntOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyIntOk() (*int64, bool)`

GetAggregationPropertyIntOk returns a tuple with the AggregationPropertyInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyInt(v int64)`

SetAggregationPropertyInt sets AggregationPropertyInt field to given value.

### HasAggregationPropertyInt

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyInt() bool`

HasAggregationPropertyInt returns a boolean if a field has been set.

### GetAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyString() string`

GetAggregationPropertyString returns the AggregationPropertyString field if non-nil, zero value otherwise.

### GetAggregationPropertyStringOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetAggregationPropertyStringOk() (*string, bool)`

GetAggregationPropertyStringOk returns a tuple with the AggregationPropertyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetAggregationPropertyString(v string)`

SetAggregationPropertyString sets AggregationPropertyString field to given value.

### HasAggregationPropertyString

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasAggregationPropertyString() bool`

HasAggregationPropertyString returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.

### HasExternalEventId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasExternalEventId() bool`

HasExternalEventId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionIds() string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionIdsOk() (*string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionIds(v string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodEnd() int64`

GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field if non-nil, zero value otherwise.

### GetSubscriptionPeriodEndOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodEndOk() (*int64, bool)`

GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionPeriodEnd(v int64)`

SetSubscriptionPeriodEnd sets SubscriptionPeriodEnd field to given value.

### HasSubscriptionPeriodEnd

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionPeriodEnd() bool`

HasSubscriptionPeriodEnd returns a boolean if a field has been set.

### GetSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodStart() int64`

GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field if non-nil, zero value otherwise.

### GetSubscriptionPeriodStartOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetSubscriptionPeriodStartOk() (*int64, bool)`

GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetSubscriptionPeriodStart(v int64)`

SetSubscriptionPeriodStart sets SubscriptionPeriodStart field to given value.

### HasSubscriptionPeriodStart

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasSubscriptionPeriodStart() bool`

HasSubscriptionPeriodStart returns a boolean if a field has been set.

### GetUsed

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanMerchantMetricEventSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


