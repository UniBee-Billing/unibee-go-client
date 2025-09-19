# UnibeeApiBeanMerchantCurrencyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoExchange** | Pointer to **bool** | using https://app.exchangerate-api.com/ to update exchange rate if true, the exchange APIKey need setup first | [optional] 
**Currency** | Pointer to **string** | target currency | [optional] 
**ExchangeRate** | Pointer to **float32** | the exchange rate of gateway, no setup required if AutoExchange is true | [optional] 

## Methods

### NewUnibeeApiBeanMerchantCurrencyConfig

`func NewUnibeeApiBeanMerchantCurrencyConfig() *UnibeeApiBeanMerchantCurrencyConfig`

NewUnibeeApiBeanMerchantCurrencyConfig instantiates a new UnibeeApiBeanMerchantCurrencyConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantCurrencyConfigWithDefaults

`func NewUnibeeApiBeanMerchantCurrencyConfigWithDefaults() *UnibeeApiBeanMerchantCurrencyConfig`

NewUnibeeApiBeanMerchantCurrencyConfigWithDefaults instantiates a new UnibeeApiBeanMerchantCurrencyConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoExchange

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetAutoExchange() bool`

GetAutoExchange returns the AutoExchange field if non-nil, zero value otherwise.

### GetAutoExchangeOk

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetAutoExchangeOk() (*bool, bool)`

GetAutoExchangeOk returns a tuple with the AutoExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExchange

`func (o *UnibeeApiBeanMerchantCurrencyConfig) SetAutoExchange(v bool)`

SetAutoExchange sets AutoExchange field to given value.

### HasAutoExchange

`func (o *UnibeeApiBeanMerchantCurrencyConfig) HasAutoExchange() bool`

HasAutoExchange returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanMerchantCurrencyConfig) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanMerchantCurrencyConfig) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanMerchantCurrencyConfig) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanMerchantCurrencyConfig) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanMerchantCurrencyConfig) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


