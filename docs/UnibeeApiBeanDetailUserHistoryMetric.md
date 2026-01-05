# UnibeeApiBeanDetailUserHistoryMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | Pointer to **string** | Invoice ID | [optional] 
**LimitStats** | Pointer to [**[]UnibeeApiBeanDetailUserHistoryMetricLimitStat**](UnibeeApiBeanDetailUserHistoryMetricLimitStat.md) | Limit metric stats (limit_metered and limit_recurring only) | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserHistoryMetric

`func NewUnibeeApiBeanDetailUserHistoryMetric() *UnibeeApiBeanDetailUserHistoryMetric`

NewUnibeeApiBeanDetailUserHistoryMetric instantiates a new UnibeeApiBeanDetailUserHistoryMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserHistoryMetricWithDefaults

`func NewUnibeeApiBeanDetailUserHistoryMetricWithDefaults() *UnibeeApiBeanDetailUserHistoryMetric`

NewUnibeeApiBeanDetailUserHistoryMetricWithDefaults instantiates a new UnibeeApiBeanDetailUserHistoryMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiBeanDetailUserHistoryMetric) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailUserHistoryMetric) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailUserHistoryMetric) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailUserHistoryMetric) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLimitStats

`func (o *UnibeeApiBeanDetailUserHistoryMetric) GetLimitStats() []UnibeeApiBeanDetailUserHistoryMetricLimitStat`

GetLimitStats returns the LimitStats field if non-nil, zero value otherwise.

### GetLimitStatsOk

`func (o *UnibeeApiBeanDetailUserHistoryMetric) GetLimitStatsOk() (*[]UnibeeApiBeanDetailUserHistoryMetricLimitStat, bool)`

GetLimitStatsOk returns a tuple with the LimitStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitStats

`func (o *UnibeeApiBeanDetailUserHistoryMetric) SetLimitStats(v []UnibeeApiBeanDetailUserHistoryMetricLimitStat)`

SetLimitStats sets LimitStats field to given value.

### HasLimitStats

`func (o *UnibeeApiBeanDetailUserHistoryMetric) HasLimitStats() bool`

HasLimitStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


