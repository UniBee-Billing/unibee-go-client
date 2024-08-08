# UnibeeApiBeanRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | country code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**ExternalRefundId** | Pointer to **string** | external_refund_id | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayRefundId** | Pointer to **string** | gateway refund id | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PaymentId** | Pointer to **string** | relative payment id | [optional] 
**RefundAmount** | Pointer to **int64** | refund amount, cent | [optional] 
**RefundComment** | Pointer to **string** | refund comment | [optional] 
**RefundId** | Pointer to **string** | refund id (system generate) | [optional] 
**RefundTime** | Pointer to **int64** | refund success time | [optional] 
**ReturnUrl** | Pointer to **string** | return url after refund success | [optional] 
**Status** | Pointer to **int32** | status。10-pending，20-success，30-failure, 40-cancel | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**Type** | Pointer to **int32** | 1-gateway refund,2-mark refund | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanRefund

`func NewUnibeeApiBeanRefund() *UnibeeApiBeanRefund`

NewUnibeeApiBeanRefund instantiates a new UnibeeApiBeanRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanRefundWithDefaults

`func NewUnibeeApiBeanRefundWithDefaults() *UnibeeApiBeanRefund`

NewUnibeeApiBeanRefundWithDefaults instantiates a new UnibeeApiBeanRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiBeanRefund) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanRefund) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanRefund) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanRefund) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanRefund) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanRefund) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanRefund) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanRefund) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanRefund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanRefund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanRefund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanRefund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalRefundId

`func (o *UnibeeApiBeanRefund) GetExternalRefundId() string`

GetExternalRefundId returns the ExternalRefundId field if non-nil, zero value otherwise.

### GetExternalRefundIdOk

`func (o *UnibeeApiBeanRefund) GetExternalRefundIdOk() (*string, bool)`

GetExternalRefundIdOk returns a tuple with the ExternalRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefundId

`func (o *UnibeeApiBeanRefund) SetExternalRefundId(v string)`

SetExternalRefundId sets ExternalRefundId field to given value.

### HasExternalRefundId

`func (o *UnibeeApiBeanRefund) HasExternalRefundId() bool`

HasExternalRefundId returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanRefund) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanRefund) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanRefund) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanRefund) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayRefundId

`func (o *UnibeeApiBeanRefund) GetGatewayRefundId() string`

GetGatewayRefundId returns the GatewayRefundId field if non-nil, zero value otherwise.

### GetGatewayRefundIdOk

`func (o *UnibeeApiBeanRefund) GetGatewayRefundIdOk() (*string, bool)`

GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRefundId

`func (o *UnibeeApiBeanRefund) SetGatewayRefundId(v string)`

SetGatewayRefundId sets GatewayRefundId field to given value.

### HasGatewayRefundId

`func (o *UnibeeApiBeanRefund) HasGatewayRefundId() bool`

HasGatewayRefundId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanRefund) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanRefund) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanRefund) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanRefund) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanRefund) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanRefund) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanRefund) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanRefund) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanRefund) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanRefund) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanRefund) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanRefund) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanRefund) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanRefund) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanRefund) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanRefund) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeApiBeanRefund) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeApiBeanRefund) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeApiBeanRefund) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeApiBeanRefund) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundComment

`func (o *UnibeeApiBeanRefund) GetRefundComment() string`

GetRefundComment returns the RefundComment field if non-nil, zero value otherwise.

### GetRefundCommentOk

`func (o *UnibeeApiBeanRefund) GetRefundCommentOk() (*string, bool)`

GetRefundCommentOk returns a tuple with the RefundComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundComment

`func (o *UnibeeApiBeanRefund) SetRefundComment(v string)`

SetRefundComment sets RefundComment field to given value.

### HasRefundComment

`func (o *UnibeeApiBeanRefund) HasRefundComment() bool`

HasRefundComment returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanRefund) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanRefund) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanRefund) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanRefund) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetRefundTime

`func (o *UnibeeApiBeanRefund) GetRefundTime() int64`

GetRefundTime returns the RefundTime field if non-nil, zero value otherwise.

### GetRefundTimeOk

`func (o *UnibeeApiBeanRefund) GetRefundTimeOk() (*int64, bool)`

GetRefundTimeOk returns a tuple with the RefundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTime

`func (o *UnibeeApiBeanRefund) SetRefundTime(v int64)`

SetRefundTime sets RefundTime field to given value.

### HasRefundTime

`func (o *UnibeeApiBeanRefund) HasRefundTime() bool`

HasRefundTime returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiBeanRefund) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiBeanRefund) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiBeanRefund) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiBeanRefund) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanRefund) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanRefund) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanRefund) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanRefund) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanRefund) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanRefund) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanRefund) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanRefund) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanRefund) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanRefund) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanRefund) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanRefund) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanRefund) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanRefund) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanRefund) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanRefund) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


