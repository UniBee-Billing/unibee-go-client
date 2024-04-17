# UnibeeApiMerchantMemberNewMemberReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of member | 
**FirstName** | Pointer to **string** | The firstName of member | [optional] 
**LastName** | Pointer to **string** | The lastName of member | [optional] 
**Role** | **string** | The permission role of member | 

## Methods

### NewUnibeeApiMerchantMemberNewMemberReq

`func NewUnibeeApiMerchantMemberNewMemberReq(email string, role string, ) *UnibeeApiMerchantMemberNewMemberReq`

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

### GetRole

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UnibeeApiMerchantMemberNewMemberReq) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UnibeeApiMerchantMemberNewMemberReq) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


