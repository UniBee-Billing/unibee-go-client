# UnibeeApiMerchantMetricNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationProperty** | Pointer to **string** | AggregationProperty, Will Needed When AggregationType !&#x3D; count | [optional] 
**AggregationType** | Pointer to **int32** | AggregationType,1-count，2-count unique, 3-latest, 4-max, 5-sum | [optional] 
**Code** | **string** | Code | 
**MetricDescription** | Pointer to **string** | MetricDescription | [optional] 
**MetricName** | **string** | MetricName | 

## Methods

### NewUnibeeApiMerchantMetricNewReq

`func NewUnibeeApiMerchantMetricNewReq(code string, metricName string, ) *UnibeeApiMerchantMetricNewReq`

NewUnibeeApiMerchantMetricNewReq instantiates a new UnibeeApiMerchantMetricNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricNewReqWithDefaults

`func NewUnibeeApiMerchantMetricNewReqWithDefaults() *UnibeeApiMerchantMetricNewReq`

NewUnibeeApiMerchantMetricNewReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationProperty

`func (o *UnibeeApiMerchantMetricNewReq) GetAggregationProperty() string`

GetAggregationProperty returns the AggregationProperty field if non-nil, zero value otherwise.

### GetAggregationPropertyOk

`func (o *UnibeeApiMerchantMetricNewReq) GetAggregationPropertyOk() (*string, bool)`

GetAggregationPropertyOk returns a tuple with the AggregationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProperty

`func (o *UnibeeApiMerchantMetricNewReq) SetAggregationProperty(v string)`

SetAggregationProperty sets AggregationProperty field to given value.

### HasAggregationProperty

`func (o *UnibeeApiMerchantMetricNewReq) HasAggregationProperty() bool`

HasAggregationProperty returns a boolean if a field has been set.

### GetAggregationType

`func (o *UnibeeApiMerchantMetricNewReq) GetAggregationType() int32`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UnibeeApiMerchantMetricNewReq) GetAggregationTypeOk() (*int32, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UnibeeApiMerchantMetricNewReq) SetAggregationType(v int32)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UnibeeApiMerchantMetricNewReq) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeApiMerchantMetricNewReq) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeApiMerchantMetricNewReq) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeApiMerchantMetricNewReq) SetCode(v string)`

SetCode sets Code field to given value.


### GetMetricDescription

`func (o *UnibeeApiMerchantMetricNewReq) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *UnibeeApiMerchantMetricNewReq) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *UnibeeApiMerchantMetricNewReq) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *UnibeeApiMerchantMetricNewReq) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeApiMerchantMetricNewReq) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeApiMerchantMetricNewReq) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeApiMerchantMetricNewReq) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


