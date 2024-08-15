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

// checks if the UnibeeApiMerchantWebhookResendWebhookRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookResendWebhookRes{}

// UnibeeApiMerchantWebhookResendWebhookRes struct for UnibeeApiMerchantWebhookResendWebhookRes
type UnibeeApiMerchantWebhookResendWebhookRes struct {
	// SendResult
	SendResult *bool `json:"sendResult,omitempty"`
}

// NewUnibeeApiMerchantWebhookResendWebhookRes instantiates a new UnibeeApiMerchantWebhookResendWebhookRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookResendWebhookRes() *UnibeeApiMerchantWebhookResendWebhookRes {
	this := UnibeeApiMerchantWebhookResendWebhookRes{}
	return &this
}

// NewUnibeeApiMerchantWebhookResendWebhookResWithDefaults instantiates a new UnibeeApiMerchantWebhookResendWebhookRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookResendWebhookResWithDefaults() *UnibeeApiMerchantWebhookResendWebhookRes {
	this := UnibeeApiMerchantWebhookResendWebhookRes{}
	return &this
}

// GetSendResult returns the SendResult field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookResendWebhookRes) GetSendResult() bool {
	if o == nil || IsNil(o.SendResult) {
		var ret bool
		return ret
	}
	return *o.SendResult
}

// GetSendResultOk returns a tuple with the SendResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookResendWebhookRes) GetSendResultOk() (*bool, bool) {
	if o == nil || IsNil(o.SendResult) {
		return nil, false
	}
	return o.SendResult, true
}

// HasSendResult returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookResendWebhookRes) HasSendResult() bool {
	if o != nil && !IsNil(o.SendResult) {
		return true
	}

	return false
}

// SetSendResult gets a reference to the given bool and assigns it to the SendResult field.
func (o *UnibeeApiMerchantWebhookResendWebhookRes) SetSendResult(v bool) {
	o.SendResult = &v
}

func (o UnibeeApiMerchantWebhookResendWebhookRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookResendWebhookRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SendResult) {
		toSerialize["sendResult"] = o.SendResult
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantWebhookResendWebhookRes struct {
	value *UnibeeApiMerchantWebhookResendWebhookRes
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookRes) Get() *UnibeeApiMerchantWebhookResendWebhookRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookRes) Set(val *UnibeeApiMerchantWebhookResendWebhookRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookResendWebhookRes(val *UnibeeApiMerchantWebhookResendWebhookRes) *NullableUnibeeApiMerchantWebhookResendWebhookRes {
	return &NullableUnibeeApiMerchantWebhookResendWebhookRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


