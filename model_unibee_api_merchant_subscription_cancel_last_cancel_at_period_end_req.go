/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq{}

// UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq This action should be request before subscription's period end, If subscription's flag 'cancelAtPeriodEnd' is enabled, this action will resume it to disable, and subscription will continue cycle recurring seems no cancelAtPeriod action be setting
type UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq struct {
	// SubscriptionId, id of subscription, either SubscriptionId or UserId needed, The only one active subscription of userId will effect
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// UserId, either SubscriptionId or UserId needed, The only one active subscription will effect if userId provide instead of subscriptionId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq instantiates a new UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq() *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq {
	this := UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq struct {
	value *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) Get() *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) Set(val *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq(val *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) *NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq {
	return &NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


