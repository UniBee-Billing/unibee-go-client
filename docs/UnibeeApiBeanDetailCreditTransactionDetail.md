# UnibeeApiBeanDetailCreditTransactionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **int32** | type of credit account, 1-main recharge account, 2-promo credit account | [optional] 
**AdminMember** | Pointer to [**UnibeeApiBeanMerchantMember**](UnibeeApiBeanMerchantMember.md) |  | [optional] 
**BizId** | Pointer to **string** | business id | [optional] 
**By** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**CreditAccount** | Pointer to [**UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) |  | [optional] 
**CreditAmountAfter** | Pointer to **int64** | the credit amount after transaction,cent | [optional] 
**CreditAmountBefore** | Pointer to **int64** | the credit amount before transaction,cent | [optional] 
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
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailCreditTransactionDetail

`func NewUnibeeApiBeanDetailCreditTransactionDetail() *UnibeeApiBeanDetailCreditTransactionDetail`

NewUnibeeApiBeanDetailCreditTransactionDetail instantiates a new UnibeeApiBeanDetailCreditTransactionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailCreditTransactionDetailWithDefaults

`func NewUnibeeApiBeanDetailCreditTransactionDetailWithDefaults() *UnibeeApiBeanDetailCreditTransactionDetail`

NewUnibeeApiBeanDetailCreditTransactionDetailWithDefaults instantiates a new UnibeeApiBeanDetailCreditTransactionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAccountType() int32`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAccountTypeOk() (*int32, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetAccountType(v int32)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminMember

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAdminMember() UnibeeApiBeanMerchantMember`

GetAdminMember returns the AdminMember field if non-nil, zero value otherwise.

### GetAdminMemberOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetAdminMemberOk() (*UnibeeApiBeanMerchantMember, bool)`

GetAdminMemberOk returns a tuple with the AdminMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminMember

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetAdminMember(v UnibeeApiBeanMerchantMember)`

SetAdminMember sets AdminMember field to given value.

### HasAdminMember

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasAdminMember() bool`

HasAdminMember returns a boolean if a field has been set.

### GetBizId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBizId() string`

GetBizId returns the BizId field if non-nil, zero value otherwise.

### GetBizIdOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBizIdOk() (*string, bool)`

GetBizIdOk returns a tuple with the BizId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetBizId(v string)`

SetBizId sets BizId field to given value.

### HasBizId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasBizId() bool`

HasBizId returns a boolean if a field has been set.

### GetBy

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetBy() string`

GetBy returns the By field if non-nil, zero value otherwise.

### GetByOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetByOk() (*string, bool)`

GetByOk returns a tuple with the By field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBy

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetBy(v string)`

SetBy sets By field to given value.

### HasBy

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasBy() bool`

HasBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreditAccount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAccount() UnibeeApiBeanCreditAccount`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAccount(v UnibeeApiBeanCreditAccount)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetCreditAmountAfter

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountAfter() int64`

GetCreditAmountAfter returns the CreditAmountAfter field if non-nil, zero value otherwise.

### GetCreditAmountAfterOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountAfterOk() (*int64, bool)`

GetCreditAmountAfterOk returns a tuple with the CreditAmountAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountAfter

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAmountAfter(v int64)`

SetCreditAmountAfter sets CreditAmountAfter field to given value.

### HasCreditAmountAfter

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAmountAfter() bool`

HasCreditAmountAfter returns a boolean if a field has been set.

### GetCreditAmountBefore

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountBefore() int64`

GetCreditAmountBefore returns the CreditAmountBefore field if non-nil, zero value otherwise.

### GetCreditAmountBeforeOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCreditAmountBeforeOk() (*int64, bool)`

GetCreditAmountBeforeOk returns a tuple with the CreditAmountBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountBefore

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCreditAmountBefore(v int64)`

SetCreditAmountBefore sets CreditAmountBefore field to given value.

### HasCreditAmountBefore

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCreditAmountBefore() bool`

HasCreditAmountBefore returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDeltaAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaAmount() int64`

GetDeltaAmount returns the DeltaAmount field if non-nil, zero value otherwise.

### GetDeltaAmountOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaAmountOk() (*int64, bool)`

GetDeltaAmountOk returns a tuple with the DeltaAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDeltaAmount(v int64)`

SetDeltaAmount sets DeltaAmount field to given value.

### HasDeltaAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDeltaAmount() bool`

HasDeltaAmount returns a boolean if a field has been set.

### GetDeltaCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaCurrencyAmount() int64`

GetDeltaCurrencyAmount returns the DeltaCurrencyAmount field if non-nil, zero value otherwise.

### GetDeltaCurrencyAmountOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDeltaCurrencyAmountOk() (*int64, bool)`

GetDeltaCurrencyAmountOk returns a tuple with the DeltaCurrencyAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDeltaCurrencyAmount(v int64)`

SetDeltaCurrencyAmount sets DeltaCurrencyAmount field to given value.

### HasDeltaCurrencyAmount

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDeltaCurrencyAmount() bool`

HasDeltaCurrencyAmount returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTransactionId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionType() int32`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetTransactionTypeOk() (*int32, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetTransactionType(v int32)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailCreditTransactionDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


