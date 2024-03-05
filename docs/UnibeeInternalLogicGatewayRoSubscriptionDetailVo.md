# UnibeeInternalLogicGatewayRoSubscriptionDetailVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo**](UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo.md) | AddonParams | [optional] 
**Addons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | Addon | [optional] 
**Gateway** | Pointer to [**UnibeeInternalLogicGatewayRoGatewaySimplify**](UnibeeInternalLogicGatewayRoGatewaySimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeInternalLogicGatewayRoSubscriptionSimplify**](UnibeeInternalLogicGatewayRoSubscriptionSimplify.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo**](UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo.md) |  | [optional] 
**User** | Pointer to [**UnibeeInternalLogicGatewayRoUserAccountSimplify**](UnibeeInternalLogicGatewayRoUserAccountSimplify.md) |  | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoSubscriptionDetailVo

`func NewUnibeeInternalLogicGatewayRoSubscriptionDetailVo() *UnibeeInternalLogicGatewayRoSubscriptionDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionDetailVo instantiates a new UnibeeInternalLogicGatewayRoSubscriptionDetailVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoSubscriptionDetailVoWithDefaults

`func NewUnibeeInternalLogicGatewayRoSubscriptionDetailVoWithDefaults() *UnibeeInternalLogicGatewayRoSubscriptionDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionDetailVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoSubscriptionDetailVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetAddonParams() []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetAddonParamsOk() (*[]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetAddonParams(v []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetGateway() UnibeeInternalLogicGatewayRoGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetGatewayOk() (*UnibeeInternalLogicGatewayRoGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetGateway(v UnibeeInternalLogicGatewayRoGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetSubscription() UnibeeInternalLogicGatewayRoSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetSubscriptionOk() (*UnibeeInternalLogicGatewayRoSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetSubscription(v UnibeeInternalLogicGatewayRoSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetUnfinishedSubscriptionPendingUpdate() UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetUnfinishedSubscriptionPendingUpdate(v UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeInternalLogicGatewayRoSubscriptionDetailVo) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


