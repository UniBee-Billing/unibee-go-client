# UnibeeApiBeanDetailMerchantMetricPlanChargeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeType** | Pointer to **int32** | ChargeType,0-standard pricing 1-graduated pricing | [optional] 
**GraduatedAmounts** | Pointer to [**[]UnibeeApiBeanMetricPlanChargeGraduatedStep**](UnibeeApiBeanMetricPlanChargeGraduatedStep.md) | GraduatedAmounts | [optional] 
**MerchantMetric** | Pointer to [**UnibeeApiBeanMerchantMetric**](UnibeeApiBeanMerchantMetric.md) |  | [optional] 
**MetricId** | **int64** | MetricId | 
**StandardAmount** | Pointer to **int64** | StandardAmount, used for standard pricing,cent | [optional] 
**StandardStartValue** | Pointer to **int64** | StandardStartValue, used for standard pricing | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetail

`func NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetail(metricId int64, ) *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail`

NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetail instantiates a new UnibeeApiBeanDetailMerchantMetricPlanChargeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetailWithDefaults() *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail`

NewUnibeeApiBeanDetailMerchantMetricPlanChargeDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantMetricPlanChargeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeType

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetChargeType() int32`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetChargeTypeOk() (*int32, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetChargeType(v int32)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetGraduatedAmounts

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetGraduatedAmounts() []UnibeeApiBeanMetricPlanChargeGraduatedStep`

GetGraduatedAmounts returns the GraduatedAmounts field if non-nil, zero value otherwise.

### GetGraduatedAmountsOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetGraduatedAmountsOk() (*[]UnibeeApiBeanMetricPlanChargeGraduatedStep, bool)`

GetGraduatedAmountsOk returns a tuple with the GraduatedAmounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedAmounts

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetGraduatedAmounts(v []UnibeeApiBeanMetricPlanChargeGraduatedStep)`

SetGraduatedAmounts sets GraduatedAmounts field to given value.

### HasGraduatedAmounts

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) HasGraduatedAmounts() bool`

HasGraduatedAmounts returns a boolean if a field has been set.

### GetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetMerchantMetric() UnibeeApiBeanMerchantMetric`

GetMerchantMetric returns the MerchantMetric field if non-nil, zero value otherwise.

### GetMerchantMetricOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool)`

GetMerchantMetricOk returns a tuple with the MerchantMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetMerchantMetric(v UnibeeApiBeanMerchantMetric)`

SetMerchantMetric sets MerchantMetric field to given value.

### HasMerchantMetric

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) HasMerchantMetric() bool`

HasMerchantMetric returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.


### GetStandardAmount

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetStandardAmount() int64`

GetStandardAmount returns the StandardAmount field if non-nil, zero value otherwise.

### GetStandardAmountOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetStandardAmountOk() (*int64, bool)`

GetStandardAmountOk returns a tuple with the StandardAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardAmount

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetStandardAmount(v int64)`

SetStandardAmount sets StandardAmount field to given value.

### HasStandardAmount

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) HasStandardAmount() bool`

HasStandardAmount returns a boolean if a field has been set.

### GetStandardStartValue

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetStandardStartValue() int64`

GetStandardStartValue returns the StandardStartValue field if non-nil, zero value otherwise.

### GetStandardStartValueOk

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) GetStandardStartValueOk() (*int64, bool)`

GetStandardStartValueOk returns a tuple with the StandardStartValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardStartValue

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) SetStandardStartValue(v int64)`

SetStandardStartValue sets StandardStartValue field to given value.

### HasStandardStartValue

`func (o *UnibeeApiBeanDetailMerchantMetricPlanChargeDetail) HasStandardStartValue() bool`

HasStandardStartValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


