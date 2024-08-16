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

// checks if the UnibeeApiMerchantInvoiceRefundRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceRefundRes{}

// UnibeeApiMerchantInvoiceRefundRes struct for UnibeeApiMerchantInvoiceRefundRes
type UnibeeApiMerchantInvoiceRefundRes struct {
	Refund *UnibeeApiBeanRefund `json:"refund,omitempty"`
}

// NewUnibeeApiMerchantInvoiceRefundRes instantiates a new UnibeeApiMerchantInvoiceRefundRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceRefundRes() *UnibeeApiMerchantInvoiceRefundRes {
	this := UnibeeApiMerchantInvoiceRefundRes{}
	return &this
}

// NewUnibeeApiMerchantInvoiceRefundResWithDefaults instantiates a new UnibeeApiMerchantInvoiceRefundRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceRefundResWithDefaults() *UnibeeApiMerchantInvoiceRefundRes {
	this := UnibeeApiMerchantInvoiceRefundRes{}
	return &this
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceRefundRes) GetRefund() UnibeeApiBeanRefund {
	if o == nil || IsNil(o.Refund) {
		var ret UnibeeApiBeanRefund
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceRefundRes) GetRefundOk() (*UnibeeApiBeanRefund, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceRefundRes) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given UnibeeApiBeanRefund and assigns it to the Refund field.
func (o *UnibeeApiMerchantInvoiceRefundRes) SetRefund(v UnibeeApiBeanRefund) {
	o.Refund = &v
}

func (o UnibeeApiMerchantInvoiceRefundRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceRefundRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceRefundRes struct {
	value *UnibeeApiMerchantInvoiceRefundRes
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceRefundRes) Get() *UnibeeApiMerchantInvoiceRefundRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceRefundRes) Set(val *UnibeeApiMerchantInvoiceRefundRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceRefundRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceRefundRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceRefundRes(val *UnibeeApiMerchantInvoiceRefundRes) *NullableUnibeeApiMerchantInvoiceRefundRes {
	return &NullableUnibeeApiMerchantInvoiceRefundRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceRefundRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceRefundRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


