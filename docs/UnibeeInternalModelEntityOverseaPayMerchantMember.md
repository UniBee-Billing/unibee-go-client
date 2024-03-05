# UnibeeInternalModelEntityOverseaPayMerchantMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int32** | userId | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**MerchantId** | Pointer to **int32** | merchant id | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 
**Password** | Pointer to **string** | password | [optional] 
**UserName** | Pointer to **string** | user name | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayMerchantMember

`func NewUnibeeInternalModelEntityOverseaPayMerchantMember() *UnibeeInternalModelEntityOverseaPayMerchantMember`

NewUnibeeInternalModelEntityOverseaPayMerchantMember instantiates a new UnibeeInternalModelEntityOverseaPayMerchantMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayMerchantMemberWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayMerchantMemberWithDefaults() *UnibeeInternalModelEntityOverseaPayMerchantMember`

NewUnibeeInternalModelEntityOverseaPayMerchantMemberWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayMerchantMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetMerchantId() int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetMerchantIdOk() (*int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetMerchantId(v int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetPassword

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUserName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UnibeeInternalModelEntityOverseaPayMerchantMember) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


