# UnibeeApiMerchantMemberOperationLogListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count Of Page | [optional] 
**DiscountCode** | Pointer to **string** | discount_code | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**FirstName** | Pointer to **string** | FirstName | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**LastName** | Pointer to **string** | LastName | [optional] 
**MemberEmail** | Pointer to **string** | Filter Member&#39;s Email, Default All | [optional] 
**MemberFirstName** | Pointer to **string** | Filter Member&#39;s FirstName Default All | [optional] 
**MemberLastName** | Pointer to **string** | Filter Member&#39;s LastName, Default All | [optional] 
**Page** | Pointer to **int32** | Page, Start With 0 | [optional] 
**PlanId** | Pointer to **int64** | plan id | [optional] 
**SubscriptionId** | Pointer to **string** | subscription_id | [optional] 

## Methods

### NewUnibeeApiMerchantMemberOperationLogListReq

`func NewUnibeeApiMerchantMemberOperationLogListReq() *UnibeeApiMerchantMemberOperationLogListReq`

NewUnibeeApiMerchantMemberOperationLogListReq instantiates a new UnibeeApiMerchantMemberOperationLogListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberOperationLogListReqWithDefaults

`func NewUnibeeApiMerchantMemberOperationLogListReqWithDefaults() *UnibeeApiMerchantMemberOperationLogListReq`

NewUnibeeApiMerchantMemberOperationLogListReqWithDefaults instantiates a new UnibeeApiMerchantMemberOperationLogListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDiscountCode

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetDiscountCode() string`

GetDiscountCode returns the DiscountCode field if non-nil, zero value otherwise.

### GetDiscountCodeOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetDiscountCodeOk() (*string, bool)`

GetDiscountCodeOk returns a tuple with the DiscountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCode

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetDiscountCode(v string)`

SetDiscountCode sets DiscountCode field to given value.

### HasDiscountCode

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasDiscountCode() bool`

HasDiscountCode returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMemberEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberEmail() string`

GetMemberEmail returns the MemberEmail field if non-nil, zero value otherwise.

### GetMemberEmailOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberEmailOk() (*string, bool)`

GetMemberEmailOk returns a tuple with the MemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetMemberEmail(v string)`

SetMemberEmail sets MemberEmail field to given value.

### HasMemberEmail

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasMemberEmail() bool`

HasMemberEmail returns a boolean if a field has been set.

### GetMemberFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberFirstName() string`

GetMemberFirstName returns the MemberFirstName field if non-nil, zero value otherwise.

### GetMemberFirstNameOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberFirstNameOk() (*string, bool)`

GetMemberFirstNameOk returns a tuple with the MemberFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetMemberFirstName(v string)`

SetMemberFirstName sets MemberFirstName field to given value.

### HasMemberFirstName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasMemberFirstName() bool`

HasMemberFirstName returns a boolean if a field has been set.

### GetMemberLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberLastName() string`

GetMemberLastName returns the MemberLastName field if non-nil, zero value otherwise.

### GetMemberLastNameOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetMemberLastNameOk() (*string, bool)`

GetMemberLastNameOk returns a tuple with the MemberLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetMemberLastName(v string)`

SetMemberLastName sets MemberLastName field to given value.

### HasMemberLastName

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasMemberLastName() bool`

HasMemberLastName returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeApiMerchantMemberOperationLogListReq) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeApiMerchantMemberOperationLogListReq) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


