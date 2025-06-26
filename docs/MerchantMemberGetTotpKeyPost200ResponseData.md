# MerchantMemberGetTotpKeyPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotpKey** | Pointer to **string** | TotpKey, used on next confirm step | [optional] 
**TotpResumeCode** | Pointer to **string** | TotpResumeCode | [optional] 
**TotpType** | Pointer to **int32** | 1-General, Google Authenticator | 2-Microsoft Authenticator | 3-Authy | 4-1Password | 5-LastPass | 6-FreeOTP | 7-Other TOTP | [optional] 
**TotpUrl** | Pointer to **string** | TotpUrl， Used to generate QR Image | [optional] 

## Methods

### NewMerchantMemberGetTotpKeyPost200ResponseData

`func NewMerchantMemberGetTotpKeyPost200ResponseData() *MerchantMemberGetTotpKeyPost200ResponseData`

NewMerchantMemberGetTotpKeyPost200ResponseData instantiates a new MerchantMemberGetTotpKeyPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantMemberGetTotpKeyPost200ResponseDataWithDefaults

`func NewMerchantMemberGetTotpKeyPost200ResponseDataWithDefaults() *MerchantMemberGetTotpKeyPost200ResponseData`

NewMerchantMemberGetTotpKeyPost200ResponseDataWithDefaults instantiates a new MerchantMemberGetTotpKeyPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotpKey

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpKey() string`

GetTotpKey returns the TotpKey field if non-nil, zero value otherwise.

### GetTotpKeyOk

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpKeyOk() (*string, bool)`

GetTotpKeyOk returns a tuple with the TotpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpKey

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) SetTotpKey(v string)`

SetTotpKey sets TotpKey field to given value.

### HasTotpKey

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) HasTotpKey() bool`

HasTotpKey returns a boolean if a field has been set.

### GetTotpResumeCode

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpResumeCode() string`

GetTotpResumeCode returns the TotpResumeCode field if non-nil, zero value otherwise.

### GetTotpResumeCodeOk

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpResumeCodeOk() (*string, bool)`

GetTotpResumeCodeOk returns a tuple with the TotpResumeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpResumeCode

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) SetTotpResumeCode(v string)`

SetTotpResumeCode sets TotpResumeCode field to given value.

### HasTotpResumeCode

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) HasTotpResumeCode() bool`

HasTotpResumeCode returns a boolean if a field has been set.

### GetTotpType

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpType() int32`

GetTotpType returns the TotpType field if non-nil, zero value otherwise.

### GetTotpTypeOk

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpTypeOk() (*int32, bool)`

GetTotpTypeOk returns a tuple with the TotpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpType

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) SetTotpType(v int32)`

SetTotpType sets TotpType field to given value.

### HasTotpType

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) HasTotpType() bool`

HasTotpType returns a boolean if a field has been set.

### GetTotpUrl

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpUrl() string`

GetTotpUrl returns the TotpUrl field if non-nil, zero value otherwise.

### GetTotpUrlOk

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) GetTotpUrlOk() (*string, bool)`

GetTotpUrlOk returns a tuple with the TotpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpUrl

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) SetTotpUrl(v string)`

SetTotpUrl sets TotpUrl field to given value.

### HasTotpUrl

`func (o *MerchantMemberGetTotpKeyPost200ResponseData) HasTotpUrl() bool`

HasTotpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


