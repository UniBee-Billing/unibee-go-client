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

// checks if the MerchantWebhookResendPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantWebhookResendPost200ResponseData{}

// MerchantWebhookResendPost200ResponseData struct for MerchantWebhookResendPost200ResponseData
type MerchantWebhookResendPost200ResponseData struct {
	// SendResult
	SendResult *bool `json:"sendResult,omitempty"`
}

// NewMerchantWebhookResendPost200ResponseData instantiates a new MerchantWebhookResendPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantWebhookResendPost200ResponseData() *MerchantWebhookResendPost200ResponseData {
	this := MerchantWebhookResendPost200ResponseData{}
	return &this
}

// NewMerchantWebhookResendPost200ResponseDataWithDefaults instantiates a new MerchantWebhookResendPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWebhookResendPost200ResponseDataWithDefaults() *MerchantWebhookResendPost200ResponseData {
	this := MerchantWebhookResendPost200ResponseData{}
	return &this
}

// GetSendResult returns the SendResult field value if set, zero value otherwise.
func (o *MerchantWebhookResendPost200ResponseData) GetSendResult() bool {
	if o == nil || IsNil(o.SendResult) {
		var ret bool
		return ret
	}
	return *o.SendResult
}

// GetSendResultOk returns a tuple with the SendResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookResendPost200ResponseData) GetSendResultOk() (*bool, bool) {
	if o == nil || IsNil(o.SendResult) {
		return nil, false
	}
	return o.SendResult, true
}

// HasSendResult returns a boolean if a field has been set.
func (o *MerchantWebhookResendPost200ResponseData) HasSendResult() bool {
	if o != nil && !IsNil(o.SendResult) {
		return true
	}

	return false
}

// SetSendResult gets a reference to the given bool and assigns it to the SendResult field.
func (o *MerchantWebhookResendPost200ResponseData) SetSendResult(v bool) {
	o.SendResult = &v
}

func (o MerchantWebhookResendPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantWebhookResendPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SendResult) {
		toSerialize["sendResult"] = o.SendResult
	}
	return toSerialize, nil
}

type NullableMerchantWebhookResendPost200ResponseData struct {
	value *MerchantWebhookResendPost200ResponseData
	isSet bool
}

func (v NullableMerchantWebhookResendPost200ResponseData) Get() *MerchantWebhookResendPost200ResponseData {
	return v.value
}

func (v *NullableMerchantWebhookResendPost200ResponseData) Set(val *MerchantWebhookResendPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantWebhookResendPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantWebhookResendPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantWebhookResendPost200ResponseData(val *MerchantWebhookResendPost200ResponseData) *NullableMerchantWebhookResendPost200ResponseData {
	return &NullableMerchantWebhookResendPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantWebhookResendPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantWebhookResendPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


