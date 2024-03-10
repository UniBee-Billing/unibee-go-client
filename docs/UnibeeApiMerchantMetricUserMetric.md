# UnibeeApiMerchantMetricUserMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | Addon | [optional] 
**IsPaid** | Pointer to **bool** | IsPaid | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeInternalLogicGatewayRoSubscriptionSimplify**](UnibeeInternalLogicGatewayRoSubscriptionSimplify.md) |  | [optional] 
**User** | Pointer to [**UnibeeInternalLogicGatewayRoUserAccountSimplify**](UnibeeInternalLogicGatewayRoUserAccountSimplify.md) |  | [optional] 
**UserMerchantMetricStats** | Pointer to [**[]UnibeeInternalLogicGatewayRoUserMerchantMetricStat**](UnibeeInternalLogicGatewayRoUserMerchantMetricStat.md) | UserMerchantMetricStats | [optional] 

## Methods

### NewUnibeeApiMerchantMetricUserMetric

`func NewUnibeeApiMerchantMetricUserMetric() *UnibeeApiMerchantMetricUserMetric`

NewUnibeeApiMerchantMetricUserMetric instantiates a new UnibeeApiMerchantMetricUserMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricUserMetricWithDefaults

`func NewUnibeeApiMerchantMetricUserMetricWithDefaults() *UnibeeApiMerchantMetricUserMetric`

NewUnibeeApiMerchantMetricUserMetricWithDefaults instantiates a new UnibeeApiMerchantMetricUserMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiMerchantMetricUserMetric) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiMerchantMetricUserMetric) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiMerchantMetricUserMetric) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeApiMerchantMetricUserMetric) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiMerchantMetricUserMetric) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiMerchantMetricUserMetric) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiMerchantMetricUserMetric) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) GetSubscription() UnibeeInternalLogicGatewayRoSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetSubscriptionOk() (*UnibeeInternalLogicGatewayRoSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) SetSubscription(v UnibeeInternalLogicGatewayRoSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiMerchantMetricUserMetric) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantMetricUserMetric) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantMetricUserMetric) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantMetricUserMetric) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStats() []UnibeeInternalLogicGatewayRoUserMerchantMetricStat`

GetUserMerchantMetricStats returns the UserMerchantMetricStats field if non-nil, zero value otherwise.

### GetUserMerchantMetricStatsOk

`func (o *UnibeeApiMerchantMetricUserMetric) GetUserMerchantMetricStatsOk() (*[]UnibeeInternalLogicGatewayRoUserMerchantMetricStat, bool)`

GetUserMerchantMetricStatsOk returns a tuple with the UserMerchantMetricStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) SetUserMerchantMetricStats(v []UnibeeInternalLogicGatewayRoUserMerchantMetricStat)`

SetUserMerchantMetricStats sets UserMerchantMetricStats field to given value.

### HasUserMerchantMetricStats

`func (o *UnibeeApiMerchantMetricUserMetric) HasUserMerchantMetricStats() bool`

HasUserMerchantMetricStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


