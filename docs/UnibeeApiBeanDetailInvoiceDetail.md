# UnibeeApiBeanDetailInvoiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycleAnchor** | Pointer to **int64** | billing_cycle_anchor | [optional] 
**CreateFrom** | Pointer to **string** | create from | [optional] 
**CryptoAmount** | Pointer to **int64** | crypto_amount, cent | [optional] 
**CryptoCurrency** | Pointer to **string** | crypto_currency | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**DayUtilDue** | Pointer to **int64** | day util due after finish | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCodeSimplify**](UnibeeApiBeanMerchantDiscountCodeSimplify.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** | DiscountAmount,Cents | [optional] 
**DiscountCode** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) |  | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**GatewayInvoiceId** | Pointer to **string** | GatewayInvoiceId | [optional] 
**GatewayInvoicePdf** | Pointer to **string** | GatewayInvoicePdf pdf | [optional] 
**GatewayPaymentId** | Pointer to **string** | GatewayPaymentId PaymentId | [optional] 
**GatewayStatus** | Pointer to **string** | GatewayStatus，Stripe：https://stripe.com/docs/api/invoices/object | [optional] 
**GatewayUserId** | Pointer to **string** | GatewayUserId Id | [optional] 
**GmtCreate** | Pointer to **string** | GmtCreate | [optional] 
**GmtModify** | Pointer to **string** | GmtModify | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | InvoiceId | [optional] 
**InvoiceName** | Pointer to **string** | InvoiceName | [optional] 
**IsDeleted** | Pointer to **int32** |  | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanInvoiceItemSimplify**](UnibeeApiBeanInvoiceItemSimplify.md) | lines json data | [optional] 
**Link** | Pointer to **string** | Link | [optional] 
**Merchant** | Pointer to [**UnibeeApiBeanMerchantSimplify**](UnibeeApiBeanMerchantSimplify.md) |  | [optional] 
**MerchantId** | Pointer to **int64** | MerchantId | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**OriginAmount** | Pointer to **int64** | OriginAmount,Cents | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPaymentSimplify**](UnibeeApiBeanPaymentSimplify.md) |  | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end | [optional] 
**PeriodStart** | Pointer to **int64** | period_start | [optional] 
**Refund** | Pointer to [**UnibeeApiBeanRefundSimplify**](UnibeeApiBeanRefundSimplify.md) |  | [optional] 
**RefundId** | Pointer to **string** | refundId | [optional] 
**SendEmail** | Pointer to **string** | SendEmail | [optional] 
**SendNote** | Pointer to **string** | SendNote | [optional] 
**SendPdf** | Pointer to **string** | SendPdf | [optional] 
**SendStatus** | Pointer to **int32** | SendStatus，0-No | 1- YES | [optional] 
**Status** | Pointer to **int32** | Status，1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscriptionSimplify**](UnibeeApiBeanSubscriptionSimplify.md) |  | [optional] 
**SubscriptionAmount** | Pointer to **int64** | SubscriptionAmount,Cents | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** | SubscriptionAmountExcludingTax,Cents | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**TaxAmount** | Pointer to **int64** | TaxAmount,Cents | [optional] 
**TaxPercentage** | Pointer to **int64** | TaxPercentage，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | TotalAmount,Cents | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** | TotalAmountExcludingTax,Cents | [optional] 
**UniqueId** | Pointer to **string** | UniqueId | [optional] 
**UserAccount** | Pointer to [**UnibeeApiBeanUserAccountSimplify**](UnibeeApiBeanUserAccountSimplify.md) |  | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiBeanDetailInvoiceDetail

`func NewUnibeeApiBeanDetailInvoiceDetail() *UnibeeApiBeanDetailInvoiceDetail`

NewUnibeeApiBeanDetailInvoiceDetail instantiates a new UnibeeApiBeanDetailInvoiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailInvoiceDetailWithDefaults

`func NewUnibeeApiBeanDetailInvoiceDetailWithDefaults() *UnibeeApiBeanDetailInvoiceDetail`

NewUnibeeApiBeanDetailInvoiceDetailWithDefaults instantiates a new UnibeeApiBeanDetailInvoiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycleAnchor

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetCreateFrom

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCreateFrom() string`

GetCreateFrom returns the CreateFrom field if non-nil, zero value otherwise.

### GetCreateFromOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCreateFromOk() (*string, bool)`

GetCreateFromOk returns a tuple with the CreateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFrom

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetCreateFrom(v string)`

SetCreateFrom sets CreateFrom field to given value.

### HasCreateFrom

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasCreateFrom() bool`

HasCreateFrom returns a boolean if a field has been set.

### GetCryptoAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCryptoAmount() int64`

GetCryptoAmount returns the CryptoAmount field if non-nil, zero value otherwise.

### GetCryptoAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCryptoAmountOk() (*int64, bool)`

GetCryptoAmountOk returns a tuple with the CryptoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetCryptoAmount(v int64)`

SetCryptoAmount sets CryptoAmount field to given value.

### HasCryptoAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasCryptoAmount() bool`

HasCryptoAmount returns a boolean if a field has been set.

### GetCryptoCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCryptoCurrency() string`

GetCryptoCurrency returns the CryptoCurrency field if non-nil, zero value otherwise.

### GetCryptoCurrencyOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCryptoCurrencyOk() (*string, bool)`

GetCryptoCurrencyOk returns a tuple with the CryptoCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetCryptoCurrency(v string)`

SetCryptoCurrency sets CryptoCurrency field to given value.

### HasCryptoCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasCryptoCurrency() bool`

HasCryptoCurrency returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDayUtilDue

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDayUtilDue() int64`

GetDayUtilDue returns the DayUtilDue field if non-nil, zero value otherwise.

### GetDayUtilDueOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDayUtilDueOk() (*int64, bool)`

GetDayUtilDueOk returns a tuple with the DayUtilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayUtilDue

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetDayUtilDue(v int64)`

SetDayUtilDue sets DayUtilDue field to given value.

### HasDayUtilDue

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasDayUtilDue() bool`

HasDayUtilDue returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscount() UnibeeApiBeanMerchantDiscountCodeSimplify`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCodeSimplify, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetDiscount(v UnibeeApiBeanMerchantDiscountCodeSimplify)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayInvoiceId() string`

GetGatewayInvoiceId returns the GatewayInvoiceId field if non-nil, zero value otherwise.

### GetGatewayInvoiceIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayInvoiceIdOk() (*string, bool)`

GetGatewayInvoiceIdOk returns a tuple with the GatewayInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayInvoiceId(v string)`

SetGatewayInvoiceId sets GatewayInvoiceId field to given value.

### HasGatewayInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayInvoiceId() bool`

HasGatewayInvoiceId returns a boolean if a field has been set.

### GetGatewayInvoicePdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayInvoicePdf() string`

GetGatewayInvoicePdf returns the GatewayInvoicePdf field if non-nil, zero value otherwise.

### GetGatewayInvoicePdfOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayInvoicePdfOk() (*string, bool)`

GetGatewayInvoicePdfOk returns a tuple with the GatewayInvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoicePdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayInvoicePdf(v string)`

SetGatewayInvoicePdf sets GatewayInvoicePdf field to given value.

### HasGatewayInvoicePdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayInvoicePdf() bool`

HasGatewayInvoicePdf returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGatewayUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayUserId() string`

GetGatewayUserId returns the GatewayUserId field if non-nil, zero value otherwise.

### GetGatewayUserIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGatewayUserIdOk() (*string, bool)`

GetGatewayUserIdOk returns a tuple with the GatewayUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGatewayUserId(v string)`

SetGatewayUserId sets GatewayUserId field to given value.

### HasGatewayUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGatewayUserId() bool`

HasGatewayUserId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetLines() []UnibeeApiBeanInvoiceItemSimplify`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetLinesOk() (*[]UnibeeApiBeanInvoiceItemSimplify, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetLines(v []UnibeeApiBeanInvoiceItemSimplify)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMerchant() UnibeeApiBeanMerchantSimplify`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMerchantOk() (*UnibeeApiBeanMerchantSimplify, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetMerchant(v UnibeeApiBeanMerchantSimplify)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPayment() UnibeeApiBeanPaymentSimplify`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPaymentOk() (*UnibeeApiBeanPaymentSimplify, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetPayment(v UnibeeApiBeanPaymentSimplify)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetRefund() UnibeeApiBeanRefundSimplify`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetRefundOk() (*UnibeeApiBeanRefundSimplify, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetRefund(v UnibeeApiBeanRefundSimplify)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendPdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendPdf() string`

GetSendPdf returns the SendPdf field if non-nil, zero value otherwise.

### GetSendPdfOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendPdfOk() (*string, bool)`

GetSendPdfOk returns a tuple with the SendPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSendPdf(v string)`

SetSendPdf sets SendPdf field to given value.

### HasSendPdf

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSendPdf() bool`

HasSendPdf returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscription() UnibeeApiBeanSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSubscription(v UnibeeApiBeanSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserAccount

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUserAccount() UnibeeApiBeanUserAccountSimplify`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUserAccountOk() (*UnibeeApiBeanUserAccountSimplify, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetUserAccount(v UnibeeApiBeanUserAccountSimplify)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailInvoiceDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailInvoiceDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


