# UnibeeApiSystemInvoiceInternalWebhookSyncReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndId** | Pointer to **string** | The end Id of invoice to sync data | [optional] 
**EndTime** | Pointer to **int32** | The end time to sync data, ignore if EndId provided | [optional] 
**IsSynchronous** | Pointer to **bool** | Synchronous or not, default false | [optional] 
**StartId** | Pointer to **string** | The start Id of invoice to sync data | [optional] 
**StartTime** | Pointer to **int32** | The start time to sync data, ignore if StartId provided | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceInternalWebhookSyncReq

`func NewUnibeeApiSystemInvoiceInternalWebhookSyncReq() *UnibeeApiSystemInvoiceInternalWebhookSyncReq`

NewUnibeeApiSystemInvoiceInternalWebhookSyncReq instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceInternalWebhookSyncReqWithDefaults

`func NewUnibeeApiSystemInvoiceInternalWebhookSyncReqWithDefaults() *UnibeeApiSystemInvoiceInternalWebhookSyncReq`

NewUnibeeApiSystemInvoiceInternalWebhookSyncReqWithDefaults instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetEndId() string`

GetEndId returns the EndId field if non-nil, zero value otherwise.

### GetEndIdOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetEndIdOk() (*string, bool)`

GetEndIdOk returns a tuple with the EndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) SetEndId(v string)`

SetEndId sets EndId field to given value.

### HasEndId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) HasEndId() bool`

HasEndId returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsSynchronous

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetIsSynchronous() bool`

GetIsSynchronous returns the IsSynchronous field if non-nil, zero value otherwise.

### GetIsSynchronousOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetIsSynchronousOk() (*bool, bool)`

GetIsSynchronousOk returns a tuple with the IsSynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynchronous

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) SetIsSynchronous(v bool)`

SetIsSynchronous sets IsSynchronous field to given value.

### HasIsSynchronous

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) HasIsSynchronous() bool`

HasIsSynchronous returns a boolean if a field has been set.

### GetStartId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetStartId() string`

GetStartId returns the StartId field if non-nil, zero value otherwise.

### GetStartIdOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetStartIdOk() (*string, bool)`

GetStartIdOk returns a tuple with the StartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) SetStartId(v string)`

SetStartId sets StartId field to given value.

### HasStartId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) HasStartId() bool`

HasStartId returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncReq) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


