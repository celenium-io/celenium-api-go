# ResponsesState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**LastHeight** | Pointer to **int64** |  | [optional] 
**LastTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Synced** | Pointer to **bool** |  | [optional] 
**TotalAccounts** | Pointer to **int64** |  | [optional] 
**TotalBlobsSize** | Pointer to **int64** |  | [optional] 
**TotalFee** | Pointer to **string** |  | [optional] 
**TotalStake** | Pointer to **string** |  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**TotalTx** | Pointer to **int64** |  | [optional] 
**TotalValidators** | Pointer to **int64** |  | [optional] 
**TotalVotingPower** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesState

`func NewResponsesState() *ResponsesState`

NewResponsesState instantiates a new ResponsesState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesStateWithDefaults

`func NewResponsesStateWithDefaults() *ResponsesState`

NewResponsesStateWithDefaults instantiates a new ResponsesState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ResponsesState) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ResponsesState) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ResponsesState) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ResponsesState) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetHash

`func (o *ResponsesState) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponsesState) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponsesState) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponsesState) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ResponsesState) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesState) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesState) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesState) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastHeight

`func (o *ResponsesState) GetLastHeight() int64`

GetLastHeight returns the LastHeight field if non-nil, zero value otherwise.

### GetLastHeightOk

`func (o *ResponsesState) GetLastHeightOk() (*int64, bool)`

GetLastHeightOk returns a tuple with the LastHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeight

`func (o *ResponsesState) SetLastHeight(v int64)`

SetLastHeight sets LastHeight field to given value.

### HasLastHeight

`func (o *ResponsesState) HasLastHeight() bool`

HasLastHeight returns a boolean if a field has been set.

### GetLastTime

`func (o *ResponsesState) GetLastTime() time.Time`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *ResponsesState) GetLastTimeOk() (*time.Time, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *ResponsesState) SetLastTime(v time.Time)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *ResponsesState) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetName

`func (o *ResponsesState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResponsesState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResponsesState) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResponsesState) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSynced

`func (o *ResponsesState) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *ResponsesState) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *ResponsesState) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *ResponsesState) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetTotalAccounts

`func (o *ResponsesState) GetTotalAccounts() int64`

GetTotalAccounts returns the TotalAccounts field if non-nil, zero value otherwise.

### GetTotalAccountsOk

`func (o *ResponsesState) GetTotalAccountsOk() (*int64, bool)`

GetTotalAccountsOk returns a tuple with the TotalAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAccounts

`func (o *ResponsesState) SetTotalAccounts(v int64)`

SetTotalAccounts sets TotalAccounts field to given value.

### HasTotalAccounts

`func (o *ResponsesState) HasTotalAccounts() bool`

HasTotalAccounts returns a boolean if a field has been set.

### GetTotalBlobsSize

`func (o *ResponsesState) GetTotalBlobsSize() int64`

GetTotalBlobsSize returns the TotalBlobsSize field if non-nil, zero value otherwise.

### GetTotalBlobsSizeOk

`func (o *ResponsesState) GetTotalBlobsSizeOk() (*int64, bool)`

GetTotalBlobsSizeOk returns a tuple with the TotalBlobsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBlobsSize

`func (o *ResponsesState) SetTotalBlobsSize(v int64)`

SetTotalBlobsSize sets TotalBlobsSize field to given value.

### HasTotalBlobsSize

`func (o *ResponsesState) HasTotalBlobsSize() bool`

HasTotalBlobsSize returns a boolean if a field has been set.

### GetTotalFee

`func (o *ResponsesState) GetTotalFee() string`

GetTotalFee returns the TotalFee field if non-nil, zero value otherwise.

### GetTotalFeeOk

`func (o *ResponsesState) GetTotalFeeOk() (*string, bool)`

GetTotalFeeOk returns a tuple with the TotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFee

`func (o *ResponsesState) SetTotalFee(v string)`

SetTotalFee sets TotalFee field to given value.

### HasTotalFee

`func (o *ResponsesState) HasTotalFee() bool`

HasTotalFee returns a boolean if a field has been set.

### GetTotalStake

`func (o *ResponsesState) GetTotalStake() string`

GetTotalStake returns the TotalStake field if non-nil, zero value otherwise.

### GetTotalStakeOk

`func (o *ResponsesState) GetTotalStakeOk() (*string, bool)`

GetTotalStakeOk returns a tuple with the TotalStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStake

`func (o *ResponsesState) SetTotalStake(v string)`

SetTotalStake sets TotalStake field to given value.

### HasTotalStake

`func (o *ResponsesState) HasTotalStake() bool`

HasTotalStake returns a boolean if a field has been set.

### GetTotalSupply

`func (o *ResponsesState) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *ResponsesState) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *ResponsesState) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *ResponsesState) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetTotalTx

`func (o *ResponsesState) GetTotalTx() int64`

GetTotalTx returns the TotalTx field if non-nil, zero value otherwise.

### GetTotalTxOk

`func (o *ResponsesState) GetTotalTxOk() (*int64, bool)`

GetTotalTxOk returns a tuple with the TotalTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTx

`func (o *ResponsesState) SetTotalTx(v int64)`

SetTotalTx sets TotalTx field to given value.

### HasTotalTx

`func (o *ResponsesState) HasTotalTx() bool`

HasTotalTx returns a boolean if a field has been set.

### GetTotalValidators

`func (o *ResponsesState) GetTotalValidators() int64`

GetTotalValidators returns the TotalValidators field if non-nil, zero value otherwise.

### GetTotalValidatorsOk

`func (o *ResponsesState) GetTotalValidatorsOk() (*int64, bool)`

GetTotalValidatorsOk returns a tuple with the TotalValidators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalValidators

`func (o *ResponsesState) SetTotalValidators(v int64)`

SetTotalValidators sets TotalValidators field to given value.

### HasTotalValidators

`func (o *ResponsesState) HasTotalValidators() bool`

HasTotalValidators returns a boolean if a field has been set.

### GetTotalVotingPower

`func (o *ResponsesState) GetTotalVotingPower() string`

GetTotalVotingPower returns the TotalVotingPower field if non-nil, zero value otherwise.

### GetTotalVotingPowerOk

`func (o *ResponsesState) GetTotalVotingPowerOk() (*string, bool)`

GetTotalVotingPowerOk returns a tuple with the TotalVotingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVotingPower

`func (o *ResponsesState) SetTotalVotingPower(v string)`

SetTotalVotingPower sets TotalVotingPower field to given value.

### HasTotalVotingPower

`func (o *ResponsesState) HasTotalVotingPower() bool`

HasTotalVotingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


