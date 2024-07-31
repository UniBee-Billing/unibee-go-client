# UnibeeApiBeanMerchantWebhookLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | body(json) | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**EndpointId** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**Mamo** | Pointer to **string** | mamo | [optional] 
**MerchantId** | Pointer to **int64** | webhook url | [optional] 
**ReconsumeCount** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** | request_id | [optional] 
**Response** | Pointer to **string** | response | [optional] 
**WebhookEvent** | Pointer to **string** | webhook_event | [optional] 
**WebhookUrl** | Pointer to **string** | webhook url | [optional] 

## Methods

### NewUnibeeApiBeanMerchantWebhookLog

`func NewUnibeeApiBeanMerchantWebhookLog() *UnibeeApiBeanMerchantWebhookLog`

NewUnibeeApiBeanMerchantWebhookLog instantiates a new UnibeeApiBeanMerchantWebhookLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantWebhookLogWithDefaults

`func NewUnibeeApiBeanMerchantWebhookLogWithDefaults() *UnibeeApiBeanMerchantWebhookLog`

NewUnibeeApiBeanMerchantWebhookLogWithDefaults instantiates a new UnibeeApiBeanMerchantWebhookLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *UnibeeApiBeanMerchantWebhookLog) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UnibeeApiBeanMerchantWebhookLog) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UnibeeApiBeanMerchantWebhookLog) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanMerchantWebhookLog) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantWebhookLog) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantWebhookLog) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEndpointId

`func (o *UnibeeApiBeanMerchantWebhookLog) GetEndpointId() int64`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetEndpointIdOk() (*int64, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *UnibeeApiBeanMerchantWebhookLog) SetEndpointId(v int64)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *UnibeeApiBeanMerchantWebhookLog) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantWebhookLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantWebhookLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantWebhookLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMamo

`func (o *UnibeeApiBeanMerchantWebhookLog) GetMamo() string`

GetMamo returns the Mamo field if non-nil, zero value otherwise.

### GetMamoOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetMamoOk() (*string, bool)`

GetMamoOk returns a tuple with the Mamo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMamo

`func (o *UnibeeApiBeanMerchantWebhookLog) SetMamo(v string)`

SetMamo sets Mamo field to given value.

### HasMamo

`func (o *UnibeeApiBeanMerchantWebhookLog) HasMamo() bool`

HasMamo returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantWebhookLog) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantWebhookLog) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantWebhookLog) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetReconsumeCount

`func (o *UnibeeApiBeanMerchantWebhookLog) GetReconsumeCount() int32`

GetReconsumeCount returns the ReconsumeCount field if non-nil, zero value otherwise.

### GetReconsumeCountOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetReconsumeCountOk() (*int32, bool)`

GetReconsumeCountOk returns a tuple with the ReconsumeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconsumeCount

`func (o *UnibeeApiBeanMerchantWebhookLog) SetReconsumeCount(v int32)`

SetReconsumeCount sets ReconsumeCount field to given value.

### HasReconsumeCount

`func (o *UnibeeApiBeanMerchantWebhookLog) HasReconsumeCount() bool`

HasReconsumeCount returns a boolean if a field has been set.

### GetRequestId

`func (o *UnibeeApiBeanMerchantWebhookLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *UnibeeApiBeanMerchantWebhookLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *UnibeeApiBeanMerchantWebhookLog) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetResponse

`func (o *UnibeeApiBeanMerchantWebhookLog) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *UnibeeApiBeanMerchantWebhookLog) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *UnibeeApiBeanMerchantWebhookLog) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetWebhookEvent

`func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookEvent() string`

GetWebhookEvent returns the WebhookEvent field if non-nil, zero value otherwise.

### GetWebhookEventOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookEventOk() (*string, bool)`

GetWebhookEventOk returns a tuple with the WebhookEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEvent

`func (o *UnibeeApiBeanMerchantWebhookLog) SetWebhookEvent(v string)`

SetWebhookEvent sets WebhookEvent field to given value.

### HasWebhookEvent

`func (o *UnibeeApiBeanMerchantWebhookLog) HasWebhookEvent() bool`

HasWebhookEvent returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UnibeeApiBeanMerchantWebhookLog) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookLog) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UnibeeApiBeanMerchantWebhookLog) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


