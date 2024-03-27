# UnibeeApiMerchantSubscriptionUpdateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**ConfirmCurrency** | Pointer to **string** | Currency To Be Confirmed，Get From Preview | [optional] 
**ConfirmTotalAmount** | Pointer to **int64** | TotalAmount To Be Confirmed，Get From Preview | [optional] 
**EffectImmediate** | Pointer to **int32** | Effect Immediate，1-Immediate，2-Next Period | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**NewPlanId** | **int64** | New PlanId | 
**ProrationDate** | **int64** | prorationDate date to start Proration，Get From Preview | 
**Quantity** | **int64** | Quantity | 
**SubscriptionId** | **string** | SubscriptionId | 

## Methods

### NewUnibeeApiMerchantSubscriptionUpdateReq

`func NewUnibeeApiMerchantSubscriptionUpdateReq(newPlanId int64, prorationDate int64, quantity int64, subscriptionId string, ) *UnibeeApiMerchantSubscriptionUpdateReq`

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

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

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

### HasConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasConfirmCurrency() bool`

HasConfirmCurrency returns a boolean if a field has been set.

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

### HasConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasConfirmTotalAmount() bool`

HasConfirmTotalAmount returns a boolean if a field has been set.

### GetEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetEffectImmediate() int32`

GetEffectImmediate returns the EffectImmediate field if non-nil, zero value otherwise.

### GetEffectImmediateOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetEffectImmediateOk() (*int32, bool)`

GetEffectImmediateOk returns a tuple with the EffectImmediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetEffectImmediate(v int32)`

SetEffectImmediate sets EffectImmediate field to given value.

### HasEffectImmediate

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasEffectImmediate() bool`

HasEffectImmediate returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanId() int64`

GetNewPlanId returns the NewPlanId field if non-nil, zero value otherwise.

### GetNewPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) GetNewPlanIdOk() (*int64, bool)`

GetNewPlanIdOk returns a tuple with the NewPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanId

`func (o *UnibeeApiMerchantSubscriptionUpdateReq) SetNewPlanId(v int64)`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


