/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionActiveTemporarilyReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionActiveTemporarilyReq{}

// UnibeeApiMerchantSubscriptionActiveTemporarilyReq Subscription active temporarily, status will transmit from pending to incomplete
type UnibeeApiMerchantSubscriptionActiveTemporarilyReq struct {
	// ExpireTime, the expire utc time if not paid
	ExpireTime int64 `json:"expireTime"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionActiveTemporarilyReq UnibeeApiMerchantSubscriptionActiveTemporarilyReq

// NewUnibeeApiMerchantSubscriptionActiveTemporarilyReq instantiates a new UnibeeApiMerchantSubscriptionActiveTemporarilyReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionActiveTemporarilyReq(expireTime int64, subscriptionId string) *UnibeeApiMerchantSubscriptionActiveTemporarilyReq {
	this := UnibeeApiMerchantSubscriptionActiveTemporarilyReq{}
	this.ExpireTime = expireTime
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionActiveTemporarilyReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionActiveTemporarilyReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionActiveTemporarilyReqWithDefaults() *UnibeeApiMerchantSubscriptionActiveTemporarilyReq {
	this := UnibeeApiMerchantSubscriptionActiveTemporarilyReq{}
	return &this
}

// GetExpireTime returns the ExpireTime field value
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) GetExpireTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpireTime
}

// GetExpireTimeOk returns a tuple with the ExpireTime field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) GetExpireTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpireTime, true
}

// SetExpireTime sets field value
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) SetExpireTime(v int64) {
	o.ExpireTime = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionActiveTemporarilyReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionActiveTemporarilyReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expireTime"] = o.ExpireTime
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expireTime",
		"subscriptionId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantSubscriptionActiveTemporarilyReq := _UnibeeApiMerchantSubscriptionActiveTemporarilyReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionActiveTemporarilyReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionActiveTemporarilyReq(varUnibeeApiMerchantSubscriptionActiveTemporarilyReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq struct {
	value *UnibeeApiMerchantSubscriptionActiveTemporarilyReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) Get() *UnibeeApiMerchantSubscriptionActiveTemporarilyReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) Set(val *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq(val *UnibeeApiMerchantSubscriptionActiveTemporarilyReq) *NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq {
	return &NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionActiveTemporarilyReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


