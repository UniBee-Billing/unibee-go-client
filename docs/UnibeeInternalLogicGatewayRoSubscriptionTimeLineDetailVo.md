# UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeInternalLogicGatewayRoPlanAddonVo**](UnibeeInternalLogicGatewayRoPlanAddonVo.md) | Addon | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end | [optional] 
**PeriodEndTime** | Pointer to **string** | period end (datatime) | [optional] 
**PeriodStart** | Pointer to **int64** | period_start | [optional] 
**PeriodStartTime** | Pointer to **string** | period start (datetime) | [optional] 
**Plan** | Pointer to [**UnibeeInternalLogicGatewayRoPlanSimplify**](UnibeeInternalLogicGatewayRoPlanSimplify.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo

`func NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo() *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo instantiates a new UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVoWithDefaults

`func NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVoWithDefaults() *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo`

NewUnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetAddons() []UnibeeInternalLogicGatewayRoPlanAddonVo`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetAddonsOk() (*[]UnibeeInternalLogicGatewayRoPlanAddonVo, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetAddons(v []UnibeeInternalLogicGatewayRoPlanAddonVo)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodEndTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndTime() string`

GetPeriodEndTime returns the PeriodEndTime field if non-nil, zero value otherwise.

### GetPeriodEndTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodEndTimeOk() (*string, bool)`

GetPeriodEndTimeOk returns a tuple with the PeriodEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodEndTime(v string)`

SetPeriodEndTime sets PeriodEndTime field to given value.

### HasPeriodEndTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodEndTime() bool`

HasPeriodEndTime returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodStartTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartTime() string`

GetPeriodStartTime returns the PeriodStartTime field if non-nil, zero value otherwise.

### GetPeriodStartTimeOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPeriodStartTimeOk() (*string, bool)`

GetPeriodStartTimeOk returns a tuple with the PeriodStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPeriodStartTime(v string)`

SetPeriodStartTime sets PeriodStartTime field to given value.

### HasPeriodStartTime

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPeriodStartTime() bool`

HasPeriodStartTime returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlan() UnibeeInternalLogicGatewayRoPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanOk() (*UnibeeInternalLogicGatewayRoPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPlan(v UnibeeInternalLogicGatewayRoPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoSubscriptionTimeLineDetailVo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


