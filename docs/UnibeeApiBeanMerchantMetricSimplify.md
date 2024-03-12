# UnibeeApiBeanMerchantMetricSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationProperty** | Pointer to **string** | aggregation property | [optional] 
**AggregationType** | Pointer to **int32** | 1-count，2-count unique, 3-latest, 4-max, 5-sum | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MetricDescription** | Pointer to **string** | metric description | [optional] 
**MetricName** | Pointer to **string** | metric name | [optional] 
**Type** | Pointer to **int32** | 1-limit_metered，2-charge_metered(come later),3-charge_recurring(come later) | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMetricSimplify

`func NewUnibeeApiBeanMerchantMetricSimplify() *UnibeeApiBeanMerchantMetricSimplify`

NewUnibeeApiBeanMerchantMetricSimplify instantiates a new UnibeeApiBeanMerchantMetricSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMetricSimplifyWithDefaults

`func NewUnibeeApiBeanMerchantMetricSimplifyWithDefaults() *UnibeeApiBeanMerchantMetricSimplify`

NewUnibeeApiBeanMerchantMetricSimplifyWithDefaults instantiates a new UnibeeApiBeanMerchantMetricSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationProperty

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationProperty() string`

GetAggregationProperty returns the AggregationProperty field if non-nil, zero value otherwise.

### GetAggregationPropertyOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationPropertyOk() (*string, bool)`

GetAggregationPropertyOk returns a tuple with the AggregationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProperty

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetAggregationProperty(v string)`

SetAggregationProperty sets AggregationProperty field to given value.

### HasAggregationProperty

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasAggregationProperty() bool`

HasAggregationProperty returns a boolean if a field has been set.

### GetAggregationType

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationType() int32`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetAggregationTypeOk() (*int32, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetAggregationType(v int32)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricDescription

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanMerchantMetricSimplify) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanMerchantMetricSimplify) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanMerchantMetricSimplify) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


