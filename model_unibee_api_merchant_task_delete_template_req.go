/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantTaskDeleteTemplateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskDeleteTemplateReq{}

// UnibeeApiMerchantTaskDeleteTemplateReq struct for UnibeeApiMerchantTaskDeleteTemplateReq
type UnibeeApiMerchantTaskDeleteTemplateReq struct {
	// templateId
	TemplateId int64 `json:"templateId"`
}

type _UnibeeApiMerchantTaskDeleteTemplateReq UnibeeApiMerchantTaskDeleteTemplateReq

// NewUnibeeApiMerchantTaskDeleteTemplateReq instantiates a new UnibeeApiMerchantTaskDeleteTemplateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskDeleteTemplateReq(templateId int64) *UnibeeApiMerchantTaskDeleteTemplateReq {
	this := UnibeeApiMerchantTaskDeleteTemplateReq{}
	this.TemplateId = templateId
	return &this
}

// NewUnibeeApiMerchantTaskDeleteTemplateReqWithDefaults instantiates a new UnibeeApiMerchantTaskDeleteTemplateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskDeleteTemplateReqWithDefaults() *UnibeeApiMerchantTaskDeleteTemplateReq {
	this := UnibeeApiMerchantTaskDeleteTemplateReq{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *UnibeeApiMerchantTaskDeleteTemplateReq) GetTemplateId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskDeleteTemplateReq) GetTemplateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *UnibeeApiMerchantTaskDeleteTemplateReq) SetTemplateId(v int64) {
	o.TemplateId = v
}

func (o UnibeeApiMerchantTaskDeleteTemplateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskDeleteTemplateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantTaskDeleteTemplateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
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

	varUnibeeApiMerchantTaskDeleteTemplateReq := _UnibeeApiMerchantTaskDeleteTemplateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantTaskDeleteTemplateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantTaskDeleteTemplateReq(varUnibeeApiMerchantTaskDeleteTemplateReq)

	return err
}

type NullableUnibeeApiMerchantTaskDeleteTemplateReq struct {
	value *UnibeeApiMerchantTaskDeleteTemplateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskDeleteTemplateReq) Get() *UnibeeApiMerchantTaskDeleteTemplateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskDeleteTemplateReq) Set(val *UnibeeApiMerchantTaskDeleteTemplateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskDeleteTemplateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskDeleteTemplateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskDeleteTemplateReq(val *UnibeeApiMerchantTaskDeleteTemplateReq) *NullableUnibeeApiMerchantTaskDeleteTemplateReq {
	return &NullableUnibeeApiMerchantTaskDeleteTemplateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskDeleteTemplateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskDeleteTemplateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


