/*
OpenAPI UniBee

UniBee Api Server

API version: daily,buildtime:202502180830
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package unibee

import (
	"encoding/json"
)

// checks if the MerchantMetricUserMetricGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantMetricUserMetricGet200Response{}

// MerchantMetricUserMetricGet200Response struct for MerchantMetricUserMetricGet200Response
type MerchantMetricUserMetricGet200Response struct {
	Code *int32 `json:"code,omitempty"`
	Data *MerchantMetricUserMetricGet200ResponseData `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	Redirect *string `json:"redirect,omitempty"`
	RequestId *string `json:"requestId,omitempty"`
}

// NewMerchantMetricUserMetricGet200Response instantiates a new MerchantMetricUserMetricGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantMetricUserMetricGet200Response() *MerchantMetricUserMetricGet200Response {
	this := MerchantMetricUserMetricGet200Response{}
	return &this
}

// NewMerchantMetricUserMetricGet200ResponseWithDefaults instantiates a new MerchantMetricUserMetricGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantMetricUserMetricGet200ResponseWithDefaults() *MerchantMetricUserMetricGet200Response {
	this := MerchantMetricUserMetricGet200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200Response) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200Response) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *MerchantMetricUserMetricGet200Response) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200Response) GetData() MerchantMetricUserMetricGet200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret MerchantMetricUserMetricGet200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200Response) GetDataOk() (*MerchantMetricUserMetricGet200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MerchantMetricUserMetricGet200ResponseData and assigns it to the Data field.
func (o *MerchantMetricUserMetricGet200Response) SetData(v MerchantMetricUserMetricGet200ResponseData) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MerchantMetricUserMetricGet200Response) SetMessage(v string) {
	o.Message = &v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200Response) GetRedirect() string {
	if o == nil || IsNil(o.Redirect) {
		var ret string
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200Response) GetRedirectOk() (*string, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200Response) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given string and assigns it to the Redirect field.
func (o *MerchantMetricUserMetricGet200Response) SetRedirect(v string) {
	o.Redirect = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *MerchantMetricUserMetricGet200Response) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantMetricUserMetricGet200Response) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *MerchantMetricUserMetricGet200Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *MerchantMetricUserMetricGet200Response) SetRequestId(v string) {
	o.RequestId = &v
}

func (o MerchantMetricUserMetricGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantMetricUserMetricGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Redirect) {
		toSerialize["redirect"] = o.Redirect
	}
	if !IsNil(o.RequestId) {
		toSerialize["requestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableMerchantMetricUserMetricGet200Response struct {
	value *MerchantMetricUserMetricGet200Response
	isSet bool
}

func (v NullableMerchantMetricUserMetricGet200Response) Get() *MerchantMetricUserMetricGet200Response {
	return v.value
}

func (v *NullableMerchantMetricUserMetricGet200Response) Set(val *MerchantMetricUserMetricGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantMetricUserMetricGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantMetricUserMetricGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantMetricUserMetricGet200Response(val *MerchantMetricUserMetricGet200Response) *NullableMerchantMetricUserMetricGet200Response {
	return &NullableMerchantMetricUserMetricGet200Response{value: val, isSet: true}
}

func (v NullableMerchantMetricUserMetricGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantMetricUserMetricGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


