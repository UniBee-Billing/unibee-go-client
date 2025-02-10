/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailGateway{}

// UnibeeApiBeanDetailGateway struct for UnibeeApiBeanDetailGateway
type UnibeeApiBeanDetailGateway struct {
	// Whether the gateway finished setup process
	IsSetupFinished *bool `json:"IsSetupFinished,omitempty"`
	Archive *bool `json:"archive,omitempty"`
	Bank *UnibeeApiBeanDetailGatewayBank `json:"bank,omitempty"`
	CountryConfig *map[string]bool `json:"countryConfig,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// The currency of wire transfer 
	Currency *string `json:"currency,omitempty"`
	// The currency exchange for gateway payment, effect at start of payment creation when currency matched
	CurrencyExchange []UnibeeApiBeanDetailGatewayCurrencyExchange `json:"currencyExchange,omitempty"`
	// whether to enable currency exchange
	CurrencyExchangeEnabled *bool `json:"currencyExchangeEnabled,omitempty"`
	// The description of gateway
	Description *string `json:"description,omitempty"`
	// The gateway display name, used at user portal
	DisplayName *string `json:"displayName,omitempty"`
	// The gateway display name, used at user portal
	GatewayIcons []string `json:"gatewayIcons,omitempty"`
	GatewayId *int64 `json:"gatewayId,omitempty"`
	GatewayKey *string `json:"gatewayKey,omitempty"`
	GatewayLogo *string `json:"gatewayLogo,omitempty"`
	// The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium
	GatewayName *string `json:"gatewayName,omitempty"`
	GatewaySecret *string `json:"gatewaySecret,omitempty"`
	// gateway type，1-Bank Card ｜ 2-Crypto | 3 - Wire Transfer
	GatewayType *int64 `json:"gatewayType,omitempty"`
	// The gateway webhook integration guide link, gateway webhook need setup if not blank
	GatewayWebhookIntegrationLink *string `json:"gatewayWebhookIntegrationLink,omitempty"`
	// The gateway website link
	GatewayWebsiteLink *string `json:"gatewayWebsiteLink,omitempty"`
	// The minimum amount of wire transfer
	MinimumAmount *int64 `json:"minimumAmount,omitempty"`
	// The name of gateway
	Name *string `json:"name,omitempty"`
	PrivateSecretName *string `json:"privateSecretName,omitempty"`
	PublicKeyName *string `json:"publicKeyName,omitempty"`
	// The sort value of payment gateway, The bigger, the closer to the front
	Sort *int64 `json:"sort,omitempty"`
	SubGateway *string `json:"subGateway,omitempty"`
	SubGatewayConfigs []UnibeeInternalInterfaceSubGatewayConfig `json:"subGatewayConfigs,omitempty"`
	SubGatewayName *string `json:"subGatewayName,omitempty"`
	// The endpoint url of gateway webhook 
	WebhookEndpointUrl *string `json:"webhookEndpointUrl,omitempty"`
	// The secret of gateway webhook
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

// NewUnibeeApiBeanDetailGateway instantiates a new UnibeeApiBeanDetailGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailGateway() *UnibeeApiBeanDetailGateway {
	this := UnibeeApiBeanDetailGateway{}
	return &this
}

// NewUnibeeApiBeanDetailGatewayWithDefaults instantiates a new UnibeeApiBeanDetailGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailGatewayWithDefaults() *UnibeeApiBeanDetailGateway {
	this := UnibeeApiBeanDetailGateway{}
	return &this
}

// GetIsSetupFinished returns the IsSetupFinished field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetIsSetupFinished() bool {
	if o == nil || IsNil(o.IsSetupFinished) {
		var ret bool
		return ret
	}
	return *o.IsSetupFinished
}

// GetIsSetupFinishedOk returns a tuple with the IsSetupFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetIsSetupFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSetupFinished) {
		return nil, false
	}
	return o.IsSetupFinished, true
}

// HasIsSetupFinished returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasIsSetupFinished() bool {
	if o != nil && !IsNil(o.IsSetupFinished) {
		return true
	}

	return false
}

// SetIsSetupFinished gets a reference to the given bool and assigns it to the IsSetupFinished field.
func (o *UnibeeApiBeanDetailGateway) SetIsSetupFinished(v bool) {
	o.IsSetupFinished = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetArchive() bool {
	if o == nil || IsNil(o.Archive) {
		var ret bool
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetArchiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Archive) {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasArchive() bool {
	if o != nil && !IsNil(o.Archive) {
		return true
	}

	return false
}

// SetArchive gets a reference to the given bool and assigns it to the Archive field.
func (o *UnibeeApiBeanDetailGateway) SetArchive(v bool) {
	o.Archive = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetBank() UnibeeApiBeanDetailGatewayBank {
	if o == nil || IsNil(o.Bank) {
		var ret UnibeeApiBeanDetailGatewayBank
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetBankOk() (*UnibeeApiBeanDetailGatewayBank, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given UnibeeApiBeanDetailGatewayBank and assigns it to the Bank field.
func (o *UnibeeApiBeanDetailGateway) SetBank(v UnibeeApiBeanDetailGatewayBank) {
	o.Bank = &v
}

// GetCountryConfig returns the CountryConfig field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetCountryConfig() map[string]bool {
	if o == nil || IsNil(o.CountryConfig) {
		var ret map[string]bool
		return ret
	}
	return *o.CountryConfig
}

// GetCountryConfigOk returns a tuple with the CountryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetCountryConfigOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.CountryConfig) {
		return nil, false
	}
	return o.CountryConfig, true
}

// HasCountryConfig returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasCountryConfig() bool {
	if o != nil && !IsNil(o.CountryConfig) {
		return true
	}

	return false
}

// SetCountryConfig gets a reference to the given map[string]bool and assigns it to the CountryConfig field.
func (o *UnibeeApiBeanDetailGateway) SetCountryConfig(v map[string]bool) {
	o.CountryConfig = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailGateway) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanDetailGateway) SetCurrency(v string) {
	o.Currency = &v
}

// GetCurrencyExchange returns the CurrencyExchange field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchange() []UnibeeApiBeanDetailGatewayCurrencyExchange {
	if o == nil || IsNil(o.CurrencyExchange) {
		var ret []UnibeeApiBeanDetailGatewayCurrencyExchange
		return ret
	}
	return o.CurrencyExchange
}

// GetCurrencyExchangeOk returns a tuple with the CurrencyExchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeOk() ([]UnibeeApiBeanDetailGatewayCurrencyExchange, bool) {
	if o == nil || IsNil(o.CurrencyExchange) {
		return nil, false
	}
	return o.CurrencyExchange, true
}

// HasCurrencyExchange returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasCurrencyExchange() bool {
	if o != nil && !IsNil(o.CurrencyExchange) {
		return true
	}

	return false
}

// SetCurrencyExchange gets a reference to the given []UnibeeApiBeanDetailGatewayCurrencyExchange and assigns it to the CurrencyExchange field.
func (o *UnibeeApiBeanDetailGateway) SetCurrencyExchange(v []UnibeeApiBeanDetailGatewayCurrencyExchange) {
	o.CurrencyExchange = v
}

// GetCurrencyExchangeEnabled returns the CurrencyExchangeEnabled field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeEnabled() bool {
	if o == nil || IsNil(o.CurrencyExchangeEnabled) {
		var ret bool
		return ret
	}
	return *o.CurrencyExchangeEnabled
}

// GetCurrencyExchangeEnabledOk returns a tuple with the CurrencyExchangeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CurrencyExchangeEnabled) {
		return nil, false
	}
	return o.CurrencyExchangeEnabled, true
}

// HasCurrencyExchangeEnabled returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasCurrencyExchangeEnabled() bool {
	if o != nil && !IsNil(o.CurrencyExchangeEnabled) {
		return true
	}

	return false
}

// SetCurrencyExchangeEnabled gets a reference to the given bool and assigns it to the CurrencyExchangeEnabled field.
func (o *UnibeeApiBeanDetailGateway) SetCurrencyExchangeEnabled(v bool) {
	o.CurrencyExchangeEnabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanDetailGateway) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UnibeeApiBeanDetailGateway) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGatewayIcons returns the GatewayIcons field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayIcons() []string {
	if o == nil || IsNil(o.GatewayIcons) {
		var ret []string
		return ret
	}
	return o.GatewayIcons
}

// GetGatewayIconsOk returns a tuple with the GatewayIcons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayIconsOk() ([]string, bool) {
	if o == nil || IsNil(o.GatewayIcons) {
		return nil, false
	}
	return o.GatewayIcons, true
}

// HasGatewayIcons returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayIcons() bool {
	if o != nil && !IsNil(o.GatewayIcons) {
		return true
	}

	return false
}

// SetGatewayIcons gets a reference to the given []string and assigns it to the GatewayIcons field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayIcons(v []string) {
	o.GatewayIcons = v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetGatewayLogo returns the GatewayLogo field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayLogo() string {
	if o == nil || IsNil(o.GatewayLogo) {
		var ret string
		return ret
	}
	return *o.GatewayLogo
}

// GetGatewayLogoOk returns a tuple with the GatewayLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayLogoOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayLogo) {
		return nil, false
	}
	return o.GatewayLogo, true
}

// HasGatewayLogo returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayLogo() bool {
	if o != nil && !IsNil(o.GatewayLogo) {
		return true
	}

	return false
}

// SetGatewayLogo gets a reference to the given string and assigns it to the GatewayLogo field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayLogo(v string) {
	o.GatewayLogo = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetGatewaySecret returns the GatewaySecret field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewaySecret() string {
	if o == nil || IsNil(o.GatewaySecret) {
		var ret string
		return ret
	}
	return *o.GatewaySecret
}

// GetGatewaySecretOk returns a tuple with the GatewaySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewaySecretOk() (*string, bool) {
	if o == nil || IsNil(o.GatewaySecret) {
		return nil, false
	}
	return o.GatewaySecret, true
}

// HasGatewaySecret returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewaySecret() bool {
	if o != nil && !IsNil(o.GatewaySecret) {
		return true
	}

	return false
}

// SetGatewaySecret gets a reference to the given string and assigns it to the GatewaySecret field.
func (o *UnibeeApiBeanDetailGateway) SetGatewaySecret(v string) {
	o.GatewaySecret = &v
}

// GetGatewayType returns the GatewayType field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayType() int64 {
	if o == nil || IsNil(o.GatewayType) {
		var ret int64
		return ret
	}
	return *o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayType) {
		return nil, false
	}
	return o.GatewayType, true
}

// HasGatewayType returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayType() bool {
	if o != nil && !IsNil(o.GatewayType) {
		return true
	}

	return false
}

// SetGatewayType gets a reference to the given int64 and assigns it to the GatewayType field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayType(v int64) {
	o.GatewayType = &v
}

// GetGatewayWebhookIntegrationLink returns the GatewayWebhookIntegrationLink field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayWebhookIntegrationLink() string {
	if o == nil || IsNil(o.GatewayWebhookIntegrationLink) {
		var ret string
		return ret
	}
	return *o.GatewayWebhookIntegrationLink
}

// GetGatewayWebhookIntegrationLinkOk returns a tuple with the GatewayWebhookIntegrationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayWebhookIntegrationLinkOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayWebhookIntegrationLink) {
		return nil, false
	}
	return o.GatewayWebhookIntegrationLink, true
}

// HasGatewayWebhookIntegrationLink returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayWebhookIntegrationLink() bool {
	if o != nil && !IsNil(o.GatewayWebhookIntegrationLink) {
		return true
	}

	return false
}

// SetGatewayWebhookIntegrationLink gets a reference to the given string and assigns it to the GatewayWebhookIntegrationLink field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayWebhookIntegrationLink(v string) {
	o.GatewayWebhookIntegrationLink = &v
}

// GetGatewayWebsiteLink returns the GatewayWebsiteLink field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetGatewayWebsiteLink() string {
	if o == nil || IsNil(o.GatewayWebsiteLink) {
		var ret string
		return ret
	}
	return *o.GatewayWebsiteLink
}

// GetGatewayWebsiteLinkOk returns a tuple with the GatewayWebsiteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetGatewayWebsiteLinkOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayWebsiteLink) {
		return nil, false
	}
	return o.GatewayWebsiteLink, true
}

// HasGatewayWebsiteLink returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasGatewayWebsiteLink() bool {
	if o != nil && !IsNil(o.GatewayWebsiteLink) {
		return true
	}

	return false
}

// SetGatewayWebsiteLink gets a reference to the given string and assigns it to the GatewayWebsiteLink field.
func (o *UnibeeApiBeanDetailGateway) SetGatewayWebsiteLink(v string) {
	o.GatewayWebsiteLink = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetMinimumAmount() int64 {
	if o == nil || IsNil(o.MinimumAmount) {
		var ret int64
		return ret
	}
	return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetMinimumAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumAmount) {
		return nil, false
	}
	return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasMinimumAmount() bool {
	if o != nil && !IsNil(o.MinimumAmount) {
		return true
	}

	return false
}

// SetMinimumAmount gets a reference to the given int64 and assigns it to the MinimumAmount field.
func (o *UnibeeApiBeanDetailGateway) SetMinimumAmount(v int64) {
	o.MinimumAmount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanDetailGateway) SetName(v string) {
	o.Name = &v
}

// GetPrivateSecretName returns the PrivateSecretName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetPrivateSecretName() string {
	if o == nil || IsNil(o.PrivateSecretName) {
		var ret string
		return ret
	}
	return *o.PrivateSecretName
}

// GetPrivateSecretNameOk returns a tuple with the PrivateSecretName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetPrivateSecretNameOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateSecretName) {
		return nil, false
	}
	return o.PrivateSecretName, true
}

// HasPrivateSecretName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasPrivateSecretName() bool {
	if o != nil && !IsNil(o.PrivateSecretName) {
		return true
	}

	return false
}

// SetPrivateSecretName gets a reference to the given string and assigns it to the PrivateSecretName field.
func (o *UnibeeApiBeanDetailGateway) SetPrivateSecretName(v string) {
	o.PrivateSecretName = &v
}

// GetPublicKeyName returns the PublicKeyName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetPublicKeyName() string {
	if o == nil || IsNil(o.PublicKeyName) {
		var ret string
		return ret
	}
	return *o.PublicKeyName
}

// GetPublicKeyNameOk returns a tuple with the PublicKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetPublicKeyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKeyName) {
		return nil, false
	}
	return o.PublicKeyName, true
}

// HasPublicKeyName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasPublicKeyName() bool {
	if o != nil && !IsNil(o.PublicKeyName) {
		return true
	}

	return false
}

// SetPublicKeyName gets a reference to the given string and assigns it to the PublicKeyName field.
func (o *UnibeeApiBeanDetailGateway) SetPublicKeyName(v string) {
	o.PublicKeyName = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetSort() int64 {
	if o == nil || IsNil(o.Sort) {
		var ret int64
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetSortOk() (*int64, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int64 and assigns it to the Sort field.
func (o *UnibeeApiBeanDetailGateway) SetSort(v int64) {
	o.Sort = &v
}

// GetSubGateway returns the SubGateway field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetSubGateway() string {
	if o == nil || IsNil(o.SubGateway) {
		var ret string
		return ret
	}
	return *o.SubGateway
}

// GetSubGatewayOk returns a tuple with the SubGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetSubGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.SubGateway) {
		return nil, false
	}
	return o.SubGateway, true
}

// HasSubGateway returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasSubGateway() bool {
	if o != nil && !IsNil(o.SubGateway) {
		return true
	}

	return false
}

// SetSubGateway gets a reference to the given string and assigns it to the SubGateway field.
func (o *UnibeeApiBeanDetailGateway) SetSubGateway(v string) {
	o.SubGateway = &v
}

// GetSubGatewayConfigs returns the SubGatewayConfigs field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetSubGatewayConfigs() []UnibeeInternalInterfaceSubGatewayConfig {
	if o == nil || IsNil(o.SubGatewayConfigs) {
		var ret []UnibeeInternalInterfaceSubGatewayConfig
		return ret
	}
	return o.SubGatewayConfigs
}

// GetSubGatewayConfigsOk returns a tuple with the SubGatewayConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetSubGatewayConfigsOk() ([]UnibeeInternalInterfaceSubGatewayConfig, bool) {
	if o == nil || IsNil(o.SubGatewayConfigs) {
		return nil, false
	}
	return o.SubGatewayConfigs, true
}

// HasSubGatewayConfigs returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasSubGatewayConfigs() bool {
	if o != nil && !IsNil(o.SubGatewayConfigs) {
		return true
	}

	return false
}

// SetSubGatewayConfigs gets a reference to the given []UnibeeInternalInterfaceSubGatewayConfig and assigns it to the SubGatewayConfigs field.
func (o *UnibeeApiBeanDetailGateway) SetSubGatewayConfigs(v []UnibeeInternalInterfaceSubGatewayConfig) {
	o.SubGatewayConfigs = v
}

// GetSubGatewayName returns the SubGatewayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetSubGatewayName() string {
	if o == nil || IsNil(o.SubGatewayName) {
		var ret string
		return ret
	}
	return *o.SubGatewayName
}

// GetSubGatewayNameOk returns a tuple with the SubGatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetSubGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubGatewayName) {
		return nil, false
	}
	return o.SubGatewayName, true
}

// HasSubGatewayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasSubGatewayName() bool {
	if o != nil && !IsNil(o.SubGatewayName) {
		return true
	}

	return false
}

// SetSubGatewayName gets a reference to the given string and assigns it to the SubGatewayName field.
func (o *UnibeeApiBeanDetailGateway) SetSubGatewayName(v string) {
	o.SubGatewayName = &v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetWebhookEndpointUrl() string {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		var ret string
		return ret
	}
	return *o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetWebhookEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		return nil, false
	}
	return o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasWebhookEndpointUrl() bool {
	if o != nil && !IsNil(o.WebhookEndpointUrl) {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given string and assigns it to the WebhookEndpointUrl field.
func (o *UnibeeApiBeanDetailGateway) SetWebhookEndpointUrl(v string) {
	o.WebhookEndpointUrl = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailGateway) GetWebhookSecret() string {
	if o == nil || IsNil(o.WebhookSecret) {
		var ret string
		return ret
	}
	return *o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailGateway) GetWebhookSecretOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookSecret) {
		return nil, false
	}
	return o.WebhookSecret, true
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailGateway) HasWebhookSecret() bool {
	if o != nil && !IsNil(o.WebhookSecret) {
		return true
	}

	return false
}

// SetWebhookSecret gets a reference to the given string and assigns it to the WebhookSecret field.
func (o *UnibeeApiBeanDetailGateway) SetWebhookSecret(v string) {
	o.WebhookSecret = &v
}

func (o UnibeeApiBeanDetailGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSetupFinished) {
		toSerialize["IsSetupFinished"] = o.IsSetupFinished
	}
	if !IsNil(o.Archive) {
		toSerialize["archive"] = o.Archive
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.CountryConfig) {
		toSerialize["countryConfig"] = o.CountryConfig
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CurrencyExchange) {
		toSerialize["currencyExchange"] = o.CurrencyExchange
	}
	if !IsNil(o.CurrencyExchangeEnabled) {
		toSerialize["currencyExchangeEnabled"] = o.CurrencyExchangeEnabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.GatewayIcons) {
		toSerialize["gatewayIcons"] = o.GatewayIcons
	}
	if !IsNil(o.GatewayId) {
		toSerialize["gatewayId"] = o.GatewayId
	}
	if !IsNil(o.GatewayKey) {
		toSerialize["gatewayKey"] = o.GatewayKey
	}
	if !IsNil(o.GatewayLogo) {
		toSerialize["gatewayLogo"] = o.GatewayLogo
	}
	if !IsNil(o.GatewayName) {
		toSerialize["gatewayName"] = o.GatewayName
	}
	if !IsNil(o.GatewaySecret) {
		toSerialize["gatewaySecret"] = o.GatewaySecret
	}
	if !IsNil(o.GatewayType) {
		toSerialize["gatewayType"] = o.GatewayType
	}
	if !IsNil(o.GatewayWebhookIntegrationLink) {
		toSerialize["gatewayWebhookIntegrationLink"] = o.GatewayWebhookIntegrationLink
	}
	if !IsNil(o.GatewayWebsiteLink) {
		toSerialize["gatewayWebsiteLink"] = o.GatewayWebsiteLink
	}
	if !IsNil(o.MinimumAmount) {
		toSerialize["minimumAmount"] = o.MinimumAmount
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PrivateSecretName) {
		toSerialize["privateSecretName"] = o.PrivateSecretName
	}
	if !IsNil(o.PublicKeyName) {
		toSerialize["publicKeyName"] = o.PublicKeyName
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !IsNil(o.SubGateway) {
		toSerialize["subGateway"] = o.SubGateway
	}
	if !IsNil(o.SubGatewayConfigs) {
		toSerialize["subGatewayConfigs"] = o.SubGatewayConfigs
	}
	if !IsNil(o.SubGatewayName) {
		toSerialize["subGatewayName"] = o.SubGatewayName
	}
	if !IsNil(o.WebhookEndpointUrl) {
		toSerialize["webhookEndpointUrl"] = o.WebhookEndpointUrl
	}
	if !IsNil(o.WebhookSecret) {
		toSerialize["webhookSecret"] = o.WebhookSecret
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailGateway struct {
	value *UnibeeApiBeanDetailGateway
	isSet bool
}

func (v NullableUnibeeApiBeanDetailGateway) Get() *UnibeeApiBeanDetailGateway {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailGateway) Set(val *UnibeeApiBeanDetailGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailGateway(val *UnibeeApiBeanDetailGateway) *NullableUnibeeApiBeanDetailGateway {
	return &NullableUnibeeApiBeanDetailGateway{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


