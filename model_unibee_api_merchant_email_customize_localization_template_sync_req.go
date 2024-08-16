/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq{}

// UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq Sync the custom localization email template to gateway (sendgrid)
type UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq struct {
	// TemplateData
	TemplateData []UnibeeApiBeanMerchantLocalizationEmailTemplate `json:"templateData"`
	// VersionEnable
	VersionEnable *bool `json:"versionEnable,omitempty"`
}

type _UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq

// NewUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq instantiates a new UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq(templateData []UnibeeApiBeanMerchantLocalizationEmailTemplate) *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq {
	this := UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq{}
	this.TemplateData = templateData
	return &this
}

// NewUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReqWithDefaults instantiates a new UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReqWithDefaults() *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq {
	this := UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq{}
	return &this
}

// GetTemplateData returns the TemplateData field value
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) GetTemplateData() []UnibeeApiBeanMerchantLocalizationEmailTemplate {
	if o == nil {
		var ret []UnibeeApiBeanMerchantLocalizationEmailTemplate
		return ret
	}

	return o.TemplateData
}

// GetTemplateDataOk returns a tuple with the TemplateData field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) GetTemplateDataOk() ([]UnibeeApiBeanMerchantLocalizationEmailTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateData, true
}

// SetTemplateData sets field value
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) SetTemplateData(v []UnibeeApiBeanMerchantLocalizationEmailTemplate) {
	o.TemplateData = v
}

// GetVersionEnable returns the VersionEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) GetVersionEnable() bool {
	if o == nil || IsNil(o.VersionEnable) {
		var ret bool
		return ret
	}
	return *o.VersionEnable
}

// GetVersionEnableOk returns a tuple with the VersionEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) GetVersionEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.VersionEnable) {
		return nil, false
	}
	return o.VersionEnable, true
}

// HasVersionEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) HasVersionEnable() bool {
	if o != nil && !IsNil(o.VersionEnable) {
		return true
	}

	return false
}

// SetVersionEnable gets a reference to the given bool and assigns it to the VersionEnable field.
func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) SetVersionEnable(v bool) {
	o.VersionEnable = &v
}

func (o UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateData"] = o.TemplateData
	if !IsNil(o.VersionEnable) {
		toSerialize["versionEnable"] = o.VersionEnable
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateData",
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

	varUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq := _UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq(varUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq)

	return err
}

type NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq struct {
	value *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) Get() *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) Set(val *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq(val *UnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) *NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq {
	return &NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailCustomizeLocalizationTemplateSyncReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


