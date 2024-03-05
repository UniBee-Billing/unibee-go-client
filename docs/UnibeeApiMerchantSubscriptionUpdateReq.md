# UnibeeApiMerchantSubscriptionUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo**](UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo.md) | addonParams | [optional] 
**ConfirmCurrency** | **string** | Currency To Be Confirmed，Get From Preview | 
**ConfirmTotalAmount** | **int64** | TotalAmount To Be Confirmed，Get From Preview | 
**NewPlanId** | **int32** | New PlanId | 
**ProrationDate** | **int64** | prorationDate date to start Proration，Get From Preview | 
**Quantity** | Pointer to **int64** | Quantity，Default 1 | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 
**WithImmediateEffect** | Pointer to **int32** | Effect Immediate，1-Immediate，2-Next Period | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionUpdateReq

`func NewUnibeeApiMerchantSubscriptionUpdateReq(confirmCurrency string, confirmTotalAmount int64, newPlanId int32, prorationDate int64, subscriptionId string, ) *UnibeeApiMerchantSubscriptionUpdateReq`

NewUnibeeApiMerchantSubscriptionUpdateReq instantiates a new UnibeeApiMerchantSubscriptionUpdateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionUpdateReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionUpdateReqWithDefaults() *UnibeeApiMerchantSubscriptionUpdateReq`

NewUnibeeApiMerchantSubscriptionUpdateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionUpdateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParams() []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParamsOk() (*[]UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetAddonParams(v []UnibeeInternalLogicGatewayRoSubscriptionPlanAddonParamRo)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmCurrency() string`

GetConfirmCurrency returns the ConfirmCurrency field if non-nil, zero value otherwise.

### GetConfirmCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmCurrencyOk() (*string, bool)`

GetConfirmCurrencyOk returns a tuple with the ConfirmCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetConfirmCurrency(v string)`

SetConfirmCurrency sets ConfirmCurrency field to given value.


### GetConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmTotalAmount() int64`

GetConfirmTotalAmount returns the ConfirmTotalAmount field if non-nil, zero value otherwise.

### GetConfirmTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetConfirmTotalAmountOk() (*int64, bool)`

GetConfirmTotalAmountOk returns a tuple with the ConfirmTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetConfirmTotalAmount(v int64)`

SetConfirmTotalAmount sets ConfirmTotalAmount field to given value.


### GetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanId() int32`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanIdOk() (*int32, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetNewPlanId(v int32)`

SetNewPlanId sets NewPlanId field to given value.


### GetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetProrationDate() int64`

GetProrationDate returns the ProrationDate field if non-nil, zero value otherwise.

### GetProrationDateOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetProrationDateOk() (*int64, bool)`

GetProrationDateOk returns a tuple with the ProrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrationDate

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetProrationDate(v int64)`

SetProrationDate sets ProrationDate field to given value.


### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetWithImmediateEffect() int32`

GetWithImmediateEffect returns the WithImmediateEffect field if non-nil, zero value otherwise.

### GetWithImmediateEffectOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetWithImmediateEffectOk() (*int32, bool)`

GetWithImmediateEffectOk returns a tuple with the WithImmediateEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetWithImmediateEffect(v int32)`

SetWithImmediateEffect sets WithImmediateEffect field to given value.

### HasWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasWithImmediateEffect() bool`

HasWithImmediateEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


