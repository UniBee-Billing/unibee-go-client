# UnibeeApiBeanDetailMerchantMetricPlanLimitDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MerchantMetric** | Pointer to [**UnibeeApiBeanMerchantMetric**](UnibeeApiBeanMerchantMetric.md) |  | [optional] 
**MetricId** | Pointer to **int64** | metricId | [optional] 
**MetricLimit** | Pointer to **int64** | plan metric limit | [optional] 
**PlanId** | Pointer to **int64** | plan_id | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetail

`func NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetail() *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail`

NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetail instantiates a new UnibeeApiBeanDetailMerchantMetricPlanLimitDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetailWithDefaults() *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail`

NewUnibeeApiBeanDetailMerchantMetricPlanLimitDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantMetricPlanLimitDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMerchantMetric() UnibeeApiBeanMerchantMetric`

GetMerchantMetric returns the MerchantMetric field if non-nil, zero value otherwise.

### GetMerchantMetricOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool)`

GetMerchantMetricOk returns a tuple with the MerchantMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetMerchantMetric(v UnibeeApiBeanMerchantMetric)`

SetMerchantMetric sets MerchantMetric field to given value.

### HasMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasMerchantMetric() bool`

HasMerchantMetric returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanLimitDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


