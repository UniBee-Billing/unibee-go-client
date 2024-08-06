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

// checks if the UnibeeApiMerchantTaskExportColumnListRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantTaskExportColumnListRes{}

// UnibeeApiMerchantTaskExportColumnListRes struct for UnibeeApiMerchantTaskExportColumnListRes
type UnibeeApiMerchantTaskExportColumnListRes struct {
	// Export Column Comments
	ColumnComments *map[string]string `json:"columnComments,omitempty"`
	// Export Column Headers
	ColumnHeaders *map[string]string `json:"columnHeaders,omitempty"`
	// Export Column List
	Columns []map[string]interface{} `json:"columns,omitempty"`
}

// NewUnibeeApiMerchantTaskExportColumnListRes instantiates a new UnibeeApiMerchantTaskExportColumnListRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantTaskExportColumnListRes() *UnibeeApiMerchantTaskExportColumnListRes {
	this := UnibeeApiMerchantTaskExportColumnListRes{}
	return &this
}

// NewUnibeeApiMerchantTaskExportColumnListResWithDefaults instantiates a new UnibeeApiMerchantTaskExportColumnListRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantTaskExportColumnListResWithDefaults() *UnibeeApiMerchantTaskExportColumnListRes {
	this := UnibeeApiMerchantTaskExportColumnListRes{}
	return &this
}

// GetColumnComments returns the ColumnComments field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumnComments() map[string]string {
	if o == nil || IsNil(o.ColumnComments) {
		var ret map[string]string
		return ret
	}
	return *o.ColumnComments
}

// GetColumnCommentsOk returns a tuple with the ColumnComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumnCommentsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ColumnComments) {
		return nil, false
	}
	return o.ColumnComments, true
}

// HasColumnComments returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) HasColumnComments() bool {
	if o != nil && !IsNil(o.ColumnComments) {
		return true
	}

	return false
}

// SetColumnComments gets a reference to the given map[string]string and assigns it to the ColumnComments field.
func (o *UnibeeApiMerchantTaskExportColumnListRes) SetColumnComments(v map[string]string) {
	o.ColumnComments = &v
}

// GetColumnHeaders returns the ColumnHeaders field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumnHeaders() map[string]string {
	if o == nil || IsNil(o.ColumnHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.ColumnHeaders
}

// GetColumnHeadersOk returns a tuple with the ColumnHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumnHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ColumnHeaders) {
		return nil, false
	}
	return o.ColumnHeaders, true
}

// HasColumnHeaders returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) HasColumnHeaders() bool {
	if o != nil && !IsNil(o.ColumnHeaders) {
		return true
	}

	return false
}

// SetColumnHeaders gets a reference to the given map[string]string and assigns it to the ColumnHeaders field.
func (o *UnibeeApiMerchantTaskExportColumnListRes) SetColumnHeaders(v map[string]string) {
	o.ColumnHeaders = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumns() []map[string]interface{} {
	if o == nil || IsNil(o.Columns) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) GetColumnsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *UnibeeApiMerchantTaskExportColumnListRes) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []map[string]interface{} and assigns it to the Columns field.
func (o *UnibeeApiMerchantTaskExportColumnListRes) SetColumns(v []map[string]interface{}) {
	o.Columns = v
}

func (o UnibeeApiMerchantTaskExportColumnListRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantTaskExportColumnListRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ColumnComments) {
		toSerialize["columnComments"] = o.ColumnComments
	}
	if !IsNil(o.ColumnHeaders) {
		toSerialize["columnHeaders"] = o.ColumnHeaders
	}
	if !IsNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantTaskExportColumnListRes struct {
	value *UnibeeApiMerchantTaskExportColumnListRes
	isSet bool
}

func (v NullableUnibeeApiMerchantTaskExportColumnListRes) Get() *UnibeeApiMerchantTaskExportColumnListRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListRes) Set(val *UnibeeApiMerchantTaskExportColumnListRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantTaskExportColumnListRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantTaskExportColumnListRes(val *UnibeeApiMerchantTaskExportColumnListRes) *NullableUnibeeApiMerchantTaskExportColumnListRes {
	return &NullableUnibeeApiMerchantTaskExportColumnListRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantTaskExportColumnListRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantTaskExportColumnListRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


