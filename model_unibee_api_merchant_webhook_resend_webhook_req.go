/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantWebhookResendWebhookReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookResendWebhookReq{}

// UnibeeApiMerchantWebhookResendWebhookReq struct for UnibeeApiMerchantWebhookResendWebhookReq
type UnibeeApiMerchantWebhookResendWebhookReq struct {
	// LogId
	LogId int64 `json:"logId"`
}

type _UnibeeApiMerchantWebhookResendWebhookReq UnibeeApiMerchantWebhookResendWebhookReq

// NewUnibeeApiMerchantWebhookResendWebhookReq instantiates a new UnibeeApiMerchantWebhookResendWebhookReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookResendWebhookReq(logId int64) *UnibeeApiMerchantWebhookResendWebhookReq {
	this := UnibeeApiMerchantWebhookResendWebhookReq{}
	this.LogId = logId
	return &this
}

// NewUnibeeApiMerchantWebhookResendWebhookReqWithDefaults instantiates a new UnibeeApiMerchantWebhookResendWebhookReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookResendWebhookReqWithDefaults() *UnibeeApiMerchantWebhookResendWebhookReq {
	this := UnibeeApiMerchantWebhookResendWebhookReq{}
	return &this
}

// GetLogId returns the LogId field value
func (o *UnibeeApiMerchantWebhookResendWebhookReq) GetLogId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LogId
}

// GetLogIdOk returns a tuple with the LogId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookResendWebhookReq) GetLogIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogId, true
}

// SetLogId sets field value
func (o *UnibeeApiMerchantWebhookResendWebhookReq) SetLogId(v int64) {
	o.LogId = v
}

func (o UnibeeApiMerchantWebhookResendWebhookReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookResendWebhookReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logId"] = o.LogId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantWebhookResendWebhookReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantWebhookResendWebhookReq := _UnibeeApiMerchantWebhookResendWebhookReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantWebhookResendWebhookReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantWebhookResendWebhookReq(varUnibeeApiMerchantWebhookResendWebhookReq)

	return err
}

type NullableUnibeeApiMerchantWebhookResendWebhookReq struct {
	value *UnibeeApiMerchantWebhookResendWebhookReq
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookReq) Get() *UnibeeApiMerchantWebhookResendWebhookReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookReq) Set(val *UnibeeApiMerchantWebhookResendWebhookReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookResendWebhookReq(val *UnibeeApiMerchantWebhookResendWebhookReq) *NullableUnibeeApiMerchantWebhookResendWebhookReq {
	return &NullableUnibeeApiMerchantWebhookResendWebhookReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookResendWebhookReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookResendWebhookReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


