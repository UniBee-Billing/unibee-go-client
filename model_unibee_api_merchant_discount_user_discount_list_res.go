/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantDiscountUserDiscountListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountUserDiscountListRes{}

// UnibeeApiMerchantDiscountUserDiscountListRes struct for UnibeeApiMerchantDiscountUserDiscountListRes
type UnibeeApiMerchantDiscountUserDiscountListRes struct {
	// Total
	Total *int32 `json:"total,omitempty"`
	// User Discount Object List
	UserDiscounts []UnibeeApiBeanDetailMerchantUserDiscountCodeDetail `json:"userDiscounts,omitempty"`
}

// NewUnibeeApiMerchantDiscountUserDiscountListRes instantiates a new UnibeeApiMerchantDiscountUserDiscountListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountUserDiscountListRes() *UnibeeApiMerchantDiscountUserDiscountListRes {
	this := UnibeeApiMerchantDiscountUserDiscountListRes{}
	return &this
}

// NewUnibeeApiMerchantDiscountUserDiscountListResWithDefaults instantiates a new UnibeeApiMerchantDiscountUserDiscountListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountUserDiscountListResWithDefaults() *UnibeeApiMerchantDiscountUserDiscountListRes {
	this := UnibeeApiMerchantDiscountUserDiscountListRes{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) SetTotal(v int32) {
	o.Total = &v
}

// GetUserDiscounts returns the UserDiscounts field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) GetUserDiscounts() []UnibeeApiBeanDetailMerchantUserDiscountCodeDetail {
	if o == nil || IsNil(o.UserDiscounts) {
		var ret []UnibeeApiBeanDetailMerchantUserDiscountCodeDetail
		return ret
	}
	return o.UserDiscounts
}

// GetUserDiscountsOk returns a tuple with the UserDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) GetUserDiscountsOk() ([]UnibeeApiBeanDetailMerchantUserDiscountCodeDetail, bool) {
	if o == nil || IsNil(o.UserDiscounts) {
		return nil, false
	}
	return o.UserDiscounts, true
}

// HasUserDiscounts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) HasUserDiscounts() bool {
	if o != nil && !IsNil(o.UserDiscounts) {
		return true
	}

	return false
}

// SetUserDiscounts gets a reference to the given []UnibeeApiBeanDetailMerchantUserDiscountCodeDetail and assigns it to the UserDiscounts field.
func (o *UnibeeApiMerchantDiscountUserDiscountListRes) SetUserDiscounts(v []UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) {
	o.UserDiscounts = v
}

func (o UnibeeApiMerchantDiscountUserDiscountListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountUserDiscountListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.UserDiscounts) {
		toSerialize["userDiscounts"] = o.UserDiscounts
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountUserDiscountListRes struct {
	value *UnibeeApiMerchantDiscountUserDiscountListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListRes) Get() *UnibeeApiMerchantDiscountUserDiscountListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListRes) Set(val *UnibeeApiMerchantDiscountUserDiscountListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountUserDiscountListRes(val *UnibeeApiMerchantDiscountUserDiscountListRes) *NullableUnibeeApiMerchantDiscountUserDiscountListRes {
	return &NullableUnibeeApiMerchantDiscountUserDiscountListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


