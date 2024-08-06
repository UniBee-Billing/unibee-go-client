/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060911 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantTaskListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskListReq{}

// UnibeeApiMerchantTaskListReq struct for UnibeeApiMerchantTaskListReq
type UnibeeApiMerchantTaskListReq struct {
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
}

// NewUnibeeApiMerchantTaskListReq instantiates a new UnibeeApiMerchantTaskListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskListReq() *UnibeeApiMerchantTaskListReq {
	this := UnibeeApiMerchantTaskListReq{}
	return &this
}

// NewUnibeeApiMerchantTaskListReqWithDefaults instantiates a new UnibeeApiMerchantTaskListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskListReqWithDefaults() *UnibeeApiMerchantTaskListReq {
	this := UnibeeApiMerchantTaskListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantTaskListReq) SetCount(v int32) {
	o.Count = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantTaskListReq) SetPage(v int32) {
	o.Page = &v
}

func (o UnibeeApiMerchantTaskListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskListReq struct {
	value *UnibeeApiMerchantTaskListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskListReq) Get() *UnibeeApiMerchantTaskListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskListReq) Set(val *UnibeeApiMerchantTaskListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskListReq(val *UnibeeApiMerchantTaskListReq) *NullableUnibeeApiMerchantTaskListReq {
	return &NullableUnibeeApiMerchantTaskListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


