/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProductListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductListRes{}

// UnibeeApiMerchantProductListRes struct for UnibeeApiMerchantProductListRes
type UnibeeApiMerchantProductListRes struct {
	// Product Object List
	Products []UnibeeApiBeanProduct `json:"products,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantProductListRes instantiates a new UnibeeApiMerchantProductListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductListRes() *UnibeeApiMerchantProductListRes {
	this := UnibeeApiMerchantProductListRes{}
	return &this
}

// NewUnibeeApiMerchantProductListResWithDefaults instantiates a new UnibeeApiMerchantProductListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductListResWithDefaults() *UnibeeApiMerchantProductListRes {
	this := UnibeeApiMerchantProductListRes{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductListRes) GetProducts() []UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Products) {
		var ret []UnibeeApiBeanProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductListRes) GetProductsOk() ([]UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductListRes) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []UnibeeApiBeanProduct and assigns it to the Products field.
func (o *UnibeeApiMerchantProductListRes) SetProducts(v []UnibeeApiBeanProduct) {
	o.Products = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantProductListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantProductListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductListRes struct {
	value *UnibeeApiMerchantProductListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantProductListRes) Get() *UnibeeApiMerchantProductListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductListRes) Set(val *UnibeeApiMerchantProductListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductListRes(val *UnibeeApiMerchantProductListRes) *NullableUnibeeApiMerchantProductListRes {
	return &NullableUnibeeApiMerchantProductListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


