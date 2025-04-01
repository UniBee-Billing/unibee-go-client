# UnibeeInternalLogicMiddlewareLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPaid** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** | License, Premium Version will contain License | [optional] 
**OwnerEmail** | Pointer to **string** | OwnerEmail | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**Subscription** | Pointer to [**UnibeeApiBeanSubscription**](UnibeeApiBeanSubscription.md) |  | [optional] 
**Version** | Pointer to [**UnibeeInternalLogicMiddlewareMerchantVersion**](UnibeeInternalLogicMiddlewareMerchantVersion.md) |  | [optional] 

## Methods

### NewUnibeeInternalLogicMiddlewareLicense

`func NewUnibeeInternalLogicMiddlewareLicense() *UnibeeInternalLogicMiddlewareLicense`

NewUnibeeInternalLogicMiddlewareLicense instantiates a new UnibeeInternalLogicMiddlewareLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalLogicMiddlewareLicenseWithDefaults

`func NewUnibeeInternalLogicMiddlewareLicenseWithDefaults() *UnibeeInternalLogicMiddlewareLicense`

NewUnibeeInternalLogicMiddlewareLicenseWithDefaults instantiates a new UnibeeInternalLogicMiddlewareLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPaid

`func (o *UnibeeInternalLogicMiddlewareLicense) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *UnibeeInternalLogicMiddlewareLicense) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.

### HasIsPaid

`func (o *UnibeeInternalLogicMiddlewareLicense) HasIsPaid() bool`

HasIsPaid returns a boolean if a field has been set.

### GetLicense

`func (o *UnibeeInternalLogicMiddlewareLicense) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *UnibeeInternalLogicMiddlewareLicense) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *UnibeeInternalLogicMiddlewareLicense) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *UnibeeInternalLogicMiddlewareLicense) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *UnibeeInternalLogicMiddlewareLicense) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *UnibeeInternalLogicMiddlewareLicense) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeInternalLogicMiddlewareLicense) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeInternalLogicMiddlewareLicense) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeInternalLogicMiddlewareLicense) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetSubscription

`func (o *UnibeeInternalLogicMiddlewareLicense) GetSubscription() UnibeeApiBeanSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetSubscriptionOk() (*UnibeeApiBeanSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnibeeInternalLogicMiddlewareLicense) SetSubscription(v UnibeeApiBeanSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UnibeeInternalLogicMiddlewareLicense) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetVersion

`func (o *UnibeeInternalLogicMiddlewareLicense) GetVersion() UnibeeInternalLogicMiddlewareMerchantVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UnibeeInternalLogicMiddlewareLicense) GetVersionOk() (*UnibeeInternalLogicMiddlewareMerchantVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UnibeeInternalLogicMiddlewareLicense) SetVersion(v UnibeeInternalLogicMiddlewareMerchantVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UnibeeInternalLogicMiddlewareLicense) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


