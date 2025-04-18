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

// checks if the MerchantInvoiceDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantInvoiceDetailGet200ResponseData{}

// MerchantInvoiceDetailGet200ResponseData struct for MerchantInvoiceDetailGet200ResponseData
type MerchantInvoiceDetailGet200ResponseData struct {
	// CreditNotes Object List Link To Invoice
	CreditNotes []UnibeeApiBeanDetailInvoiceDetail `json:"creditNotes,omitempty"`
	Invoice *UnibeeApiBeanDetailInvoiceDetail `json:"invoice,omitempty"`
}

// NewMerchantInvoiceDetailGet200ResponseData instantiates a new MerchantInvoiceDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantInvoiceDetailGet200ResponseData() *MerchantInvoiceDetailGet200ResponseData {
	this := MerchantInvoiceDetailGet200ResponseData{}
	return &this
}

// NewMerchantInvoiceDetailGet200ResponseDataWithDefaults instantiates a new MerchantInvoiceDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantInvoiceDetailGet200ResponseDataWithDefaults() *MerchantInvoiceDetailGet200ResponseData {
	this := MerchantInvoiceDetailGet200ResponseData{}
	return &this
}

// GetCreditNotes returns the CreditNotes field value if set, zero value otherwise.
func (o *MerchantInvoiceDetailGet200ResponseData) GetCreditNotes() []UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.CreditNotes) {
		var ret []UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return o.CreditNotes
}

// GetCreditNotesOk returns a tuple with the CreditNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceDetailGet200ResponseData) GetCreditNotesOk() ([]UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.CreditNotes) {
		return nil, false
	}
	return o.CreditNotes, true
}

// HasCreditNotes returns a boolean if a field has been set.
func (o *MerchantInvoiceDetailGet200ResponseData) HasCreditNotes() bool {
	if o != nil && !IsNil(o.CreditNotes) {
		return true
	}

	return false
}

// SetCreditNotes gets a reference to the given []UnibeeApiBeanDetailInvoiceDetail and assigns it to the CreditNotes field.
func (o *MerchantInvoiceDetailGet200ResponseData) SetCreditNotes(v []UnibeeApiBeanDetailInvoiceDetail) {
	o.CreditNotes = v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantInvoiceDetailGet200ResponseData) GetInvoice() UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantInvoiceDetailGet200ResponseData) GetInvoiceOk() (*UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantInvoiceDetailGet200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanDetailInvoiceDetail and assigns it to the Invoice field.
func (o *MerchantInvoiceDetailGet200ResponseData) SetInvoice(v UnibeeApiBeanDetailInvoiceDetail) {
	o.Invoice = &v
}

func (o MerchantInvoiceDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantInvoiceDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditNotes) {
		toSerialize["creditNotes"] = o.CreditNotes
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableMerchantInvoiceDetailGet200ResponseData struct {
	value *MerchantInvoiceDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantInvoiceDetailGet200ResponseData) Get() *MerchantInvoiceDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantInvoiceDetailGet200ResponseData) Set(val *MerchantInvoiceDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantInvoiceDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantInvoiceDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantInvoiceDetailGet200ResponseData(val *MerchantInvoiceDetailGet200ResponseData) *NullableMerchantInvoiceDetailGet200ResponseData {
	return &NullableMerchantInvoiceDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantInvoiceDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantInvoiceDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


