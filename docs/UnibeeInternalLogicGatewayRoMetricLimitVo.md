# UnibeeInternalLogicGatewayRoMetricLimitVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **int32** |  | [optional] 
**MetricId** | Pointer to **int64** |  | [optional] 
**PlanLimits** | Pointer to [**[]UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo**](UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo.md) |  | [optional] 
**TotalLimit** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**AggregationProperty** | Pointer to **string** | aggregation property | [optional] 
**AggregationType** | Pointer to **int32** | 0-count，1-count unique, 2-latest, 3-max, 4-sum | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**MetricName** | Pointer to **string** | metric name | [optional] 
**Type** | Pointer to **int32** | 1-limit_metered，2-charge_metered(come later),3-charge_recurring(come later) | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoMetricLimitVo

`func NewUnibeeInternalLogicGatewayRoMetricLimitVo() *UnibeeInternalLogicGatewayRoMetricLimitVo`

NewUnibeeInternalLogicGatewayRoMetricLimitVo instantiates a new UnibeeInternalLogicGatewayRoMetricLimitVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoMetricLimitVoWithDefaults

`func NewUnibeeInternalLogicGatewayRoMetricLimitVoWithDefaults() *UnibeeInternalLogicGatewayRoMetricLimitVo`

NewUnibeeInternalLogicGatewayRoMetricLimitVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoMetricLimitVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMerchantId() int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMerchantIdOk() (*int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetMerchantId(v int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetPlanLimits

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetPlanLimits() []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo`

GetPlanLimits returns the PlanLimits field if non-nil, zero value otherwise.

### GetPlanLimitsOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetPlanLimitsOk() (*[]UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo, bool)`

GetPlanLimitsOk returns a tuple with the PlanLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLimits

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetPlanLimits(v []UnibeeInternalLogicGatewayRoMerchantMetricPlanLimitVo)`

SetPlanLimits sets PlanLimits field to given value.

### HasPlanLimits

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasPlanLimits() bool`

HasPlanLimits returns a boolean if a field has been set.

### GetTotalLimit

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetTotalLimit() int32`

GetTotalLimit returns the TotalLimit field if non-nil, zero value otherwise.

### GetTotalLimitOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetTotalLimitOk() (*int32, bool)`

GetTotalLimitOk returns a tuple with the TotalLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLimit

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetTotalLimit(v int32)`

SetTotalLimit sets TotalLimit field to given value.

### HasTotalLimit

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasTotalLimit() bool`

HasTotalLimit returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetAggregationProperty() string`

GetAggregationProperty returns the AggregationProperty field if non-nil, zero value otherwise.

### GetAggregationPropertyOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetAggregationPropertyOk() (*string, bool)`

GetAggregationPropertyOk returns a tuple with the AggregationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetAggregationProperty(v string)`

SetAggregationProperty sets AggregationProperty field to given value.

### HasAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasAggregationProperty() bool`

HasAggregationProperty returns a boolean if a field has been set.

### GetAggregationType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetAggregationType() int32`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetAggregationTypeOk() (*int32, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetAggregationType(v int32)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalLogicGatewayRoMetricLimitVo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


