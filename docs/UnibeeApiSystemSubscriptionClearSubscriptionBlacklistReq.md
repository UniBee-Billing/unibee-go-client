# UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Admin password for blacklist management operations | 
**RemoteIP** | **string** | The IP address to clear from blacklist (clears both Preview and Create endpoints, including rate limit counters) | 

## Methods

### NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq

`func NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq(password string, remoteIP string, ) *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq`

NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq instantiates a new UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReqWithDefaults

`func NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReqWithDefaults() *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq`

NewUnibeeApiSystemSubscriptionClearSubscriptionBlacklistReqWithDefaults instantiates a new UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRemoteIP

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) GetRemoteIP() string`

GetRemoteIP returns the RemoteIP field if non-nil, zero value otherwise.

### GetRemoteIPOk

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) GetRemoteIPOk() (*string, bool)`

GetRemoteIPOk returns a tuple with the RemoteIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIP

`func (o *UnibeeApiSystemSubscriptionClearSubscriptionBlacklistReq) SetRemoteIP(v string)`

SetRemoteIP sets RemoteIP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


