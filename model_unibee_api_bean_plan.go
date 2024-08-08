/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408081000 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanPlan{}

// UnibeeApiBeanPlan struct for UnibeeApiBeanPlan
type UnibeeApiBeanPlan struct {
	// amount, cent, without tax
	Amount *int64 `json:"amount,omitempty"`
	// binded recurring addon planIds，split with ,
	BindingAddonIds *string `json:"bindingAddonIds,omitempty"`
	// binded onetime addon planIds，split with ,
	BindingOnetimeAddonIds *string `json:"bindingOnetimeAddonIds,omitempty"`
	// whether cancel at subscripiton first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription
	CancelAtTrialEnd *int32 `json:"cancelAtTrialEnd,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	// external_user_id
	ExternalPlanId *string `json:"externalPlanId,omitempty"`
	ExtraMetricData *string `json:"extraMetricData,omitempty"`
	// who pay the gas, merchant|user
	GasPayer *string `json:"gasPayer,omitempty"`
	// home_url
	HomeUrl *string `json:"homeUrl,omitempty"`
	Id *int64 `json:"id,omitempty"`
	// image_url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// period unit count
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// period unit,day|month|year|week
	IntervalUnit *string `json:"intervalUnit,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// PlanName
	PlanName *string `json:"planName,omitempty"`
	// product id
	ProductId *int64 `json:"productId,omitempty"`
	// 1-UnPublish,2-Publish, Use For Display Plan At UserPortal
	PublishStatus *int32 `json:"publishStatus,omitempty"`
	// status，1-editing，2-active，3-inactive，4-expired
	Status *int32 `json:"status,omitempty"`
	// TaxPercentage 1000 = 10%
	TaxPercentage *int32 `json:"taxPercentage,omitempty"`
	// price of trial period
	TrialAmount *int64 `json:"trialAmount,omitempty"`
	TrialDemand *string `json:"trialDemand,omitempty"`
	// duration of trial
	TrialDurationTime *int64 `json:"trialDurationTime,omitempty"`
	// type，1-main plan，2-addon plan
	Type *int32 `json:"type,omitempty"`
}

// NewUnibeeApiBeanPlan instantiates a new UnibeeApiBeanPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanPlan() *UnibeeApiBeanPlan {
	this := UnibeeApiBeanPlan{}
	return &this
}

// NewUnibeeApiBeanPlanWithDefaults instantiates a new UnibeeApiBeanPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanPlanWithDefaults() *UnibeeApiBeanPlan {
	this := UnibeeApiBeanPlan{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeApiBeanPlan) SetAmount(v int64) {
	o.Amount = &v
}

// GetBindingAddonIds returns the BindingAddonIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetBindingAddonIds() string {
	if o == nil || IsNil(o.BindingAddonIds) {
		var ret string
		return ret
	}
	return *o.BindingAddonIds
}

// GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetBindingAddonIdsOk() (*string, bool) {
	if o == nil || IsNil(o.BindingAddonIds) {
		return nil, false
	}
	return o.BindingAddonIds, true
}

// HasBindingAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasBindingAddonIds() bool {
	if o != nil && !IsNil(o.BindingAddonIds) {
		return true
	}

	return false
}

// SetBindingAddonIds gets a reference to the given string and assigns it to the BindingAddonIds field.
func (o *UnibeeApiBeanPlan) SetBindingAddonIds(v string) {
	o.BindingAddonIds = &v
}

// GetBindingOnetimeAddonIds returns the BindingOnetimeAddonIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetBindingOnetimeAddonIds() string {
	if o == nil || IsNil(o.BindingOnetimeAddonIds) {
		var ret string
		return ret
	}
	return *o.BindingOnetimeAddonIds
}

// GetBindingOnetimeAddonIdsOk returns a tuple with the BindingOnetimeAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetBindingOnetimeAddonIdsOk() (*string, bool) {
	if o == nil || IsNil(o.BindingOnetimeAddonIds) {
		return nil, false
	}
	return o.BindingOnetimeAddonIds, true
}

// HasBindingOnetimeAddonIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasBindingOnetimeAddonIds() bool {
	if o != nil && !IsNil(o.BindingOnetimeAddonIds) {
		return true
	}

	return false
}

// SetBindingOnetimeAddonIds gets a reference to the given string and assigns it to the BindingOnetimeAddonIds field.
func (o *UnibeeApiBeanPlan) SetBindingOnetimeAddonIds(v string) {
	o.BindingOnetimeAddonIds = &v
}

// GetCancelAtTrialEnd returns the CancelAtTrialEnd field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetCancelAtTrialEnd() int32 {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		var ret int32
		return ret
	}
	return *o.CancelAtTrialEnd
}

// GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetCancelAtTrialEndOk() (*int32, bool) {
	if o == nil || IsNil(o.CancelAtTrialEnd) {
		return nil, false
	}
	return o.CancelAtTrialEnd, true
}

// HasCancelAtTrialEnd returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasCancelAtTrialEnd() bool {
	if o != nil && !IsNil(o.CancelAtTrialEnd) {
		return true
	}

	return false
}

// SetCancelAtTrialEnd gets a reference to the given int32 and assigns it to the CancelAtTrialEnd field.
func (o *UnibeeApiBeanPlan) SetCancelAtTrialEnd(v int32) {
	o.CancelAtTrialEnd = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanPlan) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanPlan) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanPlan) SetDescription(v string) {
	o.Description = &v
}

// GetExternalPlanId returns the ExternalPlanId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetExternalPlanId() string {
	if o == nil || IsNil(o.ExternalPlanId) {
		var ret string
		return ret
	}
	return *o.ExternalPlanId
}

// GetExternalPlanIdOk returns a tuple with the ExternalPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetExternalPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPlanId) {
		return nil, false
	}
	return o.ExternalPlanId, true
}

// HasExternalPlanId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasExternalPlanId() bool {
	if o != nil && !IsNil(o.ExternalPlanId) {
		return true
	}

	return false
}

// SetExternalPlanId gets a reference to the given string and assigns it to the ExternalPlanId field.
func (o *UnibeeApiBeanPlan) SetExternalPlanId(v string) {
	o.ExternalPlanId = &v
}

// GetExtraMetricData returns the ExtraMetricData field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetExtraMetricData() string {
	if o == nil || IsNil(o.ExtraMetricData) {
		var ret string
		return ret
	}
	return *o.ExtraMetricData
}

// GetExtraMetricDataOk returns a tuple with the ExtraMetricData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetExtraMetricDataOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraMetricData) {
		return nil, false
	}
	return o.ExtraMetricData, true
}

// HasExtraMetricData returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasExtraMetricData() bool {
	if o != nil && !IsNil(o.ExtraMetricData) {
		return true
	}

	return false
}

// SetExtraMetricData gets a reference to the given string and assigns it to the ExtraMetricData field.
func (o *UnibeeApiBeanPlan) SetExtraMetricData(v string) {
	o.ExtraMetricData = &v
}

// GetGasPayer returns the GasPayer field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetGasPayer() string {
	if o == nil || IsNil(o.GasPayer) {
		var ret string
		return ret
	}
	return *o.GasPayer
}

// GetGasPayerOk returns a tuple with the GasPayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetGasPayerOk() (*string, bool) {
	if o == nil || IsNil(o.GasPayer) {
		return nil, false
	}
	return o.GasPayer, true
}

// HasGasPayer returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasGasPayer() bool {
	if o != nil && !IsNil(o.GasPayer) {
		return true
	}

	return false
}

// SetGasPayer gets a reference to the given string and assigns it to the GasPayer field.
func (o *UnibeeApiBeanPlan) SetGasPayer(v string) {
	o.GasPayer = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiBeanPlan) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanPlan) SetId(v int64) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiBeanPlan) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *UnibeeApiBeanPlan) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetIntervalUnit returns the IntervalUnit field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetIntervalUnit() string {
	if o == nil || IsNil(o.IntervalUnit) {
		var ret string
		return ret
	}
	return *o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetIntervalUnitOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalUnit) {
		return nil, false
	}
	return o.IntervalUnit, true
}

// HasIntervalUnit returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasIntervalUnit() bool {
	if o != nil && !IsNil(o.IntervalUnit) {
		return true
	}

	return false
}

// SetIntervalUnit gets a reference to the given string and assigns it to the IntervalUnit field.
func (o *UnibeeApiBeanPlan) SetIntervalUnit(v string) {
	o.IntervalUnit = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanPlan) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanPlan) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *UnibeeApiBeanPlan) SetPlanName(v string) {
	o.PlanName = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiBeanPlan) SetProductId(v int64) {
	o.ProductId = &v
}

// GetPublishStatus returns the PublishStatus field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetPublishStatus() int32 {
	if o == nil || IsNil(o.PublishStatus) {
		var ret int32
		return ret
	}
	return *o.PublishStatus
}

// GetPublishStatusOk returns a tuple with the PublishStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetPublishStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.PublishStatus) {
		return nil, false
	}
	return o.PublishStatus, true
}

// HasPublishStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasPublishStatus() bool {
	if o != nil && !IsNil(o.PublishStatus) {
		return true
	}

	return false
}

// SetPublishStatus gets a reference to the given int32 and assigns it to the PublishStatus field.
func (o *UnibeeApiBeanPlan) SetPublishStatus(v int32) {
	o.PublishStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanPlan) SetStatus(v int32) {
	o.Status = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetTaxPercentage() int32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret int32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetTaxPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given int32 and assigns it to the TaxPercentage field.
func (o *UnibeeApiBeanPlan) SetTaxPercentage(v int32) {
	o.TaxPercentage = &v
}

// GetTrialAmount returns the TrialAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetTrialAmount() int64 {
	if o == nil || IsNil(o.TrialAmount) {
		var ret int64
		return ret
	}
	return *o.TrialAmount
}

// GetTrialAmountOk returns a tuple with the TrialAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetTrialAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialAmount) {
		return nil, false
	}
	return o.TrialAmount, true
}

// HasTrialAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasTrialAmount() bool {
	if o != nil && !IsNil(o.TrialAmount) {
		return true
	}

	return false
}

// SetTrialAmount gets a reference to the given int64 and assigns it to the TrialAmount field.
func (o *UnibeeApiBeanPlan) SetTrialAmount(v int64) {
	o.TrialAmount = &v
}

// GetTrialDemand returns the TrialDemand field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetTrialDemand() string {
	if o == nil || IsNil(o.TrialDemand) {
		var ret string
		return ret
	}
	return *o.TrialDemand
}

// GetTrialDemandOk returns a tuple with the TrialDemand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetTrialDemandOk() (*string, bool) {
	if o == nil || IsNil(o.TrialDemand) {
		return nil, false
	}
	return o.TrialDemand, true
}

// HasTrialDemand returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasTrialDemand() bool {
	if o != nil && !IsNil(o.TrialDemand) {
		return true
	}

	return false
}

// SetTrialDemand gets a reference to the given string and assigns it to the TrialDemand field.
func (o *UnibeeApiBeanPlan) SetTrialDemand(v string) {
	o.TrialDemand = &v
}

// GetTrialDurationTime returns the TrialDurationTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetTrialDurationTime() int64 {
	if o == nil || IsNil(o.TrialDurationTime) {
		var ret int64
		return ret
	}
	return *o.TrialDurationTime
}

// GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetTrialDurationTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TrialDurationTime) {
		return nil, false
	}
	return o.TrialDurationTime, true
}

// HasTrialDurationTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasTrialDurationTime() bool {
	if o != nil && !IsNil(o.TrialDurationTime) {
		return true
	}

	return false
}

// SetTrialDurationTime gets a reference to the given int64 and assigns it to the TrialDurationTime field.
func (o *UnibeeApiBeanPlan) SetTrialDurationTime(v int64) {
	o.TrialDurationTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiBeanPlan) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanPlan) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiBeanPlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiBeanPlan) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeApiBeanPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BindingAddonIds) {
		toSerialize["bindingAddonIds"] = o.BindingAddonIds
	}
	if !IsNil(o.BindingOnetimeAddonIds) {
		toSerialize["bindingOnetimeAddonIds"] = o.BindingOnetimeAddonIds
	}
	if !IsNil(o.CancelAtTrialEnd) {
		toSerialize["cancelAtTrialEnd"] = o.CancelAtTrialEnd
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
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
	if !IsNil(o.ExtraMetricData) {
		toSerialize["extraMetricData"] = o.ExtraMetricData
	}
	if !IsNil(o.GasPayer) {
		toSerialize["gasPayer"] = o.GasPayer
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.PublishStatus) {
		toSerialize["publishStatus"] = o.PublishStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
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

type NullableUnibeeApiBeanPlan struct {
	value *UnibeeApiBeanPlan
	isSet bool
}

func (v NullableUnibeeApiBeanPlan) Get() *UnibeeApiBeanPlan {
	return v.value
}

func (v *NullableUnibeeApiBeanPlan) Set(val *UnibeeApiBeanPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanPlan(val *UnibeeApiBeanPlan) *NullableUnibeeApiBeanPlan {
	return &NullableUnibeeApiBeanPlan{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


