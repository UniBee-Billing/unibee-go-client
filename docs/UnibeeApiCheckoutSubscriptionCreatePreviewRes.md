# UnibeeApiCheckoutSubscriptionCreatePreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) |  | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) |  | [optional] 
**ApplyPromoCredit** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**DiscountMessage** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanDetailGateway**](UnibeeApiBeanDetailGateway.md) |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**OtherActiveSubscriptionId** | Pointer to **string** | other active or incomplete subscription id  | [optional] 
**OtherPendingCryptoSubscription** | Pointer to [**UnibeeApiBeanDetailSubscriptionDetail**](UnibeeApiBeanDetailSubscriptionDetail.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**SignIn** | Pointer to [**UnibeeApiBeanCheckoutSignIn**](UnibeeApiBeanCheckoutSignIn.md) |  | [optional] 
**SubscriptionAmountExcludingTax** | Pointer to **int64** |  | [optional] 
**TaxAmount** | Pointer to **int64** |  | [optional] 
**TaxPercentage** | Pointer to **int64** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**VatCountryCode** | Pointer to **string** |  | [optional] 
**VatCountryName** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**VatNumberValidate** | Pointer to [**UnibeeApiBeanValidResult**](UnibeeApiBeanValidResult.md) |  | [optional] 
**VatNumberValidateMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCreatePreviewRes

`func NewUnibeeApiCheckoutSubscriptionCreatePreviewRes() *UnibeeApiCheckoutSubscriptionCreatePreviewRes`

NewUnibeeApiCheckoutSubscriptionCreatePreviewRes instantiates a new UnibeeApiCheckoutSubscriptionCreatePreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCreatePreviewResWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCreatePreviewResWithDefaults() *UnibeeApiCheckoutSubscriptionCreatePreviewRes`

NewUnibeeApiCheckoutSubscriptionCreatePreviewResWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCreatePreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscountMessage() string`

GetDiscountMessage returns the DiscountMessage field if non-nil, zero value otherwise.

### GetDiscountMessageOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetDiscountMessageOk() (*string, bool)`

GetDiscountMessageOk returns a tuple with the DiscountMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetDiscountMessage(v string)`

SetDiscountMessage sets DiscountMessage field to given value.

### HasDiscountMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasDiscountMessage() bool`

HasDiscountMessage returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetGateway() UnibeeApiBeanDetailGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetGatewayOk() (*UnibeeApiBeanDetailGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetGateway(v UnibeeApiBeanDetailGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetOtherActiveSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionId() string`

GetOtherActiveSubscriptionId returns the OtherActiveSubscriptionId field if non-nil, zero value otherwise.

### GetOtherActiveSubscriptionIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionIdOk() (*string, bool)`

GetOtherActiveSubscriptionIdOk returns a tuple with the OtherActiveSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherActiveSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetOtherActiveSubscriptionId(v string)`

SetOtherActiveSubscriptionId sets OtherActiveSubscriptionId field to given value.

### HasOtherActiveSubscriptionId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasOtherActiveSubscriptionId() bool`

HasOtherActiveSubscriptionId returns a boolean if a field has been set.

### GetOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscription() UnibeeApiBeanDetailSubscriptionDetail`

GetOtherPendingCryptoSubscription returns the OtherPendingCryptoSubscription field if non-nil, zero value otherwise.

### GetOtherPendingCryptoSubscriptionOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool)`

GetOtherPendingCryptoSubscriptionOk returns a tuple with the OtherPendingCryptoSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetOtherPendingCryptoSubscription(v UnibeeApiBeanDetailSubscriptionDetail)`

SetOtherPendingCryptoSubscription sets OtherPendingCryptoSubscription field to given value.

### HasOtherPendingCryptoSubscription

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasOtherPendingCryptoSubscription() bool`

HasOtherPendingCryptoSubscription returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSignIn

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetSignIn() UnibeeApiBeanCheckoutSignIn`

GetSignIn returns the SignIn field if non-nil, zero value otherwise.

### GetSignInOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetSignInOk() (*UnibeeApiBeanCheckoutSignIn, bool)`

GetSignInOk returns a tuple with the SignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignIn

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetSignIn(v UnibeeApiBeanCheckoutSignIn)`

SetSignIn sets SignIn field to given value.

### HasSignIn

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasSignIn() bool`

HasSignIn returns a boolean if a field has been set.

### GetSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetSubscriptionAmountExcludingTax() int64`

GetSubscriptionAmountExcludingTax returns the SubscriptionAmountExcludingTax field if non-nil, zero value otherwise.

### GetSubscriptionAmountExcludingTaxOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetSubscriptionAmountExcludingTaxOk() (*int64, bool)`

GetSubscriptionAmountExcludingTaxOk returns a tuple with the SubscriptionAmountExcludingTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetSubscriptionAmountExcludingTax(v int64)`

SetSubscriptionAmountExcludingTax sets SubscriptionAmountExcludingTax field to given value.

### HasSubscriptionAmountExcludingTax

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasSubscriptionAmountExcludingTax() bool`

HasSubscriptionAmountExcludingTax returns a boolean if a field has been set.

### GetTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTaxAmount() int64`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTaxAmountOk() (*int64, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetTaxAmount(v int64)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatCountryName

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatCountryName() string`

GetVatCountryName returns the VatCountryName field if non-nil, zero value otherwise.

### GetVatCountryNameOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatCountryNameOk() (*string, bool)`

GetVatCountryNameOk returns a tuple with the VatCountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryName

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetVatCountryName(v string)`

SetVatCountryName sets VatCountryName field to given value.

### HasVatCountryName

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasVatCountryName() bool`

HasVatCountryName returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetVatNumberValidate

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumberValidate() UnibeeApiBeanValidResult`

GetVatNumberValidate returns the VatNumberValidate field if non-nil, zero value otherwise.

### GetVatNumberValidateOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool)`

GetVatNumberValidateOk returns a tuple with the VatNumberValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumberValidate

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetVatNumberValidate(v UnibeeApiBeanValidResult)`

SetVatNumberValidate sets VatNumberValidate field to given value.

### HasVatNumberValidate

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasVatNumberValidate() bool`

HasVatNumberValidate returns a boolean if a field has been set.

### GetVatNumberValidateMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumberValidateMessage() string`

GetVatNumberValidateMessage returns the VatNumberValidateMessage field if non-nil, zero value otherwise.

### GetVatNumberValidateMessageOk

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) GetVatNumberValidateMessageOk() (*string, bool)`

GetVatNumberValidateMessageOk returns a tuple with the VatNumberValidateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumberValidateMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) SetVatNumberValidateMessage(v string)`

SetVatNumberValidateMessage sets VatNumberValidateMessage field to given value.

### HasVatNumberValidateMessage

`func (o *UnibeeApiCheckoutSubscriptionCreatePreviewRes) HasVatNumberValidateMessage() bool`

HasVatNumberValidateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


