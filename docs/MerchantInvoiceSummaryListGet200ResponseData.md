# MerchantInvoiceSummaryListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Invoices** | Pointer to [**[]UnibeeApiMerchantInvoiceInvoiceSummary**](UnibeeApiMerchantInvoiceInvoiceSummary.md) |  | [optional] 

## Methods

### NewMerchantInvoiceSummaryListGet200ResponseData

`func NewMerchantInvoiceSummaryListGet200ResponseData() *MerchantInvoiceSummaryListGet200ResponseData`

NewMerchantInvoiceSummaryListGet200ResponseData instantiates a new MerchantInvoiceSummaryListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantInvoiceSummaryListGet200ResponseDataWithDefaults

`func NewMerchantInvoiceSummaryListGet200ResponseDataWithDefaults() *MerchantInvoiceSummaryListGet200ResponseData`

NewMerchantInvoiceSummaryListGet200ResponseDataWithDefaults instantiates a new MerchantInvoiceSummaryListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *MerchantInvoiceSummaryListGet200ResponseData) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *MerchantInvoiceSummaryListGet200ResponseData) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *MerchantInvoiceSummaryListGet200ResponseData) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *MerchantInvoiceSummaryListGet200ResponseData) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetInvoices

`func (o *MerchantInvoiceSummaryListGet200ResponseData) GetInvoices() []UnibeeApiMerchantInvoiceInvoiceSummary`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *MerchantInvoiceSummaryListGet200ResponseData) GetInvoicesOk() (*[]UnibeeApiMerchantInvoiceInvoiceSummary, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *MerchantInvoiceSummaryListGet200ResponseData) SetInvoices(v []UnibeeApiMerchantInvoiceInvoiceSummary)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *MerchantInvoiceSummaryListGet200ResponseData) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


