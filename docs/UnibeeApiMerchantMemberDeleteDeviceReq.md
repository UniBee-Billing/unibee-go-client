# UnibeeApiMerchantMemberDeleteDeviceReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIdentity** | Pointer to **string** | Device Identity | [optional] 
**MemberId** | Pointer to **int64** | The unique id of member | [optional] 

## Methods

### NewUnibeeApiMerchantMemberDeleteDeviceReq

`func NewUnibeeApiMerchantMemberDeleteDeviceReq() *UnibeeApiMerchantMemberDeleteDeviceReq`

NewUnibeeApiMerchantMemberDeleteDeviceReq instantiates a new UnibeeApiMerchantMemberDeleteDeviceReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberDeleteDeviceReqWithDefaults

`func NewUnibeeApiMerchantMemberDeleteDeviceReqWithDefaults() *UnibeeApiMerchantMemberDeleteDeviceReq`

NewUnibeeApiMerchantMemberDeleteDeviceReqWithDefaults instantiates a new UnibeeApiMerchantMemberDeleteDeviceReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIdentity

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) GetDeviceIdentity() string`

GetDeviceIdentity returns the DeviceIdentity field if non-nil, zero value otherwise.

### GetDeviceIdentityOk

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) GetDeviceIdentityOk() (*string, bool)`

GetDeviceIdentityOk returns a tuple with the DeviceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentity

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) SetDeviceIdentity(v string)`

SetDeviceIdentity sets DeviceIdentity field to given value.

### HasDeviceIdentity

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) HasDeviceIdentity() bool`

HasDeviceIdentity returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiMerchantMemberDeleteDeviceReq) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


