/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100809
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiSystemInvoiceInternalWebhookSyncRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiSystemInvoiceInternalWebhookSyncRes{}

// UnibeeApiSystemInvoiceInternalWebhookSyncRes struct for UnibeeApiSystemInvoiceInternalWebhookSyncRes
type UnibeeApiSystemInvoiceInternalWebhookSyncRes struct {
	// The first invoiceId of sync data, only output when isSynchronous=true
	FirstId *string `json:"firstId,omitempty"`
	// The last invoiceId of sync data, only output when isSynchronous=true
	LastId *string `json:"lastId,omitempty"`
	// The total of sync, only output when isSynchronous=true
	Total *int32 `json:"total,omitempty"`
}

// NewUnibeeApiSystemInvoiceInternalWebhookSyncRes instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiSystemInvoiceInternalWebhookSyncRes() *UnibeeApiSystemInvoiceInternalWebhookSyncRes {
	this := UnibeeApiSystemInvoiceInternalWebhookSyncRes{}
	return &this
}

// NewUnibeeApiSystemInvoiceInternalWebhookSyncResWithDefaults instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiSystemInvoiceInternalWebhookSyncResWithDefaults() *UnibeeApiSystemInvoiceInternalWebhookSyncRes {
	this := UnibeeApiSystemInvoiceInternalWebhookSyncRes{}
	return &this
}

// GetFirstId returns the FirstId field value if set, zero value otherwise.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetFirstId() string {
	if o == nil || IsNil(o.FirstId) {
		var ret string
		return ret
	}
	return *o.FirstId
}

// GetFirstIdOk returns a tuple with the FirstId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetFirstIdOk() (*string, bool) {
	if o == nil || IsNil(o.FirstId) {
		return nil, false
	}
	return o.FirstId, true
}

// HasFirstId returns a boolean if a field has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasFirstId() bool {
	if o != nil && !IsNil(o.FirstId) {
		return true
	}

	return false
}

// SetFirstId gets a reference to the given string and assigns it to the FirstId field.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetFirstId(v string) {
	o.FirstId = &v
}

// GetLastId returns the LastId field value if set, zero value otherwise.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetLastId() string {
	if o == nil || IsNil(o.LastId) {
		var ret string
		return ret
	}
	return *o.LastId
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetLastIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastId) {
		return nil, false
	}
	return o.LastId, true
}

// HasLastId returns a boolean if a field has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasLastId() bool {
	if o != nil && !IsNil(o.LastId) {
		return true
	}

	return false
}

// SetLastId gets a reference to the given string and assigns it to the LastId field.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetLastId(v string) {
	o.LastId = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetTotal(v int32) {
	o.Total = &v
}

func (o UnibeeApiSystemInvoiceInternalWebhookSyncRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiSystemInvoiceInternalWebhookSyncRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstId) {
		toSerialize["firstId"] = o.FirstId
	}
	if !IsNil(o.LastId) {
		toSerialize["lastId"] = o.LastId
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes struct {
	value *UnibeeApiSystemInvoiceInternalWebhookSyncRes
	isSet bool
}

func (v NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) Get() *UnibeeApiSystemInvoiceInternalWebhookSyncRes {
	return v.value
}

func (v *NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) Set(val *UnibeeApiSystemInvoiceInternalWebhookSyncRes) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiSystemInvoiceInternalWebhookSyncRes(val *UnibeeApiSystemInvoiceInternalWebhookSyncRes) *NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes {
	return &NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes{value: val, isSet: true}
}

func (v NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiSystemInvoiceInternalWebhookSyncRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


