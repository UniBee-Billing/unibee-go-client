# UnibeeInternalModelEntityOverseaPayMerchantMetricEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationPropertyData** | Pointer to **string** | aggregation property data (Json) | [optional] 
**AggregationPropertyInt** | Pointer to **int64** | aggregation property int, use for metric of max|sum type | [optional] 
**AggregationPropertyString** | Pointer to **string** | aggregation property string, use for metric of count|count_unique type | [optional] 
**AggregationPropertyUniqueId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**ExternalEventId** | Pointer to **string** | external_event_id, should be unique | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MetricId** | Pointer to **int64** | metric_id | [optional] 
**MetricLimit** | Pointer to **int64** |  | [optional] 
**SubscriptionIds** | Pointer to **string** |  | [optional] 
**SubscriptionPeriodEnd** | Pointer to **int64** | matched subscription&#39;s current_period_end | [optional] 
**SubscriptionPeriodStart** | Pointer to **int64** | matched subscription&#39;s current_period_start | [optional] 
**Used** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayMerchantMetricEvent

`func NewUnibeeInternalModelEntityOverseaPayMerchantMetricEvent() *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent`

NewUnibeeInternalModelEntityOverseaPayMerchantMetricEvent instantiates a new UnibeeInternalModelEntityOverseaPayMerchantMetricEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayMerchantMetricEventWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayMerchantMetricEventWithDefaults() *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent`

NewUnibeeInternalModelEntityOverseaPayMerchantMetricEventWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayMerchantMetricEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationPropertyData

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyData() string`

GetAggregationPropertyData returns the AggregationPropertyData field if non-nil, zero value otherwise.

### GetAggregationPropertyDataOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyDataOk() (*string, bool)`

GetAggregationPropertyDataOk returns a tuple with the AggregationPropertyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyData

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetAggregationPropertyData(v string)`

SetAggregationPropertyData sets AggregationPropertyData field to given value.

### HasAggregationPropertyData

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasAggregationPropertyData() bool`

HasAggregationPropertyData returns a boolean if a field has been set.

### GetAggregationPropertyInt

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyInt() int64`

GetAggregationPropertyInt returns the AggregationPropertyInt field if non-nil, zero value otherwise.

### GetAggregationPropertyIntOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyIntOk() (*int64, bool)`

GetAggregationPropertyIntOk returns a tuple with the AggregationPropertyInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyInt

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetAggregationPropertyInt(v int64)`

SetAggregationPropertyInt sets AggregationPropertyInt field to given value.

### HasAggregationPropertyInt

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasAggregationPropertyInt() bool`

HasAggregationPropertyInt returns a boolean if a field has been set.

### GetAggregationPropertyString

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyString() string`

GetAggregationPropertyString returns the AggregationPropertyString field if non-nil, zero value otherwise.

### GetAggregationPropertyStringOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyStringOk() (*string, bool)`

GetAggregationPropertyStringOk returns a tuple with the AggregationPropertyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyString

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetAggregationPropertyString(v string)`

SetAggregationPropertyString sets AggregationPropertyString field to given value.

### HasAggregationPropertyString

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasAggregationPropertyString() bool`

HasAggregationPropertyString returns a boolean if a field has been set.

### GetAggregationPropertyUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyUniqueId() string`

GetAggregationPropertyUniqueId returns the AggregationPropertyUniqueId field if non-nil, zero value otherwise.

### GetAggregationPropertyUniqueIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetAggregationPropertyUniqueIdOk() (*string, bool)`

GetAggregationPropertyUniqueIdOk returns a tuple with the AggregationPropertyUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetAggregationPropertyUniqueId(v string)`

SetAggregationPropertyUniqueId sets AggregationPropertyUniqueId field to given value.

### HasAggregationPropertyUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasAggregationPropertyUniqueId() bool`

HasAggregationPropertyUniqueId returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.

### HasExternalEventId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasExternalEventId() bool`

HasExternalEventId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionIds() string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionIdsOk() (*string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetSubscriptionIds(v string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptionPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionPeriodEnd() int64`

GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field if non-nil, zero value otherwise.

### GetSubscriptionPeriodEndOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionPeriodEndOk() (*int64, bool)`

GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetSubscriptionPeriodEnd(v int64)`

SetSubscriptionPeriodEnd sets SubscriptionPeriodEnd field to given value.

### HasSubscriptionPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasSubscriptionPeriodEnd() bool`

HasSubscriptionPeriodEnd returns a boolean if a field has been set.

### GetSubscriptionPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionPeriodStart() int64`

GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field if non-nil, zero value otherwise.

### GetSubscriptionPeriodStartOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetSubscriptionPeriodStartOk() (*int64, bool)`

GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetSubscriptionPeriodStart(v int64)`

SetSubscriptionPeriodStart sets SubscriptionPeriodStart field to given value.

### HasSubscriptionPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasSubscriptionPeriodStart() bool`

HasSubscriptionPeriodStart returns a boolean if a field has been set.

### GetUsed

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMetricEvent) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


