# UnibeeApiBeanInvoiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency | [optional] 
**Data** | Pointer to **string** | Data | [optional] 
**DiscountAmount** | Pointer to **int64** | DiscountAmount,Cents | [optional] 
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
**SendTerms** | Pointer to **string** | SendTerms | [optional] 
**Status** | Pointer to **int32** | Status，0-Init | 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscriptionSimplify**](UnibeeApiBeanSubscriptionSimplify.md) |  | [optional] 
**SubscriptionAmount** | Pointer to **int64** | SubscriptionAmount,Cents | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** | SubscriptionAmountExcludingTax,Cents | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**TaxAmount** | Pointer to **int64** | TaxAmount,Cents | [optional] 
**TaxScale** | Pointer to **int64** | TaxScale，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | TotalAmount,Cents | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** | TotalAmountExcludingTax,Cents | [optional] 
**UniqueId** | Pointer to **string** | UniqueId | [optional] 
**UserAccount** | Pointer to [**UnibeeApiBeanUserAccountSimplify**](UnibeeApiBeanUserAccountSimplify.md) |  | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeApiBeanInvoiceDetail

`func NewUnibeeApiBeanInvoiceDetail() *UnibeeApiBeanInvoiceDetail`

NewUnibeeApiBeanInvoiceDetail instantiates a new UnibeeApiBeanInvoiceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanInvoiceDetailWithDefaults

`func NewUnibeeApiBeanInvoiceDetailWithDefaults() *UnibeeApiBeanInvoiceDetail`

NewUnibeeApiBeanInvoiceDetailWithDefaults instantiates a new UnibeeApiBeanInvoiceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiBeanInvoiceDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanInvoiceDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanInvoiceDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanInvoiceDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeApiBeanInvoiceDetail) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeApiBeanInvoiceDetail) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeApiBeanInvoiceDetail) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeApiBeanInvoiceDetail) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanInvoiceDetail) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanInvoiceDetail) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanInvoiceDetail) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanInvoiceDetail) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanInvoiceDetail) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanInvoiceDetail) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanInvoiceDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayInvoiceId() string`

GetGatewayInvoiceId returns the GatewayInvoiceId field if non-nil, zero value otherwise.

### GetGatewayInvoiceIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayInvoiceIdOk() (*string, bool)`

GetGatewayInvoiceIdOk returns a tuple with the GatewayInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayInvoiceId(v string)`

SetGatewayInvoiceId sets GatewayInvoiceId field to given value.

### HasGatewayInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayInvoiceId() bool`

HasGatewayInvoiceId returns a boolean if a field has been set.

### GetGatewayInvoicePdf

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayInvoicePdf() string`

GetGatewayInvoicePdf returns the GatewayInvoicePdf field if non-nil, zero value otherwise.

### GetGatewayInvoicePdfOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayInvoicePdfOk() (*string, bool)`

GetGatewayInvoicePdfOk returns a tuple with the GatewayInvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoicePdf

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayInvoicePdf(v string)`

SetGatewayInvoicePdf sets GatewayInvoicePdf field to given value.

### HasGatewayInvoicePdf

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayInvoicePdf() bool`

HasGatewayInvoicePdf returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGatewayUserId

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayUserId() string`

GetGatewayUserId returns the GatewayUserId field if non-nil, zero value otherwise.

### GetGatewayUserIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGatewayUserIdOk() (*string, bool)`

GetGatewayUserIdOk returns a tuple with the GatewayUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUserId

`func (o *UnibeeApiBeanInvoiceDetail) SetGatewayUserId(v string)`

SetGatewayUserId sets GatewayUserId field to given value.

### HasGatewayUserId

`func (o *UnibeeApiBeanInvoiceDetail) HasGatewayUserId() bool`

HasGatewayUserId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeApiBeanInvoiceDetail) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeApiBeanInvoiceDetail) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeApiBeanInvoiceDetail) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeApiBeanInvoiceDetail) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeApiBeanInvoiceDetail) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeApiBeanInvoiceDetail) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeApiBeanInvoiceDetail) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanInvoiceDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanInvoiceDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanInvoiceDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanInvoiceDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiBeanInvoiceDetail) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiBeanInvoiceDetail) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiBeanInvoiceDetail) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiBeanInvoiceDetail) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanInvoiceDetail) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanInvoiceDetail) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanInvoiceDetail) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanInvoiceDetail) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanInvoiceDetail) GetLines() []UnibeeApiBeanInvoiceItemSimplify`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanInvoiceDetail) GetLinesOk() (*[]UnibeeApiBeanInvoiceItemSimplify, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanInvoiceDetail) SetLines(v []UnibeeApiBeanInvoiceItemSimplify)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanInvoiceDetail) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanInvoiceDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanInvoiceDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanInvoiceDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanInvoiceDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeApiBeanInvoiceDetail) GetMerchant() UnibeeApiBeanMerchantSimplify`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiBeanInvoiceDetail) GetMerchantOk() (*UnibeeApiBeanMerchantSimplify, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiBeanInvoiceDetail) SetMerchant(v UnibeeApiBeanMerchantSimplify)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiBeanInvoiceDetail) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanInvoiceDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanInvoiceDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanInvoiceDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanInvoiceDetail) GetPayment() UnibeeApiBeanPaymentSimplify`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanInvoiceDetail) GetPaymentOk() (*UnibeeApiBeanPaymentSimplify, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanInvoiceDetail) SetPayment(v UnibeeApiBeanPaymentSimplify)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanInvoiceDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanInvoiceDetail) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanInvoiceDetail) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanInvoiceDetail) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanInvoiceDetail) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanInvoiceDetail) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanInvoiceDetail) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanInvoiceDetail) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanInvoiceDetail) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanInvoiceDetail) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeApiBeanInvoiceDetail) GetRefund() UnibeeApiBeanRefundSimplify`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeApiBeanInvoiceDetail) GetRefundOk() (*UnibeeApiBeanRefundSimplify, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeApiBeanInvoiceDetail) SetRefund(v UnibeeApiBeanRefundSimplify)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeApiBeanInvoiceDetail) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanInvoiceDetail) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanInvoiceDetail) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanInvoiceDetail) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeApiBeanInvoiceDetail) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeApiBeanInvoiceDetail) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeApiBeanInvoiceDetail) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeApiBeanInvoiceDetail) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeApiBeanInvoiceDetail) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeApiBeanInvoiceDetail) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendPdf

`func (o *UnibeeApiBeanInvoiceDetail) GetSendPdf() string`

GetSendPdf returns the SendPdf field if non-nil, zero value otherwise.

### GetSendPdfOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSendPdfOk() (*string, bool)`

GetSendPdfOk returns a tuple with the SendPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPdf

`func (o *UnibeeApiBeanInvoiceDetail) SetSendPdf(v string)`

SetSendPdf sets SendPdf field to given value.

### HasSendPdf

`func (o *UnibeeApiBeanInvoiceDetail) HasSendPdf() bool`

HasSendPdf returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeApiBeanInvoiceDetail) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeApiBeanInvoiceDetail) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeApiBeanInvoiceDetail) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetSendTerms

`func (o *UnibeeApiBeanInvoiceDetail) GetSendTerms() string`

GetSendTerms returns the SendTerms field if non-nil, zero value otherwise.

### GetSendTermsOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSendTermsOk() (*string, bool)`

GetSendTermsOk returns a tuple with the SendTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTerms

`func (o *UnibeeApiBeanInvoiceDetail) SetSendTerms(v string)`

SetSendTerms sets SendTerms field to given value.

### HasSendTerms

`func (o *UnibeeApiBeanInvoiceDetail) HasSendTerms() bool`

HasSendTerms returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanInvoiceDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanInvoiceDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanInvoiceDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanInvoiceDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscription() UnibeeApiBeanSubscriptionSimplify`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscriptionSimplify, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanInvoiceDetail) SetSubscription(v UnibeeApiBeanSubscriptionSimplify)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanInvoiceDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceDetail) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeApiBeanInvoiceDetail) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanInvoiceDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanInvoiceDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiBeanInvoiceDetail) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiBeanInvoiceDetail) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiBeanInvoiceDetail) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiBeanInvoiceDetail) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeApiBeanInvoiceDetail) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeApiBeanInvoiceDetail) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeApiBeanInvoiceDetail) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeApiBeanInvoiceDetail) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanInvoiceDetail) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanInvoiceDetail) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanInvoiceDetail) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanInvoiceDetail) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeApiBeanInvoiceDetail) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeApiBeanInvoiceDetail) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeApiBeanInvoiceDetail) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeApiBeanInvoiceDetail) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeApiBeanInvoiceDetail) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserAccount

`func (o *UnibeeApiBeanInvoiceDetail) GetUserAccount() UnibeeApiBeanUserAccountSimplify`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *UnibeeApiBeanInvoiceDetail) GetUserAccountOk() (*UnibeeApiBeanUserAccountSimplify, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *UnibeeApiBeanInvoiceDetail) SetUserAccount(v UnibeeApiBeanUserAccountSimplify)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *UnibeeApiBeanInvoiceDetail) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanInvoiceDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanInvoiceDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanInvoiceDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanInvoiceDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


