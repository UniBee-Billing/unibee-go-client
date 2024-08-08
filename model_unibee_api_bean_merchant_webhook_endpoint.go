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

// checks if the UnibeeApiBeanMerchantWebhookEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantWebhookEndpoint{}

// UnibeeApiBeanMerchantWebhookEndpoint struct for UnibeeApiBeanMerchantWebhookEndpoint
type UnibeeApiBeanMerchantWebhookEndpoint struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// update time
	GmtModify *int64 `json:"gmtModify,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// webhook url
	MerchantId *int64 `json:"merchantId,omitempty"`
	// webhook_events,split dot
	WebhookEvents []string `json:"webhookEvents,omitempty"`
	// webhook url
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

// NewUnibeeApiBeanMerchantWebhookEndpoint instantiates a new UnibeeApiBeanMerchantWebhookEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantWebhookEndpoint() *UnibeeApiBeanMerchantWebhookEndpoint {
	this := UnibeeApiBeanMerchantWebhookEndpoint{}
	return &this
}

// NewUnibeeApiBeanMerchantWebhookEndpointWithDefaults instantiates a new UnibeeApiBeanMerchantWebhookEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantWebhookEndpointWithDefaults() *UnibeeApiBeanMerchantWebhookEndpoint {
	this := UnibeeApiBeanMerchantWebhookEndpoint{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetGmtModify returns the GmtModify field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetGmtModify() int64 {
	if o == nil || IsNil(o.GmtModify) {
		var ret int64
		return ret
	}
	return *o.GmtModify
}

// GetGmtModifyOk returns a tuple with the GmtModify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetGmtModifyOk() (*int64, bool) {
	if o == nil || IsNil(o.GmtModify) {
		return nil, false
	}
	return o.GmtModify, true
}

// HasGmtModify returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasGmtModify() bool {
	if o != nil && !IsNil(o.GmtModify) {
		return true
	}

	return false
}

// SetGmtModify gets a reference to the given int64 and assigns it to the GmtModify field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetGmtModify(v int64) {
	o.GmtModify = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetId(v int64) {
	o.Id = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetWebhookEvents returns the WebhookEvents field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookEvents() []string {
	if o == nil || IsNil(o.WebhookEvents) {
		var ret []string
		return ret
	}
	return o.WebhookEvents
}

// GetWebhookEventsOk returns a tuple with the WebhookEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.WebhookEvents) {
		return nil, false
	}
	return o.WebhookEvents, true
}

// HasWebhookEvents returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasWebhookEvents() bool {
	if o != nil && !IsNil(o.WebhookEvents) {
		return true
	}

	return false
}

// SetWebhookEvents gets a reference to the given []string and assigns it to the WebhookEvents field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetWebhookEvents(v []string) {
	o.WebhookEvents = v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o UnibeeApiBeanMerchantWebhookEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantWebhookEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.GmtModify) {
		toSerialize["gmtModify"] = o.GmtModify
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.WebhookEvents) {
		toSerialize["webhookEvents"] = o.WebhookEvents
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantWebhookEndpoint struct {
	value *UnibeeApiBeanMerchantWebhookEndpoint
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantWebhookEndpoint) Get() *UnibeeApiBeanMerchantWebhookEndpoint {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantWebhookEndpoint) Set(val *UnibeeApiBeanMerchantWebhookEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantWebhookEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantWebhookEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantWebhookEndpoint(val *UnibeeApiBeanMerchantWebhookEndpoint) *NullableUnibeeApiBeanMerchantWebhookEndpoint {
	return &NullableUnibeeApiBeanMerchantWebhookEndpoint{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantWebhookEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantWebhookEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


