# UnibeeInternalModelEntityOverseaPayRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalData** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** | app id | [optional] 
**BizId** | Pointer to **string** | biz id,copy from payment.biz_id | [optional] 
**BizType** | Pointer to **int32** | biz type, copy from payment.biz_type | [optional] 
**CompanyId** | Pointer to **int64** | company id | [optional] 
**CountryCode** | Pointer to **string** | country code | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**GatewayId** | Pointer to **int32** | gateway_id | [optional] 
**GatewayRefundId** | Pointer to **string** | gateway refund id | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**MerchantId** | Pointer to **int32** | merchant id | [optional] 
**OpenApiId** | Pointer to **int64** | open api id | [optional] 
**PaymentId** | Pointer to **string** | relative payment id | [optional] 
**RefundAmount** | Pointer to **int64** | refund amount, cent | [optional] 
**RefundComment** | Pointer to **string** | refund comment | [optional] 
**RefundCommentExplain** | Pointer to **string** | refund comment | [optional] 
**RefundId** | Pointer to **string** | refund id (system generate) | [optional] 
**RefundTime** | Pointer to **int64** | refund success time | [optional] 
**ReturnUrl** | Pointer to **string** | return url after refund success | [optional] 
**Status** | Pointer to **int32** | status。10-pending，20-success，30-failure, 40-cancel | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayRefund

`func NewUnibeeInternalModelEntityOverseaPayRefund() *UnibeeInternalModelEntityOverseaPayRefund`

NewUnibeeInternalModelEntityOverseaPayRefund instantiates a new UnibeeInternalModelEntityOverseaPayRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayRefundWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayRefundWithDefaults() *UnibeeInternalModelEntityOverseaPayRefund`

NewUnibeeInternalModelEntityOverseaPayRefundWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalData

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetAdditionalData() string`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetAdditionalDataOk() (*string, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetAdditionalData(v string)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.

### GetAppId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetBizId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetBizId() string`

GetBizId returns the BizId field if non-nil, zero value otherwise.

### GetBizIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetBizIdOk() (*string, bool)`

GetBizIdOk returns a tuple with the BizId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetBizId(v string)`

SetBizId sets BizId field to given value.

### HasBizId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasBizId() bool`

HasBizId returns a boolean if a field has been set.

### GetBizType

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGatewayId() int32`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGatewayIdOk() (*int32, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetGatewayId(v int32)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGatewayRefundId() string`

GetGatewayRefundId returns the GatewayRefundId field if non-nil, zero value otherwise.

### GetGatewayRefundIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGatewayRefundIdOk() (*string, bool)`

GetGatewayRefundIdOk returns a tuple with the GatewayRefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetGatewayRefundId(v string)`

SetGatewayRefundId sets GatewayRefundId field to given value.

### HasGatewayRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasGatewayRefundId() bool`

HasGatewayRefundId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetMerchantId() int32`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetMerchantIdOk() (*int32, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetMerchantId(v int32)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetOpenApiId() int64`

GetOpenApiId returns the OpenApiId field if non-nil, zero value otherwise.

### GetOpenApiIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetOpenApiIdOk() (*int64, bool)`

GetOpenApiIdOk returns a tuple with the OpenApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetOpenApiId(v int64)`

SetOpenApiId sets OpenApiId field to given value.

### HasOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasOpenApiId() bool`

HasOpenApiId returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetRefundComment

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundComment() string`

GetRefundComment returns the RefundComment field if non-nil, zero value otherwise.

### GetRefundCommentOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundCommentOk() (*string, bool)`

GetRefundCommentOk returns a tuple with the RefundComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundComment

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetRefundComment(v string)`

SetRefundComment sets RefundComment field to given value.

### HasRefundComment

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasRefundComment() bool`

HasRefundComment returns a boolean if a field has been set.

### GetRefundCommentExplain

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundCommentExplain() string`

GetRefundCommentExplain returns the RefundCommentExplain field if non-nil, zero value otherwise.

### GetRefundCommentExplainOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundCommentExplainOk() (*string, bool)`

GetRefundCommentExplainOk returns a tuple with the RefundCommentExplain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCommentExplain

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetRefundCommentExplain(v string)`

SetRefundCommentExplain sets RefundCommentExplain field to given value.

### HasRefundCommentExplain

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasRefundCommentExplain() bool`

HasRefundCommentExplain returns a boolean if a field has been set.

### GetRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.

### HasRefundId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasRefundId() bool`

HasRefundId returns a boolean if a field has been set.

### GetRefundTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundTime() int64`

GetRefundTime returns the RefundTime field if non-nil, zero value otherwise.

### GetRefundTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetRefundTimeOk() (*int64, bool)`

GetRefundTimeOk returns a tuple with the RefundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetRefundTime(v int64)`

SetRefundTime sets RefundTime field to given value.

### HasRefundTime

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasRefundTime() bool`

HasRefundTime returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayRefund) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPayRefund) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


