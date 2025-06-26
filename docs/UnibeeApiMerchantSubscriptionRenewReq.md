# UnibeeApiMerchantSubscriptionRenewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | apply promo credit amount, auto compute if not specified | [optional] 
**CancelUrl** | Pointer to **string** | CancelUrl, back to cancelUrl if renew cancelled | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode, override subscription discount | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId, use subscription&#39;s gateway if not provide | [optional] 
**GatewayPaymentType** | Pointer to **string** | Gateway Payment Type | [optional] 
**ManualPayment** | Pointer to **bool** | ManualPayment | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**ProductData** | Pointer to [**UnibeeApiBeanPlanProductParam**](UnibeeApiBeanPlanProductParam.md) |  | [optional] 
**ProductId** | Pointer to **int64** | default product will use if not specified | [optional] 
**ReturnUrl** | Pointer to **string** | ReturnUrl, back to returnUrl if renew completed | [optional] 
**SubscriptionId** | Pointer to **string** | SubscriptionId, id of subscription which addon will attached, either SubscriptionId or UserId needed, The only one active subscription or latest subscription will renew if userId provide instead of subscriptionId | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, override subscription taxPercentage if provide | [optional] 
**UserId** | Pointer to **int64** | UserId, either SubscriptionId or UserId needed, The only one active subscription or latest cancel|expire subscription will renew if userId provide instead of subscriptionId | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionRenewReq

`func NewUnibeeApiMerchantSubscriptionRenewReq() *UnibeeApiMerchantSubscriptionRenewReq`

NewUnibeeApiMerchantSubscriptionRenewReq instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults() *UnibeeApiMerchantSubscriptionRenewReq`

NewUnibeeApiMerchantSubscriptionRenewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionRenewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetCancelUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscount() UnibeeApiBeanExternalDiscountParam`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetManualPayment

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetManualPayment() bool`

GetManualPayment returns the ManualPayment field if non-nil, zero value otherwise.

### GetManualPaymentOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetManualPaymentOk() (*bool, bool)`

GetManualPaymentOk returns a tuple with the ManualPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualPayment

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetManualPayment(v bool)`

SetManualPayment sets ManualPayment field to given value.

### HasManualPayment

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasManualPayment() bool`

HasManualPayment returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProductData

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductData() UnibeeApiBeanPlanProductParam`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductDataOk() (*UnibeeApiBeanPlanProductParam, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetProductData(v UnibeeApiBeanPlanProductParam)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


