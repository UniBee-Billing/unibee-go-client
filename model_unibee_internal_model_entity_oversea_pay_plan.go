/*
OpenAPI UniBee

This is UniBee api server, For this sample, you can use the api key `EUXAgwv3Vcr1PFWt2SgBumMHXn3ImBqM` to test the authorization filters

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UnibeeInternalModelEntityOverseaPayPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeInternalModelEntityOverseaPayPlan{}

// UnibeeInternalModelEntityOverseaPayPlan struct for UnibeeInternalModelEntityOverseaPayPlan
type UnibeeInternalModelEntityOverseaPayPlan struct {
	// amount, cent, without tax
	Amount *int64 `json:"amount,omitempty"`
	// binded addon planIds，split with ,
	BindingAddonIds *string `json:"bindingAddonIds,omitempty"`
	// company id
	CompanyId *int64 `json:"companyId,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency
	Currency *string `json:"currency,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	// gateway product description
	GatewayProductDescription *string `json:"gatewayProductDescription,omitempty"`
	// gateway product name
	GatewayProductName *string `json:"gatewayProductName,omitempty"`
	// create time
	GmtCreate *string `json:"gmtCreate,omitempty"`
	// update time
	GmtModify *string `json:"gmtModify,omitempty"`
	// home_url
	HomeUrl *string `json:"homeUrl,omitempty"`
	Id *int32 `json:"id,omitempty"`
	// image_url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// period unit count
	IntervalCount *int32 `json:"intervalCount,omitempty"`
	// period unit,day|month|year|week
	IntervalUnit *string `json:"intervalUnit,omitempty"`
	// 0-UnDeleted，1-Deleted
	IsDeleted *int32 `json:"isDeleted,omitempty"`
	// merchant id
	MerchantId *int32 `json:"merchantId,omitempty"`
	// PlanName
	PlanName *string `json:"planName,omitempty"`
	// 1-UnPublish,2-Publish, Use For Display Plan At UserPortal
	PublishStatus *int32 `json:"publishStatus,omitempty"`
	// status，1-editing，2-active，3-inactive，4-expired
	Status *int32 `json:"status,omitempty"`
	// deperated
	TaxInclusive *int32 `json:"taxInclusive,omitempty"`
	// tax scale 1000 = 10%
	TaxScale *int32 `json:"taxScale,omitempty"`
	// type，1-main plan，2-addon plan
	Type *int32 `json:"type,omitempty"`
}

// NewUnibeeInternalModelEntityOverseaPayPlan instantiates a new UnibeeInternalModelEntityOverseaPayPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeInternalModelEntityOverseaPayPlan() *UnibeeInternalModelEntityOverseaPayPlan {
	this := UnibeeInternalModelEntityOverseaPayPlan{}
	return &this
}

// NewUnibeeInternalModelEntityOverseaPayPlanWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeInternalModelEntityOverseaPayPlanWithDefaults() *UnibeeInternalModelEntityOverseaPayPlan {
	this := UnibeeInternalModelEntityOverseaPayPlan{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetAmount() int64 {
	if o == nil || IsNil(o.Amount) {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetAmount(v int64) {
	o.Amount = &v
}

// GetBindingAddonIds returns the BindingAddonIds field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetBindingAddonIds() string {
	if o == nil || IsNil(o.BindingAddonIds) {
		var ret string
		return ret
	}
	return *o.BindingAddonIds
}

// GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetBindingAddonIdsOk() (*string, bool) {
	if o == nil || IsNil(o.BindingAddonIds) {
		return nil, false
	}
	return o.BindingAddonIds, true
}

// HasBindingAddonIds returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasBindingAddonIds() bool {
	if o != nil && !IsNil(o.BindingAddonIds) {
		return true
	}

	return false
}

// SetBindingAddonIds gets a reference to the given string and assigns it to the BindingAddonIds field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetBindingAddonIds(v string) {
	o.BindingAddonIds = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCompanyId() int64 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int64
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCompanyIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int64 and assigns it to the CompanyId field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCompanyId(v int64) {
	o.CompanyId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetCurrency(v string) {
	o.Currency = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetDescription(v string) {
	o.Description = &v
}

// GetGatewayProductDescription returns the GatewayProductDescription field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductDescription() string {
	if o == nil || IsNil(o.GatewayProductDescription) {
		var ret string
		return ret
	}
	return *o.GatewayProductDescription
}

// GetGatewayProductDescriptionOk returns a tuple with the GatewayProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayProductDescription) {
		return nil, false
	}
	return o.GatewayProductDescription, true
}

// HasGatewayProductDescription returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGatewayProductDescription() bool {
	if o != nil && !IsNil(o.GatewayProductDescription) {
		return true
	}

	return false
}

// SetGatewayProductDescription gets a reference to the given string and assigns it to the GatewayProductDescription field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGatewayProductDescription(v string) {
	o.GatewayProductDescription = &v
}

// GetGatewayProductName returns the GatewayProductName field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductName() string {
	if o == nil || IsNil(o.GatewayProductName) {
		var ret string
		return ret
	}
	return *o.GatewayProductName
}

// GetGatewayProductNameOk returns a tuple with the GatewayProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGatewayProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayProductName) {
		return nil, false
	}
	return o.GatewayProductName, true
}

// HasGatewayProductName returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGatewayProductName() bool {
	if o != nil && !IsNil(o.GatewayProductName) {
		return true
	}

	return false
}

// SetGatewayProductName gets a reference to the given string and assigns it to the GatewayProductName field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGatewayProductName(v string) {
	o.GatewayProductName = &v
}

// GetGmtCreate returns the GmtCreate field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtCreate() string {
	if o == nil || IsNil(o.GmtCreate) {
		var ret string
		return ret
	}
	return *o.GmtCreate
}

// GetGmtCreateOk returns a tuple with the GmtCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtCreateOk() (*string, bool) {
	if o == nil || IsNil(o.GmtCreate) {
		return nil, false
	}
	return o.GmtCreate, true
}

// HasGmtCreate returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGmtCreate() bool {
	if o != nil && !IsNil(o.GmtCreate) {
		return true
	}

	return false
}

// SetGmtCreate gets a reference to the given string and assigns it to the GmtCreate field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGmtCreate(v string) {
	o.GmtCreate = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtModify() string {
	if o == nil || IsNil(o.GmtModify) {
		var ret string
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetGmtModifyOk() (*string, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given string and assigns it to the GmtModify field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetGmtModify(v string) {
	o.GmtModify = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetId(v int32) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIntervalCount returns the IntervalCount field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalCount() int32 {
	if o == nil || IsNil(o.IntervalCount) {
		var ret int32
		return ret
	}
	return *o.IntervalCount
}

// GetIntervalCountOk returns a tuple with the IntervalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalCount) {
		return nil, false
	}
	return o.IntervalCount, true
}

// HasIntervalCount returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIntervalCount() bool {
	if o != nil && !IsNil(o.IntervalCount) {
		return true
	}

	return false
}

// SetIntervalCount gets a reference to the given int32 and assigns it to the IntervalCount field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIntervalCount(v int32) {
	o.IntervalCount = &v
}

// GetIntervalUnit returns the IntervalUnit field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalUnit() string {
	if o == nil || IsNil(o.IntervalUnit) {
		var ret string
		return ret
	}
	return *o.IntervalUnit
}

// GetIntervalUnitOk returns a tuple with the IntervalUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIntervalUnitOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalUnit) {
		return nil, false
	}
	return o.IntervalUnit, true
}

// HasIntervalUnit returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIntervalUnit() bool {
	if o != nil && !IsNil(o.IntervalUnit) {
		return true
	}

	return false
}

// SetIntervalUnit gets a reference to the given string and assigns it to the IntervalUnit field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIntervalUnit(v string) {
	o.IntervalUnit = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIsDeleted() int32 {
	if o == nil || IsNil(o.IsDeleted) {
		var ret int32
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetIsDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given int32 and assigns it to the IsDeleted field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetIsDeleted(v int32) {
	o.IsDeleted = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMerchantId() int32 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int32
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetMerchantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int32 and assigns it to the MerchantId field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetMerchantId(v int32) {
	o.MerchantId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPublishStatus returns the PublishStatus field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPublishStatus() int32 {
	if o == nil || IsNil(o.PublishStatus) {
		var ret int32
		return ret
	}
	return *o.PublishStatus
}

// GetPublishStatusOk returns a tuple with the PublishStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetPublishStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.PublishStatus) {
		return nil, false
	}
	return o.PublishStatus, true
}

// HasPublishStatus returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasPublishStatus() bool {
	if o != nil && !IsNil(o.PublishStatus) {
		return true
	}

	return false
}

// SetPublishStatus gets a reference to the given int32 and assigns it to the PublishStatus field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetPublishStatus(v int32) {
	o.PublishStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetStatus(v int32) {
	o.Status = &v
}

// GetTaxInclusive returns the TaxInclusive field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxInclusive() int32 {
	if o == nil || IsNil(o.TaxInclusive) {
		var ret int32
		return ret
	}
	return *o.TaxInclusive
}

// GetTaxInclusiveOk returns a tuple with the TaxInclusive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxInclusiveOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxInclusive) {
		return nil, false
	}
	return o.TaxInclusive, true
}

// HasTaxInclusive returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasTaxInclusive() bool {
	if o != nil && !IsNil(o.TaxInclusive) {
		return true
	}

	return false
}

// SetTaxInclusive gets a reference to the given int32 and assigns it to the TaxInclusive field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetTaxInclusive(v int32) {
	o.TaxInclusive = &v
}

// GetTaxScale returns the TaxScale field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxScale() int32 {
	if o == nil || IsNil(o.TaxScale) {
		var ret int32
		return ret
	}
	return *o.TaxScale
}

// GetTaxScaleOk returns a tuple with the TaxScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTaxScaleOk() (*int32, bool) {
	if o == nil || IsNil(o.TaxScale) {
		return nil, false
	}
	return o.TaxScale, true
}

// HasTaxScale returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasTaxScale() bool {
	if o != nil && !IsNil(o.TaxScale) {
		return true
	}

	return false
}

// SetTaxScale gets a reference to the given int32 and assigns it to the TaxScale field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetTaxScale(v int32) {
	o.TaxScale = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeInternalModelEntityOverseaPayPlan) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeInternalModelEntityOverseaPayPlan) SetType(v int32) {
	o.Type = &v
}

func (o UnibeeInternalModelEntityOverseaPayPlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeInternalModelEntityOverseaPayPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.BindingAddonIds) {
		toSerialize["bindingAddonIds"] = o.BindingAddonIds
	}
	if !IsNil(o.CompanyId) {
		toSerialize["companyId"] = o.CompanyId
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
	if !IsNil(o.GatewayProductDescription) {
		toSerialize["gatewayProductDescription"] = o.GatewayProductDescription
	}
	if !IsNil(o.GatewayProductName) {
		toSerialize["gatewayProductName"] = o.GatewayProductName
	}
	if !IsNil(o.GmtCreate) {
		toSerialize["gmtCreate"] = o.GmtCreate
	}
	if !IsNil(o.GmtModify) {
		toSerialize["gmtModify"] = o.GmtModify
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
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.PublishStatus) {
		toSerialize["publishStatus"] = o.PublishStatus
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaxInclusive) {
		toSerialize["taxInclusive"] = o.TaxInclusive
	}
	if !IsNil(o.TaxScale) {
		toSerialize["taxScale"] = o.TaxScale
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeInternalModelEntityOverseaPayPlan struct {
	value *UnibeeInternalModelEntityOverseaPayPlan
	isSet bool
}

func (v NullableUnibeeInternalModelEntityOverseaPayPlan) Get() *UnibeeInternalModelEntityOverseaPayPlan {
	return v.value
}

func (v *NullableUnibeeInternalModelEntityOverseaPayPlan) Set(val *UnibeeInternalModelEntityOverseaPayPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeInternalModelEntityOverseaPayPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeInternalModelEntityOverseaPayPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeInternalModelEntityOverseaPayPlan(val *UnibeeInternalModelEntityOverseaPayPlan) *NullableUnibeeInternalModelEntityOverseaPayPlan {
	return &NullableUnibeeInternalModelEntityOverseaPayPlan{value: val, isSet: true}
}

func (v NullableUnibeeInternalModelEntityOverseaPayPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeInternalModelEntityOverseaPayPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


