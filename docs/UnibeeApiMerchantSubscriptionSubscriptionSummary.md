# UnibeeApiMerchantSubscriptionSubscriptionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**CancelAtPeriodEnd** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**CurrentPeriodEnd** | Pointer to **int64** |  | [optional] 
**CurrentPeriodPaid** | Pointer to **int64** |  | [optional] 
**CurrentPeriodStart** | Pointer to **int64** |  | [optional] 
**ExternalSubscriptionId** | Pointer to **string** |  | [optional] 
**FirstPaidTime** | Pointer to **int64** |  | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LatestInvoiceId** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PendingUpdateId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **int64** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TrialEnd** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionSubscriptionSummary

`func NewUnibeeApiMerchantSubscriptionSubscriptionSummary() *UnibeeApiMerchantSubscriptionSubscriptionSummary`

NewUnibeeApiMerchantSubscriptionSubscriptionSummary instantiates a new UnibeeApiMerchantSubscriptionSubscriptionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionSubscriptionSummaryWithDefaults

`func NewUnibeeApiMerchantSubscriptionSubscriptionSummaryWithDefaults() *UnibeeApiMerchantSubscriptionSubscriptionSummary`

NewUnibeeApiMerchantSubscriptionSubscriptionSummaryWithDefaults instantiates a new UnibeeApiMerchantSubscriptionSubscriptionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCancelAtPeriodEnd() int32`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCancelAtPeriodEndOk() (*int32, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCancelAtPeriodEnd(v int32)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodPaid

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodPaid() int64`

GetCurrentPeriodPaid returns the CurrentPeriodPaid field if non-nil, zero value otherwise.

### GetCurrentPeriodPaidOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodPaidOk() (*int64, bool)`

GetCurrentPeriodPaidOk returns a tuple with the CurrentPeriodPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodPaid

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCurrentPeriodPaid(v int64)`

SetCurrentPeriodPaid sets CurrentPeriodPaid field to given value.

### HasCurrentPeriodPaid

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCurrentPeriodPaid() bool`

HasCurrentPeriodPaid returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodStart() int64`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetCurrentPeriodStartOk() (*int64, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetCurrentPeriodStart(v int64)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### GetFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetFirstPaidTime() int64`

GetFirstPaidTime returns the FirstPaidTime field if non-nil, zero value otherwise.

### GetFirstPaidTimeOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetFirstPaidTimeOk() (*int64, bool)`

GetFirstPaidTimeOk returns a tuple with the FirstPaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetFirstPaidTime(v int64)`

SetFirstPaidTime sets FirstPaidTime field to given value.

### HasFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasFirstPaidTime() bool`

HasFirstPaidTime returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestInvoiceId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetLatestInvoiceId() string`

GetLatestInvoiceId returns the LatestInvoiceId field if non-nil, zero value otherwise.

### GetLatestInvoiceIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetLatestInvoiceIdOk() (*string, bool)`

GetLatestInvoiceIdOk returns a tuple with the LatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoiceId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetLatestInvoiceId(v string)`

SetLatestInvoiceId sets LatestInvoiceId field to given value.

### HasLatestInvoiceId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasLatestInvoiceId() bool`

HasLatestInvoiceId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPendingUpdateId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetPendingUpdateId() string`

GetPendingUpdateId returns the PendingUpdateId field if non-nil, zero value otherwise.

### GetPendingUpdateIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetPendingUpdateIdOk() (*string, bool)`

GetPendingUpdateIdOk returns a tuple with the PendingUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetPendingUpdateId(v string)`

SetPendingUpdateId sets PendingUpdateId field to given value.

### HasPendingUpdateId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasPendingUpdateId() bool`

HasPendingUpdateId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionSubscriptionSummary) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


