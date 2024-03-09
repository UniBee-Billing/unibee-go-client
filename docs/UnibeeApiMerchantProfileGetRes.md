# UnibeeApiMerchantProfileGetRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]UnibeeInternalLogicGatewayRoCurrency**](UnibeeInternalLogicGatewayRoCurrency.md) | Currency List | [optional] 
**TimeZone** | Pointer to **[]string** | TimeZone List | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**Gateway** | Pointer to [**[]UnibeeInternalLogicGatewayRoGatewaySimplify**](UnibeeInternalLogicGatewayRoGatewaySimplify.md) | Gateway List | [optional] 
**IsProd** | Pointer to **bool** | Check System Env Is Prod, true|false | [optional] 
**Merchant** | Pointer to [**UnibeeInternalModelEntityOverseaPayMerchant**](UnibeeInternalModelEntityOverseaPayMerchant.md) |  | [optional] 

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

`func (o *UnibeeApiMerchantProfileGetRes) GetCurrency() []UnibeeInternalLogicGatewayRoCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantProfileGetRes) GetCurrencyOk() (*[]UnibeeInternalLogicGatewayRoCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantProfileGetRes) SetCurrency(v []UnibeeInternalLogicGatewayRoCurrency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantProfileGetRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

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

### GetGateway

`func (o *UnibeeApiMerchantProfileGetRes) GetGateway() []UnibeeInternalLogicGatewayRoGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantProfileGetRes) GetGatewayOk() (*[]UnibeeInternalLogicGatewayRoGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantProfileGetRes) SetGateway(v []UnibeeInternalLogicGatewayRoGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantProfileGetRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

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

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchant() UnibeeInternalModelEntityOverseaPayMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchantOk() (*UnibeeInternalModelEntityOverseaPayMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiMerchantProfileGetRes) SetMerchant(v UnibeeInternalModelEntityOverseaPayMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiMerchantProfileGetRes) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


