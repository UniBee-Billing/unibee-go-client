# UnibeeApiBeanCreditTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **int32** | type of credit account, 1-main recharge account, 2-promo credit account | [optional] 
**BizId** | Pointer to **string** | business id | [optional] 
**By** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**CreditAmountAfter** | Pointer to **int64** | the credit amount after transaction,cent | [optional] 
**CreditAmountBefore** | Pointer to **int64** | the credit amount before transaction,cent | [optional] 
**CreditId** | Pointer to **int64** | id of credit account | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**DeltaAmount** | Pointer to **int64** | delta amount,cent | [optional] 
**DeltaCurrencyAmount** | Pointer to **int64** | delta currency amount, in cent | [optional] 
**Description** | Pointer to **string** | recharge transaction description | [optional] 
**ExchangeRate** | Pointer to **int64** | ExchangeRate for transaction, keep two decimal places，multiply by 100 saved, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**InvoiceId** | Pointer to **string** | invoice_id | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**Name** | Pointer to **string** | recharge transaction title | [optional] 
**TransactionId** | Pointer to **string** | unique id for timeline | [optional] 
**TransactionType** | Pointer to **int32** | transaction type。1-recharge income，2-payment out，3-refund income，4-withdraw out，5-withdraw failed income, 6-admin change，7-recharge refund out | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanCreditTransaction

`func NewUnibeeApiBeanCreditTransaction() *UnibeeApiBeanCreditTransaction`

NewUnibeeApiBeanCreditTransaction instantiates a new UnibeeApiBeanCreditTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreditTransactionWithDefaults

`func NewUnibeeApiBeanCreditTransactionWithDefaults() *UnibeeApiBeanCreditTransaction`

NewUnibeeApiBeanCreditTransactionWithDefaults instantiates a new UnibeeApiBeanCreditTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *UnibeeApiBeanCreditTransaction) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UnibeeApiBeanCreditTransaction) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UnibeeApiBeanCreditTransaction) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UnibeeApiBeanCreditTransaction) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBizId

`func (o *UnibeeApiBeanCreditTransaction) GetBizId() string`

GetBizId returns the BizId field if non-nil, zero value otherwise.

### GetBizIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetBizIdOk() (*string, bool)`

GetBizIdOk returns a tuple with the BizId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizId

`func (o *UnibeeApiBeanCreditTransaction) SetBizId(v string)`

SetBizId sets BizId field to given value.

### HasBizId

`func (o *UnibeeApiBeanCreditTransaction) HasBizId() bool`

HasBizId returns a boolean if a field has been set.

### GetBy

`func (o *UnibeeApiBeanCreditTransaction) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *UnibeeApiBeanCreditTransaction) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *UnibeeApiBeanCreditTransaction) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *UnibeeApiBeanCreditTransaction) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanCreditTransaction) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanCreditTransaction) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanCreditTransaction) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanCreditTransaction) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreditAmountAfter

`func (o *UnibeeApiBeanCreditTransaction) GetCreditAmountAfter() int64`

GetCreditAmountAfter returns the CreditAmountAfter field if non-nil, zero value otherwise.

### GetCreditAmountAfterOk

`func (o *UnibeeApiBeanCreditTransaction) GetCreditAmountAfterOk() (*int64, bool)`

GetCreditAmountAfterOk returns a tuple with the CreditAmountAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountAfter

`func (o *UnibeeApiBeanCreditTransaction) SetCreditAmountAfter(v int64)`

SetCreditAmountAfter sets CreditAmountAfter field to given value.

### HasCreditAmountAfter

`func (o *UnibeeApiBeanCreditTransaction) HasCreditAmountAfter() bool`

HasCreditAmountAfter returns a boolean if a field has been set.

### GetCreditAmountBefore

`func (o *UnibeeApiBeanCreditTransaction) GetCreditAmountBefore() int64`

GetCreditAmountBefore returns the CreditAmountBefore field if non-nil, zero value otherwise.

### GetCreditAmountBeforeOk

`func (o *UnibeeApiBeanCreditTransaction) GetCreditAmountBeforeOk() (*int64, bool)`

GetCreditAmountBeforeOk returns a tuple with the CreditAmountBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountBefore

`func (o *UnibeeApiBeanCreditTransaction) SetCreditAmountBefore(v int64)`

SetCreditAmountBefore sets CreditAmountBefore field to given value.

### HasCreditAmountBefore

`func (o *UnibeeApiBeanCreditTransaction) HasCreditAmountBefore() bool`

HasCreditAmountBefore returns a boolean if a field has been set.

### GetCreditId

`func (o *UnibeeApiBeanCreditTransaction) GetCreditId() int64`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetCreditIdOk() (*int64, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *UnibeeApiBeanCreditTransaction) SetCreditId(v int64)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *UnibeeApiBeanCreditTransaction) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanCreditTransaction) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanCreditTransaction) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanCreditTransaction) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanCreditTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeltaAmount

`func (o *UnibeeApiBeanCreditTransaction) GetDeltaAmount() int64`

GetDeltaAmount returns the DeltaAmount field if non-nil, zero value otherwise.

### GetDeltaAmountOk

`func (o *UnibeeApiBeanCreditTransaction) GetDeltaAmountOk() (*int64, bool)`

GetDeltaAmountOk returns a tuple with the DeltaAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaAmount

`func (o *UnibeeApiBeanCreditTransaction) SetDeltaAmount(v int64)`

SetDeltaAmount sets DeltaAmount field to given value.

### HasDeltaAmount

`func (o *UnibeeApiBeanCreditTransaction) HasDeltaAmount() bool`

HasDeltaAmount returns a boolean if a field has been set.

### GetDeltaCurrencyAmount

`func (o *UnibeeApiBeanCreditTransaction) GetDeltaCurrencyAmount() int64`

GetDeltaCurrencyAmount returns the DeltaCurrencyAmount field if non-nil, zero value otherwise.

### GetDeltaCurrencyAmountOk

`func (o *UnibeeApiBeanCreditTransaction) GetDeltaCurrencyAmountOk() (*int64, bool)`

GetDeltaCurrencyAmountOk returns a tuple with the DeltaCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCurrencyAmount

`func (o *UnibeeApiBeanCreditTransaction) SetDeltaCurrencyAmount(v int64)`

SetDeltaCurrencyAmount sets DeltaCurrencyAmount field to given value.

### HasDeltaCurrencyAmount

`func (o *UnibeeApiBeanCreditTransaction) HasDeltaCurrencyAmount() bool`

HasDeltaCurrencyAmount returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanCreditTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanCreditTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanCreditTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanCreditTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanCreditTransaction) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanCreditTransaction) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanCreditTransaction) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanCreditTransaction) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanCreditTransaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanCreditTransaction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanCreditTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanCreditTransaction) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanCreditTransaction) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanCreditTransaction) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanCreditTransaction) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanCreditTransaction) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanCreditTransaction) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanCreditTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanCreditTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanCreditTransaction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanCreditTransaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransactionId

`func (o *UnibeeApiBeanCreditTransaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *UnibeeApiBeanCreditTransaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *UnibeeApiBeanCreditTransaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *UnibeeApiBeanCreditTransaction) GetTransactionType() int32`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *UnibeeApiBeanCreditTransaction) GetTransactionTypeOk() (*int32, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *UnibeeApiBeanCreditTransaction) SetTransactionType(v int32)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *UnibeeApiBeanCreditTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanCreditTransaction) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanCreditTransaction) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanCreditTransaction) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanCreditTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


