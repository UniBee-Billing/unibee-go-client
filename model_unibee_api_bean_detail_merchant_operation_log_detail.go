/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailMerchantOperationLogDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailMerchantOperationLogDetail{}

// UnibeeApiBeanDetailMerchantOperationLogDetail struct for UnibeeApiBeanDetailMerchantOperationLogDetail
type UnibeeApiBeanDetailMerchantOperationLogDetail struct {
	// operation create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// discount_code
	DiscountCode *string `json:"discountCode,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// invoice id
	InvoiceId *string `json:"invoiceId,omitempty"`
	Member *UnibeeApiBeanDetailMerchantMemberDetail `json:"member,omitempty"`
	// member_id
	MemberId *int64 `json:"memberId,omitempty"`
	// merchant Id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// operation Account
	OptAccount *string `json:"optAccount,omitempty"`
	// operation account id
	OptAccountId *string `json:"optAccountId,omitempty"`
	// opt_account_type, 0-Member|1-User|2-OpenApi|3-System
	OptAccountType *int32 `json:"optAccountType,omitempty"`
	// operation content
	OptContent *string `json:"optContent,omitempty"`
	// operation target
	OptTarget *string `json:"optTarget,omitempty"`
	// plan id
	PlanId *int64 `json:"planId,omitempty"`
	// subscription_id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	// user_id
	UserId *int64 `json:"userId,omitempty"`
}

// NewUnibeeApiBeanDetailMerchantOperationLogDetail instantiates a new UnibeeApiBeanDetailMerchantOperationLogDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailMerchantOperationLogDetail() *UnibeeApiBeanDetailMerchantOperationLogDetail {
	this := UnibeeApiBeanDetailMerchantOperationLogDetail{}
	return &this
}

// NewUnibeeApiBeanDetailMerchantOperationLogDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantOperationLogDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailMerchantOperationLogDetailWithDefaults() *UnibeeApiBeanDetailMerchantOperationLogDetail {
	this := UnibeeApiBeanDetailMerchantOperationLogDetail{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetDiscountCode returns the DiscountCode field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetDiscountCode() string {
	if o == nil || IsNil(o.DiscountCode) {
		var ret string
		return ret
	}
	return *o.DiscountCode
}

// GetDiscountCodeOk returns a tuple with the DiscountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetDiscountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountCode) {
		return nil, false
	}
	return o.DiscountCode, true
}

// HasDiscountCode returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasDiscountCode() bool {
	if o != nil && !IsNil(o.DiscountCode) {
		return true
	}

	return false
}

// SetDiscountCode gets a reference to the given string and assigns it to the DiscountCode field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetDiscountCode(v string) {
	o.DiscountCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetId(v int64) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetMember returns the Member field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMember() UnibeeApiBeanDetailMerchantMemberDetail {
	if o == nil || IsNil(o.Member) {
		var ret UnibeeApiBeanDetailMerchantMemberDetail
		return ret
	}
	return *o.Member
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool) {
	if o == nil || IsNil(o.Member) {
		return nil, false
	}
	return o.Member, true
}

// HasMember returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMember() bool {
	if o != nil && !IsNil(o.Member) {
		return true
	}

	return false
}

// SetMember gets a reference to the given UnibeeApiBeanDetailMerchantMemberDetail and assigns it to the Member field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMember(v UnibeeApiBeanDetailMerchantMemberDetail) {
	o.Member = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberId() int64 {
	if o == nil || IsNil(o.MemberId) {
		var ret int64
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int64 and assigns it to the MemberId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMemberId(v int64) {
	o.MemberId = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetOptAccount returns the OptAccount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccount() string {
	if o == nil || IsNil(o.OptAccount) {
		var ret string
		return ret
	}
	return *o.OptAccount
}

// GetOptAccountOk returns a tuple with the OptAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountOk() (*string, bool) {
	if o == nil || IsNil(o.OptAccount) {
		return nil, false
	}
	return o.OptAccount, true
}

// HasOptAccount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccount() bool {
	if o != nil && !IsNil(o.OptAccount) {
		return true
	}

	return false
}

// SetOptAccount gets a reference to the given string and assigns it to the OptAccount field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccount(v string) {
	o.OptAccount = &v
}

// GetOptAccountId returns the OptAccountId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountId() string {
	if o == nil || IsNil(o.OptAccountId) {
		var ret string
		return ret
	}
	return *o.OptAccountId
}

// GetOptAccountIdOk returns a tuple with the OptAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptAccountId) {
		return nil, false
	}
	return o.OptAccountId, true
}

// HasOptAccountId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccountId() bool {
	if o != nil && !IsNil(o.OptAccountId) {
		return true
	}

	return false
}

// SetOptAccountId gets a reference to the given string and assigns it to the OptAccountId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccountId(v string) {
	o.OptAccountId = &v
}

// GetOptAccountType returns the OptAccountType field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountType() int32 {
	if o == nil || IsNil(o.OptAccountType) {
		var ret int32
		return ret
	}
	return *o.OptAccountType
}

// GetOptAccountTypeOk returns a tuple with the OptAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.OptAccountType) {
		return nil, false
	}
	return o.OptAccountType, true
}

// HasOptAccountType returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccountType() bool {
	if o != nil && !IsNil(o.OptAccountType) {
		return true
	}

	return false
}

// SetOptAccountType gets a reference to the given int32 and assigns it to the OptAccountType field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccountType(v int32) {
	o.OptAccountType = &v
}

// GetOptContent returns the OptContent field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptContent() string {
	if o == nil || IsNil(o.OptContent) {
		var ret string
		return ret
	}
	return *o.OptContent
}

// GetOptContentOk returns a tuple with the OptContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptContentOk() (*string, bool) {
	if o == nil || IsNil(o.OptContent) {
		return nil, false
	}
	return o.OptContent, true
}

// HasOptContent returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptContent() bool {
	if o != nil && !IsNil(o.OptContent) {
		return true
	}

	return false
}

// SetOptContent gets a reference to the given string and assigns it to the OptContent field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptContent(v string) {
	o.OptContent = &v
}

// GetOptTarget returns the OptTarget field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptTarget() string {
	if o == nil || IsNil(o.OptTarget) {
		var ret string
		return ret
	}
	return *o.OptTarget
}

// GetOptTargetOk returns a tuple with the OptTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptTargetOk() (*string, bool) {
	if o == nil || IsNil(o.OptTarget) {
		return nil, false
	}
	return o.OptTarget, true
}

// HasOptTarget returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptTarget() bool {
	if o != nil && !IsNil(o.OptTarget) {
		return true
	}

	return false
}

// SetOptTarget gets a reference to the given string and assigns it to the OptTarget field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptTarget(v string) {
	o.OptTarget = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetPlanId() int64 {
	if o == nil || IsNil(o.PlanId) {
		var ret int64
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetPlanIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given int64 and assigns it to the PlanId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetPlanId(v int64) {
	o.PlanId = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetUserId(v int64) {
	o.UserId = &v
}

func (o UnibeeApiBeanDetailMerchantOperationLogDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailMerchantOperationLogDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.DiscountCode) {
		toSerialize["discountCode"] = o.DiscountCode
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.Member) {
		toSerialize["member"] = o.Member
	}
	if !IsNil(o.MemberId) {
		toSerialize["memberId"] = o.MemberId
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.OptAccount) {
		toSerialize["optAccount"] = o.OptAccount
	}
	if !IsNil(o.OptAccountId) {
		toSerialize["optAccountId"] = o.OptAccountId
	}
	if !IsNil(o.OptAccountType) {
		toSerialize["optAccountType"] = o.OptAccountType
	}
	if !IsNil(o.OptContent) {
		toSerialize["optContent"] = o.OptContent
	}
	if !IsNil(o.OptTarget) {
		toSerialize["optTarget"] = o.OptTarget
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailMerchantOperationLogDetail struct {
	value *UnibeeApiBeanDetailMerchantOperationLogDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailMerchantOperationLogDetail) Get() *UnibeeApiBeanDetailMerchantOperationLogDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailMerchantOperationLogDetail) Set(val *UnibeeApiBeanDetailMerchantOperationLogDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailMerchantOperationLogDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailMerchantOperationLogDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailMerchantOperationLogDetail(val *UnibeeApiBeanDetailMerchantOperationLogDetail) *NullableUnibeeApiBeanDetailMerchantOperationLogDetail {
	return &NullableUnibeeApiBeanDetailMerchantOperationLogDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailMerchantOperationLogDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailMerchantOperationLogDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


