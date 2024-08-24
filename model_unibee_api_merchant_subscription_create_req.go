/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantSubscriptionCreateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCreateReq{}

// UnibeeApiMerchantSubscriptionCreateReq struct for UnibeeApiMerchantSubscriptionCreateReq
type UnibeeApiMerchantSubscriptionCreateReq struct {
	// addonParams
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	// CancelUrl
	CancelUrl *string `json:"cancelUrl,omitempty"`
	// Currency to verify if provide
	ConfirmCurrency *string `json:"confirmCurrency,omitempty"`
	// TotalAmount to verify if provide
	ConfirmTotalAmount *int64 `json:"confirmTotalAmount,omitempty"`
	Discount *UnibeeApiBeanExternalDiscountParam `json:"discount,omitempty"`
	// DiscountCode
	DiscountCode *string `json:"discountCode,omitempty"`
	// Email, one of ExternalUserId&Email, UserId or User needed
	Email *string `json:"email,omitempty"`
	// ExternalUserId, unique, one of ExternalUserId&Email, UserId or User needed
	ExternalUserId *string `json:"externalUserId,omitempty"`
	// GatewayId
	GatewayId *int32 `json:"gatewayId,omitempty"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// PaymentMethodId
	PaymentMethodId *string `json:"paymentMethodId,omitempty"`
	// PlanId
	PlanId int64 `json:"planId"`
	ProductData *UnibeeApiBeanPlanProductParam `json:"productData,omitempty"`
	// Quantity，Default 1
	Quantity *int64 `json:"quantity,omitempty"`
	// ReturnUrl
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// StartIncomplete, use now pay later, subscription will generate invoice and start with incomplete status if set
	StartIncomplete *bool `json:"startIncomplete,omitempty"`
	// TaxPercentage，1000 = 10%, override subscription taxPercentage if provide
	TaxPercentage *int32 `json:"taxPercentage,omitempty"`
	// trial_end, utc time
	TrialEnd *int64 `json:"trialEnd,omitempty"`
	User *UnibeeApiBeanNewUser `json:"user,omitempty"`
	// UserId
	UserId *int64 `json:"userId,omitempty"`
	// VatCountryCode, CountryName
	VatCountryCode *string `json:"vatCountryCode,omitempty"`
	// VatNumber
	VatNumber *string `json:"vatNumber,omitempty"`
}

type _UnibeeApiMerchantSubscriptionCreateReq UnibeeApiMerchantSubscriptionCreateReq

// NewUnibeeApiMerchantSubscriptionCreateReq instantiates a new UnibeeApiMerchantSubscriptionCreateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCreateReq(planId int64) *UnibeeApiMerchantSubscriptionCreateReq {
	this := UnibeeApiMerchantSubscriptionCreateReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantSubscriptionCreateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCreateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCreateReqWithDefaults() *UnibeeApiMerchantSubscriptionCreateReq {
	this := UnibeeApiMerchantSubscriptionCreateReq{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetCancelUrl returns the CancelUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetCancelUrl() string {
	if o == nil || IsNil(o.CancelUrl) {
		var ret string
		return ret
	}
	return *o.CancelUrl
}

// GetCancelUrlOk returns a tuple with the CancelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetCancelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CancelUrl) {
		return nil, false
	}
	return o.CancelUrl, true
}

// HasCancelUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasCancelUrl() bool {
	if o != nil && !IsNil(o.CancelUrl) {
		return true
	}

	return false
}

// SetCancelUrl gets a reference to the given string and assigns it to the CancelUrl field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetCancelUrl(v string) {
	o.CancelUrl = &v
}

// GetConfirmCurrency returns the ConfirmCurrency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmCurrency() string {
	if o == nil || IsNil(o.ConfirmCurrency) {
		var ret string
		return ret
	}
	return *o.ConfirmCurrency
}

// GetConfirmCurrencyOk returns a tuple with the ConfirmCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.ConfirmCurrency) {
		return nil, false
	}
	return o.ConfirmCurrency, true
}

// HasConfirmCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasConfirmCurrency() bool {
	if o != nil && !IsNil(o.ConfirmCurrency) {
		return true
	}

	return false
}

// SetConfirmCurrency gets a reference to the given string and assigns it to the ConfirmCurrency field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetConfirmCurrency(v string) {
	o.ConfirmCurrency = &v
}

// GetConfirmTotalAmount returns the ConfirmTotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmTotalAmount() int64 {
	if o == nil || IsNil(o.ConfirmTotalAmount) {
		var ret int64
		return ret
	}
	return *o.ConfirmTotalAmount
}

// GetConfirmTotalAmountOk returns a tuple with the ConfirmTotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.ConfirmTotalAmount) {
		return nil, false
	}
	return o.ConfirmTotalAmount, true
}

// HasConfirmTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasConfirmTotalAmount() bool {
	if o != nil && !IsNil(o.ConfirmTotalAmount) {
		return true
	}

	return false
}

// SetConfirmTotalAmount gets a reference to the given int64 and assigns it to the ConfirmTotalAmount field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetConfirmTotalAmount(v int64) {
	o.ConfirmTotalAmount = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscount() UnibeeApiBeanExternalDiscountParam {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanExternalDiscountParam
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanExternalDiscountParam and assigns it to the Discount field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam) {
	o.Discount = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetEmail(v string) {
	o.Email = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetGatewayId() int32 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int32
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetGatewayIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int32 and assigns it to the GatewayId field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetGatewayId(v int32) {
	o.GatewayId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPaymentMethodId returns the PaymentMethodId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPaymentMethodId() string {
	if o == nil || IsNil(o.PaymentMethodId) {
		var ret string
		return ret
	}
	return *o.PaymentMethodId
}

// GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPaymentMethodIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodId) {
		return nil, false
	}
	return o.PaymentMethodId, true
}

// HasPaymentMethodId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasPaymentMethodId() bool {
	if o != nil && !IsNil(o.PaymentMethodId) {
		return true
	}

	return false
}

// SetPaymentMethodId gets a reference to the given string and assigns it to the PaymentMethodId field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetPaymentMethodId(v string) {
	o.PaymentMethodId = &v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetPlanId(v int64) {
	o.PlanId = v
}

// GetProductData returns the ProductData field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetProductData() UnibeeApiBeanPlanProductParam {
	if o == nil || IsNil(o.ProductData) {
		var ret UnibeeApiBeanPlanProductParam
		return ret
	}
	return *o.ProductData
}

// GetProductDataOk returns a tuple with the ProductData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetProductDataOk() (*UnibeeApiBeanPlanProductParam, bool) {
	if o == nil || IsNil(o.ProductData) {
		return nil, false
	}
	return o.ProductData, true
}

// HasProductData returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasProductData() bool {
	if o != nil && !IsNil(o.ProductData) {
		return true
	}

	return false
}

// SetProductData gets a reference to the given UnibeeApiBeanPlanProductParam and assigns it to the ProductData field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetProductData(v UnibeeApiBeanPlanProductParam) {
	o.ProductData = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetStartIncomplete returns the StartIncomplete field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetStartIncomplete() bool {
	if o == nil || IsNil(o.StartIncomplete) {
		var ret bool
		return ret
	}
	return *o.StartIncomplete
}

// GetStartIncompleteOk returns a tuple with the StartIncomplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetStartIncompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.StartIncomplete) {
		return nil, false
	}
	return o.StartIncomplete, true
}

// HasStartIncomplete returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasStartIncomplete() bool {
	if o != nil && !IsNil(o.StartIncomplete) {
		return true
	}

	return false
}

// SetStartIncomplete gets a reference to the given bool and assigns it to the StartIncomplete field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetStartIncomplete(v bool) {
	o.StartIncomplete = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTaxPercentage() int32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTaxPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int32 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetTaxPercentage(v int32) {
	o.TaxPercentage = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTrialEnd() int64 {
	if o == nil || IsNil(o.TrialEnd) {
		var ret int64
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTrialEndOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int64 and assigns it to the TrialEnd field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetTrialEnd(v int64) {
	o.TrialEnd = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUser() UnibeeApiBeanNewUser {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanNewUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUserOk() (*UnibeeApiBeanNewUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanNewUser and assigns it to the User field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetUser(v UnibeeApiBeanNewUser) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetUserId(v int64) {
	o.UserId = &v
}

// GetVatCountryCode returns the VatCountryCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatCountryCode() string {
	if o == nil || IsNil(o.VatCountryCode) {
		var ret string
		return ret
	}
	return *o.VatCountryCode
}

// GetVatCountryCodeOk returns a tuple with the VatCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryCode) {
		return nil, false
	}
	return o.VatCountryCode, true
}

// HasVatCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasVatCountryCode() bool {
	if o != nil && !IsNil(o.VatCountryCode) {
		return true
	}

	return false
}

// SetVatCountryCode gets a reference to the given string and assigns it to the VatCountryCode field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetVatCountryCode(v string) {
	o.VatCountryCode = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreateReq) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UnibeeApiMerchantSubscriptionCreateReq) SetVatNumber(v string) {
	o.VatNumber = &v
}

func (o UnibeeApiMerchantSubscriptionCreateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCreateReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.CancelUrl) {
		toSerialize["cancelUrl"] = o.CancelUrl
	}
	if !IsNil(o.ConfirmCurrency) {
		toSerialize["confirmCurrency"] = o.ConfirmCurrency
	}
	if !IsNil(o.ConfirmTotalAmount) {
		toSerialize["confirmTotalAmount"] = o.ConfirmTotalAmount
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaymentMethodId) {
		toSerialize["paymentMethodId"] = o.PaymentMethodId
	}
	toSerialize["planId"] = o.PlanId
	if !IsNil(o.ProductData) {
		toSerialize["productData"] = o.ProductData
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.StartIncomplete) {
		toSerialize["startIncomplete"] = o.StartIncomplete
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.TrialEnd) {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.VatCountryCode) {
		toSerialize["vatCountryCode"] = o.VatCountryCode
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantSubscriptionCreateReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
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

	varUnibeeApiMerchantSubscriptionCreateReq := _UnibeeApiMerchantSubscriptionCreateReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantSubscriptionCreateReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantSubscriptionCreateReq(varUnibeeApiMerchantSubscriptionCreateReq)

	return err
}

type NullableUnibeeApiMerchantSubscriptionCreateReq struct {
	value *UnibeeApiMerchantSubscriptionCreateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCreateReq) Get() *UnibeeApiMerchantSubscriptionCreateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCreateReq) Set(val *UnibeeApiMerchantSubscriptionCreateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCreateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCreateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCreateReq(val *UnibeeApiMerchantSubscriptionCreateReq) *NullableUnibeeApiMerchantSubscriptionCreateReq {
	return &NullableUnibeeApiMerchantSubscriptionCreateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCreateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCreateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


