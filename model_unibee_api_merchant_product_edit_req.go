/*
OpenAPI UniBee

This is UniBee Api Server

API version: develop
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UnibeeApiMerchantProductEditReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductEditReq{}

// UnibeeApiMerchantProductEditReq Edit exist product, product is editable for both active and inactive status 
type UnibeeApiMerchantProductEditReq struct {
	// description
	Description *string `json:"description,omitempty"`
	// home_url
	HomeUrl *string `json:"homeUrl,omitempty"`
	// image_url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// Id of product
	ProductId int64 `json:"productId"`
	// ProductName
	ProductName *string `json:"productName,omitempty"`
	// status，1-active，2-inactive, default active
	Status *int32 `json:"status,omitempty"`
}

type _UnibeeApiMerchantProductEditReq UnibeeApiMerchantProductEditReq

// NewUnibeeApiMerchantProductEditReq instantiates a new UnibeeApiMerchantProductEditReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductEditReq(productId int64) *UnibeeApiMerchantProductEditReq {
	this := UnibeeApiMerchantProductEditReq{}
	this.ProductId = productId
	return &this
}

// NewUnibeeApiMerchantProductEditReqWithDefaults instantiates a new UnibeeApiMerchantProductEditReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductEditReqWithDefaults() *UnibeeApiMerchantProductEditReq {
	this := UnibeeApiMerchantProductEditReq{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantProductEditReq) SetDescription(v string) {
	o.Description = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiMerchantProductEditReq) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiMerchantProductEditReq) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantProductEditReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetProductId returns the ProductId field value
func (o *UnibeeApiMerchantProductEditReq) GetProductId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetProductIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *UnibeeApiMerchantProductEditReq) SetProductId(v int64) {
	o.ProductId = v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiMerchantProductEditReq) SetProductName(v string) {
	o.ProductName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductEditReq) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductEditReq) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductEditReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantProductEditReq) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiMerchantProductEditReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductEditReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HomeUrl) {
		toSerialize["homeUrl"] = o.HomeUrl
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["productId"] = o.ProductId
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *UnibeeApiMerchantProductEditReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUnibeeApiMerchantProductEditReq := _UnibeeApiMerchantProductEditReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnibeeApiMerchantProductEditReq)

	if err != nil {
		return err
	}

	*o = UnibeeApiMerchantProductEditReq(varUnibeeApiMerchantProductEditReq)

	return err
}

type NullableUnibeeApiMerchantProductEditReq struct {
	value *UnibeeApiMerchantProductEditReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductEditReq) Get() *UnibeeApiMerchantProductEditReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductEditReq) Set(val *UnibeeApiMerchantProductEditReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductEditReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductEditReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductEditReq(val *UnibeeApiMerchantProductEditReq) *NullableUnibeeApiMerchantProductEditReq {
	return &NullableUnibeeApiMerchantProductEditReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductEditReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductEditReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


