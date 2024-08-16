/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantProductCopyPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantProductCopyPost200ResponseData{}

// MerchantProductCopyPost200ResponseData struct for MerchantProductCopyPost200ResponseData
type MerchantProductCopyPost200ResponseData struct {
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewMerchantProductCopyPost200ResponseData instantiates a new MerchantProductCopyPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantProductCopyPost200ResponseData() *MerchantProductCopyPost200ResponseData {
	this := MerchantProductCopyPost200ResponseData{}
	return &this
}

// NewMerchantProductCopyPost200ResponseDataWithDefaults instantiates a new MerchantProductCopyPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantProductCopyPost200ResponseDataWithDefaults() *MerchantProductCopyPost200ResponseData {
	this := MerchantProductCopyPost200ResponseData{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *MerchantProductCopyPost200ResponseData) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantProductCopyPost200ResponseData) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *MerchantProductCopyPost200ResponseData) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *MerchantProductCopyPost200ResponseData) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o MerchantProductCopyPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantProductCopyPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableMerchantProductCopyPost200ResponseData struct {
	value *MerchantProductCopyPost200ResponseData
	isSet bool
}

func (v NullableMerchantProductCopyPost200ResponseData) Get() *MerchantProductCopyPost200ResponseData {
	return v.value
}

func (v *NullableMerchantProductCopyPost200ResponseData) Set(val *MerchantProductCopyPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantProductCopyPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantProductCopyPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantProductCopyPost200ResponseData(val *MerchantProductCopyPost200ResponseData) *NullableMerchantProductCopyPost200ResponseData {
	return &NullableMerchantProductCopyPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantProductCopyPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantProductCopyPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


