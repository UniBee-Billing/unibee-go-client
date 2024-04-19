# UnibeeApiMerchantSubscriptionCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**ConfirmCurrency** | Pointer to **string** | Currency to verify if provide | [optional] 
**ConfirmTotalAmount** | Pointer to **int64** | TotalAmount to verify if provide | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**GatewayId** | **int64** | Id | 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**PaymentMethodId** | Pointer to **string** | PaymentMethodId | [optional] 
**PlanId** | **int64** | PlanId | 
**Quantity** | Pointer to **int64** | Quantity，Default 1 | [optional] 
**ReturnUrl** | Pointer to **string** | ReturnUrl | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, override subscription taxPercentage if provide | [optional] 
**UserId** | **int64** | UserId | 
**VatCountryCode** | Pointer to **string** | VatCountryCode, CountryName | [optional] 
**VatNumber** | Pointer to **string** | VatNumber | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionCreateReq

`func NewUnibeeApiMerchantSubscriptionCreateReq(gatewayId int64, planId int64, userId int64, ) *UnibeeApiMerchantSubscriptionCreateReq`

NewUnibeeApiMerchantSubscriptionCreateReq instantiates a new UnibeeApiMerchantSubscriptionCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCreateReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionCreateReqWithDefaults() *UnibeeApiMerchantSubscriptionCreateReq`

NewUnibeeApiMerchantSubscriptionCreateReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmCurrency() string`

GetConfirmCurrency returns the ConfirmCurrency field if non-nil, zero value otherwise.

### GetConfirmCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmCurrencyOk() (*string, bool)`

GetConfirmCurrencyOk returns a tuple with the ConfirmCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetConfirmCurrency(v string)`

SetConfirmCurrency sets ConfirmCurrency field to given value.

### HasConfirmCurrency

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasConfirmCurrency() bool`

HasConfirmCurrency returns a boolean if a field has been set.

### GetConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmTotalAmount() int64`

GetConfirmTotalAmount returns the ConfirmTotalAmount field if non-nil, zero value otherwise.

### GetConfirmTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetConfirmTotalAmountOk() (*int64, bool)`

GetConfirmTotalAmountOk returns a tuple with the ConfirmTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetConfirmTotalAmount(v int64)`

SetConfirmTotalAmount sets ConfirmTotalAmount field to given value.

### HasConfirmTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasConfirmTotalAmount() bool`

HasConfirmTotalAmount returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscount() UnibeeApiBeanExternalDiscountParam`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantSubscriptionCreateReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreateReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreateReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


