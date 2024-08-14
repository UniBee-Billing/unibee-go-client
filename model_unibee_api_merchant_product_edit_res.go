/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProductEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductEditRes{}

// UnibeeApiMerchantProductEditRes struct for UnibeeApiMerchantProductEditRes
type UnibeeApiMerchantProductEditRes struct {
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewUnibeeApiMerchantProductEditRes instantiates a new UnibeeApiMerchantProductEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductEditRes() *UnibeeApiMerchantProductEditRes {
	this := UnibeeApiMerchantProductEditRes{}
	return &this
}

// NewUnibeeApiMerchantProductEditResWithDefaults instantiates a new UnibeeApiMerchantProductEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductEditResWithDefaults() *UnibeeApiMerchantProductEditRes {
	this := UnibeeApiMerchantProductEditRes{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditRes) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditRes) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditRes) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *UnibeeApiMerchantProductEditRes) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o UnibeeApiMerchantProductEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductEditRes struct {
	value *UnibeeApiMerchantProductEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProductEditRes) Get() *UnibeeApiMerchantProductEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductEditRes) Set(val *UnibeeApiMerchantProductEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductEditRes(val *UnibeeApiMerchantProductEditRes) *NullableUnibeeApiMerchantProductEditRes {
	return &NullableUnibeeApiMerchantProductEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


