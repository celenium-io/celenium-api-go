# ResponsesForwardingInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to [**ResponsesChainMetadata**](ResponsesChainMetadata.md) |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Denom** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Received** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TxHash** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesForwardingInput

`func NewResponsesForwardingInput() *ResponsesForwardingInput`

NewResponsesForwardingInput instantiates a new ResponsesForwardingInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesForwardingInputWithDefaults

`func NewResponsesForwardingInputWithDefaults() *ResponsesForwardingInput`

NewResponsesForwardingInputWithDefaults instantiates a new ResponsesForwardingInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ResponsesForwardingInput) GetChain() ResponsesChainMetadata`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ResponsesForwardingInput) GetChainOk() (*ResponsesChainMetadata, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ResponsesForwardingInput) SetChain(v ResponsesChainMetadata)`

SetChain sets Chain field to given value.

### HasChain

`func (o *ResponsesForwardingInput) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *ResponsesForwardingInput) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ResponsesForwardingInput) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ResponsesForwardingInput) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ResponsesForwardingInput) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChannelId

`func (o *ResponsesForwardingInput) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ResponsesForwardingInput) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ResponsesForwardingInput) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ResponsesForwardingInput) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetDenom

`func (o *ResponsesForwardingInput) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesForwardingInput) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesForwardingInput) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesForwardingInput) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

### GetFrom

`func (o *ResponsesForwardingInput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ResponsesForwardingInput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ResponsesForwardingInput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ResponsesForwardingInput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesForwardingInput) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesForwardingInput) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesForwardingInput) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesForwardingInput) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetReceived

`func (o *ResponsesForwardingInput) GetReceived() string`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ResponsesForwardingInput) GetReceivedOk() (*string, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ResponsesForwardingInput) SetReceived(v string)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ResponsesForwardingInput) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesForwardingInput) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesForwardingInput) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesForwardingInput) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesForwardingInput) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesForwardingInput) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesForwardingInput) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesForwardingInput) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesForwardingInput) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.

### GetType

`func (o *ResponsesForwardingInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesForwardingInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesForwardingInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesForwardingInput) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


