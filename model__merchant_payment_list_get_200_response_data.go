/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentListGet200ResponseData{}

// MerchantPaymentListGet200ResponseData struct for MerchantPaymentListGet200ResponseData
type MerchantPaymentListGet200ResponseData struct {
	// Payment Detail Object List
	PaymentDetails []UnibeeApiBeanDetailPaymentDetail `json:"paymentDetails,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantPaymentListGet200ResponseData instantiates a new MerchantPaymentListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentListGet200ResponseData() *MerchantPaymentListGet200ResponseData {
	this := MerchantPaymentListGet200ResponseData{}
	return &this
}

// NewMerchantPaymentListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentListGet200ResponseDataWithDefaults() *MerchantPaymentListGet200ResponseData {
	this := MerchantPaymentListGet200ResponseData{}
	return &this
}

// GetPaymentDetails returns the PaymentDetails field value if set, zero value otherwise.
func (o *MerchantPaymentListGet200ResponseData) GetPaymentDetails() []UnibeeApiBeanDetailPaymentDetail {
	if o == nil || IsNil(o.PaymentDetails) {
		var ret []UnibeeApiBeanDetailPaymentDetail
		return ret
	}
	return o.PaymentDetails
}

// GetPaymentDetailsOk returns a tuple with the PaymentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentListGet200ResponseData) GetPaymentDetailsOk() ([]UnibeeApiBeanDetailPaymentDetail, bool) {
	if o == nil || IsNil(o.PaymentDetails) {
		return nil, false
	}
	return o.PaymentDetails, true
}

// HasPaymentDetails returns a boolean if a field has been set.
func (o *MerchantPaymentListGet200ResponseData) HasPaymentDetails() bool {
	if o != nil && !IsNil(o.PaymentDetails) {
		return true
	}

	return false
}

// SetPaymentDetails gets a reference to the given []UnibeeApiBeanDetailPaymentDetail and assigns it to the PaymentDetails field.
func (o *MerchantPaymentListGet200ResponseData) SetPaymentDetails(v []UnibeeApiBeanDetailPaymentDetail) {
	o.PaymentDetails = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantPaymentListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantPaymentListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantPaymentListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantPaymentListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentDetails) {
		toSerialize["paymentDetails"] = o.PaymentDetails
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantPaymentListGet200ResponseData struct {
	value *MerchantPaymentListGet200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentListGet200ResponseData) Get() *MerchantPaymentListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentListGet200ResponseData) Set(val *MerchantPaymentListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentListGet200ResponseData(val *MerchantPaymentListGet200ResponseData) *NullableMerchantPaymentListGet200ResponseData {
	return &NullableMerchantPaymentListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


