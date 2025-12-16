# UnibeeApiMerchantPlanMetricLimitOverrideReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataOverride** | Pointer to **map[string]map[string]interface{}** | Metadata to override for plan, only provided keys will be added/updated | [optional] 
**MetricLimit** | Pointer to [**[]UnibeeApiBeanPlanMetricLimitOverrideParam**](UnibeeApiBeanPlanMetricLimitOverrideParam.md) | MetricLimit items to insert or override, identify metric by id or code | [optional] 
**PlanId** | **int64** | Id of plan | 

## Methods

### NewUnibeeApiMerchantPlanMetricLimitOverrideReq

`func NewUnibeeApiMerchantPlanMetricLimitOverrideReq(planId int64, ) *UnibeeApiMerchantPlanMetricLimitOverrideReq`

NewUnibeeApiMerchantPlanMetricLimitOverrideReq instantiates a new UnibeeApiMerchantPlanMetricLimitOverrideReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanMetricLimitOverrideReqWithDefaults

`func NewUnibeeApiMerchantPlanMetricLimitOverrideReqWithDefaults() *UnibeeApiMerchantPlanMetricLimitOverrideReq`

NewUnibeeApiMerchantPlanMetricLimitOverrideReqWithDefaults instantiates a new UnibeeApiMerchantPlanMetricLimitOverrideReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataOverride

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetMetadataOverride() map[string]map[string]interface{}`

GetMetadataOverride returns the MetadataOverride field if non-nil, zero value otherwise.

### GetMetadataOverrideOk

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetMetadataOverrideOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOverrideOk returns a tuple with the MetadataOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataOverride

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) SetMetadataOverride(v map[string]map[string]interface{})`

SetMetadataOverride sets MetadataOverride field to given value.

### HasMetadataOverride

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) HasMetadataOverride() bool`

HasMetadataOverride returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetMetricLimit() []UnibeeApiBeanPlanMetricLimitOverrideParam`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetMetricLimitOk() (*[]UnibeeApiBeanPlanMetricLimitOverrideParam, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) SetMetricLimit(v []UnibeeApiBeanPlanMetricLimitOverrideParam)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanMetricLimitOverrideReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


