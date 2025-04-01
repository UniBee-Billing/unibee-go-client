# UnibeeApiBeanEventMetricCharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeAmount** | Pointer to **int64** | ChargeAmount | [optional] 
**ChargePricing** | Pointer to [**UnibeeApiBeanPlanMetricMeteredChargeParam**](UnibeeApiBeanPlanMetricMeteredChargeParam.md) |  | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**CurrentValue** | Pointer to **int64** | CurrentUsedValue | [optional] 
**GraduatedStep** | Pointer to [**UnibeeApiBeanMetricPlanChargeGraduatedStep**](UnibeeApiBeanMetricPlanChargeGraduatedStep.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**TotalChargeAmount** | Pointer to **int64** | TotalChargeAmount | [optional] 
**UnitAmount** | Pointer to **int64** | UnitAmount | [optional] 

## Methods

### NewUnibeeApiBeanEventMetricCharge

`func NewUnibeeApiBeanEventMetricCharge() *UnibeeApiBeanEventMetricCharge`

NewUnibeeApiBeanEventMetricCharge instantiates a new UnibeeApiBeanEventMetricCharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanEventMetricChargeWithDefaults

`func NewUnibeeApiBeanEventMetricChargeWithDefaults() *UnibeeApiBeanEventMetricCharge`

NewUnibeeApiBeanEventMetricChargeWithDefaults instantiates a new UnibeeApiBeanEventMetricCharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) GetChargeAmount() int64`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *UnibeeApiBeanEventMetricCharge) GetChargeAmountOk() (*int64, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) SetChargeAmount(v int64)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetChargePricing

`func (o *UnibeeApiBeanEventMetricCharge) GetChargePricing() UnibeeApiBeanPlanMetricMeteredChargeParam`

GetChargePricing returns the ChargePricing field if non-nil, zero value otherwise.

### GetChargePricingOk

`func (o *UnibeeApiBeanEventMetricCharge) GetChargePricingOk() (*UnibeeApiBeanPlanMetricMeteredChargeParam, bool)`

GetChargePricingOk returns a tuple with the ChargePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargePricing

`func (o *UnibeeApiBeanEventMetricCharge) SetChargePricing(v UnibeeApiBeanPlanMetricMeteredChargeParam)`

SetChargePricing sets ChargePricing field to given value.

### HasChargePricing

`func (o *UnibeeApiBeanEventMetricCharge) HasChargePricing() bool`

HasChargePricing returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanEventMetricCharge) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanEventMetricCharge) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanEventMetricCharge) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanEventMetricCharge) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentValue

`func (o *UnibeeApiBeanEventMetricCharge) GetCurrentValue() int64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *UnibeeApiBeanEventMetricCharge) GetCurrentValueOk() (*int64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *UnibeeApiBeanEventMetricCharge) SetCurrentValue(v int64)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *UnibeeApiBeanEventMetricCharge) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetGraduatedStep

`func (o *UnibeeApiBeanEventMetricCharge) GetGraduatedStep() UnibeeApiBeanMetricPlanChargeGraduatedStep`

GetGraduatedStep returns the GraduatedStep field if non-nil, zero value otherwise.

### GetGraduatedStepOk

`func (o *UnibeeApiBeanEventMetricCharge) GetGraduatedStepOk() (*UnibeeApiBeanMetricPlanChargeGraduatedStep, bool)`

GetGraduatedStepOk returns a tuple with the GraduatedStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraduatedStep

`func (o *UnibeeApiBeanEventMetricCharge) SetGraduatedStep(v UnibeeApiBeanMetricPlanChargeGraduatedStep)`

SetGraduatedStep sets GraduatedStep field to given value.

### HasGraduatedStep

`func (o *UnibeeApiBeanEventMetricCharge) HasGraduatedStep() bool`

HasGraduatedStep returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanEventMetricCharge) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanEventMetricCharge) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanEventMetricCharge) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanEventMetricCharge) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTotalChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) GetTotalChargeAmount() int64`

GetTotalChargeAmount returns the TotalChargeAmount field if non-nil, zero value otherwise.

### GetTotalChargeAmountOk

`func (o *UnibeeApiBeanEventMetricCharge) GetTotalChargeAmountOk() (*int64, bool)`

GetTotalChargeAmountOk returns a tuple with the TotalChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) SetTotalChargeAmount(v int64)`

SetTotalChargeAmount sets TotalChargeAmount field to given value.

### HasTotalChargeAmount

`func (o *UnibeeApiBeanEventMetricCharge) HasTotalChargeAmount() bool`

HasTotalChargeAmount returns a boolean if a field has been set.

### GetUnitAmount

`func (o *UnibeeApiBeanEventMetricCharge) GetUnitAmount() int64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *UnibeeApiBeanEventMetricCharge) GetUnitAmountOk() (*int64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *UnibeeApiBeanEventMetricCharge) SetUnitAmount(v int64)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *UnibeeApiBeanEventMetricCharge) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


