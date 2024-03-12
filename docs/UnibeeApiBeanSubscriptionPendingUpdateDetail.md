# UnibeeApiBeanSubscriptionPendingUpdateDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonData** | Pointer to **string** | plan addon json data | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addons | [optional] 
**Amount** | Pointer to **int64** | CaptureAmount, Cent | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**EffectImmediate** | Pointer to **int32** | EffectImmediate | [optional] 
**EffectTime** | Pointer to **int64** | effect_immediate&#x3D;0, EffectTime unit_time | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**GmtCreate** | Pointer to **string** | GmtCreate | [optional] 
**GmtModify** | Pointer to **string** | GmtModify | [optional] 
**Link** | Pointer to **string** | Link | [optional] 
**MerchantId** | Pointer to **int64** | MerchantId | [optional] 
**MerchantMember** | Pointer to [**UnibeeApiBeanMerchantMemberSimplify**](UnibeeApiBeanMerchantMemberSimplify.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Note** | Pointer to **string** | Update Note | [optional] 
**Paid** | Pointer to **int32** | Paid | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**ProrationAmount** | Pointer to **int64** | ProrationAmount,Cents | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | Status，0-Init | 1-Create｜2-Finished｜3-Cancelled | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**UpdateAddonData** | Pointer to **string** | UpdateAddonData | [optional] 
**UpdateAddons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | UpdateAddons | [optional] 
**UpdateAmount** | Pointer to **int64** | UpdateAmount, Cents | [optional] 
**UpdateCurrency** | Pointer to **string** | UpdateCurrency | [optional] 
**UpdatePlan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**UpdatePlanId** | Pointer to **int64** | UpdatePlanId | [optional] 
**UpdateQuantity** | Pointer to **int64** | UpdateQuantity | [optional] 
**UpdateSubscriptionId** | Pointer to **string** | UpdateSubscriptionId | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiBeanSubscriptionPendingUpdateDetail

`func NewUnibeeApiBeanSubscriptionPendingUpdateDetail() *UnibeeApiBeanSubscriptionPendingUpdateDetail`

NewUnibeeApiBeanSubscriptionPendingUpdateDetail instantiates a new UnibeeApiBeanSubscriptionPendingUpdateDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionPendingUpdateDetailWithDefaults

`func NewUnibeeApiBeanSubscriptionPendingUpdateDetailWithDefaults() *UnibeeApiBeanSubscriptionPendingUpdateDetail`

NewUnibeeApiBeanSubscriptionPendingUpdateDetailWithDefaults instantiates a new UnibeeApiBeanSubscriptionPendingUpdateDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetEffectTime

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetEffectTime() int64`

GetEffectTime returns the EffectTime field if non-nil, zero value otherwise.

### GetEffectTimeOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetEffectTimeOk() (*int64, bool)`

GetEffectTimeOk returns a tuple with the EffectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectTime

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetEffectTime(v int64)`

SetEffectTime sets EffectTime field to given value.

### HasEffectTime

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasEffectTime() bool`

HasEffectTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMember

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMerchantMember() UnibeeApiBeanMerchantMemberSimplify`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMerchantMemberOk() (*UnibeeApiBeanMerchantMemberSimplify, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetMerchantMember(v UnibeeApiBeanMerchantMemberSimplify)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetProrationAmount() int64`

GetProrationAmount returns the ProrationAmount field if non-nil, zero value otherwise.

### GetProrationAmountOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetProrationAmountOk() (*int64, bool)`

GetProrationAmountOk returns a tuple with the ProrationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetProrationAmount(v int64)`

SetProrationAmount sets ProrationAmount field to given value.

### HasProrationAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasProrationAmount() bool`

HasProrationAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdateAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAddonData() string`

GetUpdateAddonData returns the UpdateAddonData field if non-nil, zero value otherwise.

### GetUpdateAddonDataOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAddonDataOk() (*string, bool)`

GetUpdateAddonDataOk returns a tuple with the UpdateAddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateAddonData(v string)`

SetUpdateAddonData sets UpdateAddonData field to given value.

### HasUpdateAddonData

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateAddonData() bool`

HasUpdateAddonData returns a boolean if a field has been set.

### GetUpdateAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAddons() []UnibeeApiBeanPlanAddonDetail`

GetUpdateAddons returns the UpdateAddons field if non-nil, zero value otherwise.

### GetUpdateAddonsOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetUpdateAddonsOk returns a tuple with the UpdateAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetUpdateAddons sets UpdateAddons field to given value.

### HasUpdateAddons

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateAddons() bool`

HasUpdateAddons returns a boolean if a field has been set.

### GetUpdateAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAmount() int64`

GetUpdateAmount returns the UpdateAmount field if non-nil, zero value otherwise.

### GetUpdateAmountOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateAmountOk() (*int64, bool)`

GetUpdateAmountOk returns a tuple with the UpdateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateAmount(v int64)`

SetUpdateAmount sets UpdateAmount field to given value.

### HasUpdateAmount

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateAmount() bool`

HasUpdateAmount returns a boolean if a field has been set.

### GetUpdateCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateCurrency() string`

GetUpdateCurrency returns the UpdateCurrency field if non-nil, zero value otherwise.

### GetUpdateCurrencyOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateCurrencyOk() (*string, bool)`

GetUpdateCurrencyOk returns a tuple with the UpdateCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateCurrency(v string)`

SetUpdateCurrency sets UpdateCurrency field to given value.

### HasUpdateCurrency

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateCurrency() bool`

HasUpdateCurrency returns a boolean if a field has been set.

### GetUpdatePlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdatePlan() UnibeeApiBeanPlanSimplify`

GetUpdatePlan returns the UpdatePlan field if non-nil, zero value otherwise.

### GetUpdatePlanOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdatePlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetUpdatePlanOk returns a tuple with the UpdatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdatePlan(v UnibeeApiBeanPlanSimplify)`

SetUpdatePlan sets UpdatePlan field to given value.

### HasUpdatePlan

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdatePlan() bool`

HasUpdatePlan returns a boolean if a field has been set.

### GetUpdatePlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdatePlanId() int64`

GetUpdatePlanId returns the UpdatePlanId field if non-nil, zero value otherwise.

### GetUpdatePlanIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdatePlanIdOk() (*int64, bool)`

GetUpdatePlanIdOk returns a tuple with the UpdatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdatePlanId(v int64)`

SetUpdatePlanId sets UpdatePlanId field to given value.

### HasUpdatePlanId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdatePlanId() bool`

HasUpdatePlanId returns a boolean if a field has been set.

### GetUpdateQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateQuantity() int64`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateQuantityOk() (*int64, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateQuantity(v int64)`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### GetUpdateSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateSubscriptionId() string`

GetUpdateSubscriptionId returns the UpdateSubscriptionId field if non-nil, zero value otherwise.

### GetUpdateSubscriptionIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUpdateSubscriptionIdOk() (*string, bool)`

GetUpdateSubscriptionIdOk returns a tuple with the UpdateSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUpdateSubscriptionId(v string)`

SetUpdateSubscriptionId sets UpdateSubscriptionId field to given value.

### HasUpdateSubscriptionId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUpdateSubscriptionId() bool`

HasUpdateSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanSubscriptionPendingUpdateDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


