# UnibeeApiSystemInvoiceGetSplitPaymentsRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidAmount** | Pointer to **int64** | Total paid amount from successful split payments in cents | [optional] 
**Payments** | Pointer to [**[]UnibeeApiBeanSplitPaymentInfo**](UnibeeApiBeanSplitPaymentInfo.md) | List of split payments | [optional] 
**RemainingAmount** | Pointer to **int64** | Remaining amount to be paid in cents | [optional] 
**TotalAmount** | Pointer to **int64** | Total invoice amount in cents | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceGetSplitPaymentsRes

`func NewUnibeeApiSystemInvoiceGetSplitPaymentsRes() *UnibeeApiSystemInvoiceGetSplitPaymentsRes`

NewUnibeeApiSystemInvoiceGetSplitPaymentsRes instantiates a new UnibeeApiSystemInvoiceGetSplitPaymentsRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceGetSplitPaymentsResWithDefaults

`func NewUnibeeApiSystemInvoiceGetSplitPaymentsResWithDefaults() *UnibeeApiSystemInvoiceGetSplitPaymentsRes`

NewUnibeeApiSystemInvoiceGetSplitPaymentsResWithDefaults instantiates a new UnibeeApiSystemInvoiceGetSplitPaymentsRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetPaidAmount() int64`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetPaidAmountOk() (*int64, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) SetPaidAmount(v int64)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### GetPayments

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetPayments() []UnibeeApiBeanSplitPaymentInfo`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetPaymentsOk() (*[]UnibeeApiBeanSplitPaymentInfo, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) SetPayments(v []UnibeeApiBeanSplitPaymentInfo)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetRemainingAmount() int64`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetRemainingAmountOk() (*int64, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) SetRemainingAmount(v int64)`

SetRemainingAmount sets RemainingAmount field to given value.

### HasRemainingAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) HasRemainingAmount() bool`

HasRemainingAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiSystemInvoiceGetSplitPaymentsRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


