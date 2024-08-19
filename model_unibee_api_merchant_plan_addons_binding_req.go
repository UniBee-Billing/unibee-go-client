/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanAddonsBindingReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanAddonsBindingReq{}

// UnibeeApiMerchantPlanAddonsBindingReq struct for UnibeeApiMerchantPlanAddonsBindingReq
type UnibeeApiMerchantPlanAddonsBindingReq struct {
	// Action Type，0-override,1-add，2-delete
	Action int64 `json:"action"`
	// Plan Ids Of Recurring Addon Type
	AddonIds []int64 `json:"addonIds"`
	// Plan Ids Of Onetime Addon Type
	OnetimeAddonIds []int64 `json:"onetimeAddonIds"`
	// PlanID
	PlanId int64 `json:"planId"`
}

type _UnibeeApiMerchantPlanAddonsBindingReq UnibeeApiMerchantPlanAddonsBindingReq

// NewUnibeeApiMerchantPlanAddonsBindingReq instantiates a new UnibeeApiMerchantPlanAddonsBindingReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanAddonsBindingReq(action int64, addonIds []int64, onetimeAddonIds []int64, planId int64) *UnibeeApiMerchantPlanAddonsBindingReq {
	this := UnibeeApiMerchantPlanAddonsBindingReq{}
	this.Action = action
	this.AddonIds = addonIds
	this.OnetimeAddonIds = onetimeAddonIds
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanAddonsBindingReqWithDefaults instantiates a new UnibeeApiMerchantPlanAddonsBindingReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanAddonsBindingReqWithDefaults() *UnibeeApiMerchantPlanAddonsBindingReq {
	this := UnibeeApiMerchantPlanAddonsBindingReq{}
	return &this
}

// GetAction returns the Action field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAction() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetActionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetAction(v int64) {
	o.Action = v
}

// GetAddonIds returns the AddonIds field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAddonIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetAddonIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddonIds, true
}

// SetAddonIds sets field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetOnetimeAddonIds returns the OnetimeAddonIds field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetOnetimeAddonIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.OnetimeAddonIds
}

// GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetOnetimeAddonIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnetimeAddonIds, true
}

// SetOnetimeAddonIds sets field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetOnetimeAddonIds(v []int64) {
	o.OnetimeAddonIds = v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanAddonsBindingReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanAddonsBindingReq) SetPlanId(v int64) {
	o.PlanId = v
}

func (o UnibeeApiMerchantPlanAddonsBindingReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanAddonsBindingReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["addonIds"] = o.AddonIds
	toSerialize["onetimeAddonIds"] = o.OnetimeAddonIds
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanAddonsBindingReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"addonIds",
		"onetimeAddonIds",
		"planId",
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

	varUnibeeApiMerchantPlanAddonsBindingReq := _UnibeeApiMerchantPlanAddonsBindingReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanAddonsBindingReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanAddonsBindingReq(varUnibeeApiMerchantPlanAddonsBindingReq)

	return err
}

type NullableUnibeeApiMerchantPlanAddonsBindingReq struct {
	value *UnibeeApiMerchantPlanAddonsBindingReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingReq) Get() *UnibeeApiMerchantPlanAddonsBindingReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingReq) Set(val *UnibeeApiMerchantPlanAddonsBindingReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanAddonsBindingReq(val *UnibeeApiMerchantPlanAddonsBindingReq) *NullableUnibeeApiMerchantPlanAddonsBindingReq {
	return &NullableUnibeeApiMerchantPlanAddonsBindingReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanAddonsBindingReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanAddonsBindingReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


