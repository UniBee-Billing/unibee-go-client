# UnibeeApiBeanProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MetaData** | Pointer to **string** | meta_data(json) | [optional] 
**ProductName** | Pointer to **string** | PlanName | [optional] 
**Status** | Pointer to **int32** | status，1-active，2-inactive, default active | [optional] 

## Methods

### NewUnibeeApiBeanProduct

`func NewUnibeeApiBeanProduct() *UnibeeApiBeanProduct`

NewUnibeeApiBeanProduct instantiates a new UnibeeApiBeanProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanProductWithDefaults

`func NewUnibeeApiBeanProductWithDefaults() *UnibeeApiBeanProduct`

NewUnibeeApiBeanProductWithDefaults instantiates a new UnibeeApiBeanProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanProduct) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanProduct) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanProduct) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanProduct) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiBeanProduct) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiBeanProduct) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiBeanProduct) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiBeanProduct) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanProduct) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanProduct) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanProduct) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiBeanProduct) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiBeanProduct) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiBeanProduct) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiBeanProduct) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeApiBeanProduct) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeApiBeanProduct) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeApiBeanProduct) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeApiBeanProduct) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanProduct) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanProduct) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanProduct) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanProduct) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiBeanProduct) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiBeanProduct) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiBeanProduct) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiBeanProduct) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanProduct) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanProduct) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanProduct) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanProduct) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanProduct) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanProduct) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanProduct) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanProduct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


