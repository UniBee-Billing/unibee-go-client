# MerchantGetGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]UnibeeApiBeanCurrency**](UnibeeApiBeanCurrency.md) | Currency List | [optional] 
**MemberRoles** | Pointer to [**[]UnibeeApiBeanMerchantRole**](UnibeeApiBeanMerchantRole.md) | The member role list&#39; | [optional] 
**TimeZone** | Pointer to **[]string** | TimeZone List | [optional] 
**AnalyticsHost** | Pointer to **string** | Analytics Host | [optional] 
**DefaultCurrency** | Pointer to **string** | Default Currency | [optional] 
**EmailSender** | Pointer to [**UnibeeApiBeanSender**](UnibeeApiBeanSender.md) |  | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**ExchangeRateApiKey** | Pointer to **string** | ExchangeRateApiKey | [optional] 
**Gateways** | Pointer to [**[]UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) | Gateway List | [optional] 
**GlobalTOPTEnabled** | Pointer to **bool** | GlobalTOPTEnabled | [optional] 
**IsOwner** | Pointer to **bool** | Check Member is Owner | [optional] 
**IsProd** | Pointer to **bool** | Check System Env Is Prod, true|false | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchant**](UnibeeApiBeanMerchant.md) |  | [optional] 
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**MultiCurrencyConfigs** | Pointer to [**[]UnibeeApiBeanMerchantMultiCurrencyConfig**](UnibeeApiBeanMerchantMultiCurrencyConfig.md) | Merchant&#39;s MultiCurrency Configs | [optional] 
**OpenApiHost** | Pointer to **string** | OpenApi Host | [optional] 
**OpenApiKey** | Pointer to **string** | OpenAPIKey | [optional] 
**QuickBooksCompanyName** | Pointer to **string** | QuickBooksCompanyName | [optional] 
**QuickBooksLastSyncError** | Pointer to **string** | QuickBooksLastSyncError | [optional] 
**QuickBooksLastSynchronized** | Pointer to **string** | QuickBooksLastSynchronized | [optional] 
**SegmentServerSideKey** | Pointer to **string** | SegmentServerSideKey | [optional] 
**SegmentUserPortalKey** | Pointer to **string** | SegmentUserPortalKey | [optional] 
**SendGridKey** | Pointer to **string** | SendGridKey | [optional] 
**VatSenseKey** | Pointer to **string** | VatSenseKey | [optional] 

## Methods

### NewMerchantGetGet200ResponseData

`func NewMerchantGetGet200ResponseData() *MerchantGetGet200ResponseData`

NewMerchantGetGet200ResponseData instantiates a new MerchantGetGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantGetGet200ResponseDataWithDefaults

`func NewMerchantGetGet200ResponseDataWithDefaults() *MerchantGetGet200ResponseData`

NewMerchantGetGet200ResponseDataWithDefaults instantiates a new MerchantGetGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *MerchantGetGet200ResponseData) GetCurrency() []UnibeeApiBeanCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MerchantGetGet200ResponseData) GetCurrencyOk() (*[]UnibeeApiBeanCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MerchantGetGet200ResponseData) SetCurrency(v []UnibeeApiBeanCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MerchantGetGet200ResponseData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMemberRoles

`func (o *MerchantGetGet200ResponseData) GetMemberRoles() []UnibeeApiBeanMerchantRole`

GetMemberRoles returns the MemberRoles field if non-nil, zero value otherwise.

### GetMemberRolesOk

`func (o *MerchantGetGet200ResponseData) GetMemberRolesOk() (*[]UnibeeApiBeanMerchantRole, bool)`

GetMemberRolesOk returns a tuple with the MemberRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRoles

`func (o *MerchantGetGet200ResponseData) SetMemberRoles(v []UnibeeApiBeanMerchantRole)`

SetMemberRoles sets MemberRoles field to given value.

### HasMemberRoles

`func (o *MerchantGetGet200ResponseData) HasMemberRoles() bool`

HasMemberRoles returns a boolean if a field has been set.

### GetTimeZone

`func (o *MerchantGetGet200ResponseData) GetTimeZone() []string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MerchantGetGet200ResponseData) GetTimeZoneOk() (*[]string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MerchantGetGet200ResponseData) SetTimeZone(v []string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MerchantGetGet200ResponseData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetAnalyticsHost

`func (o *MerchantGetGet200ResponseData) GetAnalyticsHost() string`

GetAnalyticsHost returns the AnalyticsHost field if non-nil, zero value otherwise.

### GetAnalyticsHostOk

`func (o *MerchantGetGet200ResponseData) GetAnalyticsHostOk() (*string, bool)`

GetAnalyticsHostOk returns a tuple with the AnalyticsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsHost

`func (o *MerchantGetGet200ResponseData) SetAnalyticsHost(v string)`

SetAnalyticsHost sets AnalyticsHost field to given value.

### HasAnalyticsHost

`func (o *MerchantGetGet200ResponseData) HasAnalyticsHost() bool`

HasAnalyticsHost returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *MerchantGetGet200ResponseData) GetDefaultCurrency() string`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *MerchantGetGet200ResponseData) GetDefaultCurrencyOk() (*string, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *MerchantGetGet200ResponseData) SetDefaultCurrency(v string)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *MerchantGetGet200ResponseData) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetEmailSender

`func (o *MerchantGetGet200ResponseData) GetEmailSender() UnibeeApiBeanSender`

GetEmailSender returns the EmailSender field if non-nil, zero value otherwise.

### GetEmailSenderOk

`func (o *MerchantGetGet200ResponseData) GetEmailSenderOk() (*UnibeeApiBeanSender, bool)`

GetEmailSenderOk returns a tuple with the EmailSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSender

`func (o *MerchantGetGet200ResponseData) SetEmailSender(v UnibeeApiBeanSender)`

SetEmailSender sets EmailSender field to given value.

### HasEmailSender

`func (o *MerchantGetGet200ResponseData) HasEmailSender() bool`

HasEmailSender returns a boolean if a field has been set.

### GetEnv

`func (o *MerchantGetGet200ResponseData) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *MerchantGetGet200ResponseData) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *MerchantGetGet200ResponseData) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *MerchantGetGet200ResponseData) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExchangeRateApiKey

`func (o *MerchantGetGet200ResponseData) GetExchangeRateApiKey() string`

GetExchangeRateApiKey returns the ExchangeRateApiKey field if non-nil, zero value otherwise.

### GetExchangeRateApiKeyOk

`func (o *MerchantGetGet200ResponseData) GetExchangeRateApiKeyOk() (*string, bool)`

GetExchangeRateApiKeyOk returns a tuple with the ExchangeRateApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateApiKey

`func (o *MerchantGetGet200ResponseData) SetExchangeRateApiKey(v string)`

SetExchangeRateApiKey sets ExchangeRateApiKey field to given value.

### HasExchangeRateApiKey

`func (o *MerchantGetGet200ResponseData) HasExchangeRateApiKey() bool`

HasExchangeRateApiKey returns a boolean if a field has been set.

### GetGateways

`func (o *MerchantGetGet200ResponseData) GetGateways() []UnibeeApiBeanDetailGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *MerchantGetGet200ResponseData) GetGatewaysOk() (*[]UnibeeApiBeanDetailGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *MerchantGetGet200ResponseData) SetGateways(v []UnibeeApiBeanDetailGateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *MerchantGetGet200ResponseData) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetGlobalTOPTEnabled

`func (o *MerchantGetGet200ResponseData) GetGlobalTOPTEnabled() bool`

GetGlobalTOPTEnabled returns the GlobalTOPTEnabled field if non-nil, zero value otherwise.

### GetGlobalTOPTEnabledOk

`func (o *MerchantGetGet200ResponseData) GetGlobalTOPTEnabledOk() (*bool, bool)`

GetGlobalTOPTEnabledOk returns a tuple with the GlobalTOPTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTOPTEnabled

`func (o *MerchantGetGet200ResponseData) SetGlobalTOPTEnabled(v bool)`

SetGlobalTOPTEnabled sets GlobalTOPTEnabled field to given value.

### HasGlobalTOPTEnabled

`func (o *MerchantGetGet200ResponseData) HasGlobalTOPTEnabled() bool`

HasGlobalTOPTEnabled returns a boolean if a field has been set.

### GetIsOwner

`func (o *MerchantGetGet200ResponseData) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *MerchantGetGet200ResponseData) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *MerchantGetGet200ResponseData) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *MerchantGetGet200ResponseData) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsProd

`func (o *MerchantGetGet200ResponseData) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *MerchantGetGet200ResponseData) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *MerchantGetGet200ResponseData) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *MerchantGetGet200ResponseData) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetMerchant

`func (o *MerchantGetGet200ResponseData) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MerchantGetGet200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MerchantGetGet200ResponseData) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *MerchantGetGet200ResponseData) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMerchantMember

`func (o *MerchantGetGet200ResponseData) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *MerchantGetGet200ResponseData) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *MerchantGetGet200ResponseData) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *MerchantGetGet200ResponseData) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetMultiCurrencyConfigs

`func (o *MerchantGetGet200ResponseData) GetMultiCurrencyConfigs() []UnibeeApiBeanMerchantMultiCurrencyConfig`

GetMultiCurrencyConfigs returns the MultiCurrencyConfigs field if non-nil, zero value otherwise.

### GetMultiCurrencyConfigsOk

`func (o *MerchantGetGet200ResponseData) GetMultiCurrencyConfigsOk() (*[]UnibeeApiBeanMerchantMultiCurrencyConfig, bool)`

GetMultiCurrencyConfigsOk returns a tuple with the MultiCurrencyConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCurrencyConfigs

`func (o *MerchantGetGet200ResponseData) SetMultiCurrencyConfigs(v []UnibeeApiBeanMerchantMultiCurrencyConfig)`

SetMultiCurrencyConfigs sets MultiCurrencyConfigs field to given value.

### HasMultiCurrencyConfigs

`func (o *MerchantGetGet200ResponseData) HasMultiCurrencyConfigs() bool`

HasMultiCurrencyConfigs returns a boolean if a field has been set.

### GetOpenApiHost

`func (o *MerchantGetGet200ResponseData) GetOpenApiHost() string`

GetOpenApiHost returns the OpenApiHost field if non-nil, zero value otherwise.

### GetOpenApiHostOk

`func (o *MerchantGetGet200ResponseData) GetOpenApiHostOk() (*string, bool)`

GetOpenApiHostOk returns a tuple with the OpenApiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiHost

`func (o *MerchantGetGet200ResponseData) SetOpenApiHost(v string)`

SetOpenApiHost sets OpenApiHost field to given value.

### HasOpenApiHost

`func (o *MerchantGetGet200ResponseData) HasOpenApiHost() bool`

HasOpenApiHost returns a boolean if a field has been set.

### GetOpenApiKey

`func (o *MerchantGetGet200ResponseData) GetOpenApiKey() string`

GetOpenApiKey returns the OpenApiKey field if non-nil, zero value otherwise.

### GetOpenApiKeyOk

`func (o *MerchantGetGet200ResponseData) GetOpenApiKeyOk() (*string, bool)`

GetOpenApiKeyOk returns a tuple with the OpenApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiKey

`func (o *MerchantGetGet200ResponseData) SetOpenApiKey(v string)`

SetOpenApiKey sets OpenApiKey field to given value.

### HasOpenApiKey

`func (o *MerchantGetGet200ResponseData) HasOpenApiKey() bool`

HasOpenApiKey returns a boolean if a field has been set.

### GetQuickBooksCompanyName

`func (o *MerchantGetGet200ResponseData) GetQuickBooksCompanyName() string`

GetQuickBooksCompanyName returns the QuickBooksCompanyName field if non-nil, zero value otherwise.

### GetQuickBooksCompanyNameOk

`func (o *MerchantGetGet200ResponseData) GetQuickBooksCompanyNameOk() (*string, bool)`

GetQuickBooksCompanyNameOk returns a tuple with the QuickBooksCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksCompanyName

`func (o *MerchantGetGet200ResponseData) SetQuickBooksCompanyName(v string)`

SetQuickBooksCompanyName sets QuickBooksCompanyName field to given value.

### HasQuickBooksCompanyName

`func (o *MerchantGetGet200ResponseData) HasQuickBooksCompanyName() bool`

HasQuickBooksCompanyName returns a boolean if a field has been set.

### GetQuickBooksLastSyncError

`func (o *MerchantGetGet200ResponseData) GetQuickBooksLastSyncError() string`

GetQuickBooksLastSyncError returns the QuickBooksLastSyncError field if non-nil, zero value otherwise.

### GetQuickBooksLastSyncErrorOk

`func (o *MerchantGetGet200ResponseData) GetQuickBooksLastSyncErrorOk() (*string, bool)`

GetQuickBooksLastSyncErrorOk returns a tuple with the QuickBooksLastSyncError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksLastSyncError

`func (o *MerchantGetGet200ResponseData) SetQuickBooksLastSyncError(v string)`

SetQuickBooksLastSyncError sets QuickBooksLastSyncError field to given value.

### HasQuickBooksLastSyncError

`func (o *MerchantGetGet200ResponseData) HasQuickBooksLastSyncError() bool`

HasQuickBooksLastSyncError returns a boolean if a field has been set.

### GetQuickBooksLastSynchronized

`func (o *MerchantGetGet200ResponseData) GetQuickBooksLastSynchronized() string`

GetQuickBooksLastSynchronized returns the QuickBooksLastSynchronized field if non-nil, zero value otherwise.

### GetQuickBooksLastSynchronizedOk

`func (o *MerchantGetGet200ResponseData) GetQuickBooksLastSynchronizedOk() (*string, bool)`

GetQuickBooksLastSynchronizedOk returns a tuple with the QuickBooksLastSynchronized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksLastSynchronized

`func (o *MerchantGetGet200ResponseData) SetQuickBooksLastSynchronized(v string)`

SetQuickBooksLastSynchronized sets QuickBooksLastSynchronized field to given value.

### HasQuickBooksLastSynchronized

`func (o *MerchantGetGet200ResponseData) HasQuickBooksLastSynchronized() bool`

HasQuickBooksLastSynchronized returns a boolean if a field has been set.

### GetSegmentServerSideKey

`func (o *MerchantGetGet200ResponseData) GetSegmentServerSideKey() string`

GetSegmentServerSideKey returns the SegmentServerSideKey field if non-nil, zero value otherwise.

### GetSegmentServerSideKeyOk

`func (o *MerchantGetGet200ResponseData) GetSegmentServerSideKeyOk() (*string, bool)`

GetSegmentServerSideKeyOk returns a tuple with the SegmentServerSideKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentServerSideKey

`func (o *MerchantGetGet200ResponseData) SetSegmentServerSideKey(v string)`

SetSegmentServerSideKey sets SegmentServerSideKey field to given value.

### HasSegmentServerSideKey

`func (o *MerchantGetGet200ResponseData) HasSegmentServerSideKey() bool`

HasSegmentServerSideKey returns a boolean if a field has been set.

### GetSegmentUserPortalKey

`func (o *MerchantGetGet200ResponseData) GetSegmentUserPortalKey() string`

GetSegmentUserPortalKey returns the SegmentUserPortalKey field if non-nil, zero value otherwise.

### GetSegmentUserPortalKeyOk

`func (o *MerchantGetGet200ResponseData) GetSegmentUserPortalKeyOk() (*string, bool)`

GetSegmentUserPortalKeyOk returns a tuple with the SegmentUserPortalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentUserPortalKey

`func (o *MerchantGetGet200ResponseData) SetSegmentUserPortalKey(v string)`

SetSegmentUserPortalKey sets SegmentUserPortalKey field to given value.

### HasSegmentUserPortalKey

`func (o *MerchantGetGet200ResponseData) HasSegmentUserPortalKey() bool`

HasSegmentUserPortalKey returns a boolean if a field has been set.

### GetSendGridKey

`func (o *MerchantGetGet200ResponseData) GetSendGridKey() string`

GetSendGridKey returns the SendGridKey field if non-nil, zero value otherwise.

### GetSendGridKeyOk

`func (o *MerchantGetGet200ResponseData) GetSendGridKeyOk() (*string, bool)`

GetSendGridKeyOk returns a tuple with the SendGridKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendGridKey

`func (o *MerchantGetGet200ResponseData) SetSendGridKey(v string)`

SetSendGridKey sets SendGridKey field to given value.

### HasSendGridKey

`func (o *MerchantGetGet200ResponseData) HasSendGridKey() bool`

HasSendGridKey returns a boolean if a field has been set.

### GetVatSenseKey

`func (o *MerchantGetGet200ResponseData) GetVatSenseKey() string`

GetVatSenseKey returns the VatSenseKey field if non-nil, zero value otherwise.

### GetVatSenseKeyOk

`func (o *MerchantGetGet200ResponseData) GetVatSenseKeyOk() (*string, bool)`

GetVatSenseKeyOk returns a tuple with the VatSenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatSenseKey

`func (o *MerchantGetGet200ResponseData) SetVatSenseKey(v string)`

SetVatSenseKey sets VatSenseKey field to given value.

### HasVatSenseKey

`func (o *MerchantGetGet200ResponseData) HasVatSenseKey() bool`

HasVatSenseKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


