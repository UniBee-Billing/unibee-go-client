/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthLoginOtpReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthLoginOtpReq{}

// UnibeeApiMerchantAuthLoginOtpReq OTP login for merchant member, send email to member's email address with OTP code'
type UnibeeApiMerchantAuthLoginOtpReq struct {
	// The merchant member's email address
	Email string `json:"email"`
}

type _UnibeeApiMerchantAuthLoginOtpReq UnibeeApiMerchantAuthLoginOtpReq

// NewUnibeeApiMerchantAuthLoginOtpReq instantiates a new UnibeeApiMerchantAuthLoginOtpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthLoginOtpReq(email string) *UnibeeApiMerchantAuthLoginOtpReq {
	this := UnibeeApiMerchantAuthLoginOtpReq{}
	this.Email = email
	return &this
}

// NewUnibeeApiMerchantAuthLoginOtpReqWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOtpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthLoginOtpReqWithDefaults() *UnibeeApiMerchantAuthLoginOtpReq {
	this := UnibeeApiMerchantAuthLoginOtpReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthLoginOtpReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginOtpReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthLoginOtpReq) SetEmail(v string) {
	o.Email = v
}

func (o UnibeeApiMerchantAuthLoginOtpReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthLoginOtpReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthLoginOtpReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUnibeeApiMerchantAuthLoginOtpReq := _UnibeeApiMerchantAuthLoginOtpReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthLoginOtpReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthLoginOtpReq(varUnibeeApiMerchantAuthLoginOtpReq)

	return err
}

type NullableUnibeeApiMerchantAuthLoginOtpReq struct {
	value *UnibeeApiMerchantAuthLoginOtpReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthLoginOtpReq) Get() *UnibeeApiMerchantAuthLoginOtpReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpReq) Set(val *UnibeeApiMerchantAuthLoginOtpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthLoginOtpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthLoginOtpReq(val *UnibeeApiMerchantAuthLoginOtpReq) *NullableUnibeeApiMerchantAuthLoginOtpReq {
	return &NullableUnibeeApiMerchantAuthLoginOtpReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthLoginOtpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


