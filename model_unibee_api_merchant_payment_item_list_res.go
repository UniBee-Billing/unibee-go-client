/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentItemListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentItemListRes{}

// UnibeeApiMerchantPaymentItemListRes struct for UnibeeApiMerchantPaymentItemListRes
type UnibeeApiMerchantPaymentItemListRes struct {
	// Payment Item Object List
	PaymentItems []UnibeeApiBeanPaymentItem `json:"paymentItems,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantPaymentItemListRes instantiates a new UnibeeApiMerchantPaymentItemListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentItemListRes() *UnibeeApiMerchantPaymentItemListRes {
	this := UnibeeApiMerchantPaymentItemListRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentItemListResWithDefaults instantiates a new UnibeeApiMerchantPaymentItemListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentItemListResWithDefaults() *UnibeeApiMerchantPaymentItemListRes {
	this := UnibeeApiMerchantPaymentItemListRes{}
	return &this
}

// GetPaymentItems returns the PaymentItems field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItemListRes) GetPaymentItems() []UnibeeApiBeanPaymentItem {
	if o == nil || IsNil(o.PaymentItems) {
		var ret []UnibeeApiBeanPaymentItem
		return ret
	}
	return o.PaymentItems
}

// GetPaymentItemsOk returns a tuple with the PaymentItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItemListRes) GetPaymentItemsOk() ([]UnibeeApiBeanPaymentItem, bool) {
	if o == nil || IsNil(o.PaymentItems) {
		return nil, false
	}
	return o.PaymentItems, true
}

// HasPaymentItems returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItemListRes) HasPaymentItems() bool {
	if o != nil && !IsNil(o.PaymentItems) {
		return true
	}

	return false
}

// SetPaymentItems gets a reference to the given []UnibeeApiBeanPaymentItem and assigns it to the PaymentItems field.
func (o *UnibeeApiMerchantPaymentItemListRes) SetPaymentItems(v []UnibeeApiBeanPaymentItem) {
	o.PaymentItems = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentItemListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentItemListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentItemListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantPaymentItemListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantPaymentItemListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentItemListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentItems) {
		toSerialize["paymentItems"] = o.PaymentItems
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentItemListRes struct {
	value *UnibeeApiMerchantPaymentItemListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentItemListRes) Get() *UnibeeApiMerchantPaymentItemListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentItemListRes) Set(val *UnibeeApiMerchantPaymentItemListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentItemListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentItemListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentItemListRes(val *UnibeeApiMerchantPaymentItemListRes) *NullableUnibeeApiMerchantPaymentItemListRes {
	return &NullableUnibeeApiMerchantPaymentItemListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentItemListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentItemListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


