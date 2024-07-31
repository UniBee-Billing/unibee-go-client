# UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ProductId** | Pointer to **int64** | default product will use if productId not specified and subscriptionId is blank | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq

`func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq`

NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReq instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults() *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq`

NewUnibeeApiMerchantSubscriptionUserSubscriptionDetailReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionUserSubscriptionDetailReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


