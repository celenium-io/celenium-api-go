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

// checks if the ResponsesRollupGroupedStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesRollupGroupedStats{}

// ResponsesRollupGroupedStats struct for ResponsesRollupGroupedStats
type ResponsesRollupGroupedStats struct {
	BlobsCount *int32 `json:"blobs_count,omitempty"`
	Fee *string `json:"fee,omitempty"`
	Group *string `json:"group,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewResponsesRollupGroupedStats instantiates a new ResponsesRollupGroupedStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesRollupGroupedStats() *ResponsesRollupGroupedStats {
	this := ResponsesRollupGroupedStats{}
	return &this
}

// NewResponsesRollupGroupedStatsWithDefaults instantiates a new ResponsesRollupGroupedStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesRollupGroupedStatsWithDefaults() *ResponsesRollupGroupedStats {
	this := ResponsesRollupGroupedStats{}
	return &this
}

// GetBlobsCount returns the BlobsCount field value if set, zero value otherwise.
func (o *ResponsesRollupGroupedStats) GetBlobsCount() int32 {
	if o == nil || IsNil(o.BlobsCount) {
		var ret int32
		return ret
	}
	return *o.BlobsCount
}

// GetBlobsCountOk returns a tuple with the BlobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupGroupedStats) GetBlobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BlobsCount) {
		return nil, false
	}
	return o.BlobsCount, true
}

// HasBlobsCount returns a boolean if a field has been set.
func (o *ResponsesRollupGroupedStats) HasBlobsCount() bool {
	if o != nil && !IsNil(o.BlobsCount) {
		return true
	}

	return false
}

// SetBlobsCount gets a reference to the given int32 and assigns it to the BlobsCount field.
func (o *ResponsesRollupGroupedStats) SetBlobsCount(v int32) {
	o.BlobsCount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ResponsesRollupGroupedStats) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupGroupedStats) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ResponsesRollupGroupedStats) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *ResponsesRollupGroupedStats) SetFee(v string) {
	o.Fee = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ResponsesRollupGroupedStats) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupGroupedStats) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ResponsesRollupGroupedStats) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ResponsesRollupGroupedStats) SetGroup(v string) {
	o.Group = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesRollupGroupedStats) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupGroupedStats) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesRollupGroupedStats) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ResponsesRollupGroupedStats) SetSize(v int32) {
	o.Size = &v
}

func (o ResponsesRollupGroupedStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesRollupGroupedStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobsCount) {
		toSerialize["blobs_count"] = o.BlobsCount
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableResponsesRollupGroupedStats struct {
	value *ResponsesRollupGroupedStats
	isSet bool
}

func (v NullableResponsesRollupGroupedStats) Get() *ResponsesRollupGroupedStats {
	return v.value
}

func (v *NullableResponsesRollupGroupedStats) Set(val *ResponsesRollupGroupedStats) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesRollupGroupedStats) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesRollupGroupedStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesRollupGroupedStats(val *ResponsesRollupGroupedStats) *NullableResponsesRollupGroupedStats {
	return &NullableResponsesRollupGroupedStats{value: val, isSet: true}
}

func (v NullableResponsesRollupGroupedStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesRollupGroupedStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


