# UnibeeApiMerchantInvoiceFinishReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysUtilDue** | **int32** | Due Day Of Invoice Payment | 
**InvoiceId** | **string** | The unique id of invoice | 

## Methods

### NewUnibeeApiMerchantInvoiceFinishReq

`func NewUnibeeApiMerchantInvoiceFinishReq(daysUtilDue int32, invoiceId string, ) *UnibeeApiMerchantInvoiceFinishReq`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


