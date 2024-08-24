/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantUserFrozenReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserFrozenReq{}

// UnibeeApiMerchantUserFrozenReq struct for UnibeeApiMerchantUserFrozenReq
type UnibeeApiMerchantUserFrozenReq struct {
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantUserFrozenReq instantiates a new UnibeeApiMerchantUserFrozenReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserFrozenReq() *UnibeeApiMerchantUserFrozenReq {
	this := UnibeeApiMerchantUserFrozenReq{}
	return &this
}

// NewUnibeeApiMerchantUserFrozenReqWithDefaults instantiates a new UnibeeApiMerchantUserFrozenReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserFrozenReqWithDefaults() *UnibeeApiMerchantUserFrozenReq {
	this := UnibeeApiMerchantUserFrozenReq{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserFrozenReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserFrozenReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserFrozenReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantUserFrozenReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantUserFrozenReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserFrozenReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserFrozenReq struct {
	value *UnibeeApiMerchantUserFrozenReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserFrozenReq) Get() *UnibeeApiMerchantUserFrozenReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserFrozenReq) Set(val *UnibeeApiMerchantUserFrozenReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserFrozenReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserFrozenReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserFrozenReq(val *UnibeeApiMerchantUserFrozenReq) *NullableUnibeeApiMerchantUserFrozenReq {
	return &NullableUnibeeApiMerchantUserFrozenReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserFrozenReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserFrozenReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


