/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantProductListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantProductListGet200ResponseData{}

// MerchantProductListGet200ResponseData struct for MerchantProductListGet200ResponseData
type MerchantProductListGet200ResponseData struct {
	// Product Object List
	Products []UnibeeApiBeanProduct `json:"products,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantProductListGet200ResponseData instantiates a new MerchantProductListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantProductListGet200ResponseData() *MerchantProductListGet200ResponseData {
	this := MerchantProductListGet200ResponseData{}
	return &this
}

// NewMerchantProductListGet200ResponseDataWithDefaults instantiates a new MerchantProductListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantProductListGet200ResponseDataWithDefaults() *MerchantProductListGet200ResponseData {
	this := MerchantProductListGet200ResponseData{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *MerchantProductListGet200ResponseData) GetProducts() []UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Products) {
		var ret []UnibeeApiBeanProduct
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductListGet200ResponseData) GetProductsOk() ([]UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *MerchantProductListGet200ResponseData) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []UnibeeApiBeanProduct and assigns it to the Products field.
func (o *MerchantProductListGet200ResponseData) SetProducts(v []UnibeeApiBeanProduct) {
	o.Products = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantProductListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantProductListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantProductListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantProductListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantProductListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantProductListGet200ResponseData struct {
	value *MerchantProductListGet200ResponseData
	isSet bool
}

func (v NullableMerchantProductListGet200ResponseData) Get() *MerchantProductListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantProductListGet200ResponseData) Set(val *MerchantProductListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantProductListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantProductListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantProductListGet200ResponseData(val *MerchantProductListGet200ResponseData) *NullableMerchantProductListGet200ResponseData {
	return &NullableMerchantProductListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantProductListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantProductListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


