/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
)

// checks if the ResponsesSignedBlocks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesSignedBlocks{}

// ResponsesSignedBlocks struct for ResponsesSignedBlocks
type ResponsesSignedBlocks struct {
	Height *int32 `json:"height,omitempty"`
	Signed *bool `json:"signed,omitempty"`
}

// NewResponsesSignedBlocks instantiates a new ResponsesSignedBlocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesSignedBlocks() *ResponsesSignedBlocks {
	this := ResponsesSignedBlocks{}
	return &this
}

// NewResponsesSignedBlocksWithDefaults instantiates a new ResponsesSignedBlocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesSignedBlocksWithDefaults() *ResponsesSignedBlocks {
	this := ResponsesSignedBlocks{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesSignedBlocks) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesSignedBlocks) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesSignedBlocks) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ResponsesSignedBlocks) SetHeight(v int32) {
	o.Height = &v
}

// GetSigned returns the Signed field value if set, zero value otherwise.
func (o *ResponsesSignedBlocks) GetSigned() bool {
	if o == nil || IsNil(o.Signed) {
		var ret bool
		return ret
	}
	return *o.Signed
}

// GetSignedOk returns a tuple with the Signed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesSignedBlocks) GetSignedOk() (*bool, bool) {
	if o == nil || IsNil(o.Signed) {
		return nil, false
	}
	return o.Signed, true
}

// HasSigned returns a boolean if a field has been set.
func (o *ResponsesSignedBlocks) HasSigned() bool {
	if o != nil && !IsNil(o.Signed) {
		return true
	}

	return false
}

// SetSigned gets a reference to the given bool and assigns it to the Signed field.
func (o *ResponsesSignedBlocks) SetSigned(v bool) {
	o.Signed = &v
}

func (o ResponsesSignedBlocks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesSignedBlocks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Signed) {
		toSerialize["signed"] = o.Signed
	}
	return toSerialize, nil
}

type NullableResponsesSignedBlocks struct {
	value *ResponsesSignedBlocks
	isSet bool
}

func (v NullableResponsesSignedBlocks) Get() *ResponsesSignedBlocks {
	return v.value
}

func (v *NullableResponsesSignedBlocks) Set(val *ResponsesSignedBlocks) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesSignedBlocks) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesSignedBlocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesSignedBlocks(val *ResponsesSignedBlocks) *NullableResponsesSignedBlocks {
	return &NullableResponsesSignedBlocks{value: val, isSet: true}
}

func (v NullableResponsesSignedBlocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesSignedBlocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


