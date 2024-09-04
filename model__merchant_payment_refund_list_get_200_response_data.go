/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentRefundListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentRefundListGet200ResponseData{}

// MerchantPaymentRefundListGet200ResponseData struct for MerchantPaymentRefundListGet200ResponseData
type MerchantPaymentRefundListGet200ResponseData struct {
	// Refund Detail Object List
	RefundDetails []UnibeeApiBeanDetailRefundDetail `json:"refundDetails,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantPaymentRefundListGet200ResponseData instantiates a new MerchantPaymentRefundListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentRefundListGet200ResponseData() *MerchantPaymentRefundListGet200ResponseData {
	this := MerchantPaymentRefundListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentRefundListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentRefundListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentRefundListGet200ResponseDataWithDefaults() *MerchantPaymentRefundListGet200ResponseData {
	this := MerchantPaymentRefundListGet200ResponseData{}
	return &this
}

// GetRefundDetails returns the RefundDetails field value if set, zero value otherwise.
func (o *MerchantPaymentRefundListGet200ResponseData) GetRefundDetails() []UnibeeApiBeanDetailRefundDetail {
	if o == nil || IsNil(o.RefundDetails) {
		var ret []UnibeeApiBeanDetailRefundDetail
		return ret
	}
	return o.RefundDetails
}

// GetRefundDetailsOk returns a tuple with the RefundDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundListGet200ResponseData) GetRefundDetailsOk() ([]UnibeeApiBeanDetailRefundDetail, bool) {
	if o == nil || IsNil(o.RefundDetails) {
		return nil, false
	}
	return o.RefundDetails, true
}

// HasRefundDetails returns a boolean if a field has been set.
func (o *MerchantPaymentRefundListGet200ResponseData) HasRefundDetails() bool {
	if o != nil && !IsNil(o.RefundDetails) {
		return true
	}

	return false
}

// SetRefundDetails gets a reference to the given []UnibeeApiBeanDetailRefundDetail and assigns it to the RefundDetails field.
func (o *MerchantPaymentRefundListGet200ResponseData) SetRefundDetails(v []UnibeeApiBeanDetailRefundDetail) {
	o.RefundDetails = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantPaymentRefundListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantPaymentRefundListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantPaymentRefundListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantPaymentRefundListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentRefundListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RefundDetails) {
		toSerialize["refundDetails"] = o.RefundDetails
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantPaymentRefundListGet200ResponseData struct {
	value *MerchantPaymentRefundListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentRefundListGet200ResponseData) Get() *MerchantPaymentRefundListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentRefundListGet200ResponseData) Set(val *MerchantPaymentRefundListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentRefundListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentRefundListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentRefundListGet200ResponseData(val *MerchantPaymentRefundListGet200ResponseData) *NullableMerchantPaymentRefundListGet200ResponseData {
	return &NullableMerchantPaymentRefundListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentRefundListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentRefundListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


