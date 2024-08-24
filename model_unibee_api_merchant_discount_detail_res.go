/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantDiscountDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountDetailRes{}

// UnibeeApiMerchantDiscountDetailRes struct for UnibeeApiMerchantDiscountDetailRes
type UnibeeApiMerchantDiscountDetailRes struct {
	Discount *UnibeeApiBeanDetailMerchantDiscountCodeDetail `json:"discount,omitempty"`
}

// NewUnibeeApiMerchantDiscountDetailRes instantiates a new UnibeeApiMerchantDiscountDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountDetailRes() *UnibeeApiMerchantDiscountDetailRes {
	this := UnibeeApiMerchantDiscountDetailRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountDetailResWithDefaults instantiates a new UnibeeApiMerchantDiscountDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountDetailResWithDefaults() *UnibeeApiMerchantDiscountDetailRes {
	this := UnibeeApiMerchantDiscountDetailRes{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountDetailRes) GetDiscount() UnibeeApiBeanDetailMerchantDiscountCodeDetail {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanDetailMerchantDiscountCodeDetail
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountDetailRes) GetDiscountOk() (*UnibeeApiBeanDetailMerchantDiscountCodeDetail, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountDetailRes) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanDetailMerchantDiscountCodeDetail and assigns it to the Discount field.
func (o *UnibeeApiMerchantDiscountDetailRes) SetDiscount(v UnibeeApiBeanDetailMerchantDiscountCodeDetail) {
	o.Discount = &v
}

func (o UnibeeApiMerchantDiscountDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountDetailRes struct {
	value *UnibeeApiMerchantDiscountDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountDetailRes) Get() *UnibeeApiMerchantDiscountDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountDetailRes) Set(val *UnibeeApiMerchantDiscountDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountDetailRes(val *UnibeeApiMerchantDiscountDetailRes) *NullableUnibeeApiMerchantDiscountDetailRes {
	return &NullableUnibeeApiMerchantDiscountDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


