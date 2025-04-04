/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202504011414
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiBeanUserMetricChargeInvoiceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanUserMetricChargeInvoiceItem{}

// UnibeeApiBeanUserMetricChargeInvoiceItem struct for UnibeeApiBeanUserMetricChargeInvoiceItem
type UnibeeApiBeanUserMetricChargeInvoiceItem struct {
	// CurrentUsedValue
	CurrentUsedValue *int64 `json:"CurrentUsedValue,omitempty"`
	ChargePricing *UnibeeApiBeanPlanMetricMeteredChargeParam `json:"chargePricing,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
	MaxEventId *int64 `json:"maxEventId,omitempty"`
	// MetricId
	MetricId int64 `json:"metricId"`
	MinEventId *int64 `json:"minEventId,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// TotalChargeAmount
	TotalChargeAmount *int64 `json:"totalChargeAmount,omitempty"`
}

type _UnibeeApiBeanUserMetricChargeInvoiceItem UnibeeApiBeanUserMetricChargeInvoiceItem

// NewUnibeeApiBeanUserMetricChargeInvoiceItem instantiates a new UnibeeApiBeanUserMetricChargeInvoiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanUserMetricChargeInvoiceItem(metricId int64) *UnibeeApiBeanUserMetricChargeInvoiceItem {
	this := UnibeeApiBeanUserMetricChargeInvoiceItem{}
	this.MetricId = metricId
	return &this
}

// NewUnibeeApiBeanUserMetricChargeInvoiceItemWithDefaults instantiates a new UnibeeApiBeanUserMetricChargeInvoiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanUserMetricChargeInvoiceItemWithDefaults() *UnibeeApiBeanUserMetricChargeInvoiceItem {
	this := UnibeeApiBeanUserMetricChargeInvoiceItem{}
	return &this
}

// GetCurrentUsedValue returns the CurrentUsedValue field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetCurrentUsedValue() int64 {
	if o == nil || IsNil(o.CurrentUsedValue) {
		var ret int64
		return ret
	}
	return *o.CurrentUsedValue
}

// GetCurrentUsedValueOk returns a tuple with the CurrentUsedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetCurrentUsedValueOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentUsedValue) {
		return nil, false
	}
	return o.CurrentUsedValue, true
}

// HasCurrentUsedValue returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasCurrentUsedValue() bool {
	if o != nil && !IsNil(o.CurrentUsedValue) {
		return true
	}

	return false
}

// SetCurrentUsedValue gets a reference to the given int64 and assigns it to the CurrentUsedValue field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetCurrentUsedValue(v int64) {
	o.CurrentUsedValue = &v
}

// GetChargePricing returns the ChargePricing field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetChargePricing() UnibeeApiBeanPlanMetricMeteredChargeParam {
	if o == nil || IsNil(o.ChargePricing) {
		var ret UnibeeApiBeanPlanMetricMeteredChargeParam
		return ret
	}
	return *o.ChargePricing
}

// GetChargePricingOk returns a tuple with the ChargePricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetChargePricingOk() (*UnibeeApiBeanPlanMetricMeteredChargeParam, bool) {
	if o == nil || IsNil(o.ChargePricing) {
		return nil, false
	}
	return o.ChargePricing, true
}

// HasChargePricing returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasChargePricing() bool {
	if o != nil && !IsNil(o.ChargePricing) {
		return true
	}

	return false
}

// SetChargePricing gets a reference to the given UnibeeApiBeanPlanMetricMeteredChargeParam and assigns it to the ChargePricing field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetChargePricing(v UnibeeApiBeanPlanMetricMeteredChargeParam) {
	o.ChargePricing = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetDescription(v string) {
	o.Description = &v
}

// GetMaxEventId returns the MaxEventId field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMaxEventId() int64 {
	if o == nil || IsNil(o.MaxEventId) {
		var ret int64
		return ret
	}
	return *o.MaxEventId
}

// GetMaxEventIdOk returns a tuple with the MaxEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMaxEventIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxEventId) {
		return nil, false
	}
	return o.MaxEventId, true
}

// HasMaxEventId returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasMaxEventId() bool {
	if o != nil && !IsNil(o.MaxEventId) {
		return true
	}

	return false
}

// SetMaxEventId gets a reference to the given int64 and assigns it to the MaxEventId field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMaxEventId(v int64) {
	o.MaxEventId = &v
}

// GetMetricId returns the MetricId field value
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMetricId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMetricIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMetricId(v int64) {
	o.MetricId = v
}

// GetMinEventId returns the MinEventId field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMinEventId() int64 {
	if o == nil || IsNil(o.MinEventId) {
		var ret int64
		return ret
	}
	return *o.MinEventId
}

// GetMinEventIdOk returns a tuple with the MinEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetMinEventIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MinEventId) {
		return nil, false
	}
	return o.MinEventId, true
}

// HasMinEventId returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasMinEventId() bool {
	if o != nil && !IsNil(o.MinEventId) {
		return true
	}

	return false
}

// SetMinEventId gets a reference to the given int64 and assigns it to the MinEventId field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetMinEventId(v int64) {
	o.MinEventId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetName(v string) {
	o.Name = &v
}

// GetTotalChargeAmount returns the TotalChargeAmount field value if set, zero value otherwise.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetTotalChargeAmount() int64 {
	if o == nil || IsNil(o.TotalChargeAmount) {
		var ret int64
		return ret
	}
	return *o.TotalChargeAmount
}

// GetTotalChargeAmountOk returns a tuple with the TotalChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) GetTotalChargeAmountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalChargeAmount) {
		return nil, false
	}
	return o.TotalChargeAmount, true
}

// HasTotalChargeAmount returns a boolean if a field has been set.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) HasTotalChargeAmount() bool {
	if o != nil && !IsNil(o.TotalChargeAmount) {
		return true
	}

	return false
}

// SetTotalChargeAmount gets a reference to the given int64 and assigns it to the TotalChargeAmount field.
func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) SetTotalChargeAmount(v int64) {
	o.TotalChargeAmount = &v
}

func (o UnibeeApiBeanUserMetricChargeInvoiceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanUserMetricChargeInvoiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentUsedValue) {
		toSerialize["CurrentUsedValue"] = o.CurrentUsedValue
	}
	if !IsNil(o.ChargePricing) {
		toSerialize["chargePricing"] = o.ChargePricing
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MaxEventId) {
		toSerialize["maxEventId"] = o.MaxEventId
	}
	toSerialize["metricId"] = o.MetricId
	if !IsNil(o.MinEventId) {
		toSerialize["minEventId"] = o.MinEventId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TotalChargeAmount) {
		toSerialize["totalChargeAmount"] = o.TotalChargeAmount
	}
	return toSerialize, nil
}

func (o *UnibeeApiBeanUserMetricChargeInvoiceItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metricId",
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

	varUnibeeApiBeanUserMetricChargeInvoiceItem := _UnibeeApiBeanUserMetricChargeInvoiceItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiBeanUserMetricChargeInvoiceItem)

	if err != nil {
		return err
	}

	*o = UnibeeApiBeanUserMetricChargeInvoiceItem(varUnibeeApiBeanUserMetricChargeInvoiceItem)

	return err
}

type NullableUnibeeApiBeanUserMetricChargeInvoiceItem struct {
	value *UnibeeApiBeanUserMetricChargeInvoiceItem
	isSet bool
}

func (v NullableUnibeeApiBeanUserMetricChargeInvoiceItem) Get() *UnibeeApiBeanUserMetricChargeInvoiceItem {
	return v.value
}

func (v *NullableUnibeeApiBeanUserMetricChargeInvoiceItem) Set(val *UnibeeApiBeanUserMetricChargeInvoiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanUserMetricChargeInvoiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanUserMetricChargeInvoiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanUserMetricChargeInvoiceItem(val *UnibeeApiBeanUserMetricChargeInvoiceItem) *NullableUnibeeApiBeanUserMetricChargeInvoiceItem {
	return &NullableUnibeeApiBeanUserMetricChargeInvoiceItem{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanUserMetricChargeInvoiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanUserMetricChargeInvoiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


