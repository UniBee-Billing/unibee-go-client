# UnibeeApiBeanDetailSubscriptionNextInvoiceApply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCreditAmount** | Pointer to **int64** | Promo credit amount scheduled to apply on the next invoice, if any. | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailSubscriptionNextInvoiceApply

`func NewUnibeeApiBeanDetailSubscriptionNextInvoiceApply() *UnibeeApiBeanDetailSubscriptionNextInvoiceApply`

NewUnibeeApiBeanDetailSubscriptionNextInvoiceApply instantiates a new UnibeeApiBeanDetailSubscriptionNextInvoiceApply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailSubscriptionNextInvoiceApplyWithDefaults

`func NewUnibeeApiBeanDetailSubscriptionNextInvoiceApplyWithDefaults() *UnibeeApiBeanDetailSubscriptionNextInvoiceApply`

NewUnibeeApiBeanDetailSubscriptionNextInvoiceApplyWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionNextInvoiceApply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCreditAmount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) GetApplyPromoCreditAmount() int64`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) GetApplyPromoCreditAmountOk() (*int64, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) SetApplyPromoCreditAmount(v int64)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanDetailSubscriptionNextInvoiceApply) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


