# UnibeeApiBeanDetailSubscriptionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | AddonParams | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addon | [optional] 
**DayLeft** | Pointer to **int32** | DayLeft util the period end, only available for webhook | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**LatestInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**Note** | Pointer to **string** | note | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailSubscriptionDetail

`func NewUnibeeApiBeanDetailSubscriptionDetail() *UnibeeApiBeanDetailSubscriptionDetail`

NewUnibeeApiBeanDetailSubscriptionDetail instantiates a new UnibeeApiBeanDetailSubscriptionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailSubscriptionDetailWithDefaults

`func NewUnibeeApiBeanDetailSubscriptionDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionDetail`

NewUnibeeApiBeanDetailSubscriptionDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetDayLeft

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDayLeft() int32`

GetDayLeft returns the DayLeft field if non-nil, zero value otherwise.

### GetDayLeftOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDayLeftOk() (*int32, bool)`

GetDayLeftOk returns a tuple with the DayLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayLeft

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetDayLeft(v int32)`

SetDayLeft sets DayLeft field to given value.

### HasDayLeft

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasDayLeft() bool`

HasDayLeft returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetLatestInvoice

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetLatestInvoice() UnibeeApiBeanInvoice`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetLatestInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetLatestInvoice(v UnibeeApiBeanInvoice)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

GetUnfinishedSubscriptionPendingUpdate returns the UnfinishedSubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetUnfinishedSubscriptionPendingUpdateOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUnfinishedSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool)`

GetUnfinishedSubscriptionPendingUpdateOk returns a tuple with the UnfinishedSubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetUnfinishedSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail)`

SetUnfinishedSubscriptionPendingUpdate sets UnfinishedSubscriptionPendingUpdate field to given value.

### HasUnfinishedSubscriptionPendingUpdate

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasUnfinishedSubscriptionPendingUpdate() bool`

HasUnfinishedSubscriptionPendingUpdate returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailSubscriptionDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailSubscriptionDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailSubscriptionDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


