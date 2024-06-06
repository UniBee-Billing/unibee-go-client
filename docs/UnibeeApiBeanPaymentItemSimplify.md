# UnibeeApiBeanPaymentItemSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | amount | [optional] 
**BizType** | Pointer to **int32** | biz_type 1-onetime payment, 3-subscription | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failure | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UnitAmount** | Pointer to **int64** | unit_amount | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeApiBeanPaymentItemSimplify

`func NewUnibeeApiBeanPaymentItemSimplify() *UnibeeApiBeanPaymentItemSimplify`

NewUnibeeApiBeanPaymentItemSimplify instantiates a new UnibeeApiBeanPaymentItemSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPaymentItemSimplifyWithDefaults

`func NewUnibeeApiBeanPaymentItemSimplifyWithDefaults() *UnibeeApiBeanPaymentItemSimplify`

NewUnibeeApiBeanPaymentItemSimplifyWithDefaults instantiates a new UnibeeApiBeanPaymentItemSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBizType

`func (o *UnibeeApiBeanPaymentItemSimplify) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeApiBeanPaymentItemSimplify) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeApiBeanPaymentItemSimplify) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanPaymentItemSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanPaymentItemSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanPaymentItemSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPaymentItemSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPaymentItemSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPaymentItemSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanPaymentItemSimplify) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanPaymentItemSimplify) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanPaymentItemSimplify) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanPaymentItemSimplify) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanPaymentItemSimplify) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanPaymentItemSimplify) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanPaymentItemSimplify) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanPaymentItemSimplify) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanPaymentItemSimplify) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPaymentItemSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPaymentItemSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPaymentItemSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUnitAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUnitAmount() int64`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUnitAmountOk() (*int64, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) SetUnitAmount(v int64)`

SetUnitAmount sets UnitAmount field to given value.

### HasUnitAmount

`func (o *UnibeeApiBeanPaymentItemSimplify) HasUnitAmount() bool`

HasUnitAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanPaymentItemSimplify) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanPaymentItemSimplify) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanPaymentItemSimplify) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


