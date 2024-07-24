# UnibeeApiBeanDetailMerchantOperationLogDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** | operation create utc time | [optional] 
**DiscountCode** | Pointer to **string** | discount_code | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**Member** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**MemberId** | Pointer to **int64** | member_id | [optional] 
**MerchantId** | Pointer to **int64** | merchant Id | [optional] 
**OptAccount** | Pointer to **string** | operation Account | [optional] 
**OptAccountId** | Pointer to **string** | operation account id | [optional] 
**OptAccountType** | Pointer to **int32** | opt_account_type, 0-Member|1-User|2-OpenApi|3-System | [optional] 
**OptContent** | Pointer to **string** | operation content | [optional] 
**OptTarget** | Pointer to **string** | operation target | [optional] 
**PlanId** | Pointer to **int64** | plan id | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeApiBeanDetailMerchantOperationLogDetail

`func NewUnibeeApiBeanDetailMerchantOperationLogDetail() *UnibeeApiBeanDetailMerchantOperationLogDetail`

NewUnibeeApiBeanDetailMerchantOperationLogDetail instantiates a new UnibeeApiBeanDetailMerchantOperationLogDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailMerchantOperationLogDetailWithDefaults

`func NewUnibeeApiBeanDetailMerchantOperationLogDetailWithDefaults() *UnibeeApiBeanDetailMerchantOperationLogDetail`

NewUnibeeApiBeanDetailMerchantOperationLogDetailWithDefaults instantiates a new UnibeeApiBeanDetailMerchantOperationLogDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetMember

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMember sets Member field to given value.

### HasMember

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMemberId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetOptAccount

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccount() string`

GetOptAccount returns the OptAccount field if non-nil, zero value otherwise.

### GetOptAccountOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountOk() (*string, bool)`

GetOptAccountOk returns a tuple with the OptAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptAccount

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccount(v string)`

SetOptAccount sets OptAccount field to given value.

### HasOptAccount

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccount() bool`

HasOptAccount returns a boolean if a field has been set.

### GetOptAccountId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountId() string`

GetOptAccountId returns the OptAccountId field if non-nil, zero value otherwise.

### GetOptAccountIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountIdOk() (*string, bool)`

GetOptAccountIdOk returns a tuple with the OptAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptAccountId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccountId(v string)`

SetOptAccountId sets OptAccountId field to given value.

### HasOptAccountId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccountId() bool`

HasOptAccountId returns a boolean if a field has been set.

### GetOptAccountType

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountType() int32`

GetOptAccountType returns the OptAccountType field if non-nil, zero value otherwise.

### GetOptAccountTypeOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptAccountTypeOk() (*int32, bool)`

GetOptAccountTypeOk returns a tuple with the OptAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptAccountType

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptAccountType(v int32)`

SetOptAccountType sets OptAccountType field to given value.

### HasOptAccountType

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptAccountType() bool`

HasOptAccountType returns a boolean if a field has been set.

### GetOptContent

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptContent() string`

GetOptContent returns the OptContent field if non-nil, zero value otherwise.

### GetOptContentOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptContentOk() (*string, bool)`

GetOptContentOk returns a tuple with the OptContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptContent

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptContent(v string)`

SetOptContent sets OptContent field to given value.

### HasOptContent

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptContent() bool`

HasOptContent returns a boolean if a field has been set.

### GetOptTarget

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptTarget() string`

GetOptTarget returns the OptTarget field if non-nil, zero value otherwise.

### GetOptTargetOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetOptTargetOk() (*string, bool)`

GetOptTargetOk returns a tuple with the OptTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptTarget

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetOptTarget(v string)`

SetOptTarget sets OptTarget field to given value.

### HasOptTarget

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasOptTarget() bool`

HasOptTarget returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiBeanDetailMerchantOperationLogDetail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


