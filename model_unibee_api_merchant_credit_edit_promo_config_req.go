/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantCreditEditPromoConfigReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantCreditEditPromoConfigReq{}

// UnibeeApiMerchantCreditEditPromoConfigReq struct for UnibeeApiMerchantCreditEditPromoConfigReq
type UnibeeApiMerchantCreditEditPromoConfigReq struct {
	// currency
	Currency string `json:"currency"`
	// description
	Description *string `json:"description,omitempty"`
	// discount code exclusive when purchase, default no, 0-no, 1-yes
	DiscountCodeExclusive *int32 `json:"discountCodeExclusive,omitempty"`
	// keep two decimal places，scale = 100, 1 currency = 1 credit * (exchange_rate/100), main account fixed rate to 100
	ExchangeRate *int32 `json:"exchangeRate,omitempty"`
	// logo image base64, show when user purchase
	Logo *string `json:"logo,omitempty"`
	// logo url, show when user purchase
	LogoUrl *string `json:"logoUrl,omitempty"`
	// meta_data(json)
	MetaData *map[string]interface{} `json:"metaData,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// credit account can payout or not, default no, 0-no, 1-yes
	PayoutEnable *int32 `json:"payoutEnable,omitempty"`
	// is default used when in purchase preview, default no, 0-no, 1-yes
	PreviewDefaultUsed *int32 `json:"previewDefaultUsed,omitempty"`
	// credit account can be recharged or not, 0-no, 1-yes
	RechargeEnable *int32 `json:"rechargeEnable,omitempty"`
	// apply to recurring, default no, 0-no,1-yes
	Recurring *int32 `json:"recurring,omitempty"`
}

type _UnibeeApiMerchantCreditEditPromoConfigReq UnibeeApiMerchantCreditEditPromoConfigReq

// NewUnibeeApiMerchantCreditEditPromoConfigReq instantiates a new UnibeeApiMerchantCreditEditPromoConfigReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantCreditEditPromoConfigReq(currency string) *UnibeeApiMerchantCreditEditPromoConfigReq {
	this := UnibeeApiMerchantCreditEditPromoConfigReq{}
	this.Currency = currency
	return &this
}

// NewUnibeeApiMerchantCreditEditPromoConfigReqWithDefaults instantiates a new UnibeeApiMerchantCreditEditPromoConfigReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantCreditEditPromoConfigReqWithDefaults() *UnibeeApiMerchantCreditEditPromoConfigReq {
	this := UnibeeApiMerchantCreditEditPromoConfigReq{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountCodeExclusive returns the DiscountCodeExclusive field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDiscountCodeExclusive() int32 {
	if o == nil || IsNil(o.DiscountCodeExclusive) {
		var ret int32
		return ret
	}
	return *o.DiscountCodeExclusive
}

// GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDiscountCodeExclusiveOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountCodeExclusive) {
		return nil, false
	}
	return o.DiscountCodeExclusive, true
}

// HasDiscountCodeExclusive returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasDiscountCodeExclusive() bool {
	if o != nil && !IsNil(o.DiscountCodeExclusive) {
		return true
	}

	return false
}

// SetDiscountCodeExclusive gets a reference to the given int32 and assigns it to the DiscountCodeExclusive field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetDiscountCodeExclusive(v int32) {
	o.DiscountCodeExclusive = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetExchangeRate() int32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret int32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetExchangeRateOk() (*int32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given int32 and assigns it to the ExchangeRate field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetExchangeRate(v int32) {
	o.ExchangeRate = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetMetaDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetMetaData(v map[string]interface{}) {
	o.MetaData = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetName(v string) {
	o.Name = &v
}

// GetPayoutEnable returns the PayoutEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPayoutEnable() int32 {
	if o == nil || IsNil(o.PayoutEnable) {
		var ret int32
		return ret
	}
	return *o.PayoutEnable
}

// GetPayoutEnableOk returns a tuple with the PayoutEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPayoutEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.PayoutEnable) {
		return nil, false
	}
	return o.PayoutEnable, true
}

// HasPayoutEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasPayoutEnable() bool {
	if o != nil && !IsNil(o.PayoutEnable) {
		return true
	}

	return false
}

// SetPayoutEnable gets a reference to the given int32 and assigns it to the PayoutEnable field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetPayoutEnable(v int32) {
	o.PayoutEnable = &v
}

// GetPreviewDefaultUsed returns the PreviewDefaultUsed field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPreviewDefaultUsed() int32 {
	if o == nil || IsNil(o.PreviewDefaultUsed) {
		var ret int32
		return ret
	}
	return *o.PreviewDefaultUsed
}

// GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPreviewDefaultUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.PreviewDefaultUsed) {
		return nil, false
	}
	return o.PreviewDefaultUsed, true
}

// HasPreviewDefaultUsed returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasPreviewDefaultUsed() bool {
	if o != nil && !IsNil(o.PreviewDefaultUsed) {
		return true
	}

	return false
}

// SetPreviewDefaultUsed gets a reference to the given int32 and assigns it to the PreviewDefaultUsed field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetPreviewDefaultUsed(v int32) {
	o.PreviewDefaultUsed = &v
}

// GetRechargeEnable returns the RechargeEnable field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRechargeEnable() int32 {
	if o == nil || IsNil(o.RechargeEnable) {
		var ret int32
		return ret
	}
	return *o.RechargeEnable
}

// GetRechargeEnableOk returns a tuple with the RechargeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRechargeEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.RechargeEnable) {
		return nil, false
	}
	return o.RechargeEnable, true
}

// HasRechargeEnable returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasRechargeEnable() bool {
	if o != nil && !IsNil(o.RechargeEnable) {
		return true
	}

	return false
}

// SetRechargeEnable gets a reference to the given int32 and assigns it to the RechargeEnable field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetRechargeEnable(v int32) {
	o.RechargeEnable = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRecurring() int32 {
	if o == nil || IsNil(o.Recurring) {
		var ret int32
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRecurringOk() (*int32, bool) {
	if o == nil || IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasRecurring() bool {
	if o != nil && !IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given int32 and assigns it to the Recurring field.
func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetRecurring(v int32) {
	o.Recurring = &v
}

func (o UnibeeApiMerchantCreditEditPromoConfigReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantCreditEditPromoConfigReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DiscountCodeExclusive) {
		toSerialize["discountCodeExclusive"] = o.DiscountCodeExclusive
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PayoutEnable) {
		toSerialize["payoutEnable"] = o.PayoutEnable
	}
	if !IsNil(o.PreviewDefaultUsed) {
		toSerialize["previewDefaultUsed"] = o.PreviewDefaultUsed
	}
	if !IsNil(o.RechargeEnable) {
		toSerialize["rechargeEnable"] = o.RechargeEnable
	}
	if !IsNil(o.Recurring) {
		toSerialize["recurring"] = o.Recurring
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantCreditEditPromoConfigReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currency",
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

	varUnibeeApiMerchantCreditEditPromoConfigReq := _UnibeeApiMerchantCreditEditPromoConfigReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantCreditEditPromoConfigReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantCreditEditPromoConfigReq(varUnibeeApiMerchantCreditEditPromoConfigReq)

	return err
}

type NullableUnibeeApiMerchantCreditEditPromoConfigReq struct {
	value *UnibeeApiMerchantCreditEditPromoConfigReq
	isSet bool
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigReq) Get() *UnibeeApiMerchantCreditEditPromoConfigReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigReq) Set(val *UnibeeApiMerchantCreditEditPromoConfigReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantCreditEditPromoConfigReq(val *UnibeeApiMerchantCreditEditPromoConfigReq) *NullableUnibeeApiMerchantCreditEditPromoConfigReq {
	return &NullableUnibeeApiMerchantCreditEditPromoConfigReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantCreditEditPromoConfigReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantCreditEditPromoConfigReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


