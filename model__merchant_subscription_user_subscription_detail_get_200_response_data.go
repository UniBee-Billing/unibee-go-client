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

// checks if the MerchantSubscriptionUserSubscriptionDetailGet200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantSubscriptionUserSubscriptionDetailGet200ResponseData{}

// MerchantSubscriptionUserSubscriptionDetailGet200ResponseData struct for MerchantSubscriptionUserSubscriptionDetailGet200ResponseData
type MerchantSubscriptionUserSubscriptionDetailGet200ResponseData struct {
	// Plan Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
	LatestInvoice *UnibeeApiBeanInvoice `json:"latestInvoice,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
	UnfinishedSubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"unfinishedSubscriptionPendingUpdate,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseData instantiates a new MerchantSubscriptionUserSubscriptionDetailGet200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseData() *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionUserSubscriptionDetailGet200ResponseData{}
	return &this
}

// NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUserSubscriptionDetailGet200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseDataWithDefaults() *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData {
	this := MerchantSubscriptionUserSubscriptionDetailGet200ResponseData{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

// GetLatestInvoice returns the LatestInvoice field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetLatestInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.LatestInvoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.LatestInvoice
}

// GetLatestInvoiceOk returns a tuple with the LatestInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetLatestInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.LatestInvoice) {
		return nil, false
	}
	return o.LatestInvoice, true
}

// HasLatestInvoice returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasLatestInvoice() bool {
	if o != nil && !IsNil(o.LatestInvoice) {
		return true
	}

	return false
}

// SetLatestInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the LatestInvoice field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetLatestInvoice(v UnibeeApiBeanInvoice) {
	o.LatestInvoice = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

// GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.UnfinishedSubscriptionPendingUpdate
}

// GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return nil, false
	}
	return o.UnfinishedSubscriptionPendingUpdate, true
}

// HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasUnfinishedSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetUnfinishedSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the UnfinishedSubscriptionPendingUpdate field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.UnfinishedSubscriptionPendingUpdate = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.LatestInvoice) {
		toSerialize["latestInvoice"] = o.LatestInvoice
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		toSerialize["unfinishedSubscriptionPendingUpdate"] = o.UnfinishedSubscriptionPendingUpdate
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData struct {
	value *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData
	isSet bool
}

func (v NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) Get() *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData {
	return v.value
}

func (v *NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) Set(val *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData(val *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) *NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData {
	return &NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData{value: val, isSet: true}
}

func (v NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantSubscriptionUserSubscriptionDetailGet200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


