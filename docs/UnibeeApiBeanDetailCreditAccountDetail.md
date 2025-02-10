# UnibeeApiBeanDetailCreditAccountDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | credit amount, in cent if type is main | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**CurrencyAmount** | Pointer to **int64** | currency amount, in cent | [optional] 
**ExchangeRate** | Pointer to **int64** | keep two decimal places，multiply by 100 saved, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**TotalDecrementAmount** | Pointer to **int64** | the total decrement amount | [optional] 
**TotalIncrementAmount** | Pointer to **int64** | the total increment amount | [optional] 
**Type** | Pointer to **int32** | type of credit account, 1-main account, 2-gift account | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailCreditAccountDetail

`func NewUnibeeApiBeanDetailCreditAccountDetail() *UnibeeApiBeanDetailCreditAccountDetail`

NewUnibeeApiBeanDetailCreditAccountDetail instantiates a new UnibeeApiBeanDetailCreditAccountDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailCreditAccountDetailWithDefaults

`func NewUnibeeApiBeanDetailCreditAccountDetailWithDefaults() *UnibeeApiBeanDetailCreditAccountDetail`

NewUnibeeApiBeanDetailCreditAccountDetailWithDefaults instantiates a new UnibeeApiBeanDetailCreditAccountDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCurrencyAmount() int64`

GetCurrencyAmount returns the CurrencyAmount field if non-nil, zero value otherwise.

### GetCurrencyAmountOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetCurrencyAmountOk() (*int64, bool)`

GetCurrencyAmountOk returns a tuple with the CurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetCurrencyAmount(v int64)`

SetCurrencyAmount sets CurrencyAmount field to given value.

### HasCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasCurrencyAmount() bool`

HasCurrencyAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTotalDecrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetTotalDecrementAmount() int64`

GetTotalDecrementAmount returns the TotalDecrementAmount field if non-nil, zero value otherwise.

### GetTotalDecrementAmountOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetTotalDecrementAmountOk() (*int64, bool)`

GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDecrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetTotalDecrementAmount(v int64)`

SetTotalDecrementAmount sets TotalDecrementAmount field to given value.

### HasTotalDecrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasTotalDecrementAmount() bool`

HasTotalDecrementAmount returns a boolean if a field has been set.

### GetTotalIncrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetTotalIncrementAmount() int64`

GetTotalIncrementAmount returns the TotalIncrementAmount field if non-nil, zero value otherwise.

### GetTotalIncrementAmountOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetTotalIncrementAmountOk() (*int64, bool)`

GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetTotalIncrementAmount(v int64)`

SetTotalIncrementAmount sets TotalIncrementAmount field to given value.

### HasTotalIncrementAmount

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasTotalIncrementAmount() bool`

HasTotalIncrementAmount returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailCreditAccountDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailCreditAccountDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailCreditAccountDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


