/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPaymentItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPaymentItem{}

// UnibeeApiBeanPaymentItem struct for UnibeeApiBeanPaymentItem
type UnibeeApiBeanPaymentItem struct {
	// amount
	Amount *int64 `json:"amount,omitempty"`
	// biz_type 1-onetime payment, 3-subscription
	BizType *int32 `json:"bizType,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	Id *int64 `json:"id,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// PaymentId
	PaymentId *string `json:"paymentId,omitempty"`
	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// 0-pending, 1-success, 2-failure
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// unique id
	UniqueId *string `json:"uniqueId,omitempty"`
	// unit_amount
	UnitAmount *int64 `json:"unitAmount,omitempty"`
	// userId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanPaymentItem instantiates a new UnibeeApiBeanPaymentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPaymentItem() *UnibeeApiBeanPaymentItem {
	this := UnibeeApiBeanPaymentItem{}
	return &this
}

// NewUnibeeApiBeanPaymentItemWithDefaults instantiates a new UnibeeApiBeanPaymentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPaymentItemWithDefaults() *UnibeeApiBeanPaymentItem {
	this := UnibeeApiBeanPaymentItem{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeApiBeanPaymentItem) SetAmount(v int64) {
	o.Amount = &v
}

// GetBizType returns the BizType field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetBizType() int32 {
	if o == nil || IsNil(o.BizType) {
		var ret int32
		return ret
	}
	return *o.BizType
}

// GetBizTypeOk returns a tuple with the BizType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetBizTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.BizType) {
		return nil, false
	}
	return o.BizType, true
}

// HasBizType returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasBizType() bool {
	if o != nil && !IsNil(o.BizType) {
		return true
	}

	return false
}

// SetBizType gets a reference to the given int32 and assigns it to the BizType field.
func (o *UnibeeApiBeanPaymentItem) SetBizType(v int32) {
	o.BizType = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanPaymentItem) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanPaymentItem) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanPaymentItem) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanPaymentItem) SetId(v int64) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanPaymentItem) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanPaymentItem) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanPaymentItem) SetName(v string) {
	o.Name = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanPaymentItem) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanPaymentItem) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanPaymentItem) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanPaymentItem) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetUniqueId() string {
	if o == nil || IsNil(o.UniqueId) {
		var ret string
		return ret
	}
	return *o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetUniqueIdOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueId) {
		return nil, false
	}
	return o.UniqueId, true
}

// HasUniqueId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasUniqueId() bool {
	if o != nil && !IsNil(o.UniqueId) {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given string and assigns it to the UniqueId field.
func (o *UnibeeApiBeanPaymentItem) SetUniqueId(v string) {
	o.UniqueId = &v
}

// GetUnitAmount returns the UnitAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetUnitAmount() int64 {
	if o == nil || IsNil(o.UnitAmount) {
		var ret int64
		return ret
	}
	return *o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetUnitAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.UnitAmount) {
		return nil, false
	}
	return o.UnitAmount, true
}

// HasUnitAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasUnitAmount() bool {
	if o != nil && !IsNil(o.UnitAmount) {
		return true
	}

	return false
}

// SetUnitAmount gets a reference to the given int64 and assigns it to the UnitAmount field.
func (o *UnibeeApiBeanPaymentItem) SetUnitAmount(v int64) {
	o.UnitAmount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPaymentItem) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPaymentItem) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPaymentItem) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanPaymentItem) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanPaymentItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPaymentItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BizType) {
		toSerialize["bizType"] = o.BizType
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
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
	if !IsNil(o.UnitAmount) {
		toSerialize["unitAmount"] = o.UnitAmount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPaymentItem struct {
	value *UnibeeApiBeanPaymentItem
	isSet bool
}

func (v NullableUnibeeApiBeanPaymentItem) Get() *UnibeeApiBeanPaymentItem {
	return v.value
}

func (v *NullableUnibeeApiBeanPaymentItem) Set(val *UnibeeApiBeanPaymentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPaymentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPaymentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPaymentItem(val *UnibeeApiBeanPaymentItem) *NullableUnibeeApiBeanPaymentItem {
	return &NullableUnibeeApiBeanPaymentItem{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPaymentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPaymentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


