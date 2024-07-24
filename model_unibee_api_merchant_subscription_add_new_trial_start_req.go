/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionAddNewTrialStartReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionAddNewTrialStartReq{}

// UnibeeApiMerchantSubscriptionAddNewTrialStartReq struct for UnibeeApiMerchantSubscriptionAddNewTrialStartReq
type UnibeeApiMerchantSubscriptionAddNewTrialStartReq struct {
	// add appendTrialEndHour For Free
	AppendTrialEndHour int64 `json:"appendTrialEndHour"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionAddNewTrialStartReq UnibeeApiMerchantSubscriptionAddNewTrialStartReq

// NewUnibeeApiMerchantSubscriptionAddNewTrialStartReq instantiates a new UnibeeApiMerchantSubscriptionAddNewTrialStartReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionAddNewTrialStartReq(appendTrialEndHour int64, subscriptionId string) *UnibeeApiMerchantSubscriptionAddNewTrialStartReq {
	this := UnibeeApiMerchantSubscriptionAddNewTrialStartReq{}
	this.AppendTrialEndHour = appendTrialEndHour
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionAddNewTrialStartReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionAddNewTrialStartReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionAddNewTrialStartReqWithDefaults() *UnibeeApiMerchantSubscriptionAddNewTrialStartReq {
	this := UnibeeApiMerchantSubscriptionAddNewTrialStartReq{}
	return &this
}

// GetAppendTrialEndHour returns the AppendTrialEndHour field value
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) GetAppendTrialEndHour() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AppendTrialEndHour
}

// GetAppendTrialEndHourOk returns a tuple with the AppendTrialEndHour field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) GetAppendTrialEndHourOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppendTrialEndHour, true
}

// SetAppendTrialEndHour sets field value
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) SetAppendTrialEndHour(v int64) {
	o.AppendTrialEndHour = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionAddNewTrialStartReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionAddNewTrialStartReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appendTrialEndHour"] = o.AppendTrialEndHour
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"appendTrialEndHour",
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

	varUnibeeApiMerchantSubscriptionAddNewTrialStartReq := _UnibeeApiMerchantSubscriptionAddNewTrialStartReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionAddNewTrialStartReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionAddNewTrialStartReq(varUnibeeApiMerchantSubscriptionAddNewTrialStartReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq struct {
	value *UnibeeApiMerchantSubscriptionAddNewTrialStartReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) Get() *UnibeeApiMerchantSubscriptionAddNewTrialStartReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) Set(val *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq(val *UnibeeApiMerchantSubscriptionAddNewTrialStartReq) *NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq {
	return &NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionAddNewTrialStartReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


