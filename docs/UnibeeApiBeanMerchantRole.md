# UnibeeApiBeanMerchantRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Permissions** | Pointer to [**[]UnibeeApiBeanMerchantRolePermission**](UnibeeApiBeanMerchantRolePermission.md) | permissions | [optional] 
**Role** | Pointer to **string** | role | [optional] 

## Methods

### NewUnibeeApiBeanMerchantRole

`func NewUnibeeApiBeanMerchantRole() *UnibeeApiBeanMerchantRole`

NewUnibeeApiBeanMerchantRole instantiates a new UnibeeApiBeanMerchantRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantRoleWithDefaults

`func NewUnibeeApiBeanMerchantRoleWithDefaults() *UnibeeApiBeanMerchantRole`

NewUnibeeApiBeanMerchantRoleWithDefaults instantiates a new UnibeeApiBeanMerchantRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantRole) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantRole) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantRole) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantRole) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantRole) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantRole) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantRole) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantRole) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantRole) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantRole) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantRole) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPermissions

`func (o *UnibeeApiBeanMerchantRole) GetPermissions() []UnibeeApiBeanMerchantRolePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UnibeeApiBeanMerchantRole) GetPermissionsOk() (*[]UnibeeApiBeanMerchantRolePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UnibeeApiBeanMerchantRole) SetPermissions(v []UnibeeApiBeanMerchantRolePermission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UnibeeApiBeanMerchantRole) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRole

`func (o *UnibeeApiBeanMerchantRole) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UnibeeApiBeanMerchantRole) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UnibeeApiBeanMerchantRole) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UnibeeApiBeanMerchantRole) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


