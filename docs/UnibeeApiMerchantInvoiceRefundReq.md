# UnibeeApiMerchantInvoiceRefundReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceId** | **string** | The unique id of invoice | 
**Reason** | **string** | The reason of refund | 
**RefundAmount** | **int64** | The amount of refund | 
**RefundNo** | Pointer to **string** | The out refund number | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceRefundReq

`func NewUnibeeApiMerchantInvoiceRefundReq(invoiceId string, reason string, refundAmount int64, ) *UnibeeApiMerchantInvoiceRefundReq`

NewUnibeeApiMerchantInvoiceRefundReq instantiates a new UnibeeApiMerchantInvoiceRefundReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceRefundReqWithDefaults

`func NewUnibeeApiMerchantInvoiceRefundReqWithDefaults() *UnibeeApiMerchantInvoiceRefundReq`

NewUnibeeApiMerchantInvoiceRefundReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceRefundReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceRefundReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetReason

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnibeeApiMerchantInvoiceRefundReq) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRefundAmount

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeApiMerchantInvoiceRefundReq) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.


### GetRefundNo

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundNo() string`

GetRefundNo returns the RefundNo field if non-nil, zero value otherwise.

### GetRefundNoOk

`func (o *UnibeeApiMerchantInvoiceRefundReq) GetRefundNoOk() (*string, bool)`

GetRefundNoOk returns a tuple with the RefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNo

`func (o *UnibeeApiMerchantInvoiceRefundReq) SetRefundNo(v string)`

SetRefundNo sets RefundNo field to given value.

### HasRefundNo

`func (o *UnibeeApiMerchantInvoiceRefundReq) HasRefundNo() bool`

HasRefundNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


