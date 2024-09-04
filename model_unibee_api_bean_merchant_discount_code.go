/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202409040645 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantDiscountCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantDiscountCode{}

// UnibeeApiBeanMerchantDiscountCode struct for UnibeeApiBeanMerchantDiscountCode
type UnibeeApiBeanMerchantDiscountCode struct {
	// billing_type, 1-one-time, 2-recurring
	BillingType *int32 `json:"billingType,omitempty"`
	// code
	Code *string `json:"code,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// currency of discount, available when discount_type is fixed_amount
	Currency *string `json:"currency,omitempty"`
	// the count limitation of subscription cycle , 0-no limit
	CycleLimit *int32 `json:"cycleLimit,omitempty"`
	// amount of discount, available when discount_type is fixed_amount
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// percentage of discount, 100=1%, available when discount_type is percentage
	DiscountPercentage *int64 `json:"discountPercentage,omitempty"`
	// discount_type, 1-percentage, 2-fixed_amount
	DiscountType *int32 `json:"discountType,omitempty"`
	// end of discount available utc time, 0-invalid
	EndTime *int64 `json:"endTime,omitempty"`
	// Id
	Id *int64 `json:"id,omitempty"`
	// merchantId
	MerchantId *int64 `json:"merchantId,omitempty"`
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// Ids of plan which discount code can effect, default effect all plans if not set
	PlanIds []int64 `json:"planIds,omitempty"`
	// start of discount available utc time
	StartTime *int64 `json:"startTime,omitempty"`
	// status, 1-editable, 2-active, 3-deactive, 4-expire
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiBeanMerchantDiscountCode instantiates a new UnibeeApiBeanMerchantDiscountCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantDiscountCode() *UnibeeApiBeanMerchantDiscountCode {
	this := UnibeeApiBeanMerchantDiscountCode{}
	return &this
}

// NewUnibeeApiBeanMerchantDiscountCodeWithDefaults instantiates a new UnibeeApiBeanMerchantDiscountCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantDiscountCodeWithDefaults() *UnibeeApiBeanMerchantDiscountCode {
	this := UnibeeApiBeanMerchantDiscountCode{}
	return &this
}

// GetBillingType returns the BillingType field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetBillingType() int32 {
	if o == nil || IsNil(o.BillingType) {
		var ret int32
		return ret
	}
	return *o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetBillingTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingType) {
		return nil, false
	}
	return o.BillingType, true
}

// HasBillingType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasBillingType() bool {
	if o != nil && !IsNil(o.BillingType) {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given int32 and assigns it to the BillingType field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetBillingType(v int32) {
	o.BillingType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetCode(v string) {
	o.Code = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetCurrency(v string) {
	o.Currency = &v
}

// GetCycleLimit returns the CycleLimit field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCycleLimit() int32 {
	if o == nil || IsNil(o.CycleLimit) {
		var ret int32
		return ret
	}
	return *o.CycleLimit
}

// GetCycleLimitOk returns a tuple with the CycleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetCycleLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CycleLimit) {
		return nil, false
	}
	return o.CycleLimit, true
}

// HasCycleLimit returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasCycleLimit() bool {
	if o != nil && !IsNil(o.CycleLimit) {
		return true
	}

	return false
}

// SetCycleLimit gets a reference to the given int32 and assigns it to the CycleLimit field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetCycleLimit(v int32) {
	o.CycleLimit = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountPercentage() int64 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int64
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int64 and assigns it to the DiscountPercentage field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetDiscountPercentage(v int64) {
	o.DiscountPercentage = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountType() int32 {
	if o == nil || IsNil(o.DiscountType) {
		var ret int32
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetDiscountTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given int32 and assigns it to the DiscountType field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetDiscountType(v int32) {
	o.DiscountType = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetName(v string) {
	o.Name = &v
}

// GetPlanIds returns the PlanIds field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetPlanIds() []int64 {
	if o == nil || IsNil(o.PlanIds) {
		var ret []int64
		return ret
	}
	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetPlanIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PlanIds) {
		return nil, false
	}
	return o.PlanIds, true
}

// HasPlanIds returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasPlanIds() bool {
	if o != nil && !IsNil(o.PlanIds) {
		return true
	}

	return false
}

// SetPlanIds gets a reference to the given []int64 and assigns it to the PlanIds field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetPlanIds(v []int64) {
	o.PlanIds = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantDiscountCode) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantDiscountCode) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanMerchantDiscountCode) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiBeanMerchantDiscountCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantDiscountCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingType) {
		toSerialize["billingType"] = o.BillingType
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.CycleLimit) {
		toSerialize["cycleLimit"] = o.CycleLimit
	}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.DiscountType) {
		toSerialize["discountType"] = o.DiscountType
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanIds) {
		toSerialize["planIds"] = o.PlanIds
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantDiscountCode struct {
	value *UnibeeApiBeanMerchantDiscountCode
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantDiscountCode) Get() *UnibeeApiBeanMerchantDiscountCode {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantDiscountCode) Set(val *UnibeeApiBeanMerchantDiscountCode) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantDiscountCode) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantDiscountCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantDiscountCode(val *UnibeeApiBeanMerchantDiscountCode) *NullableUnibeeApiBeanMerchantDiscountCode {
	return &NullableUnibeeApiBeanMerchantDiscountCode{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantDiscountCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantDiscountCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


