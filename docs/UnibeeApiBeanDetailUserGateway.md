# UnibeeApiBeanDetailUserGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bank** | Pointer to [**UnibeeApiBeanDetailGatewayBank**](UnibeeApiBeanDetailGatewayBank.md) |  | [optional] 
**Currency** | Pointer to **string** | The currency of wire transfer | [optional] 
**DefaultGatewayPaymentType** | Pointer to [**UnibeeApiBeanDetailUserGatewayPaymentType**](UnibeeApiBeanDetailUserGatewayPaymentType.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The gateway display name, used at user portal | [optional] 
**GatewayIcons** | Pointer to **[]string** | The gateway icons, used at user portal | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**GatewayLogo** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** | The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium | [optional] 
**GatewayPaymentTypes** | Pointer to [**[]UnibeeApiBeanDetailUserGatewayPaymentType**](UnibeeApiBeanDetailUserGatewayPaymentType.md) | gatewayPaymentTypes | [optional] 
**GatewayType** | Pointer to **int64** | gateway type: 1-Card | 2-Crypto | 3-Wire Transfer | 4-PayPal | 5-Credit | 6-Alipay+ | 7-Hyperswitch | 8-External | [optional] 
**MinimumAmount** | Pointer to **int64** | The minimum amount of wire transfer | [optional] 
**Name** | Pointer to **string** | The name of gateway | [optional] 

## Methods

### NewUnibeeApiBeanDetailUserGateway

`func NewUnibeeApiBeanDetailUserGateway() *UnibeeApiBeanDetailUserGateway`

NewUnibeeApiBeanDetailUserGateway instantiates a new UnibeeApiBeanDetailUserGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailUserGatewayWithDefaults

`func NewUnibeeApiBeanDetailUserGatewayWithDefaults() *UnibeeApiBeanDetailUserGateway`

NewUnibeeApiBeanDetailUserGatewayWithDefaults instantiates a new UnibeeApiBeanDetailUserGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBank

`func (o *UnibeeApiBeanDetailUserGateway) GetBank() UnibeeApiBeanDetailGatewayBank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *UnibeeApiBeanDetailUserGateway) GetBankOk() (*UnibeeApiBeanDetailGatewayBank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *UnibeeApiBeanDetailUserGateway) SetBank(v UnibeeApiBeanDetailGatewayBank)`

SetBank sets Bank field to given value.

### HasBank

`func (o *UnibeeApiBeanDetailUserGateway) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailUserGateway) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailUserGateway) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailUserGateway) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailUserGateway) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserGateway) GetDefaultGatewayPaymentType() UnibeeApiBeanDetailUserGatewayPaymentType`

GetDefaultGatewayPaymentType returns the DefaultGatewayPaymentType field if non-nil, zero value otherwise.

### GetDefaultGatewayPaymentTypeOk

`func (o *UnibeeApiBeanDetailUserGateway) GetDefaultGatewayPaymentTypeOk() (*UnibeeApiBeanDetailUserGatewayPaymentType, bool)`

GetDefaultGatewayPaymentTypeOk returns a tuple with the DefaultGatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserGateway) SetDefaultGatewayPaymentType(v UnibeeApiBeanDetailUserGatewayPaymentType)`

SetDefaultGatewayPaymentType sets DefaultGatewayPaymentType field to given value.

### HasDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailUserGateway) HasDefaultGatewayPaymentType() bool`

HasDefaultGatewayPaymentType returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiBeanDetailUserGateway) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiBeanDetailUserGateway) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiBeanDetailUserGateway) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiBeanDetailUserGateway) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayIcons

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayIcons() []string`

GetGatewayIcons returns the GatewayIcons field if non-nil, zero value otherwise.

### GetGatewayIconsOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayIconsOk() (*[]string, bool)`

GetGatewayIconsOk returns a tuple with the GatewayIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIcons

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayIcons(v []string)`

SetGatewayIcons sets GatewayIcons field to given value.

### HasGatewayIcons

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayIcons() bool`

HasGatewayIcons returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayLogo

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayLogo() string`

GetGatewayLogo returns the GatewayLogo field if non-nil, zero value otherwise.

### GetGatewayLogoOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayLogoOk() (*string, bool)`

GetGatewayLogoOk returns a tuple with the GatewayLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLogo

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayLogo(v string)`

SetGatewayLogo sets GatewayLogo field to given value.

### HasGatewayLogo

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayLogo() bool`

HasGatewayLogo returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayPaymentTypes() []UnibeeApiBeanDetailUserGatewayPaymentType`

GetGatewayPaymentTypes returns the GatewayPaymentTypes field if non-nil, zero value otherwise.

### GetGatewayPaymentTypesOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayPaymentTypesOk() (*[]UnibeeApiBeanDetailUserGatewayPaymentType, bool)`

GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayPaymentTypes(v []UnibeeApiBeanDetailUserGatewayPaymentType)`

SetGatewayPaymentTypes sets GatewayPaymentTypes field to given value.

### HasGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayPaymentTypes() bool`

HasGatewayPaymentTypes returns a boolean if a field has been set.

### GetGatewayType

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayType() int64`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *UnibeeApiBeanDetailUserGateway) GetGatewayTypeOk() (*int64, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *UnibeeApiBeanDetailUserGateway) SetGatewayType(v int64)`

SetGatewayType sets GatewayType field to given value.

### HasGatewayType

`func (o *UnibeeApiBeanDetailUserGateway) HasGatewayType() bool`

HasGatewayType returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *UnibeeApiBeanDetailUserGateway) GetMinimumAmount() int64`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *UnibeeApiBeanDetailUserGateway) GetMinimumAmountOk() (*int64, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *UnibeeApiBeanDetailUserGateway) SetMinimumAmount(v int64)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *UnibeeApiBeanDetailUserGateway) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanDetailUserGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanDetailUserGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanDetailUserGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanDetailUserGateway) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


