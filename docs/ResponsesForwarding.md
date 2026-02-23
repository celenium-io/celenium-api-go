# ResponsesForwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to [**ResponsesChainMetadata**](ResponsesChainMetadata.md) |  | [optional] 
**DestAddress** | Pointer to ***os.File** |  | [optional] 
**DestDomain** | Pointer to **int64** |  | [optional] 
**FailedCount** | Pointer to **int64** |  | [optional] 
**ForwardAddress** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Inputs** | Pointer to [**[]ResponsesForwardingInput**](ResponsesForwardingInput.md) |  | [optional] 
**SuccessCount** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Transfers** | Pointer to **[]int32** |  | [optional] 
**TxHash** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewResponsesForwarding

`func NewResponsesForwarding() *ResponsesForwarding`

NewResponsesForwarding instantiates a new ResponsesForwarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesForwardingWithDefaults

`func NewResponsesForwardingWithDefaults() *ResponsesForwarding`

NewResponsesForwardingWithDefaults instantiates a new ResponsesForwarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *ResponsesForwarding) GetChain() ResponsesChainMetadata`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *ResponsesForwarding) GetChainOk() (*ResponsesChainMetadata, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *ResponsesForwarding) SetChain(v ResponsesChainMetadata)`

SetChain sets Chain field to given value.

### HasChain

`func (o *ResponsesForwarding) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetDestAddress

`func (o *ResponsesForwarding) GetDestAddress() *os.File`

GetDestAddress returns the DestAddress field if non-nil, zero value otherwise.

### GetDestAddressOk

`func (o *ResponsesForwarding) GetDestAddressOk() (**os.File, bool)`

GetDestAddressOk returns a tuple with the DestAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAddress

`func (o *ResponsesForwarding) SetDestAddress(v *os.File)`

SetDestAddress sets DestAddress field to given value.

### HasDestAddress

`func (o *ResponsesForwarding) HasDestAddress() bool`

HasDestAddress returns a boolean if a field has been set.

### GetDestDomain

`func (o *ResponsesForwarding) GetDestDomain() int64`

GetDestDomain returns the DestDomain field if non-nil, zero value otherwise.

### GetDestDomainOk

`func (o *ResponsesForwarding) GetDestDomainOk() (*int64, bool)`

GetDestDomainOk returns a tuple with the DestDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestDomain

`func (o *ResponsesForwarding) SetDestDomain(v int64)`

SetDestDomain sets DestDomain field to given value.

### HasDestDomain

`func (o *ResponsesForwarding) HasDestDomain() bool`

HasDestDomain returns a boolean if a field has been set.

### GetFailedCount

`func (o *ResponsesForwarding) GetFailedCount() int64`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *ResponsesForwarding) GetFailedCountOk() (*int64, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *ResponsesForwarding) SetFailedCount(v int64)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *ResponsesForwarding) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### GetForwardAddress

`func (o *ResponsesForwarding) GetForwardAddress() ResponsesShortAddress`

GetForwardAddress returns the ForwardAddress field if non-nil, zero value otherwise.

### GetForwardAddressOk

`func (o *ResponsesForwarding) GetForwardAddressOk() (*ResponsesShortAddress, bool)`

GetForwardAddressOk returns a tuple with the ForwardAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAddress

`func (o *ResponsesForwarding) SetForwardAddress(v ResponsesShortAddress)`

SetForwardAddress sets ForwardAddress field to given value.

### HasForwardAddress

`func (o *ResponsesForwarding) HasForwardAddress() bool`

HasForwardAddress returns a boolean if a field has been set.

### GetHeight

`func (o *ResponsesForwarding) GetHeight() int64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ResponsesForwarding) GetHeightOk() (*int64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ResponsesForwarding) SetHeight(v int64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ResponsesForwarding) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetId

`func (o *ResponsesForwarding) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesForwarding) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesForwarding) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResponsesForwarding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *ResponsesForwarding) GetInputs() []ResponsesForwardingInput`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ResponsesForwarding) GetInputsOk() (*[]ResponsesForwardingInput, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ResponsesForwarding) SetInputs(v []ResponsesForwardingInput)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ResponsesForwarding) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetSuccessCount

`func (o *ResponsesForwarding) GetSuccessCount() int64`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *ResponsesForwarding) GetSuccessCountOk() (*int64, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *ResponsesForwarding) SetSuccessCount(v int64)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *ResponsesForwarding) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTime

`func (o *ResponsesForwarding) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ResponsesForwarding) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ResponsesForwarding) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ResponsesForwarding) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTransfers

`func (o *ResponsesForwarding) GetTransfers() []int32`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *ResponsesForwarding) GetTransfersOk() (*[]int32, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *ResponsesForwarding) SetTransfers(v []int32)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *ResponsesForwarding) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.

### GetTxHash

`func (o *ResponsesForwarding) GetTxHash() *os.File`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *ResponsesForwarding) GetTxHashOk() (**os.File, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *ResponsesForwarding) SetTxHash(v *os.File)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *ResponsesForwarding) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


