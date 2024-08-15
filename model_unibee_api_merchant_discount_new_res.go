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

// checks if the UnibeeApiMerchantDiscountNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountNewRes{}

// UnibeeApiMerchantDiscountNewRes struct for UnibeeApiMerchantDiscountNewRes
type UnibeeApiMerchantDiscountNewRes struct {
	Discount *UnibeeApiBeanMerchantDiscountCode `json:"discount,omitempty"`
}

// NewUnibeeApiMerchantDiscountNewRes instantiates a new UnibeeApiMerchantDiscountNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountNewRes() *UnibeeApiMerchantDiscountNewRes {
	this := UnibeeApiMerchantDiscountNewRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountNewResWithDefaults instantiates a new UnibeeApiMerchantDiscountNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountNewResWithDefaults() *UnibeeApiMerchantDiscountNewRes {
	this := UnibeeApiMerchantDiscountNewRes{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewRes) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the Discount field.
func (o *UnibeeApiMerchantDiscountNewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode) {
	o.Discount = &v
}

func (o UnibeeApiMerchantDiscountNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountNewRes struct {
	value *UnibeeApiMerchantDiscountNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountNewRes) Get() *UnibeeApiMerchantDiscountNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountNewRes) Set(val *UnibeeApiMerchantDiscountNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountNewRes(val *UnibeeApiMerchantDiscountNewRes) *NullableUnibeeApiMerchantDiscountNewRes {
	return &NullableUnibeeApiMerchantDiscountNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


