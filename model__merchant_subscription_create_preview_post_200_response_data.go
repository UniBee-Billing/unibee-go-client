/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantSubscriptionCreatePreviewPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionCreatePreviewPost200ResponseData{}

// MerchantSubscriptionCreatePreviewPost200ResponseData struct for MerchantSubscriptionCreatePreviewPost200ResponseData
type MerchantSubscriptionCreatePreviewPost200ResponseData struct {
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

// NewMerchantSubscriptionCreatePreviewPost200ResponseData instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionCreatePreviewPost200ResponseData() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	this := MerchantSubscriptionCreatePreviewPost200ResponseData{}
	return &this
}

// NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	this := MerchantSubscriptionCreatePreviewPost200ResponseData{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetCurrency(v string) {
	o.Currency = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscount() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the Discount field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetDiscount(v UnibeeApiBeanMerchantDiscountCode) {
	o.Discount = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountMessage returns the DiscountMessage field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscountMessage() string {
	if o == nil || IsNil(o.DiscountMessage) {
		var ret string
		return ret
	}
	return *o.DiscountMessage
}

// GetDiscountMessageOk returns a tuple with the DiscountMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetDiscountMessageOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountMessage) {
		return nil, false
	}
	return o.DiscountMessage, true
}

// HasDiscountMessage returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasDiscountMessage() bool {
	if o != nil && !IsNil(o.DiscountMessage) {
		return true
	}

	return false
}

// SetDiscountMessage gets a reference to the given string and assigns it to the DiscountMessage field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetDiscountMessage(v string) {
	o.DiscountMessage = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetEmail(v string) {
	o.Email = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGateway() UnibeeApiBeanGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGatewayOk() (*UnibeeApiBeanGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanGateway and assigns it to the Gateway field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetGateway(v UnibeeApiBeanGateway) {
	o.Gateway = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.Invoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the Invoice field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoice) {
	o.Invoice = &v
}

// GetOriginAmount returns the OriginAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetOriginAmount() int64 {
	if o == nil || IsNil(o.OriginAmount) {
		var ret int64
		return ret
	}
	return *o.OriginAmount
}

// GetOriginAmountOk returns a tuple with the OriginAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetOriginAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.OriginAmount) {
		return nil, false
	}
	return o.OriginAmount, true
}

// HasOriginAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasOriginAmount() bool {
	if o != nil && !IsNil(o.OriginAmount) {
		return true
	}

	return false
}

// SetOriginAmount gets a reference to the given int64 and assigns it to the OriginAmount field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetOriginAmount(v int64) {
	o.OriginAmount = &v
}

// GetOtherPendingCryptoSubscription returns the OtherPendingCryptoSubscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetOtherPendingCryptoSubscription() UnibeeApiBeanDetailSubscriptionDetail {
	if o == nil || IsNil(o.OtherPendingCryptoSubscription) {
		var ret UnibeeApiBeanDetailSubscriptionDetail
		return ret
	}
	return *o.OtherPendingCryptoSubscription
}

// GetOtherPendingCryptoSubscriptionOk returns a tuple with the OtherPendingCryptoSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetOtherPendingCryptoSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool) {
	if o == nil || IsNil(o.OtherPendingCryptoSubscription) {
		return nil, false
	}
	return o.OtherPendingCryptoSubscription, true
}

// HasOtherPendingCryptoSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasOtherPendingCryptoSubscription() bool {
	if o != nil && !IsNil(o.OtherPendingCryptoSubscription) {
		return true
	}

	return false
}

// SetOtherPendingCryptoSubscription gets a reference to the given UnibeeApiBeanDetailSubscriptionDetail and assigns it to the OtherPendingCryptoSubscription field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetOtherPendingCryptoSubscription(v UnibeeApiBeanDetailSubscriptionDetail) {
	o.OtherPendingCryptoSubscription = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxPercentage() int64 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int64
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int64 and assigns it to the TaxPercentage field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTaxPercentage(v int64) {
	o.TaxPercentage = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTrialEnd() int64 {
	if o == nil || IsNil(o.TrialEnd) {
		var ret int64
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTrialEndOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialEnd) {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTrialEnd() bool {
	if o != nil && !IsNil(o.TrialEnd) {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given int64 and assigns it to the TrialEnd field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTrialEnd(v int64) {
	o.TrialEnd = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetUserId(v int64) {
	o.UserId = &v
}

// GetVatCountryCode returns the VatCountryCode field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCode() string {
	if o == nil || IsNil(o.VatCountryCode) {
		var ret string
		return ret
	}
	return *o.VatCountryCode
}

// GetVatCountryCodeOk returns a tuple with the VatCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryCode) {
		return nil, false
	}
	return o.VatCountryCode, true
}

// HasVatCountryCode returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryCode() bool {
	if o != nil && !IsNil(o.VatCountryCode) {
		return true
	}

	return false
}

// SetVatCountryCode gets a reference to the given string and assigns it to the VatCountryCode field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryCode(v string) {
	o.VatCountryCode = &v
}

// GetVatCountryName returns the VatCountryName field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryName() string {
	if o == nil || IsNil(o.VatCountryName) {
		var ret string
		return ret
	}
	return *o.VatCountryName
}

// GetVatCountryNameOk returns a tuple with the VatCountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryNameOk() (*string, bool) {
	if o == nil || IsNil(o.VatCountryName) {
		return nil, false
	}
	return o.VatCountryName, true
}

// HasVatCountryName returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryName() bool {
	if o != nil && !IsNil(o.VatCountryName) {
		return true
	}

	return false
}

// SetVatCountryName gets a reference to the given string and assigns it to the VatCountryName field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryName(v string) {
	o.VatCountryName = &v
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber) {
		var ret string
		return ret
	}
	return *o.VatNumber
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumber) {
		return nil, false
	}
	return o.VatNumber, true
}

// HasVatNumber returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumber() bool {
	if o != nil && !IsNil(o.VatNumber) {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given string and assigns it to the VatNumber field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumber(v string) {
	o.VatNumber = &v
}

// GetVatNumberValidate returns the VatNumberValidate field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidate() UnibeeApiBeanValidResult {
	if o == nil || IsNil(o.VatNumberValidate) {
		var ret UnibeeApiBeanValidResult
		return ret
	}
	return *o.VatNumberValidate
}

// GetVatNumberValidateOk returns a tuple with the VatNumberValidate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool) {
	if o == nil || IsNil(o.VatNumberValidate) {
		return nil, false
	}
	return o.VatNumberValidate, true
}

// HasVatNumberValidate returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumberValidate() bool {
	if o != nil && !IsNil(o.VatNumberValidate) {
		return true
	}

	return false
}

// SetVatNumberValidate gets a reference to the given UnibeeApiBeanValidResult and assigns it to the VatNumberValidate field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumberValidate(v UnibeeApiBeanValidResult) {
	o.VatNumberValidate = &v
}

// GetVatNumberValidateMessage returns the VatNumberValidateMessage field value if set, zero value otherwise.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidateMessage() string {
	if o == nil || IsNil(o.VatNumberValidateMessage) {
		var ret string
		return ret
	}
	return *o.VatNumberValidateMessage
}

// GetVatNumberValidateMessageOk returns a tuple with the VatNumberValidateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.VatNumberValidateMessage) {
		return nil, false
	}
	return o.VatNumberValidateMessage, true
}

// HasVatNumberValidateMessage returns a boolean if a field has been set.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumberValidateMessage() bool {
	if o != nil && !IsNil(o.VatNumberValidateMessage) {
		return true
	}

	return false
}

// SetVatNumberValidateMessage gets a reference to the given string and assigns it to the VatNumberValidateMessage field.
func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumberValidateMessage(v string) {
	o.VatNumberValidateMessage = &v
}

func (o MerchantSubscriptionCreatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionCreatePreviewPost200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullableMerchantSubscriptionCreatePreviewPost200ResponseData struct {
	value *MerchantSubscriptionCreatePreviewPost200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Get() *MerchantSubscriptionCreatePreviewPost200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Set(val *MerchantSubscriptionCreatePreviewPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionCreatePreviewPost200ResponseData(val *MerchantSubscriptionCreatePreviewPost200ResponseData) *NullableMerchantSubscriptionCreatePreviewPost200ResponseData {
	return &NullableMerchantSubscriptionCreatePreviewPost200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionCreatePreviewPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionCreatePreviewPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


