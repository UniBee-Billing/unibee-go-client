# UnibeeInternalModelEntityOverseaPaySubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddonData** | Pointer to **string** | plan addon json data | [optional] 
**Amount** | Pointer to **int64** | amount, cent | [optional] 
**BillingCycleAnchor** | Pointer to **int64** | billing_cycle_anchor | [optional] 
**CancelAtPeriodEnd** | Pointer to **int32** | whether cancel at period end，0-false | 1-true | [optional] 
**CancelReason** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** | create utc time | [optional] 
**Currency** | Pointer to **string** | currency | [optional] 
**CurrentPeriodEnd** | Pointer to **int64** | current_period_end, utc time | [optional] 
**CurrentPeriodEndTime** | Pointer to **string** |  | [optional] 
**CurrentPeriodStart** | Pointer to **int64** | current_period_start, utc time | [optional] 
**CurrentPeriodStartTime** | Pointer to **string** |  | [optional] 
**CustomerEmail** | Pointer to **string** | customer_email | [optional] 
**CustomerName** | Pointer to **string** | customer_name | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**DunningTime** | Pointer to **int64** | dunning_time, utc time | [optional] 
**FirstPaidTime** | Pointer to **int64** | first success payment time | [optional] 
**GatewayDefaultPaymentMethod** | Pointer to **string** |  | [optional] 
**GatewayId** | Pointer to **int64** | gateway_id | [optional] 
**GatewayItemData** | Pointer to **string** | gateway_item_data | [optional] 
**GatewayLatestInvoiceId** | Pointer to **string** | gateway latest invoice id | [optional] 
**GatewayStatus** | Pointer to **string** | gateway status，Stripe：https://stripe.com/docs/billing/subscriptions/webhooks  Paypal：https://developer.paypal.com/docs/api/subscriptions/v1/#subscriptions_get | [optional] 
**GatewaySubscriptionId** | Pointer to **string** | gateway subscription id | [optional] 
**GmtCreate** | Pointer to **string** | create time | [optional] 
**GmtModify** | Pointer to **string** | update time | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsDeleted** | Pointer to **int32** | 0-UnDeleted，1-Deleted | [optional] 
**LastUpdateTime** | Pointer to **int64** |  | [optional] 
**LatestInvoiceId** | Pointer to **string** | latest_invoice_id | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**MerchantId** | Pointer to **int64** | merchant id | [optional] 
**MetaData** | Pointer to **string** | meta_data(json) | [optional] 
**PendingUpdateId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **int64** | plan id | [optional] 
**Quantity** | Pointer to **int64** | quantity | [optional] 
**ResponseData** | Pointer to **string** |  | [optional] 
**ReturnUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | status，0-Init | 1-Create｜2-Active｜3-PendingInActive | 4-Cancel | 5-Expire | 6- Suspend| 7-Incomplete | [optional] 
**SubscriptionId** | Pointer to **string** | subscription id | [optional] 
**TaskTime** | Pointer to **string** | task_time | [optional] 
**TaxScale** | Pointer to **int64** | Tax Scale，1000 &#x3D; 10% | [optional] 
**TestClock** | Pointer to **int64** | test_clock, simulator clock for subscription, if set , sub will out of cronjob controll | [optional] 
**TrialEnd** | Pointer to **int64** | trial_end, utc time | [optional] 
**Type** | Pointer to **int32** | sub type, 0-gateway sub, 1-unibee sub | [optional] 
**UserId** | Pointer to **int64** | userId | [optional] 
**VatNumber** | Pointer to **string** |  | [optional] 
**VatVerifyData** | Pointer to **string** |  | [optional] 

## Methods

### NewUnibeeInternalModelEntityOverseaPaySubscription

`func NewUnibeeInternalModelEntityOverseaPaySubscription() *UnibeeInternalModelEntityOverseaPaySubscription`

NewUnibeeInternalModelEntityOverseaPaySubscription instantiates a new UnibeeInternalModelEntityOverseaPaySubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnibeeInternalModelEntityOverseaPaySubscriptionWithDefaults

`func NewUnibeeInternalModelEntityOverseaPaySubscriptionWithDefaults() *UnibeeInternalModelEntityOverseaPaySubscription`

NewUnibeeInternalModelEntityOverseaPaySubscriptionWithDefaults instantiates a new UnibeeInternalModelEntityOverseaPaySubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetAddonData() string`

GetAddonData returns the AddonData field if non-nil, zero value otherwise.

### GetAddonDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetAddonDataOk() (*string, bool)`

GetAddonDataOk returns a tuple with the AddonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetAddonData(v string)`

SetAddonData sets AddonData field to given value.

### HasAddonData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasAddonData() bool`

HasAddonData returns a boolean if a field has been set.

### GetAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingCycleAnchor

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetBillingCycleAnchor() int64`

GetBillingCycleAnchor returns the BillingCycleAnchor field if non-nil, zero value otherwise.

### GetBillingCycleAnchorOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetBillingCycleAnchorOk() (*int64, bool)`

GetBillingCycleAnchorOk returns a tuple with the BillingCycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycleAnchor

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetBillingCycleAnchor(v int64)`

SetBillingCycleAnchor sets BillingCycleAnchor field to given value.

### HasBillingCycleAnchor

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasBillingCycleAnchor() bool`

HasBillingCycleAnchor returns a boolean if a field has been set.

### GetCancelAtPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCancelAtPeriodEnd() int32`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCancelAtPeriodEndOk() (*int32, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCancelAtPeriodEnd(v int32)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.

### HasCancelAtPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCancelAtPeriodEnd() bool`

HasCancelAtPeriodEnd returns a boolean if a field has been set.

### GetCancelReason

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.

### GetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCurrentPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodEnd() int64`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodEndOk() (*int64, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCurrentPeriodEnd(v int64)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.

### HasCurrentPeriodEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCurrentPeriodEnd() bool`

HasCurrentPeriodEnd returns a boolean if a field has been set.

### GetCurrentPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodEndTime() string`

GetCurrentPeriodEndTime returns the CurrentPeriodEndTime field if non-nil, zero value otherwise.

### GetCurrentPeriodEndTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodEndTimeOk() (*string, bool)`

GetCurrentPeriodEndTimeOk returns a tuple with the CurrentPeriodEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCurrentPeriodEndTime(v string)`

SetCurrentPeriodEndTime sets CurrentPeriodEndTime field to given value.

### HasCurrentPeriodEndTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCurrentPeriodEndTime() bool`

HasCurrentPeriodEndTime returns a boolean if a field has been set.

### GetCurrentPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodStart() int64`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodStartOk() (*int64, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCurrentPeriodStart(v int64)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.

### HasCurrentPeriodStart

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCurrentPeriodStart() bool`

HasCurrentPeriodStart returns a boolean if a field has been set.

### GetCurrentPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodStartTime() string`

GetCurrentPeriodStartTime returns the CurrentPeriodStartTime field if non-nil, zero value otherwise.

### GetCurrentPeriodStartTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCurrentPeriodStartTimeOk() (*string, bool)`

GetCurrentPeriodStartTimeOk returns a tuple with the CurrentPeriodStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCurrentPeriodStartTime(v string)`

SetCurrentPeriodStartTime sets CurrentPeriodStartTime field to given value.

### HasCurrentPeriodStartTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCurrentPeriodStartTime() bool`

HasCurrentPeriodStartTime returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerName

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDunningTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetDunningTime() int64`

GetDunningTime returns the DunningTime field if non-nil, zero value otherwise.

### GetDunningTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetDunningTimeOk() (*int64, bool)`

GetDunningTimeOk returns a tuple with the DunningTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDunningTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetDunningTime(v int64)`

SetDunningTime sets DunningTime field to given value.

### HasDunningTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasDunningTime() bool`

HasDunningTime returns a boolean if a field has been set.

### GetFirstPaidTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetFirstPaidTime() int64`

GetFirstPaidTime returns the FirstPaidTime field if non-nil, zero value otherwise.

### GetFirstPaidTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetFirstPaidTimeOk() (*int64, bool)`

GetFirstPaidTimeOk returns a tuple with the FirstPaidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPaidTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetFirstPaidTime(v int64)`

SetFirstPaidTime sets FirstPaidTime field to given value.

### HasFirstPaidTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasFirstPaidTime() bool`

HasFirstPaidTime returns a boolean if a field has been set.

### GetGatewayDefaultPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayDefaultPaymentMethod() string`

GetGatewayDefaultPaymentMethod returns the GatewayDefaultPaymentMethod field if non-nil, zero value otherwise.

### GetGatewayDefaultPaymentMethodOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayDefaultPaymentMethodOk() (*string, bool)`

GetGatewayDefaultPaymentMethodOk returns a tuple with the GatewayDefaultPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDefaultPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewayDefaultPaymentMethod(v string)`

SetGatewayDefaultPaymentMethod sets GatewayDefaultPaymentMethod field to given value.

### HasGatewayDefaultPaymentMethod

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewayDefaultPaymentMethod() bool`

HasGatewayDefaultPaymentMethod returns a boolean if a field has been set.

### GetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayId() int64`

GetGatewayId returns the GatewayId field if non-nil, zero value otherwise.

### GetGatewayIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayIdOk() (*int64, bool)`

GetGatewayIdOk returns a tuple with the GatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewayId(v int64)`

SetGatewayId sets GatewayId field to given value.

### HasGatewayId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewayId() bool`

HasGatewayId returns a boolean if a field has been set.

### GetGatewayItemData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayItemData() string`

GetGatewayItemData returns the GatewayItemData field if non-nil, zero value otherwise.

### GetGatewayItemDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayItemDataOk() (*string, bool)`

GetGatewayItemDataOk returns a tuple with the GatewayItemData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayItemData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewayItemData(v string)`

SetGatewayItemData sets GatewayItemData field to given value.

### HasGatewayItemData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewayItemData() bool`

HasGatewayItemData returns a boolean if a field has been set.

### GetGatewayLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayLatestInvoiceId() string`

GetGatewayLatestInvoiceId returns the GatewayLatestInvoiceId field if non-nil, zero value otherwise.

### GetGatewayLatestInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayLatestInvoiceIdOk() (*string, bool)`

GetGatewayLatestInvoiceIdOk returns a tuple with the GatewayLatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewayLatestInvoiceId(v string)`

SetGatewayLatestInvoiceId sets GatewayLatestInvoiceId field to given value.

### HasGatewayLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewayLatestInvoiceId() bool`

HasGatewayLatestInvoiceId returns a boolean if a field has been set.

### GetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayStatus() string`

GetGatewayStatus returns the GatewayStatus field if non-nil, zero value otherwise.

### GetGatewayStatusOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewayStatusOk() (*string, bool)`

GetGatewayStatusOk returns a tuple with the GatewayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewayStatus(v string)`

SetGatewayStatus sets GatewayStatus field to given value.

### HasGatewayStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewayStatus() bool`

HasGatewayStatus returns a boolean if a field has been set.

### GetGatewaySubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewaySubscriptionId() string`

GetGatewaySubscriptionId returns the GatewaySubscriptionId field if non-nil, zero value otherwise.

### GetGatewaySubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGatewaySubscriptionIdOk() (*string, bool)`

GetGatewaySubscriptionIdOk returns a tuple with the GatewaySubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGatewaySubscriptionId(v string)`

SetGatewaySubscriptionId sets GatewaySubscriptionId field to given value.

### HasGatewaySubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGatewaySubscriptionId() bool`

HasGatewaySubscriptionId returns a boolean if a field has been set.

### GetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGmtCreate() string`

GetGmtCreate returns the GmtCreate field if non-nil, zero value otherwise.

### GetGmtCreateOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGmtCreateOk() (*string, bool)`

GetGmtCreateOk returns a tuple with the GmtCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGmtCreate(v string)`

SetGmtCreate sets GmtCreate field to given value.

### HasGmtCreate

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGmtCreate() bool`

HasGmtCreate returns a boolean if a field has been set.

### GetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGmtModify() string`

GetGmtModify returns the GmtModify field if non-nil, zero value otherwise.

### GetGmtModifyOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetGmtModifyOk() (*string, bool)`

GetGmtModifyOk returns a tuple with the GmtModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetGmtModify(v string)`

SetGmtModify sets GmtModify field to given value.

### HasGmtModify

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasGmtModify() bool`

HasGmtModify returns a boolean if a field has been set.

### GetId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetIsDeleted() int32`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetIsDeletedOk() (*int32, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetIsDeleted(v int32)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLastUpdateTime() int64`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLastUpdateTimeOk() (*int64, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetLastUpdateTime(v int64)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLatestInvoiceId() string`

GetLatestInvoiceId returns the LatestInvoiceId field if non-nil, zero value otherwise.

### GetLatestInvoiceIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLatestInvoiceIdOk() (*string, bool)`

GetLatestInvoiceIdOk returns a tuple with the LatestInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetLatestInvoiceId(v string)`

SetLatestInvoiceId sets LatestInvoiceId field to given value.

### HasLatestInvoiceId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasLatestInvoiceId() bool`

HasLatestInvoiceId returns a boolean if a field has been set.

### GetLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetMerchantId() int64`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetMerchantIdOk() (*int64, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetMerchantId(v int64)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetMetaData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetPendingUpdateId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetPendingUpdateId() string`

GetPendingUpdateId returns the PendingUpdateId field if non-nil, zero value otherwise.

### GetPendingUpdateIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetPendingUpdateIdOk() (*string, bool)`

GetPendingUpdateIdOk returns a tuple with the PendingUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdateId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetPendingUpdateId(v string)`

SetPendingUpdateId sets PendingUpdateId field to given value.

### HasPendingUpdateId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasPendingUpdateId() bool`

HasPendingUpdateId returns a boolean if a field has been set.

### GetPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetPlanId() int64`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetPlanIdOk() (*int64, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetPlanId(v int64)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetResponseData() string`

GetResponseData returns the ResponseData field if non-nil, zero value otherwise.

### GetResponseDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetResponseDataOk() (*string, bool)`

GetResponseDataOk returns a tuple with the ResponseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetResponseData(v string)`

SetResponseData sets ResponseData field to given value.

### HasResponseData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasResponseData() bool`

HasResponseData returns a boolean if a field has been set.

### GetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTaskTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTaskTime() string`

GetTaskTime returns the TaskTime field if non-nil, zero value otherwise.

### GetTaskTimeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTaskTimeOk() (*string, bool)`

GetTaskTimeOk returns a tuple with the TaskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetTaskTime(v string)`

SetTaskTime sets TaskTime field to given value.

### HasTaskTime

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasTaskTime() bool`

HasTaskTime returns a boolean if a field has been set.

### GetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTaxScale() int64`

GetTaxScale returns the TaxScale field if non-nil, zero value otherwise.

### GetTaxScaleOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTaxScaleOk() (*int64, bool)`

GetTaxScaleOk returns a tuple with the TaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxScale

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetTaxScale(v int64)`

SetTaxScale sets TaxScale field to given value.

### HasTaxScale

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasTaxScale() bool`

HasTaxScale returns a boolean if a field has been set.

### GetTestClock

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTestClock() int64`

GetTestClock returns the TestClock field if non-nil, zero value otherwise.

### GetTestClockOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTestClockOk() (*int64, bool)`

GetTestClockOk returns a tuple with the TestClock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestClock

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetTestClock(v int64)`

SetTestClock sets TestClock field to given value.

### HasTestClock

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasTestClock() bool`

HasTestClock returns a boolean if a field has been set.

### GetTrialEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTrialEnd() int64`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTrialEndOk() (*int64, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetTrialEnd(v int64)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### GetType

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVatNumber

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetVatNumber() string`

GetVatNumber returns the VatNumber field if non-nil, zero value otherwise.

### GetVatNumberOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetVatNumberOk() (*string, bool)`

GetVatNumberOk returns a tuple with the VatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatNumber

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetVatNumber(v string)`

SetVatNumber sets VatNumber field to given value.

### HasVatNumber

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasVatNumber() bool`

HasVatNumber returns a boolean if a field has been set.

### GetVatVerifyData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetVatVerifyData() string`

GetVatVerifyData returns the VatVerifyData field if non-nil, zero value otherwise.

### GetVatVerifyDataOk

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) GetVatVerifyDataOk() (*string, bool)`

GetVatVerifyDataOk returns a tuple with the VatVerifyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatVerifyData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) SetVatVerifyData(v string)`

SetVatVerifyData sets VatVerifyData field to given value.

### HasVatVerifyData

`func (o *UnibeeInternalModelEntityOverseaPaySubscription) HasVatVerifyData() bool`

HasVatVerifyData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


