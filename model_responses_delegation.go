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

// checks if the ResponsesDelegation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesDelegation{}

// ResponsesDelegation struct for ResponsesDelegation
type ResponsesDelegation struct {
	Amount *string `json:"amount,omitempty"`
	Delegator *ResponsesShortAddress `json:"delegator,omitempty"`
	Validator *ResponsesShortValidator `json:"validator,omitempty"`
}

// NewResponsesDelegation instantiates a new ResponsesDelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesDelegation() *ResponsesDelegation {
	this := ResponsesDelegation{}
	return &this
}

// NewResponsesDelegationWithDefaults instantiates a new ResponsesDelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesDelegationWithDefaults() *ResponsesDelegation {
	this := ResponsesDelegation{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ResponsesDelegation) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDelegation) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ResponsesDelegation) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *ResponsesDelegation) SetAmount(v string) {
	o.Amount = &v
}

// GetDelegator returns the Delegator field value if set, zero value otherwise.
func (o *ResponsesDelegation) GetDelegator() ResponsesShortAddress {
	if o == nil || IsNil(o.Delegator) {
		var ret ResponsesShortAddress
		return ret
	}
	return *o.Delegator
}

// GetDelegatorOk returns a tuple with the Delegator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDelegation) GetDelegatorOk() (*ResponsesShortAddress, bool) {
	if o == nil || IsNil(o.Delegator) {
		return nil, false
	}
	return o.Delegator, true
}

// HasDelegator returns a boolean if a field has been set.
func (o *ResponsesDelegation) HasDelegator() bool {
	if o != nil && !IsNil(o.Delegator) {
		return true
	}

	return false
}

// SetDelegator gets a reference to the given ResponsesShortAddress and assigns it to the Delegator field.
func (o *ResponsesDelegation) SetDelegator(v ResponsesShortAddress) {
	o.Delegator = &v
}

// GetValidator returns the Validator field value if set, zero value otherwise.
func (o *ResponsesDelegation) GetValidator() ResponsesShortValidator {
	if o == nil || IsNil(o.Validator) {
		var ret ResponsesShortValidator
		return ret
	}
	return *o.Validator
}

// GetValidatorOk returns a tuple with the Validator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDelegation) GetValidatorOk() (*ResponsesShortValidator, bool) {
	if o == nil || IsNil(o.Validator) {
		return nil, false
	}
	return o.Validator, true
}

// HasValidator returns a boolean if a field has been set.
func (o *ResponsesDelegation) HasValidator() bool {
	if o != nil && !IsNil(o.Validator) {
		return true
	}

	return false
}

// SetValidator gets a reference to the given ResponsesShortValidator and assigns it to the Validator field.
func (o *ResponsesDelegation) SetValidator(v ResponsesShortValidator) {
	o.Validator = &v
}

func (o ResponsesDelegation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesDelegation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Delegator) {
		toSerialize["delegator"] = o.Delegator
	}
	if !IsNil(o.Validator) {
		toSerialize["validator"] = o.Validator
	}
	return toSerialize, nil
}

type NullableResponsesDelegation struct {
	value *ResponsesDelegation
	isSet bool
}

func (v NullableResponsesDelegation) Get() *ResponsesDelegation {
	return v.value
}

func (v *NullableResponsesDelegation) Set(val *ResponsesDelegation) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesDelegation) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesDelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesDelegation(val *ResponsesDelegation) *NullableResponsesDelegation {
	return &NullableResponsesDelegation{value: val, isSet: true}
}

func (v NullableResponsesDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesDelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


