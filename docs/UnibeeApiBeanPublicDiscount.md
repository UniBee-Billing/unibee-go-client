# UnibeeApiBeanPublicDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Discount currency, available when discountType is fixed_amount. | [optional] 
**CycleLimit** | Pointer to **int32** | Subscription cycle count limit for this discount, 0 &#x3D; no limit. | [optional] 
**DiscountAmount** | Pointer to **int64** | Configured discount amount, available when discountType is fixed_amount, minor units. | [optional] 
**DiscountPercentage** | Pointer to **int64** | Configured discount percentage, 100 &#x3D; 1%, available when discountType is percentage. | [optional] 
**DiscountType** | Pointer to **int32** | 1-percentage, 2-fixed_amount. | [optional] 

## Methods

### NewUnibeeApiBeanPublicDiscount

`func NewUnibeeApiBeanPublicDiscount() *UnibeeApiBeanPublicDiscount`

NewUnibeeApiBeanPublicDiscount instantiates a new UnibeeApiBeanPublicDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanPublicDiscountWithDefaults

`func NewUnibeeApiBeanPublicDiscountWithDefaults() *UnibeeApiBeanPublicDiscount`

NewUnibeeApiBeanPublicDiscountWithDefaults instantiates a new UnibeeApiBeanPublicDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *UnibeeApiBeanPublicDiscount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiBeanPublicDiscount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiBeanPublicDiscount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiBeanPublicDiscount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiBeanPublicDiscount) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiBeanPublicDiscount) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiBeanPublicDiscount) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiBeanPublicDiscount) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiBeanPublicDiscount) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiBeanPublicDiscount) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiBeanPublicDiscount) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiBeanPublicDiscount) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiBeanPublicDiscount) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiBeanPublicDiscount) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiBeanPublicDiscount) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


