# UnibeeApiMerchantInvoiceHistoryInvoiceImportReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | Optional, ISO 3166-1 alpha-2 country code | [optional] 
**CreateTime** | Pointer to **string** | Optional, invoice create time &#39;2006-01-02 15:04:05&#39; (UTC) | [optional] 
**Currency** | **string** | Required, currency code (e.g. USD) | 
**DiscountAmount** | Pointer to **int64** | Optional, discount amount in cents | [optional] 
**Email** | Pointer to **string** | User email, one of Email or ExternalUserId is required | [optional] 
**ExternalUserId** | Pointer to **string** | External user id, one of Email or ExternalUserId is required | [optional] 
**GatewayId** | Pointer to **int64** | Optional, gateway id. If 0 or omitted, invoice will be stored with gateway_id&#x3D;0. | [optional] 
**InvoiceId** | **string** | Required, external invoice id | 
**InvoiceName** | Pointer to **string** | Optional, human readable invoice name (for display/search). If empty, system will use a generic name. | [optional] 
**InvoiceStatus** | Pointer to **string** | Optional, paid|failed|cancelled, default paid if blank | [optional] 
**Lines** | Pointer to [**[]UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem**](UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem.md) | Optional (API only), simplified invoice line items. If omitted or empty, system will automatically generate a single aggregated line from invoice-level amounts. | [optional] 
**LinkToLatestSubscription** | Pointer to **bool** | Optional, if true, link to user&#39;s latest active subscription in current merchant (ignores SubscriptionId). | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Optional, metadata map to be stored into invoice.meta_data. | [optional] 
**OriginalInvoiceId** | Pointer to **string** | Required when totalAmount &lt; 0, original payment invoice&#39;s InvoiceId (of the payment invoice). | [optional] 
**PeriodEnd** | Pointer to **string** | Optional, period end &#39;2006-01-02 15:04:05&#39; (UTC) | [optional] 
**PeriodStart** | Pointer to **string** | Optional, period start &#39;2006-01-02 15:04:05&#39; (UTC) | [optional] 
**ProductName** | Pointer to **string** | Optional, product name shown on invoice. If empty, may fall back to InvoiceName. | [optional] 
**SubscriptionId** | Pointer to **string** | Optional, subscription id. If provided, must belong to current user. | [optional] 
**TaxAmount** | Pointer to **int64** | Optional, tax amount in cents | [optional] 
**TotalAmount** | **int64** | Required, total amount in cents, positive for payment, negative for refund | 

## Methods

### NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReq

`func NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReq(currency string, invoiceId string, totalAmount int64, ) *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq`

NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReq instantiates a new UnibeeApiMerchantInvoiceHistoryInvoiceImportReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReqWithDefaults

`func NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReqWithDefaults() *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq`

NewUnibeeApiMerchantInvoiceHistoryInvoiceImportReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceHistoryInvoiceImportReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetInvoiceName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetInvoiceStatus

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceStatus() string`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetInvoiceStatusOk() (*string, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetInvoiceStatus(v string)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetLines

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetLines() []UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetLinesOk() (*[]UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetLines(v []UnibeeApiMerchantInvoiceHistoryInvoiceImportLineItem)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLinkToLatestSubscription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetLinkToLatestSubscription() bool`

GetLinkToLatestSubscription returns the LinkToLatestSubscription field if non-nil, zero value otherwise.

### GetLinkToLatestSubscriptionOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetLinkToLatestSubscriptionOk() (*bool, bool)`

GetLinkToLatestSubscriptionOk returns a tuple with the LinkToLatestSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToLatestSubscription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetLinkToLatestSubscription(v bool)`

SetLinkToLatestSubscription sets LinkToLatestSubscription field to given value.

### HasLinkToLatestSubscription

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasLinkToLatestSubscription() bool`

HasLinkToLatestSubscription returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOriginalInvoiceId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetOriginalInvoiceId() string`

GetOriginalInvoiceId returns the OriginalInvoiceId field if non-nil, zero value otherwise.

### GetOriginalInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetOriginalInvoiceIdOk() (*string, bool)`

GetOriginalInvoiceIdOk returns a tuple with the OriginalInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalInvoiceId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetOriginalInvoiceId(v string)`

SetOriginalInvoiceId sets OriginalInvoiceId field to given value.

### HasOriginalInvoiceId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasOriginalInvoiceId() bool`

HasOriginalInvoiceId returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPeriodStart

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantInvoiceHistoryInvoiceImportReq) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


