/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoCurrency type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoCurrency{}

// UnibeeInternalLogicGatewayRoCurrency struct for UnibeeInternalLogicGatewayRoCurrency
type UnibeeInternalLogicGatewayRoCurrency struct {
	Currency *string `json:"Currency,omitempty"`
	Scale *int32 `json:"Scale,omitempty"`
	Symbol *string `json:"Symbol,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoCurrency instantiates a new UnibeeInternalLogicGatewayRoCurrency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoCurrency() *UnibeeInternalLogicGatewayRoCurrency {
	this := UnibeeInternalLogicGatewayRoCurrency{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoCurrencyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoCurrency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoCurrencyWithDefaults() *UnibeeInternalLogicGatewayRoCurrency {
	this := UnibeeInternalLogicGatewayRoCurrency{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeInternalLogicGatewayRoCurrency) SetCurrency(v string) {
	o.Currency = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetScale() int32 {
	if o == nil || IsNil(o.Scale) {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.Scale) {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) HasScale() bool {
	if o != nil && !IsNil(o.Scale) {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *UnibeeInternalLogicGatewayRoCurrency) SetScale(v int32) {
	o.Scale = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoCurrency) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UnibeeInternalLogicGatewayRoCurrency) SetSymbol(v string) {
	o.Symbol = &v
}

func (o UnibeeInternalLogicGatewayRoCurrency) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoCurrency) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["Currency"] = o.Currency
	}
	if !IsNil(o.Scale) {
		toSerialize["Scale"] = o.Scale
	}
	if !IsNil(o.Symbol) {
		toSerialize["Symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoCurrency struct {
	value *UnibeeInternalLogicGatewayRoCurrency
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoCurrency) Get() *UnibeeInternalLogicGatewayRoCurrency {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoCurrency) Set(val *UnibeeInternalLogicGatewayRoCurrency) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoCurrency(val *UnibeeInternalLogicGatewayRoCurrency) *NullableUnibeeInternalLogicGatewayRoCurrency {
	return &NullableUnibeeInternalLogicGatewayRoCurrency{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


