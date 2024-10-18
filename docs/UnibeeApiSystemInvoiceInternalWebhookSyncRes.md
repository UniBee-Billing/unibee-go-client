# UnibeeApiSystemInvoiceInternalWebhookSyncRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstId** | Pointer to **string** | The first invoiceId of sync data, only output when isSynchronous&#x3D;true | [optional] 
**LastId** | Pointer to **string** | The last invoiceId of sync data, only output when isSynchronous&#x3D;true | [optional] 
**Total** | Pointer to **int32** | The total of sync, only output when isSynchronous&#x3D;true | [optional] 

## Methods

### NewUnibeeApiSystemInvoiceInternalWebhookSyncRes

`func NewUnibeeApiSystemInvoiceInternalWebhookSyncRes() *UnibeeApiSystemInvoiceInternalWebhookSyncRes`

NewUnibeeApiSystemInvoiceInternalWebhookSyncRes instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemInvoiceInternalWebhookSyncResWithDefaults

`func NewUnibeeApiSystemInvoiceInternalWebhookSyncResWithDefaults() *UnibeeApiSystemInvoiceInternalWebhookSyncRes`

NewUnibeeApiSystemInvoiceInternalWebhookSyncResWithDefaults instantiates a new UnibeeApiSystemInvoiceInternalWebhookSyncRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetFirstId() string`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetFirstIdOk() (*string, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetFirstId(v string)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetLastId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetLastId() string`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetLastIdOk() (*string, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetLastId(v string)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiSystemInvoiceInternalWebhookSyncRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


