# UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int64** | Product id for plan scope; default product when omitted or 0, same as merchant user_subscription_detail | [optional] 
**UserId** | **int64** | User id | 

## Methods

### NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq

`func NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq(userId int64, ) *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq`

NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq instantiates a new UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReqWithDefaults

`func NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReqWithDefaults() *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq`

NewUnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReqWithDefaults instantiates a new UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiSystemSubscriptionSendUserSubUpdateWebhookEventReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


