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

// checks if the UnibeeApiMerchantInvoiceDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceDetailRes{}

// UnibeeApiMerchantInvoiceDetailRes struct for UnibeeApiMerchantInvoiceDetailRes
type UnibeeApiMerchantInvoiceDetailRes struct {
	Invoice *UnibeeInternalLogicGatewayRoInvoiceDetailRo `json:"invoice,omitempty"`
}

// NewUnibeeApiMerchantInvoiceDetailRes instantiates a new UnibeeApiMerchantInvoiceDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceDetailRes() *UnibeeApiMerchantInvoiceDetailRes {
	this := UnibeeApiMerchantInvoiceDetailRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceDetailResWithDefaults instantiates a new UnibeeApiMerchantInvoiceDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceDetailResWithDefaults() *UnibeeApiMerchantInvoiceDetailRes {
	this := UnibeeApiMerchantInvoiceDetailRes{}
	return &this
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceDetailRes) GetInvoice() UnibeeInternalLogicGatewayRoInvoiceDetailRo {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeInternalLogicGatewayRoInvoiceDetailRo
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceDetailRes) GetInvoiceOk() (*UnibeeInternalLogicGatewayRoInvoiceDetailRo, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceDetailRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeInternalLogicGatewayRoInvoiceDetailRo and assigns it to the Invoice field.
func (o *UnibeeApiMerchantInvoiceDetailRes) SetInvoice(v UnibeeInternalLogicGatewayRoInvoiceDetailRo) {
	o.Invoice = &v
}

func (o UnibeeApiMerchantInvoiceDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceDetailRes struct {
	value *UnibeeApiMerchantInvoiceDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceDetailRes) Get() *UnibeeApiMerchantInvoiceDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceDetailRes) Set(val *UnibeeApiMerchantInvoiceDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceDetailRes(val *UnibeeApiMerchantInvoiceDetailRes) *NullableUnibeeApiMerchantInvoiceDetailRes {
	return &NullableUnibeeApiMerchantInvoiceDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


