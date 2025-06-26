# UnibeeApiBeanInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**AutoCharge** | Pointer to **bool** |  | [optional] 
**BillingCycleAnchor** | Pointer to **int64** | billing_cycle_anchor | [optional] 
**BizType** | Pointer to **int32** | biz type from payment 1-onetime payment, 3-subscription | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreateFrom** | Pointer to **string** | create from | [optional] 
**CreditAccount** | Pointer to [**UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) |  | [optional] 
**CreditPayout** | Pointer to [**UnibeeApiBeanCreditPayout**](UnibeeApiBeanCreditPayout.md) |  | [optional] 
**CryptoAmount** | Pointer to **int64** | crypto_amount, cent | [optional] 
**CryptoCurrency** | Pointer to **string** | crypto_currency | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DayUtilDue** | Pointer to **int64** | day util due after finish | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**DiscountCode** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceName** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanInvoiceItemSimplify**](UnibeeApiBeanInvoiceItemSimplify.md) |  | [optional] 
**Link** | Pointer to **string** | invoice link | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**PartialCreditPaidAmount** | Pointer to **int64** | partial credit paid amount | [optional] 
**PaymentId** | Pointer to **string** | paymentId | [optional] 
**PaymentLink** | Pointer to **string** | invoice payment link | [optional] 
**PaymentType** | Pointer to **string** |  | [optional] 
**PeriodEnd** | Pointer to **int64** |  | [optional] 
**PeriodStart** | Pointer to **int64** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**PromoCreditAccount** | Pointer to [**UnibeeApiBeanCreditAccount**](UnibeeApiBeanCreditAccount.md) |  | [optional] 
**PromoCreditDiscountAmount** | Pointer to **int64** | promo credit discount amount | [optional] 
**PromoCreditPayout** | Pointer to [**UnibeeApiBeanCreditPayout**](UnibeeApiBeanCreditPayout.md) |  | [optional] 
**PromoCreditTransaction** | Pointer to [**UnibeeApiBeanCreditTransaction**](UnibeeApiBeanCreditTransaction.md) |  | [optional] 
**ProrationDate** | Pointer to **int64** |  | [optional] 
**ProrationScale** | Pointer to **int64** |  | [optional] 
**RefundId** | Pointer to **string** | refundId | [optional] 
**SendNote** | Pointer to **string** | send_note | [optional] 
**SendStatus** | Pointer to **int32** | email send status，0-No | 1- YES| 2-Unnecessary | [optional] 
**Status** | Pointer to **int32** | status，1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**SubscriptionAmount** | Pointer to **int64** |  | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 
**TaxAmount** | Pointer to **int64** |  | [optional] 
**TaxPercentage** | Pointer to **int64** | TaxPercentage，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 
**UserMetricChargeForInvoice** | Pointer to [**UnibeeApiBeanUserMetricChargeInvoiceItemEntity**](UnibeeApiBeanUserMetricChargeInvoiceItemEntity.md) |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiBeanInvoice

`func NewUnibeeApiBeanInvoice() *UnibeeApiBeanInvoice`

NewUnibeeApiBeanInvoice instantiates a new UnibeeApiBeanInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanInvoiceWithDefaults

`func NewUnibeeApiBeanInvoiceWithDefaults() *UnibeeApiBeanInvoice`

NewUnibeeApiBeanInvoiceWithDefaults instantiates a new UnibeeApiBeanInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *UnibeeApiBeanInvoice) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiBeanInvoice) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiBeanInvoice) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiBeanInvoice) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetAutoCharge

`func (o *UnibeeApiBeanInvoice) GetAutoCharge() bool`

GetAutoCharge returns the AutoCharge field if non-nil, zero value otherwise.

### GetAutoChargeOk

`func (o *UnibeeApiBeanInvoice) GetAutoChargeOk() (*bool, bool)`

GetAutoChargeOk returns a tuple with the AutoCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCharge

`func (o *UnibeeApiBeanInvoice) SetAutoCharge(v bool)`

SetAutoCharge sets AutoCharge field to given value.

### HasAutoCharge

`func (o *UnibeeApiBeanInvoice) HasAutoCharge() bool`

HasAutoCharge returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *UnibeeApiBeanInvoice) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeApiBeanInvoice) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeApiBeanInvoice) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeApiBeanInvoice) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetBizType

`func (o *UnibeeApiBeanInvoice) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeApiBeanInvoice) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeApiBeanInvoice) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeApiBeanInvoice) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanInvoice) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanInvoice) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanInvoice) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanInvoice) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateFrom

`func (o *UnibeeApiBeanInvoice) GetCreateFrom() string`

GetCreateFrom returns the CreateFrom field if non-nil, zero value otherwise.

### GetCreateFromOk

`func (o *UnibeeApiBeanInvoice) GetCreateFromOk() (*string, bool)`

GetCreateFromOk returns a tuple with the CreateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFrom

`func (o *UnibeeApiBeanInvoice) SetCreateFrom(v string)`

SetCreateFrom sets CreateFrom field to given value.

### HasCreateFrom

`func (o *UnibeeApiBeanInvoice) HasCreateFrom() bool`

HasCreateFrom returns a boolean if a field has been set.

### GetCreditAccount

`func (o *UnibeeApiBeanInvoice) GetCreditAccount() UnibeeApiBeanCreditAccount`

GetCreditAccount returns the CreditAccount field if non-nil, zero value otherwise.

### GetCreditAccountOk

`func (o *UnibeeApiBeanInvoice) GetCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool)`

GetCreditAccountOk returns a tuple with the CreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAccount

`func (o *UnibeeApiBeanInvoice) SetCreditAccount(v UnibeeApiBeanCreditAccount)`

SetCreditAccount sets CreditAccount field to given value.

### HasCreditAccount

`func (o *UnibeeApiBeanInvoice) HasCreditAccount() bool`

HasCreditAccount returns a boolean if a field has been set.

### GetCreditPayout

`func (o *UnibeeApiBeanInvoice) GetCreditPayout() UnibeeApiBeanCreditPayout`

GetCreditPayout returns the CreditPayout field if non-nil, zero value otherwise.

### GetCreditPayoutOk

`func (o *UnibeeApiBeanInvoice) GetCreditPayoutOk() (*UnibeeApiBeanCreditPayout, bool)`

GetCreditPayoutOk returns a tuple with the CreditPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPayout

`func (o *UnibeeApiBeanInvoice) SetCreditPayout(v UnibeeApiBeanCreditPayout)`

SetCreditPayout sets CreditPayout field to given value.

### HasCreditPayout

`func (o *UnibeeApiBeanInvoice) HasCreditPayout() bool`

HasCreditPayout returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *UnibeeApiBeanInvoice) GetCryptoAmount() int64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *UnibeeApiBeanInvoice) GetCryptoAmountOk() (*int64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *UnibeeApiBeanInvoice) SetCryptoAmount(v int64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *UnibeeApiBeanInvoice) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoCurrency

`func (o *UnibeeApiBeanInvoice) GetCryptoCurrency() string`

GetCryptoCurrency returns the CryptoCurrency field if non-nil, zero value otherwise.

### GetCryptoCurrencyOk

`func (o *UnibeeApiBeanInvoice) GetCryptoCurrencyOk() (*string, bool)`

GetCryptoCurrencyOk returns a tuple with the CryptoCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrency

`func (o *UnibeeApiBeanInvoice) SetCryptoCurrency(v string)`

SetCryptoCurrency sets CryptoCurrency field to given value.

### HasCryptoCurrency

`func (o *UnibeeApiBeanInvoice) HasCryptoCurrency() bool`

HasCryptoCurrency returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeApiBeanInvoice) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiBeanInvoice) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiBeanInvoice) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeApiBeanInvoice) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDayUtilDue

`func (o *UnibeeApiBeanInvoice) GetDayUtilDue() int64`

GetDayUtilDue returns the DayUtilDue field if non-nil, zero value otherwise.

### GetDayUtilDueOk

`func (o *UnibeeApiBeanInvoice) GetDayUtilDueOk() (*int64, bool)`

GetDayUtilDueOk returns a tuple with the DayUtilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayUtilDue

`func (o *UnibeeApiBeanInvoice) SetDayUtilDue(v int64)`

SetDayUtilDue sets DayUtilDue field to given value.

### HasDayUtilDue

`func (o *UnibeeApiBeanInvoice) HasDayUtilDue() bool`

HasDayUtilDue returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanInvoice) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanInvoice) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanInvoice) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanInvoice) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanInvoice) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanInvoice) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanInvoice) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiBeanInvoice) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiBeanInvoice) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiBeanInvoice) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiBeanInvoice) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanInvoice) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanInvoice) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanInvoice) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanInvoice) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanInvoice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanInvoice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanInvoice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanInvoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanInvoice) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanInvoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanInvoice) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiBeanInvoice) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiBeanInvoice) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiBeanInvoice) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiBeanInvoice) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanInvoice) GetLines() []UnibeeApiBeanInvoiceItemSimplify`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanInvoice) GetLinesOk() (*[]UnibeeApiBeanInvoiceItemSimplify, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanInvoice) SetLines(v []UnibeeApiBeanInvoiceItemSimplify)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanInvoice) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanInvoice) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanInvoice) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanInvoice) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanInvoice) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanInvoice) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanInvoice) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanInvoice) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanInvoice) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiBeanInvoice) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiBeanInvoice) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiBeanInvoice) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiBeanInvoice) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetPartialCreditPaidAmount

`func (o *UnibeeApiBeanInvoice) GetPartialCreditPaidAmount() int64`

GetPartialCreditPaidAmount returns the PartialCreditPaidAmount field if non-nil, zero value otherwise.

### GetPartialCreditPaidAmountOk

`func (o *UnibeeApiBeanInvoice) GetPartialCreditPaidAmountOk() (*int64, bool)`

GetPartialCreditPaidAmountOk returns a tuple with the PartialCreditPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialCreditPaidAmount

`func (o *UnibeeApiBeanInvoice) SetPartialCreditPaidAmount(v int64)`

SetPartialCreditPaidAmount sets PartialCreditPaidAmount field to given value.

### HasPartialCreditPaidAmount

`func (o *UnibeeApiBeanInvoice) HasPartialCreditPaidAmount() bool`

HasPartialCreditPaidAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanInvoice) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanInvoice) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanInvoice) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanInvoice) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentLink

`func (o *UnibeeApiBeanInvoice) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeApiBeanInvoice) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeApiBeanInvoice) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *UnibeeApiBeanInvoice) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetPaymentType

`func (o *UnibeeApiBeanInvoice) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *UnibeeApiBeanInvoice) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *UnibeeApiBeanInvoice) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *UnibeeApiBeanInvoice) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanInvoice) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanInvoice) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanInvoice) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanInvoice) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanInvoice) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanInvoice) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanInvoice) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanInvoice) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanInvoice) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanInvoice) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanInvoice) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanInvoice) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetPromoCreditAccount

`func (o *UnibeeApiBeanInvoice) GetPromoCreditAccount() UnibeeApiBeanCreditAccount`

GetPromoCreditAccount returns the PromoCreditAccount field if non-nil, zero value otherwise.

### GetPromoCreditAccountOk

`func (o *UnibeeApiBeanInvoice) GetPromoCreditAccountOk() (*UnibeeApiBeanCreditAccount, bool)`

GetPromoCreditAccountOk returns a tuple with the PromoCreditAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditAccount

`func (o *UnibeeApiBeanInvoice) SetPromoCreditAccount(v UnibeeApiBeanCreditAccount)`

SetPromoCreditAccount sets PromoCreditAccount field to given value.

### HasPromoCreditAccount

`func (o *UnibeeApiBeanInvoice) HasPromoCreditAccount() bool`

HasPromoCreditAccount returns a boolean if a field has been set.

### GetPromoCreditDiscountAmount

`func (o *UnibeeApiBeanInvoice) GetPromoCreditDiscountAmount() int64`

GetPromoCreditDiscountAmount returns the PromoCreditDiscountAmount field if non-nil, zero value otherwise.

### GetPromoCreditDiscountAmountOk

`func (o *UnibeeApiBeanInvoice) GetPromoCreditDiscountAmountOk() (*int64, bool)`

GetPromoCreditDiscountAmountOk returns a tuple with the PromoCreditDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditDiscountAmount

`func (o *UnibeeApiBeanInvoice) SetPromoCreditDiscountAmount(v int64)`

SetPromoCreditDiscountAmount sets PromoCreditDiscountAmount field to given value.

### HasPromoCreditDiscountAmount

`func (o *UnibeeApiBeanInvoice) HasPromoCreditDiscountAmount() bool`

HasPromoCreditDiscountAmount returns a boolean if a field has been set.

### GetPromoCreditPayout

`func (o *UnibeeApiBeanInvoice) GetPromoCreditPayout() UnibeeApiBeanCreditPayout`

GetPromoCreditPayout returns the PromoCreditPayout field if non-nil, zero value otherwise.

### GetPromoCreditPayoutOk

`func (o *UnibeeApiBeanInvoice) GetPromoCreditPayoutOk() (*UnibeeApiBeanCreditPayout, bool)`

GetPromoCreditPayoutOk returns a tuple with the PromoCreditPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditPayout

`func (o *UnibeeApiBeanInvoice) SetPromoCreditPayout(v UnibeeApiBeanCreditPayout)`

SetPromoCreditPayout sets PromoCreditPayout field to given value.

### HasPromoCreditPayout

`func (o *UnibeeApiBeanInvoice) HasPromoCreditPayout() bool`

HasPromoCreditPayout returns a boolean if a field has been set.

### GetPromoCreditTransaction

`func (o *UnibeeApiBeanInvoice) GetPromoCreditTransaction() UnibeeApiBeanCreditTransaction`

GetPromoCreditTransaction returns the PromoCreditTransaction field if non-nil, zero value otherwise.

### GetPromoCreditTransactionOk

`func (o *UnibeeApiBeanInvoice) GetPromoCreditTransactionOk() (*UnibeeApiBeanCreditTransaction, bool)`

GetPromoCreditTransactionOk returns a tuple with the PromoCreditTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditTransaction

`func (o *UnibeeApiBeanInvoice) SetPromoCreditTransaction(v UnibeeApiBeanCreditTransaction)`

SetPromoCreditTransaction sets PromoCreditTransaction field to given value.

### HasPromoCreditTransaction

`func (o *UnibeeApiBeanInvoice) HasPromoCreditTransaction() bool`

HasPromoCreditTransaction returns a boolean if a field has been set.

### GetProrationDate

`func (o *UnibeeApiBeanInvoice) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeApiBeanInvoice) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeApiBeanInvoice) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeApiBeanInvoice) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetProrationScale

`func (o *UnibeeApiBeanInvoice) GetProrationScale() int64`

GetProrationScale returns the ProrationScale field if non-nil, zero value otherwise.

### GetProrationScaleOk

`func (o *UnibeeApiBeanInvoice) GetProrationScaleOk() (*int64, bool)`

GetProrationScaleOk returns a tuple with the ProrationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationScale

`func (o *UnibeeApiBeanInvoice) SetProrationScale(v int64)`

SetProrationScale sets ProrationScale field to given value.

### HasProrationScale

`func (o *UnibeeApiBeanInvoice) HasProrationScale() bool`

HasProrationScale returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanInvoice) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanInvoice) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanInvoice) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanInvoice) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeApiBeanInvoice) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeApiBeanInvoice) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeApiBeanInvoice) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeApiBeanInvoice) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeApiBeanInvoice) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeApiBeanInvoice) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeApiBeanInvoice) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeApiBeanInvoice) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanInvoice) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanInvoice) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanInvoice) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeApiBeanInvoice) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeApiBeanInvoice) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeApiBeanInvoice) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeApiBeanInvoice) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoice) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanInvoice) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanInvoice) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanInvoice) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanInvoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiBeanInvoice) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiBeanInvoice) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiBeanInvoice) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiBeanInvoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanInvoice) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanInvoice) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanInvoice) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanInvoice) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanInvoice) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanInvoice) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanInvoice) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoice) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoice) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiBeanInvoice) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiBeanInvoice) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiBeanInvoice) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiBeanInvoice) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanInvoice) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanInvoice) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanInvoice) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanInvoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserMetricChargeForInvoice

`func (o *UnibeeApiBeanInvoice) GetUserMetricChargeForInvoice() UnibeeApiBeanUserMetricChargeInvoiceItemEntity`

GetUserMetricChargeForInvoice returns the UserMetricChargeForInvoice field if non-nil, zero value otherwise.

### GetUserMetricChargeForInvoiceOk

`func (o *UnibeeApiBeanInvoice) GetUserMetricChargeForInvoiceOk() (*UnibeeApiBeanUserMetricChargeInvoiceItemEntity, bool)`

GetUserMetricChargeForInvoiceOk returns a tuple with the UserMetricChargeForInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMetricChargeForInvoice

`func (o *UnibeeApiBeanInvoice) SetUserMetricChargeForInvoice(v UnibeeApiBeanUserMetricChargeInvoiceItemEntity)`

SetUserMetricChargeForInvoice sets UserMetricChargeForInvoice field to given value.

### HasUserMetricChargeForInvoice

`func (o *UnibeeApiBeanInvoice) HasUserMetricChargeForInvoice() bool`

HasUserMetricChargeForInvoice returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanInvoice) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanInvoice) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanInvoice) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanInvoice) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


