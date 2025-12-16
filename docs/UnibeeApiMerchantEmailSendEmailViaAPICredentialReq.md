# UnibeeApiMerchantEmailSendEmailViaAPICredentialReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCredential** | Pointer to [**UnibeeApiBeanEmailGatewayConnectionAPIKeys**](UnibeeApiBeanEmailGatewayConnectionAPIKeys.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Email** | **string** | Email | 
**GatewayName** | **string** | The name of email gateway, available for sendgrid|smtp, default gateway will use if not provide | 
**Language** | Pointer to **string** | Language | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Variables** | Pointer to **map[string]map[string]interface{}** | Variables，Map | [optional] 

## Methods

### NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReq

`func NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReq(email string, gatewayName string, ) *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq`

NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReq instantiates a new UnibeeApiMerchantEmailSendEmailViaAPICredentialReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReqWithDefaults

`func NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReqWithDefaults() *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq`

NewUnibeeApiMerchantEmailSendEmailViaAPICredentialReqWithDefaults instantiates a new UnibeeApiMerchantEmailSendEmailViaAPICredentialReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCredential

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetApiCredential() UnibeeApiBeanEmailGatewayConnectionAPIKeys`

GetApiCredential returns the ApiCredential field if non-nil, zero value otherwise.

### GetApiCredentialOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetApiCredentialOk() (*UnibeeApiBeanEmailGatewayConnectionAPIKeys, bool)`

GetApiCredentialOk returns a tuple with the ApiCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCredential

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetApiCredential(v UnibeeApiBeanEmailGatewayConnectionAPIKeys)`

SetApiCredential sets ApiCredential field to given value.

### HasApiCredential

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) HasApiCredential() bool`

HasApiCredential returns a boolean if a field has been set.

### GetContent

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetGatewayName

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.


### GetLanguage

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSubject

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetVariables

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetVariables() map[string]map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) GetVariablesOk() (*map[string]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) SetVariables(v map[string]map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UnibeeApiMerchantEmailSendEmailViaAPICredentialReq) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


