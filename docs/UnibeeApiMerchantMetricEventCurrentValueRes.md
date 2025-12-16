# UnibeeApiMerchantMetricEventCurrentValueRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **int64** |  | [optional] 
**MetricLimit** | Pointer to [**UnibeeApiBeanDetailPlanMetricLimitDetail**](UnibeeApiBeanDetailPlanMetricLimitDetail.md) |  | [optional] 
**TotalLimit** | Pointer to **int64** | Total limit for the metric. Returns -1 if metric is not limit type. | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEventCurrentValueRes

`func NewUnibeeApiMerchantMetricEventCurrentValueRes() *UnibeeApiMerchantMetricEventCurrentValueRes`

NewUnibeeApiMerchantMetricEventCurrentValueRes instantiates a new UnibeeApiMerchantMetricEventCurrentValueRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEventCurrentValueResWithDefaults

`func NewUnibeeApiMerchantMetricEventCurrentValueResWithDefaults() *UnibeeApiMerchantMetricEventCurrentValueRes`

NewUnibeeApiMerchantMetricEventCurrentValueResWithDefaults instantiates a new UnibeeApiMerchantMetricEventCurrentValueRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetCurrentValue() int64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetCurrentValueOk() (*int64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) SetCurrentValue(v int64)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetMetricLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetMetricLimit() UnibeeApiBeanDetailPlanMetricLimitDetail`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetMetricLimitOk() (*UnibeeApiBeanDetailPlanMetricLimitDetail, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) SetMetricLimit(v UnibeeApiBeanDetailPlanMetricLimitDetail)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetTotalLimit() int64`

GetTotalLimit returns the TotalLimit field if non-nil, zero value otherwise.

### GetTotalLimitOk

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) GetTotalLimitOk() (*int64, bool)`

GetTotalLimitOk returns a tuple with the TotalLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) SetTotalLimit(v int64)`

SetTotalLimit sets TotalLimit field to given value.

### HasTotalLimit

`func (o *UnibeeApiMerchantMetricEventCurrentValueRes) HasTotalLimit() bool`

HasTotalLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


