# UnibeeApiBeanPlanSimplify

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
**ExtraMetricData** | Pointer to **string** |  | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**IntervalCount** | Pointer to **int32** | period unit count | [optional] 
**IntervalUnit** | Pointer to **string** | period unit,day|month|year|week | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**ProductDescription** | Pointer to **string** | product description | [optional] 
**ProductName** | Pointer to **string** | product name | [optional] 
**PublishStatus** | Pointer to **int32** | 1-UnPublish,2-Publish, Use For Display Plan At UserPortal | [optional] 
**Status** | Pointer to **int32** | status，1-editing，2-active，3-inactive，4-expired | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage 1000 &#x3D; 10% | [optional] 
**TrialAmount** | Pointer to **int64** | price of trial period | [optional] 
**TrialDemand** | Pointer to **string** |  | [optional] 
**TrialDurationTime** | Pointer to **int64** | duration of trial | [optional] 
**Type** | Pointer to **int32** | type，1-main plan，2-addon plan | [optional] 

## Methods

### NewUnibeeApiBeanPlanSimplify

`func NewUnibeeApiBeanPlanSimplify() *UnibeeApiBeanPlanSimplify`

NewUnibeeApiBeanPlanSimplify instantiates a new UnibeeApiBeanPlanSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPlanSimplifyWithDefaults

`func NewUnibeeApiBeanPlanSimplifyWithDefaults() *UnibeeApiBeanPlanSimplify`

NewUnibeeApiBeanPlanSimplifyWithDefaults instantiates a new UnibeeApiBeanPlanSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanPlanSimplify) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanPlanSimplify) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanPlanSimplify) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanPlanSimplify) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBindingAddonIds

`func (o *UnibeeApiBeanPlanSimplify) GetBindingAddonIds() string`

GetBindingAddonIds returns the BindingAddonIds field if non-nil, zero value otherwise.

### GetBindingAddonIdsOk

`func (o *UnibeeApiBeanPlanSimplify) GetBindingAddonIdsOk() (*string, bool)`

GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingAddonIds

`func (o *UnibeeApiBeanPlanSimplify) SetBindingAddonIds(v string)`

SetBindingAddonIds sets BindingAddonIds field to given value.

### HasBindingAddonIds

`func (o *UnibeeApiBeanPlanSimplify) HasBindingAddonIds() bool`

HasBindingAddonIds returns a boolean if a field has been set.

### GetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlanSimplify) GetBindingOnetimeAddonIds() string`

GetBindingOnetimeAddonIds returns the BindingOnetimeAddonIds field if non-nil, zero value otherwise.

### GetBindingOnetimeAddonIdsOk

`func (o *UnibeeApiBeanPlanSimplify) GetBindingOnetimeAddonIdsOk() (*string, bool)`

GetBindingOnetimeAddonIdsOk returns a tuple with the BindingOnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlanSimplify) SetBindingOnetimeAddonIds(v string)`

SetBindingOnetimeAddonIds sets BindingOnetimeAddonIds field to given value.

### HasBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPlanSimplify) HasBindingOnetimeAddonIds() bool`

HasBindingOnetimeAddonIds returns a boolean if a field has been set.

### GetCancelAtTrialEnd

`func (o *UnibeeApiBeanPlanSimplify) GetCancelAtTrialEnd() int32`

GetCancelAtTrialEnd returns the CancelAtTrialEnd field if non-nil, zero value otherwise.

### GetCancelAtTrialEndOk

`func (o *UnibeeApiBeanPlanSimplify) GetCancelAtTrialEndOk() (*int32, bool)`

GetCancelAtTrialEndOk returns a tuple with the CancelAtTrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtTrialEnd

`func (o *UnibeeApiBeanPlanSimplify) SetCancelAtTrialEnd(v int32)`

SetCancelAtTrialEnd sets CancelAtTrialEnd field to given value.

### HasCancelAtTrialEnd

`func (o *UnibeeApiBeanPlanSimplify) HasCancelAtTrialEnd() bool`

HasCancelAtTrialEnd returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanPlanSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanPlanSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanPlanSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanPlanSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPlanSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPlanSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPlanSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPlanSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanPlanSimplify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanPlanSimplify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanPlanSimplify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanPlanSimplify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtraMetricData

`func (o *UnibeeApiBeanPlanSimplify) GetExtraMetricData() string`

GetExtraMetricData returns the ExtraMetricData field if non-nil, zero value otherwise.

### GetExtraMetricDataOk

`func (o *UnibeeApiBeanPlanSimplify) GetExtraMetricDataOk() (*string, bool)`

GetExtraMetricDataOk returns a tuple with the ExtraMetricData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraMetricData

`func (o *UnibeeApiBeanPlanSimplify) SetExtraMetricData(v string)`

SetExtraMetricData sets ExtraMetricData field to given value.

### HasExtraMetricData

`func (o *UnibeeApiBeanPlanSimplify) HasExtraMetricData() bool`

HasExtraMetricData returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiBeanPlanSimplify) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiBeanPlanSimplify) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiBeanPlanSimplify) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiBeanPlanSimplify) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiBeanPlanSimplify) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiBeanPlanSimplify) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiBeanPlanSimplify) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiBeanPlanSimplify) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPlanSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPlanSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPlanSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPlanSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiBeanPlanSimplify) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiBeanPlanSimplify) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiBeanPlanSimplify) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiBeanPlanSimplify) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeApiBeanPlanSimplify) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiBeanPlanSimplify) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiBeanPlanSimplify) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiBeanPlanSimplify) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiBeanPlanSimplify) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiBeanPlanSimplify) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiBeanPlanSimplify) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiBeanPlanSimplify) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPlanSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPlanSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPlanSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPlanSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanPlanSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanPlanSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanPlanSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanPlanSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiBeanPlanSimplify) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiBeanPlanSimplify) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiBeanPlanSimplify) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiBeanPlanSimplify) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductDescription

`func (o *UnibeeApiBeanPlanSimplify) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *UnibeeApiBeanPlanSimplify) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *UnibeeApiBeanPlanSimplify) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *UnibeeApiBeanPlanSimplify) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanPlanSimplify) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanPlanSimplify) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanPlanSimplify) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanPlanSimplify) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetPublishStatus

`func (o *UnibeeApiBeanPlanSimplify) GetPublishStatus() int32`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *UnibeeApiBeanPlanSimplify) GetPublishStatusOk() (*int32, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *UnibeeApiBeanPlanSimplify) SetPublishStatus(v int32)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *UnibeeApiBeanPlanSimplify) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPlanSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPlanSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPlanSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPlanSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanPlanSimplify) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanPlanSimplify) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanPlanSimplify) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanPlanSimplify) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialAmount

`func (o *UnibeeApiBeanPlanSimplify) GetTrialAmount() int64`

GetTrialAmount returns the TrialAmount field if non-nil, zero value otherwise.

### GetTrialAmountOk

`func (o *UnibeeApiBeanPlanSimplify) GetTrialAmountOk() (*int64, bool)`

GetTrialAmountOk returns a tuple with the TrialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmount

`func (o *UnibeeApiBeanPlanSimplify) SetTrialAmount(v int64)`

SetTrialAmount sets TrialAmount field to given value.

### HasTrialAmount

`func (o *UnibeeApiBeanPlanSimplify) HasTrialAmount() bool`

HasTrialAmount returns a boolean if a field has been set.

### GetTrialDemand

`func (o *UnibeeApiBeanPlanSimplify) GetTrialDemand() string`

GetTrialDemand returns the TrialDemand field if non-nil, zero value otherwise.

### GetTrialDemandOk

`func (o *UnibeeApiBeanPlanSimplify) GetTrialDemandOk() (*string, bool)`

GetTrialDemandOk returns a tuple with the TrialDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDemand

`func (o *UnibeeApiBeanPlanSimplify) SetTrialDemand(v string)`

SetTrialDemand sets TrialDemand field to given value.

### HasTrialDemand

`func (o *UnibeeApiBeanPlanSimplify) HasTrialDemand() bool`

HasTrialDemand returns a boolean if a field has been set.

### GetTrialDurationTime

`func (o *UnibeeApiBeanPlanSimplify) GetTrialDurationTime() int64`

GetTrialDurationTime returns the TrialDurationTime field if non-nil, zero value otherwise.

### GetTrialDurationTimeOk

`func (o *UnibeeApiBeanPlanSimplify) GetTrialDurationTimeOk() (*int64, bool)`

GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDurationTime

`func (o *UnibeeApiBeanPlanSimplify) SetTrialDurationTime(v int64)`

SetTrialDurationTime sets TrialDurationTime field to given value.

### HasTrialDurationTime

`func (o *UnibeeApiBeanPlanSimplify) HasTrialDurationTime() bool`

HasTrialDurationTime returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanPlanSimplify) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanPlanSimplify) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanPlanSimplify) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanPlanSimplify) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


