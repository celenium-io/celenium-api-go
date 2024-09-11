# ResponsesTx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codespace** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**EventsCount** | Pointer to **int64** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**GasUsed** | Pointer to **int64** |  | [optional] 
**GasWanted** | Pointer to **int64** |  | [optional] 
**Hash** | Pointer to ***os.File** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**MessageTypes** | Pointer to [**[]TypesMsgType**](TypesMsgType.md) |  | [optional] 
**Messages** | Pointer to [**[]ResponsesMessage**](ResponsesMessage.md) |  | [optional] 
**MessagesCount** | Pointer to **int64** |  | [optional] 
**Position** | Pointer to **int64** |  | [optional] 
**Signers** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus**](GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TimeoutHeight** | Pointer to **int64** |  | [optional] 

## Methods

### NewResponsesTx

`func NewResponsesTx() *ResponsesTx`

NewResponsesTx instantiates a new ResponsesTx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTxWithDefaults

`func NewResponsesTxWithDefaults() *ResponsesTx`

NewResponsesTxWithDefaults instantiates a new ResponsesTx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodespace

`func (o *ResponsesTx) GetCodespace() string`

GetCodespace returns the Codespace field if non-nil, zero value otherwise.

### GetCodespaceOk

`func (o *ResponsesTx) GetCodespaceOk() (*string, bool)`

GetCodespaceOk returns a tuple with the Codespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodespace

`func (o *ResponsesTx) SetCodespace(v string)`

SetCodespace sets Codespace field to given value.

### HasCodespace

`func (o *ResponsesTx) HasCodespace() bool`

HasCodespace returns a boolean if a field has been set.

### GetError

`func (o *ResponsesTx) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResponsesTx) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResponsesTx) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ResponsesTx) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEventsCount

`func (o *ResponsesTx) GetEventsCount() int64`

GetEventsCount returns the EventsCount field if non-nil, zero value otherwise.

### GetEventsCountOk

`func (o *ResponsesTx) GetEventsCountOk() (*int64, bool)`

GetEventsCountOk returns a tuple with the EventsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsCount

`func (o *ResponsesTx) SetEventsCount(v int64)`

SetEventsCount sets EventsCount field to given value.

### HasEventsCount

`func (o *ResponsesTx) HasEventsCount() bool`

HasEventsCount returns a boolean if a field has been set.

### GetFee

`func (o *ResponsesTx) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *ResponsesTx) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *ResponsesTx) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *ResponsesTx) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetGasUsed

`func (o *ResponsesTx) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ResponsesTx) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ResponsesTx) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ResponsesTx) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### GetGasWanted

`func (o *ResponsesTx) GetGasWanted() int64`

GetGasWanted returns the GasWanted field if non-nil, zero value otherwise.

### GetGasWantedOk

`func (o *ResponsesTx) GetGasWantedOk() (*int64, bool)`

GetGasWantedOk returns a tuple with the GasWanted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasWanted

`func (o *ResponsesTx) SetGasWanted(v int64)`

SetGasWanted sets GasWanted field to given value.

### HasGasWanted

`func (o *ResponsesTx) HasGasWanted() bool`

HasGasWanted returns a boolean if a field has been set.

### GetHash

`func (o *ResponsesTx) GetHash() *os.File`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ResponsesTx) GetHashOk() (**os.File, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ResponsesTx) SetHash(v *os.File)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ResponsesTx) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesTx) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesTx) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesTx) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesTx) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesTx) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesTx) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesTx) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesTx) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemo

`func (o *ResponsesTx) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ResponsesTx) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ResponsesTx) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ResponsesTx) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetMessageTypes

`func (o *ResponsesTx) GetMessageTypes() []TypesMsgType`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *ResponsesTx) GetMessageTypesOk() (*[]TypesMsgType, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *ResponsesTx) SetMessageTypes(v []TypesMsgType)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *ResponsesTx) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.

### GetMessages

`func (o *ResponsesTx) GetMessages() []ResponsesMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ResponsesTx) GetMessagesOk() (*[]ResponsesMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ResponsesTx) SetMessages(v []ResponsesMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ResponsesTx) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMessagesCount

`func (o *ResponsesTx) GetMessagesCount() int64`

GetMessagesCount returns the MessagesCount field if non-nil, zero value otherwise.

### GetMessagesCountOk

`func (o *ResponsesTx) GetMessagesCountOk() (*int64, bool)`

GetMessagesCountOk returns a tuple with the MessagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesCount

`func (o *ResponsesTx) SetMessagesCount(v int64)`

SetMessagesCount sets MessagesCount field to given value.

### HasMessagesCount

`func (o *ResponsesTx) HasMessagesCount() bool`

HasMessagesCount returns a boolean if a field has been set.

### GetPosition

`func (o *ResponsesTx) GetPosition() int64`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ResponsesTx) GetPositionOk() (*int64, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ResponsesTx) SetPosition(v int64)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ResponsesTx) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSigners

`func (o *ResponsesTx) GetSigners() []string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *ResponsesTx) GetSignersOk() (*[]string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *ResponsesTx) SetSigners(v []string)`

SetSigners sets Signers field to given value.

### HasSigners

`func (o *ResponsesTx) HasSigners() bool`

HasSigners returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesTx) GetStatus() GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesTx) GetStatusOk() (*GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesTx) SetStatus(v GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesTx) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesTx) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesTx) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesTx) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesTx) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeoutHeight

`func (o *ResponsesTx) GetTimeoutHeight() int64`

GetTimeoutHeight returns the TimeoutHeight field if non-nil, zero value otherwise.

### GetTimeoutHeightOk

`func (o *ResponsesTx) GetTimeoutHeightOk() (*int64, bool)`

GetTimeoutHeightOk returns a tuple with the TimeoutHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutHeight

`func (o *ResponsesTx) SetTimeoutHeight(v int64)`

SetTimeoutHeight sets TimeoutHeight field to given value.

### HasTimeoutHeight

`func (o *ResponsesTx) HasTimeoutHeight() bool`

HasTimeoutHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


