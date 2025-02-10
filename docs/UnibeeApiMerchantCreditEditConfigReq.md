# UnibeeApiMerchantCreditEditConfigReq

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
**RechargeEnable** | Pointer to **int32** | credit account can recharge or not, default no, 0-no, 1-yes | [optional] 
**Recurring** | Pointer to **int32** | apply to recurring, default no, 0-no,1-yes | [optional] 
**Type** | **int32** | type of credit account, 1-main account, 2-promo credit account | 

## Methods

### NewUnibeeApiMerchantCreditEditConfigReq

`func NewUnibeeApiMerchantCreditEditConfigReq(currency string, type_ int32, ) *UnibeeApiMerchantCreditEditConfigReq`

NewUnibeeApiMerchantCreditEditConfigReq instantiates a new UnibeeApiMerchantCreditEditConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditEditConfigReqWithDefaults

`func NewUnibeeApiMerchantCreditEditConfigReqWithDefaults() *UnibeeApiMerchantCreditEditConfigReq`

NewUnibeeApiMerchantCreditEditConfigReqWithDefaults instantiates a new UnibeeApiMerchantCreditEditConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetDiscountCodeExclusive() int32`

GetDiscountCodeExclusive returns the DiscountCodeExclusive field if non-nil, zero value otherwise.

### GetDiscountCodeExclusiveOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetDiscountCodeExclusiveOk() (*int32, bool)`

GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetDiscountCodeExclusive(v int32)`

SetDiscountCodeExclusive sets DiscountCodeExclusive field to given value.

### HasDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasDiscountCodeExclusive() bool`

HasDiscountCodeExclusive returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetExchangeRate() int32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetExchangeRateOk() (*int32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetExchangeRate(v int32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLogo

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoUrl

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetPreviewDefaultUsed() int32`

GetPreviewDefaultUsed returns the PreviewDefaultUsed field if non-nil, zero value otherwise.

### GetPreviewDefaultUsedOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetPreviewDefaultUsedOk() (*int32, bool)`

GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetPreviewDefaultUsed(v int32)`

SetPreviewDefaultUsed sets PreviewDefaultUsed field to given value.

### HasPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasPreviewDefaultUsed() bool`

HasPreviewDefaultUsed returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.

### GetRecurring

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetRecurring() int32`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetRecurringOk() (*int32, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetRecurring(v int32)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *UnibeeApiMerchantCreditEditConfigReq) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantCreditEditConfigReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantCreditEditConfigReq) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


