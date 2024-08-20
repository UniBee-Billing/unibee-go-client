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

// checks if the UnibeeApiMerchantProductActivateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductActivateReq{}

// UnibeeApiMerchantProductActivateReq struct for UnibeeApiMerchantProductActivateReq
type UnibeeApiMerchantProductActivateReq struct {
	// ProductId
	ProductId int64 `json:"productId"`
}

type _UnibeeApiMerchantProductActivateReq UnibeeApiMerchantProductActivateReq

// NewUnibeeApiMerchantProductActivateReq instantiates a new UnibeeApiMerchantProductActivateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductActivateReq(productId int64) *UnibeeApiMerchantProductActivateReq {
	this := UnibeeApiMerchantProductActivateReq{}
	this.ProductId = productId
	return &this
}

// NewUnibeeApiMerchantProductActivateReqWithDefaults instantiates a new UnibeeApiMerchantProductActivateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductActivateReqWithDefaults() *UnibeeApiMerchantProductActivateReq {
	this := UnibeeApiMerchantProductActivateReq{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *UnibeeApiMerchantProductActivateReq) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductActivateReq) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *UnibeeApiMerchantProductActivateReq) SetProductId(v int64) {
	o.ProductId = v
}

func (o UnibeeApiMerchantProductActivateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductActivateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProductActivateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productId",
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

	varUnibeeApiMerchantProductActivateReq := _UnibeeApiMerchantProductActivateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProductActivateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProductActivateReq(varUnibeeApiMerchantProductActivateReq)

	return err
}

type NullableUnibeeApiMerchantProductActivateReq struct {
	value *UnibeeApiMerchantProductActivateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductActivateReq) Get() *UnibeeApiMerchantProductActivateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductActivateReq) Set(val *UnibeeApiMerchantProductActivateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductActivateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductActivateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductActivateReq(val *UnibeeApiMerchantProductActivateReq) *NullableUnibeeApiMerchantProductActivateReq {
	return &NullableUnibeeApiMerchantProductActivateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductActivateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductActivateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


