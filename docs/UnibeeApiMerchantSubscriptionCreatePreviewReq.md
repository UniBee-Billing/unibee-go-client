# UnibeeApiMerchantSubscriptionCreatePreviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | apply promo credit amount, auto compute if not specified | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**Email** | Pointer to **string** | Email, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**FreeInInitialPeriod** | Pointer to **bool** | Is free or not for the first period, true or false | [optional] 
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

### NewUnibeeApiMerchantSubscriptionCreatePreviewReq

`func NewUnibeeApiMerchantSubscriptionCreatePreviewReq(planId int64, ) *UnibeeApiMerchantSubscriptionCreatePreviewReq`

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

### GetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

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

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFreeInInitialPeriod

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetFreeInInitialPeriod() bool`

GetFreeInInitialPeriod returns the FreeInInitialPeriod field if non-nil, zero value otherwise.

### GetFreeInInitialPeriodOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetFreeInInitialPeriodOk() (*bool, bool)`

GetFreeInInitialPeriodOk returns a tuple with the FreeInInitialPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeInInitialPeriod

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetFreeInInitialPeriod(v bool)`

SetFreeInInitialPeriod sets FreeInInitialPeriod field to given value.

### HasFreeInInitialPeriod

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasFreeInInitialPeriod() bool`

HasFreeInInitialPeriod returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

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

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetUser() UnibeeApiBeanNewUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) GetUserOk() (*UnibeeApiBeanNewUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) SetUser(v UnibeeApiBeanNewUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasUser() bool`

HasUser returns a boolean if a field has been set.

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

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

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


