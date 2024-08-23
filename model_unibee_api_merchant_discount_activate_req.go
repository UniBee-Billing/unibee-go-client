/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountActivateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountActivateReq{}

// UnibeeApiMerchantDiscountActivateReq Activate discount code, the discount code can only effect to payment or subscription after activated
type UnibeeApiMerchantDiscountActivateReq struct {
	// The discount's Id
	Id int64 `json:"id"`
}

type _UnibeeApiMerchantDiscountActivateReq UnibeeApiMerchantDiscountActivateReq

// NewUnibeeApiMerchantDiscountActivateReq instantiates a new UnibeeApiMerchantDiscountActivateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountActivateReq(id int64) *UnibeeApiMerchantDiscountActivateReq {
	this := UnibeeApiMerchantDiscountActivateReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountActivateReqWithDefaults instantiates a new UnibeeApiMerchantDiscountActivateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountActivateReqWithDefaults() *UnibeeApiMerchantDiscountActivateReq {
	this := UnibeeApiMerchantDiscountActivateReq{}
	return &this
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountActivateReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountActivateReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountActivateReq) SetId(v int64) {
	o.Id = v
}

func (o UnibeeApiMerchantDiscountActivateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountActivateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountActivateReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantDiscountActivateReq := _UnibeeApiMerchantDiscountActivateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountActivateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountActivateReq(varUnibeeApiMerchantDiscountActivateReq)

	return err
}

type NullableUnibeeApiMerchantDiscountActivateReq struct {
	value *UnibeeApiMerchantDiscountActivateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountActivateReq) Get() *UnibeeApiMerchantDiscountActivateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountActivateReq) Set(val *UnibeeApiMerchantDiscountActivateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountActivateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountActivateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountActivateReq(val *UnibeeApiMerchantDiscountActivateReq) *NullableUnibeeApiMerchantDiscountActivateReq {
	return &NullableUnibeeApiMerchantDiscountActivateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountActivateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountActivateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


