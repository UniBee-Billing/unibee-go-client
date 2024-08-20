/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountDeactivateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountDeactivateReq{}

// UnibeeApiMerchantDiscountDeactivateReq Deactivate discount code
type UnibeeApiMerchantDiscountDeactivateReq struct {
	// The discount's Id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantDiscountDeactivateReq UnibeeApiMerchantDiscountDeactivateReq

// NewUnibeeApiMerchantDiscountDeactivateReq instantiates a new UnibeeApiMerchantDiscountDeactivateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountDeactivateReq(id int64) *UnibeeApiMerchantDiscountDeactivateReq {
	this := UnibeeApiMerchantDiscountDeactivateReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountDeactivateReqWithDefaults instantiates a new UnibeeApiMerchantDiscountDeactivateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountDeactivateReqWithDefaults() *UnibeeApiMerchantDiscountDeactivateReq {
	this := UnibeeApiMerchantDiscountDeactivateReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountDeactivateReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountDeactivateReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountDeactivateReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantDiscountDeactivateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountDeactivateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountDeactivateReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantDiscountDeactivateReq := _UnibeeApiMerchantDiscountDeactivateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountDeactivateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountDeactivateReq(varUnibeeApiMerchantDiscountDeactivateReq)

	return err
}

type NullableUnibeeApiMerchantDiscountDeactivateReq struct {
	value *UnibeeApiMerchantDiscountDeactivateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountDeactivateReq) Get() *UnibeeApiMerchantDiscountDeactivateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountDeactivateReq) Set(val *UnibeeApiMerchantDiscountDeactivateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountDeactivateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountDeactivateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountDeactivateReq(val *UnibeeApiMerchantDiscountDeactivateReq) *NullableUnibeeApiMerchantDiscountDeactivateReq {
	return &NullableUnibeeApiMerchantDiscountDeactivateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountDeactivateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountDeactivateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


