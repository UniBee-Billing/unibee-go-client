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

// checks if the UnibeeApiMerchantSubscriptionCreatePreviewRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionCreatePreviewRes{}

// UnibeeApiMerchantSubscriptionCreatePreviewRes struct for UnibeeApiMerchantSubscriptionCreatePreviewRes
type UnibeeApiMerchantSubscriptionCreatePreviewRes struct {
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Discount *UnibeeApiBeanMerchantDiscountCode `json:"discount,omitempty"`
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	DiscountMessage *string `json:"discountMessage,omitempty"`
	Email *string `json:"email,omitempty"`
	Gateway *UnibeeApiBeanGateway `json:"gateway,omitempty"`
	Invoice *UnibeeApiBeanInvoice `json:"invoice,omitempty"`
	OriginAmount *int64 `json:"originAmount,omitempty"`
	// other active or incomplete subscription id 
	OtherActiveSubscriptionId *string `json:"otherActiveSubscriptionId,omitempty"`
	OtherPendingCryptoSubscription *UnibeeApiBeanDetailSubscriptionDetail `json:"otherPendingCryptoSubscription,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Quantity *int64 `json:"quantity,omitempty"`
	TaxPercentage *int64 `json:"taxPercentage,omitempty"`
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// trial_end, utc time
	TrialEnd *int64 `json:"trialEnd,omitempty"`
	UserId *int64 `json:"userId,omitempty"`
	VatCountryCode *string `json:"vatCountryCode,omitempty"`
	VatCountryName *string `json:"vatCountryName,omitempty"`
	VatNumber *string `json:"vatNumber,omitempty"`
	VatNumberValidate *UnibeeApiBeanValidResult `json:"vatNumberValidate,omitempty"`
	VatNumberValidateMessage *string `json:"vatNumberValidateMessage,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionCreatePreviewRes instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionCreatePreviewRes() *UnibeeApiMerchantSubscriptionCreatePreviewRes {
	this := UnibeeApiMerchantSubscriptionCreatePreviewRes{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionCreatePreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionCreatePreviewResWithDefaults() *UnibeeApiMerchantSubscriptionCreatePreviewRes {
	this := UnibeeApiMerchantSubscriptionCreatePreviewRes{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetCurrency(v string) {
	o.Currency = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the Discount field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode) {
	o.Discount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountMessage returns the DiscountMessage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountMessage() string {
	if o == nil || IsNil(o.DiscountMessage) {
		var ret string
		return ret
	}
	return *o.DiscountMessage
}

// GetDiscountMessageOk returns a tuple with the DiscountMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountMessage) {
		return nil, false
	}
	return o.DiscountMessage, true
}

// HasDiscountMessage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscountMessage() bool {
	if o != nil && !IsNil(o.DiscountMessage) {
		return true
	}

	return false
}

// SetDiscountMessage gets a reference to the given string and assigns it to the DiscountMessage field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscountMessage(v string) {
	o.DiscountMessage = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetEmail(v string) {
	o.Email = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOriginAmount() int64 {
	if o == nil || IsNil(o.OriginAmount) {
		var ret int64
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOriginAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given int64 and assigns it to the OriginAmount field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOriginAmount(v int64) {
	o.OriginAmount = &v
}

// GetOtherActiveSubscriptionId returns the OtherActiveSubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionId() string {
	if o == nil || IsNil(o.OtherActiveSubscriptionId) {
		var ret string
		return ret
	}
	return *o.OtherActiveSubscriptionId
}

// GetOtherActiveSubscriptionIdOk returns a tuple with the OtherActiveSubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.OtherActiveSubscriptionId) {
		return nil, false
	}
	return o.OtherActiveSubscriptionId, true
}

// HasOtherActiveSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOtherActiveSubscriptionId() bool {
	if o != nil && !IsNil(o.OtherActiveSubscriptionId) {
		return true
	}

	return false
}

// SetOtherActiveSubscriptionId gets a reference to the given string and assigns it to the OtherActiveSubscriptionId field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOtherActiveSubscriptionId(v string) {
	o.OtherActiveSubscriptionId = &v
}

// GetOtherPendingCryptoSubscription returns the OtherPendingCryptoSubscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscription() UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.OtherPendingCryptoSubscription) {
		var ret UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return *o.OtherPendingCryptoSubscription
}

// GetOtherPendingCryptoSubscriptionOk returns a tuple with the OtherPendingCryptoSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.OtherPendingCryptoSubscription) {
		return nil, false
	}
	return o.OtherPendingCryptoSubscription, true
}

// HasOtherPendingCryptoSubscription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOtherPendingCryptoSubscription() bool {
	if o != nil && !IsNil(o.OtherPendingCryptoSubscription) {
		return true
	}

	return false
}

// SetOtherPendingCryptoSubscription gets a reference to the given UnibeeApiBeanDetailSubscriptionDetail and assigns it to the OtherPendingCryptoSubscription field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOtherPendingCryptoSubscription(v UnibeeApiBeanDetailSubscriptionDetail) {
	o.OtherPendingCryptoSubscription = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTaxPercentage() int64 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTrialEnd() int64 {
	if o == nil || IsNil(o.TrialEnd) {
		var ret int64
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTrialEndOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int64 and assigns it to the TrialEnd field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTrialEnd(v int64) {
	o.TrialEnd = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetUserId(v int64) {
	o.UserId = &v
}

// GetVatCountryCode returns the VatCountryCode field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryCode() string {
	if o == nil || IsNil(o.VatCountryCode) {
		var ret string
		return ret
	}
	return *o.VatCountryCode
}

// GetVatCountryCodeOk returns a tuple with the VatCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryCode) {
		return nil, false
	}
	return o.VatCountryCode, true
}

// HasVatCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatCountryCode() bool {
	if o != nil && !IsNil(o.VatCountryCode) {
		return true
	}

	return false
}

// SetVatCountryCode gets a reference to the given string and assigns it to the VatCountryCode field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatCountryCode(v string) {
	o.VatCountryCode = &v
}

// GetVatCountryName returns the VatCountryName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryName() string {
	if o == nil || IsNil(o.VatCountryName) {
		var ret string
		return ret
	}
	return *o.VatCountryName
}

// GetVatCountryNameOk returns a tuple with the VatCountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryName) {
		return nil, false
	}
	return o.VatCountryName, true
}

// HasVatCountryName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatCountryName() bool {
	if o != nil && !IsNil(o.VatCountryName) {
		return true
	}

	return false
}

// SetVatCountryName gets a reference to the given string and assigns it to the VatCountryName field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatCountryName(v string) {
	o.VatCountryName = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetVatNumberValidate returns the VatNumberValidate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidate() UnibeeApiBeanValidResult {
	if o == nil || IsNil(o.VatNumberValidate) {
		var ret UnibeeApiBeanValidResult
		return ret
	}
	return *o.VatNumberValidate
}

// GetVatNumberValidateOk returns a tuple with the VatNumberValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool) {
	if o == nil || IsNil(o.VatNumberValidate) {
		return nil, false
	}
	return o.VatNumberValidate, true
}

// HasVatNumberValidate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumberValidate() bool {
	if o != nil && !IsNil(o.VatNumberValidate) {
		return true
	}

	return false
}

// SetVatNumberValidate gets a reference to the given UnibeeApiBeanValidResult and assigns it to the VatNumberValidate field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumberValidate(v UnibeeApiBeanValidResult) {
	o.VatNumberValidate = &v
}

// GetVatNumberValidateMessage returns the VatNumberValidateMessage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateMessage() string {
	if o == nil || IsNil(o.VatNumberValidateMessage) {
		var ret string
		return ret
	}
	return *o.VatNumberValidateMessage
}

// GetVatNumberValidateMessageOk returns a tuple with the VatNumberValidateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumberValidateMessage) {
		return nil, false
	}
	return o.VatNumberValidateMessage, true
}

// HasVatNumberValidateMessage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumberValidateMessage() bool {
	if o != nil && !IsNil(o.VatNumberValidateMessage) {
		return true
	}

	return false
}

// SetVatNumberValidateMessage gets a reference to the given string and assigns it to the VatNumberValidateMessage field.
func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumberValidateMessage(v string) {
	o.VatNumberValidateMessage = &v
}

func (o UnibeeApiMerchantSubscriptionCreatePreviewRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionCreatePreviewRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountMessage) {
		toSerialize["discountMessage"] = o.DiscountMessage
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.OriginAmount) {
		toSerialize["originAmount"] = o.OriginAmount
	}
	if !IsNil(o.OtherActiveSubscriptionId) {
		toSerialize["otherActiveSubscriptionId"] = o.OtherActiveSubscriptionId
	}
	if !IsNil(o.OtherPendingCryptoSubscription) {
		toSerialize["otherPendingCryptoSubscription"] = o.OtherPendingCryptoSubscription
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.TrialEnd) {
		toSerialize["trialEnd"] = o.TrialEnd
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.VatCountryCode) {
		toSerialize["vatCountryCode"] = o.VatCountryCode
	}
	if !IsNil(o.VatCountryName) {
		toSerialize["vatCountryName"] = o.VatCountryName
	}
	if !IsNil(o.VatNumber) {
		toSerialize["vatNumber"] = o.VatNumber
	}
	if !IsNil(o.VatNumberValidate) {
		toSerialize["vatNumberValidate"] = o.VatNumberValidate
	}
	if !IsNil(o.VatNumberValidateMessage) {
		toSerialize["vatNumberValidateMessage"] = o.VatNumberValidateMessage
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionCreatePreviewRes struct {
	value *UnibeeApiMerchantSubscriptionCreatePreviewRes
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) Get() *UnibeeApiMerchantSubscriptionCreatePreviewRes {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) Set(val *UnibeeApiMerchantSubscriptionCreatePreviewRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionCreatePreviewRes(val *UnibeeApiMerchantSubscriptionCreatePreviewRes) *NullableUnibeeApiMerchantSubscriptionCreatePreviewRes {
	return &NullableUnibeeApiMerchantSubscriptionCreatePreviewRes{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionCreatePreviewRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


