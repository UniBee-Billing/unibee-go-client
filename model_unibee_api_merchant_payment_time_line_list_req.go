/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPaymentTimeLineListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPaymentTimeLineListReq{}

// UnibeeApiMerchantPaymentTimeLineListReq struct for UnibeeApiMerchantPaymentTimeLineListReq
type UnibeeApiMerchantPaymentTimeLineListReq struct {
	// The filter end amount of timeline
	AmountEnd *int32 `json:"amountEnd,omitempty"`
	// The filter start amount of timeline
	AmountStart *int32 `json:"amountStart,omitempty"`
	// Count Of Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// Currency
	Currency *string `json:"currency,omitempty"`
	// The filter ids of gateway
	GatewayIds []int64 `json:"gatewayIds,omitempty"`
	// Page,Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort，invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// The filter status, 0-pending, 1-success, 2-failure，3-cancel
	Status []int32 `json:"status,omitempty"`
	// The filter timelineType, 0-pay, 1-refund
	TimelineTypes []int32 `json:"timelineTypes,omitempty"`
	// Filter UserId, Default All
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantPaymentTimeLineListReq instantiates a new UnibeeApiMerchantPaymentTimeLineListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPaymentTimeLineListReq() *UnibeeApiMerchantPaymentTimeLineListReq {
	this := UnibeeApiMerchantPaymentTimeLineListReq{}
	return &this
}

// NewUnibeeApiMerchantPaymentTimeLineListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentTimeLineListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPaymentTimeLineListReqWithDefaults() *UnibeeApiMerchantPaymentTimeLineListReq {
	this := UnibeeApiMerchantPaymentTimeLineListReq{}
	return &this
}

// GetAmountEnd returns the AmountEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountEnd() int32 {
	if o == nil || IsNil(o.AmountEnd) {
		var ret int32
		return ret
	}
	return *o.AmountEnd
}

// GetAmountEndOk returns a tuple with the AmountEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountEndOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountEnd) {
		return nil, false
	}
	return o.AmountEnd, true
}

// HasAmountEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasAmountEnd() bool {
	if o != nil && !IsNil(o.AmountEnd) {
		return true
	}

	return false
}

// SetAmountEnd gets a reference to the given int32 and assigns it to the AmountEnd field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetAmountEnd(v int32) {
	o.AmountEnd = &v
}

// GetAmountStart returns the AmountStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountStart() int32 {
	if o == nil || IsNil(o.AmountStart) {
		var ret int32
		return ret
	}
	return *o.AmountStart
}

// GetAmountStartOk returns a tuple with the AmountStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetAmountStartOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountStart) {
		return nil, false
	}
	return o.AmountStart, true
}

// HasAmountStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasAmountStart() bool {
	if o != nil && !IsNil(o.AmountStart) {
		return true
	}

	return false
}

// SetAmountStart gets a reference to the given int32 and assigns it to the AmountStart field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetAmountStart(v int32) {
	o.AmountStart = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetGatewayIds returns the GatewayIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetGatewayIds() []int64 {
	if o == nil || IsNil(o.GatewayIds) {
		var ret []int64
		return ret
	}
	return o.GatewayIds
}

// GetGatewayIdsOk returns a tuple with the GatewayIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetGatewayIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.GatewayIds) {
		return nil, false
	}
	return o.GatewayIds, true
}

// HasGatewayIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasGatewayIds() bool {
	if o != nil && !IsNil(o.GatewayIds) {
		return true
	}

	return false
}

// SetGatewayIds gets a reference to the given []int64 and assigns it to the GatewayIds field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetGatewayIds(v []int64) {
	o.GatewayIds = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetStatus(v []int32) {
	o.Status = v
}

// GetTimelineTypes returns the TimelineTypes field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetTimelineTypes() []int32 {
	if o == nil || IsNil(o.TimelineTypes) {
		var ret []int32
		return ret
	}
	return o.TimelineTypes
}

// GetTimelineTypesOk returns a tuple with the TimelineTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetTimelineTypesOk() ([]int32, bool) {
	if o == nil || IsNil(o.TimelineTypes) {
		return nil, false
	}
	return o.TimelineTypes, true
}

// HasTimelineTypes returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasTimelineTypes() bool {
	if o != nil && !IsNil(o.TimelineTypes) {
		return true
	}

	return false
}

// SetTimelineTypes gets a reference to the given []int32 and assigns it to the TimelineTypes field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetTimelineTypes(v []int32) {
	o.TimelineTypes = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantPaymentTimeLineListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantPaymentTimeLineListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPaymentTimeLineListReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.GatewayIds) {
		toSerialize["gatewayIds"] = o.GatewayIds
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
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
	if !IsNil(o.TimelineTypes) {
		toSerialize["timelineTypes"] = o.TimelineTypes
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPaymentTimeLineListReq struct {
	value *UnibeeApiMerchantPaymentTimeLineListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListReq) Get() *UnibeeApiMerchantPaymentTimeLineListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListReq) Set(val *UnibeeApiMerchantPaymentTimeLineListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPaymentTimeLineListReq(val *UnibeeApiMerchantPaymentTimeLineListReq) *NullableUnibeeApiMerchantPaymentTimeLineListReq {
	return &NullableUnibeeApiMerchantPaymentTimeLineListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPaymentTimeLineListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPaymentTimeLineListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


