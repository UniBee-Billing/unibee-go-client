# UnibeeApiMerchantDiscountEditReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingType** | Pointer to **int32** | The billing type of the discount code, 1-one-time, 2-recurring, define the situation the code can be used, the code of one-time billing_type can used for all situation that effect only once, the code of recurring billing_tye can only used for subscription purchase | [optional] 
**Currency** | Pointer to **string** | The discount currency of discount code, available when discount_type is fixed_amount | [optional] 
**CycleLimit** | Pointer to **int32** | The count limitation of subscription cycle，each subscription is valid separately, 0-no limit | [optional] 
**DiscountAmount** | Pointer to **int64** | The discount amount of the discount code, available when discount_type is fixed_amount | [optional] 
**DiscountPercentage** | Pointer to **int64** | The discount percentage of discount code, 100&#x3D;1%, available when discount_type is percentage | [optional] 
**DiscountType** | Pointer to **int32** | The discount type of the discount code, 1-percentage, 2-fixed_amount, the discountType of code, the discountPercentage will be effect when discountType is percentage, the discountAmount and currency will be effect when discountTYpe is fixed_amount | [optional] 
**EndTime** | Pointer to **int64** | The end time of discount code can effect, utc time | [optional] 
**Id** | **int64** | The discount&#39;s Id | 
**Metadata** | Pointer to **map[string]string** | Metadata，Map | [optional] 
**Name** | Pointer to **string** | The discount&#39;s name | [optional] 
**StartTime** | Pointer to **int64** | The start time of discount code can effect, utc time | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountEditReq

`func NewUnibeeApiMerchantDiscountEditReq(id int64, ) *UnibeeApiMerchantDiscountEditReq`

NewUnibeeApiMerchantDiscountEditReq instantiates a new UnibeeApiMerchantDiscountEditReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountEditReqWithDefaults

`func NewUnibeeApiMerchantDiscountEditReqWithDefaults() *UnibeeApiMerchantDiscountEditReq`

NewUnibeeApiMerchantDiscountEditReqWithDefaults instantiates a new UnibeeApiMerchantDiscountEditReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingType

`func (o *UnibeeApiMerchantDiscountEditReq) GetBillingType() int32`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetBillingTypeOk() (*int32, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *UnibeeApiMerchantDiscountEditReq) SetBillingType(v int32)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *UnibeeApiMerchantDiscountEditReq) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeApiMerchantDiscountEditReq) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeApiMerchantDiscountEditReq) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeApiMerchantDiscountEditReq) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCycleLimit

`func (o *UnibeeApiMerchantDiscountEditReq) GetCycleLimit() int32`

GetCycleLimit returns the CycleLimit field if non-nil, zero value otherwise.

### GetCycleLimitOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetCycleLimitOk() (*int32, bool)`

GetCycleLimitOk returns a tuple with the CycleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLimit

`func (o *UnibeeApiMerchantDiscountEditReq) SetCycleLimit(v int32)`

SetCycleLimit sets CycleLimit field to given value.

### HasCycleLimit

`func (o *UnibeeApiMerchantDiscountEditReq) HasCycleLimit() bool`

HasCycleLimit returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountPercentage() int64`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountPercentageOk() (*int64, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountPercentage(v int64)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountType() int32`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetDiscountTypeOk() (*int32, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UnibeeApiMerchantDiscountEditReq) SetDiscountType(v int32)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UnibeeApiMerchantDiscountEditReq) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetEndTime

`func (o *UnibeeApiMerchantDiscountEditReq) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UnibeeApiMerchantDiscountEditReq) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UnibeeApiMerchantDiscountEditReq) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiMerchantDiscountEditReq) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiMerchantDiscountEditReq) SetId(v int64)`

SetId sets Id field to given value.


### GetMetadata

`func (o *UnibeeApiMerchantDiscountEditReq) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantDiscountEditReq) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantDiscountEditReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UnibeeApiMerchantDiscountEditReq) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UnibeeApiMerchantDiscountEditReq) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UnibeeApiMerchantDiscountEditReq) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *UnibeeApiMerchantDiscountEditReq) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UnibeeApiMerchantDiscountEditReq) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UnibeeApiMerchantDiscountEditReq) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UnibeeApiMerchantDiscountEditReq) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


