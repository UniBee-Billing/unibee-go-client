# UnibeeApiBeanCreditConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**DiscountCodeExclusive** | Pointer to **int32** | discount code exclusive when purchase, default no, 0-no, 1-yes | [optional] 
**ExchangeRate** | Pointer to **int64** | keep two decimal places，multiply by 100 saved, 1 currency &#x3D; 1 credit * (exchange_rate/100), main account fixed rate to 100 | [optional] 
**Id** | Pointer to **int64** | Id | [optional] 
**Logo** | Pointer to **string** | logo image base64, show when user purchase | [optional] 
**LogoUrl** | Pointer to **string** | logo url, show when user purchase | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MetaData** | Pointer to **map[string]map[string]interface{}** | meta_data(json) | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**PayoutEnable** | Pointer to **int32** | 0-no, 1-yes | [optional] 
**PreviewDefaultUsed** | Pointer to **int32** | is default used when in purchase preview, default no, 0-no, 1-yes | [optional] 
**RechargeEnable** | Pointer to **int32** | 0-no, 1-yes | [optional] 
**Recurring** | Pointer to **int32** | apply to recurring, default no, 0-no,1-yes | [optional] 
**TotalDecrementAmount** | Pointer to **int64** | the total decrement amount | [optional] 
**TotalIncrementAmount** | Pointer to **int64** | the total increment amount | [optional] 
**Type** | Pointer to **int32** | type of credit account, 1-main account, 2-promo credit account | [optional] 

## Methods

### NewUnibeeApiBeanCreditConfig

`func NewUnibeeApiBeanCreditConfig() *UnibeeApiBeanCreditConfig`

NewUnibeeApiBeanCreditConfig instantiates a new UnibeeApiBeanCreditConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanCreditConfigWithDefaults

`func NewUnibeeApiBeanCreditConfigWithDefaults() *UnibeeApiBeanCreditConfig`

NewUnibeeApiBeanCreditConfigWithDefaults instantiates a new UnibeeApiBeanCreditConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanCreditConfig) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanCreditConfig) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanCreditConfig) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanCreditConfig) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanCreditConfig) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanCreditConfig) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanCreditConfig) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanCreditConfig) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanCreditConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanCreditConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanCreditConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanCreditConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscountCodeExclusive

`func (o *UnibeeApiBeanCreditConfig) GetDiscountCodeExclusive() int32`

GetDiscountCodeExclusive returns the DiscountCodeExclusive field if non-nil, zero value otherwise.

### GetDiscountCodeExclusiveOk

`func (o *UnibeeApiBeanCreditConfig) GetDiscountCodeExclusiveOk() (*int32, bool)`

GetDiscountCodeExclusiveOk returns a tuple with the DiscountCodeExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCodeExclusive

`func (o *UnibeeApiBeanCreditConfig) SetDiscountCodeExclusive(v int32)`

SetDiscountCodeExclusive sets DiscountCodeExclusive field to given value.

### HasDiscountCodeExclusive

`func (o *UnibeeApiBeanCreditConfig) HasDiscountCodeExclusive() bool`

HasDiscountCodeExclusive returns a boolean if a field has been set.

### GetExchangeRate

`func (o *UnibeeApiBeanCreditConfig) GetExchangeRate() int64`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *UnibeeApiBeanCreditConfig) GetExchangeRateOk() (*int64, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *UnibeeApiBeanCreditConfig) SetExchangeRate(v int64)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *UnibeeApiBeanCreditConfig) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanCreditConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanCreditConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanCreditConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanCreditConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogo

`func (o *UnibeeApiBeanCreditConfig) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UnibeeApiBeanCreditConfig) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UnibeeApiBeanCreditConfig) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UnibeeApiBeanCreditConfig) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoUrl

`func (o *UnibeeApiBeanCreditConfig) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *UnibeeApiBeanCreditConfig) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *UnibeeApiBeanCreditConfig) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *UnibeeApiBeanCreditConfig) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanCreditConfig) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanCreditConfig) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanCreditConfig) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanCreditConfig) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeApiBeanCreditConfig) GetMetaData() map[string]map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeApiBeanCreditConfig) GetMetaDataOk() (*map[string]map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeApiBeanCreditConfig) SetMetaData(v map[string]map[string]interface{})`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeApiBeanCreditConfig) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiBeanCreditConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiBeanCreditConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiBeanCreditConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiBeanCreditConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPayoutEnable

`func (o *UnibeeApiBeanCreditConfig) GetPayoutEnable() int32`

GetPayoutEnable returns the PayoutEnable field if non-nil, zero value otherwise.

### GetPayoutEnableOk

`func (o *UnibeeApiBeanCreditConfig) GetPayoutEnableOk() (*int32, bool)`

GetPayoutEnableOk returns a tuple with the PayoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutEnable

`func (o *UnibeeApiBeanCreditConfig) SetPayoutEnable(v int32)`

SetPayoutEnable sets PayoutEnable field to given value.

### HasPayoutEnable

`func (o *UnibeeApiBeanCreditConfig) HasPayoutEnable() bool`

HasPayoutEnable returns a boolean if a field has been set.

### GetPreviewDefaultUsed

`func (o *UnibeeApiBeanCreditConfig) GetPreviewDefaultUsed() int32`

GetPreviewDefaultUsed returns the PreviewDefaultUsed field if non-nil, zero value otherwise.

### GetPreviewDefaultUsedOk

`func (o *UnibeeApiBeanCreditConfig) GetPreviewDefaultUsedOk() (*int32, bool)`

GetPreviewDefaultUsedOk returns a tuple with the PreviewDefaultUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewDefaultUsed

`func (o *UnibeeApiBeanCreditConfig) SetPreviewDefaultUsed(v int32)`

SetPreviewDefaultUsed sets PreviewDefaultUsed field to given value.

### HasPreviewDefaultUsed

`func (o *UnibeeApiBeanCreditConfig) HasPreviewDefaultUsed() bool`

HasPreviewDefaultUsed returns a boolean if a field has been set.

### GetRechargeEnable

`func (o *UnibeeApiBeanCreditConfig) GetRechargeEnable() int32`

GetRechargeEnable returns the RechargeEnable field if non-nil, zero value otherwise.

### GetRechargeEnableOk

`func (o *UnibeeApiBeanCreditConfig) GetRechargeEnableOk() (*int32, bool)`

GetRechargeEnableOk returns a tuple with the RechargeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRechargeEnable

`func (o *UnibeeApiBeanCreditConfig) SetRechargeEnable(v int32)`

SetRechargeEnable sets RechargeEnable field to given value.

### HasRechargeEnable

`func (o *UnibeeApiBeanCreditConfig) HasRechargeEnable() bool`

HasRechargeEnable returns a boolean if a field has been set.

### GetRecurring

`func (o *UnibeeApiBeanCreditConfig) GetRecurring() int32`

GetRecurring returns the Recurring field if non-nil, zero value otherwise.

### GetRecurringOk

`func (o *UnibeeApiBeanCreditConfig) GetRecurringOk() (*int32, bool)`

GetRecurringOk returns a tuple with the Recurring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurring

`func (o *UnibeeApiBeanCreditConfig) SetRecurring(v int32)`

SetRecurring sets Recurring field to given value.

### HasRecurring

`func (o *UnibeeApiBeanCreditConfig) HasRecurring() bool`

HasRecurring returns a boolean if a field has been set.

### GetTotalDecrementAmount

`func (o *UnibeeApiBeanCreditConfig) GetTotalDecrementAmount() int64`

GetTotalDecrementAmount returns the TotalDecrementAmount field if non-nil, zero value otherwise.

### GetTotalDecrementAmountOk

`func (o *UnibeeApiBeanCreditConfig) GetTotalDecrementAmountOk() (*int64, bool)`

GetTotalDecrementAmountOk returns a tuple with the TotalDecrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDecrementAmount

`func (o *UnibeeApiBeanCreditConfig) SetTotalDecrementAmount(v int64)`

SetTotalDecrementAmount sets TotalDecrementAmount field to given value.

### HasTotalDecrementAmount

`func (o *UnibeeApiBeanCreditConfig) HasTotalDecrementAmount() bool`

HasTotalDecrementAmount returns a boolean if a field has been set.

### GetTotalIncrementAmount

`func (o *UnibeeApiBeanCreditConfig) GetTotalIncrementAmount() int64`

GetTotalIncrementAmount returns the TotalIncrementAmount field if non-nil, zero value otherwise.

### GetTotalIncrementAmountOk

`func (o *UnibeeApiBeanCreditConfig) GetTotalIncrementAmountOk() (*int64, bool)`

GetTotalIncrementAmountOk returns a tuple with the TotalIncrementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIncrementAmount

`func (o *UnibeeApiBeanCreditConfig) SetTotalIncrementAmount(v int64)`

SetTotalIncrementAmount sets TotalIncrementAmount field to given value.

### HasTotalIncrementAmount

`func (o *UnibeeApiBeanCreditConfig) HasTotalIncrementAmount() bool`

HasTotalIncrementAmount returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanCreditConfig) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanCreditConfig) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanCreditConfig) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanCreditConfig) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


