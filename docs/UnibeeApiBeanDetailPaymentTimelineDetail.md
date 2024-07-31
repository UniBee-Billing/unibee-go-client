# UnibeeApiBeanDetailPaymentTimelineDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**ExternalTransactionId** | Pointer to **string** | ExternalTransactionId | [optional] 
**FullRefund** | Pointer to **int32** | 0-no, 1-yes | [optional] 
**GatewayId** | Pointer to **int64** | gateway id | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**Refund** | Pointer to [**UnibeeApiBeanRefund**](UnibeeApiBeanRefund.md) |  | [optional] 
**RefundId** | Pointer to **string** | refund id | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failure | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TimelineType** | Pointer to **int32** | 0-pay, 1-refund | [optional] 
**TotalAmount** | Pointer to **int64** | total amount | [optional] 
**TransactionId** | Pointer to **string** | TransactionId | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeApiBeanDetailPaymentTimelineDetail

`func NewUnibeeApiBeanDetailPaymentTimelineDetail() *UnibeeApiBeanDetailPaymentTimelineDetail`

NewUnibeeApiBeanDetailPaymentTimelineDetail instantiates a new UnibeeApiBeanDetailPaymentTimelineDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailPaymentTimelineDetailWithDefaults

`func NewUnibeeApiBeanDetailPaymentTimelineDetailWithDefaults() *UnibeeApiBeanDetailPaymentTimelineDetail`

NewUnibeeApiBeanDetailPaymentTimelineDetailWithDefaults instantiates a new UnibeeApiBeanDetailPaymentTimelineDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetExternalTransactionId() string`

GetExternalTransactionId returns the ExternalTransactionId field if non-nil, zero value otherwise.

### GetExternalTransactionIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetExternalTransactionIdOk() (*string, bool)`

GetExternalTransactionIdOk returns a tuple with the ExternalTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetExternalTransactionId(v string)`

SetExternalTransactionId sets ExternalTransactionId field to given value.

### HasExternalTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasExternalTransactionId() bool`

HasExternalTransactionId returns a boolean if a field has been set.

### GetFullRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetFullRefund() int32`

GetFullRefund returns the FullRefund field if non-nil, zero value otherwise.

### GetFullRefundOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetFullRefundOk() (*int32, bool)`

GetFullRefundOk returns a tuple with the FullRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetFullRefund(v int32)`

SetFullRefund sets FullRefund field to given value.

### HasFullRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasFullRefund() bool`

HasFullRefund returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefund() UnibeeApiBeanRefund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundOk() (*UnibeeApiBeanRefund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetRefund(v UnibeeApiBeanRefund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTimelineType

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTimelineType() int32`

GetTimelineType returns the TimelineType field if non-nil, zero value otherwise.

### GetTimelineTypeOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTimelineTypeOk() (*int32, bool)`

GetTimelineTypeOk returns a tuple with the TimelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineType

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTimelineType(v int32)`

SetTimelineType sets TimelineType field to given value.

### HasTimelineType

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTimelineType() bool`

HasTimelineType returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailPaymentTimelineDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


