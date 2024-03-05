# UnibeeInternalModelEntityOverseaPayInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizType** | Pointer to **int32** | biz type from payment 1-single payment, 3-subscription | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Data** | Pointer to **string** | data (json) | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayInvoiceId** | Pointer to **string** |  | [optional] 
**GatewayInvoicePdf** | Pointer to **string** |  | [optional] 
**GatewayPaymentId** | Pointer to **string** |  | [optional] 
**GatewayStatus** | Pointer to **string** |  | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice_id | [optional] 
**InvoiceName** | Pointer to **string** | invoice name | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**Lines** | Pointer to **string** | lines( json) | [optional] 
**Link** | Pointer to **string** | invoice link | [optional] 
**MerchantId** | Pointer to **int64** | merchant_id | [optional] 
**PaymentId** | Pointer to **string** | paymentId | [optional] 
**PaymentLink** | Pointer to **string** | invoice payment link | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end utc time | [optional] 
**PeriodEndTime** | Pointer to **string** |  | [optional] 
**PeriodStart** | Pointer to **int64** | period_start, utc time | [optional] 
**PeriodStartTime** | Pointer to **string** |  | [optional] 
**RefundId** | Pointer to **string** | refundId | [optional] 
**SendEmail** | Pointer to **string** | email | [optional] 
**SendNote** | Pointer to **string** | send_note | [optional] 
**SendPdf** | Pointer to **string** | pdf link | [optional] 
**SendStatus** | Pointer to **int32** | email send status，0-No | 1- YES | [optional] 
**SendTerms** | Pointer to **string** | send_terms | [optional] 
**Status** | Pointer to **int32** | status，0-Init | 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**SubscriptionAmount** | Pointer to **int64** | sub amount,cent | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 
**TaxAmount** | Pointer to **int64** | tax amount,cent | [optional] 
**TaxScale** | Pointer to **int64** | Tax scale，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | total amount, cent | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** | unique_id | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayInvoice

`func NewUnibeeInternalModelEntityOverseaPayInvoice() *UnibeeInternalModelEntityOverseaPayInvoice`

NewUnibeeInternalModelEntityOverseaPayInvoice instantiates a new UnibeeInternalModelEntityOverseaPayInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayInvoiceWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayInvoiceWithDefaults() *UnibeeInternalModelEntityOverseaPayInvoice`

NewUnibeeInternalModelEntityOverseaPayInvoiceWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizType

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasData() bool`

HasData returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayInvoiceId() string`

GetGatewayInvoiceId returns the GatewayInvoiceId field if non-nil, zero value otherwise.

### GetGatewayInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayInvoiceIdOk() (*string, bool)`

GetGatewayInvoiceIdOk returns a tuple with the GatewayInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGatewayInvoiceId(v string)`

SetGatewayInvoiceId sets GatewayInvoiceId field to given value.

### HasGatewayInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGatewayInvoiceId() bool`

HasGatewayInvoiceId returns a boolean if a field has been set.

### GetGatewayInvoicePdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayInvoicePdf() string`

GetGatewayInvoicePdf returns the GatewayInvoicePdf field if non-nil, zero value otherwise.

### GetGatewayInvoicePdfOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayInvoicePdfOk() (*string, bool)`

GetGatewayInvoicePdfOk returns a tuple with the GatewayInvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoicePdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGatewayInvoicePdf(v string)`

SetGatewayInvoicePdf sets GatewayInvoicePdf field to given value.

### HasGatewayInvoicePdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGatewayInvoicePdf() bool`

HasGatewayInvoicePdf returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetLines() string`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetLinesOk() (*string, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetLines(v string)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodEndTime() string`

GetPeriodEndTime returns the PeriodEndTime field if non-nil, zero value otherwise.

### GetPeriodEndTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodEndTimeOk() (*string, bool)`

GetPeriodEndTimeOk returns a tuple with the PeriodEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPeriodEndTime(v string)`

SetPeriodEndTime sets PeriodEndTime field to given value.

### HasPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPeriodEndTime() bool`

HasPeriodEndTime returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodStartTime() string`

GetPeriodStartTime returns the PeriodStartTime field if non-nil, zero value otherwise.

### GetPeriodStartTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetPeriodStartTimeOk() (*string, bool)`

GetPeriodStartTimeOk returns a tuple with the PeriodStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetPeriodStartTime(v string)`

SetPeriodStartTime sets PeriodStartTime field to given value.

### HasPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasPeriodStartTime() bool`

HasPeriodStartTime returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendPdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendPdf() string`

GetSendPdf returns the SendPdf field if non-nil, zero value otherwise.

### GetSendPdfOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendPdfOk() (*string, bool)`

GetSendPdfOk returns a tuple with the SendPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSendPdf(v string)`

SetSendPdf sets SendPdf field to given value.

### HasSendPdf

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSendPdf() bool`

HasSendPdf returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetSendTerms

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendTerms() string`

GetSendTerms returns the SendTerms field if non-nil, zero value otherwise.

### GetSendTermsOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSendTermsOk() (*string, bool)`

GetSendTermsOk returns a tuple with the SendTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTerms

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSendTerms(v string)`

SetSendTerms sets SendTerms field to given value.

### HasSendTerms

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSendTerms() bool`

HasSendTerms returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPayInvoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


