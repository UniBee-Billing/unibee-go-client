/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentDetailGet200ResponseData{}

// MerchantPaymentDetailGet200ResponseData struct for MerchantPaymentDetailGet200ResponseData
type MerchantPaymentDetailGet200ResponseData struct {
	PaymentDetail *UnibeeApiBeanDetailPaymentDetail `json:"paymentDetail,omitempty"`
}

// NewMerchantPaymentDetailGet200ResponseData instantiates a new MerchantPaymentDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentDetailGet200ResponseData() *MerchantPaymentDetailGet200ResponseData {
	this := MerchantPaymentDetailGet200ResponseData{}
	return &this
}

// NewMerchantPaymentDetailGet200ResponseDataWithDefaults instantiates a new MerchantPaymentDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentDetailGet200ResponseDataWithDefaults() *MerchantPaymentDetailGet200ResponseData {
	this := MerchantPaymentDetailGet200ResponseData{}
	return &this
}

// GetPaymentDetail returns the PaymentDetail field value if set, zero value otherwise.
func (o *MerchantPaymentDetailGet200ResponseData) GetPaymentDetail() UnibeeApiBeanDetailPaymentDetail {
	if o == nil || IsNil(o.PaymentDetail) {
		var ret UnibeeApiBeanDetailPaymentDetail
		return ret
	}
	return *o.PaymentDetail
}

// GetPaymentDetailOk returns a tuple with the PaymentDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentDetailGet200ResponseData) GetPaymentDetailOk() (*UnibeeApiBeanDetailPaymentDetail, bool) {
	if o == nil || IsNil(o.PaymentDetail) {
		return nil, false
	}
	return o.PaymentDetail, true
}

// HasPaymentDetail returns a boolean if a field has been set.
func (o *MerchantPaymentDetailGet200ResponseData) HasPaymentDetail() bool {
	if o != nil && !IsNil(o.PaymentDetail) {
		return true
	}

	return false
}

// SetPaymentDetail gets a reference to the given UnibeeApiBeanDetailPaymentDetail and assigns it to the PaymentDetail field.
func (o *MerchantPaymentDetailGet200ResponseData) SetPaymentDetail(v UnibeeApiBeanDetailPaymentDetail) {
	o.PaymentDetail = &v
}

func (o MerchantPaymentDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentDetail) {
		toSerialize["paymentDetail"] = o.PaymentDetail
	}
	return toSerialize, nil
}

type NullableMerchantPaymentDetailGet200ResponseData struct {
	value *MerchantPaymentDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentDetailGet200ResponseData) Get() *MerchantPaymentDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentDetailGet200ResponseData) Set(val *MerchantPaymentDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentDetailGet200ResponseData(val *MerchantPaymentDetailGet200ResponseData) *NullableMerchantPaymentDetailGet200ResponseData {
	return &NullableMerchantPaymentDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


