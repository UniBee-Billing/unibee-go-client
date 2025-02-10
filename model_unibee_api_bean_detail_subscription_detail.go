/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanDetailSubscriptionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanDetailSubscriptionDetail{}

// UnibeeApiBeanDetailSubscriptionDetail struct for UnibeeApiBeanDetailSubscriptionDetail
type UnibeeApiBeanDetailSubscriptionDetail struct {
	// AddonParams
	AddonParams []UnibeeApiBeanPlanAddonParam `json:"addonParams,omitempty"`
	// Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	// DayLeft util the period end, only available for webhook
	DayLeft *int32 `json:"dayLeft,omitempty"`
	Discount *UnibeeApiBeanMerchantDiscountCode `json:"discount,omitempty"`
	Gateway *UnibeeApiBeanDetailGateway `json:"gateway,omitempty"`
	LatestInvoice *UnibeeApiBeanInvoice `json:"latestInvoice,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
	UnfinishedSubscriptionPendingUpdate *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail `json:"unfinishedSubscriptionPendingUpdate,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
}

// NewUnibeeApiBeanDetailSubscriptionDetail instantiates a new UnibeeApiBeanDetailSubscriptionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanDetailSubscriptionDetail() *UnibeeApiBeanDetailSubscriptionDetail {
	this := UnibeeApiBeanDetailSubscriptionDetail{}
	return &this
}

// NewUnibeeApiBeanDetailSubscriptionDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanDetailSubscriptionDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionDetail {
	this := UnibeeApiBeanDetailSubscriptionDetail{}
	return &this
}

// GetAddonParams returns the AddonParams field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonParams() []UnibeeApiBeanPlanAddonParam {
	if o == nil || IsNil(o.AddonParams) {
		var ret []UnibeeApiBeanPlanAddonParam
		return ret
	}
	return o.AddonParams
}

// GetAddonParamsOk returns a tuple with the AddonParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonParamsOk() ([]UnibeeApiBeanPlanAddonParam, bool) {
	if o == nil || IsNil(o.AddonParams) {
		return nil, false
	}
	return o.AddonParams, true
}

// HasAddonParams returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasAddonParams() bool {
	if o != nil && !IsNil(o.AddonParams) {
		return true
	}

	return false
}

// SetAddonParams gets a reference to the given []UnibeeApiBeanPlanAddonParam and assigns it to the AddonParams field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetAddonParams(v []UnibeeApiBeanPlanAddonParam) {
	o.AddonParams = v
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetDayLeft returns the DayLeft field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDayLeft() int32 {
	if o == nil || IsNil(o.DayLeft) {
		var ret int32
		return ret
	}
	return *o.DayLeft
}

// GetDayLeftOk returns a tuple with the DayLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDayLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.DayLeft) {
		return nil, false
	}
	return o.DayLeft, true
}

// HasDayLeft returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasDayLeft() bool {
	if o != nil && !IsNil(o.DayLeft) {
		return true
	}

	return false
}

// SetDayLeft gets a reference to the given int32 and assigns it to the DayLeft field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetDayLeft(v int32) {
	o.DayLeft = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDiscount() UnibeeApiBeanMerchantDiscountCode {
	if o == nil || IsNil(o.Discount) {
		var ret UnibeeApiBeanMerchantDiscountCode
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given UnibeeApiBeanMerchantDiscountCode and assigns it to the Discount field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetDiscount(v UnibeeApiBeanMerchantDiscountCode) {
	o.Discount = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetGateway() UnibeeApiBeanDetailGateway {
	if o == nil || IsNil(o.Gateway) {
		var ret UnibeeApiBeanDetailGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given UnibeeApiBeanDetailGateway and assigns it to the Gateway field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetGateway(v UnibeeApiBeanDetailGateway) {
	o.Gateway = &v
}

// GetLatestInvoice returns the LatestInvoice field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetLatestInvoice() UnibeeApiBeanInvoice {
	if o == nil || IsNil(o.LatestInvoice) {
		var ret UnibeeApiBeanInvoice
		return ret
	}
	return *o.LatestInvoice
}

// GetLatestInvoiceOk returns a tuple with the LatestInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetLatestInvoiceOk() (*UnibeeApiBeanInvoice, bool) {
	if o == nil || IsNil(o.LatestInvoice) {
		return nil, false
	}
	return o.LatestInvoice, true
}

// HasLatestInvoice returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasLatestInvoice() bool {
	if o != nil && !IsNil(o.LatestInvoice) {
		return true
	}

	return false
}

// SetLatestInvoice gets a reference to the given UnibeeApiBeanInvoice and assigns it to the LatestInvoice field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetLatestInvoice(v UnibeeApiBeanInvoice) {
	o.LatestInvoice = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

// GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		var ret UnibeeApiBeanDetailSubscriptionPendingUpdateDetail
		return ret
	}
	return *o.UnfinishedSubscriptionPendingUpdate
}

// GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool) {
	if o == nil || IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return nil, false
	}
	return o.UnfinishedSubscriptionPendingUpdate, true
}

// HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasUnfinishedSubscriptionPendingUpdate() bool {
	if o != nil && !IsNil(o.UnfinishedSubscriptionPendingUpdate) {
		return true
	}

	return false
}

// SetUnfinishedSubscriptionPendingUpdate gets a reference to the given UnibeeApiBeanDetailSubscriptionPendingUpdateDetail and assigns it to the UnfinishedSubscriptionPendingUpdate field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) {
	o.UnfinishedSubscriptionPendingUpdate = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiBeanDetailSubscriptionDetail) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiBeanDetailSubscriptionDetail) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

func (o UnibeeApiBeanDetailSubscriptionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanDetailSubscriptionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddonParams) {
		toSerialize["addonParams"] = o.AddonParams
	}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.DayLeft) {
		toSerialize["dayLeft"] = o.DayLeft
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
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

type NullableUnibeeApiBeanDetailSubscriptionDetail struct {
	value *UnibeeApiBeanDetailSubscriptionDetail
	isSet bool
}

func (v NullableUnibeeApiBeanDetailSubscriptionDetail) Get() *UnibeeApiBeanDetailSubscriptionDetail {
	return v.value
}

func (v *NullableUnibeeApiBeanDetailSubscriptionDetail) Set(val *UnibeeApiBeanDetailSubscriptionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanDetailSubscriptionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanDetailSubscriptionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanDetailSubscriptionDetail(val *UnibeeApiBeanDetailSubscriptionDetail) *NullableUnibeeApiBeanDetailSubscriptionDetail {
	return &NullableUnibeeApiBeanDetailSubscriptionDetail{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanDetailSubscriptionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanDetailSubscriptionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


