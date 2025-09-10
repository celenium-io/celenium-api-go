# ResponsesSignalVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Validator** | Pointer to [**ResponsesShortValidator**](ResponsesShortValidator.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**VotingPower** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesSignalVersion

`func NewResponsesSignalVersion() *ResponsesSignalVersion`

NewResponsesSignalVersion instantiates a new ResponsesSignalVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesSignalVersionWithDefaults

`func NewResponsesSignalVersionWithDefaults() *ResponsesSignalVersion`

NewResponsesSignalVersionWithDefaults instantiates a new ResponsesSignalVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ResponsesSignalVersion) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesSignalVersion) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesSignalVersion) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesSignalVersion) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesSignalVersion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesSignalVersion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesSignalVersion) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesSignalVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesSignalVersion) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesSignalVersion) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesSignalVersion) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesSignalVersion) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesSignalVersion) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesSignalVersion) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesSignalVersion) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesSignalVersion) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetValidator

`func (o *ResponsesSignalVersion) GetValidator() ResponsesShortValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *ResponsesSignalVersion) GetValidatorOk() (*ResponsesShortValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *ResponsesSignalVersion) SetValidator(v ResponsesShortValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *ResponsesSignalVersion) HasValidator() bool`

HasValidator returns a boolean if a field has been set.

### GetVersion

`func (o *ResponsesSignalVersion) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponsesSignalVersion) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponsesSignalVersion) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponsesSignalVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVotingPower

`func (o *ResponsesSignalVersion) GetVotingPower() string`

GetVotingPower returns the VotingPower field if non-nil, zero value otherwise.

### GetVotingPowerOk

`func (o *ResponsesSignalVersion) GetVotingPowerOk() (*string, bool)`

GetVotingPowerOk returns a tuple with the VotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotingPower

`func (o *ResponsesSignalVersion) SetVotingPower(v string)`

SetVotingPower sets VotingPower field to given value.

### HasVotingPower

`func (o *ResponsesSignalVersion) HasVotingPower() bool`

HasVotingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


