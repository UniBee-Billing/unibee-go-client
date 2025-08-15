# UnibeeApiMerchantSubscriptionUpdateRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **map[string]interface{}** |  | [optional] 
**InvoiceId** | Pointer to **string** | The unique id of invoice | [optional] 
**Link** | Pointer to **string** | The payment link, need redirect customer to link if paid&#x3D;false | [optional] 
**Note** | Pointer to **string** | note | [optional] 
**Paid** | Pointer to **bool** | Paid or not，true|false | [optional] 
**PaymentId** | Pointer to **string** | The unique id of payment | [optional] 
**SubscriptionPendingUpdate** | Pointer to [**UnibeeApiBeanDetailSubscriptionPendingUpdateDetail**](UnibeeApiBeanDetailSubscriptionPendingUpdateDetail.md) |  | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionUpdateRes

`func NewUnibeeApiMerchantSubscriptionUpdateRes() *UnibeeApiMerchantSubscriptionUpdateRes`

NewUnibeeApiMerchantSubscriptionUpdateRes instantiates a new UnibeeApiMerchantSubscriptionUpdateRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionUpdateResWithDefaults

`func NewUnibeeApiMerchantSubscriptionUpdateResWithDefaults() *UnibeeApiMerchantSubscriptionUpdateRes`

NewUnibeeApiMerchantSubscriptionUpdateResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdateRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetNote

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPaid

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetSubscriptionPendingUpdate() UnibeeApiBeanDetailSubscriptionPendingUpdateDetail`

GetSubscriptionPendingUpdate returns the SubscriptionPendingUpdate field if non-nil, zero value otherwise.

### GetSubscriptionPendingUpdateOk

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) GetSubscriptionPendingUpdateOk() (*UnibeeApiBeanDetailSubscriptionPendingUpdateDetail, bool)`

GetSubscriptionPendingUpdateOk returns a tuple with the SubscriptionPendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) SetSubscriptionPendingUpdate(v UnibeeApiBeanDetailSubscriptionPendingUpdateDetail)`

SetSubscriptionPendingUpdate sets SubscriptionPendingUpdate field to given value.

### HasSubscriptionPendingUpdate

`func (o *UnibeeApiMerchantSubscriptionUpdateRes) HasSubscriptionPendingUpdate() bool`

HasSubscriptionPendingUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


