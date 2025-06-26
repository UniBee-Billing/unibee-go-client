# UnibeeApiCheckoutSubscriptionCreateReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) | addonParams | [optional] 
**ApplyPromoCredit** | Pointer to **bool** | apply promo credit or not | [optional] 
**ApplyPromoCreditAmount** | Pointer to **int32** | apply promo credit amount, auto compute if not specified | [optional] 
**CancelUrl** | Pointer to **string** | CancelUrl, back to cancelUrl if checkout cancelled | [optional] 
**ConfirmCurrency** | Pointer to **string** | Currency to verify if provide | [optional] 
**ConfirmTotalAmount** | Pointer to **int64** | TotalAmount to verify if provide | [optional] 
**Discount** | Pointer to [**UnibeeApiBeanExternalDiscountParam**](UnibeeApiBeanExternalDiscountParam.md) |  | [optional] 
**DiscountCode** | Pointer to **string** | DiscountCode | [optional] 
**Email** | Pointer to **string** | Email, one of ExternalUserId&amp;Email, UserId or User needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, one of ExternalUserId&amp;Email, UserId or User needed | [optional] 
**GatewayId** | Pointer to **int32** | GatewayId | [optional] 
**GatewayPaymentType** | Pointer to **string** | Gateway Payment Type | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**PaymentMethodId** | Pointer to **string** | PaymentMethodId | [optional] 
**PlanId** | **int64** | PlanId | 
**ProductData** | Pointer to [**UnibeeApiBeanPlanProductParam**](UnibeeApiBeanPlanProductParam.md) |  | [optional] 
**Quantity** | Pointer to **int64** | Quantity，Default 1 | [optional] 
**ReturnUrl** | Pointer to **string** | ReturnUrl, back to returnUrl if checkout completed | [optional] 
**StartIncomplete** | Pointer to **bool** | StartIncomplete, use now pay later, subscription will generate invoice and start with incomplete status if set | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage，1000 &#x3D; 10%, override subscription taxPercentage if provide | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**User** | Pointer to [**UnibeeApiBeanNewUser**](UnibeeApiBeanNewUser.md) |  | [optional] 
**UserId** | Pointer to **int64** | UserId | [optional] 
**VatCountryCode** | Pointer to **string** | VatCountryCode, CountryName | [optional] 
**VatNumber** | Pointer to **string** | VatNumber | [optional] 

## Methods

### NewUnibeeApiCheckoutSubscriptionCreateReq

`func NewUnibeeApiCheckoutSubscriptionCreateReq(planId int64, ) *UnibeeApiCheckoutSubscriptionCreateReq`

NewUnibeeApiCheckoutSubscriptionCreateReq instantiates a new UnibeeApiCheckoutSubscriptionCreateReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutSubscriptionCreateReqWithDefaults

`func NewUnibeeApiCheckoutSubscriptionCreateReqWithDefaults() *UnibeeApiCheckoutSubscriptionCreateReq`

NewUnibeeApiCheckoutSubscriptionCreateReqWithDefaults instantiates a new UnibeeApiCheckoutSubscriptionCreateReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetApplyPromoCredit() bool`

GetApplyPromoCredit returns the ApplyPromoCredit field if non-nil, zero value otherwise.

### GetApplyPromoCreditOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetApplyPromoCreditOk() (*bool, bool)`

GetApplyPromoCreditOk returns a tuple with the ApplyPromoCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetApplyPromoCredit(v bool)`

SetApplyPromoCredit sets ApplyPromoCredit field to given value.

### HasApplyPromoCredit

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasApplyPromoCredit() bool`

HasApplyPromoCredit returns a boolean if a field has been set.

### GetApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetApplyPromoCreditAmount() int32`

GetApplyPromoCreditAmount returns the ApplyPromoCreditAmount field if non-nil, zero value otherwise.

### GetApplyPromoCreditAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetApplyPromoCreditAmountOk() (*int32, bool)`

GetApplyPromoCreditAmountOk returns a tuple with the ApplyPromoCreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetApplyPromoCreditAmount(v int32)`

SetApplyPromoCreditAmount sets ApplyPromoCreditAmount field to given value.

### HasApplyPromoCreditAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasApplyPromoCreditAmount() bool`

HasApplyPromoCreditAmount returns a boolean if a field has been set.

### GetCancelUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetConfirmCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetConfirmCurrency() string`

GetConfirmCurrency returns the ConfirmCurrency field if non-nil, zero value otherwise.

### GetConfirmCurrencyOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetConfirmCurrencyOk() (*string, bool)`

GetConfirmCurrencyOk returns a tuple with the ConfirmCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetConfirmCurrency(v string)`

SetConfirmCurrency sets ConfirmCurrency field to given value.

### HasConfirmCurrency

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasConfirmCurrency() bool`

HasConfirmCurrency returns a boolean if a field has been set.

### GetConfirmTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetConfirmTotalAmount() int64`

GetConfirmTotalAmount returns the ConfirmTotalAmount field if non-nil, zero value otherwise.

### GetConfirmTotalAmountOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetConfirmTotalAmountOk() (*int64, bool)`

GetConfirmTotalAmountOk returns a tuple with the ConfirmTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetConfirmTotalAmount(v int64)`

SetConfirmTotalAmount sets ConfirmTotalAmount field to given value.

### HasConfirmTotalAmount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasConfirmTotalAmount() bool`

HasConfirmTotalAmount returns a boolean if a field has been set.

### GetDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetDiscount() UnibeeApiBeanExternalDiscountParam`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetDiscountOk() (*UnibeeApiBeanExternalDiscountParam, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetDiscount(v UnibeeApiBeanExternalDiscountParam)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.


### GetProductData

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetProductData() UnibeeApiBeanPlanProductParam`

GetProductData returns the ProductData field if non-nil, zero value otherwise.

### GetProductDataOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetProductDataOk() (*UnibeeApiBeanPlanProductParam, bool)`

GetProductDataOk returns a tuple with the ProductData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductData

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetProductData(v UnibeeApiBeanPlanProductParam)`

SetProductData sets ProductData field to given value.

### HasProductData

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasProductData() bool`

HasProductData returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStartIncomplete

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetStartIncomplete() bool`

GetStartIncomplete returns the StartIncomplete field if non-nil, zero value otherwise.

### GetStartIncompleteOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetStartIncompleteOk() (*bool, bool)`

GetStartIncompleteOk returns a tuple with the StartIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIncomplete

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetStartIncomplete(v bool)`

SetStartIncomplete sets StartIncomplete field to given value.

### HasStartIncomplete

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasStartIncomplete() bool`

HasStartIncomplete returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetUser

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetUser() UnibeeApiBeanNewUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetUserOk() (*UnibeeApiBeanNewUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetUser(v UnibeeApiBeanNewUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiCheckoutSubscriptionCreateReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


