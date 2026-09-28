# UnibeeApiMerchantPlanActivePriceChangePreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAffectedSubscriptions** | Pointer to **int64** | Active/processing subscriptions whose future billing will be affected | [optional] 
**NewAmount** | Pointer to **int64** | New amount of plan | [optional] 
**NewCurrency** | Pointer to **string** |  | [optional] 
**OldAmount** | Pointer to **int64** | Current amount of plan | [optional] 
**OldCurrency** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **int64** | Target planId | [optional] 
**TotalAffectedSubscriptions** | Pointer to **int64** | Total subscriptions that have used or are using this plan (by status set) | [optional] 
**UnsupportedCurrencies** | Pointer to **[]string** | Currencies used by active subscriptions or pending updates that cannot be priced after the change | [optional] 

## Methods

### NewUnibeeApiMerchantPlanActivePriceChangePreviewRes

`func NewUnibeeApiMerchantPlanActivePriceChangePreviewRes() *UnibeeApiMerchantPlanActivePriceChangePreviewRes`

NewUnibeeApiMerchantPlanActivePriceChangePreviewRes instantiates a new UnibeeApiMerchantPlanActivePriceChangePreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanActivePriceChangePreviewResWithDefaults

`func NewUnibeeApiMerchantPlanActivePriceChangePreviewResWithDefaults() *UnibeeApiMerchantPlanActivePriceChangePreviewRes`

NewUnibeeApiMerchantPlanActivePriceChangePreviewResWithDefaults instantiates a new UnibeeApiMerchantPlanActivePriceChangePreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetActiveAffectedSubscriptions() int64`

GetActiveAffectedSubscriptions returns the ActiveAffectedSubscriptions field if non-nil, zero value otherwise.

### GetActiveAffectedSubscriptionsOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetActiveAffectedSubscriptionsOk() (*int64, bool)`

GetActiveAffectedSubscriptionsOk returns a tuple with the ActiveAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetActiveAffectedSubscriptions(v int64)`

SetActiveAffectedSubscriptions sets ActiveAffectedSubscriptions field to given value.

### HasActiveAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasActiveAffectedSubscriptions() bool`

HasActiveAffectedSubscriptions returns a boolean if a field has been set.

### GetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetNewAmount() int64`

GetNewAmount returns the NewAmount field if non-nil, zero value otherwise.

### GetNewAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetNewAmountOk() (*int64, bool)`

GetNewAmountOk returns a tuple with the NewAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetNewAmount(v int64)`

SetNewAmount sets NewAmount field to given value.

### HasNewAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasNewAmount() bool`

HasNewAmount returns a boolean if a field has been set.

### GetNewCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetNewCurrency() string`

GetNewCurrency returns the NewCurrency field if non-nil, zero value otherwise.

### GetNewCurrencyOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetNewCurrencyOk() (*string, bool)`

GetNewCurrencyOk returns a tuple with the NewCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetNewCurrency(v string)`

SetNewCurrency sets NewCurrency field to given value.

### HasNewCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasNewCurrency() bool`

HasNewCurrency returns a boolean if a field has been set.

### GetOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetOldAmount() int64`

GetOldAmount returns the OldAmount field if non-nil, zero value otherwise.

### GetOldAmountOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetOldAmountOk() (*int64, bool)`

GetOldAmountOk returns a tuple with the OldAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetOldAmount(v int64)`

SetOldAmount sets OldAmount field to given value.

### HasOldAmount

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasOldAmount() bool`

HasOldAmount returns a boolean if a field has been set.

### GetOldCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetOldCurrency() string`

GetOldCurrency returns the OldCurrency field if non-nil, zero value otherwise.

### GetOldCurrencyOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetOldCurrencyOk() (*string, bool)`

GetOldCurrencyOk returns a tuple with the OldCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetOldCurrency(v string)`

SetOldCurrency sets OldCurrency field to given value.

### HasOldCurrency

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasOldCurrency() bool`

HasOldCurrency returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetTotalAffectedSubscriptions() int64`

GetTotalAffectedSubscriptions returns the TotalAffectedSubscriptions field if non-nil, zero value otherwise.

### GetTotalAffectedSubscriptionsOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetTotalAffectedSubscriptionsOk() (*int64, bool)`

GetTotalAffectedSubscriptionsOk returns a tuple with the TotalAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetTotalAffectedSubscriptions(v int64)`

SetTotalAffectedSubscriptions sets TotalAffectedSubscriptions field to given value.

### HasTotalAffectedSubscriptions

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasTotalAffectedSubscriptions() bool`

HasTotalAffectedSubscriptions returns a boolean if a field has been set.

### GetUnsupportedCurrencies

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetUnsupportedCurrencies() []string`

GetUnsupportedCurrencies returns the UnsupportedCurrencies field if non-nil, zero value otherwise.

### GetUnsupportedCurrenciesOk

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) GetUnsupportedCurrenciesOk() (*[]string, bool)`

GetUnsupportedCurrenciesOk returns a tuple with the UnsupportedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedCurrencies

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) SetUnsupportedCurrencies(v []string)`

SetUnsupportedCurrencies sets UnsupportedCurrencies field to given value.

### HasUnsupportedCurrencies

`func (o *UnibeeApiMerchantPlanActivePriceChangePreviewRes) HasUnsupportedCurrencies() bool`

HasUnsupportedCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


