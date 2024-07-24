# UnibeeApiBeanInvoiceSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleAnchor** | Pointer to **int64** | billing_cycle_anchor | [optional] 
**BizType** | Pointer to **int32** | biz type from payment 1-onetime payment, 3-subscription | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreateFrom** | Pointer to **string** | create from | [optional] 
**CryptoAmount** | Pointer to **int64** | crypto_amount, cent | [optional] 
**CryptoCurrency** | Pointer to **string** | crypto_currency | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DayUtilDue** | Pointer to **int64** | day util due after finish | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCodeSimplify**](UnibeeApiBeanMerchantDiscountCodeSimplify.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**DiscountCode** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceName** | Pointer to **string** |  | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanInvoiceItemSimplify**](UnibeeApiBeanInvoiceItemSimplify.md) |  | [optional] 
**Link** | Pointer to **string** | invoice link | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**PaymentId** | Pointer to **string** | paymentId | [optional] 
**PaymentLink** | Pointer to **string** | invoice payment link | [optional] 
**PeriodEnd** | Pointer to **int64** |  | [optional] 
**PeriodStart** | Pointer to **int64** |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
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
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiBeanInvoiceSimplify

`func NewUnibeeApiBeanInvoiceSimplify() *UnibeeApiBeanInvoiceSimplify`

NewUnibeeApiBeanInvoiceSimplify instantiates a new UnibeeApiBeanInvoiceSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanInvoiceSimplifyWithDefaults

`func NewUnibeeApiBeanInvoiceSimplifyWithDefaults() *UnibeeApiBeanInvoiceSimplify`

NewUnibeeApiBeanInvoiceSimplifyWithDefaults instantiates a new UnibeeApiBeanInvoiceSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleAnchor

`func (o *UnibeeApiBeanInvoiceSimplify) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeApiBeanInvoiceSimplify) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeApiBeanInvoiceSimplify) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetBizType

`func (o *UnibeeApiBeanInvoiceSimplify) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeApiBeanInvoiceSimplify) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeApiBeanInvoiceSimplify) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanInvoiceSimplify) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanInvoiceSimplify) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanInvoiceSimplify) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateFrom

`func (o *UnibeeApiBeanInvoiceSimplify) GetCreateFrom() string`

GetCreateFrom returns the CreateFrom field if non-nil, zero value otherwise.

### GetCreateFromOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetCreateFromOk() (*string, bool)`

GetCreateFromOk returns a tuple with the CreateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFrom

`func (o *UnibeeApiBeanInvoiceSimplify) SetCreateFrom(v string)`

SetCreateFrom sets CreateFrom field to given value.

### HasCreateFrom

`func (o *UnibeeApiBeanInvoiceSimplify) HasCreateFrom() bool`

HasCreateFrom returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetCryptoAmount() int64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetCryptoAmountOk() (*int64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetCryptoAmount(v int64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) GetCryptoCurrency() string`

GetCryptoCurrency returns the CryptoCurrency field if non-nil, zero value otherwise.

### GetCryptoCurrencyOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetCryptoCurrencyOk() (*string, bool)`

GetCryptoCurrencyOk returns a tuple with the CryptoCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) SetCryptoCurrency(v string)`

SetCryptoCurrency sets CryptoCurrency field to given value.

### HasCryptoCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) HasCryptoCurrency() bool`

HasCryptoCurrency returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanInvoiceSimplify) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeApiBeanInvoiceSimplify) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiBeanInvoiceSimplify) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeApiBeanInvoiceSimplify) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDayUtilDue

`func (o *UnibeeApiBeanInvoiceSimplify) GetDayUtilDue() int64`

GetDayUtilDue returns the DayUtilDue field if non-nil, zero value otherwise.

### GetDayUtilDueOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetDayUtilDueOk() (*int64, bool)`

GetDayUtilDueOk returns a tuple with the DayUtilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayUtilDue

`func (o *UnibeeApiBeanInvoiceSimplify) SetDayUtilDue(v int64)`

SetDayUtilDue sets DayUtilDue field to given value.

### HasDayUtilDue

`func (o *UnibeeApiBeanInvoiceSimplify) HasDayUtilDue() bool`

HasDayUtilDue returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscount() UnibeeApiBeanMerchantDiscountCodeSimplify`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCodeSimplify, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanInvoiceSimplify) SetDiscount(v UnibeeApiBeanMerchantDiscountCodeSimplify)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanInvoiceSimplify) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiBeanInvoiceSimplify) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiBeanInvoiceSimplify) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanInvoiceSimplify) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanInvoiceSimplify) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanInvoiceSimplify) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanInvoiceSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanInvoiceSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanInvoiceSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanInvoiceSimplify) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanInvoiceSimplify) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanInvoiceSimplify) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiBeanInvoiceSimplify) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiBeanInvoiceSimplify) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiBeanInvoiceSimplify) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanInvoiceSimplify) GetLines() []UnibeeApiBeanInvoiceItemSimplify`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetLinesOk() (*[]UnibeeApiBeanInvoiceItemSimplify, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanInvoiceSimplify) SetLines(v []UnibeeApiBeanInvoiceItemSimplify)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanInvoiceSimplify) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanInvoiceSimplify) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanInvoiceSimplify) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanInvoiceSimplify) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanInvoiceSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanInvoiceSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanInvoiceSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanInvoiceSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanInvoiceSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanInvoiceSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentLink

`func (o *UnibeeApiBeanInvoiceSimplify) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeApiBeanInvoiceSimplify) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *UnibeeApiBeanInvoiceSimplify) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanInvoiceSimplify) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanInvoiceSimplify) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanInvoiceSimplify) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanInvoiceSimplify) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanInvoiceSimplify) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanInvoiceSimplify) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanInvoiceSimplify) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanInvoiceSimplify) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanInvoiceSimplify) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProrationDate

`func (o *UnibeeApiBeanInvoiceSimplify) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeApiBeanInvoiceSimplify) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeApiBeanInvoiceSimplify) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetProrationScale

`func (o *UnibeeApiBeanInvoiceSimplify) GetProrationScale() int64`

GetProrationScale returns the ProrationScale field if non-nil, zero value otherwise.

### GetProrationScaleOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetProrationScaleOk() (*int64, bool)`

GetProrationScaleOk returns a tuple with the ProrationScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationScale

`func (o *UnibeeApiBeanInvoiceSimplify) SetProrationScale(v int64)`

SetProrationScale sets ProrationScale field to given value.

### HasProrationScale

`func (o *UnibeeApiBeanInvoiceSimplify) HasProrationScale() bool`

HasProrationScale returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanInvoiceSimplify) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanInvoiceSimplify) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanInvoiceSimplify) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeApiBeanInvoiceSimplify) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeApiBeanInvoiceSimplify) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeApiBeanInvoiceSimplify) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeApiBeanInvoiceSimplify) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeApiBeanInvoiceSimplify) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeApiBeanInvoiceSimplify) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanInvoiceSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanInvoiceSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanInvoiceSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanInvoiceSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanInvoiceSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanInvoiceSimplify) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanInvoiceSimplify) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanInvoiceSimplify) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanInvoiceSimplify) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanInvoiceSimplify) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanInvoiceSimplify) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceSimplify) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiBeanInvoiceSimplify) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiBeanInvoiceSimplify) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiBeanInvoiceSimplify) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanInvoiceSimplify) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanInvoiceSimplify) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanInvoiceSimplify) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanInvoiceSimplify) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


