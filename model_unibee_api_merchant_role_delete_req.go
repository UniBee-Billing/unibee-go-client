/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantRoleDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantRoleDeleteReq{}

// UnibeeApiMerchantRoleDeleteReq struct for UnibeeApiMerchantRoleDeleteReq
type UnibeeApiMerchantRoleDeleteReq struct {
	// id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantRoleDeleteReq UnibeeApiMerchantRoleDeleteReq

// NewUnibeeApiMerchantRoleDeleteReq instantiates a new UnibeeApiMerchantRoleDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantRoleDeleteReq(id int64) *UnibeeApiMerchantRoleDeleteReq {
	this := UnibeeApiMerchantRoleDeleteReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantRoleDeleteReqWithDefaults instantiates a new UnibeeApiMerchantRoleDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantRoleDeleteReqWithDefaults() *UnibeeApiMerchantRoleDeleteReq {
	this := UnibeeApiMerchantRoleDeleteReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantRoleDeleteReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantRoleDeleteReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantRoleDeleteReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantRoleDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantRoleDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantRoleDeleteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUnibeeApiMerchantRoleDeleteReq := _UnibeeApiMerchantRoleDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantRoleDeleteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantRoleDeleteReq(varUnibeeApiMerchantRoleDeleteReq)

	return err
}

type NullableUnibeeApiMerchantRoleDeleteReq struct {
	value *UnibeeApiMerchantRoleDeleteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantRoleDeleteReq) Get() *UnibeeApiMerchantRoleDeleteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantRoleDeleteReq) Set(val *UnibeeApiMerchantRoleDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantRoleDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantRoleDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantRoleDeleteReq(val *UnibeeApiMerchantRoleDeleteReq) *NullableUnibeeApiMerchantRoleDeleteReq {
	return &NullableUnibeeApiMerchantRoleDeleteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantRoleDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantRoleDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


