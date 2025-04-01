# UnibeeApiBeanDetailUserMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addon | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**IsPaid** | Pointer to **bool** | IsPaid | [optional] 
**LimitStats** | Pointer to [**[]UnibeeApiBeanDetailUserMerchantMetricLimitStat**](UnibeeApiBeanDetailUserMerchantMetricLimitStat.md) | LimitStats | [optional] 
**MeteredChargeStats** | Pointer to [**[]UnibeeApiBeanDetailUserMerchantMetricChargeStat**](UnibeeApiBeanDetailUserMerchantMetricChargeStat.md) | MeteredChargeStats | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Product** | Pointer to [**UnibeeApiBeanProduct**](UnibeeApiBeanProduct.md) |  | [optional] 
**RecurringChargeStats** | Pointer to [**[]UnibeeApiBeanDetailUserMerchantMetricChargeStat**](UnibeeApiBeanDetailUserMerchantMetricChargeStat.md) | RecurringChargeStats | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserMetric

`func NewUnibeeApiBeanDetailUserMetric() *UnibeeApiBeanDetailUserMetric`

NewUnibeeApiBeanDetailUserMetric instantiates a new UnibeeApiBeanDetailUserMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserMetricWithDefaults

`func NewUnibeeApiBeanDetailUserMetricWithDefaults() *UnibeeApiBeanDetailUserMetric`

NewUnibeeApiBeanDetailUserMetricWithDefaults instantiates a new UnibeeApiBeanDetailUserMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiBeanDetailUserMetric) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailUserMetric) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailUserMetric) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailUserMetric) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanDetailUserMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanDetailUserMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanDetailUserMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanDetailUserMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPaid

`func (o *UnibeeApiBeanDetailUserMetric) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeApiBeanDetailUserMetric) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeApiBeanDetailUserMetric) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeApiBeanDetailUserMetric) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetLimitStats

`func (o *UnibeeApiBeanDetailUserMetric) GetLimitStats() []UnibeeApiBeanDetailUserMerchantMetricLimitStat`

GetLimitStats returns the LimitStats field if non-nil, zero value otherwise.

### GetLimitStatsOk

`func (o *UnibeeApiBeanDetailUserMetric) GetLimitStatsOk() (*[]UnibeeApiBeanDetailUserMerchantMetricLimitStat, bool)`

GetLimitStatsOk returns a tuple with the LimitStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitStats

`func (o *UnibeeApiBeanDetailUserMetric) SetLimitStats(v []UnibeeApiBeanDetailUserMerchantMetricLimitStat)`

SetLimitStats sets LimitStats field to given value.

### HasLimitStats

`func (o *UnibeeApiBeanDetailUserMetric) HasLimitStats() bool`

HasLimitStats returns a boolean if a field has been set.

### GetMeteredChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) GetMeteredChargeStats() []UnibeeApiBeanDetailUserMerchantMetricChargeStat`

GetMeteredChargeStats returns the MeteredChargeStats field if non-nil, zero value otherwise.

### GetMeteredChargeStatsOk

`func (o *UnibeeApiBeanDetailUserMetric) GetMeteredChargeStatsOk() (*[]UnibeeApiBeanDetailUserMerchantMetricChargeStat, bool)`

GetMeteredChargeStatsOk returns a tuple with the MeteredChargeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) SetMeteredChargeStats(v []UnibeeApiBeanDetailUserMerchantMetricChargeStat)`

SetMeteredChargeStats sets MeteredChargeStats field to given value.

### HasMeteredChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) HasMeteredChargeStats() bool`

HasMeteredChargeStats returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailUserMetric) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailUserMetric) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailUserMetric) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailUserMetric) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProduct

`func (o *UnibeeApiBeanDetailUserMetric) GetProduct() UnibeeApiBeanProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UnibeeApiBeanDetailUserMetric) GetProductOk() (*UnibeeApiBeanProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UnibeeApiBeanDetailUserMetric) SetProduct(v UnibeeApiBeanProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UnibeeApiBeanDetailUserMetric) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRecurringChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) GetRecurringChargeStats() []UnibeeApiBeanDetailUserMerchantMetricChargeStat`

GetRecurringChargeStats returns the RecurringChargeStats field if non-nil, zero value otherwise.

### GetRecurringChargeStatsOk

`func (o *UnibeeApiBeanDetailUserMetric) GetRecurringChargeStatsOk() (*[]UnibeeApiBeanDetailUserMerchantMetricChargeStat, bool)`

GetRecurringChargeStatsOk returns a tuple with the RecurringChargeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) SetRecurringChargeStats(v []UnibeeApiBeanDetailUserMerchantMetricChargeStat)`

SetRecurringChargeStats sets RecurringChargeStats field to given value.

### HasRecurringChargeStats

`func (o *UnibeeApiBeanDetailUserMetric) HasRecurringChargeStats() bool`

HasRecurringChargeStats returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanDetailUserMetric) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanDetailUserMetric) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanDetailUserMetric) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanDetailUserMetric) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailUserMetric) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailUserMetric) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailUserMetric) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailUserMetric) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


