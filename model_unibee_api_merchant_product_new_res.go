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

// checks if the UnibeeApiMerchantProductNewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductNewRes{}

// UnibeeApiMerchantProductNewRes struct for UnibeeApiMerchantProductNewRes
type UnibeeApiMerchantProductNewRes struct {
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewUnibeeApiMerchantProductNewRes instantiates a new UnibeeApiMerchantProductNewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductNewRes() *UnibeeApiMerchantProductNewRes {
	this := UnibeeApiMerchantProductNewRes{}
	return &this
}

// NewUnibeeApiMerchantProductNewResWithDefaults instantiates a new UnibeeApiMerchantProductNewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductNewResWithDefaults() *UnibeeApiMerchantProductNewRes {
	this := UnibeeApiMerchantProductNewRes{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewRes) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewRes) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewRes) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *UnibeeApiMerchantProductNewRes) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o UnibeeApiMerchantProductNewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductNewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductNewRes struct {
	value *UnibeeApiMerchantProductNewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProductNewRes) Get() *UnibeeApiMerchantProductNewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductNewRes) Set(val *UnibeeApiMerchantProductNewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductNewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductNewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductNewRes(val *UnibeeApiMerchantProductNewRes) *NullableUnibeeApiMerchantProductNewRes {
	return &NullableUnibeeApiMerchantProductNewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductNewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductNewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


