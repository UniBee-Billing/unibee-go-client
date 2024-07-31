# UnibeeApiMerchantSubscriptionDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | AddonParams | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Plan Addon | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGateway**](UnibeeApiBeanGateway.md) |  | [optional] 
**LatestInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

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

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGateway() UnibeeApiBeanGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetGatewayOk() (*UnibeeApiBeanGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetGateway(v UnibeeApiBeanGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetLatestInvoice

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetLatestInvoice() UnibeeApiBeanInvoice`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetLatestInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetLatestInvoice(v UnibeeApiBeanInvoice)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantSubscriptionDetailRes) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantSubscriptionDetailRes) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


