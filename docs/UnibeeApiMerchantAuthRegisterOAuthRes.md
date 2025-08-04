# UnibeeApiMerchantAuthRegisterOAuthRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**Token** | Pointer to **string** | Access token of admin portal | [optional] 

## Methods

### NewUnibeeApiMerchantAuthRegisterOAuthRes

`func NewUnibeeApiMerchantAuthRegisterOAuthRes() *UnibeeApiMerchantAuthRegisterOAuthRes`

NewUnibeeApiMerchantAuthRegisterOAuthRes instantiates a new UnibeeApiMerchantAuthRegisterOAuthRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthRegisterOAuthResWithDefaults

`func NewUnibeeApiMerchantAuthRegisterOAuthResWithDefaults() *UnibeeApiMerchantAuthRegisterOAuthRes`

NewUnibeeApiMerchantAuthRegisterOAuthResWithDefaults instantiates a new UnibeeApiMerchantAuthRegisterOAuthRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMember

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeApiMerchantAuthRegisterOAuthRes) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


