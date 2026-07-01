# UnibeeApiMerchantMetricMetricLimitAdjustBatchRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DuplicateCount** | Pointer to **int32** |  | [optional] 
**FailedCount** | Pointer to **int32** |  | [optional] 
**ProcessedCount** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]UnibeeApiMerchantMetricMetricLimitAdjustBatchItemResult**](UnibeeApiMerchantMetricMetricLimitAdjustBatchItemResult.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**SuccessCount** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchRes

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchRes() *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchRes instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchResWithDefaults

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchResWithDefaults() *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchResWithDefaults instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplicateCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetDuplicateCount() int32`

GetDuplicateCount returns the DuplicateCount field if non-nil, zero value otherwise.

### GetDuplicateCountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetDuplicateCountOk() (*int32, bool)`

GetDuplicateCountOk returns a tuple with the DuplicateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetDuplicateCount(v int32)`

SetDuplicateCount sets DuplicateCount field to given value.

### HasDuplicateCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasDuplicateCount() bool`

HasDuplicateCount returns a boolean if a field has been set.

### GetFailedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetProcessedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetProcessedCount() int32`

GetProcessedCount returns the ProcessedCount field if non-nil, zero value otherwise.

### GetProcessedCountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetProcessedCountOk() (*int32, bool)`

GetProcessedCountOk returns a tuple with the ProcessedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetProcessedCount(v int32)`

SetProcessedCount sets ProcessedCount field to given value.

### HasProcessedCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasProcessedCount() bool`

HasProcessedCount returns a boolean if a field has been set.

### GetRequestId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResults

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetResults() []UnibeeApiMerchantMetricMetricLimitAdjustBatchItemResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetResultsOk() (*[]UnibeeApiMerchantMetricMetricLimitAdjustBatchItemResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetResults(v []UnibeeApiMerchantMetricMetricLimitAdjustBatchItemResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccess

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSuccessCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


