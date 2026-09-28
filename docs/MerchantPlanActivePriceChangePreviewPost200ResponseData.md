# MerchantPlanActivePriceChangePreviewPost200ResponseData

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

### NewMerchantPlanActivePriceChangePreviewPost200ResponseData

`func NewMerchantPlanActivePriceChangePreviewPost200ResponseData() *MerchantPlanActivePriceChangePreviewPost200ResponseData`

NewMerchantPlanActivePriceChangePreviewPost200ResponseData instantiates a new MerchantPlanActivePriceChangePreviewPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPlanActivePriceChangePreviewPost200ResponseDataWithDefaults

`func NewMerchantPlanActivePriceChangePreviewPost200ResponseDataWithDefaults() *MerchantPlanActivePriceChangePreviewPost200ResponseData`

NewMerchantPlanActivePriceChangePreviewPost200ResponseDataWithDefaults instantiates a new MerchantPlanActivePriceChangePreviewPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetActiveAffectedSubscriptions() int64`

GetActiveAffectedSubscriptions returns the ActiveAffectedSubscriptions field if non-nil, zero value otherwise.

### GetActiveAffectedSubscriptionsOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetActiveAffectedSubscriptionsOk() (*int64, bool)`

GetActiveAffectedSubscriptionsOk returns a tuple with the ActiveAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetActiveAffectedSubscriptions(v int64)`

SetActiveAffectedSubscriptions sets ActiveAffectedSubscriptions field to given value.

### HasActiveAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasActiveAffectedSubscriptions() bool`

HasActiveAffectedSubscriptions returns a boolean if a field has been set.

### GetNewAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetNewAmount() int64`

GetNewAmount returns the NewAmount field if non-nil, zero value otherwise.

### GetNewAmountOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetNewAmountOk() (*int64, bool)`

GetNewAmountOk returns a tuple with the NewAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetNewAmount(v int64)`

SetNewAmount sets NewAmount field to given value.

### HasNewAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasNewAmount() bool`

HasNewAmount returns a boolean if a field has been set.

### GetNewCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetNewCurrency() string`

GetNewCurrency returns the NewCurrency field if non-nil, zero value otherwise.

### GetNewCurrencyOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetNewCurrencyOk() (*string, bool)`

GetNewCurrencyOk returns a tuple with the NewCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetNewCurrency(v string)`

SetNewCurrency sets NewCurrency field to given value.

### HasNewCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasNewCurrency() bool`

HasNewCurrency returns a boolean if a field has been set.

### GetOldAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetOldAmount() int64`

GetOldAmount returns the OldAmount field if non-nil, zero value otherwise.

### GetOldAmountOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetOldAmountOk() (*int64, bool)`

GetOldAmountOk returns a tuple with the OldAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetOldAmount(v int64)`

SetOldAmount sets OldAmount field to given value.

### HasOldAmount

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasOldAmount() bool`

HasOldAmount returns a boolean if a field has been set.

### GetOldCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetOldCurrency() string`

GetOldCurrency returns the OldCurrency field if non-nil, zero value otherwise.

### GetOldCurrencyOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetOldCurrencyOk() (*string, bool)`

GetOldCurrencyOk returns a tuple with the OldCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetOldCurrency(v string)`

SetOldCurrency sets OldCurrency field to given value.

### HasOldCurrency

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasOldCurrency() bool`

HasOldCurrency returns a boolean if a field has been set.

### GetPlanId

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetTotalAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetTotalAffectedSubscriptions() int64`

GetTotalAffectedSubscriptions returns the TotalAffectedSubscriptions field if non-nil, zero value otherwise.

### GetTotalAffectedSubscriptionsOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetTotalAffectedSubscriptionsOk() (*int64, bool)`

GetTotalAffectedSubscriptionsOk returns a tuple with the TotalAffectedSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetTotalAffectedSubscriptions(v int64)`

SetTotalAffectedSubscriptions sets TotalAffectedSubscriptions field to given value.

### HasTotalAffectedSubscriptions

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasTotalAffectedSubscriptions() bool`

HasTotalAffectedSubscriptions returns a boolean if a field has been set.

### GetUnsupportedCurrencies

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetUnsupportedCurrencies() []string`

GetUnsupportedCurrencies returns the UnsupportedCurrencies field if non-nil, zero value otherwise.

### GetUnsupportedCurrenciesOk

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) GetUnsupportedCurrenciesOk() (*[]string, bool)`

GetUnsupportedCurrenciesOk returns a tuple with the UnsupportedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportedCurrencies

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) SetUnsupportedCurrencies(v []string)`

SetUnsupportedCurrencies sets UnsupportedCurrencies field to given value.

### HasUnsupportedCurrencies

`func (o *MerchantPlanActivePriceChangePreviewPost200ResponseData) HasUnsupportedCurrencies() bool`

HasUnsupportedCurrencies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


