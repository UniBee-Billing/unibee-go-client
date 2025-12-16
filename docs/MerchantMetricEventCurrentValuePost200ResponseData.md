# MerchantMetricEventCurrentValuePost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **int64** |  | [optional] 
**MetricLimit** | Pointer to [**UnibeeApiBeanDetailPlanMetricLimitDetail**](UnibeeApiBeanDetailPlanMetricLimitDetail.md) |  | [optional] 
**TotalLimit** | Pointer to **int64** | Total limit for the metric. Returns -1 if metric is not limit type. | [optional] 

## Methods

### NewMerchantMetricEventCurrentValuePost200ResponseData

`func NewMerchantMetricEventCurrentValuePost200ResponseData() *MerchantMetricEventCurrentValuePost200ResponseData`

NewMerchantMetricEventCurrentValuePost200ResponseData instantiates a new MerchantMetricEventCurrentValuePost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantMetricEventCurrentValuePost200ResponseDataWithDefaults

`func NewMerchantMetricEventCurrentValuePost200ResponseDataWithDefaults() *MerchantMetricEventCurrentValuePost200ResponseData`

NewMerchantMetricEventCurrentValuePost200ResponseDataWithDefaults instantiates a new MerchantMetricEventCurrentValuePost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetCurrentValue() int64`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetCurrentValueOk() (*int64, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) SetCurrentValue(v int64)`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### GetMetricLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetMetricLimit() UnibeeApiBeanDetailPlanMetricLimitDetail`

GetMetricLimit returns the MetricLimit field if non-nil, zero value otherwise.

### GetMetricLimitOk

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetMetricLimitOk() (*UnibeeApiBeanDetailPlanMetricLimitDetail, bool)`

GetMetricLimitOk returns a tuple with the MetricLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) SetMetricLimit(v UnibeeApiBeanDetailPlanMetricLimitDetail)`

SetMetricLimit sets MetricLimit field to given value.

### HasMetricLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) HasMetricLimit() bool`

HasMetricLimit returns a boolean if a field has been set.

### GetTotalLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetTotalLimit() int64`

GetTotalLimit returns the TotalLimit field if non-nil, zero value otherwise.

### GetTotalLimitOk

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) GetTotalLimitOk() (*int64, bool)`

GetTotalLimitOk returns a tuple with the TotalLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) SetTotalLimit(v int64)`

SetTotalLimit sets TotalLimit field to given value.

### HasTotalLimit

`func (o *MerchantMetricEventCurrentValuePost200ResponseData) HasTotalLimit() bool`

HasTotalLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


