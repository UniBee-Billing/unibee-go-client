# UnibeeApiMerchantWebhookUpdateEndpointReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | **int64** | EndpointId | 
**Events** | Pointer to **[]string** | Events To Update | [optional] 
**Url** | **string** | Url To Update | 

## Methods

### NewUnibeeApiMerchantWebhookUpdateEndpointReq

`func NewUnibeeApiMerchantWebhookUpdateEndpointReq(endpointId int64, url string, ) *UnibeeApiMerchantWebhookUpdateEndpointReq`

NewUnibeeApiMerchantWebhookUpdateEndpointReq instantiates a new UnibeeApiMerchantWebhookUpdateEndpointReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantWebhookUpdateEndpointReqWithDefaults

`func NewUnibeeApiMerchantWebhookUpdateEndpointReqWithDefaults() *UnibeeApiMerchantWebhookUpdateEndpointReq`

NewUnibeeApiMerchantWebhookUpdateEndpointReqWithDefaults instantiates a new UnibeeApiMerchantWebhookUpdateEndpointReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEndpointId() int64`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEndpointIdOk() (*int64, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetEndpointId(v int64)`

SetEndpointId sets EndpointId field to given value.


### GetEvents

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetUrl

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UnibeeApiMerchantWebhookUpdateEndpointReq) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


