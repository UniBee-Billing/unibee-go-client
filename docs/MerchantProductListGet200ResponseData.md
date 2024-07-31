# MerchantProductListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]UnibeeApiBeanProduct**](UnibeeApiBeanProduct.md) | Product Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewMerchantProductListGet200ResponseData

`func NewMerchantProductListGet200ResponseData() *MerchantProductListGet200ResponseData`

NewMerchantProductListGet200ResponseData instantiates a new MerchantProductListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantProductListGet200ResponseDataWithDefaults

`func NewMerchantProductListGet200ResponseDataWithDefaults() *MerchantProductListGet200ResponseData`

NewMerchantProductListGet200ResponseDataWithDefaults instantiates a new MerchantProductListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *MerchantProductListGet200ResponseData) GetProducts() []UnibeeApiBeanProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *MerchantProductListGet200ResponseData) GetProductsOk() (*[]UnibeeApiBeanProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *MerchantProductListGet200ResponseData) SetProducts(v []UnibeeApiBeanProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *MerchantProductListGet200ResponseData) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantProductListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantProductListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantProductListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantProductListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


