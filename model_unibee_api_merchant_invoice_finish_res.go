/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceFinishRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceFinishRes{}

// UnibeeApiMerchantInvoiceFinishRes struct for UnibeeApiMerchantInvoiceFinishRes
type UnibeeApiMerchantInvoiceFinishRes struct {
	Invoice *UnibeeInternalModelEntityOverseaPayInvoice `json:"invoice,omitempty"`
}

// NewUnibeeApiMerchantInvoiceFinishRes instantiates a new UnibeeApiMerchantInvoiceFinishRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceFinishRes() *UnibeeApiMerchantInvoiceFinishRes {
	this := UnibeeApiMerchantInvoiceFinishRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceFinishResWithDefaults instantiates a new UnibeeApiMerchantInvoiceFinishRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceFinishResWithDefaults() *UnibeeApiMerchantInvoiceFinishRes {
	this := UnibeeApiMerchantInvoiceFinishRes{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceFinishRes) GetInvoice() UnibeeInternalModelEntityOverseaPayInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeInternalModelEntityOverseaPayInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceFinishRes) GetInvoiceOk() (*UnibeeInternalModelEntityOverseaPayInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceFinishRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeInternalModelEntityOverseaPayInvoice and assigns it to the Invoice field.
func (o *UnibeeApiMerchantInvoiceFinishRes) SetInvoice(v UnibeeInternalModelEntityOverseaPayInvoice) {
	o.Invoice = &v
}

func (o UnibeeApiMerchantInvoiceFinishRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceFinishRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceFinishRes struct {
	value *UnibeeApiMerchantInvoiceFinishRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceFinishRes) Get() *UnibeeApiMerchantInvoiceFinishRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceFinishRes) Set(val *UnibeeApiMerchantInvoiceFinishRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceFinishRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceFinishRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceFinishRes(val *UnibeeApiMerchantInvoiceFinishRes) *NullableUnibeeApiMerchantInvoiceFinishRes {
	return &NullableUnibeeApiMerchantInvoiceFinishRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceFinishRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceFinishRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


