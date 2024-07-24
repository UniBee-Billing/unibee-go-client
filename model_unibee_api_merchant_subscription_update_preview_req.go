/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240454 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionUpdatePreviewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdatePreviewReq{}

// UnibeeApiMerchantSubscriptionUpdatePreviewReq struct for UnibeeApiMerchantSubscriptionUpdatePreviewReq
type UnibeeApiMerchantSubscriptionUpdatePreviewReq struct {
	// addonParams
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	// DiscountCode
	DiscountCode *string `json:"discountCode,omitempty"`
	// Effect Immediate，1-Immediate，2-Next Period
	EffectImmediate *int32 `json:"effectImmediate,omitempty"`
	// Id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// New PlanId
	NewPlanId int64 `json:"newPlanId"`
	// Quantity，Default 1
	Quantity *int64 `json:"quantity,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
}

type _UnibeeApiMerchantSubscriptionUpdatePreviewReq UnibeeApiMerchantSubscriptionUpdatePreviewReq

// NewUnibeeApiMerchantSubscriptionUpdatePreviewReq instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(newPlanId int64, subscriptionId string) *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewReq{}
	this.NewPlanId = newPlanId
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	this := UnibeeApiMerchantSubscriptionUpdatePreviewReq{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetEffectImmediate returns the EffectImmediate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetEffectImmediate() int32 {
	if o == nil || IsNil(o.EffectImmediate) {
		var ret int32
		return ret
	}
	return *o.EffectImmediate
}

// GetEffectImmediateOk returns a tuple with the EffectImmediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetEffectImmediateOk() (*int32, bool) {
	if o == nil || IsNil(o.EffectImmediate) {
		return nil, false
	}
	return o.EffectImmediate, true
}

// HasEffectImmediate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasEffectImmediate() bool {
	if o != nil && !IsNil(o.EffectImmediate) {
		return true
	}

	return false
}

// SetEffectImmediate gets a reference to the given int32 and assigns it to the EffectImmediate field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetEffectImmediate(v int32) {
	o.EffectImmediate = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetNewPlanId returns the NewPlanId field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewPlanId
}

// GetNewPlanIdOk returns a tuple with the NewPlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPlanId, true
}

// SetNewPlanId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetNewPlanId(v int64) {
	o.NewPlanId = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUpdatePreviewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.EffectImmediate) {
		toSerialize["effectImmediate"] = o.EffectImmediate
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	toSerialize["newPlanId"] = o.NewPlanId
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newPlanId",
		"subscriptionId",
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

	varUnibeeApiMerchantSubscriptionUpdatePreviewReq := _UnibeeApiMerchantSubscriptionUpdatePreviewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionUpdatePreviewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionUpdatePreviewReq(varUnibeeApiMerchantSubscriptionUpdatePreviewReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq struct {
	value *UnibeeApiMerchantSubscriptionUpdatePreviewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Get() *UnibeeApiMerchantSubscriptionUpdatePreviewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Set(val *UnibeeApiMerchantSubscriptionUpdatePreviewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUpdatePreviewReq(val *UnibeeApiMerchantSubscriptionUpdatePreviewReq) *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq {
	return &NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdatePreviewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


