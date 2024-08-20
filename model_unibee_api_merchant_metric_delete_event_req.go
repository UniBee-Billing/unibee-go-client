/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408200913 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantMetricDeleteEventReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricDeleteEventReq{}

// UnibeeApiMerchantMetricDeleteEventReq struct for UnibeeApiMerchantMetricDeleteEventReq
type UnibeeApiMerchantMetricDeleteEventReq struct {
	// ExternalEventId
	ExternalEventId string `json:"externalEventId"`
	// ExternalUserId
	ExternalUserId string `json:"externalUserId"`
	// MetricCode
	MetricCode string `json:"metricCode"`
}

type _UnibeeApiMerchantMetricDeleteEventReq UnibeeApiMerchantMetricDeleteEventReq

// NewUnibeeApiMerchantMetricDeleteEventReq instantiates a new UnibeeApiMerchantMetricDeleteEventReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricDeleteEventReq(externalEventId string, externalUserId string, metricCode string) *UnibeeApiMerchantMetricDeleteEventReq {
	this := UnibeeApiMerchantMetricDeleteEventReq{}
	this.ExternalEventId = externalEventId
	this.ExternalUserId = externalUserId
	this.MetricCode = metricCode
	return &this
}

// NewUnibeeApiMerchantMetricDeleteEventReqWithDefaults instantiates a new UnibeeApiMerchantMetricDeleteEventReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricDeleteEventReqWithDefaults() *UnibeeApiMerchantMetricDeleteEventReq {
	this := UnibeeApiMerchantMetricDeleteEventReq{}
	return &this
}

// GetExternalEventId returns the ExternalEventId field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventId, true
}

// SetExternalEventId sets field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) SetExternalEventId(v string) {
	o.ExternalEventId = v
}

// GetExternalUserId returns the ExternalUserId field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUserId, true
}

// SetExternalUserId sets field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) SetExternalUserId(v string) {
	o.ExternalUserId = v
}

// GetMetricCode returns the MetricCode field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetMetricCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricCode
}

// GetMetricCodeOk returns a tuple with the MetricCode field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricDeleteEventReq) GetMetricCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricCode, true
}

// SetMetricCode sets field value
func (o *UnibeeApiMerchantMetricDeleteEventReq) SetMetricCode(v string) {
	o.MetricCode = v
}

func (o UnibeeApiMerchantMetricDeleteEventReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricDeleteEventReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["externalEventId"] = o.ExternalEventId
	toSerialize["externalUserId"] = o.ExternalUserId
	toSerialize["metricCode"] = o.MetricCode
	return toSerialize, nil
}

func (o *UnibeeApiMerchantMetricDeleteEventReq) UnmarshalJSON(data []byte) (err error) {
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

	varUnibeeApiMerchantMetricDeleteEventReq := _UnibeeApiMerchantMetricDeleteEventReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantMetricDeleteEventReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantMetricDeleteEventReq(varUnibeeApiMerchantMetricDeleteEventReq)

	return err
}

type NullableUnibeeApiMerchantMetricDeleteEventReq struct {
	value *UnibeeApiMerchantMetricDeleteEventReq
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricDeleteEventReq) Get() *UnibeeApiMerchantMetricDeleteEventReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricDeleteEventReq) Set(val *UnibeeApiMerchantMetricDeleteEventReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricDeleteEventReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricDeleteEventReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricDeleteEventReq(val *UnibeeApiMerchantMetricDeleteEventReq) *NullableUnibeeApiMerchantMetricDeleteEventReq {
	return &NullableUnibeeApiMerchantMetricDeleteEventReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricDeleteEventReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricDeleteEventReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


