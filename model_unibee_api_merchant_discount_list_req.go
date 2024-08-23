/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408231424 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantDiscountListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountListReq{}

// UnibeeApiMerchantDiscountListReq Get discountCode list
type UnibeeApiMerchantDiscountListReq struct {
	// billing_type, 1-one-time, 2-recurring
	BillingType []int32 `json:"billingType,omitempty"`
	// Filter Code
	Code *string `json:"code,omitempty"`
	// Count Of Per Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// Filter Currency
	Currency *string `json:"currency,omitempty"`
	// discount_type, 1-percentage, 2-fixed_amount
	DiscountType []int32 `json:"discountType,omitempty"`
	// Page, Start 0
	Page *int32 `json:"page,omitempty"`
	// Sort Field，gmt_create|gmt_modify，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort Type，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// status, 1-editable, 2-active, 3-deactive, 4-expire
	Status []int32 `json:"status,omitempty"`
}

// NewUnibeeApiMerchantDiscountListReq instantiates a new UnibeeApiMerchantDiscountListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountListReq() *UnibeeApiMerchantDiscountListReq {
	this := UnibeeApiMerchantDiscountListReq{}
	return &this
}

// NewUnibeeApiMerchantDiscountListReqWithDefaults instantiates a new UnibeeApiMerchantDiscountListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountListReqWithDefaults() *UnibeeApiMerchantDiscountListReq {
	this := UnibeeApiMerchantDiscountListReq{}
	return &this
}

// GetBillingType returns the BillingType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetBillingType() []int32 {
	if o == nil || IsNil(o.BillingType) {
		var ret []int32
		return ret
	}
	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetBillingTypeOk() ([]int32, bool) {
	if o == nil || IsNil(o.BillingType) {
		return nil, false
	}
	return o.BillingType, true
}

// HasBillingType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasBillingType() bool {
	if o != nil && !IsNil(o.BillingType) {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given []int32 and assigns it to the BillingType field.
func (o *UnibeeApiMerchantDiscountListReq) SetBillingType(v []int32) {
	o.BillingType = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UnibeeApiMerchantDiscountListReq) SetCode(v string) {
	o.Code = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantDiscountListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantDiscountListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantDiscountListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantDiscountListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetDiscountType() []int32 {
	if o == nil || IsNil(o.DiscountType) {
		var ret []int32
		return ret
	}
	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetDiscountTypeOk() ([]int32, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given []int32 and assigns it to the DiscountType field.
func (o *UnibeeApiMerchantDiscountListReq) SetDiscountType(v []int32) {
	o.DiscountType = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantDiscountListReq) SetPage(v int32) {
	o.Page = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantDiscountListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantDiscountListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantDiscountListReq) SetStatus(v []int32) {
	o.Status = v
}

func (o UnibeeApiMerchantDiscountListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingType) {
		toSerialize["billingType"] = o.BillingType
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
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
	if !IsNil(o.DiscountType) {
		toSerialize["discountType"] = o.DiscountType
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
	return toSerialize, nil
}

type NullableUnibeeApiMerchantDiscountListReq struct {
	value *UnibeeApiMerchantDiscountListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountListReq) Get() *UnibeeApiMerchantDiscountListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountListReq) Set(val *UnibeeApiMerchantDiscountListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountListReq(val *UnibeeApiMerchantDiscountListReq) *NullableUnibeeApiMerchantDiscountListReq {
	return &NullableUnibeeApiMerchantDiscountListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


