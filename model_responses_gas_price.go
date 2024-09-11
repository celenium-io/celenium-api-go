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

// checks if the ResponsesGasPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesGasPrice{}

// ResponsesGasPrice struct for ResponsesGasPrice
type ResponsesGasPrice struct {
	Fast *string `json:"fast,omitempty"`
	Median *string `json:"median,omitempty"`
	Slow *string `json:"slow,omitempty"`
}

// NewResponsesGasPrice instantiates a new ResponsesGasPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesGasPrice() *ResponsesGasPrice {
	this := ResponsesGasPrice{}
	return &this
}

// NewResponsesGasPriceWithDefaults instantiates a new ResponsesGasPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesGasPriceWithDefaults() *ResponsesGasPrice {
	this := ResponsesGasPrice{}
	return &this
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *ResponsesGasPrice) GetFast() string {
	if o == nil || IsNil(o.Fast) {
		var ret string
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesGasPrice) GetFastOk() (*string, bool) {
	if o == nil || IsNil(o.Fast) {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *ResponsesGasPrice) HasFast() bool {
	if o != nil && !IsNil(o.Fast) {
		return true
	}

	return false
}

// SetFast gets a reference to the given string and assigns it to the Fast field.
func (o *ResponsesGasPrice) SetFast(v string) {
	o.Fast = &v
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *ResponsesGasPrice) GetMedian() string {
	if o == nil || IsNil(o.Median) {
		var ret string
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesGasPrice) GetMedianOk() (*string, bool) {
	if o == nil || IsNil(o.Median) {
		return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *ResponsesGasPrice) HasMedian() bool {
	if o != nil && !IsNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given string and assigns it to the Median field.
func (o *ResponsesGasPrice) SetMedian(v string) {
	o.Median = &v
}

// GetSlow returns the Slow field value if set, zero value otherwise.
func (o *ResponsesGasPrice) GetSlow() string {
	if o == nil || IsNil(o.Slow) {
		var ret string
		return ret
	}
	return *o.Slow
}

// GetSlowOk returns a tuple with the Slow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesGasPrice) GetSlowOk() (*string, bool) {
	if o == nil || IsNil(o.Slow) {
		return nil, false
	}
	return o.Slow, true
}

// HasSlow returns a boolean if a field has been set.
func (o *ResponsesGasPrice) HasSlow() bool {
	if o != nil && !IsNil(o.Slow) {
		return true
	}

	return false
}

// SetSlow gets a reference to the given string and assigns it to the Slow field.
func (o *ResponsesGasPrice) SetSlow(v string) {
	o.Slow = &v
}

func (o ResponsesGasPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesGasPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fast) {
		toSerialize["fast"] = o.Fast
	}
	if !IsNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	if !IsNil(o.Slow) {
		toSerialize["slow"] = o.Slow
	}
	return toSerialize, nil
}

type NullableResponsesGasPrice struct {
	value *ResponsesGasPrice
	isSet bool
}

func (v NullableResponsesGasPrice) Get() *ResponsesGasPrice {
	return v.value
}

func (v *NullableResponsesGasPrice) Set(val *ResponsesGasPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesGasPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesGasPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesGasPrice(val *ResponsesGasPrice) *NullableResponsesGasPrice {
	return &NullableResponsesGasPrice{value: val, isSet: true}
}

func (v NullableResponsesGasPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesGasPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


