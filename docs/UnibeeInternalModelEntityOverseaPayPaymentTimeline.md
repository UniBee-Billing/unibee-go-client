# UnibeeInternalModelEntityOverseaPayPaymentTimeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**GatewayId** | Pointer to **int32** | gateway id | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**MerchantId** | Pointer to **int32** | merchant id | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**Status** | Pointer to **int32** | 0-pending, 1-success, 2-failure | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TimelineType** | Pointer to **int32** | 0-pay, 1-refund | [optional] 
**TotalAmount** | Pointer to **int64** | total amount | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayPaymentTimeline

`func NewUnibeeInternalModelEntityOverseaPayPaymentTimeline() *UnibeeInternalModelEntityOverseaPayPaymentTimeline`

NewUnibeeInternalModelEntityOverseaPayPaymentTimeline instantiates a new UnibeeInternalModelEntityOverseaPayPaymentTimeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayPaymentTimelineWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayPaymentTimelineWithDefaults() *UnibeeInternalModelEntityOverseaPayPaymentTimeline`

NewUnibeeInternalModelEntityOverseaPayPaymentTimelineWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayPaymentTimeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetMerchantId() int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetMerchantIdOk() (*int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetMerchantId(v int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTimelineType

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetTimelineType() int32`

GetTimelineType returns the TimelineType field if non-nil, zero value otherwise.

### GetTimelineTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetTimelineTypeOk() (*int32, bool)`

GetTimelineTypeOk returns a tuple with the TimelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineType

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetTimelineType(v int32)`

SetTimelineType sets TimelineType field to given value.

### HasTimelineType

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasTimelineType() bool`

HasTimelineType returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPayPaymentTimeline) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


