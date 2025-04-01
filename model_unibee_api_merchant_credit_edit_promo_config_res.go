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

// checks if the UnibeeApiMerchantCreditEditPromoConfigRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditEditPromoConfigRes{}

// UnibeeApiMerchantCreditEditPromoConfigRes struct for UnibeeApiMerchantCreditEditPromoConfigRes
type UnibeeApiMerchantCreditEditPromoConfigRes struct {
	CreditConfig *UnibeeApiBeanCreditConfig `json:"creditConfig,omitempty"`
}

// NewUnibeeApiMerchantCreditEditPromoConfigRes instantiates a new UnibeeApiMerchantCreditEditPromoConfigRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditEditPromoConfigRes() *UnibeeApiMerchantCreditEditPromoConfigRes {
	this := UnibeeApiMerchantCreditEditPromoConfigRes{}
	return &this
}

// NewUnibeeApiMerchantCreditEditPromoConfigResWithDefaults instantiates a new UnibeeApiMerchantCreditEditPromoConfigRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditEditPromoConfigResWithDefaults() *UnibeeApiMerchantCreditEditPromoConfigRes {
	this := UnibeeApiMerchantCreditEditPromoConfigRes{}
	return &this
}

// GetCreditConfig returns the CreditConfig field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigRes) GetCreditConfig() UnibeeApiBeanCreditConfig {
	if o == nil || IsNil(o.CreditConfig) {
		var ret UnibeeApiBeanCreditConfig
		return ret
	}
	return *o.CreditConfig
}

// GetCreditConfigOk returns a tuple with the CreditConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigRes) GetCreditConfigOk() (*UnibeeApiBeanCreditConfig, bool) {
	if o == nil || IsNil(o.CreditConfig) {
		return nil, false
	}
	return o.CreditConfig, true
}

// HasCreditConfig returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigRes) HasCreditConfig() bool {
	if o != nil && !IsNil(o.CreditConfig) {
		return true
	}

	return false
}

// SetCreditConfig gets a reference to the given UnibeeApiBeanCreditConfig and assigns it to the CreditConfig field.
func (o *UnibeeApiMerchantCreditEditPromoConfigRes) SetCreditConfig(v UnibeeApiBeanCreditConfig) {
	o.CreditConfig = &v
}

func (o UnibeeApiMerchantCreditEditPromoConfigRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditEditPromoConfigRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditConfig) {
		toSerialize["creditConfig"] = o.CreditConfig
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantCreditEditPromoConfigRes struct {
	value *UnibeeApiMerchantCreditEditPromoConfigRes
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigRes) Get() *UnibeeApiMerchantCreditEditPromoConfigRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigRes) Set(val *UnibeeApiMerchantCreditEditPromoConfigRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditEditPromoConfigRes(val *UnibeeApiMerchantCreditEditPromoConfigRes) *NullableUnibeeApiMerchantCreditEditPromoConfigRes {
	return &NullableUnibeeApiMerchantCreditEditPromoConfigRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


