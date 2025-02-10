/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthPasswordSetupOtpReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthPasswordSetupOtpReq{}

// UnibeeApiMerchantAuthPasswordSetupOtpReq Member Password Setup
type UnibeeApiMerchantAuthPasswordSetupOtpReq struct {
	// The merchant member email address
	Email string `json:"email"`
	// The new password
	NewPassword string `json:"newPassword"`
	// The merchant member password setup token
	SetupToken string `json:"setupToken"`
}

type _UnibeeApiMerchantAuthPasswordSetupOtpReq UnibeeApiMerchantAuthPasswordSetupOtpReq

// NewUnibeeApiMerchantAuthPasswordSetupOtpReq instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthPasswordSetupOtpReq(email string, newPassword string, setupToken string) *UnibeeApiMerchantAuthPasswordSetupOtpReq {
	this := UnibeeApiMerchantAuthPasswordSetupOtpReq{}
	this.Email = email
	this.NewPassword = newPassword
	this.SetupToken = setupToken
	return &this
}

// NewUnibeeApiMerchantAuthPasswordSetupOtpReqWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthPasswordSetupOtpReqWithDefaults() *UnibeeApiMerchantAuthPasswordSetupOtpReq {
	this := UnibeeApiMerchantAuthPasswordSetupOtpReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetEmail(v string) {
	o.Email = v
}

// GetNewPassword returns the NewPassword field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetSetupToken returns the SetupToken field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetSetupToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SetupToken
}

// GetSetupTokenOk returns a tuple with the SetupToken field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) GetSetupTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetupToken, true
}

// SetSetupToken sets field value
func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) SetSetupToken(v string) {
	o.SetupToken = v
}

func (o UnibeeApiMerchantAuthPasswordSetupOtpReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthPasswordSetupOtpReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["newPassword"] = o.NewPassword
	toSerialize["setupToken"] = o.SetupToken
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthPasswordSetupOtpReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"newPassword",
		"setupToken",
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

	varUnibeeApiMerchantAuthPasswordSetupOtpReq := _UnibeeApiMerchantAuthPasswordSetupOtpReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthPasswordSetupOtpReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthPasswordSetupOtpReq(varUnibeeApiMerchantAuthPasswordSetupOtpReq)

	return err
}

type NullableUnibeeApiMerchantAuthPasswordSetupOtpReq struct {
	value *UnibeeApiMerchantAuthPasswordSetupOtpReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) Get() *UnibeeApiMerchantAuthPasswordSetupOtpReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) Set(val *UnibeeApiMerchantAuthPasswordSetupOtpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthPasswordSetupOtpReq(val *UnibeeApiMerchantAuthPasswordSetupOtpReq) *NullableUnibeeApiMerchantAuthPasswordSetupOtpReq {
	return &NullableUnibeeApiMerchantAuthPasswordSetupOtpReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthPasswordSetupOtpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


