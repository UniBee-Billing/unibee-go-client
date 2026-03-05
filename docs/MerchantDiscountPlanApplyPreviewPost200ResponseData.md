# MerchantDiscountPlanApplyPreviewPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllPlansAllowed** | Pointer to **bool** | True &#x3D; discount applies to all plans in cart; false &#x3D; only to allowedPlanIds | [optional] 
**AllowedPlanIds** | Pointer to **[]int64** | Plan ids that get this discount (when partial apply); empty or same as main+addons when all plans allowed | [optional] 
**DiscountAmount** | Pointer to **int64** | The discount amount can apply to plan | [optional] 
**DiscountCode** | Pointer to [**UnibeeApiBeanMerchantDiscountCode**](UnibeeApiBeanMerchantDiscountCode.md) |  | [optional] 
**FailureReason** | Pointer to **string** | The apply preview failure reason | [optional] 
**Valid** | Pointer to **bool** | The apply preview result, true or false | [optional] 

## Methods

### NewMerchantDiscountPlanApplyPreviewPost200ResponseData

`func NewMerchantDiscountPlanApplyPreviewPost200ResponseData() *MerchantDiscountPlanApplyPreviewPost200ResponseData`

NewMerchantDiscountPlanApplyPreviewPost200ResponseData instantiates a new MerchantDiscountPlanApplyPreviewPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantDiscountPlanApplyPreviewPost200ResponseDataWithDefaults

`func NewMerchantDiscountPlanApplyPreviewPost200ResponseDataWithDefaults() *MerchantDiscountPlanApplyPreviewPost200ResponseData`

NewMerchantDiscountPlanApplyPreviewPost200ResponseDataWithDefaults instantiates a new MerchantDiscountPlanApplyPreviewPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllPlansAllowed

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetAllPlansAllowed() bool`

GetAllPlansAllowed returns the AllPlansAllowed field if non-nil, zero value otherwise.

### GetAllPlansAllowedOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetAllPlansAllowedOk() (*bool, bool)`

GetAllPlansAllowedOk returns a tuple with the AllPlansAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPlansAllowed

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetAllPlansAllowed(v bool)`

SetAllPlansAllowed sets AllPlansAllowed field to given value.

### HasAllPlansAllowed

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasAllPlansAllowed() bool`

HasAllPlansAllowed returns a boolean if a field has been set.

### GetAllowedPlanIds

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetAllowedPlanIds() []int64`

GetAllowedPlanIds returns the AllowedPlanIds field if non-nil, zero value otherwise.

### GetAllowedPlanIdsOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetAllowedPlanIdsOk() (*[]int64, bool)`

GetAllowedPlanIdsOk returns a tuple with the AllowedPlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlanIds

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetAllowedPlanIds(v []int64)`

SetAllowedPlanIds sets AllowedPlanIds field to given value.

### HasAllowedPlanIds

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasAllowedPlanIds() bool`

HasAllowedPlanIds returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountAmount() int64`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountAmountOk() (*int64, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetDiscountAmount(v int64)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountCode() UnibeeApiBeanMerchantDiscountCode`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetDiscountCodeOk() (*UnibeeApiBeanMerchantDiscountCode, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetDiscountCode(v UnibeeApiBeanMerchantDiscountCode)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetFailureReason

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetValid

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *MerchantDiscountPlanApplyPreviewPost200ResponseData) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


