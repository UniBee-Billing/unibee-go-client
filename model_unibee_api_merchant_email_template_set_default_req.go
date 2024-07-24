/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailTemplateSetDefaultReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailTemplateSetDefaultReq{}

// UnibeeApiMerchantEmailTemplateSetDefaultReq Setup email template as default
type UnibeeApiMerchantEmailTemplateSetDefaultReq struct {
	// The name of email template
	TemplateName string `json:"templateName"`
}

type _UnibeeApiMerchantEmailTemplateSetDefaultReq UnibeeApiMerchantEmailTemplateSetDefaultReq

// NewUnibeeApiMerchantEmailTemplateSetDefaultReq instantiates a new UnibeeApiMerchantEmailTemplateSetDefaultReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailTemplateSetDefaultReq(templateName string) *UnibeeApiMerchantEmailTemplateSetDefaultReq {
	this := UnibeeApiMerchantEmailTemplateSetDefaultReq{}
	this.TemplateName = templateName
	return &this
}

// NewUnibeeApiMerchantEmailTemplateSetDefaultReqWithDefaults instantiates a new UnibeeApiMerchantEmailTemplateSetDefaultReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailTemplateSetDefaultReqWithDefaults() *UnibeeApiMerchantEmailTemplateSetDefaultReq {
	this := UnibeeApiMerchantEmailTemplateSetDefaultReq{}
	return &this
}

// GetTemplateName returns the TemplateName field value
func (o *UnibeeApiMerchantEmailTemplateSetDefaultReq) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateSetDefaultReq) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *UnibeeApiMerchantEmailTemplateSetDefaultReq) SetTemplateName(v string) {
	o.TemplateName = v
}

func (o UnibeeApiMerchantEmailTemplateSetDefaultReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailTemplateSetDefaultReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateName"] = o.TemplateName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailTemplateSetDefaultReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateName",
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

	varUnibeeApiMerchantEmailTemplateSetDefaultReq := _UnibeeApiMerchantEmailTemplateSetDefaultReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailTemplateSetDefaultReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailTemplateSetDefaultReq(varUnibeeApiMerchantEmailTemplateSetDefaultReq)

	return err
}

type NullableUnibeeApiMerchantEmailTemplateSetDefaultReq struct {
	value *UnibeeApiMerchantEmailTemplateSetDefaultReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) Get() *UnibeeApiMerchantEmailTemplateSetDefaultReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) Set(val *UnibeeApiMerchantEmailTemplateSetDefaultReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailTemplateSetDefaultReq(val *UnibeeApiMerchantEmailTemplateSetDefaultReq) *NullableUnibeeApiMerchantEmailTemplateSetDefaultReq {
	return &NullableUnibeeApiMerchantEmailTemplateSetDefaultReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailTemplateSetDefaultReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


