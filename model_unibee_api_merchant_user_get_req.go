/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserGetReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserGetReq{}

// UnibeeApiMerchantUserGetReq struct for UnibeeApiMerchantUserGetReq
type UnibeeApiMerchantUserGetReq struct {
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantUserGetReq instantiates a new UnibeeApiMerchantUserGetReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserGetReq() *UnibeeApiMerchantUserGetReq {
	this := UnibeeApiMerchantUserGetReq{}
	return &this
}

// NewUnibeeApiMerchantUserGetReqWithDefaults instantiates a new UnibeeApiMerchantUserGetReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserGetReqWithDefaults() *UnibeeApiMerchantUserGetReq {
	this := UnibeeApiMerchantUserGetReq{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserGetReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserGetReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserGetReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantUserGetReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantUserGetReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserGetReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserGetReq struct {
	value *UnibeeApiMerchantUserGetReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserGetReq) Get() *UnibeeApiMerchantUserGetReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserGetReq) Set(val *UnibeeApiMerchantUserGetReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserGetReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserGetReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserGetReq(val *UnibeeApiMerchantUserGetReq) *NullableUnibeeApiMerchantUserGetReq {
	return &NullableUnibeeApiMerchantUserGetReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserGetReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserGetReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


