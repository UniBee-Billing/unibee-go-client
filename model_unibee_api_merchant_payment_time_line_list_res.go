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

// checks if the UnibeeApiMerchantPaymentTimeLineListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentTimeLineListRes{}

// UnibeeApiMerchantPaymentTimeLineListRes struct for UnibeeApiMerchantPaymentTimeLineListRes
type UnibeeApiMerchantPaymentTimeLineListRes struct {
	// Payment TimeLine Object List
	PaymentTimeLines []UnibeeApiBeanDetailPaymentTimelineDetail `json:"paymentTimeLines,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantPaymentTimeLineListRes instantiates a new UnibeeApiMerchantPaymentTimeLineListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentTimeLineListRes() *UnibeeApiMerchantPaymentTimeLineListRes {
	this := UnibeeApiMerchantPaymentTimeLineListRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentTimeLineListResWithDefaults instantiates a new UnibeeApiMerchantPaymentTimeLineListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentTimeLineListResWithDefaults() *UnibeeApiMerchantPaymentTimeLineListRes {
	this := UnibeeApiMerchantPaymentTimeLineListRes{}
	return &this
}

// GetPaymentTimeLines returns the PaymentTimeLines field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetPaymentTimeLines() []UnibeeApiBeanDetailPaymentTimelineDetail {
	if o == nil || IsNil(o.PaymentTimeLines) {
		var ret []UnibeeApiBeanDetailPaymentTimelineDetail
		return ret
	}
	return o.PaymentTimeLines
}

// GetPaymentTimeLinesOk returns a tuple with the PaymentTimeLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetPaymentTimeLinesOk() ([]UnibeeApiBeanDetailPaymentTimelineDetail, bool) {
	if o == nil || IsNil(o.PaymentTimeLines) {
		return nil, false
	}
	return o.PaymentTimeLines, true
}

// HasPaymentTimeLines returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) HasPaymentTimeLines() bool {
	if o != nil && !IsNil(o.PaymentTimeLines) {
		return true
	}

	return false
}

// SetPaymentTimeLines gets a reference to the given []UnibeeApiBeanDetailPaymentTimelineDetail and assigns it to the PaymentTimeLines field.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) SetPaymentTimeLines(v []UnibeeApiBeanDetailPaymentTimelineDetail) {
	o.PaymentTimeLines = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantPaymentTimeLineListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantPaymentTimeLineListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentTimeLineListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentTimeLines) {
		toSerialize["paymentTimeLines"] = o.PaymentTimeLines
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentTimeLineListRes struct {
	value *UnibeeApiMerchantPaymentTimeLineListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListRes) Get() *UnibeeApiMerchantPaymentTimeLineListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListRes) Set(val *UnibeeApiMerchantPaymentTimeLineListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentTimeLineListRes(val *UnibeeApiMerchantPaymentTimeLineListRes) *NullableUnibeeApiMerchantPaymentTimeLineListRes {
	return &NullableUnibeeApiMerchantPaymentTimeLineListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


