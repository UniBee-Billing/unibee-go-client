/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantGatewayDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayDetailRes{}

// UnibeeApiMerchantGatewayDetailRes struct for UnibeeApiMerchantGatewayDetailRes
type UnibeeApiMerchantGatewayDetailRes struct {
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
}

// NewUnibeeApiMerchantGatewayDetailRes instantiates a new UnibeeApiMerchantGatewayDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayDetailRes() *UnibeeApiMerchantGatewayDetailRes {
	this := UnibeeApiMerchantGatewayDetailRes{}
	return &this
}

// NewUnibeeApiMerchantGatewayDetailResWithDefaults instantiates a new UnibeeApiMerchantGatewayDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayDetailResWithDefaults() *UnibeeApiMerchantGatewayDetailRes {
	this := UnibeeApiMerchantGatewayDetailRes{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayDetailRes) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayDetailRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayDetailRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantGatewayDetailRes) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

func (o UnibeeApiMerchantGatewayDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayDetailRes struct {
	value *UnibeeApiMerchantGatewayDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayDetailRes) Get() *UnibeeApiMerchantGatewayDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayDetailRes) Set(val *UnibeeApiMerchantGatewayDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayDetailRes(val *UnibeeApiMerchantGatewayDetailRes) *NullableUnibeeApiMerchantGatewayDetailRes {
	return &NullableUnibeeApiMerchantGatewayDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


