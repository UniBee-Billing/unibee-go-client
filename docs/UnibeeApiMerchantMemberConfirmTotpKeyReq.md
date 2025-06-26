# UnibeeApiMerchantMemberConfirmTotpKeyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpCode** | Pointer to **string** | The totp code | [optional] 
**TotpKey** | Pointer to **string** | TotpKey | [optional] 
**TotpType** | Pointer to **int32** | 1-General, Google Authenticator | 2-Microsoft Authenticator | 3-Authy | 4-1Password | 5-LastPass | 6-FreeOTP | 7-Other TOTP | [optional] 

## Methods

### NewUnibeeApiMerchantMemberConfirmTotpKeyReq

`func NewUnibeeApiMerchantMemberConfirmTotpKeyReq() *UnibeeApiMerchantMemberConfirmTotpKeyReq`

NewUnibeeApiMerchantMemberConfirmTotpKeyReq instantiates a new UnibeeApiMerchantMemberConfirmTotpKeyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberConfirmTotpKeyReqWithDefaults

`func NewUnibeeApiMerchantMemberConfirmTotpKeyReqWithDefaults() *UnibeeApiMerchantMemberConfirmTotpKeyReq`

NewUnibeeApiMerchantMemberConfirmTotpKeyReqWithDefaults instantiates a new UnibeeApiMerchantMemberConfirmTotpKeyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpCode

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.

### GetTotpKey

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpKey() string`

GetTotpKey returns the TotpKey field if non-nil, zero value otherwise.

### GetTotpKeyOk

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpKeyOk() (*string, bool)`

GetTotpKeyOk returns a tuple with the TotpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKey

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) SetTotpKey(v string)`

SetTotpKey sets TotpKey field to given value.

### HasTotpKey

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) HasTotpKey() bool`

HasTotpKey returns a boolean if a field has been set.

### GetTotpType

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpType() int32`

GetTotpType returns the TotpType field if non-nil, zero value otherwise.

### GetTotpTypeOk

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) GetTotpTypeOk() (*int32, bool)`

GetTotpTypeOk returns a tuple with the TotpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpType

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) SetTotpType(v int32)`

SetTotpType sets TotpType field to given value.

### HasTotpType

`func (o *UnibeeApiMerchantMemberConfirmTotpKeyReq) HasTotpType() bool`

HasTotpType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


