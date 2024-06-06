/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202406061803 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentRefundListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentRefundListRes{}

// UnibeeApiMerchantPaymentRefundListRes struct for UnibeeApiMerchantPaymentRefundListRes
type UnibeeApiMerchantPaymentRefundListRes struct {
	// Refund Detail Object List
	RefundDetails []UnibeeApiBeanDetailRefundDetail `json:"refundDetails,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantPaymentRefundListRes instantiates a new UnibeeApiMerchantPaymentRefundListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentRefundListRes() *UnibeeApiMerchantPaymentRefundListRes {
	this := UnibeeApiMerchantPaymentRefundListRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentRefundListResWithDefaults instantiates a new UnibeeApiMerchantPaymentRefundListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentRefundListResWithDefaults() *UnibeeApiMerchantPaymentRefundListRes {
	this := UnibeeApiMerchantPaymentRefundListRes{}
	return &this
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListRes) GetRefundDetails() []UnibeeApiBeanDetailRefundDetail {
	if o == nil || IsNil(o.RefundDetails) {
		var ret []UnibeeApiBeanDetailRefundDetail
		return ret
	}
	return o.RefundDetails
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListRes) GetRefundDetailsOk() ([]UnibeeApiBeanDetailRefundDetail, bool) {
	if o == nil || IsNil(o.RefundDetails) {
		return nil, false
	}
	return o.RefundDetails, true
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListRes) HasRefundDetails() bool {
	if o != nil && !IsNil(o.RefundDetails) {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given []UnibeeApiBeanDetailRefundDetail and assigns it to the RefundDetails field.
func (o *UnibeeApiMerchantPaymentRefundListRes) SetRefundDetails(v []UnibeeApiBeanDetailRefundDetail) {
	o.RefundDetails = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentRefundListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentRefundListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentRefundListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantPaymentRefundListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantPaymentRefundListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentRefundListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetails) {
		toSerialize["refundDetails"] = o.RefundDetails
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentRefundListRes struct {
	value *UnibeeApiMerchantPaymentRefundListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentRefundListRes) Get() *UnibeeApiMerchantPaymentRefundListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentRefundListRes) Set(val *UnibeeApiMerchantPaymentRefundListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentRefundListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentRefundListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentRefundListRes(val *UnibeeApiMerchantPaymentRefundListRes) *NullableUnibeeApiMerchantPaymentRefundListRes {
	return &NullableUnibeeApiMerchantPaymentRefundListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentRefundListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentRefundListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


