# UnibeeApiBeanSubscriptionOnetimeAddonSimplify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | Pointer to **int64** | onetime addonId | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PaymentId** | Pointer to **string** | PaymentId | [optional] 
**PaymentLink** | Pointer to **string** | PaymentLink | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | status, 1-create, 2-paid, 3-cancel, 4-expired | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 

## Methods

### NewUnibeeApiBeanSubscriptionOnetimeAddonSimplify

`func NewUnibeeApiBeanSubscriptionOnetimeAddonSimplify() *UnibeeApiBeanSubscriptionOnetimeAddonSimplify`

NewUnibeeApiBeanSubscriptionOnetimeAddonSimplify instantiates a new UnibeeApiBeanSubscriptionOnetimeAddonSimplify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanSubscriptionOnetimeAddonSimplifyWithDefaults

`func NewUnibeeApiBeanSubscriptionOnetimeAddonSimplifyWithDefaults() *UnibeeApiBeanSubscriptionOnetimeAddonSimplify`

NewUnibeeApiBeanSubscriptionOnetimeAddonSimplifyWithDefaults instantiates a new UnibeeApiBeanSubscriptionOnetimeAddonSimplify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetAddonId() int64`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetAddonIdOk() (*int64, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetAddonId(v int64)`

SetAddonId sets AddonId field to given value.

### HasAddonId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasAddonId() bool`

HasAddonId returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentLink

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanSubscriptionOnetimeAddonSimplify) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


