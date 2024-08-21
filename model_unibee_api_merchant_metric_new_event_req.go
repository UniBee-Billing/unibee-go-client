/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408210718 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricNewEventReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricNewEventReq{}

// UnibeeApiMerchantMetricNewEventReq struct for UnibeeApiMerchantMetricNewEventReq
type UnibeeApiMerchantMetricNewEventReq struct {
	// ExternalEventId, __unique__
	ExternalEventId string `json:"externalEventId"`
	// ExternalUserId
	ExternalUserId string `json:"externalUserId"`
	// MetricCode
	MetricCode string `json:"metricCode"`
	MetricProperties map[string]interface{} `json:"metricProperties,omitempty"`
	// default product will use if productId not specified and subscriptionId is blank
	ProductId *int64 `json:"productId,omitempty"`
}

type _UnibeeApiMerchantMetricNewEventReq UnibeeApiMerchantMetricNewEventReq

// NewUnibeeApiMerchantMetricNewEventReq instantiates a new UnibeeApiMerchantMetricNewEventReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricNewEventReq(externalEventId string, externalUserId string, metricCode string) *UnibeeApiMerchantMetricNewEventReq {
	this := UnibeeApiMerchantMetricNewEventReq{}
	this.ExternalEventId = externalEventId
	this.ExternalUserId = externalUserId
	this.MetricCode = metricCode
	return &this
}

// NewUnibeeApiMerchantMetricNewEventReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewEventReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricNewEventReqWithDefaults() *UnibeeApiMerchantMetricNewEventReq {
	this := UnibeeApiMerchantMetricNewEventReq{}
	return &this
}

// GetExternalEventId returns the ExternalEventId field value
func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventId, true
}

// SetExternalEventId sets field value
func (o *UnibeeApiMerchantMetricNewEventReq) SetExternalEventId(v string) {
	o.ExternalEventId = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *UnibeeApiMerchantMetricNewEventReq) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetMetricCode returns the MetricCode field value
func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricCode
}

// GetMetricCodeOk returns a tuple with the MetricCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricCode, true
}

// SetMetricCode sets field value
func (o *UnibeeApiMerchantMetricNewEventReq) SetMetricCode(v string) {
	o.MetricCode = v
}

// GetMetricProperties returns the MetricProperties field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricProperties() map[string]interface{} {
	if o == nil || IsNil(o.MetricProperties) {
		var ret map[string]interface{}
		return ret
	}
	return o.MetricProperties
}

// GetMetricPropertiesOk returns a tuple with the MetricProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MetricProperties) {
		return map[string]interface{}{}, false
	}
	return o.MetricProperties, true
}

// HasMetricProperties returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) HasMetricProperties() bool {
	if o != nil && !IsNil(o.MetricProperties) {
		return true
	}

	return false
}

// SetMetricProperties gets a reference to the given map[string]interface{} and assigns it to the MetricProperties field.
func (o *UnibeeApiMerchantMetricNewEventReq) SetMetricProperties(v map[string]interface{}) {
	o.MetricProperties = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricNewEventReq) GetProductId() int64 {
	if o == nil || IsNil(o.ProductId) {
		var ret int64
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) GetProductIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricNewEventReq) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int64 and assigns it to the ProductId field.
func (o *UnibeeApiMerchantMetricNewEventReq) SetProductId(v int64) {
	o.ProductId = &v
}

func (o UnibeeApiMerchantMetricNewEventReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricNewEventReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalEventId"] = o.ExternalEventId
	toSerialize["externalUserId"] = o.ExternalUserId
	toSerialize["metricCode"] = o.MetricCode
	if !IsNil(o.MetricProperties) {
		toSerialize["metricProperties"] = o.MetricProperties
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricNewEventReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"externalEventId",
		"externalUserId",
		"metricCode",
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

	varUnibeeApiMerchantMetricNewEventReq := _UnibeeApiMerchantMetricNewEventReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricNewEventReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricNewEventReq(varUnibeeApiMerchantMetricNewEventReq)

	return err
}

type NullableUnibeeApiMerchantMetricNewEventReq struct {
	value *UnibeeApiMerchantMetricNewEventReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricNewEventReq) Get() *UnibeeApiMerchantMetricNewEventReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricNewEventReq) Set(val *UnibeeApiMerchantMetricNewEventReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricNewEventReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricNewEventReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricNewEventReq(val *UnibeeApiMerchantMetricNewEventReq) *NullableUnibeeApiMerchantMetricNewEventReq {
	return &NullableUnibeeApiMerchantMetricNewEventReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricNewEventReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricNewEventReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


