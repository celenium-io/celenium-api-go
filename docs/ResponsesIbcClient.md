# ResponsesIbcClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChainId** | Pointer to ***os.File** |  | [optional] 
**ConnectionCount** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Creator** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**FrozenRevisionHeight** | Pointer to **int32** |  | [optional] 
**FrozenRevisionNumber** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LatestRevisionHeight** | Pointer to **int32** |  | [optional] 
**LatestRevisionNumber** | Pointer to **int32** |  | [optional] 
**MaxClockDrift** | Pointer to **int32** |  | [optional] 
**TrustLevelDenominator** | Pointer to **int32** |  | [optional] 
**TrustLevelNumerator** | Pointer to **int32** |  | [optional] 
**TrustingPeriod** | Pointer to **int32** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnbondingPeriod** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewResponsesIbcClient

`func NewResponsesIbcClient() *ResponsesIbcClient`

NewResponsesIbcClient instantiates a new ResponsesIbcClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesIbcClientWithDefaults

`func NewResponsesIbcClientWithDefaults() *ResponsesIbcClient`

NewResponsesIbcClientWithDefaults instantiates a new ResponsesIbcClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChainId

`func (o *ResponsesIbcClient) GetChainId() *os.File`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ResponsesIbcClient) GetChainIdOk() (**os.File, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ResponsesIbcClient) SetChainId(v *os.File)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ResponsesIbcClient) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetConnectionCount

`func (o *ResponsesIbcClient) GetConnectionCount() int32`

GetConnectionCount returns the ConnectionCount field if non-nil, zero value otherwise.

### GetConnectionCountOk

`func (o *ResponsesIbcClient) GetConnectionCountOk() (*int32, bool)`

GetConnectionCountOk returns a tuple with the ConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCount

`func (o *ResponsesIbcClient) SetConnectionCount(v int32)`

SetConnectionCount sets ConnectionCount field to given value.

### HasConnectionCount

`func (o *ResponsesIbcClient) HasConnectionCount() bool`

HasConnectionCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponsesIbcClient) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponsesIbcClient) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponsesIbcClient) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponsesIbcClient) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *ResponsesIbcClient) GetCreator() ResponsesShortAddress`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ResponsesIbcClient) GetCreatorOk() (*ResponsesShortAddress, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ResponsesIbcClient) SetCreator(v ResponsesShortAddress)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ResponsesIbcClient) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetFrozenRevisionHeight

`func (o *ResponsesIbcClient) GetFrozenRevisionHeight() int32`

GetFrozenRevisionHeight returns the FrozenRevisionHeight field if non-nil, zero value otherwise.

### GetFrozenRevisionHeightOk

`func (o *ResponsesIbcClient) GetFrozenRevisionHeightOk() (*int32, bool)`

GetFrozenRevisionHeightOk returns a tuple with the FrozenRevisionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenRevisionHeight

`func (o *ResponsesIbcClient) SetFrozenRevisionHeight(v int32)`

SetFrozenRevisionHeight sets FrozenRevisionHeight field to given value.

### HasFrozenRevisionHeight

`func (o *ResponsesIbcClient) HasFrozenRevisionHeight() bool`

HasFrozenRevisionHeight returns a boolean if a field has been set.

### GetFrozenRevisionNumber

`func (o *ResponsesIbcClient) GetFrozenRevisionNumber() int32`

GetFrozenRevisionNumber returns the FrozenRevisionNumber field if non-nil, zero value otherwise.

### GetFrozenRevisionNumberOk

`func (o *ResponsesIbcClient) GetFrozenRevisionNumberOk() (*int32, bool)`

GetFrozenRevisionNumberOk returns a tuple with the FrozenRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenRevisionNumber

`func (o *ResponsesIbcClient) SetFrozenRevisionNumber(v int32)`

SetFrozenRevisionNumber sets FrozenRevisionNumber field to given value.

### HasFrozenRevisionNumber

`func (o *ResponsesIbcClient) HasFrozenRevisionNumber() bool`

HasFrozenRevisionNumber returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesIbcClient) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesIbcClient) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesIbcClient) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesIbcClient) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesIbcClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesIbcClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesIbcClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesIbcClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestRevisionHeight

`func (o *ResponsesIbcClient) GetLatestRevisionHeight() int32`

GetLatestRevisionHeight returns the LatestRevisionHeight field if non-nil, zero value otherwise.

### GetLatestRevisionHeightOk

`func (o *ResponsesIbcClient) GetLatestRevisionHeightOk() (*int32, bool)`

GetLatestRevisionHeightOk returns a tuple with the LatestRevisionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionHeight

`func (o *ResponsesIbcClient) SetLatestRevisionHeight(v int32)`

SetLatestRevisionHeight sets LatestRevisionHeight field to given value.

### HasLatestRevisionHeight

`func (o *ResponsesIbcClient) HasLatestRevisionHeight() bool`

HasLatestRevisionHeight returns a boolean if a field has been set.

### GetLatestRevisionNumber

`func (o *ResponsesIbcClient) GetLatestRevisionNumber() int32`

GetLatestRevisionNumber returns the LatestRevisionNumber field if non-nil, zero value otherwise.

### GetLatestRevisionNumberOk

`func (o *ResponsesIbcClient) GetLatestRevisionNumberOk() (*int32, bool)`

GetLatestRevisionNumberOk returns a tuple with the LatestRevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionNumber

`func (o *ResponsesIbcClient) SetLatestRevisionNumber(v int32)`

SetLatestRevisionNumber sets LatestRevisionNumber field to given value.

### HasLatestRevisionNumber

`func (o *ResponsesIbcClient) HasLatestRevisionNumber() bool`

HasLatestRevisionNumber returns a boolean if a field has been set.

### GetMaxClockDrift

`func (o *ResponsesIbcClient) GetMaxClockDrift() int32`

GetMaxClockDrift returns the MaxClockDrift field if non-nil, zero value otherwise.

### GetMaxClockDriftOk

`func (o *ResponsesIbcClient) GetMaxClockDriftOk() (*int32, bool)`

GetMaxClockDriftOk returns a tuple with the MaxClockDrift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClockDrift

`func (o *ResponsesIbcClient) SetMaxClockDrift(v int32)`

SetMaxClockDrift sets MaxClockDrift field to given value.

### HasMaxClockDrift

`func (o *ResponsesIbcClient) HasMaxClockDrift() bool`

HasMaxClockDrift returns a boolean if a field has been set.

### GetTrustLevelDenominator

`func (o *ResponsesIbcClient) GetTrustLevelDenominator() int32`

GetTrustLevelDenominator returns the TrustLevelDenominator field if non-nil, zero value otherwise.

### GetTrustLevelDenominatorOk

`func (o *ResponsesIbcClient) GetTrustLevelDenominatorOk() (*int32, bool)`

GetTrustLevelDenominatorOk returns a tuple with the TrustLevelDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevelDenominator

`func (o *ResponsesIbcClient) SetTrustLevelDenominator(v int32)`

SetTrustLevelDenominator sets TrustLevelDenominator field to given value.

### HasTrustLevelDenominator

`func (o *ResponsesIbcClient) HasTrustLevelDenominator() bool`

HasTrustLevelDenominator returns a boolean if a field has been set.

### GetTrustLevelNumerator

`func (o *ResponsesIbcClient) GetTrustLevelNumerator() int32`

GetTrustLevelNumerator returns the TrustLevelNumerator field if non-nil, zero value otherwise.

### GetTrustLevelNumeratorOk

`func (o *ResponsesIbcClient) GetTrustLevelNumeratorOk() (*int32, bool)`

GetTrustLevelNumeratorOk returns a tuple with the TrustLevelNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevelNumerator

`func (o *ResponsesIbcClient) SetTrustLevelNumerator(v int32)`

SetTrustLevelNumerator sets TrustLevelNumerator field to given value.

### HasTrustLevelNumerator

`func (o *ResponsesIbcClient) HasTrustLevelNumerator() bool`

HasTrustLevelNumerator returns a boolean if a field has been set.

### GetTrustingPeriod

`func (o *ResponsesIbcClient) GetTrustingPeriod() int32`

GetTrustingPeriod returns the TrustingPeriod field if non-nil, zero value otherwise.

### GetTrustingPeriodOk

`func (o *ResponsesIbcClient) GetTrustingPeriodOk() (*int32, bool)`

GetTrustingPeriodOk returns a tuple with the TrustingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustingPeriod

`func (o *ResponsesIbcClient) SetTrustingPeriod(v int32)`

SetTrustingPeriod sets TrustingPeriod field to given value.

### HasTrustingPeriod

`func (o *ResponsesIbcClient) HasTrustingPeriod() bool`

HasTrustingPeriod returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesIbcClient) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesIbcClient) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesIbcClient) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesIbcClient) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetType

`func (o *ResponsesIbcClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesIbcClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesIbcClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesIbcClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnbondingPeriod

`func (o *ResponsesIbcClient) GetUnbondingPeriod() int32`

GetUnbondingPeriod returns the UnbondingPeriod field if non-nil, zero value otherwise.

### GetUnbondingPeriodOk

`func (o *ResponsesIbcClient) GetUnbondingPeriodOk() (*int32, bool)`

GetUnbondingPeriodOk returns a tuple with the UnbondingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbondingPeriod

`func (o *ResponsesIbcClient) SetUnbondingPeriod(v int32)`

SetUnbondingPeriod sets UnbondingPeriod field to given value.

### HasUnbondingPeriod

`func (o *ResponsesIbcClient) HasUnbondingPeriod() bool`

HasUnbondingPeriod returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResponsesIbcClient) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResponsesIbcClient) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResponsesIbcClient) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResponsesIbcClient) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


