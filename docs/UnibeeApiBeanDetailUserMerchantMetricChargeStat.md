# UnibeeApiBeanDetailUserMerchantMetricChargeStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUsedValue** | Pointer to **int64** | CurrentUsedValue | [optional] 
**ChargePricing** | Pointer to [**UnibeeApiBeanPlanMetricMeteredChargeParam**](UnibeeApiBeanPlanMetricMeteredChargeParam.md) |  | [optional] 
**GraduatedStep** | Pointer to [**UnibeeApiBeanMetricPlanChargeGraduatedStep**](UnibeeApiBeanMetricPlanChargeGraduatedStep.md) |  | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanMetricPlanChargeLine**](UnibeeApiBeanMetricPlanChargeLine.md) | Lines | [optional] 
**MaxEventId** | Pointer to **int64** |  | [optional] 
**MerchantMetric** | Pointer to [**UnibeeApiBeanMerchantMetric**](UnibeeApiBeanMerchantMetric.md) |  | [optional] 
**MetricId** | **int64** | MetricId | 
**MinEventId** | Pointer to **int64** |  | [optional] 
**TotalChargeAmount** | Pointer to **int64** | TotalChargeAmount | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserMerchantMetricChargeStat

`func NewUnibeeApiBeanDetailUserMerchantMetricChargeStat(metricId int64, ) *UnibeeApiBeanDetailUserMerchantMetricChargeStat`

NewUnibeeApiBeanDetailUserMerchantMetricChargeStat instantiates a new UnibeeApiBeanDetailUserMerchantMetricChargeStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserMerchantMetricChargeStatWithDefaults

`func NewUnibeeApiBeanDetailUserMerchantMetricChargeStatWithDefaults() *UnibeeApiBeanDetailUserMerchantMetricChargeStat`

NewUnibeeApiBeanDetailUserMerchantMetricChargeStatWithDefaults instantiates a new UnibeeApiBeanDetailUserMerchantMetricChargeStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUsedValue

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetCurrentUsedValue() int64`

GetCurrentUsedValue returns the CurrentUsedValue field if non-nil, zero value otherwise.

### GetCurrentUsedValueOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetCurrentUsedValueOk() (*int64, bool)`

GetCurrentUsedValueOk returns a tuple with the CurrentUsedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUsedValue

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetCurrentUsedValue(v int64)`

SetCurrentUsedValue sets CurrentUsedValue field to given value.

### HasCurrentUsedValue

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasCurrentUsedValue() bool`

HasCurrentUsedValue returns a boolean if a field has been set.

### GetChargePricing

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetChargePricing() UnibeeApiBeanPlanMetricMeteredChargeParam`

GetChargePricing returns the ChargePricing field if non-nil, zero value otherwise.

### GetChargePricingOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetChargePricingOk() (*UnibeeApiBeanPlanMetricMeteredChargeParam, bool)`

GetChargePricingOk returns a tuple with the ChargePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargePricing

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetChargePricing(v UnibeeApiBeanPlanMetricMeteredChargeParam)`

SetChargePricing sets ChargePricing field to given value.

### HasChargePricing

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasChargePricing() bool`

HasChargePricing returns a boolean if a field has been set.

### GetGraduatedStep

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetGraduatedStep() UnibeeApiBeanMetricPlanChargeGraduatedStep`

GetGraduatedStep returns the GraduatedStep field if non-nil, zero value otherwise.

### GetGraduatedStepOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetGraduatedStepOk() (*UnibeeApiBeanMetricPlanChargeGraduatedStep, bool)`

GetGraduatedStepOk returns a tuple with the GraduatedStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedStep

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetGraduatedStep(v UnibeeApiBeanMetricPlanChargeGraduatedStep)`

SetGraduatedStep sets GraduatedStep field to given value.

### HasGraduatedStep

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasGraduatedStep() bool`

HasGraduatedStep returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetLines() []UnibeeApiBeanMetricPlanChargeLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetLinesOk() (*[]UnibeeApiBeanMetricPlanChargeLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetLines(v []UnibeeApiBeanMetricPlanChargeLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMaxEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMaxEventId() int64`

GetMaxEventId returns the MaxEventId field if non-nil, zero value otherwise.

### GetMaxEventIdOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMaxEventIdOk() (*int64, bool)`

GetMaxEventIdOk returns a tuple with the MaxEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetMaxEventId(v int64)`

SetMaxEventId sets MaxEventId field to given value.

### HasMaxEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasMaxEventId() bool`

HasMaxEventId returns a boolean if a field has been set.

### GetMerchantMetric

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMerchantMetric() UnibeeApiBeanMerchantMetric`

GetMerchantMetric returns the MerchantMetric field if non-nil, zero value otherwise.

### GetMerchantMetricOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMerchantMetricOk() (*UnibeeApiBeanMerchantMetric, bool)`

GetMerchantMetricOk returns a tuple with the MerchantMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMetric

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetMerchantMetric(v UnibeeApiBeanMerchantMetric)`

SetMerchantMetric sets MerchantMetric field to given value.

### HasMerchantMetric

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasMerchantMetric() bool`

HasMerchantMetric returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.


### GetMinEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMinEventId() int64`

GetMinEventId returns the MinEventId field if non-nil, zero value otherwise.

### GetMinEventIdOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetMinEventIdOk() (*int64, bool)`

GetMinEventIdOk returns a tuple with the MinEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetMinEventId(v int64)`

SetMinEventId sets MinEventId field to given value.

### HasMinEventId

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasMinEventId() bool`

HasMinEventId returns a boolean if a field has been set.

### GetTotalChargeAmount

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetTotalChargeAmount() int64`

GetTotalChargeAmount returns the TotalChargeAmount field if non-nil, zero value otherwise.

### GetTotalChargeAmountOk

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) GetTotalChargeAmountOk() (*int64, bool)`

GetTotalChargeAmountOk returns a tuple with the TotalChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChargeAmount

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) SetTotalChargeAmount(v int64)`

SetTotalChargeAmount sets TotalChargeAmount field to given value.

### HasTotalChargeAmount

`func (o *UnibeeApiBeanDetailUserMerchantMetricChargeStat) HasTotalChargeAmount() bool`

HasTotalChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


