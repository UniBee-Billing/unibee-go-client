/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407240509 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanNewReq{}

// UnibeeApiMerchantPlanNewReq struct for UnibeeApiMerchantPlanNewReq
type UnibeeApiMerchantPlanNewReq struct {
	// Plan Ids Of Recurring Addon Type
	AddonIds []int64 `json:"addonIds,omitempty"`
	// Plan CaptureAmount
	Amount int64 `json:"amount"`
	// whether cancel at subscription first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription
	CancelAtTrialEnd *int32 `json:"cancelAtTrialEnd,omitempty"`
	// Plan Currency
	Currency string `json:"currency"`
	// Description
	Description *string `json:"description,omitempty"`
	// ExternalPlanId
	ExternalPlanId *string `json:"externalPlanId,omitempty"`
	// who pay the gas for crypto payment, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	// HomeUrl,Start With: http
	HomeUrl *string `json:"homeUrl,omitempty"`
	// ImageUrl,Start With: http
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Number Of IntervalUnit，em: day|month|year|week
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// Plan Interval Unit，em: day|month|year|week
	IntervalUnit *string `json:"intervalUnit,omitempty"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Plan's MetricLimit List
	MetricLimits []UnibeeApiBeanBulkMetricLimitPlanBindingParam `json:"metricLimits,omitempty"`
	// Plan Ids Of Onetime Addon Type
	OnetimeAddonIds []int64 `json:"onetimeAddonIds,omitempty"`
	// Plan Name
	PlanName string `json:"planName"`
	// Default Copy Description
	ProductDescription *string `json:"productDescription,omitempty"`
	// Default Copy PlanName
	ProductName *string `json:"productName,omitempty"`
	// price of trial period， not available for addon
	TrialAmount *int64 `json:"trialAmount,omitempty"`
	// demand of trial， not available for addon, example, paymentMethod, payment method will ask for subscription trial start
	TrialDemand *string `json:"trialDemand,omitempty"`
	// duration of trial， not available for addon
	TrialDurationTime *int64 `json:"trialDurationTime,omitempty"`
	// The type of plan, 1-main plan，2-addon plan, 3-onetime plan, default main plan
	Type *int32 `json:"type,omitempty"`
}

type _UnibeeApiMerchantPlanNewReq UnibeeApiMerchantPlanNewReq

// NewUnibeeApiMerchantPlanNewReq instantiates a new UnibeeApiMerchantPlanNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanNewReq(amount int64, currency string, planName string) *UnibeeApiMerchantPlanNewReq {
	this := UnibeeApiMerchantPlanNewReq{}
	this.Amount = amount
	this.Currency = currency
	this.PlanName = planName
	var type_ int32 = 1
	this.Type = &type_
	return &this
}

// NewUnibeeApiMerchantPlanNewReqWithDefaults instantiates a new UnibeeApiMerchantPlanNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanNewReqWithDefaults() *UnibeeApiMerchantPlanNewReq {
	this := UnibeeApiMerchantPlanNewReq{}
	var type_ int32 = 1
	this.Type = &type_
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeApiMerchantPlanNewReq) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAmount returns the Amount field value
func (o *UnibeeApiMerchantPlanNewReq) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UnibeeApiMerchantPlanNewReq) SetAmount(v int64) {
	o.Amount = v
}

// GetCancelAtTrialEnd returns the CancelAtTrialEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetCancelAtTrialEnd() int32 {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		var ret int32
		return ret
	}
	return *o.CancelAtTrialEnd
}

// GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetCancelAtTrialEndOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		return nil, false
	}
	return o.CancelAtTrialEnd, true
}

// HasCancelAtTrialEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasCancelAtTrialEnd() bool {
	if o != nil && !IsNil(o.CancelAtTrialEnd) {
		return true
	}

	return false
}

// SetCancelAtTrialEnd gets a reference to the given int32 and assigns it to the CancelAtTrialEnd field.
func (o *UnibeeApiMerchantPlanNewReq) SetCancelAtTrialEnd(v int32) {
	o.CancelAtTrialEnd = &v
}

// GetCurrency returns the Currency field value
func (o *UnibeeApiMerchantPlanNewReq) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UnibeeApiMerchantPlanNewReq) SetCurrency(v string) {
	o.Currency = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantPlanNewReq) SetDescription(v string) {
	o.Description = &v
}

// GetExternalPlanId returns the ExternalPlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetExternalPlanId() string {
	if o == nil || IsNil(o.ExternalPlanId) {
		var ret string
		return ret
	}
	return *o.ExternalPlanId
}

// GetExternalPlanIdOk returns a tuple with the ExternalPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetExternalPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPlanId) {
		return nil, false
	}
	return o.ExternalPlanId, true
}

// HasExternalPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasExternalPlanId() bool {
	if o != nil && !IsNil(o.ExternalPlanId) {
		return true
	}

	return false
}

// SetExternalPlanId gets a reference to the given string and assigns it to the ExternalPlanId field.
func (o *UnibeeApiMerchantPlanNewReq) SetExternalPlanId(v string) {
	o.ExternalPlanId = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiMerchantPlanNewReq) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiMerchantPlanNewReq) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiMerchantPlanNewReq) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *UnibeeApiMerchantPlanNewReq) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetIntervalUnit returns the IntervalUnit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetIntervalUnit() string {
	if o == nil || IsNil(o.IntervalUnit) {
		var ret string
		return ret
	}
	return *o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetIntervalUnitOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalUnit) {
		return nil, false
	}
	return o.IntervalUnit, true
}

// HasIntervalUnit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasIntervalUnit() bool {
	if o != nil && !IsNil(o.IntervalUnit) {
		return true
	}

	return false
}

// SetIntervalUnit gets a reference to the given string and assigns it to the IntervalUnit field.
func (o *UnibeeApiMerchantPlanNewReq) SetIntervalUnit(v string) {
	o.IntervalUnit = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPlanNewReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMetricLimits returns the MetricLimits field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetMetricLimits() []UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	if o == nil || IsNil(o.MetricLimits) {
		var ret []UnibeeApiBeanBulkMetricLimitPlanBindingParam
		return ret
	}
	return o.MetricLimits
}

// GetMetricLimitsOk returns a tuple with the MetricLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetMetricLimitsOk() ([]UnibeeApiBeanBulkMetricLimitPlanBindingParam, bool) {
	if o == nil || IsNil(o.MetricLimits) {
		return nil, false
	}
	return o.MetricLimits, true
}

// HasMetricLimits returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasMetricLimits() bool {
	if o != nil && !IsNil(o.MetricLimits) {
		return true
	}

	return false
}

// SetMetricLimits gets a reference to the given []UnibeeApiBeanBulkMetricLimitPlanBindingParam and assigns it to the MetricLimits field.
func (o *UnibeeApiMerchantPlanNewReq) SetMetricLimits(v []UnibeeApiBeanBulkMetricLimitPlanBindingParam) {
	o.MetricLimits = v
}

// GetOnetimeAddonIds returns the OnetimeAddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetOnetimeAddonIds() []int64 {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		var ret []int64
		return ret
	}
	return o.OnetimeAddonIds
}

// GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetOnetimeAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		return nil, false
	}
	return o.OnetimeAddonIds, true
}

// HasOnetimeAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasOnetimeAddonIds() bool {
	if o != nil && !IsNil(o.OnetimeAddonIds) {
		return true
	}

	return false
}

// SetOnetimeAddonIds gets a reference to the given []int64 and assigns it to the OnetimeAddonIds field.
func (o *UnibeeApiMerchantPlanNewReq) SetOnetimeAddonIds(v []int64) {
	o.OnetimeAddonIds = v
}

// GetPlanName returns the PlanName field value
func (o *UnibeeApiMerchantPlanNewReq) GetPlanName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetPlanNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanName, true
}

// SetPlanName sets field value
func (o *UnibeeApiMerchantPlanNewReq) SetPlanName(v string) {
	o.PlanName = v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *UnibeeApiMerchantPlanNewReq) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiMerchantPlanNewReq) SetProductName(v string) {
	o.ProductName = &v
}

// GetTrialAmount returns the TrialAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialAmount() int64 {
	if o == nil || IsNil(o.TrialAmount) {
		var ret int64
		return ret
	}
	return *o.TrialAmount
}

// GetTrialAmountOk returns a tuple with the TrialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialAmount) {
		return nil, false
	}
	return o.TrialAmount, true
}

// HasTrialAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasTrialAmount() bool {
	if o != nil && !IsNil(o.TrialAmount) {
		return true
	}

	return false
}

// SetTrialAmount gets a reference to the given int64 and assigns it to the TrialAmount field.
func (o *UnibeeApiMerchantPlanNewReq) SetTrialAmount(v int64) {
	o.TrialAmount = &v
}

// GetTrialDemand returns the TrialDemand field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialDemand() string {
	if o == nil || IsNil(o.TrialDemand) {
		var ret string
		return ret
	}
	return *o.TrialDemand
}

// GetTrialDemandOk returns a tuple with the TrialDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialDemandOk() (*string, bool) {
	if o == nil || IsNil(o.TrialDemand) {
		return nil, false
	}
	return o.TrialDemand, true
}

// HasTrialDemand returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasTrialDemand() bool {
	if o != nil && !IsNil(o.TrialDemand) {
		return true
	}

	return false
}

// SetTrialDemand gets a reference to the given string and assigns it to the TrialDemand field.
func (o *UnibeeApiMerchantPlanNewReq) SetTrialDemand(v string) {
	o.TrialDemand = &v
}

// GetTrialDurationTime returns the TrialDurationTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialDurationTime() int64 {
	if o == nil || IsNil(o.TrialDurationTime) {
		var ret int64
		return ret
	}
	return *o.TrialDurationTime
}

// GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetTrialDurationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialDurationTime) {
		return nil, false
	}
	return o.TrialDurationTime, true
}

// HasTrialDurationTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasTrialDurationTime() bool {
	if o != nil && !IsNil(o.TrialDurationTime) {
		return true
	}

	return false
}

// SetTrialDurationTime gets a reference to the given int64 and assigns it to the TrialDurationTime field.
func (o *UnibeeApiMerchantPlanNewReq) SetTrialDurationTime(v int64) {
	o.TrialDurationTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanNewReq) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanNewReq) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanNewReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiMerchantPlanNewReq) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeApiMerchantPlanNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonIds) {
		toSerialize["addonIds"] = o.AddonIds
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CancelAtTrialEnd) {
		toSerialize["cancelAtTrialEnd"] = o.CancelAtTrialEnd
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalPlanId) {
		toSerialize["externalPlanId"] = o.ExternalPlanId
	}
	if !IsNil(o.GasPayer) {
		toSerialize["gasPayer"] = o.GasPayer
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.IntervalCount) {
		toSerialize["intervalCount"] = o.IntervalCount
	}
	if !IsNil(o.IntervalUnit) {
		toSerialize["intervalUnit"] = o.IntervalUnit
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetricLimits) {
		toSerialize["metricLimits"] = o.MetricLimits
	}
	if !IsNil(o.OnetimeAddonIds) {
		toSerialize["onetimeAddonIds"] = o.OnetimeAddonIds
	}
	toSerialize["planName"] = o.PlanName
	if !IsNil(o.ProductDescription) {
		toSerialize["productDescription"] = o.ProductDescription
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.TrialAmount) {
		toSerialize["trialAmount"] = o.TrialAmount
	}
	if !IsNil(o.TrialDemand) {
		toSerialize["trialDemand"] = o.TrialDemand
	}
	if !IsNil(o.TrialDurationTime) {
		toSerialize["trialDurationTime"] = o.TrialDurationTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
		"planName",
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

	varUnibeeApiMerchantPlanNewReq := _UnibeeApiMerchantPlanNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanNewReq(varUnibeeApiMerchantPlanNewReq)

	return err
}

type NullableUnibeeApiMerchantPlanNewReq struct {
	value *UnibeeApiMerchantPlanNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanNewReq) Get() *UnibeeApiMerchantPlanNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanNewReq) Set(val *UnibeeApiMerchantPlanNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanNewReq(val *UnibeeApiMerchantPlanNewReq) *NullableUnibeeApiMerchantPlanNewReq {
	return &NullableUnibeeApiMerchantPlanNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


