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

// checks if the UnibeeApiMerchantDiscountQuantityDecrementRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountQuantityDecrementRes{}

// UnibeeApiMerchantDiscountQuantityDecrementRes struct for UnibeeApiMerchantDiscountQuantityDecrementRes
type UnibeeApiMerchantDiscountQuantityDecrementRes struct {
	DiscountCode *UnibeeApiBeanMerchantDiscountCode `json:"discountCode,omitempty"`
}

// NewUnibeeApiMerchantDiscountQuantityDecrementRes instantiates a new UnibeeApiMerchantDiscountQuantityDecrementRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountQuantityDecrementRes() *UnibeeApiMerchantDiscountQuantityDecrementRes {
	this := UnibeeApiMerchantDiscountQuantityDecrementRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountQuantityDecrementResWithDefaults instantiates a new UnibeeApiMerchantDiscountQuantityDecrementRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountQuantityDecrementResWithDefaults() *UnibeeApiMerchantDiscountQuantityDecrementRes {
	this := UnibeeApiMerchantDiscountQuantityDecrementRes{}
	return &this
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountQuantityDecrementRes) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.DiscountCode) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountQuantityDecrementRes) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountQuantityDecrementRes) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantDiscountQuantityDecrementRes) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode) {
	o.DiscountCode = &v
}

func (o UnibeeApiMerchantDiscountQuantityDecrementRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountQuantityDecrementRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountQuantityDecrementRes struct {
	value *UnibeeApiMerchantDiscountQuantityDecrementRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountQuantityDecrementRes) Get() *UnibeeApiMerchantDiscountQuantityDecrementRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountQuantityDecrementRes) Set(val *UnibeeApiMerchantDiscountQuantityDecrementRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountQuantityDecrementRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountQuantityDecrementRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountQuantityDecrementRes(val *UnibeeApiMerchantDiscountQuantityDecrementRes) *NullableUnibeeApiMerchantDiscountQuantityDecrementRes {
	return &NullableUnibeeApiMerchantDiscountQuantityDecrementRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountQuantityDecrementRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountQuantityDecrementRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


