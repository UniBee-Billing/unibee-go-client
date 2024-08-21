/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanGateway{}

// UnibeeApiBeanGateway struct for UnibeeApiBeanGateway
type UnibeeApiBeanGateway struct {
	Bank *UnibeeApiBeanGatewayBank `json:"bank,omitempty"`
	CountryConfig *map[string]bool `json:"countryConfig,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// The currency of wire transfer 
	Currency *string `json:"currency,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	GatewayId *int64 `json:"gatewayId,omitempty"`
	GatewayKey *string `json:"gatewayKey,omitempty"`
	GatewayLogo *string `json:"gatewayLogo,omitempty"`
	GatewayName *string `json:"gatewayName,omitempty"`
	// gateway type，1-Bank Card ｜ 2-Crypto | 3 - Wire Transfer
	GatewayType *int64 `json:"gatewayType,omitempty"`
	// The minimum amount of wire transfer
	MinimumAmount *int64 `json:"minimumAmount,omitempty"`
	// The endpoint url of gateway webhook 
	WebhookEndpointUrl *string `json:"webhookEndpointUrl,omitempty"`
	// The secret of gateway webhook
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

// NewUnibeeApiBeanGateway instantiates a new UnibeeApiBeanGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanGateway() *UnibeeApiBeanGateway {
	this := UnibeeApiBeanGateway{}
	return &this
}

// NewUnibeeApiBeanGatewayWithDefaults instantiates a new UnibeeApiBeanGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanGatewayWithDefaults() *UnibeeApiBeanGateway {
	this := UnibeeApiBeanGateway{}
	return &this
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetBank() UnibeeApiBeanGatewayBank {
	if o == nil || IsNil(o.Bank) {
		var ret UnibeeApiBeanGatewayBank
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetBankOk() (*UnibeeApiBeanGatewayBank, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given UnibeeApiBeanGatewayBank and assigns it to the Bank field.
func (o *UnibeeApiBeanGateway) SetBank(v UnibeeApiBeanGatewayBank) {
	o.Bank = &v
}

// GetCountryConfig returns the CountryConfig field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetCountryConfig() map[string]bool {
	if o == nil || IsNil(o.CountryConfig) {
		var ret map[string]bool
		return ret
	}
	return *o.CountryConfig
}

// GetCountryConfigOk returns a tuple with the CountryConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetCountryConfigOk() (*map[string]bool, bool) {
	if o == nil || IsNil(o.CountryConfig) {
		return nil, false
	}
	return o.CountryConfig, true
}

// HasCountryConfig returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasCountryConfig() bool {
	if o != nil && !IsNil(o.CountryConfig) {
		return true
	}

	return false
}

// SetCountryConfig gets a reference to the given map[string]bool and assigns it to the CountryConfig field.
func (o *UnibeeApiBeanGateway) SetCountryConfig(v map[string]bool) {
	o.CountryConfig = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanGateway) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanGateway) SetCurrency(v string) {
	o.Currency = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UnibeeApiBeanGateway) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetGatewayId returns the GatewayId field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetGatewayId() int64 {
	if o == nil || IsNil(o.GatewayId) {
		var ret int64
		return ret
	}
	return *o.GatewayId
}

// GetGatewayIdOk returns a tuple with the GatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetGatewayIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayId) {
		return nil, false
	}
	return o.GatewayId, true
}

// HasGatewayId returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasGatewayId() bool {
	if o != nil && !IsNil(o.GatewayId) {
		return true
	}

	return false
}

// SetGatewayId gets a reference to the given int64 and assigns it to the GatewayId field.
func (o *UnibeeApiBeanGateway) SetGatewayId(v int64) {
	o.GatewayId = &v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *UnibeeApiBeanGateway) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetGatewayLogo returns the GatewayLogo field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetGatewayLogo() string {
	if o == nil || IsNil(o.GatewayLogo) {
		var ret string
		return ret
	}
	return *o.GatewayLogo
}

// GetGatewayLogoOk returns a tuple with the GatewayLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetGatewayLogoOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayLogo) {
		return nil, false
	}
	return o.GatewayLogo, true
}

// HasGatewayLogo returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasGatewayLogo() bool {
	if o != nil && !IsNil(o.GatewayLogo) {
		return true
	}

	return false
}

// SetGatewayLogo gets a reference to the given string and assigns it to the GatewayLogo field.
func (o *UnibeeApiBeanGateway) SetGatewayLogo(v string) {
	o.GatewayLogo = &v
}

// GetGatewayName returns the GatewayName field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetGatewayName() string {
	if o == nil || IsNil(o.GatewayName) {
		var ret string
		return ret
	}
	return *o.GatewayName
}

// GetGatewayNameOk returns a tuple with the GatewayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetGatewayNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayName) {
		return nil, false
	}
	return o.GatewayName, true
}

// HasGatewayName returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasGatewayName() bool {
	if o != nil && !IsNil(o.GatewayName) {
		return true
	}

	return false
}

// SetGatewayName gets a reference to the given string and assigns it to the GatewayName field.
func (o *UnibeeApiBeanGateway) SetGatewayName(v string) {
	o.GatewayName = &v
}

// GetGatewayType returns the GatewayType field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetGatewayType() int64 {
	if o == nil || IsNil(o.GatewayType) {
		var ret int64
		return ret
	}
	return *o.GatewayType
}

// GetGatewayTypeOk returns a tuple with the GatewayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetGatewayTypeOk() (*int64, bool) {
	if o == nil || IsNil(o.GatewayType) {
		return nil, false
	}
	return o.GatewayType, true
}

// HasGatewayType returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasGatewayType() bool {
	if o != nil && !IsNil(o.GatewayType) {
		return true
	}

	return false
}

// SetGatewayType gets a reference to the given int64 and assigns it to the GatewayType field.
func (o *UnibeeApiBeanGateway) SetGatewayType(v int64) {
	o.GatewayType = &v
}

// GetMinimumAmount returns the MinimumAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetMinimumAmount() int64 {
	if o == nil || IsNil(o.MinimumAmount) {
		var ret int64
		return ret
	}
	return *o.MinimumAmount
}

// GetMinimumAmountOk returns a tuple with the MinimumAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetMinimumAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.MinimumAmount) {
		return nil, false
	}
	return o.MinimumAmount, true
}

// HasMinimumAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasMinimumAmount() bool {
	if o != nil && !IsNil(o.MinimumAmount) {
		return true
	}

	return false
}

// SetMinimumAmount gets a reference to the given int64 and assigns it to the MinimumAmount field.
func (o *UnibeeApiBeanGateway) SetMinimumAmount(v int64) {
	o.MinimumAmount = &v
}

// GetWebhookEndpointUrl returns the WebhookEndpointUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetWebhookEndpointUrl() string {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		var ret string
		return ret
	}
	return *o.WebhookEndpointUrl
}

// GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetWebhookEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEndpointUrl) {
		return nil, false
	}
	return o.WebhookEndpointUrl, true
}

// HasWebhookEndpointUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasWebhookEndpointUrl() bool {
	if o != nil && !IsNil(o.WebhookEndpointUrl) {
		return true
	}

	return false
}

// SetWebhookEndpointUrl gets a reference to the given string and assigns it to the WebhookEndpointUrl field.
func (o *UnibeeApiBeanGateway) SetWebhookEndpointUrl(v string) {
	o.WebhookEndpointUrl = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise.
func (o *UnibeeApiBeanGateway) GetWebhookSecret() string {
	if o == nil || IsNil(o.WebhookSecret) {
		var ret string
		return ret
	}
	return *o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanGateway) GetWebhookSecretOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookSecret) {
		return nil, false
	}
	return o.WebhookSecret, true
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *UnibeeApiBeanGateway) HasWebhookSecret() bool {
	if o != nil && !IsNil(o.WebhookSecret) {
		return true
	}

	return false
}

// SetWebhookSecret gets a reference to the given string and assigns it to the WebhookSecret field.
func (o *UnibeeApiBeanGateway) SetWebhookSecret(v string) {
	o.WebhookSecret = &v
}

func (o UnibeeApiBeanGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
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
	if !IsNil(o.GatewayType) {
		toSerialize["gatewayType"] = o.GatewayType
	}
	if !IsNil(o.MinimumAmount) {
		toSerialize["minimumAmount"] = o.MinimumAmount
	}
	if !IsNil(o.WebhookEndpointUrl) {
		toSerialize["webhookEndpointUrl"] = o.WebhookEndpointUrl
	}
	if !IsNil(o.WebhookSecret) {
		toSerialize["webhookSecret"] = o.WebhookSecret
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanGateway struct {
	value *UnibeeApiBeanGateway
	isSet bool
}

func (v NullableUnibeeApiBeanGateway) Get() *UnibeeApiBeanGateway {
	return v.value
}

func (v *NullableUnibeeApiBeanGateway) Set(val *UnibeeApiBeanGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanGateway(val *UnibeeApiBeanGateway) *NullableUnibeeApiBeanGateway {
	return &NullableUnibeeApiBeanGateway{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


