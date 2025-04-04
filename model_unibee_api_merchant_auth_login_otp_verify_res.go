/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantAuthLoginOtpVerifyRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantAuthLoginOtpVerifyRes{}

// UnibeeApiMerchantAuthLoginOtpVerifyRes struct for UnibeeApiMerchantAuthLoginOtpVerifyRes
type UnibeeApiMerchantAuthLoginOtpVerifyRes struct {
	MerchantMember *UnibeeApiBeanDetailMerchantMemberDetail `json:"merchantMember,omitempty"`
	// Access token of admin portal
	Token *string `json:"token,omitempty"`
}

// NewUnibeeApiMerchantAuthLoginOtpVerifyRes instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantAuthLoginOtpVerifyRes() *UnibeeApiMerchantAuthLoginOtpVerifyRes {
	this := UnibeeApiMerchantAuthLoginOtpVerifyRes{}
	return &this
}

// NewUnibeeApiMerchantAuthLoginOtpVerifyResWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOtpVerifyRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantAuthLoginOtpVerifyResWithDefaults() *UnibeeApiMerchantAuthLoginOtpVerifyRes {
	this := UnibeeApiMerchantAuthLoginOtpVerifyRes{}
	return &this
}

// GetMerchantMember returns the MerchantMember field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.MerchantMember) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.MerchantMember
}

// GetMerchantMemberOk returns a tuple with the MerchantMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.MerchantMember) {
		return nil, false
	}
	return o.MerchantMember, true
}

// HasMerchantMember returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) HasMerchantMember() bool {
	if o != nil && !IsNil(o.MerchantMember) {
		return true
	}

	return false
}

// SetMerchantMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the MerchantMember field.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.MerchantMember = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UnibeeApiMerchantAuthLoginOtpVerifyRes) SetToken(v string) {
	o.Token = &v
}

func (o UnibeeApiMerchantAuthLoginOtpVerifyRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantAuthLoginOtpVerifyRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantMember) {
		toSerialize["merchantMember"] = o.MerchantMember
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantAuthLoginOtpVerifyRes struct {
	value *UnibeeApiMerchantAuthLoginOtpVerifyRes
	isSet bool
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) Get() *UnibeeApiMerchantAuthLoginOtpVerifyRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) Set(val *UnibeeApiMerchantAuthLoginOtpVerifyRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantAuthLoginOtpVerifyRes(val *UnibeeApiMerchantAuthLoginOtpVerifyRes) *NullableUnibeeApiMerchantAuthLoginOtpVerifyRes {
	return &NullableUnibeeApiMerchantAuthLoginOtpVerifyRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantAuthLoginOtpVerifyRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


