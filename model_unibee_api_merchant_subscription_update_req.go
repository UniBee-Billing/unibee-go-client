/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionUpdateReq{}

// UnibeeApiMerchantSubscriptionUpdateReq struct for UnibeeApiMerchantSubscriptionUpdateReq
type UnibeeApiMerchantSubscriptionUpdateReq struct {
	// addonParams
	AddonParams []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo `json:"addonParams,omitempty"`
	// Currency To Be Confirmed，Get From Preview
	ConfirmCurrency string `json:"confirmCurrency"`
	// TotalAmount To Be Confirmed，Get From Preview
	ConfirmTotalAmount int64 `json:"confirmTotalAmount"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// New PlanId
	NewPlanId int64 `json:"newPlanId"`
	// prorationDate date to start Proration，Get From Preview
	ProrationDate int64 `json:"prorationDate"`
	// Quantity，Default 1
	Quantity *int64 `json:"quantity,omitempty"`
	// SubscriptionId
	SubscriptionId string `json:"subscriptionId"`
	// Effect Immediate，1-Immediate，2-Next Period
	WithImmediateEffect *int32 `json:"withImmediateEffect,omitempty"`
}

type _UnibeeApiMerchantSubscriptionUpdateReq UnibeeApiMerchantSubscriptionUpdateReq

// NewUnibeeApiMerchantSubscriptionUpdateReq instantiates a new UnibeeApiMerchantSubscriptionUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionUpdateReq(confirmCurrency string, confirmTotalAmount int64, newPlanId int64, prorationDate int64, subscriptionId string) *UnibeeApiMerchantSubscriptionUpdateReq {
	this := UnibeeApiMerchantSubscriptionUpdateReq{}
	this.ConfirmCurrency = confirmCurrency
	this.ConfirmTotalAmount = confirmTotalAmount
	this.NewPlanId = newPlanId
	this.ProrationDate = prorationDate
	this.SubscriptionId = subscriptionId
	return &this
}

// NewUnibeeApiMerchantSubscriptionUpdateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionUpdateReqWithDefaults() *UnibeeApiMerchantSubscriptionUpdateReq {
	this := UnibeeApiMerchantSubscriptionUpdateReq{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParams() []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParamsOk() ([]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo and assigns it to the AddonParams field.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetAddonParams(v []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo) {
	o.AddonParams = v
}

// GetConfirmCurrency returns the ConfirmCurrency field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmCurrency
}

// GetConfirmCurrencyOk returns a tuple with the ConfirmCurrency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmCurrency, true
}

// SetConfirmCurrency sets field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetConfirmCurrency(v string) {
	o.ConfirmCurrency = v
}

// GetConfirmTotalAmount returns the ConfirmTotalAmount field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmTotalAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ConfirmTotalAmount
}

// GetConfirmTotalAmountOk returns a tuple with the ConfirmTotalAmount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmTotalAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmTotalAmount, true
}

// SetConfirmTotalAmount sets field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetConfirmTotalAmount(v int64) {
	o.ConfirmTotalAmount = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetNewPlanId returns the NewPlanId field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewPlanId
}

// GetNewPlanIdOk returns a tuple with the NewPlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPlanId, true
}

// SetNewPlanId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetNewPlanId(v int64) {
	o.NewPlanId = v
}

// GetProrationDate returns the ProrationDate field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetProrationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProrationDate
}

// GetProrationDateOk returns a tuple with the ProrationDate field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetProrationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProrationDate, true
}

// SetProrationDate sets field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetProrationDate(v int64) {
	o.ProrationDate = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetWithImmediateEffect returns the WithImmediateEffect field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetWithImmediateEffect() int32 {
	if o == nil || IsNil(o.WithImmediateEffect) {
		var ret int32
		return ret
	}
	return *o.WithImmediateEffect
}

// GetWithImmediateEffectOk returns a tuple with the WithImmediateEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetWithImmediateEffectOk() (*int32, bool) {
	if o == nil || IsNil(o.WithImmediateEffect) {
		return nil, false
	}
	return o.WithImmediateEffect, true
}

// HasWithImmediateEffect returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasWithImmediateEffect() bool {
	if o != nil && !IsNil(o.WithImmediateEffect) {
		return true
	}

	return false
}

// SetWithImmediateEffect gets a reference to the given int32 and assigns it to the WithImmediateEffect field.
func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetWithImmediateEffect(v int32) {
	o.WithImmediateEffect = &v
}

func (o UnibeeApiMerchantSubscriptionUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionUpdateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	toSerialize["confirmCurrency"] = o.ConfirmCurrency
	toSerialize["confirmTotalAmount"] = o.ConfirmTotalAmount
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["newPlanId"] = o.NewPlanId
	toSerialize["prorationDate"] = o.ProrationDate
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	toSerialize["subscriptionId"] = o.SubscriptionId
	if !IsNil(o.WithImmediateEffect) {
		toSerialize["withImmediateEffect"] = o.WithImmediateEffect
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionUpdateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"confirmCurrency",
		"confirmTotalAmount",
		"newPlanId",
		"prorationDate",
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

	varUnibeeApiMerchantSubscriptionUpdateReq := _UnibeeApiMerchantSubscriptionUpdateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionUpdateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionUpdateReq(varUnibeeApiMerchantSubscriptionUpdateReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionUpdateReq struct {
	value *UnibeeApiMerchantSubscriptionUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateReq) Get() *UnibeeApiMerchantSubscriptionUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateReq) Set(val *UnibeeApiMerchantSubscriptionUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionUpdateReq(val *UnibeeApiMerchantSubscriptionUpdateReq) *NullableUnibeeApiMerchantSubscriptionUpdateReq {
	return &NullableUnibeeApiMerchantSubscriptionUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


