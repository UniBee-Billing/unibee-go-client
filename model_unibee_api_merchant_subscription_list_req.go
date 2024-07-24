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

// checks if the UnibeeApiMerchantSubscriptionListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantSubscriptionListReq{}

// UnibeeApiMerchantSubscriptionListReq struct for UnibeeApiMerchantSubscriptionListReq
type UnibeeApiMerchantSubscriptionListReq struct {
	// The filter end amount of subscription
	AmountEnd *int32 `json:"amountEnd,omitempty"`
	// The filter start amount of subscription
	AmountStart *int32 `json:"amountStart,omitempty"`
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// The currency of subscription
	Currency *string `json:"currency,omitempty"`
	// Page, Start With 0
	Page *int32 `json:"page,omitempty"`
	// The filter ids of plan
	PlanIds []int32 `json:"planIds,omitempty"`
	// Sort Field，gmt_create|gmt_modify，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// Filter, Default All，Status，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed
	Status []int32 `json:"status,omitempty"`
	// UserId
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantSubscriptionListReq instantiates a new UnibeeApiMerchantSubscriptionListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantSubscriptionListReq() *UnibeeApiMerchantSubscriptionListReq {
	this := UnibeeApiMerchantSubscriptionListReq{}
	return &this
}

// NewUnibeeApiMerchantSubscriptionListReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantSubscriptionListReqWithDefaults() *UnibeeApiMerchantSubscriptionListReq {
	this := UnibeeApiMerchantSubscriptionListReq{}
	return &this
}

// GetAmountEnd returns the AmountEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountEnd() int32 {
	if o == nil || IsNil(o.AmountEnd) {
		var ret int32
		return ret
	}
	return *o.AmountEnd
}

// GetAmountEndOk returns a tuple with the AmountEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountEndOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountEnd) {
		return nil, false
	}
	return o.AmountEnd, true
}

// HasAmountEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasAmountEnd() bool {
	if o != nil && !IsNil(o.AmountEnd) {
		return true
	}

	return false
}

// SetAmountEnd gets a reference to the given int32 and assigns it to the AmountEnd field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetAmountEnd(v int32) {
	o.AmountEnd = &v
}

// GetAmountStart returns the AmountStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountStart() int32 {
	if o == nil || IsNil(o.AmountStart) {
		var ret int32
		return ret
	}
	return *o.AmountStart
}

// GetAmountStartOk returns a tuple with the AmountStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetAmountStartOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountStart) {
		return nil, false
	}
	return o.AmountStart, true
}

// HasAmountStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasAmountStart() bool {
	if o != nil && !IsNil(o.AmountStart) {
		return true
	}

	return false
}

// SetAmountStart gets a reference to the given int32 and assigns it to the AmountStart field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetAmountStart(v int32) {
	o.AmountStart = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetPage(v int32) {
	o.Page = &v
}

// GetPlanIds returns the PlanIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetPlanIds() []int32 {
	if o == nil || IsNil(o.PlanIds) {
		var ret []int32
		return ret
	}
	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetPlanIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.PlanIds) {
		return nil, false
	}
	return o.PlanIds, true
}

// HasPlanIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasPlanIds() bool {
	if o != nil && !IsNil(o.PlanIds) {
		return true
	}

	return false
}

// SetPlanIds gets a reference to the given []int32 and assigns it to the PlanIds field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetPlanIds(v []int32) {
	o.PlanIds = v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetStatus(v []int32) {
	o.Status = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantSubscriptionListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantSubscriptionListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantSubscriptionListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantSubscriptionListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantSubscriptionListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountEnd) {
		toSerialize["amountEnd"] = o.AmountEnd
	}
	if !IsNil(o.AmountStart) {
		toSerialize["amountStart"] = o.AmountStart
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.CreateTimeEnd) {
		toSerialize["createTimeEnd"] = o.CreateTimeEnd
	}
	if !IsNil(o.CreateTimeStart) {
		toSerialize["createTimeStart"] = o.CreateTimeStart
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.PlanIds) {
		toSerialize["planIds"] = o.PlanIds
	}
	if !IsNil(o.SortField) {
		toSerialize["sortField"] = o.SortField
	}
	if !IsNil(o.SortType) {
		toSerialize["sortType"] = o.SortType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantSubscriptionListReq struct {
	value *UnibeeApiMerchantSubscriptionListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantSubscriptionListReq) Get() *UnibeeApiMerchantSubscriptionListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantSubscriptionListReq) Set(val *UnibeeApiMerchantSubscriptionListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantSubscriptionListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantSubscriptionListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantSubscriptionListReq(val *UnibeeApiMerchantSubscriptionListReq) *NullableUnibeeApiMerchantSubscriptionListReq {
	return &NullableUnibeeApiMerchantSubscriptionListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantSubscriptionListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantSubscriptionListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


