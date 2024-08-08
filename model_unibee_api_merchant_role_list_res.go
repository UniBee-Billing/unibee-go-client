/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantRoleListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantRoleListRes{}

// UnibeeApiMerchantRoleListRes struct for UnibeeApiMerchantRoleListRes
type UnibeeApiMerchantRoleListRes struct {
	// Merchant Roles
	MerchantRoles []UnibeeApiBeanMerchantRole `json:"merchantRoles,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantRoleListRes instantiates a new UnibeeApiMerchantRoleListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantRoleListRes() *UnibeeApiMerchantRoleListRes {
	this := UnibeeApiMerchantRoleListRes{}
	return &this
}

// NewUnibeeApiMerchantRoleListResWithDefaults instantiates a new UnibeeApiMerchantRoleListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantRoleListResWithDefaults() *UnibeeApiMerchantRoleListRes {
	this := UnibeeApiMerchantRoleListRes{}
	return &this
}

// GetMerchantRoles returns the MerchantRoles field value if set, zero value otherwise.
func (o *UnibeeApiMerchantRoleListRes) GetMerchantRoles() []UnibeeApiBeanMerchantRole {
	if o == nil || IsNil(o.MerchantRoles) {
		var ret []UnibeeApiBeanMerchantRole
		return ret
	}
	return o.MerchantRoles
}

// GetMerchantRolesOk returns a tuple with the MerchantRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleListRes) GetMerchantRolesOk() ([]UnibeeApiBeanMerchantRole, bool) {
	if o == nil || IsNil(o.MerchantRoles) {
		return nil, false
	}
	return o.MerchantRoles, true
}

// HasMerchantRoles returns a boolean if a field has been set.
func (o *UnibeeApiMerchantRoleListRes) HasMerchantRoles() bool {
	if o != nil && !IsNil(o.MerchantRoles) {
		return true
	}

	return false
}

// SetMerchantRoles gets a reference to the given []UnibeeApiBeanMerchantRole and assigns it to the MerchantRoles field.
func (o *UnibeeApiMerchantRoleListRes) SetMerchantRoles(v []UnibeeApiBeanMerchantRole) {
	o.MerchantRoles = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantRoleListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantRoleListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantRoleListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantRoleListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantRoleListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantRoles) {
		toSerialize["merchantRoles"] = o.MerchantRoles
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantRoleListRes struct {
	value *UnibeeApiMerchantRoleListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantRoleListRes) Get() *UnibeeApiMerchantRoleListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantRoleListRes) Set(val *UnibeeApiMerchantRoleListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantRoleListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantRoleListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantRoleListRes(val *UnibeeApiMerchantRoleListRes) *NullableUnibeeApiMerchantRoleListRes {
	return &NullableUnibeeApiMerchantRoleListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantRoleListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantRoleListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


