# UnibeeApiBeanMerchantMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationProperty** | Pointer to **string** | aggregation property | [optional] 
**AggregationType** | Pointer to **int32** | 1-count，2-count unique, 3-latest, 4-max, 5-sum | [optional] 
**Archived** | Pointer to **bool** | archived | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | merchantId | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** | meta_data(json) | [optional] 
**MetricDescription** | Pointer to **string** | metric description | [optional] 
**MetricName** | Pointer to **string** | metric name | [optional] 
**Type** | Pointer to **int32** | 1-limit_metered，2-charge_metered,3-charge_recurring | [optional] 
**Unit** | Pointer to **string** | unit | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMetric

`func NewUnibeeApiBeanMerchantMetric() *UnibeeApiBeanMerchantMetric`

NewUnibeeApiBeanMerchantMetric instantiates a new UnibeeApiBeanMerchantMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMetricWithDefaults

`func NewUnibeeApiBeanMerchantMetricWithDefaults() *UnibeeApiBeanMerchantMetric`

NewUnibeeApiBeanMerchantMetricWithDefaults instantiates a new UnibeeApiBeanMerchantMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationProperty

`func (o *UnibeeApiBeanMerchantMetric) GetAggregationProperty() string`

GetAggregationProperty returns the AggregationProperty field if non-nil, zero value otherwise.

### GetAggregationPropertyOk

`func (o *UnibeeApiBeanMerchantMetric) GetAggregationPropertyOk() (*string, bool)`

GetAggregationPropertyOk returns a tuple with the AggregationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProperty

`func (o *UnibeeApiBeanMerchantMetric) SetAggregationProperty(v string)`

SetAggregationProperty sets AggregationProperty field to given value.

### HasAggregationProperty

`func (o *UnibeeApiBeanMerchantMetric) HasAggregationProperty() bool`

HasAggregationProperty returns a boolean if a field has been set.

### GetAggregationType

`func (o *UnibeeApiBeanMerchantMetric) GetAggregationType() int32`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UnibeeApiBeanMerchantMetric) GetAggregationTypeOk() (*int32, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UnibeeApiBeanMerchantMetric) SetAggregationType(v int32)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UnibeeApiBeanMerchantMetric) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetArchived

`func (o *UnibeeApiBeanMerchantMetric) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *UnibeeApiBeanMerchantMetric) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *UnibeeApiBeanMerchantMetric) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *UnibeeApiBeanMerchantMetric) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiBeanMerchantMetric) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiBeanMerchantMetric) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiBeanMerchantMetric) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeApiBeanMerchantMetric) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMetric) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMetric) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMetric) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMetric) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanMerchantMetric) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanMerchantMetric) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanMerchantMetric) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanMerchantMetric) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMetric) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMetric) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMetric) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMetric) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMetric) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMetric) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMetric) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiBeanMerchantMetric) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiBeanMerchantMetric) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiBeanMerchantMetric) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiBeanMerchantMetric) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMetricDescription

`func (o *UnibeeApiBeanMerchantMetric) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *UnibeeApiBeanMerchantMetric) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *UnibeeApiBeanMerchantMetric) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *UnibeeApiBeanMerchantMetric) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeApiBeanMerchantMetric) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeApiBeanMerchantMetric) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeApiBeanMerchantMetric) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UnibeeApiBeanMerchantMetric) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanMerchantMetric) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanMerchantMetric) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanMerchantMetric) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanMerchantMetric) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *UnibeeApiBeanMerchantMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UnibeeApiBeanMerchantMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UnibeeApiBeanMerchantMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UnibeeApiBeanMerchantMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


