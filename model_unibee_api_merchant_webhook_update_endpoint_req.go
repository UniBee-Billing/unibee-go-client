/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240451 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantWebhookUpdateEndpointReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantWebhookUpdateEndpointReq{}

// UnibeeApiMerchantWebhookUpdateEndpointReq struct for UnibeeApiMerchantWebhookUpdateEndpointReq
type UnibeeApiMerchantWebhookUpdateEndpointReq struct {
	// EndpointId
	EndpointId int64 `json:"endpointId"`
	// Events To Update
	Events []string `json:"events,omitempty"`
	// Url To Update
	Url string `json:"url"`
}

type _UnibeeApiMerchantWebhookUpdateEndpointReq UnibeeApiMerchantWebhookUpdateEndpointReq

// NewUnibeeApiMerchantWebhookUpdateEndpointReq instantiates a new UnibeeApiMerchantWebhookUpdateEndpointReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantWebhookUpdateEndpointReq(endpointId int64, url string) *UnibeeApiMerchantWebhookUpdateEndpointReq {
	this := UnibeeApiMerchantWebhookUpdateEndpointReq{}
	this.EndpointId = endpointId
	this.Url = url
	return &this
}

// NewUnibeeApiMerchantWebhookUpdateEndpointReqWithDefaults instantiates a new UnibeeApiMerchantWebhookUpdateEndpointReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantWebhookUpdateEndpointReqWithDefaults() *UnibeeApiMerchantWebhookUpdateEndpointReq {
	this := UnibeeApiMerchantWebhookUpdateEndpointReq{}
	return &this
}

// GetEndpointId returns the EndpointId field value
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEndpointId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEndpointIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointId, true
}

// SetEndpointId sets field value
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetEndpointId(v int64) {
	o.EndpointId = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetEvents(v []string) {
	o.Events = v
}

// GetUrl returns the Url field value
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetUrl(v string) {
	o.Url = v
}

func (o UnibeeApiMerchantWebhookUpdateEndpointReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantWebhookUpdateEndpointReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpointId"] = o.EndpointId
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpointId",
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

	varUnibeeApiMerchantWebhookUpdateEndpointReq := _UnibeeApiMerchantWebhookUpdateEndpointReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantWebhookUpdateEndpointReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantWebhookUpdateEndpointReq(varUnibeeApiMerchantWebhookUpdateEndpointReq)

	return err
}

type NullableUnibeeApiMerchantWebhookUpdateEndpointReq struct {
	value *UnibeeApiMerchantWebhookUpdateEndpointReq
	isSet bool
}

func (v NullableUnibeeApiMerchantWebhookUpdateEndpointReq) Get() *UnibeeApiMerchantWebhookUpdateEndpointReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantWebhookUpdateEndpointReq) Set(val *UnibeeApiMerchantWebhookUpdateEndpointReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantWebhookUpdateEndpointReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantWebhookUpdateEndpointReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantWebhookUpdateEndpointReq(val *UnibeeApiMerchantWebhookUpdateEndpointReq) *NullableUnibeeApiMerchantWebhookUpdateEndpointReq {
	return &NullableUnibeeApiMerchantWebhookUpdateEndpointReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantWebhookUpdateEndpointReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantWebhookUpdateEndpointReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


