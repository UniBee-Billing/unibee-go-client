/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantTaskListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskListRes{}

// UnibeeApiMerchantTaskListRes struct for UnibeeApiMerchantTaskListRes
type UnibeeApiMerchantTaskListRes struct {
	// Merchant Member Task List
	Downloads []UnibeeApiBeanMerchantBatchTask `json:"downloads,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiMerchantTaskListRes instantiates a new UnibeeApiMerchantTaskListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskListRes() *UnibeeApiMerchantTaskListRes {
	this := UnibeeApiMerchantTaskListRes{}
	return &this
}

// NewUnibeeApiMerchantTaskListResWithDefaults instantiates a new UnibeeApiMerchantTaskListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskListResWithDefaults() *UnibeeApiMerchantTaskListRes {
	this := UnibeeApiMerchantTaskListRes{}
	return &this
}

// GetDownloads returns the Downloads field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskListRes) GetDownloads() []UnibeeApiBeanMerchantBatchTask {
	if o == nil || IsNil(o.Downloads) {
		var ret []UnibeeApiBeanMerchantBatchTask
		return ret
	}
	return o.Downloads
}

// GetDownloadsOk returns a tuple with the Downloads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskListRes) GetDownloadsOk() ([]UnibeeApiBeanMerchantBatchTask, bool) {
	if o == nil || IsNil(o.Downloads) {
		return nil, false
	}
	return o.Downloads, true
}

// HasDownloads returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskListRes) HasDownloads() bool {
	if o != nil && !IsNil(o.Downloads) {
		return true
	}

	return false
}

// SetDownloads gets a reference to the given []UnibeeApiBeanMerchantBatchTask and assigns it to the Downloads field.
func (o *UnibeeApiMerchantTaskListRes) SetDownloads(v []UnibeeApiBeanMerchantBatchTask) {
	o.Downloads = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskListRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskListRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskListRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiMerchantTaskListRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiMerchantTaskListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Downloads) {
		toSerialize["downloads"] = o.Downloads
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskListRes struct {
	value *UnibeeApiMerchantTaskListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskListRes) Get() *UnibeeApiMerchantTaskListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskListRes) Set(val *UnibeeApiMerchantTaskListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskListRes(val *UnibeeApiMerchantTaskListRes) *NullableUnibeeApiMerchantTaskListRes {
	return &NullableUnibeeApiMerchantTaskListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


