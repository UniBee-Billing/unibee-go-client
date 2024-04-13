# UnibeeApiMerchantSubscriptionCreatePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**GatewayId** | **int64** | Id | 
**PlanId** | **int64** | PlanId | 
**Quantity** | Pointer to **int64** | Quantity | [optional] 
**UserId** | **int64** | UserId | 
**VatCountryCode** | Pointer to **string** | VatCountryCode, CountryName | [optional] 
**VatNumber** | Pointer to **string** | VatNumber | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionCreatePreviewReq

`func NewUnibeeApiMerchantSubscriptionCreatePreviewReq(gatewayId int64, planId int64, userId int64, ) *UnibeeApiMerchantSubscriptionCreatePreviewReq`

NewUnibeeApiMerchantSubscriptionCreatePreviewReq instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCreatePreviewReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionCreatePreviewReqWithDefaults() *UnibeeApiMerchantSubscriptionCreatePreviewReq`

NewUnibeeApiMerchantSubscriptionCreatePreviewReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


