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

// checks if the UnibeeApiBeanRefund type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanRefund{}

// UnibeeApiBeanRefund struct for UnibeeApiBeanRefund
type UnibeeApiBeanRefund struct {
	// country code
	CountryCode *string `json:"countryCode,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// external_refund_id
	ExternalRefundId *string `json:"externalRefundId,omitempty"`
	// gateway_id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// gateway refund id
	GatewayRefundId *string `json:"gatewayRefundId,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// relative payment id
	PaymentId *string `json:"paymentId,omitempty"`
	// refund amount, cent
	RefundAmount *int64 `json:"refundAmount,omitempty"`
	// refund comment
	RefundComment *string `json:"refundComment,omitempty"`
	// refund id (system generate)
	RefundId *string `json:"refundId,omitempty"`
	// refund success time
	RefundTime *int64 `json:"refundTime,omitempty"`
	// return url after refund success
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// status。10-pending，20-success，30-failure, 40-cancel
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// 1-gateway refund,2-mark refund
	Type *int32 `json:"type,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanRefund instantiates a new UnibeeApiBeanRefund object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanRefund() *UnibeeApiBeanRefund {
	this := UnibeeApiBeanRefund{}
	return &this
}

// NewUnibeeApiBeanRefundWithDefaults instantiates a new UnibeeApiBeanRefund object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanRefundWithDefaults() *UnibeeApiBeanRefund {
	this := UnibeeApiBeanRefund{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanRefund) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanRefund) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanRefund) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalRefundId returns the ExternalRefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetExternalRefundId() string {
	if o == nil || IsNil(o.ExternalRefundId) {
		var ret string
		return ret
	}
	return *o.ExternalRefundId
}

// GetExternalRefundIdOk returns a tuple with the ExternalRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetExternalRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalRefundId) {
		return nil, false
	}
	return o.ExternalRefundId, true
}

// HasExternalRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasExternalRefundId() bool {
	if o != nil && !IsNil(o.ExternalRefundId) {
		return true
	}

	return false
}

// SetExternalRefundId gets a reference to the given string and assigns it to the ExternalRefundId field.
func (o *UnibeeApiBeanRefund) SetExternalRefundId(v string) {
	o.ExternalRefundId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanRefund) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayRefundId returns the GatewayRefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetGatewayRefundId() string {
	if o == nil || IsNil(o.GatewayRefundId) {
		var ret string
		return ret
	}
	return *o.GatewayRefundId
}

// GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetGatewayRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayRefundId) {
		return nil, false
	}
	return o.GatewayRefundId, true
}

// HasGatewayRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasGatewayRefundId() bool {
	if o != nil && !IsNil(o.GatewayRefundId) {
		return true
	}

	return false
}

// SetGatewayRefundId gets a reference to the given string and assigns it to the GatewayRefundId field.
func (o *UnibeeApiBeanRefund) SetGatewayRefundId(v string) {
	o.GatewayRefundId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanRefund) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanRefund) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanRefund) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanRefund) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetRefundAmount() int64 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret int64
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetRefundAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given int64 and assigns it to the RefundAmount field.
func (o *UnibeeApiBeanRefund) SetRefundAmount(v int64) {
	o.RefundAmount = &v
}

// GetRefundComment returns the RefundComment field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetRefundComment() string {
	if o == nil || IsNil(o.RefundComment) {
		var ret string
		return ret
	}
	return *o.RefundComment
}

// GetRefundCommentOk returns a tuple with the RefundComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetRefundCommentOk() (*string, bool) {
	if o == nil || IsNil(o.RefundComment) {
		return nil, false
	}
	return o.RefundComment, true
}

// HasRefundComment returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasRefundComment() bool {
	if o != nil && !IsNil(o.RefundComment) {
		return true
	}

	return false
}

// SetRefundComment gets a reference to the given string and assigns it to the RefundComment field.
func (o *UnibeeApiBeanRefund) SetRefundComment(v string) {
	o.RefundComment = &v
}

// GetRefundId returns the RefundId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetRefundId() string {
	if o == nil || IsNil(o.RefundId) {
		var ret string
		return ret
	}
	return *o.RefundId
}

// GetRefundIdOk returns a tuple with the RefundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetRefundIdOk() (*string, bool) {
	if o == nil || IsNil(o.RefundId) {
		return nil, false
	}
	return o.RefundId, true
}

// HasRefundId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasRefundId() bool {
	if o != nil && !IsNil(o.RefundId) {
		return true
	}

	return false
}

// SetRefundId gets a reference to the given string and assigns it to the RefundId field.
func (o *UnibeeApiBeanRefund) SetRefundId(v string) {
	o.RefundId = &v
}

// GetRefundTime returns the RefundTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetRefundTime() int64 {
	if o == nil || IsNil(o.RefundTime) {
		var ret int64
		return ret
	}
	return *o.RefundTime
}

// GetRefundTimeOk returns a tuple with the RefundTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetRefundTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundTime) {
		return nil, false
	}
	return o.RefundTime, true
}

// HasRefundTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasRefundTime() bool {
	if o != nil && !IsNil(o.RefundTime) {
		return true
	}

	return false
}

// SetRefundTime gets a reference to the given int64 and assigns it to the RefundTime field.
func (o *UnibeeApiBeanRefund) SetRefundTime(v int64) {
	o.RefundTime = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiBeanRefund) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanRefund) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanRefund) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanRefund) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanRefund) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanRefund) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanRefund) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanRefund) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanRefund) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanRefund) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExternalRefundId) {
		toSerialize["externalRefundId"] = o.ExternalRefundId
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayRefundId) {
		toSerialize["gatewayRefundId"] = o.GatewayRefundId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refundAmount"] = o.RefundAmount
	}
	if !IsNil(o.RefundComment) {
		toSerialize["refundComment"] = o.RefundComment
	}
	if !IsNil(o.RefundId) {
		toSerialize["refundId"] = o.RefundId
	}
	if !IsNil(o.RefundTime) {
		toSerialize["refundTime"] = o.RefundTime
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanRefund struct {
	value *UnibeeApiBeanRefund
	isSet bool
}

func (v NullableUnibeeApiBeanRefund) Get() *UnibeeApiBeanRefund {
	return v.value
}

func (v *NullableUnibeeApiBeanRefund) Set(val *UnibeeApiBeanRefund) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanRefund) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanRefund) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanRefund(val *UnibeeApiBeanRefund) *NullableUnibeeApiBeanRefund {
	return &NullableUnibeeApiBeanRefund{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanRefund) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanRefund) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


