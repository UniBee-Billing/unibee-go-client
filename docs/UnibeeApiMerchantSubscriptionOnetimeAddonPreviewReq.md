# UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonId** | **int64** | AddonId, id of one-time addon, the new payment will created base on the addon&#39;s amount&#39; | 
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | apply promo credit amount, auto compute if not specified | [optional] 
**DiscountAmount** | Pointer to **int32** | Amount of discount | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**DiscountPercentage** | Pointer to **int32** | Percentage of discount, 100&#x3D;1%, ignore if discountAmount specified | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId, use user&#39;s gateway if not provide | [optional] 
**GatewayPaymentType** | Pointer to **string** | Gateway Payment Type | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，custom data | [optional] 
**Quantity** | **int64** | Quantity, quantity of the new payment which one-time addon purchased | 
**SubscriptionId** | Pointer to **string** | SubscriptionId, id of subscription which addon will attached, either SubscriptionId or UserId needed, The only one active subscription of userId will attach the addon | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, default user&#39;s taxPercentage if not specified | [optional] 
**UserId** | Pointer to **int64** | UserId, either SubscriptionId or UserId needed, The only one active subscription will update if userId provide instead of subscriptionId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq(addonId int64, quantity int64, ) *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq`

NewUnibeeApiMerchantSubscriptionOnetimeAddonPreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetAddonId() int64`

GetAddonId returns the AddonId field if non-nil, zero value otherwise.

### GetAddonIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetAddonIdOk() (*int64, bool)`

GetAddonIdOk returns a tuple with the AddonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetAddonId(v int64)`

SetAddonId sets AddonId field to given value.


### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountPercentage() int32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetDiscountPercentageOk() (*int32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetDiscountPercentage(v int32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.


### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionOnetimeAddonPreviewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


