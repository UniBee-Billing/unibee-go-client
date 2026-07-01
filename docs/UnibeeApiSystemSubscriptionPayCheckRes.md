# UnibeeApiSystemSubscriptionPayCheckRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestInvoiceId** | Pointer to **string** | Latest invoice id of the subscription | [optional] 
**PayStatus** | Pointer to **int32** | Pay Status，1-Pending，2-Active(Paid)，3-PendingInActive，4-Cancel, 5-Expired, 6-Suspended, 8-Processing, 9-Failed | [optional] 

## Methods

### NewUnibeeApiSystemSubscriptionPayCheckRes

`func NewUnibeeApiSystemSubscriptionPayCheckRes() *UnibeeApiSystemSubscriptionPayCheckRes`

NewUnibeeApiSystemSubscriptionPayCheckRes instantiates a new UnibeeApiSystemSubscriptionPayCheckRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiSystemSubscriptionPayCheckResWithDefaults

`func NewUnibeeApiSystemSubscriptionPayCheckResWithDefaults() *UnibeeApiSystemSubscriptionPayCheckRes`

NewUnibeeApiSystemSubscriptionPayCheckResWithDefaults instantiates a new UnibeeApiSystemSubscriptionPayCheckRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestInvoiceId

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) GetLatestInvoiceId() string`

GetLatestInvoiceId returns the LatestInvoiceId field if non-nil, zero value otherwise.

### GetLatestInvoiceIdOk

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) GetLatestInvoiceIdOk() (*string, bool)`

GetLatestInvoiceIdOk returns a tuple with the LatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoiceId

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) SetLatestInvoiceId(v string)`

SetLatestInvoiceId sets LatestInvoiceId field to given value.

### HasLatestInvoiceId

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) HasLatestInvoiceId() bool`

HasLatestInvoiceId returns a boolean if a field has been set.

### GetPayStatus

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) GetPayStatus() int32`

GetPayStatus returns the PayStatus field if non-nil, zero value otherwise.

### GetPayStatusOk

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) GetPayStatusOk() (*int32, bool)`

GetPayStatusOk returns a tuple with the PayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayStatus

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) SetPayStatus(v int32)`

SetPayStatus sets PayStatus field to given value.

### HasPayStatus

`func (o *UnibeeApiSystemSubscriptionPayCheckRes) HasPayStatus() bool`

HasPayStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


