# UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonData** | Pointer to **string** | plan addon json data | [optional] 
**Addons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | Addons | [optional] 
**Amount** | Pointer to **int64** | CaptureAmount, Cent | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**EffectImmediate** | Pointer to **int32** | EffectImmediate | [optional] 
**EffectTime** | Pointer to **int64** | effect_immediate&#x3D;0, EffectTime unit_time | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**GmtCreate** | Pointer to **string** | GmtCreate | [optional] 
**GmtModify** | Pointer to **string** | GmtModify | [optional] 
**Link** | Pointer to **string** | Link | [optional] 
**MerchantId** | Pointer to **int64** | MerchantId | [optional] 
**MerchantMember** | Pointer to [**UnibeeInternalLogicGatewayRoMerchantMemberSimplify**](UnibeeInternalLogicGatewayRoMerchantMemberSimplify.md) |  | [optional] 
**Note** | Pointer to **string** | Update Note | [optional] 
**Paid** | Pointer to **int32** | Paid | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**ProrationAmount** | Pointer to **int64** | ProrationAmount,Cents | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | Status，0-Init | 1-Create｜2-Finished｜3-Cancelled | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**UpdateAddonData** | Pointer to **string** | UpdateAddonData | [optional] 
**UpdateAddons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | UpdateAddons | [optional] 
**UpdateAmount** | Pointer to **int64** | UpdateAmount, Cents | [optional] 
**UpdateCurrency** | Pointer to **string** | UpdateCurrency | [optional] 
**UpdatePlan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**UpdatePlanId** | Pointer to **int64** | UpdatePlanId | [optional] 
**UpdateQuantity** | Pointer to **int64** | UpdateQuantity | [optional] 
**UpdateSubscriptionId** | Pointer to **string** | UpdateSubscriptionId | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo

`func NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo() *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo instantiates a new UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVoWithDefaults

`func NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVoWithDefaults() *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetEffectTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetEffectTime() int64`

GetEffectTime returns the EffectTime field if non-nil, zero value otherwise.

### GetEffectTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetEffectTimeOk() (*int64, bool)`

GetEffectTimeOk returns a tuple with the EffectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetEffectTime(v int64)`

SetEffectTime sets EffectTime field to given value.

### HasEffectTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasEffectTime() bool`

HasEffectTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMember

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetMerchantMember() UnibeeInternalLogicGatewayRoMerchantMemberSimplify`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetMerchantMemberOk() (*UnibeeInternalLogicGatewayRoMerchantMemberSimplify, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetMerchantMember(v UnibeeInternalLogicGatewayRoMerchantMemberSimplify)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetProrationAmount() int64`

GetProrationAmount returns the ProrationAmount field if non-nil, zero value otherwise.

### GetProrationAmountOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetProrationAmountOk() (*int64, bool)`

GetProrationAmountOk returns a tuple with the ProrationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetProrationAmount(v int64)`

SetProrationAmount sets ProrationAmount field to given value.

### HasProrationAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasProrationAmount() bool`

HasProrationAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdateAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAddonData() string`

GetUpdateAddonData returns the UpdateAddonData field if non-nil, zero value otherwise.

### GetUpdateAddonDataOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAddonDataOk() (*string, bool)`

GetUpdateAddonDataOk returns a tuple with the UpdateAddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateAddonData(v string)`

SetUpdateAddonData sets UpdateAddonData field to given value.

### HasUpdateAddonData

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateAddonData() bool`

HasUpdateAddonData returns a boolean if a field has been set.

### GetUpdateAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetUpdateAddons returns the UpdateAddons field if non-nil, zero value otherwise.

### GetUpdateAddonsOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetUpdateAddonsOk returns a tuple with the UpdateAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetUpdateAddons sets UpdateAddons field to given value.

### HasUpdateAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateAddons() bool`

HasUpdateAddons returns a boolean if a field has been set.

### GetUpdateAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAmount() int64`

GetUpdateAmount returns the UpdateAmount field if non-nil, zero value otherwise.

### GetUpdateAmountOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateAmountOk() (*int64, bool)`

GetUpdateAmountOk returns a tuple with the UpdateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateAmount(v int64)`

SetUpdateAmount sets UpdateAmount field to given value.

### HasUpdateAmount

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateAmount() bool`

HasUpdateAmount returns a boolean if a field has been set.

### GetUpdateCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateCurrency() string`

GetUpdateCurrency returns the UpdateCurrency field if non-nil, zero value otherwise.

### GetUpdateCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateCurrencyOk() (*string, bool)`

GetUpdateCurrencyOk returns a tuple with the UpdateCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateCurrency(v string)`

SetUpdateCurrency sets UpdateCurrency field to given value.

### HasUpdateCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateCurrency() bool`

HasUpdateCurrency returns a boolean if a field has been set.

### GetUpdatePlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdatePlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetUpdatePlan returns the UpdatePlan field if non-nil, zero value otherwise.

### GetUpdatePlanOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdatePlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetUpdatePlanOk returns a tuple with the UpdatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdatePlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetUpdatePlan sets UpdatePlan field to given value.

### HasUpdatePlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdatePlan() bool`

HasUpdatePlan returns a boolean if a field has been set.

### GetUpdatePlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdatePlanId() int64`

GetUpdatePlanId returns the UpdatePlanId field if non-nil, zero value otherwise.

### GetUpdatePlanIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdatePlanIdOk() (*int64, bool)`

GetUpdatePlanIdOk returns a tuple with the UpdatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdatePlanId(v int64)`

SetUpdatePlanId sets UpdatePlanId field to given value.

### HasUpdatePlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdatePlanId() bool`

HasUpdatePlanId returns a boolean if a field has been set.

### GetUpdateQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateQuantity() int64`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateQuantityOk() (*int64, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateQuantity(v int64)`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### GetUpdateSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateSubscriptionId() string`

GetUpdateSubscriptionId returns the UpdateSubscriptionId field if non-nil, zero value otherwise.

### GetUpdateSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUpdateSubscriptionIdOk() (*string, bool)`

GetUpdateSubscriptionIdOk returns a tuple with the UpdateSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUpdateSubscriptionId(v string)`

SetUpdateSubscriptionId sets UpdateSubscriptionId field to given value.

### HasUpdateSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUpdateSubscriptionId() bool`

HasUpdateSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionPendingUpdateDetailVo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


