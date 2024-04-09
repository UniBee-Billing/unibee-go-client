# UnibeeApiMerchantAuthLoginRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**Token** | Pointer to **string** | Token | [optional] 

## Methods

### NewUnibeeApiMerchantAuthLoginRes

`func NewUnibeeApiMerchantAuthLoginRes() *UnibeeApiMerchantAuthLoginRes`

NewUnibeeApiMerchantAuthLoginRes instantiates a new UnibeeApiMerchantAuthLoginRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthLoginResWithDefaults

`func NewUnibeeApiMerchantAuthLoginResWithDefaults() *UnibeeApiMerchantAuthLoginRes`

NewUnibeeApiMerchantAuthLoginResWithDefaults instantiates a new UnibeeApiMerchantAuthLoginRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMember

`func (o *UnibeeApiMerchantAuthLoginRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiMerchantAuthLoginRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiMerchantAuthLoginRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiMerchantAuthLoginRes) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeApiMerchantAuthLoginRes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeApiMerchantAuthLoginRes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeApiMerchantAuthLoginRes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeApiMerchantAuthLoginRes) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


