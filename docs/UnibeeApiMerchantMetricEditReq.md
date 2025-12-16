# UnibeeApiMerchantMetricEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to **map[string]map[string]interface{}** | meta_data(json) | [optional] 
**MetricDescription** | Pointer to **string** | MetricDescription | [optional] 
**MetricId** | **int64** | MetricId | 
**MetricName** | **string** | MetricName | 
**Type** | Pointer to **int32** | 1-limit_metered，2-charge_metered，3-charge_recurring，4-limit_recurring | [optional] 
**Unit** | Pointer to **string** | unit | [optional] 

## Methods

### NewUnibeeApiMerchantMetricEditReq

`func NewUnibeeApiMerchantMetricEditReq(metricId int64, metricName string, ) *UnibeeApiMerchantMetricEditReq`

NewUnibeeApiMerchantMetricEditReq instantiates a new UnibeeApiMerchantMetricEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricEditReqWithDefaults

`func NewUnibeeApiMerchantMetricEditReqWithDefaults() *UnibeeApiMerchantMetricEditReq`

NewUnibeeApiMerchantMetricEditReqWithDefaults instantiates a new UnibeeApiMerchantMetricEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *UnibeeApiMerchantMetricEditReq) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiMerchantMetricEditReq) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiMerchantMetricEditReq) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiMerchantMetricEditReq) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetMetricDescription

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *UnibeeApiMerchantMetricEditReq) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *UnibeeApiMerchantMetricEditReq) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricId

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricId() int64`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricIdOk() (*int64, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *UnibeeApiMerchantMetricEditReq) SetMetricId(v int64)`

SetMetricId sets MetricId field to given value.


### GetMetricName

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeApiMerchantMetricEditReq) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeApiMerchantMetricEditReq) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetType

`func (o *UnibeeApiMerchantMetricEditReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantMetricEditReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantMetricEditReq) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiMerchantMetricEditReq) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *UnibeeApiMerchantMetricEditReq) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UnibeeApiMerchantMetricEditReq) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UnibeeApiMerchantMetricEditReq) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *UnibeeApiMerchantMetricEditReq) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


