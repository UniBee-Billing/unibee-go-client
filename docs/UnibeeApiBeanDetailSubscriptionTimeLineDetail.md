# UnibeeApiBeanDetailSubscriptionTimeLineDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addon | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end | [optional] 
**PeriodEndTime** | Pointer to **string** | period end (datatime) | [optional] 
**PeriodStart** | Pointer to **int64** | period_start | [optional] 
**PeriodStartTime** | Pointer to **string** | period start (datetime) | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**PlanId** | Pointer to **int64** | PlanId | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | 1-processing,2-finish | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeApiBeanDetailSubscriptionTimeLineDetail

`func NewUnibeeApiBeanDetailSubscriptionTimeLineDetail() *UnibeeApiBeanDetailSubscriptionTimeLineDetail`

NewUnibeeApiBeanDetailSubscriptionTimeLineDetail instantiates a new UnibeeApiBeanDetailSubscriptionTimeLineDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailSubscriptionTimeLineDetailWithDefaults

`func NewUnibeeApiBeanDetailSubscriptionTimeLineDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionTimeLineDetail`

NewUnibeeApiBeanDetailSubscriptionTimeLineDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionTimeLineDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodEndTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndTime() string`

GetPeriodEndTime returns the PeriodEndTime field if non-nil, zero value otherwise.

### GetPeriodEndTimeOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodEndTimeOk() (*string, bool)`

GetPeriodEndTimeOk returns a tuple with the PeriodEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodEndTime(v string)`

SetPeriodEndTime sets PeriodEndTime field to given value.

### HasPeriodEndTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodEndTime() bool`

HasPeriodEndTime returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodStartTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartTime() string`

GetPeriodStartTime returns the PeriodStartTime field if non-nil, zero value otherwise.

### GetPeriodStartTimeOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPeriodStartTimeOk() (*string, bool)`

GetPeriodStartTimeOk returns a tuple with the PeriodStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPeriodStartTime(v string)`

SetPeriodStartTime sets PeriodStartTime field to given value.

### HasPeriodStartTime

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPeriodStartTime() bool`

HasPeriodStartTime returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailSubscriptionTimeLineDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


