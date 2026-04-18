# ResponsesForwarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to [**ResponsesChainMetadata**](ResponsesChainMetadata.md) |  | [optional] 
**Denom** | Pointer to **string** |  | [optional] 
**DestAddress** | Pointer to ***os.File** |  | [optional] 
**DestDomain** | Pointer to **int64** |  | [optional] 
**ForwardAddress** | Pointer to [**ResponsesShortAddress**](ResponsesShortAddress.md) |  | [optional] 
**Height** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Inputs** | Pointer to [**[]ResponsesForwardingInput**](ResponsesForwardingInput.md) |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
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

### GetAmount

`func (o *ResponsesForwarding) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponsesForwarding) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponsesForwarding) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponsesForwarding) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

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

### GetDenom

`func (o *ResponsesForwarding) GetDenom() string`

GetDenom returns the Denom field if non-nil, zero value otherwise.

### GetDenomOk

`func (o *ResponsesForwarding) GetDenomOk() (*string, bool)`

GetDenomOk returns a tuple with the Denom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenom

`func (o *ResponsesForwarding) SetDenom(v string)`

SetDenom sets Denom field to given value.

### HasDenom

`func (o *ResponsesForwarding) HasDenom() bool`

HasDenom returns a boolean if a field has been set.

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

### GetMessageId

`func (o *ResponsesForwarding) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ResponsesForwarding) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ResponsesForwarding) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ResponsesForwarding) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

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

### GetTokenId

`func (o *ResponsesForwarding) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ResponsesForwarding) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ResponsesForwarding) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ResponsesForwarding) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

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


