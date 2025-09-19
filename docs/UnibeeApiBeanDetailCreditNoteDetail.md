# UnibeeApiBeanDetailCreditNoteDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizType** | Pointer to **int32** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreateFrom** | Pointer to **string** | create from | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** | Currency | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** | DiscountAmount,Cents | [optional] 
**DiscountCode** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **int64** |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** | InvoiceId | [optional] 
**InvoiceName** | Pointer to **string** | InvoiceName | [optional] 
**Lines** | Pointer to [**[]UnibeeApiBeanInvoiceItemSimplify**](UnibeeApiBeanInvoiceItemSimplify.md) | lines json data | [optional] 
**Link** | Pointer to **string** | Link | [optional] 
**MerchantId** | Pointer to **int64** | MerchantId | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**OriginAmount** | Pointer to **int64** | OriginAmount,Cents | [optional] 
**OriginalPaymentInvoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**PaidTime** | Pointer to **int64** |  | [optional] 
**PartialCreditPaidAmount** | Pointer to **int64** | partial credit paid amount | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**PeriodEnd** | Pointer to **int64** | period_end | [optional] 
**PeriodStart** | Pointer to **int64** | period_start | [optional] 
**PlanSnapshot** | Pointer to [**UnibeeApiBeanInvoicePlanSnapshot**](UnibeeApiBeanInvoicePlanSnapshot.md) |  | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**PromoCreditDiscountAmount** | Pointer to **int64** | promo credit discount amount | [optional] 
**PromoCreditTransaction** | Pointer to [**UnibeeApiBeanCreditTransaction**](UnibeeApiBeanCreditTransaction.md) |  | [optional] 
**Refund** | Pointer to [**UnibeeApiBeanRefund**](UnibeeApiBeanRefund.md) |  | [optional] 
**RefundId** | Pointer to **string** | refundId | [optional] 
**Status** | Pointer to **int32** | Status，1-pending｜2-processing｜3-paid | 4-failed | 5-cancelled | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**SubscriptionAmount** | Pointer to **int64** | SubscriptionAmount,Cents | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** | SubscriptionAmountExcludingTax,Cents | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId | [optional] 
**TaxAmount** | Pointer to **int64** | TaxAmount,Cents | [optional] 
**TaxPercentage** | Pointer to **int64** | TaxPercentage，1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | TotalAmount,Cents | [optional] 
**TotalAmountExcludingTax** | Pointer to **int64** | TotalAmountExcludingTax,Cents | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 
**UserSnapshot** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailCreditNoteDetail

`func NewUnibeeApiBeanDetailCreditNoteDetail() *UnibeeApiBeanDetailCreditNoteDetail`

NewUnibeeApiBeanDetailCreditNoteDetail instantiates a new UnibeeApiBeanDetailCreditNoteDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailCreditNoteDetailWithDefaults

`func NewUnibeeApiBeanDetailCreditNoteDetailWithDefaults() *UnibeeApiBeanDetailCreditNoteDetail`

NewUnibeeApiBeanDetailCreditNoteDetailWithDefaults instantiates a new UnibeeApiBeanDetailCreditNoteDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizType

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateFrom

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCreateFrom() string`

GetCreateFrom returns the CreateFrom field if non-nil, zero value otherwise.

### GetCreateFromOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCreateFromOk() (*string, bool)`

GetCreateFromOk returns a tuple with the CreateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateFrom

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetCreateFrom(v string)`

SetCreateFrom sets CreateFrom field to given value.

### HasCreateFrom

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasCreateFrom() bool`

HasCreateFrom returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetLines() []UnibeeApiBeanInvoiceItemSimplify`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetLinesOk() (*[]UnibeeApiBeanInvoiceItemSimplify, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetLines(v []UnibeeApiBeanInvoiceItemSimplify)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMessage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetOriginalPaymentInvoice

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetOriginalPaymentInvoice() UnibeeApiBeanInvoice`

GetOriginalPaymentInvoice returns the OriginalPaymentInvoice field if non-nil, zero value otherwise.

### GetOriginalPaymentInvoiceOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetOriginalPaymentInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetOriginalPaymentInvoiceOk returns a tuple with the OriginalPaymentInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPaymentInvoice

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetOriginalPaymentInvoice(v UnibeeApiBeanInvoice)`

SetOriginalPaymentInvoice sets OriginalPaymentInvoice field to given value.

### HasOriginalPaymentInvoice

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasOriginalPaymentInvoice() bool`

HasOriginalPaymentInvoice returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPartialCreditPaidAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPartialCreditPaidAmount() int64`

GetPartialCreditPaidAmount returns the PartialCreditPaidAmount field if non-nil, zero value otherwise.

### GetPartialCreditPaidAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPartialCreditPaidAmountOk() (*int64, bool)`

GetPartialCreditPaidAmountOk returns a tuple with the PartialCreditPaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialCreditPaidAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPartialCreditPaidAmount(v int64)`

SetPartialCreditPaidAmount sets PartialCreditPaidAmount field to given value.

### HasPartialCreditPaidAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPartialCreditPaidAmount() bool`

HasPartialCreditPaidAmount returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPeriodEnd() int64`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPeriodEndOk() (*int64, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPeriodEnd(v int64)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPeriodStart() int64`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPeriodStartOk() (*int64, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPeriodStart(v int64)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPlanSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPlanSnapshot() UnibeeApiBeanInvoicePlanSnapshot`

GetPlanSnapshot returns the PlanSnapshot field if non-nil, zero value otherwise.

### GetPlanSnapshotOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPlanSnapshotOk() (*UnibeeApiBeanInvoicePlanSnapshot, bool)`

GetPlanSnapshotOk returns a tuple with the PlanSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPlanSnapshot(v UnibeeApiBeanInvoicePlanSnapshot)`

SetPlanSnapshot sets PlanSnapshot field to given value.

### HasPlanSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPlanSnapshot() bool`

HasPlanSnapshot returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetPromoCreditDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPromoCreditDiscountAmount() int64`

GetPromoCreditDiscountAmount returns the PromoCreditDiscountAmount field if non-nil, zero value otherwise.

### GetPromoCreditDiscountAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPromoCreditDiscountAmountOk() (*int64, bool)`

GetPromoCreditDiscountAmountOk returns a tuple with the PromoCreditDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPromoCreditDiscountAmount(v int64)`

SetPromoCreditDiscountAmount sets PromoCreditDiscountAmount field to given value.

### HasPromoCreditDiscountAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPromoCreditDiscountAmount() bool`

HasPromoCreditDiscountAmount returns a boolean if a field has been set.

### GetPromoCreditTransaction

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPromoCreditTransaction() UnibeeApiBeanCreditTransaction`

GetPromoCreditTransaction returns the PromoCreditTransaction field if non-nil, zero value otherwise.

### GetPromoCreditTransactionOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetPromoCreditTransactionOk() (*UnibeeApiBeanCreditTransaction, bool)`

GetPromoCreditTransactionOk returns a tuple with the PromoCreditTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoCreditTransaction

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetPromoCreditTransaction(v UnibeeApiBeanCreditTransaction)`

SetPromoCreditTransaction sets PromoCreditTransaction field to given value.

### HasPromoCreditTransaction

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasPromoCreditTransaction() bool`

HasPromoCreditTransaction returns a boolean if a field has been set.

### GetRefund

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetRefund() UnibeeApiBeanRefund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetRefundOk() (*UnibeeApiBeanRefund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetRefund(v UnibeeApiBeanRefund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetSubscriptionAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionAmount() int64`

GetSubscriptionAmount returns the SubscriptionAmount field if non-nil, zero value otherwise.

### GetSubscriptionAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionAmountOk() (*int64, bool)`

GetSubscriptionAmountOk returns a tuple with the SubscriptionAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetSubscriptionAmount(v int64)`

SetSubscriptionAmount sets SubscriptionAmount field to given value.

### HasSubscriptionAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasSubscriptionAmount() bool`

HasSubscriptionAmount returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTotalAmountExcludingTax() int64`

GetTotalAmountExcludingTax returns the TotalAmountExcludingTax field if non-nil, zero value otherwise.

### GetTotalAmountExcludingTaxOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetTotalAmountExcludingTaxOk() (*int64, bool)`

GetTotalAmountExcludingTaxOk returns a tuple with the TotalAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetTotalAmountExcludingTax(v int64)`

SetTotalAmountExcludingTax sets TotalAmountExcludingTax field to given value.

### HasTotalAmountExcludingTax

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasTotalAmountExcludingTax() bool`

HasTotalAmountExcludingTax returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetUserSnapshot() UnibeeApiBeanUserAccount`

GetUserSnapshot returns the UserSnapshot field if non-nil, zero value otherwise.

### GetUserSnapshotOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetUserSnapshotOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserSnapshotOk returns a tuple with the UserSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetUserSnapshot(v UnibeeApiBeanUserAccount)`

SetUserSnapshot sets UserSnapshot field to given value.

### HasUserSnapshot

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasUserSnapshot() bool`

HasUserSnapshot returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiBeanDetailCreditNoteDetail) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiBeanDetailCreditNoteDetail) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiBeanDetailCreditNoteDetail) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


