/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes{}

// UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes struct for UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes
type UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes struct {
	Subscription *UnibeeApiBeanDetailSubscriptionDetail `json:"subscription,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes instantiates a new UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes {
	this := UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailResWithDefaults() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes {
	this := UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) GetSubscription() UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) GetSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanDetailSubscriptionDetail and assigns it to the Subscription field.
func (o *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) SetSubscription(v UnibeeApiBeanDetailSubscriptionDetail) {
	o.Subscription = &v
}

func (o UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes struct {
	value *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) Get() *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) Set(val *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes(val *UnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes {
	return &NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUserPendingCryptoSubscriptionDetailRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


