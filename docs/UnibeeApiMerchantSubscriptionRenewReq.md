# UnibeeApiMerchantSubscriptionRenewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyPromoCredit** | Pointer to **bool** | Optional. Whether to apply available promo credit to this renewal invoice. | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | Optional. Maximum promo credit amount to apply. If omitted and applyPromoCredit is true, the system auto-computes the usable amount. | [optional] 
**CancelUrl** | Pointer to **string** | Optional. URL to redirect the customer to when the renewal payment is cancelled or fails. | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | Optional. Discount or coupon code applied only to this renewal. Overrides the subscription&#39;s recurring discount for this invoice. | [optional] 
**GatewayId** | Pointer to **int32** | Optional. Payment gateway ID used for the renewal invoice. If omitted, the subscription&#39;s original gateway configuration is used. | [optional] 
**GatewayPaymentType** | Pointer to **string** | Optional. Payment type for the selected gateway, such as card, wallet, etc. | [optional] 
**ManualPayment** | Pointer to **bool** | Optional. If true, do not create an automatic payment for the renewal invoice; the invoice will be created in open status for manual collection. | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Optional. Custom metadata map that will be stored on the renewal invoice and subscription timeline. | [optional] 
**PaymentUIMode** | Pointer to **string** | Optional. Checkout UI mode: hosted | embedded | custom. Default is hosted. | [optional] 
**ProductData** | Pointer to [**UnibeeApiBeanPlanProductParam**](UnibeeApiBeanPlanProductParam.md) |  | [optional] 
**ProductId** | Pointer to **int64** | Optional. Product ID used together with userId when subscriptionId is not specified, to narrow down which subscription to renew. If 0, the system uses its default product selection rules. | [optional] 
**ReturnUrl** | Pointer to **string** | Optional. URL to redirect the customer to after successful renewal or payment completion. | [optional] 
**SubscriptionId** | Pointer to **string** | Optional. SubscriptionId to be renewed. Either subscriptionId or userId must be provided. When subscriptionId is omitted, the system first tries to find the latest active or incomplete subscription for the user (and productId if provided), otherwise falls back to the latest subscription. | [optional] 
**TaxPercentage** | Pointer to **int32** | Optional. External tax percentage override for the renewal invoice, in basis points (e.g. 1000 &#x3D; 10%%). Overrides the subscription taxPercentage when provided. | [optional] 
**UserId** | Pointer to **int64** | Optional. UserId associated with the subscription to renew. Either subscriptionId or userId must be provided. Used to locate the target subscription when subscriptionId is not provided. | [optional] 

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

### GetPaymentUIMode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetPaymentUIMode() string`

GetPaymentUIMode returns the PaymentUIMode field if non-nil, zero value otherwise.

### GetPaymentUIModeOk

`func (o *UnibeeApiMerchantSubscriptionRenewReq) GetPaymentUIModeOk() (*string, bool)`

GetPaymentUIModeOk returns a tuple with the PaymentUIMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentUIMode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) SetPaymentUIMode(v string)`

SetPaymentUIMode sets PaymentUIMode field to given value.

### HasPaymentUIMode

`func (o *UnibeeApiMerchantSubscriptionRenewReq) HasPaymentUIMode() bool`

HasPaymentUIMode returns a boolean if a field has been set.

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


