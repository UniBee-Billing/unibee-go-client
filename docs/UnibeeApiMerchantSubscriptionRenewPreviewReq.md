# UnibeeApiMerchantSubscriptionRenewPreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCredit** | Pointer to **bool** | Optional. Whether to apply available promo credit to this renewal invoice. | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | Optional. Maximum promo credit amount to apply. If omitted and applyPromoCredit is true, the system auto-computes the usable amount. | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | Optional. Discount or coupon code applied only to this renewal. Overrides the subscription&#39;s recurring discount for this invoice. | [optional] 
**GatewayId** | Pointer to **int32** | Optional. Payment gateway ID used for the renewal invoice. If omitted, the subscription&#39;s original gateway configuration is used. | [optional] 
**GatewayPaymentType** | Pointer to **string** | Optional. Payment type for the selected gateway, such as card, wallet, etc. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Optional. Custom metadata map that will be stored on the renewal invoice and subscription timeline. | [optional] 
**ProductData** | Pointer to [**UnibeeApiBeanPlanProductParam**](UnibeeApiBeanPlanProductParam.md) |  | [optional] 
**ProductId** | Pointer to **int64** | Optional. Product ID used together with userId when subscriptionId is not specified, to narrow down which subscription to renew. If 0, the system uses its default product selection rules. | [optional] 
**SubscriptionId** | Pointer to **string** | Optional. SubscriptionId to be renewed. Either subscriptionId or userId must be provided. When subscriptionId is omitted, the system first tries to find the latest active or incomplete subscription for the user (and productId if provided), otherwise falls back to the latest subscription. | [optional] 
**TaxPercentage** | Pointer to **int32** | Optional. External tax percentage override for the renewal invoice, in basis points (e.g. 1000 &#x3D; 10%%). Overrides the subscription taxPercentage when provided. | [optional] 
**UserId** | Pointer to **int64** | Optional. UserId associated with the subscription to renew. Either subscriptionId or userId must be provided. Used to locate the target subscription when subscriptionId is not provided. | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionRenewPreviewReq

`func NewUnibeeApiMerchantSubscriptionRenewPreviewReq() *UnibeeApiMerchantSubscriptionRenewPreviewReq`

NewUnibeeApiMerchantSubscriptionRenewPreviewReq instantiates a new UnibeeApiMerchantSubscriptionRenewPreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionRenewPreviewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionRenewPreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionRenewPreviewReq`

NewUnibeeApiMerchantSubscriptionRenewPreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewPreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetDiscount() UnibeeApiBeanExternalDiscountParam`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProductData

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetProductData() UnibeeApiBeanPlanProductParam`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetProductDataOk() (*UnibeeApiBeanPlanProductParam, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetProductData(v UnibeeApiBeanPlanProductParam)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionRenewPreviewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


