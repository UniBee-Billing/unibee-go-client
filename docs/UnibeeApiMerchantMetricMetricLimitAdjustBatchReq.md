# UnibeeApiMerchantMetricMetricLimitAdjustBatchReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | Pointer to **string** | Optional batch-level idempotency key, max 256 chars. Used to derive item idempotency keys when item idempotencyKey is empty. | [optional] 
**Items** | [**[]UnibeeApiMerchantMetricMetricLimitAdjustBatchItem**](UnibeeApiMerchantMetricMetricLimitAdjustBatchItem.md) | Adjustment items, max 500 | 
**RequestId** | **string** | Batch request ID for logs and reconciliation | 

## Methods

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReq

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReq(items []UnibeeApiMerchantMetricMetricLimitAdjustBatchItem, requestId string, ) *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReq instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReqWithDefaults

`func NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReqWithDefaults() *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq`

NewUnibeeApiMerchantMetricMetricLimitAdjustBatchReqWithDefaults instantiates a new UnibeeApiMerchantMetricMetricLimitAdjustBatchReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### HasIdempotencyKey

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) HasIdempotencyKey() bool`

HasIdempotencyKey returns a boolean if a field has been set.

### GetItems

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetItems() []UnibeeApiMerchantMetricMetricLimitAdjustBatchItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetItemsOk() (*[]UnibeeApiMerchantMetricMetricLimitAdjustBatchItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) SetItems(v []UnibeeApiMerchantMetricMetricLimitAdjustBatchItem)`

SetItems sets Items field to given value.


### GetRequestId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UnibeeApiMerchantMetricMetricLimitAdjustBatchReq) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


