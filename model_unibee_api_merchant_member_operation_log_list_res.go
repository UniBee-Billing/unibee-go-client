/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMemberOperationLogListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMemberOperationLogListRes{}

// UnibeeApiMerchantMemberOperationLogListRes struct for UnibeeApiMerchantMemberOperationLogListRes
type UnibeeApiMerchantMemberOperationLogListRes struct {
	// Merchant Member Operation Log List
	MerchantOperationLogs []UnibeeApiBeanDetailMerchantOperationLogDetail `json:"merchantOperationLogs,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantMemberOperationLogListRes instantiates a new UnibeeApiMerchantMemberOperationLogListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMemberOperationLogListRes() *UnibeeApiMerchantMemberOperationLogListRes {
	this := UnibeeApiMerchantMemberOperationLogListRes{}
	return &this
}

// NewUnibeeApiMerchantMemberOperationLogListResWithDefaults instantiates a new UnibeeApiMerchantMemberOperationLogListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMemberOperationLogListResWithDefaults() *UnibeeApiMerchantMemberOperationLogListRes {
	this := UnibeeApiMerchantMemberOperationLogListRes{}
	return &this
}

// GetMerchantOperationLogs returns the MerchantOperationLogs field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberOperationLogListRes) GetMerchantOperationLogs() []UnibeeApiBeanDetailMerchantOperationLogDetail {
	if o == nil || IsNil(o.MerchantOperationLogs) {
		var ret []UnibeeApiBeanDetailMerchantOperationLogDetail
		return ret
	}
	return o.MerchantOperationLogs
}

// GetMerchantOperationLogsOk returns a tuple with the MerchantOperationLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberOperationLogListRes) GetMerchantOperationLogsOk() ([]UnibeeApiBeanDetailMerchantOperationLogDetail, bool) {
	if o == nil || IsNil(o.MerchantOperationLogs) {
		return nil, false
	}
	return o.MerchantOperationLogs, true
}

// HasMerchantOperationLogs returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberOperationLogListRes) HasMerchantOperationLogs() bool {
	if o != nil && !IsNil(o.MerchantOperationLogs) {
		return true
	}

	return false
}

// SetMerchantOperationLogs gets a reference to the given []UnibeeApiBeanDetailMerchantOperationLogDetail and assigns it to the MerchantOperationLogs field.
func (o *UnibeeApiMerchantMemberOperationLogListRes) SetMerchantOperationLogs(v []UnibeeApiBeanDetailMerchantOperationLogDetail) {
	o.MerchantOperationLogs = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMemberOperationLogListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMemberOperationLogListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMemberOperationLogListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantMemberOperationLogListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantMemberOperationLogListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMemberOperationLogListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MerchantOperationLogs) {
		toSerialize["merchantOperationLogs"] = o.MerchantOperationLogs
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMemberOperationLogListRes struct {
	value *UnibeeApiMerchantMemberOperationLogListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantMemberOperationLogListRes) Get() *UnibeeApiMerchantMemberOperationLogListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantMemberOperationLogListRes) Set(val *UnibeeApiMerchantMemberOperationLogListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMemberOperationLogListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMemberOperationLogListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMemberOperationLogListRes(val *UnibeeApiMerchantMemberOperationLogListRes) *NullableUnibeeApiMerchantMemberOperationLogListRes {
	return &NullableUnibeeApiMerchantMemberOperationLogListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMemberOperationLogListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMemberOperationLogListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


