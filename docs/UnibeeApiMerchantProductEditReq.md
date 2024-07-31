# UnibeeApiMerchantProductEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**ProductId** | **int64** | Id of product | 
**ProductName** | Pointer to **string** | ProductName | [optional] 
**Status** | Pointer to **int32** | status，1-active，2-inactive, default active | [optional] 

## Methods

### NewUnibeeApiMerchantProductEditReq

`func NewUnibeeApiMerchantProductEditReq(productId int64, ) *UnibeeApiMerchantProductEditReq`

NewUnibeeApiMerchantProductEditReq instantiates a new UnibeeApiMerchantProductEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProductEditReqWithDefaults

`func NewUnibeeApiMerchantProductEditReqWithDefaults() *UnibeeApiMerchantProductEditReq`

NewUnibeeApiMerchantProductEditReqWithDefaults instantiates a new UnibeeApiMerchantProductEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnibeeApiMerchantProductEditReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantProductEditReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantProductEditReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantProductEditReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiMerchantProductEditReq) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiMerchantProductEditReq) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiMerchantProductEditReq) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiMerchantProductEditReq) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiMerchantProductEditReq) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiMerchantProductEditReq) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiMerchantProductEditReq) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiMerchantProductEditReq) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantProductEditReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantProductEditReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantProductEditReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantProductEditReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantProductEditReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantProductEditReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantProductEditReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *UnibeeApiMerchantProductEditReq) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiMerchantProductEditReq) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiMerchantProductEditReq) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiMerchantProductEditReq) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantProductEditReq) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantProductEditReq) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantProductEditReq) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantProductEditReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


