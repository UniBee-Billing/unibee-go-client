# UnibeeApiMerchantAuthPasswordSetupOtpRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**Token** | Pointer to **string** | Access token of admin portal | [optional] 

## Methods

### NewUnibeeApiMerchantAuthPasswordSetupOtpRes

`func NewUnibeeApiMerchantAuthPasswordSetupOtpRes() *UnibeeApiMerchantAuthPasswordSetupOtpRes`

NewUnibeeApiMerchantAuthPasswordSetupOtpRes instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantAuthPasswordSetupOtpResWithDefaults

`func NewUnibeeApiMerchantAuthPasswordSetupOtpResWithDefaults() *UnibeeApiMerchantAuthPasswordSetupOtpRes`

NewUnibeeApiMerchantAuthPasswordSetupOtpResWithDefaults instantiates a new UnibeeApiMerchantAuthPasswordSetupOtpRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMember

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeApiMerchantAuthPasswordSetupOtpRes) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


