/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407080801 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantMetricSimplify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantMetricSimplify{}

// UnibeeApiBeanMerchantMetricSimplify struct for UnibeeApiBeanMerchantMetricSimplify
type UnibeeApiBeanMerchantMetricSimplify struct {
	// aggregation property
	AggregationProperty *string `json:"aggregationProperty,omitempty"`
	// 1-count，2-count unique, 3-latest, 4-max, 5-sum
	AggregationType *int32 `json:"aggregationType,omitempty"`
	// code
	Code *string `json:"code,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// update time
	GmtModify *int64 `json:"gmtModify,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// merchantId
	MerchantId *int64 `json:"merchantId,omitempty"`
	// metric description
	MetricDescription *string `json:"metricDescription,omitempty"`
	// metric name
	MetricName *string `json:"metricName,omitempty"`
	// 1-limit_metered，2-charge_metered(come later),3-charge_recurring(come later)
	Type *int32 `json:"type,omitempty"`
}

// NewUnibeeApiBeanMerchantMetricSimplify instantiates a new UnibeeApiBeanMerchantMetricSimplify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantMetricSimplify() *UnibeeApiBeanMerchantMetricSimplify {
	this := UnibeeApiBeanMerchantMetricSimplify{}
	return &this
}

// NewUnibeeApiBeanMerchantMetricSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantMetricSimplify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantMetricSimplifyWithDefaults() *UnibeeApiBeanMerchantMetricSimplify {
	this := UnibeeApiBeanMerchantMetricSimplify{}
	return &this
}

// GetAggregationProperty returns the AggregationProperty field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationProperty() string {
	if o == nil || IsNil(o.AggregationProperty) {
		var ret string
		return ret
	}
	return *o.AggregationProperty
}

// GetAggregationPropertyOk returns a tuple with the AggregationProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationProperty) {
		return nil, false
	}
	return o.AggregationProperty, true
}

// HasAggregationProperty returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasAggregationProperty() bool {
	if o != nil && !IsNil(o.AggregationProperty) {
		return true
	}

	return false
}

// SetAggregationProperty gets a reference to the given string and assigns it to the AggregationProperty field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetAggregationProperty(v string) {
	o.AggregationProperty = &v
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationType() int32 {
	if o == nil || IsNil(o.AggregationType) {
		var ret int32
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given int32 and assigns it to the AggregationType field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetAggregationType(v int32) {
	o.AggregationType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetCode(v string) {
	o.Code = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetGmtModify() int64 {
	if o == nil || IsNil(o.GmtModify) {
		var ret int64
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetGmtModifyOk() (*int64, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given int64 and assigns it to the GmtModify field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetGmtModify(v int64) {
	o.GmtModify = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetricDescription returns the MetricDescription field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricDescription() string {
	if o == nil || IsNil(o.MetricDescription) {
		var ret string
		return ret
	}
	return *o.MetricDescription
}

// GetMetricDescriptionOk returns a tuple with the MetricDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDescription) {
		return nil, false
	}
	return o.MetricDescription, true
}

// HasMetricDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasMetricDescription() bool {
	if o != nil && !IsNil(o.MetricDescription) {
		return true
	}

	return false
}

// SetMetricDescription gets a reference to the given string and assigns it to the MetricDescription field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetMetricDescription(v string) {
	o.MetricDescription = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetMetricName(v string) {
	o.MetricName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantMetricSimplify) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanMerchantMetricSimplify) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeApiBeanMerchantMetricSimplify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantMetricSimplify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationProperty) {
		toSerialize["aggregationProperty"] = o.AggregationProperty
	}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
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
	if !IsNil(o.MetricDescription) {
		toSerialize["metricDescription"] = o.MetricDescription
	}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantMetricSimplify struct {
	value *UnibeeApiBeanMerchantMetricSimplify
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantMetricSimplify) Get() *UnibeeApiBeanMerchantMetricSimplify {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantMetricSimplify) Set(val *UnibeeApiBeanMerchantMetricSimplify) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantMetricSimplify) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantMetricSimplify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantMetricSimplify(val *UnibeeApiBeanMerchantMetricSimplify) *NullableUnibeeApiBeanMerchantMetricSimplify {
	return &NullableUnibeeApiBeanMerchantMetricSimplify{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantMetricSimplify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantMetricSimplify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


