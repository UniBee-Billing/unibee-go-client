# UnibeeApiMerchantGatewayEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyExchange** | Pointer to [**[]UnibeeApiBeanDetailGatewayCurrencyExchange**](UnibeeApiBeanDetailGatewayCurrencyExchange.md) | The currency exchange for gateway payment, effect at start of payment creation when currency matched | [optional] 
**DisplayName** | Pointer to **string** | The displayName of payment gateway | [optional] 
**GatewayId** | **int64** | The id of payment gateway | 
**GatewayKey** | Pointer to **string** | The key of payment gateway | [optional] 
**GatewayLogo** | Pointer to **[][]string** | The logo of payment gateway | [optional] 
**GatewayPaymentTypes** | Pointer to **[]string** | Selected gateway payment types | [optional] 
**GatewaySecret** | Pointer to **string** | The secret of payment gateway | [optional] 
**Sort** | Pointer to **int32** | The sort value of payment gateway, The higher the value, the lower the ranking | [optional] 
**SubGateway** | Pointer to **string** | The sub gateway of payment gateway | [optional] 

## Methods

### NewUnibeeApiMerchantGatewayEditReq

`func NewUnibeeApiMerchantGatewayEditReq(gatewayId int64, ) *UnibeeApiMerchantGatewayEditReq`

NewUnibeeApiMerchantGatewayEditReq instantiates a new UnibeeApiMerchantGatewayEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantGatewayEditReqWithDefaults

`func NewUnibeeApiMerchantGatewayEditReqWithDefaults() *UnibeeApiMerchantGatewayEditReq`

NewUnibeeApiMerchantGatewayEditReqWithDefaults instantiates a new UnibeeApiMerchantGatewayEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyExchange

`func (o *UnibeeApiMerchantGatewayEditReq) GetCurrencyExchange() []UnibeeApiBeanDetailGatewayCurrencyExchange`

GetCurrencyExchange returns the CurrencyExchange field if non-nil, zero value otherwise.

### GetCurrencyExchangeOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetCurrencyExchangeOk() (*[]UnibeeApiBeanDetailGatewayCurrencyExchange, bool)`

GetCurrencyExchangeOk returns a tuple with the CurrencyExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyExchange

`func (o *UnibeeApiMerchantGatewayEditReq) SetCurrencyExchange(v []UnibeeApiBeanDetailGatewayCurrencyExchange)`

SetCurrencyExchange sets CurrencyExchange field to given value.

### HasCurrencyExchange

`func (o *UnibeeApiMerchantGatewayEditReq) HasCurrencyExchange() bool`

HasCurrencyExchange returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiMerchantGatewayEditReq) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiMerchantGatewayEditReq) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiMerchantGatewayEditReq) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetGatewayKey

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKey() string`

GetGatewayKey returns the GatewayKey field if non-nil, zero value otherwise.

### GetGatewayKeyOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayKeyOk() (*string, bool)`

GetGatewayKeyOk returns a tuple with the GatewayKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayKey

`func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayKey(v string)`

SetGatewayKey sets GatewayKey field to given value.

### HasGatewayKey

`func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayKey() bool`

HasGatewayKey returns a boolean if a field has been set.

### GetGatewayLogo

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayLogo() [][]string`

GetGatewayLogo returns the GatewayLogo field if non-nil, zero value otherwise.

### GetGatewayLogoOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayLogoOk() (*[][]string, bool)`

GetGatewayLogoOk returns a tuple with the GatewayLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLogo

`func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayLogo(v [][]string)`

SetGatewayLogo sets GatewayLogo field to given value.

### HasGatewayLogo

`func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayLogo() bool`

HasGatewayLogo returns a boolean if a field has been set.

### GetGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayPaymentTypes() []string`

GetGatewayPaymentTypes returns the GatewayPaymentTypes field if non-nil, zero value otherwise.

### GetGatewayPaymentTypesOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewayPaymentTypesOk() (*[]string, bool)`

GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewayEditReq) SetGatewayPaymentTypes(v []string)`

SetGatewayPaymentTypes sets GatewayPaymentTypes field to given value.

### HasGatewayPaymentTypes

`func (o *UnibeeApiMerchantGatewayEditReq) HasGatewayPaymentTypes() bool`

HasGatewayPaymentTypes returns a boolean if a field has been set.

### GetGatewaySecret

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecret() string`

GetGatewaySecret returns the GatewaySecret field if non-nil, zero value otherwise.

### GetGatewaySecretOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetGatewaySecretOk() (*string, bool)`

GetGatewaySecretOk returns a tuple with the GatewaySecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySecret

`func (o *UnibeeApiMerchantGatewayEditReq) SetGatewaySecret(v string)`

SetGatewaySecret sets GatewaySecret field to given value.

### HasGatewaySecret

`func (o *UnibeeApiMerchantGatewayEditReq) HasGatewaySecret() bool`

HasGatewaySecret returns a boolean if a field has been set.

### GetSort

`func (o *UnibeeApiMerchantGatewayEditReq) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UnibeeApiMerchantGatewayEditReq) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UnibeeApiMerchantGatewayEditReq) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSubGateway

`func (o *UnibeeApiMerchantGatewayEditReq) GetSubGateway() string`

GetSubGateway returns the SubGateway field if non-nil, zero value otherwise.

### GetSubGatewayOk

`func (o *UnibeeApiMerchantGatewayEditReq) GetSubGatewayOk() (*string, bool)`

GetSubGatewayOk returns a tuple with the SubGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubGateway

`func (o *UnibeeApiMerchantGatewayEditReq) SetSubGateway(v string)`

SetSubGateway sets SubGateway field to given value.

### HasSubGateway

`func (o *UnibeeApiMerchantGatewayEditReq) HasSubGateway() bool`

HasSubGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


