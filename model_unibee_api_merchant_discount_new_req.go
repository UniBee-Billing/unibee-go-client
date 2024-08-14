/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408141003 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountNewReq{}

// UnibeeApiMerchantDiscountNewReq Create a new discount code, code can used in onetime or subscription purchase to make discount
type UnibeeApiMerchantDiscountNewReq struct {
	// The billing type of the discount code, 1-one-time, 2-recurring, define the situation the code can be used, the code of one-time billing_type can used for all situation that effect only once, the code of recurring billing_tye can only used for subscription purchase
	BillingType int32 `json:"billingType"`
	// The discount's unique code, customize by merchant
	Code string `json:"code"`
	// The discount currency of discount code, available when discount_type is fixed_amount
	Currency *string `json:"currency,omitempty"`
	// The count limitation of subscription cycle, each subscription is valid separately , 0-no limit
	CycleLimit *int32 `json:"cycleLimit,omitempty"`
	// The discount amount of the discount code, available when discount_type is fixed_amount
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// The discount percentage of discount code, 100=1%, available when discount_type is percentage
	DiscountPercentage *int64 `json:"discountPercentage,omitempty"`
	// The discount type of the discount code, 1-percentage, 2-fixed_amount, the discountType of code, the discountPercentage will be effect when discountType is percentage, the discountAmount and currency will be effect when discountTYpe is fixed_amount
	DiscountType int32 `json:"discountType"`
	// The end time of discount code can effect, utc time
	EndTime int64 `json:"endTime"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// The discount's name
	Name *string `json:"name,omitempty"`
	// Ids of plan which discount code can effect, default effect all plans if not set
	PlanIds []int64 `json:"planIds,omitempty"`
	// The start time of discount code can effect, utc time
	StartTime int64 `json:"startTime"`
}

type _UnibeeApiMerchantDiscountNewReq UnibeeApiMerchantDiscountNewReq

// NewUnibeeApiMerchantDiscountNewReq instantiates a new UnibeeApiMerchantDiscountNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountNewReq(billingType int32, code string, discountType int32, endTime int64, startTime int64) *UnibeeApiMerchantDiscountNewReq {
	this := UnibeeApiMerchantDiscountNewReq{}
	this.BillingType = billingType
	this.Code = code
	this.DiscountType = discountType
	this.EndTime = endTime
	this.StartTime = startTime
	return &this
}

// NewUnibeeApiMerchantDiscountNewReqWithDefaults instantiates a new UnibeeApiMerchantDiscountNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountNewReqWithDefaults() *UnibeeApiMerchantDiscountNewReq {
	this := UnibeeApiMerchantDiscountNewReq{}
	return &this
}

// GetBillingType returns the BillingType field value
func (o *UnibeeApiMerchantDiscountNewReq) GetBillingType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetBillingTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillingType, true
}

// SetBillingType sets field value
func (o *UnibeeApiMerchantDiscountNewReq) SetBillingType(v int32) {
	o.BillingType = v
}

// GetCode returns the Code field value
func (o *UnibeeApiMerchantDiscountNewReq) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *UnibeeApiMerchantDiscountNewReq) SetCode(v string) {
	o.Code = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantDiscountNewReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetCycleLimit returns the CycleLimit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetCycleLimit() int32 {
	if o == nil || IsNil(o.CycleLimit) {
		var ret int32
		return ret
	}
	return *o.CycleLimit
}

// GetCycleLimitOk returns a tuple with the CycleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetCycleLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CycleLimit) {
		return nil, false
	}
	return o.CycleLimit, true
}

// HasCycleLimit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasCycleLimit() bool {
	if o != nil && !IsNil(o.CycleLimit) {
		return true
	}

	return false
}

// SetCycleLimit gets a reference to the given int32 and assigns it to the CycleLimit field.
func (o *UnibeeApiMerchantDiscountNewReq) SetCycleLimit(v int32) {
	o.CycleLimit = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountPercentage() int64 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int64
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int64 and assigns it to the DiscountPercentage field.
func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountPercentage(v int64) {
	o.DiscountPercentage = &v
}

// GetDiscountType returns the DiscountType field value
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetDiscountTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscountType, true
}

// SetDiscountType sets field value
func (o *UnibeeApiMerchantDiscountNewReq) SetDiscountType(v int32) {
	o.DiscountType = v
}

// GetEndTime returns the EndTime field value
func (o *UnibeeApiMerchantDiscountNewReq) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *UnibeeApiMerchantDiscountNewReq) SetEndTime(v int64) {
	o.EndTime = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantDiscountNewReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantDiscountNewReq) SetName(v string) {
	o.Name = &v
}

// GetPlanIds returns the PlanIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountNewReq) GetPlanIds() []int64 {
	if o == nil || IsNil(o.PlanIds) {
		var ret []int64
		return ret
	}
	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetPlanIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PlanIds) {
		return nil, false
	}
	return o.PlanIds, true
}

// HasPlanIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountNewReq) HasPlanIds() bool {
	if o != nil && !IsNil(o.PlanIds) {
		return true
	}

	return false
}

// SetPlanIds gets a reference to the given []int64 and assigns it to the PlanIds field.
func (o *UnibeeApiMerchantDiscountNewReq) SetPlanIds(v []int64) {
	o.PlanIds = v
}

// GetStartTime returns the StartTime field value
func (o *UnibeeApiMerchantDiscountNewReq) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountNewReq) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *UnibeeApiMerchantDiscountNewReq) SetStartTime(v int64) {
	o.StartTime = v
}

func (o UnibeeApiMerchantDiscountNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountNewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["billingType"] = o.BillingType
	toSerialize["code"] = o.Code
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
	toSerialize["discountType"] = o.DiscountType
	toSerialize["endTime"] = o.EndTime
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanIds) {
		toSerialize["planIds"] = o.PlanIds
	}
	toSerialize["startTime"] = o.StartTime
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountNewReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"billingType",
		"code",
		"discountType",
		"endTime",
		"startTime",
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

	varUnibeeApiMerchantDiscountNewReq := _UnibeeApiMerchantDiscountNewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountNewReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountNewReq(varUnibeeApiMerchantDiscountNewReq)

	return err
}

type NullableUnibeeApiMerchantDiscountNewReq struct {
	value *UnibeeApiMerchantDiscountNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountNewReq) Get() *UnibeeApiMerchantDiscountNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountNewReq) Set(val *UnibeeApiMerchantDiscountNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountNewReq(val *UnibeeApiMerchantDiscountNewReq) *NullableUnibeeApiMerchantDiscountNewReq {
	return &NullableUnibeeApiMerchantDiscountNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


