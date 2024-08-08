# UnibeeApiMerchantSubscriptionNewPaymentReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** | CancelUrl | [optional] 
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**Currency** | Pointer to **string** | Currency, either Currency&amp;TotalAmount or PlanId needed | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Email** | Pointer to **string** | Email, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ExternalPaymentId** | Pointer to **string** | ExternalPaymentId should unique for payment | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**GatewayId** | **int64** | GatewayId | 
**Items** | Pointer to [**[]UnibeeApiMerchantPaymentItem**](UnibeeApiMerchantPaymentItem.md) | Items | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**PlanId** | Pointer to **int64** | PlanId, either TotalAmount&amp;Currency or PlanId needed | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect Url | [optional] 
**TotalAmount** | Pointer to **int64** | Total PaymentAmount, Cent, either TotalAmount&amp;Currency or PlanId needed | [optional] 
**UserId** | Pointer to **int64** | UserId, either ExternalUserId&amp;Email or UserId needed | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionNewPaymentReq

`func NewUnibeeApiMerchantSubscriptionNewPaymentReq(gatewayId int64, ) *UnibeeApiMerchantSubscriptionNewPaymentReq`

NewUnibeeApiMerchantSubscriptionNewPaymentReq instantiates a new UnibeeApiMerchantSubscriptionNewPaymentReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionNewPaymentReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionNewPaymentReqWithDefaults() *UnibeeApiMerchantSubscriptionNewPaymentReq`

NewUnibeeApiMerchantSubscriptionNewPaymentReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionNewPaymentReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetItems

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetItems() []UnibeeApiMerchantPaymentItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetItemsOk() (*[]UnibeeApiMerchantPaymentItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetItems(v []UnibeeApiMerchantPaymentItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionNewPaymentReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


