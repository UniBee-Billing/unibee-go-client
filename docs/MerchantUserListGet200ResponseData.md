# MerchantUserListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total | [optional] 
**UserAccounts** | Pointer to [**[]UnibeeApiBeanDetailUserAccountDetail**](UnibeeApiBeanDetailUserAccountDetail.md) | User Account Object List | [optional] 

## Methods

### NewMerchantUserListGet200ResponseData

`func NewMerchantUserListGet200ResponseData() *MerchantUserListGet200ResponseData`

NewMerchantUserListGet200ResponseData instantiates a new MerchantUserListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantUserListGet200ResponseDataWithDefaults

`func NewMerchantUserListGet200ResponseDataWithDefaults() *MerchantUserListGet200ResponseData`

NewMerchantUserListGet200ResponseDataWithDefaults instantiates a new MerchantUserListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *MerchantUserListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantUserListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantUserListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantUserListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUserAccounts

`func (o *MerchantUserListGet200ResponseData) GetUserAccounts() []UnibeeApiBeanDetailUserAccountDetail`

GetUserAccounts returns the UserAccounts field if non-nil, zero value otherwise.

### GetUserAccountsOk

`func (o *MerchantUserListGet200ResponseData) GetUserAccountsOk() (*[]UnibeeApiBeanDetailUserAccountDetail, bool)`

GetUserAccountsOk returns a tuple with the UserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccounts

`func (o *MerchantUserListGet200ResponseData) SetUserAccounts(v []UnibeeApiBeanDetailUserAccountDetail)`

SetUserAccounts sets UserAccounts field to given value.

### HasUserAccounts

`func (o *MerchantUserListGet200ResponseData) HasUserAccounts() bool`

HasUserAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


