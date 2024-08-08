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

// checks if the MerchantPaymentTimelineListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentTimelineListGet200ResponseData{}

// MerchantPaymentTimelineListGet200ResponseData struct for MerchantPaymentTimelineListGet200ResponseData
type MerchantPaymentTimelineListGet200ResponseData struct {
	// Payment TimeLine Object List
	PaymentTimeLines []UnibeeApiBeanDetailPaymentTimelineDetail `json:"paymentTimeLines,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantPaymentTimelineListGet200ResponseData instantiates a new MerchantPaymentTimelineListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentTimelineListGet200ResponseData() *MerchantPaymentTimelineListGet200ResponseData {
	this := MerchantPaymentTimelineListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentTimelineListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentTimelineListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentTimelineListGet200ResponseDataWithDefaults() *MerchantPaymentTimelineListGet200ResponseData {
	this := MerchantPaymentTimelineListGet200ResponseData{}
	return &this
}

// GetPaymentTimeLines returns the PaymentTimeLines field value if set, zero value otherwise.
func (o *MerchantPaymentTimelineListGet200ResponseData) GetPaymentTimeLines() []UnibeeApiBeanDetailPaymentTimelineDetail {
	if o == nil || IsNil(o.PaymentTimeLines) {
		var ret []UnibeeApiBeanDetailPaymentTimelineDetail
		return ret
	}
	return o.PaymentTimeLines
}

// GetPaymentTimeLinesOk returns a tuple with the PaymentTimeLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentTimelineListGet200ResponseData) GetPaymentTimeLinesOk() ([]UnibeeApiBeanDetailPaymentTimelineDetail, bool) {
	if o == nil || IsNil(o.PaymentTimeLines) {
		return nil, false
	}
	return o.PaymentTimeLines, true
}

// HasPaymentTimeLines returns a boolean if a field has been set.
func (o *MerchantPaymentTimelineListGet200ResponseData) HasPaymentTimeLines() bool {
	if o != nil && !IsNil(o.PaymentTimeLines) {
		return true
	}

	return false
}

// SetPaymentTimeLines gets a reference to the given []UnibeeApiBeanDetailPaymentTimelineDetail and assigns it to the PaymentTimeLines field.
func (o *MerchantPaymentTimelineListGet200ResponseData) SetPaymentTimeLines(v []UnibeeApiBeanDetailPaymentTimelineDetail) {
	o.PaymentTimeLines = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantPaymentTimelineListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentTimelineListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantPaymentTimelineListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantPaymentTimelineListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantPaymentTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentTimelineListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentTimeLines) {
		toSerialize["paymentTimeLines"] = o.PaymentTimeLines
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantPaymentTimelineListGet200ResponseData struct {
	value *MerchantPaymentTimelineListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentTimelineListGet200ResponseData) Get() *MerchantPaymentTimelineListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentTimelineListGet200ResponseData) Set(val *MerchantPaymentTimelineListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentTimelineListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentTimelineListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentTimelineListGet200ResponseData(val *MerchantPaymentTimelineListGet200ResponseData) *NullableMerchantPaymentTimelineListGet200ResponseData {
	return &NullableMerchantPaymentTimelineListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentTimelineListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


