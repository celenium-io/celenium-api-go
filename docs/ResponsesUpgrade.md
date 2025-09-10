# ResponsesUpgrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**MsgId** | Pointer to **int64** |  | [optional] 
**Signer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

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

### GetId

`func (o *ResponsesUpgrade) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesUpgrade) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesUpgrade) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesUpgrade) HasId() bool`

HasId returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


