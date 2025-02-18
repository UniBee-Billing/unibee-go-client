/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiSystemAuthTokenGeneratorRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemAuthTokenGeneratorRes{}

// UnibeeApiSystemAuthTokenGeneratorRes struct for UnibeeApiSystemAuthTokenGeneratorRes
type UnibeeApiSystemAuthTokenGeneratorRes struct {
	Token *string `json:"token,omitempty"`
}

// NewUnibeeApiSystemAuthTokenGeneratorRes instantiates a new UnibeeApiSystemAuthTokenGeneratorRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemAuthTokenGeneratorRes() *UnibeeApiSystemAuthTokenGeneratorRes {
	this := UnibeeApiSystemAuthTokenGeneratorRes{}
	return &this
}

// NewUnibeeApiSystemAuthTokenGeneratorResWithDefaults instantiates a new UnibeeApiSystemAuthTokenGeneratorRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemAuthTokenGeneratorResWithDefaults() *UnibeeApiSystemAuthTokenGeneratorRes {
	this := UnibeeApiSystemAuthTokenGeneratorRes{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UnibeeApiSystemAuthTokenGeneratorRes) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorRes) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UnibeeApiSystemAuthTokenGeneratorRes) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UnibeeApiSystemAuthTokenGeneratorRes) SetToken(v string) {
	o.Token = &v
}

func (o UnibeeApiSystemAuthTokenGeneratorRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemAuthTokenGeneratorRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUnibeeApiSystemAuthTokenGeneratorRes struct {
	value *UnibeeApiSystemAuthTokenGeneratorRes
	isSet bool
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorRes) Get() *UnibeeApiSystemAuthTokenGeneratorRes {
	return v.value
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorRes) Set(val *UnibeeApiSystemAuthTokenGeneratorRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemAuthTokenGeneratorRes(val *UnibeeApiSystemAuthTokenGeneratorRes) *NullableUnibeeApiSystemAuthTokenGeneratorRes {
	return &NullableUnibeeApiSystemAuthTokenGeneratorRes{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemAuthTokenGeneratorRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemAuthTokenGeneratorRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


