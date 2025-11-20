# ResponsesUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedAt** | Pointer to **time.Time** |  | [optional] 
**AppliedAtLevel** | Pointer to **int64** |  | [optional] 
**EndHeight** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**MsgId** | Pointer to **int64** |  | [optional] 
**SignalsCount** | Pointer to **int64** |  | [optional] 
**Signer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**VotedPower** | Pointer to **string** |  | [optional] 
**VotingPower** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesUpgrade

`func NewResponsesUpgrade() *ResponsesUpgrade`

NewResponsesUpgrade instantiates a new ResponsesUpgrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesUpgradeWithDefaults

`func NewResponsesUpgradeWithDefaults() *ResponsesUpgrade`

NewResponsesUpgradeWithDefaults instantiates a new ResponsesUpgrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedAt

`func (o *ResponsesUpgrade) GetAppliedAt() time.Time`

GetAppliedAt returns the AppliedAt field if non-nil, zero value otherwise.

### GetAppliedAtOk

`func (o *ResponsesUpgrade) GetAppliedAtOk() (*time.Time, bool)`

GetAppliedAtOk returns a tuple with the AppliedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAt

`func (o *ResponsesUpgrade) SetAppliedAt(v time.Time)`

SetAppliedAt sets AppliedAt field to given value.

### HasAppliedAt

`func (o *ResponsesUpgrade) HasAppliedAt() bool`

HasAppliedAt returns a boolean if a field has been set.

### GetAppliedAtLevel

`func (o *ResponsesUpgrade) GetAppliedAtLevel() int64`

GetAppliedAtLevel returns the AppliedAtLevel field if non-nil, zero value otherwise.

### GetAppliedAtLevelOk

`func (o *ResponsesUpgrade) GetAppliedAtLevelOk() (*int64, bool)`

GetAppliedAtLevelOk returns a tuple with the AppliedAtLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAtLevel

`func (o *ResponsesUpgrade) SetAppliedAtLevel(v int64)`

SetAppliedAtLevel sets AppliedAtLevel field to given value.

### HasAppliedAtLevel

`func (o *ResponsesUpgrade) HasAppliedAtLevel() bool`

HasAppliedAtLevel returns a boolean if a field has been set.

### GetEndHeight

`func (o *ResponsesUpgrade) GetEndHeight() int64`

GetEndHeight returns the EndHeight field if non-nil, zero value otherwise.

### GetEndHeightOk

`func (o *ResponsesUpgrade) GetEndHeightOk() (*int64, bool)`

GetEndHeightOk returns a tuple with the EndHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHeight

`func (o *ResponsesUpgrade) SetEndHeight(v int64)`

SetEndHeight sets EndHeight field to given value.

### HasEndHeight

`func (o *ResponsesUpgrade) HasEndHeight() bool`

HasEndHeight returns a boolean if a field has been set.

### GetEndTime

`func (o *ResponsesUpgrade) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResponsesUpgrade) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResponsesUpgrade) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ResponsesUpgrade) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesUpgrade) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesUpgrade) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesUpgrade) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesUpgrade) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMsgId

`func (o *ResponsesUpgrade) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *ResponsesUpgrade) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *ResponsesUpgrade) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *ResponsesUpgrade) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetSignalsCount

`func (o *ResponsesUpgrade) GetSignalsCount() int64`

GetSignalsCount returns the SignalsCount field if non-nil, zero value otherwise.

### GetSignalsCountOk

`func (o *ResponsesUpgrade) GetSignalsCountOk() (*int64, bool)`

GetSignalsCountOk returns a tuple with the SignalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalsCount

`func (o *ResponsesUpgrade) SetSignalsCount(v int64)`

SetSignalsCount sets SignalsCount field to given value.

### HasSignalsCount

`func (o *ResponsesUpgrade) HasSignalsCount() bool`

HasSignalsCount returns a boolean if a field has been set.

### GetSigner

`func (o *ResponsesUpgrade) GetSigner() ResponsesShortAddress`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ResponsesUpgrade) GetSignerOk() (*ResponsesShortAddress, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ResponsesUpgrade) SetSigner(v ResponsesShortAddress)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ResponsesUpgrade) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesUpgrade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesUpgrade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesUpgrade) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesUpgrade) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesUpgrade) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesUpgrade) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesUpgrade) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesUpgrade) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesUpgrade) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesUpgrade) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesUpgrade) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesUpgrade) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetVersion

`func (o *ResponsesUpgrade) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponsesUpgrade) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponsesUpgrade) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponsesUpgrade) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotedPower

`func (o *ResponsesUpgrade) GetVotedPower() string`

GetVotedPower returns the VotedPower field if non-nil, zero value otherwise.

### GetVotedPowerOk

`func (o *ResponsesUpgrade) GetVotedPowerOk() (*string, bool)`

GetVotedPowerOk returns a tuple with the VotedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotedPower

`func (o *ResponsesUpgrade) SetVotedPower(v string)`

SetVotedPower sets VotedPower field to given value.

### HasVotedPower

`func (o *ResponsesUpgrade) HasVotedPower() bool`

HasVotedPower returns a boolean if a field has been set.

### GetVotingPower

`func (o *ResponsesUpgrade) GetVotingPower() string`

GetVotingPower returns the VotingPower field if non-nil, zero value otherwise.

### GetVotingPowerOk

`func (o *ResponsesUpgrade) GetVotingPowerOk() (*string, bool)`

GetVotingPowerOk returns a tuple with the VotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingPower

`func (o *ResponsesUpgrade) SetVotingPower(v string)`

SetVotingPower sets VotingPower field to given value.

### HasVotingPower

`func (o *ResponsesUpgrade) HasVotingPower() bool`

HasVotingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


