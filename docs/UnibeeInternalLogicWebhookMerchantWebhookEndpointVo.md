# UnibeeInternalLogicWebhookMerchantWebhookEndpointVo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**GmtModify** | Pointer to **int64** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | webhook url | [optional] 
**WebhookEvents** | Pointer to **[]string** | webhook_events,split dot | [optional] 
**WebhookUrl** | Pointer to **string** | webhook url | [optional] 

## Methods

### NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVo

`func NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVo() *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo`

NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVo instantiates a new UnibeeInternalLogicWebhookMerchantWebhookEndpointVo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVoWithDefaults

`func NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVoWithDefaults() *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo`

NewUnibeeInternalLogicWebhookMerchantWebhookEndpointVoWithDefaults instantiates a new UnibeeInternalLogicWebhookMerchantWebhookEndpointVo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetGmtModify() int64`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetGmtModifyOk() (*int64, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetGmtModify(v int64)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetWebhookEvents

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetWebhookEvents() []string`

GetWebhookEvents returns the WebhookEvents field if non-nil, zero value otherwise.

### GetWebhookEventsOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetWebhookEventsOk() (*[]string, bool)`

GetWebhookEventsOk returns a tuple with the WebhookEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvents

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetWebhookEvents(v []string)`

SetWebhookEvents sets WebhookEvents field to given value.

### HasWebhookEvents

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasWebhookEvents() bool`

HasWebhookEvents returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UnibeeInternalLogicWebhookMerchantWebhookEndpointVo) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


