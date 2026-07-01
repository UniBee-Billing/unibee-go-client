# UnibeeApiBeanDetailPublicPlanDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonIds** | Pointer to **[]int64** | AddonIds | [optional] 
**Addons** | Pointer to [**[]UnibeeApiBeanPublicPlan**](UnibeeApiBeanPublicPlan.md) | Addons | [optional] 
**CheckAddressViaGateway** | Pointer to **bool** |  | [optional] 
**GlobalUSVATActive** | Pointer to **bool** |  | [optional] 
**OnetimeAddonIds** | Pointer to **[]int64** | OneTimeAddonIds | [optional] 
**OnetimeAddons** | Pointer to [**[]UnibeeApiBeanPublicPlan**](UnibeeApiBeanPublicPlan.md) | OneTimeAddons | [optional] 
**Plan** | Pointer to [**UnibeeApiBeanPublicPlan**](UnibeeApiBeanPublicPlan.md) |  | [optional] 
**Product** | Pointer to [**UnibeeApiBeanPublicProduct**](UnibeeApiBeanPublicProduct.md) |  | [optional] 
**SellOnUSOnly** | Pointer to **bool** | SellOnUSOnly, customer must select US Country And provider US Address when its true | [optional] 

## Methods

### NewUnibeeApiBeanDetailPublicPlanDetail

`func NewUnibeeApiBeanDetailPublicPlanDetail() *UnibeeApiBeanDetailPublicPlanDetail`

NewUnibeeApiBeanDetailPublicPlanDetail instantiates a new UnibeeApiBeanDetailPublicPlanDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeApiBeanDetailPublicPlanDetailWithDefaults

`func NewUnibeeApiBeanDetailPublicPlanDetailWithDefaults() *UnibeeApiBeanDetailPublicPlanDetail`

NewUnibeeApiBeanDetailPublicPlanDetailWithDefaults instantiates a new UnibeeApiBeanDetailPublicPlanDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetAddonIds() []int64`

GetAddonIds returns the AddonIds field if non-nil, zero value otherwise.

### GetAddonIdsOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetAddonIdsOk() (*[]int64, bool)`

GetAddonIdsOk returns a tuple with the AddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetAddonIds(v []int64)`

SetAddonIds sets AddonIds field to given value.

### HasAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasAddonIds() bool`

HasAddonIds returns a boolean if a field has been set.

### GetAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetAddons() []UnibeeApiBeanPublicPlan`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetAddonsOk() (*[]UnibeeApiBeanPublicPlan, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetAddons(v []UnibeeApiBeanPublicPlan)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetCheckAddressViaGateway

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetCheckAddressViaGateway() bool`

GetCheckAddressViaGateway returns the CheckAddressViaGateway field if non-nil, zero value otherwise.

### GetCheckAddressViaGatewayOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetCheckAddressViaGatewayOk() (*bool, bool)`

GetCheckAddressViaGatewayOk returns a tuple with the CheckAddressViaGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAddressViaGateway

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetCheckAddressViaGateway(v bool)`

SetCheckAddressViaGateway sets CheckAddressViaGateway field to given value.

### HasCheckAddressViaGateway

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasCheckAddressViaGateway() bool`

HasCheckAddressViaGateway returns a boolean if a field has been set.

### GetGlobalUSVATActive

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetGlobalUSVATActive() bool`

GetGlobalUSVATActive returns the GlobalUSVATActive field if non-nil, zero value otherwise.

### GetGlobalUSVATActiveOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetGlobalUSVATActiveOk() (*bool, bool)`

GetGlobalUSVATActiveOk returns a tuple with the GlobalUSVATActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalUSVATActive

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetGlobalUSVATActive(v bool)`

SetGlobalUSVATActive sets GlobalUSVATActive field to given value.

### HasGlobalUSVATActive

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasGlobalUSVATActive() bool`

HasGlobalUSVATActive returns a boolean if a field has been set.

### GetOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetOnetimeAddonIds() []int64`

GetOnetimeAddonIds returns the OnetimeAddonIds field if non-nil, zero value otherwise.

### GetOnetimeAddonIdsOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetOnetimeAddonIdsOk() (*[]int64, bool)`

GetOnetimeAddonIdsOk returns a tuple with the OnetimeAddonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetOnetimeAddonIds(v []int64)`

SetOnetimeAddonIds sets OnetimeAddonIds field to given value.

### HasOnetimeAddonIds

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasOnetimeAddonIds() bool`

HasOnetimeAddonIds returns a boolean if a field has been set.

### GetOnetimeAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetOnetimeAddons() []UnibeeApiBeanPublicPlan`

GetOnetimeAddons returns the OnetimeAddons field if non-nil, zero value otherwise.

### GetOnetimeAddonsOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetOnetimeAddonsOk() (*[]UnibeeApiBeanPublicPlan, bool)`

GetOnetimeAddonsOk returns a tuple with the OnetimeAddons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnetimeAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetOnetimeAddons(v []UnibeeApiBeanPublicPlan)`

SetOnetimeAddons sets OnetimeAddons field to given value.

### HasOnetimeAddons

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasOnetimeAddons() bool`

HasOnetimeAddons returns a boolean if a field has been set.

### GetPlan

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetPlan() UnibeeApiBeanPublicPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetPlanOk() (*UnibeeApiBeanPublicPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetPlan(v UnibeeApiBeanPublicPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProduct

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetProduct() UnibeeApiBeanPublicProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetProductOk() (*UnibeeApiBeanPublicProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetProduct(v UnibeeApiBeanPublicProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSellOnUSOnly

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetSellOnUSOnly() bool`

GetSellOnUSOnly returns the SellOnUSOnly field if non-nil, zero value otherwise.

### GetSellOnUSOnlyOk

`func (o *UnibeeApiBeanDetailPublicPlanDetail) GetSellOnUSOnlyOk() (*bool, bool)`

GetSellOnUSOnlyOk returns a tuple with the SellOnUSOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellOnUSOnly

`func (o *UnibeeApiBeanDetailPublicPlanDetail) SetSellOnUSOnly(v bool)`

SetSellOnUSOnly sets SellOnUSOnly field to given value.

### HasSellOnUSOnly

`func (o *UnibeeApiBeanDetailPublicPlanDetail) HasSellOnUSOnly() bool`

HasSellOnUSOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


