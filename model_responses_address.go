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

// checks if the ResponsesAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesAddress{}

// ResponsesAddress Celestia address information
type ResponsesAddress struct {
	Balance *ResponsesBalance `json:"balance,omitempty"`
	FirstHeight *int32 `json:"first_height,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LastHeight *int32 `json:"last_height,omitempty"`
}

// NewResponsesAddress instantiates a new ResponsesAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesAddress() *ResponsesAddress {
	this := ResponsesAddress{}
	return &this
}

// NewResponsesAddressWithDefaults instantiates a new ResponsesAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesAddressWithDefaults() *ResponsesAddress {
	this := ResponsesAddress{}
	return &this
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *ResponsesAddress) GetBalance() ResponsesBalance {
	if o == nil || IsNil(o.Balance) {
		var ret ResponsesBalance
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesAddress) GetBalanceOk() (*ResponsesBalance, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *ResponsesAddress) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given ResponsesBalance and assigns it to the Balance field.
func (o *ResponsesAddress) SetBalance(v ResponsesBalance) {
	o.Balance = &v
}

// GetFirstHeight returns the FirstHeight field value if set, zero value otherwise.
func (o *ResponsesAddress) GetFirstHeight() int32 {
	if o == nil || IsNil(o.FirstHeight) {
		var ret int32
		return ret
	}
	return *o.FirstHeight
}

// GetFirstHeightOk returns a tuple with the FirstHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesAddress) GetFirstHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.FirstHeight) {
		return nil, false
	}
	return o.FirstHeight, true
}

// HasFirstHeight returns a boolean if a field has been set.
func (o *ResponsesAddress) HasFirstHeight() bool {
	if o != nil && !IsNil(o.FirstHeight) {
		return true
	}

	return false
}

// SetFirstHeight gets a reference to the given int32 and assigns it to the FirstHeight field.
func (o *ResponsesAddress) SetFirstHeight(v int32) {
	o.FirstHeight = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ResponsesAddress) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesAddress) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ResponsesAddress) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ResponsesAddress) SetHash(v string) {
	o.Hash = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesAddress) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesAddress) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesAddress) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResponsesAddress) SetId(v int32) {
	o.Id = &v
}

// GetLastHeight returns the LastHeight field value if set, zero value otherwise.
func (o *ResponsesAddress) GetLastHeight() int32 {
	if o == nil || IsNil(o.LastHeight) {
		var ret int32
		return ret
	}
	return *o.LastHeight
}

// GetLastHeightOk returns a tuple with the LastHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesAddress) GetLastHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.LastHeight) {
		return nil, false
	}
	return o.LastHeight, true
}

// HasLastHeight returns a boolean if a field has been set.
func (o *ResponsesAddress) HasLastHeight() bool {
	if o != nil && !IsNil(o.LastHeight) {
		return true
	}

	return false
}

// SetLastHeight gets a reference to the given int32 and assigns it to the LastHeight field.
func (o *ResponsesAddress) SetLastHeight(v int32) {
	o.LastHeight = &v
}

func (o ResponsesAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.FirstHeight) {
		toSerialize["first_height"] = o.FirstHeight
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastHeight) {
		toSerialize["last_height"] = o.LastHeight
	}
	return toSerialize, nil
}

type NullableResponsesAddress struct {
	value *ResponsesAddress
	isSet bool
}

func (v NullableResponsesAddress) Get() *ResponsesAddress {
	return v.value
}

func (v *NullableResponsesAddress) Set(val *ResponsesAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesAddress(val *ResponsesAddress) *NullableResponsesAddress {
	return &NullableResponsesAddress{value: val, isSet: true}
}

func (v NullableResponsesAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


