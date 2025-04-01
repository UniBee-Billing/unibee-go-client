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

// checks if the UnibeeApiMerchantGatewayListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayListRes{}

// UnibeeApiMerchantGatewayListRes struct for UnibeeApiMerchantGatewayListRes
type UnibeeApiMerchantGatewayListRes struct {
	// Payment Gateway Object List
	Gateways []UnibeeApiBeanDetailGateway `json:"gateways,omitempty"`
}

// NewUnibeeApiMerchantGatewayListRes instantiates a new UnibeeApiMerchantGatewayListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayListRes() *UnibeeApiMerchantGatewayListRes {
	this := UnibeeApiMerchantGatewayListRes{}
	return &this
}

// NewUnibeeApiMerchantGatewayListResWithDefaults instantiates a new UnibeeApiMerchantGatewayListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayListResWithDefaults() *UnibeeApiMerchantGatewayListRes {
	this := UnibeeApiMerchantGatewayListRes{}
	return &this
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayListRes) GetGateways() []UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateways) {
		var ret []UnibeeApiBeanDetailGateway
		return ret
	}
	return o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayListRes) GetGatewaysOk() ([]UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayListRes) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given []UnibeeApiBeanDetailGateway and assigns it to the Gateways field.
func (o *UnibeeApiMerchantGatewayListRes) SetGateways(v []UnibeeApiBeanDetailGateway) {
	o.Gateways = v
}

func (o UnibeeApiMerchantGatewayListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayListRes struct {
	value *UnibeeApiMerchantGatewayListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayListRes) Get() *UnibeeApiMerchantGatewayListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayListRes) Set(val *UnibeeApiMerchantGatewayListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayListRes(val *UnibeeApiMerchantGatewayListRes) *NullableUnibeeApiMerchantGatewayListRes {
	return &NullableUnibeeApiMerchantGatewayListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


