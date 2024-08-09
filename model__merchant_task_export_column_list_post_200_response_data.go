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

// checks if the MerchantTaskExportColumnListPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantTaskExportColumnListPost200ResponseData{}

// MerchantTaskExportColumnListPost200ResponseData struct for MerchantTaskExportColumnListPost200ResponseData
type MerchantTaskExportColumnListPost200ResponseData struct {
	// Export Column Comments
	ColumnComments *map[string]string `json:"columnComments,omitempty"`
	// Export Column Headers
	ColumnHeaders *map[string]string `json:"columnHeaders,omitempty"`
	// Export Column List
	Columns []map[string]interface{} `json:"columns,omitempty"`
}

// NewMerchantTaskExportColumnListPost200ResponseData instantiates a new MerchantTaskExportColumnListPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantTaskExportColumnListPost200ResponseData() *MerchantTaskExportColumnListPost200ResponseData {
	this := MerchantTaskExportColumnListPost200ResponseData{}
	return &this
}

// NewMerchantTaskExportColumnListPost200ResponseDataWithDefaults instantiates a new MerchantTaskExportColumnListPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantTaskExportColumnListPost200ResponseDataWithDefaults() *MerchantTaskExportColumnListPost200ResponseData {
	this := MerchantTaskExportColumnListPost200ResponseData{}
	return &this
}

// GetColumnComments returns the ColumnComments field value if set, zero value otherwise.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumnComments() map[string]string {
	if o == nil || IsNil(o.ColumnComments) {
		var ret map[string]string
		return ret
	}
	return *o.ColumnComments
}

// GetColumnCommentsOk returns a tuple with the ColumnComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumnCommentsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ColumnComments) {
		return nil, false
	}
	return o.ColumnComments, true
}

// HasColumnComments returns a boolean if a field has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) HasColumnComments() bool {
	if o != nil && !IsNil(o.ColumnComments) {
		return true
	}

	return false
}

// SetColumnComments gets a reference to the given map[string]string and assigns it to the ColumnComments field.
func (o *MerchantTaskExportColumnListPost200ResponseData) SetColumnComments(v map[string]string) {
	o.ColumnComments = &v
}

// GetColumnHeaders returns the ColumnHeaders field value if set, zero value otherwise.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumnHeaders() map[string]string {
	if o == nil || IsNil(o.ColumnHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.ColumnHeaders
}

// GetColumnHeadersOk returns a tuple with the ColumnHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumnHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ColumnHeaders) {
		return nil, false
	}
	return o.ColumnHeaders, true
}

// HasColumnHeaders returns a boolean if a field has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) HasColumnHeaders() bool {
	if o != nil && !IsNil(o.ColumnHeaders) {
		return true
	}

	return false
}

// SetColumnHeaders gets a reference to the given map[string]string and assigns it to the ColumnHeaders field.
func (o *MerchantTaskExportColumnListPost200ResponseData) SetColumnHeaders(v map[string]string) {
	o.ColumnHeaders = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumns() []map[string]interface{} {
	if o == nil || IsNil(o.Columns) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) GetColumnsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Columns) {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *MerchantTaskExportColumnListPost200ResponseData) HasColumns() bool {
	if o != nil && !IsNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []map[string]interface{} and assigns it to the Columns field.
func (o *MerchantTaskExportColumnListPost200ResponseData) SetColumns(v []map[string]interface{}) {
	o.Columns = v
}

func (o MerchantTaskExportColumnListPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantTaskExportColumnListPost200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantTaskExportColumnListPost200ResponseData struct {
	value *MerchantTaskExportColumnListPost200ResponseData
	isSet bool
}

func (v NullableMerchantTaskExportColumnListPost200ResponseData) Get() *MerchantTaskExportColumnListPost200ResponseData {
	return v.value
}

func (v *NullableMerchantTaskExportColumnListPost200ResponseData) Set(val *MerchantTaskExportColumnListPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantTaskExportColumnListPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantTaskExportColumnListPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantTaskExportColumnListPost200ResponseData(val *MerchantTaskExportColumnListPost200ResponseData) *NullableMerchantTaskExportColumnListPost200ResponseData {
	return &NullableMerchantTaskExportColumnListPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantTaskExportColumnListPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantTaskExportColumnListPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


