# UnibeeApiMerchantPlanNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | Plan Ids Of Recurring Addon Type | [optional] 
**Amount** | **int64** | Plan CaptureAmount | 
**CancelAtTrialEnd** | Pointer to **int32** | whether cancel at subscription first trial end，0-false | 1-true, will pass to cancelAtPeriodEnd of subscription | [optional] 
**Currency** | **string** | Plan Currency | 
**Description** | Pointer to **string** | Description | [optional] 
**ExternalPlanId** | Pointer to **string** | ExternalPlanId | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas for crypto payment, merchant|user | [optional] 
**HomeUrl** | Pointer to **string** | HomeUrl,Start With: http | [optional] 
**ImageUrl** | Pointer to **string** | ImageUrl,Start With: http | [optional] 
**IntervalCount** | Pointer to **int32** | Number Of IntervalUnit，em: day|month|year|week | [optional] 
**IntervalUnit** | Pointer to **string** | Plan Interval Unit，em: day|month|year|week | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**MetricLimits** | Pointer to [**[]UnibeeApiBeanBulkMetricLimitPlanBindingParam**](UnibeeApiBeanBulkMetricLimitPlanBindingParam.md) | Plan&#39;s MetricLimit List | [optional] 
**OnetimeAddonIds** | Pointer to **[]int64** | Plan Ids Of Onetime Addon Type | [optional] 
**PlanName** | **string** | Plan Name | 
**ProductDescription** | Pointer to **string** | Default Copy Description | [optional] 
**ProductId** | Pointer to **int64** | Id of product which plan to linked | [optional] 
**ProductName** | Pointer to **string** | Default Copy PlanName | [optional] 
**TrialAmount** | Pointer to **int64** | price of trial period， not available for addon | [optional] 
**TrialDemand** | Pointer to **string** | demand of trial， not available for addon, example, paymentMethod, payment method will ask for subscription trial start | [optional] 
**TrialDurationTime** | Pointer to **int64** | duration of trial， not available for addon | [optional] 
**Type** | Pointer to **int32** | The type of plan, 1-main plan，2-addon plan, 3-onetime plan, default main plan | [optional] [default to 1]

## Methods

### NewUnibeeApiMerchantPlanNewReq

`func NewUnibeeApiMerchantPlanNewReq(amount int64, currency string, planName string, ) *UnibeeApiMerchantPlanNewReq`

NewUnibeeApiMerchantPlanNewReq instantiates a new UnibeeApiMerchantPlanNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanNewReqWithDefaults

`func NewUnibeeApiMerchantPlanNewReqWithDefaults() *UnibeeApiMerchantPlanNewReq`

NewUnibeeApiMerchantPlanNewReqWithDefaults instantiates a new UnibeeApiMerchantPlanNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiMerchantPlanNewReq) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeApiMerchantPlanNewReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantPlanNewReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantPlanNewReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanNewReq) GetCancelAtTrialEnd() int32`

GetCancelAtTrialEnd returns the CancelAtTrialEnd field if non-nil, zero value otherwise.

### GetCancelAtTrialEndOk

`func (o *UnibeeApiMerchantPlanNewReq) GetCancelAtTrialEndOk() (*int32, bool)`

GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanNewReq) SetCancelAtTrialEnd(v int32)`

SetCancelAtTrialEnd sets CancelAtTrialEnd field to given value.

### HasCancelAtTrialEnd

`func (o *UnibeeApiMerchantPlanNewReq) HasCancelAtTrialEnd() bool`

HasCancelAtTrialEnd returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPlanNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPlanNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPlanNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantPlanNewReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantPlanNewReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantPlanNewReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantPlanNewReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiMerchantPlanNewReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantPlanNewReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantPlanNewReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantPlanNewReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiMerchantPlanNewReq) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiMerchantPlanNewReq) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiMerchantPlanNewReq) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiMerchantPlanNewReq) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiMerchantPlanNewReq) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiMerchantPlanNewReq) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiMerchantPlanNewReq) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiMerchantPlanNewReq) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiMerchantPlanNewReq) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiMerchantPlanNewReq) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiMerchantPlanNewReq) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiMerchantPlanNewReq) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeApiMerchantPlanNewReq) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiMerchantPlanNewReq) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiMerchantPlanNewReq) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiMerchantPlanNewReq) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiMerchantPlanNewReq) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiMerchantPlanNewReq) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiMerchantPlanNewReq) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiMerchantPlanNewReq) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantPlanNewReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantPlanNewReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantPlanNewReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantPlanNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetricLimits

`func (o *UnibeeApiMerchantPlanNewReq) GetMetricLimits() []UnibeeApiBeanBulkMetricLimitPlanBindingParam`

GetMetricLimits returns the MetricLimits field if non-nil, zero value otherwise.

### GetMetricLimitsOk

`func (o *UnibeeApiMerchantPlanNewReq) GetMetricLimitsOk() (*[]UnibeeApiBeanBulkMetricLimitPlanBindingParam, bool)`

GetMetricLimitsOk returns a tuple with the MetricLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimits

`func (o *UnibeeApiMerchantPlanNewReq) SetMetricLimits(v []UnibeeApiBeanBulkMetricLimitPlanBindingParam)`

SetMetricLimits sets MetricLimits field to given value.

### HasMetricLimits

`func (o *UnibeeApiMerchantPlanNewReq) HasMetricLimits() bool`

HasMetricLimits returns a boolean if a field has been set.

### GetOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) GetOnetimeAddonIds() []int64`

GetOnetimeAddonIds returns the OnetimeAddonIds field if non-nil, zero value otherwise.

### GetOnetimeAddonIdsOk

`func (o *UnibeeApiMerchantPlanNewReq) GetOnetimeAddonIdsOk() (*[]int64, bool)`

GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) SetOnetimeAddonIds(v []int64)`

SetOnetimeAddonIds sets OnetimeAddonIds field to given value.

### HasOnetimeAddonIds

`func (o *UnibeeApiMerchantPlanNewReq) HasOnetimeAddonIds() bool`

HasOnetimeAddonIds returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiMerchantPlanNewReq) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiMerchantPlanNewReq) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiMerchantPlanNewReq) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetProductDescription

`func (o *UnibeeApiMerchantPlanNewReq) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *UnibeeApiMerchantPlanNewReq) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *UnibeeApiMerchantPlanNewReq) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *UnibeeApiMerchantPlanNewReq) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantPlanNewReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantPlanNewReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantPlanNewReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantPlanNewReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiMerchantPlanNewReq) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiMerchantPlanNewReq) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiMerchantPlanNewReq) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiMerchantPlanNewReq) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetTrialAmount

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialAmount() int64`

GetTrialAmount returns the TrialAmount field if non-nil, zero value otherwise.

### GetTrialAmountOk

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialAmountOk() (*int64, bool)`

GetTrialAmountOk returns a tuple with the TrialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmount

`func (o *UnibeeApiMerchantPlanNewReq) SetTrialAmount(v int64)`

SetTrialAmount sets TrialAmount field to given value.

### HasTrialAmount

`func (o *UnibeeApiMerchantPlanNewReq) HasTrialAmount() bool`

HasTrialAmount returns a boolean if a field has been set.

### GetTrialDemand

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialDemand() string`

GetTrialDemand returns the TrialDemand field if non-nil, zero value otherwise.

### GetTrialDemandOk

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialDemandOk() (*string, bool)`

GetTrialDemandOk returns a tuple with the TrialDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDemand

`func (o *UnibeeApiMerchantPlanNewReq) SetTrialDemand(v string)`

SetTrialDemand sets TrialDemand field to given value.

### HasTrialDemand

`func (o *UnibeeApiMerchantPlanNewReq) HasTrialDemand() bool`

HasTrialDemand returns a boolean if a field has been set.

### GetTrialDurationTime

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialDurationTime() int64`

GetTrialDurationTime returns the TrialDurationTime field if non-nil, zero value otherwise.

### GetTrialDurationTimeOk

`func (o *UnibeeApiMerchantPlanNewReq) GetTrialDurationTimeOk() (*int64, bool)`

GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDurationTime

`func (o *UnibeeApiMerchantPlanNewReq) SetTrialDurationTime(v int64)`

SetTrialDurationTime sets TrialDurationTime field to given value.

### HasTrialDurationTime

`func (o *UnibeeApiMerchantPlanNewReq) HasTrialDurationTime() bool`

HasTrialDurationTime returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantPlanNewReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantPlanNewReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantPlanNewReq) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantPlanNewReq) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


