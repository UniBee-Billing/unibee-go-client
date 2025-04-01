# UnibeeApiBeanPlanMetricMeteredChargeParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **int32** | ChargeType,0-standard pricing 1-graduated pricing | [optional] 
**GraduatedAmounts** | Pointer to [**[]UnibeeApiBeanMetricPlanChargeGraduatedStep**](UnibeeApiBeanMetricPlanChargeGraduatedStep.md) | GraduatedAmounts, used for graduated pricing | [optional] 
**MetricId** | Pointer to **int64** | MetricId | [optional] 
**StandardAmount** | Pointer to **int64** | StandardAmount, used for standard pricing,cent | [optional] 
**StandardStartValue** | Pointer to **int64** | StandardStartValue, used for standard pricing | [optional] 

## Methods

### NewUnibeeApiBeanPlanMetricMeteredChargeParam

`func NewUnibeeApiBeanPlanMetricMeteredChargeParam() *UnibeeApiBeanPlanMetricMeteredChargeParam`

NewUnibeeApiBeanPlanMetricMeteredChargeParam instantiates a new UnibeeApiBeanPlanMetricMeteredChargeParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanMetricMeteredChargeParamWithDefaults

`func NewUnibeeApiBeanPlanMetricMeteredChargeParamWithDefaults() *UnibeeApiBeanPlanMetricMeteredChargeParam`

NewUnibeeApiBeanPlanMetricMeteredChargeParamWithDefaults instantiates a new UnibeeApiBeanPlanMetricMeteredChargeParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetChargeType() int32`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetChargeTypeOk() (*int32, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) SetChargeType(v int32)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetGraduatedAmounts

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetGraduatedAmounts() []UnibeeApiBeanMetricPlanChargeGraduatedStep`

GetGraduatedAmounts returns the GraduatedAmounts field if non-nil, zero value otherwise.

### GetGraduatedAmountsOk

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetGraduatedAmountsOk() (*[]UnibeeApiBeanMetricPlanChargeGraduatedStep, bool)`

GetGraduatedAmountsOk returns a tuple with the GraduatedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedAmounts

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) SetGraduatedAmounts(v []UnibeeApiBeanMetricPlanChargeGraduatedStep)`

SetGraduatedAmounts sets GraduatedAmounts field to given value.

### HasGraduatedAmounts

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) HasGraduatedAmounts() bool`

HasGraduatedAmounts returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetStandardAmount

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetStandardAmount() int64`

GetStandardAmount returns the StandardAmount field if non-nil, zero value otherwise.

### GetStandardAmountOk

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetStandardAmountOk() (*int64, bool)`

GetStandardAmountOk returns a tuple with the StandardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardAmount

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) SetStandardAmount(v int64)`

SetStandardAmount sets StandardAmount field to given value.

### HasStandardAmount

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) HasStandardAmount() bool`

HasStandardAmount returns a boolean if a field has been set.

### GetStandardStartValue

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetStandardStartValue() int64`

GetStandardStartValue returns the StandardStartValue field if non-nil, zero value otherwise.

### GetStandardStartValueOk

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) GetStandardStartValueOk() (*int64, bool)`

GetStandardStartValueOk returns a tuple with the StandardStartValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardStartValue

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) SetStandardStartValue(v int64)`

SetStandardStartValue sets StandardStartValue field to given value.

### HasStandardStartValue

`func (o *UnibeeApiBeanPlanMetricMeteredChargeParam) HasStandardStartValue() bool`

HasStandardStartValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


