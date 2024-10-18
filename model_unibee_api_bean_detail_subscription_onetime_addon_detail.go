/*
OpenAPI UniBee

This is UniBee Api Server

API version: daily,buildtime:202410181820
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail{}

// UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail struct for UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail
type UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail struct {
	Addon *UnibeeApiBeanPlan `json:"addon,omitempty"`
	// onetime addonId
	AddonId *int64 `json:"addonId,omitempty"`
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// id
	Id *int64 `json:"id,omitempty"`
	// Metadata
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	Payment *UnibeeApiBeanPayment `json:"payment,omitempty"`
	// quantity
	Quantity *int64 `json:"quantity,omitempty"`
	// status, 1-create, 2-paid, 3-cancel, 4-expired
	Status *int32 `json:"status,omitempty"`
	// subscription_id
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail instantiates a new UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail() *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	this := UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail{}
	return &this
}

// NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	this := UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail{}
	return &this
}

// GetAddon returns the Addon field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddon() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Addon) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Addon
}

// GetAddonOk returns a tuple with the Addon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Addon) {
		return nil, false
	}
	return o.Addon, true
}

// HasAddon returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasAddon() bool {
	if o != nil && !IsNil(o.Addon) {
		return true
	}

	return false
}

// SetAddon gets a reference to the given UnibeeApiBeanPlan and assigns it to the Addon field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetAddon(v UnibeeApiBeanPlan) {
	o.Addon = &v
}

// GetAddonId returns the AddonId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonId() int64 {
	if o == nil || IsNil(o.AddonId) {
		var ret int64
		return ret
	}
	return *o.AddonId
}

// GetAddonIdOk returns a tuple with the AddonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AddonId) {
		return nil, false
	}
	return o.AddonId, true
}

// HasAddonId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasAddonId() bool {
	if o != nil && !IsNil(o.AddonId) {
		return true
	}

	return false
}

// SetAddonId gets a reference to the given int64 and assigns it to the AddonId field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetAddonId(v int64) {
	o.AddonId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetId(v int64) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetPayment() UnibeeApiBeanPayment {
	if o == nil || IsNil(o.Payment) {
		var ret UnibeeApiBeanPayment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given UnibeeApiBeanPayment and assigns it to the Payment field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetPayment(v UnibeeApiBeanPayment) {
	o.Payment = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetQuantity(v int64) {
	o.Quantity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetStatus(v int32) {
	o.Status = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addon) {
		toSerialize["addon"] = o.Addon
	}
	if !IsNil(o.AddonId) {
		toSerialize["addonId"] = o.AddonId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail struct {
	value *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) Get() *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) Set(val *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail(val *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) *NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail {
	return &NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


