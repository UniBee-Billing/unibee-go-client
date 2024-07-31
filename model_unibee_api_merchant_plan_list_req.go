/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202407311159 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantPlanListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantPlanListReq{}

// UnibeeApiMerchantPlanListReq struct for UnibeeApiMerchantPlanListReq
type UnibeeApiMerchantPlanListReq struct {
	// Count Of Per Page
	Count *int32 `json:"count,omitempty"`
	// Filter Currency
	Currency *string `json:"currency,omitempty"`
	// Page, Start 0
	Page *int32 `json:"page,omitempty"`
	// filter id list of product, default all
	ProductIds []int64 `json:"productIds,omitempty"`
	// Filter, Default All，PublishStatus，1-UnPublished，2-Published
	PublishStatus *int32 `json:"publishStatus,omitempty"`
	// Sort Field，gmt_create|gmt_modify，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// Filter, Default All，,Status，1-Editing，2-Active，3-InActive，4-Expired
	Status []int32 `json:"status,omitempty"`
	// 1-main plan，2-addon plan
	Type []int32 `json:"type,omitempty"`
}

// NewUnibeeApiMerchantPlanListReq instantiates a new UnibeeApiMerchantPlanListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantPlanListReq() *UnibeeApiMerchantPlanListReq {
	this := UnibeeApiMerchantPlanListReq{}
	return &this
}

// NewUnibeeApiMerchantPlanListReqWithDefaults instantiates a new UnibeeApiMerchantPlanListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantPlanListReqWithDefaults() *UnibeeApiMerchantPlanListReq {
	this := UnibeeApiMerchantPlanListReq{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantPlanListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantPlanListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantPlanListReq) SetPage(v int32) {
	o.Page = &v
}

// GetProductIds returns the ProductIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetProductIds() []int64 {
	if o == nil || IsNil(o.ProductIds) {
		var ret []int64
		return ret
	}
	return o.ProductIds
}

// GetProductIdsOk returns a tuple with the ProductIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetProductIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ProductIds) {
		return nil, false
	}
	return o.ProductIds, true
}

// HasProductIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasProductIds() bool {
	if o != nil && !IsNil(o.ProductIds) {
		return true
	}

	return false
}

// SetProductIds gets a reference to the given []int64 and assigns it to the ProductIds field.
func (o *UnibeeApiMerchantPlanListReq) SetProductIds(v []int64) {
	o.ProductIds = v
}

// GetPublishStatus returns the PublishStatus field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetPublishStatus() int32 {
	if o == nil || IsNil(o.PublishStatus) {
		var ret int32
		return ret
	}
	return *o.PublishStatus
}

// GetPublishStatusOk returns a tuple with the PublishStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetPublishStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.PublishStatus) {
		return nil, false
	}
	return o.PublishStatus, true
}

// HasPublishStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasPublishStatus() bool {
	if o != nil && !IsNil(o.PublishStatus) {
		return true
	}

	return false
}

// SetPublishStatus gets a reference to the given int32 and assigns it to the PublishStatus field.
func (o *UnibeeApiMerchantPlanListReq) SetPublishStatus(v int32) {
	o.PublishStatus = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantPlanListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantPlanListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantPlanListReq) SetStatus(v []int32) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantPlanListReq) GetType() []int32 {
	if o == nil || IsNil(o.Type) {
		var ret []int32
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantPlanListReq) GetTypeOk() ([]int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantPlanListReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []int32 and assigns it to the Type field.
func (o *UnibeeApiMerchantPlanListReq) SetType(v []int32) {
	o.Type = v
}

func (o UnibeeApiMerchantPlanListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantPlanListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.ProductIds) {
		toSerialize["productIds"] = o.ProductIds
	}
	if !IsNil(o.PublishStatus) {
		toSerialize["publishStatus"] = o.PublishStatus
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantPlanListReq struct {
	value *UnibeeApiMerchantPlanListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantPlanListReq) Get() *UnibeeApiMerchantPlanListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantPlanListReq) Set(val *UnibeeApiMerchantPlanListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantPlanListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantPlanListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantPlanListReq(val *UnibeeApiMerchantPlanListReq) *NullableUnibeeApiMerchantPlanListReq {
	return &NullableUnibeeApiMerchantPlanListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantPlanListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantPlanListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


