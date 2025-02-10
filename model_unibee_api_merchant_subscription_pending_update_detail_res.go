/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionPendingUpdateDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionPendingUpdateDetailRes{}

// UnibeeApiMerchantSubscriptionPendingUpdateDetailRes struct for UnibeeApiMerchantSubscriptionPendingUpdateDetailRes
type UnibeeApiMerchantSubscriptionPendingUpdateDetailRes struct {
	SubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"SubscriptionPendingUpdate,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionPendingUpdateDetailRes instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionPendingUpdateDetailRes() *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes {
	this := UnibeeApiMerchantSubscriptionPendingUpdateDetailRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionPendingUpdateDetailResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionPendingUpdateDetailResWithDefaults() *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes {
	this := UnibeeApiMerchantSubscriptionPendingUpdateDetailRes{}
	return &this
}

// GetSubscriptionPendingUpdate returns the SubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) GetSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.SubscriptionPendingUpdate
}

// GetSubscriptionPendingUpdateOk returns a tuple with the SubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) GetSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		return nil, false
	}
	return o.SubscriptionPendingUpdate, true
}

// HasSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) HasSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.SubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the SubscriptionPendingUpdate field.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) SetSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.SubscriptionPendingUpdate = &v
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionPendingUpdate) {
		toSerialize["SubscriptionPendingUpdate"] = o.SubscriptionPendingUpdate
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes struct {
	value *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) Get() *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) Set(val *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes(val *UnibeeApiMerchantSubscriptionPendingUpdateDetailRes) *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes {
	return &NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


