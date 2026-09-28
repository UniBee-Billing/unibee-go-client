# UnibeeApiMerchantPaymentExpiringCardUserListReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Count of page | [optional] 
**ExpireDateEnd** | Pointer to **string** | UTC YYYY-MM-DD inclusive | [optional] 
**ExpireDateStart** | Pointer to **string** | UTC YYYY-MM-DD inclusive, cannot be before today | [optional] 
**ExpireInDays** | Pointer to **int32** | Exact remaining days, e.g. 14|7|3 | [optional] 
**GatewayId** | Pointer to **int64** | Optional Stripe gateway id; default all Stripe gateways of this merchant | [optional] 
**Page** | Pointer to **int32** | Page, start 0 | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentExpiringCardUserListReq

`func NewUnibeeApiMerchantPaymentExpiringCardUserListReq() *UnibeeApiMerchantPaymentExpiringCardUserListReq`

NewUnibeeApiMerchantPaymentExpiringCardUserListReq instantiates a new UnibeeApiMerchantPaymentExpiringCardUserListReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExpiringCardUserListReqWithDefaults

`func NewUnibeeApiMerchantPaymentExpiringCardUserListReqWithDefaults() *UnibeeApiMerchantPaymentExpiringCardUserListReq`

NewUnibeeApiMerchantPaymentExpiringCardUserListReqWithDefaults instantiates a new UnibeeApiMerchantPaymentExpiringCardUserListReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetExpireDateEnd

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireDateEnd() string`

GetExpireDateEnd returns the ExpireDateEnd field if non-nil, zero value otherwise.

### GetExpireDateEndOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireDateEndOk() (*string, bool)`

GetExpireDateEndOk returns a tuple with the ExpireDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDateEnd

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetExpireDateEnd(v string)`

SetExpireDateEnd sets ExpireDateEnd field to given value.

### HasExpireDateEnd

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasExpireDateEnd() bool`

HasExpireDateEnd returns a boolean if a field has been set.

### GetExpireDateStart

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireDateStart() string`

GetExpireDateStart returns the ExpireDateStart field if non-nil, zero value otherwise.

### GetExpireDateStartOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireDateStartOk() (*string, bool)`

GetExpireDateStartOk returns a tuple with the ExpireDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDateStart

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetExpireDateStart(v string)`

SetExpireDateStart sets ExpireDateStart field to given value.

### HasExpireDateStart

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasExpireDateStart() bool`

HasExpireDateStart returns a boolean if a field has been set.

### GetExpireInDays

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireInDays() int32`

GetExpireInDays returns the ExpireInDays field if non-nil, zero value otherwise.

### GetExpireInDaysOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetExpireInDaysOk() (*int32, bool)`

GetExpireInDaysOk returns a tuple with the ExpireInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireInDays

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetExpireInDays(v int32)`

SetExpireInDays sets ExpireInDays field to given value.

### HasExpireInDays

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasExpireInDays() bool`

HasExpireInDays returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetPage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListReq) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


