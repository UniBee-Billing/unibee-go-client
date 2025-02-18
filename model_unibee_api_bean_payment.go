/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPayment{}

// UnibeeApiBeanPayment struct for UnibeeApiBeanPayment
type UnibeeApiBeanPayment struct {
	AuthorizeReason *string `json:"authorizeReason,omitempty"`
	// authorize status，0-waiting authorize，1-authorized，2-authorized_request
	AuthorizeStatus *int32 `json:"authorizeStatus,omitempty"`
	AutoCharge *bool `json:"autoCharge,omitempty"`
	Automatic *int32 `json:"automatic,omitempty"`
	// balance_amount
	BalanceAmount *int64 `json:"balanceAmount,omitempty"`
	BillingReason *string `json:"billingReason,omitempty"`
	// cancel time, utc time
	CancelTime *int64 `json:"cancelTime,omitempty"`
	// country code
	CountryCode *string `json:"countryCode,omitempty"`
	// create time, utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency，“SGD” “MYR” “PHP” “IDR” “THB”
	Currency *string `json:"currency,omitempty"`
	// external_payment_id
	ExternalPaymentId *string `json:"externalPaymentId,omitempty"`
	FailureReason *string `json:"failureReason,omitempty"`
	// who pay the gas, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	GatewayCurrencyExchange *UnibeeApiBeanGatewayCurrencyExchange `json:"gatewayCurrencyExchange,omitempty"`
	// gateway_id
	GatewayId *int64 `json:"gatewayId,omitempty"`
	// gateway_payment_id
	GatewayPaymentId *string `json:"gatewayPaymentId,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	Link *string `json:"link,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// paid time, utc time
	PaidTime *int64 `json:"paidTime,omitempty"`
	// payment_amount
	PaymentAmount *int64 `json:"paymentAmount,omitempty"`
	// payment id
	PaymentId *string `json:"paymentId,omitempty"`
	// total refund amount
	RefundAmount *int64 `json:"refundAmount,omitempty"`
	// return url
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// status  10-pending，20-success，30-failure, 40-cancel
	Status *int32 `json:"status,omitempty"`
	// subscription id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// total amount
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanPayment instantiates a new UnibeeApiBeanPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPayment() *UnibeeApiBeanPayment {
	this := UnibeeApiBeanPayment{}
	return &this
}

// NewUnibeeApiBeanPaymentWithDefaults instantiates a new UnibeeApiBeanPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPaymentWithDefaults() *UnibeeApiBeanPayment {
	this := UnibeeApiBeanPayment{}
	return &this
}

// GetAuthorizeReason returns the AuthorizeReason field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetAuthorizeReason() string {
	if o == nil || IsNil(o.AuthorizeReason) {
		var ret string
		return ret
	}
	return *o.AuthorizeReason
}

// GetAuthorizeReasonOk returns a tuple with the AuthorizeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetAuthorizeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizeReason) {
		return nil, false
	}
	return o.AuthorizeReason, true
}

// HasAuthorizeReason returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasAuthorizeReason() bool {
	if o != nil && !IsNil(o.AuthorizeReason) {
		return true
	}

	return false
}

// SetAuthorizeReason gets a reference to the given string and assigns it to the AuthorizeReason field.
func (o *UnibeeApiBeanPayment) SetAuthorizeReason(v string) {
	o.AuthorizeReason = &v
}

// GetAuthorizeStatus returns the AuthorizeStatus field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetAuthorizeStatus() int32 {
	if o == nil || IsNil(o.AuthorizeStatus) {
		var ret int32
		return ret
	}
	return *o.AuthorizeStatus
}

// GetAuthorizeStatusOk returns a tuple with the AuthorizeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetAuthorizeStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.AuthorizeStatus) {
		return nil, false
	}
	return o.AuthorizeStatus, true
}

// HasAuthorizeStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasAuthorizeStatus() bool {
	if o != nil && !IsNil(o.AuthorizeStatus) {
		return true
	}

	return false
}

// SetAuthorizeStatus gets a reference to the given int32 and assigns it to the AuthorizeStatus field.
func (o *UnibeeApiBeanPayment) SetAuthorizeStatus(v int32) {
	o.AuthorizeStatus = &v
}

// GetAutoCharge returns the AutoCharge field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetAutoCharge() bool {
	if o == nil || IsNil(o.AutoCharge) {
		var ret bool
		return ret
	}
	return *o.AutoCharge
}

// GetAutoChargeOk returns a tuple with the AutoCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetAutoChargeOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCharge) {
		return nil, false
	}
	return o.AutoCharge, true
}

// HasAutoCharge returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasAutoCharge() bool {
	if o != nil && !IsNil(o.AutoCharge) {
		return true
	}

	return false
}

// SetAutoCharge gets a reference to the given bool and assigns it to the AutoCharge field.
func (o *UnibeeApiBeanPayment) SetAutoCharge(v bool) {
	o.AutoCharge = &v
}

// GetAutomatic returns the Automatic field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetAutomatic() int32 {
	if o == nil || IsNil(o.Automatic) {
		var ret int32
		return ret
	}
	return *o.Automatic
}

// GetAutomaticOk returns a tuple with the Automatic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetAutomaticOk() (*int32, bool) {
	if o == nil || IsNil(o.Automatic) {
		return nil, false
	}
	return o.Automatic, true
}

// HasAutomatic returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasAutomatic() bool {
	if o != nil && !IsNil(o.Automatic) {
		return true
	}

	return false
}

// SetAutomatic gets a reference to the given int32 and assigns it to the Automatic field.
func (o *UnibeeApiBeanPayment) SetAutomatic(v int32) {
	o.Automatic = &v
}

// GetBalanceAmount returns the BalanceAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetBalanceAmount() int64 {
	if o == nil || IsNil(o.BalanceAmount) {
		var ret int64
		return ret
	}
	return *o.BalanceAmount
}

// GetBalanceAmountOk returns a tuple with the BalanceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetBalanceAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.BalanceAmount) {
		return nil, false
	}
	return o.BalanceAmount, true
}

// HasBalanceAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasBalanceAmount() bool {
	if o != nil && !IsNil(o.BalanceAmount) {
		return true
	}

	return false
}

// SetBalanceAmount gets a reference to the given int64 and assigns it to the BalanceAmount field.
func (o *UnibeeApiBeanPayment) SetBalanceAmount(v int64) {
	o.BalanceAmount = &v
}

// GetBillingReason returns the BillingReason field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetBillingReason() string {
	if o == nil || IsNil(o.BillingReason) {
		var ret string
		return ret
	}
	return *o.BillingReason
}

// GetBillingReasonOk returns a tuple with the BillingReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetBillingReasonOk() (*string, bool) {
	if o == nil || IsNil(o.BillingReason) {
		return nil, false
	}
	return o.BillingReason, true
}

// HasBillingReason returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasBillingReason() bool {
	if o != nil && !IsNil(o.BillingReason) {
		return true
	}

	return false
}

// SetBillingReason gets a reference to the given string and assigns it to the BillingReason field.
func (o *UnibeeApiBeanPayment) SetBillingReason(v string) {
	o.BillingReason = &v
}

// GetCancelTime returns the CancelTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetCancelTime() int64 {
	if o == nil || IsNil(o.CancelTime) {
		var ret int64
		return ret
	}
	return *o.CancelTime
}

// GetCancelTimeOk returns a tuple with the CancelTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetCancelTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CancelTime) {
		return nil, false
	}
	return o.CancelTime, true
}

// HasCancelTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasCancelTime() bool {
	if o != nil && !IsNil(o.CancelTime) {
		return true
	}

	return false
}

// SetCancelTime gets a reference to the given int64 and assigns it to the CancelTime field.
func (o *UnibeeApiBeanPayment) SetCancelTime(v int64) {
	o.CancelTime = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *UnibeeApiBeanPayment) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanPayment) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanPayment) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalPaymentId returns the ExternalPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetExternalPaymentId() string {
	if o == nil || IsNil(o.ExternalPaymentId) {
		var ret string
		return ret
	}
	return *o.ExternalPaymentId
}

// GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetExternalPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPaymentId) {
		return nil, false
	}
	return o.ExternalPaymentId, true
}

// HasExternalPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasExternalPaymentId() bool {
	if o != nil && !IsNil(o.ExternalPaymentId) {
		return true
	}

	return false
}

// SetExternalPaymentId gets a reference to the given string and assigns it to the ExternalPaymentId field.
func (o *UnibeeApiBeanPayment) SetExternalPaymentId(v string) {
	o.ExternalPaymentId = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *UnibeeApiBeanPayment) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiBeanPayment) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetGatewayCurrencyExchange returns the GatewayCurrencyExchange field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetGatewayCurrencyExchange() UnibeeApiBeanGatewayCurrencyExchange {
	if o == nil || IsNil(o.GatewayCurrencyExchange) {
		var ret UnibeeApiBeanGatewayCurrencyExchange
		return ret
	}
	return *o.GatewayCurrencyExchange
}

// GetGatewayCurrencyExchangeOk returns a tuple with the GatewayCurrencyExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetGatewayCurrencyExchangeOk() (*UnibeeApiBeanGatewayCurrencyExchange, bool) {
	if o == nil || IsNil(o.GatewayCurrencyExchange) {
		return nil, false
	}
	return o.GatewayCurrencyExchange, true
}

// HasGatewayCurrencyExchange returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasGatewayCurrencyExchange() bool {
	if o != nil && !IsNil(o.GatewayCurrencyExchange) {
		return true
	}

	return false
}

// SetGatewayCurrencyExchange gets a reference to the given UnibeeApiBeanGatewayCurrencyExchange and assigns it to the GatewayCurrencyExchange field.
func (o *UnibeeApiBeanPayment) SetGatewayCurrencyExchange(v UnibeeApiBeanGatewayCurrencyExchange) {
	o.GatewayCurrencyExchange = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanPayment) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayPaymentId returns the GatewayPaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetGatewayPaymentId() string {
	if o == nil || IsNil(o.GatewayPaymentId) {
		var ret string
		return ret
	}
	return *o.GatewayPaymentId
}

// GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetGatewayPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayPaymentId) {
		return nil, false
	}
	return o.GatewayPaymentId, true
}

// HasGatewayPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasGatewayPaymentId() bool {
	if o != nil && !IsNil(o.GatewayPaymentId) {
		return true
	}

	return false
}

// SetGatewayPaymentId gets a reference to the given string and assigns it to the GatewayPaymentId field.
func (o *UnibeeApiBeanPayment) SetGatewayPaymentId(v string) {
	o.GatewayPaymentId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanPayment) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *UnibeeApiBeanPayment) SetLink(v string) {
	o.Link = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanPayment) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanPayment) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPaidTime returns the PaidTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetPaidTime() int64 {
	if o == nil || IsNil(o.PaidTime) {
		var ret int64
		return ret
	}
	return *o.PaidTime
}

// GetPaidTimeOk returns a tuple with the PaidTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetPaidTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PaidTime) {
		return nil, false
	}
	return o.PaidTime, true
}

// HasPaidTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasPaidTime() bool {
	if o != nil && !IsNil(o.PaidTime) {
		return true
	}

	return false
}

// SetPaidTime gets a reference to the given int64 and assigns it to the PaidTime field.
func (o *UnibeeApiBeanPayment) SetPaidTime(v int64) {
	o.PaidTime = &v
}

// GetPaymentAmount returns the PaymentAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetPaymentAmount() int64 {
	if o == nil || IsNil(o.PaymentAmount) {
		var ret int64
		return ret
	}
	return *o.PaymentAmount
}

// GetPaymentAmountOk returns a tuple with the PaymentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetPaymentAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.PaymentAmount) {
		return nil, false
	}
	return o.PaymentAmount, true
}

// HasPaymentAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasPaymentAmount() bool {
	if o != nil && !IsNil(o.PaymentAmount) {
		return true
	}

	return false
}

// SetPaymentAmount gets a reference to the given int64 and assigns it to the PaymentAmount field.
func (o *UnibeeApiBeanPayment) SetPaymentAmount(v int64) {
	o.PaymentAmount = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *UnibeeApiBeanPayment) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetRefundAmount returns the RefundAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetRefundAmount() int64 {
	if o == nil || IsNil(o.RefundAmount) {
		var ret int64
		return ret
	}
	return *o.RefundAmount
}

// GetRefundAmountOk returns a tuple with the RefundAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetRefundAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.RefundAmount) {
		return nil, false
	}
	return o.RefundAmount, true
}

// HasRefundAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasRefundAmount() bool {
	if o != nil && !IsNil(o.RefundAmount) {
		return true
	}

	return false
}

// SetRefundAmount gets a reference to the given int64 and assigns it to the RefundAmount field.
func (o *UnibeeApiBeanPayment) SetRefundAmount(v int64) {
	o.RefundAmount = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *UnibeeApiBeanPayment) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanPayment) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanPayment) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetTotalAmount() int64 {
	if o == nil || IsNil(o.TotalAmount) {
		var ret int64
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetTotalAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given int64 and assigns it to the TotalAmount field.
func (o *UnibeeApiBeanPayment) SetTotalAmount(v int64) {
	o.TotalAmount = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPayment) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPayment) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPayment) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanPayment) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizeReason) {
		toSerialize["authorizeReason"] = o.AuthorizeReason
	}
	if !IsNil(o.AuthorizeStatus) {
		toSerialize["authorizeStatus"] = o.AuthorizeStatus
	}
	if !IsNil(o.AutoCharge) {
		toSerialize["autoCharge"] = o.AutoCharge
	}
	if !IsNil(o.Automatic) {
		toSerialize["automatic"] = o.Automatic
	}
	if !IsNil(o.BalanceAmount) {
		toSerialize["balanceAmount"] = o.BalanceAmount
	}
	if !IsNil(o.BillingReason) {
		toSerialize["billingReason"] = o.BillingReason
	}
	if !IsNil(o.CancelTime) {
		toSerialize["cancelTime"] = o.CancelTime
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExternalPaymentId) {
		toSerialize["externalPaymentId"] = o.ExternalPaymentId
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.GasPayer) {
		toSerialize["gasPayer"] = o.GasPayer
	}
	if !IsNil(o.GatewayCurrencyExchange) {
		toSerialize["gatewayCurrencyExchange"] = o.GatewayCurrencyExchange
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayPaymentId) {
		toSerialize["gatewayPaymentId"] = o.GatewayPaymentId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PaidTime) {
		toSerialize["paidTime"] = o.PaidTime
	}
	if !IsNil(o.PaymentAmount) {
		toSerialize["paymentAmount"] = o.PaymentAmount
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.RefundAmount) {
		toSerialize["refundAmount"] = o.RefundAmount
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
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanPayment struct {
	value *UnibeeApiBeanPayment
	isSet bool
}

func (v NullableUnibeeApiBeanPayment) Get() *UnibeeApiBeanPayment {
	return v.value
}

func (v *NullableUnibeeApiBeanPayment) Set(val *UnibeeApiBeanPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPayment(val *UnibeeApiBeanPayment) *NullableUnibeeApiBeanPayment {
	return &NullableUnibeeApiBeanPayment{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


