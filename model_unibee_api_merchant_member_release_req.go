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

// checks if the UnibeeApiMerchantMemberReleaseReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberReleaseReq{}

// UnibeeApiMerchantMemberReleaseReq struct for UnibeeApiMerchantMemberReleaseReq
type UnibeeApiMerchantMemberReleaseReq struct {
	// The unique id of member
	MemberId *int64 `json:"memberId,omitempty"`
}

// NewUnibeeApiMerchantMemberReleaseReq instantiates a new UnibeeApiMerchantMemberReleaseReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberReleaseReq() *UnibeeApiMerchantMemberReleaseReq {
	this := UnibeeApiMerchantMemberReleaseReq{}
	return &this
}

// NewUnibeeApiMerchantMemberReleaseReqWithDefaults instantiates a new UnibeeApiMerchantMemberReleaseReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberReleaseReqWithDefaults() *UnibeeApiMerchantMemberReleaseReq {
	this := UnibeeApiMerchantMemberReleaseReq{}
	return &this
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberReleaseReq) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberReleaseReq) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberReleaseReq) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiMerchantMemberReleaseReq) SetMemberId(v int64) {
	o.MemberId = &v
}

func (o UnibeeApiMerchantMemberReleaseReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberReleaseReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberReleaseReq struct {
	value *UnibeeApiMerchantMemberReleaseReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberReleaseReq) Get() *UnibeeApiMerchantMemberReleaseReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberReleaseReq) Set(val *UnibeeApiMerchantMemberReleaseReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberReleaseReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberReleaseReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberReleaseReq(val *UnibeeApiMerchantMemberReleaseReq) *NullableUnibeeApiMerchantMemberReleaseReq {
	return &NullableUnibeeApiMerchantMemberReleaseReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberReleaseReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberReleaseReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


