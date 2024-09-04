/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantDiscountListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountListRes{}

// UnibeeApiMerchantDiscountListRes struct for UnibeeApiMerchantDiscountListRes
type UnibeeApiMerchantDiscountListRes struct {
	// Discount Object List
	Discounts []UnibeeApiBeanMerchantDiscountCode `json:"discounts,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantDiscountListRes instantiates a new UnibeeApiMerchantDiscountListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountListRes() *UnibeeApiMerchantDiscountListRes {
	this := UnibeeApiMerchantDiscountListRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountListResWithDefaults instantiates a new UnibeeApiMerchantDiscountListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountListResWithDefaults() *UnibeeApiMerchantDiscountListRes {
	this := UnibeeApiMerchantDiscountListRes{}
	return &this
}

// GetDiscounts returns the Discounts field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListRes) GetDiscounts() []UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discounts) {
		var ret []UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListRes) GetDiscountsOk() ([]UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discounts) {
		return nil, false
	}
	return o.Discounts, true
}

// HasDiscounts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListRes) HasDiscounts() bool {
	if o != nil && !IsNil(o.Discounts) {
		return true
	}

	return false
}

// SetDiscounts gets a reference to the given []UnibeeApiBeanMerchantDiscountCode and assigns it to the Discounts field.
func (o *UnibeeApiMerchantDiscountListRes) SetDiscounts(v []UnibeeApiBeanMerchantDiscountCode) {
	o.Discounts = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantDiscountListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantDiscountListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discounts) {
		toSerialize["discounts"] = o.Discounts
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountListRes struct {
	value *UnibeeApiMerchantDiscountListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountListRes) Get() *UnibeeApiMerchantDiscountListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountListRes) Set(val *UnibeeApiMerchantDiscountListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountListRes(val *UnibeeApiMerchantDiscountListRes) *NullableUnibeeApiMerchantDiscountListRes {
	return &NullableUnibeeApiMerchantDiscountListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


