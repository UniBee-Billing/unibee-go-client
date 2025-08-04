# UnibeeApiSystemInformationGetRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildVersion** | Pointer to **string** | System Build Version | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**Gateway** | Pointer to [**[]UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) | Support Gateway List | [optional] 
**IsProd** | Pointer to **bool** | Check System Env Is Prod, true|false | [optional] 
**Mode** | Pointer to **string** | System Mode | [optional] 
**Oauth** | Pointer to [**UnibeeApiSystemInformationOAuthConfig**](UnibeeApiSystemInformationOAuthConfig.md) |  | [optional] 
**SupportCountryCode** | Pointer to **[]string** | Support Country Code List (ISO 3166-1 alpha-2) | [optional] 
**SupportCurrency** | Pointer to [**[]UnibeeApiBeanCurrency**](UnibeeApiBeanCurrency.md) | Support Currency List | [optional] 
**SupportLanguage** | Pointer to [**[]UnibeeApiSystemInformationLanguage**](UnibeeApiSystemInformationLanguage.md) | Support Language List | [optional] 
**SupportTimeZone** | Pointer to **[]string** | Support TimeZone List | [optional] 

## Methods

### NewUnibeeApiSystemInformationGetRes

`func NewUnibeeApiSystemInformationGetRes() *UnibeeApiSystemInformationGetRes`

NewUnibeeApiSystemInformationGetRes instantiates a new UnibeeApiSystemInformationGetRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInformationGetResWithDefaults

`func NewUnibeeApiSystemInformationGetResWithDefaults() *UnibeeApiSystemInformationGetRes`

NewUnibeeApiSystemInformationGetResWithDefaults instantiates a new UnibeeApiSystemInformationGetRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildVersion

`func (o *UnibeeApiSystemInformationGetRes) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *UnibeeApiSystemInformationGetRes) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *UnibeeApiSystemInformationGetRes) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *UnibeeApiSystemInformationGetRes) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetEnv

`func (o *UnibeeApiSystemInformationGetRes) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *UnibeeApiSystemInformationGetRes) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *UnibeeApiSystemInformationGetRes) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *UnibeeApiSystemInformationGetRes) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiSystemInformationGetRes) GetGateway() []UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiSystemInformationGetRes) GetGatewayOk() (*[]UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiSystemInformationGetRes) SetGateway(v []UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiSystemInformationGetRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIsProd

`func (o *UnibeeApiSystemInformationGetRes) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *UnibeeApiSystemInformationGetRes) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *UnibeeApiSystemInformationGetRes) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *UnibeeApiSystemInformationGetRes) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetMode

`func (o *UnibeeApiSystemInformationGetRes) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UnibeeApiSystemInformationGetRes) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UnibeeApiSystemInformationGetRes) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UnibeeApiSystemInformationGetRes) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetOauth

`func (o *UnibeeApiSystemInformationGetRes) GetOauth() UnibeeApiSystemInformationOAuthConfig`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *UnibeeApiSystemInformationGetRes) GetOauthOk() (*UnibeeApiSystemInformationOAuthConfig, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *UnibeeApiSystemInformationGetRes) SetOauth(v UnibeeApiSystemInformationOAuthConfig)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *UnibeeApiSystemInformationGetRes) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetSupportCountryCode

`func (o *UnibeeApiSystemInformationGetRes) GetSupportCountryCode() []string`

GetSupportCountryCode returns the SupportCountryCode field if non-nil, zero value otherwise.

### GetSupportCountryCodeOk

`func (o *UnibeeApiSystemInformationGetRes) GetSupportCountryCodeOk() (*[]string, bool)`

GetSupportCountryCodeOk returns a tuple with the SupportCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCountryCode

`func (o *UnibeeApiSystemInformationGetRes) SetSupportCountryCode(v []string)`

SetSupportCountryCode sets SupportCountryCode field to given value.

### HasSupportCountryCode

`func (o *UnibeeApiSystemInformationGetRes) HasSupportCountryCode() bool`

HasSupportCountryCode returns a boolean if a field has been set.

### GetSupportCurrency

`func (o *UnibeeApiSystemInformationGetRes) GetSupportCurrency() []UnibeeApiBeanCurrency`

GetSupportCurrency returns the SupportCurrency field if non-nil, zero value otherwise.

### GetSupportCurrencyOk

`func (o *UnibeeApiSystemInformationGetRes) GetSupportCurrencyOk() (*[]UnibeeApiBeanCurrency, bool)`

GetSupportCurrencyOk returns a tuple with the SupportCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCurrency

`func (o *UnibeeApiSystemInformationGetRes) SetSupportCurrency(v []UnibeeApiBeanCurrency)`

SetSupportCurrency sets SupportCurrency field to given value.

### HasSupportCurrency

`func (o *UnibeeApiSystemInformationGetRes) HasSupportCurrency() bool`

HasSupportCurrency returns a boolean if a field has been set.

### GetSupportLanguage

`func (o *UnibeeApiSystemInformationGetRes) GetSupportLanguage() []UnibeeApiSystemInformationLanguage`

GetSupportLanguage returns the SupportLanguage field if non-nil, zero value otherwise.

### GetSupportLanguageOk

`func (o *UnibeeApiSystemInformationGetRes) GetSupportLanguageOk() (*[]UnibeeApiSystemInformationLanguage, bool)`

GetSupportLanguageOk returns a tuple with the SupportLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLanguage

`func (o *UnibeeApiSystemInformationGetRes) SetSupportLanguage(v []UnibeeApiSystemInformationLanguage)`

SetSupportLanguage sets SupportLanguage field to given value.

### HasSupportLanguage

`func (o *UnibeeApiSystemInformationGetRes) HasSupportLanguage() bool`

HasSupportLanguage returns a boolean if a field has been set.

### GetSupportTimeZone

`func (o *UnibeeApiSystemInformationGetRes) GetSupportTimeZone() []string`

GetSupportTimeZone returns the SupportTimeZone field if non-nil, zero value otherwise.

### GetSupportTimeZoneOk

`func (o *UnibeeApiSystemInformationGetRes) GetSupportTimeZoneOk() (*[]string, bool)`

GetSupportTimeZoneOk returns a tuple with the SupportTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTimeZone

`func (o *UnibeeApiSystemInformationGetRes) SetSupportTimeZone(v []string)`

SetSupportTimeZone sets SupportTimeZone field to given value.

### HasSupportTimeZone

`func (o *UnibeeApiSystemInformationGetRes) HasSupportTimeZone() bool`

HasSupportTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


