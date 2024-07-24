/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantMetricPlanLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantMetricPlanLimit{}

// UnibeeApiBeanMerchantMetricPlanLimit struct for UnibeeApiBeanMerchantMetricPlanLimit
type UnibeeApiBeanMerchantMetricPlanLimit struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// update time
	GmtModify *int64 `json:"gmtModify,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// merchantId
	MerchantId *int64 `json:"merchantId,omitempty"`
	MerchantMetricVo *UnibeeApiBeanMerchantMetricSimplify `json:"merchantMetricVo,omitempty"`
	// metricId
	MetricId *int64 `json:"metricId,omitempty"`
	// plan metric limit
	MetricLimit *int64 `json:"metricLimit,omitempty"`
	// plan_id
	PlanId *int64 `json:"planId,omitempty"`
}

// NewUnibeeApiBeanMerchantMetricPlanLimit instantiates a new UnibeeApiBeanMerchantMetricPlanLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantMetricPlanLimit() *UnibeeApiBeanMerchantMetricPlanLimit {
	this := UnibeeApiBeanMerchantMetricPlanLimit{}
	return &this
}

// NewUnibeeApiBeanMerchantMetricPlanLimitWithDefaults instantiates a new UnibeeApiBeanMerchantMetricPlanLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantMetricPlanLimitWithDefaults() *UnibeeApiBeanMerchantMetricPlanLimit {
	this := UnibeeApiBeanMerchantMetricPlanLimit{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetGmtModify() int64 {
	if o == nil || IsNil(o.GmtModify) {
		var ret int64
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetGmtModifyOk() (*int64, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given int64 and assigns it to the GmtModify field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetGmtModify(v int64) {
	o.GmtModify = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMerchantMetricVo returns the MerchantMetricVo field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantMetricVo() UnibeeApiBeanMerchantMetricSimplify {
	if o == nil || IsNil(o.MerchantMetricVo) {
		var ret UnibeeApiBeanMerchantMetricSimplify
		return ret
	}
	return *o.MerchantMetricVo
}

// GetMerchantMetricVoOk returns a tuple with the MerchantMetricVo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMerchantMetricVoOk() (*UnibeeApiBeanMerchantMetricSimplify, bool) {
	if o == nil || IsNil(o.MerchantMetricVo) {
		return nil, false
	}
	return o.MerchantMetricVo, true
}

// HasMerchantMetricVo returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMerchantMetricVo() bool {
	if o != nil && !IsNil(o.MerchantMetricVo) {
		return true
	}

	return false
}

// SetMerchantMetricVo gets a reference to the given UnibeeApiBeanMerchantMetricSimplify and assigns it to the MerchantMetricVo field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMerchantMetricVo(v UnibeeApiBeanMerchantMetricSimplify) {
	o.MerchantMetricVo = &v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricId() int64 {
	if o == nil || IsNil(o.MetricId) {
		var ret int64
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricId) {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMetricId() bool {
	if o != nil && !IsNil(o.MetricId) {
		return true
	}

	return false
}

// SetMetricId gets a reference to the given int64 and assigns it to the MetricId field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMetricId(v int64) {
	o.MetricId = &v
}

// GetMetricLimit returns the MetricLimit field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricLimit() int64 {
	if o == nil || IsNil(o.MetricLimit) {
		var ret int64
		return ret
	}
	return *o.MetricLimit
}

// GetMetricLimitOk returns a tuple with the MetricLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetMetricLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.MetricLimit) {
		return nil, false
	}
	return o.MetricLimit, true
}

// HasMetricLimit returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasMetricLimit() bool {
	if o != nil && !IsNil(o.MetricLimit) {
		return true
	}

	return false
}

// SetMetricLimit gets a reference to the given int64 and assigns it to the MetricLimit field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetMetricLimit(v int64) {
	o.MetricLimit = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiBeanMerchantMetricPlanLimit) SetPlanId(v int64) {
	o.PlanId = &v
}

func (o UnibeeApiBeanMerchantMetricPlanLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantMetricPlanLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.GmtModify) {
		toSerialize["gmtModify"] = o.GmtModify
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.MerchantMetricVo) {
		toSerialize["merchantMetricVo"] = o.MerchantMetricVo
	}
	if !IsNil(o.MetricId) {
		toSerialize["metricId"] = o.MetricId
	}
	if !IsNil(o.MetricLimit) {
		toSerialize["metricLimit"] = o.MetricLimit
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantMetricPlanLimit struct {
	value *UnibeeApiBeanMerchantMetricPlanLimit
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantMetricPlanLimit) Get() *UnibeeApiBeanMerchantMetricPlanLimit {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantMetricPlanLimit) Set(val *UnibeeApiBeanMerchantMetricPlanLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantMetricPlanLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantMetricPlanLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantMetricPlanLimit(val *UnibeeApiBeanMerchantMetricPlanLimit) *NullableUnibeeApiBeanMerchantMetricPlanLimit {
	return &NullableUnibeeApiBeanMerchantMetricPlanLimit{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantMetricPlanLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantMetricPlanLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


