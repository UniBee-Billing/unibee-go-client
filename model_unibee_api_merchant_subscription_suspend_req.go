/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionSuspendReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionSuspendReq{}

// UnibeeApiMerchantSubscriptionSuspendReq struct for UnibeeApiMerchantSubscriptionSuspendReq
type UnibeeApiMerchantSubscriptionSuspendReq struct {
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionSuspendReq UnibeeApiMerchantSubscriptionSuspendReq

// NewUnibeeApiMerchantSubscriptionSuspendReq instantiates a new UnibeeApiMerchantSubscriptionSuspendReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionSuspendReq(subscriptionId string) *UnibeeApiMerchantSubscriptionSuspendReq {
	this := UnibeeApiMerchantSubscriptionSuspendReq{}
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionSuspendReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionSuspendReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionSuspendReqWithDefaults() *UnibeeApiMerchantSubscriptionSuspendReq {
	this := UnibeeApiMerchantSubscriptionSuspendReq{}
	return &this
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionSuspendReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionSuspendReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionSuspendReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionSuspendReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionSuspendReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionSuspendReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUnibeeApiMerchantSubscriptionSuspendReq := _UnibeeApiMerchantSubscriptionSuspendReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionSuspendReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionSuspendReq(varUnibeeApiMerchantSubscriptionSuspendReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionSuspendReq struct {
	value *UnibeeApiMerchantSubscriptionSuspendReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionSuspendReq) Get() *UnibeeApiMerchantSubscriptionSuspendReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionSuspendReq) Set(val *UnibeeApiMerchantSubscriptionSuspendReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionSuspendReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionSuspendReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionSuspendReq(val *UnibeeApiMerchantSubscriptionSuspendReq) *NullableUnibeeApiMerchantSubscriptionSuspendReq {
	return &NullableUnibeeApiMerchantSubscriptionSuspendReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionSuspendReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionSuspendReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


