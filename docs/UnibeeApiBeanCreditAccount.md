# UnibeeApiBeanCreditAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | credit amount, in cent if type is main | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**CurrencyAmount** | Pointer to **int64** | currency amount, in cent | [optional] 
**ExchangeRate** | Pointer to **int64** | keep two decimal places，multiply by 100 saved, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**PayoutEnable** | Pointer to **int32** | 0-no, 1-yes | [optional] 
**RechargeEnable** | Pointer to **int32** | 0-no, 1-yes | [optional] 
**TotalDecrementAmount** | Pointer to **int64** | the total decrement amount | [optional] 
**TotalIncrementAmount** | Pointer to **int64** | the total increment amount | [optional] 
**Type** | Pointer to **int32** | type of credit account, 1-main account, 2-gift account | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanCreditAccount

`func NewUnibeeApiBeanCreditAccount() *UnibeeApiBeanCreditAccount`

NewUnibeeApiBeanCreditAccount instantiates a new UnibeeApiBeanCreditAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreditAccountWithDefaults

`func NewUnibeeApiBeanCreditAccountWithDefaults() *UnibeeApiBeanCreditAccount`

NewUnibeeApiBeanCreditAccountWithDefaults instantiates a new UnibeeApiBeanCreditAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanCreditAccount) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanCreditAccount) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanCreditAccount) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanCreditAccount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanCreditAccount) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanCreditAccount) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanCreditAccount) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanCreditAccount) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanCreditAccount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanCreditAccount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanCreditAccount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanCreditAccount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyAmount

`func (o *UnibeeApiBeanCreditAccount) GetCurrencyAmount() int64`

GetCurrencyAmount returns the CurrencyAmount field if non-nil, zero value otherwise.

### GetCurrencyAmountOk

`func (o *UnibeeApiBeanCreditAccount) GetCurrencyAmountOk() (*int64, bool)`

GetCurrencyAmountOk returns a tuple with the CurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAmount

`func (o *UnibeeApiBeanCreditAccount) SetCurrencyAmount(v int64)`

SetCurrencyAmount sets CurrencyAmount field to given value.

### HasCurrencyAmount

`func (o *UnibeeApiBeanCreditAccount) HasCurrencyAmount() bool`

HasCurrencyAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanCreditAccount) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanCreditAccount) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanCreditAccount) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanCreditAccount) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanCreditAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanCreditAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanCreditAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanCreditAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayoutEnable

`func (o *UnibeeApiBeanCreditAccount) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiBeanCreditAccount) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiBeanCreditAccount) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiBeanCreditAccount) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiBeanCreditAccount) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiBeanCreditAccount) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiBeanCreditAccount) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiBeanCreditAccount) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.

### GetTotalDecrementAmount

`func (o *UnibeeApiBeanCreditAccount) GetTotalDecrementAmount() int64`

GetTotalDecrementAmount returns the TotalDecrementAmount field if non-nil, zero value otherwise.

### GetTotalDecrementAmountOk

`func (o *UnibeeApiBeanCreditAccount) GetTotalDecrementAmountOk() (*int64, bool)`

GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDecrementAmount

`func (o *UnibeeApiBeanCreditAccount) SetTotalDecrementAmount(v int64)`

SetTotalDecrementAmount sets TotalDecrementAmount field to given value.

### HasTotalDecrementAmount

`func (o *UnibeeApiBeanCreditAccount) HasTotalDecrementAmount() bool`

HasTotalDecrementAmount returns a boolean if a field has been set.

### GetTotalIncrementAmount

`func (o *UnibeeApiBeanCreditAccount) GetTotalIncrementAmount() int64`

GetTotalIncrementAmount returns the TotalIncrementAmount field if non-nil, zero value otherwise.

### GetTotalIncrementAmountOk

`func (o *UnibeeApiBeanCreditAccount) GetTotalIncrementAmountOk() (*int64, bool)`

GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncrementAmount

`func (o *UnibeeApiBeanCreditAccount) SetTotalIncrementAmount(v int64)`

SetTotalIncrementAmount sets TotalIncrementAmount field to given value.

### HasTotalIncrementAmount

`func (o *UnibeeApiBeanCreditAccount) HasTotalIncrementAmount() bool`

HasTotalIncrementAmount returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanCreditAccount) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanCreditAccount) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanCreditAccount) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanCreditAccount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanCreditAccount) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanCreditAccount) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanCreditAccount) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanCreditAccount) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


