/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProductCopyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductCopyRes{}

// UnibeeApiMerchantProductCopyRes struct for UnibeeApiMerchantProductCopyRes
type UnibeeApiMerchantProductCopyRes struct {
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewUnibeeApiMerchantProductCopyRes instantiates a new UnibeeApiMerchantProductCopyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductCopyRes() *UnibeeApiMerchantProductCopyRes {
	this := UnibeeApiMerchantProductCopyRes{}
	return &this
}

// NewUnibeeApiMerchantProductCopyResWithDefaults instantiates a new UnibeeApiMerchantProductCopyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductCopyResWithDefaults() *UnibeeApiMerchantProductCopyRes {
	this := UnibeeApiMerchantProductCopyRes{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductCopyRes) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductCopyRes) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductCopyRes) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *UnibeeApiMerchantProductCopyRes) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o UnibeeApiMerchantProductCopyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductCopyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductCopyRes struct {
	value *UnibeeApiMerchantProductCopyRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProductCopyRes) Get() *UnibeeApiMerchantProductCopyRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductCopyRes) Set(val *UnibeeApiMerchantProductCopyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductCopyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductCopyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductCopyRes(val *UnibeeApiMerchantProductCopyRes) *NullableUnibeeApiMerchantProductCopyRes {
	return &NullableUnibeeApiMerchantProductCopyRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductCopyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductCopyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


