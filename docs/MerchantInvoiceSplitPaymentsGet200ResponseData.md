# MerchantInvoiceSplitPaymentsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidAmount** | Pointer to **int64** | Total paid amount from successful split payments in cents | [optional] 
**Payments** | Pointer to [**[]UnibeeApiBeanSplitPaymentInfo**](UnibeeApiBeanSplitPaymentInfo.md) | List of split payments | [optional] 
**RemainingAmount** | Pointer to **int64** | Remaining amount to be paid in cents | [optional] 
**TotalAmount** | Pointer to **int64** | Total invoice amount in cents | [optional] 

## Methods

### NewMerchantInvoiceSplitPaymentsGet200ResponseData

`func NewMerchantInvoiceSplitPaymentsGet200ResponseData() *MerchantInvoiceSplitPaymentsGet200ResponseData`

NewMerchantInvoiceSplitPaymentsGet200ResponseData instantiates a new MerchantInvoiceSplitPaymentsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantInvoiceSplitPaymentsGet200ResponseDataWithDefaults

`func NewMerchantInvoiceSplitPaymentsGet200ResponseDataWithDefaults() *MerchantInvoiceSplitPaymentsGet200ResponseData`

NewMerchantInvoiceSplitPaymentsGet200ResponseDataWithDefaults instantiates a new MerchantInvoiceSplitPaymentsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetPaidAmount() int64`

GetPaidAmount returns the PaidAmount field if non-nil, zero value otherwise.

### GetPaidAmountOk

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetPaidAmountOk() (*int64, bool)`

GetPaidAmountOk returns a tuple with the PaidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) SetPaidAmount(v int64)`

SetPaidAmount sets PaidAmount field to given value.

### HasPaidAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) HasPaidAmount() bool`

HasPaidAmount returns a boolean if a field has been set.

### GetPayments

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetPayments() []UnibeeApiBeanSplitPaymentInfo`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetPaymentsOk() (*[]UnibeeApiBeanSplitPaymentInfo, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) SetPayments(v []UnibeeApiBeanSplitPaymentInfo)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetRemainingAmount() int64`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetRemainingAmountOk() (*int64, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) SetRemainingAmount(v int64)`

SetRemainingAmount sets RemainingAmount field to given value.

### HasRemainingAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) HasRemainingAmount() bool`

HasRemainingAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MerchantInvoiceSplitPaymentsGet200ResponseData) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


