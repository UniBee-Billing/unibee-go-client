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

// checks if the UnibeeApiMerchantUserReleaseReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantUserReleaseReq{}

// UnibeeApiMerchantUserReleaseReq struct for UnibeeApiMerchantUserReleaseReq
type UnibeeApiMerchantUserReleaseReq struct {
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantUserReleaseReq instantiates a new UnibeeApiMerchantUserReleaseReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantUserReleaseReq() *UnibeeApiMerchantUserReleaseReq {
	this := UnibeeApiMerchantUserReleaseReq{}
	return &this
}

// NewUnibeeApiMerchantUserReleaseReqWithDefaults instantiates a new UnibeeApiMerchantUserReleaseReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantUserReleaseReqWithDefaults() *UnibeeApiMerchantUserReleaseReq {
	this := UnibeeApiMerchantUserReleaseReq{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantUserReleaseReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantUserReleaseReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantUserReleaseReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantUserReleaseReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantUserReleaseReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantUserReleaseReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantUserReleaseReq struct {
	value *UnibeeApiMerchantUserReleaseReq
	isSet bool
}

func (v NullableUnibeeApiMerchantUserReleaseReq) Get() *UnibeeApiMerchantUserReleaseReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantUserReleaseReq) Set(val *UnibeeApiMerchantUserReleaseReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantUserReleaseReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantUserReleaseReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantUserReleaseReq(val *UnibeeApiMerchantUserReleaseReq) *NullableUnibeeApiMerchantUserReleaseReq {
	return &NullableUnibeeApiMerchantUserReleaseReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantUserReleaseReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantUserReleaseReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


