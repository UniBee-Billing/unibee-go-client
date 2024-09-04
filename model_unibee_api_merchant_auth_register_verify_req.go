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

// checks if the UnibeeApiMerchantAuthRegisterVerifyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthRegisterVerifyReq{}

// UnibeeApiMerchantAuthRegisterVerifyReq Merchant Register, verify OTP code 
type UnibeeApiMerchantAuthRegisterVerifyReq struct {
	// The merchant member's email address
	Email string `json:"email"`
	// OTP Code, received from email
	VerificationCode string `json:"verificationCode"`
}

type _UnibeeApiMerchantAuthRegisterVerifyReq UnibeeApiMerchantAuthRegisterVerifyReq

// NewUnibeeApiMerchantAuthRegisterVerifyReq instantiates a new UnibeeApiMerchantAuthRegisterVerifyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthRegisterVerifyReq(email string, verificationCode string) *UnibeeApiMerchantAuthRegisterVerifyReq {
	this := UnibeeApiMerchantAuthRegisterVerifyReq{}
	this.Email = email
	this.VerificationCode = verificationCode
	return &this
}

// NewUnibeeApiMerchantAuthRegisterVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterVerifyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthRegisterVerifyReqWithDefaults() *UnibeeApiMerchantAuthRegisterVerifyReq {
	this := UnibeeApiMerchantAuthRegisterVerifyReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) SetEmail(v string) {
	o.Email = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *UnibeeApiMerchantAuthRegisterVerifyReq) SetVerificationCode(v string) {
	o.VerificationCode = v
}

func (o UnibeeApiMerchantAuthRegisterVerifyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthRegisterVerifyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["verificationCode"] = o.VerificationCode
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthRegisterVerifyReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"verificationCode",
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

	varUnibeeApiMerchantAuthRegisterVerifyReq := _UnibeeApiMerchantAuthRegisterVerifyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthRegisterVerifyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthRegisterVerifyReq(varUnibeeApiMerchantAuthRegisterVerifyReq)

	return err
}

type NullableUnibeeApiMerchantAuthRegisterVerifyReq struct {
	value *UnibeeApiMerchantAuthRegisterVerifyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyReq) Get() *UnibeeApiMerchantAuthRegisterVerifyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyReq) Set(val *UnibeeApiMerchantAuthRegisterVerifyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthRegisterVerifyReq(val *UnibeeApiMerchantAuthRegisterVerifyReq) *NullableUnibeeApiMerchantAuthRegisterVerifyReq {
	return &NullableUnibeeApiMerchantAuthRegisterVerifyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthRegisterVerifyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthRegisterVerifyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


