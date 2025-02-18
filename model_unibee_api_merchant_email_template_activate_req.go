/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailTemplateActivateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailTemplateActivateReq{}

// UnibeeApiMerchantEmailTemplateActivateReq Activate the email template
type UnibeeApiMerchantEmailTemplateActivateReq struct {
	// The name of email template
	TemplateName string `json:"templateName"`
}

type _UnibeeApiMerchantEmailTemplateActivateReq UnibeeApiMerchantEmailTemplateActivateReq

// NewUnibeeApiMerchantEmailTemplateActivateReq instantiates a new UnibeeApiMerchantEmailTemplateActivateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailTemplateActivateReq(templateName string) *UnibeeApiMerchantEmailTemplateActivateReq {
	this := UnibeeApiMerchantEmailTemplateActivateReq{}
	this.TemplateName = templateName
	return &this
}

// NewUnibeeApiMerchantEmailTemplateActivateReqWithDefaults instantiates a new UnibeeApiMerchantEmailTemplateActivateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailTemplateActivateReqWithDefaults() *UnibeeApiMerchantEmailTemplateActivateReq {
	this := UnibeeApiMerchantEmailTemplateActivateReq{}
	return &this
}

// GetTemplateName returns the TemplateName field value
func (o *UnibeeApiMerchantEmailTemplateActivateReq) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateActivateReq) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *UnibeeApiMerchantEmailTemplateActivateReq) SetTemplateName(v string) {
	o.TemplateName = v
}

func (o UnibeeApiMerchantEmailTemplateActivateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailTemplateActivateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateName"] = o.TemplateName
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailTemplateActivateReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantEmailTemplateActivateReq := _UnibeeApiMerchantEmailTemplateActivateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailTemplateActivateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailTemplateActivateReq(varUnibeeApiMerchantEmailTemplateActivateReq)

	return err
}

type NullableUnibeeApiMerchantEmailTemplateActivateReq struct {
	value *UnibeeApiMerchantEmailTemplateActivateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailTemplateActivateReq) Get() *UnibeeApiMerchantEmailTemplateActivateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailTemplateActivateReq) Set(val *UnibeeApiMerchantEmailTemplateActivateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailTemplateActivateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailTemplateActivateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailTemplateActivateReq(val *UnibeeApiMerchantEmailTemplateActivateReq) *NullableUnibeeApiMerchantEmailTemplateActivateReq {
	return &NullableUnibeeApiMerchantEmailTemplateActivateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailTemplateActivateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailTemplateActivateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


