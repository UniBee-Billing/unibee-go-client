# UnibeeApiBeanGatewayCurrencyExchange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeAmount** | Pointer to **int64** | the exchange amount of gateway | [optional] 
**ExchangeRate** | Pointer to **float32** | the exchange rate of gateway, set to 0 if using https://app.exchangerate-api.com/ instead of fixed exchange rate | [optional] 
**FromCurrency** | Pointer to **string** | the currency of gateway exchange from | [optional] 
**ToCurrency** | Pointer to **string** | the currency of gateway exchange to | [optional] 

## Methods

### NewUnibeeApiBeanGatewayCurrencyExchange

`func NewUnibeeApiBeanGatewayCurrencyExchange() *UnibeeApiBeanGatewayCurrencyExchange`

NewUnibeeApiBeanGatewayCurrencyExchange instantiates a new UnibeeApiBeanGatewayCurrencyExchange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanGatewayCurrencyExchangeWithDefaults

`func NewUnibeeApiBeanGatewayCurrencyExchangeWithDefaults() *UnibeeApiBeanGatewayCurrencyExchange`

NewUnibeeApiBeanGatewayCurrencyExchangeWithDefaults instantiates a new UnibeeApiBeanGatewayCurrencyExchange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeAmount

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetExchangeAmount() int64`

GetExchangeAmount returns the ExchangeAmount field if non-nil, zero value otherwise.

### GetExchangeAmountOk

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetExchangeAmountOk() (*int64, bool)`

GetExchangeAmountOk returns a tuple with the ExchangeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAmount

`func (o *UnibeeApiBeanGatewayCurrencyExchange) SetExchangeAmount(v int64)`

SetExchangeAmount sets ExchangeAmount field to given value.

### HasExchangeAmount

`func (o *UnibeeApiBeanGatewayCurrencyExchange) HasExchangeAmount() bool`

HasExchangeAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanGatewayCurrencyExchange) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanGatewayCurrencyExchange) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetFromCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.

### HasFromCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) HasFromCurrency() bool`

HasFromCurrency returns a boolean if a field has been set.

### GetToCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *UnibeeApiBeanGatewayCurrencyExchange) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.

### HasToCurrency

`func (o *UnibeeApiBeanGatewayCurrencyExchange) HasToCurrency() bool`

HasToCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


