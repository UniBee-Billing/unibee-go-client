# UnibeeApiBeanPlanMultiCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | the amount of exchange rate | [optional] 
**AutoExchange** | Pointer to **bool** | using https://app.exchangerate-api.com/ to update exchange rate if true, the exchange APIKey need setup first | [optional] 
**Currency** | Pointer to **string** | target currency | [optional] 
**Disable** | Pointer to **bool** | disable currency exchange | [optional] 
**ExchangeRate** | Pointer to **float32** | exchange rate, no setup required if AutoExchange is true | [optional] 

## Methods

### NewUnibeeApiBeanPlanMultiCurrency

`func NewUnibeeApiBeanPlanMultiCurrency() *UnibeeApiBeanPlanMultiCurrency`

NewUnibeeApiBeanPlanMultiCurrency instantiates a new UnibeeApiBeanPlanMultiCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanMultiCurrencyWithDefaults

`func NewUnibeeApiBeanPlanMultiCurrencyWithDefaults() *UnibeeApiBeanPlanMultiCurrency`

NewUnibeeApiBeanPlanMultiCurrencyWithDefaults instantiates a new UnibeeApiBeanPlanMultiCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanPlanMultiCurrency) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanPlanMultiCurrency) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanPlanMultiCurrency) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanPlanMultiCurrency) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAutoExchange

`func (o *UnibeeApiBeanPlanMultiCurrency) GetAutoExchange() bool`

GetAutoExchange returns the AutoExchange field if non-nil, zero value otherwise.

### GetAutoExchangeOk

`func (o *UnibeeApiBeanPlanMultiCurrency) GetAutoExchangeOk() (*bool, bool)`

GetAutoExchangeOk returns a tuple with the AutoExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchange

`func (o *UnibeeApiBeanPlanMultiCurrency) SetAutoExchange(v bool)`

SetAutoExchange sets AutoExchange field to given value.

### HasAutoExchange

`func (o *UnibeeApiBeanPlanMultiCurrency) HasAutoExchange() bool`

HasAutoExchange returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPlanMultiCurrency) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPlanMultiCurrency) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPlanMultiCurrency) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPlanMultiCurrency) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisable

`func (o *UnibeeApiBeanPlanMultiCurrency) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *UnibeeApiBeanPlanMultiCurrency) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *UnibeeApiBeanPlanMultiCurrency) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *UnibeeApiBeanPlanMultiCurrency) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanPlanMultiCurrency) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanPlanMultiCurrency) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanPlanMultiCurrency) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanPlanMultiCurrency) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


