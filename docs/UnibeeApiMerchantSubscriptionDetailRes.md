# UnibeeApiMerchantSubscriptionDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | Plan Addon | [optional] 
**Gateway** | Pointer to [**UnibeeInternalLogicGatewayRoGatewaySimplify**](UnibeeInternalLogicGatewayRoGatewaySimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeInternalLogicGatewayRoSubscriptionSimplify**](UnibeeInternalLogicGatewayRoSubscriptionSimplify.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo**](UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo.md) |  | [optional] 
**User** | Pointer to [**UnibeeInternalLogicGatewayRoUserAccountSimplify**](UnibeeInternalLogicGatewayRoUserAccountSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionDetailRes

`func NewUnibeeApiMerchantSubscriptionDetailRes() *UnibeeApiMerchantSubscriptionDetailRes`

NewUnibeeApiMerchantSubscriptionDetailRes instantiates a new UnibeeApiMerchantSubscriptionDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionDetailResWithDefaults

`func NewUnibeeApiMerchantSubscriptionDetailResWithDefaults() *UnibeeApiMerchantSubscriptionDetailRes`

NewUnibeeApiMerchantSubscriptionDetailResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGateway() UnibeeInternalLogicGatewayRoGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGatewayOk() (*UnibeeInternalLogicGatewayRoGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetGateway(v UnibeeInternalLogicGatewayRoGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscription() UnibeeInternalLogicGatewayRoSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscriptionOk() (*UnibeeInternalLogicGatewayRoSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetSubscription(v UnibeeInternalLogicGatewayRoSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdate() UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUnfinishedSubscriptionPendingUpdate(v UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUser() UnibeeInternalLogicGatewayRoUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUserOk() (*UnibeeInternalLogicGatewayRoUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUser(v UnibeeInternalLogicGatewayRoUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


