# UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**AddonId** | Pointer to **int64** | onetime addonId | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata | [optional] 
**Payment** | Pointer to [**UnibeeApiBeanPayment**](UnibeeApiBeanPayment.md) |  | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**Status** | Pointer to **int32** | status, 1-create, 2-paid, 3-cancel, 4-expired | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 
**User** | Pointer to [**UnibeeApiBeanUserAccount**](UnibeeApiBeanUserAccount.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail

`func NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail() *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail`

NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetail instantiates a new UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetailWithDefaults

`func NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetailWithDefaults() *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail`

NewUnibeeApiBeanDetailSubscriptionOnetimeAddonDetailWithDefaults instantiates a new UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddon() UnibeeApiBeanPlan`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonOk() (*UnibeeApiBeanPlan, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetAddon(v UnibeeApiBeanPlan)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddonId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonId() int64`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetAddonIdOk() (*int64, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetAddonId(v int64)`

SetAddonId sets AddonId field to given value.

### HasAddonId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasAddonId() bool`

HasAddonId returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPayment

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetPayment() UnibeeApiBeanPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetPaymentOk() (*UnibeeApiBeanPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetPayment(v UnibeeApiBeanPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetUser() UnibeeApiBeanUserAccount`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) GetUserOk() (*UnibeeApiBeanUserAccount, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) SetUser(v UnibeeApiBeanUserAccount)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiBeanDetailSubscriptionOnetimeAddonDetail) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


