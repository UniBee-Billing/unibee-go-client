/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantDiscountEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantDiscountEditReq{}

// UnibeeApiMerchantDiscountEditReq Edit the discount code before activate
type UnibeeApiMerchantDiscountEditReq struct {
	// The billing type of the discount code, 1-one-time, 2-recurring, define the situation the code can be used, the code of one-time billing_type can used for all situation that effect only once, the code of recurring billing_tye can only used for subscription purchase
	BillingType *int32 `json:"billingType,omitempty"`
	// The discount currency of discount code, available when discount_type is fixed_amount
	Currency *string `json:"currency,omitempty"`
	// The count limitation of subscription cycle，each subscription is valid separately, 0-no limit
	CycleLimit *int32 `json:"cycleLimit,omitempty"`
	// The discount amount of the discount code, available when discount_type is fixed_amount
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// The discount percentage of discount code, 100=1%, available when discount_type is percentage
	DiscountPercentage *int64 `json:"discountPercentage,omitempty"`
	// The discount type of the discount code, 1-percentage, 2-fixed_amount, the discountType of code, the discountPercentage will be effect when discountType is percentage, the discountAmount and currency will be effect when discountTYpe is fixed_amount
	DiscountType *int32 `json:"discountType,omitempty"`
	// The end time of discount code can effect, utc time
	EndTime *int64 `json:"endTime,omitempty"`
	// The discount's Id
	Id int64 `json:"id"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// The discount's name
	Name *string `json:"name,omitempty"`
	// Ids of plan which discount code can effect, default effect all plans if not set
	PlanIds []int64 `json:"planIds,omitempty"`
	// The start time of discount code can effect, utc time
	StartTime *int64 `json:"startTime,omitempty"`
}

type _UnibeeApiMerchantDiscountEditReq UnibeeApiMerchantDiscountEditReq

// NewUnibeeApiMerchantDiscountEditReq instantiates a new UnibeeApiMerchantDiscountEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantDiscountEditReq(id int64) *UnibeeApiMerchantDiscountEditReq {
	this := UnibeeApiMerchantDiscountEditReq{}
	this.Id = id
	return &this
}

// NewUnibeeApiMerchantDiscountEditReqWithDefaults instantiates a new UnibeeApiMerchantDiscountEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantDiscountEditReqWithDefaults() *UnibeeApiMerchantDiscountEditReq {
	this := UnibeeApiMerchantDiscountEditReq{}
	return &this
}

// GetBillingType returns the BillingType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetBillingType() int32 {
	if o == nil || IsNil(o.BillingType) {
		var ret int32
		return ret
	}
	return *o.BillingType
}

// GetBillingTypeOk returns a tuple with the BillingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetBillingTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.BillingType) {
		return nil, false
	}
	return o.BillingType, true
}

// HasBillingType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasBillingType() bool {
	if o != nil && !IsNil(o.BillingType) {
		return true
	}

	return false
}

// SetBillingType gets a reference to the given int32 and assigns it to the BillingType field.
func (o *UnibeeApiMerchantDiscountEditReq) SetBillingType(v int32) {
	o.BillingType = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UnibeeApiMerchantDiscountEditReq) SetCurrency(v string) {
	o.Currency = &v
}

// GetCycleLimit returns the CycleLimit field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetCycleLimit() int32 {
	if o == nil || IsNil(o.CycleLimit) {
		var ret int32
		return ret
	}
	return *o.CycleLimit
}

// GetCycleLimitOk returns a tuple with the CycleLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetCycleLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.CycleLimit) {
		return nil, false
	}
	return o.CycleLimit, true
}

// HasCycleLimit returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasCycleLimit() bool {
	if o != nil && !IsNil(o.CycleLimit) {
		return true
	}

	return false
}

// SetCycleLimit gets a reference to the given int32 and assigns it to the CycleLimit field.
func (o *UnibeeApiMerchantDiscountEditReq) SetCycleLimit(v int32) {
	o.CycleLimit = &v
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountAmount() int64 {
	if o == nil || IsNil(o.DiscountAmount) {
		var ret int64
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given int64 and assigns it to the DiscountAmount field.
func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountAmount(v int64) {
	o.DiscountAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountPercentage() int64 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret int64
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountPercentageOk() (*int64, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given int64 and assigns it to the DiscountPercentage field.
func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountPercentage(v int64) {
	o.DiscountPercentage = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountType() int32 {
	if o == nil || IsNil(o.DiscountType) {
		var ret int32
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given int32 and assigns it to the DiscountType field.
func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountType(v int32) {
	o.DiscountType = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *UnibeeApiMerchantDiscountEditReq) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetId returns the Id field value
func (o *UnibeeApiMerchantDiscountEditReq) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UnibeeApiMerchantDiscountEditReq) SetId(v int64) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantDiscountEditReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiMerchantDiscountEditReq) SetName(v string) {
	o.Name = &v
}

// GetPlanIds returns the PlanIds field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetPlanIds() []int64 {
	if o == nil || IsNil(o.PlanIds) {
		var ret []int64
		return ret
	}
	return o.PlanIds
}

// GetPlanIdsOk returns a tuple with the PlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetPlanIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PlanIds) {
		return nil, false
	}
	return o.PlanIds, true
}

// HasPlanIds returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasPlanIds() bool {
	if o != nil && !IsNil(o.PlanIds) {
		return true
	}

	return false
}

// SetPlanIds gets a reference to the given []int64 and assigns it to the PlanIds field.
func (o *UnibeeApiMerchantDiscountEditReq) SetPlanIds(v []int64) {
	o.PlanIds = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UnibeeApiMerchantDiscountEditReq) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantDiscountEditReq) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UnibeeApiMerchantDiscountEditReq) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *UnibeeApiMerchantDiscountEditReq) SetStartTime(v int64) {
	o.StartTime = &v
}

func (o UnibeeApiMerchantDiscountEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantDiscountEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingType) {
		toSerialize["billingType"] = o.BillingType
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
	toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

func (o *UnibeeApiMerchantDiscountEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varUnibeeApiMerchantDiscountEditReq := _UnibeeApiMerchantDiscountEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantDiscountEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantDiscountEditReq(varUnibeeApiMerchantDiscountEditReq)

	return err
}

type NullableUnibeeApiMerchantDiscountEditReq struct {
	value *UnibeeApiMerchantDiscountEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantDiscountEditReq) Get() *UnibeeApiMerchantDiscountEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantDiscountEditReq) Set(val *UnibeeApiMerchantDiscountEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantDiscountEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantDiscountEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantDiscountEditReq(val *UnibeeApiMerchantDiscountEditReq) *NullableUnibeeApiMerchantDiscountEditReq {
	return &NullableUnibeeApiMerchantDiscountEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantDiscountEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantDiscountEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


