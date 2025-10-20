# UnibeeApiMerchantEmailSendEmailToUserReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachInvoiceId** | Pointer to **string** | AttachInvoiceId | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Email** | **string** | Email | 
**GatewayTemplateId** | Pointer to **string** | GatewayTemplateId | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** | Variables，Map | [optional] 

## Methods

### NewUnibeeApiMerchantEmailSendEmailToUserReq

`func NewUnibeeApiMerchantEmailSendEmailToUserReq(email string, ) *UnibeeApiMerchantEmailSendEmailToUserReq`

NewUnibeeApiMerchantEmailSendEmailToUserReq instantiates a new UnibeeApiMerchantEmailSendEmailToUserReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailSendEmailToUserReqWithDefaults

`func NewUnibeeApiMerchantEmailSendEmailToUserReqWithDefaults() *UnibeeApiMerchantEmailSendEmailToUserReq`

NewUnibeeApiMerchantEmailSendEmailToUserReqWithDefaults instantiates a new UnibeeApiMerchantEmailSendEmailToUserReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetAttachInvoiceId() string`

GetAttachInvoiceId returns the AttachInvoiceId field if non-nil, zero value otherwise.

### GetAttachInvoiceIdOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetAttachInvoiceIdOk() (*string, bool)`

GetAttachInvoiceIdOk returns a tuple with the AttachInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetAttachInvoiceId(v string)`

SetAttachInvoiceId sets AttachInvoiceId field to given value.

### HasAttachInvoiceId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) HasAttachInvoiceId() bool`

HasAttachInvoiceId returns a boolean if a field has been set.

### GetContent

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGatewayTemplateId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetGatewayTemplateId() string`

GetGatewayTemplateId returns the GatewayTemplateId field if non-nil, zero value otherwise.

### GetGatewayTemplateIdOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetGatewayTemplateIdOk() (*string, bool)`

GetGatewayTemplateIdOk returns a tuple with the GatewayTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayTemplateId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetGatewayTemplateId(v string)`

SetGatewayTemplateId sets GatewayTemplateId field to given value.

### HasGatewayTemplateId

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) HasGatewayTemplateId() bool`

HasGatewayTemplateId returns a boolean if a field has been set.

### GetSubject

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVariables

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UnibeeApiMerchantEmailSendEmailToUserReq) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


