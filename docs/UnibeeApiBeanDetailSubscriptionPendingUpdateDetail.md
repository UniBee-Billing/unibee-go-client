# UnibeeApiBeanDetailSubscriptionPendingUpdateDetail

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
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Note** | Pointer to **string** | Update Note | [optional] 
**Paid** | Pointer to **int32** | Paid | [optional] 
**PendingUpdateId** | Pointer to **string** | PendingUpdateId | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**ProrationAmount** | Pointer to **int64** | ProrationAmount,Cents | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | Status，1-Pending｜2-Finished｜3-Cancelled | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**UpdateAddonData** | Pointer to **string** | UpdateAddonData | [optional] 
**UpdateAddons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | UpdateAddons | [optional] 
**UpdateAmount** | Pointer to **int64** | UpdateAmount, Cents | [optional] 
**UpdateCurrency** | Pointer to **string** | UpdateCurrency | [optional] 
**UpdatePlan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**UpdatePlanId** | Pointer to **int64** | UpdatePlanId | [optional] 
**UpdateQuantity** | Pointer to **int64** | UpdateQuantity | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetail

`func NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetail() *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetail instantiates a new UnibeeApiBeanDetailSubscriptionPendingUpdateDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetailWithDefaults

`func NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

NewUnibeeApiBeanDetailSubscriptionPendingUpdateDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionPendingUpdateDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetEffectTime

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetEffectTime() int64`

GetEffectTime returns the EffectTime field if non-nil, zero value otherwise.

### GetEffectTimeOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetEffectTimeOk() (*int64, bool)`

GetEffectTimeOk returns a tuple with the EffectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectTime

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetEffectTime(v int64)`

SetEffectTime sets EffectTime field to given value.

### HasEffectTime

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasEffectTime() bool`

HasEffectTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMember

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPendingUpdateId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPendingUpdateId() string`

GetPendingUpdateId returns the PendingUpdateId field if non-nil, zero value otherwise.

### GetPendingUpdateIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPendingUpdateIdOk() (*string, bool)`

GetPendingUpdateIdOk returns a tuple with the PendingUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetPendingUpdateId(v string)`

SetPendingUpdateId sets PendingUpdateId field to given value.

### HasPendingUpdateId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasPendingUpdateId() bool`

HasPendingUpdateId returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetProrationAmount() int64`

GetProrationAmount returns the ProrationAmount field if non-nil, zero value otherwise.

### GetProrationAmountOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetProrationAmountOk() (*int64, bool)`

GetProrationAmountOk returns a tuple with the ProrationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetProrationAmount(v int64)`

SetProrationAmount sets ProrationAmount field to given value.

### HasProrationAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasProrationAmount() bool`

HasProrationAmount returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdateAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAddonData() string`

GetUpdateAddonData returns the UpdateAddonData field if non-nil, zero value otherwise.

### GetUpdateAddonDataOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAddonDataOk() (*string, bool)`

GetUpdateAddonDataOk returns a tuple with the UpdateAddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdateAddonData(v string)`

SetUpdateAddonData sets UpdateAddonData field to given value.

### HasUpdateAddonData

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdateAddonData() bool`

HasUpdateAddonData returns a boolean if a field has been set.

### GetUpdateAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAddons() []UnibeeApiBeanPlanAddonDetail`

GetUpdateAddons returns the UpdateAddons field if non-nil, zero value otherwise.

### GetUpdateAddonsOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetUpdateAddonsOk returns a tuple with the UpdateAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdateAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetUpdateAddons sets UpdateAddons field to given value.

### HasUpdateAddons

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdateAddons() bool`

HasUpdateAddons returns a boolean if a field has been set.

### GetUpdateAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAmount() int64`

GetUpdateAmount returns the UpdateAmount field if non-nil, zero value otherwise.

### GetUpdateAmountOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateAmountOk() (*int64, bool)`

GetUpdateAmountOk returns a tuple with the UpdateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdateAmount(v int64)`

SetUpdateAmount sets UpdateAmount field to given value.

### HasUpdateAmount

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdateAmount() bool`

HasUpdateAmount returns a boolean if a field has been set.

### GetUpdateCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateCurrency() string`

GetUpdateCurrency returns the UpdateCurrency field if non-nil, zero value otherwise.

### GetUpdateCurrencyOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateCurrencyOk() (*string, bool)`

GetUpdateCurrencyOk returns a tuple with the UpdateCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdateCurrency(v string)`

SetUpdateCurrency sets UpdateCurrency field to given value.

### HasUpdateCurrency

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdateCurrency() bool`

HasUpdateCurrency returns a boolean if a field has been set.

### GetUpdatePlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdatePlan() UnibeeApiBeanPlan`

GetUpdatePlan returns the UpdatePlan field if non-nil, zero value otherwise.

### GetUpdatePlanOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdatePlanOk() (*UnibeeApiBeanPlan, bool)`

GetUpdatePlanOk returns a tuple with the UpdatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdatePlan(v UnibeeApiBeanPlan)`

SetUpdatePlan sets UpdatePlan field to given value.

### HasUpdatePlan

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdatePlan() bool`

HasUpdatePlan returns a boolean if a field has been set.

### GetUpdatePlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdatePlanId() int64`

GetUpdatePlanId returns the UpdatePlanId field if non-nil, zero value otherwise.

### GetUpdatePlanIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdatePlanIdOk() (*int64, bool)`

GetUpdatePlanIdOk returns a tuple with the UpdatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdatePlanId(v int64)`

SetUpdatePlanId sets UpdatePlanId field to given value.

### HasUpdatePlanId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdatePlanId() bool`

HasUpdatePlanId returns a boolean if a field has been set.

### GetUpdateQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateQuantity() int64`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUpdateQuantityOk() (*int64, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUpdateQuantity(v int64)`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailSubscriptionPendingUpdateDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


