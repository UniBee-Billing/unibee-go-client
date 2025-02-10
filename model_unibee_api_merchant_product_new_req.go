/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502100408
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the UnibeeApiMerchantProductNewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnibeeApiMerchantProductNewReq{}

// UnibeeApiMerchantProductNewReq struct for UnibeeApiMerchantProductNewReq
type UnibeeApiMerchantProductNewReq struct {
	// description
	Description *string `json:"description,omitempty"`
	// home_url
	HomeUrl *string `json:"homeUrl,omitempty"`
	// image_url
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Metadata，Map
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// ProductName
	ProductName *string `json:"productName,omitempty"`
	// status，1-active，2-inactive, default active
	Status *int32 `json:"status,omitempty"`
}

// NewUnibeeApiMerchantProductNewReq instantiates a new UnibeeApiMerchantProductNewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnibeeApiMerchantProductNewReq() *UnibeeApiMerchantProductNewReq {
	this := UnibeeApiMerchantProductNewReq{}
	return &this
}

// NewUnibeeApiMerchantProductNewReqWithDefaults instantiates a new UnibeeApiMerchantProductNewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnibeeApiMerchantProductNewReqWithDefaults() *UnibeeApiMerchantProductNewReq {
	this := UnibeeApiMerchantProductNewReq{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnibeeApiMerchantProductNewReq) SetDescription(v string) {
	o.Description = &v
}

// GetHomeUrl returns the HomeUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetHomeUrl() string {
	if o == nil || IsNil(o.HomeUrl) {
		var ret string
		return ret
	}
	return *o.HomeUrl
}

// GetHomeUrlOk returns a tuple with the HomeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetHomeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HomeUrl) {
		return nil, false
	}
	return o.HomeUrl, true
}

// HasHomeUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasHomeUrl() bool {
	if o != nil && !IsNil(o.HomeUrl) {
		return true
	}

	return false
}

// SetHomeUrl gets a reference to the given string and assigns it to the HomeUrl field.
func (o *UnibeeApiMerchantProductNewReq) SetHomeUrl(v string) {
	o.HomeUrl = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *UnibeeApiMerchantProductNewReq) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *UnibeeApiMerchantProductNewReq) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *UnibeeApiMerchantProductNewReq) SetProductName(v string) {
	o.ProductName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UnibeeApiMerchantProductNewReq) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnibeeApiMerchantProductNewReq) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UnibeeApiMerchantProductNewReq) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *UnibeeApiMerchantProductNewReq) SetStatus(v int32) {
	o.Status = &v
}

func (o UnibeeApiMerchantProductNewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnibeeApiMerchantProductNewReq) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUnibeeApiMerchantProductNewReq struct {
	value *UnibeeApiMerchantProductNewReq
	isSet bool
}

func (v NullableUnibeeApiMerchantProductNewReq) Get() *UnibeeApiMerchantProductNewReq {
	return v.value
}

func (v *NullableUnibeeApiMerchantProductNewReq) Set(val *UnibeeApiMerchantProductNewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableUnibeeApiMerchantProductNewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableUnibeeApiMerchantProductNewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnibeeApiMerchantProductNewReq(val *UnibeeApiMerchantProductNewReq) *NullableUnibeeApiMerchantProductNewReq {
	return &NullableUnibeeApiMerchantProductNewReq{value: val, isSet: true}
}

func (v NullableUnibeeApiMerchantProductNewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnibeeApiMerchantProductNewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


