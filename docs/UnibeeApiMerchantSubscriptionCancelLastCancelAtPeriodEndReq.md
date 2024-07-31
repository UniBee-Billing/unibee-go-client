# UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int64** | default product will use if productId not specified and subscriptionId is blank | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId, id of subscription, either SubscriptionId or UserId needed, The only one active subscription of userId will effect | [optional] 
**UserId** | Pointer to **int64** | UserId, either SubscriptionId or UserId needed, The only one active subscription will effect if userId provide instead of subscriptionId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq

`func NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq() *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq`

NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq instantiates a new UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReqWithDefaults() *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq`

NewUnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionCancelLastCancelAtPeriodEndReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


