/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountPlanApplyPreviewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountPlanApplyPreviewReq{}

// UnibeeApiMerchantDiscountPlanApplyPreviewReq Check discount can apply to plan, Only check rules about plan，the actual usage is subject to the subscription interface
type UnibeeApiMerchantDiscountPlanApplyPreviewReq struct {
	// The discount's unique code, customize by merchant
	Code string `json:"code"`
	// Email
	Email *string `json:"email,omitempty"`
	// The externalId of plan which code to apply, either planId or externalPlanId is needed
	ExternalPlanId *string `json:"externalPlanId,omitempty"`
	// IsChangeToLongPlan
	IsChangeToLongPlan *bool `json:"isChangeToLongPlan,omitempty"`
	// IsChangeToSameIntervalPlan
	IsChangeToSameIntervalPlan *bool `json:"isChangeToSameIntervalPlan,omitempty"`
	// IsUpgrade
	IsUpgrade *bool `json:"isUpgrade,omitempty"`
	// The id of plan which code to apply, either planId or externalPlanId is needed
	PlanId *int64 `json:"planId,omitempty"`
}

type _UnibeeApiMerchantDiscountPlanApplyPreviewReq UnibeeApiMerchantDiscountPlanApplyPreviewReq

// NewUnibeeApiMerchantDiscountPlanApplyPreviewReq instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountPlanApplyPreviewReq(code string) *UnibeeApiMerchantDiscountPlanApplyPreviewReq {
	this := UnibeeApiMerchantDiscountPlanApplyPreviewReq{}
	this.Code = code
	return &this
}

// NewUnibeeApiMerchantDiscountPlanApplyPreviewReqWithDefaults instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountPlanApplyPreviewReqWithDefaults() *UnibeeApiMerchantDiscountPlanApplyPreviewReq {
	this := UnibeeApiMerchantDiscountPlanApplyPreviewReq{}
	return &this
}

// GetCode returns the Code field value
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetCode(v string) {
	o.Code = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetEmail(v string) {
	o.Email = &v
}

// GetExternalPlanId returns the ExternalPlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetExternalPlanId() string {
	if o == nil || IsNil(o.ExternalPlanId) {
		var ret string
		return ret
	}
	return *o.ExternalPlanId
}

// GetExternalPlanIdOk returns a tuple with the ExternalPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetExternalPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPlanId) {
		return nil, false
	}
	return o.ExternalPlanId, true
}

// HasExternalPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasExternalPlanId() bool {
	if o != nil && !IsNil(o.ExternalPlanId) {
		return true
	}

	return false
}

// SetExternalPlanId gets a reference to the given string and assigns it to the ExternalPlanId field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetExternalPlanId(v string) {
	o.ExternalPlanId = &v
}

// GetIsChangeToLongPlan returns the IsChangeToLongPlan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToLongPlan() bool {
	if o == nil || IsNil(o.IsChangeToLongPlan) {
		var ret bool
		return ret
	}
	return *o.IsChangeToLongPlan
}

// GetIsChangeToLongPlanOk returns a tuple with the IsChangeToLongPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToLongPlanOk() (*bool, bool) {
	if o == nil || IsNil(o.IsChangeToLongPlan) {
		return nil, false
	}
	return o.IsChangeToLongPlan, true
}

// HasIsChangeToLongPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsChangeToLongPlan() bool {
	if o != nil && !IsNil(o.IsChangeToLongPlan) {
		return true
	}

	return false
}

// SetIsChangeToLongPlan gets a reference to the given bool and assigns it to the IsChangeToLongPlan field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsChangeToLongPlan(v bool) {
	o.IsChangeToLongPlan = &v
}

// GetIsChangeToSameIntervalPlan returns the IsChangeToSameIntervalPlan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToSameIntervalPlan() bool {
	if o == nil || IsNil(o.IsChangeToSameIntervalPlan) {
		var ret bool
		return ret
	}
	return *o.IsChangeToSameIntervalPlan
}

// GetIsChangeToSameIntervalPlanOk returns a tuple with the IsChangeToSameIntervalPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsChangeToSameIntervalPlanOk() (*bool, bool) {
	if o == nil || IsNil(o.IsChangeToSameIntervalPlan) {
		return nil, false
	}
	return o.IsChangeToSameIntervalPlan, true
}

// HasIsChangeToSameIntervalPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsChangeToSameIntervalPlan() bool {
	if o != nil && !IsNil(o.IsChangeToSameIntervalPlan) {
		return true
	}

	return false
}

// SetIsChangeToSameIntervalPlan gets a reference to the given bool and assigns it to the IsChangeToSameIntervalPlan field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsChangeToSameIntervalPlan(v bool) {
	o.IsChangeToSameIntervalPlan = &v
}

// GetIsUpgrade returns the IsUpgrade field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsUpgrade() bool {
	if o == nil || IsNil(o.IsUpgrade) {
		var ret bool
		return ret
	}
	return *o.IsUpgrade
}

// GetIsUpgradeOk returns a tuple with the IsUpgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetIsUpgradeOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUpgrade) {
		return nil, false
	}
	return o.IsUpgrade, true
}

// HasIsUpgrade returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasIsUpgrade() bool {
	if o != nil && !IsNil(o.IsUpgrade) {
		return true
	}

	return false
}

// SetIsUpgrade gets a reference to the given bool and assigns it to the IsUpgrade field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetIsUpgrade(v bool) {
	o.IsUpgrade = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) SetPlanId(v int64) {
	o.PlanId = &v
}

func (o UnibeeApiMerchantDiscountPlanApplyPreviewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountPlanApplyPreviewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalPlanId) {
		toSerialize["externalPlanId"] = o.ExternalPlanId
	}
	if !IsNil(o.IsChangeToLongPlan) {
		toSerialize["isChangeToLongPlan"] = o.IsChangeToLongPlan
	}
	if !IsNil(o.IsChangeToSameIntervalPlan) {
		toSerialize["isChangeToSameIntervalPlan"] = o.IsChangeToSameIntervalPlan
	}
	if !IsNil(o.IsUpgrade) {
		toSerialize["isUpgrade"] = o.IsUpgrade
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountPlanApplyPreviewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varUnibeeApiMerchantDiscountPlanApplyPreviewReq := _UnibeeApiMerchantDiscountPlanApplyPreviewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountPlanApplyPreviewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountPlanApplyPreviewReq(varUnibeeApiMerchantDiscountPlanApplyPreviewReq)

	return err
}

type NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq struct {
	value *UnibeeApiMerchantDiscountPlanApplyPreviewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) Get() *UnibeeApiMerchantDiscountPlanApplyPreviewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) Set(val *UnibeeApiMerchantDiscountPlanApplyPreviewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountPlanApplyPreviewReq(val *UnibeeApiMerchantDiscountPlanApplyPreviewReq) *NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq {
	return &NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountPlanApplyPreviewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


