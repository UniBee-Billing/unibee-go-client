/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionOnetimeAddonListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionOnetimeAddonListRes{}

// UnibeeApiMerchantSubscriptionOnetimeAddonListRes struct for UnibeeApiMerchantSubscriptionOnetimeAddonListRes
type UnibeeApiMerchantSubscriptionOnetimeAddonListRes struct {
	// SubscriptionOnetimeAddons
	SubscriptionOnetimeAddons []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail `json:"subscriptionOnetimeAddons,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonListRes instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionOnetimeAddonListRes() *UnibeeApiMerchantSubscriptionOnetimeAddonListRes {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonListRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionOnetimeAddonListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionOnetimeAddonListResWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonListRes {
	this := UnibeeApiMerchantSubscriptionOnetimeAddonListRes{}
	return &this
}

// GetSubscriptionOnetimeAddons returns the SubscriptionOnetimeAddons field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) GetSubscriptionOnetimeAddons() []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	if o == nil || IsNil(o.SubscriptionOnetimeAddons) {
		var ret []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail
		return ret
	}
	return o.SubscriptionOnetimeAddons
}

// GetSubscriptionOnetimeAddonsOk returns a tuple with the SubscriptionOnetimeAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) GetSubscriptionOnetimeAddonsOk() ([]UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail, bool) {
	if o == nil || IsNil(o.SubscriptionOnetimeAddons) {
		return nil, false
	}
	return o.SubscriptionOnetimeAddons, true
}

// HasSubscriptionOnetimeAddons returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) HasSubscriptionOnetimeAddons() bool {
	if o != nil && !IsNil(o.SubscriptionOnetimeAddons) {
		return true
	}

	return false
}

// SetSubscriptionOnetimeAddons gets a reference to the given []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail and assigns it to the SubscriptionOnetimeAddons field.
func (o *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) SetSubscriptionOnetimeAddons(v []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) {
	o.SubscriptionOnetimeAddons = v
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionOnetimeAddonListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionOnetimeAddons) {
		toSerialize["subscriptionOnetimeAddons"] = o.SubscriptionOnetimeAddons
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes struct {
	value *UnibeeApiMerchantSubscriptionOnetimeAddonListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) Get() *UnibeeApiMerchantSubscriptionOnetimeAddonListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) Set(val *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes(val *UnibeeApiMerchantSubscriptionOnetimeAddonListRes) *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes {
	return &NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionOnetimeAddonListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


