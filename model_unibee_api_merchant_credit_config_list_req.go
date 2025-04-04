/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantCreditConfigListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditConfigListReq{}

// UnibeeApiMerchantCreditConfigListReq struct for UnibeeApiMerchantCreditConfigListReq
type UnibeeApiMerchantCreditConfigListReq struct {
	// currency
	Currency *string `json:"currency,omitempty"`
	// type list of credit account, 1-main account, 2-promo credit account
	Types []int32 `json:"types,omitempty"`
}

// NewUnibeeApiMerchantCreditConfigListReq instantiates a new UnibeeApiMerchantCreditConfigListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditConfigListReq() *UnibeeApiMerchantCreditConfigListReq {
	this := UnibeeApiMerchantCreditConfigListReq{}
	return &this
}

// NewUnibeeApiMerchantCreditConfigListReqWithDefaults instantiates a new UnibeeApiMerchantCreditConfigListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditConfigListReqWithDefaults() *UnibeeApiMerchantCreditConfigListReq {
	this := UnibeeApiMerchantCreditConfigListReq{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditConfigListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditConfigListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditConfigListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantCreditConfigListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditConfigListReq) GetTypes() []int32 {
	if o == nil || IsNil(o.Types) {
		var ret []int32
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditConfigListReq) GetTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditConfigListReq) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []int32 and assigns it to the Types field.
func (o *UnibeeApiMerchantCreditConfigListReq) SetTypes(v []int32) {
	o.Types = v
}

func (o UnibeeApiMerchantCreditConfigListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditConfigListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditConfigListReq struct {
	value *UnibeeApiMerchantCreditConfigListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditConfigListReq) Get() *UnibeeApiMerchantCreditConfigListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditConfigListReq) Set(val *UnibeeApiMerchantCreditConfigListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditConfigListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditConfigListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditConfigListReq(val *UnibeeApiMerchantCreditConfigListReq) *NullableUnibeeApiMerchantCreditConfigListReq {
	return &NullableUnibeeApiMerchantCreditConfigListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditConfigListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditConfigListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


