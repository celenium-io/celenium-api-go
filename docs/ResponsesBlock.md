# ResponsesBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppHash** | Pointer to **string** |  | [optional] 
**ConsensusHash** | Pointer to **string** |  | [optional] 
**DataHash** | Pointer to **string** |  | [optional] 
**EvidenceHash** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**LastCommitHash** | Pointer to **string** |  | [optional] 
**LastResultsHash** | Pointer to **string** |  | [optional] 
**MessageTypes** | Pointer to **[]string** |  | [optional] 
**NextValidatorsHash** | Pointer to **string** |  | [optional] 
**ParentHash** | Pointer to **string** |  | [optional] 
**Proposer** | Pointer to [**ResponsesShortValidator**](ResponsesShortValidator.md) |  | [optional] 
**Stats** | Pointer to [**ResponsesBlockStats**](ResponsesBlockStats.md) |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**ValidatorsHash** | Pointer to **string** |  | [optional] 
**VersionApp** | Pointer to **string** |  | [optional] 
**VersionBlock** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesBlock

`func NewResponsesBlock() *ResponsesBlock`

NewResponsesBlock instantiates a new ResponsesBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesBlockWithDefaults

`func NewResponsesBlockWithDefaults() *ResponsesBlock`

NewResponsesBlockWithDefaults instantiates a new ResponsesBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppHash

`func (o *ResponsesBlock) GetAppHash() string`

GetAppHash returns the AppHash field if non-nil, zero value otherwise.

### GetAppHashOk

`func (o *ResponsesBlock) GetAppHashOk() (*string, bool)`

GetAppHashOk returns a tuple with the AppHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppHash

`func (o *ResponsesBlock) SetAppHash(v string)`

SetAppHash sets AppHash field to given value.

### HasAppHash

`func (o *ResponsesBlock) HasAppHash() bool`

HasAppHash returns a boolean if a field has been set.

### GetConsensusHash

`func (o *ResponsesBlock) GetConsensusHash() string`

GetConsensusHash returns the ConsensusHash field if non-nil, zero value otherwise.

### GetConsensusHashOk

`func (o *ResponsesBlock) GetConsensusHashOk() (*string, bool)`

GetConsensusHashOk returns a tuple with the ConsensusHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsensusHash

`func (o *ResponsesBlock) SetConsensusHash(v string)`

SetConsensusHash sets ConsensusHash field to given value.

### HasConsensusHash

`func (o *ResponsesBlock) HasConsensusHash() bool`

HasConsensusHash returns a boolean if a field has been set.

### GetDataHash

`func (o *ResponsesBlock) GetDataHash() string`

GetDataHash returns the DataHash field if non-nil, zero value otherwise.

### GetDataHashOk

`func (o *ResponsesBlock) GetDataHashOk() (*string, bool)`

GetDataHashOk returns a tuple with the DataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHash

`func (o *ResponsesBlock) SetDataHash(v string)`

SetDataHash sets DataHash field to given value.

### HasDataHash

`func (o *ResponsesBlock) HasDataHash() bool`

HasDataHash returns a boolean if a field has been set.

### GetEvidenceHash

`func (o *ResponsesBlock) GetEvidenceHash() string`

GetEvidenceHash returns the EvidenceHash field if non-nil, zero value otherwise.

### GetEvidenceHashOk

`func (o *ResponsesBlock) GetEvidenceHashOk() (*string, bool)`

GetEvidenceHashOk returns a tuple with the EvidenceHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidenceHash

`func (o *ResponsesBlock) SetEvidenceHash(v string)`

SetEvidenceHash sets EvidenceHash field to given value.

### HasEvidenceHash

`func (o *ResponsesBlock) HasEvidenceHash() bool`

HasEvidenceHash returns a boolean if a field has been set.

### GetHash

`func (o *ResponsesBlock) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponsesBlock) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponsesBlock) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponsesBlock) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesBlock) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesBlock) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesBlock) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesBlock) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesBlock) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesBlock) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesBlock) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastCommitHash

`func (o *ResponsesBlock) GetLastCommitHash() string`

GetLastCommitHash returns the LastCommitHash field if non-nil, zero value otherwise.

### GetLastCommitHashOk

`func (o *ResponsesBlock) GetLastCommitHashOk() (*string, bool)`

GetLastCommitHashOk returns a tuple with the LastCommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitHash

`func (o *ResponsesBlock) SetLastCommitHash(v string)`

SetLastCommitHash sets LastCommitHash field to given value.

### HasLastCommitHash

`func (o *ResponsesBlock) HasLastCommitHash() bool`

HasLastCommitHash returns a boolean if a field has been set.

### GetLastResultsHash

`func (o *ResponsesBlock) GetLastResultsHash() string`

GetLastResultsHash returns the LastResultsHash field if non-nil, zero value otherwise.

### GetLastResultsHashOk

`func (o *ResponsesBlock) GetLastResultsHashOk() (*string, bool)`

GetLastResultsHashOk returns a tuple with the LastResultsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResultsHash

`func (o *ResponsesBlock) SetLastResultsHash(v string)`

SetLastResultsHash sets LastResultsHash field to given value.

### HasLastResultsHash

`func (o *ResponsesBlock) HasLastResultsHash() bool`

HasLastResultsHash returns a boolean if a field has been set.

### GetMessageTypes

`func (o *ResponsesBlock) GetMessageTypes() []string`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *ResponsesBlock) GetMessageTypesOk() (*[]string, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *ResponsesBlock) SetMessageTypes(v []string)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *ResponsesBlock) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### GetNextValidatorsHash

`func (o *ResponsesBlock) GetNextValidatorsHash() string`

GetNextValidatorsHash returns the NextValidatorsHash field if non-nil, zero value otherwise.

### GetNextValidatorsHashOk

`func (o *ResponsesBlock) GetNextValidatorsHashOk() (*string, bool)`

GetNextValidatorsHashOk returns a tuple with the NextValidatorsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextValidatorsHash

`func (o *ResponsesBlock) SetNextValidatorsHash(v string)`

SetNextValidatorsHash sets NextValidatorsHash field to given value.

### HasNextValidatorsHash

`func (o *ResponsesBlock) HasNextValidatorsHash() bool`

HasNextValidatorsHash returns a boolean if a field has been set.

### GetParentHash

`func (o *ResponsesBlock) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *ResponsesBlock) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *ResponsesBlock) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.

### HasParentHash

`func (o *ResponsesBlock) HasParentHash() bool`

HasParentHash returns a boolean if a field has been set.

### GetProposer

`func (o *ResponsesBlock) GetProposer() ResponsesShortValidator`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *ResponsesBlock) GetProposerOk() (*ResponsesShortValidator, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *ResponsesBlock) SetProposer(v ResponsesShortValidator)`

SetProposer sets Proposer field to given value.

### HasProposer

`func (o *ResponsesBlock) HasProposer() bool`

HasProposer returns a boolean if a field has been set.

### GetStats

`func (o *ResponsesBlock) GetStats() ResponsesBlockStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ResponsesBlock) GetStatsOk() (*ResponsesBlockStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ResponsesBlock) SetStats(v ResponsesBlockStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ResponsesBlock) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesBlock) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesBlock) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesBlock) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesBlock) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetValidatorsHash

`func (o *ResponsesBlock) GetValidatorsHash() string`

GetValidatorsHash returns the ValidatorsHash field if non-nil, zero value otherwise.

### GetValidatorsHashOk

`func (o *ResponsesBlock) GetValidatorsHashOk() (*string, bool)`

GetValidatorsHashOk returns a tuple with the ValidatorsHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorsHash

`func (o *ResponsesBlock) SetValidatorsHash(v string)`

SetValidatorsHash sets ValidatorsHash field to given value.

### HasValidatorsHash

`func (o *ResponsesBlock) HasValidatorsHash() bool`

HasValidatorsHash returns a boolean if a field has been set.

### GetVersionApp

`func (o *ResponsesBlock) GetVersionApp() string`

GetVersionApp returns the VersionApp field if non-nil, zero value otherwise.

### GetVersionAppOk

`func (o *ResponsesBlock) GetVersionAppOk() (*string, bool)`

GetVersionAppOk returns a tuple with the VersionApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionApp

`func (o *ResponsesBlock) SetVersionApp(v string)`

SetVersionApp sets VersionApp field to given value.

### HasVersionApp

`func (o *ResponsesBlock) HasVersionApp() bool`

HasVersionApp returns a boolean if a field has been set.

### GetVersionBlock

`func (o *ResponsesBlock) GetVersionBlock() string`

GetVersionBlock returns the VersionBlock field if non-nil, zero value otherwise.

### GetVersionBlockOk

`func (o *ResponsesBlock) GetVersionBlockOk() (*string, bool)`

GetVersionBlockOk returns a tuple with the VersionBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionBlock

`func (o *ResponsesBlock) SetVersionBlock(v string)`

SetVersionBlock sets VersionBlock field to given value.

### HasVersionBlock

`func (o *ResponsesBlock) HasVersionBlock() bool`

HasVersionBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


