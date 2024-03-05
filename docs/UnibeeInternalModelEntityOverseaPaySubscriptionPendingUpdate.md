# UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonData** | Pointer to **string** | plan addon data (json) of this period | [optional] 
**Amount** | Pointer to **int64** | amount of this period, cent | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency of this period | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**EffectImmediate** | Pointer to **int32** | effect immediate，0-no，1-yes | [optional] 
**EffectTime** | Pointer to **int64** | effect_immediate&#x3D;0, effect time, utc_time | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayStatus** | Pointer to **string** | gateway status | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**InvoiceId** | Pointer to **string** | gateway update payment id assosiate to this update, use payment.paymentId | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**Link** | Pointer to **string** | payment link | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MerchantMemberId** | Pointer to **int64** | merchant_user_id | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Note** | Pointer to **string** | note | [optional] 
**Paid** | Pointer to **int32** | paid，0-no，1-yes | [optional] 
**PlanId** | Pointer to **int64** | the plan id of this period | [optional] 
**ProrationAmount** | Pointer to **int64** | proration amount of this pending update , cent | [optional] 
**ProrationDate** | Pointer to **int64** | merchant_user_id | [optional] 
**Quantity** | Pointer to **int64** | quantity of this period | [optional] 
**ResponseData** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | status，0-Init | 1-Create｜2-Finished｜3-Cancelled | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UpdateAddonData** | Pointer to **string** | plan addon data (json) after update | [optional] 
**UpdateAmount** | Pointer to **int64** | the amount after update | [optional] 
**UpdateCurrency** | Pointer to **string** | the currency after update | [optional] 
**UpdatePlanId** | Pointer to **int64** | the plan id after update | [optional] 
**UpdateQuantity** | Pointer to **int64** | quantity after update | [optional] 
**UpdateSubscriptionId** | Pointer to **string** | pending update unique id | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate

`func NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate() *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate`

NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate instantiates a new UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdateWithDefaults

`func NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdateWithDefaults() *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate`

NewUnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdateWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetEffectTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetEffectTime() int64`

GetEffectTime returns the EffectTime field if non-nil, zero value otherwise.

### GetEffectTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetEffectTimeOk() (*int64, bool)`

GetEffectTimeOk returns a tuple with the EffectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetEffectTime(v int64)`

SetEffectTime sets EffectTime field to given value.

### HasEffectTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasEffectTime() bool`

HasEffectTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMerchantMemberId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetMerchantMemberId() int64`

GetMerchantMemberId returns the MerchantMemberId field if non-nil, zero value otherwise.

### GetMerchantMemberIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetMerchantMemberIdOk() (*int64, bool)`

GetMerchantMemberIdOk returns a tuple with the MerchantMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetMerchantMemberId(v int64)`

SetMerchantMemberId sets MerchantMemberId field to given value.

### HasMerchantMemberId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasMerchantMemberId() bool`

HasMerchantMemberId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetPaid() int32`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetPaidOk() (*int32, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetPaid(v int32)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProrationAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetProrationAmount() int64`

GetProrationAmount returns the ProrationAmount field if non-nil, zero value otherwise.

### GetProrationAmountOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetProrationAmountOk() (*int64, bool)`

GetProrationAmountOk returns a tuple with the ProrationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetProrationAmount(v int64)`

SetProrationAmount sets ProrationAmount field to given value.

### HasProrationAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasProrationAmount() bool`

HasProrationAmount returns a boolean if a field has been set.

### GetProrationDate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetResponseData() string`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetResponseDataOk() (*string, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetResponseData(v string)`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdateAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateAddonData() string`

GetUpdateAddonData returns the UpdateAddonData field if non-nil, zero value otherwise.

### GetUpdateAddonDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateAddonDataOk() (*string, bool)`

GetUpdateAddonDataOk returns a tuple with the UpdateAddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdateAddonData(v string)`

SetUpdateAddonData sets UpdateAddonData field to given value.

### HasUpdateAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdateAddonData() bool`

HasUpdateAddonData returns a boolean if a field has been set.

### GetUpdateAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateAmount() int64`

GetUpdateAmount returns the UpdateAmount field if non-nil, zero value otherwise.

### GetUpdateAmountOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateAmountOk() (*int64, bool)`

GetUpdateAmountOk returns a tuple with the UpdateAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdateAmount(v int64)`

SetUpdateAmount sets UpdateAmount field to given value.

### HasUpdateAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdateAmount() bool`

HasUpdateAmount returns a boolean if a field has been set.

### GetUpdateCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateCurrency() string`

GetUpdateCurrency returns the UpdateCurrency field if non-nil, zero value otherwise.

### GetUpdateCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateCurrencyOk() (*string, bool)`

GetUpdateCurrencyOk returns a tuple with the UpdateCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdateCurrency(v string)`

SetUpdateCurrency sets UpdateCurrency field to given value.

### HasUpdateCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdateCurrency() bool`

HasUpdateCurrency returns a boolean if a field has been set.

### GetUpdatePlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdatePlanId() int64`

GetUpdatePlanId returns the UpdatePlanId field if non-nil, zero value otherwise.

### GetUpdatePlanIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdatePlanIdOk() (*int64, bool)`

GetUpdatePlanIdOk returns a tuple with the UpdatePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdatePlanId(v int64)`

SetUpdatePlanId sets UpdatePlanId field to given value.

### HasUpdatePlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdatePlanId() bool`

HasUpdatePlanId returns a boolean if a field has been set.

### GetUpdateQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateQuantity() int64`

GetUpdateQuantity returns the UpdateQuantity field if non-nil, zero value otherwise.

### GetUpdateQuantityOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateQuantityOk() (*int64, bool)`

GetUpdateQuantityOk returns a tuple with the UpdateQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdateQuantity(v int64)`

SetUpdateQuantity sets UpdateQuantity field to given value.

### HasUpdateQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdateQuantity() bool`

HasUpdateQuantity returns a boolean if a field has been set.

### GetUpdateSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateSubscriptionId() string`

GetUpdateSubscriptionId returns the UpdateSubscriptionId field if non-nil, zero value otherwise.

### GetUpdateSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUpdateSubscriptionIdOk() (*string, bool)`

GetUpdateSubscriptionIdOk returns a tuple with the UpdateSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUpdateSubscriptionId(v string)`

SetUpdateSubscriptionId sets UpdateSubscriptionId field to given value.

### HasUpdateSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUpdateSubscriptionId() bool`

HasUpdateSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscriptionPendingUpdate) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


