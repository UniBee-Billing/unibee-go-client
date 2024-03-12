/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantBalanceDetailQueryRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantBalanceDetailQueryRes{}

// UnibeeApiMerchantBalanceDetailQueryRes struct for UnibeeApiMerchantBalanceDetailQueryRes
type UnibeeApiMerchantBalanceDetailQueryRes struct {
	Available []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"available,omitempty"`
	ConnectReserved []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"connectReserved,omitempty"`
	Pending []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"pending,omitempty"`
}

// NewUnibeeApiMerchantBalanceDetailQueryRes instantiates a new UnibeeApiMerchantBalanceDetailQueryRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantBalanceDetailQueryRes() *UnibeeApiMerchantBalanceDetailQueryRes {
	this := UnibeeApiMerchantBalanceDetailQueryRes{}
	return &this
}

// NewUnibeeApiMerchantBalanceDetailQueryResWithDefaults instantiates a new UnibeeApiMerchantBalanceDetailQueryRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantBalanceDetailQueryResWithDefaults() *UnibeeApiMerchantBalanceDetailQueryRes {
	this := UnibeeApiMerchantBalanceDetailQueryRes{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetAvailable() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Available) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetAvailableOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Available field.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) SetAvailable(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Available = v
}

// GetConnectReserved returns the ConnectReserved field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetConnectReserved() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.ConnectReserved) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.ConnectReserved
}

// GetConnectReservedOk returns a tuple with the ConnectReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetConnectReservedOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.ConnectReserved) {
		return nil, false
	}
	return o.ConnectReserved, true
}

// HasConnectReserved returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) HasConnectReserved() bool {
	if o != nil && !IsNil(o.ConnectReserved) {
		return true
	}

	return false
}

// SetConnectReserved gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the ConnectReserved field.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) SetConnectReserved(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.ConnectReserved = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetPending() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Pending) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) GetPendingOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Pending field.
func (o *UnibeeApiMerchantBalanceDetailQueryRes) SetPending(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Pending = v
}

func (o UnibeeApiMerchantBalanceDetailQueryRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantBalanceDetailQueryRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.ConnectReserved) {
		toSerialize["connectReserved"] = o.ConnectReserved
	}
	if !IsNil(o.Pending) {
		toSerialize["pending"] = o.Pending
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantBalanceDetailQueryRes struct {
	value *UnibeeApiMerchantBalanceDetailQueryRes
	isSet bool
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryRes) Get() *UnibeeApiMerchantBalanceDetailQueryRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryRes) Set(val *UnibeeApiMerchantBalanceDetailQueryRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantBalanceDetailQueryRes(val *UnibeeApiMerchantBalanceDetailQueryRes) *NullableUnibeeApiMerchantBalanceDetailQueryRes {
	return &NullableUnibeeApiMerchantBalanceDetailQueryRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantBalanceDetailQueryRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantBalanceDetailQueryRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


