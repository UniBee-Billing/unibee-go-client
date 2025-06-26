# UnibeeApiMerchantMemberGetTotpKeyReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpType** | Pointer to **int32** | 1-General, Google Authenticator | 2-Microsoft Authenticator | 3-Authy | 4-1Password | 5-LastPass | 6-FreeOTP | 7-Other TOTP | [optional] 

## Methods

### NewUnibeeApiMerchantMemberGetTotpKeyReq

`func NewUnibeeApiMerchantMemberGetTotpKeyReq() *UnibeeApiMerchantMemberGetTotpKeyReq`

NewUnibeeApiMerchantMemberGetTotpKeyReq instantiates a new UnibeeApiMerchantMemberGetTotpKeyReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberGetTotpKeyReqWithDefaults

`func NewUnibeeApiMerchantMemberGetTotpKeyReqWithDefaults() *UnibeeApiMerchantMemberGetTotpKeyReq`

NewUnibeeApiMerchantMemberGetTotpKeyReqWithDefaults instantiates a new UnibeeApiMerchantMemberGetTotpKeyReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyReq) GetTotpType() int32`

GetTotpType returns the TotpType field if non-nil, zero value otherwise.

### GetTotpTypeOk

`func (o *UnibeeApiMerchantMemberGetTotpKeyReq) GetTotpTypeOk() (*int32, bool)`

GetTotpTypeOk returns a tuple with the TotpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyReq) SetTotpType(v int32)`

SetTotpType sets TotpType field to given value.

### HasTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyReq) HasTotpType() bool`

HasTotpType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


