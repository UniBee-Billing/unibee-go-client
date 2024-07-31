# UnibeeApiMerchantDiscountPlanApplyPreviewRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountAmount** | Pointer to **int64** | The discount amount can apply to plan | [optional] 
**DiscountCode** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**FailureReason** | Pointer to **string** | The apply preview failure reason | [optional] 
**Valid** | Pointer to **bool** | The apply preview result, true or false | [optional] 

## Methods

### NewUnibeeApiMerchantDiscountPlanApplyPreviewRes

`func NewUnibeeApiMerchantDiscountPlanApplyPreviewRes() *UnibeeApiMerchantDiscountPlanApplyPreviewRes`

NewUnibeeApiMerchantDiscountPlanApplyPreviewRes instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantDiscountPlanApplyPreviewResWithDefaults

`func NewUnibeeApiMerchantDiscountPlanApplyPreviewResWithDefaults() *UnibeeApiMerchantDiscountPlanApplyPreviewRes`

NewUnibeeApiMerchantDiscountPlanApplyPreviewResWithDefaults instantiates a new UnibeeApiMerchantDiscountPlanApplyPreviewRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountAmount

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetValid

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *UnibeeApiMerchantDiscountPlanApplyPreviewRes) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


