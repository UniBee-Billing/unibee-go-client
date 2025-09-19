# UnibeeApiBeanDetailGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSetupFinished** | Pointer to **bool** | Whether the gateway finished setup process | [optional] 
**Archive** | Pointer to **bool** |  | [optional] 
**AutoChargeEnabled** | Pointer to **bool** |  | [optional] 
**Bank** | Pointer to [**UnibeeApiBeanDetailGatewayBank**](UnibeeApiBeanDetailGatewayBank.md) |  | [optional] 
**CountryConfig** | Pointer to **map[string]bool** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | The currency of wire transfer  | [optional] 
**CurrencyExchange** | Pointer to [**[]UnibeeApiBeanDetailGatewayCurrencyExchange**](UnibeeApiBeanDetailGatewayCurrencyExchange.md) | The currency exchange for gateway payment, effect at start of payment creation when currency matched | [optional] 
**CurrencyExchangeEnabled** | Pointer to **bool** | whether to enable currency exchange | [optional] 
**Description** | Pointer to **string** | The description of gateway | [optional] 
**DisplayName** | Pointer to **string** | The gateway display name, used at user portal | [optional] 
**GatewayIcons** | Pointer to **[]string** | The gateway display name, used at user portal | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**GatewayKey** | Pointer to **string** |  | [optional] 
**GatewayLogo** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** | The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium | [optional] 
**GatewayPaymentTypes** | Pointer to [**[]UnibeeInternalInterfaceGatewayPaymentType**](UnibeeInternalInterfaceGatewayPaymentType.md) | gatewayPaymentTypes | [optional] 
**GatewaySecret** | Pointer to **string** |  | [optional] 
**GatewayType** | Pointer to **int64** | gateway type，1-Bank Card ｜ 2-Crypto | 3 - Wire Transfer | [optional] 
**GatewayWebhookIntegrationLink** | Pointer to **string** | The gateway webhook integration guide link, gateway webhook need setup if not blank | [optional] 
**GatewayWebsiteLink** | Pointer to **string** | The gateway website link | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**MinimumAmount** | Pointer to **int64** | The minimum amount of wire transfer | [optional] 
**Name** | Pointer to **string** | The name of gateway | [optional] 
**PrivateSecretName** | Pointer to **string** |  | [optional] 
**PublicKeyName** | Pointer to **string** |  | [optional] 
**SetupGatewayPaymentTypes** | Pointer to [**[]UnibeeInternalInterfaceGatewayPaymentType**](UnibeeInternalInterfaceGatewayPaymentType.md) | The total list of gateway payment types, used for setup | [optional] 
**Sort** | Pointer to **int64** | The sort value of payment gateway, The higher the value, the lower the ranking | [optional] 
**SubGateway** | Pointer to **string** |  | [optional] 
**SubGatewayName** | Pointer to **string** |  | [optional] 
**WebhookEndpointUrl** | Pointer to **string** | The endpoint url of gateway webhook  | [optional] 
**WebhookSecret** | Pointer to **string** | The secret of gateway webhook | [optional] 

## Methods

### NewUnibeeApiBeanDetailGateway

`func NewUnibeeApiBeanDetailGateway() *UnibeeApiBeanDetailGateway`

NewUnibeeApiBeanDetailGateway instantiates a new UnibeeApiBeanDetailGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailGatewayWithDefaults

`func NewUnibeeApiBeanDetailGatewayWithDefaults() *UnibeeApiBeanDetailGateway`

NewUnibeeApiBeanDetailGatewayWithDefaults instantiates a new UnibeeApiBeanDetailGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSetupFinished

`func (o *UnibeeApiBeanDetailGateway) GetIsSetupFinished() bool`

GetIsSetupFinished returns the IsSetupFinished field if non-nil, zero value otherwise.

### GetIsSetupFinishedOk

`func (o *UnibeeApiBeanDetailGateway) GetIsSetupFinishedOk() (*bool, bool)`

GetIsSetupFinishedOk returns a tuple with the IsSetupFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetupFinished

`func (o *UnibeeApiBeanDetailGateway) SetIsSetupFinished(v bool)`

SetIsSetupFinished sets IsSetupFinished field to given value.

### HasIsSetupFinished

`func (o *UnibeeApiBeanDetailGateway) HasIsSetupFinished() bool`

HasIsSetupFinished returns a boolean if a field has been set.

### GetArchive

`func (o *UnibeeApiBeanDetailGateway) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *UnibeeApiBeanDetailGateway) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *UnibeeApiBeanDetailGateway) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *UnibeeApiBeanDetailGateway) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetAutoChargeEnabled

`func (o *UnibeeApiBeanDetailGateway) GetAutoChargeEnabled() bool`

GetAutoChargeEnabled returns the AutoChargeEnabled field if non-nil, zero value otherwise.

### GetAutoChargeEnabledOk

`func (o *UnibeeApiBeanDetailGateway) GetAutoChargeEnabledOk() (*bool, bool)`

GetAutoChargeEnabledOk returns a tuple with the AutoChargeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeEnabled

`func (o *UnibeeApiBeanDetailGateway) SetAutoChargeEnabled(v bool)`

SetAutoChargeEnabled sets AutoChargeEnabled field to given value.

### HasAutoChargeEnabled

`func (o *UnibeeApiBeanDetailGateway) HasAutoChargeEnabled() bool`

HasAutoChargeEnabled returns a boolean if a field has been set.

### GetBank

`func (o *UnibeeApiBeanDetailGateway) GetBank() UnibeeApiBeanDetailGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiBeanDetailGateway) GetBankOk() (*UnibeeApiBeanDetailGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiBeanDetailGateway) SetBank(v UnibeeApiBeanDetailGatewayBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UnibeeApiBeanDetailGateway) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCountryConfig

`func (o *UnibeeApiBeanDetailGateway) GetCountryConfig() map[string]bool`

GetCountryConfig returns the CountryConfig field if non-nil, zero value otherwise.

### GetCountryConfigOk

`func (o *UnibeeApiBeanDetailGateway) GetCountryConfigOk() (*map[string]bool, bool)`

GetCountryConfigOk returns a tuple with the CountryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryConfig

`func (o *UnibeeApiBeanDetailGateway) SetCountryConfig(v map[string]bool)`

SetCountryConfig sets CountryConfig field to given value.

### HasCountryConfig

`func (o *UnibeeApiBeanDetailGateway) HasCountryConfig() bool`

HasCountryConfig returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailGateway) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailGateway) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailGateway) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailGateway) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailGateway) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailGateway) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailGateway) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailGateway) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrencyExchange

`func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchange() []UnibeeApiBeanDetailGatewayCurrencyExchange`

GetCurrencyExchange returns the CurrencyExchange field if non-nil, zero value otherwise.

### GetCurrencyExchangeOk

`func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeOk() (*[]UnibeeApiBeanDetailGatewayCurrencyExchange, bool)`

GetCurrencyExchangeOk returns a tuple with the CurrencyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyExchange

`func (o *UnibeeApiBeanDetailGateway) SetCurrencyExchange(v []UnibeeApiBeanDetailGatewayCurrencyExchange)`

SetCurrencyExchange sets CurrencyExchange field to given value.

### HasCurrencyExchange

`func (o *UnibeeApiBeanDetailGateway) HasCurrencyExchange() bool`

HasCurrencyExchange returns a boolean if a field has been set.

### GetCurrencyExchangeEnabled

`func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeEnabled() bool`

GetCurrencyExchangeEnabled returns the CurrencyExchangeEnabled field if non-nil, zero value otherwise.

### GetCurrencyExchangeEnabledOk

`func (o *UnibeeApiBeanDetailGateway) GetCurrencyExchangeEnabledOk() (*bool, bool)`

GetCurrencyExchangeEnabledOk returns a tuple with the CurrencyExchangeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyExchangeEnabled

`func (o *UnibeeApiBeanDetailGateway) SetCurrencyExchangeEnabled(v bool)`

SetCurrencyExchangeEnabled sets CurrencyExchangeEnabled field to given value.

### HasCurrencyExchangeEnabled

`func (o *UnibeeApiBeanDetailGateway) HasCurrencyExchangeEnabled() bool`

HasCurrencyExchangeEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanDetailGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanDetailGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanDetailGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanDetailGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiBeanDetailGateway) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiBeanDetailGateway) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiBeanDetailGateway) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiBeanDetailGateway) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayIcons

`func (o *UnibeeApiBeanDetailGateway) GetGatewayIcons() []string`

GetGatewayIcons returns the GatewayIcons field if non-nil, zero value otherwise.

### GetGatewayIconsOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayIconsOk() (*[]string, bool)`

GetGatewayIconsOk returns a tuple with the GatewayIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIcons

`func (o *UnibeeApiBeanDetailGateway) SetGatewayIcons(v []string)`

SetGatewayIcons sets GatewayIcons field to given value.

### HasGatewayIcons

`func (o *UnibeeApiBeanDetailGateway) HasGatewayIcons() bool`

HasGatewayIcons returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailGateway) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailGateway) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailGateway) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayKey

`func (o *UnibeeApiBeanDetailGateway) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *UnibeeApiBeanDetailGateway) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *UnibeeApiBeanDetailGateway) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetGatewayLogo

`func (o *UnibeeApiBeanDetailGateway) GetGatewayLogo() string`

GetGatewayLogo returns the GatewayLogo field if non-nil, zero value otherwise.

### GetGatewayLogoOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayLogoOk() (*string, bool)`

GetGatewayLogoOk returns a tuple with the GatewayLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLogo

`func (o *UnibeeApiBeanDetailGateway) SetGatewayLogo(v string)`

SetGatewayLogo sets GatewayLogo field to given value.

### HasGatewayLogo

`func (o *UnibeeApiBeanDetailGateway) HasGatewayLogo() bool`

HasGatewayLogo returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanDetailGateway) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanDetailGateway) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanDetailGateway) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) GetGatewayPaymentTypes() []UnibeeInternalInterfaceGatewayPaymentType`

GetGatewayPaymentTypes returns the GatewayPaymentTypes field if non-nil, zero value otherwise.

### GetGatewayPaymentTypesOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayPaymentTypesOk() (*[]UnibeeInternalInterfaceGatewayPaymentType, bool)`

GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) SetGatewayPaymentTypes(v []UnibeeInternalInterfaceGatewayPaymentType)`

SetGatewayPaymentTypes sets GatewayPaymentTypes field to given value.

### HasGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) HasGatewayPaymentTypes() bool`

HasGatewayPaymentTypes returns a boolean if a field has been set.

### GetGatewaySecret

`func (o *UnibeeApiBeanDetailGateway) GetGatewaySecret() string`

GetGatewaySecret returns the GatewaySecret field if non-nil, zero value otherwise.

### GetGatewaySecretOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewaySecretOk() (*string, bool)`

GetGatewaySecretOk returns a tuple with the GatewaySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySecret

`func (o *UnibeeApiBeanDetailGateway) SetGatewaySecret(v string)`

SetGatewaySecret sets GatewaySecret field to given value.

### HasGatewaySecret

`func (o *UnibeeApiBeanDetailGateway) HasGatewaySecret() bool`

HasGatewaySecret returns a boolean if a field has been set.

### GetGatewayType

`func (o *UnibeeApiBeanDetailGateway) GetGatewayType() int64`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayTypeOk() (*int64, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *UnibeeApiBeanDetailGateway) SetGatewayType(v int64)`

SetGatewayType sets GatewayType field to given value.

### HasGatewayType

`func (o *UnibeeApiBeanDetailGateway) HasGatewayType() bool`

HasGatewayType returns a boolean if a field has been set.

### GetGatewayWebhookIntegrationLink

`func (o *UnibeeApiBeanDetailGateway) GetGatewayWebhookIntegrationLink() string`

GetGatewayWebhookIntegrationLink returns the GatewayWebhookIntegrationLink field if non-nil, zero value otherwise.

### GetGatewayWebhookIntegrationLinkOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayWebhookIntegrationLinkOk() (*string, bool)`

GetGatewayWebhookIntegrationLinkOk returns a tuple with the GatewayWebhookIntegrationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayWebhookIntegrationLink

`func (o *UnibeeApiBeanDetailGateway) SetGatewayWebhookIntegrationLink(v string)`

SetGatewayWebhookIntegrationLink sets GatewayWebhookIntegrationLink field to given value.

### HasGatewayWebhookIntegrationLink

`func (o *UnibeeApiBeanDetailGateway) HasGatewayWebhookIntegrationLink() bool`

HasGatewayWebhookIntegrationLink returns a boolean if a field has been set.

### GetGatewayWebsiteLink

`func (o *UnibeeApiBeanDetailGateway) GetGatewayWebsiteLink() string`

GetGatewayWebsiteLink returns the GatewayWebsiteLink field if non-nil, zero value otherwise.

### GetGatewayWebsiteLinkOk

`func (o *UnibeeApiBeanDetailGateway) GetGatewayWebsiteLinkOk() (*string, bool)`

GetGatewayWebsiteLinkOk returns a tuple with the GatewayWebsiteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayWebsiteLink

`func (o *UnibeeApiBeanDetailGateway) SetGatewayWebsiteLink(v string)`

SetGatewayWebsiteLink sets GatewayWebsiteLink field to given value.

### HasGatewayWebsiteLink

`func (o *UnibeeApiBeanDetailGateway) HasGatewayWebsiteLink() bool`

HasGatewayWebsiteLink returns a boolean if a field has been set.

### GetIsDefault

`func (o *UnibeeApiBeanDetailGateway) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiBeanDetailGateway) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiBeanDetailGateway) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiBeanDetailGateway) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *UnibeeApiBeanDetailGateway) GetMinimumAmount() int64`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnibeeApiBeanDetailGateway) GetMinimumAmountOk() (*int64, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnibeeApiBeanDetailGateway) SetMinimumAmount(v int64)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *UnibeeApiBeanDetailGateway) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanDetailGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanDetailGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanDetailGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanDetailGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateSecretName

`func (o *UnibeeApiBeanDetailGateway) GetPrivateSecretName() string`

GetPrivateSecretName returns the PrivateSecretName field if non-nil, zero value otherwise.

### GetPrivateSecretNameOk

`func (o *UnibeeApiBeanDetailGateway) GetPrivateSecretNameOk() (*string, bool)`

GetPrivateSecretNameOk returns a tuple with the PrivateSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSecretName

`func (o *UnibeeApiBeanDetailGateway) SetPrivateSecretName(v string)`

SetPrivateSecretName sets PrivateSecretName field to given value.

### HasPrivateSecretName

`func (o *UnibeeApiBeanDetailGateway) HasPrivateSecretName() bool`

HasPrivateSecretName returns a boolean if a field has been set.

### GetPublicKeyName

`func (o *UnibeeApiBeanDetailGateway) GetPublicKeyName() string`

GetPublicKeyName returns the PublicKeyName field if non-nil, zero value otherwise.

### GetPublicKeyNameOk

`func (o *UnibeeApiBeanDetailGateway) GetPublicKeyNameOk() (*string, bool)`

GetPublicKeyNameOk returns a tuple with the PublicKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyName

`func (o *UnibeeApiBeanDetailGateway) SetPublicKeyName(v string)`

SetPublicKeyName sets PublicKeyName field to given value.

### HasPublicKeyName

`func (o *UnibeeApiBeanDetailGateway) HasPublicKeyName() bool`

HasPublicKeyName returns a boolean if a field has been set.

### GetSetupGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) GetSetupGatewayPaymentTypes() []UnibeeInternalInterfaceGatewayPaymentType`

GetSetupGatewayPaymentTypes returns the SetupGatewayPaymentTypes field if non-nil, zero value otherwise.

### GetSetupGatewayPaymentTypesOk

`func (o *UnibeeApiBeanDetailGateway) GetSetupGatewayPaymentTypesOk() (*[]UnibeeInternalInterfaceGatewayPaymentType, bool)`

GetSetupGatewayPaymentTypesOk returns a tuple with the SetupGatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) SetSetupGatewayPaymentTypes(v []UnibeeInternalInterfaceGatewayPaymentType)`

SetSetupGatewayPaymentTypes sets SetupGatewayPaymentTypes field to given value.

### HasSetupGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailGateway) HasSetupGatewayPaymentTypes() bool`

HasSetupGatewayPaymentTypes returns a boolean if a field has been set.

### GetSort

`func (o *UnibeeApiBeanDetailGateway) GetSort() int64`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UnibeeApiBeanDetailGateway) GetSortOk() (*int64, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UnibeeApiBeanDetailGateway) SetSort(v int64)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UnibeeApiBeanDetailGateway) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSubGateway

`func (o *UnibeeApiBeanDetailGateway) GetSubGateway() string`

GetSubGateway returns the SubGateway field if non-nil, zero value otherwise.

### GetSubGatewayOk

`func (o *UnibeeApiBeanDetailGateway) GetSubGatewayOk() (*string, bool)`

GetSubGatewayOk returns a tuple with the SubGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGateway

`func (o *UnibeeApiBeanDetailGateway) SetSubGateway(v string)`

SetSubGateway sets SubGateway field to given value.

### HasSubGateway

`func (o *UnibeeApiBeanDetailGateway) HasSubGateway() bool`

HasSubGateway returns a boolean if a field has been set.

### GetSubGatewayName

`func (o *UnibeeApiBeanDetailGateway) GetSubGatewayName() string`

GetSubGatewayName returns the SubGatewayName field if non-nil, zero value otherwise.

### GetSubGatewayNameOk

`func (o *UnibeeApiBeanDetailGateway) GetSubGatewayNameOk() (*string, bool)`

GetSubGatewayNameOk returns a tuple with the SubGatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGatewayName

`func (o *UnibeeApiBeanDetailGateway) SetSubGatewayName(v string)`

SetSubGatewayName sets SubGatewayName field to given value.

### HasSubGatewayName

`func (o *UnibeeApiBeanDetailGateway) HasSubGatewayName() bool`

HasSubGatewayName returns a boolean if a field has been set.

### GetWebhookEndpointUrl

`func (o *UnibeeApiBeanDetailGateway) GetWebhookEndpointUrl() string`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *UnibeeApiBeanDetailGateway) GetWebhookEndpointUrlOk() (*string, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *UnibeeApiBeanDetailGateway) SetWebhookEndpointUrl(v string)`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *UnibeeApiBeanDetailGateway) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *UnibeeApiBeanDetailGateway) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *UnibeeApiBeanDetailGateway) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *UnibeeApiBeanDetailGateway) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *UnibeeApiBeanDetailGateway) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


