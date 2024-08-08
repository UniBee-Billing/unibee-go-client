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

// checks if the UnibeeApiMerchantProductCopyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductCopyReq{}

// UnibeeApiMerchantProductCopyReq struct for UnibeeApiMerchantProductCopyReq
type UnibeeApiMerchantProductCopyReq struct {
	// ProductId
	ProductId int64 `json:"productId"`
}

type _UnibeeApiMerchantProductCopyReq UnibeeApiMerchantProductCopyReq

// NewUnibeeApiMerchantProductCopyReq instantiates a new UnibeeApiMerchantProductCopyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductCopyReq(productId int64) *UnibeeApiMerchantProductCopyReq {
	this := UnibeeApiMerchantProductCopyReq{}
	this.ProductId = productId
	return &this
}

// NewUnibeeApiMerchantProductCopyReqWithDefaults instantiates a new UnibeeApiMerchantProductCopyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductCopyReqWithDefaults() *UnibeeApiMerchantProductCopyReq {
	this := UnibeeApiMerchantProductCopyReq{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *UnibeeApiMerchantProductCopyReq) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductCopyReq) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *UnibeeApiMerchantProductCopyReq) SetProductId(v int64) {
	o.ProductId = v
}

func (o UnibeeApiMerchantProductCopyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductCopyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProductCopyReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantProductCopyReq := _UnibeeApiMerchantProductCopyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProductCopyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProductCopyReq(varUnibeeApiMerchantProductCopyReq)

	return err
}

type NullableUnibeeApiMerchantProductCopyReq struct {
	value *UnibeeApiMerchantProductCopyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductCopyReq) Get() *UnibeeApiMerchantProductCopyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductCopyReq) Set(val *UnibeeApiMerchantProductCopyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductCopyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductCopyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductCopyReq(val *UnibeeApiMerchantProductCopyReq) *NullableUnibeeApiMerchantProductCopyReq {
	return &NullableUnibeeApiMerchantProductCopyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductCopyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductCopyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


