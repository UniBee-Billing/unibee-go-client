/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantInvoiceEditPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantInvoiceEditPost200ResponseData{}

// MerchantInvoiceEditPost200ResponseData struct for MerchantInvoiceEditPost200ResponseData
type MerchantInvoiceEditPost200ResponseData struct {
	Invoice *UnibeeApiBeanDetailInvoiceDetail `json:"invoice,omitempty"`
}

// NewMerchantInvoiceEditPost200ResponseData instantiates a new MerchantInvoiceEditPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantInvoiceEditPost200ResponseData() *MerchantInvoiceEditPost200ResponseData {
	this := MerchantInvoiceEditPost200ResponseData{}
	return &this
}

// NewMerchantInvoiceEditPost200ResponseDataWithDefaults instantiates a new MerchantInvoiceEditPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantInvoiceEditPost200ResponseDataWithDefaults() *MerchantInvoiceEditPost200ResponseData {
	this := MerchantInvoiceEditPost200ResponseData{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantInvoiceEditPost200ResponseData) GetInvoice() UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceEditPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantInvoiceEditPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanDetailInvoiceDetail and assigns it to the Invoice field.
func (o *MerchantInvoiceEditPost200ResponseData) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail) {
	o.Invoice = &v
}

func (o MerchantInvoiceEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantInvoiceEditPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableMerchantInvoiceEditPost200ResponseData struct {
	value *MerchantInvoiceEditPost200ResponseData
	isSet bool
}

func (v NullableMerchantInvoiceEditPost200ResponseData) Get() *MerchantInvoiceEditPost200ResponseData {
	return v.value
}

func (v *NullableMerchantInvoiceEditPost200ResponseData) Set(val *MerchantInvoiceEditPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantInvoiceEditPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantInvoiceEditPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantInvoiceEditPost200ResponseData(val *MerchantInvoiceEditPost200ResponseData) *NullableMerchantInvoiceEditPost200ResponseData {
	return &NullableMerchantInvoiceEditPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantInvoiceEditPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantInvoiceEditPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


