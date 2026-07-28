# UnibeeApiMerchantPlanBatchTagBindReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanIds** | **[]int64** | Plan ids to bind tags to | 
**TagIds** | **[]int64** | Plan tag ids to append. Existing bindings are skipped. | 

## Methods

### NewUnibeeApiMerchantPlanBatchTagBindReq

`func NewUnibeeApiMerchantPlanBatchTagBindReq(planIds []int64, tagIds []int64, ) *UnibeeApiMerchantPlanBatchTagBindReq`

NewUnibeeApiMerchantPlanBatchTagBindReq instantiates a new UnibeeApiMerchantPlanBatchTagBindReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPlanBatchTagBindReqWithDefaults

`func NewUnibeeApiMerchantPlanBatchTagBindReqWithDefaults() *UnibeeApiMerchantPlanBatchTagBindReq`

NewUnibeeApiMerchantPlanBatchTagBindReqWithDefaults instantiates a new UnibeeApiMerchantPlanBatchTagBindReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanIds

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) GetPlanIds() []int64`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) GetPlanIdsOk() (*[]int64, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) SetPlanIds(v []int64)`

SetPlanIds sets PlanIds field to given value.


### GetTagIds

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) GetTagIds() []int64`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) GetTagIdsOk() (*[]int64, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *UnibeeApiMerchantPlanBatchTagBindReq) SetTagIds(v []int64)`

SetTagIds sets TagIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


