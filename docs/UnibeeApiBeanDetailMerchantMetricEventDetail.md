# UnibeeApiBeanDetailMerchantMetricEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationPropertyData** | Pointer to **string** | aggregation property data (Json) | [optional] 
**AggregationPropertyInt** | Pointer to **int64** | aggregation property int, use for metric of max|sum type | [optional] 
**AggregationPropertyString** | Pointer to **string** | aggregation property string, use for metric of count|count_unique type | [optional] 
**ChargeInvoiceId** | Pointer to **string** | charge invoice id | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**EventCharge** | Pointer to [**UnibeeApiBeanEventMetricCharge**](UnibeeApiBeanEventMetricCharge.md) |  | [optional] 
**ExternalEventId** | Pointer to **string** | external_event_id, should be unique | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MerchantMetric** | Pointer to [**UnibeeApiBeanMerchantMetric**](UnibeeApiBeanMerchantMetric.md) |  | [optional] 
**MetricId** | Pointer to **int64** | metric_id | [optional] 
**MetricLimit** | Pointer to **int64** |  | [optional] 
**SubscriptionIds** | Pointer to **string** |  | [optional] 
**SubscriptionPeriodEnd** | Pointer to **int64** | matched subscription&#39;s current_period_end | [optional] 
**SubscriptionPeriodStart** | Pointer to **int64** | matched subscription&#39;s current_period_start | [optional] 
**Used** | Pointer to **int64** |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantMetricEventDetail

`func NewUnibeeApiBeanDetailMerchantMetricEventDetail() *UnibeeApiBeanDetailMerchantMetricEventDetail`

NewUnibeeApiBeanDetailMerchantMetricEventDetail instantiates a new UnibeeApiBeanDetailMerchantMetricEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantMetricEventDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantMetricEventDetailWithDefaults() *UnibeeApiBeanDetailMerchantMetricEventDetail`

NewUnibeeApiBeanDetailMerchantMetricEventDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantMetricEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationPropertyData

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyData() string`

GetAggregationPropertyData returns the AggregationPropertyData field if non-nil, zero value otherwise.

### GetAggregationPropertyDataOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyDataOk() (*string, bool)`

GetAggregationPropertyDataOk returns a tuple with the AggregationPropertyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyData

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetAggregationPropertyData(v string)`

SetAggregationPropertyData sets AggregationPropertyData field to given value.

### HasAggregationPropertyData

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasAggregationPropertyData() bool`

HasAggregationPropertyData returns a boolean if a field has been set.

### GetAggregationPropertyInt

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyInt() int64`

GetAggregationPropertyInt returns the AggregationPropertyInt field if non-nil, zero value otherwise.

### GetAggregationPropertyIntOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyIntOk() (*int64, bool)`

GetAggregationPropertyIntOk returns a tuple with the AggregationPropertyInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyInt

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetAggregationPropertyInt(v int64)`

SetAggregationPropertyInt sets AggregationPropertyInt field to given value.

### HasAggregationPropertyInt

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasAggregationPropertyInt() bool`

HasAggregationPropertyInt returns a boolean if a field has been set.

### GetAggregationPropertyString

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyString() string`

GetAggregationPropertyString returns the AggregationPropertyString field if non-nil, zero value otherwise.

### GetAggregationPropertyStringOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetAggregationPropertyStringOk() (*string, bool)`

GetAggregationPropertyStringOk returns a tuple with the AggregationPropertyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationPropertyString

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetAggregationPropertyString(v string)`

SetAggregationPropertyString sets AggregationPropertyString field to given value.

### HasAggregationPropertyString

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasAggregationPropertyString() bool`

HasAggregationPropertyString returns a boolean if a field has been set.

### GetChargeInvoiceId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetChargeInvoiceId() string`

GetChargeInvoiceId returns the ChargeInvoiceId field if non-nil, zero value otherwise.

### GetChargeInvoiceIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetChargeInvoiceIdOk() (*string, bool)`

GetChargeInvoiceIdOk returns a tuple with the ChargeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeInvoiceId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetChargeInvoiceId(v string)`

SetChargeInvoiceId sets ChargeInvoiceId field to given value.

### HasChargeInvoiceId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasChargeInvoiceId() bool`

HasChargeInvoiceId returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEventCharge

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetEventCharge() UnibeeApiBeanEventMetricCharge`

GetEventCharge returns the EventCharge field if non-nil, zero value otherwise.

### GetEventChargeOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetEventChargeOk() (*UnibeeApiBeanEventMetricCharge, bool)`

GetEventChargeOk returns a tuple with the EventCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCharge

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetEventCharge(v UnibeeApiBeanEventMetricCharge)`

SetEventCharge sets EventCharge field to given value.

### HasEventCharge

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasEventCharge() bool`

HasEventCharge returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.

### HasExternalEventId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasExternalEventId() bool`

HasExternalEventId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMerchantMetric() UnibeeApiBeanMerchantMetric`

GetMerchantMetric returns the MerchantMetric field if non-nil, zero value otherwise.

### GetMerchantMetricOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool)`

GetMerchantMetricOk returns a tuple with the MerchantMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetMerchantMetric(v UnibeeApiBeanMerchantMetric)`

SetMerchantMetric sets MerchantMetric field to given value.

### HasMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasMerchantMetric() bool`

HasMerchantMetric returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetSubscriptionIds

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionIds() string`

GetSubscriptionIds returns the SubscriptionIds field if non-nil, zero value otherwise.

### GetSubscriptionIdsOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionIdsOk() (*string, bool)`

GetSubscriptionIdsOk returns a tuple with the SubscriptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionIds

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetSubscriptionIds(v string)`

SetSubscriptionIds sets SubscriptionIds field to given value.

### HasSubscriptionIds

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasSubscriptionIds() bool`

HasSubscriptionIds returns a boolean if a field has been set.

### GetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionPeriodEnd() int64`

GetSubscriptionPeriodEnd returns the SubscriptionPeriodEnd field if non-nil, zero value otherwise.

### GetSubscriptionPeriodEndOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionPeriodEndOk() (*int64, bool)`

GetSubscriptionPeriodEndOk returns a tuple with the SubscriptionPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodEnd

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetSubscriptionPeriodEnd(v int64)`

SetSubscriptionPeriodEnd sets SubscriptionPeriodEnd field to given value.

### HasSubscriptionPeriodEnd

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasSubscriptionPeriodEnd() bool`

HasSubscriptionPeriodEnd returns a boolean if a field has been set.

### GetSubscriptionPeriodStart

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionPeriodStart() int64`

GetSubscriptionPeriodStart returns the SubscriptionPeriodStart field if non-nil, zero value otherwise.

### GetSubscriptionPeriodStartOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetSubscriptionPeriodStartOk() (*int64, bool)`

GetSubscriptionPeriodStartOk returns a tuple with the SubscriptionPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodStart

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetSubscriptionPeriodStart(v int64)`

SetSubscriptionPeriodStart sets SubscriptionPeriodStart field to given value.

### HasSubscriptionPeriodStart

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasSubscriptionPeriodStart() bool`

HasSubscriptionPeriodStart returns a boolean if a field has been set.

### GetUsed

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUsed() int64`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUsedOk() (*int64, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetUsed(v int64)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailMerchantMetricEventDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


