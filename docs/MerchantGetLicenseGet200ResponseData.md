# MerchantGetLicenseGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIRateLimit** | Pointer to **int32** | APIRateLimit | [optional] 
**CurrentMemberCount** | Pointer to **int32** | CurrentMemberCount | [optional] 
**License** | Pointer to [**UnibeeInternalLogicMiddlewareLicenseLicense**](UnibeeInternalLogicMiddlewareLicenseLicense.md) |  | [optional] 
**MemberLimit** | Pointer to **int32** | MemberLimit, -1&#x3D;Unlimited | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchant**](UnibeeApiBeanMerchant.md) |  | [optional] 

## Methods

### NewMerchantGetLicenseGet200ResponseData

`func NewMerchantGetLicenseGet200ResponseData() *MerchantGetLicenseGet200ResponseData`

NewMerchantGetLicenseGet200ResponseData instantiates a new MerchantGetLicenseGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGetLicenseGet200ResponseDataWithDefaults

`func NewMerchantGetLicenseGet200ResponseDataWithDefaults() *MerchantGetLicenseGet200ResponseData`

NewMerchantGetLicenseGet200ResponseDataWithDefaults instantiates a new MerchantGetLicenseGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPIRateLimit

`func (o *MerchantGetLicenseGet200ResponseData) GetAPIRateLimit() int32`

GetAPIRateLimit returns the APIRateLimit field if non-nil, zero value otherwise.

### GetAPIRateLimitOk

`func (o *MerchantGetLicenseGet200ResponseData) GetAPIRateLimitOk() (*int32, bool)`

GetAPIRateLimitOk returns a tuple with the APIRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIRateLimit

`func (o *MerchantGetLicenseGet200ResponseData) SetAPIRateLimit(v int32)`

SetAPIRateLimit sets APIRateLimit field to given value.

### HasAPIRateLimit

`func (o *MerchantGetLicenseGet200ResponseData) HasAPIRateLimit() bool`

HasAPIRateLimit returns a boolean if a field has been set.

### GetCurrentMemberCount

`func (o *MerchantGetLicenseGet200ResponseData) GetCurrentMemberCount() int32`

GetCurrentMemberCount returns the CurrentMemberCount field if non-nil, zero value otherwise.

### GetCurrentMemberCountOk

`func (o *MerchantGetLicenseGet200ResponseData) GetCurrentMemberCountOk() (*int32, bool)`

GetCurrentMemberCountOk returns a tuple with the CurrentMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberCount

`func (o *MerchantGetLicenseGet200ResponseData) SetCurrentMemberCount(v int32)`

SetCurrentMemberCount sets CurrentMemberCount field to given value.

### HasCurrentMemberCount

`func (o *MerchantGetLicenseGet200ResponseData) HasCurrentMemberCount() bool`

HasCurrentMemberCount returns a boolean if a field has been set.

### GetLicense

`func (o *MerchantGetLicenseGet200ResponseData) GetLicense() UnibeeInternalLogicMiddlewareLicenseLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MerchantGetLicenseGet200ResponseData) GetLicenseOk() (*UnibeeInternalLogicMiddlewareLicenseLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MerchantGetLicenseGet200ResponseData) SetLicense(v UnibeeInternalLogicMiddlewareLicenseLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *MerchantGetLicenseGet200ResponseData) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMemberLimit

`func (o *MerchantGetLicenseGet200ResponseData) GetMemberLimit() int32`

GetMemberLimit returns the MemberLimit field if non-nil, zero value otherwise.

### GetMemberLimitOk

`func (o *MerchantGetLicenseGet200ResponseData) GetMemberLimitOk() (*int32, bool)`

GetMemberLimitOk returns a tuple with the MemberLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLimit

`func (o *MerchantGetLicenseGet200ResponseData) SetMemberLimit(v int32)`

SetMemberLimit sets MemberLimit field to given value.

### HasMemberLimit

`func (o *MerchantGetLicenseGet200ResponseData) HasMemberLimit() bool`

HasMemberLimit returns a boolean if a field has been set.

### GetMerchant

`func (o *MerchantGetLicenseGet200ResponseData) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MerchantGetLicenseGet200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MerchantGetLicenseGet200ResponseData) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *MerchantGetLicenseGet200ResponseData) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


