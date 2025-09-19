# UnibeeApiCheckoutPlanDetailRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | Pointer to [**UnibeeApiBeanMerchant**](UnibeeApiBeanMerchant.md) |  | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanDetailPlanDetail**](UnibeeApiBeanDetailPlanDetail.md) |  | [optional] 

## Methods

### NewUnibeeApiCheckoutPlanDetailRes

`func NewUnibeeApiCheckoutPlanDetailRes() *UnibeeApiCheckoutPlanDetailRes`

NewUnibeeApiCheckoutPlanDetailRes instantiates a new UnibeeApiCheckoutPlanDetailRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiCheckoutPlanDetailResWithDefaults

`func NewUnibeeApiCheckoutPlanDetailResWithDefaults() *UnibeeApiCheckoutPlanDetailRes`

NewUnibeeApiCheckoutPlanDetailResWithDefaults instantiates a new UnibeeApiCheckoutPlanDetailRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *UnibeeApiCheckoutPlanDetailRes) GetMerchant() UnibeeApiBeanMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *UnibeeApiCheckoutPlanDetailRes) GetMerchantOk() (*UnibeeApiBeanMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *UnibeeApiCheckoutPlanDetailRes) SetMerchant(v UnibeeApiBeanMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *UnibeeApiCheckoutPlanDetailRes) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiCheckoutPlanDetailRes) GetPlan() UnibeeApiBeanDetailPlanDetail`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiCheckoutPlanDetailRes) GetPlanOk() (*UnibeeApiBeanDetailPlanDetail, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiCheckoutPlanDetailRes) SetPlan(v UnibeeApiBeanDetailPlanDetail)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiCheckoutPlanDetailRes) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


