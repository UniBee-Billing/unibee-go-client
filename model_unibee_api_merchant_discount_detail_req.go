/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountDetailReq{}

// UnibeeApiMerchantDiscountDetailReq struct for UnibeeApiMerchantDiscountDetailReq
type UnibeeApiMerchantDiscountDetailReq struct {
	// The discount's Id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantDiscountDetailReq UnibeeApiMerchantDiscountDetailReq

// NewUnibeeApiMerchantDiscountDetailReq instantiates a new UnibeeApiMerchantDiscountDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountDetailReq(id int64) *UnibeeApiMerchantDiscountDetailReq {
	this := UnibeeApiMerchantDiscountDetailReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountDetailReqWithDefaults instantiates a new UnibeeApiMerchantDiscountDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountDetailReqWithDefaults() *UnibeeApiMerchantDiscountDetailReq {
	this := UnibeeApiMerchantDiscountDetailReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountDetailReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountDetailReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountDetailReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantDiscountDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantDiscountDetailReq := _UnibeeApiMerchantDiscountDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountDetailReq(varUnibeeApiMerchantDiscountDetailReq)

	return err
}

type NullableUnibeeApiMerchantDiscountDetailReq struct {
	value *UnibeeApiMerchantDiscountDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountDetailReq) Get() *UnibeeApiMerchantDiscountDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountDetailReq) Set(val *UnibeeApiMerchantDiscountDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountDetailReq(val *UnibeeApiMerchantDiscountDetailReq) *NullableUnibeeApiMerchantDiscountDetailReq {
	return &NullableUnibeeApiMerchantDiscountDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


