# UnibeeApiMerchantProfileGetRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to [**[]UnibeeApiBeanCurrency**](UnibeeApiBeanCurrency.md) | Currency List | [optional] 
**MemberRoles** | Pointer to [**[]UnibeeApiBeanMerchantRoleSimplify**](UnibeeApiBeanMerchantRoleSimplify.md) | The member&#39;s role list&#39; | [optional] 
**TimeZone** | Pointer to **[]string** | TimeZone List | [optional] 
**Env** | Pointer to **string** | System Env, em: daily|stage|local|prod | [optional] 
**Gateways** | Pointer to [**[]UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) | Gateway List | [optional] 
**IsOwner** | Pointer to **bool** | Check Member is Owner | [optional] 
**IsProd** | Pointer to **bool** | Check System Env Is Prod, true|false | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchantSimplify**](UnibeeApiBeanMerchantSimplify.md) |  | [optional] 
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**OpenApiKey** | Pointer to **string** | OpenApiKey | [optional] 
**SegmentServerSideKey** | Pointer to **string** | SegmentServerSideKey | [optional] 
**SegmentUserPortalKey** | Pointer to **string** | SegmentUserPortalKey | [optional] 
**SendGridKey** | Pointer to **string** | SendGridKey | [optional] 
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

`func (o *UnibeeApiMerchantProfileGetRes) GetMemberRoles() []UnibeeApiBeanMerchantRoleSimplify`

GetMemberRoles returns the MemberRoles field if non-nil, zero value otherwise.

### GetMemberRolesOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMemberRolesOk() (*[]UnibeeApiBeanMerchantRoleSimplify, bool)`

GetMemberRolesOk returns a tuple with the MemberRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRoles

`func (o *UnibeeApiMerchantProfileGetRes) SetMemberRoles(v []UnibeeApiBeanMerchantRoleSimplify)`

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

### GetGateways

`func (o *UnibeeApiMerchantProfileGetRes) GetGateways() []UnibeeApiBeanGatewaySimplify`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *UnibeeApiMerchantProfileGetRes) GetGatewaysOk() (*[]UnibeeApiBeanGatewaySimplify, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *UnibeeApiMerchantProfileGetRes) SetGateways(v []UnibeeApiBeanGatewaySimplify)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *UnibeeApiMerchantProfileGetRes) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

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

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchant() UnibeeApiBeanMerchantSimplify`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiMerchantProfileGetRes) GetMerchantOk() (*UnibeeApiBeanMerchantSimplify, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiMerchantProfileGetRes) SetMerchant(v UnibeeApiBeanMerchantSimplify)`

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


