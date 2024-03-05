# UnibeeInternalLogicGatewayRoUserMetricStat

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

### NewUnibeeInternalLogicGatewayRoUserMetricStat

`func NewUnibeeInternalLogicGatewayRoUserMetricStat() *UnibeeInternalLogicGatewayRoUserMetricStat`

NewUnibeeInternalLogicGatewayRoUserMetricStat instantiates a new UnibeeInternalLogicGatewayRoUserMetricStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoUserMetricStatWithDefaults

`func NewUnibeeInternalLogicGatewayRoUserMetricStatWithDefaults() *UnibeeInternalLogicGatewayRoUserMetricStat`

NewUnibeeInternalLogicGatewayRoUserMetricStatWithDefaults instantiates a new UnibeeInternalLogicGatewayRoUserMetricStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetIsPaid

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetSubscription() UnibeeInternalLogicGatewayRoSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetSubscriptionOk() (*UnibeeInternalLogicGatewayRoSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetSubscription(v UnibeeInternalLogicGatewayRoSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserMerchantMetricStats

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetUserMerchantMetricStats() []UnibeeInternalLogicGatewayRoUserMerchantMetricStat`

GetUserMerchantMetricStats returns the UserMerchantMetricStats field if non-nil, zero value otherwise.

### GetUserMerchantMetricStatsOk

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) GetUserMerchantMetricStatsOk() (*[]UnibeeInternalLogicGatewayRoUserMerchantMetricStat, bool)`

GetUserMerchantMetricStatsOk returns a tuple with the UserMerchantMetricStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMerchantMetricStats

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) SetUserMerchantMetricStats(v []UnibeeInternalLogicGatewayRoUserMerchantMetricStat)`

SetUserMerchantMetricStats sets UserMerchantMetricStats field to given value.

### HasUserMerchantMetricStats

`func (o *UnibeeInternalLogicGatewayRoUserMetricStat) HasUserMerchantMetricStats() bool`

HasUserMerchantMetricStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


