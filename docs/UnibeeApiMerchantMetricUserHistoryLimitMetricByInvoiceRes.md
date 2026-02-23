# UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLatestPaidInvoice** | Pointer to **bool** | True when current-metric path was used (queried invoice is the subscription&#39;s latest paid invoice), false when snapshot or history path was used | [optional] 
**UserHistoryMetric** | Pointer to [**UnibeeApiBeanDetailUserHistoryMetric**](UnibeeApiBeanDetailUserHistoryMetric.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes

`func NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes() *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes`

NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes instantiates a new UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceResWithDefaults

`func NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceResWithDefaults() *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes`

NewUnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceResWithDefaults instantiates a new UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLatestPaidInvoice

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) GetIsLatestPaidInvoice() bool`

GetIsLatestPaidInvoice returns the IsLatestPaidInvoice field if non-nil, zero value otherwise.

### GetIsLatestPaidInvoiceOk

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) GetIsLatestPaidInvoiceOk() (*bool, bool)`

GetIsLatestPaidInvoiceOk returns a tuple with the IsLatestPaidInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestPaidInvoice

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) SetIsLatestPaidInvoice(v bool)`

SetIsLatestPaidInvoice sets IsLatestPaidInvoice field to given value.

### HasIsLatestPaidInvoice

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) HasIsLatestPaidInvoice() bool`

HasIsLatestPaidInvoice returns a boolean if a field has been set.

### GetUserHistoryMetric

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) GetUserHistoryMetric() UnibeeApiBeanDetailUserHistoryMetric`

GetUserHistoryMetric returns the UserHistoryMetric field if non-nil, zero value otherwise.

### GetUserHistoryMetricOk

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) GetUserHistoryMetricOk() (*UnibeeApiBeanDetailUserHistoryMetric, bool)`

GetUserHistoryMetricOk returns a tuple with the UserHistoryMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHistoryMetric

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) SetUserHistoryMetric(v UnibeeApiBeanDetailUserHistoryMetric)`

SetUserHistoryMetric sets UserHistoryMetric field to given value.

### HasUserHistoryMetric

`func (o *UnibeeApiMerchantMetricUserHistoryLimitMetricByInvoiceRes) HasUserHistoryMetric() bool`

HasUserHistoryMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


