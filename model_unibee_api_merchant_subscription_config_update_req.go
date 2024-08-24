/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantSubscriptionConfigUpdateReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionConfigUpdateReq{}

// UnibeeApiMerchantSubscriptionConfigUpdateReq struct for UnibeeApiMerchantSubscriptionConfigUpdateReq
type UnibeeApiMerchantSubscriptionConfigUpdateReq struct {
	// DowngradeEffectImmediately, whether subscription downgrade should effect immediately or at period end, default at period end
	DowngradeEffectImmediately *bool `json:"downgradeEffectImmediately,omitempty"`
	GatewayVATRule [][]UnibeeApiBeanMerchantVatRule `json:"gatewayVATRule,omitempty"`
	// IncompleteExpireTime, em.. default 1day for plan of month type
	IncompleteExpireTime *int32 `json:"incompleteExpireTime,omitempty"`
	// InvoiceEmail, whether to send invoice email to user, default yes
	InvoiceEmail *bool `json:"invoiceEmail,omitempty"`
	// TryAutomaticPaymentBeforePeriodEnd, default 30 min
	TryAutomaticPaymentBeforePeriodEnd *int32 `json:"tryAutomaticPaymentBeforePeriodEnd,omitempty"`
	// UpgradeProration, whether subscription update generation proration invoice or not, default yes
	UpgradeProration *bool `json:"upgradeProration,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionConfigUpdateReq instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionConfigUpdateReq() *UnibeeApiMerchantSubscriptionConfigUpdateReq {
	this := UnibeeApiMerchantSubscriptionConfigUpdateReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionConfigUpdateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionConfigUpdateReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionConfigUpdateReqWithDefaults() *UnibeeApiMerchantSubscriptionConfigUpdateReq {
	this := UnibeeApiMerchantSubscriptionConfigUpdateReq{}
	return &this
}

// GetDowngradeEffectImmediately returns the DowngradeEffectImmediately field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetDowngradeEffectImmediately() bool {
	if o == nil || IsNil(o.DowngradeEffectImmediately) {
		var ret bool
		return ret
	}
	return *o.DowngradeEffectImmediately
}

// GetDowngradeEffectImmediatelyOk returns a tuple with the DowngradeEffectImmediately field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetDowngradeEffectImmediatelyOk() (*bool, bool) {
	if o == nil || IsNil(o.DowngradeEffectImmediately) {
		return nil, false
	}
	return o.DowngradeEffectImmediately, true
}

// HasDowngradeEffectImmediately returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasDowngradeEffectImmediately() bool {
	if o != nil && !IsNil(o.DowngradeEffectImmediately) {
		return true
	}

	return false
}

// SetDowngradeEffectImmediately gets a reference to the given bool and assigns it to the DowngradeEffectImmediately field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetDowngradeEffectImmediately(v bool) {
	o.DowngradeEffectImmediately = &v
}

// GetGatewayVATRule returns the GatewayVATRule field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetGatewayVATRule() [][]UnibeeApiBeanMerchantVatRule {
	if o == nil || IsNil(o.GatewayVATRule) {
		var ret [][]UnibeeApiBeanMerchantVatRule
		return ret
	}
	return o.GatewayVATRule
}

// GetGatewayVATRuleOk returns a tuple with the GatewayVATRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetGatewayVATRuleOk() ([][]UnibeeApiBeanMerchantVatRule, bool) {
	if o == nil || IsNil(o.GatewayVATRule) {
		return nil, false
	}
	return o.GatewayVATRule, true
}

// HasGatewayVATRule returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasGatewayVATRule() bool {
	if o != nil && !IsNil(o.GatewayVATRule) {
		return true
	}

	return false
}

// SetGatewayVATRule gets a reference to the given [][]UnibeeApiBeanMerchantVatRule and assigns it to the GatewayVATRule field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetGatewayVATRule(v [][]UnibeeApiBeanMerchantVatRule) {
	o.GatewayVATRule = v
}

// GetIncompleteExpireTime returns the IncompleteExpireTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetIncompleteExpireTime() int32 {
	if o == nil || IsNil(o.IncompleteExpireTime) {
		var ret int32
		return ret
	}
	return *o.IncompleteExpireTime
}

// GetIncompleteExpireTimeOk returns a tuple with the IncompleteExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetIncompleteExpireTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.IncompleteExpireTime) {
		return nil, false
	}
	return o.IncompleteExpireTime, true
}

// HasIncompleteExpireTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasIncompleteExpireTime() bool {
	if o != nil && !IsNil(o.IncompleteExpireTime) {
		return true
	}

	return false
}

// SetIncompleteExpireTime gets a reference to the given int32 and assigns it to the IncompleteExpireTime field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetIncompleteExpireTime(v int32) {
	o.IncompleteExpireTime = &v
}

// GetInvoiceEmail returns the InvoiceEmail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetInvoiceEmail() bool {
	if o == nil || IsNil(o.InvoiceEmail) {
		var ret bool
		return ret
	}
	return *o.InvoiceEmail
}

// GetInvoiceEmailOk returns a tuple with the InvoiceEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetInvoiceEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceEmail) {
		return nil, false
	}
	return o.InvoiceEmail, true
}

// HasInvoiceEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasInvoiceEmail() bool {
	if o != nil && !IsNil(o.InvoiceEmail) {
		return true
	}

	return false
}

// SetInvoiceEmail gets a reference to the given bool and assigns it to the InvoiceEmail field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetInvoiceEmail(v bool) {
	o.InvoiceEmail = &v
}

// GetTryAutomaticPaymentBeforePeriodEnd returns the TryAutomaticPaymentBeforePeriodEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetTryAutomaticPaymentBeforePeriodEnd() int32 {
	if o == nil || IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		var ret int32
		return ret
	}
	return *o.TryAutomaticPaymentBeforePeriodEnd
}

// GetTryAutomaticPaymentBeforePeriodEndOk returns a tuple with the TryAutomaticPaymentBeforePeriodEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetTryAutomaticPaymentBeforePeriodEndOk() (*int32, bool) {
	if o == nil || IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		return nil, false
	}
	return o.TryAutomaticPaymentBeforePeriodEnd, true
}

// HasTryAutomaticPaymentBeforePeriodEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasTryAutomaticPaymentBeforePeriodEnd() bool {
	if o != nil && !IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		return true
	}

	return false
}

// SetTryAutomaticPaymentBeforePeriodEnd gets a reference to the given int32 and assigns it to the TryAutomaticPaymentBeforePeriodEnd field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetTryAutomaticPaymentBeforePeriodEnd(v int32) {
	o.TryAutomaticPaymentBeforePeriodEnd = &v
}

// GetUpgradeProration returns the UpgradeProration field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetUpgradeProration() bool {
	if o == nil || IsNil(o.UpgradeProration) {
		var ret bool
		return ret
	}
	return *o.UpgradeProration
}

// GetUpgradeProrationOk returns a tuple with the UpgradeProration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) GetUpgradeProrationOk() (*bool, bool) {
	if o == nil || IsNil(o.UpgradeProration) {
		return nil, false
	}
	return o.UpgradeProration, true
}

// HasUpgradeProration returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) HasUpgradeProration() bool {
	if o != nil && !IsNil(o.UpgradeProration) {
		return true
	}

	return false
}

// SetUpgradeProration gets a reference to the given bool and assigns it to the UpgradeProration field.
func (o *UnibeeApiMerchantSubscriptionConfigUpdateReq) SetUpgradeProration(v bool) {
	o.UpgradeProration = &v
}

func (o UnibeeApiMerchantSubscriptionConfigUpdateReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionConfigUpdateReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TryAutomaticPaymentBeforePeriodEnd) {
		toSerialize["tryAutomaticPaymentBeforePeriodEnd"] = o.TryAutomaticPaymentBeforePeriodEnd
	}
	if !IsNil(o.UpgradeProration) {
		toSerialize["upgradeProration"] = o.UpgradeProration
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionConfigUpdateReq struct {
	value *UnibeeApiMerchantSubscriptionConfigUpdateReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) Get() *UnibeeApiMerchantSubscriptionConfigUpdateReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) Set(val *UnibeeApiMerchantSubscriptionConfigUpdateReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionConfigUpdateReq(val *UnibeeApiMerchantSubscriptionConfigUpdateReq) *NullableUnibeeApiMerchantSubscriptionConfigUpdateReq {
	return &NullableUnibeeApiMerchantSubscriptionConfigUpdateReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionConfigUpdateReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


