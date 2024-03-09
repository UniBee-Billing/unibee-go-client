/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentDetailRes{}

// UnibeeApiMerchantPaymentDetailRes struct for UnibeeApiMerchantPaymentDetailRes
type UnibeeApiMerchantPaymentDetailRes struct {
	PaymentDetail *UnibeeInternalLogicGatewayRoPaymentDetailRo `json:"paymentDetail,omitempty"`
}

// NewUnibeeApiMerchantPaymentDetailRes instantiates a new UnibeeApiMerchantPaymentDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentDetailRes() *UnibeeApiMerchantPaymentDetailRes {
	this := UnibeeApiMerchantPaymentDetailRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentDetailResWithDefaults instantiates a new UnibeeApiMerchantPaymentDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentDetailResWithDefaults() *UnibeeApiMerchantPaymentDetailRes {
	this := UnibeeApiMerchantPaymentDetailRes{}
	return &this
}

// GetPaymentDetail returns the PaymentDetail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentDetailRes) GetPaymentDetail() UnibeeInternalLogicGatewayRoPaymentDetailRo {
	if o == nil || IsNil(o.PaymentDetail) {
		var ret UnibeeInternalLogicGatewayRoPaymentDetailRo
		return ret
	}
	return *o.PaymentDetail
}

// GetPaymentDetailOk returns a tuple with the PaymentDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentDetailRes) GetPaymentDetailOk() (*UnibeeInternalLogicGatewayRoPaymentDetailRo, bool) {
	if o == nil || IsNil(o.PaymentDetail) {
		return nil, false
	}
	return o.PaymentDetail, true
}

// HasPaymentDetail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentDetailRes) HasPaymentDetail() bool {
	if o != nil && !IsNil(o.PaymentDetail) {
		return true
	}

	return false
}

// SetPaymentDetail gets a reference to the given UnibeeInternalLogicGatewayRoPaymentDetailRo and assigns it to the PaymentDetail field.
func (o *UnibeeApiMerchantPaymentDetailRes) SetPaymentDetail(v UnibeeInternalLogicGatewayRoPaymentDetailRo) {
	o.PaymentDetail = &v
}

func (o UnibeeApiMerchantPaymentDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentDetail) {
		toSerialize["paymentDetail"] = o.PaymentDetail
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentDetailRes struct {
	value *UnibeeApiMerchantPaymentDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentDetailRes) Get() *UnibeeApiMerchantPaymentDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentDetailRes) Set(val *UnibeeApiMerchantPaymentDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentDetailRes(val *UnibeeApiMerchantPaymentDetailRes) *NullableUnibeeApiMerchantPaymentDetailRes {
	return &NullableUnibeeApiMerchantPaymentDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


