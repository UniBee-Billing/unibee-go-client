# UnibeeApiMerchantProfileGetRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]UnibeeApiBeanCurrency**](UnibeeApiBeanCurrency.md) | Currency List | [optional] 
**MemberRoles** | Pointer to [**[]UnibeeApiBeanMerchantRole**](UnibeeApiBeanMerchantRole.md) | The member role list&#39; | [optional] 
**TimeZone** | Pointer to **[]string** | TimeZone List | [optional] 
**AnalyticsHost** | Pointer to **string** | Analytics Host | [optional] 
**CloudFeatureAnalyticsEnabled** | Pointer to **bool** | Analytics Feature Enabled For Cloud Version | [optional] 
**DefaultCurrency** | Pointer to **string** | Default Currency | [optional] 
**EmailSender** | Pointer to [**UnibeeApiBeanSender**](UnibeeApiBeanSender.md) |  | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**ExchangeRateApiKey** | Pointer to **string** | ExchangeRateApiKey | [optional] 
**Gateways** | Pointer to [**[]UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) | Gateway List | [optional] 
**GlobalTOPTEnabled** | Pointer to **bool** | GlobalTOPTEnabled | [optional] 
**GlobalUSVATConfig** | Pointer to [**UnibeeApiBeanUSVATGlobalConfig**](UnibeeApiBeanUSVATGlobalConfig.md) |  | [optional] 
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
**TaxJarApiKeys** | Pointer to [**UnibeeApiBeanUSVATGatewayConnectionAPIKeys**](UnibeeApiBeanUSVATGatewayConnectionAPIKeys.md) |  | [optional] 
**VatSenseKey** | Pointer to **string** | VatSenseKey | [optional] 

## Methods

### NewUnibeeApiMerchantProfileGetRes

`func NewUnibeeApiMerchantProfileGetRes() *UnibeeApiMerchantProfileGetRes`

NewUnibeeApiMerchantProfileGetRes instantiates a new UnibeeApiMerchantProfileGetRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProfileGetResWithDefaults

`func NewUnibeeApiMerchantProfileGetResWithDefaults() *UnibeeApiMerchantProfileGetRes`

NewUnibeeApiMerchantProfileGetResWithDefaults instantiates a new UnibeeApiMerchantProfileGetRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantProfileGetRes) GetCurrency() []UnibeeApiBeanCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetCurrencyOk() (*[]UnibeeApiBeanCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantProfileGetRes) SetCurrency(v []UnibeeApiBeanCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantProfileGetRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMemberRoles

`func (o *UnibeeApiMerchantProfileGetRes) GetMemberRoles() []UnibeeApiBeanMerchantRole`

GetMemberRoles returns the MemberRoles field if non-nil, zero value otherwise.

### GetMemberRolesOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMemberRolesOk() (*[]UnibeeApiBeanMerchantRole, bool)`

GetMemberRolesOk returns a tuple with the MemberRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRoles

`func (o *UnibeeApiMerchantProfileGetRes) SetMemberRoles(v []UnibeeApiBeanMerchantRole)`

SetMemberRoles sets MemberRoles field to given value.

### HasMemberRoles

`func (o *UnibeeApiMerchantProfileGetRes) HasMemberRoles() bool`

HasMemberRoles returns a boolean if a field has been set.

### GetTimeZone

`func (o *UnibeeApiMerchantProfileGetRes) GetTimeZone() []string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UnibeeApiMerchantProfileGetRes) GetTimeZoneOk() (*[]string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UnibeeApiMerchantProfileGetRes) SetTimeZone(v []string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UnibeeApiMerchantProfileGetRes) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetAnalyticsHost

`func (o *UnibeeApiMerchantProfileGetRes) GetAnalyticsHost() string`

GetAnalyticsHost returns the AnalyticsHost field if non-nil, zero value otherwise.

### GetAnalyticsHostOk

`func (o *UnibeeApiMerchantProfileGetRes) GetAnalyticsHostOk() (*string, bool)`

GetAnalyticsHostOk returns a tuple with the AnalyticsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsHost

`func (o *UnibeeApiMerchantProfileGetRes) SetAnalyticsHost(v string)`

SetAnalyticsHost sets AnalyticsHost field to given value.

### HasAnalyticsHost

`func (o *UnibeeApiMerchantProfileGetRes) HasAnalyticsHost() bool`

HasAnalyticsHost returns a boolean if a field has been set.

### GetCloudFeatureAnalyticsEnabled

`func (o *UnibeeApiMerchantProfileGetRes) GetCloudFeatureAnalyticsEnabled() bool`

GetCloudFeatureAnalyticsEnabled returns the CloudFeatureAnalyticsEnabled field if non-nil, zero value otherwise.

### GetCloudFeatureAnalyticsEnabledOk

`func (o *UnibeeApiMerchantProfileGetRes) GetCloudFeatureAnalyticsEnabledOk() (*bool, bool)`

GetCloudFeatureAnalyticsEnabledOk returns a tuple with the CloudFeatureAnalyticsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFeatureAnalyticsEnabled

`func (o *UnibeeApiMerchantProfileGetRes) SetCloudFeatureAnalyticsEnabled(v bool)`

SetCloudFeatureAnalyticsEnabled sets CloudFeatureAnalyticsEnabled field to given value.

### HasCloudFeatureAnalyticsEnabled

`func (o *UnibeeApiMerchantProfileGetRes) HasCloudFeatureAnalyticsEnabled() bool`

HasCloudFeatureAnalyticsEnabled returns a boolean if a field has been set.

### GetDefaultCurrency

`func (o *UnibeeApiMerchantProfileGetRes) GetDefaultCurrency() string`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetDefaultCurrencyOk() (*string, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *UnibeeApiMerchantProfileGetRes) SetDefaultCurrency(v string)`

SetDefaultCurrency sets DefaultCurrency field to given value.

### HasDefaultCurrency

`func (o *UnibeeApiMerchantProfileGetRes) HasDefaultCurrency() bool`

HasDefaultCurrency returns a boolean if a field has been set.

### GetEmailSender

`func (o *UnibeeApiMerchantProfileGetRes) GetEmailSender() UnibeeApiBeanSender`

GetEmailSender returns the EmailSender field if non-nil, zero value otherwise.

### GetEmailSenderOk

`func (o *UnibeeApiMerchantProfileGetRes) GetEmailSenderOk() (*UnibeeApiBeanSender, bool)`

GetEmailSenderOk returns a tuple with the EmailSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSender

`func (o *UnibeeApiMerchantProfileGetRes) SetEmailSender(v UnibeeApiBeanSender)`

SetEmailSender sets EmailSender field to given value.

### HasEmailSender

`func (o *UnibeeApiMerchantProfileGetRes) HasEmailSender() bool`

HasEmailSender returns a boolean if a field has been set.

### GetEnv

`func (o *UnibeeApiMerchantProfileGetRes) GetEnv() string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *UnibeeApiMerchantProfileGetRes) GetEnvOk() (*string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *UnibeeApiMerchantProfileGetRes) SetEnv(v string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *UnibeeApiMerchantProfileGetRes) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExchangeRateApiKey

`func (o *UnibeeApiMerchantProfileGetRes) GetExchangeRateApiKey() string`

GetExchangeRateApiKey returns the ExchangeRateApiKey field if non-nil, zero value otherwise.

### GetExchangeRateApiKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetExchangeRateApiKeyOk() (*string, bool)`

GetExchangeRateApiKeyOk returns a tuple with the ExchangeRateApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRateApiKey

`func (o *UnibeeApiMerchantProfileGetRes) SetExchangeRateApiKey(v string)`

SetExchangeRateApiKey sets ExchangeRateApiKey field to given value.

### HasExchangeRateApiKey

`func (o *UnibeeApiMerchantProfileGetRes) HasExchangeRateApiKey() bool`

HasExchangeRateApiKey returns a boolean if a field has been set.

### GetGateways

`func (o *UnibeeApiMerchantProfileGetRes) GetGateways() []UnibeeApiBeanDetailGateway`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *UnibeeApiMerchantProfileGetRes) GetGatewaysOk() (*[]UnibeeApiBeanDetailGateway, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *UnibeeApiMerchantProfileGetRes) SetGateways(v []UnibeeApiBeanDetailGateway)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *UnibeeApiMerchantProfileGetRes) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetGlobalTOPTEnabled

`func (o *UnibeeApiMerchantProfileGetRes) GetGlobalTOPTEnabled() bool`

GetGlobalTOPTEnabled returns the GlobalTOPTEnabled field if non-nil, zero value otherwise.

### GetGlobalTOPTEnabledOk

`func (o *UnibeeApiMerchantProfileGetRes) GetGlobalTOPTEnabledOk() (*bool, bool)`

GetGlobalTOPTEnabledOk returns a tuple with the GlobalTOPTEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTOPTEnabled

`func (o *UnibeeApiMerchantProfileGetRes) SetGlobalTOPTEnabled(v bool)`

SetGlobalTOPTEnabled sets GlobalTOPTEnabled field to given value.

### HasGlobalTOPTEnabled

`func (o *UnibeeApiMerchantProfileGetRes) HasGlobalTOPTEnabled() bool`

HasGlobalTOPTEnabled returns a boolean if a field has been set.

### GetGlobalUSVATConfig

`func (o *UnibeeApiMerchantProfileGetRes) GetGlobalUSVATConfig() UnibeeApiBeanUSVATGlobalConfig`

GetGlobalUSVATConfig returns the GlobalUSVATConfig field if non-nil, zero value otherwise.

### GetGlobalUSVATConfigOk

`func (o *UnibeeApiMerchantProfileGetRes) GetGlobalUSVATConfigOk() (*UnibeeApiBeanUSVATGlobalConfig, bool)`

GetGlobalUSVATConfigOk returns a tuple with the GlobalUSVATConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalUSVATConfig

`func (o *UnibeeApiMerchantProfileGetRes) SetGlobalUSVATConfig(v UnibeeApiBeanUSVATGlobalConfig)`

SetGlobalUSVATConfig sets GlobalUSVATConfig field to given value.

### HasGlobalUSVATConfig

`func (o *UnibeeApiMerchantProfileGetRes) HasGlobalUSVATConfig() bool`

HasGlobalUSVATConfig returns a boolean if a field has been set.

### GetIsOwner

`func (o *UnibeeApiMerchantProfileGetRes) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *UnibeeApiMerchantProfileGetRes) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *UnibeeApiMerchantProfileGetRes) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *UnibeeApiMerchantProfileGetRes) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsProd

`func (o *UnibeeApiMerchantProfileGetRes) GetIsProd() bool`

GetIsProd returns the IsProd field if non-nil, zero value otherwise.

### GetIsProdOk

`func (o *UnibeeApiMerchantProfileGetRes) GetIsProdOk() (*bool, bool)`

GetIsProdOk returns a tuple with the IsProd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProd

`func (o *UnibeeApiMerchantProfileGetRes) SetIsProd(v bool)`

SetIsProd sets IsProd field to given value.

### HasIsProd

`func (o *UnibeeApiMerchantProfileGetRes) HasIsProd() bool`

HasIsProd returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiMerchantProfileGetRes) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiMerchantProfileGetRes) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMerchantMember

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *UnibeeApiMerchantProfileGetRes) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *UnibeeApiMerchantProfileGetRes) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetMultiCurrencyConfigs

`func (o *UnibeeApiMerchantProfileGetRes) GetMultiCurrencyConfigs() []UnibeeApiBeanMerchantMultiCurrencyConfig`

GetMultiCurrencyConfigs returns the MultiCurrencyConfigs field if non-nil, zero value otherwise.

### GetMultiCurrencyConfigsOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMultiCurrencyConfigsOk() (*[]UnibeeApiBeanMerchantMultiCurrencyConfig, bool)`

GetMultiCurrencyConfigsOk returns a tuple with the MultiCurrencyConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCurrencyConfigs

`func (o *UnibeeApiMerchantProfileGetRes) SetMultiCurrencyConfigs(v []UnibeeApiBeanMerchantMultiCurrencyConfig)`

SetMultiCurrencyConfigs sets MultiCurrencyConfigs field to given value.

### HasMultiCurrencyConfigs

`func (o *UnibeeApiMerchantProfileGetRes) HasMultiCurrencyConfigs() bool`

HasMultiCurrencyConfigs returns a boolean if a field has been set.

### GetOpenApiHost

`func (o *UnibeeApiMerchantProfileGetRes) GetOpenApiHost() string`

GetOpenApiHost returns the OpenApiHost field if non-nil, zero value otherwise.

### GetOpenApiHostOk

`func (o *UnibeeApiMerchantProfileGetRes) GetOpenApiHostOk() (*string, bool)`

GetOpenApiHostOk returns a tuple with the OpenApiHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiHost

`func (o *UnibeeApiMerchantProfileGetRes) SetOpenApiHost(v string)`

SetOpenApiHost sets OpenApiHost field to given value.

### HasOpenApiHost

`func (o *UnibeeApiMerchantProfileGetRes) HasOpenApiHost() bool`

HasOpenApiHost returns a boolean if a field has been set.

### GetOpenApiKey

`func (o *UnibeeApiMerchantProfileGetRes) GetOpenApiKey() string`

GetOpenApiKey returns the OpenApiKey field if non-nil, zero value otherwise.

### GetOpenApiKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetOpenApiKeyOk() (*string, bool)`

GetOpenApiKeyOk returns a tuple with the OpenApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiKey

`func (o *UnibeeApiMerchantProfileGetRes) SetOpenApiKey(v string)`

SetOpenApiKey sets OpenApiKey field to given value.

### HasOpenApiKey

`func (o *UnibeeApiMerchantProfileGetRes) HasOpenApiKey() bool`

HasOpenApiKey returns a boolean if a field has been set.

### GetQuickBooksCompanyName

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksCompanyName() string`

GetQuickBooksCompanyName returns the QuickBooksCompanyName field if non-nil, zero value otherwise.

### GetQuickBooksCompanyNameOk

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksCompanyNameOk() (*string, bool)`

GetQuickBooksCompanyNameOk returns a tuple with the QuickBooksCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksCompanyName

`func (o *UnibeeApiMerchantProfileGetRes) SetQuickBooksCompanyName(v string)`

SetQuickBooksCompanyName sets QuickBooksCompanyName field to given value.

### HasQuickBooksCompanyName

`func (o *UnibeeApiMerchantProfileGetRes) HasQuickBooksCompanyName() bool`

HasQuickBooksCompanyName returns a boolean if a field has been set.

### GetQuickBooksLastSyncError

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksLastSyncError() string`

GetQuickBooksLastSyncError returns the QuickBooksLastSyncError field if non-nil, zero value otherwise.

### GetQuickBooksLastSyncErrorOk

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksLastSyncErrorOk() (*string, bool)`

GetQuickBooksLastSyncErrorOk returns a tuple with the QuickBooksLastSyncError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksLastSyncError

`func (o *UnibeeApiMerchantProfileGetRes) SetQuickBooksLastSyncError(v string)`

SetQuickBooksLastSyncError sets QuickBooksLastSyncError field to given value.

### HasQuickBooksLastSyncError

`func (o *UnibeeApiMerchantProfileGetRes) HasQuickBooksLastSyncError() bool`

HasQuickBooksLastSyncError returns a boolean if a field has been set.

### GetQuickBooksLastSynchronized

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksLastSynchronized() string`

GetQuickBooksLastSynchronized returns the QuickBooksLastSynchronized field if non-nil, zero value otherwise.

### GetQuickBooksLastSynchronizedOk

`func (o *UnibeeApiMerchantProfileGetRes) GetQuickBooksLastSynchronizedOk() (*string, bool)`

GetQuickBooksLastSynchronizedOk returns a tuple with the QuickBooksLastSynchronized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickBooksLastSynchronized

`func (o *UnibeeApiMerchantProfileGetRes) SetQuickBooksLastSynchronized(v string)`

SetQuickBooksLastSynchronized sets QuickBooksLastSynchronized field to given value.

### HasQuickBooksLastSynchronized

`func (o *UnibeeApiMerchantProfileGetRes) HasQuickBooksLastSynchronized() bool`

HasQuickBooksLastSynchronized returns a boolean if a field has been set.

### GetSegmentServerSideKey

`func (o *UnibeeApiMerchantProfileGetRes) GetSegmentServerSideKey() string`

GetSegmentServerSideKey returns the SegmentServerSideKey field if non-nil, zero value otherwise.

### GetSegmentServerSideKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetSegmentServerSideKeyOk() (*string, bool)`

GetSegmentServerSideKeyOk returns a tuple with the SegmentServerSideKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentServerSideKey

`func (o *UnibeeApiMerchantProfileGetRes) SetSegmentServerSideKey(v string)`

SetSegmentServerSideKey sets SegmentServerSideKey field to given value.

### HasSegmentServerSideKey

`func (o *UnibeeApiMerchantProfileGetRes) HasSegmentServerSideKey() bool`

HasSegmentServerSideKey returns a boolean if a field has been set.

### GetSegmentUserPortalKey

`func (o *UnibeeApiMerchantProfileGetRes) GetSegmentUserPortalKey() string`

GetSegmentUserPortalKey returns the SegmentUserPortalKey field if non-nil, zero value otherwise.

### GetSegmentUserPortalKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetSegmentUserPortalKeyOk() (*string, bool)`

GetSegmentUserPortalKeyOk returns a tuple with the SegmentUserPortalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentUserPortalKey

`func (o *UnibeeApiMerchantProfileGetRes) SetSegmentUserPortalKey(v string)`

SetSegmentUserPortalKey sets SegmentUserPortalKey field to given value.

### HasSegmentUserPortalKey

`func (o *UnibeeApiMerchantProfileGetRes) HasSegmentUserPortalKey() bool`

HasSegmentUserPortalKey returns a boolean if a field has been set.

### GetSendGridKey

`func (o *UnibeeApiMerchantProfileGetRes) GetSendGridKey() string`

GetSendGridKey returns the SendGridKey field if non-nil, zero value otherwise.

### GetSendGridKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetSendGridKeyOk() (*string, bool)`

GetSendGridKeyOk returns a tuple with the SendGridKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendGridKey

`func (o *UnibeeApiMerchantProfileGetRes) SetSendGridKey(v string)`

SetSendGridKey sets SendGridKey field to given value.

### HasSendGridKey

`func (o *UnibeeApiMerchantProfileGetRes) HasSendGridKey() bool`

HasSendGridKey returns a boolean if a field has been set.

### GetTaxJarApiKeys

`func (o *UnibeeApiMerchantProfileGetRes) GetTaxJarApiKeys() UnibeeApiBeanUSVATGatewayConnectionAPIKeys`

GetTaxJarApiKeys returns the TaxJarApiKeys field if non-nil, zero value otherwise.

### GetTaxJarApiKeysOk

`func (o *UnibeeApiMerchantProfileGetRes) GetTaxJarApiKeysOk() (*UnibeeApiBeanUSVATGatewayConnectionAPIKeys, bool)`

GetTaxJarApiKeysOk returns a tuple with the TaxJarApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxJarApiKeys

`func (o *UnibeeApiMerchantProfileGetRes) SetTaxJarApiKeys(v UnibeeApiBeanUSVATGatewayConnectionAPIKeys)`

SetTaxJarApiKeys sets TaxJarApiKeys field to given value.

### HasTaxJarApiKeys

`func (o *UnibeeApiMerchantProfileGetRes) HasTaxJarApiKeys() bool`

HasTaxJarApiKeys returns a boolean if a field has been set.

### GetVatSenseKey

`func (o *UnibeeApiMerchantProfileGetRes) GetVatSenseKey() string`

GetVatSenseKey returns the VatSenseKey field if non-nil, zero value otherwise.

### GetVatSenseKeyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetVatSenseKeyOk() (*string, bool)`

GetVatSenseKeyOk returns a tuple with the VatSenseKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatSenseKey

`func (o *UnibeeApiMerchantProfileGetRes) SetVatSenseKey(v string)`

SetVatSenseKey sets VatSenseKey field to given value.

### HasVatSenseKey

`func (o *UnibeeApiMerchantProfileGetRes) HasVatSenseKey() bool`

HasVatSenseKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


