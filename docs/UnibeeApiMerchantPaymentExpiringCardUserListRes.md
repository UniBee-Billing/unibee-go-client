# UnibeeApiMerchantPaymentExpiringCardUserListRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | data preparing when ready&#x3D;false | [optional] 
**Ready** | Pointer to **bool** | false while historical backfill is still running | [optional] 
**Total** | Pointer to **int32** | Total users after dedup | [optional] 
**Users** | Pointer to [**[]UnibeeApiBeanExpiringCardUser**](UnibeeApiBeanExpiringCardUser.md) | User list, one row per user | [optional] 

## Methods

### NewUnibeeApiMerchantPaymentExpiringCardUserListRes

`func NewUnibeeApiMerchantPaymentExpiringCardUserListRes() *UnibeeApiMerchantPaymentExpiringCardUserListRes`

NewUnibeeApiMerchantPaymentExpiringCardUserListRes instantiates a new UnibeeApiMerchantPaymentExpiringCardUserListRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantPaymentExpiringCardUserListResWithDefaults

`func NewUnibeeApiMerchantPaymentExpiringCardUserListResWithDefaults() *UnibeeApiMerchantPaymentExpiringCardUserListRes`

NewUnibeeApiMerchantPaymentExpiringCardUserListResWithDefaults instantiates a new UnibeeApiMerchantPaymentExpiringCardUserListRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReady

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetTotal

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetUsers() []UnibeeApiBeanExpiringCardUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) GetUsersOk() (*[]UnibeeApiBeanExpiringCardUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) SetUsers(v []UnibeeApiBeanExpiringCardUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UnibeeApiMerchantPaymentExpiringCardUserListRes) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


