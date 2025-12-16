# UnibeeApiCheckoutVatUSVATValidateAddressReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**UnibeeApiBeanUSAddress**](UnibeeApiBeanUSAddress.md) |  | [optional] 
**MerchantId** | **int64** |  | 

## Methods

### NewUnibeeApiCheckoutVatUSVATValidateAddressReq

`func NewUnibeeApiCheckoutVatUSVATValidateAddressReq(merchantId int64, ) *UnibeeApiCheckoutVatUSVATValidateAddressReq`

NewUnibeeApiCheckoutVatUSVATValidateAddressReq instantiates a new UnibeeApiCheckoutVatUSVATValidateAddressReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutVatUSVATValidateAddressReqWithDefaults

`func NewUnibeeApiCheckoutVatUSVATValidateAddressReqWithDefaults() *UnibeeApiCheckoutVatUSVATValidateAddressReq`

NewUnibeeApiCheckoutVatUSVATValidateAddressReqWithDefaults instantiates a new UnibeeApiCheckoutVatUSVATValidateAddressReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) GetAddress() UnibeeApiBeanUSAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) GetAddressOk() (*UnibeeApiBeanUSAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) SetAddress(v UnibeeApiBeanUSAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiCheckoutVatUSVATValidateAddressReq) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


