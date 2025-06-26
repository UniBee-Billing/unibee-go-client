# UnibeeApiMerchantProfileGetLicenseRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIRateLimit** | Pointer to **int32** | APIRateLimit | [optional] 
**CurrentMemberCount** | Pointer to **int32** | CurrentMemberCount | [optional] 
**License** | Pointer to [**UnibeeInternalLogicMiddlewareLicenseLicense**](UnibeeInternalLogicMiddlewareLicenseLicense.md) |  | [optional] 
**MemberLimit** | Pointer to **int32** | MemberLimit, -1&#x3D;Unlimited | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchant**](UnibeeApiBeanMerchant.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantProfileGetLicenseRes

`func NewUnibeeApiMerchantProfileGetLicenseRes() *UnibeeApiMerchantProfileGetLicenseRes`

NewUnibeeApiMerchantProfileGetLicenseRes instantiates a new UnibeeApiMerchantProfileGetLicenseRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProfileGetLicenseResWithDefaults

`func NewUnibeeApiMerchantProfileGetLicenseResWithDefaults() *UnibeeApiMerchantProfileGetLicenseRes`

NewUnibeeApiMerchantProfileGetLicenseResWithDefaults instantiates a new UnibeeApiMerchantProfileGetLicenseRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPIRateLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetAPIRateLimit() int32`

GetAPIRateLimit returns the APIRateLimit field if non-nil, zero value otherwise.

### GetAPIRateLimitOk

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetAPIRateLimitOk() (*int32, bool)`

GetAPIRateLimitOk returns a tuple with the APIRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIRateLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) SetAPIRateLimit(v int32)`

SetAPIRateLimit sets APIRateLimit field to given value.

### HasAPIRateLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) HasAPIRateLimit() bool`

HasAPIRateLimit returns a boolean if a field has been set.

### GetCurrentMemberCount

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetCurrentMemberCount() int32`

GetCurrentMemberCount returns the CurrentMemberCount field if non-nil, zero value otherwise.

### GetCurrentMemberCountOk

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetCurrentMemberCountOk() (*int32, bool)`

GetCurrentMemberCountOk returns a tuple with the CurrentMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMemberCount

`func (o *UnibeeApiMerchantProfileGetLicenseRes) SetCurrentMemberCount(v int32)`

SetCurrentMemberCount sets CurrentMemberCount field to given value.

### HasCurrentMemberCount

`func (o *UnibeeApiMerchantProfileGetLicenseRes) HasCurrentMemberCount() bool`

HasCurrentMemberCount returns a boolean if a field has been set.

### GetLicense

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetLicense() UnibeeInternalLogicMiddlewareLicenseLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetLicenseOk() (*UnibeeInternalLogicMiddlewareLicenseLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *UnibeeApiMerchantProfileGetLicenseRes) SetLicense(v UnibeeInternalLogicMiddlewareLicenseLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *UnibeeApiMerchantProfileGetLicenseRes) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMemberLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMemberLimit() int32`

GetMemberLimit returns the MemberLimit field if non-nil, zero value otherwise.

### GetMemberLimitOk

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMemberLimitOk() (*int32, bool)`

GetMemberLimitOk returns a tuple with the MemberLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) SetMemberLimit(v int32)`

SetMemberLimit sets MemberLimit field to given value.

### HasMemberLimit

`func (o *UnibeeApiMerchantProfileGetLicenseRes) HasMemberLimit() bool`

HasMemberLimit returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiMerchantProfileGetLicenseRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiMerchantProfileGetLicenseRes) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiMerchantProfileGetLicenseRes) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


