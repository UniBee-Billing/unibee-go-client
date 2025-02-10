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

// checks if the UnibeeApiMerchantCreditPromoConfigRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditPromoConfigRes{}

// UnibeeApiMerchantCreditPromoConfigRes struct for UnibeeApiMerchantCreditPromoConfigRes
type UnibeeApiMerchantCreditPromoConfigRes struct {
	CreditConfig *UnibeeApiBeanCreditConfig `json:"creditConfig,omitempty"`
}

// NewUnibeeApiMerchantCreditPromoConfigRes instantiates a new UnibeeApiMerchantCreditPromoConfigRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditPromoConfigRes() *UnibeeApiMerchantCreditPromoConfigRes {
	this := UnibeeApiMerchantCreditPromoConfigRes{}
	return &this
}

// NewUnibeeApiMerchantCreditPromoConfigResWithDefaults instantiates a new UnibeeApiMerchantCreditPromoConfigRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditPromoConfigResWithDefaults() *UnibeeApiMerchantCreditPromoConfigRes {
	this := UnibeeApiMerchantCreditPromoConfigRes{}
	return &this
}

// GetCreditConfig returns the CreditConfig field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditPromoConfigRes) GetCreditConfig() UnibeeApiBeanCreditConfig {
	if o == nil || IsNil(o.CreditConfig) {
		var ret UnibeeApiBeanCreditConfig
		return ret
	}
	return *o.CreditConfig
}

// GetCreditConfigOk returns a tuple with the CreditConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditPromoConfigRes) GetCreditConfigOk() (*UnibeeApiBeanCreditConfig, bool) {
	if o == nil || IsNil(o.CreditConfig) {
		return nil, false
	}
	return o.CreditConfig, true
}

// HasCreditConfig returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditPromoConfigRes) HasCreditConfig() bool {
	if o != nil && !IsNil(o.CreditConfig) {
		return true
	}

	return false
}

// SetCreditConfig gets a reference to the given UnibeeApiBeanCreditConfig and assigns it to the CreditConfig field.
func (o *UnibeeApiMerchantCreditPromoConfigRes) SetCreditConfig(v UnibeeApiBeanCreditConfig) {
	o.CreditConfig = &v
}

func (o UnibeeApiMerchantCreditPromoConfigRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditPromoConfigRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditConfig) {
		toSerialize["creditConfig"] = o.CreditConfig
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditPromoConfigRes struct {
	value *UnibeeApiMerchantCreditPromoConfigRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditPromoConfigRes) Get() *UnibeeApiMerchantCreditPromoConfigRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigRes) Set(val *UnibeeApiMerchantCreditPromoConfigRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditPromoConfigRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditPromoConfigRes(val *UnibeeApiMerchantCreditPromoConfigRes) *NullableUnibeeApiMerchantCreditPromoConfigRes {
	return &NullableUnibeeApiMerchantCreditPromoConfigRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditPromoConfigRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditPromoConfigRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


