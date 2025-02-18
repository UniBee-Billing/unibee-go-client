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

// checks if the UnibeeApiMerchantGatewayDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantGatewayDetailReq{}

// UnibeeApiMerchantGatewayDetailReq Get Payment Gateway Detail
type UnibeeApiMerchantGatewayDetailReq struct {
	// The id of payment gateway, either gatewayId or gatewayName
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// The name of payment gateway, , either gatewayId or gatewayName, stripe|paypal|changelly|unitpay|payssion|cryptadium
	GatewayName *string `json:"gatewayName,omitempty"`
}

// NewUnibeeApiMerchantGatewayDetailReq instantiates a new UnibeeApiMerchantGatewayDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantGatewayDetailReq() *UnibeeApiMerchantGatewayDetailReq {
	this := UnibeeApiMerchantGatewayDetailReq{}
	return &this
}

// NewUnibeeApiMerchantGatewayDetailReqWithDefaults instantiates a new UnibeeApiMerchantGatewayDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantGatewayDetailReqWithDefaults() *UnibeeApiMerchantGatewayDetailReq {
	this := UnibeeApiMerchantGatewayDetailReq{}
	return &this
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayDetailReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantGatewayDetailReq) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantGatewayDetailReq) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantGatewayDetailReq) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *UnibeeApiMerchantGatewayDetailReq) SetGatewayName(v string) {
	o.GatewayName = &v
}

func (o UnibeeApiMerchantGatewayDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantGatewayDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gatewayName"] = o.GatewayName
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantGatewayDetailReq struct {
	value *UnibeeApiMerchantGatewayDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantGatewayDetailReq) Get() *UnibeeApiMerchantGatewayDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantGatewayDetailReq) Set(val *UnibeeApiMerchantGatewayDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantGatewayDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantGatewayDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantGatewayDetailReq(val *UnibeeApiMerchantGatewayDetailReq) *NullableUnibeeApiMerchantGatewayDetailReq {
	return &NullableUnibeeApiMerchantGatewayDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantGatewayDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantGatewayDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


