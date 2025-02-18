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

// checks if the UnibeeApiBeanCreditConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanCreditConfig{}

// UnibeeApiBeanCreditConfig struct for UnibeeApiBeanCreditConfig
type UnibeeApiBeanCreditConfig struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	// discount code exclusive when purchase, default no, 0-no, 1-yes
	DiscountCodeExclusive *int32 `json:"discountCodeExclusive,omitempty"`
	// keep two decimal places，multiply by 100 saved, 1 currency = 1 credit * (exchange_rate/100), main account fixed rate to 100
	ExchangeRate *int64 `json:"exchangeRate,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// logo image base64, show when user purchase
	Logo *string `json:"logo,omitempty"`
	// logo url, show when user purchase
	LogoUrl *string `json:"logoUrl,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// meta_data(json)
	MetaData *map[string]interface{} `json:"metaData,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// 0-no, 1-yes
	PayoutEnable *int32 `json:"payoutEnable,omitempty"`
	// is default used when in purchase preview, default no, 0-no, 1-yes
	PreviewDefaultUsed *int32 `json:"previewDefaultUsed,omitempty"`
	// 0-no, 1-yes
	RechargeEnable *int32 `json:"rechargeEnable,omitempty"`
	// apply to recurring, default no, 0-no,1-yes
	Recurring *int32 `json:"recurring,omitempty"`
	// the total decrement amount
	TotalDecrementAmount *int64 `json:"totalDecrementAmount,omitempty"`
	// the total increment amount
	TotalIncrementAmount *int64 `json:"totalIncrementAmount,omitempty"`
	// type of credit account, 1-main account, 2-promo credit account
	Type *int32 `json:"type,omitempty"`
}

// NewUnibeeApiBeanCreditConfig instantiates a new UnibeeApiBeanCreditConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanCreditConfig() *UnibeeApiBeanCreditConfig {
	this := UnibeeApiBeanCreditConfig{}
	return &this
}

// NewUnibeeApiBeanCreditConfigWithDefaults instantiates a new UnibeeApiBeanCreditConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanCreditConfigWithDefaults() *UnibeeApiBeanCreditConfig {
	this := UnibeeApiBeanCreditConfig{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanCreditConfig) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanCreditConfig) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanCreditConfig) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountCodeExclusive returns the DiscountCodeExclusive field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetDiscountCodeExclusive() int32 {
	if o == nil || IsNil(o.DiscountCodeExclusive) {
		var ret int32
		return ret
	}
	return *o.DiscountCodeExclusive
}

// GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetDiscountCodeExclusiveOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountCodeExclusive) {
		return nil, false
	}
	return o.DiscountCodeExclusive, true
}

// HasDiscountCodeExclusive returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasDiscountCodeExclusive() bool {
	if o != nil && !IsNil(o.DiscountCodeExclusive) {
		return true
	}

	return false
}

// SetDiscountCodeExclusive gets a reference to the given int32 and assigns it to the DiscountCodeExclusive field.
func (o *UnibeeApiBeanCreditConfig) SetDiscountCodeExclusive(v int32) {
	o.DiscountCodeExclusive = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetExchangeRate() int64 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret int64
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetExchangeRateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given int64 and assigns it to the ExchangeRate field.
func (o *UnibeeApiBeanCreditConfig) SetExchangeRate(v int64) {
	o.ExchangeRate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanCreditConfig) SetId(v int64) {
	o.Id = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *UnibeeApiBeanCreditConfig) SetLogo(v string) {
	o.Logo = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *UnibeeApiBeanCreditConfig) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanCreditConfig) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetMetaData() map[string]interface{} {
	if o == nil || IsNil(o.MetaData) {
		var ret map[string]interface{}
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetMetaDataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given map[string]interface{} and assigns it to the MetaData field.
func (o *UnibeeApiBeanCreditConfig) SetMetaData(v map[string]interface{}) {
	o.MetaData = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanCreditConfig) SetName(v string) {
	o.Name = &v
}

// GetPayoutEnable returns the PayoutEnable field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetPayoutEnable() int32 {
	if o == nil || IsNil(o.PayoutEnable) {
		var ret int32
		return ret
	}
	return *o.PayoutEnable
}

// GetPayoutEnableOk returns a tuple with the PayoutEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetPayoutEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.PayoutEnable) {
		return nil, false
	}
	return o.PayoutEnable, true
}

// HasPayoutEnable returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasPayoutEnable() bool {
	if o != nil && !IsNil(o.PayoutEnable) {
		return true
	}

	return false
}

// SetPayoutEnable gets a reference to the given int32 and assigns it to the PayoutEnable field.
func (o *UnibeeApiBeanCreditConfig) SetPayoutEnable(v int32) {
	o.PayoutEnable = &v
}

// GetPreviewDefaultUsed returns the PreviewDefaultUsed field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetPreviewDefaultUsed() int32 {
	if o == nil || IsNil(o.PreviewDefaultUsed) {
		var ret int32
		return ret
	}
	return *o.PreviewDefaultUsed
}

// GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetPreviewDefaultUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.PreviewDefaultUsed) {
		return nil, false
	}
	return o.PreviewDefaultUsed, true
}

// HasPreviewDefaultUsed returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasPreviewDefaultUsed() bool {
	if o != nil && !IsNil(o.PreviewDefaultUsed) {
		return true
	}

	return false
}

// SetPreviewDefaultUsed gets a reference to the given int32 and assigns it to the PreviewDefaultUsed field.
func (o *UnibeeApiBeanCreditConfig) SetPreviewDefaultUsed(v int32) {
	o.PreviewDefaultUsed = &v
}

// GetRechargeEnable returns the RechargeEnable field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetRechargeEnable() int32 {
	if o == nil || IsNil(o.RechargeEnable) {
		var ret int32
		return ret
	}
	return *o.RechargeEnable
}

// GetRechargeEnableOk returns a tuple with the RechargeEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetRechargeEnableOk() (*int32, bool) {
	if o == nil || IsNil(o.RechargeEnable) {
		return nil, false
	}
	return o.RechargeEnable, true
}

// HasRechargeEnable returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasRechargeEnable() bool {
	if o != nil && !IsNil(o.RechargeEnable) {
		return true
	}

	return false
}

// SetRechargeEnable gets a reference to the given int32 and assigns it to the RechargeEnable field.
func (o *UnibeeApiBeanCreditConfig) SetRechargeEnable(v int32) {
	o.RechargeEnable = &v
}

// GetRecurring returns the Recurring field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetRecurring() int32 {
	if o == nil || IsNil(o.Recurring) {
		var ret int32
		return ret
	}
	return *o.Recurring
}

// GetRecurringOk returns a tuple with the Recurring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetRecurringOk() (*int32, bool) {
	if o == nil || IsNil(o.Recurring) {
		return nil, false
	}
	return o.Recurring, true
}

// HasRecurring returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasRecurring() bool {
	if o != nil && !IsNil(o.Recurring) {
		return true
	}

	return false
}

// SetRecurring gets a reference to the given int32 and assigns it to the Recurring field.
func (o *UnibeeApiBeanCreditConfig) SetRecurring(v int32) {
	o.Recurring = &v
}

// GetTotalDecrementAmount returns the TotalDecrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetTotalDecrementAmount() int64 {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalDecrementAmount
}

// GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetTotalDecrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDecrementAmount) {
		return nil, false
	}
	return o.TotalDecrementAmount, true
}

// HasTotalDecrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasTotalDecrementAmount() bool {
	if o != nil && !IsNil(o.TotalDecrementAmount) {
		return true
	}

	return false
}

// SetTotalDecrementAmount gets a reference to the given int64 and assigns it to the TotalDecrementAmount field.
func (o *UnibeeApiBeanCreditConfig) SetTotalDecrementAmount(v int64) {
	o.TotalDecrementAmount = &v
}

// GetTotalIncrementAmount returns the TotalIncrementAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetTotalIncrementAmount() int64 {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		var ret int64
		return ret
	}
	return *o.TotalIncrementAmount
}

// GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetTotalIncrementAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalIncrementAmount) {
		return nil, false
	}
	return o.TotalIncrementAmount, true
}

// HasTotalIncrementAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasTotalIncrementAmount() bool {
	if o != nil && !IsNil(o.TotalIncrementAmount) {
		return true
	}

	return false
}

// SetTotalIncrementAmount gets a reference to the given int64 and assigns it to the TotalIncrementAmount field.
func (o *UnibeeApiBeanCreditConfig) SetTotalIncrementAmount(v int64) {
	o.TotalIncrementAmount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanCreditConfig) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanCreditConfig) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanCreditConfig) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanCreditConfig) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeApiBeanCreditConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanCreditConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DiscountCodeExclusive) {
		toSerialize["discountCodeExclusive"] = o.DiscountCodeExclusive
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
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
	if !IsNil(o.TotalDecrementAmount) {
		toSerialize["totalDecrementAmount"] = o.TotalDecrementAmount
	}
	if !IsNil(o.TotalIncrementAmount) {
		toSerialize["totalIncrementAmount"] = o.TotalIncrementAmount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanCreditConfig struct {
	value *UnibeeApiBeanCreditConfig
	isSet bool
}

func (v NullableUnibeeApiBeanCreditConfig) Get() *UnibeeApiBeanCreditConfig {
	return v.value
}

func (v *NullableUnibeeApiBeanCreditConfig) Set(val *UnibeeApiBeanCreditConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanCreditConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanCreditConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanCreditConfig(val *UnibeeApiBeanCreditConfig) *NullableUnibeeApiBeanCreditConfig {
	return &NullableUnibeeApiBeanCreditConfig{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanCreditConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanCreditConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


