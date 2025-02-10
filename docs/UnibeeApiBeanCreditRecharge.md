# UnibeeApiBeanCreditRecharge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **int32** | type of credit account, 1-main recharge account, 2-promo credit account | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**CreditId** | Pointer to **int64** | id of credit account | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | recharge description | [optional] 
**GatewayId** | Pointer to **int64** | payment gateway id | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**InvoiceId** | Pointer to **string** | invoice_id | [optional] 
**Name** | Pointer to **string** | recharge name | [optional] 
**PaidTime** | Pointer to **int64** | paid time | [optional] 
**PaymentAmount** | Pointer to **string** | the payment amount for recharge | [optional] 
**PaymentId** | Pointer to **string** | paymentId | [optional] 
**RechargeId** | Pointer to **string** | unique recharge id for credit account | [optional] 
**RechargeStatus** | Pointer to **int32** | recharge status, 10-in charging，20-recharge success，30-recharge failed | [optional] 
**TotalAmount** | Pointer to **int64** | recharge total amount, cent | [optional] 
**TotalRefundAmount** | Pointer to **int64** | total refund amount,cent | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanCreditRecharge

`func NewUnibeeApiBeanCreditRecharge() *UnibeeApiBeanCreditRecharge`

NewUnibeeApiBeanCreditRecharge instantiates a new UnibeeApiBeanCreditRecharge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreditRechargeWithDefaults

`func NewUnibeeApiBeanCreditRechargeWithDefaults() *UnibeeApiBeanCreditRecharge`

NewUnibeeApiBeanCreditRechargeWithDefaults instantiates a new UnibeeApiBeanCreditRecharge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *UnibeeApiBeanCreditRecharge) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UnibeeApiBeanCreditRecharge) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UnibeeApiBeanCreditRecharge) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UnibeeApiBeanCreditRecharge) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanCreditRecharge) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanCreditRecharge) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanCreditRecharge) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanCreditRecharge) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreditId

`func (o *UnibeeApiBeanCreditRecharge) GetCreditId() int64`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetCreditIdOk() (*int64, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *UnibeeApiBeanCreditRecharge) SetCreditId(v int64)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *UnibeeApiBeanCreditRecharge) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanCreditRecharge) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanCreditRecharge) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanCreditRecharge) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanCreditRecharge) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanCreditRecharge) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanCreditRecharge) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanCreditRecharge) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanCreditRecharge) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanCreditRecharge) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanCreditRecharge) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanCreditRecharge) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanCreditRecharge) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanCreditRecharge) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanCreditRecharge) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanCreditRecharge) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanCreditRecharge) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanCreditRecharge) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanCreditRecharge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanCreditRecharge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanCreditRecharge) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanCreditRecharge) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiBeanCreditRecharge) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiBeanCreditRecharge) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiBeanCreditRecharge) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiBeanCreditRecharge) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *UnibeeApiBeanCreditRecharge) GetPaymentAmount() string`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *UnibeeApiBeanCreditRecharge) GetPaymentAmountOk() (*string, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *UnibeeApiBeanCreditRecharge) SetPaymentAmount(v string)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *UnibeeApiBeanCreditRecharge) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanCreditRecharge) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanCreditRecharge) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanCreditRecharge) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRechargeId

`func (o *UnibeeApiBeanCreditRecharge) GetRechargeId() string`

GetRechargeId returns the RechargeId field if non-nil, zero value otherwise.

### GetRechargeIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetRechargeIdOk() (*string, bool)`

GetRechargeIdOk returns a tuple with the RechargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeId

`func (o *UnibeeApiBeanCreditRecharge) SetRechargeId(v string)`

SetRechargeId sets RechargeId field to given value.

### HasRechargeId

`func (o *UnibeeApiBeanCreditRecharge) HasRechargeId() bool`

HasRechargeId returns a boolean if a field has been set.

### GetRechargeStatus

`func (o *UnibeeApiBeanCreditRecharge) GetRechargeStatus() int32`

GetRechargeStatus returns the RechargeStatus field if non-nil, zero value otherwise.

### GetRechargeStatusOk

`func (o *UnibeeApiBeanCreditRecharge) GetRechargeStatusOk() (*int32, bool)`

GetRechargeStatusOk returns a tuple with the RechargeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeStatus

`func (o *UnibeeApiBeanCreditRecharge) SetRechargeStatus(v int32)`

SetRechargeStatus sets RechargeStatus field to given value.

### HasRechargeStatus

`func (o *UnibeeApiBeanCreditRecharge) HasRechargeStatus() bool`

HasRechargeStatus returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanCreditRecharge) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanCreditRecharge) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanCreditRecharge) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanCreditRecharge) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalRefundAmount

`func (o *UnibeeApiBeanCreditRecharge) GetTotalRefundAmount() int64`

GetTotalRefundAmount returns the TotalRefundAmount field if non-nil, zero value otherwise.

### GetTotalRefundAmountOk

`func (o *UnibeeApiBeanCreditRecharge) GetTotalRefundAmountOk() (*int64, bool)`

GetTotalRefundAmountOk returns a tuple with the TotalRefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRefundAmount

`func (o *UnibeeApiBeanCreditRecharge) SetTotalRefundAmount(v int64)`

SetTotalRefundAmount sets TotalRefundAmount field to given value.

### HasTotalRefundAmount

`func (o *UnibeeApiBeanCreditRecharge) HasTotalRefundAmount() bool`

HasTotalRefundAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanCreditRecharge) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanCreditRecharge) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanCreditRecharge) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanCreditRecharge) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


