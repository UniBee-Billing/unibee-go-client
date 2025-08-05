# UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription. | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | Required, UTC time, the current period end time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**CurrentPeriodStart** | Pointer to **string** | Required, UTC time, the current period start time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**Email** | Pointer to **string** | The email of user, one of Email or ExternalUserId is required | [optional] 
**ExternalPlanId** | Pointer to **string** | The external id of plan, one of planId or ExternalPlanId is required, plan should created at UniBee at first | [optional] 
**ExternalSubscriptionId** | Pointer to **string** | Required, The external id of subscription | [optional] 
**ExternalUserId** | Pointer to **string** | The external id of user, one of Email or ExternalUserId is required  | [optional] 
**Gateway** | Pointer to **string** | Required, should one of stripe|paypal|wire_transfer|changelly  | [optional] 
**PlanId** | Pointer to **int64** | The id of plan, one of planId or ExternalPlanId is required, plan should created at UniBee at first  | [optional] 
**Quantity** | Pointer to **int64** | the quantity of plan, default 1 if not provided  | [optional] 
**TaxPercentage** | Pointer to **int64** | The TaxPercentage of subscription, Only applicable when the system VAT gateway not setup, 1000 &#x3D; 10% | [optional] 
**TotalAmount** | Pointer to **int64** | Required. Unit: cents. | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReq

`func NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReq() *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq`

NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReq instantiates a new UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReqWithDefaults() *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq`

NewUnibeeApiMerchantSubscriptionHistorySubscriptionImportReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *UnibeeApiMerchantSubscriptionHistorySubscriptionImportReq) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


