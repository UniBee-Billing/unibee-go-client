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

// checks if the UnibeeApiMerchantInvoiceEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceEditRes{}

// UnibeeApiMerchantInvoiceEditRes struct for UnibeeApiMerchantInvoiceEditRes
type UnibeeApiMerchantInvoiceEditRes struct {
	Invoice *UnibeeApiBeanDetailInvoiceDetail `json:"invoice,omitempty"`
}

// NewUnibeeApiMerchantInvoiceEditRes instantiates a new UnibeeApiMerchantInvoiceEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceEditRes() *UnibeeApiMerchantInvoiceEditRes {
	this := UnibeeApiMerchantInvoiceEditRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceEditResWithDefaults instantiates a new UnibeeApiMerchantInvoiceEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceEditResWithDefaults() *UnibeeApiMerchantInvoiceEditRes {
	this := UnibeeApiMerchantInvoiceEditRes{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceEditRes) GetInvoice() UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceEditRes) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceEditRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanDetailInvoiceDetail and assigns it to the Invoice field.
func (o *UnibeeApiMerchantInvoiceEditRes) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail) {
	o.Invoice = &v
}

func (o UnibeeApiMerchantInvoiceEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceEditRes struct {
	value *UnibeeApiMerchantInvoiceEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceEditRes) Get() *UnibeeApiMerchantInvoiceEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceEditRes) Set(val *UnibeeApiMerchantInvoiceEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceEditRes(val *UnibeeApiMerchantInvoiceEditRes) *NullableUnibeeApiMerchantInvoiceEditRes {
	return &NullableUnibeeApiMerchantInvoiceEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


