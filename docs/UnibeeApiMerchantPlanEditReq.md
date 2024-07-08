# UnibeeApiMerchantPlanEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | Plan Ids Of Recurring Addon Type | [optional] 
**Amount** | Pointer to **int32** | CaptureAmount of plan, not editable when plan is active | [optional] 
**CancelAtTrialEnd** | Pointer to **int32** | whether cancel at subscripiton first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription | [optional] 
**Currency** | Pointer to **string** | Currency of plan, not editable when plan is active | [optional] 
**Description** | Pointer to **string** | Description of plan | [optional] 
**ExternalPlanId** | Pointer to **string** | ExternalPlanId | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas for crypto payment, merchant|user | [optional] 
**HomeUrl** | Pointer to **string** | HomeUrl,Start With: http | [optional] 
**ImageUrl** | Pointer to **string** | ImageUrl,Start With: http | [optional] 
**IntervalCount** | Pointer to **int32** | Number,intervalUnit of plan, not editable when plan is active | [optional] 
**IntervalUnit** | Pointer to **string** | Interval unit of plan，em: day|month|year|week, not editable when plan is active | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**MetricLimits** | Pointer to [**[]UnibeeApiBeanBulkMetricLimitPlanBindingParam**](UnibeeApiBeanBulkMetricLimitPlanBindingParam.md) | Plan&#39;s MetricLimit List | [optional] 
**OnetimeAddonIds** | Pointer to **[]int64** | Plan Ids Of Onetime Addon Type | [optional] 
**PlanId** | **int64** | Id of plan | 
**PlanName** | Pointer to **string** | Name of plan | [optional] 
**ProductDescription** | Pointer to **string** | ProductDescription of plan, Default copy description | [optional] 
**ProductName** | Pointer to **string** | ProductName of plan, Default copy planName | [optional] 
**TrialAmount** | Pointer to **int32** | price of trial period， not available for addon | [optional] 
**TrialDemand** | Pointer to **string** | demand of trial, not available for addon, example, paymentMethod, payment method will ask for subscription trial start | [optional] 
**TrialDurationTime** | Pointer to **int32** | duration of trial， not available for addon | [optional] 

## Methods

### NewUnibeeApiMerchantPlanEditReq

`func NewUnibeeApiMerchantPlanEditReq(planId int64, ) *UnibeeApiMerchantPlanEditReq`

NewUnibeeApiMerchantPlanEditReq instantiates a new UnibeeApiMerchantPlanEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanEditReqWithDefaults

`func NewUnibeeApiMerchantPlanEditReqWithDefaults() *UnibeeApiMerchantPlanEditReq`

NewUnibeeApiMerchantPlanEditReqWithDefaults instantiates a new UnibeeApiMerchantPlanEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiMerchantPlanEditReq) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeApiMerchantPlanEditReq) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantPlanEditReq) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantPlanEditReq) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiMerchantPlanEditReq) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanEditReq) GetCancelAtTrialEnd() int32`

GetCancelAtTrialEnd returns the CancelAtTrialEnd field if non-nil, zero value otherwise.

### GetCancelAtTrialEndOk

`func (o *UnibeeApiMerchantPlanEditReq) GetCancelAtTrialEndOk() (*int32, bool)`

GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanEditReq) SetCancelAtTrialEnd(v int32)`

SetCancelAtTrialEnd sets CancelAtTrialEnd field to given value.

### HasCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanEditReq) HasCancelAtTrialEnd() bool`

HasCancelAtTrialEnd returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPlanEditReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPlanEditReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPlanEditReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPlanEditReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantPlanEditReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantPlanEditReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantPlanEditReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantPlanEditReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiMerchantPlanEditReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantPlanEditReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantPlanEditReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantPlanEditReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiMerchantPlanEditReq) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiMerchantPlanEditReq) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiMerchantPlanEditReq) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiMerchantPlanEditReq) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiMerchantPlanEditReq) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiMerchantPlanEditReq) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiMerchantPlanEditReq) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiMerchantPlanEditReq) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiMerchantPlanEditReq) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiMerchantPlanEditReq) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiMerchantPlanEditReq) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiMerchantPlanEditReq) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiMerchantPlanEditReq) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiMerchantPlanEditReq) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiMerchantPlanEditReq) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiMerchantPlanEditReq) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiMerchantPlanEditReq) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantPlanEditReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantPlanEditReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantPlanEditReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantPlanEditReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimits() []UnibeeApiBeanBulkMetricLimitPlanBindingParam`

GetMetricLimits returns the MetricLimits field if non-nil, zero value otherwise.

### GetMetricLimitsOk

`func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimitsOk() (*[]UnibeeApiBeanBulkMetricLimitPlanBindingParam, bool)`

GetMetricLimitsOk returns a tuple with the MetricLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) SetMetricLimits(v []UnibeeApiBeanBulkMetricLimitPlanBindingParam)`

SetMetricLimits sets MetricLimits field to given value.

### HasMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) HasMetricLimits() bool`

HasMetricLimits returns a boolean if a field has been set.

### GetOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) GetOnetimeAddonIds() []int64`

GetOnetimeAddonIds returns the OnetimeAddonIds field if non-nil, zero value otherwise.

### GetOnetimeAddonIdsOk

`func (o *UnibeeApiMerchantPlanEditReq) GetOnetimeAddonIdsOk() (*[]int64, bool)`

GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) SetOnetimeAddonIds(v []int64)`

SetOnetimeAddonIds sets OnetimeAddonIds field to given value.

### HasOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanEditReq) HasOnetimeAddonIds() bool`

HasOnetimeAddonIds returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPlanEditReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPlanEditReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPlanEditReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetPlanName

`func (o *UnibeeApiMerchantPlanEditReq) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiMerchantPlanEditReq) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiMerchantPlanEditReq) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiMerchantPlanEditReq) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductDescription

`func (o *UnibeeApiMerchantPlanEditReq) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *UnibeeApiMerchantPlanEditReq) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *UnibeeApiMerchantPlanEditReq) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *UnibeeApiMerchantPlanEditReq) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiMerchantPlanEditReq) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiMerchantPlanEditReq) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiMerchantPlanEditReq) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiMerchantPlanEditReq) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTrialAmount

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialAmount() int32`

GetTrialAmount returns the TrialAmount field if non-nil, zero value otherwise.

### GetTrialAmountOk

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialAmountOk() (*int32, bool)`

GetTrialAmountOk returns a tuple with the TrialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmount

`func (o *UnibeeApiMerchantPlanEditReq) SetTrialAmount(v int32)`

SetTrialAmount sets TrialAmount field to given value.

### HasTrialAmount

`func (o *UnibeeApiMerchantPlanEditReq) HasTrialAmount() bool`

HasTrialAmount returns a boolean if a field has been set.

### GetTrialDemand

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialDemand() string`

GetTrialDemand returns the TrialDemand field if non-nil, zero value otherwise.

### GetTrialDemandOk

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialDemandOk() (*string, bool)`

GetTrialDemandOk returns a tuple with the TrialDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDemand

`func (o *UnibeeApiMerchantPlanEditReq) SetTrialDemand(v string)`

SetTrialDemand sets TrialDemand field to given value.

### HasTrialDemand

`func (o *UnibeeApiMerchantPlanEditReq) HasTrialDemand() bool`

HasTrialDemand returns a boolean if a field has been set.

### GetTrialDurationTime

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialDurationTime() int32`

GetTrialDurationTime returns the TrialDurationTime field if non-nil, zero value otherwise.

### GetTrialDurationTimeOk

`func (o *UnibeeApiMerchantPlanEditReq) GetTrialDurationTimeOk() (*int32, bool)`

GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDurationTime

`func (o *UnibeeApiMerchantPlanEditReq) SetTrialDurationTime(v int32)`

SetTrialDurationTime sets TrialDurationTime field to given value.

### HasTrialDurationTime

`func (o *UnibeeApiMerchantPlanEditReq) HasTrialDurationTime() bool`

HasTrialDurationTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


