# MerchantAuthSessionLoginPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantMember** | Pointer to [**UnibeeApiBeanDetailMerchantMemberDetail**](UnibeeApiBeanDetailMerchantMemberDetail.md) |  | [optional] 
**ReturnUrl** | Pointer to **string** | Return URL | [optional] 
**Token** | Pointer to **string** | Access token of admin portal | [optional] 

## Methods

### NewMerchantAuthSessionLoginPost200ResponseData

`func NewMerchantAuthSessionLoginPost200ResponseData() *MerchantAuthSessionLoginPost200ResponseData`

NewMerchantAuthSessionLoginPost200ResponseData instantiates a new MerchantAuthSessionLoginPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantAuthSessionLoginPost200ResponseDataWithDefaults

`func NewMerchantAuthSessionLoginPost200ResponseDataWithDefaults() *MerchantAuthSessionLoginPost200ResponseData`

NewMerchantAuthSessionLoginPost200ResponseDataWithDefaults instantiates a new MerchantAuthSessionLoginPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantMember

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetMerchantMember() UnibeeApiBeanDetailMerchantMemberDetail`

GetMerchantMember returns the MerchantMember field if non-nil, zero value otherwise.

### GetMerchantMemberOk

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetMerchantMemberOk() (*UnibeeApiBeanDetailMerchantMemberDetail, bool)`

GetMerchantMemberOk returns a tuple with the MerchantMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMember

`func (o *MerchantAuthSessionLoginPost200ResponseData) SetMerchantMember(v UnibeeApiBeanDetailMerchantMemberDetail)`

SetMerchantMember sets MerchantMember field to given value.

### HasMerchantMember

`func (o *MerchantAuthSessionLoginPost200ResponseData) HasMerchantMember() bool`

HasMerchantMember returns a boolean if a field has been set.

### GetReturnUrl

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *MerchantAuthSessionLoginPost200ResponseData) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *MerchantAuthSessionLoginPost200ResponseData) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetToken

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MerchantAuthSessionLoginPost200ResponseData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MerchantAuthSessionLoginPost200ResponseData) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *MerchantAuthSessionLoginPost200ResponseData) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


