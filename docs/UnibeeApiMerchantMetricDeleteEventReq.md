# UnibeeApiMerchantMetricDeleteEventReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email， UserId,ExternalUserId, or Email provides one of three options | [optional] 
**ExternalEventId** | **string** | ExternalEventId | 
**ExternalUserId** | Pointer to **string** | ExternalUserId， UserId, ExternalUserId, or Email provides one of three options | [optional] 
**MetricCode** | **string** | MetricCode | 
**UserId** | Pointer to **int64** | UserId， UserId,ExternalUserId, or Email provides one of three options | [optional] 

## Methods

### NewUnibeeApiMerchantMetricDeleteEventReq

`func NewUnibeeApiMerchantMetricDeleteEventReq(externalEventId string, metricCode string, ) *UnibeeApiMerchantMetricDeleteEventReq`

NewUnibeeApiMerchantMetricDeleteEventReq instantiates a new UnibeeApiMerchantMetricDeleteEventReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricDeleteEventReqWithDefaults

`func NewUnibeeApiMerchantMetricDeleteEventReqWithDefaults() *UnibeeApiMerchantMetricDeleteEventReq`

NewUnibeeApiMerchantMetricDeleteEventReqWithDefaults instantiates a new UnibeeApiMerchantMetricDeleteEventReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMetricDeleteEventReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMetricDeleteEventReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalEventId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalEventId() string`

GetExternalEventId returns the ExternalEventId field if non-nil, zero value otherwise.

### GetExternalEventIdOk

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalEventIdOk() (*string, bool)`

GetExternalEventIdOk returns a tuple with the ExternalEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEventId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) SetExternalEventId(v string)`

SetExternalEventId sets ExternalEventId field to given value.


### GetExternalUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetMetricCode

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetMetricCode() string`

GetMetricCode returns the MetricCode field if non-nil, zero value otherwise.

### GetMetricCodeOk

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetMetricCodeOk() (*string, bool)`

GetMetricCodeOk returns a tuple with the MetricCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricCode

`func (o *UnibeeApiMerchantMetricDeleteEventReq) SetMetricCode(v string)`

SetMetricCode sets MetricCode field to given value.


### GetUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricDeleteEventReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricDeleteEventReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


