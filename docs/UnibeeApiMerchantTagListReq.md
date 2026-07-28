# UnibeeApiMerchantTagListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count per page, default 100 | [optional] 
**Page** | Pointer to **int32** | Page, start 0 | [optional] 
**Scope** | **int32** | 0-plan | 
**SearchKey** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeApiMerchantTagListReq

`func NewUnibeeApiMerchantTagListReq(scope int32, ) *UnibeeApiMerchantTagListReq`

NewUnibeeApiMerchantTagListReq instantiates a new UnibeeApiMerchantTagListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantTagListReqWithDefaults

`func NewUnibeeApiMerchantTagListReqWithDefaults() *UnibeeApiMerchantTagListReq`

NewUnibeeApiMerchantTagListReqWithDefaults instantiates a new UnibeeApiMerchantTagListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantTagListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantTagListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantTagListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantTagListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantTagListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantTagListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantTagListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantTagListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetScope

`func (o *UnibeeApiMerchantTagListReq) GetScope() int32`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UnibeeApiMerchantTagListReq) GetScopeOk() (*int32, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UnibeeApiMerchantTagListReq) SetScope(v int32)`

SetScope sets Scope field to given value.


### GetSearchKey

`func (o *UnibeeApiMerchantTagListReq) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *UnibeeApiMerchantTagListReq) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *UnibeeApiMerchantTagListReq) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *UnibeeApiMerchantTagListReq) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


