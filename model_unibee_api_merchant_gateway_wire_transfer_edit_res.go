/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantGatewayWireTransferEditRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayWireTransferEditRes{}

// UnibeeApiMerchantGatewayWireTransferEditRes struct for UnibeeApiMerchantGatewayWireTransferEditRes
type UnibeeApiMerchantGatewayWireTransferEditRes struct {
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
}

// NewUnibeeApiMerchantGatewayWireTransferEditRes instantiates a new UnibeeApiMerchantGatewayWireTransferEditRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayWireTransferEditRes() *UnibeeApiMerchantGatewayWireTransferEditRes {
	this := UnibeeApiMerchantGatewayWireTransferEditRes{}
	return &this
}

// NewUnibeeApiMerchantGatewayWireTransferEditResWithDefaults instantiates a new UnibeeApiMerchantGatewayWireTransferEditRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayWireTransferEditResWithDefaults() *UnibeeApiMerchantGatewayWireTransferEditRes {
	this := UnibeeApiMerchantGatewayWireTransferEditRes{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayWireTransferEditRes) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayWireTransferEditRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantGatewayWireTransferEditRes) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

func (o UnibeeApiMerchantGatewayWireTransferEditRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayWireTransferEditRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayWireTransferEditRes struct {
	value *UnibeeApiMerchantGatewayWireTransferEditRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditRes) Get() *UnibeeApiMerchantGatewayWireTransferEditRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditRes) Set(val *UnibeeApiMerchantGatewayWireTransferEditRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayWireTransferEditRes(val *UnibeeApiMerchantGatewayWireTransferEditRes) *NullableUnibeeApiMerchantGatewayWireTransferEditRes {
	return &NullableUnibeeApiMerchantGatewayWireTransferEditRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayWireTransferEditRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayWireTransferEditRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


