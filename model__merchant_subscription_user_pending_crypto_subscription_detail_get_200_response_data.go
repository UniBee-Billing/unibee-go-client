/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData{}

// MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData struct for MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData
type MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData struct {
	Subscription *UnibeeApiBeanDetailSubscriptionDetail `json:"subscription,omitempty"`
}

// NewMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData instantiates a new MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData() *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseDataWithDefaults() *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanDetailSubscriptionDetail and assigns it to the Subscription field.
func (o *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanDetailSubscriptionDetail) {
	o.Subscription = &v
}

func (o MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData struct {
	value *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) Get() *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) Set(val *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData(val *MerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) *NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData {
	return &NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionUserPendingCryptoSubscriptionDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


