# UnibeeApiMerchantProfileGetExchangeRateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromCurrency** | **string** | From currency code, ISO 4217 (e.g. USD) | 
**ToCurrency** | **string** | To currency code, ISO 4217 (e.g. EUR) | 

## Methods

### NewUnibeeApiMerchantProfileGetExchangeRateReq

`func NewUnibeeApiMerchantProfileGetExchangeRateReq(fromCurrency string, toCurrency string, ) *UnibeeApiMerchantProfileGetExchangeRateReq`

NewUnibeeApiMerchantProfileGetExchangeRateReq instantiates a new UnibeeApiMerchantProfileGetExchangeRateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProfileGetExchangeRateReqWithDefaults

`func NewUnibeeApiMerchantProfileGetExchangeRateReqWithDefaults() *UnibeeApiMerchantProfileGetExchangeRateReq`

NewUnibeeApiMerchantProfileGetExchangeRateReqWithDefaults instantiates a new UnibeeApiMerchantProfileGetExchangeRateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromCurrency

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.


### GetToCurrency

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *UnibeeApiMerchantProfileGetExchangeRateReq) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


