/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantDiscountEditPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantDiscountEditPost200ResponseData{}

// MerchantDiscountEditPost200ResponseData struct for MerchantDiscountEditPost200ResponseData
type MerchantDiscountEditPost200ResponseData struct {
	Discount *UnibeeApiBeanMerchantDiscountCodeSimplify `json:"discount,omitempty"`
}

// NewMerchantDiscountEditPost200ResponseData instantiates a new MerchantDiscountEditPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDiscountEditPost200ResponseData() *MerchantDiscountEditPost200ResponseData {
	this := MerchantDiscountEditPost200ResponseData{}
	return &this
}

// NewMerchantDiscountEditPost200ResponseDataWithDefaults instantiates a new MerchantDiscountEditPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDiscountEditPost200ResponseDataWithDefaults() *MerchantDiscountEditPost200ResponseData {
	this := MerchantDiscountEditPost200ResponseData{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *MerchantDiscountEditPost200ResponseData) GetDiscount() UnibeeApiBeanMerchantDiscountCodeSimplify {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCodeSimplify
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountEditPost200ResponseData) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCodeSimplify, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *MerchantDiscountEditPost200ResponseData) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCodeSimplify and assigns it to the Discount field.
func (o *MerchantDiscountEditPost200ResponseData) SetDiscount(v UnibeeApiBeanMerchantDiscountCodeSimplify) {
	o.Discount = &v
}

func (o MerchantDiscountEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantDiscountEditPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	return toSerialize, nil
}

type NullableMerchantDiscountEditPost200ResponseData struct {
	value *MerchantDiscountEditPost200ResponseData
	isSet bool
}

func (v NullableMerchantDiscountEditPost200ResponseData) Get() *MerchantDiscountEditPost200ResponseData {
	return v.value
}

func (v *NullableMerchantDiscountEditPost200ResponseData) Set(val *MerchantDiscountEditPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDiscountEditPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDiscountEditPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDiscountEditPost200ResponseData(val *MerchantDiscountEditPost200ResponseData) *NullableMerchantDiscountEditPost200ResponseData {
	return &NullableMerchantDiscountEditPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantDiscountEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDiscountEditPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


