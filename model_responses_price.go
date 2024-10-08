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
	"time"
)

// checks if the ResponsesPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesPrice{}

// ResponsesPrice struct for ResponsesPrice
type ResponsesPrice struct {
	Close *string `json:"close,omitempty"`
	High *string `json:"high,omitempty"`
	Low *string `json:"low,omitempty"`
	Open *string `json:"open,omitempty"`
	Time *time.Time `json:"time,omitempty"`
}

// NewResponsesPrice instantiates a new ResponsesPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesPrice() *ResponsesPrice {
	this := ResponsesPrice{}
	return &this
}

// NewResponsesPriceWithDefaults instantiates a new ResponsesPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesPriceWithDefaults() *ResponsesPrice {
	this := ResponsesPrice{}
	return &this
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *ResponsesPrice) GetClose() string {
	if o == nil || IsNil(o.Close) {
		var ret string
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesPrice) GetCloseOk() (*string, bool) {
	if o == nil || IsNil(o.Close) {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *ResponsesPrice) HasClose() bool {
	if o != nil && !IsNil(o.Close) {
		return true
	}

	return false
}

// SetClose gets a reference to the given string and assigns it to the Close field.
func (o *ResponsesPrice) SetClose(v string) {
	o.Close = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *ResponsesPrice) GetHigh() string {
	if o == nil || IsNil(o.High) {
		var ret string
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesPrice) GetHighOk() (*string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *ResponsesPrice) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given string and assigns it to the High field.
func (o *ResponsesPrice) SetHigh(v string) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ResponsesPrice) GetLow() string {
	if o == nil || IsNil(o.Low) {
		var ret string
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesPrice) GetLowOk() (*string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ResponsesPrice) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given string and assigns it to the Low field.
func (o *ResponsesPrice) SetLow(v string) {
	o.Low = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *ResponsesPrice) GetOpen() string {
	if o == nil || IsNil(o.Open) {
		var ret string
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesPrice) GetOpenOk() (*string, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *ResponsesPrice) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given string and assigns it to the Open field.
func (o *ResponsesPrice) SetOpen(v string) {
	o.Open = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesPrice) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesPrice) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesPrice) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesPrice) SetTime(v time.Time) {
	o.Time = &v
}

func (o ResponsesPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Close) {
		toSerialize["close"] = o.Close
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableResponsesPrice struct {
	value *ResponsesPrice
	isSet bool
}

func (v NullableResponsesPrice) Get() *ResponsesPrice {
	return v.value
}

func (v *NullableResponsesPrice) Set(val *ResponsesPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesPrice(val *ResponsesPrice) *NullableResponsesPrice {
	return &NullableResponsesPrice{value: val, isSet: true}
}

func (v NullableResponsesPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


