/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountUserDiscountListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountUserDiscountListReq{}

// UnibeeApiMerchantDiscountUserDiscountListReq Get user discountCode list
type UnibeeApiMerchantDiscountUserDiscountListReq struct {
	// Count Of Per Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// The discount's Id
	Id int64 `json:"id"`
	// Page, Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort Field，gmt_create|gmt_modify，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
}

type _UnibeeApiMerchantDiscountUserDiscountListReq UnibeeApiMerchantDiscountUserDiscountListReq

// NewUnibeeApiMerchantDiscountUserDiscountListReq instantiates a new UnibeeApiMerchantDiscountUserDiscountListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountUserDiscountListReq(id int64) *UnibeeApiMerchantDiscountUserDiscountListReq {
	this := UnibeeApiMerchantDiscountUserDiscountListReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountUserDiscountListReqWithDefaults instantiates a new UnibeeApiMerchantDiscountUserDiscountListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountUserDiscountListReqWithDefaults() *UnibeeApiMerchantDiscountUserDiscountListReq {
	this := UnibeeApiMerchantDiscountUserDiscountListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetId(v int64) {
	o.Id = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantDiscountUserDiscountListReq) SetSortType(v string) {
	o.SortType = &v
}

func (o UnibeeApiMerchantDiscountUserDiscountListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountUserDiscountListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.CreateTimeEnd) {
		toSerialize["createTimeEnd"] = o.CreateTimeEnd
	}
	if !IsNil(o.CreateTimeStart) {
		toSerialize["createTimeStart"] = o.CreateTimeStart
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.SortField) {
		toSerialize["sortField"] = o.SortField
	}
	if !IsNil(o.SortType) {
		toSerialize["sortType"] = o.SortType
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountUserDiscountListReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantDiscountUserDiscountListReq := _UnibeeApiMerchantDiscountUserDiscountListReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountUserDiscountListReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountUserDiscountListReq(varUnibeeApiMerchantDiscountUserDiscountListReq)

	return err
}

type NullableUnibeeApiMerchantDiscountUserDiscountListReq struct {
	value *UnibeeApiMerchantDiscountUserDiscountListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListReq) Get() *UnibeeApiMerchantDiscountUserDiscountListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListReq) Set(val *UnibeeApiMerchantDiscountUserDiscountListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountUserDiscountListReq(val *UnibeeApiMerchantDiscountUserDiscountListReq) *NullableUnibeeApiMerchantDiscountUserDiscountListReq {
	return &NullableUnibeeApiMerchantDiscountUserDiscountListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountUserDiscountListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountUserDiscountListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


