/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentListRes{}

// UnibeeApiMerchantPaymentListRes struct for UnibeeApiMerchantPaymentListRes
type UnibeeApiMerchantPaymentListRes struct {
	// Payment Detail Object List
	PaymentDetails []UnibeeApiBeanDetailPaymentDetail `json:"paymentDetails,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantPaymentListRes instantiates a new UnibeeApiMerchantPaymentListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentListRes() *UnibeeApiMerchantPaymentListRes {
	this := UnibeeApiMerchantPaymentListRes{}
	return &this
}

// NewUnibeeApiMerchantPaymentListResWithDefaults instantiates a new UnibeeApiMerchantPaymentListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentListResWithDefaults() *UnibeeApiMerchantPaymentListRes {
	this := UnibeeApiMerchantPaymentListRes{}
	return &this
}

// GetPaymentDetails returns the PaymentDetails field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentListRes) GetPaymentDetails() []UnibeeApiBeanDetailPaymentDetail {
	if o == nil || IsNil(o.PaymentDetails) {
		var ret []UnibeeApiBeanDetailPaymentDetail
		return ret
	}
	return o.PaymentDetails
}

// GetPaymentDetailsOk returns a tuple with the PaymentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentListRes) GetPaymentDetailsOk() ([]UnibeeApiBeanDetailPaymentDetail, bool) {
	if o == nil || IsNil(o.PaymentDetails) {
		return nil, false
	}
	return o.PaymentDetails, true
}

// HasPaymentDetails returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentListRes) HasPaymentDetails() bool {
	if o != nil && !IsNil(o.PaymentDetails) {
		return true
	}

	return false
}

// SetPaymentDetails gets a reference to the given []UnibeeApiBeanDetailPaymentDetail and assigns it to the PaymentDetails field.
func (o *UnibeeApiMerchantPaymentListRes) SetPaymentDetails(v []UnibeeApiBeanDetailPaymentDetail) {
	o.PaymentDetails = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantPaymentListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantPaymentListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentDetails) {
		toSerialize["paymentDetails"] = o.PaymentDetails
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentListRes struct {
	value *UnibeeApiMerchantPaymentListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentListRes) Get() *UnibeeApiMerchantPaymentListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentListRes) Set(val *UnibeeApiMerchantPaymentListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentListRes(val *UnibeeApiMerchantPaymentListRes) *NullableUnibeeApiMerchantPaymentListRes {
	return &NullableUnibeeApiMerchantPaymentListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


