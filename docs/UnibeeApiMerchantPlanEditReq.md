# UnibeeApiMerchantPlanEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | Plan Ids Of Addon Type | [optional] 
**Amount** | **int64** | Plan Amount | 
**Currency** | **string** | Plan Currency | 
**Description** | Pointer to **string** | Description | [optional] 
**HomeUrl** | Pointer to **string** | HomeUrl,Start With: http | [optional] 
**ImageUrl** | Pointer to **string** | ImageUrl,Start With: http | [optional] 
**IntervalCount** | Pointer to **int32** | Default 1，Number Of IntervalUnit | [optional] [default to 1]
**IntervalUnit** | **string** | Plan Interval Unit，em: day|month|year|week | 
**MetricLimits** | Pointer to [**[]UnibeeInternalLogicGatewayRoBulkMetricLimitPlanBindingParam**](UnibeeInternalLogicGatewayRoBulkMetricLimitPlanBindingParam.md) | Plan&#39;s MetricLimit List | [optional] 
**PlanId** | **int64** | PlanId | 
**PlanName** | **string** | Plan Name | 
**ProductDescription** | Pointer to **string** | Default Copy Description | [optional] 
**ProductName** | Pointer to **string** | Default Copy PlanName | [optional] 

## Methods

### NewUnibeeApiMerchantPlanEditReq

`func NewUnibeeApiMerchantPlanEditReq(amount int64, currency string, intervalUnit string, planId int64, planName string, ) *UnibeeApiMerchantPlanEditReq`

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

`func (o *UnibeeApiMerchantPlanEditReq) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiMerchantPlanEditReq) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiMerchantPlanEditReq) SetAmount(v int64)`

SetAmount sets Amount field to given value.


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


### GetMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimits() []UnibeeInternalLogicGatewayRoBulkMetricLimitPlanBindingParam`

GetMetricLimits returns the MetricLimits field if non-nil, zero value otherwise.

### GetMetricLimitsOk

`func (o *UnibeeApiMerchantPlanEditReq) GetMetricLimitsOk() (*[]UnibeeInternalLogicGatewayRoBulkMetricLimitPlanBindingParam, bool)`

GetMetricLimitsOk returns a tuple with the MetricLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) SetMetricLimits(v []UnibeeInternalLogicGatewayRoBulkMetricLimitPlanBindingParam)`

SetMetricLimits sets MetricLimits field to given value.

### HasMetricLimits

`func (o *UnibeeApiMerchantPlanEditReq) HasMetricLimits() bool`

HasMetricLimits returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


