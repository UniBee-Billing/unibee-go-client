# UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | Pointer to **string** | SubscriptionId, id of subscription, either SubscriptionId or UserId needed, The only one active subscription of userId will effect | [optional] 
**UserId** | Pointer to **int64** | UserId, either SubscriptionId or UserId needed, The only one active subscription will effect if userId provide instead of subscriptionId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq

`func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq`

NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReq instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq`

NewUnibeeApiMerchantSubscriptionCancelAtPeriodEndReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionCancelAtPeriodEndReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


