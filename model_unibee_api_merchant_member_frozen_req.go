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

// checks if the UnibeeApiMerchantMemberFrozenReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberFrozenReq{}

// UnibeeApiMerchantMemberFrozenReq struct for UnibeeApiMerchantMemberFrozenReq
type UnibeeApiMerchantMemberFrozenReq struct {
	// The unique id of member
	MemberId *int64 `json:"memberId,omitempty"`
}

// NewUnibeeApiMerchantMemberFrozenReq instantiates a new UnibeeApiMerchantMemberFrozenReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberFrozenReq() *UnibeeApiMerchantMemberFrozenReq {
	this := UnibeeApiMerchantMemberFrozenReq{}
	return &this
}

// NewUnibeeApiMerchantMemberFrozenReqWithDefaults instantiates a new UnibeeApiMerchantMemberFrozenReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberFrozenReqWithDefaults() *UnibeeApiMerchantMemberFrozenReq {
	this := UnibeeApiMerchantMemberFrozenReq{}
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberFrozenReq) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberFrozenReq) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberFrozenReq) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiMerchantMemberFrozenReq) SetMemberId(v int64) {
	o.MemberId = &v
}

func (o UnibeeApiMerchantMemberFrozenReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberFrozenReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberFrozenReq struct {
	value *UnibeeApiMerchantMemberFrozenReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberFrozenReq) Get() *UnibeeApiMerchantMemberFrozenReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberFrozenReq) Set(val *UnibeeApiMerchantMemberFrozenReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberFrozenReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberFrozenReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberFrozenReq(val *UnibeeApiMerchantMemberFrozenReq) *NullableUnibeeApiMerchantMemberFrozenReq {
	return &NullableUnibeeApiMerchantMemberFrozenReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberFrozenReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberFrozenReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


