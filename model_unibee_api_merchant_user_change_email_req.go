/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantUserChangeEmailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserChangeEmailReq{}

// UnibeeApiMerchantUserChangeEmailReq struct for UnibeeApiMerchantUserChangeEmailReq
type UnibeeApiMerchantUserChangeEmailReq struct {
	// The externalUserId of user, either ExternalUserId or UserId needed
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// Target Email want to change
	NewEmail string `json:"newEmail"`
	// The id of user, either ExternalUserId or UserId needed
	UserId *int64 `json:"userId,omitempty"`
}

type _UnibeeApiMerchantUserChangeEmailReq UnibeeApiMerchantUserChangeEmailReq

// NewUnibeeApiMerchantUserChangeEmailReq instantiates a new UnibeeApiMerchantUserChangeEmailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserChangeEmailReq(newEmail string) *UnibeeApiMerchantUserChangeEmailReq {
	this := UnibeeApiMerchantUserChangeEmailReq{}
	this.NewEmail = newEmail
	return &this
}

// NewUnibeeApiMerchantUserChangeEmailReqWithDefaults instantiates a new UnibeeApiMerchantUserChangeEmailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserChangeEmailReqWithDefaults() *UnibeeApiMerchantUserChangeEmailReq {
	this := UnibeeApiMerchantUserChangeEmailReq{}
	return &this
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserChangeEmailReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeEmailReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserChangeEmailReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantUserChangeEmailReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetNewEmail returns the NewEmail field value
func (o *UnibeeApiMerchantUserChangeEmailReq) GetNewEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewEmail
}

// GetNewEmailOk returns a tuple with the NewEmail field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeEmailReq) GetNewEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewEmail, true
}

// SetNewEmail sets field value
func (o *UnibeeApiMerchantUserChangeEmailReq) SetNewEmail(v string) {
	o.NewEmail = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserChangeEmailReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserChangeEmailReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserChangeEmailReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantUserChangeEmailReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantUserChangeEmailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserChangeEmailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	toSerialize["newEmail"] = o.NewEmail
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantUserChangeEmailReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newEmail",
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

	varUnibeeApiMerchantUserChangeEmailReq := _UnibeeApiMerchantUserChangeEmailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantUserChangeEmailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantUserChangeEmailReq(varUnibeeApiMerchantUserChangeEmailReq)

	return err
}

type NullableUnibeeApiMerchantUserChangeEmailReq struct {
	value *UnibeeApiMerchantUserChangeEmailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserChangeEmailReq) Get() *UnibeeApiMerchantUserChangeEmailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserChangeEmailReq) Set(val *UnibeeApiMerchantUserChangeEmailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserChangeEmailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserChangeEmailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserChangeEmailReq(val *UnibeeApiMerchantUserChangeEmailReq) *NullableUnibeeApiMerchantUserChangeEmailReq {
	return &NullableUnibeeApiMerchantUserChangeEmailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserChangeEmailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserChangeEmailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


