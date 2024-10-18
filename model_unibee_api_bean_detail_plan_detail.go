/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailPlanDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailPlanDetail{}

// UnibeeApiBeanDetailPlanDetail struct for UnibeeApiBeanDetailPlanDetail
type UnibeeApiBeanDetailPlanDetail struct {
	// AddonIds
	AddonIds []int64 `json:"addonIds,omitempty"`
	// Addons
	Addons []UnibeeApiBeanPlan `json:"addons,omitempty"`
	// MetricPlanLimits
	MetricPlanLimits []UnibeeApiBeanMerchantMetricPlanLimit `json:"metricPlanLimits,omitempty"`
	// OneTimeAddonIds
	OnetimeAddonIds []int64 `json:"onetimeAddonIds,omitempty"`
	// OneTimeAddons
	OnetimeAddons []UnibeeApiBeanPlan `json:"onetimeAddons,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Product *UnibeeApiBeanProduct `json:"product,omitempty"`
}

// NewUnibeeApiBeanDetailPlanDetail instantiates a new UnibeeApiBeanDetailPlanDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailPlanDetail() *UnibeeApiBeanDetailPlanDetail {
	this := UnibeeApiBeanDetailPlanDetail{}
	return &this
}

// NewUnibeeApiBeanDetailPlanDetailWithDefaults instantiates a new UnibeeApiBeanDetailPlanDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailPlanDetailWithDefaults() *UnibeeApiBeanDetailPlanDetail {
	this := UnibeeApiBeanDetailPlanDetail{}
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeApiBeanDetailPlanDetail) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetAddons() []UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlan
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetAddonsOk() ([]UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlan and assigns it to the Addons field.
func (o *UnibeeApiBeanDetailPlanDetail) SetAddons(v []UnibeeApiBeanPlan) {
	o.Addons = v
}

// GetMetricPlanLimits returns the MetricPlanLimits field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetMetricPlanLimits() []UnibeeApiBeanMerchantMetricPlanLimit {
	if o == nil || IsNil(o.MetricPlanLimits) {
		var ret []UnibeeApiBeanMerchantMetricPlanLimit
		return ret
	}
	return o.MetricPlanLimits
}

// GetMetricPlanLimitsOk returns a tuple with the MetricPlanLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetMetricPlanLimitsOk() ([]UnibeeApiBeanMerchantMetricPlanLimit, bool) {
	if o == nil || IsNil(o.MetricPlanLimits) {
		return nil, false
	}
	return o.MetricPlanLimits, true
}

// HasMetricPlanLimits returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasMetricPlanLimits() bool {
	if o != nil && !IsNil(o.MetricPlanLimits) {
		return true
	}

	return false
}

// SetMetricPlanLimits gets a reference to the given []UnibeeApiBeanMerchantMetricPlanLimit and assigns it to the MetricPlanLimits field.
func (o *UnibeeApiBeanDetailPlanDetail) SetMetricPlanLimits(v []UnibeeApiBeanMerchantMetricPlanLimit) {
	o.MetricPlanLimits = v
}

// GetOnetimeAddonIds returns the OnetimeAddonIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonIds() []int64 {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		var ret []int64
		return ret
	}
	return o.OnetimeAddonIds
}

// GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		return nil, false
	}
	return o.OnetimeAddonIds, true
}

// HasOnetimeAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasOnetimeAddonIds() bool {
	if o != nil && !IsNil(o.OnetimeAddonIds) {
		return true
	}

	return false
}

// SetOnetimeAddonIds gets a reference to the given []int64 and assigns it to the OnetimeAddonIds field.
func (o *UnibeeApiBeanDetailPlanDetail) SetOnetimeAddonIds(v []int64) {
	o.OnetimeAddonIds = v
}

// GetOnetimeAddons returns the OnetimeAddons field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddons() []UnibeeApiBeanPlan {
	if o == nil || IsNil(o.OnetimeAddons) {
		var ret []UnibeeApiBeanPlan
		return ret
	}
	return o.OnetimeAddons
}

// GetOnetimeAddonsOk returns a tuple with the OnetimeAddons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonsOk() ([]UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.OnetimeAddons) {
		return nil, false
	}
	return o.OnetimeAddons, true
}

// HasOnetimeAddons returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasOnetimeAddons() bool {
	if o != nil && !IsNil(o.OnetimeAddons) {
		return true
	}

	return false
}

// SetOnetimeAddons gets a reference to the given []UnibeeApiBeanPlan and assigns it to the OnetimeAddons field.
func (o *UnibeeApiBeanDetailPlanDetail) SetOnetimeAddons(v []UnibeeApiBeanPlan) {
	o.OnetimeAddons = v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiBeanDetailPlanDetail) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPlanDetail) GetProduct() UnibeeApiBeanProduct {
	if o == nil || IsNil(o.Product) {
		var ret UnibeeApiBeanProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPlanDetail) GetProductOk() (*UnibeeApiBeanProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPlanDetail) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given UnibeeApiBeanProduct and assigns it to the Product field.
func (o *UnibeeApiBeanDetailPlanDetail) SetProduct(v UnibeeApiBeanProduct) {
	o.Product = &v
}

func (o UnibeeApiBeanDetailPlanDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailPlanDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonIds) {
		toSerialize["addonIds"] = o.AddonIds
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.MetricPlanLimits) {
		toSerialize["metricPlanLimits"] = o.MetricPlanLimits
	}
	if !IsNil(o.OnetimeAddonIds) {
		toSerialize["onetimeAddonIds"] = o.OnetimeAddonIds
	}
	if !IsNil(o.OnetimeAddons) {
		toSerialize["onetimeAddons"] = o.OnetimeAddons
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailPlanDetail struct {
	value *UnibeeApiBeanDetailPlanDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailPlanDetail) Get() *UnibeeApiBeanDetailPlanDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailPlanDetail) Set(val *UnibeeApiBeanDetailPlanDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailPlanDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailPlanDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailPlanDetail(val *UnibeeApiBeanDetailPlanDetail) *NullableUnibeeApiBeanDetailPlanDetail {
	return &NullableUnibeeApiBeanDetailPlanDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailPlanDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailPlanDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


