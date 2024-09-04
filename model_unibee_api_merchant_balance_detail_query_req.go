/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantBalanceDetailQueryReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantBalanceDetailQueryReq{}

// UnibeeApiMerchantBalanceDetailQueryReq struct for UnibeeApiMerchantBalanceDetailQueryReq
type UnibeeApiMerchantBalanceDetailQueryReq struct {
	// gatewayId
	GatewayId int64 `json:"gatewayId"`
}

type _UnibeeApiMerchantBalanceDetailQueryReq UnibeeApiMerchantBalanceDetailQueryReq

// NewUnibeeApiMerchantBalanceDetailQueryReq instantiates a new UnibeeApiMerchantBalanceDetailQueryReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantBalanceDetailQueryReq(gatewayId int64) *UnibeeApiMerchantBalanceDetailQueryReq {
	this := UnibeeApiMerchantBalanceDetailQueryReq{}
	this.GatewayId = gatewayId
	return &this
}

// NewUnibeeApiMerchantBalanceDetailQueryReqWithDefaults instantiates a new UnibeeApiMerchantBalanceDetailQueryReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantBalanceDetailQueryReqWithDefaults() *UnibeeApiMerchantBalanceDetailQueryReq {
	this := UnibeeApiMerchantBalanceDetailQueryReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantBalanceDetailQueryReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantBalanceDetailQueryReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

func (o UnibeeApiMerchantBalanceDetailQueryReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantBalanceDetailQueryReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantBalanceDetailQueryReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
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

	varUnibeeApiMerchantBalanceDetailQueryReq := _UnibeeApiMerchantBalanceDetailQueryReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantBalanceDetailQueryReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantBalanceDetailQueryReq(varUnibeeApiMerchantBalanceDetailQueryReq)

	return err
}

type NullableUnibeeApiMerchantBalanceDetailQueryReq struct {
	value *UnibeeApiMerchantBalanceDetailQueryReq
	isSet bool
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryReq) Get() *UnibeeApiMerchantBalanceDetailQueryReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryReq) Set(val *UnibeeApiMerchantBalanceDetailQueryReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantBalanceDetailQueryReq(val *UnibeeApiMerchantBalanceDetailQueryReq) *NullableUnibeeApiMerchantBalanceDetailQueryReq {
	return &NullableUnibeeApiMerchantBalanceDetailQueryReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


