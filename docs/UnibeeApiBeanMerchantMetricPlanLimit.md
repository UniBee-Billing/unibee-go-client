# UnibeeApiBeanMerchantMetricPlanLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MerchantMetricVo** | Pointer to [**UnibeeApiBeanMerchantMetric**](UnibeeApiBeanMerchantMetric.md) |  | [optional] 
**MetricId** | Pointer to **int64** | metricId | [optional] 
**MetricLimit** | Pointer to **int64** | plan metric limit | [optional] 
**PlanId** | Pointer to **int64** | plan_id | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMetricPlanLimit

`func NewUnibeeApiBeanMerchantMetricPlanLimit() *UnibeeApiBeanMerchantMetricPlanLimit`

NewUnibeeApiBeanMerchantMetricPlanLimit instantiates a new UnibeeApiBeanMerchantMetricPlanLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMetricPlanLimitWithDefaults

`func NewUnibeeApiBeanMerchantMetricPlanLimitWithDefaults() *UnibeeApiBeanMerchantMetricPlanLimit`

NewUnibeeApiBeanMerchantMetricPlanLimitWithDefaults instantiates a new UnibeeApiBeanMerchantMetricPlanLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMetricVo

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantMetricVo() UnibeeApiBeanMerchantMetric`

GetMerchantMetricVo returns the MerchantMetricVo field if non-nil, zero value otherwise.

### GetMerchantMetricVoOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantMetricVoOk() (*UnibeeApiBeanMerchantMetric, bool)`

GetMerchantMetricVoOk returns a tuple with the MerchantMetricVo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMetricVo

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMerchantMetricVo(v UnibeeApiBeanMerchantMetric)`

SetMerchantMetricVo sets MerchantMetricVo field to given value.

### HasMerchantMetricVo

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMerchantMetricVo() bool`

HasMerchantMetricVo returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricLimit() int64`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricLimitOk() (*int64, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMetricLimit(v int64)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


