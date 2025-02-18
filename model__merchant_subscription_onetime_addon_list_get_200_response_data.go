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

// checks if the MerchantSubscriptionOnetimeAddonListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionOnetimeAddonListGet200ResponseData{}

// MerchantSubscriptionOnetimeAddonListGet200ResponseData struct for MerchantSubscriptionOnetimeAddonListGet200ResponseData
type MerchantSubscriptionOnetimeAddonListGet200ResponseData struct {
	// SubscriptionOnetimeAddons
	SubscriptionOnetimeAddons []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail `json:"subscriptionOnetimeAddons,omitempty"`
}

// NewMerchantSubscriptionOnetimeAddonListGet200ResponseData instantiates a new MerchantSubscriptionOnetimeAddonListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionOnetimeAddonListGet200ResponseData() *MerchantSubscriptionOnetimeAddonListGet200ResponseData {
	this := MerchantSubscriptionOnetimeAddonListGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionOnetimeAddonListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionOnetimeAddonListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionOnetimeAddonListGet200ResponseDataWithDefaults() *MerchantSubscriptionOnetimeAddonListGet200ResponseData {
	this := MerchantSubscriptionOnetimeAddonListGet200ResponseData{}
	return &this
}

// GetSubscriptionOnetimeAddons returns the SubscriptionOnetimeAddons field value if set, zero value otherwise.
func (o *MerchantSubscriptionOnetimeAddonListGet200ResponseData) GetSubscriptionOnetimeAddons() []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	if o == nil || IsNil(o.SubscriptionOnetimeAddons) {
		var ret []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail
		return ret
	}
	return o.SubscriptionOnetimeAddons
}

// GetSubscriptionOnetimeAddonsOk returns a tuple with the SubscriptionOnetimeAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionOnetimeAddonListGet200ResponseData) GetSubscriptionOnetimeAddonsOk() ([]UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail, bool) {
	if o == nil || IsNil(o.SubscriptionOnetimeAddons) {
		return nil, false
	}
	return o.SubscriptionOnetimeAddons, true
}

// HasSubscriptionOnetimeAddons returns a boolean if a field has been set.
func (o *MerchantSubscriptionOnetimeAddonListGet200ResponseData) HasSubscriptionOnetimeAddons() bool {
	if o != nil && !IsNil(o.SubscriptionOnetimeAddons) {
		return true
	}

	return false
}

// SetSubscriptionOnetimeAddons gets a reference to the given []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail and assigns it to the SubscriptionOnetimeAddons field.
func (o *MerchantSubscriptionOnetimeAddonListGet200ResponseData) SetSubscriptionOnetimeAddons(v []UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) {
	o.SubscriptionOnetimeAddons = v
}

func (o MerchantSubscriptionOnetimeAddonListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionOnetimeAddonListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionOnetimeAddons) {
		toSerialize["subscriptionOnetimeAddons"] = o.SubscriptionOnetimeAddons
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData struct {
	value *MerchantSubscriptionOnetimeAddonListGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) Get() *MerchantSubscriptionOnetimeAddonListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) Set(val *MerchantSubscriptionOnetimeAddonListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionOnetimeAddonListGet200ResponseData(val *MerchantSubscriptionOnetimeAddonListGet200ResponseData) *NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData {
	return &NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionOnetimeAddonListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


