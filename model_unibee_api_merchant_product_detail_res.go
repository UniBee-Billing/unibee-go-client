/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProductDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductDetailRes{}

// UnibeeApiMerchantProductDetailRes struct for UnibeeApiMerchantProductDetailRes
type UnibeeApiMerchantProductDetailRes struct {
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewUnibeeApiMerchantProductDetailRes instantiates a new UnibeeApiMerchantProductDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductDetailRes() *UnibeeApiMerchantProductDetailRes {
	this := UnibeeApiMerchantProductDetailRes{}
	return &this
}

// NewUnibeeApiMerchantProductDetailResWithDefaults instantiates a new UnibeeApiMerchantProductDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductDetailResWithDefaults() *UnibeeApiMerchantProductDetailRes {
	this := UnibeeApiMerchantProductDetailRes{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductDetailRes) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductDetailRes) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductDetailRes) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *UnibeeApiMerchantProductDetailRes) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o UnibeeApiMerchantProductDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductDetailRes struct {
	value *UnibeeApiMerchantProductDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProductDetailRes) Get() *UnibeeApiMerchantProductDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductDetailRes) Set(val *UnibeeApiMerchantProductDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductDetailRes(val *UnibeeApiMerchantProductDetailRes) *NullableUnibeeApiMerchantProductDetailRes {
	return &NullableUnibeeApiMerchantProductDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


