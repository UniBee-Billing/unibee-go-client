# UnibeeApiMerchantSubscriptionOnetimeAddonNewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **int64** | AddonId, id of one-time addon, the new payment will created base on the addon&#39;s amount&#39; | 
**DiscountAmount** | Pointer to **int32** | Amount of discount | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**DiscountPercentage** | Pointer to **int32** | Percentage of discount, 100&#x3D;1%, ignore if discountAmount provide | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId, use subscription&#39;s gateway if not provide | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，custom data | [optional] 
**Quantity** | **int64** | Quantity, quantity of the new payment which one-time addon purchased | 
**ReturnUrl** | Pointer to **string** | ReturnUrl, the addon&#39;s payment will redirect based on the returnUrl provided when it&#39;s back from gateway side | [optional] 
**SubscriptionId** | **string** | SubscriptionId, id of subscription which addon will attached | 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, use subscription&#39;s taxPercentage if not provide | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq(addonId int64, quantity int64, subscriptionId string, ) *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonNewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonNewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetAddonId() int64`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetAddonIdOk() (*int64, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetAddonId(v int64)`

SetAddonId sets AddonId field to given value.


### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountPercentage() int32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetDiscountPercentageOk() (*int32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetDiscountPercentage(v int32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonNewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


