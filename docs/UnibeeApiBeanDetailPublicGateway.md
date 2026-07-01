# UnibeeApiBeanDetailPublicGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to **bool** |  | [optional] 
**AutoChargeEnabled** | Pointer to **bool** |  | [optional] 
**DefaultGatewayPaymentType** | Pointer to [**UnibeeApiBeanDetailPublicGatewayPaymentType**](UnibeeApiBeanDetailPublicGatewayPaymentType.md) |  | [optional] 
**Description** | Pointer to **string** | The description of gateway | [optional] 
**DisplayName** | Pointer to **string** | The gateway display name, used at checkout | [optional] 
**GatewayIcons** | Pointer to **[]string** | The gateway icons, used at checkout | [optional] 
**GatewayId** | Pointer to **int64** |  | [optional] 
**GatewayLogo** | Pointer to **string** |  | [optional] 
**GatewayName** | Pointer to **string** | The gateway name, stripe|paypal|changelly|unitpay|payssion|cryptadium | [optional] 
**GatewayPaymentTypes** | Pointer to [**[]UnibeeApiBeanDetailPublicGatewayPaymentType**](UnibeeApiBeanDetailPublicGatewayPaymentType.md) | gatewayPaymentTypes | [optional] 
**GatewayType** | Pointer to **int64** | gateway type: 1-Card | 2-Crypto | 3-Wire Transfer | 4-PayPal | 5-Credit | 6-Alipay+ | 7-Hyperswitch | 8-External | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The name of gateway | [optional] 

## Methods

### NewUnibeeApiBeanDetailPublicGateway

`func NewUnibeeApiBeanDetailPublicGateway() *UnibeeApiBeanDetailPublicGateway`

NewUnibeeApiBeanDetailPublicGateway instantiates a new UnibeeApiBeanDetailPublicGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailPublicGatewayWithDefaults

`func NewUnibeeApiBeanDetailPublicGatewayWithDefaults() *UnibeeApiBeanDetailPublicGateway`

NewUnibeeApiBeanDetailPublicGatewayWithDefaults instantiates a new UnibeeApiBeanDetailPublicGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *UnibeeApiBeanDetailPublicGateway) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *UnibeeApiBeanDetailPublicGateway) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *UnibeeApiBeanDetailPublicGateway) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### GetAutoChargeEnabled

`func (o *UnibeeApiBeanDetailPublicGateway) GetAutoChargeEnabled() bool`

GetAutoChargeEnabled returns the AutoChargeEnabled field if non-nil, zero value otherwise.

### GetAutoChargeEnabledOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetAutoChargeEnabledOk() (*bool, bool)`

GetAutoChargeEnabledOk returns a tuple with the AutoChargeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoChargeEnabled

`func (o *UnibeeApiBeanDetailPublicGateway) SetAutoChargeEnabled(v bool)`

SetAutoChargeEnabled sets AutoChargeEnabled field to given value.

### HasAutoChargeEnabled

`func (o *UnibeeApiBeanDetailPublicGateway) HasAutoChargeEnabled() bool`

HasAutoChargeEnabled returns a boolean if a field has been set.

### GetDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailPublicGateway) GetDefaultGatewayPaymentType() UnibeeApiBeanDetailPublicGatewayPaymentType`

GetDefaultGatewayPaymentType returns the DefaultGatewayPaymentType field if non-nil, zero value otherwise.

### GetDefaultGatewayPaymentTypeOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetDefaultGatewayPaymentTypeOk() (*UnibeeApiBeanDetailPublicGatewayPaymentType, bool)`

GetDefaultGatewayPaymentTypeOk returns a tuple with the DefaultGatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailPublicGateway) SetDefaultGatewayPaymentType(v UnibeeApiBeanDetailPublicGatewayPaymentType)`

SetDefaultGatewayPaymentType sets DefaultGatewayPaymentType field to given value.

### HasDefaultGatewayPaymentType

`func (o *UnibeeApiBeanDetailPublicGateway) HasDefaultGatewayPaymentType() bool`

HasDefaultGatewayPaymentType returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanDetailPublicGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanDetailPublicGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanDetailPublicGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UnibeeApiBeanDetailPublicGateway) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UnibeeApiBeanDetailPublicGateway) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UnibeeApiBeanDetailPublicGateway) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGatewayIcons

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayIcons() []string`

GetGatewayIcons returns the GatewayIcons field if non-nil, zero value otherwise.

### GetGatewayIconsOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayIconsOk() (*[]string, bool)`

GetGatewayIconsOk returns a tuple with the GatewayIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIcons

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayIcons(v []string)`

SetGatewayIcons sets GatewayIcons field to given value.

### HasGatewayIcons

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayIcons() bool`

HasGatewayIcons returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayLogo

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayLogo() string`

GetGatewayLogo returns the GatewayLogo field if non-nil, zero value otherwise.

### GetGatewayLogoOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayLogoOk() (*string, bool)`

GetGatewayLogoOk returns a tuple with the GatewayLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLogo

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayLogo(v string)`

SetGatewayLogo sets GatewayLogo field to given value.

### HasGatewayLogo

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayLogo() bool`

HasGatewayLogo returns a boolean if a field has been set.

### GetGatewayName

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### GetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayPaymentTypes() []UnibeeApiBeanDetailPublicGatewayPaymentType`

GetGatewayPaymentTypes returns the GatewayPaymentTypes field if non-nil, zero value otherwise.

### GetGatewayPaymentTypesOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayPaymentTypesOk() (*[]UnibeeApiBeanDetailPublicGatewayPaymentType, bool)`

GetGatewayPaymentTypesOk returns a tuple with the GatewayPaymentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayPaymentTypes(v []UnibeeApiBeanDetailPublicGatewayPaymentType)`

SetGatewayPaymentTypes sets GatewayPaymentTypes field to given value.

### HasGatewayPaymentTypes

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayPaymentTypes() bool`

HasGatewayPaymentTypes returns a boolean if a field has been set.

### GetGatewayType

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayType() int64`

GetGatewayType returns the GatewayType field if non-nil, zero value otherwise.

### GetGatewayTypeOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetGatewayTypeOk() (*int64, bool)`

GetGatewayTypeOk returns a tuple with the GatewayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayType

`func (o *UnibeeApiBeanDetailPublicGateway) SetGatewayType(v int64)`

SetGatewayType sets GatewayType field to given value.

### HasGatewayType

`func (o *UnibeeApiBeanDetailPublicGateway) HasGatewayType() bool`

HasGatewayType returns a boolean if a field has been set.

### GetIsDefault

`func (o *UnibeeApiBeanDetailPublicGateway) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UnibeeApiBeanDetailPublicGateway) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UnibeeApiBeanDetailPublicGateway) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanDetailPublicGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanDetailPublicGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanDetailPublicGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanDetailPublicGateway) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


