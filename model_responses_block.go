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

// checks if the ResponsesBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponsesBlock{}

// ResponsesBlock struct for ResponsesBlock
type ResponsesBlock struct {
	AppHash *string `json:"app_hash,omitempty"`
	ConsensusHash *string `json:"consensus_hash,omitempty"`
	DataHash *string `json:"data_hash,omitempty"`
	EvidenceHash *string `json:"evidence_hash,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Id *int32 `json:"id,omitempty"`
	LastCommitHash *string `json:"last_commit_hash,omitempty"`
	LastResultsHash *string `json:"last_results_hash,omitempty"`
	MessageTypes []string `json:"message_types,omitempty"`
	NextValidatorsHash *string `json:"next_validators_hash,omitempty"`
	ParentHash *string `json:"parent_hash,omitempty"`
	Proposer *ResponsesShortValidator `json:"proposer,omitempty"`
	Stats *ResponsesBlockStats `json:"stats,omitempty"`
	Time *string `json:"time,omitempty"`
	ValidatorsHash *string `json:"validators_hash,omitempty"`
	VersionApp *string `json:"version_app,omitempty"`
	VersionBlock *string `json:"version_block,omitempty"`
}

// NewResponsesBlock instantiates a new ResponsesBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponsesBlock() *ResponsesBlock {
	this := ResponsesBlock{}
	return &this
}

// NewResponsesBlockWithDefaults instantiates a new ResponsesBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponsesBlockWithDefaults() *ResponsesBlock {
	this := ResponsesBlock{}
	return &this
}

// GetAppHash returns the AppHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetAppHash() string {
	if o == nil || IsNil(o.AppHash) {
		var ret string
		return ret
	}
	return *o.AppHash
}

// GetAppHashOk returns a tuple with the AppHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetAppHashOk() (*string, bool) {
	if o == nil || IsNil(o.AppHash) {
		return nil, false
	}
	return o.AppHash, true
}

// HasAppHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasAppHash() bool {
	if o != nil && !IsNil(o.AppHash) {
		return true
	}

	return false
}

// SetAppHash gets a reference to the given string and assigns it to the AppHash field.
func (o *ResponsesBlock) SetAppHash(v string) {
	o.AppHash = &v
}

// GetConsensusHash returns the ConsensusHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetConsensusHash() string {
	if o == nil || IsNil(o.ConsensusHash) {
		var ret string
		return ret
	}
	return *o.ConsensusHash
}

// GetConsensusHashOk returns a tuple with the ConsensusHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetConsensusHashOk() (*string, bool) {
	if o == nil || IsNil(o.ConsensusHash) {
		return nil, false
	}
	return o.ConsensusHash, true
}

// HasConsensusHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasConsensusHash() bool {
	if o != nil && !IsNil(o.ConsensusHash) {
		return true
	}

	return false
}

// SetConsensusHash gets a reference to the given string and assigns it to the ConsensusHash field.
func (o *ResponsesBlock) SetConsensusHash(v string) {
	o.ConsensusHash = &v
}

// GetDataHash returns the DataHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetDataHash() string {
	if o == nil || IsNil(o.DataHash) {
		var ret string
		return ret
	}
	return *o.DataHash
}

// GetDataHashOk returns a tuple with the DataHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetDataHashOk() (*string, bool) {
	if o == nil || IsNil(o.DataHash) {
		return nil, false
	}
	return o.DataHash, true
}

// HasDataHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasDataHash() bool {
	if o != nil && !IsNil(o.DataHash) {
		return true
	}

	return false
}

// SetDataHash gets a reference to the given string and assigns it to the DataHash field.
func (o *ResponsesBlock) SetDataHash(v string) {
	o.DataHash = &v
}

// GetEvidenceHash returns the EvidenceHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetEvidenceHash() string {
	if o == nil || IsNil(o.EvidenceHash) {
		var ret string
		return ret
	}
	return *o.EvidenceHash
}

// GetEvidenceHashOk returns a tuple with the EvidenceHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetEvidenceHashOk() (*string, bool) {
	if o == nil || IsNil(o.EvidenceHash) {
		return nil, false
	}
	return o.EvidenceHash, true
}

// HasEvidenceHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasEvidenceHash() bool {
	if o != nil && !IsNil(o.EvidenceHash) {
		return true
	}

	return false
}

// SetEvidenceHash gets a reference to the given string and assigns it to the EvidenceHash field.
func (o *ResponsesBlock) SetEvidenceHash(v string) {
	o.EvidenceHash = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ResponsesBlock) SetHash(v string) {
	o.Hash = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ResponsesBlock) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ResponsesBlock) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ResponsesBlock) SetHeight(v int32) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponsesBlock) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponsesBlock) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ResponsesBlock) SetId(v int32) {
	o.Id = &v
}

// GetLastCommitHash returns the LastCommitHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetLastCommitHash() string {
	if o == nil || IsNil(o.LastCommitHash) {
		var ret string
		return ret
	}
	return *o.LastCommitHash
}

// GetLastCommitHashOk returns a tuple with the LastCommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetLastCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.LastCommitHash) {
		return nil, false
	}
	return o.LastCommitHash, true
}

// HasLastCommitHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasLastCommitHash() bool {
	if o != nil && !IsNil(o.LastCommitHash) {
		return true
	}

	return false
}

// SetLastCommitHash gets a reference to the given string and assigns it to the LastCommitHash field.
func (o *ResponsesBlock) SetLastCommitHash(v string) {
	o.LastCommitHash = &v
}

// GetLastResultsHash returns the LastResultsHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetLastResultsHash() string {
	if o == nil || IsNil(o.LastResultsHash) {
		var ret string
		return ret
	}
	return *o.LastResultsHash
}

// GetLastResultsHashOk returns a tuple with the LastResultsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetLastResultsHashOk() (*string, bool) {
	if o == nil || IsNil(o.LastResultsHash) {
		return nil, false
	}
	return o.LastResultsHash, true
}

// HasLastResultsHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasLastResultsHash() bool {
	if o != nil && !IsNil(o.LastResultsHash) {
		return true
	}

	return false
}

// SetLastResultsHash gets a reference to the given string and assigns it to the LastResultsHash field.
func (o *ResponsesBlock) SetLastResultsHash(v string) {
	o.LastResultsHash = &v
}

// GetMessageTypes returns the MessageTypes field value if set, zero value otherwise.
func (o *ResponsesBlock) GetMessageTypes() []string {
	if o == nil || IsNil(o.MessageTypes) {
		var ret []string
		return ret
	}
	return o.MessageTypes
}

// GetMessageTypesOk returns a tuple with the MessageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetMessageTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.MessageTypes) {
		return nil, false
	}
	return o.MessageTypes, true
}

// HasMessageTypes returns a boolean if a field has been set.
func (o *ResponsesBlock) HasMessageTypes() bool {
	if o != nil && !IsNil(o.MessageTypes) {
		return true
	}

	return false
}

// SetMessageTypes gets a reference to the given []string and assigns it to the MessageTypes field.
func (o *ResponsesBlock) SetMessageTypes(v []string) {
	o.MessageTypes = v
}

// GetNextValidatorsHash returns the NextValidatorsHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetNextValidatorsHash() string {
	if o == nil || IsNil(o.NextValidatorsHash) {
		var ret string
		return ret
	}
	return *o.NextValidatorsHash
}

// GetNextValidatorsHashOk returns a tuple with the NextValidatorsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetNextValidatorsHashOk() (*string, bool) {
	if o == nil || IsNil(o.NextValidatorsHash) {
		return nil, false
	}
	return o.NextValidatorsHash, true
}

// HasNextValidatorsHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasNextValidatorsHash() bool {
	if o != nil && !IsNil(o.NextValidatorsHash) {
		return true
	}

	return false
}

// SetNextValidatorsHash gets a reference to the given string and assigns it to the NextValidatorsHash field.
func (o *ResponsesBlock) SetNextValidatorsHash(v string) {
	o.NextValidatorsHash = &v
}

// GetParentHash returns the ParentHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetParentHash() string {
	if o == nil || IsNil(o.ParentHash) {
		var ret string
		return ret
	}
	return *o.ParentHash
}

// GetParentHashOk returns a tuple with the ParentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetParentHashOk() (*string, bool) {
	if o == nil || IsNil(o.ParentHash) {
		return nil, false
	}
	return o.ParentHash, true
}

// HasParentHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasParentHash() bool {
	if o != nil && !IsNil(o.ParentHash) {
		return true
	}

	return false
}

// SetParentHash gets a reference to the given string and assigns it to the ParentHash field.
func (o *ResponsesBlock) SetParentHash(v string) {
	o.ParentHash = &v
}

// GetProposer returns the Proposer field value if set, zero value otherwise.
func (o *ResponsesBlock) GetProposer() ResponsesShortValidator {
	if o == nil || IsNil(o.Proposer) {
		var ret ResponsesShortValidator
		return ret
	}
	return *o.Proposer
}

// GetProposerOk returns a tuple with the Proposer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetProposerOk() (*ResponsesShortValidator, bool) {
	if o == nil || IsNil(o.Proposer) {
		return nil, false
	}
	return o.Proposer, true
}

// HasProposer returns a boolean if a field has been set.
func (o *ResponsesBlock) HasProposer() bool {
	if o != nil && !IsNil(o.Proposer) {
		return true
	}

	return false
}

// SetProposer gets a reference to the given ResponsesShortValidator and assigns it to the Proposer field.
func (o *ResponsesBlock) SetProposer(v ResponsesShortValidator) {
	o.Proposer = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ResponsesBlock) GetStats() ResponsesBlockStats {
	if o == nil || IsNil(o.Stats) {
		var ret ResponsesBlockStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetStatsOk() (*ResponsesBlockStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ResponsesBlock) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ResponsesBlockStats and assigns it to the Stats field.
func (o *ResponsesBlock) SetStats(v ResponsesBlockStats) {
	o.Stats = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ResponsesBlock) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ResponsesBlock) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ResponsesBlock) SetTime(v string) {
	o.Time = &v
}

// GetValidatorsHash returns the ValidatorsHash field value if set, zero value otherwise.
func (o *ResponsesBlock) GetValidatorsHash() string {
	if o == nil || IsNil(o.ValidatorsHash) {
		var ret string
		return ret
	}
	return *o.ValidatorsHash
}

// GetValidatorsHashOk returns a tuple with the ValidatorsHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetValidatorsHashOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorsHash) {
		return nil, false
	}
	return o.ValidatorsHash, true
}

// HasValidatorsHash returns a boolean if a field has been set.
func (o *ResponsesBlock) HasValidatorsHash() bool {
	if o != nil && !IsNil(o.ValidatorsHash) {
		return true
	}

	return false
}

// SetValidatorsHash gets a reference to the given string and assigns it to the ValidatorsHash field.
func (o *ResponsesBlock) SetValidatorsHash(v string) {
	o.ValidatorsHash = &v
}

// GetVersionApp returns the VersionApp field value if set, zero value otherwise.
func (o *ResponsesBlock) GetVersionApp() string {
	if o == nil || IsNil(o.VersionApp) {
		var ret string
		return ret
	}
	return *o.VersionApp
}

// GetVersionAppOk returns a tuple with the VersionApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetVersionAppOk() (*string, bool) {
	if o == nil || IsNil(o.VersionApp) {
		return nil, false
	}
	return o.VersionApp, true
}

// HasVersionApp returns a boolean if a field has been set.
func (o *ResponsesBlock) HasVersionApp() bool {
	if o != nil && !IsNil(o.VersionApp) {
		return true
	}

	return false
}

// SetVersionApp gets a reference to the given string and assigns it to the VersionApp field.
func (o *ResponsesBlock) SetVersionApp(v string) {
	o.VersionApp = &v
}

// GetVersionBlock returns the VersionBlock field value if set, zero value otherwise.
func (o *ResponsesBlock) GetVersionBlock() string {
	if o == nil || IsNil(o.VersionBlock) {
		var ret string
		return ret
	}
	return *o.VersionBlock
}

// GetVersionBlockOk returns a tuple with the VersionBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponsesBlock) GetVersionBlockOk() (*string, bool) {
	if o == nil || IsNil(o.VersionBlock) {
		return nil, false
	}
	return o.VersionBlock, true
}

// HasVersionBlock returns a boolean if a field has been set.
func (o *ResponsesBlock) HasVersionBlock() bool {
	if o != nil && !IsNil(o.VersionBlock) {
		return true
	}

	return false
}

// SetVersionBlock gets a reference to the given string and assigns it to the VersionBlock field.
func (o *ResponsesBlock) SetVersionBlock(v string) {
	o.VersionBlock = &v
}

func (o ResponsesBlock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponsesBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppHash) {
		toSerialize["app_hash"] = o.AppHash
	}
	if !IsNil(o.ConsensusHash) {
		toSerialize["consensus_hash"] = o.ConsensusHash
	}
	if !IsNil(o.DataHash) {
		toSerialize["data_hash"] = o.DataHash
	}
	if !IsNil(o.EvidenceHash) {
		toSerialize["evidence_hash"] = o.EvidenceHash
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastCommitHash) {
		toSerialize["last_commit_hash"] = o.LastCommitHash
	}
	if !IsNil(o.LastResultsHash) {
		toSerialize["last_results_hash"] = o.LastResultsHash
	}
	if !IsNil(o.MessageTypes) {
		toSerialize["message_types"] = o.MessageTypes
	}
	if !IsNil(o.NextValidatorsHash) {
		toSerialize["next_validators_hash"] = o.NextValidatorsHash
	}
	if !IsNil(o.ParentHash) {
		toSerialize["parent_hash"] = o.ParentHash
	}
	if !IsNil(o.Proposer) {
		toSerialize["proposer"] = o.Proposer
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ValidatorsHash) {
		toSerialize["validators_hash"] = o.ValidatorsHash
	}
	if !IsNil(o.VersionApp) {
		toSerialize["version_app"] = o.VersionApp
	}
	if !IsNil(o.VersionBlock) {
		toSerialize["version_block"] = o.VersionBlock
	}
	return toSerialize, nil
}

type NullableResponsesBlock struct {
	value *ResponsesBlock
	isSet bool
}

func (v NullableResponsesBlock) Get() *ResponsesBlock {
	return v.value
}

func (v *NullableResponsesBlock) Set(val *ResponsesBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableResponsesBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableResponsesBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponsesBlock(val *ResponsesBlock) *NullableResponsesBlock {
	return &NullableResponsesBlock{value: val, isSet: true}
}

func (v NullableResponsesBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponsesBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


