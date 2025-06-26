# UnibeeApiMerchantSubscriptionUpdateRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | Pointer to **string** | The payment link, need redirect customer to link if paid&#x3D;false | [optional] 
**Note** | Pointer to **string** | note | [optional] 
**Paid** | Pointer to **bool** | Paid or not，true|false | [optional] 
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


