# UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLevel** | Pointer to **int32** | Current punishment level: 0&#x3D;No restriction, 1&#x3D;Temporary rate limit (5min), 2&#x3D;Short-term blacklist (1h), 3&#x3D;Long-term blacklist (24h) | [optional] 
**Endpoint** | Pointer to **string** | Endpoint name (Preview or Create) | [optional] 
**LevelDescription** | Pointer to **string** | Description of current level | [optional] 
**NextLevelWarning** | Pointer to **string** | Warning about what happens if violated again | [optional] 
**PunishmentExpiry** | Pointer to **int64** | Unix timestamp when current punishment expires (0 if no punishment) | [optional] 
**RateLimitCount** | Pointer to **int32** | Current rate limit counter (requests in current 1-minute window) | [optional] 
**RateLimitMax** | Pointer to **int32** | Maximum allowed requests per minute for this endpoint | [optional] 
**RemainingRequests** | Pointer to **int32** | Remaining requests available in current window (0 if blacklisted) | [optional] 
**ViolationCount** | Pointer to **int32** | Number of violations in the past 5 minutes | [optional] 
**ViolationExpiry** | Pointer to **int64** | Unix timestamp when violation count resets (0 if no violations) | [optional] 

## Methods

### NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus

`func NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus() *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus`

NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus instantiates a new UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatusWithDefaults

`func NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatusWithDefaults() *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus`

NewUnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatusWithDefaults instantiates a new UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLevel

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetCurrentLevel() int32`

GetCurrentLevel returns the CurrentLevel field if non-nil, zero value otherwise.

### GetCurrentLevelOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetCurrentLevelOk() (*int32, bool)`

GetCurrentLevelOk returns a tuple with the CurrentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLevel

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetCurrentLevel(v int32)`

SetCurrentLevel sets CurrentLevel field to given value.

### HasCurrentLevel

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasCurrentLevel() bool`

HasCurrentLevel returns a boolean if a field has been set.

### GetEndpoint

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetLevelDescription

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetLevelDescription() string`

GetLevelDescription returns the LevelDescription field if non-nil, zero value otherwise.

### GetLevelDescriptionOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetLevelDescriptionOk() (*string, bool)`

GetLevelDescriptionOk returns a tuple with the LevelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelDescription

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetLevelDescription(v string)`

SetLevelDescription sets LevelDescription field to given value.

### HasLevelDescription

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasLevelDescription() bool`

HasLevelDescription returns a boolean if a field has been set.

### GetNextLevelWarning

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetNextLevelWarning() string`

GetNextLevelWarning returns the NextLevelWarning field if non-nil, zero value otherwise.

### GetNextLevelWarningOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetNextLevelWarningOk() (*string, bool)`

GetNextLevelWarningOk returns a tuple with the NextLevelWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLevelWarning

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetNextLevelWarning(v string)`

SetNextLevelWarning sets NextLevelWarning field to given value.

### HasNextLevelWarning

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasNextLevelWarning() bool`

HasNextLevelWarning returns a boolean if a field has been set.

### GetPunishmentExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetPunishmentExpiry() int64`

GetPunishmentExpiry returns the PunishmentExpiry field if non-nil, zero value otherwise.

### GetPunishmentExpiryOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetPunishmentExpiryOk() (*int64, bool)`

GetPunishmentExpiryOk returns a tuple with the PunishmentExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPunishmentExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetPunishmentExpiry(v int64)`

SetPunishmentExpiry sets PunishmentExpiry field to given value.

### HasPunishmentExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasPunishmentExpiry() bool`

HasPunishmentExpiry returns a boolean if a field has been set.

### GetRateLimitCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRateLimitCount() int32`

GetRateLimitCount returns the RateLimitCount field if non-nil, zero value otherwise.

### GetRateLimitCountOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRateLimitCountOk() (*int32, bool)`

GetRateLimitCountOk returns a tuple with the RateLimitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetRateLimitCount(v int32)`

SetRateLimitCount sets RateLimitCount field to given value.

### HasRateLimitCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasRateLimitCount() bool`

HasRateLimitCount returns a boolean if a field has been set.

### GetRateLimitMax

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRateLimitMax() int32`

GetRateLimitMax returns the RateLimitMax field if non-nil, zero value otherwise.

### GetRateLimitMaxOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRateLimitMaxOk() (*int32, bool)`

GetRateLimitMaxOk returns a tuple with the RateLimitMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitMax

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetRateLimitMax(v int32)`

SetRateLimitMax sets RateLimitMax field to given value.

### HasRateLimitMax

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasRateLimitMax() bool`

HasRateLimitMax returns a boolean if a field has been set.

### GetRemainingRequests

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRemainingRequests() int32`

GetRemainingRequests returns the RemainingRequests field if non-nil, zero value otherwise.

### GetRemainingRequestsOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetRemainingRequestsOk() (*int32, bool)`

GetRemainingRequestsOk returns a tuple with the RemainingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRequests

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetRemainingRequests(v int32)`

SetRemainingRequests sets RemainingRequests field to given value.

### HasRemainingRequests

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasRemainingRequests() bool`

HasRemainingRequests returns a boolean if a field has been set.

### GetViolationCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetViolationCount() int32`

GetViolationCount returns the ViolationCount field if non-nil, zero value otherwise.

### GetViolationCountOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetViolationCountOk() (*int32, bool)`

GetViolationCountOk returns a tuple with the ViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetViolationCount(v int32)`

SetViolationCount sets ViolationCount field to given value.

### HasViolationCount

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasViolationCount() bool`

HasViolationCount returns a boolean if a field has been set.

### GetViolationExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetViolationExpiry() int64`

GetViolationExpiry returns the ViolationExpiry field if non-nil, zero value otherwise.

### GetViolationExpiryOk

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) GetViolationExpiryOk() (*int64, bool)`

GetViolationExpiryOk returns a tuple with the ViolationExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) SetViolationExpiry(v int64)`

SetViolationExpiry sets ViolationExpiry field to given value.

### HasViolationExpiry

`func (o *UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus) HasViolationExpiry() bool`

HasViolationExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


