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

// checks if the ResponsesDenomMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesDenomMetadata{}

// ResponsesDenomMetadata struct for ResponsesDenomMetadata
type ResponsesDenomMetadata struct {
	Base *string `json:"base,omitempty"`
	Description *string `json:"description,omitempty"`
	Display *string `json:"display,omitempty"`
	Name *string `json:"name,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Units []int32 `json:"units,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewResponsesDenomMetadata instantiates a new ResponsesDenomMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesDenomMetadata() *ResponsesDenomMetadata {
	this := ResponsesDenomMetadata{}
	return &this
}

// NewResponsesDenomMetadataWithDefaults instantiates a new ResponsesDenomMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesDenomMetadataWithDefaults() *ResponsesDenomMetadata {
	this := ResponsesDenomMetadata{}
	return &this
}

// GetBase returns the Base field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetBase() string {
	if o == nil || IsNil(o.Base) {
		var ret string
		return ret
	}
	return *o.Base
}

// GetBaseOk returns a tuple with the Base field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetBaseOk() (*string, bool) {
	if o == nil || IsNil(o.Base) {
		return nil, false
	}
	return o.Base, true
}

// HasBase returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasBase() bool {
	if o != nil && !IsNil(o.Base) {
		return true
	}

	return false
}

// SetBase gets a reference to the given string and assigns it to the Base field.
func (o *ResponsesDenomMetadata) SetBase(v string) {
	o.Base = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResponsesDenomMetadata) SetDescription(v string) {
	o.Description = &v
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetDisplay() string {
	if o == nil || IsNil(o.Display) {
		var ret string
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.Display) {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasDisplay() bool {
	if o != nil && !IsNil(o.Display) {
		return true
	}

	return false
}

// SetDisplay gets a reference to the given string and assigns it to the Display field.
func (o *ResponsesDenomMetadata) SetDisplay(v string) {
	o.Display = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponsesDenomMetadata) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ResponsesDenomMetadata) SetSymbol(v string) {
	o.Symbol = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetUnits() []int32 {
	if o == nil || IsNil(o.Units) {
		var ret []int32
		return ret
	}
	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetUnitsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given []int32 and assigns it to the Units field.
func (o *ResponsesDenomMetadata) SetUnits(v []int32) {
	o.Units = v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ResponsesDenomMetadata) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesDenomMetadata) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ResponsesDenomMetadata) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ResponsesDenomMetadata) SetUri(v string) {
	o.Uri = &v
}

func (o ResponsesDenomMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesDenomMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base) {
		toSerialize["base"] = o.Base
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Display) {
		toSerialize["display"] = o.Display
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableResponsesDenomMetadata struct {
	value *ResponsesDenomMetadata
	isSet bool
}

func (v NullableResponsesDenomMetadata) Get() *ResponsesDenomMetadata {
	return v.value
}

func (v *NullableResponsesDenomMetadata) Set(val *ResponsesDenomMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesDenomMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesDenomMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesDenomMetadata(val *ResponsesDenomMetadata) *NullableResponsesDenomMetadata {
	return &NullableResponsesDenomMetadata{value: val, isSet: true}
}

func (v NullableResponsesDenomMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesDenomMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


