# UnibeeApiMerchantProductListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]UnibeeApiBeanProduct**](UnibeeApiBeanProduct.md) | Product Object List | [optional] 
**Total** | Pointer to **int32** | Total | [optional] 

## Methods

### NewUnibeeApiMerchantProductListRes

`func NewUnibeeApiMerchantProductListRes() *UnibeeApiMerchantProductListRes`

NewUnibeeApiMerchantProductListRes instantiates a new UnibeeApiMerchantProductListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantProductListResWithDefaults

`func NewUnibeeApiMerchantProductListResWithDefaults() *UnibeeApiMerchantProductListRes`

NewUnibeeApiMerchantProductListResWithDefaults instantiates a new UnibeeApiMerchantProductListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *UnibeeApiMerchantProductListRes) GetProducts() []UnibeeApiBeanProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *UnibeeApiMerchantProductListRes) GetProductsOk() (*[]UnibeeApiBeanProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *UnibeeApiMerchantProductListRes) SetProducts(v []UnibeeApiBeanProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *UnibeeApiMerchantProductListRes) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantProductListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantProductListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantProductListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantProductListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


