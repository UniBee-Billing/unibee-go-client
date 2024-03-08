# UnibeeInternalLogicGatewayRoRefundSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | country code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**ExternalRefundId** | Pointer to **string** | external_refund_id | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayRefundId** | Pointer to **string** | gateway refund id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**PaymentId** | Pointer to **string** | relative payment id | [optional] 
**RefundAmount** | Pointer to **int64** | refund amount, cent | [optional] 
**RefundComment** | Pointer to **string** | refund comment | [optional] 
**RefundId** | Pointer to **string** | refund id (system generate) | [optional] 
**RefundTime** | Pointer to **int64** | refund success time | [optional] 
**ReturnUrl** | Pointer to **string** | return url after refund success | [optional] 
**Status** | Pointer to **int32** | status。10-pending，20-success，30-failure, 40-cancel | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoRefundSimplify

`func NewUnibeeInternalLogicGatewayRoRefundSimplify() *UnibeeInternalLogicGatewayRoRefundSimplify`

NewUnibeeInternalLogicGatewayRoRefundSimplify instantiates a new UnibeeInternalLogicGatewayRoRefundSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoRefundSimplifyWithDefaults

`func NewUnibeeInternalLogicGatewayRoRefundSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoRefundSimplify`

NewUnibeeInternalLogicGatewayRoRefundSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoRefundSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetExternalRefundId() string`

GetExternalRefundId returns the ExternalRefundId field if non-nil, zero value otherwise.

### GetExternalRefundIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetExternalRefundIdOk() (*string, bool)`

GetExternalRefundIdOk returns a tuple with the ExternalRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetExternalRefundId(v string)`

SetExternalRefundId sets ExternalRefundId field to given value.

### HasExternalRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasExternalRefundId() bool`

HasExternalRefundId returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetGatewayRefundId() string`

GetGatewayRefundId returns the GatewayRefundId field if non-nil, zero value otherwise.

### GetGatewayRefundIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetGatewayRefundIdOk() (*string, bool)`

GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetGatewayRefundId(v string)`

SetGatewayRefundId sets GatewayRefundId field to given value.

### HasGatewayRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasGatewayRefundId() bool`

HasGatewayRefundId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundComment

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundComment() string`

GetRefundComment returns the RefundComment field if non-nil, zero value otherwise.

### GetRefundCommentOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundCommentOk() (*string, bool)`

GetRefundCommentOk returns a tuple with the RefundComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundComment

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetRefundComment(v string)`

SetRefundComment sets RefundComment field to given value.

### HasRefundComment

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasRefundComment() bool`

HasRefundComment returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetRefundTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundTime() int64`

GetRefundTime returns the RefundTime field if non-nil, zero value otherwise.

### GetRefundTimeOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetRefundTimeOk() (*int64, bool)`

GetRefundTimeOk returns a tuple with the RefundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetRefundTime(v int64)`

SetRefundTime sets RefundTime field to given value.

### HasRefundTime

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasRefundTime() bool`

HasRefundTime returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoRefundSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


