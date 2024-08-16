/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionNewAdminNoteReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionNewAdminNoteReq{}

// UnibeeApiMerchantSubscriptionNewAdminNoteReq struct for UnibeeApiMerchantSubscriptionNewAdminNoteReq
type UnibeeApiMerchantSubscriptionNewAdminNoteReq struct {
	// Note
	Note string `json:"note"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionNewAdminNoteReq UnibeeApiMerchantSubscriptionNewAdminNoteReq

// NewUnibeeApiMerchantSubscriptionNewAdminNoteReq instantiates a new UnibeeApiMerchantSubscriptionNewAdminNoteReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionNewAdminNoteReq(note string, subscriptionId string) *UnibeeApiMerchantSubscriptionNewAdminNoteReq {
	this := UnibeeApiMerchantSubscriptionNewAdminNoteReq{}
	this.Note = note
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionNewAdminNoteReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionNewAdminNoteReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionNewAdminNoteReqWithDefaults() *UnibeeApiMerchantSubscriptionNewAdminNoteReq {
	this := UnibeeApiMerchantSubscriptionNewAdminNoteReq{}
	return &this
}

// GetNote returns the Note field value
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) SetNote(v string) {
	o.Note = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionNewAdminNoteReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionNewAdminNoteReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["note"] = o.Note
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionNewAdminNoteReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"note",
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

	varUnibeeApiMerchantSubscriptionNewAdminNoteReq := _UnibeeApiMerchantSubscriptionNewAdminNoteReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionNewAdminNoteReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionNewAdminNoteReq(varUnibeeApiMerchantSubscriptionNewAdminNoteReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq struct {
	value *UnibeeApiMerchantSubscriptionNewAdminNoteReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) Get() *UnibeeApiMerchantSubscriptionNewAdminNoteReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) Set(val *UnibeeApiMerchantSubscriptionNewAdminNoteReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionNewAdminNoteReq(val *UnibeeApiMerchantSubscriptionNewAdminNoteReq) *NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq {
	return &NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionNewAdminNoteReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


