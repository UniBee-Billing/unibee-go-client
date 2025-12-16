# UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateStatus** | Pointer to [**UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus**](UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus.md) |  | [optional] 
**IsPrivateIP** | Pointer to **bool** | Whether the IP is a private/internal IP (not subject to rate limiting) | [optional] 
**PreviewStatus** | Pointer to [**UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus**](UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus.md) |  | [optional] 
**RemoteIP** | Pointer to **string** | The queried IP address | [optional] 

## Methods

### NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes

`func NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes() *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes`

NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes instantiates a new UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistResWithDefaults

`func NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistResWithDefaults() *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes`

NewUnibeeApiSystemSubscriptionQuerySubscriptionBlacklistResWithDefaults instantiates a new UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetCreateStatus() UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus`

GetCreateStatus returns the CreateStatus field if non-nil, zero value otherwise.

### GetCreateStatusOk

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetCreateStatusOk() (*UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus, bool)`

GetCreateStatusOk returns a tuple with the CreateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) SetCreateStatus(v UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus)`

SetCreateStatus sets CreateStatus field to given value.

### HasCreateStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) HasCreateStatus() bool`

HasCreateStatus returns a boolean if a field has been set.

### GetIsPrivateIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetIsPrivateIP() bool`

GetIsPrivateIP returns the IsPrivateIP field if non-nil, zero value otherwise.

### GetIsPrivateIPOk

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetIsPrivateIPOk() (*bool, bool)`

GetIsPrivateIPOk returns a tuple with the IsPrivateIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) SetIsPrivateIP(v bool)`

SetIsPrivateIP sets IsPrivateIP field to given value.

### HasIsPrivateIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) HasIsPrivateIP() bool`

HasIsPrivateIP returns a boolean if a field has been set.

### GetPreviewStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetPreviewStatus() UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus`

GetPreviewStatus returns the PreviewStatus field if non-nil, zero value otherwise.

### GetPreviewStatusOk

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetPreviewStatusOk() (*UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus, bool)`

GetPreviewStatusOk returns a tuple with the PreviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) SetPreviewStatus(v UnibeeApiSystemSubscriptionSubscriptionEndpointBlacklistStatus)`

SetPreviewStatus sets PreviewStatus field to given value.

### HasPreviewStatus

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) HasPreviewStatus() bool`

HasPreviewStatus returns a boolean if a field has been set.

### GetRemoteIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetRemoteIP() string`

GetRemoteIP returns the RemoteIP field if non-nil, zero value otherwise.

### GetRemoteIPOk

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) GetRemoteIPOk() (*string, bool)`

GetRemoteIPOk returns a tuple with the RemoteIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) SetRemoteIP(v string)`

SetRemoteIP sets RemoteIP field to given value.

### HasRemoteIP

`func (o *UnibeeApiSystemSubscriptionQuerySubscriptionBlacklistRes) HasRemoteIP() bool`

HasRemoteIP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


