# UnibeeApiMerchantCreditNewConfigReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | currency | 
**Description** | Pointer to **string** | description | [optional] 
**DiscountCodeExclusive** | Pointer to **int32** | discount code exclusive when purchase, default no, 0-no, 1-yes | [optional] 
**ExchangeRate** | Pointer to **int64** | keep two decimal places，scale &#x3D; 100, 1 currency &#x3D; 1 credit * (exchange_rate/100), no effect on main account type | [optional] 
**Logo** | Pointer to **string** | logo image base64, show when user purchase | [optional] 
**LogoUrl** | Pointer to **string** | logo url, show when user purchase | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** | meta_data(json) | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PayoutEnable** | Pointer to **int32** | credit account can used or payout in purchase or not, 0-no, 1-yes | [optional] 
**PreviewDefaultUsed** | Pointer to **int32** | is default used when in purchase preview, default no, 0-no, 1-yes | [optional] 
**RechargeEnable** | Pointer to **int32** | credit account can be recharged or not, 0-no, 1-yes | [optional] 
**Recurring** | Pointer to **int32** | apply to recurring, default no, 0-no,1-yes | [optional] 
**Type** | **int32** | type of credit account, 1-main account, 2-promo credit account | 

## Methods

### NewUnibeeApiMerchantCreditNewConfigReq

`func NewUnibeeApiMerchantCreditNewConfigReq(currency string, type_ int32, ) *UnibeeApiMerchantCreditNewConfigReq`

NewUnibeeApiMerchantCreditNewConfigReq instantiates a new UnibeeApiMerchantCreditNewConfigReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantCreditNewConfigReqWithDefaults

`func NewUnibeeApiMerchantCreditNewConfigReqWithDefaults() *UnibeeApiMerchantCreditNewConfigReq`

NewUnibeeApiMerchantCreditNewConfigReqWithDefaults instantiates a new UnibeeApiMerchantCreditNewConfigReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDescription

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetDiscountCodeExclusive() int32`

GetDiscountCodeExclusive returns the DiscountCodeExclusive field if non-nil, zero value otherwise.

### GetDiscountCodeExclusiveOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetDiscountCodeExclusiveOk() (*int32, bool)`

GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetDiscountCodeExclusive(v int32)`

SetDiscountCodeExclusive sets DiscountCodeExclusive field to given value.

### HasDiscountCodeExclusive

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasDiscountCodeExclusive() bool`

HasDiscountCodeExclusive returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLogo

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoUrl

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayoutEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetPreviewDefaultUsed() int32`

GetPreviewDefaultUsed returns the PreviewDefaultUsed field if non-nil, zero value otherwise.

### GetPreviewDefaultUsedOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetPreviewDefaultUsedOk() (*int32, bool)`

GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetPreviewDefaultUsed(v int32)`

SetPreviewDefaultUsed sets PreviewDefaultUsed field to given value.

### HasPreviewDefaultUsed

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasPreviewDefaultUsed() bool`

HasPreviewDefaultUsed returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.

### GetRecurring

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetRecurring() int32`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetRecurringOk() (*int32, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetRecurring(v int32)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *UnibeeApiMerchantCreditNewConfigReq) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiMerchantCreditNewConfigReq) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiMerchantCreditNewConfigReq) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


