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

// checks if the UnibeeApiMerchantAuthLoginOtpVerifyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthLoginOtpVerifyReq{}

// UnibeeApiMerchantAuthLoginOtpVerifyReq OTP login for merchant member, verify OTP code
type UnibeeApiMerchantAuthLoginOtpVerifyReq struct {
	// The merchant member's email address
	Email string `json:"email"`
	// OTP Code, received from email
	VerificationCode string `json:"verificationCode"`
}

type _UnibeeApiMerchantAuthLoginOtpVerifyReq UnibeeApiMerchantAuthLoginOtpVerifyReq

// NewUnibeeApiMerchantAuthLoginOtpVerifyReq instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthLoginOtpVerifyReq(email string, verificationCode string) *UnibeeApiMerchantAuthLoginOtpVerifyReq {
	this := UnibeeApiMerchantAuthLoginOtpVerifyReq{}
	this.Email = email
	this.VerificationCode = verificationCode
	return &this
}

// NewUnibeeApiMerchantAuthLoginOtpVerifyReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthLoginOtpVerifyReqWithDefaults() *UnibeeApiMerchantAuthLoginOtpVerifyReq {
	this := UnibeeApiMerchantAuthLoginOtpVerifyReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) SetEmail(v string) {
	o.Email = v
}

// GetVerificationCode returns the VerificationCode field value
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetVerificationCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerificationCode
}

// GetVerificationCodeOk returns a tuple with the VerificationCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) GetVerificationCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationCode, true
}

// SetVerificationCode sets field value
func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) SetVerificationCode(v string) {
	o.VerificationCode = v
}

func (o UnibeeApiMerchantAuthLoginOtpVerifyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthLoginOtpVerifyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["verificationCode"] = o.VerificationCode
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthLoginOtpVerifyReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantAuthLoginOtpVerifyReq := _UnibeeApiMerchantAuthLoginOtpVerifyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthLoginOtpVerifyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthLoginOtpVerifyReq(varUnibeeApiMerchantAuthLoginOtpVerifyReq)

	return err
}

type NullableUnibeeApiMerchantAuthLoginOtpVerifyReq struct {
	value *UnibeeApiMerchantAuthLoginOtpVerifyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) Get() *UnibeeApiMerchantAuthLoginOtpVerifyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) Set(val *UnibeeApiMerchantAuthLoginOtpVerifyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthLoginOtpVerifyReq(val *UnibeeApiMerchantAuthLoginOtpVerifyReq) *NullableUnibeeApiMerchantAuthLoginOtpVerifyReq {
	return &NullableUnibeeApiMerchantAuthLoginOtpVerifyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


