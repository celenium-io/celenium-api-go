# ResponsesZkISMUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**NewState** | Pointer to ***os.File** |  | [optional] 
**NewStateRoot** | Pointer to ***os.File** |  | [optional] 
**Signer** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesZkISMUpdate

`func NewResponsesZkISMUpdate() *ResponsesZkISMUpdate`

NewResponsesZkISMUpdate instantiates a new ResponsesZkISMUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesZkISMUpdateWithDefaults

`func NewResponsesZkISMUpdateWithDefaults() *ResponsesZkISMUpdate`

NewResponsesZkISMUpdateWithDefaults instantiates a new ResponsesZkISMUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ResponsesZkISMUpdate) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesZkISMUpdate) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesZkISMUpdate) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesZkISMUpdate) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesZkISMUpdate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesZkISMUpdate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesZkISMUpdate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesZkISMUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNewState

`func (o *ResponsesZkISMUpdate) GetNewState() *os.File`

GetNewState returns the NewState field if non-nil, zero value otherwise.

### GetNewStateOk

`func (o *ResponsesZkISMUpdate) GetNewStateOk() (**os.File, bool)`

GetNewStateOk returns a tuple with the NewState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewState

`func (o *ResponsesZkISMUpdate) SetNewState(v *os.File)`

SetNewState sets NewState field to given value.

### HasNewState

`func (o *ResponsesZkISMUpdate) HasNewState() bool`

HasNewState returns a boolean if a field has been set.

### GetNewStateRoot

`func (o *ResponsesZkISMUpdate) GetNewStateRoot() *os.File`

GetNewStateRoot returns the NewStateRoot field if non-nil, zero value otherwise.

### GetNewStateRootOk

`func (o *ResponsesZkISMUpdate) GetNewStateRootOk() (**os.File, bool)`

GetNewStateRootOk returns a tuple with the NewStateRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStateRoot

`func (o *ResponsesZkISMUpdate) SetNewStateRoot(v *os.File)`

SetNewStateRoot sets NewStateRoot field to given value.

### HasNewStateRoot

`func (o *ResponsesZkISMUpdate) HasNewStateRoot() bool`

HasNewStateRoot returns a boolean if a field has been set.

### GetSigner

`func (o *ResponsesZkISMUpdate) GetSigner() ResponsesShortAddress`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *ResponsesZkISMUpdate) GetSignerOk() (*ResponsesShortAddress, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *ResponsesZkISMUpdate) SetSigner(v ResponsesShortAddress)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *ResponsesZkISMUpdate) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesZkISMUpdate) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesZkISMUpdate) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesZkISMUpdate) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesZkISMUpdate) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesZkISMUpdate) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesZkISMUpdate) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesZkISMUpdate) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesZkISMUpdate) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


