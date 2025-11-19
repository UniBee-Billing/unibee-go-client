# UnibeeApiBeanMerchantZapierCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | zapier access token | [optional] 
**CreatedAt** | Pointer to **int64** | token created at | [optional] 
**ExpiredAt** | Pointer to **int64** | token expired at | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** | additional zapier info | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**RefreshToken** | Pointer to **string** | zapier refresh token | [optional] 
**Scope** | Pointer to **[]string** | zapier scopes | [optional] 
**Subscriptions** | Pointer to [**[]UnibeeApiBeanMerchantZapierSubscription**](UnibeeApiBeanMerchantZapierSubscription.md) | zapier subscriptions | [optional] 
**TokenType** | Pointer to **string** | zapier token type | [optional] 
**UpdatedAt** | Pointer to **int64** | token updated at | [optional] 

## Methods

### NewUnibeeApiBeanMerchantZapierCredential

`func NewUnibeeApiBeanMerchantZapierCredential() *UnibeeApiBeanMerchantZapierCredential`

NewUnibeeApiBeanMerchantZapierCredential instantiates a new UnibeeApiBeanMerchantZapierCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantZapierCredentialWithDefaults

`func NewUnibeeApiBeanMerchantZapierCredentialWithDefaults() *UnibeeApiBeanMerchantZapierCredential`

NewUnibeeApiBeanMerchantZapierCredentialWithDefaults instantiates a new UnibeeApiBeanMerchantZapierCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *UnibeeApiBeanMerchantZapierCredential) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *UnibeeApiBeanMerchantZapierCredential) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *UnibeeApiBeanMerchantZapierCredential) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiredAt

`func (o *UnibeeApiBeanMerchantZapierCredential) GetExpiredAt() int64`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetExpiredAtOk() (*int64, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *UnibeeApiBeanMerchantZapierCredential) SetExpiredAt(v int64)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *UnibeeApiBeanMerchantZapierCredential) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetExtra

`func (o *UnibeeApiBeanMerchantZapierCredential) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *UnibeeApiBeanMerchantZapierCredential) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *UnibeeApiBeanMerchantZapierCredential) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantZapierCredential) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantZapierCredential) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantZapierCredential) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetRefreshToken

`func (o *UnibeeApiBeanMerchantZapierCredential) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *UnibeeApiBeanMerchantZapierCredential) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *UnibeeApiBeanMerchantZapierCredential) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *UnibeeApiBeanMerchantZapierCredential) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UnibeeApiBeanMerchantZapierCredential) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *UnibeeApiBeanMerchantZapierCredential) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSubscriptions

`func (o *UnibeeApiBeanMerchantZapierCredential) GetSubscriptions() []UnibeeApiBeanMerchantZapierSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetSubscriptionsOk() (*[]UnibeeApiBeanMerchantZapierSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UnibeeApiBeanMerchantZapierCredential) SetSubscriptions(v []UnibeeApiBeanMerchantZapierSubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UnibeeApiBeanMerchantZapierCredential) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTokenType

`func (o *UnibeeApiBeanMerchantZapierCredential) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *UnibeeApiBeanMerchantZapierCredential) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *UnibeeApiBeanMerchantZapierCredential) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UnibeeApiBeanMerchantZapierCredential) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UnibeeApiBeanMerchantZapierCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


