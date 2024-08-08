/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentRefundDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentRefundDetailGet200ResponseData{}

// MerchantPaymentRefundDetailGet200ResponseData struct for MerchantPaymentRefundDetailGet200ResponseData
type MerchantPaymentRefundDetailGet200ResponseData struct {
	RefundDetail *UnibeeApiBeanDetailRefundDetail `json:"refundDetail,omitempty"`
}

// NewMerchantPaymentRefundDetailGet200ResponseData instantiates a new MerchantPaymentRefundDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentRefundDetailGet200ResponseData() *MerchantPaymentRefundDetailGet200ResponseData {
	this := MerchantPaymentRefundDetailGet200ResponseData{}
	return &this
}

// NewMerchantPaymentRefundDetailGet200ResponseDataWithDefaults instantiates a new MerchantPaymentRefundDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentRefundDetailGet200ResponseDataWithDefaults() *MerchantPaymentRefundDetailGet200ResponseData {
	this := MerchantPaymentRefundDetailGet200ResponseData{}
	return &this
}

// GetRefundDetail returns the RefundDetail field value if set, zero value otherwise.
func (o *MerchantPaymentRefundDetailGet200ResponseData) GetRefundDetail() UnibeeApiBeanDetailRefundDetail {
	if o == nil || IsNil(o.RefundDetail) {
		var ret UnibeeApiBeanDetailRefundDetail
		return ret
	}
	return *o.RefundDetail
}

// GetRefundDetailOk returns a tuple with the RefundDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundDetailGet200ResponseData) GetRefundDetailOk() (*UnibeeApiBeanDetailRefundDetail, bool) {
	if o == nil || IsNil(o.RefundDetail) {
		return nil, false
	}
	return o.RefundDetail, true
}

// HasRefundDetail returns a boolean if a field has been set.
func (o *MerchantPaymentRefundDetailGet200ResponseData) HasRefundDetail() bool {
	if o != nil && !IsNil(o.RefundDetail) {
		return true
	}

	return false
}

// SetRefundDetail gets a reference to the given UnibeeApiBeanDetailRefundDetail and assigns it to the RefundDetail field.
func (o *MerchantPaymentRefundDetailGet200ResponseData) SetRefundDetail(v UnibeeApiBeanDetailRefundDetail) {
	o.RefundDetail = &v
}

func (o MerchantPaymentRefundDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentRefundDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetail) {
		toSerialize["refundDetail"] = o.RefundDetail
	}
	return toSerialize, nil
}

type NullableMerchantPaymentRefundDetailGet200ResponseData struct {
	value *MerchantPaymentRefundDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentRefundDetailGet200ResponseData) Get() *MerchantPaymentRefundDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentRefundDetailGet200ResponseData) Set(val *MerchantPaymentRefundDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentRefundDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentRefundDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentRefundDetailGet200ResponseData(val *MerchantPaymentRefundDetailGet200ResponseData) *NullableMerchantPaymentRefundDetailGet200ResponseData {
	return &NullableMerchantPaymentRefundDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentRefundDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentRefundDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


