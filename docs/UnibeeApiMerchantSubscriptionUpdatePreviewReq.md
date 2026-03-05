# UnibeeApiMerchantSubscriptionUpdatePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | Optional. Target addon configuration after the update. | [optional] 
**ApplyPromoCredit** | Pointer to **bool** | Optional. Whether to apply available promo credit to the update preview. | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | Optional. Maximum promo credit amount to apply in preview. If omitted and applyPromoCredit is true, the system auto-computes the usable amount. | [optional] 
**Currency** | Pointer to **string** | Optional. Target currency for the updated subscription. If not provided, use current subscription currency. Currency change may be restricted when proration is required. | [optional] 
**DiscountCode** | Pointer to **string** | Optional. Discount code applied to the update preview. | [optional] 
**EffectImmediate** | Pointer to **int32** | Optional. When to apply the plan change: 1 &#x3D; immediate (current period, proration will be computed), 2 &#x3D; at next period (no proration in current period). | [optional] 
**GatewayId** | Pointer to **int64** | Optional. Gateway ID for previewing payment routing. If not provided, the subscription&#39;s current gateway is used for preview. | [optional] 
**NewPlanId** | **int64** | Required. Target plan ID to switch the subscription to. | 
**ProrationDate** | Pointer to **int32** | Optional. UTC timestamp at which proration should start when effectImmediate &#x3D; 1. Defaults to current time if not provided. | [optional] 
**Quantity** | Pointer to **int64** | Optional. Target quantity for the subscription. If &lt;&#x3D; 0, it defaults to 1. | [optional] 
**SubscriptionId** | **string** | Required. SubscriptionId to be upgraded or downgraded. | 

## Methods

### NewUnibeeApiMerchantSubscriptionUpdatePreviewReq

`func NewUnibeeApiMerchantSubscriptionUpdatePreviewReq(newPlanId int64, subscriptionId string, ) *UnibeeApiMerchantSubscriptionUpdatePreviewReq`

NewUnibeeApiMerchantSubscriptionUpdatePreviewReq instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionUpdatePreviewReq`

NewUnibeeApiMerchantSubscriptionUpdatePreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdatePreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanId() int64`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetNewPlanIdOk() (*int64, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetNewPlanId(v int64)`

SetNewPlanId sets NewPlanId field to given value.


### GetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetProrationDate() int32`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetProrationDateOk() (*int32, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetProrationDate(v int32)`

SetProrationDate sets ProrationDate field to given value.

### HasProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasProrationDate() bool`

HasProrationDate returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


