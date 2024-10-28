/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410280512
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanMerchantWebhookLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanMerchantWebhookLog{}

// UnibeeApiBeanMerchantWebhookLog struct for UnibeeApiBeanMerchantWebhookLog
type UnibeeApiBeanMerchantWebhookLog struct {
	// body(json)
	Body *string `json:"body,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	EndpointId *int64 `json:"endpointId,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// mamo
	Mamo *string `json:"mamo,omitempty"`
	// webhook url
	MerchantId *int64 `json:"merchantId,omitempty"`
	ReconsumeCount *int32 `json:"reconsumeCount,omitempty"`
	// request_id
	RequestId *string `json:"requestId,omitempty"`
	// response
	Response *string `json:"response,omitempty"`
	// webhook_event
	WebhookEvent *string `json:"webhookEvent,omitempty"`
	// webhook url
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

// NewUnibeeApiBeanMerchantWebhookLog instantiates a new UnibeeApiBeanMerchantWebhookLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanMerchantWebhookLog() *UnibeeApiBeanMerchantWebhookLog {
	this := UnibeeApiBeanMerchantWebhookLog{}
	return &this
}

// NewUnibeeApiBeanMerchantWebhookLogWithDefaults instantiates a new UnibeeApiBeanMerchantWebhookLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanMerchantWebhookLogWithDefaults() *UnibeeApiBeanMerchantWebhookLog {
	this := UnibeeApiBeanMerchantWebhookLog{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetBody(v string) {
	o.Body = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetEndpointId() int64 {
	if o == nil || IsNil(o.EndpointId) {
		var ret int64
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetEndpointIdOk() (*int64, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given int64 and assigns it to the EndpointId field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetEndpointId(v int64) {
	o.EndpointId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetId(v int64) {
	o.Id = &v
}

// GetMamo returns the Mamo field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetMamo() string {
	if o == nil || IsNil(o.Mamo) {
		var ret string
		return ret
	}
	return *o.Mamo
}

// GetMamoOk returns a tuple with the Mamo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetMamoOk() (*string, bool) {
	if o == nil || IsNil(o.Mamo) {
		return nil, false
	}
	return o.Mamo, true
}

// HasMamo returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasMamo() bool {
	if o != nil && !IsNil(o.Mamo) {
		return true
	}

	return false
}

// SetMamo gets a reference to the given string and assigns it to the Mamo field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetMamo(v string) {
	o.Mamo = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetReconsumeCount returns the ReconsumeCount field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetReconsumeCount() int32 {
	if o == nil || IsNil(o.ReconsumeCount) {
		var ret int32
		return ret
	}
	return *o.ReconsumeCount
}

// GetReconsumeCountOk returns a tuple with the ReconsumeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetReconsumeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReconsumeCount) {
		return nil, false
	}
	return o.ReconsumeCount, true
}

// HasReconsumeCount returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasReconsumeCount() bool {
	if o != nil && !IsNil(o.ReconsumeCount) {
		return true
	}

	return false
}

// SetReconsumeCount gets a reference to the given int32 and assigns it to the ReconsumeCount field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetReconsumeCount(v int32) {
	o.ReconsumeCount = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetRequestId(v string) {
	o.RequestId = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetResponse(v string) {
	o.Response = &v
}

// GetWebhookEvent returns the WebhookEvent field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookEvent() string {
	if o == nil || IsNil(o.WebhookEvent) {
		var ret string
		return ret
	}
	return *o.WebhookEvent
}

// GetWebhookEventOk returns a tuple with the WebhookEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookEventOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEvent) {
		return nil, false
	}
	return o.WebhookEvent, true
}

// HasWebhookEvent returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasWebhookEvent() bool {
	if o != nil && !IsNil(o.WebhookEvent) {
		return true
	}

	return false
}

// SetWebhookEvent gets a reference to the given string and assigns it to the WebhookEvent field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetWebhookEvent(v string) {
	o.WebhookEvent = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanMerchantWebhookLog) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *UnibeeApiBeanMerchantWebhookLog) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o UnibeeApiBeanMerchantWebhookLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanMerchantWebhookLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.EndpointId) {
		toSerialize["endpointId"] = o.EndpointId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mamo) {
		toSerialize["mamo"] = o.Mamo
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.ReconsumeCount) {
		toSerialize["reconsumeCount"] = o.ReconsumeCount
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.WebhookEvent) {
		toSerialize["webhookEvent"] = o.WebhookEvent
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanMerchantWebhookLog struct {
	value *UnibeeApiBeanMerchantWebhookLog
	isSet bool
}

func (v NullableUnibeeApiBeanMerchantWebhookLog) Get() *UnibeeApiBeanMerchantWebhookLog {
	return v.value
}

func (v *NullableUnibeeApiBeanMerchantWebhookLog) Set(val *UnibeeApiBeanMerchantWebhookLog) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanMerchantWebhookLog) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanMerchantWebhookLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanMerchantWebhookLog(val *UnibeeApiBeanMerchantWebhookLog) *NullableUnibeeApiBeanMerchantWebhookLog {
	return &NullableUnibeeApiBeanMerchantWebhookLog{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanMerchantWebhookLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanMerchantWebhookLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


