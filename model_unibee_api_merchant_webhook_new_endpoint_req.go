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

// checks if the UnibeeApiMerchantWebhookNewEndpointReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookNewEndpointReq{}

// UnibeeApiMerchantWebhookNewEndpointReq struct for UnibeeApiMerchantWebhookNewEndpointReq
type UnibeeApiMerchantWebhookNewEndpointReq struct {
	// Events
	Events []string `json:"events,omitempty"`
	// Url
	Url string `json:"url"`
}

type _UnibeeApiMerchantWebhookNewEndpointReq UnibeeApiMerchantWebhookNewEndpointReq

// NewUnibeeApiMerchantWebhookNewEndpointReq instantiates a new UnibeeApiMerchantWebhookNewEndpointReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookNewEndpointReq(url string) *UnibeeApiMerchantWebhookNewEndpointReq {
	this := UnibeeApiMerchantWebhookNewEndpointReq{}
	this.Url = url
	return &this
}

// NewUnibeeApiMerchantWebhookNewEndpointReqWithDefaults instantiates a new UnibeeApiMerchantWebhookNewEndpointReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookNewEndpointReqWithDefaults() *UnibeeApiMerchantWebhookNewEndpointReq {
	this := UnibeeApiMerchantWebhookNewEndpointReq{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookNewEndpointReq) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookNewEndpointReq) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookNewEndpointReq) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *UnibeeApiMerchantWebhookNewEndpointReq) SetEvents(v []string) {
	o.Events = v
}

// GetUrl returns the Url field value
func (o *UnibeeApiMerchantWebhookNewEndpointReq) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookNewEndpointReq) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UnibeeApiMerchantWebhookNewEndpointReq) SetUrl(v string) {
	o.Url = v
}

func (o UnibeeApiMerchantWebhookNewEndpointReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookNewEndpointReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *UnibeeApiMerchantWebhookNewEndpointReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varUnibeeApiMerchantWebhookNewEndpointReq := _UnibeeApiMerchantWebhookNewEndpointReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantWebhookNewEndpointReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantWebhookNewEndpointReq(varUnibeeApiMerchantWebhookNewEndpointReq)

	return err
}

type NullableUnibeeApiMerchantWebhookNewEndpointReq struct {
	value *UnibeeApiMerchantWebhookNewEndpointReq
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookNewEndpointReq) Get() *UnibeeApiMerchantWebhookNewEndpointReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookNewEndpointReq) Set(val *UnibeeApiMerchantWebhookNewEndpointReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookNewEndpointReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookNewEndpointReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookNewEndpointReq(val *UnibeeApiMerchantWebhookNewEndpointReq) *NullableUnibeeApiMerchantWebhookNewEndpointReq {
	return &NullableUnibeeApiMerchantWebhookNewEndpointReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookNewEndpointReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookNewEndpointReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


