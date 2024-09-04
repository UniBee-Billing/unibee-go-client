/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantEmailSendTemplateEmailToUserReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantEmailSendTemplateEmailToUserReq{}

// UnibeeApiMerchantEmailSendTemplateEmailToUserReq struct for UnibeeApiMerchantEmailSendTemplateEmailToUserReq
type UnibeeApiMerchantEmailSendTemplateEmailToUserReq struct {
	// The name of email template
	TemplateName string `json:"templateName"`
	// UserId
	UserId int64 `json:"userId"`
	// Variables，Map
	Variables *map[string]interface{} `json:"variables,omitempty"`
}

type _UnibeeApiMerchantEmailSendTemplateEmailToUserReq UnibeeApiMerchantEmailSendTemplateEmailToUserReq

// NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq instantiates a new UnibeeApiMerchantEmailSendTemplateEmailToUserReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq(templateName string, userId int64) *UnibeeApiMerchantEmailSendTemplateEmailToUserReq {
	this := UnibeeApiMerchantEmailSendTemplateEmailToUserReq{}
	this.TemplateName = templateName
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantEmailSendTemplateEmailToUserReqWithDefaults instantiates a new UnibeeApiMerchantEmailSendTemplateEmailToUserReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantEmailSendTemplateEmailToUserReqWithDefaults() *UnibeeApiMerchantEmailSendTemplateEmailToUserReq {
	this := UnibeeApiMerchantEmailSendTemplateEmailToUserReq{}
	return &this
}

// GetTemplateName returns the TemplateName field value
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetUserId(v int64) {
	o.UserId = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetVariablesOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetVariables(v map[string]interface{}) {
	o.Variables = &v
}

func (o UnibeeApiMerchantEmailSendTemplateEmailToUserReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantEmailSendTemplateEmailToUserReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateName"] = o.TemplateName
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateName",
		"userId",
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

	varUnibeeApiMerchantEmailSendTemplateEmailToUserReq := _UnibeeApiMerchantEmailSendTemplateEmailToUserReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantEmailSendTemplateEmailToUserReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantEmailSendTemplateEmailToUserReq(varUnibeeApiMerchantEmailSendTemplateEmailToUserReq)

	return err
}

type NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq struct {
	value *UnibeeApiMerchantEmailSendTemplateEmailToUserReq
	isSet bool
}

func (v NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) Get() *UnibeeApiMerchantEmailSendTemplateEmailToUserReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) Set(val *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq(val *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) *NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq {
	return &NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantEmailSendTemplateEmailToUserReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


