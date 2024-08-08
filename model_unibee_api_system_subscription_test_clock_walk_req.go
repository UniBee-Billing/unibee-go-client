/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiSystemSubscriptionTestClockWalkReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemSubscriptionTestClockWalkReq{}

// UnibeeApiSystemSubscriptionTestClockWalkReq struct for UnibeeApiSystemSubscriptionTestClockWalkReq
type UnibeeApiSystemSubscriptionTestClockWalkReq struct {
	// NewTestClock
	NewTestClock int64 `json:"newTestClock"`
	// Subscription Id
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiSystemSubscriptionTestClockWalkReq UnibeeApiSystemSubscriptionTestClockWalkReq

// NewUnibeeApiSystemSubscriptionTestClockWalkReq instantiates a new UnibeeApiSystemSubscriptionTestClockWalkReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemSubscriptionTestClockWalkReq(newTestClock int64, subscriptionId string) *UnibeeApiSystemSubscriptionTestClockWalkReq {
	this := UnibeeApiSystemSubscriptionTestClockWalkReq{}
	this.NewTestClock = newTestClock
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiSystemSubscriptionTestClockWalkReqWithDefaults instantiates a new UnibeeApiSystemSubscriptionTestClockWalkReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemSubscriptionTestClockWalkReqWithDefaults() *UnibeeApiSystemSubscriptionTestClockWalkReq {
	this := UnibeeApiSystemSubscriptionTestClockWalkReq{}
	return &this
}

// GetNewTestClock returns the NewTestClock field value
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) GetNewTestClock() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewTestClock
}

// GetNewTestClockOk returns a tuple with the NewTestClock field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) GetNewTestClockOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewTestClock, true
}

// SetNewTestClock sets field value
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) SetNewTestClock(v int64) {
	o.NewTestClock = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiSystemSubscriptionTestClockWalkReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemSubscriptionTestClockWalkReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newTestClock"] = o.NewTestClock
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiSystemSubscriptionTestClockWalkReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newTestClock",
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

	varUnibeeApiSystemSubscriptionTestClockWalkReq := _UnibeeApiSystemSubscriptionTestClockWalkReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiSystemSubscriptionTestClockWalkReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiSystemSubscriptionTestClockWalkReq(varUnibeeApiSystemSubscriptionTestClockWalkReq)

	return err
}

type NullableUnibeeApiSystemSubscriptionTestClockWalkReq struct {
	value *UnibeeApiSystemSubscriptionTestClockWalkReq
	isSet bool
}

func (v NullableUnibeeApiSystemSubscriptionTestClockWalkReq) Get() *UnibeeApiSystemSubscriptionTestClockWalkReq {
	return v.value
}

func (v *NullableUnibeeApiSystemSubscriptionTestClockWalkReq) Set(val *UnibeeApiSystemSubscriptionTestClockWalkReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemSubscriptionTestClockWalkReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemSubscriptionTestClockWalkReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemSubscriptionTestClockWalkReq(val *UnibeeApiSystemSubscriptionTestClockWalkReq) *NullableUnibeeApiSystemSubscriptionTestClockWalkReq {
	return &NullableUnibeeApiSystemSubscriptionTestClockWalkReq{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemSubscriptionTestClockWalkReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemSubscriptionTestClockWalkReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


