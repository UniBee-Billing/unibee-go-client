# UnibeeApiMerchantSubscriptionCreatePreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) |  | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**DiscountAmount** | Pointer to **int64** |  | [optional] 
**DiscountMessage** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGateway**](UnibeeApiBeanGateway.md) |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoice**](UnibeeApiBeanInvoice.md) |  | [optional] 
**OriginAmount** | Pointer to **int64** |  | [optional] 
**OtherActiveSubscriptionId** | Pointer to **string** | other active or incomplete subscription id  | [optional] 
**OtherPendingCryptoSubscription** | Pointer to [**UnibeeApiBeanDetailSubscriptionDetail**](UnibeeApiBeanDetailSubscriptionDetail.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
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

### NewUnibeeApiMerchantSubscriptionCreatePreviewRes

`func NewUnibeeApiMerchantSubscriptionCreatePreviewRes() *UnibeeApiMerchantSubscriptionCreatePreviewRes`

NewUnibeeApiMerchantSubscriptionCreatePreviewRes instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionCreatePreviewResWithDefaults

`func NewUnibeeApiMerchantSubscriptionCreatePreviewResWithDefaults() *UnibeeApiMerchantSubscriptionCreatePreviewRes`

NewUnibeeApiMerchantSubscriptionCreatePreviewResWithDefaults instantiates a new UnibeeApiMerchantSubscriptionCreatePreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscount() UnibeeApiBeanMerchantDiscountCode`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscount(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountMessage() string`

GetDiscountMessage returns the DiscountMessage field if non-nil, zero value otherwise.

### GetDiscountMessageOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetDiscountMessageOk() (*string, bool)`

GetDiscountMessageOk returns a tuple with the DiscountMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetDiscountMessage(v string)`

SetDiscountMessage sets DiscountMessage field to given value.

### HasDiscountMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasDiscountMessage() bool`

HasDiscountMessage returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetGateway() UnibeeApiBeanGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetGatewayOk() (*UnibeeApiBeanGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetGateway(v UnibeeApiBeanGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInvoice

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetInvoice() UnibeeApiBeanInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetInvoiceOk() (*UnibeeApiBeanInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetInvoice(v UnibeeApiBeanInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOriginAmount() int64`

GetOriginAmount returns the OriginAmount field if non-nil, zero value otherwise.

### GetOriginAmountOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOriginAmountOk() (*int64, bool)`

GetOriginAmountOk returns a tuple with the OriginAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOriginAmount(v int64)`

SetOriginAmount sets OriginAmount field to given value.

### HasOriginAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOriginAmount() bool`

HasOriginAmount returns a boolean if a field has been set.

### GetOtherActiveSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionId() string`

GetOtherActiveSubscriptionId returns the OtherActiveSubscriptionId field if non-nil, zero value otherwise.

### GetOtherActiveSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherActiveSubscriptionIdOk() (*string, bool)`

GetOtherActiveSubscriptionIdOk returns a tuple with the OtherActiveSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherActiveSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOtherActiveSubscriptionId(v string)`

SetOtherActiveSubscriptionId sets OtherActiveSubscriptionId field to given value.

### HasOtherActiveSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOtherActiveSubscriptionId() bool`

HasOtherActiveSubscriptionId returns a boolean if a field has been set.

### GetOtherPendingCryptoSubscription

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscription() UnibeeApiBeanDetailSubscriptionDetail`

GetOtherPendingCryptoSubscription returns the OtherPendingCryptoSubscription field if non-nil, zero value otherwise.

### GetOtherPendingCryptoSubscriptionOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetOtherPendingCryptoSubscriptionOk() (*UnibeeApiBeanDetailSubscriptionDetail, bool)`

GetOtherPendingCryptoSubscriptionOk returns a tuple with the OtherPendingCryptoSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPendingCryptoSubscription

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetOtherPendingCryptoSubscription(v UnibeeApiBeanDetailSubscriptionDetail)`

SetOtherPendingCryptoSubscription sets OtherPendingCryptoSubscription field to given value.

### HasOtherPendingCryptoSubscription

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasOtherPendingCryptoSubscription() bool`

HasOtherPendingCryptoSubscription returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatCountryName

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryName() string`

GetVatCountryName returns the VatCountryName field if non-nil, zero value otherwise.

### GetVatCountryNameOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatCountryNameOk() (*string, bool)`

GetVatCountryNameOk returns a tuple with the VatCountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryName

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatCountryName(v string)`

SetVatCountryName sets VatCountryName field to given value.

### HasVatCountryName

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatCountryName() bool`

HasVatCountryName returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetVatNumberValidate

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidate() UnibeeApiBeanValidResult`

GetVatNumberValidate returns the VatNumberValidate field if non-nil, zero value otherwise.

### GetVatNumberValidateOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool)`

GetVatNumberValidateOk returns a tuple with the VatNumberValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumberValidate

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumberValidate(v UnibeeApiBeanValidResult)`

SetVatNumberValidate sets VatNumberValidate field to given value.

### HasVatNumberValidate

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumberValidate() bool`

HasVatNumberValidate returns a boolean if a field has been set.

### GetVatNumberValidateMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateMessage() string`

GetVatNumberValidateMessage returns the VatNumberValidateMessage field if non-nil, zero value otherwise.

### GetVatNumberValidateMessageOk

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) GetVatNumberValidateMessageOk() (*string, bool)`

GetVatNumberValidateMessageOk returns a tuple with the VatNumberValidateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumberValidateMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) SetVatNumberValidateMessage(v string)`

SetVatNumberValidateMessage sets VatNumberValidateMessage field to given value.

### HasVatNumberValidateMessage

`func (o *UnibeeApiMerchantSubscriptionCreatePreviewRes) HasVatNumberValidateMessage() bool`

HasVatNumberValidateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


