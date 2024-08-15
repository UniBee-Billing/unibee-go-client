/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantAuthPasswordForgetOtpReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthPasswordForgetOtpReq{}

// UnibeeApiMerchantAuthPasswordForgetOtpReq Merchant member's password forget OTP process,, send email to member's email address with OTP code'
type UnibeeApiMerchantAuthPasswordForgetOtpReq struct {
	// The merchant member's email address
	Email string `json:"email"`
}

type _UnibeeApiMerchantAuthPasswordForgetOtpReq UnibeeApiMerchantAuthPasswordForgetOtpReq

// NewUnibeeApiMerchantAuthPasswordForgetOtpReq instantiates a new UnibeeApiMerchantAuthPasswordForgetOtpReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthPasswordForgetOtpReq(email string) *UnibeeApiMerchantAuthPasswordForgetOtpReq {
	this := UnibeeApiMerchantAuthPasswordForgetOtpReq{}
	this.Email = email
	return &this
}

// NewUnibeeApiMerchantAuthPasswordForgetOtpReqWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordForgetOtpReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthPasswordForgetOtpReqWithDefaults() *UnibeeApiMerchantAuthPasswordForgetOtpReq {
	this := UnibeeApiMerchantAuthPasswordForgetOtpReq{}
	return &this
}

// GetEmail returns the Email field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpReq) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthPasswordForgetOtpReq) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UnibeeApiMerchantAuthPasswordForgetOtpReq) SetEmail(v string) {
	o.Email = v
}

func (o UnibeeApiMerchantAuthPasswordForgetOtpReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthPasswordForgetOtpReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *UnibeeApiMerchantAuthPasswordForgetOtpReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantAuthPasswordForgetOtpReq := _UnibeeApiMerchantAuthPasswordForgetOtpReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantAuthPasswordForgetOtpReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantAuthPasswordForgetOtpReq(varUnibeeApiMerchantAuthPasswordForgetOtpReq)

	return err
}

type NullableUnibeeApiMerchantAuthPasswordForgetOtpReq struct {
	value *UnibeeApiMerchantAuthPasswordForgetOtpReq
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) Get() *UnibeeApiMerchantAuthPasswordForgetOtpReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) Set(val *UnibeeApiMerchantAuthPasswordForgetOtpReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthPasswordForgetOtpReq(val *UnibeeApiMerchantAuthPasswordForgetOtpReq) *NullableUnibeeApiMerchantAuthPasswordForgetOtpReq {
	return &NullableUnibeeApiMerchantAuthPasswordForgetOtpReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthPasswordForgetOtpReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


