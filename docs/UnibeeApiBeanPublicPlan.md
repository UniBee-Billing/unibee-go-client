# UnibeeApiBeanPublicPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** | amount, cent, without tax | [optional] 
**AutoUpgradePlanId** | Pointer to **int64** | Target plan id for auto-upgrade (non-sensitive relation id); 0 if unset | [optional] 
**BindingAddonIds** | Pointer to **string** | binded recurring addon planIds，split with , (non-sensitive relation ids) | [optional] 
**BindingOnetimeAddonIds** | Pointer to **string** | binded onetime addon planIds，split with , (non-sensitive relation ids) | [optional] 
**CheckoutTemplateId** | Pointer to **int64** | Checkout template id (non-sensitive relation id; checkoutUrl already derives from it) | [optional] 
**CheckoutUrl** | Pointer to **string** | CheckoutUrl built from checkoutTemplateId and planId | [optional] 
**CompareAtAmount** | Pointer to **int64** | optional compare-at / strikethrough display price for regular plan amount in minor units (plan currency); UI only | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**Description** | Pointer to **string** | description | [optional] 
**HomeUrl** | Pointer to **string** | home_url | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ImageUrl** | Pointer to **string** | image_url | [optional] 
**IntervalCount** | Pointer to **int32** | period unit count | [optional] 
**IntervalUnit** | Pointer to **string** | period unit,day|month|year|week | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MultiCurrencies** | Pointer to [**[]UnibeeApiBeanPublicPlanMultiCurrency**](UnibeeApiBeanPublicPlanMultiCurrency.md) | Plan&#39;s public MultiCurrencies | [optional] 
**MultiTrials** | Pointer to [**[]UnibeeApiBeanPublicPlanMultiTrial**](UnibeeApiBeanPublicPlanMultiTrial.md) | multi-step trials for this plan; metadata intentionally omitted | [optional] 
**PlanName** | Pointer to **string** | PlanName | [optional] 
**ProductId** | Pointer to **int64** | product id | [optional] 
**Status** | Pointer to **int32** | status，1-editing，2-active，3-inactive，4-soft archive, 5-hard archive | [optional] 
**TaxPercentage** | Pointer to **int32** | TaxPercentage 1000 &#x3D; 10% | [optional] 
**TrialAmount** | Pointer to **int64** | price of trial period; may be overridden by multiTrials[0].trialPrice | [optional] 
**TrialDemand** | Pointer to **string** |  | [optional] 
**TrialDurationTime** | Pointer to **int64** | duration of trial (seconds); may be overridden by multiTrials[0].trialDuration | [optional] 
**Type** | Pointer to **int32** | type，1-main plan，2-addon plan | [optional] 
**UsVATConfig** | Pointer to [**UnibeeApiBeanPublicUSVATConfig**](UnibeeApiBeanPublicUSVATConfig.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanPublicPlan

`func NewUnibeeApiBeanPublicPlan() *UnibeeApiBeanPublicPlan`

NewUnibeeApiBeanPublicPlan instantiates a new UnibeeApiBeanPublicPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPublicPlanWithDefaults

`func NewUnibeeApiBeanPublicPlanWithDefaults() *UnibeeApiBeanPublicPlan`

NewUnibeeApiBeanPublicPlanWithDefaults instantiates a new UnibeeApiBeanPublicPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UnibeeApiBeanPublicPlan) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeApiBeanPublicPlan) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeApiBeanPublicPlan) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeApiBeanPublicPlan) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAutoUpgradePlanId

`func (o *UnibeeApiBeanPublicPlan) GetAutoUpgradePlanId() int64`

GetAutoUpgradePlanId returns the AutoUpgradePlanId field if non-nil, zero value otherwise.

### GetAutoUpgradePlanIdOk

`func (o *UnibeeApiBeanPublicPlan) GetAutoUpgradePlanIdOk() (*int64, bool)`

GetAutoUpgradePlanIdOk returns a tuple with the AutoUpgradePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgradePlanId

`func (o *UnibeeApiBeanPublicPlan) SetAutoUpgradePlanId(v int64)`

SetAutoUpgradePlanId sets AutoUpgradePlanId field to given value.

### HasAutoUpgradePlanId

`func (o *UnibeeApiBeanPublicPlan) HasAutoUpgradePlanId() bool`

HasAutoUpgradePlanId returns a boolean if a field has been set.

### GetBindingAddonIds

`func (o *UnibeeApiBeanPublicPlan) GetBindingAddonIds() string`

GetBindingAddonIds returns the BindingAddonIds field if non-nil, zero value otherwise.

### GetBindingAddonIdsOk

`func (o *UnibeeApiBeanPublicPlan) GetBindingAddonIdsOk() (*string, bool)`

GetBindingAddonIdsOk returns a tuple with the BindingAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingAddonIds

`func (o *UnibeeApiBeanPublicPlan) SetBindingAddonIds(v string)`

SetBindingAddonIds sets BindingAddonIds field to given value.

### HasBindingAddonIds

`func (o *UnibeeApiBeanPublicPlan) HasBindingAddonIds() bool`

HasBindingAddonIds returns a boolean if a field has been set.

### GetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPublicPlan) GetBindingOnetimeAddonIds() string`

GetBindingOnetimeAddonIds returns the BindingOnetimeAddonIds field if non-nil, zero value otherwise.

### GetBindingOnetimeAddonIdsOk

`func (o *UnibeeApiBeanPublicPlan) GetBindingOnetimeAddonIdsOk() (*string, bool)`

GetBindingOnetimeAddonIdsOk returns a tuple with the BindingOnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPublicPlan) SetBindingOnetimeAddonIds(v string)`

SetBindingOnetimeAddonIds sets BindingOnetimeAddonIds field to given value.

### HasBindingOnetimeAddonIds

`func (o *UnibeeApiBeanPublicPlan) HasBindingOnetimeAddonIds() bool`

HasBindingOnetimeAddonIds returns a boolean if a field has been set.

### GetCheckoutTemplateId

`func (o *UnibeeApiBeanPublicPlan) GetCheckoutTemplateId() int64`

GetCheckoutTemplateId returns the CheckoutTemplateId field if non-nil, zero value otherwise.

### GetCheckoutTemplateIdOk

`func (o *UnibeeApiBeanPublicPlan) GetCheckoutTemplateIdOk() (*int64, bool)`

GetCheckoutTemplateIdOk returns a tuple with the CheckoutTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutTemplateId

`func (o *UnibeeApiBeanPublicPlan) SetCheckoutTemplateId(v int64)`

SetCheckoutTemplateId sets CheckoutTemplateId field to given value.

### HasCheckoutTemplateId

`func (o *UnibeeApiBeanPublicPlan) HasCheckoutTemplateId() bool`

HasCheckoutTemplateId returns a boolean if a field has been set.

### GetCheckoutUrl

`func (o *UnibeeApiBeanPublicPlan) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *UnibeeApiBeanPublicPlan) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *UnibeeApiBeanPublicPlan) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *UnibeeApiBeanPublicPlan) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### GetCompareAtAmount

`func (o *UnibeeApiBeanPublicPlan) GetCompareAtAmount() int64`

GetCompareAtAmount returns the CompareAtAmount field if non-nil, zero value otherwise.

### GetCompareAtAmountOk

`func (o *UnibeeApiBeanPublicPlan) GetCompareAtAmountOk() (*int64, bool)`

GetCompareAtAmountOk returns a tuple with the CompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmount

`func (o *UnibeeApiBeanPublicPlan) SetCompareAtAmount(v int64)`

SetCompareAtAmount sets CompareAtAmount field to given value.

### HasCompareAtAmount

`func (o *UnibeeApiBeanPublicPlan) HasCompareAtAmount() bool`

HasCompareAtAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiBeanPublicPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPublicPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPublicPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPublicPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *UnibeeApiBeanPublicPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnibeeApiBeanPublicPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnibeeApiBeanPublicPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnibeeApiBeanPublicPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomeUrl

`func (o *UnibeeApiBeanPublicPlan) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *UnibeeApiBeanPublicPlan) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *UnibeeApiBeanPublicPlan) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *UnibeeApiBeanPublicPlan) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanPublicPlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanPublicPlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanPublicPlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanPublicPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *UnibeeApiBeanPublicPlan) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *UnibeeApiBeanPublicPlan) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *UnibeeApiBeanPublicPlan) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *UnibeeApiBeanPublicPlan) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIntervalCount

`func (o *UnibeeApiBeanPublicPlan) GetIntervalCount() int32`

GetIntervalCount returns the IntervalCount field if non-nil, zero value otherwise.

### GetIntervalCountOk

`func (o *UnibeeApiBeanPublicPlan) GetIntervalCountOk() (*int32, bool)`

GetIntervalCountOk returns a tuple with the IntervalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalCount

`func (o *UnibeeApiBeanPublicPlan) SetIntervalCount(v int32)`

SetIntervalCount sets IntervalCount field to given value.

### HasIntervalCount

`func (o *UnibeeApiBeanPublicPlan) HasIntervalCount() bool`

HasIntervalCount returns a boolean if a field has been set.

### GetIntervalUnit

`func (o *UnibeeApiBeanPublicPlan) GetIntervalUnit() string`

GetIntervalUnit returns the IntervalUnit field if non-nil, zero value otherwise.

### GetIntervalUnitOk

`func (o *UnibeeApiBeanPublicPlan) GetIntervalUnitOk() (*string, bool)`

GetIntervalUnitOk returns a tuple with the IntervalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalUnit

`func (o *UnibeeApiBeanPublicPlan) SetIntervalUnit(v string)`

SetIntervalUnit sets IntervalUnit field to given value.

### HasIntervalUnit

`func (o *UnibeeApiBeanPublicPlan) HasIntervalUnit() bool`

HasIntervalUnit returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanPublicPlan) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanPublicPlan) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanPublicPlan) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanPublicPlan) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMultiCurrencies

`func (o *UnibeeApiBeanPublicPlan) GetMultiCurrencies() []UnibeeApiBeanPublicPlanMultiCurrency`

GetMultiCurrencies returns the MultiCurrencies field if non-nil, zero value otherwise.

### GetMultiCurrenciesOk

`func (o *UnibeeApiBeanPublicPlan) GetMultiCurrenciesOk() (*[]UnibeeApiBeanPublicPlanMultiCurrency, bool)`

GetMultiCurrenciesOk returns a tuple with the MultiCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiCurrencies

`func (o *UnibeeApiBeanPublicPlan) SetMultiCurrencies(v []UnibeeApiBeanPublicPlanMultiCurrency)`

SetMultiCurrencies sets MultiCurrencies field to given value.

### HasMultiCurrencies

`func (o *UnibeeApiBeanPublicPlan) HasMultiCurrencies() bool`

HasMultiCurrencies returns a boolean if a field has been set.

### GetMultiTrials

`func (o *UnibeeApiBeanPublicPlan) GetMultiTrials() []UnibeeApiBeanPublicPlanMultiTrial`

GetMultiTrials returns the MultiTrials field if non-nil, zero value otherwise.

### GetMultiTrialsOk

`func (o *UnibeeApiBeanPublicPlan) GetMultiTrialsOk() (*[]UnibeeApiBeanPublicPlanMultiTrial, bool)`

GetMultiTrialsOk returns a tuple with the MultiTrials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiTrials

`func (o *UnibeeApiBeanPublicPlan) SetMultiTrials(v []UnibeeApiBeanPublicPlanMultiTrial)`

SetMultiTrials sets MultiTrials field to given value.

### HasMultiTrials

`func (o *UnibeeApiBeanPublicPlan) HasMultiTrials() bool`

HasMultiTrials returns a boolean if a field has been set.

### GetPlanName

`func (o *UnibeeApiBeanPublicPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UnibeeApiBeanPublicPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UnibeeApiBeanPublicPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UnibeeApiBeanPublicPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiBeanPublicPlan) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiBeanPublicPlan) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiBeanPublicPlan) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiBeanPublicPlan) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiBeanPublicPlan) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiBeanPublicPlan) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiBeanPublicPlan) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiBeanPublicPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiBeanPublicPlan) GetTaxPercentage() int32`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiBeanPublicPlan) GetTaxPercentageOk() (*int32, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiBeanPublicPlan) SetTaxPercentage(v int32)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiBeanPublicPlan) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTrialAmount

`func (o *UnibeeApiBeanPublicPlan) GetTrialAmount() int64`

GetTrialAmount returns the TrialAmount field if non-nil, zero value otherwise.

### GetTrialAmountOk

`func (o *UnibeeApiBeanPublicPlan) GetTrialAmountOk() (*int64, bool)`

GetTrialAmountOk returns a tuple with the TrialAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmount

`func (o *UnibeeApiBeanPublicPlan) SetTrialAmount(v int64)`

SetTrialAmount sets TrialAmount field to given value.

### HasTrialAmount

`func (o *UnibeeApiBeanPublicPlan) HasTrialAmount() bool`

HasTrialAmount returns a boolean if a field has been set.

### GetTrialDemand

`func (o *UnibeeApiBeanPublicPlan) GetTrialDemand() string`

GetTrialDemand returns the TrialDemand field if non-nil, zero value otherwise.

### GetTrialDemandOk

`func (o *UnibeeApiBeanPublicPlan) GetTrialDemandOk() (*string, bool)`

GetTrialDemandOk returns a tuple with the TrialDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDemand

`func (o *UnibeeApiBeanPublicPlan) SetTrialDemand(v string)`

SetTrialDemand sets TrialDemand field to given value.

### HasTrialDemand

`func (o *UnibeeApiBeanPublicPlan) HasTrialDemand() bool`

HasTrialDemand returns a boolean if a field has been set.

### GetTrialDurationTime

`func (o *UnibeeApiBeanPublicPlan) GetTrialDurationTime() int64`

GetTrialDurationTime returns the TrialDurationTime field if non-nil, zero value otherwise.

### GetTrialDurationTimeOk

`func (o *UnibeeApiBeanPublicPlan) GetTrialDurationTimeOk() (*int64, bool)`

GetTrialDurationTimeOk returns a tuple with the TrialDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDurationTime

`func (o *UnibeeApiBeanPublicPlan) SetTrialDurationTime(v int64)`

SetTrialDurationTime sets TrialDurationTime field to given value.

### HasTrialDurationTime

`func (o *UnibeeApiBeanPublicPlan) HasTrialDurationTime() bool`

HasTrialDurationTime returns a boolean if a field has been set.

### GetType

`func (o *UnibeeApiBeanPublicPlan) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeApiBeanPublicPlan) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeApiBeanPublicPlan) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeApiBeanPublicPlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsVATConfig

`func (o *UnibeeApiBeanPublicPlan) GetUsVATConfig() UnibeeApiBeanPublicUSVATConfig`

GetUsVATConfig returns the UsVATConfig field if non-nil, zero value otherwise.

### GetUsVATConfigOk

`func (o *UnibeeApiBeanPublicPlan) GetUsVATConfigOk() (*UnibeeApiBeanPublicUSVATConfig, bool)`

GetUsVATConfigOk returns a tuple with the UsVATConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsVATConfig

`func (o *UnibeeApiBeanPublicPlan) SetUsVATConfig(v UnibeeApiBeanPublicUSVATConfig)`

SetUsVATConfig sets UsVATConfig field to given value.

### HasUsVATConfig

`func (o *UnibeeApiBeanPublicPlan) HasUsVATConfig() bool`

HasUsVATConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


