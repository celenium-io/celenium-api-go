# ResponsesProposal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstain** | Pointer to **int32** |  | [optional] 
**AbstainAddrs** | Pointer to **int32** |  | [optional] 
**AbstainVals** | Pointer to **int32** |  | [optional] 
**AbstainVotingPower** | Pointer to **string** |  | [optional] 
**ActivationTime** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Deposit** | Pointer to **string** |  | [optional] 
**DepositTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Metadata** | Pointer to **string** |  | [optional] 
**MinDeposit** | Pointer to **string** |  | [optional] 
**No** | Pointer to **int32** |  | [optional] 
**NoAddrs** | Pointer to **int32** |  | [optional] 
**NoVals** | Pointer to **int32** |  | [optional] 
**NoVotingPower** | Pointer to **string** |  | [optional] 
**NoWithVeto** | Pointer to **int32** |  | [optional] 
**NoWithVetoAddrs** | Pointer to **int32** |  | [optional] 
**NoWithVetoVals** | Pointer to **int32** |  | [optional] 
**NoWithVetoVotingPower** | Pointer to **string** |  | [optional] 
**Proposer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Quorum** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TotalVotingPower** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VetoQuorum** | Pointer to **string** |  | [optional] 
**VotesCount** | Pointer to **int32** |  | [optional] 
**VotingPower** | Pointer to **string** |  | [optional] 
**Yes** | Pointer to **int32** |  | [optional] 
**YesAddrs** | Pointer to **int32** |  | [optional] 
**YesVals** | Pointer to **int32** |  | [optional] 
**YesVotingPower** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesProposal

`func NewResponsesProposal() *ResponsesProposal`

NewResponsesProposal instantiates a new ResponsesProposal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesProposalWithDefaults

`func NewResponsesProposalWithDefaults() *ResponsesProposal`

NewResponsesProposalWithDefaults instantiates a new ResponsesProposal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstain

`func (o *ResponsesProposal) GetAbstain() int32`

GetAbstain returns the Abstain field if non-nil, zero value otherwise.

### GetAbstainOk

`func (o *ResponsesProposal) GetAbstainOk() (*int32, bool)`

GetAbstainOk returns a tuple with the Abstain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstain

`func (o *ResponsesProposal) SetAbstain(v int32)`

SetAbstain sets Abstain field to given value.

### HasAbstain

`func (o *ResponsesProposal) HasAbstain() bool`

HasAbstain returns a boolean if a field has been set.

### GetAbstainAddrs

`func (o *ResponsesProposal) GetAbstainAddrs() int32`

GetAbstainAddrs returns the AbstainAddrs field if non-nil, zero value otherwise.

### GetAbstainAddrsOk

`func (o *ResponsesProposal) GetAbstainAddrsOk() (*int32, bool)`

GetAbstainAddrsOk returns a tuple with the AbstainAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstainAddrs

`func (o *ResponsesProposal) SetAbstainAddrs(v int32)`

SetAbstainAddrs sets AbstainAddrs field to given value.

### HasAbstainAddrs

`func (o *ResponsesProposal) HasAbstainAddrs() bool`

HasAbstainAddrs returns a boolean if a field has been set.

### GetAbstainVals

`func (o *ResponsesProposal) GetAbstainVals() int32`

GetAbstainVals returns the AbstainVals field if non-nil, zero value otherwise.

### GetAbstainValsOk

`func (o *ResponsesProposal) GetAbstainValsOk() (*int32, bool)`

GetAbstainValsOk returns a tuple with the AbstainVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstainVals

`func (o *ResponsesProposal) SetAbstainVals(v int32)`

SetAbstainVals sets AbstainVals field to given value.

### HasAbstainVals

`func (o *ResponsesProposal) HasAbstainVals() bool`

HasAbstainVals returns a boolean if a field has been set.

### GetAbstainVotingPower

`func (o *ResponsesProposal) GetAbstainVotingPower() string`

GetAbstainVotingPower returns the AbstainVotingPower field if non-nil, zero value otherwise.

### GetAbstainVotingPowerOk

`func (o *ResponsesProposal) GetAbstainVotingPowerOk() (*string, bool)`

GetAbstainVotingPowerOk returns a tuple with the AbstainVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstainVotingPower

`func (o *ResponsesProposal) SetAbstainVotingPower(v string)`

SetAbstainVotingPower sets AbstainVotingPower field to given value.

### HasAbstainVotingPower

`func (o *ResponsesProposal) HasAbstainVotingPower() bool`

HasAbstainVotingPower returns a boolean if a field has been set.

### GetActivationTime

`func (o *ResponsesProposal) GetActivationTime() time.Time`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *ResponsesProposal) GetActivationTimeOk() (*time.Time, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *ResponsesProposal) SetActivationTime(v time.Time)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *ResponsesProposal) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponsesProposal) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponsesProposal) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponsesProposal) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponsesProposal) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeposit

`func (o *ResponsesProposal) GetDeposit() string`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *ResponsesProposal) GetDepositOk() (*string, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *ResponsesProposal) SetDeposit(v string)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *ResponsesProposal) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDepositTime

`func (o *ResponsesProposal) GetDepositTime() time.Time`

GetDepositTime returns the DepositTime field if non-nil, zero value otherwise.

### GetDepositTimeOk

`func (o *ResponsesProposal) GetDepositTimeOk() (*time.Time, bool)`

GetDepositTimeOk returns a tuple with the DepositTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositTime

`func (o *ResponsesProposal) SetDepositTime(v time.Time)`

SetDepositTime sets DepositTime field to given value.

### HasDepositTime

`func (o *ResponsesProposal) HasDepositTime() bool`

HasDepositTime returns a boolean if a field has been set.

### GetDescription

`func (o *ResponsesProposal) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResponsesProposal) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResponsesProposal) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResponsesProposal) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndTime

`func (o *ResponsesProposal) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResponsesProposal) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResponsesProposal) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResponsesProposal) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetError

`func (o *ResponsesProposal) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponsesProposal) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponsesProposal) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponsesProposal) HasError() bool`

HasError returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesProposal) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesProposal) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesProposal) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesProposal) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesProposal) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesProposal) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesProposal) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesProposal) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ResponsesProposal) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResponsesProposal) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResponsesProposal) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResponsesProposal) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinDeposit

`func (o *ResponsesProposal) GetMinDeposit() string`

GetMinDeposit returns the MinDeposit field if non-nil, zero value otherwise.

### GetMinDepositOk

`func (o *ResponsesProposal) GetMinDepositOk() (*string, bool)`

GetMinDepositOk returns a tuple with the MinDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDeposit

`func (o *ResponsesProposal) SetMinDeposit(v string)`

SetMinDeposit sets MinDeposit field to given value.

### HasMinDeposit

`func (o *ResponsesProposal) HasMinDeposit() bool`

HasMinDeposit returns a boolean if a field has been set.

### GetNo

`func (o *ResponsesProposal) GetNo() int32`

GetNo returns the No field if non-nil, zero value otherwise.

### GetNoOk

`func (o *ResponsesProposal) GetNoOk() (*int32, bool)`

GetNoOk returns a tuple with the No field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNo

`func (o *ResponsesProposal) SetNo(v int32)`

SetNo sets No field to given value.

### HasNo

`func (o *ResponsesProposal) HasNo() bool`

HasNo returns a boolean if a field has been set.

### GetNoAddrs

`func (o *ResponsesProposal) GetNoAddrs() int32`

GetNoAddrs returns the NoAddrs field if non-nil, zero value otherwise.

### GetNoAddrsOk

`func (o *ResponsesProposal) GetNoAddrsOk() (*int32, bool)`

GetNoAddrsOk returns a tuple with the NoAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAddrs

`func (o *ResponsesProposal) SetNoAddrs(v int32)`

SetNoAddrs sets NoAddrs field to given value.

### HasNoAddrs

`func (o *ResponsesProposal) HasNoAddrs() bool`

HasNoAddrs returns a boolean if a field has been set.

### GetNoVals

`func (o *ResponsesProposal) GetNoVals() int32`

GetNoVals returns the NoVals field if non-nil, zero value otherwise.

### GetNoValsOk

`func (o *ResponsesProposal) GetNoValsOk() (*int32, bool)`

GetNoValsOk returns a tuple with the NoVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVals

`func (o *ResponsesProposal) SetNoVals(v int32)`

SetNoVals sets NoVals field to given value.

### HasNoVals

`func (o *ResponsesProposal) HasNoVals() bool`

HasNoVals returns a boolean if a field has been set.

### GetNoVotingPower

`func (o *ResponsesProposal) GetNoVotingPower() string`

GetNoVotingPower returns the NoVotingPower field if non-nil, zero value otherwise.

### GetNoVotingPowerOk

`func (o *ResponsesProposal) GetNoVotingPowerOk() (*string, bool)`

GetNoVotingPowerOk returns a tuple with the NoVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoVotingPower

`func (o *ResponsesProposal) SetNoVotingPower(v string)`

SetNoVotingPower sets NoVotingPower field to given value.

### HasNoVotingPower

`func (o *ResponsesProposal) HasNoVotingPower() bool`

HasNoVotingPower returns a boolean if a field has been set.

### GetNoWithVeto

`func (o *ResponsesProposal) GetNoWithVeto() int32`

GetNoWithVeto returns the NoWithVeto field if non-nil, zero value otherwise.

### GetNoWithVetoOk

`func (o *ResponsesProposal) GetNoWithVetoOk() (*int32, bool)`

GetNoWithVetoOk returns a tuple with the NoWithVeto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWithVeto

`func (o *ResponsesProposal) SetNoWithVeto(v int32)`

SetNoWithVeto sets NoWithVeto field to given value.

### HasNoWithVeto

`func (o *ResponsesProposal) HasNoWithVeto() bool`

HasNoWithVeto returns a boolean if a field has been set.

### GetNoWithVetoAddrs

`func (o *ResponsesProposal) GetNoWithVetoAddrs() int32`

GetNoWithVetoAddrs returns the NoWithVetoAddrs field if non-nil, zero value otherwise.

### GetNoWithVetoAddrsOk

`func (o *ResponsesProposal) GetNoWithVetoAddrsOk() (*int32, bool)`

GetNoWithVetoAddrsOk returns a tuple with the NoWithVetoAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWithVetoAddrs

`func (o *ResponsesProposal) SetNoWithVetoAddrs(v int32)`

SetNoWithVetoAddrs sets NoWithVetoAddrs field to given value.

### HasNoWithVetoAddrs

`func (o *ResponsesProposal) HasNoWithVetoAddrs() bool`

HasNoWithVetoAddrs returns a boolean if a field has been set.

### GetNoWithVetoVals

`func (o *ResponsesProposal) GetNoWithVetoVals() int32`

GetNoWithVetoVals returns the NoWithVetoVals field if non-nil, zero value otherwise.

### GetNoWithVetoValsOk

`func (o *ResponsesProposal) GetNoWithVetoValsOk() (*int32, bool)`

GetNoWithVetoValsOk returns a tuple with the NoWithVetoVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWithVetoVals

`func (o *ResponsesProposal) SetNoWithVetoVals(v int32)`

SetNoWithVetoVals sets NoWithVetoVals field to given value.

### HasNoWithVetoVals

`func (o *ResponsesProposal) HasNoWithVetoVals() bool`

HasNoWithVetoVals returns a boolean if a field has been set.

### GetNoWithVetoVotingPower

`func (o *ResponsesProposal) GetNoWithVetoVotingPower() string`

GetNoWithVetoVotingPower returns the NoWithVetoVotingPower field if non-nil, zero value otherwise.

### GetNoWithVetoVotingPowerOk

`func (o *ResponsesProposal) GetNoWithVetoVotingPowerOk() (*string, bool)`

GetNoWithVetoVotingPowerOk returns a tuple with the NoWithVetoVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWithVetoVotingPower

`func (o *ResponsesProposal) SetNoWithVetoVotingPower(v string)`

SetNoWithVetoVotingPower sets NoWithVetoVotingPower field to given value.

### HasNoWithVetoVotingPower

`func (o *ResponsesProposal) HasNoWithVetoVotingPower() bool`

HasNoWithVetoVotingPower returns a boolean if a field has been set.

### GetProposer

`func (o *ResponsesProposal) GetProposer() ResponsesShortAddress`

GetProposer returns the Proposer field if non-nil, zero value otherwise.

### GetProposerOk

`func (o *ResponsesProposal) GetProposerOk() (*ResponsesShortAddress, bool)`

GetProposerOk returns a tuple with the Proposer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposer

`func (o *ResponsesProposal) SetProposer(v ResponsesShortAddress)`

SetProposer sets Proposer field to given value.

### HasProposer

`func (o *ResponsesProposal) HasProposer() bool`

HasProposer returns a boolean if a field has been set.

### GetQuorum

`func (o *ResponsesProposal) GetQuorum() string`

GetQuorum returns the Quorum field if non-nil, zero value otherwise.

### GetQuorumOk

`func (o *ResponsesProposal) GetQuorumOk() (*string, bool)`

GetQuorumOk returns a tuple with the Quorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuorum

`func (o *ResponsesProposal) SetQuorum(v string)`

SetQuorum sets Quorum field to given value.

### HasQuorum

`func (o *ResponsesProposal) HasQuorum() bool`

HasQuorum returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesProposal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesProposal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesProposal) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesProposal) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreshold

`func (o *ResponsesProposal) GetThreshold() string`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ResponsesProposal) GetThresholdOk() (*string, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ResponsesProposal) SetThreshold(v string)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ResponsesProposal) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTitle

`func (o *ResponsesProposal) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResponsesProposal) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResponsesProposal) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResponsesProposal) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTotalVotingPower

`func (o *ResponsesProposal) GetTotalVotingPower() string`

GetTotalVotingPower returns the TotalVotingPower field if non-nil, zero value otherwise.

### GetTotalVotingPowerOk

`func (o *ResponsesProposal) GetTotalVotingPowerOk() (*string, bool)`

GetTotalVotingPowerOk returns a tuple with the TotalVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVotingPower

`func (o *ResponsesProposal) SetTotalVotingPower(v string)`

SetTotalVotingPower sets TotalVotingPower field to given value.

### HasTotalVotingPower

`func (o *ResponsesProposal) HasTotalVotingPower() bool`

HasTotalVotingPower returns a boolean if a field has been set.

### GetType

`func (o *ResponsesProposal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesProposal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesProposal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesProposal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVetoQuorum

`func (o *ResponsesProposal) GetVetoQuorum() string`

GetVetoQuorum returns the VetoQuorum field if non-nil, zero value otherwise.

### GetVetoQuorumOk

`func (o *ResponsesProposal) GetVetoQuorumOk() (*string, bool)`

GetVetoQuorumOk returns a tuple with the VetoQuorum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVetoQuorum

`func (o *ResponsesProposal) SetVetoQuorum(v string)`

SetVetoQuorum sets VetoQuorum field to given value.

### HasVetoQuorum

`func (o *ResponsesProposal) HasVetoQuorum() bool`

HasVetoQuorum returns a boolean if a field has been set.

### GetVotesCount

`func (o *ResponsesProposal) GetVotesCount() int32`

GetVotesCount returns the VotesCount field if non-nil, zero value otherwise.

### GetVotesCountOk

`func (o *ResponsesProposal) GetVotesCountOk() (*int32, bool)`

GetVotesCountOk returns a tuple with the VotesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotesCount

`func (o *ResponsesProposal) SetVotesCount(v int32)`

SetVotesCount sets VotesCount field to given value.

### HasVotesCount

`func (o *ResponsesProposal) HasVotesCount() bool`

HasVotesCount returns a boolean if a field has been set.

### GetVotingPower

`func (o *ResponsesProposal) GetVotingPower() string`

GetVotingPower returns the VotingPower field if non-nil, zero value otherwise.

### GetVotingPowerOk

`func (o *ResponsesProposal) GetVotingPowerOk() (*string, bool)`

GetVotingPowerOk returns a tuple with the VotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingPower

`func (o *ResponsesProposal) SetVotingPower(v string)`

SetVotingPower sets VotingPower field to given value.

### HasVotingPower

`func (o *ResponsesProposal) HasVotingPower() bool`

HasVotingPower returns a boolean if a field has been set.

### GetYes

`func (o *ResponsesProposal) GetYes() int32`

GetYes returns the Yes field if non-nil, zero value otherwise.

### GetYesOk

`func (o *ResponsesProposal) GetYesOk() (*int32, bool)`

GetYesOk returns a tuple with the Yes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYes

`func (o *ResponsesProposal) SetYes(v int32)`

SetYes sets Yes field to given value.

### HasYes

`func (o *ResponsesProposal) HasYes() bool`

HasYes returns a boolean if a field has been set.

### GetYesAddrs

`func (o *ResponsesProposal) GetYesAddrs() int32`

GetYesAddrs returns the YesAddrs field if non-nil, zero value otherwise.

### GetYesAddrsOk

`func (o *ResponsesProposal) GetYesAddrsOk() (*int32, bool)`

GetYesAddrsOk returns a tuple with the YesAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesAddrs

`func (o *ResponsesProposal) SetYesAddrs(v int32)`

SetYesAddrs sets YesAddrs field to given value.

### HasYesAddrs

`func (o *ResponsesProposal) HasYesAddrs() bool`

HasYesAddrs returns a boolean if a field has been set.

### GetYesVals

`func (o *ResponsesProposal) GetYesVals() int32`

GetYesVals returns the YesVals field if non-nil, zero value otherwise.

### GetYesValsOk

`func (o *ResponsesProposal) GetYesValsOk() (*int32, bool)`

GetYesValsOk returns a tuple with the YesVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesVals

`func (o *ResponsesProposal) SetYesVals(v int32)`

SetYesVals sets YesVals field to given value.

### HasYesVals

`func (o *ResponsesProposal) HasYesVals() bool`

HasYesVals returns a boolean if a field has been set.

### GetYesVotingPower

`func (o *ResponsesProposal) GetYesVotingPower() string`

GetYesVotingPower returns the YesVotingPower field if non-nil, zero value otherwise.

### GetYesVotingPowerOk

`func (o *ResponsesProposal) GetYesVotingPowerOk() (*string, bool)`

GetYesVotingPowerOk returns a tuple with the YesVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYesVotingPower

`func (o *ResponsesProposal) SetYesVotingPower(v string)`

SetYesVotingPower sets YesVotingPower field to given value.

### HasYesVotingPower

`func (o *ResponsesProposal) HasYesVotingPower() bool`

HasYesVotingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


