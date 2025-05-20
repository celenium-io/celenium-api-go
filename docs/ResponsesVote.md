# ResponsesVote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositTime** | Pointer to **time.Time** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**ProposalId** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Validator** | Pointer to [**ResponsesShortValidator**](ResponsesShortValidator.md) |  | [optional] 
**Voter** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**VoterId** | Pointer to **int64** |  | [optional] 
**Weight** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponsesVote

`func NewResponsesVote() *ResponsesVote`

NewResponsesVote instantiates a new ResponsesVote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesVoteWithDefaults

`func NewResponsesVoteWithDefaults() *ResponsesVote`

NewResponsesVoteWithDefaults instantiates a new ResponsesVote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositTime

`func (o *ResponsesVote) GetDepositTime() time.Time`

GetDepositTime returns the DepositTime field if non-nil, zero value otherwise.

### GetDepositTimeOk

`func (o *ResponsesVote) GetDepositTimeOk() (*time.Time, bool)`

GetDepositTimeOk returns a tuple with the DepositTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositTime

`func (o *ResponsesVote) SetDepositTime(v time.Time)`

SetDepositTime sets DepositTime field to given value.

### HasDepositTime

`func (o *ResponsesVote) HasDepositTime() bool`

HasDepositTime returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesVote) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesVote) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesVote) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesVote) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesVote) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesVote) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesVote) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesVote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProposalId

`func (o *ResponsesVote) GetProposalId() int64`

GetProposalId returns the ProposalId field if non-nil, zero value otherwise.

### GetProposalIdOk

`func (o *ResponsesVote) GetProposalIdOk() (*int64, bool)`

GetProposalIdOk returns a tuple with the ProposalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposalId

`func (o *ResponsesVote) SetProposalId(v int64)`

SetProposalId sets ProposalId field to given value.

### HasProposalId

`func (o *ResponsesVote) HasProposalId() bool`

HasProposalId returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesVote) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesVote) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesVote) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesVote) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetValidator

`func (o *ResponsesVote) GetValidator() ResponsesShortValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *ResponsesVote) GetValidatorOk() (*ResponsesShortValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *ResponsesVote) SetValidator(v ResponsesShortValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *ResponsesVote) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetVoter

`func (o *ResponsesVote) GetVoter() ResponsesShortAddress`

GetVoter returns the Voter field if non-nil, zero value otherwise.

### GetVoterOk

`func (o *ResponsesVote) GetVoterOk() (*ResponsesShortAddress, bool)`

GetVoterOk returns a tuple with the Voter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoter

`func (o *ResponsesVote) SetVoter(v ResponsesShortAddress)`

SetVoter sets Voter field to given value.

### HasVoter

`func (o *ResponsesVote) HasVoter() bool`

HasVoter returns a boolean if a field has been set.

### GetVoterId

`func (o *ResponsesVote) GetVoterId() int64`

GetVoterId returns the VoterId field if non-nil, zero value otherwise.

### GetVoterIdOk

`func (o *ResponsesVote) GetVoterIdOk() (*int64, bool)`

GetVoterIdOk returns a tuple with the VoterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoterId

`func (o *ResponsesVote) SetVoterId(v int64)`

SetVoterId sets VoterId field to given value.

### HasVoterId

`func (o *ResponsesVote) HasVoterId() bool`

HasVoterId returns a boolean if a field has been set.

### GetWeight

`func (o *ResponsesVote) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ResponsesVote) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ResponsesVote) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ResponsesVote) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


