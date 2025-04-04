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

// checks if the UnibeeApiMerchantMetricEventListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricEventListReq{}

// UnibeeApiMerchantMetricEventListReq struct for UnibeeApiMerchantMetricEventListReq
type UnibeeApiMerchantMetricEventListReq struct {
	// Count OF Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// Filter MetricIds, Default All
	MetricIds []int64 `json:"metricIds,omitempty"`
	// Page,Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort，user_id|gmt_create，Default gmt_create
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// Filter UserIds, Default All
	UserIds []int64 `json:"userIds,omitempty"`
}

// NewUnibeeApiMerchantMetricEventListReq instantiates a new UnibeeApiMerchantMetricEventListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricEventListReq() *UnibeeApiMerchantMetricEventListReq {
	this := UnibeeApiMerchantMetricEventListReq{}
	return &this
}

// NewUnibeeApiMerchantMetricEventListReqWithDefaults instantiates a new UnibeeApiMerchantMetricEventListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricEventListReqWithDefaults() *UnibeeApiMerchantMetricEventListReq {
	this := UnibeeApiMerchantMetricEventListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantMetricEventListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantMetricEventListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantMetricEventListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetMetricIds returns the MetricIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetMetricIds() []int64 {
	if o == nil || IsNil(o.MetricIds) {
		var ret []int64
		return ret
	}
	return o.MetricIds
}

// GetMetricIdsOk returns a tuple with the MetricIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetMetricIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.MetricIds) {
		return nil, false
	}
	return o.MetricIds, true
}

// HasMetricIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasMetricIds() bool {
	if o != nil && !IsNil(o.MetricIds) {
		return true
	}

	return false
}

// SetMetricIds gets a reference to the given []int64 and assigns it to the MetricIds field.
func (o *UnibeeApiMerchantMetricEventListReq) SetMetricIds(v []int64) {
	o.MetricIds = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantMetricEventListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantMetricEventListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantMetricEventListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetUserIds returns the UserIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricEventListReq) GetUserIds() []int64 {
	if o == nil || IsNil(o.UserIds) {
		var ret []int64
		return ret
	}
	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricEventListReq) GetUserIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.UserIds) {
		return nil, false
	}
	return o.UserIds, true
}

// HasUserIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricEventListReq) HasUserIds() bool {
	if o != nil && !IsNil(o.UserIds) {
		return true
	}

	return false
}

// SetUserIds gets a reference to the given []int64 and assigns it to the UserIds field.
func (o *UnibeeApiMerchantMetricEventListReq) SetUserIds(v []int64) {
	o.UserIds = v
}

func (o UnibeeApiMerchantMetricEventListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricEventListReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.MetricIds) {
		toSerialize["metricIds"] = o.MetricIds
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.SortField) {
		toSerialize["sortField"] = o.SortField
	}
	if !IsNil(o.SortType) {
		toSerialize["sortType"] = o.SortType
	}
	if !IsNil(o.UserIds) {
		toSerialize["userIds"] = o.UserIds
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricEventListReq struct {
	value *UnibeeApiMerchantMetricEventListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricEventListReq) Get() *UnibeeApiMerchantMetricEventListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricEventListReq) Set(val *UnibeeApiMerchantMetricEventListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricEventListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricEventListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricEventListReq(val *UnibeeApiMerchantMetricEventListReq) *NullableUnibeeApiMerchantMetricEventListReq {
	return &NullableUnibeeApiMerchantMetricEventListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricEventListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricEventListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


