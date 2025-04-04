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

// checks if the MerchantSubscriptionTimelineListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionTimelineListGet200ResponseData{}

// MerchantSubscriptionTimelineListGet200ResponseData struct for MerchantSubscriptionTimelineListGet200ResponseData
type MerchantSubscriptionTimelineListGet200ResponseData struct {
	// SubscriptionTimeLines
	SubscriptionTimeLines []UnibeeApiBeanDetailSubscriptionTimeLineDetail `json:"subscriptionTimeLines,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantSubscriptionTimelineListGet200ResponseData instantiates a new MerchantSubscriptionTimelineListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionTimelineListGet200ResponseData() *MerchantSubscriptionTimelineListGet200ResponseData {
	this := MerchantSubscriptionTimelineListGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionTimelineListGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionTimelineListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionTimelineListGet200ResponseDataWithDefaults() *MerchantSubscriptionTimelineListGet200ResponseData {
	this := MerchantSubscriptionTimelineListGet200ResponseData{}
	return &this
}

// GetSubscriptionTimeLines returns the SubscriptionTimeLines field value if set, zero value otherwise.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) GetSubscriptionTimeLines() []UnibeeApiBeanDetailSubscriptionTimeLineDetail {
	if o == nil || IsNil(o.SubscriptionTimeLines) {
		var ret []UnibeeApiBeanDetailSubscriptionTimeLineDetail
		return ret
	}
	return o.SubscriptionTimeLines
}

// GetSubscriptionTimeLinesOk returns a tuple with the SubscriptionTimeLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) GetSubscriptionTimeLinesOk() ([]UnibeeApiBeanDetailSubscriptionTimeLineDetail, bool) {
	if o == nil || IsNil(o.SubscriptionTimeLines) {
		return nil, false
	}
	return o.SubscriptionTimeLines, true
}

// HasSubscriptionTimeLines returns a boolean if a field has been set.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) HasSubscriptionTimeLines() bool {
	if o != nil && !IsNil(o.SubscriptionTimeLines) {
		return true
	}

	return false
}

// SetSubscriptionTimeLines gets a reference to the given []UnibeeApiBeanDetailSubscriptionTimeLineDetail and assigns it to the SubscriptionTimeLines field.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) SetSubscriptionTimeLines(v []UnibeeApiBeanDetailSubscriptionTimeLineDetail) {
	o.SubscriptionTimeLines = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantSubscriptionTimelineListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantSubscriptionTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionTimelineListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionTimeLines) {
		toSerialize["subscriptionTimeLines"] = o.SubscriptionTimeLines
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionTimelineListGet200ResponseData struct {
	value *MerchantSubscriptionTimelineListGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionTimelineListGet200ResponseData) Get() *MerchantSubscriptionTimelineListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionTimelineListGet200ResponseData) Set(val *MerchantSubscriptionTimelineListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionTimelineListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionTimelineListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionTimelineListGet200ResponseData(val *MerchantSubscriptionTimelineListGet200ResponseData) *NullableMerchantSubscriptionTimelineListGet200ResponseData {
	return &NullableMerchantSubscriptionTimelineListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionTimelineListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionTimelineListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


