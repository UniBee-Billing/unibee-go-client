/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoGatewayBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoGatewayBalance{}

// UnibeeInternalLogicGatewayRoGatewayBalance struct for UnibeeInternalLogicGatewayRoGatewayBalance
type UnibeeInternalLogicGatewayRoGatewayBalance struct {
	Amount *int64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoGatewayBalance instantiates a new UnibeeInternalLogicGatewayRoGatewayBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoGatewayBalance() *UnibeeInternalLogicGatewayRoGatewayBalance {
	this := UnibeeInternalLogicGatewayRoGatewayBalance{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoGatewayBalanceWithDefaults instantiates a new UnibeeInternalLogicGatewayRoGatewayBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoGatewayBalanceWithDefaults() *UnibeeInternalLogicGatewayRoGatewayBalance {
	this := UnibeeInternalLogicGatewayRoGatewayBalance{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeInternalLogicGatewayRoGatewayBalance) SetCurrency(v string) {
	o.Currency = &v
}

func (o UnibeeInternalLogicGatewayRoGatewayBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoGatewayBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoGatewayBalance struct {
	value *UnibeeInternalLogicGatewayRoGatewayBalance
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoGatewayBalance) Get() *UnibeeInternalLogicGatewayRoGatewayBalance {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewayBalance) Set(val *UnibeeInternalLogicGatewayRoGatewayBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoGatewayBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewayBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoGatewayBalance(val *UnibeeInternalLogicGatewayRoGatewayBalance) *NullableUnibeeInternalLogicGatewayRoGatewayBalance {
	return &NullableUnibeeInternalLogicGatewayRoGatewayBalance{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoGatewayBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoGatewayBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


