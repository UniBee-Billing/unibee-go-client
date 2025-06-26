# UnibeeApiMerchantMemberClearTotpReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | Pointer to **int64** | The unique id of member | [optional] 
**TotpCode** | Pointer to **string** | The admin totp code | [optional] 

## Methods

### NewUnibeeApiMerchantMemberClearTotpReq

`func NewUnibeeApiMerchantMemberClearTotpReq() *UnibeeApiMerchantMemberClearTotpReq`

NewUnibeeApiMerchantMemberClearTotpReq instantiates a new UnibeeApiMerchantMemberClearTotpReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantMemberClearTotpReqWithDefaults

`func NewUnibeeApiMerchantMemberClearTotpReqWithDefaults() *UnibeeApiMerchantMemberClearTotpReq`

NewUnibeeApiMerchantMemberClearTotpReqWithDefaults instantiates a new UnibeeApiMerchantMemberClearTotpReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *UnibeeApiMerchantMemberClearTotpReq) GetMemberId() int64`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *UnibeeApiMerchantMemberClearTotpReq) GetMemberIdOk() (*int64, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *UnibeeApiMerchantMemberClearTotpReq) SetMemberId(v int64)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *UnibeeApiMerchantMemberClearTotpReq) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetTotpCode

`func (o *UnibeeApiMerchantMemberClearTotpReq) GetTotpCode() string`

GetTotpCode returns the TotpCode field if non-nil, zero value otherwise.

### GetTotpCodeOk

`func (o *UnibeeApiMerchantMemberClearTotpReq) GetTotpCodeOk() (*string, bool)`

GetTotpCodeOk returns a tuple with the TotpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotpCode

`func (o *UnibeeApiMerchantMemberClearTotpReq) SetTotpCode(v string)`

SetTotpCode sets TotpCode field to given value.

### HasTotpCode

`func (o *UnibeeApiMerchantMemberClearTotpReq) HasTotpCode() bool`

HasTotpCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


