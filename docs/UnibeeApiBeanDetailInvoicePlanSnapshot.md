# UnibeeApiBeanDetailInvoicePlanSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addons associated with the current plan. | [optional] 
**AutoCharge** | Pointer to **bool** | Billing charge | [optional] 
**ChargeType** | Pointer to **int32** | Billing charge type. 0: One-time, 1: New Subscription, 2: Upgrade, 3: Downgrade, 4: Renewal, 5: Billing Cycle Charge. | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 
**PreviousAddons** | Pointer to [**[]UnibeeApiBeanPlanAddonDetail**](UnibeeApiBeanPlanAddonDetail.md) | Addons from the previous plan, relevant for upgrade or downgrade (paidType &#x3D; 2 or 3). | [optional] 
**PreviousPlan** | Pointer to [**UnibeeApiBeanPlan**](UnibeeApiBeanPlan.md) |  | [optional] 

## Methods

### NewUnibeeApiBeanDetailInvoicePlanSnapshot

`func NewUnibeeApiBeanDetailInvoicePlanSnapshot() *UnibeeApiBeanDetailInvoicePlanSnapshot`

NewUnibeeApiBeanDetailInvoicePlanSnapshot instantiates a new UnibeeApiBeanDetailInvoicePlanSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailInvoicePlanSnapshotWithDefaults

`func NewUnibeeApiBeanDetailInvoicePlanSnapshotWithDefaults() *UnibeeApiBeanDetailInvoicePlanSnapshot`

NewUnibeeApiBeanDetailInvoicePlanSnapshotWithDefaults instantiates a new UnibeeApiBeanDetailInvoicePlanSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetAddons() []UnibeeApiBeanPlanAddonDetail`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAutoCharge

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetAutoCharge() bool`

GetAutoCharge returns the AutoCharge field if non-nil, zero value otherwise.

### GetAutoChargeOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetAutoChargeOk() (*bool, bool)`

GetAutoChargeOk returns a tuple with the AutoCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCharge

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetAutoCharge(v bool)`

SetAutoCharge sets AutoCharge field to given value.

### HasAutoCharge

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasAutoCharge() bool`

HasAutoCharge returns a boolean if a field has been set.

### GetChargeType

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetChargeType() int32`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetChargeTypeOk() (*int32, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetChargeType(v int32)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPlan() UnibeeApiBeanPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetPlan(v UnibeeApiBeanPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPreviousAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPreviousAddons() []UnibeeApiBeanPlanAddonDetail`

GetPreviousAddons returns the PreviousAddons field if non-nil, zero value otherwise.

### GetPreviousAddonsOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPreviousAddonsOk() (*[]UnibeeApiBeanPlanAddonDetail, bool)`

GetPreviousAddonsOk returns a tuple with the PreviousAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetPreviousAddons(v []UnibeeApiBeanPlanAddonDetail)`

SetPreviousAddons sets PreviousAddons field to given value.

### HasPreviousAddons

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasPreviousAddons() bool`

HasPreviousAddons returns a boolean if a field has been set.

### GetPreviousPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPreviousPlan() UnibeeApiBeanPlan`

GetPreviousPlan returns the PreviousPlan field if non-nil, zero value otherwise.

### GetPreviousPlanOk

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) GetPreviousPlanOk() (*UnibeeApiBeanPlan, bool)`

GetPreviousPlanOk returns a tuple with the PreviousPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) SetPreviousPlan(v UnibeeApiBeanPlan)`

SetPreviousPlan sets PreviousPlan field to given value.

### HasPreviousPlan

`func (o *UnibeeApiBeanDetailInvoicePlanSnapshot) HasPreviousPlan() bool`

HasPreviousPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


