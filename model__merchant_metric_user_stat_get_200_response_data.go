/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MerchantMetricUserStatGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricUserStatGet200ResponseData{}

// MerchantMetricUserStatGet200ResponseData struct for MerchantMetricUserStatGet200ResponseData
type MerchantMetricUserStatGet200ResponseData struct {
	UserMetricStat *UnibeeInternalLogicGatewayRoUserMetricStat `json:"userMetricStat,omitempty"`
}

// NewMerchantMetricUserStatGet200ResponseData instantiates a new MerchantMetricUserStatGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricUserStatGet200ResponseData() *MerchantMetricUserStatGet200ResponseData {
	this := MerchantMetricUserStatGet200ResponseData{}
	return &this
}

// NewMerchantMetricUserStatGet200ResponseDataWithDefaults instantiates a new MerchantMetricUserStatGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricUserStatGet200ResponseDataWithDefaults() *MerchantMetricUserStatGet200ResponseData {
	this := MerchantMetricUserStatGet200ResponseData{}
	return &this
}

// GetUserMetricStat returns the UserMetricStat field value if set, zero value otherwise.
func (o *MerchantMetricUserStatGet200ResponseData) GetUserMetricStat() UnibeeInternalLogicGatewayRoUserMetricStat {
	if o == nil || IsNil(o.UserMetricStat) {
		var ret UnibeeInternalLogicGatewayRoUserMetricStat
		return ret
	}
	return *o.UserMetricStat
}

// GetUserMetricStatOk returns a tuple with the UserMetricStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserStatGet200ResponseData) GetUserMetricStatOk() (*UnibeeInternalLogicGatewayRoUserMetricStat, bool) {
	if o == nil || IsNil(o.UserMetricStat) {
		return nil, false
	}
	return o.UserMetricStat, true
}

// HasUserMetricStat returns a boolean if a field has been set.
func (o *MerchantMetricUserStatGet200ResponseData) HasUserMetricStat() bool {
	if o != nil && !IsNil(o.UserMetricStat) {
		return true
	}

	return false
}

// SetUserMetricStat gets a reference to the given UnibeeInternalLogicGatewayRoUserMetricStat and assigns it to the UserMetricStat field.
func (o *MerchantMetricUserStatGet200ResponseData) SetUserMetricStat(v UnibeeInternalLogicGatewayRoUserMetricStat) {
	o.UserMetricStat = &v
}

func (o MerchantMetricUserStatGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricUserStatGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserMetricStat) {
		toSerialize["userMetricStat"] = o.UserMetricStat
	}
	return toSerialize, nil
}

type NullableMerchantMetricUserStatGet200ResponseData struct {
	value *MerchantMetricUserStatGet200ResponseData
	isSet bool
}

func (v NullableMerchantMetricUserStatGet200ResponseData) Get() *MerchantMetricUserStatGet200ResponseData {
	return v.value
}

func (v *NullableMerchantMetricUserStatGet200ResponseData) Set(val *MerchantMetricUserStatGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricUserStatGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricUserStatGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricUserStatGet200ResponseData(val *MerchantMetricUserStatGet200ResponseData) *NullableMerchantMetricUserStatGet200ResponseData {
	return &NullableMerchantMetricUserStatGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMetricUserStatGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricUserStatGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


