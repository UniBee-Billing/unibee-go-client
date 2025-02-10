# UnibeeApiMerchantCreditEditPromoConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | currency | 
**Description** | Pointer to **string** | description | [optional] 
**DiscountCodeExclusive** | Pointer to **int32** | discount code exclusive when purchase, default no, 0-no, 1-yes | [optional] 
**ExchangeRate** | Pointer to **int32** | keep two decimal places，scale &#x3D; 100, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 
**Logo** | Pointer to **string** | logo image base64, show when user purchase | [optional] 
**LogoUrl** | Pointer to **string** | logo url, show when user purchase | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** | meta_data(json) | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PayoutEnable** | Pointer to **int32** | credit account can payout or not, default no, 0-no, 1-yes | [optional] 
**PreviewDefaultUsed** | Pointer to **int32** | is default used when in purchase preview, default no, 0-no, 1-yes | [optional] 
**RechargeEnable** | Pointer to **int32** | credit account can be recharged or not, 0-no, 1-yes | [optional] 
**Recurring** | Pointer to **int32** | apply to recurring, default no, 0-no,1-yes | [optional] 

## Methods

### NewUnibeeApiMerchantCreditEditPromoConfigReq

`func NewUnibeeApiMerchantCreditEditPromoConfigReq(currency string, ) *UnibeeApiMerchantCreditEditPromoConfigReq`

NewUnibeeApiMerchantCreditEditPromoConfigReq instantiates a new UnibeeApiMerchantCreditEditPromoConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditEditPromoConfigReqWithDefaults

`func NewUnibeeApiMerchantCreditEditPromoConfigReqWithDefaults() *UnibeeApiMerchantCreditEditPromoConfigReq`

NewUnibeeApiMerchantCreditEditPromoConfigReqWithDefaults instantiates a new UnibeeApiMerchantCreditEditPromoConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDiscountCodeExclusive() int32`

GetDiscountCodeExclusive returns the DiscountCodeExclusive field if non-nil, zero value otherwise.

### GetDiscountCodeExclusiveOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetDiscountCodeExclusiveOk() (*int32, bool)`

GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetDiscountCodeExclusive(v int32)`

SetDiscountCodeExclusive sets DiscountCodeExclusive field to given value.

### HasDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasDiscountCodeExclusive() bool`

HasDiscountCodeExclusive returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetExchangeRate() int32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetExchangeRateOk() (*int32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetExchangeRate(v int32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLogo

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoUrl

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPreviewDefaultUsed() int32`

GetPreviewDefaultUsed returns the PreviewDefaultUsed field if non-nil, zero value otherwise.

### GetPreviewDefaultUsedOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetPreviewDefaultUsedOk() (*int32, bool)`

GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetPreviewDefaultUsed(v int32)`

SetPreviewDefaultUsed sets PreviewDefaultUsed field to given value.

### HasPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasPreviewDefaultUsed() bool`

HasPreviewDefaultUsed returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.

### GetRecurring

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRecurring() int32`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) GetRecurringOk() (*int32, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) SetRecurring(v int32)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *UnibeeApiMerchantCreditEditPromoConfigReq) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


