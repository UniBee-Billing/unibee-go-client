/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionUpdateRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdateRes{}

// UnibeeApiMerchantSubscriptionUpdateRes struct for UnibeeApiMerchantSubscriptionUpdateRes
type UnibeeApiMerchantSubscriptionUpdateRes struct {
	Link *string `json:"link,omitempty"`
	// note
	Note *string `json:"note,omitempty"`
	Paid *bool `json:"paid,omitempty"`
	SubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"subscriptionPendingUpdate,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionUpdateRes instantiates a new UnibeeApiMerchantSubscriptionUpdateRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUpdateRes() *UnibeeApiMerchantSubscriptionUpdateRes {
	this := UnibeeApiMerchantSubscriptionUpdateRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionUpdateResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdateRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUpdateResWithDefaults() *UnibeeApiMerchantSubscriptionUpdateRes {
	this := UnibeeApiMerchantSubscriptionUpdateRes{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetLink(v string) {
	o.Link = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetNote(v string) {
	o.Note = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasPaid() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetPaid(v bool) {
	o.Paid = &v
}

// GetSubscriptionPendingUpdate returns the SubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.SubscriptionPendingUpdate
}

// GetSubscriptionPendingUpdateOk returns a tuple with the SubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.SubscriptionPendingUpdate) {
		return nil, false
	}
	return o.SubscriptionPendingUpdate, true
}

// HasSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.SubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the SubscriptionPendingUpdate field.
func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.SubscriptionPendingUpdate = &v
}

func (o UnibeeApiMerchantSubscriptionUpdateRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUpdateRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if !IsNil(o.SubscriptionPendingUpdate) {
		toSerialize["subscriptionPendingUpdate"] = o.SubscriptionPendingUpdate
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionUpdateRes struct {
	value *UnibeeApiMerchantSubscriptionUpdateRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateRes) Get() *UnibeeApiMerchantSubscriptionUpdateRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateRes) Set(val *UnibeeApiMerchantSubscriptionUpdateRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUpdateRes(val *UnibeeApiMerchantSubscriptionUpdateRes) *NullableUnibeeApiMerchantSubscriptionUpdateRes {
	return &NullableUnibeeApiMerchantSubscriptionUpdateRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


