# UnibeeApiBeanCreditPayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAmount** | Pointer to **int64** | credit amount, scale &#x3D; 100 | [optional] 
**CurrencyAmount** | Pointer to **int64** | currency amount,cent | [optional] 
**ExchangeRate** | Pointer to **int64** | exchange rate, keep two decimal places，scale &#x3D; 100, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 

## Methods

### NewUnibeeApiBeanCreditPayout

`func NewUnibeeApiBeanCreditPayout() *UnibeeApiBeanCreditPayout`

NewUnibeeApiBeanCreditPayout instantiates a new UnibeeApiBeanCreditPayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreditPayoutWithDefaults

`func NewUnibeeApiBeanCreditPayoutWithDefaults() *UnibeeApiBeanCreditPayout`

NewUnibeeApiBeanCreditPayoutWithDefaults instantiates a new UnibeeApiBeanCreditPayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAmount

`func (o *UnibeeApiBeanCreditPayout) GetCreditAmount() int64`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *UnibeeApiBeanCreditPayout) GetCreditAmountOk() (*int64, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *UnibeeApiBeanCreditPayout) SetCreditAmount(v int64)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *UnibeeApiBeanCreditPayout) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

### GetCurrencyAmount

`func (o *UnibeeApiBeanCreditPayout) GetCurrencyAmount() int64`

GetCurrencyAmount returns the CurrencyAmount field if non-nil, zero value otherwise.

### GetCurrencyAmountOk

`func (o *UnibeeApiBeanCreditPayout) GetCurrencyAmountOk() (*int64, bool)`

GetCurrencyAmountOk returns a tuple with the CurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAmount

`func (o *UnibeeApiBeanCreditPayout) SetCurrencyAmount(v int64)`

SetCurrencyAmount sets CurrencyAmount field to given value.

### HasCurrencyAmount

`func (o *UnibeeApiBeanCreditPayout) HasCurrencyAmount() bool`

HasCurrencyAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanCreditPayout) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanCreditPayout) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanCreditPayout) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanCreditPayout) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


