# UnibeeApiBeanDetailMerchantMemberDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberGroupPermission** | Pointer to [**map[string]UnibeeApiBeanMerchantRolePermission**](UnibeeApiBeanMerchantRolePermission.md) | The member group permission map&#39; | [optional] 
**MemberRoles** | Pointer to [**[]UnibeeApiBeanMerchantRole**](UnibeeApiBeanMerchantRole.md) | The member role list&#39; | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**Id** | Pointer to **int64** | userId | [optional] 
**IsBlankPasswd** | Pointer to **bool** | is blank password | [optional] 
**IsOwner** | Pointer to **bool** | Check Member is Owner | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 
**Status** | Pointer to **int32** | 0-Active, 2-Suspend | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantMemberDetail

`func NewUnibeeApiBeanDetailMerchantMemberDetail() *UnibeeApiBeanDetailMerchantMemberDetail`

NewUnibeeApiBeanDetailMerchantMemberDetail instantiates a new UnibeeApiBeanDetailMerchantMemberDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantMemberDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantMemberDetailWithDefaults() *UnibeeApiBeanDetailMerchantMemberDetail`

NewUnibeeApiBeanDetailMerchantMemberDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantMemberDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberGroupPermission

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMemberGroupPermission() map[string]UnibeeApiBeanMerchantRolePermission`

GetMemberGroupPermission returns the MemberGroupPermission field if non-nil, zero value otherwise.

### GetMemberGroupPermissionOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMemberGroupPermissionOk() (*map[string]UnibeeApiBeanMerchantRolePermission, bool)`

GetMemberGroupPermissionOk returns a tuple with the MemberGroupPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroupPermission

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMemberGroupPermission(v map[string]UnibeeApiBeanMerchantRolePermission)`

SetMemberGroupPermission sets MemberGroupPermission field to given value.

### HasMemberGroupPermission

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMemberGroupPermission() bool`

HasMemberGroupPermission returns a boolean if a field has been set.

### GetMemberRoles

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMemberRoles() []UnibeeApiBeanMerchantRole`

GetMemberRoles returns the MemberRoles field if non-nil, zero value otherwise.

### GetMemberRolesOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMemberRolesOk() (*[]UnibeeApiBeanMerchantRole, bool)`

GetMemberRolesOk returns a tuple with the MemberRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRoles

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMemberRoles(v []UnibeeApiBeanMerchantRole)`

SetMemberRoles sets MemberRoles field to given value.

### HasMemberRoles

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMemberRoles() bool`

HasMemberRoles returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBlankPasswd

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIsBlankPasswd() bool`

GetIsBlankPasswd returns the IsBlankPasswd field if non-nil, zero value otherwise.

### GetIsBlankPasswdOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIsBlankPasswdOk() (*bool, bool)`

GetIsBlankPasswdOk returns a tuple with the IsBlankPasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlankPasswd

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetIsBlankPasswd(v bool)`

SetIsBlankPasswd sets IsBlankPasswd field to given value.

### HasIsBlankPasswd

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasIsBlankPasswd() bool`

HasIsBlankPasswd returns a boolean if a field has been set.

### GetIsOwner

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailMerchantMemberDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


