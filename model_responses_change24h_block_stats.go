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

// checks if the ResponsesChange24hBlockStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesChange24hBlockStats{}

// ResponsesChange24hBlockStats struct for ResponsesChange24hBlockStats
type ResponsesChange24hBlockStats struct {
	BlobsSize24h *float32 `json:"blobs_size_24h,omitempty"`
	BytesInBlock24h *float32 `json:"bytes_in_block_24h,omitempty"`
	Fee24h *float32 `json:"fee_24h,omitempty"`
	TxCount24h *float32 `json:"tx_count_24h,omitempty"`
}

// NewResponsesChange24hBlockStats instantiates a new ResponsesChange24hBlockStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesChange24hBlockStats() *ResponsesChange24hBlockStats {
	this := ResponsesChange24hBlockStats{}
	return &this
}

// NewResponsesChange24hBlockStatsWithDefaults instantiates a new ResponsesChange24hBlockStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesChange24hBlockStatsWithDefaults() *ResponsesChange24hBlockStats {
	this := ResponsesChange24hBlockStats{}
	return &this
}

// GetBlobsSize24h returns the BlobsSize24h field value if set, zero value otherwise.
func (o *ResponsesChange24hBlockStats) GetBlobsSize24h() float32 {
	if o == nil || IsNil(o.BlobsSize24h) {
		var ret float32
		return ret
	}
	return *o.BlobsSize24h
}

// GetBlobsSize24hOk returns a tuple with the BlobsSize24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesChange24hBlockStats) GetBlobsSize24hOk() (*float32, bool) {
	if o == nil || IsNil(o.BlobsSize24h) {
		return nil, false
	}
	return o.BlobsSize24h, true
}

// HasBlobsSize24h returns a boolean if a field has been set.
func (o *ResponsesChange24hBlockStats) HasBlobsSize24h() bool {
	if o != nil && !IsNil(o.BlobsSize24h) {
		return true
	}

	return false
}

// SetBlobsSize24h gets a reference to the given float32 and assigns it to the BlobsSize24h field.
func (o *ResponsesChange24hBlockStats) SetBlobsSize24h(v float32) {
	o.BlobsSize24h = &v
}

// GetBytesInBlock24h returns the BytesInBlock24h field value if set, zero value otherwise.
func (o *ResponsesChange24hBlockStats) GetBytesInBlock24h() float32 {
	if o == nil || IsNil(o.BytesInBlock24h) {
		var ret float32
		return ret
	}
	return *o.BytesInBlock24h
}

// GetBytesInBlock24hOk returns a tuple with the BytesInBlock24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesChange24hBlockStats) GetBytesInBlock24hOk() (*float32, bool) {
	if o == nil || IsNil(o.BytesInBlock24h) {
		return nil, false
	}
	return o.BytesInBlock24h, true
}

// HasBytesInBlock24h returns a boolean if a field has been set.
func (o *ResponsesChange24hBlockStats) HasBytesInBlock24h() bool {
	if o != nil && !IsNil(o.BytesInBlock24h) {
		return true
	}

	return false
}

// SetBytesInBlock24h gets a reference to the given float32 and assigns it to the BytesInBlock24h field.
func (o *ResponsesChange24hBlockStats) SetBytesInBlock24h(v float32) {
	o.BytesInBlock24h = &v
}

// GetFee24h returns the Fee24h field value if set, zero value otherwise.
func (o *ResponsesChange24hBlockStats) GetFee24h() float32 {
	if o == nil || IsNil(o.Fee24h) {
		var ret float32
		return ret
	}
	return *o.Fee24h
}

// GetFee24hOk returns a tuple with the Fee24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesChange24hBlockStats) GetFee24hOk() (*float32, bool) {
	if o == nil || IsNil(o.Fee24h) {
		return nil, false
	}
	return o.Fee24h, true
}

// HasFee24h returns a boolean if a field has been set.
func (o *ResponsesChange24hBlockStats) HasFee24h() bool {
	if o != nil && !IsNil(o.Fee24h) {
		return true
	}

	return false
}

// SetFee24h gets a reference to the given float32 and assigns it to the Fee24h field.
func (o *ResponsesChange24hBlockStats) SetFee24h(v float32) {
	o.Fee24h = &v
}

// GetTxCount24h returns the TxCount24h field value if set, zero value otherwise.
func (o *ResponsesChange24hBlockStats) GetTxCount24h() float32 {
	if o == nil || IsNil(o.TxCount24h) {
		var ret float32
		return ret
	}
	return *o.TxCount24h
}

// GetTxCount24hOk returns a tuple with the TxCount24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesChange24hBlockStats) GetTxCount24hOk() (*float32, bool) {
	if o == nil || IsNil(o.TxCount24h) {
		return nil, false
	}
	return o.TxCount24h, true
}

// HasTxCount24h returns a boolean if a field has been set.
func (o *ResponsesChange24hBlockStats) HasTxCount24h() bool {
	if o != nil && !IsNil(o.TxCount24h) {
		return true
	}

	return false
}

// SetTxCount24h gets a reference to the given float32 and assigns it to the TxCount24h field.
func (o *ResponsesChange24hBlockStats) SetTxCount24h(v float32) {
	o.TxCount24h = &v
}

func (o ResponsesChange24hBlockStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesChange24hBlockStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobsSize24h) {
		toSerialize["blobs_size_24h"] = o.BlobsSize24h
	}
	if !IsNil(o.BytesInBlock24h) {
		toSerialize["bytes_in_block_24h"] = o.BytesInBlock24h
	}
	if !IsNil(o.Fee24h) {
		toSerialize["fee_24h"] = o.Fee24h
	}
	if !IsNil(o.TxCount24h) {
		toSerialize["tx_count_24h"] = o.TxCount24h
	}
	return toSerialize, nil
}

type NullableResponsesChange24hBlockStats struct {
	value *ResponsesChange24hBlockStats
	isSet bool
}

func (v NullableResponsesChange24hBlockStats) Get() *ResponsesChange24hBlockStats {
	return v.value
}

func (v *NullableResponsesChange24hBlockStats) Set(val *ResponsesChange24hBlockStats) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesChange24hBlockStats) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesChange24hBlockStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesChange24hBlockStats(val *ResponsesChange24hBlockStats) *NullableResponsesChange24hBlockStats {
	return &NullableResponsesChange24hBlockStats{value: val, isSet: true}
}

func (v NullableResponsesChange24hBlockStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesChange24hBlockStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


