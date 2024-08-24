/*
OpenAPI UniBee

This is UniBee Api Server

API version: buildtime:202408240519 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiBeanProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiBeanProduct{}

// UnibeeApiBeanProduct struct for UnibeeApiBeanProduct
type UnibeeApiBeanProduct struct {
	// create utc time
	CreateTime *int64 `json:"createTime,omitempty"`
	// description
	Description *string `json:"description,omitempty"`
	// home_url
	HomeUrl *string `json:"homeUrl,omitempty"`
	Id *int64 `json:"id,omitempty"`
	// image_url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// 0-UnDeleted，1-Deleted
	IsDeleted *int32 `json:"isDeleted,omitempty"`
	// merchant id
	MerchantId *int64 `json:"merchantId,omitempty"`
	// meta_data(json)
	MetaData *string `json:"metaData,omitempty"`
	// PlanName
	ProductName *string `json:"productName,omitempty"`
	// status，1-active，2-inactive, default active
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiBeanProduct instantiates a new UnibeeApiBeanProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiBeanProduct() *UnibeeApiBeanProduct {
	this := UnibeeApiBeanProduct{}
	return &this
}

// NewUnibeeApiBeanProductWithDefaults instantiates a new UnibeeApiBeanProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiBeanProductWithDefaults() *UnibeeApiBeanProduct {
	this := UnibeeApiBeanProduct{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetCreateTime() int64 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetCreateTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *UnibeeApiBeanProduct) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiBeanProduct) SetDescription(v string) {
	o.Description = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiBeanProduct) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UnibeeApiBeanProduct) SetId(v int64) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiBeanProduct) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetIsDeleted() int32 {
	if o == nil || IsNil(o.IsDeleted) {
		var ret int32
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetIsDeletedOk() (*int32, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given int32 and assigns it to the IsDeleted field.
func (o *UnibeeApiBeanProduct) SetIsDeleted(v int32) {
	o.IsDeleted = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetMerchantId() int64 {
	if o == nil || IsNil(o.MerchantId) {
		var ret int64
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetMerchantIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given int64 and assigns it to the MerchantId field.
func (o *UnibeeApiBeanProduct) SetMerchantId(v int64) {
	o.MerchantId = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetMetaData() string {
	if o == nil || IsNil(o.MetaData) {
		var ret string
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetMetaDataOk() (*string, bool) {
	if o == nil || IsNil(o.MetaData) {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasMetaData() bool {
	if o != nil && !IsNil(o.MetaData) {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given string and assigns it to the MetaData field.
func (o *UnibeeApiBeanProduct) SetMetaData(v string) {
	o.MetaData = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiBeanProduct) SetProductName(v string) {
	o.ProductName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiBeanProduct) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiBeanProduct) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiBeanProduct) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiBeanProduct) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiBeanProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiBeanProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	if !IsNil(o.MetaData) {
		toSerialize["metaData"] = o.MetaData
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiBeanProduct struct {
	value *UnibeeApiBeanProduct
	isSet bool
}

func (v NullableUnibeeApiBeanProduct) Get() *UnibeeApiBeanProduct {
	return v.value
}

func (v *NullableUnibeeApiBeanProduct) Set(val *UnibeeApiBeanProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiBeanProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiBeanProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiBeanProduct(val *UnibeeApiBeanProduct) *NullableUnibeeApiBeanProduct {
	return &NullableUnibeeApiBeanProduct{value: val, isSet: true}
}

func (v NullableUnibeeApiBeanProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiBeanProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


