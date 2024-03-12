# UnibeeApiBeanPaymentTimelineSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**GatewayId** | Pointer to **int64** | gateway id | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failure | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TimelineType** | Pointer to **int32** | 0-pay, 1-refund | [optional] 
**TotalAmount** | Pointer to **int64** | total amount | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeApiBeanPaymentTimelineSimplify

`func NewUnibeeApiBeanPaymentTimelineSimplify() *UnibeeApiBeanPaymentTimelineSimplify`

NewUnibeeApiBeanPaymentTimelineSimplify instantiates a new UnibeeApiBeanPaymentTimelineSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPaymentTimelineSimplifyWithDefaults

`func NewUnibeeApiBeanPaymentTimelineSimplifyWithDefaults() *UnibeeApiBeanPaymentTimelineSimplify`

NewUnibeeApiBeanPaymentTimelineSimplifyWithDefaults instantiates a new UnibeeApiBeanPaymentTimelineSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTimelineType

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetTimelineType() int32`

GetTimelineType returns the TimelineType field if non-nil, zero value otherwise.

### GetTimelineTypeOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetTimelineTypeOk() (*int32, bool)`

GetTimelineTypeOk returns a tuple with the TimelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineType

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetTimelineType(v int32)`

SetTimelineType sets TimelineType field to given value.

### HasTimelineType

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasTimelineType() bool`

HasTimelineType returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanPaymentTimelineSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanPaymentTimelineSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


