# UnibeeApiMerchantMetricUserMetricReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Email, One Of UserId|Email|ExternalUserId Needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, One Of UserId|Email|ExternalUserId Needed | [optional] 
**ProductId** | Pointer to **int64** | Id of product, default product will use if productId not specified and subscriptionId is blank | [optional] 
**ReloadCache** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **int64** | UserId, One Of UserId|Email|ExternalUserId Needed | [optional] 

## Methods

### NewUnibeeApiMerchantMetricUserMetricReq

`func NewUnibeeApiMerchantMetricUserMetricReq() *UnibeeApiMerchantMetricUserMetricReq`

NewUnibeeApiMerchantMetricUserMetricReq instantiates a new UnibeeApiMerchantMetricUserMetricReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMetricUserMetricReqWithDefaults

`func NewUnibeeApiMerchantMetricUserMetricReqWithDefaults() *UnibeeApiMerchantMetricUserMetricReq`

NewUnibeeApiMerchantMetricUserMetricReqWithDefaults instantiates a new UnibeeApiMerchantMetricUserMetricReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMetricUserMetricReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMetricUserMetricReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantMetricUserMetricReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantMetricUserMetricReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReloadCache

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetReloadCache() bool`

GetReloadCache returns the ReloadCache field if non-nil, zero value otherwise.

### GetReloadCacheOk

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetReloadCacheOk() (*bool, bool)`

GetReloadCacheOk returns a tuple with the ReloadCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReloadCache

`func (o *UnibeeApiMerchantMetricUserMetricReq) SetReloadCache(v bool)`

SetReloadCache sets ReloadCache field to given value.

### HasReloadCache

`func (o *UnibeeApiMerchantMetricUserMetricReq) HasReloadCache() bool`

HasReloadCache returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantMetricUserMetricReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantMetricUserMetricReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


