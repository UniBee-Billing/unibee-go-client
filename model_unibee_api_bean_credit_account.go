/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanCreditAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCreditAccount{}

// UnibeeApiBeanCreditAccount struct for UnibeeApiBeanCreditAccount
type UnibeeApiBeanCreditAccount struct {
	// credit amount, in cent if type is main
	Amount *int64 `json:"amount,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// currency amount, in cent
	CurrencyAmount *int64 `json:"currencyAmount,omitempty"`
	// keep two decimal places，multiply by 100 saved, 1 currency = 1 credit * (exchange_rate/100), main account fixed rate to 100
	ExchangeRate *int64 `json:"exchangeRate,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// 0-no, 1-yes
	PayoutEnable *int32 `json:"payoutEnable,omitempty"`
	// 0-no, 1-yes
	RechargeEnable *int32 `json:"rechargeEnable,omitempty"`
	// the total decrement amount
	TotalDecrementAmount *int64 `json:"totalDecrementAmount,omitempty"`
	// the total increment amount
	TotalIncrementAmount *int64 `json:"totalIncrementAmount,omitempty"`
	// type of credit account, 1-main account, 2-gift account
	Type *int32 `json:"type,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanCreditAccount instantiates a new UnibeeApiBeanCreditAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCreditAccount() *UnibeeApiBeanCreditAccount {
	this := UnibeeApiBeanCreditAccount{}
	return &this
}

// NewUnibeeApiBeanCreditAccountWithDefaults instantiates a new UnibeeApiBeanCreditAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCreditAccountWithDefaults() *UnibeeApiBeanCreditAccount {
	this := UnibeeApiBeanCreditAccount{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeApiBeanCreditAccount) SetAmount(v int64) {
	o.Amount = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanCreditAccount) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanCreditAccount) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencyAmount returns the CurrencyAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetCurrencyAmount() int64 {
	if o == nil || IsNil(o.CurrencyAmount) {
		var ret int64
		return ret
	}
	return *o.CurrencyAmount
}

// GetCurrencyAmountOk returns a tuple with the CurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetCurrencyAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrencyAmount) {
		return nil, false
	}
	return o.CurrencyAmount, true
}

// HasCurrencyAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasCurrencyAmount() bool {
	if o != nil && !IsNil(o.CurrencyAmount) {
		return true
	}

	return false
}

// SetCurrencyAmount gets a reference to the given int64 and assigns it to the CurrencyAmount field.
func (o *UnibeeApiBeanCreditAccount) SetCurrencyAmount(v int64) {
	o.CurrencyAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetExchangeRate() int64 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret int64
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetExchangeRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given int64 and assigns it to the ExchangeRate field.
func (o *UnibeeApiBeanCreditAccount) SetExchangeRate(v int64) {
	o.ExchangeRate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanCreditAccount) SetId(v int64) {
	o.Id = &v
}

// GetPayoutEnable returns the PayoutEnable field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetPayoutEnable() int32 {
	if o == nil || IsNil(o.PayoutEnable) {
		var ret int32
		return ret
	}
	return *o.PayoutEnable
}

// GetPayoutEnableOk returns a tuple with the PayoutEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetPayoutEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.PayoutEnable) {
		return nil, false
	}
	return o.PayoutEnable, true
}

// HasPayoutEnable returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasPayoutEnable() bool {
	if o != nil && !IsNil(o.PayoutEnable) {
		return true
	}

	return false
}

// SetPayoutEnable gets a reference to the given int32 and assigns it to the PayoutEnable field.
func (o *UnibeeApiBeanCreditAccount) SetPayoutEnable(v int32) {
	o.PayoutEnable = &v
}

// GetRechargeEnable returns the RechargeEnable field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetRechargeEnable() int32 {
	if o == nil || IsNil(o.RechargeEnable) {
		var ret int32
		return ret
	}
	return *o.RechargeEnable
}

// GetRechargeEnableOk returns a tuple with the RechargeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetRechargeEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.RechargeEnable) {
		return nil, false
	}
	return o.RechargeEnable, true
}

// HasRechargeEnable returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasRechargeEnable() bool {
	if o != nil && !IsNil(o.RechargeEnable) {
		return true
	}

	return false
}

// SetRechargeEnable gets a reference to the given int32 and assigns it to the RechargeEnable field.
func (o *UnibeeApiBeanCreditAccount) SetRechargeEnable(v int32) {
	o.RechargeEnable = &v
}

// GetTotalDecrementAmount returns the TotalDecrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetTotalDecrementAmount() int64 {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalDecrementAmount
}

// GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetTotalDecrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		return nil, false
	}
	return o.TotalDecrementAmount, true
}

// HasTotalDecrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasTotalDecrementAmount() bool {
	if o != nil && !IsNil(o.TotalDecrementAmount) {
		return true
	}

	return false
}

// SetTotalDecrementAmount gets a reference to the given int64 and assigns it to the TotalDecrementAmount field.
func (o *UnibeeApiBeanCreditAccount) SetTotalDecrementAmount(v int64) {
	o.TotalDecrementAmount = &v
}

// GetTotalIncrementAmount returns the TotalIncrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetTotalIncrementAmount() int64 {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalIncrementAmount
}

// GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetTotalIncrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		return nil, false
	}
	return o.TotalIncrementAmount, true
}

// HasTotalIncrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasTotalIncrementAmount() bool {
	if o != nil && !IsNil(o.TotalIncrementAmount) {
		return true
	}

	return false
}

// SetTotalIncrementAmount gets a reference to the given int64 and assigns it to the TotalIncrementAmount field.
func (o *UnibeeApiBeanCreditAccount) SetTotalIncrementAmount(v int64) {
	o.TotalIncrementAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanCreditAccount) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditAccount) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditAccount) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditAccount) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanCreditAccount) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanCreditAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCreditAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CurrencyAmount) {
		toSerialize["currencyAmount"] = o.CurrencyAmount
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PayoutEnable) {
		toSerialize["payoutEnable"] = o.PayoutEnable
	}
	if !IsNil(o.RechargeEnable) {
		toSerialize["rechargeEnable"] = o.RechargeEnable
	}
	if !IsNil(o.TotalDecrementAmount) {
		toSerialize["totalDecrementAmount"] = o.TotalDecrementAmount
	}
	if !IsNil(o.TotalIncrementAmount) {
		toSerialize["totalIncrementAmount"] = o.TotalIncrementAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCreditAccount struct {
	value *UnibeeApiBeanCreditAccount
	isSet bool
}

func (v NullableUnibeeApiBeanCreditAccount) Get() *UnibeeApiBeanCreditAccount {
	return v.value
}

func (v *NullableUnibeeApiBeanCreditAccount) Set(val *UnibeeApiBeanCreditAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCreditAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCreditAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCreditAccount(val *UnibeeApiBeanCreditAccount) *NullableUnibeeApiBeanCreditAccount {
	return &NullableUnibeeApiBeanCreditAccount{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCreditAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCreditAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


