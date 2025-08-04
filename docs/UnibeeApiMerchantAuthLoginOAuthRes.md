# UnibeeApiMerchantAuthLoginOAuthRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**Token** | Pointer to **string** | Access token of admin portal | [optional] 

## Methods

### NewUnibeeApiMerchantAuthLoginOAuthRes

`func NewUnibeeApiMerchantAuthLoginOAuthRes() *UnibeeApiMerchantAuthLoginOAuthRes`

NewUnibeeApiMerchantAuthLoginOAuthRes instantiates a new UnibeeApiMerchantAuthLoginOAuthRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthLoginOAuthResWithDefaults

`func NewUnibeeApiMerchantAuthLoginOAuthResWithDefaults() *UnibeeApiMerchantAuthLoginOAuthRes`

NewUnibeeApiMerchantAuthLoginOAuthResWithDefaults instantiates a new UnibeeApiMerchantAuthLoginOAuthRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMember

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeApiMerchantAuthLoginOAuthRes) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


