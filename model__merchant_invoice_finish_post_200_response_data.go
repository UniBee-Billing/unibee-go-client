/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantInvoiceFinishPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantInvoiceFinishPost200ResponseData{}

// MerchantInvoiceFinishPost200ResponseData struct for MerchantInvoiceFinishPost200ResponseData
type MerchantInvoiceFinishPost200ResponseData struct {
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
}

// NewMerchantInvoiceFinishPost200ResponseData instantiates a new MerchantInvoiceFinishPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantInvoiceFinishPost200ResponseData() *MerchantInvoiceFinishPost200ResponseData {
	this := MerchantInvoiceFinishPost200ResponseData{}
	return &this
}

// NewMerchantInvoiceFinishPost200ResponseDataWithDefaults instantiates a new MerchantInvoiceFinishPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantInvoiceFinishPost200ResponseDataWithDefaults() *MerchantInvoiceFinishPost200ResponseData {
	this := MerchantInvoiceFinishPost200ResponseData{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantInvoiceFinishPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceFinishPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantInvoiceFinishPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *MerchantInvoiceFinishPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

func (o MerchantInvoiceFinishPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantInvoiceFinishPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableMerchantInvoiceFinishPost200ResponseData struct {
	value *MerchantInvoiceFinishPost200ResponseData
	isSet bool
}

func (v NullableMerchantInvoiceFinishPost200ResponseData) Get() *MerchantInvoiceFinishPost200ResponseData {
	return v.value
}

func (v *NullableMerchantInvoiceFinishPost200ResponseData) Set(val *MerchantInvoiceFinishPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantInvoiceFinishPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantInvoiceFinishPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantInvoiceFinishPost200ResponseData(val *MerchantInvoiceFinishPost200ResponseData) *NullableMerchantInvoiceFinishPost200ResponseData {
	return &NullableMerchantInvoiceFinishPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantInvoiceFinishPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantInvoiceFinishPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


