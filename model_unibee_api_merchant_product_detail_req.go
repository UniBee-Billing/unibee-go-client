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

// checks if the UnibeeApiMerchantProductDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductDetailReq{}

// UnibeeApiMerchantProductDetailReq struct for UnibeeApiMerchantProductDetailReq
type UnibeeApiMerchantProductDetailReq struct {
	// ProductId
	ProductId int64 `json:"productId"`
}

type _UnibeeApiMerchantProductDetailReq UnibeeApiMerchantProductDetailReq

// NewUnibeeApiMerchantProductDetailReq instantiates a new UnibeeApiMerchantProductDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductDetailReq(productId int64) *UnibeeApiMerchantProductDetailReq {
	this := UnibeeApiMerchantProductDetailReq{}
	this.ProductId = productId
	return &this
}

// NewUnibeeApiMerchantProductDetailReqWithDefaults instantiates a new UnibeeApiMerchantProductDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductDetailReqWithDefaults() *UnibeeApiMerchantProductDetailReq {
	this := UnibeeApiMerchantProductDetailReq{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *UnibeeApiMerchantProductDetailReq) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductDetailReq) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *UnibeeApiMerchantProductDetailReq) SetProductId(v int64) {
	o.ProductId = v
}

func (o UnibeeApiMerchantProductDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProductDetailReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantProductDetailReq := _UnibeeApiMerchantProductDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProductDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProductDetailReq(varUnibeeApiMerchantProductDetailReq)

	return err
}

type NullableUnibeeApiMerchantProductDetailReq struct {
	value *UnibeeApiMerchantProductDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductDetailReq) Get() *UnibeeApiMerchantProductDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductDetailReq) Set(val *UnibeeApiMerchantProductDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductDetailReq(val *UnibeeApiMerchantProductDetailReq) *NullableUnibeeApiMerchantProductDetailReq {
	return &NullableUnibeeApiMerchantProductDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


