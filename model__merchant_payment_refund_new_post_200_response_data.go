/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantPaymentRefundNewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantPaymentRefundNewPost200ResponseData{}

// MerchantPaymentRefundNewPost200ResponseData struct for MerchantPaymentRefundNewPost200ResponseData
type MerchantPaymentRefundNewPost200ResponseData struct {
	// ExternalRefundId
	ExternalRefundId *string `json:"externalRefundId,omitempty"`
	// PaymentId
	PaymentId *string `json:"paymentId,omitempty"`
	// RefundId
	RefundId *string `json:"refundId,omitempty"`
	// Status,10-create|20-success|30-Failed|40-Reverse
	Status *int32 `json:"status,omitempty"`
}

// NewMerchantPaymentRefundNewPost200ResponseData instantiates a new MerchantPaymentRefundNewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantPaymentRefundNewPost200ResponseData() *MerchantPaymentRefundNewPost200ResponseData {
	this := MerchantPaymentRefundNewPost200ResponseData{}
	return &this
}

// NewMerchantPaymentRefundNewPost200ResponseDataWithDefaults instantiates a new MerchantPaymentRefundNewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantPaymentRefundNewPost200ResponseDataWithDefaults() *MerchantPaymentRefundNewPost200ResponseData {
	this := MerchantPaymentRefundNewPost200ResponseData{}
	return &this
}

// GetExternalRefundId returns the ExternalRefundId field value if set, zero value otherwise.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetExternalRefundId() string {
	if o == nil || IsNil(o.ExternalRefundId) {
		var ret string
		return ret
	}
	return *o.ExternalRefundId
}

// GetExternalRefundIdOk returns a tuple with the ExternalRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetExternalRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalRefundId) {
		return nil, false
	}
	return o.ExternalRefundId, true
}

// HasExternalRefundId returns a boolean if a field has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) HasExternalRefundId() bool {
	if o != nil && !IsNil(o.ExternalRefundId) {
		return true
	}

	return false
}

// SetExternalRefundId gets a reference to the given string and assigns it to the ExternalRefundId field.
func (o *MerchantPaymentRefundNewPost200ResponseData) SetExternalRefundId(v string) {
	o.ExternalRefundId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *MerchantPaymentRefundNewPost200ResponseData) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *MerchantPaymentRefundNewPost200ResponseData) SetRefundId(v string) {
	o.RefundId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MerchantPaymentRefundNewPost200ResponseData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *MerchantPaymentRefundNewPost200ResponseData) SetStatus(v int32) {
	o.Status = &v
}

func (o MerchantPaymentRefundNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantPaymentRefundNewPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalRefundId) {
		toSerialize["externalRefundId"] = o.ExternalRefundId
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableMerchantPaymentRefundNewPost200ResponseData struct {
	value *MerchantPaymentRefundNewPost200ResponseData
	isSet bool
}

func (v NullableMerchantPaymentRefundNewPost200ResponseData) Get() *MerchantPaymentRefundNewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantPaymentRefundNewPost200ResponseData) Set(val *MerchantPaymentRefundNewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantPaymentRefundNewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantPaymentRefundNewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantPaymentRefundNewPost200ResponseData(val *MerchantPaymentRefundNewPost200ResponseData) *NullableMerchantPaymentRefundNewPost200ResponseData {
	return &NullableMerchantPaymentRefundNewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantPaymentRefundNewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantPaymentRefundNewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


