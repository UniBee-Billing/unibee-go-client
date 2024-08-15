/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayGatewayBeanGatewayBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayGatewayBeanGatewayBalance{}

// UnibeeInternalLogicGatewayGatewayBeanGatewayBalance struct for UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
type UnibeeInternalLogicGatewayGatewayBeanGatewayBalance struct {
	Amount *int64 `json:"amount,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// NewUnibeeInternalLogicGatewayGatewayBeanGatewayBalance instantiates a new UnibeeInternalLogicGatewayGatewayBeanGatewayBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayGatewayBeanGatewayBalance() *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	this := UnibeeInternalLogicGatewayGatewayBeanGatewayBalance{}
	return &this
}

// NewUnibeeInternalLogicGatewayGatewayBeanGatewayBalanceWithDefaults instantiates a new UnibeeInternalLogicGatewayGatewayBeanGatewayBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayGatewayBeanGatewayBalanceWithDefaults() *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	this := UnibeeInternalLogicGatewayGatewayBeanGatewayBalance{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) SetAmount(v int64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) SetCurrency(v string) {
	o.Currency = &v
}

func (o UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance struct {
	value *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) Get() *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) Set(val *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance(val *UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) *NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	return &NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayGatewayBeanGatewayBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


