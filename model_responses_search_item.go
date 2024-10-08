/*
Celenium API

Celenium API is a powerful tool to access all blockchain data that is processed and indexed by our proprietary indexer. With Celenium API you can retrieve all historical data, off-chain data, blobs and statistics through our REST API. Celenium API indexer are open source, which allows you to not depend on third-party services. You can clone, build and run them independently, giving you full control over all components. If you have any questions or feature requests, please feel free to contact us. We appreciate your feedback!

API version: 1.0
Contact: celenium@pklabs.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
)

// checks if the ResponsesSearchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesSearchItem{}

// ResponsesSearchItem struct for ResponsesSearchItem
type ResponsesSearchItem struct {
	// Search result. Can be one of folowwing types: Block, Address, Namespace, Tx, Validator, Rollup
	Result map[string]interface{} `json:"result,omitempty"`
	// Result type which is in the result. Can be 'block', 'address', 'namespace', 'tx', 'validator', 'rollup'
	Type *string `json:"type,omitempty"`
}

// NewResponsesSearchItem instantiates a new ResponsesSearchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesSearchItem() *ResponsesSearchItem {
	this := ResponsesSearchItem{}
	return &this
}

// NewResponsesSearchItemWithDefaults instantiates a new ResponsesSearchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesSearchItemWithDefaults() *ResponsesSearchItem {
	this := ResponsesSearchItem{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ResponsesSearchItem) GetResult() map[string]interface{} {
	if o == nil || IsNil(o.Result) {
		var ret map[string]interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesSearchItem) GetResultOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return map[string]interface{}{}, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ResponsesSearchItem) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *ResponsesSearchItem) SetResult(v map[string]interface{}) {
	o.Result = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponsesSearchItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesSearchItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponsesSearchItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResponsesSearchItem) SetType(v string) {
	o.Type = &v
}

func (o ResponsesSearchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesSearchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableResponsesSearchItem struct {
	value *ResponsesSearchItem
	isSet bool
}

func (v NullableResponsesSearchItem) Get() *ResponsesSearchItem {
	return v.value
}

func (v *NullableResponsesSearchItem) Set(val *ResponsesSearchItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesSearchItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesSearchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesSearchItem(val *ResponsesSearchItem) *NullableResponsesSearchItem {
	return &NullableResponsesSearchItem{value: val, isSet: true}
}

func (v NullableResponsesSearchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesSearchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


