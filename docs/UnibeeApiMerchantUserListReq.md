# UnibeeApiMerchantUserListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count OF Page | [optional] 
**CreateTimeEnd** | Pointer to **int64** | CreateTimeEnd，UTC timestamp，seconds | [optional] 
**CreateTimeStart** | Pointer to **int64** | CreateTimeStart，UTC timestamp，seconds | [optional] 
**DeleteInclude** | Pointer to **bool** | Deleted Involved，Need Admin | [optional] 
**Email** | Pointer to **string** | Search Filter Email | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId | [optional] 
**FirstName** | Pointer to **string** | Search FirstName | [optional] 
**GatewayIds** | Pointer to **[]int64** | GatewayIds, Search Filter GatewayIds | [optional] 
**LastName** | Pointer to **string** | Search LastName | [optional] 
**Page** | Pointer to **int32** | Page,Start 0 | [optional] 
**PlanIds** | Pointer to **[]int32** | PlanIds, Search Filter PlanIds | [optional] 
**SortField** | Pointer to **string** | Sort，user_id|gmt_create|email|user_name|subscription_name|subscription_status|payment_method|recurring_amount|billing_type，Default gmt_create | [optional] 
**SortType** | Pointer to **string** | Sort Type，asc|desc，Default desc | [optional] 
**Status** | Pointer to **[]int32** | Status, 0-Active｜2-Frozen | [optional] 
**SubStatus** | Pointer to **[]int32** | Filter, Default All，1-Pending｜2-Active｜3-Suspend | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | 8-Processing | 9-Failed | [optional] 
**SubscriptionId** | Pointer to **string** | Search Filter SubscriptionId | [optional] 
**UserId** | Pointer to **int64** | Filter UserId | [optional] 

## Methods

### NewUnibeeApiMerchantUserListReq

`func NewUnibeeApiMerchantUserListReq() *UnibeeApiMerchantUserListReq`

NewUnibeeApiMerchantUserListReq instantiates a new UnibeeApiMerchantUserListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserListReqWithDefaults

`func NewUnibeeApiMerchantUserListReqWithDefaults() *UnibeeApiMerchantUserListReq`

NewUnibeeApiMerchantUserListReqWithDefaults instantiates a new UnibeeApiMerchantUserListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantUserListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantUserListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantUserListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantUserListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateTimeEnd

`func (o *UnibeeApiMerchantUserListReq) GetCreateTimeEnd() int64`

GetCreateTimeEnd returns the CreateTimeEnd field if non-nil, zero value otherwise.

### GetCreateTimeEndOk

`func (o *UnibeeApiMerchantUserListReq) GetCreateTimeEndOk() (*int64, bool)`

GetCreateTimeEndOk returns a tuple with the CreateTimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeEnd

`func (o *UnibeeApiMerchantUserListReq) SetCreateTimeEnd(v int64)`

SetCreateTimeEnd sets CreateTimeEnd field to given value.

### HasCreateTimeEnd

`func (o *UnibeeApiMerchantUserListReq) HasCreateTimeEnd() bool`

HasCreateTimeEnd returns a boolean if a field has been set.

### GetCreateTimeStart

`func (o *UnibeeApiMerchantUserListReq) GetCreateTimeStart() int64`

GetCreateTimeStart returns the CreateTimeStart field if non-nil, zero value otherwise.

### GetCreateTimeStartOk

`func (o *UnibeeApiMerchantUserListReq) GetCreateTimeStartOk() (*int64, bool)`

GetCreateTimeStartOk returns a tuple with the CreateTimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeStart

`func (o *UnibeeApiMerchantUserListReq) SetCreateTimeStart(v int64)`

SetCreateTimeStart sets CreateTimeStart field to given value.

### HasCreateTimeStart

`func (o *UnibeeApiMerchantUserListReq) HasCreateTimeStart() bool`

HasCreateTimeStart returns a boolean if a field has been set.

### GetDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) GetDeleteInclude() bool`

GetDeleteInclude returns the DeleteInclude field if non-nil, zero value otherwise.

### GetDeleteIncludeOk

`func (o *UnibeeApiMerchantUserListReq) GetDeleteIncludeOk() (*bool, bool)`

GetDeleteIncludeOk returns a tuple with the DeleteInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) SetDeleteInclude(v bool)`

SetDeleteInclude sets DeleteInclude field to given value.

### HasDeleteInclude

`func (o *UnibeeApiMerchantUserListReq) HasDeleteInclude() bool`

HasDeleteInclude returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantUserListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantUserListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantUserListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantUserListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantUserListReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantUserListReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantUserListReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantUserListReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiMerchantUserListReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantUserListReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantUserListReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantUserListReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGatewayIds

`func (o *UnibeeApiMerchantUserListReq) GetGatewayIds() []int64`

GetGatewayIds returns the GatewayIds field if non-nil, zero value otherwise.

### GetGatewayIdsOk

`func (o *UnibeeApiMerchantUserListReq) GetGatewayIdsOk() (*[]int64, bool)`

GetGatewayIdsOk returns a tuple with the GatewayIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIds

`func (o *UnibeeApiMerchantUserListReq) SetGatewayIds(v []int64)`

SetGatewayIds sets GatewayIds field to given value.

### HasGatewayIds

`func (o *UnibeeApiMerchantUserListReq) HasGatewayIds() bool`

HasGatewayIds returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantUserListReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantUserListReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantUserListReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantUserListReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantUserListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantUserListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantUserListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantUserListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlanIds

`func (o *UnibeeApiMerchantUserListReq) GetPlanIds() []int32`

GetPlanIds returns the PlanIds field if non-nil, zero value otherwise.

### GetPlanIdsOk

`func (o *UnibeeApiMerchantUserListReq) GetPlanIdsOk() (*[]int32, bool)`

GetPlanIdsOk returns a tuple with the PlanIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanIds

`func (o *UnibeeApiMerchantUserListReq) SetPlanIds(v []int32)`

SetPlanIds sets PlanIds field to given value.

### HasPlanIds

`func (o *UnibeeApiMerchantUserListReq) HasPlanIds() bool`

HasPlanIds returns a boolean if a field has been set.

### GetSortField

`func (o *UnibeeApiMerchantUserListReq) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *UnibeeApiMerchantUserListReq) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *UnibeeApiMerchantUserListReq) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *UnibeeApiMerchantUserListReq) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortType

`func (o *UnibeeApiMerchantUserListReq) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *UnibeeApiMerchantUserListReq) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *UnibeeApiMerchantUserListReq) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *UnibeeApiMerchantUserListReq) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeApiMerchantUserListReq) GetStatus() []int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeApiMerchantUserListReq) GetStatusOk() (*[]int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeApiMerchantUserListReq) SetStatus(v []int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeApiMerchantUserListReq) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubStatus

`func (o *UnibeeApiMerchantUserListReq) GetSubStatus() []int32`

GetSubStatus returns the SubStatus field if non-nil, zero value otherwise.

### GetSubStatusOk

`func (o *UnibeeApiMerchantUserListReq) GetSubStatusOk() (*[]int32, bool)`

GetSubStatusOk returns a tuple with the SubStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStatus

`func (o *UnibeeApiMerchantUserListReq) SetSubStatus(v []int32)`

SetSubStatus sets SubStatus field to given value.

### HasSubStatus

`func (o *UnibeeApiMerchantUserListReq) HasSubStatus() bool`

HasSubStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantUserListReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantUserListReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantUserListReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantUserListReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantUserListReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantUserListReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantUserListReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantUserListReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


