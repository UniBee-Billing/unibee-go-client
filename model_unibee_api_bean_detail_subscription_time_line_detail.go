/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408090936 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailSubscriptionTimeLineDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailSubscriptionTimeLineDetail{}

// UnibeeApiBeanDetailSubscriptionTimeLineDetail struct for UnibeeApiBeanDetailSubscriptionTimeLineDetail
type UnibeeApiBeanDetailSubscriptionTimeLineDetail struct {
	// Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// gateway_id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// payment id
	PaymentId *string `json:"paymentId,omitempty"`
	// period_end
	PeriodEnd *int64 `json:"periodEnd,omitempty"`
	// period end (datatime)
	PeriodEndTime *string `json:"periodEndTime,omitempty"`
	// period_start
	PeriodStart *int64 `json:"periodStart,omitempty"`
	// period start (datetime)
	PeriodStartTime *string `json:"periodStartTime,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	// PlanId
	PlanId *int64 `json:"planId,omitempty"`
	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// 1-processing,2-finish,3-cancelled,4-expired
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// unique id
	UniqueId *string `json:"uniqueId,omitempty"`
	// userId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanDetailSubscriptionTimeLineDetail instantiates a new UnibeeApiBeanDetailSubscriptionTimeLineDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailSubscriptionTimeLineDetail() *UnibeeApiBeanDetailSubscriptionTimeLineDetail {
	this := UnibeeApiBeanDetailSubscriptionTimeLineDetail{}
	return &this
}

// NewUnibeeApiBeanDetailSubscriptionTimeLineDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionTimeLineDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailSubscriptionTimeLineDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionTimeLineDetail {
	this := UnibeeApiBeanDetailSubscriptionTimeLineDetail{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetPeriodEnd returns the PeriodEnd field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEnd() int64 {
	if o == nil || IsNil(o.PeriodEnd) {
		var ret int64
		return ret
	}
	return *o.PeriodEnd
}

// GetPeriodEndOk returns a tuple with the PeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodEnd) {
		return nil, false
	}
	return o.PeriodEnd, true
}

// HasPeriodEnd returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodEnd() bool {
	if o != nil && !IsNil(o.PeriodEnd) {
		return true
	}

	return false
}

// SetPeriodEnd gets a reference to the given int64 and assigns it to the PeriodEnd field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodEnd(v int64) {
	o.PeriodEnd = &v
}

// GetPeriodEndTime returns the PeriodEndTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndTime() string {
	if o == nil || IsNil(o.PeriodEndTime) {
		var ret string
		return ret
	}
	return *o.PeriodEndTime
}

// GetPeriodEndTimeOk returns a tuple with the PeriodEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodEndTime) {
		return nil, false
	}
	return o.PeriodEndTime, true
}

// HasPeriodEndTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodEndTime() bool {
	if o != nil && !IsNil(o.PeriodEndTime) {
		return true
	}

	return false
}

// SetPeriodEndTime gets a reference to the given string and assigns it to the PeriodEndTime field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodEndTime(v string) {
	o.PeriodEndTime = &v
}

// GetPeriodStart returns the PeriodStart field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStart() int64 {
	if o == nil || IsNil(o.PeriodStart) {
		var ret int64
		return ret
	}
	return *o.PeriodStart
}

// GetPeriodStartOk returns a tuple with the PeriodStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartOk() (*int64, bool) {
	if o == nil || IsNil(o.PeriodStart) {
		return nil, false
	}
	return o.PeriodStart, true
}

// HasPeriodStart returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodStart() bool {
	if o != nil && !IsNil(o.PeriodStart) {
		return true
	}

	return false
}

// SetPeriodStart gets a reference to the given int64 and assigns it to the PeriodStart field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodStart(v int64) {
	o.PeriodStart = &v
}

// GetPeriodStartTime returns the PeriodStartTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartTime() string {
	if o == nil || IsNil(o.PeriodStartTime) {
		var ret string
		return ret
	}
	return *o.PeriodStartTime
}

// GetPeriodStartTimeOk returns a tuple with the PeriodStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PeriodStartTime) {
		return nil, false
	}
	return o.PeriodStartTime, true
}

// HasPeriodStartTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodStartTime() bool {
	if o != nil && !IsNil(o.PeriodStartTime) {
		return true
	}

	return false
}

// SetPeriodStartTime gets a reference to the given string and assigns it to the PeriodStartTime field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodStartTime(v string) {
	o.PeriodStartTime = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPlanId(v int64) {
	o.PlanId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanDetailSubscriptionTimeLineDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailSubscriptionTimeLineDetail) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

type NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail struct {
	value *UnibeeApiBeanDetailSubscriptionTimeLineDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) Get() *UnibeeApiBeanDetailSubscriptionTimeLineDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) Set(val *UnibeeApiBeanDetailSubscriptionTimeLineDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailSubscriptionTimeLineDetail(val *UnibeeApiBeanDetailSubscriptionTimeLineDetail) *NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail {
	return &NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailSubscriptionTimeLineDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


