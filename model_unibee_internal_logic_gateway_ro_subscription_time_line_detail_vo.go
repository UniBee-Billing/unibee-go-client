/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo{}

// UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo struct for UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo
type UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo struct {
	// Addon
	Addons []UnibeeInternalLogicGatewayRoPlanAddonVo `json:"addons,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// gateway_id
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int32 `json:"merchantId,omitempty"`
	// period_end
	PeriodEnd *int64 `json:"periodEnd,omitempty"`
	// period end (datatime)
	PeriodEndTime *string `json:"periodEndTime,omitempty"`
	// period_start
	PeriodStart *int64 `json:"periodStart,omitempty"`
	// period start (datetime)
	PeriodStartTime *string `json:"periodStartTime,omitempty"`
	Plan *UnibeeInternalLogicGatewayRoPlanSimplify `json:"plan,omitempty"`
	// PlanId
	PlanId *int32 `json:"planId,omitempty"`
	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// unique id
	UniqueId *string `json:"uniqueId,omitempty"`
	// userId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo instantiates a new UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo() *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo {
	this := UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo{}
	return &this
}

// NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVoWithDefaults() *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo {
	this := UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeInternalLogicGatewayRoPlanAddonVo
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetAddonsOk() ([]UnibeeInternalLogicGatewayRoPlanAddonVo, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeInternalLogicGatewayRoPlanAddonVo and assigns it to the Addons field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo) {
	o.Addons = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetCurrency(v string) {
	o.Currency = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetMerchantId() int32 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int32
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetMerchantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int32 and assigns it to the MerchantId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetMerchantId(v int32) {
	o.MerchantId = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEnd() int64 {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret int64
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given int64 and assigns it to the PeriodEnd field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodEnd(v int64) {
	o.PeriodEnd = &v
}

// GetPeriodEndTime returns the PeriodEndTime field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndTime() string {
	if o == nil || IsNil(o.PeriodEndTime) {
		var ret string
		return ret
	}
	return *o.PeriodEndTime
}

// GetPeriodEndTimeOk returns a tuple with the PeriodEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEndTime) {
		return nil, false
	}
	return o.PeriodEndTime, true
}

// HasPeriodEndTime returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodEndTime() bool {
	if o != nil && !IsNil(o.PeriodEndTime) {
		return true
	}

	return false
}

// SetPeriodEndTime gets a reference to the given string and assigns it to the PeriodEndTime field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodEndTime(v string) {
	o.PeriodEndTime = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStart() int64 {
	if o == nil || IsNil(o.PeriodStart) {
		var ret int64
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given int64 and assigns it to the PeriodStart field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodStart(v int64) {
	o.PeriodStart = &v
}

// GetPeriodStartTime returns the PeriodStartTime field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartTime() string {
	if o == nil || IsNil(o.PeriodStartTime) {
		var ret string
		return ret
	}
	return *o.PeriodStartTime
}

// GetPeriodStartTimeOk returns a tuple with the PeriodStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodStartTime) {
		return nil, false
	}
	return o.PeriodStartTime, true
}

// HasPeriodStartTime returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodStartTime() bool {
	if o != nil && !IsNil(o.PeriodStartTime) {
		return true
	}

	return false
}

// SetPeriodStartTime gets a reference to the given string and assigns it to the PeriodStartTime field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodStartTime(v string) {
	o.PeriodStartTime = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeInternalLogicGatewayRoPlanSimplify
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeInternalLogicGatewayRoPlanSimplify and assigns it to the Plan field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify) {
	o.Plan = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanId() int32 {
	if o == nil || IsNil(o.PlanId) {
		var ret int32
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int32 and assigns it to the PlanId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPlanId(v int32) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.PeriodEnd) {
		toSerialize["periodEnd"] = o.PeriodEnd
	}
	if !IsNil(o.PeriodEndTime) {
		toSerialize["periodEndTime"] = o.PeriodEndTime
	}
	if !IsNil(o.PeriodStart) {
		toSerialize["periodStart"] = o.PeriodStart
	}
	if !IsNil(o.PeriodStartTime) {
		toSerialize["periodStartTime"] = o.PeriodStartTime
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UniqueId) {
		toSerialize["uniqueId"] = o.UniqueId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo struct {
	value *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo
	isSet bool
}

func (v NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) Get() *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo {
	return v.value
}

func (v *NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) Set(val *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo(val *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) *NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo {
	return &NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo{value: val, isSet: true}
}

func (v NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


