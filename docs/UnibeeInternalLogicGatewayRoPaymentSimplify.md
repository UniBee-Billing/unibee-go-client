# UnibeeInternalLogicGatewayRoPaymentSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizeReason** | Pointer to **string** |  | [optional] 
**AuthorizeStatus** | Pointer to **int32** | authorize status，0-waiting authorize，1-authorized，2-authorized_request | [optional] 
**Automatic** | Pointer to **int32** |  | [optional] 
**BalanceAmount** | Pointer to **int64** | balance_amount | [optional] 
**BillingReason** | Pointer to **string** |  | [optional] 
**CancelTime** | Pointer to **int64** | cancel time, utc time | [optional] 
**CountryCode** | Pointer to **string** | country code | [optional] 
**CreateTime** | Pointer to **int64** | create time, utc time | [optional] 
**Currency** | Pointer to **string** | currency，“SGD” “MYR” “PHP” “IDR” “THB” | [optional] 
**ExternalPaymentId** | Pointer to **string** | external_payment_id | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**GasPayer** | Pointer to **string** | who pay the gas, merchant|user | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayPaymentId** | Pointer to **string** | gateway_payment_id | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaidTime** | Pointer to **int64** | paid time, utc time | [optional] 
**PaymentAmount** | Pointer to **int64** | payment_amount | [optional] 
**PaymentId** | Pointer to **string** | payment id | [optional] 
**RefundAmount** | Pointer to **int64** | total refund amount | [optional] 
**ReturnUrl** | Pointer to **string** | return url | [optional] 
**Status** | Pointer to **int32** | status  10-pending，20-success，30-failure, 40-cancel | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TotalAmount** | Pointer to **int64** | total amount | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoPaymentSimplify

`func NewUnibeeInternalLogicGatewayRoPaymentSimplify() *UnibeeInternalLogicGatewayRoPaymentSimplify`

NewUnibeeInternalLogicGatewayRoPaymentSimplify instantiates a new UnibeeInternalLogicGatewayRoPaymentSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoPaymentSimplifyWithDefaults

`func NewUnibeeInternalLogicGatewayRoPaymentSimplifyWithDefaults() *UnibeeInternalLogicGatewayRoPaymentSimplify`

NewUnibeeInternalLogicGatewayRoPaymentSimplifyWithDefaults instantiates a new UnibeeInternalLogicGatewayRoPaymentSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAuthorizeReason() string`

GetAuthorizeReason returns the AuthorizeReason field if non-nil, zero value otherwise.

### GetAuthorizeReasonOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAuthorizeReasonOk() (*string, bool)`

GetAuthorizeReasonOk returns a tuple with the AuthorizeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetAuthorizeReason(v string)`

SetAuthorizeReason sets AuthorizeReason field to given value.

### HasAuthorizeReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasAuthorizeReason() bool`

HasAuthorizeReason returns a boolean if a field has been set.

### GetAuthorizeStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAuthorizeStatus() int32`

GetAuthorizeStatus returns the AuthorizeStatus field if non-nil, zero value otherwise.

### GetAuthorizeStatusOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAuthorizeStatusOk() (*int32, bool)`

GetAuthorizeStatusOk returns a tuple with the AuthorizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetAuthorizeStatus(v int32)`

SetAuthorizeStatus sets AuthorizeStatus field to given value.

### HasAuthorizeStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasAuthorizeStatus() bool`

HasAuthorizeStatus returns a boolean if a field has been set.

### GetAutomatic

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAutomatic() int32`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetAutomaticOk() (*int32, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetAutomatic(v int32)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetBalanceAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetBalanceAmount() int64`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetBalanceAmountOk() (*int64, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetBalanceAmount(v int64)`

SetBalanceAmount sets BalanceAmount field to given value.

### HasBalanceAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasBalanceAmount() bool`

HasBalanceAmount returns a boolean if a field has been set.

### GetBillingReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetBillingReason() string`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetBillingReasonOk() (*string, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetBillingReason(v string)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetCancelTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoPaymentSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


