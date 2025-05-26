# ResponsesIbcChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**ResponsesShortIbcClient**](ResponsesShortIbcClient.md) |  | [optional] 
**ConfirmationHeight** | Pointer to **int32** |  | [optional] 
**ConfirmationTxHash** | Pointer to ***os.File** |  | [optional] 
**ConfirmedAt** | Pointer to **time.Time** |  | [optional] 
**ConnectionId** | Pointer to **string** |  | [optional] 
**CounterpartyChannelId** | Pointer to **string** |  | [optional] 
**CounterpartyPortId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedTxHash** | Pointer to ***os.File** |  | [optional] 
**Creator** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ordering** | Pointer to **bool** |  | [optional] 
**PortId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewResponsesIbcChannel

`func NewResponsesIbcChannel() *ResponsesIbcChannel`

NewResponsesIbcChannel instantiates a new ResponsesIbcChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesIbcChannelWithDefaults

`func NewResponsesIbcChannelWithDefaults() *ResponsesIbcChannel`

NewResponsesIbcChannelWithDefaults instantiates a new ResponsesIbcChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *ResponsesIbcChannel) GetClient() ResponsesShortIbcClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ResponsesIbcChannel) GetClientOk() (*ResponsesShortIbcClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ResponsesIbcChannel) SetClient(v ResponsesShortIbcClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *ResponsesIbcChannel) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetConfirmationHeight

`func (o *ResponsesIbcChannel) GetConfirmationHeight() int32`

GetConfirmationHeight returns the ConfirmationHeight field if non-nil, zero value otherwise.

### GetConfirmationHeightOk

`func (o *ResponsesIbcChannel) GetConfirmationHeightOk() (*int32, bool)`

GetConfirmationHeightOk returns a tuple with the ConfirmationHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationHeight

`func (o *ResponsesIbcChannel) SetConfirmationHeight(v int32)`

SetConfirmationHeight sets ConfirmationHeight field to given value.

### HasConfirmationHeight

`func (o *ResponsesIbcChannel) HasConfirmationHeight() bool`

HasConfirmationHeight returns a boolean if a field has been set.

### GetConfirmationTxHash

`func (o *ResponsesIbcChannel) GetConfirmationTxHash() *os.File`

GetConfirmationTxHash returns the ConfirmationTxHash field if non-nil, zero value otherwise.

### GetConfirmationTxHashOk

`func (o *ResponsesIbcChannel) GetConfirmationTxHashOk() (**os.File, bool)`

GetConfirmationTxHashOk returns a tuple with the ConfirmationTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationTxHash

`func (o *ResponsesIbcChannel) SetConfirmationTxHash(v *os.File)`

SetConfirmationTxHash sets ConfirmationTxHash field to given value.

### HasConfirmationTxHash

`func (o *ResponsesIbcChannel) HasConfirmationTxHash() bool`

HasConfirmationTxHash returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *ResponsesIbcChannel) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *ResponsesIbcChannel) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *ResponsesIbcChannel) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *ResponsesIbcChannel) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetConnectionId

`func (o *ResponsesIbcChannel) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ResponsesIbcChannel) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ResponsesIbcChannel) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ResponsesIbcChannel) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCounterpartyChannelId

`func (o *ResponsesIbcChannel) GetCounterpartyChannelId() string`

GetCounterpartyChannelId returns the CounterpartyChannelId field if non-nil, zero value otherwise.

### GetCounterpartyChannelIdOk

`func (o *ResponsesIbcChannel) GetCounterpartyChannelIdOk() (*string, bool)`

GetCounterpartyChannelIdOk returns a tuple with the CounterpartyChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyChannelId

`func (o *ResponsesIbcChannel) SetCounterpartyChannelId(v string)`

SetCounterpartyChannelId sets CounterpartyChannelId field to given value.

### HasCounterpartyChannelId

`func (o *ResponsesIbcChannel) HasCounterpartyChannelId() bool`

HasCounterpartyChannelId returns a boolean if a field has been set.

### GetCounterpartyPortId

`func (o *ResponsesIbcChannel) GetCounterpartyPortId() string`

GetCounterpartyPortId returns the CounterpartyPortId field if non-nil, zero value otherwise.

### GetCounterpartyPortIdOk

`func (o *ResponsesIbcChannel) GetCounterpartyPortIdOk() (*string, bool)`

GetCounterpartyPortIdOk returns a tuple with the CounterpartyPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartyPortId

`func (o *ResponsesIbcChannel) SetCounterpartyPortId(v string)`

SetCounterpartyPortId sets CounterpartyPortId field to given value.

### HasCounterpartyPortId

`func (o *ResponsesIbcChannel) HasCounterpartyPortId() bool`

HasCounterpartyPortId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResponsesIbcChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResponsesIbcChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResponsesIbcChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResponsesIbcChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedTxHash

`func (o *ResponsesIbcChannel) GetCreatedTxHash() *os.File`

GetCreatedTxHash returns the CreatedTxHash field if non-nil, zero value otherwise.

### GetCreatedTxHashOk

`func (o *ResponsesIbcChannel) GetCreatedTxHashOk() (**os.File, bool)`

GetCreatedTxHashOk returns a tuple with the CreatedTxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTxHash

`func (o *ResponsesIbcChannel) SetCreatedTxHash(v *os.File)`

SetCreatedTxHash sets CreatedTxHash field to given value.

### HasCreatedTxHash

`func (o *ResponsesIbcChannel) HasCreatedTxHash() bool`

HasCreatedTxHash returns a boolean if a field has been set.

### GetCreator

`func (o *ResponsesIbcChannel) GetCreator() ResponsesShortAddress`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ResponsesIbcChannel) GetCreatorOk() (*ResponsesShortAddress, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ResponsesIbcChannel) SetCreator(v ResponsesShortAddress)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ResponsesIbcChannel) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesIbcChannel) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesIbcChannel) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesIbcChannel) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesIbcChannel) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesIbcChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesIbcChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesIbcChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesIbcChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrdering

`func (o *ResponsesIbcChannel) GetOrdering() bool`

GetOrdering returns the Ordering field if non-nil, zero value otherwise.

### GetOrderingOk

`func (o *ResponsesIbcChannel) GetOrderingOk() (*bool, bool)`

GetOrderingOk returns a tuple with the Ordering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdering

`func (o *ResponsesIbcChannel) SetOrdering(v bool)`

SetOrdering sets Ordering field to given value.

### HasOrdering

`func (o *ResponsesIbcChannel) HasOrdering() bool`

HasOrdering returns a boolean if a field has been set.

### GetPortId

`func (o *ResponsesIbcChannel) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *ResponsesIbcChannel) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *ResponsesIbcChannel) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *ResponsesIbcChannel) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetStatus

`func (o *ResponsesIbcChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesIbcChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesIbcChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesIbcChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ResponsesIbcChannel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ResponsesIbcChannel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ResponsesIbcChannel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ResponsesIbcChannel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


