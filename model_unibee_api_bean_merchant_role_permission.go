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

// checks if the UnibeeApiBeanMerchantRolePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantRolePermission{}

// UnibeeApiBeanMerchantRolePermission struct for UnibeeApiBeanMerchantRolePermission
type UnibeeApiBeanMerchantRolePermission struct {
	// Group
	Group *string `json:"group,omitempty"`
	// Permissions
	Permissions []string `json:"permissions,omitempty"`
}

// NewUnibeeApiBeanMerchantRolePermission instantiates a new UnibeeApiBeanMerchantRolePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantRolePermission() *UnibeeApiBeanMerchantRolePermission {
	this := UnibeeApiBeanMerchantRolePermission{}
	return &this
}

// NewUnibeeApiBeanMerchantRolePermissionWithDefaults instantiates a new UnibeeApiBeanMerchantRolePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantRolePermissionWithDefaults() *UnibeeApiBeanMerchantRolePermission {
	this := UnibeeApiBeanMerchantRolePermission{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRolePermission) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRolePermission) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRolePermission) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *UnibeeApiBeanMerchantRolePermission) SetGroup(v string) {
	o.Group = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantRolePermission) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantRolePermission) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantRolePermission) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *UnibeeApiBeanMerchantRolePermission) SetPermissions(v []string) {
	o.Permissions = v
}

func (o UnibeeApiBeanMerchantRolePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantRolePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantRolePermission struct {
	value *UnibeeApiBeanMerchantRolePermission
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantRolePermission) Get() *UnibeeApiBeanMerchantRolePermission {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantRolePermission) Set(val *UnibeeApiBeanMerchantRolePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantRolePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantRolePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantRolePermission(val *UnibeeApiBeanMerchantRolePermission) *NullableUnibeeApiBeanMerchantRolePermission {
	return &NullableUnibeeApiBeanMerchantRolePermission{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantRolePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantRolePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


