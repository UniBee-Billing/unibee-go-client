# UnibeeApiBeanDetailGatewayCurrencyExchange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExchangeRate** | Pointer to **float32** | the exchange rate of gateway, set to 0 if using https://app.exchangerate-api.com/ instead of fixed exchange rate | [optional] 
**FromCurrency** | Pointer to **string** | the currency of gateway exchange from | [optional] 
**ToCurrency** | Pointer to **string** | the currency of gateway exchange to | [optional] 

## Methods

### NewUnibeeApiBeanDetailGatewayCurrencyExchange

`func NewUnibeeApiBeanDetailGatewayCurrencyExchange() *UnibeeApiBeanDetailGatewayCurrencyExchange`

NewUnibeeApiBeanDetailGatewayCurrencyExchange instantiates a new UnibeeApiBeanDetailGatewayCurrencyExchange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailGatewayCurrencyExchangeWithDefaults

`func NewUnibeeApiBeanDetailGatewayCurrencyExchangeWithDefaults() *UnibeeApiBeanDetailGatewayCurrencyExchange`

NewUnibeeApiBeanDetailGatewayCurrencyExchangeWithDefaults instantiates a new UnibeeApiBeanDetailGatewayCurrencyExchange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExchangeRate

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetFromCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.

### HasFromCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) HasFromCurrency() bool`

HasFromCurrency returns a boolean if a field has been set.

### GetToCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.

### HasToCurrency

`func (o *UnibeeApiBeanDetailGatewayCurrencyExchange) HasToCurrency() bool`

HasToCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


