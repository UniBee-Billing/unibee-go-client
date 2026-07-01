# UnibeeApiBeanPublicProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**ProductName** | Pointer to **string** | ProductName | [optional] 
**Status** | Pointer to **int32** | status，1-active，2-inactive, default active | [optional] 
**UsVATConfig** | Pointer to [**UnibeeApiBeanPublicUSVATConfig**](UnibeeApiBeanPublicUSVATConfig.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanPublicProduct

`func NewUnibeeApiBeanPublicProduct() *UnibeeApiBeanPublicProduct`

NewUnibeeApiBeanPublicProduct instantiates a new UnibeeApiBeanPublicProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPublicProductWithDefaults

`func NewUnibeeApiBeanPublicProductWithDefaults() *UnibeeApiBeanPublicProduct`

NewUnibeeApiBeanPublicProductWithDefaults instantiates a new UnibeeApiBeanPublicProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UnibeeApiBeanPublicProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanPublicProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanPublicProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanPublicProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiBeanPublicProduct) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiBeanPublicProduct) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiBeanPublicProduct) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiBeanPublicProduct) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPublicProduct) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPublicProduct) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPublicProduct) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPublicProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiBeanPublicProduct) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiBeanPublicProduct) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiBeanPublicProduct) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiBeanPublicProduct) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPublicProduct) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPublicProduct) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPublicProduct) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPublicProduct) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetProductName

`func (o *UnibeeApiBeanPublicProduct) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *UnibeeApiBeanPublicProduct) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *UnibeeApiBeanPublicProduct) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *UnibeeApiBeanPublicProduct) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPublicProduct) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPublicProduct) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPublicProduct) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPublicProduct) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsVATConfig

`func (o *UnibeeApiBeanPublicProduct) GetUsVATConfig() UnibeeApiBeanPublicUSVATConfig`

GetUsVATConfig returns the UsVATConfig field if non-nil, zero value otherwise.

### GetUsVATConfigOk

`func (o *UnibeeApiBeanPublicProduct) GetUsVATConfigOk() (*UnibeeApiBeanPublicUSVATConfig, bool)`

GetUsVATConfigOk returns a tuple with the UsVATConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVATConfig

`func (o *UnibeeApiBeanPublicProduct) SetUsVATConfig(v UnibeeApiBeanPublicUSVATConfig)`

SetUsVATConfig sets UsVATConfig field to given value.

### HasUsVATConfig

`func (o *UnibeeApiBeanPublicProduct) HasUsVATConfig() bool`

HasUsVATConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


