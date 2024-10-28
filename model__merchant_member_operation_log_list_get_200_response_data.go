/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMemberOperationLogListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMemberOperationLogListGet200ResponseData{}

// MerchantMemberOperationLogListGet200ResponseData struct for MerchantMemberOperationLogListGet200ResponseData
type MerchantMemberOperationLogListGet200ResponseData struct {
	// Merchant Member Operation Log List
	MerchantOperationLogs []UnibeeApiBeanDetailMerchantOperationLogDetail `json:"merchantOperationLogs,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantMemberOperationLogListGet200ResponseData instantiates a new MerchantMemberOperationLogListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMemberOperationLogListGet200ResponseData() *MerchantMemberOperationLogListGet200ResponseData {
	this := MerchantMemberOperationLogListGet200ResponseData{}
	return &this
}

// NewMerchantMemberOperationLogListGet200ResponseDataWithDefaults instantiates a new MerchantMemberOperationLogListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMemberOperationLogListGet200ResponseDataWithDefaults() *MerchantMemberOperationLogListGet200ResponseData {
	this := MerchantMemberOperationLogListGet200ResponseData{}
	return &this
}

// GetMerchantOperationLogs returns the MerchantOperationLogs field value if set, zero value otherwise.
func (o *MerchantMemberOperationLogListGet200ResponseData) GetMerchantOperationLogs() []UnibeeApiBeanDetailMerchantOperationLogDetail {
	if o == nil || IsNil(o.MerchantOperationLogs) {
		var ret []UnibeeApiBeanDetailMerchantOperationLogDetail
		return ret
	}
	return o.MerchantOperationLogs
}

// GetMerchantOperationLogsOk returns a tuple with the MerchantOperationLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMemberOperationLogListGet200ResponseData) GetMerchantOperationLogsOk() ([]UnibeeApiBeanDetailMerchantOperationLogDetail, bool) {
	if o == nil || IsNil(o.MerchantOperationLogs) {
		return nil, false
	}
	return o.MerchantOperationLogs, true
}

// HasMerchantOperationLogs returns a boolean if a field has been set.
func (o *MerchantMemberOperationLogListGet200ResponseData) HasMerchantOperationLogs() bool {
	if o != nil && !IsNil(o.MerchantOperationLogs) {
		return true
	}

	return false
}

// SetMerchantOperationLogs gets a reference to the given []UnibeeApiBeanDetailMerchantOperationLogDetail and assigns it to the MerchantOperationLogs field.
func (o *MerchantMemberOperationLogListGet200ResponseData) SetMerchantOperationLogs(v []UnibeeApiBeanDetailMerchantOperationLogDetail) {
	o.MerchantOperationLogs = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantMemberOperationLogListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMemberOperationLogListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantMemberOperationLogListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantMemberOperationLogListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantMemberOperationLogListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMemberOperationLogListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantOperationLogs) {
		toSerialize["merchantOperationLogs"] = o.MerchantOperationLogs
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantMemberOperationLogListGet200ResponseData struct {
	value *MerchantMemberOperationLogListGet200ResponseData
	isSet bool
}

func (v NullableMerchantMemberOperationLogListGet200ResponseData) Get() *MerchantMemberOperationLogListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantMemberOperationLogListGet200ResponseData) Set(val *MerchantMemberOperationLogListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMemberOperationLogListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMemberOperationLogListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMemberOperationLogListGet200ResponseData(val *MerchantMemberOperationLogListGet200ResponseData) *NullableMerchantMemberOperationLogListGet200ResponseData {
	return &NullableMerchantMemberOperationLogListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantMemberOperationLogListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMemberOperationLogListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


