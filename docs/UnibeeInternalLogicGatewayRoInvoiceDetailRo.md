# UnibeeInternalLogicGatewayRoInvoiceDetailRo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency | [optional] 
**Data** | Pointer to **string** | Data | [optional] 
**DiscountAmount** | Pointer to **int64** | DiscountAmount,Cents | [optional] 
**Gateway** | Pointer to [**UnibeeInternalLogicGatewayRoGatewaySimplify**](UnibeeInternalLogicGatewayRoGatewaySimplify.md) |  | [optional] 
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
**Lines** | Pointer to [**[]UnibeeInternalLogicGatewayRoInvoiceItemDetailRo**](UnibeeInternalLogicGatewayRoInvoiceItemDetailRo.md) | lines json data | [optional] 
**Link** | Pointer to **string** | Link | [optional] 
**Merchant** | Pointer to [**UnibeeInternalModelEntityOverseaPayMerchant**](UnibeeInternalModelEntityOverseaPayMerchant.md) |  | [optional] 
**MerchantId** | Pointer to **int64** | MerchantId | [optional] 
**Payment** | Pointer to [**UnibeeInternalModelEntityOverseaPayPayment**](UnibeeInternalModelEntityOverseaPayPayment.md) |  | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end | [optional] 
**PeriodStart** | Pointer to **int64** | period_start | [optional] 
**Refund** | Pointer to [**UnibeeInternalModelEntityOverseaPayRefund**](UnibeeInternalModelEntityOverseaPayRefund.md) |  | [optional] 
**RefundId** | Pointer to **string** | refundId | [optional] 
**SendEmail** | Pointer to **string** | SendEmail | [optional] 
**SendNote** | Pointer to **string** | SendNote | [optional] 
**SendPdf** | Pointer to **string** | SendPdf | [optional] 
**SendStatus** | Pointer to **int32** | SendStatus，0-No | 1- YES | [optional] 
**SendTerms** | Pointer to **string** | SendTerms | [optional] 
**Status** | Pointer to **int32** | Status，0-Init | 1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**Subscription** | Pointer to [**UnibeeInternalModelEntityOverseaPaySubscription**](UnibeeInternalModelEntityOverseaPaySubscription.md) |  | [optional] 
**SubscriptionAmount** | Pointer to **int64** | SubscriptionAmount,Cents | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** | SubscriptionAmountExcludingTax,Cents | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**TaxAmount** | Pointer to **int64** | TaxAmount,Cents | [optional] 
**TaxScale** | Pointer to **int64** | TaxScale，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | TotalAmount,Cents | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** | TotalAmountExcludingTax,Cents | [optional] 
**UniqueId** | Pointer to **string** | UniqueId | [optional] 
**UserAccount** | Pointer to [**UnibeeInternalModelEntityOverseaPayUserAccount**](UnibeeInternalModelEntityOverseaPayUserAccount.md) |  | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 

## Methods

### NewUnibeeInternalLogicGatewayRoInvoiceDetailRo

`func NewUnibeeInternalLogicGatewayRoInvoiceDetailRo() *UnibeeInternalLogicGatewayRoInvoiceDetailRo`

NewUnibeeInternalLogicGatewayRoInvoiceDetailRo instantiates a new UnibeeInternalLogicGatewayRoInvoiceDetailRo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicGatewayRoInvoiceDetailRoWithDefaults

`func NewUnibeeInternalLogicGatewayRoInvoiceDetailRoWithDefaults() *UnibeeInternalLogicGatewayRoInvoiceDetailRo`

NewUnibeeInternalLogicGatewayRoInvoiceDetailRoWithDefaults instantiates a new UnibeeInternalLogicGatewayRoInvoiceDetailRo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetData

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGateway() UnibeeInternalLogicGatewayRoGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayOk() (*UnibeeInternalLogicGatewayRoGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGateway(v UnibeeInternalLogicGatewayRoGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayInvoiceId() string`

GetGatewayInvoiceId returns the GatewayInvoiceId field if non-nil, zero value otherwise.

### GetGatewayInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayInvoiceIdOk() (*string, bool)`

GetGatewayInvoiceIdOk returns a tuple with the GatewayInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayInvoiceId(v string)`

SetGatewayInvoiceId sets GatewayInvoiceId field to given value.

### HasGatewayInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayInvoiceId() bool`

HasGatewayInvoiceId returns a boolean if a field has been set.

### GetGatewayInvoicePdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayInvoicePdf() string`

GetGatewayInvoicePdf returns the GatewayInvoicePdf field if non-nil, zero value otherwise.

### GetGatewayInvoicePdfOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayInvoicePdfOk() (*string, bool)`

GetGatewayInvoicePdfOk returns a tuple with the GatewayInvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayInvoicePdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayInvoicePdf(v string)`

SetGatewayInvoicePdf sets GatewayInvoicePdf field to given value.

### HasGatewayInvoicePdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayInvoicePdf() bool`

HasGatewayInvoicePdf returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGatewayUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayUserId() string`

GetGatewayUserId returns the GatewayUserId field if non-nil, zero value otherwise.

### GetGatewayUserIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGatewayUserIdOk() (*string, bool)`

GetGatewayUserIdOk returns a tuple with the GatewayUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGatewayUserId(v string)`

SetGatewayUserId sets GatewayUserId field to given value.

### HasGatewayUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGatewayUserId() bool`

HasGatewayUserId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetLines() []UnibeeInternalLogicGatewayRoInvoiceItemDetailRo`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetLinesOk() (*[]UnibeeInternalLogicGatewayRoInvoiceItemDetailRo, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetLines(v []UnibeeInternalLogicGatewayRoInvoiceItemDetailRo)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchant

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetMerchant() UnibeeInternalModelEntityOverseaPayMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetMerchantOk() (*UnibeeInternalModelEntityOverseaPayMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetMerchant(v UnibeeInternalModelEntityOverseaPayMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPayment() UnibeeInternalModelEntityOverseaPayPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPaymentOk() (*UnibeeInternalModelEntityOverseaPayPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetPayment(v UnibeeInternalModelEntityOverseaPayPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetRefund() UnibeeInternalModelEntityOverseaPayRefund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetRefundOk() (*UnibeeInternalModelEntityOverseaPayRefund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetRefund(v UnibeeInternalModelEntityOverseaPayRefund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetSendNote

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendNote() string`

GetSendNote returns the SendNote field if non-nil, zero value otherwise.

### GetSendNoteOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendNoteOk() (*string, bool)`

GetSendNoteOk returns a tuple with the SendNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNote

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSendNote(v string)`

SetSendNote sets SendNote field to given value.

### HasSendNote

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSendNote() bool`

HasSendNote returns a boolean if a field has been set.

### GetSendPdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendPdf() string`

GetSendPdf returns the SendPdf field if non-nil, zero value otherwise.

### GetSendPdfOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendPdfOk() (*string, bool)`

GetSendPdfOk returns a tuple with the SendPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendPdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSendPdf(v string)`

SetSendPdf sets SendPdf field to given value.

### HasSendPdf

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSendPdf() bool`

HasSendPdf returns a boolean if a field has been set.

### GetSendStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendStatus() int32`

GetSendStatus returns the SendStatus field if non-nil, zero value otherwise.

### GetSendStatusOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendStatusOk() (*int32, bool)`

GetSendStatusOk returns a tuple with the SendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSendStatus(v int32)`

SetSendStatus sets SendStatus field to given value.

### HasSendStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSendStatus() bool`

HasSendStatus returns a boolean if a field has been set.

### GetSendTerms

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendTerms() string`

GetSendTerms returns the SendTerms field if non-nil, zero value otherwise.

### GetSendTermsOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSendTermsOk() (*string, bool)`

GetSendTermsOk returns a tuple with the SendTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTerms

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSendTerms(v string)`

SetSendTerms sets SendTerms field to given value.

### HasSendTerms

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSendTerms() bool`

HasSendTerms returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscription() UnibeeInternalModelEntityOverseaPaySubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionOk() (*UnibeeInternalModelEntityOverseaPaySubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSubscription(v UnibeeInternalModelEntityOverseaPaySubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserAccount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUserAccount() UnibeeInternalModelEntityOverseaPayUserAccount`

GetUserAccount returns the UserAccount field if non-nil, zero value otherwise.

### GetUserAccountOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUserAccountOk() (*UnibeeInternalModelEntityOverseaPayUserAccount, bool)`

GetUserAccountOk returns a tuple with the UserAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetUserAccount(v UnibeeInternalModelEntityOverseaPayUserAccount)`

SetUserAccount sets UserAccount field to given value.

### HasUserAccount

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasUserAccount() bool`

HasUserAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalLogicGatewayRoInvoiceDetailRo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


