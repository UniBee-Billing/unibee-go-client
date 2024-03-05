# UnibeeApiMerchantMetricNewEventReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalEventId** | **string** | ExternalEventId, __unique__ | 
**ExternalUserId** | **string** | ExternalUserId | 
**MetricCode** | **string** | MetricCode | 
**MetricProperties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUnibeeApiMerchantMetricNewEventReq

`func NewUnibeeApiMerchantMetricNewEventReq(externalEventId string, externalUserId string, metricCode string, ) *UnibeeApiMerchantMetricNewEventReq`

NewUnibeeApiMerchantMetricNewEventReq instantiates a new UnibeeApiMerchantMetricNewEventReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricNewEventReqWithDefaults

`func NewUnibeeApiMerchantMetricNewEventReqWithDefaults() *UnibeeApiMerchantMetricNewEventReq`

NewUnibeeApiMerchantMetricNewEventReqWithDefaults instantiates a new UnibeeApiMerchantMetricNewEventReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


