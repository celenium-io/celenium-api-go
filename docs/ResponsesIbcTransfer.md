# ResponsesIbcTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**Denom** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Receiver** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Sender** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Sequence** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Timeout** | Pointer to **time.Time** |  | [optional] 
**TimeoutHeight** | Pointer to **int32** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesIbcTransfer

`func NewResponsesIbcTransfer() *ResponsesIbcTransfer`

NewResponsesIbcTransfer instantiates a new ResponsesIbcTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesIbcTransferWithDefaults

`func NewResponsesIbcTransferWithDefaults() *ResponsesIbcTransfer`

NewResponsesIbcTransferWithDefaults instantiates a new ResponsesIbcTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponsesIbcTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponsesIbcTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponsesIbcTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponsesIbcTransfer) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChannelId

`func (o *ResponsesIbcTransfer) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ResponsesIbcTransfer) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ResponsesIbcTransfer) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ResponsesIbcTransfer) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetConnectionId

`func (o *ResponsesIbcTransfer) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ResponsesIbcTransfer) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ResponsesIbcTransfer) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ResponsesIbcTransfer) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetDenom

`func (o *ResponsesIbcTransfer) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesIbcTransfer) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesIbcTransfer) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesIbcTransfer) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesIbcTransfer) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesIbcTransfer) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesIbcTransfer) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesIbcTransfer) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesIbcTransfer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesIbcTransfer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesIbcTransfer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesIbcTransfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemo

`func (o *ResponsesIbcTransfer) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ResponsesIbcTransfer) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ResponsesIbcTransfer) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ResponsesIbcTransfer) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPort

`func (o *ResponsesIbcTransfer) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ResponsesIbcTransfer) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ResponsesIbcTransfer) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ResponsesIbcTransfer) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReceiver

`func (o *ResponsesIbcTransfer) GetReceiver() ResponsesShortAddress`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *ResponsesIbcTransfer) GetReceiverOk() (*ResponsesShortAddress, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *ResponsesIbcTransfer) SetReceiver(v ResponsesShortAddress)`

SetReceiver sets Receiver field to given value.

### HasReceiver

`func (o *ResponsesIbcTransfer) HasReceiver() bool`

HasReceiver returns a boolean if a field has been set.

### GetSender

`func (o *ResponsesIbcTransfer) GetSender() ResponsesShortAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ResponsesIbcTransfer) GetSenderOk() (*ResponsesShortAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ResponsesIbcTransfer) SetSender(v ResponsesShortAddress)`

SetSender sets Sender field to given value.

### HasSender

`func (o *ResponsesIbcTransfer) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSequence

`func (o *ResponsesIbcTransfer) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ResponsesIbcTransfer) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ResponsesIbcTransfer) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *ResponsesIbcTransfer) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesIbcTransfer) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesIbcTransfer) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesIbcTransfer) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesIbcTransfer) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimeout

`func (o *ResponsesIbcTransfer) GetTimeout() time.Time`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ResponsesIbcTransfer) GetTimeoutOk() (*time.Time, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ResponsesIbcTransfer) SetTimeout(v time.Time)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ResponsesIbcTransfer) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTimeoutHeight

`func (o *ResponsesIbcTransfer) GetTimeoutHeight() int32`

GetTimeoutHeight returns the TimeoutHeight field if non-nil, zero value otherwise.

### GetTimeoutHeightOk

`func (o *ResponsesIbcTransfer) GetTimeoutHeightOk() (*int32, bool)`

GetTimeoutHeightOk returns a tuple with the TimeoutHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutHeight

`func (o *ResponsesIbcTransfer) SetTimeoutHeight(v int32)`

SetTimeoutHeight sets TimeoutHeight field to given value.

### HasTimeoutHeight

`func (o *ResponsesIbcTransfer) HasTimeoutHeight() bool`

HasTimeoutHeight returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesIbcTransfer) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesIbcTransfer) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesIbcTransfer) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesIbcTransfer) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


