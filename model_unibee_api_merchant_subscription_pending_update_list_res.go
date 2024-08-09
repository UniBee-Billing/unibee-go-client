/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionPendingUpdateListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionPendingUpdateListRes{}

// UnibeeApiMerchantSubscriptionPendingUpdateListRes struct for UnibeeApiMerchantSubscriptionPendingUpdateListRes
type UnibeeApiMerchantSubscriptionPendingUpdateListRes struct {
	// SubscriptionPendingUpdateDetails
	SubscriptionPendingUpdateDetails []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"subscriptionPendingUpdateDetails,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionPendingUpdateListRes instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionPendingUpdateListRes() *UnibeeApiMerchantSubscriptionPendingUpdateListRes {
	this := UnibeeApiMerchantSubscriptionPendingUpdateListRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionPendingUpdateListResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionPendingUpdateListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionPendingUpdateListResWithDefaults() *UnibeeApiMerchantSubscriptionPendingUpdateListRes {
	this := UnibeeApiMerchantSubscriptionPendingUpdateListRes{}
	return &this
}

// GetSubscriptionPendingUpdateDetails returns the SubscriptionPendingUpdateDetails field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) GetSubscriptionPendingUpdateDetails() []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.SubscriptionPendingUpdateDetails) {
		var ret []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return o.SubscriptionPendingUpdateDetails
}

// GetSubscriptionPendingUpdateDetailsOk returns a tuple with the SubscriptionPendingUpdateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) GetSubscriptionPendingUpdateDetailsOk() ([]UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.SubscriptionPendingUpdateDetails) {
		return nil, false
	}
	return o.SubscriptionPendingUpdateDetails, true
}

// HasSubscriptionPendingUpdateDetails returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) HasSubscriptionPendingUpdateDetails() bool {
	if o != nil && !IsNil(o.SubscriptionPendingUpdateDetails) {
		return true
	}

	return false
}

// SetSubscriptionPendingUpdateDetails gets a reference to the given []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the SubscriptionPendingUpdateDetails field.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) SetSubscriptionPendingUpdateDetails(v []UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.SubscriptionPendingUpdateDetails = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantSubscriptionPendingUpdateListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionPendingUpdateListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionPendingUpdateDetails) {
		toSerialize["subscriptionPendingUpdateDetails"] = o.SubscriptionPendingUpdateDetails
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes struct {
	value *UnibeeApiMerchantSubscriptionPendingUpdateListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) Get() *UnibeeApiMerchantSubscriptionPendingUpdateListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) Set(val *UnibeeApiMerchantSubscriptionPendingUpdateListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionPendingUpdateListRes(val *UnibeeApiMerchantSubscriptionPendingUpdateListRes) *NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes {
	return &NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionPendingUpdateListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


