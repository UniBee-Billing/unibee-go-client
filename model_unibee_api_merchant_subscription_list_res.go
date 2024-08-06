/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionListRes{}

// UnibeeApiMerchantSubscriptionListRes struct for UnibeeApiMerchantSubscriptionListRes
type UnibeeApiMerchantSubscriptionListRes struct {
	// Subscriptions
	Subscriptions []UnibeeApiBeanDetailSubscriptionDetail `json:"subscriptions,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionListRes instantiates a new UnibeeApiMerchantSubscriptionListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionListRes() *UnibeeApiMerchantSubscriptionListRes {
	this := UnibeeApiMerchantSubscriptionListRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionListResWithDefaults() *UnibeeApiMerchantSubscriptionListRes {
	this := UnibeeApiMerchantSubscriptionListRes{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListRes) GetSubscriptions() []UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.Subscriptions) {
		var ret []UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListRes) GetSubscriptionsOk() ([]UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListRes) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []UnibeeApiBeanDetailSubscriptionDetail and assigns it to the Subscriptions field.
func (o *UnibeeApiMerchantSubscriptionListRes) SetSubscriptions(v []UnibeeApiBeanDetailSubscriptionDetail) {
	o.Subscriptions = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantSubscriptionListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantSubscriptionListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionListRes struct {
	value *UnibeeApiMerchantSubscriptionListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionListRes) Get() *UnibeeApiMerchantSubscriptionListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionListRes) Set(val *UnibeeApiMerchantSubscriptionListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionListRes(val *UnibeeApiMerchantSubscriptionListRes) *NullableUnibeeApiMerchantSubscriptionListRes {
	return &NullableUnibeeApiMerchantSubscriptionListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


