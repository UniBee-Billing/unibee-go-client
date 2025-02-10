/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantDiscountDecreaseQuantityPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantDiscountDecreaseQuantityPost200ResponseData{}

// MerchantDiscountDecreaseQuantityPost200ResponseData struct for MerchantDiscountDecreaseQuantityPost200ResponseData
type MerchantDiscountDecreaseQuantityPost200ResponseData struct {
	DiscountCode *UnibeeApiBeanMerchantDiscountCode `json:"discountCode,omitempty"`
}

// NewMerchantDiscountDecreaseQuantityPost200ResponseData instantiates a new MerchantDiscountDecreaseQuantityPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantDiscountDecreaseQuantityPost200ResponseData() *MerchantDiscountDecreaseQuantityPost200ResponseData {
	this := MerchantDiscountDecreaseQuantityPost200ResponseData{}
	return &this
}

// NewMerchantDiscountDecreaseQuantityPost200ResponseDataWithDefaults instantiates a new MerchantDiscountDecreaseQuantityPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantDiscountDecreaseQuantityPost200ResponseDataWithDefaults() *MerchantDiscountDecreaseQuantityPost200ResponseData {
	this := MerchantDiscountDecreaseQuantityPost200ResponseData{}
	return &this
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *MerchantDiscountDecreaseQuantityPost200ResponseData) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.DiscountCode) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantDiscountDecreaseQuantityPost200ResponseData) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *MerchantDiscountDecreaseQuantityPost200ResponseData) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the DiscountCode field.
func (o *MerchantDiscountDecreaseQuantityPost200ResponseData) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode) {
	o.DiscountCode = &v
}

func (o MerchantDiscountDecreaseQuantityPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantDiscountDecreaseQuantityPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	return toSerialize, nil
}

type NullableMerchantDiscountDecreaseQuantityPost200ResponseData struct {
	value *MerchantDiscountDecreaseQuantityPost200ResponseData
	isSet bool
}

func (v NullableMerchantDiscountDecreaseQuantityPost200ResponseData) Get() *MerchantDiscountDecreaseQuantityPost200ResponseData {
	return v.value
}

func (v *NullableMerchantDiscountDecreaseQuantityPost200ResponseData) Set(val *MerchantDiscountDecreaseQuantityPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantDiscountDecreaseQuantityPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantDiscountDecreaseQuantityPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantDiscountDecreaseQuantityPost200ResponseData(val *MerchantDiscountDecreaseQuantityPost200ResponseData) *NullableMerchantDiscountDecreaseQuantityPost200ResponseData {
	return &NullableMerchantDiscountDecreaseQuantityPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantDiscountDecreaseQuantityPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantDiscountDecreaseQuantityPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


