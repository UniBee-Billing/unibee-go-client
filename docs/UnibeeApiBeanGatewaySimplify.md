# UnibeeApiBeanGatewaySimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**UnibeeApiBeanGatewayBank**](UnibeeApiBeanGatewayBank.md) |  | [optional] 
**CountryConfig** | Pointer to **map[string]bool** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | The currency of wire transfer  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**GatewayKey** | Pointer to **string** |  | [optional] 
**GatewayLogo** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** |  | [optional] 
**GatewayType** | Pointer to **int64** | gateway type，1-Bank Card ｜ 2-Crypto | 3 - Wire Transfer | [optional] 
**MinimumAmount** | Pointer to **int64** | The minimum amount of wire transfer | [optional] 
**WebhookEndpointUrl** | Pointer to **string** | The endpoint url of gateway webhook  | [optional] 
**WebhookSecret** | Pointer to **string** | The secret of gateway webhook | [optional] 

## Methods

### NewUnibeeApiBeanGatewaySimplify

`func NewUnibeeApiBeanGatewaySimplify() *UnibeeApiBeanGatewaySimplify`

NewUnibeeApiBeanGatewaySimplify instantiates a new UnibeeApiBeanGatewaySimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanGatewaySimplifyWithDefaults

`func NewUnibeeApiBeanGatewaySimplifyWithDefaults() *UnibeeApiBeanGatewaySimplify`

NewUnibeeApiBeanGatewaySimplifyWithDefaults instantiates a new UnibeeApiBeanGatewaySimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *UnibeeApiBeanGatewaySimplify) GetBank() UnibeeApiBeanGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiBeanGatewaySimplify) GetBankOk() (*UnibeeApiBeanGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiBeanGatewaySimplify) SetBank(v UnibeeApiBeanGatewayBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UnibeeApiBeanGatewaySimplify) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCountryConfig

`func (o *UnibeeApiBeanGatewaySimplify) GetCountryConfig() map[string]bool`

GetCountryConfig returns the CountryConfig field if non-nil, zero value otherwise.

### GetCountryConfigOk

`func (o *UnibeeApiBeanGatewaySimplify) GetCountryConfigOk() (*map[string]bool, bool)`

GetCountryConfigOk returns a tuple with the CountryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryConfig

`func (o *UnibeeApiBeanGatewaySimplify) SetCountryConfig(v map[string]bool)`

SetCountryConfig sets CountryConfig field to given value.

### HasCountryConfig

`func (o *UnibeeApiBeanGatewaySimplify) HasCountryConfig() bool`

HasCountryConfig returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanGatewaySimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanGatewaySimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanGatewaySimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanGatewaySimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanGatewaySimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanGatewaySimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanGatewaySimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanGatewaySimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiBeanGatewaySimplify) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiBeanGatewaySimplify) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiBeanGatewaySimplify) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiBeanGatewaySimplify) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanGatewaySimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanGatewaySimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayKey

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *UnibeeApiBeanGatewaySimplify) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *UnibeeApiBeanGatewaySimplify) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetGatewayLogo

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayLogo() string`

GetGatewayLogo returns the GatewayLogo field if non-nil, zero value otherwise.

### GetGatewayLogoOk

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayLogoOk() (*string, bool)`

GetGatewayLogoOk returns a tuple with the GatewayLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLogo

`func (o *UnibeeApiBeanGatewaySimplify) SetGatewayLogo(v string)`

SetGatewayLogo sets GatewayLogo field to given value.

### HasGatewayLogo

`func (o *UnibeeApiBeanGatewaySimplify) HasGatewayLogo() bool`

HasGatewayLogo returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanGatewaySimplify) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanGatewaySimplify) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGatewayType

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayType() int64`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *UnibeeApiBeanGatewaySimplify) GetGatewayTypeOk() (*int64, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *UnibeeApiBeanGatewaySimplify) SetGatewayType(v int64)`

SetGatewayType sets GatewayType field to given value.

### HasGatewayType

`func (o *UnibeeApiBeanGatewaySimplify) HasGatewayType() bool`

HasGatewayType returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *UnibeeApiBeanGatewaySimplify) GetMinimumAmount() int64`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnibeeApiBeanGatewaySimplify) GetMinimumAmountOk() (*int64, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnibeeApiBeanGatewaySimplify) SetMinimumAmount(v int64)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *UnibeeApiBeanGatewaySimplify) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetWebhookEndpointUrl

`func (o *UnibeeApiBeanGatewaySimplify) GetWebhookEndpointUrl() string`

GetWebhookEndpointUrl returns the WebhookEndpointUrl field if non-nil, zero value otherwise.

### GetWebhookEndpointUrlOk

`func (o *UnibeeApiBeanGatewaySimplify) GetWebhookEndpointUrlOk() (*string, bool)`

GetWebhookEndpointUrlOk returns a tuple with the WebhookEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookEndpointUrl

`func (o *UnibeeApiBeanGatewaySimplify) SetWebhookEndpointUrl(v string)`

SetWebhookEndpointUrl sets WebhookEndpointUrl field to given value.

### HasWebhookEndpointUrl

`func (o *UnibeeApiBeanGatewaySimplify) HasWebhookEndpointUrl() bool`

HasWebhookEndpointUrl returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *UnibeeApiBeanGatewaySimplify) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *UnibeeApiBeanGatewaySimplify) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *UnibeeApiBeanGatewaySimplify) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *UnibeeApiBeanGatewaySimplify) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


