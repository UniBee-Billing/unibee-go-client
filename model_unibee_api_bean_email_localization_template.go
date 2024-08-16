/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanEmailLocalizationTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanEmailLocalizationTemplate{}

// UnibeeApiBeanEmailLocalizationTemplate struct for UnibeeApiBeanEmailLocalizationTemplate
type UnibeeApiBeanEmailLocalizationTemplate struct {
	Content *string `json:"content,omitempty"`
	Language *string `json:"language,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewUnibeeApiBeanEmailLocalizationTemplate instantiates a new UnibeeApiBeanEmailLocalizationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanEmailLocalizationTemplate() *UnibeeApiBeanEmailLocalizationTemplate {
	this := UnibeeApiBeanEmailLocalizationTemplate{}
	return &this
}

// NewUnibeeApiBeanEmailLocalizationTemplateWithDefaults instantiates a new UnibeeApiBeanEmailLocalizationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanEmailLocalizationTemplateWithDefaults() *UnibeeApiBeanEmailLocalizationTemplate {
	this := UnibeeApiBeanEmailLocalizationTemplate{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *UnibeeApiBeanEmailLocalizationTemplate) SetContent(v string) {
	o.Content = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UnibeeApiBeanEmailLocalizationTemplate) SetLanguage(v string) {
	o.Language = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UnibeeApiBeanEmailLocalizationTemplate) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UnibeeApiBeanEmailLocalizationTemplate) SetTitle(v string) {
	o.Title = &v
}

func (o UnibeeApiBeanEmailLocalizationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanEmailLocalizationTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanEmailLocalizationTemplate struct {
	value *UnibeeApiBeanEmailLocalizationTemplate
	isSet bool
}

func (v NullableUnibeeApiBeanEmailLocalizationTemplate) Get() *UnibeeApiBeanEmailLocalizationTemplate {
	return v.value
}

func (v *NullableUnibeeApiBeanEmailLocalizationTemplate) Set(val *UnibeeApiBeanEmailLocalizationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanEmailLocalizationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanEmailLocalizationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanEmailLocalizationTemplate(val *UnibeeApiBeanEmailLocalizationTemplate) *NullableUnibeeApiBeanEmailLocalizationTemplate {
	return &NullableUnibeeApiBeanEmailLocalizationTemplate{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanEmailLocalizationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanEmailLocalizationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


