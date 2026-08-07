# UnibeeApiMerchantInvoiceInvoiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BizType** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FinishTime** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**InvoiceId** | Pointer to **string** |  | [optional] 
**InvoiceName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**PaidTime** | Pointer to **int64** |  | [optional] 
**SendEmail** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnibeeApiMerchantInvoiceInvoiceSummary

`func NewUnibeeApiMerchantInvoiceInvoiceSummary() *UnibeeApiMerchantInvoiceInvoiceSummary`

NewUnibeeApiMerchantInvoiceInvoiceSummary instantiates a new UnibeeApiMerchantInvoiceInvoiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantInvoiceInvoiceSummaryWithDefaults

`func NewUnibeeApiMerchantInvoiceInvoiceSummaryWithDefaults() *UnibeeApiMerchantInvoiceInvoiceSummary`

NewUnibeeApiMerchantInvoiceInvoiceSummaryWithDefaults instantiates a new UnibeeApiMerchantInvoiceInvoiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBizType

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFinishTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetFinishTime() int64`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetFinishTimeOk() (*int64, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetFinishTime(v int64)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetInvoiceName

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetInvoiceName() string`

GetInvoiceName returns the InvoiceName field if non-nil, zero value otherwise.

### GetInvoiceNameOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetInvoiceNameOk() (*string, bool)`

GetInvoiceNameOk returns a tuple with the InvoiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceName

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetInvoiceName(v string)`

SetInvoiceName sets InvoiceName field to given value.

### HasInvoiceName

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasInvoiceName() bool`

HasInvoiceName returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetSendEmail

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetSendEmail() string`

GetSendEmail returns the SendEmail field if non-nil, zero value otherwise.

### GetSendEmailOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetSendEmailOk() (*string, bool)`

GetSendEmailOk returns a tuple with the SendEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendEmail

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetSendEmail(v string)`

SetSendEmail sets SendEmail field to given value.

### HasSendEmail

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasSendEmail() bool`

HasSendEmail returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantInvoiceInvoiceSummary) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


