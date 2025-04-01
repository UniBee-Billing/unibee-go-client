# UnibeeApiBeanPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | amount, cent, without tax | [optional] 
**BindingAddonIds** | Pointer to **string** | binded recurring addon planIds，split with , | [optional] 
**BindingOnetimeAddonIds** | Pointer to **string** | binded onetime addon planIds，split with , | [optional] 
**CancelAtTrialEnd** | Pointer to **int32** | whether cancel at subscripiton first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**ExternalPlanId** | Pointer to **string** | external_user_id | [optional] 
**ExtraMetricData** | Pointer to **string** |  | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**InternalName** | Pointer to **string** | PlanInternalName | [optional] 
**IntervalCount** | Pointer to **int32** | period unit count | [optional] 
**IntervalUnit** | Pointer to **string** | period unit,day|month|year|week | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MetricLimits** | Pointer to [**[]UnibeeApiBeanPlanMetricLimitParam**](UnibeeApiBeanPlanMetricLimitParam.md) | Plan&#39;s MetricLimit List | [optional] 
**MetricMeteredCharge** | Pointer to [**[]UnibeeApiBeanPlanMetricMeteredChargeParam**](UnibeeApiBeanPlanMetricMeteredChargeParam.md) | Plan&#39;s MetricMeteredCharge | [optional] 
**MetricRecurringCharge** | Pointer to [**[]UnibeeApiBeanPlanMetricMeteredChargeParam**](UnibeeApiBeanPlanMetricMeteredChargeParam.md) | Plan&#39;s MetricRecurringCharge | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**ProductId** | Pointer to **int64** | product id | [optional] 
**PublishStatus** | Pointer to **int32** | 1-UnPublish,2-Publish, Use For Display Plan At UserPortal | [optional] 
**Status** | Pointer to **int32** | status，1-editing，2-active，3-inactive，4-soft archive, 5-hard archive | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage 1000 &#x3D; 10% | [optional] 
**TrialAmount** | Pointer to **int64** | price of trial period | [optional] 
**TrialDemand** | Pointer to **string** |  | [optional] 
**TrialDurationTime** | Pointer to **int64** | duration of trial | [optional] 
**Type** | Pointer to **int32** | type，1-main plan，2-addon plan | [optional] 

## Methods

### NewUnibeeApiBeanPlan

`func NewUnibeeApiBeanPlan() *UnibeeApiBeanPlan`

NewUnibeeApiBeanPlan instantiates a new UnibeeApiBeanPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanWithDefaults

`func NewUnibeeApiBeanPlanWithDefaults() *UnibeeApiBeanPlan`

NewUnibeeApiBeanPlanWithDefaults instantiates a new UnibeeApiBeanPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanPlan) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanPlan) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanPlan) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanPlan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBindingAddonIds

`func (o *UnibeeApiBeanPlan) GetBindingAddonIds() string`

GetBindingAddonIds returns the BindingAddonIds field if non-nil, zero value otherwise.

### GetBindingAddonIdsOk

`func (o *UnibeeApiBeanPlan) GetBindingAddonIdsOk() (*string, bool)`

GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingAddonIds

`func (o *UnibeeApiBeanPlan) SetBindingAddonIds(v string)`

SetBindingAddonIds sets BindingAddonIds field to given value.

### HasBindingAddonIds

`func (o *UnibeeApiBeanPlan) HasBindingAddonIds() bool`

HasBindingAddonIds returns a boolean if a field has been set.

### GetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlan) GetBindingOnetimeAddonIds() string`

GetBindingOnetimeAddonIds returns the BindingOnetimeAddonIds field if non-nil, zero value otherwise.

### GetBindingOnetimeAddonIdsOk

`func (o *UnibeeApiBeanPlan) GetBindingOnetimeAddonIdsOk() (*string, bool)`

GetBindingOnetimeAddonIdsOk returns a tuple with the BindingOnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlan) SetBindingOnetimeAddonIds(v string)`

SetBindingOnetimeAddonIds sets BindingOnetimeAddonIds field to given value.

### HasBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlan) HasBindingOnetimeAddonIds() bool`

HasBindingOnetimeAddonIds returns a boolean if a field has been set.

### GetCancelAtTrialEnd

`func (o *UnibeeApiBeanPlan) GetCancelAtTrialEnd() int32`

GetCancelAtTrialEnd returns the CancelAtTrialEnd field if non-nil, zero value otherwise.

### GetCancelAtTrialEndOk

`func (o *UnibeeApiBeanPlan) GetCancelAtTrialEndOk() (*int32, bool)`

GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtTrialEnd

`func (o *UnibeeApiBeanPlan) SetCancelAtTrialEnd(v int32)`

SetCancelAtTrialEnd sets CancelAtTrialEnd field to given value.

### HasCancelAtTrialEnd

`func (o *UnibeeApiBeanPlan) HasCancelAtTrialEnd() bool`

HasCancelAtTrialEnd returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanPlan) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanPlan) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanPlan) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanPlan) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiBeanPlan) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiBeanPlan) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiBeanPlan) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiBeanPlan) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetExtraMetricData

`func (o *UnibeeApiBeanPlan) GetExtraMetricData() string`

GetExtraMetricData returns the ExtraMetricData field if non-nil, zero value otherwise.

### GetExtraMetricDataOk

`func (o *UnibeeApiBeanPlan) GetExtraMetricDataOk() (*string, bool)`

GetExtraMetricDataOk returns a tuple with the ExtraMetricData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMetricData

`func (o *UnibeeApiBeanPlan) SetExtraMetricData(v string)`

SetExtraMetricData sets ExtraMetricData field to given value.

### HasExtraMetricData

`func (o *UnibeeApiBeanPlan) HasExtraMetricData() bool`

HasExtraMetricData returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiBeanPlan) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiBeanPlan) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiBeanPlan) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiBeanPlan) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiBeanPlan) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiBeanPlan) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiBeanPlan) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiBeanPlan) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiBeanPlan) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiBeanPlan) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiBeanPlan) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiBeanPlan) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetInternalName

`func (o *UnibeeApiBeanPlan) GetInternalName() string`

GetInternalName returns the InternalName field if non-nil, zero value otherwise.

### GetInternalNameOk

`func (o *UnibeeApiBeanPlan) GetInternalNameOk() (*string, bool)`

GetInternalNameOk returns a tuple with the InternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalName

`func (o *UnibeeApiBeanPlan) SetInternalName(v string)`

SetInternalName sets InternalName field to given value.

### HasInternalName

`func (o *UnibeeApiBeanPlan) HasInternalName() bool`

HasInternalName returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeApiBeanPlan) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiBeanPlan) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiBeanPlan) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiBeanPlan) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiBeanPlan) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiBeanPlan) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiBeanPlan) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiBeanPlan) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPlan) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPlan) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPlan) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPlan) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanPlan) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanPlan) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanPlan) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanPlan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetricLimits

`func (o *UnibeeApiBeanPlan) GetMetricLimits() []UnibeeApiBeanPlanMetricLimitParam`

GetMetricLimits returns the MetricLimits field if non-nil, zero value otherwise.

### GetMetricLimitsOk

`func (o *UnibeeApiBeanPlan) GetMetricLimitsOk() (*[]UnibeeApiBeanPlanMetricLimitParam, bool)`

GetMetricLimitsOk returns a tuple with the MetricLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimits

`func (o *UnibeeApiBeanPlan) SetMetricLimits(v []UnibeeApiBeanPlanMetricLimitParam)`

SetMetricLimits sets MetricLimits field to given value.

### HasMetricLimits

`func (o *UnibeeApiBeanPlan) HasMetricLimits() bool`

HasMetricLimits returns a boolean if a field has been set.

### GetMetricMeteredCharge

`func (o *UnibeeApiBeanPlan) GetMetricMeteredCharge() []UnibeeApiBeanPlanMetricMeteredChargeParam`

GetMetricMeteredCharge returns the MetricMeteredCharge field if non-nil, zero value otherwise.

### GetMetricMeteredChargeOk

`func (o *UnibeeApiBeanPlan) GetMetricMeteredChargeOk() (*[]UnibeeApiBeanPlanMetricMeteredChargeParam, bool)`

GetMetricMeteredChargeOk returns a tuple with the MetricMeteredCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricMeteredCharge

`func (o *UnibeeApiBeanPlan) SetMetricMeteredCharge(v []UnibeeApiBeanPlanMetricMeteredChargeParam)`

SetMetricMeteredCharge sets MetricMeteredCharge field to given value.

### HasMetricMeteredCharge

`func (o *UnibeeApiBeanPlan) HasMetricMeteredCharge() bool`

HasMetricMeteredCharge returns a boolean if a field has been set.

### GetMetricRecurringCharge

`func (o *UnibeeApiBeanPlan) GetMetricRecurringCharge() []UnibeeApiBeanPlanMetricMeteredChargeParam`

GetMetricRecurringCharge returns the MetricRecurringCharge field if non-nil, zero value otherwise.

### GetMetricRecurringChargeOk

`func (o *UnibeeApiBeanPlan) GetMetricRecurringChargeOk() (*[]UnibeeApiBeanPlanMetricMeteredChargeParam, bool)`

GetMetricRecurringChargeOk returns a tuple with the MetricRecurringCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRecurringCharge

`func (o *UnibeeApiBeanPlan) SetMetricRecurringCharge(v []UnibeeApiBeanPlanMetricMeteredChargeParam)`

SetMetricRecurringCharge sets MetricRecurringCharge field to given value.

### HasMetricRecurringCharge

`func (o *UnibeeApiBeanPlan) HasMetricRecurringCharge() bool`

HasMetricRecurringCharge returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiBeanPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiBeanPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiBeanPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiBeanPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiBeanPlan) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiBeanPlan) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiBeanPlan) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiBeanPlan) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetPublishStatus

`func (o *UnibeeApiBeanPlan) GetPublishStatus() int32`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *UnibeeApiBeanPlan) GetPublishStatusOk() (*int32, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *UnibeeApiBeanPlan) SetPublishStatus(v int32)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *UnibeeApiBeanPlan) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPlan) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPlan) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPlan) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanPlan) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanPlan) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanPlan) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanPlan) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialAmount

`func (o *UnibeeApiBeanPlan) GetTrialAmount() int64`

GetTrialAmount returns the TrialAmount field if non-nil, zero value otherwise.

### GetTrialAmountOk

`func (o *UnibeeApiBeanPlan) GetTrialAmountOk() (*int64, bool)`

GetTrialAmountOk returns a tuple with the TrialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmount

`func (o *UnibeeApiBeanPlan) SetTrialAmount(v int64)`

SetTrialAmount sets TrialAmount field to given value.

### HasTrialAmount

`func (o *UnibeeApiBeanPlan) HasTrialAmount() bool`

HasTrialAmount returns a boolean if a field has been set.

### GetTrialDemand

`func (o *UnibeeApiBeanPlan) GetTrialDemand() string`

GetTrialDemand returns the TrialDemand field if non-nil, zero value otherwise.

### GetTrialDemandOk

`func (o *UnibeeApiBeanPlan) GetTrialDemandOk() (*string, bool)`

GetTrialDemandOk returns a tuple with the TrialDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDemand

`func (o *UnibeeApiBeanPlan) SetTrialDemand(v string)`

SetTrialDemand sets TrialDemand field to given value.

### HasTrialDemand

`func (o *UnibeeApiBeanPlan) HasTrialDemand() bool`

HasTrialDemand returns a boolean if a field has been set.

### GetTrialDurationTime

`func (o *UnibeeApiBeanPlan) GetTrialDurationTime() int64`

GetTrialDurationTime returns the TrialDurationTime field if non-nil, zero value otherwise.

### GetTrialDurationTimeOk

`func (o *UnibeeApiBeanPlan) GetTrialDurationTimeOk() (*int64, bool)`

GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDurationTime

`func (o *UnibeeApiBeanPlan) SetTrialDurationTime(v int64)`

SetTrialDurationTime sets TrialDurationTime field to given value.

### HasTrialDurationTime

`func (o *UnibeeApiBeanPlan) HasTrialDurationTime() bool`

HasTrialDurationTime returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanPlan) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanPlan) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanPlan) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanPlan) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


