# UnibeeApiMerchantEmailSendTemplateEmailToUserReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachInvoiceId** | Pointer to **string** | AttachInvoiceId | [optional] 
**TemplateName** | **string** | The name of email template | 
**UserId** | **int64** | UserId | 
**Variables** | Pointer to **map[string]map[string]interface{}** | Variables，Map | [optional] 

## Methods

### NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq

`func NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq(templateName string, userId int64, ) *UnibeeApiMerchantEmailSendTemplateEmailToUserReq`

NewUnibeeApiMerchantEmailSendTemplateEmailToUserReq instantiates a new UnibeeApiMerchantEmailSendTemplateEmailToUserReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailSendTemplateEmailToUserReqWithDefaults

`func NewUnibeeApiMerchantEmailSendTemplateEmailToUserReqWithDefaults() *UnibeeApiMerchantEmailSendTemplateEmailToUserReq`

NewUnibeeApiMerchantEmailSendTemplateEmailToUserReqWithDefaults instantiates a new UnibeeApiMerchantEmailSendTemplateEmailToUserReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetAttachInvoiceId() string`

GetAttachInvoiceId returns the AttachInvoiceId field if non-nil, zero value otherwise.

### GetAttachInvoiceIdOk

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetAttachInvoiceIdOk() (*string, bool)`

GetAttachInvoiceIdOk returns a tuple with the AttachInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetAttachInvoiceId(v string)`

SetAttachInvoiceId sets AttachInvoiceId field to given value.

### HasAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) HasAttachInvoiceId() bool`

HasAttachInvoiceId returns a boolean if a field has been set.

### GetTemplateName

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.


### GetUserId

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetVariables

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UnibeeApiMerchantEmailSendTemplateEmailToUserReq) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


