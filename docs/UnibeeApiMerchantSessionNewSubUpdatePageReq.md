# UnibeeApiMerchantSessionNewSubUpdatePageReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | Pointer to **string** | CancelUrl | [optional] 
**Email** | Pointer to **string** | Email, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**ExternalUserId** | Pointer to **string** | ExternalUserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**PlanId** | Pointer to **int64** | Id of plan to update | [optional] 
**ProductId** | Pointer to **int64** | default product will use if productId not specified | [optional] 
**ReturnUrl** | Pointer to **string** | ReturnUrl | [optional] 
**UserId** | Pointer to **int64** | UserId, unique, either ExternalUserId&amp;Email or UserId needed | [optional] 
**VatCountryCode** | Pointer to **string** | Vat Country Code | [optional] 

## Methods

### NewUnibeeApiMerchantSessionNewSubUpdatePageReq

`func NewUnibeeApiMerchantSessionNewSubUpdatePageReq() *UnibeeApiMerchantSessionNewSubUpdatePageReq`

NewUnibeeApiMerchantSessionNewSubUpdatePageReq instantiates a new UnibeeApiMerchantSessionNewSubUpdatePageReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSessionNewSubUpdatePageReqWithDefaults

`func NewUnibeeApiMerchantSessionNewSubUpdatePageReqWithDefaults() *UnibeeApiMerchantSessionNewSubUpdatePageReq`

NewUnibeeApiMerchantSessionNewSubUpdatePageReqWithDefaults instantiates a new UnibeeApiMerchantSessionNewSubUpdatePageReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProductId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatCountryCode

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetVatCountryCode() string`

GetVatCountryCode returns the VatCountryCode field if non-nil, zero value otherwise.

### GetVatCountryCodeOk

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) GetVatCountryCodeOk() (*string, bool)`

GetVatCountryCodeOk returns a tuple with the VatCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatCountryCode

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) SetVatCountryCode(v string)`

SetVatCountryCode sets VatCountryCode field to given value.

### HasVatCountryCode

`func (o *UnibeeApiMerchantSessionNewSubUpdatePageReq) HasVatCountryCode() bool`

HasVatCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


