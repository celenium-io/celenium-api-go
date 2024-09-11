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

// GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus the model 'GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus'
type GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus string

// List of github_com_celenium-io_celestia-indexer_internal_storage_types.Status
const (
	StatusSuccess GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus = "success"
	StatusFailed GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus = "failed"
)

// All allowed values of GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus enum
var AllowedGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusEnumValues = []GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus{
	"success",
	"failed",
}

func (v *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus(value)
	for _, existing := range AllowedGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus", value)
}

// NewGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusFromValue returns a pointer to a valid GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusFromValue(v string) (*GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus, error) {
	ev := GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus: valid values are %v", v, AllowedGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) IsValid() bool {
	for _, existing := range AllowedGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to github_com_celenium-io_celestia-indexer_internal_storage_types.Status value
func (v GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) Ptr() *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus {
	return &v
}

type NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus struct {
	value *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus
	isSet bool
}

func (v NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) Get() *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus {
	return v.value
}

func (v *NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) Set(val *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus(val *GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) *NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus {
	return &NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus{value: val, isSet: true}
}

func (v NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
