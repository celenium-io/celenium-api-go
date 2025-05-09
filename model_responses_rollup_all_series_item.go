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

// checks if the ResponsesRollupAllSeriesItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesRollupAllSeriesItem{}

// ResponsesRollupAllSeriesItem struct for ResponsesRollupAllSeriesItem
type ResponsesRollupAllSeriesItem struct {
	BlobsCount *int32 `json:"blobs_count,omitempty"`
	Fee *string `json:"fee,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Name *string `json:"name,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewResponsesRollupAllSeriesItem instantiates a new ResponsesRollupAllSeriesItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesRollupAllSeriesItem() *ResponsesRollupAllSeriesItem {
	this := ResponsesRollupAllSeriesItem{}
	return &this
}

// NewResponsesRollupAllSeriesItemWithDefaults instantiates a new ResponsesRollupAllSeriesItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesRollupAllSeriesItemWithDefaults() *ResponsesRollupAllSeriesItem {
	this := ResponsesRollupAllSeriesItem{}
	return &this
}

// GetBlobsCount returns the BlobsCount field value if set, zero value otherwise.
func (o *ResponsesRollupAllSeriesItem) GetBlobsCount() int32 {
	if o == nil || IsNil(o.BlobsCount) {
		var ret int32
		return ret
	}
	return *o.BlobsCount
}

// GetBlobsCountOk returns a tuple with the BlobsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupAllSeriesItem) GetBlobsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BlobsCount) {
		return nil, false
	}
	return o.BlobsCount, true
}

// HasBlobsCount returns a boolean if a field has been set.
func (o *ResponsesRollupAllSeriesItem) HasBlobsCount() bool {
	if o != nil && !IsNil(o.BlobsCount) {
		return true
	}

	return false
}

// SetBlobsCount gets a reference to the given int32 and assigns it to the BlobsCount field.
func (o *ResponsesRollupAllSeriesItem) SetBlobsCount(v int32) {
	o.BlobsCount = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *ResponsesRollupAllSeriesItem) GetFee() string {
	if o == nil || IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupAllSeriesItem) GetFeeOk() (*string, bool) {
	if o == nil || IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *ResponsesRollupAllSeriesItem) HasFee() bool {
	if o != nil && !IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *ResponsesRollupAllSeriesItem) SetFee(v string) {
	o.Fee = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ResponsesRollupAllSeriesItem) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupAllSeriesItem) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ResponsesRollupAllSeriesItem) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *ResponsesRollupAllSeriesItem) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponsesRollupAllSeriesItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupAllSeriesItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponsesRollupAllSeriesItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponsesRollupAllSeriesItem) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ResponsesRollupAllSeriesItem) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesRollupAllSeriesItem) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ResponsesRollupAllSeriesItem) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *ResponsesRollupAllSeriesItem) SetSize(v int32) {
	o.Size = &v
}

func (o ResponsesRollupAllSeriesItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesRollupAllSeriesItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlobsCount) {
		toSerialize["blobs_count"] = o.BlobsCount
	}
	if !IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableResponsesRollupAllSeriesItem struct {
	value *ResponsesRollupAllSeriesItem
	isSet bool
}

func (v NullableResponsesRollupAllSeriesItem) Get() *ResponsesRollupAllSeriesItem {
	return v.value
}

func (v *NullableResponsesRollupAllSeriesItem) Set(val *ResponsesRollupAllSeriesItem) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesRollupAllSeriesItem) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesRollupAllSeriesItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesRollupAllSeriesItem(val *ResponsesRollupAllSeriesItem) *NullableResponsesRollupAllSeriesItem {
	return &NullableResponsesRollupAllSeriesItem{value: val, isSet: true}
}

func (v NullableResponsesRollupAllSeriesItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesRollupAllSeriesItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


