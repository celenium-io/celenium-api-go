# ResponsesZkISMMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MessageId** | Pointer to ***os.File** |  | [optional] 
**Signer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**StateRoot** | Pointer to ***os.File** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesZkISMMessage

`func NewResponsesZkISMMessage() *ResponsesZkISMMessage`

NewResponsesZkISMMessage instantiates a new ResponsesZkISMMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesZkISMMessageWithDefaults

`func NewResponsesZkISMMessageWithDefaults() *ResponsesZkISMMessage`

NewResponsesZkISMMessageWithDefaults instantiates a new ResponsesZkISMMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ResponsesZkISMMessage) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesZkISMMessage) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesZkISMMessage) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesZkISMMessage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesZkISMMessage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesZkISMMessage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesZkISMMessage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesZkISMMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageId

`func (o *ResponsesZkISMMessage) GetMessageId() *os.File`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ResponsesZkISMMessage) GetMessageIdOk() (**os.File, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ResponsesZkISMMessage) SetMessageId(v *os.File)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ResponsesZkISMMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetSigner

`func (o *ResponsesZkISMMessage) GetSigner() ResponsesShortAddress`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ResponsesZkISMMessage) GetSignerOk() (*ResponsesShortAddress, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ResponsesZkISMMessage) SetSigner(v ResponsesShortAddress)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ResponsesZkISMMessage) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetStateRoot

`func (o *ResponsesZkISMMessage) GetStateRoot() *os.File`

GetStateRoot returns the StateRoot field if non-nil, zero value otherwise.

### GetStateRootOk

`func (o *ResponsesZkISMMessage) GetStateRootOk() (**os.File, bool)`

GetStateRootOk returns a tuple with the StateRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRoot

`func (o *ResponsesZkISMMessage) SetStateRoot(v *os.File)`

SetStateRoot sets StateRoot field to given value.

### HasStateRoot

`func (o *ResponsesZkISMMessage) HasStateRoot() bool`

HasStateRoot returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesZkISMMessage) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesZkISMMessage) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesZkISMMessage) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesZkISMMessage) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesZkISMMessage) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesZkISMMessage) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesZkISMMessage) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesZkISMMessage) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


