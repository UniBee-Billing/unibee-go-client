# UnibeeApiMerchantUserDeleteReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDelete** | Pointer to **bool** | Force delete, if true, will cancel all active subscriptions before deleting the user | [optional] 
**UserId** | **int64** | UserId | 

## Methods

### NewUnibeeApiMerchantUserDeleteReq

`func NewUnibeeApiMerchantUserDeleteReq(userId int64, ) *UnibeeApiMerchantUserDeleteReq`

NewUnibeeApiMerchantUserDeleteReq instantiates a new UnibeeApiMerchantUserDeleteReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserDeleteReqWithDefaults

`func NewUnibeeApiMerchantUserDeleteReqWithDefaults() *UnibeeApiMerchantUserDeleteReq`

NewUnibeeApiMerchantUserDeleteReqWithDefaults instantiates a new UnibeeApiMerchantUserDeleteReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDelete

`func (o *UnibeeApiMerchantUserDeleteReq) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *UnibeeApiMerchantUserDeleteReq) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *UnibeeApiMerchantUserDeleteReq) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.

### HasForceDelete

`func (o *UnibeeApiMerchantUserDeleteReq) HasForceDelete() bool`

HasForceDelete returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserDeleteReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserDeleteReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserDeleteReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


