# UnibeeApiMerchantGatewaySetupReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyExchange** | Pointer to [**[]UnibeeApiBeanDetailGatewayCurrencyExchange**](UnibeeApiBeanDetailGatewayCurrencyExchange.md) | The currency exchange for gateway payment, effect at start of payment creation when currency matched | [optional] 
**DisplayName** | Pointer to **string** | The displayName of payment gateway | [optional] 
**GatewayIcons** | Pointer to **[][]string** | The icons of payment gateway | [optional] 
**GatewayKey** | Pointer to **string** | The key of payment gateway | [optional] 
**GatewayName** | **string** | The name of payment gateway, stripe|paypal|changelly|unitpay|payssion|cryptadium | 
**GatewayPaymentTypes** | Pointer to **[]string** | Selected gateway payment types | [optional] 
**GatewaySecret** | Pointer to **string** | The secret of payment gateway | [optional] 
**Sort** | Pointer to **int32** | The sort value of payment gateway, The higher the value, the lower the ranking | [optional] 
**SubGateway** | Pointer to **string** | The sub gateway of payment gateway | [optional] 

## Methods

### NewUnibeeApiMerchantGatewaySetupReq

`func NewUnibeeApiMerchantGatewaySetupReq(gatewayName string, ) *UnibeeApiMerchantGatewaySetupReq`

NewUnibeeApiMerchantGatewaySetupReq instantiates a new UnibeeApiMerchantGatewaySetupReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewaySetupReqWithDefaults

`func NewUnibeeApiMerchantGatewaySetupReqWithDefaults() *UnibeeApiMerchantGatewaySetupReq`

NewUnibeeApiMerchantGatewaySetupReqWithDefaults instantiates a new UnibeeApiMerchantGatewaySetupReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyExchange

`func (o *UnibeeApiMerchantGatewaySetupReq) GetCurrencyExchange() []UnibeeApiBeanDetailGatewayCurrencyExchange`

GetCurrencyExchange returns the CurrencyExchange field if non-nil, zero value otherwise.

### GetCurrencyExchangeOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetCurrencyExchangeOk() (*[]UnibeeApiBeanDetailGatewayCurrencyExchange, bool)`

GetCurrencyExchangeOk returns a tuple with the CurrencyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyExchange

`func (o *UnibeeApiMerchantGatewaySetupReq) SetCurrencyExchange(v []UnibeeApiBeanDetailGatewayCurrencyExchange)`

SetCurrencyExchange sets CurrencyExchange field to given value.

### HasCurrencyExchange

`func (o *UnibeeApiMerchantGatewaySetupReq) HasCurrencyExchange() bool`

HasCurrencyExchange returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiMerchantGatewaySetupReq) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiMerchantGatewaySetupReq) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiMerchantGatewaySetupReq) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayIcons

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayIcons() [][]string`

GetGatewayIcons returns the GatewayIcons field if non-nil, zero value otherwise.

### GetGatewayIconsOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayIconsOk() (*[][]string, bool)`

GetGatewayIconsOk returns a tuple with the GatewayIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIcons

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayIcons(v [][]string)`

SetGatewayIcons sets GatewayIcons field to given value.

### HasGatewayIcons

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewayIcons() bool`

HasGatewayIcons returns a boolean if a field has been set.

### GetGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.


### GetGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayPaymentTypes() []string`

GetGatewayPaymentTypes returns the GatewayPaymentTypes field if non-nil, zero value otherwise.

### GetGatewayPaymentTypesOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewayPaymentTypesOk() (*[]string, bool)`

GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewayPaymentTypes(v []string)`

SetGatewayPaymentTypes sets GatewayPaymentTypes field to given value.

### HasGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewayPaymentTypes() bool`

HasGatewayPaymentTypes returns a boolean if a field has been set.

### GetGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecret() string`

GetGatewaySecret returns the GatewaySecret field if non-nil, zero value otherwise.

### GetGatewaySecretOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetGatewaySecretOk() (*string, bool)`

GetGatewaySecretOk returns a tuple with the GatewaySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) SetGatewaySecret(v string)`

SetGatewaySecret sets GatewaySecret field to given value.

### HasGatewaySecret

`func (o *UnibeeApiMerchantGatewaySetupReq) HasGatewaySecret() bool`

HasGatewaySecret returns a boolean if a field has been set.

### GetSort

`func (o *UnibeeApiMerchantGatewaySetupReq) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UnibeeApiMerchantGatewaySetupReq) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UnibeeApiMerchantGatewaySetupReq) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSubGateway

`func (o *UnibeeApiMerchantGatewaySetupReq) GetSubGateway() string`

GetSubGateway returns the SubGateway field if non-nil, zero value otherwise.

### GetSubGatewayOk

`func (o *UnibeeApiMerchantGatewaySetupReq) GetSubGatewayOk() (*string, bool)`

GetSubGatewayOk returns a tuple with the SubGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGateway

`func (o *UnibeeApiMerchantGatewaySetupReq) SetSubGateway(v string)`

SetSubGateway sets SubGateway field to given value.

### HasSubGateway

`func (o *UnibeeApiMerchantGatewaySetupReq) HasSubGateway() bool`

HasSubGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


