/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceListRes{}

// UnibeeApiMerchantInvoiceListRes struct for UnibeeApiMerchantInvoiceListRes
type UnibeeApiMerchantInvoiceListRes struct {
	// Invoice Detail Object List
	Invoices []UnibeeApiBeanDetailInvoiceDetail `json:"invoices,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantInvoiceListRes instantiates a new UnibeeApiMerchantInvoiceListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceListRes() *UnibeeApiMerchantInvoiceListRes {
	this := UnibeeApiMerchantInvoiceListRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceListResWithDefaults instantiates a new UnibeeApiMerchantInvoiceListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceListResWithDefaults() *UnibeeApiMerchantInvoiceListRes {
	this := UnibeeApiMerchantInvoiceListRes{}
	return &this
}

// GetInvoices returns the Invoices field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListRes) GetInvoices() []UnibeeApiBeanDetailInvoiceDetail {
	if o == nil || IsNil(o.Invoices) {
		var ret []UnibeeApiBeanDetailInvoiceDetail
		return ret
	}
	return o.Invoices
}

// GetInvoicesOk returns a tuple with the Invoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListRes) GetInvoicesOk() ([]UnibeeApiBeanDetailInvoiceDetail, bool) {
	if o == nil || IsNil(o.Invoices) {
		return nil, false
	}
	return o.Invoices, true
}

// HasInvoices returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListRes) HasInvoices() bool {
	if o != nil && !IsNil(o.Invoices) {
		return true
	}

	return false
}

// SetInvoices gets a reference to the given []UnibeeApiBeanDetailInvoiceDetail and assigns it to the Invoices field.
func (o *UnibeeApiMerchantInvoiceListRes) SetInvoices(v []UnibeeApiBeanDetailInvoiceDetail) {
	o.Invoices = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantInvoiceListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantInvoiceListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoices) {
		toSerialize["invoices"] = o.Invoices
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceListRes struct {
	value *UnibeeApiMerchantInvoiceListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceListRes) Get() *UnibeeApiMerchantInvoiceListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceListRes) Set(val *UnibeeApiMerchantInvoiceListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceListRes(val *UnibeeApiMerchantInvoiceListRes) *NullableUnibeeApiMerchantInvoiceListRes {
	return &NullableUnibeeApiMerchantInvoiceListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


