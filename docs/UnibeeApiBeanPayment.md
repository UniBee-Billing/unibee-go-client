# UnibeeApiBeanPayment

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

### NewUnibeeApiBeanPayment

`func NewUnibeeApiBeanPayment() *UnibeeApiBeanPayment`

NewUnibeeApiBeanPayment instantiates a new UnibeeApiBeanPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPaymentWithDefaults

`func NewUnibeeApiBeanPaymentWithDefaults() *UnibeeApiBeanPayment`

NewUnibeeApiBeanPaymentWithDefaults instantiates a new UnibeeApiBeanPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizeReason

`func (o *UnibeeApiBeanPayment) GetAuthorizeReason() string`

GetAuthorizeReason returns the AuthorizeReason field if non-nil, zero value otherwise.

### GetAuthorizeReasonOk

`func (o *UnibeeApiBeanPayment) GetAuthorizeReasonOk() (*string, bool)`

GetAuthorizeReasonOk returns a tuple with the AuthorizeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeReason

`func (o *UnibeeApiBeanPayment) SetAuthorizeReason(v string)`

SetAuthorizeReason sets AuthorizeReason field to given value.

### HasAuthorizeReason

`func (o *UnibeeApiBeanPayment) HasAuthorizeReason() bool`

HasAuthorizeReason returns a boolean if a field has been set.

### GetAuthorizeStatus

`func (o *UnibeeApiBeanPayment) GetAuthorizeStatus() int32`

GetAuthorizeStatus returns the AuthorizeStatus field if non-nil, zero value otherwise.

### GetAuthorizeStatusOk

`func (o *UnibeeApiBeanPayment) GetAuthorizeStatusOk() (*int32, bool)`

GetAuthorizeStatusOk returns a tuple with the AuthorizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeStatus

`func (o *UnibeeApiBeanPayment) SetAuthorizeStatus(v int32)`

SetAuthorizeStatus sets AuthorizeStatus field to given value.

### HasAuthorizeStatus

`func (o *UnibeeApiBeanPayment) HasAuthorizeStatus() bool`

HasAuthorizeStatus returns a boolean if a field has been set.

### GetAutomatic

`func (o *UnibeeApiBeanPayment) GetAutomatic() int32`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *UnibeeApiBeanPayment) GetAutomaticOk() (*int32, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *UnibeeApiBeanPayment) SetAutomatic(v int32)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *UnibeeApiBeanPayment) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetBalanceAmount

`func (o *UnibeeApiBeanPayment) GetBalanceAmount() int64`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *UnibeeApiBeanPayment) GetBalanceAmountOk() (*int64, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *UnibeeApiBeanPayment) SetBalanceAmount(v int64)`

SetBalanceAmount sets BalanceAmount field to given value.

### HasBalanceAmount

`func (o *UnibeeApiBeanPayment) HasBalanceAmount() bool`

HasBalanceAmount returns a boolean if a field has been set.

### GetBillingReason

`func (o *UnibeeApiBeanPayment) GetBillingReason() string`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *UnibeeApiBeanPayment) GetBillingReasonOk() (*string, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *UnibeeApiBeanPayment) SetBillingReason(v string)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *UnibeeApiBeanPayment) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetCancelTime

`func (o *UnibeeApiBeanPayment) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *UnibeeApiBeanPayment) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *UnibeeApiBeanPayment) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *UnibeeApiBeanPayment) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanPayment) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanPayment) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanPayment) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanPayment) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanPayment) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanPayment) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanPayment) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanPayment) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeApiBeanPayment) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeApiBeanPayment) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeApiBeanPayment) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeApiBeanPayment) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeApiBeanPayment) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiBeanPayment) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiBeanPayment) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeApiBeanPayment) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGasPayer

`func (o *UnibeeApiBeanPayment) GetGasPayer() string`

GetGasPayer returns the GasPayer field if non-nil, zero value otherwise.

### GetGasPayerOk

`func (o *UnibeeApiBeanPayment) GetGasPayerOk() (*string, bool)`

GetGasPayerOk returns a tuple with the GasPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPayer

`func (o *UnibeeApiBeanPayment) SetGasPayer(v string)`

SetGasPayer sets GasPayer field to given value.

### HasGasPayer

`func (o *UnibeeApiBeanPayment) HasGasPayer() bool`

HasGasPayer returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanPayment) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanPayment) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanPayment) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanPayment) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeApiBeanPayment) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeApiBeanPayment) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeApiBeanPayment) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeApiBeanPayment) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanPayment) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanPayment) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanPayment) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanPayment) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanPayment) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanPayment) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanPayment) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanPayment) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPayment) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPayment) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPayment) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPayment) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanPayment) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanPayment) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanPayment) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanPayment) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiBeanPayment) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiBeanPayment) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiBeanPayment) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiBeanPayment) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *UnibeeApiBeanPayment) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *UnibeeApiBeanPayment) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *UnibeeApiBeanPayment) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *UnibeeApiBeanPayment) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanPayment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeApiBeanPayment) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeApiBeanPayment) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeApiBeanPayment) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeApiBeanPayment) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiBeanPayment) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiBeanPayment) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiBeanPayment) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiBeanPayment) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPayment) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPayment) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPayment) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanPayment) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanPayment) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanPayment) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanPayment) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanPayment) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanPayment) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanPayment) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanPayment) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanPayment) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanPayment) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanPayment) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanPayment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


