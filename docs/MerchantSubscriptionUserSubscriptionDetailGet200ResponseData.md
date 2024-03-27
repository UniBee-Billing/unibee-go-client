# MerchantSubscriptionUserSubscriptionDetailGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Plan Addon | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscriptionSimplify**](UnibeeApiBeanSubscriptionSimplify.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccountSimplify**](UnibeeApiBeanUserAccountSimplify.md) |  | [optional] 

## Methods

### NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseData

`func NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseData() *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData`

NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseData instantiates a new MerchantSubscriptionUserSubscriptionDetailGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseDataWithDefaults

`func NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseDataWithDefaults() *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData`

NewMerchantSubscriptionUserSubscriptionDetailGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionUserSubscriptionDetailGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUser() UnibeeApiBeanUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetUser(v UnibeeApiBeanUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


