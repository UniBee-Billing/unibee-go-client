# UnibeeApiMerchantPaymentNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | CountryCode | [optional] 
**Currency** | **string** | Currency | 
**Email** | **string** | Email | 
**ExternalPaymentId** | **string** | ExternalPaymentId should unique for payment | 
**ExternalUserId** | **string** | ExternalUserId, should unique for user | 
**GatewayId** | **int64** | GatewayId | 
**LineItems** | Pointer to [**[]UnibeeApiMerchantPaymentItem**](UnibeeApiMerchantPaymentItem.md) | Items | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect Url | [optional] 
**TotalAmount** | **int64** | Total PaymentAmount, Cent | 

## Methods

### NewUnibeeApiMerchantPaymentNewReq

`func NewUnibeeApiMerchantPaymentNewReq(currency string, email string, externalPaymentId string, externalUserId string, gatewayId int64, totalAmount int64, ) *UnibeeApiMerchantPaymentNewReq`

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


### GetLineItems

`func (o *UnibeeApiMerchantPaymentNewReq) GetLineItems() []UnibeeApiMerchantPaymentItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *UnibeeApiMerchantPaymentNewReq) GetLineItemsOk() (*[]UnibeeApiMerchantPaymentItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *UnibeeApiMerchantPaymentNewReq) SetLineItems(v []UnibeeApiMerchantPaymentItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *UnibeeApiMerchantPaymentNewReq) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


