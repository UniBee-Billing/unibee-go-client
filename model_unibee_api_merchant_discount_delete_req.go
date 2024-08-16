/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408161355 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountDeleteReq{}

// UnibeeApiMerchantDiscountDeleteReq Delete discount code before activate
type UnibeeApiMerchantDiscountDeleteReq struct {
	// The discount's Id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantDiscountDeleteReq UnibeeApiMerchantDiscountDeleteReq

// NewUnibeeApiMerchantDiscountDeleteReq instantiates a new UnibeeApiMerchantDiscountDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountDeleteReq(id int64) *UnibeeApiMerchantDiscountDeleteReq {
	this := UnibeeApiMerchantDiscountDeleteReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountDeleteReqWithDefaults instantiates a new UnibeeApiMerchantDiscountDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountDeleteReqWithDefaults() *UnibeeApiMerchantDiscountDeleteReq {
	this := UnibeeApiMerchantDiscountDeleteReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountDeleteReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountDeleteReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountDeleteReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantDiscountDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountDeleteReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantDiscountDeleteReq := _UnibeeApiMerchantDiscountDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountDeleteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountDeleteReq(varUnibeeApiMerchantDiscountDeleteReq)

	return err
}

type NullableUnibeeApiMerchantDiscountDeleteReq struct {
	value *UnibeeApiMerchantDiscountDeleteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountDeleteReq) Get() *UnibeeApiMerchantDiscountDeleteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountDeleteReq) Set(val *UnibeeApiMerchantDiscountDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountDeleteReq(val *UnibeeApiMerchantDiscountDeleteReq) *NullableUnibeeApiMerchantDiscountDeleteReq {
	return &NullableUnibeeApiMerchantDiscountDeleteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


