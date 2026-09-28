# UnibeeApiMerchantUserTagListUsersReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count per page, default 20, max 100 | [optional] 
**Mode** | Pointer to **string** | any|all, default any | [optional] 
**Page** | Pointer to **int32** | Page, start 0 | [optional] 
**TagCodes** | **[]string** | System tag codes, e.g. billing.first_monetary_payment | 

## Methods

### NewUnibeeApiMerchantUserTagListUsersReq

`func NewUnibeeApiMerchantUserTagListUsersReq(tagCodes []string, ) *UnibeeApiMerchantUserTagListUsersReq`

NewUnibeeApiMerchantUserTagListUsersReq instantiates a new UnibeeApiMerchantUserTagListUsersReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserTagListUsersReqWithDefaults

`func NewUnibeeApiMerchantUserTagListUsersReqWithDefaults() *UnibeeApiMerchantUserTagListUsersReq`

NewUnibeeApiMerchantUserTagListUsersReqWithDefaults instantiates a new UnibeeApiMerchantUserTagListUsersReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantUserTagListUsersReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantUserTagListUsersReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMode

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *UnibeeApiMerchantUserTagListUsersReq) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *UnibeeApiMerchantUserTagListUsersReq) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantUserTagListUsersReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantUserTagListUsersReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTagCodes

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetTagCodes() []string`

GetTagCodes returns the TagCodes field if non-nil, zero value otherwise.

### GetTagCodesOk

`func (o *UnibeeApiMerchantUserTagListUsersReq) GetTagCodesOk() (*[]string, bool)`

GetTagCodesOk returns a tuple with the TagCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCodes

`func (o *UnibeeApiMerchantUserTagListUsersReq) SetTagCodes(v []string)`

SetTagCodes sets TagCodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


