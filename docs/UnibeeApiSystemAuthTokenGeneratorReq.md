# UnibeeApiSystemAuthTokenGeneratorReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | [default to ""]
**Env** | Pointer to **string** | default daily | [optional] [default to "daily"]
**MerchantId** | Pointer to **int64** |  | [optional] [default to 15621]
**PortalType** | Pointer to **int64** | 0-Admin Portal, 1-User Portal, Default 0 | [optional] [default to 0]
**RedisDatabase** | Pointer to **int32** | default 1 | [optional] [default to 1]

## Methods

### NewUnibeeApiSystemAuthTokenGeneratorReq

`func NewUnibeeApiSystemAuthTokenGeneratorReq(email string, ) *UnibeeApiSystemAuthTokenGeneratorReq`

NewUnibeeApiSystemAuthTokenGeneratorReq instantiates a new UnibeeApiSystemAuthTokenGeneratorReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemAuthTokenGeneratorReqWithDefaults

`func NewUnibeeApiSystemAuthTokenGeneratorReqWithDefaults() *UnibeeApiSystemAuthTokenGeneratorReq`

NewUnibeeApiSystemAuthTokenGeneratorReqWithDefaults instantiates a new UnibeeApiSystemAuthTokenGeneratorReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEnv

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPortalType

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetPortalType() int64`

GetPortalType returns the PortalType field if non-nil, zero value otherwise.

### GetPortalTypeOk

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetPortalTypeOk() (*int64, bool)`

GetPortalTypeOk returns a tuple with the PortalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalType

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetPortalType(v int64)`

SetPortalType sets PortalType field to given value.

### HasPortalType

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasPortalType() bool`

HasPortalType returns a boolean if a field has been set.

### GetRedisDatabase

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetRedisDatabase() int32`

GetRedisDatabase returns the RedisDatabase field if non-nil, zero value otherwise.

### GetRedisDatabaseOk

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) GetRedisDatabaseOk() (*int32, bool)`

GetRedisDatabaseOk returns a tuple with the RedisDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedisDatabase

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) SetRedisDatabase(v int32)`

SetRedisDatabase sets RedisDatabase field to given value.

### HasRedisDatabase

`func (o *UnibeeApiSystemAuthTokenGeneratorReq) HasRedisDatabase() bool`

HasRedisDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


