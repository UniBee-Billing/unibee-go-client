# UnibeeApiMerchantMetricNewEventReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationUniqueId** | Pointer to **string** | AggregationUniqueId, valid when AggregationType is count unique | [optional] 
**AggregationValue** | Pointer to **int32** | AggregationValue, valid when AggregationType latest, max or sum | [optional] 
**Email** | Pointer to **string** | Email， UserId, ExternalUserId, or Email provides one of three options | [optional] [default to "account@unibee.dev"]
**ExternalEventId** | **string** | ExternalEventId, __unique__ | 
**ExternalUserId** | Pointer to **string** | ExternalUserId， UserId, ExternalUserId, or Email provides one of three options | [optional] 
**MetricCode** | **string** | MetricCode | 
**MetricProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**ProductId** | Pointer to **int64** | default product will use if productId not specified and subscriptionId is blank | [optional] 
**UserId** | Pointer to **int64** | UserId， UserId, ExternalUserId, or Email provides one of three options | [optional] 

## Methods

### NewUnibeeApiMerchantMetricNewEventReq

`func NewUnibeeApiMerchantMetricNewEventReq(externalEventId string, metricCode string, ) *UnibeeApiMerchantMetricNewEventReq`

NewUnibeeApiMerchantMetricNewEventReq instantiates a new UnibeeApiMerchantMetricNewEventReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricNewEventReqWithDefaults

`func NewUnibeeApiMerchantMetricNewEventReqWithDefaults() *UnibeeApiMerchantMetricNewEventReq`

NewUnibeeApiMerchantMetricNewEventReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewEventReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationUniqueId

`func (o *UnibeeApiMerchantMetricNewEventReq) GetAggregationUniqueId() string`

GetAggregationUniqueId returns the AggregationUniqueId field if non-nil, zero value otherwise.

### GetAggregationUniqueIdOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetAggregationUniqueIdOk() (*string, bool)`

GetAggregationUniqueIdOk returns a tuple with the AggregationUniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationUniqueId

`func (o *UnibeeApiMerchantMetricNewEventReq) SetAggregationUniqueId(v string)`

SetAggregationUniqueId sets AggregationUniqueId field to given value.

### HasAggregationUniqueId

`func (o *UnibeeApiMerchantMetricNewEventReq) HasAggregationUniqueId() bool`

HasAggregationUniqueId returns a boolean if a field has been set.

### GetAggregationValue

`func (o *UnibeeApiMerchantMetricNewEventReq) GetAggregationValue() int32`

GetAggregationValue returns the AggregationValue field if non-nil, zero value otherwise.

### GetAggregationValueOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetAggregationValueOk() (*int32, bool)`

GetAggregationValueOk returns a tuple with the AggregationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationValue

`func (o *UnibeeApiMerchantMetricNewEventReq) SetAggregationValue(v int32)`

SetAggregationValue sets AggregationValue field to given value.

### HasAggregationValue

`func (o *UnibeeApiMerchantMetricNewEventReq) HasAggregationValue() bool`

HasAggregationValue returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantMetricNewEventReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMetricNewEventReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMetricNewEventReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeApiMerchantMetricNewEventReq) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.


### GetExternalUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricNewEventReq) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.


### GetMetricProperties

`func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricProperties() map[string]interface{}`

GetMetricProperties returns the MetricProperties field if non-nil, zero value otherwise.

### GetMetricPropertiesOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetMetricPropertiesOk() (*map[string]interface{}, bool)`

GetMetricPropertiesOk returns a tuple with the MetricProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricProperties

`func (o *UnibeeApiMerchantMetricNewEventReq) SetMetricProperties(v map[string]interface{})`

SetMetricProperties sets MetricProperties field to given value.

### HasMetricProperties

`func (o *UnibeeApiMerchantMetricNewEventReq) HasMetricProperties() bool`

HasMetricProperties returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantMetricNewEventReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricNewEventReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricNewEventReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricNewEventReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricNewEventReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


