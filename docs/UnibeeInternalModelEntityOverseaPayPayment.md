# UnibeeInternalModelEntityOverseaPayPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | app id | [optional] 
**AuthorizeReason** | Pointer to **string** |  | [optional] 
**AuthorizeStatus** | Pointer to **int32** | authorize status，0-waiting authorize，1-authorized，2-authorized_request | [optional] 
**Automatic** | Pointer to **int32** | 0-no,1-yes | [optional] 
**BalanceAmount** | Pointer to **int64** | balance_amount | [optional] 
**BalanceEnd** | Pointer to **int64** | balance_end, utc time | [optional] 
**BalanceStart** | Pointer to **int64** | balance_start, utc time | [optional] 
**BillingReason** | Pointer to **string** |  | [optional] 
**BizType** | Pointer to **int32** | biz_type 1-single payment, 3-subscription | [optional] 
**CancelTime** | Pointer to **int64** | cancel time, utc time | [optional] 
**CaptureDelayHours** | Pointer to **int32** | capture_delay_hours | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CompanyId** | Pointer to **int64** | company id | [optional] 
**CountryCode** | Pointer to **string** | country code | [optional] 
**CreateTime** | Pointer to **int64** | create time, utc time | [optional] 
**Currency** | Pointer to **string** | currency，“SGD” “MYR” “PHP” “IDR” “THB” | [optional] 
**ExternalPaymentId** | Pointer to **string** | external_payment_id | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**GatewayEdition** | Pointer to **string** | gateway edition | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayPaymentId** | Pointer to **string** | gateway_payment_id | [optional] 
**GatewayPaymentIntentId** | Pointer to **string** | gateway_payment_intent_id | [optional] 
**GatewayPaymentMethod** | Pointer to **string** |  | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**HidePaymentMethods** | Pointer to **string** | hide_payment_methods | [optional] 
**Id** | Pointer to **int64** | id | [optional] 
**InvoiceData** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** | invoice id | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MetaData** | Pointer to **string** | meta_data (json) | [optional] 
**OpenApiId** | Pointer to **int64** | open api id | [optional] 
**PaidTime** | Pointer to **int64** | paid time, utc time | [optional] 
**PaymentAmount** | Pointer to **int64** | payment_amount | [optional] 
**PaymentData** | Pointer to **string** | payment create data (json) | [optional] 
**PaymentId** | Pointer to **string** | payment id | [optional] 
**RefundAmount** | Pointer to **int64** | total refund amount | [optional] 
**ReturnUrl** | Pointer to **string** | return url | [optional] 
**Status** | Pointer to **int32** | status  10-pending，20-success，30-failure, 40-cancel | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TerminalIp** | Pointer to **string** | client ip | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TotalAmount** | Pointer to **int64** | total amount | [optional] 
**UniqueId** | Pointer to **string** | unique id | [optional] 
**UserId** | Pointer to **int64** | user_id | [optional] 
**Verify** | Pointer to **string** | codeVerify | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPayPayment

`func NewUnibeeInternalModelEntityOverseaPayPayment() *UnibeeInternalModelEntityOverseaPayPayment`

NewUnibeeInternalModelEntityOverseaPayPayment instantiates a new UnibeeInternalModelEntityOverseaPayPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPayPaymentWithDefaults

`func NewUnibeeInternalModelEntityOverseaPayPaymentWithDefaults() *UnibeeInternalModelEntityOverseaPayPayment`

NewUnibeeInternalModelEntityOverseaPayPaymentWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPayPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAuthorizeReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAuthorizeReason() string`

GetAuthorizeReason returns the AuthorizeReason field if non-nil, zero value otherwise.

### GetAuthorizeReasonOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAuthorizeReasonOk() (*string, bool)`

GetAuthorizeReasonOk returns a tuple with the AuthorizeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetAuthorizeReason(v string)`

SetAuthorizeReason sets AuthorizeReason field to given value.

### HasAuthorizeReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasAuthorizeReason() bool`

HasAuthorizeReason returns a boolean if a field has been set.

### GetAuthorizeStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAuthorizeStatus() int32`

GetAuthorizeStatus returns the AuthorizeStatus field if non-nil, zero value otherwise.

### GetAuthorizeStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAuthorizeStatusOk() (*int32, bool)`

GetAuthorizeStatusOk returns a tuple with the AuthorizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetAuthorizeStatus(v int32)`

SetAuthorizeStatus sets AuthorizeStatus field to given value.

### HasAuthorizeStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasAuthorizeStatus() bool`

HasAuthorizeStatus returns a boolean if a field has been set.

### GetAutomatic

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAutomatic() int32`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetAutomaticOk() (*int32, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetAutomatic(v int32)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetBalanceAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceAmount() int64`

GetBalanceAmount returns the BalanceAmount field if non-nil, zero value otherwise.

### GetBalanceAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceAmountOk() (*int64, bool)`

GetBalanceAmountOk returns a tuple with the BalanceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetBalanceAmount(v int64)`

SetBalanceAmount sets BalanceAmount field to given value.

### HasBalanceAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasBalanceAmount() bool`

HasBalanceAmount returns a boolean if a field has been set.

### GetBalanceEnd

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceEnd() int64`

GetBalanceEnd returns the BalanceEnd field if non-nil, zero value otherwise.

### GetBalanceEndOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceEndOk() (*int64, bool)`

GetBalanceEndOk returns a tuple with the BalanceEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceEnd

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetBalanceEnd(v int64)`

SetBalanceEnd sets BalanceEnd field to given value.

### HasBalanceEnd

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasBalanceEnd() bool`

HasBalanceEnd returns a boolean if a field has been set.

### GetBalanceStart

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceStart() int64`

GetBalanceStart returns the BalanceStart field if non-nil, zero value otherwise.

### GetBalanceStartOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBalanceStartOk() (*int64, bool)`

GetBalanceStartOk returns a tuple with the BalanceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceStart

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetBalanceStart(v int64)`

SetBalanceStart sets BalanceStart field to given value.

### HasBalanceStart

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasBalanceStart() bool`

HasBalanceStart returns a boolean if a field has been set.

### GetBillingReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBillingReason() string`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBillingReasonOk() (*string, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetBillingReason(v string)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetBizType

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBizType() int32`

GetBizType returns the BizType field if non-nil, zero value otherwise.

### GetBizTypeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetBizTypeOk() (*int32, bool)`

GetBizTypeOk returns a tuple with the BizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizType

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetBizType(v int32)`

SetBizType sets BizType field to given value.

### HasBizType

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasBizType() bool`

HasBizType returns a boolean if a field has been set.

### GetCancelTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCancelTime() int64`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCancelTimeOk() (*int64, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCancelTime(v int64)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCaptureDelayHours

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCaptureDelayHours() int32`

GetCaptureDelayHours returns the CaptureDelayHours field if non-nil, zero value otherwise.

### GetCaptureDelayHoursOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCaptureDelayHoursOk() (*int32, bool)`

GetCaptureDelayHoursOk returns a tuple with the CaptureDelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureDelayHours

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCaptureDelayHours(v int32)`

SetCaptureDelayHours sets CaptureDelayHours field to given value.

### HasCaptureDelayHours

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCaptureDelayHours() bool`

HasCaptureDelayHours returns a boolean if a field has been set.

### GetCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCompanyId() int64`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCompanyIdOk() (*int64, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCompanyId(v int64)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExternalPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetExternalPaymentId() string`

GetExternalPaymentId returns the ExternalPaymentId field if non-nil, zero value otherwise.

### GetExternalPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetExternalPaymentIdOk() (*string, bool)`

GetExternalPaymentIdOk returns a tuple with the ExternalPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetExternalPaymentId(v string)`

SetExternalPaymentId sets ExternalPaymentId field to given value.

### HasExternalPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasExternalPaymentId() bool`

HasExternalPaymentId returns a boolean if a field has been set.

### GetFailureReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGatewayEdition

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayEdition() string`

GetGatewayEdition returns the GatewayEdition field if non-nil, zero value otherwise.

### GetGatewayEditionOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayEditionOk() (*string, bool)`

GetGatewayEditionOk returns a tuple with the GatewayEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayEdition

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGatewayEdition(v string)`

SetGatewayEdition sets GatewayEdition field to given value.

### HasGatewayEdition

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGatewayEdition() bool`

HasGatewayEdition returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentId() string`

GetGatewayPaymentId returns the GatewayPaymentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentIdOk() (*string, bool)`

GetGatewayPaymentIdOk returns a tuple with the GatewayPaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGatewayPaymentId(v string)`

SetGatewayPaymentId sets GatewayPaymentId field to given value.

### HasGatewayPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGatewayPaymentId() bool`

HasGatewayPaymentId returns a boolean if a field has been set.

### GetGatewayPaymentIntentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentIntentId() string`

GetGatewayPaymentIntentId returns the GatewayPaymentIntentId field if non-nil, zero value otherwise.

### GetGatewayPaymentIntentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentIntentIdOk() (*string, bool)`

GetGatewayPaymentIntentIdOk returns a tuple with the GatewayPaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentIntentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGatewayPaymentIntentId(v string)`

SetGatewayPaymentIntentId sets GatewayPaymentIntentId field to given value.

### HasGatewayPaymentIntentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGatewayPaymentIntentId() bool`

HasGatewayPaymentIntentId returns a boolean if a field has been set.

### GetGatewayPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentMethod() string`

GetGatewayPaymentMethod returns the GatewayPaymentMethod field if non-nil, zero value otherwise.

### GetGatewayPaymentMethodOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGatewayPaymentMethodOk() (*string, bool)`

GetGatewayPaymentMethodOk returns a tuple with the GatewayPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGatewayPaymentMethod(v string)`

SetGatewayPaymentMethod sets GatewayPaymentMethod field to given value.

### HasGatewayPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGatewayPaymentMethod() bool`

HasGatewayPaymentMethod returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetHidePaymentMethods

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetHidePaymentMethods() string`

GetHidePaymentMethods returns the HidePaymentMethods field if non-nil, zero value otherwise.

### GetHidePaymentMethodsOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetHidePaymentMethodsOk() (*string, bool)`

GetHidePaymentMethodsOk returns a tuple with the HidePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePaymentMethods

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetHidePaymentMethods(v string)`

SetHidePaymentMethods sets HidePaymentMethods field to given value.

### HasHidePaymentMethods

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasHidePaymentMethods() bool`

HasHidePaymentMethods returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetInvoiceData() string`

GetInvoiceData returns the InvoiceData field if non-nil, zero value otherwise.

### GetInvoiceDataOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetInvoiceDataOk() (*string, bool)`

GetInvoiceDataOk returns a tuple with the InvoiceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetInvoiceData(v string)`

SetInvoiceData sets InvoiceData field to given value.

### HasInvoiceData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasInvoiceData() bool`

HasInvoiceData returns a boolean if a field has been set.

### GetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetOpenApiId() int64`

GetOpenApiId returns the OpenApiId field if non-nil, zero value otherwise.

### GetOpenApiIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetOpenApiIdOk() (*int64, bool)`

GetOpenApiIdOk returns a tuple with the OpenApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetOpenApiId(v int64)`

SetOpenApiId sets OpenApiId field to given value.

### HasOpenApiId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasOpenApiId() bool`

HasOpenApiId returns a boolean if a field has been set.

### GetPaidTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaidTime() int64`

GetPaidTime returns the PaidTime field if non-nil, zero value otherwise.

### GetPaidTimeOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaidTimeOk() (*int64, bool)`

GetPaidTimeOk returns a tuple with the PaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetPaidTime(v int64)`

SetPaidTime sets PaidTime field to given value.

### HasPaidTime

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasPaidTime() bool`

HasPaidTime returns a boolean if a field has been set.

### GetPaymentAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentAmount() int64`

GetPaymentAmount returns the PaymentAmount field if non-nil, zero value otherwise.

### GetPaymentAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentAmountOk() (*int64, bool)`

GetPaymentAmountOk returns a tuple with the PaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetPaymentAmount(v int64)`

SetPaymentAmount sets PaymentAmount field to given value.

### HasPaymentAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasPaymentAmount() bool`

HasPaymentAmount returns a boolean if a field has been set.

### GetPaymentData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentData() string`

GetPaymentData returns the PaymentData field if non-nil, zero value otherwise.

### GetPaymentDataOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentDataOk() (*string, bool)`

GetPaymentDataOk returns a tuple with the PaymentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetPaymentData(v string)`

SetPaymentData sets PaymentData field to given value.

### HasPaymentData

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasPaymentData() bool`

HasPaymentData returns a boolean if a field has been set.

### GetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetRefundAmount() int64`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetRefundAmountOk() (*int64, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetRefundAmount(v int64)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTerminalIp

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetTerminalIp() string`

GetTerminalIp returns the TerminalIp field if non-nil, zero value otherwise.

### GetTerminalIpOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetTerminalIpOk() (*string, bool)`

GetTerminalIpOk returns a tuple with the TerminalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalIp

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetTerminalIp(v string)`

SetTerminalIp sets TerminalIp field to given value.

### HasTerminalIp

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasTerminalIp() bool`

HasTerminalIp returns a boolean if a field has been set.

### GetToken

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVerify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetVerify() string`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *UnibeeInternalModelEntityOverseaPayPayment) GetVerifyOk() (*string, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) SetVerify(v string)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *UnibeeInternalModelEntityOverseaPayPayment) HasVerify() bool`

HasVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


