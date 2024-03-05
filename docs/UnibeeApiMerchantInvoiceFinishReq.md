# UnibeeApiMerchantInvoiceFinishReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysUtilDue** | **int32** | DaysUtilDue,Due Day Of Pay | 
**InvoiceId** | **string** | InvoiceId | 
**PayMethod** | **int32** | PayMethod,1-manual，2-auto | 

## Methods

### NewUnibeeApiMerchantInvoiceFinishReq

`func NewUnibeeApiMerchantInvoiceFinishReq(daysUtilDue int32, invoiceId string, payMethod int32, ) *UnibeeApiMerchantInvoiceFinishReq`

NewUnibeeApiMerchantInvoiceFinishReq instantiates a new UnibeeApiMerchantInvoiceFinishReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceFinishReqWithDefaults

`func NewUnibeeApiMerchantInvoiceFinishReqWithDefaults() *UnibeeApiMerchantInvoiceFinishReq`

NewUnibeeApiMerchantInvoiceFinishReqWithDefaults instantiates a new UnibeeApiMerchantInvoiceFinishReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysUtilDue

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetDaysUtilDue() int32`

GetDaysUtilDue returns the DaysUtilDue field if non-nil, zero value otherwise.

### GetDaysUtilDueOk

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetDaysUtilDueOk() (*int32, bool)`

GetDaysUtilDueOk returns a tuple with the DaysUtilDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysUtilDue

`func (o *UnibeeApiMerchantInvoiceFinishReq) SetDaysUtilDue(v int32)`

SetDaysUtilDue sets DaysUtilDue field to given value.


### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceFinishReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.


### GetPayMethod

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetPayMethod() int32`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *UnibeeApiMerchantInvoiceFinishReq) GetPayMethodOk() (*int32, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *UnibeeApiMerchantInvoiceFinishReq) SetPayMethod(v int32)`

SetPayMethod sets PayMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


