# MerchantSubscriptionCreatePreviewPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonParams** | Pointer to [**[]UnibeeApiBeanPlanAddonParam**](UnibeeApiBeanPlanAddonParam.md) |  | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to [**UnibeeApiBeanGatewaySimplify**](UnibeeApiBeanGatewaySimplify.md) |  | [optional] 
**Invoice** | Pointer to [**UnibeeApiBeanInvoiceSimplify**](UnibeeApiBeanInvoiceSimplify.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlanSimplify**](UnibeeApiBeanPlanSimplify.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**TaxPercentage** | Pointer to **int64** |  | [optional] 
**TotalAmount** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**VatCountryCode** | Pointer to **string** |  | [optional] 
**VatCountryName** | Pointer to **string** |  | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**VatNumberValidate** | Pointer to [**UnibeeApiBeanValidResult**](UnibeeApiBeanValidResult.md) |  | [optional] 

## Methods

### NewMerchantSubscriptionCreatePreviewPost200ResponseData

`func NewMerchantSubscriptionCreatePreviewPost200ResponseData() *MerchantSubscriptionCreatePreviewPost200ResponseData`

NewMerchantSubscriptionCreatePreviewPost200ResponseData instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults

`func NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults() *MerchantSubscriptionCreatePreviewPost200ResponseData`

NewMerchantSubscriptionCreatePreviewPost200ResponseDataWithDefaults instantiates a new MerchantSubscriptionCreatePreviewPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonParams

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParams() []UnibeeApiBeanPlanAddonParam`

GetAddonParams returns the AddonParams field if non-nil, zero value otherwise.

### GetAddonParamsOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonParamsOk() (*[]UnibeeApiBeanPlanAddonParam, bool)`

GetAddonParamsOk returns a tuple with the AddonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonParams

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddonParams(v []UnibeeApiBeanPlanAddonParam)`

SetAddonParams sets AddonParams field to given value.

### HasAddonParams

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddonParams() bool`

HasAddonParams returns a boolean if a field has been set.

### GetAddons

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCurrency

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEmail

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGateway

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGateway() UnibeeApiBeanGatewaySimplify`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetGatewayOk() (*UnibeeApiBeanGatewaySimplify, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetGateway(v UnibeeApiBeanGatewaySimplify)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInvoice

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoice() UnibeeApiBeanInvoiceSimplify`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetInvoiceOk() (*UnibeeApiBeanInvoiceSimplify, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetInvoice(v UnibeeApiBeanInvoiceSimplify)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetPlan

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlan() UnibeeApiBeanPlanSimplify`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetPlanOk() (*UnibeeApiBeanPlanSimplify, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetPlan(v UnibeeApiBeanPlanSimplify)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetQuantity

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUserId

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.

### GetVatCountryName

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryName() string`

GetVatCountryName returns the VatCountryName field if non-nil, zero value otherwise.

### GetVatCountryNameOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatCountryNameOk() (*string, bool)`

GetVatCountryNameOk returns a tuple with the VatCountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryName

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatCountryName(v string)`

SetVatCountryName sets VatCountryName field to given value.

### HasVatCountryName

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatCountryName() bool`

HasVatCountryName returns a boolean if a field has been set.

### GetVatNumber

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetVatNumberValidate

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidate() UnibeeApiBeanValidResult`

GetVatNumberValidate returns the VatNumberValidate field if non-nil, zero value otherwise.

### GetVatNumberValidateOk

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) GetVatNumberValidateOk() (*UnibeeApiBeanValidResult, bool)`

GetVatNumberValidateOk returns a tuple with the VatNumberValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumberValidate

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) SetVatNumberValidate(v UnibeeApiBeanValidResult)`

SetVatNumberValidate sets VatNumberValidate field to given value.

### HasVatNumberValidate

`func (o *MerchantSubscriptionCreatePreviewPost200ResponseData) HasVatNumberValidate() bool`

HasVatNumberValidate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


