# MerchantSubscriptionUserSubscriptionDetailGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Plan Addon | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**LatestInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**NextInvoiceApply** | Pointer to [**UnibeeApiBeanDetailSubscriptionNextInvoiceApply**](UnibeeApiBeanDetailSubscriptionNextInvoiceApply.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**UnfinishedSubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

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

### GetDiscount

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetLatestInvoice

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetLatestInvoice() UnibeeApiBeanInvoice`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetLatestInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetLatestInvoice(v UnibeeApiBeanInvoice)`

SetLatestInvoice sets LatestInvoice field to given value.

### HasLatestInvoice

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasLatestInvoice() bool`

HasLatestInvoice returns a boolean if a field has been set.

### GetNextInvoiceApply

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetNextInvoiceApply() UnibeeApiBeanDetailSubscriptionNextInvoiceApply`

GetNextInvoiceApply returns the NextInvoiceApply field if non-nil, zero value otherwise.

### GetNextInvoiceApplyOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetNextInvoiceApplyOk() (*UnibeeApiBeanDetailSubscriptionNextInvoiceApply, bool)`

GetNextInvoiceApplyOk returns a tuple with the NextInvoiceApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceApply

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetNextInvoiceApply(v UnibeeApiBeanDetailSubscriptionNextInvoiceApply)`

SetNextInvoiceApply sets NextInvoiceApply field to given value.

### HasNextInvoiceApply

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasNextInvoiceApply() bool`

HasNextInvoiceApply returns a boolean if a field has been set.

### GetPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetSubscription(v UnibeeApiBeanSubscription)`

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

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *MerchantSubscriptionUserSubscriptionDetailGet200ResponseData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


