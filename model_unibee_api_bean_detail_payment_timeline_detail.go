/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408160545 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailPaymentTimelineDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailPaymentTimelineDetail{}

// UnibeeApiBeanDetailPaymentTimelineDetail struct for UnibeeApiBeanDetailPaymentTimelineDetail
type UnibeeApiBeanDetailPaymentTimelineDetail struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// ExternalTransactionId
	ExternalTransactionId *string `json:"externalTransactionId,omitempty"`
	// 0-no, 1-yes
	FullRefund *int32 `json:"fullRefund,omitempty"`
	// gateway id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	Id *int64 `json:"id,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Payment *UnibeeApiBeanPayment `json:"payment,omitempty"`
	// PaymentId
	PaymentId *string `json:"paymentId,omitempty"`
	Refund *UnibeeApiBeanRefund `json:"refund,omitempty"`
	// refund id
	RefundId *string `json:"refundId,omitempty"`
	// 0-pending, 1-success, 2-failure
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// 0-pay, 1-refund
	TimelineType *int32 `json:"timelineType,omitempty"`
	// total amount
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// TransactionId
	TransactionId *string `json:"transactionId,omitempty"`
	// userId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanDetailPaymentTimelineDetail instantiates a new UnibeeApiBeanDetailPaymentTimelineDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailPaymentTimelineDetail() *UnibeeApiBeanDetailPaymentTimelineDetail {
	this := UnibeeApiBeanDetailPaymentTimelineDetail{}
	return &this
}

// NewUnibeeApiBeanDetailPaymentTimelineDetailWithDefaults instantiates a new UnibeeApiBeanDetailPaymentTimelineDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailPaymentTimelineDetailWithDefaults() *UnibeeApiBeanDetailPaymentTimelineDetail {
	this := UnibeeApiBeanDetailPaymentTimelineDetail{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalTransactionId returns the ExternalTransactionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetExternalTransactionId() string {
	if o == nil || IsNil(o.ExternalTransactionId) {
		var ret string
		return ret
	}
	return *o.ExternalTransactionId
}

// GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetExternalTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalTransactionId) {
		return nil, false
	}
	return o.ExternalTransactionId, true
}

// HasExternalTransactionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasExternalTransactionId() bool {
	if o != nil && !IsNil(o.ExternalTransactionId) {
		return true
	}

	return false
}

// SetExternalTransactionId gets a reference to the given string and assigns it to the ExternalTransactionId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetExternalTransactionId(v string) {
	o.ExternalTransactionId = &v
}

// GetFullRefund returns the FullRefund field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetFullRefund() int32 {
	if o == nil || IsNil(o.FullRefund) {
		var ret int32
		return ret
	}
	return *o.FullRefund
}

// GetFullRefundOk returns a tuple with the FullRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetFullRefundOk() (*int32, bool) {
	if o == nil || IsNil(o.FullRefund) {
		return nil, false
	}
	return o.FullRefund, true
}

// HasFullRefund returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasFullRefund() bool {
	if o != nil && !IsNil(o.FullRefund) {
		return true
	}

	return false
}

// SetFullRefund gets a reference to the given int32 and assigns it to the FullRefund field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetFullRefund(v int32) {
	o.FullRefund = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetId(v int64) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPayment() UnibeeApiBeanPayment {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPayment and assigns it to the Payment field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetPayment(v UnibeeApiBeanPayment) {
	o.Payment = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefund() UnibeeApiBeanRefund {
	if o == nil || IsNil(o.Refund) {
		var ret UnibeeApiBeanRefund
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundOk() (*UnibeeApiBeanRefund, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given UnibeeApiBeanRefund and assigns it to the Refund field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetRefund(v UnibeeApiBeanRefund) {
	o.Refund = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetRefundId(v string) {
	o.RefundId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTimelineType returns the TimelineType field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTimelineType() int32 {
	if o == nil || IsNil(o.TimelineType) {
		var ret int32
		return ret
	}
	return *o.TimelineType
}

// GetTimelineTypeOk returns a tuple with the TimelineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTimelineTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.TimelineType) {
		return nil, false
	}
	return o.TimelineType, true
}

// HasTimelineType returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTimelineType() bool {
	if o != nil && !IsNil(o.TimelineType) {
		return true
	}

	return false
}

// SetTimelineType gets a reference to the given int32 and assigns it to the TimelineType field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTimelineType(v int32) {
	o.TimelineType = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanDetailPaymentTimelineDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailPaymentTimelineDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExternalTransactionId) {
		toSerialize["externalTransactionId"] = o.ExternalTransactionId
	}
	if !IsNil(o.FullRefund) {
		toSerialize["fullRefund"] = o.FullRefund
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
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
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.TimelineType) {
		toSerialize["timelineType"] = o.TimelineType
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailPaymentTimelineDetail struct {
	value *UnibeeApiBeanDetailPaymentTimelineDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailPaymentTimelineDetail) Get() *UnibeeApiBeanDetailPaymentTimelineDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailPaymentTimelineDetail) Set(val *UnibeeApiBeanDetailPaymentTimelineDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailPaymentTimelineDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailPaymentTimelineDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailPaymentTimelineDetail(val *UnibeeApiBeanDetailPaymentTimelineDetail) *NullableUnibeeApiBeanDetailPaymentTimelineDetail {
	return &NullableUnibeeApiBeanDetailPaymentTimelineDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailPaymentTimelineDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailPaymentTimelineDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


