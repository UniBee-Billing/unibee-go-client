# UnibeeApiMerchantPaymentNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**Currency** | Pointer to **string** | Currency, either Currency&amp;Currency or PlanId needed | [optional] 
**Email** | Pointer to **string** | Email, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ExternalPaymentId** | Pointer to **string** | ExternalPaymentId should unique for payment | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**GatewayId** | **int64** | GatewayId | 
**Items** | Pointer to [**[]UnibeeApiMerchantPaymentItem**](UnibeeApiMerchantPaymentItem.md) | Items | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**PlanId** | Pointer to **int64** | PlanId, either TotalAmount&amp;Currency or PlanId needed | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect Url | [optional] 
**TotalAmount** | Pointer to **int64** | Total PaymentAmount, Cent, either TotalAmount&amp;Currency or PlanId needed | [optional] 
**UserId** | Pointer to **int64** | UserId, either ExternalUserId&amp;Email or UserId needed | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentNewReq

`func NewUnibeeApiMerchantPaymentNewReq(gatewayId int64, ) *UnibeeApiMerchantPaymentNewReq`

NewUnibeeApiMerchantPaymentNewReq instantiates a new UnibeeApiMerchantPaymentNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentNewReqWithDefaults

`func NewUnibeeApiMerchantPaymentNewReqWithDefaults() *UnibeeApiMerchantPaymentNewReq`

NewUnibeeApiMerchantPaymentNewReqWithDefaults instantiates a new UnibeeApiMerchantPaymentNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantPaymentNewReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantPaymentNewReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantPaymentNewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantPaymentNewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantPaymentNewReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantPaymentNewReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantPaymentNewReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantPaymentNewReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewReq) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeApiMerchantPaymentNewReq) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantPaymentNewReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantPaymentNewReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiMerchantPaymentNewReq) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiMerchantPaymentNewReq) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiMerchantPaymentNewReq) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentNewReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetItems

`func (o *UnibeeApiMerchantPaymentNewReq) GetItems() []UnibeeApiMerchantPaymentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetItemsOk() (*[]UnibeeApiMerchantPaymentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiMerchantPaymentNewReq) SetItems(v []UnibeeApiMerchantPaymentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UnibeeApiMerchantPaymentNewReq) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantPaymentNewReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantPaymentNewReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantPaymentNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantPaymentNewReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantPaymentNewReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantPaymentNewReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UnibeeApiMerchantPaymentNewReq) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UnibeeApiMerchantPaymentNewReq) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantPaymentNewReq) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantPaymentNewReq) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantPaymentNewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantPaymentNewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantPaymentNewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


