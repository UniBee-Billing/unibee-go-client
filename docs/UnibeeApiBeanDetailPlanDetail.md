# UnibeeApiBeanDetailPlanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | AddonIds | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) | Addons | [optional] 
**MetricPlanLimits** | Pointer to [**[]UnibeeApiBeanMerchantMetricPlanLimit**](UnibeeApiBeanMerchantMetricPlanLimit.md) | MetricPlanLimits | [optional] 
**OnetimeAddonIds** | Pointer to **[]int64** | OneTimeAddonIds | [optional] 
**OnetimeAddons** | Pointer to [**[]UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) | OneTimeAddons | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailPlanDetail

`func NewUnibeeApiBeanDetailPlanDetail() *UnibeeApiBeanDetailPlanDetail`

NewUnibeeApiBeanDetailPlanDetail instantiates a new UnibeeApiBeanDetailPlanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailPlanDetailWithDefaults

`func NewUnibeeApiBeanDetailPlanDetailWithDefaults() *UnibeeApiBeanDetailPlanDetail`

NewUnibeeApiBeanDetailPlanDetailWithDefaults instantiates a new UnibeeApiBeanDetailPlanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanDetailPlanDetail) GetAddons() []UnibeeApiBeanPlanSimplify`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetAddonsOk() (*[]UnibeeApiBeanPlanSimplify, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailPlanDetail) SetAddons(v []UnibeeApiBeanPlanSimplify)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailPlanDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetMetricPlanLimits

`func (o *UnibeeApiBeanDetailPlanDetail) GetMetricPlanLimits() []UnibeeApiBeanMerchantMetricPlanLimit`

GetMetricPlanLimits returns the MetricPlanLimits field if non-nil, zero value otherwise.

### GetMetricPlanLimitsOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetMetricPlanLimitsOk() (*[]UnibeeApiBeanMerchantMetricPlanLimit, bool)`

GetMetricPlanLimitsOk returns a tuple with the MetricPlanLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPlanLimits

`func (o *UnibeeApiBeanDetailPlanDetail) SetMetricPlanLimits(v []UnibeeApiBeanMerchantMetricPlanLimit)`

SetMetricPlanLimits sets MetricPlanLimits field to given value.

### HasMetricPlanLimits

`func (o *UnibeeApiBeanDetailPlanDetail) HasMetricPlanLimits() bool`

HasMetricPlanLimits returns a boolean if a field has been set.

### GetOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonIds() []int64`

GetOnetimeAddonIds returns the OnetimeAddonIds field if non-nil, zero value otherwise.

### GetOnetimeAddonIdsOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonIdsOk() (*[]int64, bool)`

GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) SetOnetimeAddonIds(v []int64)`

SetOnetimeAddonIds sets OnetimeAddonIds field to given value.

### HasOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPlanDetail) HasOnetimeAddonIds() bool`

HasOnetimeAddonIds returns a boolean if a field has been set.

### GetOnetimeAddons

`func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddons() []UnibeeApiBeanPlanSimplify`

GetOnetimeAddons returns the OnetimeAddons field if non-nil, zero value otherwise.

### GetOnetimeAddonsOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetOnetimeAddonsOk() (*[]UnibeeApiBeanPlanSimplify, bool)`

GetOnetimeAddonsOk returns a tuple with the OnetimeAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddons

`func (o *UnibeeApiBeanDetailPlanDetail) SetOnetimeAddons(v []UnibeeApiBeanPlanSimplify)`

SetOnetimeAddons sets OnetimeAddons field to given value.

### HasOnetimeAddons

`func (o *UnibeeApiBeanDetailPlanDetail) HasOnetimeAddons() bool`

HasOnetimeAddons returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailPlanDetail) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailPlanDetail) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailPlanDetail) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailPlanDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


