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

// checks if the UnibeeApiBeanCreditPayout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCreditPayout{}

// UnibeeApiBeanCreditPayout struct for UnibeeApiBeanCreditPayout
type UnibeeApiBeanCreditPayout struct {
	// credit amount, scale = 100
	CreditAmount *int64 `json:"creditAmount,omitempty"`
	// currency amount,cent
	CurrencyAmount *int64 `json:"currencyAmount,omitempty"`
	// exchange rate, keep two decimal places，scale = 100, 1 currency = 1 credit * (exchange_rate/100), main account fixed rate to 100
	ExchangeRate *int64 `json:"exchangeRate,omitempty"`
}

// NewUnibeeApiBeanCreditPayout instantiates a new UnibeeApiBeanCreditPayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCreditPayout() *UnibeeApiBeanCreditPayout {
	this := UnibeeApiBeanCreditPayout{}
	return &this
}

// NewUnibeeApiBeanCreditPayoutWithDefaults instantiates a new UnibeeApiBeanCreditPayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCreditPayoutWithDefaults() *UnibeeApiBeanCreditPayout {
	this := UnibeeApiBeanCreditPayout{}
	return &this
}

// GetCreditAmount returns the CreditAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditPayout) GetCreditAmount() int64 {
	if o == nil || IsNil(o.CreditAmount) {
		var ret int64
		return ret
	}
	return *o.CreditAmount
}

// GetCreditAmountOk returns a tuple with the CreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditPayout) GetCreditAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.CreditAmount) {
		return nil, false
	}
	return o.CreditAmount, true
}

// HasCreditAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditPayout) HasCreditAmount() bool {
	if o != nil && !IsNil(o.CreditAmount) {
		return true
	}

	return false
}

// SetCreditAmount gets a reference to the given int64 and assigns it to the CreditAmount field.
func (o *UnibeeApiBeanCreditPayout) SetCreditAmount(v int64) {
	o.CreditAmount = &v
}

// GetCurrencyAmount returns the CurrencyAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditPayout) GetCurrencyAmount() int64 {
	if o == nil || IsNil(o.CurrencyAmount) {
		var ret int64
		return ret
	}
	return *o.CurrencyAmount
}

// GetCurrencyAmountOk returns a tuple with the CurrencyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditPayout) GetCurrencyAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrencyAmount) {
		return nil, false
	}
	return o.CurrencyAmount, true
}

// HasCurrencyAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditPayout) HasCurrencyAmount() bool {
	if o != nil && !IsNil(o.CurrencyAmount) {
		return true
	}

	return false
}

// SetCurrencyAmount gets a reference to the given int64 and assigns it to the CurrencyAmount field.
func (o *UnibeeApiBeanCreditPayout) SetCurrencyAmount(v int64) {
	o.CurrencyAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditPayout) GetExchangeRate() int64 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret int64
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditPayout) GetExchangeRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditPayout) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given int64 and assigns it to the ExchangeRate field.
func (o *UnibeeApiBeanCreditPayout) SetExchangeRate(v int64) {
	o.ExchangeRate = &v
}

func (o UnibeeApiBeanCreditPayout) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCreditPayout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditAmount) {
		toSerialize["creditAmount"] = o.CreditAmount
	}
	if !IsNil(o.CurrencyAmount) {
		toSerialize["currencyAmount"] = o.CurrencyAmount
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCreditPayout struct {
	value *UnibeeApiBeanCreditPayout
	isSet bool
}

func (v NullableUnibeeApiBeanCreditPayout) Get() *UnibeeApiBeanCreditPayout {
	return v.value
}

func (v *NullableUnibeeApiBeanCreditPayout) Set(val *UnibeeApiBeanCreditPayout) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCreditPayout) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCreditPayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCreditPayout(val *UnibeeApiBeanCreditPayout) *NullableUnibeeApiBeanCreditPayout {
	return &NullableUnibeeApiBeanCreditPayout{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCreditPayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCreditPayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


