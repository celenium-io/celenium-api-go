# ResponsesIbcConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelsCount** | Pointer to **int32** |  | [optional] 
**Client** | Pointer to [**ResponsesShortIbcClient**](ResponsesShortIbcClient.md) |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 
**ConnectedHeight** | Pointer to **int32** |  | [optional] 
**ConnectedTxHash** | Pointer to ***os.File** |  | [optional] 
**CounterpartyClientId** | Pointer to **string** |  | [optional] 
**CounterpartyConnectionId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTxHash** | Pointer to ***os.File** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesIbcConnection

`func NewResponsesIbcConnection() *ResponsesIbcConnection`

NewResponsesIbcConnection instantiates a new ResponsesIbcConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesIbcConnectionWithDefaults

`func NewResponsesIbcConnectionWithDefaults() *ResponsesIbcConnection`

NewResponsesIbcConnectionWithDefaults instantiates a new ResponsesIbcConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelsCount

`func (o *ResponsesIbcConnection) GetChannelsCount() int32`

GetChannelsCount returns the ChannelsCount field if non-nil, zero value otherwise.

### GetChannelsCountOk

`func (o *ResponsesIbcConnection) GetChannelsCountOk() (*int32, bool)`

GetChannelsCountOk returns a tuple with the ChannelsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelsCount

`func (o *ResponsesIbcConnection) SetChannelsCount(v int32)`

SetChannelsCount sets ChannelsCount field to given value.

### HasChannelsCount

`func (o *ResponsesIbcConnection) HasChannelsCount() bool`

HasChannelsCount returns a boolean if a field has been set.

### GetClient

`func (o *ResponsesIbcConnection) GetClient() ResponsesShortIbcClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ResponsesIbcConnection) GetClientOk() (*ResponsesShortIbcClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ResponsesIbcConnection) SetClient(v ResponsesShortIbcClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *ResponsesIbcConnection) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetConnectedAt

`func (o *ResponsesIbcConnection) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ResponsesIbcConnection) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ResponsesIbcConnection) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *ResponsesIbcConnection) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetConnectedHeight

`func (o *ResponsesIbcConnection) GetConnectedHeight() int32`

GetConnectedHeight returns the ConnectedHeight field if non-nil, zero value otherwise.

### GetConnectedHeightOk

`func (o *ResponsesIbcConnection) GetConnectedHeightOk() (*int32, bool)`

GetConnectedHeightOk returns a tuple with the ConnectedHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedHeight

`func (o *ResponsesIbcConnection) SetConnectedHeight(v int32)`

SetConnectedHeight sets ConnectedHeight field to given value.

### HasConnectedHeight

`func (o *ResponsesIbcConnection) HasConnectedHeight() bool`

HasConnectedHeight returns a boolean if a field has been set.

### GetConnectedTxHash

`func (o *ResponsesIbcConnection) GetConnectedTxHash() *os.File`

GetConnectedTxHash returns the ConnectedTxHash field if non-nil, zero value otherwise.

### GetConnectedTxHashOk

`func (o *ResponsesIbcConnection) GetConnectedTxHashOk() (**os.File, bool)`

GetConnectedTxHashOk returns a tuple with the ConnectedTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedTxHash

`func (o *ResponsesIbcConnection) SetConnectedTxHash(v *os.File)`

SetConnectedTxHash sets ConnectedTxHash field to given value.

### HasConnectedTxHash

`func (o *ResponsesIbcConnection) HasConnectedTxHash() bool`

HasConnectedTxHash returns a boolean if a field has been set.

### GetCounterpartyClientId

`func (o *ResponsesIbcConnection) GetCounterpartyClientId() string`

GetCounterpartyClientId returns the CounterpartyClientId field if non-nil, zero value otherwise.

### GetCounterpartyClientIdOk

`func (o *ResponsesIbcConnection) GetCounterpartyClientIdOk() (*string, bool)`

GetCounterpartyClientIdOk returns a tuple with the CounterpartyClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyClientId

`func (o *ResponsesIbcConnection) SetCounterpartyClientId(v string)`

SetCounterpartyClientId sets CounterpartyClientId field to given value.

### HasCounterpartyClientId

`func (o *ResponsesIbcConnection) HasCounterpartyClientId() bool`

HasCounterpartyClientId returns a boolean if a field has been set.

### GetCounterpartyConnectionId

`func (o *ResponsesIbcConnection) GetCounterpartyConnectionId() string`

GetCounterpartyConnectionId returns the CounterpartyConnectionId field if non-nil, zero value otherwise.

### GetCounterpartyConnectionIdOk

`func (o *ResponsesIbcConnection) GetCounterpartyConnectionIdOk() (*string, bool)`

GetCounterpartyConnectionIdOk returns a tuple with the CounterpartyConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyConnectionId

`func (o *ResponsesIbcConnection) SetCounterpartyConnectionId(v string)`

SetCounterpartyConnectionId sets CounterpartyConnectionId field to given value.

### HasCounterpartyConnectionId

`func (o *ResponsesIbcConnection) HasCounterpartyConnectionId() bool`

HasCounterpartyConnectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponsesIbcConnection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponsesIbcConnection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponsesIbcConnection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponsesIbcConnection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTxHash

`func (o *ResponsesIbcConnection) GetCreatedTxHash() *os.File`

GetCreatedTxHash returns the CreatedTxHash field if non-nil, zero value otherwise.

### GetCreatedTxHashOk

`func (o *ResponsesIbcConnection) GetCreatedTxHashOk() (**os.File, bool)`

GetCreatedTxHashOk returns a tuple with the CreatedTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTxHash

`func (o *ResponsesIbcConnection) SetCreatedTxHash(v *os.File)`

SetCreatedTxHash sets CreatedTxHash field to given value.

### HasCreatedTxHash

`func (o *ResponsesIbcConnection) HasCreatedTxHash() bool`

HasCreatedTxHash returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesIbcConnection) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesIbcConnection) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesIbcConnection) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesIbcConnection) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesIbcConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesIbcConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesIbcConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesIbcConnection) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


