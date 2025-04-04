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

// checks if the UnibeeApiMerchantPaymentRefundDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentRefundDetailRes{}

// UnibeeApiMerchantPaymentRefundDetailRes struct for UnibeeApiMerchantPaymentRefundDetailRes
type UnibeeApiMerchantPaymentRefundDetailRes struct {
	RefundDetail *UnibeeApiBeanDetailRefundDetail `json:"refundDetail,omitempty"`
}

// NewUnibeeApiMerchantPaymentRefundDetailRes instantiates a new UnibeeApiMerchantPaymentRefundDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentRefundDetailRes() *UnibeeApiMerchantPaymentRefundDetailRes {
	this := UnibeeApiMerchantPaymentRefundDetailRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentRefundDetailResWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentRefundDetailResWithDefaults() *UnibeeApiMerchantPaymentRefundDetailRes {
	this := UnibeeApiMerchantPaymentRefundDetailRes{}
	return &this
}

// GetRefundDetail returns the RefundDetail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundDetailRes) GetRefundDetail() UnibeeApiBeanDetailRefundDetail {
	if o == nil || IsNil(o.RefundDetail) {
		var ret UnibeeApiBeanDetailRefundDetail
		return ret
	}
	return *o.RefundDetail
}

// GetRefundDetailOk returns a tuple with the RefundDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundDetailRes) GetRefundDetailOk() (*UnibeeApiBeanDetailRefundDetail, bool) {
	if o == nil || IsNil(o.RefundDetail) {
		return nil, false
	}
	return o.RefundDetail, true
}

// HasRefundDetail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundDetailRes) HasRefundDetail() bool {
	if o != nil && !IsNil(o.RefundDetail) {
		return true
	}

	return false
}

// SetRefundDetail gets a reference to the given UnibeeApiBeanDetailRefundDetail and assigns it to the RefundDetail field.
func (o *UnibeeApiMerchantPaymentRefundDetailRes) SetRefundDetail(v UnibeeApiBeanDetailRefundDetail) {
	o.RefundDetail = &v
}

func (o UnibeeApiMerchantPaymentRefundDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentRefundDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetail) {
		toSerialize["refundDetail"] = o.RefundDetail
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentRefundDetailRes struct {
	value *UnibeeApiMerchantPaymentRefundDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailRes) Get() *UnibeeApiMerchantPaymentRefundDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailRes) Set(val *UnibeeApiMerchantPaymentRefundDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentRefundDetailRes(val *UnibeeApiMerchantPaymentRefundDetailRes) *NullableUnibeeApiMerchantPaymentRefundDetailRes {
	return &NullableUnibeeApiMerchantPaymentRefundDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentRefundDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentRefundDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


