# UnibeeApiMerchantMemberGetTotpKeyRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpKey** | Pointer to **string** | TotpKey, used on next confirm step | [optional] 
**TotpResumeCode** | Pointer to **string** | TotpResumeCode | [optional] 
**TotpType** | Pointer to **int32** | 1-General, Google Authenticator | 2-Microsoft Authenticator | 3-Authy | 4-1Password | 5-LastPass | 6-FreeOTP | 7-Other TOTP | [optional] 
**TotpUrl** | Pointer to **string** | TotpUrl， Used to generate QR Image | [optional] 

## Methods

### NewUnibeeApiMerchantMemberGetTotpKeyRes

`func NewUnibeeApiMerchantMemberGetTotpKeyRes() *UnibeeApiMerchantMemberGetTotpKeyRes`

NewUnibeeApiMerchantMemberGetTotpKeyRes instantiates a new UnibeeApiMerchantMemberGetTotpKeyRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberGetTotpKeyResWithDefaults

`func NewUnibeeApiMerchantMemberGetTotpKeyResWithDefaults() *UnibeeApiMerchantMemberGetTotpKeyRes`

NewUnibeeApiMerchantMemberGetTotpKeyResWithDefaults instantiates a new UnibeeApiMerchantMemberGetTotpKeyRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpKey

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpKey() string`

GetTotpKey returns the TotpKey field if non-nil, zero value otherwise.

### GetTotpKeyOk

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpKeyOk() (*string, bool)`

GetTotpKeyOk returns a tuple with the TotpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKey

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) SetTotpKey(v string)`

SetTotpKey sets TotpKey field to given value.

### HasTotpKey

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) HasTotpKey() bool`

HasTotpKey returns a boolean if a field has been set.

### GetTotpResumeCode

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpResumeCode() string`

GetTotpResumeCode returns the TotpResumeCode field if non-nil, zero value otherwise.

### GetTotpResumeCodeOk

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpResumeCodeOk() (*string, bool)`

GetTotpResumeCodeOk returns a tuple with the TotpResumeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpResumeCode

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) SetTotpResumeCode(v string)`

SetTotpResumeCode sets TotpResumeCode field to given value.

### HasTotpResumeCode

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) HasTotpResumeCode() bool`

HasTotpResumeCode returns a boolean if a field has been set.

### GetTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpType() int32`

GetTotpType returns the TotpType field if non-nil, zero value otherwise.

### GetTotpTypeOk

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpTypeOk() (*int32, bool)`

GetTotpTypeOk returns a tuple with the TotpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) SetTotpType(v int32)`

SetTotpType sets TotpType field to given value.

### HasTotpType

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) HasTotpType() bool`

HasTotpType returns a boolean if a field has been set.

### GetTotpUrl

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpUrl() string`

GetTotpUrl returns the TotpUrl field if non-nil, zero value otherwise.

### GetTotpUrlOk

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) GetTotpUrlOk() (*string, bool)`

GetTotpUrlOk returns a tuple with the TotpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUrl

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) SetTotpUrl(v string)`

SetTotpUrl sets TotpUrl field to given value.

### HasTotpUrl

`func (o *UnibeeApiMerchantMemberGetTotpKeyRes) HasTotpUrl() bool`

HasTotpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


