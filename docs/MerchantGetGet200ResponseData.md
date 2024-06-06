# MerchantGetGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]UnibeeApiBeanCurrency**](UnibeeApiBeanCurrency.md) | Currency List | [optional] 
**TimeZone** | Pointer to **[]string** | TimeZone List | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**Gateways** | Pointer to [**[]UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) | Gateway List | [optional] 
**IsProd** | Pointer to **bool** | Check System Env Is Prod, true|false | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchantSimplify**](UnibeeApiBeanMerchantSimplify.md) |  | [optional] 
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**OpenApiKey** | Pointer to **string** | OpenApiKey | [optional] 
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

### GetGateways

`func (o *MerchantGetGet200ResponseData) GetGateways() []UnibeeApiBeanGatewaySimplify`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *MerchantGetGet200ResponseData) GetGatewaysOk() (*[]UnibeeApiBeanGatewaySimplify, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *MerchantGetGet200ResponseData) SetGateways(v []UnibeeApiBeanGatewaySimplify)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *MerchantGetGet200ResponseData) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

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

`func (o *MerchantGetGet200ResponseData) GetMerchant() UnibeeApiBeanMerchantSimplify`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *MerchantGetGet200ResponseData) GetMerchantOk() (*UnibeeApiBeanMerchantSimplify, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *MerchantGetGet200ResponseData) SetMerchant(v UnibeeApiBeanMerchantSimplify)`

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


