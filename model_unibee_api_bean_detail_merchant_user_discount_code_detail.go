/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailMerchantUserDiscountCodeDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailMerchantUserDiscountCodeDetail{}

// UnibeeApiBeanDetailMerchantUserDiscountCodeDetail struct for UnibeeApiBeanDetailMerchantUserDiscountCodeDetail
type UnibeeApiBeanDetailMerchantUserDiscountCodeDetail struct {
	// apply_amount
	ApplyAmount *int64 `json:"applyAmount,omitempty"`
	// code
	Code *string `json:"code,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// ID
	Id *int64 `json:"id,omitempty"`
	// invoice_id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchantId
	MerchantId *int64 `json:"merchantId,omitempty"`
	// payment_id
	PaymentId *string `json:"paymentId,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	// subscription_id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailMerchantUserDiscountCodeDetail instantiates a new UnibeeApiBeanDetailMerchantUserDiscountCodeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailMerchantUserDiscountCodeDetail() *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail {
	this := UnibeeApiBeanDetailMerchantUserDiscountCodeDetail{}
	return &this
}

// NewUnibeeApiBeanDetailMerchantUserDiscountCodeDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantUserDiscountCodeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailMerchantUserDiscountCodeDetailWithDefaults() *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail {
	this := UnibeeApiBeanDetailMerchantUserDiscountCodeDetail{}
	return &this
}

// GetApplyAmount returns the ApplyAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetApplyAmount() int64 {
	if o == nil || IsNil(o.ApplyAmount) {
		var ret int64
		return ret
	}
	return *o.ApplyAmount
}

// GetApplyAmountOk returns a tuple with the ApplyAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetApplyAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.ApplyAmount) {
		return nil, false
	}
	return o.ApplyAmount, true
}

// HasApplyAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasApplyAmount() bool {
	if o != nil && !IsNil(o.ApplyAmount) {
		return true
	}

	return false
}

// SetApplyAmount gets a reference to the given int64 and assigns it to the ApplyAmount field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetApplyAmount(v int64) {
	o.ApplyAmount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetCode(v string) {
	o.Code = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetId(v int64) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyAmount) {
		toSerialize["applyAmount"] = o.ApplyAmount
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail struct {
	value *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) Get() *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) Set(val *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail(val *UnibeeApiBeanDetailMerchantUserDiscountCodeDetail) *NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail {
	return &NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailMerchantUserDiscountCodeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


