# UnibeeApiMerchantRoleEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | [**[]UnibeeApiBeanMerchantRolePermission**](UnibeeApiBeanMerchantRolePermission.md) | Permissions | 
**Role** | **string** | Role | 

## Methods

### NewUnibeeApiMerchantRoleEditReq

`func NewUnibeeApiMerchantRoleEditReq(permissions []UnibeeApiBeanMerchantRolePermission, role string, ) *UnibeeApiMerchantRoleEditReq`

NewUnibeeApiMerchantRoleEditReq instantiates a new UnibeeApiMerchantRoleEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantRoleEditReqWithDefaults

`func NewUnibeeApiMerchantRoleEditReqWithDefaults() *UnibeeApiMerchantRoleEditReq`

NewUnibeeApiMerchantRoleEditReqWithDefaults instantiates a new UnibeeApiMerchantRoleEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *UnibeeApiMerchantRoleEditReq) GetPermissions() []UnibeeApiBeanMerchantRolePermission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UnibeeApiMerchantRoleEditReq) GetPermissionsOk() (*[]UnibeeApiBeanMerchantRolePermission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UnibeeApiMerchantRoleEditReq) SetPermissions(v []UnibeeApiBeanMerchantRolePermission)`

SetPermissions sets Permissions field to given value.


### GetRole

`func (o *UnibeeApiMerchantRoleEditReq) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UnibeeApiMerchantRoleEditReq) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UnibeeApiMerchantRoleEditReq) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


