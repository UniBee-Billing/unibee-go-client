# UnibeeApiBeanSubscriptionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | AddonParams | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addon | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscriptionSimplify**](UnibeeApiBeanSubscriptionSimplify.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanSubscriptionPendingUpdateDetail**](UnibeeApiBeanSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccountSimplify**](UnibeeApiBeanUserAccountSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanSubscriptionDetail

`func NewUnibeeApiBeanSubscriptionDetail() *UnibeeApiBeanSubscriptionDetail`

NewUnibeeApiBeanSubscriptionDetail instantiates a new UnibeeApiBeanSubscriptionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionDetailWithDefaults

`func NewUnibeeApiBeanSubscriptionDetailWithDefaults() *UnibeeApiBeanSubscriptionDetail`

NewUnibeeApiBeanSubscriptionDetailWithDefaults instantiates a new UnibeeApiBeanSubscriptionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiBeanSubscriptionDetail) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiBeanSubscriptionDetail) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiBeanSubscriptionDetail) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanSubscriptionDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanSubscriptionDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanSubscriptionDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanSubscriptionDetail) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanSubscriptionDetail) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanSubscriptionDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanSubscriptionDetail) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanSubscriptionDetail) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanSubscriptionDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanSubscriptionDetail) GetSubscription() UnibeeApiBeanSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanSubscriptionDetail) SetSubscription(v UnibeeApiBeanSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanSubscriptionDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanSubscriptionPendingUpdateDetail`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanSubscriptionPendingUpdateDetail, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanSubscriptionDetail) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanSubscriptionPendingUpdateDetail)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanSubscriptionDetail) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanSubscriptionDetail) GetUser() UnibeeApiBeanUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanSubscriptionDetail) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanSubscriptionDetail) SetUser(v UnibeeApiBeanUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanSubscriptionDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


