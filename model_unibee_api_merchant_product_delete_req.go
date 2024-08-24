/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantProductDeleteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductDeleteReq{}

// UnibeeApiMerchantProductDeleteReq Product can being deleted when has no plan linked
type UnibeeApiMerchantProductDeleteReq struct {
	// ProductId
	ProductId int64 `json:"productId"`
}

type _UnibeeApiMerchantProductDeleteReq UnibeeApiMerchantProductDeleteReq

// NewUnibeeApiMerchantProductDeleteReq instantiates a new UnibeeApiMerchantProductDeleteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductDeleteReq(productId int64) *UnibeeApiMerchantProductDeleteReq {
	this := UnibeeApiMerchantProductDeleteReq{}
	this.ProductId = productId
	return &this
}

// NewUnibeeApiMerchantProductDeleteReqWithDefaults instantiates a new UnibeeApiMerchantProductDeleteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductDeleteReqWithDefaults() *UnibeeApiMerchantProductDeleteReq {
	this := UnibeeApiMerchantProductDeleteReq{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *UnibeeApiMerchantProductDeleteReq) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductDeleteReq) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *UnibeeApiMerchantProductDeleteReq) SetProductId(v int64) {
	o.ProductId = v
}

func (o UnibeeApiMerchantProductDeleteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductDeleteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProductDeleteReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantProductDeleteReq := _UnibeeApiMerchantProductDeleteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProductDeleteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProductDeleteReq(varUnibeeApiMerchantProductDeleteReq)

	return err
}

type NullableUnibeeApiMerchantProductDeleteReq struct {
	value *UnibeeApiMerchantProductDeleteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductDeleteReq) Get() *UnibeeApiMerchantProductDeleteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductDeleteReq) Set(val *UnibeeApiMerchantProductDeleteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductDeleteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductDeleteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductDeleteReq(val *UnibeeApiMerchantProductDeleteReq) *NullableUnibeeApiMerchantProductDeleteReq {
	return &NullableUnibeeApiMerchantProductDeleteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductDeleteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductDeleteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


