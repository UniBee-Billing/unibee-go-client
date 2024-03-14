# UnibeeApiMerchantSubscriptionUpdatePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**GatewayId** | Pointer to **int64** | Id | [optional] 
**NewPlanId** | **int64** | New PlanId | 
**Quantity** | Pointer to **int64** | Quantity，Default 1 | [optional] 
**SubscriptionId** | **string** | SubscriptionId | 
**WithImmediateEffect** | Pointer to **int32** | Effect Immediate，1-Immediate，2-Next Period | [optional] 

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


### GetWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetWithImmediateEffect() int32`

GetWithImmediateEffect returns the WithImmediateEffect field if non-nil, zero value otherwise.

### GetWithImmediateEffectOk

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) GetWithImmediateEffectOk() (*int32, bool)`

GetWithImmediateEffectOk returns a tuple with the WithImmediateEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) SetWithImmediateEffect(v int32)`

SetWithImmediateEffect sets WithImmediateEffect field to given value.

### HasWithImmediateEffect

`func (o *UnibeeApiMerchantSubscriptionUpdatePreviewReq) HasWithImmediateEffect() bool`

HasWithImmediateEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


