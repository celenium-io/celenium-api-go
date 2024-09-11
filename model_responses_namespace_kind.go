/*
Celenium API

This is docs of Celenium API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"encoding/json"
	"fmt"
)

// ResponsesNamespaceKind the model 'ResponsesNamespaceKind'
type ResponsesNamespaceKind string

// List of responses.NamespaceKind
const (
	PayForBlobNamespace ResponsesNamespaceKind = "pay_for_blob"
	TailPaddingNamespace ResponsesNamespaceKind = "tail_padding"
	TxNamespace ResponsesNamespaceKind = "tx"
	ParitySharesNamespace ResponsesNamespaceKind = "parity_shares"
	PrimaryReservedNamespace ResponsesNamespaceKind = "primary_reserved_padding"
	DefaultNamespace ResponsesNamespaceKind = "namespace"
)

// All allowed values of ResponsesNamespaceKind enum
var AllowedResponsesNamespaceKindEnumValues = []ResponsesNamespaceKind{
	"pay_for_blob",
	"tail_padding",
	"tx",
	"parity_shares",
	"primary_reserved_padding",
	"namespace",
}

func (v *ResponsesNamespaceKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResponsesNamespaceKind(value)
	for _, existing := range AllowedResponsesNamespaceKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResponsesNamespaceKind", value)
}

// NewResponsesNamespaceKindFromValue returns a pointer to a valid ResponsesNamespaceKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResponsesNamespaceKindFromValue(v string) (*ResponsesNamespaceKind, error) {
	ev := ResponsesNamespaceKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResponsesNamespaceKind: valid values are %v", v, AllowedResponsesNamespaceKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResponsesNamespaceKind) IsValid() bool {
	for _, existing := range AllowedResponsesNamespaceKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to responses.NamespaceKind value
func (v ResponsesNamespaceKind) Ptr() *ResponsesNamespaceKind {
	return &v
}

type NullableResponsesNamespaceKind struct {
	value *ResponsesNamespaceKind
	isSet bool
}

func (v NullableResponsesNamespaceKind) Get() *ResponsesNamespaceKind {
	return v.value
}

func (v *NullableResponsesNamespaceKind) Set(val *ResponsesNamespaceKind) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesNamespaceKind) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesNamespaceKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesNamespaceKind(val *ResponsesNamespaceKind) *NullableResponsesNamespaceKind {
	return &NullableResponsesNamespaceKind{value: val, isSet: true}
}

func (v NullableResponsesNamespaceKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesNamespaceKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
