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
	"os"
	"time"
)

// checks if the ResponsesHyperlaneTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesHyperlaneTransfer{}

// ResponsesHyperlaneTransfer struct for ResponsesHyperlaneTransfer
type ResponsesHyperlaneTransfer struct {
	Address *ResponsesShortAddress `json:"address,omitempty"`
	Body *string `json:"body,omitempty"`
	Counterparty *ResponsesHyperlaneCounterparty `json:"counterparty,omitempty"`
	Denom *string `json:"denom,omitempty"`
	Height *int64 `json:"height,omitempty"`
	Id *int64 `json:"id,omitempty"`
	Mailbox **os.File `json:"mailbox,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
	Nonce *int64 `json:"nonce,omitempty"`
	Received *string `json:"received,omitempty"`
	Relayer *ResponsesShortAddress `json:"relayer,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	TokenId **os.File `json:"token_id,omitempty"`
	TxHash **os.File `json:"tx_hash,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *int64 `json:"version,omitempty"`
}

// NewResponsesHyperlaneTransfer instantiates a new ResponsesHyperlaneTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesHyperlaneTransfer() *ResponsesHyperlaneTransfer {
	this := ResponsesHyperlaneTransfer{}
	return &this
}

// NewResponsesHyperlaneTransferWithDefaults instantiates a new ResponsesHyperlaneTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesHyperlaneTransferWithDefaults() *ResponsesHyperlaneTransfer {
	this := ResponsesHyperlaneTransfer{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetAddress() ResponsesShortAddress {
	if o == nil || IsNil(o.Address) {
		var ret ResponsesShortAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetAddressOk() (*ResponsesShortAddress, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given ResponsesShortAddress and assigns it to the Address field.
func (o *ResponsesHyperlaneTransfer) SetAddress(v ResponsesShortAddress) {
	o.Address = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ResponsesHyperlaneTransfer) SetBody(v string) {
	o.Body = &v
}

// GetCounterparty returns the Counterparty field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetCounterparty() ResponsesHyperlaneCounterparty {
	if o == nil || IsNil(o.Counterparty) {
		var ret ResponsesHyperlaneCounterparty
		return ret
	}
	return *o.Counterparty
}

// GetCounterpartyOk returns a tuple with the Counterparty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetCounterpartyOk() (*ResponsesHyperlaneCounterparty, bool) {
	if o == nil || IsNil(o.Counterparty) {
		return nil, false
	}
	return o.Counterparty, true
}

// HasCounterparty returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasCounterparty() bool {
	if o != nil && !IsNil(o.Counterparty) {
		return true
	}

	return false
}

// SetCounterparty gets a reference to the given ResponsesHyperlaneCounterparty and assigns it to the Counterparty field.
func (o *ResponsesHyperlaneTransfer) SetCounterparty(v ResponsesHyperlaneCounterparty) {
	o.Counterparty = &v
}

// GetDenom returns the Denom field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetDenom() string {
	if o == nil || IsNil(o.Denom) {
		var ret string
		return ret
	}
	return *o.Denom
}

// GetDenomOk returns a tuple with the Denom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetDenomOk() (*string, bool) {
	if o == nil || IsNil(o.Denom) {
		return nil, false
	}
	return o.Denom, true
}

// HasDenom returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasDenom() bool {
	if o != nil && !IsNil(o.Denom) {
		return true
	}

	return false
}

// SetDenom gets a reference to the given string and assigns it to the Denom field.
func (o *ResponsesHyperlaneTransfer) SetDenom(v string) {
	o.Denom = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ResponsesHyperlaneTransfer) SetHeight(v int64) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ResponsesHyperlaneTransfer) SetId(v int64) {
	o.Id = &v
}

// GetMailbox returns the Mailbox field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetMailbox() *os.File {
	if o == nil || IsNil(o.Mailbox) {
		var ret *os.File
		return ret
	}
	return *o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetMailboxOk() (**os.File, bool) {
	if o == nil || IsNil(o.Mailbox) {
		return nil, false
	}
	return o.Mailbox, true
}

// HasMailbox returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasMailbox() bool {
	if o != nil && !IsNil(o.Mailbox) {
		return true
	}

	return false
}

// SetMailbox gets a reference to the given *os.File and assigns it to the Mailbox field.
func (o *ResponsesHyperlaneTransfer) SetMailbox(v *os.File) {
	o.Mailbox = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *ResponsesHyperlaneTransfer) SetMetadata(v string) {
	o.Metadata = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetNonce() int64 {
	if o == nil || IsNil(o.Nonce) {
		var ret int64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetNonceOk() (*int64, bool) {
	if o == nil || IsNil(o.Nonce) {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasNonce() bool {
	if o != nil && !IsNil(o.Nonce) {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int64 and assigns it to the Nonce field.
func (o *ResponsesHyperlaneTransfer) SetNonce(v int64) {
	o.Nonce = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetReceived() string {
	if o == nil || IsNil(o.Received) {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetReceivedOk() (*string, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given string and assigns it to the Received field.
func (o *ResponsesHyperlaneTransfer) SetReceived(v string) {
	o.Received = &v
}

// GetRelayer returns the Relayer field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetRelayer() ResponsesShortAddress {
	if o == nil || IsNil(o.Relayer) {
		var ret ResponsesShortAddress
		return ret
	}
	return *o.Relayer
}

// GetRelayerOk returns a tuple with the Relayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetRelayerOk() (*ResponsesShortAddress, bool) {
	if o == nil || IsNil(o.Relayer) {
		return nil, false
	}
	return o.Relayer, true
}

// HasRelayer returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasRelayer() bool {
	if o != nil && !IsNil(o.Relayer) {
		return true
	}

	return false
}

// SetRelayer gets a reference to the given ResponsesShortAddress and assigns it to the Relayer field.
func (o *ResponsesHyperlaneTransfer) SetRelayer(v ResponsesShortAddress) {
	o.Relayer = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *ResponsesHyperlaneTransfer) SetTime(v time.Time) {
	o.Time = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetTokenId() *os.File {
	if o == nil || IsNil(o.TokenId) {
		var ret *os.File
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetTokenIdOk() (**os.File, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given *os.File and assigns it to the TokenId field.
func (o *ResponsesHyperlaneTransfer) SetTokenId(v *os.File) {
	o.TokenId = &v
}

// GetTxHash returns the TxHash field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetTxHash() *os.File {
	if o == nil || IsNil(o.TxHash) {
		var ret *os.File
		return ret
	}
	return *o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetTxHashOk() (**os.File, bool) {
	if o == nil || IsNil(o.TxHash) {
		return nil, false
	}
	return o.TxHash, true
}

// HasTxHash returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasTxHash() bool {
	if o != nil && !IsNil(o.TxHash) {
		return true
	}

	return false
}

// SetTxHash gets a reference to the given *os.File and assigns it to the TxHash field.
func (o *ResponsesHyperlaneTransfer) SetTxHash(v *os.File) {
	o.TxHash = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResponsesHyperlaneTransfer) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ResponsesHyperlaneTransfer) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesHyperlaneTransfer) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ResponsesHyperlaneTransfer) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ResponsesHyperlaneTransfer) SetVersion(v int64) {
	o.Version = &v
}

func (o ResponsesHyperlaneTransfer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesHyperlaneTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !IsNil(o.Counterparty) {
		toSerialize["counterparty"] = o.Counterparty
	}
	if !IsNil(o.Denom) {
		toSerialize["denom"] = o.Denom
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mailbox) {
		toSerialize["mailbox"] = o.Mailbox
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Nonce) {
		toSerialize["nonce"] = o.Nonce
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Relayer) {
		toSerialize["relayer"] = o.Relayer
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.TxHash) {
		toSerialize["tx_hash"] = o.TxHash
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableResponsesHyperlaneTransfer struct {
	value *ResponsesHyperlaneTransfer
	isSet bool
}

func (v NullableResponsesHyperlaneTransfer) Get() *ResponsesHyperlaneTransfer {
	return v.value
}

func (v *NullableResponsesHyperlaneTransfer) Set(val *ResponsesHyperlaneTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesHyperlaneTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesHyperlaneTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesHyperlaneTransfer(val *ResponsesHyperlaneTransfer) *NullableResponsesHyperlaneTransfer {
	return &NullableResponsesHyperlaneTransfer{value: val, isSet: true}
}

func (v NullableResponsesHyperlaneTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesHyperlaneTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


