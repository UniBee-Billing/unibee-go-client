# UnibeeApiCheckoutSubscriptionCreatePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**ApplyPromoCredit** | Pointer to **bool** |  | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | apply promo credit amount, auto compute if not specified | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**Email** | Pointer to **string** | Email, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId | [optional] 
**GatewayPaymentType** | Pointer to **string** | Gateway Payment Type | [optional] 
**PlanId** | **int64** | PlanId | 
**Quantity** | Pointer to **int64** | Quantity | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10% | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**User** | Pointer to [**UnibeeApiBeanNewUser**](UnibeeApiBeanNewUser.md) |  | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 
**VatCountryCode** | Pointer to **string** | VatCountryCode, CountryName | [optional] 
**VatNumber** | Pointer to **string** | VatNumber | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCreatePreviewReq

`func NewUnibeeApiCheckoutSubscriptionCreatePreviewReq(planId int64, ) *UnibeeApiCheckoutSubscriptionCreatePreviewReq`

NewUnibeeApiCheckoutSubscriptionCreatePreviewReq instantiates a new UnibeeApiCheckoutSubscriptionCreatePreviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCreatePreviewReqWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCreatePreviewReqWithDefaults() *UnibeeApiCheckoutSubscriptionCreatePreviewReq`

NewUnibeeApiCheckoutSubscriptionCreatePreviewReqWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCreatePreviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetUser() UnibeeApiBeanNewUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetUserOk() (*UnibeeApiBeanNewUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetUser(v UnibeeApiBeanNewUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


