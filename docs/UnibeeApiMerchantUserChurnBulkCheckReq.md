# UnibeeApiMerchantUserChurnBulkCheckReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **string** | Email list, split by comma / semicolon / newline / tab, max 1000. Deduped case-insensitively, input order kept. | 

## Methods

### NewUnibeeApiMerchantUserChurnBulkCheckReq

`func NewUnibeeApiMerchantUserChurnBulkCheckReq(emails string, ) *UnibeeApiMerchantUserChurnBulkCheckReq`

NewUnibeeApiMerchantUserChurnBulkCheckReq instantiates a new UnibeeApiMerchantUserChurnBulkCheckReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantUserChurnBulkCheckReqWithDefaults

`func NewUnibeeApiMerchantUserChurnBulkCheckReqWithDefaults() *UnibeeApiMerchantUserChurnBulkCheckReq`

NewUnibeeApiMerchantUserChurnBulkCheckReqWithDefaults instantiates a new UnibeeApiMerchantUserChurnBulkCheckReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UnibeeApiMerchantUserChurnBulkCheckReq) GetEmails() string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UnibeeApiMerchantUserChurnBulkCheckReq) GetEmailsOk() (*string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UnibeeApiMerchantUserChurnBulkCheckReq) SetEmails(v string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


