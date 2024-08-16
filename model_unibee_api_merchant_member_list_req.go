/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMemberListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberListReq{}

// UnibeeApiMerchantMemberListReq struct for UnibeeApiMerchantMemberListReq
type UnibeeApiMerchantMemberListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
	// The member's roleId if specified'
	RoleIds []int64 `json:"roleIds,omitempty"`
}

// NewUnibeeApiMerchantMemberListReq instantiates a new UnibeeApiMerchantMemberListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberListReq() *UnibeeApiMerchantMemberListReq {
	this := UnibeeApiMerchantMemberListReq{}
	return &this
}

// NewUnibeeApiMerchantMemberListReqWithDefaults instantiates a new UnibeeApiMerchantMemberListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberListReqWithDefaults() *UnibeeApiMerchantMemberListReq {
	this := UnibeeApiMerchantMemberListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantMemberListReq) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantMemberListReq) SetPage(v int32) {
	o.Page = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberListReq) GetRoleIds() []int64 {
	if o == nil || IsNil(o.RoleIds) {
		var ret []int64
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberListReq) GetRoleIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberListReq) HasRoleIds() bool {
	if o != nil && !IsNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []int64 and assigns it to the RoleIds field.
func (o *UnibeeApiMerchantMemberListReq) SetRoleIds(v []int64) {
	o.RoleIds = v
}

func (o UnibeeApiMerchantMemberListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.RoleIds) {
		toSerialize["roleIds"] = o.RoleIds
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberListReq struct {
	value *UnibeeApiMerchantMemberListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberListReq) Get() *UnibeeApiMerchantMemberListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberListReq) Set(val *UnibeeApiMerchantMemberListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberListReq(val *UnibeeApiMerchantMemberListReq) *NullableUnibeeApiMerchantMemberListReq {
	return &NullableUnibeeApiMerchantMemberListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


