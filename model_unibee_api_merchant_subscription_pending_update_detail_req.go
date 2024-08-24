/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionPendingUpdateDetailReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionPendingUpdateDetailReq{}

// UnibeeApiMerchantSubscriptionPendingUpdateDetailReq struct for UnibeeApiMerchantSubscriptionPendingUpdateDetailReq
type UnibeeApiMerchantSubscriptionPendingUpdateDetailReq struct {
	// SubscriptionPendingUpdateId
	SubscriptionPendingUpdateId string `json:"subscriptionPendingUpdateId"`
}

type _UnibeeApiMerchantSubscriptionPendingUpdateDetailReq UnibeeApiMerchantSubscriptionPendingUpdateDetailReq

// NewUnibeeApiMerchantSubscriptionPendingUpdateDetailReq instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateDetailReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionPendingUpdateDetailReq(subscriptionPendingUpdateId string) *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq {
	this := UnibeeApiMerchantSubscriptionPendingUpdateDetailReq{}
	this.SubscriptionPendingUpdateId = subscriptionPendingUpdateId
	return &this
}

// NewUnibeeApiMerchantSubscriptionPendingUpdateDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateDetailReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionPendingUpdateDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq {
	this := UnibeeApiMerchantSubscriptionPendingUpdateDetailReq{}
	return &this
}

// GetSubscriptionPendingUpdateId returns the SubscriptionPendingUpdateId field value
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) GetSubscriptionPendingUpdateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionPendingUpdateId
}

// GetSubscriptionPendingUpdateIdOk returns a tuple with the SubscriptionPendingUpdateId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) GetSubscriptionPendingUpdateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionPendingUpdateId, true
}

// SetSubscriptionPendingUpdateId sets field value
func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) SetSubscriptionPendingUpdateId(v string) {
	o.SubscriptionPendingUpdateId = v
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionPendingUpdateId"] = o.SubscriptionPendingUpdateId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionPendingUpdateId",
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

	varUnibeeApiMerchantSubscriptionPendingUpdateDetailReq := _UnibeeApiMerchantSubscriptionPendingUpdateDetailReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionPendingUpdateDetailReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionPendingUpdateDetailReq(varUnibeeApiMerchantSubscriptionPendingUpdateDetailReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq struct {
	value *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) Get() *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) Set(val *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq(val *UnibeeApiMerchantSubscriptionPendingUpdateDetailReq) *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq {
	return &NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateDetailReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


