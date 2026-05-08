# UnibeeApiBeanPlanMultiCurrencySnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeCurrency** | Pointer to **string** | Invoice charge currency used for this snapshot | [optional] 
**Rates** | Pointer to [**[]UnibeeApiBeanPlanMultiCurrencyRate**](UnibeeApiBeanPlanMultiCurrencyRate.md) | FX rates from various default currencies into chargeCurrency | [optional] 

## Methods

### NewUnibeeApiBeanPlanMultiCurrencySnapshot

`func NewUnibeeApiBeanPlanMultiCurrencySnapshot() *UnibeeApiBeanPlanMultiCurrencySnapshot`

NewUnibeeApiBeanPlanMultiCurrencySnapshot instantiates a new UnibeeApiBeanPlanMultiCurrencySnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanMultiCurrencySnapshotWithDefaults

`func NewUnibeeApiBeanPlanMultiCurrencySnapshotWithDefaults() *UnibeeApiBeanPlanMultiCurrencySnapshot`

NewUnibeeApiBeanPlanMultiCurrencySnapshotWithDefaults instantiates a new UnibeeApiBeanPlanMultiCurrencySnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeCurrency

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) GetChargeCurrency() string`

GetChargeCurrency returns the ChargeCurrency field if non-nil, zero value otherwise.

### GetChargeCurrencyOk

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) GetChargeCurrencyOk() (*string, bool)`

GetChargeCurrencyOk returns a tuple with the ChargeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCurrency

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) SetChargeCurrency(v string)`

SetChargeCurrency sets ChargeCurrency field to given value.

### HasChargeCurrency

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) HasChargeCurrency() bool`

HasChargeCurrency returns a boolean if a field has been set.

### GetRates

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) GetRates() []UnibeeApiBeanPlanMultiCurrencyRate`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) GetRatesOk() (*[]UnibeeApiBeanPlanMultiCurrencyRate, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) SetRates(v []UnibeeApiBeanPlanMultiCurrencyRate)`

SetRates sets Rates field to given value.

### HasRates

`func (o *UnibeeApiBeanPlanMultiCurrencySnapshot) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


