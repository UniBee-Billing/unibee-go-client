/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408060614 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantInvoiceListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantInvoiceListReq{}

// UnibeeApiMerchantInvoiceListReq Get invoice list
type UnibeeApiMerchantInvoiceListReq struct {
	// The filter end amount of invoice
	AmountEnd *int32 `json:"amountEnd,omitempty"`
	// The filter start amount of invoice
	AmountStart *int32 `json:"amountStart,omitempty"`
	// Count By Page
	Count *int32 `json:"count,omitempty"`
	// CreateTimeEnd
	CreateTimeEnd *int64 `json:"createTimeEnd,omitempty"`
	// CreateTimeStart
	CreateTimeStart *int64 `json:"createTimeStart,omitempty"`
	// The currency of invoice
	Currency *string `json:"currency,omitempty"`
	// Deleted Involved，Need Admin Permission
	DeleteInclude *bool `json:"deleteInclude,omitempty"`
	// The firstName of invoice
	FirstName *string `json:"firstName,omitempty"`
	// The lastName of invoice
	LastName *string `json:"lastName,omitempty"`
	// Page, Start 0
	Page *int32 `json:"page,omitempty"`
	// ReportTimeEnd
	ReportTimeEnd *int64 `json:"reportTimeEnd,omitempty"`
	// ReportTimeStart
	ReportTimeStart *int64 `json:"reportTimeStart,omitempty"`
	// The filter email of invoice
	SendEmail *string `json:"sendEmail,omitempty"`
	// Filter，em. invoice_id|gmt_create|gmt_modify|period_end|total_amount，Default gmt_modify
	SortField *string `json:"sortField,omitempty"`
	// Sort，asc|desc，Default desc
	SortType *string `json:"sortType,omitempty"`
	// The status of invoice, 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled
	Status []int32 `json:"status,omitempty"`
	// invoice Type, 0-payment, 1-refund
	Type *int32 `json:"type,omitempty"`
	// The filter userid of invoice
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiMerchantInvoiceListReq instantiates a new UnibeeApiMerchantInvoiceListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantInvoiceListReq() *UnibeeApiMerchantInvoiceListReq {
	this := UnibeeApiMerchantInvoiceListReq{}
	return &this
}

// NewUnibeeApiMerchantInvoiceListReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantInvoiceListReqWithDefaults() *UnibeeApiMerchantInvoiceListReq {
	this := UnibeeApiMerchantInvoiceListReq{}
	return &this
}

// GetAmountEnd returns the AmountEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetAmountEnd() int32 {
	if o == nil || IsNil(o.AmountEnd) {
		var ret int32
		return ret
	}
	return *o.AmountEnd
}

// GetAmountEndOk returns a tuple with the AmountEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetAmountEndOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountEnd) {
		return nil, false
	}
	return o.AmountEnd, true
}

// HasAmountEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasAmountEnd() bool {
	if o != nil && !IsNil(o.AmountEnd) {
		return true
	}

	return false
}

// SetAmountEnd gets a reference to the given int32 and assigns it to the AmountEnd field.
func (o *UnibeeApiMerchantInvoiceListReq) SetAmountEnd(v int32) {
	o.AmountEnd = &v
}

// GetAmountStart returns the AmountStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetAmountStart() int32 {
	if o == nil || IsNil(o.AmountStart) {
		var ret int32
		return ret
	}
	return *o.AmountStart
}

// GetAmountStartOk returns a tuple with the AmountStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetAmountStartOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountStart) {
		return nil, false
	}
	return o.AmountStart, true
}

// HasAmountStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasAmountStart() bool {
	if o != nil && !IsNil(o.AmountStart) {
		return true
	}

	return false
}

// SetAmountStart gets a reference to the given int32 and assigns it to the AmountStart field.
func (o *UnibeeApiMerchantInvoiceListReq) SetAmountStart(v int32) {
	o.AmountStart = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UnibeeApiMerchantInvoiceListReq) SetCount(v int32) {
	o.Count = &v
}

// GetCreateTimeEnd returns the CreateTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeEnd() int64 {
	if o == nil || IsNil(o.CreateTimeEnd) {
		var ret int64
		return ret
	}
	return *o.CreateTimeEnd
}

// GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeEnd) {
		return nil, false
	}
	return o.CreateTimeEnd, true
}

// HasCreateTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasCreateTimeEnd() bool {
	if o != nil && !IsNil(o.CreateTimeEnd) {
		return true
	}

	return false
}

// SetCreateTimeEnd gets a reference to the given int64 and assigns it to the CreateTimeEnd field.
func (o *UnibeeApiMerchantInvoiceListReq) SetCreateTimeEnd(v int64) {
	o.CreateTimeEnd = &v
}

// GetCreateTimeStart returns the CreateTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeStart() int64 {
	if o == nil || IsNil(o.CreateTimeStart) {
		var ret int64
		return ret
	}
	return *o.CreateTimeStart
}

// GetCreateTimeStartOk returns a tuple with the CreateTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetCreateTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTimeStart) {
		return nil, false
	}
	return o.CreateTimeStart, true
}

// HasCreateTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasCreateTimeStart() bool {
	if o != nil && !IsNil(o.CreateTimeStart) {
		return true
	}

	return false
}

// SetCreateTimeStart gets a reference to the given int64 and assigns it to the CreateTimeStart field.
func (o *UnibeeApiMerchantInvoiceListReq) SetCreateTimeStart(v int64) {
	o.CreateTimeStart = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantInvoiceListReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetDeleteInclude returns the DeleteInclude field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetDeleteInclude() bool {
	if o == nil || IsNil(o.DeleteInclude) {
		var ret bool
		return ret
	}
	return *o.DeleteInclude
}

// GetDeleteIncludeOk returns a tuple with the DeleteInclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetDeleteIncludeOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteInclude) {
		return nil, false
	}
	return o.DeleteInclude, true
}

// HasDeleteInclude returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasDeleteInclude() bool {
	if o != nil && !IsNil(o.DeleteInclude) {
		return true
	}

	return false
}

// SetDeleteInclude gets a reference to the given bool and assigns it to the DeleteInclude field.
func (o *UnibeeApiMerchantInvoiceListReq) SetDeleteInclude(v bool) {
	o.DeleteInclude = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UnibeeApiMerchantInvoiceListReq) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UnibeeApiMerchantInvoiceListReq) SetLastName(v string) {
	o.LastName = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *UnibeeApiMerchantInvoiceListReq) SetPage(v int32) {
	o.Page = &v
}

// GetReportTimeEnd returns the ReportTimeEnd field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeEnd() int64 {
	if o == nil || IsNil(o.ReportTimeEnd) {
		var ret int64
		return ret
	}
	return *o.ReportTimeEnd
}

// GetReportTimeEndOk returns a tuple with the ReportTimeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeEndOk() (*int64, bool) {
	if o == nil || IsNil(o.ReportTimeEnd) {
		return nil, false
	}
	return o.ReportTimeEnd, true
}

// HasReportTimeEnd returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasReportTimeEnd() bool {
	if o != nil && !IsNil(o.ReportTimeEnd) {
		return true
	}

	return false
}

// SetReportTimeEnd gets a reference to the given int64 and assigns it to the ReportTimeEnd field.
func (o *UnibeeApiMerchantInvoiceListReq) SetReportTimeEnd(v int64) {
	o.ReportTimeEnd = &v
}

// GetReportTimeStart returns the ReportTimeStart field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeStart() int64 {
	if o == nil || IsNil(o.ReportTimeStart) {
		var ret int64
		return ret
	}
	return *o.ReportTimeStart
}

// GetReportTimeStartOk returns a tuple with the ReportTimeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetReportTimeStartOk() (*int64, bool) {
	if o == nil || IsNil(o.ReportTimeStart) {
		return nil, false
	}
	return o.ReportTimeStart, true
}

// HasReportTimeStart returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasReportTimeStart() bool {
	if o != nil && !IsNil(o.ReportTimeStart) {
		return true
	}

	return false
}

// SetReportTimeStart gets a reference to the given int64 and assigns it to the ReportTimeStart field.
func (o *UnibeeApiMerchantInvoiceListReq) SetReportTimeStart(v int64) {
	o.ReportTimeStart = &v
}

// GetSendEmail returns the SendEmail field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetSendEmail() string {
	if o == nil || IsNil(o.SendEmail) {
		var ret string
		return ret
	}
	return *o.SendEmail
}

// GetSendEmailOk returns a tuple with the SendEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetSendEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SendEmail) {
		return nil, false
	}
	return o.SendEmail, true
}

// HasSendEmail returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasSendEmail() bool {
	if o != nil && !IsNil(o.SendEmail) {
		return true
	}

	return false
}

// SetSendEmail gets a reference to the given string and assigns it to the SendEmail field.
func (o *UnibeeApiMerchantInvoiceListReq) SetSendEmail(v string) {
	o.SendEmail = &v
}

// GetSortField returns the SortField field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetSortField() string {
	if o == nil || IsNil(o.SortField) {
		var ret string
		return ret
	}
	return *o.SortField
}

// GetSortFieldOk returns a tuple with the SortField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetSortFieldOk() (*string, bool) {
	if o == nil || IsNil(o.SortField) {
		return nil, false
	}
	return o.SortField, true
}

// HasSortField returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasSortField() bool {
	if o != nil && !IsNil(o.SortField) {
		return true
	}

	return false
}

// SetSortField gets a reference to the given string and assigns it to the SortField field.
func (o *UnibeeApiMerchantInvoiceListReq) SetSortField(v string) {
	o.SortField = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *UnibeeApiMerchantInvoiceListReq) SetSortType(v string) {
	o.SortType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetStatus() []int32 {
	if o == nil || IsNil(o.Status) {
		var ret []int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetStatusOk() ([]int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantInvoiceListReq) SetStatus(v []int32) {
	o.Status = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *UnibeeApiMerchantInvoiceListReq) SetType(v int32) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantInvoiceListReq) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantInvoiceListReq) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantInvoiceListReq) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiMerchantInvoiceListReq) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiMerchantInvoiceListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantInvoiceListReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DeleteInclude) {
		toSerialize["deleteInclude"] = o.DeleteInclude
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.ReportTimeEnd) {
		toSerialize["reportTimeEnd"] = o.ReportTimeEnd
	}
	if !IsNil(o.ReportTimeStart) {
		toSerialize["reportTimeStart"] = o.ReportTimeStart
	}
	if !IsNil(o.SendEmail) {
		toSerialize["sendEmail"] = o.SendEmail
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
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantInvoiceListReq struct {
	value *UnibeeApiMerchantInvoiceListReq
	isSet bool
}

func (v NullableUnibeeApiMerchantInvoiceListReq) Get() *UnibeeApiMerchantInvoiceListReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantInvoiceListReq) Set(val *UnibeeApiMerchantInvoiceListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantInvoiceListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantInvoiceListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantInvoiceListReq(val *UnibeeApiMerchantInvoiceListReq) *NullableUnibeeApiMerchantInvoiceListReq {
	return &NullableUnibeeApiMerchantInvoiceListReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantInvoiceListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantInvoiceListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


