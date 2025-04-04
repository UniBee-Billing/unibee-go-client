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

// checks if the UnibeeApiSystemRefundGatewayDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemRefundGatewayDetailRes{}

// UnibeeApiSystemRefundGatewayDetailRes struct for UnibeeApiSystemRefundGatewayDetailRes
type UnibeeApiSystemRefundGatewayDetailRes struct {
	RefundDetail map[string]interface{} `json:"refundDetail,omitempty"`
}

// NewUnibeeApiSystemRefundGatewayDetailRes instantiates a new UnibeeApiSystemRefundGatewayDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemRefundGatewayDetailRes() *UnibeeApiSystemRefundGatewayDetailRes {
	this := UnibeeApiSystemRefundGatewayDetailRes{}
	return &this
}

// NewUnibeeApiSystemRefundGatewayDetailResWithDefaults instantiates a new UnibeeApiSystemRefundGatewayDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemRefundGatewayDetailResWithDefaults() *UnibeeApiSystemRefundGatewayDetailRes {
	this := UnibeeApiSystemRefundGatewayDetailRes{}
	return &this
}

// GetRefundDetail returns the RefundDetail field value if set, zero value otherwise.
func (o *UnibeeApiSystemRefundGatewayDetailRes) GetRefundDetail() map[string]interface{} {
	if o == nil || IsNil(o.RefundDetail) {
		var ret map[string]interface{}
		return ret
	}
	return o.RefundDetail
}

// GetRefundDetailOk returns a tuple with the RefundDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemRefundGatewayDetailRes) GetRefundDetailOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.RefundDetail) {
		return map[string]interface{}{}, false
	}
	return o.RefundDetail, true
}

// HasRefundDetail returns a boolean if a field has been set.
func (o *UnibeeApiSystemRefundGatewayDetailRes) HasRefundDetail() bool {
	if o != nil && !IsNil(o.RefundDetail) {
		return true
	}

	return false
}

// SetRefundDetail gets a reference to the given map[string]interface{} and assigns it to the RefundDetail field.
func (o *UnibeeApiSystemRefundGatewayDetailRes) SetRefundDetail(v map[string]interface{}) {
	o.RefundDetail = v
}

func (o UnibeeApiSystemRefundGatewayDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemRefundGatewayDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetail) {
		toSerialize["refundDetail"] = o.RefundDetail
	}
	return toSerialize, nil
}

type NullableUnibeeApiSystemRefundGatewayDetailRes struct {
	value *UnibeeApiSystemRefundGatewayDetailRes
	isSet bool
}

func (v NullableUnibeeApiSystemRefundGatewayDetailRes) Get() *UnibeeApiSystemRefundGatewayDetailRes {
	return v.value
}

func (v *NullableUnibeeApiSystemRefundGatewayDetailRes) Set(val *UnibeeApiSystemRefundGatewayDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemRefundGatewayDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemRefundGatewayDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemRefundGatewayDetailRes(val *UnibeeApiSystemRefundGatewayDetailRes) *NullableUnibeeApiSystemRefundGatewayDetailRes {
	return &NullableUnibeeApiSystemRefundGatewayDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemRefundGatewayDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemRefundGatewayDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


