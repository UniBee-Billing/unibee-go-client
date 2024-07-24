/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanSubscriptionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanSubscriptionConfig{}

// UnibeeApiBeanSubscriptionConfig struct for UnibeeApiBeanSubscriptionConfig
type UnibeeApiBeanSubscriptionConfig struct {
	// DowngradeEffectImmediately, whether subscription update should effect immediately or at period end, default at period end
	DowngradeEffectImmediately *bool `json:"downgradeEffectImmediately,omitempty"`
	GatewayVATRule *string `json:"gatewayVATRule,omitempty"`
	// IncompleteExpireTime, em.. default 1day for plan of month type
	IncompleteExpireTime *int64 `json:"incompleteExpireTime,omitempty"`
	// InvoiceEmail, whether to send invoice email to user, default yes
	InvoiceEmail *bool `json:"invoiceEmail,omitempty"`
	// InvoicePdfGenerate, whether to generate invoice pdf to user, default yes
	InvoicePdfGenerate *bool `json:"invoicePdfGenerate,omitempty"`
	// TryAutomaticPaymentBeforePeriodEnd, default 30 min
	TryAutomaticPaymentBeforePeriodEnd *int64 `json:"tryAutomaticPaymentBeforePeriodEnd,omitempty"`
	// UpgradeProration, whether subscription update generation proration invoice or not, default yes
	UpgradeProration *bool `json:"upgradeProration,omitempty"`
}

// NewUnibeeApiBeanSubscriptionConfig instantiates a new UnibeeApiBeanSubscriptionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanSubscriptionConfig() *UnibeeApiBeanSubscriptionConfig {
	this := UnibeeApiBeanSubscriptionConfig{}
	return &this
}

// NewUnibeeApiBeanSubscriptionConfigWithDefaults instantiates a new UnibeeApiBeanSubscriptionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanSubscriptionConfigWithDefaults() *UnibeeApiBeanSubscriptionConfig {
	this := UnibeeApiBeanSubscriptionConfig{}
	return &this
}

// GetDowngradeEffectImmediately returns the DowngradeEffectImmediately field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeEffectImmediately() bool {
	if o == nil || IsNil(o.DowngradeEffectImmediately) {
		var ret bool
		return ret
	}
	return *o.DowngradeEffectImmediately
}

// GetDowngradeEffectImmediatelyOk returns a tuple with the DowngradeEffectImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetDowngradeEffectImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.DowngradeEffectImmediately) {
		return nil, false
	}
	return o.DowngradeEffectImmediately, true
}

// HasDowngradeEffectImmediately returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasDowngradeEffectImmediately() bool {
	if o != nil && !IsNil(o.DowngradeEffectImmediately) {
		return true
	}

	return false
}

// SetDowngradeEffectImmediately gets a reference to the given bool and assigns it to the DowngradeEffectImmediately field.
func (o *UnibeeApiBeanSubscriptionConfig) SetDowngradeEffectImmediately(v bool) {
	o.DowngradeEffectImmediately = &v
}

// GetGatewayVATRule returns the GatewayVATRule field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetGatewayVATRule() string {
	if o == nil || IsNil(o.GatewayVATRule) {
		var ret string
		return ret
	}
	return *o.GatewayVATRule
}

// GetGatewayVATRuleOk returns a tuple with the GatewayVATRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetGatewayVATRuleOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayVATRule) {
		return nil, false
	}
	return o.GatewayVATRule, true
}

// HasGatewayVATRule returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasGatewayVATRule() bool {
	if o != nil && !IsNil(o.GatewayVATRule) {
		return true
	}

	return false
}

// SetGatewayVATRule gets a reference to the given string and assigns it to the GatewayVATRule field.
func (o *UnibeeApiBeanSubscriptionConfig) SetGatewayVATRule(v string) {
	o.GatewayVATRule = &v
}

// GetIncompleteExpireTime returns the IncompleteExpireTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetIncompleteExpireTime() int64 {
	if o == nil || IsNil(o.IncompleteExpireTime) {
		var ret int64
		return ret
	}
	return *o.IncompleteExpireTime
}

// GetIncompleteExpireTimeOk returns a tuple with the IncompleteExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetIncompleteExpireTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.IncompleteExpireTime) {
		return nil, false
	}
	return o.IncompleteExpireTime, true
}

// HasIncompleteExpireTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasIncompleteExpireTime() bool {
	if o != nil && !IsNil(o.IncompleteExpireTime) {
		return true
	}

	return false
}

// SetIncompleteExpireTime gets a reference to the given int64 and assigns it to the IncompleteExpireTime field.
func (o *UnibeeApiBeanSubscriptionConfig) SetIncompleteExpireTime(v int64) {
	o.IncompleteExpireTime = &v
}

// GetInvoiceEmail returns the InvoiceEmail field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetInvoiceEmail() bool {
	if o == nil || IsNil(o.InvoiceEmail) {
		var ret bool
		return ret
	}
	return *o.InvoiceEmail
}

// GetInvoiceEmailOk returns a tuple with the InvoiceEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetInvoiceEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceEmail) {
		return nil, false
	}
	return o.InvoiceEmail, true
}

// HasInvoiceEmail returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasInvoiceEmail() bool {
	if o != nil && !IsNil(o.InvoiceEmail) {
		return true
	}

	return false
}

// SetInvoiceEmail gets a reference to the given bool and assigns it to the InvoiceEmail field.
func (o *UnibeeApiBeanSubscriptionConfig) SetInvoiceEmail(v bool) {
	o.InvoiceEmail = &v
}

// GetInvoicePdfGenerate returns the InvoicePdfGenerate field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetInvoicePdfGenerate() bool {
	if o == nil || IsNil(o.InvoicePdfGenerate) {
		var ret bool
		return ret
	}
	return *o.InvoicePdfGenerate
}

// GetInvoicePdfGenerateOk returns a tuple with the InvoicePdfGenerate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetInvoicePdfGenerateOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoicePdfGenerate) {
		return nil, false
	}
	return o.InvoicePdfGenerate, true
}

// HasInvoicePdfGenerate returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasInvoicePdfGenerate() bool {
	if o != nil && !IsNil(o.InvoicePdfGenerate) {
		return true
	}

	return false
}

// SetInvoicePdfGenerate gets a reference to the given bool and assigns it to the InvoicePdfGenerate field.
func (o *UnibeeApiBeanSubscriptionConfig) SetInvoicePdfGenerate(v bool) {
	o.InvoicePdfGenerate = &v
}

// GetTryAutomaticPaymentBeforePeriodEnd returns the TryAutomaticPaymentBeforePeriodEnd field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetTryAutomaticPaymentBeforePeriodEnd() int64 {
	if o == nil || IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		var ret int64
		return ret
	}
	return *o.TryAutomaticPaymentBeforePeriodEnd
}

// GetTryAutomaticPaymentBeforePeriodEndOk returns a tuple with the TryAutomaticPaymentBeforePeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetTryAutomaticPaymentBeforePeriodEndOk() (*int64, bool) {
	if o == nil || IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		return nil, false
	}
	return o.TryAutomaticPaymentBeforePeriodEnd, true
}

// HasTryAutomaticPaymentBeforePeriodEnd returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasTryAutomaticPaymentBeforePeriodEnd() bool {
	if o != nil && !IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		return true
	}

	return false
}

// SetTryAutomaticPaymentBeforePeriodEnd gets a reference to the given int64 and assigns it to the TryAutomaticPaymentBeforePeriodEnd field.
func (o *UnibeeApiBeanSubscriptionConfig) SetTryAutomaticPaymentBeforePeriodEnd(v int64) {
	o.TryAutomaticPaymentBeforePeriodEnd = &v
}

// GetUpgradeProration returns the UpgradeProration field value if set, zero value otherwise.
func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeProration() bool {
	if o == nil || IsNil(o.UpgradeProration) {
		var ret bool
		return ret
	}
	return *o.UpgradeProration
}

// GetUpgradeProrationOk returns a tuple with the UpgradeProration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanSubscriptionConfig) GetUpgradeProrationOk() (*bool, bool) {
	if o == nil || IsNil(o.UpgradeProration) {
		return nil, false
	}
	return o.UpgradeProration, true
}

// HasUpgradeProration returns a boolean if a field has been set.
func (o *UnibeeApiBeanSubscriptionConfig) HasUpgradeProration() bool {
	if o != nil && !IsNil(o.UpgradeProration) {
		return true
	}

	return false
}

// SetUpgradeProration gets a reference to the given bool and assigns it to the UpgradeProration field.
func (o *UnibeeApiBeanSubscriptionConfig) SetUpgradeProration(v bool) {
	o.UpgradeProration = &v
}

func (o UnibeeApiBeanSubscriptionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanSubscriptionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DowngradeEffectImmediately) {
		toSerialize["downgradeEffectImmediately"] = o.DowngradeEffectImmediately
	}
	if !IsNil(o.GatewayVATRule) {
		toSerialize["gatewayVATRule"] = o.GatewayVATRule
	}
	if !IsNil(o.IncompleteExpireTime) {
		toSerialize["incompleteExpireTime"] = o.IncompleteExpireTime
	}
	if !IsNil(o.InvoiceEmail) {
		toSerialize["invoiceEmail"] = o.InvoiceEmail
	}
	if !IsNil(o.InvoicePdfGenerate) {
		toSerialize["invoicePdfGenerate"] = o.InvoicePdfGenerate
	}
	if !IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		toSerialize["tryAutomaticPaymentBeforePeriodEnd"] = o.TryAutomaticPaymentBeforePeriodEnd
	}
	if !IsNil(o.UpgradeProration) {
		toSerialize["upgradeProration"] = o.UpgradeProration
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanSubscriptionConfig struct {
	value *UnibeeApiBeanSubscriptionConfig
	isSet bool
}

func (v NullableUnibeeApiBeanSubscriptionConfig) Get() *UnibeeApiBeanSubscriptionConfig {
	return v.value
}

func (v *NullableUnibeeApiBeanSubscriptionConfig) Set(val *UnibeeApiBeanSubscriptionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanSubscriptionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanSubscriptionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanSubscriptionConfig(val *UnibeeApiBeanSubscriptionConfig) *NullableUnibeeApiBeanSubscriptionConfig {
	return &NullableUnibeeApiBeanSubscriptionConfig{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanSubscriptionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanSubscriptionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


