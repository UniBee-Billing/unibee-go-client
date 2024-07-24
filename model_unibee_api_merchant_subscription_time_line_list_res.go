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

// checks if the UnibeeApiMerchantSubscriptionTimeLineListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionTimeLineListRes{}

// UnibeeApiMerchantSubscriptionTimeLineListRes struct for UnibeeApiMerchantSubscriptionTimeLineListRes
type UnibeeApiMerchantSubscriptionTimeLineListRes struct {
	// SubscriptionTimeLines
	SubscriptionTimeLines []UnibeeApiBeanDetailSubscriptionTimeLineDetail `json:"subscriptionTimeLines,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionTimeLineListRes instantiates a new UnibeeApiMerchantSubscriptionTimeLineListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionTimeLineListRes() *UnibeeApiMerchantSubscriptionTimeLineListRes {
	this := UnibeeApiMerchantSubscriptionTimeLineListRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionTimeLineListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionTimeLineListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionTimeLineListResWithDefaults() *UnibeeApiMerchantSubscriptionTimeLineListRes {
	this := UnibeeApiMerchantSubscriptionTimeLineListRes{}
	return &this
}

// GetSubscriptionTimeLines returns the SubscriptionTimeLines field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) GetSubscriptionTimeLines() []UnibeeApiBeanDetailSubscriptionTimeLineDetail {
	if o == nil || IsNil(o.SubscriptionTimeLines) {
		var ret []UnibeeApiBeanDetailSubscriptionTimeLineDetail
		return ret
	}
	return o.SubscriptionTimeLines
}

// GetSubscriptionTimeLinesOk returns a tuple with the SubscriptionTimeLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) GetSubscriptionTimeLinesOk() ([]UnibeeApiBeanDetailSubscriptionTimeLineDetail, bool) {
	if o == nil || IsNil(o.SubscriptionTimeLines) {
		return nil, false
	}
	return o.SubscriptionTimeLines, true
}

// HasSubscriptionTimeLines returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) HasSubscriptionTimeLines() bool {
	if o != nil && !IsNil(o.SubscriptionTimeLines) {
		return true
	}

	return false
}

// SetSubscriptionTimeLines gets a reference to the given []UnibeeApiBeanDetailSubscriptionTimeLineDetail and assigns it to the SubscriptionTimeLines field.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) SetSubscriptionTimeLines(v []UnibeeApiBeanDetailSubscriptionTimeLineDetail) {
	o.SubscriptionTimeLines = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantSubscriptionTimeLineListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantSubscriptionTimeLineListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionTimeLineListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionTimeLines) {
		toSerialize["subscriptionTimeLines"] = o.SubscriptionTimeLines
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionTimeLineListRes struct {
	value *UnibeeApiMerchantSubscriptionTimeLineListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionTimeLineListRes) Get() *UnibeeApiMerchantSubscriptionTimeLineListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionTimeLineListRes) Set(val *UnibeeApiMerchantSubscriptionTimeLineListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionTimeLineListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionTimeLineListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionTimeLineListRes(val *UnibeeApiMerchantSubscriptionTimeLineListRes) *NullableUnibeeApiMerchantSubscriptionTimeLineListRes {
	return &NullableUnibeeApiMerchantSubscriptionTimeLineListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionTimeLineListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionTimeLineListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


