/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantRoleEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantRoleEditReq{}

// UnibeeApiMerchantRoleEditReq struct for UnibeeApiMerchantRoleEditReq
type UnibeeApiMerchantRoleEditReq struct {
	// id
	Id int64 `json:"id"`
	// Permissions
	Permissions []UnibeeApiBeanMerchantRolePermission `json:"permissions"`
	// Role
	Role string `json:"role"`
}

type _UnibeeApiMerchantRoleEditReq UnibeeApiMerchantRoleEditReq

// NewUnibeeApiMerchantRoleEditReq instantiates a new UnibeeApiMerchantRoleEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantRoleEditReq(id int64, permissions []UnibeeApiBeanMerchantRolePermission, role string) *UnibeeApiMerchantRoleEditReq {
	this := UnibeeApiMerchantRoleEditReq{}
	this.Id = id
	this.Permissions = permissions
	this.Role = role
	return &this
}

// NewUnibeeApiMerchantRoleEditReqWithDefaults instantiates a new UnibeeApiMerchantRoleEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantRoleEditReqWithDefaults() *UnibeeApiMerchantRoleEditReq {
	this := UnibeeApiMerchantRoleEditReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantRoleEditReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleEditReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantRoleEditReq) SetId(v int64) {
	o.Id = v
}

// GetPermissions returns the Permissions field value
func (o *UnibeeApiMerchantRoleEditReq) GetPermissions() []UnibeeApiBeanMerchantRolePermission {
	if o == nil {
		var ret []UnibeeApiBeanMerchantRolePermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleEditReq) GetPermissionsOk() ([]UnibeeApiBeanMerchantRolePermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *UnibeeApiMerchantRoleEditReq) SetPermissions(v []UnibeeApiBeanMerchantRolePermission) {
	o.Permissions = v
}

// GetRole returns the Role field value
func (o *UnibeeApiMerchantRoleEditReq) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleEditReq) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UnibeeApiMerchantRoleEditReq) SetRole(v string) {
	o.Role = v
}

func (o UnibeeApiMerchantRoleEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantRoleEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["permissions"] = o.Permissions
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *UnibeeApiMerchantRoleEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"permissions",
		"role",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantRoleEditReq := _UnibeeApiMerchantRoleEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantRoleEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantRoleEditReq(varUnibeeApiMerchantRoleEditReq)

	return err
}

type NullableUnibeeApiMerchantRoleEditReq struct {
	value *UnibeeApiMerchantRoleEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantRoleEditReq) Get() *UnibeeApiMerchantRoleEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantRoleEditReq) Set(val *UnibeeApiMerchantRoleEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantRoleEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantRoleEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantRoleEditReq(val *UnibeeApiMerchantRoleEditReq) *NullableUnibeeApiMerchantRoleEditReq {
	return &NullableUnibeeApiMerchantRoleEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantRoleEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantRoleEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


