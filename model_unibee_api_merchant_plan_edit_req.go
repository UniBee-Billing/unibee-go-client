/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantPlanEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanEditReq{}

// UnibeeApiMerchantPlanEditReq Edit exist plan, amount|currency|intervalUnit|intervalCount is not editable when plan is active 
type UnibeeApiMerchantPlanEditReq struct {
	// Plan Ids Of Recurring Addon Type
	AddonIds []int64 `json:"addonIds,omitempty"`
	// CaptureAmount of plan, not editable when plan is active
	Amount *int32 `json:"amount,omitempty"`
	// whether cancel at subscripiton first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription
	CancelAtTrialEnd *int32 `json:"cancelAtTrialEnd,omitempty"`
	// Currency of plan, not editable when plan is active
	Currency *string `json:"currency,omitempty"`
	// Description of plan
	Description *string `json:"description,omitempty"`
	// ExternalPlanId
	ExternalPlanId *string `json:"externalPlanId,omitempty"`
	// who pay the gas for crypto payment, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	// HomeUrl,Start With: http
	HomeUrl *string `json:"homeUrl,omitempty"`
	// ImageUrl,Start With: http
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Number,intervalUnit of plan, not editable when plan is active
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// Interval unit of plan，em: day|month|year|week, not editable when plan is active
	IntervalUnit *string `json:"intervalUnit,omitempty"`
	// Metadata，Map
	Metadata *map[string]string `json:"metadata,omitempty"`
	// Plan's MetricLimit List
	MetricLimits []UnibeeApiBeanBulkMetricLimitPlanBindingParam `json:"metricLimits,omitempty"`
	// Plan Ids Of Onetime Addon Type
	OnetimeAddonIds []int64 `json:"onetimeAddonIds,omitempty"`
	// Id of plan
	PlanId int64 `json:"planId"`
	// Name of plan
	PlanName *string `json:"planName,omitempty"`
	// ProductDescription of plan, Default copy description
	ProductDescription *string `json:"productDescription,omitempty"`
	// Id of product which plan to linked
	ProductId *int32 `json:"productId,omitempty"`
	// ProductName of plan, Default copy planName
	ProductName *string `json:"productName,omitempty"`
	// price of trial period， not available for addon
	TrialAmount *int32 `json:"trialAmount,omitempty"`
	// demand of trial, not available for addon, example, paymentMethod, payment method will ask for subscription trial start
	TrialDemand *string `json:"trialDemand,omitempty"`
	// duration of trial， not available for addon
	TrialDurationTime *int32 `json:"trialDurationTime,omitempty"`
}

type _UnibeeApiMerchantPlanEditReq UnibeeApiMerchantPlanEditReq

// NewUnibeeApiMerchantPlanEditReq instantiates a new UnibeeApiMerchantPlanEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanEditReq(planId int64) *UnibeeApiMerchantPlanEditReq {
	this := UnibeeApiMerchantPlanEditReq{}
	this.PlanId = planId
	return &this
}

// NewUnibeeApiMerchantPlanEditReqWithDefaults instantiates a new UnibeeApiMerchantPlanEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanEditReqWithDefaults() *UnibeeApiMerchantPlanEditReq {
	this := UnibeeApiMerchantPlanEditReq{}
	return &this
}

// GetAddonIds returns the AddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetAddonIds() []int64 {
	if o == nil || IsNil(o.AddonIds) {
		var ret []int64
		return ret
	}
	return o.AddonIds
}

// GetAddonIdsOk returns a tuple with the AddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AddonIds) {
		return nil, false
	}
	return o.AddonIds, true
}

// HasAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasAddonIds() bool {
	if o != nil && !IsNil(o.AddonIds) {
		return true
	}

	return false
}

// SetAddonIds gets a reference to the given []int64 and assigns it to the AddonIds field.
func (o *UnibeeApiMerchantPlanEditReq) SetAddonIds(v []int64) {
	o.AddonIds = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetAmount() int32 {
	if o == nil || IsNil(o.Amount) {
		var ret int32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int32 and assigns it to the Amount field.
func (o *UnibeeApiMerchantPlanEditReq) SetAmount(v int32) {
	o.Amount = &v
}

// GetCancelAtTrialEnd returns the CancelAtTrialEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetCancelAtTrialEnd() int32 {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		var ret int32
		return ret
	}
	return *o.CancelAtTrialEnd
}

// GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetCancelAtTrialEndOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		return nil, false
	}
	return o.CancelAtTrialEnd, true
}

// HasCancelAtTrialEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasCancelAtTrialEnd() bool {
	if o != nil && !IsNil(o.CancelAtTrialEnd) {
		return true
	}

	return false
}

// SetCancelAtTrialEnd gets a reference to the given int32 and assigns it to the CancelAtTrialEnd field.
func (o *UnibeeApiMerchantPlanEditReq) SetCancelAtTrialEnd(v int32) {
	o.CancelAtTrialEnd = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPlanEditReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantPlanEditReq) SetDescription(v string) {
	o.Description = &v
}

// GetExternalPlanId returns the ExternalPlanId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetExternalPlanId() string {
	if o == nil || IsNil(o.ExternalPlanId) {
		var ret string
		return ret
	}
	return *o.ExternalPlanId
}

// GetExternalPlanIdOk returns a tuple with the ExternalPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetExternalPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPlanId) {
		return nil, false
	}
	return o.ExternalPlanId, true
}

// HasExternalPlanId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasExternalPlanId() bool {
	if o != nil && !IsNil(o.ExternalPlanId) {
		return true
	}

	return false
}

// SetExternalPlanId gets a reference to the given string and assigns it to the ExternalPlanId field.
func (o *UnibeeApiMerchantPlanEditReq) SetExternalPlanId(v string) {
	o.ExternalPlanId = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiMerchantPlanEditReq) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiMerchantPlanEditReq) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiMerchantPlanEditReq) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *UnibeeApiMerchantPlanEditReq) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetIntervalUnit returns the IntervalUnit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnit() string {
	if o == nil || IsNil(o.IntervalUnit) {
		var ret string
		return ret
	}
	return *o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnitOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalUnit) {
		return nil, false
	}
	return o.IntervalUnit, true
}

// HasIntervalUnit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasIntervalUnit() bool {
	if o != nil && !IsNil(o.IntervalUnit) {
		return true
	}

	return false
}

// SetIntervalUnit gets a reference to the given string and assigns it to the IntervalUnit field.
func (o *UnibeeApiMerchantPlanEditReq) SetIntervalUnit(v string) {
	o.IntervalUnit = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UnibeeApiMerchantPlanEditReq) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMetricLimits returns the MetricLimits field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimits() []UnibeeApiBeanBulkMetricLimitPlanBindingParam {
	if o == nil || IsNil(o.MetricLimits) {
		var ret []UnibeeApiBeanBulkMetricLimitPlanBindingParam
		return ret
	}
	return o.MetricLimits
}

// GetMetricLimitsOk returns a tuple with the MetricLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimitsOk() ([]UnibeeApiBeanBulkMetricLimitPlanBindingParam, bool) {
	if o == nil || IsNil(o.MetricLimits) {
		return nil, false
	}
	return o.MetricLimits, true
}

// HasMetricLimits returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasMetricLimits() bool {
	if o != nil && !IsNil(o.MetricLimits) {
		return true
	}

	return false
}

// SetMetricLimits gets a reference to the given []UnibeeApiBeanBulkMetricLimitPlanBindingParam and assigns it to the MetricLimits field.
func (o *UnibeeApiMerchantPlanEditReq) SetMetricLimits(v []UnibeeApiBeanBulkMetricLimitPlanBindingParam) {
	o.MetricLimits = v
}

// GetOnetimeAddonIds returns the OnetimeAddonIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetOnetimeAddonIds() []int64 {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		var ret []int64
		return ret
	}
	return o.OnetimeAddonIds
}

// GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetOnetimeAddonIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.OnetimeAddonIds) {
		return nil, false
	}
	return o.OnetimeAddonIds, true
}

// HasOnetimeAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasOnetimeAddonIds() bool {
	if o != nil && !IsNil(o.OnetimeAddonIds) {
		return true
	}

	return false
}

// SetOnetimeAddonIds gets a reference to the given []int64 and assigns it to the OnetimeAddonIds field.
func (o *UnibeeApiMerchantPlanEditReq) SetOnetimeAddonIds(v []int64) {
	o.OnetimeAddonIds = v
}

// GetPlanId returns the PlanId field value
func (o *UnibeeApiMerchantPlanEditReq) GetPlanId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetPlanIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *UnibeeApiMerchantPlanEditReq) SetPlanId(v int64) {
	o.PlanId = v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *UnibeeApiMerchantPlanEditReq) SetPlanName(v string) {
	o.PlanName = &v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *UnibeeApiMerchantPlanEditReq) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantPlanEditReq) SetProductId(v int32) {
	o.ProductId = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiMerchantPlanEditReq) SetProductName(v string) {
	o.ProductName = &v
}

// GetTrialAmount returns the TrialAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialAmount() int32 {
	if o == nil || IsNil(o.TrialAmount) {
		var ret int32
		return ret
	}
	return *o.TrialAmount
}

// GetTrialAmountOk returns a tuple with the TrialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialAmount) {
		return nil, false
	}
	return o.TrialAmount, true
}

// HasTrialAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasTrialAmount() bool {
	if o != nil && !IsNil(o.TrialAmount) {
		return true
	}

	return false
}

// SetTrialAmount gets a reference to the given int32 and assigns it to the TrialAmount field.
func (o *UnibeeApiMerchantPlanEditReq) SetTrialAmount(v int32) {
	o.TrialAmount = &v
}

// GetTrialDemand returns the TrialDemand field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialDemand() string {
	if o == nil || IsNil(o.TrialDemand) {
		var ret string
		return ret
	}
	return *o.TrialDemand
}

// GetTrialDemandOk returns a tuple with the TrialDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialDemandOk() (*string, bool) {
	if o == nil || IsNil(o.TrialDemand) {
		return nil, false
	}
	return o.TrialDemand, true
}

// HasTrialDemand returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasTrialDemand() bool {
	if o != nil && !IsNil(o.TrialDemand) {
		return true
	}

	return false
}

// SetTrialDemand gets a reference to the given string and assigns it to the TrialDemand field.
func (o *UnibeeApiMerchantPlanEditReq) SetTrialDemand(v string) {
	o.TrialDemand = &v
}

// GetTrialDurationTime returns the TrialDurationTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialDurationTime() int32 {
	if o == nil || IsNil(o.TrialDurationTime) {
		var ret int32
		return ret
	}
	return *o.TrialDurationTime
}

// GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanEditReq) GetTrialDurationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialDurationTime) {
		return nil, false
	}
	return o.TrialDurationTime, true
}

// HasTrialDurationTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanEditReq) HasTrialDurationTime() bool {
	if o != nil && !IsNil(o.TrialDurationTime) {
		return true
	}

	return false
}

// SetTrialDurationTime gets a reference to the given int32 and assigns it to the TrialDurationTime field.
func (o *UnibeeApiMerchantPlanEditReq) SetTrialDurationTime(v int32) {
	o.TrialDurationTime = &v
}

func (o UnibeeApiMerchantPlanEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonIds) {
		toSerialize["addonIds"] = o.AddonIds
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CancelAtTrialEnd) {
		toSerialize["cancelAtTrialEnd"] = o.CancelAtTrialEnd
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
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
	toSerialize["planId"] = o.PlanId
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.ProductDescription) {
		toSerialize["productDescription"] = o.ProductDescription
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
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
	return toSerialize, nil
}

func (o *UnibeeApiMerchantPlanEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
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

	varUnibeeApiMerchantPlanEditReq := _UnibeeApiMerchantPlanEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantPlanEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantPlanEditReq(varUnibeeApiMerchantPlanEditReq)

	return err
}

type NullableUnibeeApiMerchantPlanEditReq struct {
	value *UnibeeApiMerchantPlanEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanEditReq) Get() *UnibeeApiMerchantPlanEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanEditReq) Set(val *UnibeeApiMerchantPlanEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanEditReq(val *UnibeeApiMerchantPlanEditReq) *NullableUnibeeApiMerchantPlanEditReq {
	return &NullableUnibeeApiMerchantPlanEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


