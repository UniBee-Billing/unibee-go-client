/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408190753 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantMetricUserMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantMetricUserMetric{}

// UnibeeApiMerchantMetricUserMetric struct for UnibeeApiMerchantMetricUserMetric
type UnibeeApiMerchantMetricUserMetric struct {
	// Addon
	Addons []UnibeeApiBeanPlanAddonDetail `json:"addons,omitempty"`
	// IsPaid
	IsPaid *bool `json:"isPaid,omitempty"`
	Plan *UnibeeApiBeanPlan `json:"plan,omitempty"`
	Subscription *UnibeeApiBeanSubscription `json:"subscription,omitempty"`
	User *UnibeeApiBeanUserAccount `json:"user,omitempty"`
	// UserMerchantMetricStats
	UserMerchantMetricStats []UnibeeApiBeanUserMerchantMetricStat `json:"userMerchantMetricStats,omitempty"`
}

// NewUnibeeApiMerchantMetricUserMetric instantiates a new UnibeeApiMerchantMetricUserMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantMetricUserMetric() *UnibeeApiMerchantMetricUserMetric {
	this := UnibeeApiMerchantMetricUserMetric{}
	return &this
}

// NewUnibeeApiMerchantMetricUserMetricWithDefaults instantiates a new UnibeeApiMerchantMetricUserMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantMetricUserMetricWithDefaults() *UnibeeApiMerchantMetricUserMetric {
	this := UnibeeApiMerchantMetricUserMetric{}
	return &this
}

// GetAddons returns the Addons field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetAddons() []UnibeeApiBeanPlanAddonDetail {
	if o == nil || IsNil(o.Addons) {
		var ret []UnibeeApiBeanPlanAddonDetail
		return ret
	}
	return o.Addons
}

// GetAddonsOk returns a tuple with the Addons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetAddonsOk() ([]UnibeeApiBeanPlanAddonDetail, bool) {
	if o == nil || IsNil(o.Addons) {
		return nil, false
	}
	return o.Addons, true
}

// HasAddons returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasAddons() bool {
	if o != nil && !IsNil(o.Addons) {
		return true
	}

	return false
}

// SetAddons gets a reference to the given []UnibeeApiBeanPlanAddonDetail and assigns it to the Addons field.
func (o *UnibeeApiMerchantMetricUserMetric) SetAddons(v []UnibeeApiBeanPlanAddonDetail) {
	o.Addons = v
}

// GetIsPaid returns the IsPaid field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaid() bool {
	if o == nil || IsNil(o.IsPaid) {
		var ret bool
		return ret
	}
	return *o.IsPaid
}

// GetIsPaidOk returns a tuple with the IsPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaid) {
		return nil, false
	}
	return o.IsPaid, true
}

// HasIsPaid returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasIsPaid() bool {
	if o != nil && !IsNil(o.IsPaid) {
		return true
	}

	return false
}

// SetIsPaid gets a reference to the given bool and assigns it to the IsPaid field.
func (o *UnibeeApiMerchantMetricUserMetric) SetIsPaid(v bool) {
	o.IsPaid = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetPlan() UnibeeApiBeanPlan {
	if o == nil || IsNil(o.Plan) {
		var ret UnibeeApiBeanPlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetPlanOk() (*UnibeeApiBeanPlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given UnibeeApiBeanPlan and assigns it to the Plan field.
func (o *UnibeeApiMerchantMetricUserMetric) SetPlan(v UnibeeApiBeanPlan) {
	o.Plan = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetSubscription() UnibeeApiBeanSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret UnibeeApiBeanSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given UnibeeApiBeanSubscription and assigns it to the Subscription field.
func (o *UnibeeApiMerchantMetricUserMetric) SetSubscription(v UnibeeApiBeanSubscription) {
	o.Subscription = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetUser() UnibeeApiBeanUserAccount {
	if o == nil || IsNil(o.User) {
		var ret UnibeeApiBeanUserAccount
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetUserOk() (*UnibeeApiBeanUserAccount, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UnibeeApiBeanUserAccount and assigns it to the User field.
func (o *UnibeeApiMerchantMetricUserMetric) SetUser(v UnibeeApiBeanUserAccount) {
	o.User = &v
}

// GetUserMerchantMetricStats returns the UserMerchantMetricStats field value if set, zero value otherwise.
func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStats() []UnibeeApiBeanUserMerchantMetricStat {
	if o == nil || IsNil(o.UserMerchantMetricStats) {
		var ret []UnibeeApiBeanUserMerchantMetricStat
		return ret
	}
	return o.UserMerchantMetricStats
}

// GetUserMerchantMetricStatsOk returns a tuple with the UserMerchantMetricStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStatsOk() ([]UnibeeApiBeanUserMerchantMetricStat, bool) {
	if o == nil || IsNil(o.UserMerchantMetricStats) {
		return nil, false
	}
	return o.UserMerchantMetricStats, true
}

// HasUserMerchantMetricStats returns a boolean if a field has been set.
func (o *UnibeeApiMerchantMetricUserMetric) HasUserMerchantMetricStats() bool {
	if o != nil && !IsNil(o.UserMerchantMetricStats) {
		return true
	}

	return false
}

// SetUserMerchantMetricStats gets a reference to the given []UnibeeApiBeanUserMerchantMetricStat and assigns it to the UserMerchantMetricStats field.
func (o *UnibeeApiMerchantMetricUserMetric) SetUserMerchantMetricStats(v []UnibeeApiBeanUserMerchantMetricStat) {
	o.UserMerchantMetricStats = v
}

func (o UnibeeApiMerchantMetricUserMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantMetricUserMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addons) {
		toSerialize["addons"] = o.Addons
	}
	if !IsNil(o.IsPaid) {
		toSerialize["isPaid"] = o.IsPaid
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserMerchantMetricStats) {
		toSerialize["userMerchantMetricStats"] = o.UserMerchantMetricStats
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantMetricUserMetric struct {
	value *UnibeeApiMerchantMetricUserMetric
	isSet bool
}

func (v NullableUnibeeApiMerchantMetricUserMetric) Get() *UnibeeApiMerchantMetricUserMetric {
	return v.value
}

func (v *NullableUnibeeApiMerchantMetricUserMetric) Set(val *UnibeeApiMerchantMetricUserMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantMetricUserMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantMetricUserMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantMetricUserMetric(val *UnibeeApiMerchantMetricUserMetric) *NullableUnibeeApiMerchantMetricUserMetric {
	return &NullableUnibeeApiMerchantMetricUserMetric{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantMetricUserMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantMetricUserMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


