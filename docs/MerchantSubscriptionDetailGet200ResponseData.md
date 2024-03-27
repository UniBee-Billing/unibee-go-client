# MerchantSubscriptionDetailGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | AddonParams | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Plan Addon | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) |  | [optional] 
**LatestInvoice** | Pointer to [**UnibeeApiBeanInvoiceSimplify**](UnibeeApiBeanInvoiceSimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscriptionSimplify**](UnibeeApiBeanSubscriptionSimplify.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccountSimplify**](UnibeeApiBeanUserAccountSimplify.md) |  | [optional] 

## Methods

### NewMerchantSubscriptionDetailGet200ResponseData

`func NewMerchantSubscriptionDetailGet200ResponseData() *MerchantSubscriptionDetailGet200ResponseData`

NewMerchantSubscriptionDetailGet200ResponseData instantiates a new MerchantSubscriptionDetailGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionDetailGet200ResponseDataWithDefaults

`func NewMerchantSubscriptionDetailGet200ResponseDataWithDefaults() *MerchantSubscriptionDetailGet200ResponseData`

NewMerchantSubscriptionDetailGet200ResponseDataWithDefaults instantiates a new MerchantSubscriptionDetailGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetGateway

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetLatestInvoice

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetLatestInvoice() UnibeeApiBeanInvoiceSimplify`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetLatestInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetLatestInvoice(v UnibeeApiBeanInvoiceSimplify)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetPlan

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetUser() UnibeeApiBeanUserAccountSimplify`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MerchantSubscriptionDetailGet200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccountSimplify, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MerchantSubscriptionDetailGet200ResponseData) SetUser(v UnibeeApiBeanUserAccountSimplify)`

SetUser sets User field to given value.

### HasUser

`func (o *MerchantSubscriptionDetailGet200ResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


