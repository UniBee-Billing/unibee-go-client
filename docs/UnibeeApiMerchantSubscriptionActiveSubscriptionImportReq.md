# UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripePaymentMethod** | Pointer to **string** |  | [optional] 
**StripeUserId** | Pointer to **string** |  | [optional] 
**BillingCycleAnchor** | Pointer to **string** | Required, UTC time, The reference point that aligns future billing cycle dates. It sets the day of week for week intervals, the day of month for month and year intervals, and the month of year for year intervals, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**CountryCode** | Pointer to **string** | Required. Specifies the ISO 3166-1 alpha-2 country code for the subscription (e.g., EE, RU). This code determines the applicable tax rules for the subscription. | [optional] 
**CreateTime** | Pointer to **string** | Required, UTC time, the creation time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**CurrentPeriodEnd** | Pointer to **string** | Required, UTC time, the current period end time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**CurrentPeriodStart** | Pointer to **string** | Required, UTC time, the current period start time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**Email** | Pointer to **string** | The email of user, one of Email or ExternalUserId is required | [optional] 
**ExpectedTotalAmount** | Pointer to **int64** | Optional. Unit: cents. If greater than 0, the system will verify the calculated total amount against this value | [optional] 
**ExternalPlanId** | Pointer to **string** | The external id of plan, one of planId or ExternalPlanId is required, plan should created at UniBee at first | [optional] 
**ExternalSubscriptionId** | Pointer to **string** | Required, The external id of subscription | [optional] 
**ExternalUserId** | Pointer to **string** | The external id of user, one of Email or ExternalUserId is required  | [optional] 
**Features** | Pointer to **string** | In json format, additional features data of subscription, will join user&#39;s metric data in user api if provided&#39; | [optional] 
**FirstPaidTime** | Pointer to **string** | UTC time, the first payment success time of subscription, format &#39;2006-01-02 15:04:05&#39; | [optional] 
**Gateway** | Pointer to **string** | Required, should one of stripe|paypal|wire_transfer|changelly  | [optional] 
**GatewayPaymentType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]map[string]interface{}** | Metadata，Map | [optional] 
**PlanId** | Pointer to **int64** | The id of plan, one of planId or ExternalPlanId is required, plan should created at UniBee at first  | [optional] 
**Quantity** | Pointer to **int64** | the quantity of plan, default 1 if not provided  | [optional] 
**TaxPercentage** | Pointer to **int64** | The tax percentage. Only applicable when the system VAT gateway not setup. Value is in thousandths (e.g., 1000 &#x3D; 10%). | [optional] 
**VatNumber** | Pointer to **string** | The Vat Number of user | [optional] 

## Methods

### NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReq

`func NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReq() *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq`

NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReq instantiates a new UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReqWithDefaults

`func NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReqWithDefaults() *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq`

NewUnibeeApiMerchantSubscriptionActiveSubscriptionImportReqWithDefaults instantiates a new UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripePaymentMethod

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetStripePaymentMethod() string`

GetStripePaymentMethod returns the StripePaymentMethod field if non-nil, zero value otherwise.

### GetStripePaymentMethodOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetStripePaymentMethodOk() (*string, bool)`

GetStripePaymentMethodOk returns a tuple with the StripePaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePaymentMethod

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetStripePaymentMethod(v string)`

SetStripePaymentMethod sets StripePaymentMethod field to given value.

### HasStripePaymentMethod

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasStripePaymentMethod() bool`

HasStripePaymentMethod returns a boolean if a field has been set.

### GetStripeUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetStripeUserId() string`

GetStripeUserId returns the StripeUserId field if non-nil, zero value otherwise.

### GetStripeUserIdOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetStripeUserIdOk() (*string, bool)`

GetStripeUserIdOk returns a tuple with the StripeUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetStripeUserId(v string)`

SetStripeUserId sets StripeUserId field to given value.

### HasStripeUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasStripeUserId() bool`

HasStripeUserId returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetBillingCycleAnchor() string`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetBillingCycleAnchorOk() (*string, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetBillingCycleAnchor(v string)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetEmail

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpectedTotalAmount

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExpectedTotalAmount() int64`

GetExpectedTotalAmount returns the ExpectedTotalAmount field if non-nil, zero value otherwise.

### GetExpectedTotalAmountOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExpectedTotalAmountOk() (*int64, bool)`

GetExpectedTotalAmountOk returns a tuple with the ExpectedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedTotalAmount

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetExpectedTotalAmount(v int64)`

SetExpectedTotalAmount sets ExpectedTotalAmount field to given value.

### HasExpectedTotalAmount

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasExpectedTotalAmount() bool`

HasExpectedTotalAmount returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalSubscriptionId() string`

GetExternalSubscriptionId returns the ExternalSubscriptionId field if non-nil, zero value otherwise.

### GetExternalSubscriptionIdOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalSubscriptionIdOk() (*string, bool)`

GetExternalSubscriptionIdOk returns a tuple with the ExternalSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetExternalSubscriptionId(v string)`

SetExternalSubscriptionId sets ExternalSubscriptionId field to given value.

### HasExternalSubscriptionId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasExternalSubscriptionId() bool`

HasExternalSubscriptionId returns a boolean if a field has been set.

### GetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetFeatures

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetFirstPaidTime() string`

GetFirstPaidTime returns the FirstPaidTime field if non-nil, zero value otherwise.

### GetFirstPaidTimeOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetFirstPaidTimeOk() (*string, bool)`

GetFirstPaidTimeOk returns a tuple with the FirstPaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetFirstPaidTime(v string)`

SetFirstPaidTime sets FirstPaidTime field to given value.

### HasFirstPaidTime

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasFirstPaidTime() bool`

HasFirstPaidTime returns a boolean if a field has been set.

### GetGateway

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetGatewayPaymentType() string`

GetGatewayPaymentType returns the GatewayPaymentType field if non-nil, zero value otherwise.

### GetGatewayPaymentTypeOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetGatewayPaymentTypeOk() (*string, bool)`

GetGatewayPaymentTypeOk returns a tuple with the GatewayPaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetGatewayPaymentType(v string)`

SetGatewayPaymentType sets GatewayPaymentType field to given value.

### HasGatewayPaymentType

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasGatewayPaymentType() bool`

HasGatewayPaymentType returns a boolean if a field has been set.

### GetMetadata

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetTaxPercentage() int64`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetTaxPercentageOk() (*int64, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetTaxPercentage(v int64)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeApiMerchantSubscriptionActiveSubscriptionImportReq) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


