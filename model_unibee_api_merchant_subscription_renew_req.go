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

// checks if the UnibeeApiMerchantSubscriptionRenewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionRenewReq{}

// UnibeeApiMerchantSubscriptionRenewReq renew an exist subscription 
type UnibeeApiMerchantSubscriptionRenewReq struct {
	// apply promo credit or not
	ApplyPromoCredit *bool `json:"applyPromoCredit,omitempty"`
	// apply promo credit amount, auto compute if not specified
	ApplyPromoCreditAmount *int32 `json:"applyPromoCreditAmount,omitempty"`
	// CancelUrl
	CancelUrl *string `json:"cancelUrl,omitempty"`
	Discount *UnibeeApiBeanExternalDiscountParam `json:"discount,omitempty"`
	// DiscountCode, override subscription discount
	DiscountCode *string `json:"discountCode,omitempty"`
	// GatewayId, use subscription's gateway if not provide
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// Gateway Payment Type
	GatewayPaymentType *string `json:"gatewayPaymentType,omitempty"`
	// ManualPayment
	ManualPayment *bool `json:"manualPayment,omitempty"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	ProductData *UnibeeApiBeanPlanProductParam `json:"productData,omitempty"`
	// default product will use if not specified
	ProductId *int64 `json:"productId,omitempty"`
	// ReturnUrl
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// SubscriptionId, id of subscription which addon will attached, either SubscriptionId or UserId needed, The only one active subscription or latest subscription will renew if userId provide instead of subscriptionId
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// TaxPercentage，1000 = 10%, override subscription taxPercentage if provide
	TaxPercentage *int32 `json:"taxPercentage,omitempty"`
	// UserId, either SubscriptionId or UserId needed, The only one active subscription or latest cancel|expire subscription will renew if userId provide instead of subscriptionId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionRenewReq instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionRenewReq() *UnibeeApiMerchantSubscriptionRenewReq {
	this := UnibeeApiMerchantSubscriptionRenewReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults() *UnibeeApiMerchantSubscriptionRenewReq {
	this := UnibeeApiMerchantSubscriptionRenewReq{}
	return &this
}

// GetApplyPromoCredit returns the ApplyPromoCredit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCredit() bool {
	if o == nil || IsNil(o.ApplyPromoCredit) {
		var ret bool
		return ret
	}
	return *o.ApplyPromoCredit
}

// GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyPromoCredit) {
		return nil, false
	}
	return o.ApplyPromoCredit, true
}

// HasApplyPromoCredit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasApplyPromoCredit() bool {
	if o != nil && !IsNil(o.ApplyPromoCredit) {
		return true
	}

	return false
}

// SetApplyPromoCredit gets a reference to the given bool and assigns it to the ApplyPromoCredit field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetApplyPromoCredit(v bool) {
	o.ApplyPromoCredit = &v
}

// GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditAmount() int32 {
	if o == nil || IsNil(o.ApplyPromoCreditAmount) {
		var ret int32
		return ret
	}
	return *o.ApplyPromoCreditAmount
}

// GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApplyPromoCreditAmount) {
		return nil, false
	}
	return o.ApplyPromoCreditAmount, true
}

// HasApplyPromoCreditAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasApplyPromoCreditAmount() bool {
	if o != nil && !IsNil(o.ApplyPromoCreditAmount) {
		return true
	}

	return false
}

// SetApplyPromoCreditAmount gets a reference to the given int32 and assigns it to the ApplyPromoCreditAmount field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetApplyPromoCreditAmount(v int32) {
	o.ApplyPromoCreditAmount = &v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscount() UnibeeApiBeanExternalDiscountParam {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanExternalDiscountParam
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanExternalDiscountParam and assigns it to the Discount field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam) {
	o.Discount = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetGatewayPaymentType returns the GatewayPaymentType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayPaymentType() string {
	if o == nil || IsNil(o.GatewayPaymentType) {
		var ret string
		return ret
	}
	return *o.GatewayPaymentType
}

// GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayPaymentType) {
		return nil, false
	}
	return o.GatewayPaymentType, true
}

// HasGatewayPaymentType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasGatewayPaymentType() bool {
	if o != nil && !IsNil(o.GatewayPaymentType) {
		return true
	}

	return false
}

// SetGatewayPaymentType gets a reference to the given string and assigns it to the GatewayPaymentType field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetGatewayPaymentType(v string) {
	o.GatewayPaymentType = &v
}

// GetManualPayment returns the ManualPayment field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetManualPayment() bool {
	if o == nil || IsNil(o.ManualPayment) {
		var ret bool
		return ret
	}
	return *o.ManualPayment
}

// GetManualPaymentOk returns a tuple with the ManualPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetManualPaymentOk() (*bool, bool) {
	if o == nil || IsNil(o.ManualPayment) {
		return nil, false
	}
	return o.ManualPayment, true
}

// HasManualPayment returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasManualPayment() bool {
	if o != nil && !IsNil(o.ManualPayment) {
		return true
	}

	return false
}

// SetManualPayment gets a reference to the given bool and assigns it to the ManualPayment field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetManualPayment(v bool) {
	o.ManualPayment = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetProductData returns the ProductData field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductData() UnibeeApiBeanPlanProductParam {
	if o == nil || IsNil(o.ProductData) {
		var ret UnibeeApiBeanPlanProductParam
		return ret
	}
	return *o.ProductData
}

// GetProductDataOk returns a tuple with the ProductData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductDataOk() (*UnibeeApiBeanPlanProductParam, bool) {
	if o == nil || IsNil(o.ProductData) {
		return nil, false
	}
	return o.ProductData, true
}

// HasProductData returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasProductData() bool {
	if o != nil && !IsNil(o.ProductData) {
		return true
	}

	return false
}

// SetProductData gets a reference to the given UnibeeApiBeanPlanProductParam and assigns it to the ProductData field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetProductData(v UnibeeApiBeanPlanProductParam) {
	o.ProductData = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetProductId(v int64) {
	o.ProductId = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentage() int32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int32 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetTaxPercentage(v int32) {
	o.TaxPercentage = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionRenewReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionRenewReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionRenewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionRenewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyPromoCredit) {
		toSerialize["applyPromoCredit"] = o.ApplyPromoCredit
	}
	if !IsNil(o.ApplyPromoCreditAmount) {
		toSerialize["applyPromoCreditAmount"] = o.ApplyPromoCreditAmount
	}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayPaymentType) {
		toSerialize["gatewayPaymentType"] = o.GatewayPaymentType
	}
	if !IsNil(o.ManualPayment) {
		toSerialize["manualPayment"] = o.ManualPayment
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ProductData) {
		toSerialize["productData"] = o.ProductData
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionRenewReq struct {
	value *UnibeeApiMerchantSubscriptionRenewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) Get() *UnibeeApiMerchantSubscriptionRenewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) Set(val *UnibeeApiMerchantSubscriptionRenewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionRenewReq(val *UnibeeApiMerchantSubscriptionRenewReq) *NullableUnibeeApiMerchantSubscriptionRenewReq {
	return &NullableUnibeeApiMerchantSubscriptionRenewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionRenewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionRenewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


