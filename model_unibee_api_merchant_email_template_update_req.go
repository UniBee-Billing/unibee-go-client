/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailTemplateUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailTemplateUpdateReq{}

// UnibeeApiMerchantEmailTemplateUpdateReq Update the email template
type UnibeeApiMerchantEmailTemplateUpdateReq struct {
	// The content of email template
	TemplateContent string `json:"templateContent"`
	// The name of email template
	TemplateName string `json:"templateName"`
	// The title of email template
	TemplateTitle string `json:"templateTitle"`
}

type _UnibeeApiMerchantEmailTemplateUpdateReq UnibeeApiMerchantEmailTemplateUpdateReq

// NewUnibeeApiMerchantEmailTemplateUpdateReq instantiates a new UnibeeApiMerchantEmailTemplateUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailTemplateUpdateReq(templateContent string, templateName string, templateTitle string) *UnibeeApiMerchantEmailTemplateUpdateReq {
	this := UnibeeApiMerchantEmailTemplateUpdateReq{}
	this.TemplateContent = templateContent
	this.TemplateName = templateName
	this.TemplateTitle = templateTitle
	return &this
}

// NewUnibeeApiMerchantEmailTemplateUpdateReqWithDefaults instantiates a new UnibeeApiMerchantEmailTemplateUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailTemplateUpdateReqWithDefaults() *UnibeeApiMerchantEmailTemplateUpdateReq {
	this := UnibeeApiMerchantEmailTemplateUpdateReq{}
	return &this
}

// GetTemplateContent returns the TemplateContent field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateContent
}

// GetTemplateContentOk returns a tuple with the TemplateContent field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateContent, true
}

// SetTemplateContent sets field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) SetTemplateContent(v string) {
	o.TemplateContent = v
}

// GetTemplateName returns the TemplateName field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetTemplateTitle returns the TemplateTitle field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateTitle
}

// GetTemplateTitleOk returns a tuple with the TemplateTitle field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) GetTemplateTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateTitle, true
}

// SetTemplateTitle sets field value
func (o *UnibeeApiMerchantEmailTemplateUpdateReq) SetTemplateTitle(v string) {
	o.TemplateTitle = v
}

func (o UnibeeApiMerchantEmailTemplateUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailTemplateUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateContent"] = o.TemplateContent
	toSerialize["templateName"] = o.TemplateName
	toSerialize["templateTitle"] = o.TemplateTitle
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailTemplateUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateContent",
		"templateName",
		"templateTitle",
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

	varUnibeeApiMerchantEmailTemplateUpdateReq := _UnibeeApiMerchantEmailTemplateUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailTemplateUpdateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailTemplateUpdateReq(varUnibeeApiMerchantEmailTemplateUpdateReq)

	return err
}

type NullableUnibeeApiMerchantEmailTemplateUpdateReq struct {
	value *UnibeeApiMerchantEmailTemplateUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailTemplateUpdateReq) Get() *UnibeeApiMerchantEmailTemplateUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailTemplateUpdateReq) Set(val *UnibeeApiMerchantEmailTemplateUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailTemplateUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailTemplateUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailTemplateUpdateReq(val *UnibeeApiMerchantEmailTemplateUpdateReq) *NullableUnibeeApiMerchantEmailTemplateUpdateReq {
	return &NullableUnibeeApiMerchantEmailTemplateUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailTemplateUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailTemplateUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


