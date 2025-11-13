# UnibeeApiBeanMerchantWebhookEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | webhook url | [optional] 
**Type** | Pointer to **int32** | webhook type,0-Webhook,1-zapier | [optional] 
**WebhookEvents** | Pointer to **[]string** | webhook_events,split dot | [optional] 
**WebhookUrl** | Pointer to **string** | webhook url | [optional] 

## Methods

### NewUnibeeApiBeanMerchantWebhookEndpoint

`func NewUnibeeApiBeanMerchantWebhookEndpoint() *UnibeeApiBeanMerchantWebhookEndpoint`

NewUnibeeApiBeanMerchantWebhookEndpoint instantiates a new UnibeeApiBeanMerchantWebhookEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantWebhookEndpointWithDefaults

`func NewUnibeeApiBeanMerchantWebhookEndpointWithDefaults() *UnibeeApiBeanMerchantWebhookEndpoint`

NewUnibeeApiBeanMerchantWebhookEndpointWithDefaults instantiates a new UnibeeApiBeanMerchantWebhookEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWebhookEvents

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.

### HasWebhookEvents

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasWebhookEvents() bool`

HasWebhookEvents returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookEndpoint) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


