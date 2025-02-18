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

// checks if the UnibeeApiMerchantMemberPasswordResetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberPasswordResetReq{}

// UnibeeApiMerchantMemberPasswordResetReq struct for UnibeeApiMerchantMemberPasswordResetReq
type UnibeeApiMerchantMemberPasswordResetReq struct {
	// The new password of member
	NewPassword string `json:"newPassword"`
	// The old password of member
	OldPassword string `json:"oldPassword"`
}

type _UnibeeApiMerchantMemberPasswordResetReq UnibeeApiMerchantMemberPasswordResetReq

// NewUnibeeApiMerchantMemberPasswordResetReq instantiates a new UnibeeApiMerchantMemberPasswordResetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberPasswordResetReq(newPassword string, oldPassword string) *UnibeeApiMerchantMemberPasswordResetReq {
	this := UnibeeApiMerchantMemberPasswordResetReq{}
	this.NewPassword = newPassword
	this.OldPassword = oldPassword
	return &this
}

// NewUnibeeApiMerchantMemberPasswordResetReqWithDefaults instantiates a new UnibeeApiMerchantMemberPasswordResetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberPasswordResetReqWithDefaults() *UnibeeApiMerchantMemberPasswordResetReq {
	this := UnibeeApiMerchantMemberPasswordResetReq{}
	return &this
}

// GetNewPassword returns the NewPassword field value
func (o *UnibeeApiMerchantMemberPasswordResetReq) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberPasswordResetReq) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *UnibeeApiMerchantMemberPasswordResetReq) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetOldPassword returns the OldPassword field value
func (o *UnibeeApiMerchantMemberPasswordResetReq) GetOldPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberPasswordResetReq) GetOldPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldPassword, true
}

// SetOldPassword sets field value
func (o *UnibeeApiMerchantMemberPasswordResetReq) SetOldPassword(v string) {
	o.OldPassword = v
}

func (o UnibeeApiMerchantMemberPasswordResetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberPasswordResetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newPassword"] = o.NewPassword
	toSerialize["oldPassword"] = o.OldPassword
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMemberPasswordResetReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPassword",
		"oldPassword",
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

	varUnibeeApiMerchantMemberPasswordResetReq := _UnibeeApiMerchantMemberPasswordResetReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMemberPasswordResetReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMemberPasswordResetReq(varUnibeeApiMerchantMemberPasswordResetReq)

	return err
}

type NullableUnibeeApiMerchantMemberPasswordResetReq struct {
	value *UnibeeApiMerchantMemberPasswordResetReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberPasswordResetReq) Get() *UnibeeApiMerchantMemberPasswordResetReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberPasswordResetReq) Set(val *UnibeeApiMerchantMemberPasswordResetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberPasswordResetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberPasswordResetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberPasswordResetReq(val *UnibeeApiMerchantMemberPasswordResetReq) *NullableUnibeeApiMerchantMemberPasswordResetReq {
	return &NullableUnibeeApiMerchantMemberPasswordResetReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberPasswordResetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberPasswordResetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


