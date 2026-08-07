# UnibeeApiMerchantInvoiceSummaryListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Invoices** | Pointer to [**[]UnibeeApiMerchantInvoiceInvoiceSummary**](UnibeeApiMerchantInvoiceInvoiceSummary.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceSummaryListRes

`func NewUnibeeApiMerchantInvoiceSummaryListRes() *UnibeeApiMerchantInvoiceSummaryListRes`

NewUnibeeApiMerchantInvoiceSummaryListRes instantiates a new UnibeeApiMerchantInvoiceSummaryListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceSummaryListResWithDefaults

`func NewUnibeeApiMerchantInvoiceSummaryListResWithDefaults() *UnibeeApiMerchantInvoiceSummaryListRes`

NewUnibeeApiMerchantInvoiceSummaryListResWithDefaults instantiates a new UnibeeApiMerchantInvoiceSummaryListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetInvoices

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) GetInvoices() []UnibeeApiMerchantInvoiceInvoiceSummary`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) GetInvoicesOk() (*[]UnibeeApiMerchantInvoiceInvoiceSummary, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) SetInvoices(v []UnibeeApiMerchantInvoiceInvoiceSummary)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *UnibeeApiMerchantInvoiceSummaryListRes) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


