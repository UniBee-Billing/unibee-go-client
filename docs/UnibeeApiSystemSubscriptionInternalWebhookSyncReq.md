# UnibeeApiSystemSubscriptionInternalWebhookSyncReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndId** | Pointer to **string** | The end auto-increase Id of Subscription to sync data | [optional] 
**EndTime** | Pointer to **int32** | The end time to sync data, ignore if EndId provided | [optional] 
**IsSynchronous** | Pointer to **bool** | Synchronous or not, default false | [optional] 
**StartId** | Pointer to **string** | The start auto-increase Id of Subscription to sync data | [optional] 
**StartTime** | Pointer to **int32** | The start time to sync data, ignore if StartId provided | [optional] 

## Methods

### NewUnibeeApiSystemSubscriptionInternalWebhookSyncReq

`func NewUnibeeApiSystemSubscriptionInternalWebhookSyncReq() *UnibeeApiSystemSubscriptionInternalWebhookSyncReq`

NewUnibeeApiSystemSubscriptionInternalWebhookSyncReq instantiates a new UnibeeApiSystemSubscriptionInternalWebhookSyncReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionInternalWebhookSyncReqWithDefaults

`func NewUnibeeApiSystemSubscriptionInternalWebhookSyncReqWithDefaults() *UnibeeApiSystemSubscriptionInternalWebhookSyncReq`

NewUnibeeApiSystemSubscriptionInternalWebhookSyncReqWithDefaults instantiates a new UnibeeApiSystemSubscriptionInternalWebhookSyncReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetEndId() string`

GetEndId returns the EndId field if non-nil, zero value otherwise.

### GetEndIdOk

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetEndIdOk() (*string, bool)`

GetEndIdOk returns a tuple with the EndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) SetEndId(v string)`

SetEndId sets EndId field to given value.

### HasEndId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) HasEndId() bool`

HasEndId returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsSynchronous

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetIsSynchronous() bool`

GetIsSynchronous returns the IsSynchronous field if non-nil, zero value otherwise.

### GetIsSynchronousOk

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetIsSynchronousOk() (*bool, bool)`

GetIsSynchronousOk returns a tuple with the IsSynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynchronous

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) SetIsSynchronous(v bool)`

SetIsSynchronous sets IsSynchronous field to given value.

### HasIsSynchronous

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) HasIsSynchronous() bool`

HasIsSynchronous returns a boolean if a field has been set.

### GetStartId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetStartId() string`

GetStartId returns the StartId field if non-nil, zero value otherwise.

### GetStartIdOk

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetStartIdOk() (*string, bool)`

GetStartIdOk returns a tuple with the StartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) SetStartId(v string)`

SetStartId sets StartId field to given value.

### HasStartId

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) HasStartId() bool`

HasStartId returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiSystemSubscriptionInternalWebhookSyncReq) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


