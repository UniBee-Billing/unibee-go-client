# UnibeeApiBeanMerchantMemberDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentDevice** | Pointer to **bool** | Is CurrentDevice | [optional] 
**Identity** | Pointer to **string** | Identity | [optional] 
**IpAddress** | Pointer to **string** | IP Address | [optional] 
**LastActiveTime** | Pointer to **int64** | Last Active Time | [optional] 
**LastLoginTime** | Pointer to **int64** | Last Login Time | [optional] 
**LastTotpVerificationTime** | Pointer to **int64** | Last Totp Verification Time | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Status** | Pointer to **bool** | true-Active, false-Offline | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMemberDevice

`func NewUnibeeApiBeanMerchantMemberDevice() *UnibeeApiBeanMerchantMemberDevice`

NewUnibeeApiBeanMerchantMemberDevice instantiates a new UnibeeApiBeanMerchantMemberDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMemberDeviceWithDefaults

`func NewUnibeeApiBeanMerchantMemberDeviceWithDefaults() *UnibeeApiBeanMerchantMemberDevice`

NewUnibeeApiBeanMerchantMemberDeviceWithDefaults instantiates a new UnibeeApiBeanMerchantMemberDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentDevice

`func (o *UnibeeApiBeanMerchantMemberDevice) GetCurrentDevice() bool`

GetCurrentDevice returns the CurrentDevice field if non-nil, zero value otherwise.

### GetCurrentDeviceOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetCurrentDeviceOk() (*bool, bool)`

GetCurrentDeviceOk returns a tuple with the CurrentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDevice

`func (o *UnibeeApiBeanMerchantMemberDevice) SetCurrentDevice(v bool)`

SetCurrentDevice sets CurrentDevice field to given value.

### HasCurrentDevice

`func (o *UnibeeApiBeanMerchantMemberDevice) HasCurrentDevice() bool`

HasCurrentDevice returns a boolean if a field has been set.

### GetIdentity

`func (o *UnibeeApiBeanMerchantMemberDevice) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *UnibeeApiBeanMerchantMemberDevice) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *UnibeeApiBeanMerchantMemberDevice) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpAddress

`func (o *UnibeeApiBeanMerchantMemberDevice) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UnibeeApiBeanMerchantMemberDevice) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UnibeeApiBeanMerchantMemberDevice) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetLastActiveTime

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastActiveTime() int64`

GetLastActiveTime returns the LastActiveTime field if non-nil, zero value otherwise.

### GetLastActiveTimeOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastActiveTimeOk() (*int64, bool)`

GetLastActiveTimeOk returns a tuple with the LastActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveTime

`func (o *UnibeeApiBeanMerchantMemberDevice) SetLastActiveTime(v int64)`

SetLastActiveTime sets LastActiveTime field to given value.

### HasLastActiveTime

`func (o *UnibeeApiBeanMerchantMemberDevice) HasLastActiveTime() bool`

HasLastActiveTime returns a boolean if a field has been set.

### GetLastLoginTime

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastLoginTime() int64`

GetLastLoginTime returns the LastLoginTime field if non-nil, zero value otherwise.

### GetLastLoginTimeOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastLoginTimeOk() (*int64, bool)`

GetLastLoginTimeOk returns a tuple with the LastLoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTime

`func (o *UnibeeApiBeanMerchantMemberDevice) SetLastLoginTime(v int64)`

SetLastLoginTime sets LastLoginTime field to given value.

### HasLastLoginTime

`func (o *UnibeeApiBeanMerchantMemberDevice) HasLastLoginTime() bool`

HasLastLoginTime returns a boolean if a field has been set.

### GetLastTotpVerificationTime

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastTotpVerificationTime() int64`

GetLastTotpVerificationTime returns the LastTotpVerificationTime field if non-nil, zero value otherwise.

### GetLastTotpVerificationTimeOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetLastTotpVerificationTimeOk() (*int64, bool)`

GetLastTotpVerificationTimeOk returns a tuple with the LastTotpVerificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTotpVerificationTime

`func (o *UnibeeApiBeanMerchantMemberDevice) SetLastTotpVerificationTime(v int64)`

SetLastTotpVerificationTime sets LastTotpVerificationTime field to given value.

### HasLastTotpVerificationTime

`func (o *UnibeeApiBeanMerchantMemberDevice) HasLastTotpVerificationTime() bool`

HasLastTotpVerificationTime returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanMerchantMemberDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanMerchantMemberDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanMerchantMemberDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanMerchantMemberDevice) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanMerchantMemberDevice) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanMerchantMemberDevice) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanMerchantMemberDevice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


