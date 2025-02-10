# UnibeeApiBeanMerchantMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Email** | Pointer to **string** | email | [optional] 
**FirstName** | Pointer to **string** | first name | [optional] 
**Id** | Pointer to **int64** | userId | [optional] 
**IsBlankPasswd** | Pointer to **bool** | is blank password | [optional] 
**IsOwner** | Pointer to **bool** | Check Member is Owner | [optional] 
**LastName** | Pointer to **string** | last name | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Mobile** | Pointer to **string** | mobile | [optional] 

## Methods

### NewUnibeeApiBeanMerchantMember

`func NewUnibeeApiBeanMerchantMember() *UnibeeApiBeanMerchantMember`

NewUnibeeApiBeanMerchantMember instantiates a new UnibeeApiBeanMerchantMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanMerchantMemberWithDefaults

`func NewUnibeeApiBeanMerchantMemberWithDefaults() *UnibeeApiBeanMerchantMember`

NewUnibeeApiBeanMerchantMemberWithDefaults instantiates a new UnibeeApiBeanMerchantMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanMerchantMember) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanMerchantMember) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanMerchantMember) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanMerchantMember) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiBeanMerchantMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiBeanMerchantMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiBeanMerchantMember) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiBeanMerchantMember) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiBeanMerchantMember) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiBeanMerchantMember) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiBeanMerchantMember) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiBeanMerchantMember) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanMerchantMember) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanMerchantMember) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanMerchantMember) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanMerchantMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBlankPasswd

`func (o *UnibeeApiBeanMerchantMember) GetIsBlankPasswd() bool`

GetIsBlankPasswd returns the IsBlankPasswd field if non-nil, zero value otherwise.

### GetIsBlankPasswdOk

`func (o *UnibeeApiBeanMerchantMember) GetIsBlankPasswdOk() (*bool, bool)`

GetIsBlankPasswdOk returns a tuple with the IsBlankPasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlankPasswd

`func (o *UnibeeApiBeanMerchantMember) SetIsBlankPasswd(v bool)`

SetIsBlankPasswd sets IsBlankPasswd field to given value.

### HasIsBlankPasswd

`func (o *UnibeeApiBeanMerchantMember) HasIsBlankPasswd() bool`

HasIsBlankPasswd returns a boolean if a field has been set.

### GetIsOwner

`func (o *UnibeeApiBeanMerchantMember) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *UnibeeApiBeanMerchantMember) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *UnibeeApiBeanMerchantMember) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *UnibeeApiBeanMerchantMember) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiBeanMerchantMember) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiBeanMerchantMember) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiBeanMerchantMember) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiBeanMerchantMember) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanMerchantMember) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanMerchantMember) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanMerchantMember) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanMerchantMember) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMobile

`func (o *UnibeeApiBeanMerchantMember) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UnibeeApiBeanMerchantMember) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UnibeeApiBeanMerchantMember) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UnibeeApiBeanMerchantMember) HasMobile() bool`

HasMobile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


