# UnibeeInternalLogicGatewayRoMerchantMetricVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationProperty** | Pointer to **string** | aggregation property | [optional] 
**AggregationType** | Pointer to **int32** | 1-count，2-count unique, 3-latest, 4-max, 5-sum | [optional] 
**Code** | Pointer to **string** | code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int32** | id | [optional] 
**MerchantId** | Pointer to **int32** | merchantId | [optional] 
**MetricDescription** | Pointer to **string** | metric description | [optional] 
**MetricName** | Pointer to **string** | metric name | [optional] 
**Type** | Pointer to **int32** | 1-limit_metered，2-charge_metered(come later),3-charge_recurring(come later) | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoMerchantMetricVo

`func NewUnibeeInternalLogicGatewayRoMerchantMetricVo() *UnibeeInternalLogicGatewayRoMerchantMetricVo`

NewUnibeeInternalLogicGatewayRoMerchantMetricVo instantiates a new UnibeeInternalLogicGatewayRoMerchantMetricVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoMerchantMetricVoWithDefaults

`func NewUnibeeInternalLogicGatewayRoMerchantMetricVoWithDefaults() *UnibeeInternalLogicGatewayRoMerchantMetricVo`

NewUnibeeInternalLogicGatewayRoMerchantMetricVoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoMerchantMetricVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetAggregationProperty() string`

GetAggregationProperty returns the AggregationProperty field if non-nil, zero value otherwise.

### GetAggregationPropertyOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetAggregationPropertyOk() (*string, bool)`

GetAggregationPropertyOk returns a tuple with the AggregationProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetAggregationProperty(v string)`

SetAggregationProperty sets AggregationProperty field to given value.

### HasAggregationProperty

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasAggregationProperty() bool`

HasAggregationProperty returns a boolean if a field has been set.

### GetAggregationType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetAggregationType() int32`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetAggregationTypeOk() (*int32, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetAggregationType(v int32)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMerchantId() int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMerchantIdOk() (*int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetMerchantId(v int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetricDescription

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMetricDescription() string`

GetMetricDescription returns the MetricDescription field if non-nil, zero value otherwise.

### GetMetricDescriptionOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMetricDescriptionOk() (*string, bool)`

GetMetricDescriptionOk returns a tuple with the MetricDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDescription

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetMetricDescription(v string)`

SetMetricDescription sets MetricDescription field to given value.

### HasMetricDescription

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasMetricDescription() bool`

HasMetricDescription returns a boolean if a field has been set.

### GetMetricName

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalLogicGatewayRoMerchantMetricVo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


