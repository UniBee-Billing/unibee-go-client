# MerchantPaymentExpiringCardUserListGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | data preparing when ready&#x3D;false | [optional] 
**Ready** | Pointer to **bool** | false while historical backfill is still running | [optional] 
**Total** | Pointer to **int32** | Total users after dedup | [optional] 
**Users** | Pointer to [**[]UnibeeApiBeanExpiringCardUser**](UnibeeApiBeanExpiringCardUser.md) | User list, one row per user | [optional] 

## Methods

### NewMerchantPaymentExpiringCardUserListGet200ResponseData

`func NewMerchantPaymentExpiringCardUserListGet200ResponseData() *MerchantPaymentExpiringCardUserListGet200ResponseData`

NewMerchantPaymentExpiringCardUserListGet200ResponseData instantiates a new MerchantPaymentExpiringCardUserListGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantPaymentExpiringCardUserListGet200ResponseDataWithDefaults

`func NewMerchantPaymentExpiringCardUserListGet200ResponseDataWithDefaults() *MerchantPaymentExpiringCardUserListGet200ResponseData`

NewMerchantPaymentExpiringCardUserListGet200ResponseDataWithDefaults instantiates a new MerchantPaymentExpiringCardUserListGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReady

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetTotal

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetUsers() []UnibeeApiBeanExpiringCardUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) GetUsersOk() (*[]UnibeeApiBeanExpiringCardUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) SetUsers(v []UnibeeApiBeanExpiringCardUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *MerchantPaymentExpiringCardUserListGet200ResponseData) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


