/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408150845 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantWebhookEndpointLogListGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantWebhookEndpointLogListGet200ResponseData{}

// MerchantWebhookEndpointLogListGet200ResponseData struct for MerchantWebhookEndpointLogListGet200ResponseData
type MerchantWebhookEndpointLogListGet200ResponseData struct {
	// EndpointLogList
	EndpointLogList []UnibeeApiBeanMerchantWebhookLog `json:"endpointLogList,omitempty"`
	// Total
	Total *int32 `json:"total,omitempty"`
}

// NewMerchantWebhookEndpointLogListGet200ResponseData instantiates a new MerchantWebhookEndpointLogListGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantWebhookEndpointLogListGet200ResponseData() *MerchantWebhookEndpointLogListGet200ResponseData {
	this := MerchantWebhookEndpointLogListGet200ResponseData{}
	return &this
}

// NewMerchantWebhookEndpointLogListGet200ResponseDataWithDefaults instantiates a new MerchantWebhookEndpointLogListGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantWebhookEndpointLogListGet200ResponseDataWithDefaults() *MerchantWebhookEndpointLogListGet200ResponseData {
	this := MerchantWebhookEndpointLogListGet200ResponseData{}
	return &this
}

// GetEndpointLogList returns the EndpointLogList field value if set, zero value otherwise.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) GetEndpointLogList() []UnibeeApiBeanMerchantWebhookLog {
	if o == nil || IsNil(o.EndpointLogList) {
		var ret []UnibeeApiBeanMerchantWebhookLog
		return ret
	}
	return o.EndpointLogList
}

// GetEndpointLogListOk returns a tuple with the EndpointLogList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) GetEndpointLogListOk() ([]UnibeeApiBeanMerchantWebhookLog, bool) {
	if o == nil || IsNil(o.EndpointLogList) {
		return nil, false
	}
	return o.EndpointLogList, true
}

// HasEndpointLogList returns a boolean if a field has been set.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) HasEndpointLogList() bool {
	if o != nil && !IsNil(o.EndpointLogList) {
		return true
	}

	return false
}

// SetEndpointLogList gets a reference to the given []UnibeeApiBeanMerchantWebhookLog and assigns it to the EndpointLogList field.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) SetEndpointLogList(v []UnibeeApiBeanMerchantWebhookLog) {
	o.EndpointLogList = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *MerchantWebhookEndpointLogListGet200ResponseData) SetTotal(v int32) {
	o.Total = &v
}

func (o MerchantWebhookEndpointLogListGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantWebhookEndpointLogListGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointLogList) {
		toSerialize["endpointLogList"] = o.EndpointLogList
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableMerchantWebhookEndpointLogListGet200ResponseData struct {
	value *MerchantWebhookEndpointLogListGet200ResponseData
	isSet bool
}

func (v NullableMerchantWebhookEndpointLogListGet200ResponseData) Get() *MerchantWebhookEndpointLogListGet200ResponseData {
	return v.value
}

func (v *NullableMerchantWebhookEndpointLogListGet200ResponseData) Set(val *MerchantWebhookEndpointLogListGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantWebhookEndpointLogListGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantWebhookEndpointLogListGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantWebhookEndpointLogListGet200ResponseData(val *MerchantWebhookEndpointLogListGet200ResponseData) *NullableMerchantWebhookEndpointLogListGet200ResponseData {
	return &NullableMerchantWebhookEndpointLogListGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantWebhookEndpointLogListGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantWebhookEndpointLogListGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


