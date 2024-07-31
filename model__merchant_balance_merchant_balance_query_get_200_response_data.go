/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantBalanceMerchantBalanceQueryGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantBalanceMerchantBalanceQueryGet200ResponseData{}

// MerchantBalanceMerchantBalanceQueryGet200ResponseData struct for MerchantBalanceMerchantBalanceQueryGet200ResponseData
type MerchantBalanceMerchantBalanceQueryGet200ResponseData struct {
	Available []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"available,omitempty"`
	ConnectReserved []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"connectReserved,omitempty"`
	Pending []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance `json:"pending,omitempty"`
}

// NewMerchantBalanceMerchantBalanceQueryGet200ResponseData instantiates a new MerchantBalanceMerchantBalanceQueryGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantBalanceMerchantBalanceQueryGet200ResponseData() *MerchantBalanceMerchantBalanceQueryGet200ResponseData {
	this := MerchantBalanceMerchantBalanceQueryGet200ResponseData{}
	return &this
}

// NewMerchantBalanceMerchantBalanceQueryGet200ResponseDataWithDefaults instantiates a new MerchantBalanceMerchantBalanceQueryGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantBalanceMerchantBalanceQueryGet200ResponseDataWithDefaults() *MerchantBalanceMerchantBalanceQueryGet200ResponseData {
	this := MerchantBalanceMerchantBalanceQueryGet200ResponseData{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetAvailable() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Available) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetAvailableOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Available field.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) SetAvailable(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Available = v
}

// GetConnectReserved returns the ConnectReserved field value if set, zero value otherwise.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetConnectReserved() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.ConnectReserved) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.ConnectReserved
}

// GetConnectReservedOk returns a tuple with the ConnectReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetConnectReservedOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.ConnectReserved) {
		return nil, false
	}
	return o.ConnectReserved, true
}

// HasConnectReserved returns a boolean if a field has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) HasConnectReserved() bool {
	if o != nil && !IsNil(o.ConnectReserved) {
		return true
	}

	return false
}

// SetConnectReserved gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the ConnectReserved field.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) SetConnectReserved(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.ConnectReserved = v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetPending() []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance {
	if o == nil || IsNil(o.Pending) {
		var ret []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance
		return ret
	}
	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) GetPendingOk() ([]UnibeeInternalLogicGatewayGatewayBeanGatewayBalance, bool) {
	if o == nil || IsNil(o.Pending) {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) HasPending() bool {
	if o != nil && !IsNil(o.Pending) {
		return true
	}

	return false
}

// SetPending gets a reference to the given []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance and assigns it to the Pending field.
func (o *MerchantBalanceMerchantBalanceQueryGet200ResponseData) SetPending(v []UnibeeInternalLogicGatewayGatewayBeanGatewayBalance) {
	o.Pending = v
}

func (o MerchantBalanceMerchantBalanceQueryGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantBalanceMerchantBalanceQueryGet200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData struct {
	value *MerchantBalanceMerchantBalanceQueryGet200ResponseData
	isSet bool
}

func (v NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) Get() *MerchantBalanceMerchantBalanceQueryGet200ResponseData {
	return v.value
}

func (v *NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) Set(val *MerchantBalanceMerchantBalanceQueryGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantBalanceMerchantBalanceQueryGet200ResponseData(val *MerchantBalanceMerchantBalanceQueryGet200ResponseData) *NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData {
	return &NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantBalanceMerchantBalanceQueryGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


