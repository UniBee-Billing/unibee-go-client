# UnibeeApiMerchantMemberNewMemberReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of member | 
**FirstName** | Pointer to **string** | The firstName of member | [optional] 
**LastName** | Pointer to **string** | The lastName of member | [optional] 
**RoleIds** | **[]int64** | The id list of role | 

## Methods

### NewUnibeeApiMerchantMemberNewMemberReq

`func NewUnibeeApiMerchantMemberNewMemberReq(email string, roleIds []int64, ) *UnibeeApiMerchantMemberNewMemberReq`

NewUnibeeApiMerchantMemberNewMemberReq instantiates a new UnibeeApiMerchantMemberNewMemberReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberNewMemberReqWithDefaults

`func NewUnibeeApiMerchantMemberNewMemberReqWithDefaults() *UnibeeApiMerchantMemberNewMemberReq`

NewUnibeeApiMerchantMemberNewMemberReqWithDefaults instantiates a new UnibeeApiMerchantMemberNewMemberReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMemberNewMemberReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantMemberNewMemberReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantMemberNewMemberReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantMemberNewMemberReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantMemberNewMemberReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRoleIds

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetRoleIds() []int64`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetRoleIdsOk() (*[]int64, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UnibeeApiMerchantMemberNewMemberReq) SetRoleIds(v []int64)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


