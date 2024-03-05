/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantBalanceUserDetailQueryReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantBalanceUserDetailQueryReq{}

// UnibeeApiMerchantBalanceUserDetailQueryReq struct for UnibeeApiMerchantBalanceUserDetailQueryReq
type UnibeeApiMerchantBalanceUserDetailQueryReq struct {
	// gatewayId
	GatewayId int64 `json:"gatewayId"`
	// userId
	UserId int64 `json:"userId"`
}

type _UnibeeApiMerchantBalanceUserDetailQueryReq UnibeeApiMerchantBalanceUserDetailQueryReq

// NewUnibeeApiMerchantBalanceUserDetailQueryReq instantiates a new UnibeeApiMerchantBalanceUserDetailQueryReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantBalanceUserDetailQueryReq(gatewayId int64, userId int64) *UnibeeApiMerchantBalanceUserDetailQueryReq {
	this := UnibeeApiMerchantBalanceUserDetailQueryReq{}
	this.GatewayId = gatewayId
	this.UserId = userId
	return &this
}

// NewUnibeeApiMerchantBalanceUserDetailQueryReqWithDefaults instantiates a new UnibeeApiMerchantBalanceUserDetailQueryReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantBalanceUserDetailQueryReqWithDefaults() *UnibeeApiMerchantBalanceUserDetailQueryReq {
	this := UnibeeApiMerchantBalanceUserDetailQueryReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) GetGatewayId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GatewayId, true
}

// SetGatewayId sets field value
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) SetGatewayId(v int64) {
	o.GatewayId = v
}

// GetUserId returns the UserId field value
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) SetUserId(v int64) {
	o.UserId = v
}

func (o UnibeeApiMerchantBalanceUserDetailQueryReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantBalanceUserDetailQueryReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gatewayId"] = o.GatewayId
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantBalanceUserDetailQueryReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"gatewayId",
		"userId",
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

	varUnibeeApiMerchantBalanceUserDetailQueryReq := _UnibeeApiMerchantBalanceUserDetailQueryReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantBalanceUserDetailQueryReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantBalanceUserDetailQueryReq(varUnibeeApiMerchantBalanceUserDetailQueryReq)

	return err
}

type NullableUnibeeApiMerchantBalanceUserDetailQueryReq struct {
	value *UnibeeApiMerchantBalanceUserDetailQueryReq
	isSet bool
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryReq) Get() *UnibeeApiMerchantBalanceUserDetailQueryReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryReq) Set(val *UnibeeApiMerchantBalanceUserDetailQueryReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantBalanceUserDetailQueryReq(val *UnibeeApiMerchantBalanceUserDetailQueryReq) *NullableUnibeeApiMerchantBalanceUserDetailQueryReq {
	return &NullableUnibeeApiMerchantBalanceUserDetailQueryReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantBalanceUserDetailQueryReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantBalanceUserDetailQueryReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


